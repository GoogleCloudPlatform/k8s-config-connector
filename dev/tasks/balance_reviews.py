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

import subprocess
import json
import re
import sys

# Define the team members
TEAM = {
    "acpana",
    "anfernee",
    "anhdle-sso",
    "barney-s",
    "gemmahou",
    "maqiuyujoyce",
}

def get_open_prs():
    """Fetches all open PRs in GoogleCloudPlatform/k8s-config-connector."""
    print("Fetching open PRs from GitHub REST API...")
    cmd = [
        "gh", "api", "--paginate",
        "repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100"
    ]
    res = subprocess.run(cmd, capture_output=True, text=True, check=True)
    raw = res.stdout.strip()
    if not raw:
        return []
    
    prs = []
    decoder = json.JSONDecoder()
    pos = 0
    while pos < len(raw):
        # Skip leading whitespace
        while pos < len(raw) and raw[pos].isspace():
            pos += 1
        if pos >= len(raw):
            break
        try:
            obj, count = decoder.raw_decode(raw[pos:])
            prs.extend(obj)
            pos += count
        except json.JSONDecodeError as e:
            print(f"Error decoding JSON at position {pos}: {e}", file=sys.stderr)
            break
    return prs

def extract_issue_ids(title, body):
    """Extracts all issue references from PR title and body via regex matching #<NUMBER>."""
    text = f"{title or ''}\n{body or ''}"
    return set(re.findall(r'#(\d+)', text))

def main():
    try:
        prs = get_open_prs()
    except subprocess.CalledProcessError as e:
        print(f"Failed to fetch open PRs from GitHub: {e}", file=sys.stderr)
        print(f"stdout: {e.stdout}", file=sys.stderr)
        print(f"stderr: {e.stderr}", file=sys.stderr)
        sys.exit(1)

    print(f"Retrieved {len(prs)} open PRs.")

    # Initialize data structures
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}

    # Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping
    for pr in prs:
        requested_logins = [
            r["login"] for r in pr.get("requested_reviewers", [])
            if r.get("login") in TEAM
        ]
        for login in requested_logins:
            workload[login] += 1
            
        if requested_logins:
            # Maintain context continuity: associate referenced issue IDs with the assigned reviewer
            reviewer = requested_logins[0]
            issue_ids = extract_issue_ids(pr.get("title", ""), pr.get("body", ""))
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = reviewer

    print("\n=== Initial Workload Audit ===")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned open reviews")

    # Step 2: Fetch Unassigned Candidate PRs
    candidates = []
    for pr in prs:
        labels = [lbl["name"] for lbl in pr.get("labels", [])]
        if "ready-for-human" not in labels:
            continue
        
        # Candidate must NOT have any member of k8s-config-connector-team listed in requested_reviewers
        requested_logins = [
            r["login"] for r in pr.get("requested_reviewers", [])
            if r.get("login") in TEAM
        ]
        if not requested_logins:
            candidates.append(pr)

    # Sort candidates by PR number ascending (oldest first)
    candidates.sort(key=lambda x: x["number"])
    print(f"\nIdentified {len(candidates)} unassigned candidate 'ready-for-human' PR(s).")

    # Step 3: Assignment Algorithm
    assignments = []

    for pr in candidates:
        # Check termination condition: all team members have reached the hard ceiling of 10
        if all(workload[m] >= 10 for m in TEAM):
            print("\nTermination Condition Met: All team members have reached the maximum ceiling of 10 reviews.")
            break

        pr_num = pr["number"]
        title = pr.get("title", "")
        body = pr.get("body", "")
        issue_ids = extract_issue_ids(title, body)
        
        selected_reviewer = None
        priority_reason = ""

        # Priority 1 (Workflow Affinity)
        # If any extracted issue ID matches a key in tracking_issue_to_reviewer, and that reviewer's count < 10, select them.
        affinity_reviewers = []
        for issue_id in issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    affinity_reviewers.append(rev)
        if affinity_reviewers:
            # Select the reviewer with the lowest count (break ties alphabetically)
            selected_reviewer = min(affinity_reviewers, key=lambda x: (workload[x], x))
            priority_reason = "Workflow Affinity (Priority 1)"

        # Priority 2 (Underloaded Balancing, C_user < 5)
        if not selected_reviewer:
            underloaded = [m for m in TEAM if workload[m] < 5]
            if underloaded:
                selected_reviewer = min(underloaded, key=lambda x: (workload[x], x))
                priority_reason = "Underloaded Balancing (Priority 2)"

        # Priority 3 (Capacity Absorption, 5 <= C_user < 10)
        if not selected_reviewer:
            available = [m for m in TEAM if workload[m] < 10]
            if available:
                selected_reviewer = min(available, key=lambda x: (workload[x], x))
                priority_reason = "Capacity Absorption (Priority 3)"

        # Ceiling Reached for all members (fallback check)
        if not selected_reviewer:
            print(f"Skipping PR #{pr_num} ('{title}'): No reviewer has capacity under the hard ceiling of 10.")
            continue

        # Update State & Queue Assignment
        workload[selected_reviewer] += 1
        for issue_id in issue_ids:
            tracking_issue_to_reviewer[issue_id] = selected_reviewer

        assignments.append({
            "pr_number": pr_num,
            "reviewer": selected_reviewer,
            "title": title,
            "reason": priority_reason
        })

    # Step 4: Execute Assignments
    if not assignments:
        print("\nNo assignments queued.")
    else:
        print(f"\nExecuting {len(assignments)} queued review assignment(s):")
        for a in assignments:
            pr_num = a["pr_number"]
            reviewer = a["reviewer"]
            print(f"  Assigning PR #{pr_num} ('{a['title']}') to {reviewer} via {a['reason']}...")
            try:
                subprocess.run([
                    "gh", "pr", "edit", str(pr_num),
                    "--repo", "GoogleCloudPlatform/k8s-config-connector",
                    "--add-reviewer", reviewer
                ], check=True)
            except subprocess.CalledProcessError as e:
                print(f"  Error assigning PR #{pr_num} to {reviewer}: {e}", file=sys.stderr)

    # Print a summary log of assignments made
    print("\n=== Assignment Summary ===")
    if not assignments:
        print("  No new review assignments were made.")
    for a in assignments:
        print(f"  PR #{a['pr_number']}: Assigned to {a['reviewer']} ({a['reason']})")

    print("\n=== Final Workload Status ===")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned open reviews")

if __name__ == "__main__":
    main()
