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
import sys
import subprocess
import argparse

TEAM = ["acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"]
REPO = "GoogleCloudPlatform/k8s-config-connector"

def run_command(cmd):
    """Utility to run a shell command and return its output."""
    res = subprocess.run(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    if res.returncode != 0:
        print(f"Command failed: {cmd}\nError: {res.stderr}", file=sys.stderr)
        raise RuntimeError(f"Command failed with exit code {res.returncode}")
    return res.stdout

def fetch_open_prs():
    """Fetches all open PRs using paginated GitHub REST API."""
    print("Fetching open PRs from GitHub...")
    # We query the pulls REST API to avoid read:org scope issues in GraphQL
    cmd = (
        f'gh api --paginate "repos/{REPO}/pulls?state=open&per_page=100" '
        '--jq \'.[] | {number: .number, title: .title, body: .body, labels: [.labels[].name], requested_reviewers: [.requested_reviewers[].login]}\''
    )
    stdout = run_command(cmd)
    
    prs = []
    for line in stdout.strip().split("\n"):
        if line.strip():
            try:
                prs.append(json.loads(line))
            except json.JSONDecodeError as e:
                print(f"Failed to decode line: {line}\nError: {e}", file=sys.stderr)
    return prs

def extract_issue_ids(title, body):
    """Extracts all referenced issue IDs of the form #<NUMBER> from title and body."""
    text = (title or "") + " " + (body or "")
    # Find all pattern like #1234
    matches = re.findall(r'#(\d+)', text)
    return sorted(list(set(int(num) for num in matches)))

def balance_workloads(prs, dry_run=False):
    # Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}

    for pr in prs:
        req_reviewers = pr.get("requested_reviewers", [])
        team_reviewers = [r for r in req_reviewers if r in TEAM]
        
        # Calculate workload
        for r in team_reviewers:
            workload[r] += 1
            
        # Map tracking issues
        if team_reviewers:
            issue_ids = extract_issue_ids(pr.get("title"), pr.get("body"))
            # If multiple team members are assigned, map to all of them
            for r in team_reviewers:
                for iid in issue_ids:
                    tracking_issue_to_reviewer[iid] = r

    print("\n=== Initial Audit Summary ===")
    print("Current Workloads (Count of open assigned reviews):")
    for r in sorted(TEAM):
        print(f"  - {r}: {workload[r]}")
    
    print(f"Mapped tracking issues to current reviewers: {len(tracking_issue_to_reviewer)}")

    # Step 2: Fetch Unassigned Candidate PRs
    candidates = []
    for pr in prs:
        labels = pr.get("labels", [])
        if "ready-for-human" in labels:
            req_reviewers = pr.get("requested_reviewers", [])
            team_reviewers = [r for r in req_reviewers if r in TEAM]
            if not team_reviewers:
                candidates.append(pr)

    print(f"\nUnassigned 'ready-for-human' candidate PRs found: {len(candidates)}")
    for pr in candidates:
        print(f"  - PR #{pr['number']}: {pr['title']}")

    # Step 3: Assignment Algorithm
    assignments = []
    
    for pr in candidates:
        pr_num = pr["number"]
        pr_title = pr["title"]
        issue_ids = extract_issue_ids(pr_title, pr.get("body"))
        
        # Check termination condition: Are all team members at the ceiling?
        if all(workload[r] >= 10 for r in TEAM):
            print("\n[Ceiling Reached] All team members have reached the maximum hard ceiling of 10 assigned open reviews. Stopping assignments.")
            break

        selected_user = None
        priority_level = None

        # Priority 1: Workflow Affinity
        # If ANY extracted issue ID matches a key in tracking_issue_to_reviewer, and that reviewer's count < 10, select that reviewer.
        for iid in issue_ids:
            if iid in tracking_issue_to_reviewer:
                reviewer = tracking_issue_to_reviewer[iid]
                if workload[reviewer] < 10:
                    selected_user = reviewer
                    priority_level = f"Priority 1 (Workflow Affinity via Issue #{iid})"
                    break
        
        # Priority 2: Underloaded Balancing (C_user < 5)
        if not selected_user:
            underloaded = [r for r in TEAM if workload[r] < 5]
            if underloaded:
                # Select the member with the lowest workload. Use handle alphabetically to break ties.
                selected_user = min(underloaded, key=lambda x: (workload[x], x))
                priority_level = f"Priority 2 (Underloaded Balancing, workload = {workload[selected_user]} < 5)"

        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        if not selected_user:
            eligible = [r for r in TEAM if workload[r] < 10]
            if eligible:
                # Select the member with the lowest workload. Use handle alphabetically to break ties.
                selected_user = min(eligible, key=lambda x: (workload[x], x))
                priority_level = f"Priority 3 (Capacity Absorption, workload = {workload[selected_user]} < 10)"

        if selected_user:
            # Update state
            workload[selected_user] += 1
            for iid in issue_ids:
                tracking_issue_to_reviewer[iid] = selected_user
            
            assignments.append((pr_num, pr_title, selected_user, priority_level))
            print(f"Scheduled PR #{pr_num} to {selected_user} via {priority_level}")
        else:
            print(f"Could not assign PR #{pr_num} (no eligible team member under the ceiling of 10).")

    # Step 4: Execute Assignments
    print("\n=== Assignment Execution ===")
    if not assignments:
        print("No new assignments to execute.")
    else:
        for pr_num, title, user, priority in assignments:
            edit_cmd = f"gh pr edit {pr_num} --repo {REPO} --add-reviewer {user}"
            print(f"Executing: {edit_cmd}")
            if not dry_run:
                try:
                    run_command(edit_cmd)
                    print(f"  Successfully assigned PR #{pr_num} to {user}")
                except Exception as e:
                    print(f"  Failed to assign PR #{pr_num}: {e}", file=sys.stderr)
            else:
                print(f"  [Dry Run] Would assign PR #{pr_num} to {user}")

    print("\n=== Final Workload Summary ===")
    for r in sorted(TEAM):
        print(f"  - {r}: {workload[r]}")

def main():
    parser = argparse.ArgumentParser(description="PR Review Workload Balancer")
    parser.add_argument("--dry-run", action="store_true", help="Run assignment logic and print summary without making changes on GitHub")
    args = parser.parse_args()

    prs = fetch_open_prs()
    balance_workloads(prs, dry_run=args.dry_run)

if __name__ == "__main__":
    main()
