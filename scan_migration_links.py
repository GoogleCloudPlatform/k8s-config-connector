#!/usr/bin/env python3
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os
import yaml
import json
import re
import subprocess
import sys

def main():
    print("Starting KCC Direct Migration Issues Cross-Reference Linker...")

    # Step 1: Build kind-to-group mapping from CRDs
    kind_to_gvk = {}
    crds_dir = "config/crds/resources"
    if not os.path.exists(crds_dir):
        print(f"Error: CRDs directory {crds_dir} not found.", file=sys.stderr)
        sys.exit(1)

    print("Parsing CRDs to build Kind-to-Group mapping...")
    for filename in os.listdir(crds_dir):
        if filename.endswith(".yaml"):
            filepath = os.path.join(crds_dir, filename)
            with open(filepath, "r") as f:
                try:
                    crd = yaml.safe_load(f)
                    spec = crd.get("spec", {})
                    group = spec.get("group", "")
                    if group.endswith(".cnrm.cloud.google.com"):
                        group_prefix = group[:-len(".cnrm.cloud.google.com")]
                    else:
                        group_prefix = group
                    names = spec.get("names", {})
                    kind = names.get("kind", "")
                    if kind:
                        kind_to_gvk[kind.lower()] = (kind, group_prefix)
                except Exception as e:
                    pass
    print(f"Loaded {len(kind_to_gvk)} CRD kind-to-group mappings.")

    # Step 2: Fetch all issues with area/direct and overseer labels
    print("Fetching issues from GitHub...")
    try:
        res = subprocess.run([
            "gh", "issue", "list", 
            "--state", "all", 
            "--label", "area/direct,overseer", 
            "--limit", "500", 
            "--json", "number,title,labels,state"
        ], capture_output=True, text=True, check=True)
        issues = json.loads(res.stdout)
    except Exception as e:
        print(f"Error fetching issues: {e}", file=sys.stderr)
        sys.exit(1)
    print(f"Fetched {len(issues)} direct migration issues.")

    # Step 3: Classify into Step 1 and Step 2 issues
    step1_issues = []
    step2_issues = []
    for issue in issues:
        labels = [l["name"] for l in issue.get("labels", [])]
        if "step/gen-types" in labels:
            step1_issues.append(issue)
        elif "step/identity-reference" in labels:
            step2_issues.append(issue)

    print(f"Classified {len(step1_issues)} Step 1 issues and {len(step2_issues)} Step 2 issues.")

    # Build matchers for substring matching (group + kind, and kind)
    matchers = []
    for kind_lower, (kind, group) in kind_to_gvk.items():
        matchers.append((group + kind_lower, kind, group))
        matchers.append((kind_lower, kind, group))
    matchers.sort(key=lambda x: len(x[0]), reverse=True)

    def find_resource(title):
        norm_title = re.sub(r'[^a-zA-Z0-9]', '', title).lower()
        for pattern, kind, group in matchers:
            if pattern in norm_title:
                return kind, group
        return None

    # Group Step 1 issues by (Kind, Group)
    step1_by_gvk = {}
    for iss in step1_issues:
        res_info = find_resource(iss["title"])
        if res_info:
            key = (res_info[0].lower(), res_info[1].lower())
            if key not in step1_by_gvk:
                step1_by_gvk[key] = []
            step1_by_gvk[key].append(iss)

    # Step 4: Identify missing links for open Step 2 issues
    print("Scanning open Step 2 issues for missing Step 1 links...")
    comments_posted = 0
    max_comments = 10

    for s2_iss in step2_issues:
        if s2_iss["state"] != "OPEN":
            continue

        res_info = find_resource(s2_iss["title"])
        if not res_info:
            print(f"Could not identify resource GVK for Step 2 issue #{s2_iss['number']}: '{s2_iss['title']}'")
            continue

        kind, group = res_info
        key = (kind.lower(), group.lower())
        s1_list = step1_by_gvk.get(key, [])
        if not s1_list:
            # No corresponding Step 1 issue exists in the overseer label-matched set
            continue

        s1_iss = s1_list[0] # Prefer the first matched Step 1 issue

        # Check if the Step 2 issue already contains a reference to the Step 1 issue number
        print(f"Checking links for #{s2_iss['number']} ({group} {kind})...")
        try:
            view_res = subprocess.run([
                "gh", "issue", "view", str(s2_iss["number"]), "--json", "body,comments"
            ], capture_output=True, text=True, check=True)
            view_data = json.loads(view_res.stdout)
        except Exception as e:
            print(f"Error fetching issue #{s2_iss['number']} details: {e}", file=sys.stderr)
            continue

        body = view_data.get("body", "")
        comments = [c.get("body", "") for c in view_data.get("comments", [])]
        mention = f"#{s1_iss['number']}"

        if mention in body or any(mention in c for c in comments):
            print(f"  Step 2 issue #{s2_iss['number']} already links to Step 1 issue #{s1_iss['number']}.")
            continue

        # Link is missing! Proceed to post comment if within throttle limits
        print(f"  Missing link identified! Step 2 #{s2_iss['number']} is missing reference to Step 1 #{s1_iss['number']}.")
        if comments_posted >= max_comments:
            print(f"  [Throttle] Already identified {max_comments} missing links. Skipping posting comment.")
            continue

        # Post the comment
        step1_done = "x" if s1_iss["state"] == "CLOSED" else " "
        comment_body = f"""### Migration Context Chain
This issue is part of the direct migration for `{group} {kind}` (Epic #5954).

- [{step1_done}] **Step 1: Types & Scaffolding** - #{s1_iss['number']}
- [ ] **Step 2: Identity & Reference** - #{s2_iss['number']}"""

        print(f"  Posting cross-reference comment on #{s2_iss['number']}...")
        try:
            subprocess.run([
                "gh", "issue", "comment", str(s2_iss["number"]), "--body", comment_body
            ], check=True)
            print(f"  Successfully commented on #{s2_iss['number']}!")
            comments_posted += 1
        except Exception as e:
            print(f"  Failed to post comment on #{s2_iss['number']}: {e}", file=sys.stderr)

    print(f"Scan complete. Total linking comments posted: {comments_posted}.")

if __name__ == "__main__":
    main()
