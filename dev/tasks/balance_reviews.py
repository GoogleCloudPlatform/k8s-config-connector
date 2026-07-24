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

import json
import re
import subprocess
import sys

# Define the team members
TEAM = {"acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"}

def extract_issue_ids(text):
    if not text:
        return []
    # Find all references of the form #<NUMBER>
    # Note: tracking issues, sub-issues, fixes are all matched by #<NUMBER>
    return [int(x) for x in re.findall(r"#(\d+)", text)]

def fetch_all_open_prs():
    print("Fetching all open PRs from GoogleCloudPlatform/k8s-config-connector...")
    all_prs = []
    page = 1
    while True:
        # Fetching 100 per page to minimize API calls
        cmd = [
            "gh", "api",
            f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"
        ]
        res = subprocess.run(cmd, capture_output=True, text=True)
        if res.returncode != 0:
            print(f"Error fetching page {page}: {res.stderr}", file=sys.stderr)
            sys.exit(1)
        data = json.loads(res.stdout)
        if not data:
            break
        all_prs.extend(data)
        page += 1
    print(f"Fetched {len(all_prs)} open PRs.")
    return all_prs

def main():
    prs = fetch_all_open_prs()

    # Step 1: Audit workloads & Map tracking issues to reviewers
    workloads = {user: 0 for user in TEAM}
    tracking_issue_to_reviewer = {}
    candidates = []

    for pr in prs:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "") or ""
        labels = [label.get("name") for label in pr.get("labels", []) if label.get("name")]
        
        # Extract issues
        issue_ids = extract_issue_ids(title + "\n" + body)

        # Audit workloads based on requested reviewers
        review_requests = pr.get("requested_reviewers", [])
        team_reviewers_requested = []
        for r in review_requests:
            login = r.get("login")
            if login in TEAM:
                workloads[login] += 1
                team_reviewers_requested.append(login)

        # Map tracked issues to the active reviewer
        if team_reviewers_requested:
            primary_reviewer = team_reviewers_requested[0]
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = primary_reviewer

        # Step 2: Fetch Unassigned Candidate PRs
        # Must have "ready-for-human" label and no existing team reviewers requested
        if "ready-for-human" in labels:
            if not team_reviewers_requested:
                candidates.append({
                    "number": number,
                    "title": title,
                    "body": body,
                    "issue_ids": issue_ids,
                })

    print("\nInitial Workloads:")
    for user in sorted(TEAM):
        print(f"  {user}: {workloads[user]}")

    print(f"\nFound {len(candidates)} unassigned 'ready-for-human' candidate PRs.")
    for c in candidates:
        print(f"  #{c['number']}: \"{c['title']}\" (Issues: {c['issue_ids']})")

    # Step 3: Assignment Algorithm
    assignments = []
    
    # Sort candidates by PR number (ascending) to be deterministic and process oldest first
    candidates.sort(key=lambda x: x["number"])

    for pr in candidates:
        # Check termination condition: every team member has reached the hard ceiling of 10
        if all(workloads[user] >= 10 for user in TEAM):
            print("\nHard ceiling of 10 reached for all team members. Stopping assignments.")
            break

        issue_ids = pr["issue_ids"]
        selected_user = None

        # Priority 1 (Workflow Affinity)
        # Check if any extracted issue ID matches an existing reviewer whose workload is < 10
        for issue_id in issue_ids:
            reviewer = tracking_issue_to_reviewer.get(issue_id)
            if reviewer and workloads[reviewer] < 10:
                selected_user = reviewer
                print(f"  [Affinity Match] PR #{pr['number']} matches issue #{issue_id} reviewed by {reviewer}")
                break

        # Priority 2 (Underloaded Balancing, C_user < 5)
        if not selected_user:
            underloaded = [u for u in TEAM if workloads[u] < 5]
            if underloaded:
                # Select member with lowest workload, deterministic tie-breaking alphabetically
                selected_user = min(underloaded, key=lambda u: (workloads[u], u))
                print(f"  [Underloaded Match] PR #{pr['number']} assigned to {selected_user} with workload {workloads[selected_user]}")

        # Priority 3 (Capacity Absorption, 5 <= C_user < 10)
        if not selected_user:
            capable = [u for u in TEAM if 5 <= workloads[u] < 10]
            if capable:
                selected_user = min(capable, key=lambda u: (workloads[u], u))
                print(f"  [Capacity Match] PR #{pr['number']} assigned to {selected_user} with workload {workloads[selected_user]}")

        # Execute selected assignment
        if selected_user:
            workloads[selected_user] += 1
            # Update tracking for all issues referenced by this PR
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_user
            
            assignments.append((pr["number"], selected_user))
        else:
            print(f"  [No Reviewer] Could not find a suitable reviewer for PR #{pr['number']}")

    # Step 4: Execute Assignments & Print Summary
    print("\nProposed Assignments:")
    if not assignments:
        print("  No assignments to make.")
    else:
        for pr_num, user in assignments:
            print(f"  Assigning PR #{pr_num} to {user}")

        # Actually execute assignments
        print("\nExecuting assignments on GitHub...")
        for pr_num, user in assignments:
            cmd = [
                "gh", "pr", "edit", str(pr_num),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", user
            ]
            print(f"Running: {' '.join(cmd)}")
            res = subprocess.run(cmd, capture_output=True, text=True)
            if res.returncode != 0:
                print(f"Failed to assign PR #{pr_num} to {user}: {res.stderr}", file=sys.stderr)
            else:
                print(f"Successfully assigned PR #{pr_num} to {user}")

    print("\nFinal Workloads:")
    for user in sorted(TEAM):
        print(f"  {user}: {workloads[user]}")

if __name__ == "__main__":
    main()
