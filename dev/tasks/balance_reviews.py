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
TEAM = [
    "acpana",
    "anfernee",
    "anhdle-sso",
    "barney-s",
    "gemmahou",
    "maqiuyujoyce"
]

def extract_issue_ids(title, body):
    text = f"{title}\n{body or ''}"
    # Match all instances of #<number>
    return set(int(num) for num in re.findall(r'#(\d+)', text))

def fetch_open_prs():
    print("Fetching all open pull requests...", flush=True)
    cmd = [
        "gh", "api", "--paginate",
        "repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100",
        "--jq", ".[] | {number, title, body, labels: [.labels[].name], requested_reviewers: [.requested_reviewers[].login]}"
    ]
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, check=True)
    prs = []
    for line in result.stdout.strip().split("\n"):
        if line.strip():
            try:
                prs.append(json.loads(line))
            except Exception as e:
                print(f"Error parsing line: {line}\nError: {e}", file=sys.stderr)
    return prs

def compute_assignments(prs, team):
    """
    Computes review assignments based on the current workload, workflow affinity, and bounds.
    Returns:
        workload (dict): updated workload dictionary
        tracking_issue_to_reviewer (dict): updated issue-to-reviewer dictionary
        candidates (list): list of unassigned candidate PRs identified
        assignments (list): list of tuples (pr_number, selected_reviewer, title)
    """
    workload = {member: 0 for member in team}
    tracking_issue_to_reviewer = {}

    # Step 1: Audit workloads and map tracking issues
    for pr in prs:
        pr_reviewers = pr.get("requested_reviewers", [])
        assigned_team_reviewers = [r for r in pr_reviewers if r in workload]
        
        if assigned_team_reviewers:
            for reviewer in assigned_team_reviewers:
                workload[reviewer] += 1
                
            issue_ids = extract_issue_ids(pr.get("title", ""), pr.get("body", ""))
            for issue_id in issue_ids:
                for reviewer in assigned_team_reviewers:
                    tracking_issue_to_reviewer[issue_id] = reviewer

    # Step 2: Fetch unassigned candidate PRs
    candidates = []
    for pr in prs:
        labels = pr.get("labels", [])
        if "ready-for-human" not in labels:
            continue
            
        pr_reviewers = pr.get("requested_reviewers", [])
        assigned_team_reviewers = [r for r in pr_reviewers if r in workload]
        
        if not assigned_team_reviewers:
            candidates.append(pr)

    # Step 3: Assignment Algorithm
    assignments = []
    
    for pr in candidates:
        pr_num = pr["number"]
        title = pr["title"]
        body = pr["body"]
        
        # Extract workflow tracking issues
        issue_ids = extract_issue_ids(title, body)
        
        selected_user = None
        
        # Priority 1: Workflow Affinity
        for issue_id in issue_ids:
            reviewer = tracking_issue_to_reviewer.get(issue_id)
            if reviewer and workload[reviewer] < 10:
                selected_user = reviewer
                break
                
        # Priority 2: Underloaded Balancing (C_user < 5)
        if not selected_user:
            underloaded = [m for m in team if workload[m] < 5]
            if underloaded:
                selected_user = min(underloaded, key=lambda m: workload[m])

        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        if not selected_user:
            available = [m for m in team if workload[m] < 10]
            if available:
                selected_user = min(available, key=lambda m: workload[m])
                
        if selected_user:
            workload[selected_user] += 1
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_user
            assignments.append((pr_num, selected_user, title))

        # Termination Condition: Stop assigning if every team member has reached 10 assigned open reviews
        if all(workload[m] >= 10 for m in team):
            break

    return workload, tracking_issue_to_reviewer, candidates, assignments

def main():
    prs = fetch_open_prs()
    print(f"Retrieved {len(prs)} open pull requests.", flush=True)

    # Compute assignments
    workload, tracking_issue_to_reviewer, candidates, assignments = compute_assignments(prs, TEAM)

    print("\nCurrent workload audit:")
    for member, count in workload.items():
        print(f"  {member}: {count} open reviews")

    print(f"\nFound {len(candidates)} unassigned ready-for-human candidate PR(s).")
    for c in candidates:
        print(f"  Candidate PR #{c['number']}: {c['title']}")

    # Print planned assignments
    print("\nPlanned assignments:")
    for pr_num, user, title in assignments:
        print(f"  PR #{pr_num} -> {user} ('{title}')")

    # Step 4: Execute Assignments & Exit
    if assignments:
        print(f"\nExecuting {len(assignments)} assignment(s)...")
        for pr_num, user, title in assignments:
            print(f"Assigning {user} to PR #{pr_num}...")
            cmd = [
                "gh", "pr", "edit", str(pr_num),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", user
            ]
            res = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            if res.returncode == 0:
                print(f"Successfully assigned {user} to PR #{pr_num}.")
            else:
                print(f"Failed to assign {user} to PR #{pr_num}: {res.stderr.strip()}", file=sys.stderr)
        
        print("\nUpdated workload after assignments:")
        for member, count in workload.items():
            print(f"  {member}: {count} open reviews")
    else:
        print("\nNo new assignments made.")

if __name__ == "__main__":
    main()
