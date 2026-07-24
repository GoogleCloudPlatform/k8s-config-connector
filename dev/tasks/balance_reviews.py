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

# Define the team of reviewers
TEAM = {"acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"}

def run_command(cmd):
    result = subprocess.run(cmd, capture_output=True, text=True, check=True)
    return result.stdout

def extract_issues(text):
    if not text:
        return set()
    # Find all pattern like #12345
    matches = re.findall(r'#(\d+)', text)
    return {int(num) for num in matches}

def main():
    dry_run = "--dry-run" in sys.argv or "-n" in sys.argv
    if dry_run:
        print("Executing in DRY RUN mode. No assignments will be written to GitHub.")

    # Fetch all open PRs using REST API to bypass GraphQL read:org scope limits
    print("Fetching all open Pull Requests from GoogleCloudPlatform/k8s-config-connector...")
    cmd = ["gh", "api", "repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100", "--paginate", "--slurp"]
    try:
        output = run_command(cmd)
        pages = json.loads(output)
    except Exception as e:
        print(f"Error fetching open PRs: {e}", file=sys.stderr)
        sys.exit(1)

    all_prs = []
    for page in pages:
        if page:
            all_prs.extend(page)

    print(f"Found {len(all_prs)} open PRs.")

    # 1. Audit current workloads and build tracking_issue_to_reviewer
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}

    for pr in all_prs:
        reviewers = [r['login'] for r in pr.get('requested_reviewers', []) or [] if r.get('login') in TEAM]
        # Track workload
        for rev in reviewers:
            workload[rev] += 1
        
        # Build tracking_issue_to_reviewer if there are team reviewers
        if reviewers:
            primary_reviewer = reviewers[0]
            issues = extract_issues(pr.get('title', '')) | extract_issues(pr.get('body', ''))
            for issue_id in issues:
                tracking_issue_to_reviewer[issue_id] = primary_reviewer

    print("\n--- Initial Team Workloads ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned reviews")

    print(f"\nFound {len(tracking_issue_to_reviewer)} issue-to-reviewer mappings in open PRs.")

    # 2. Filter unassigned candidate PRs labeled 'ready-for-human'
    candidate_prs = []
    for pr in all_prs:
        labels = [l.get('name') for l in pr.get('labels', []) or [] if l.get('name')]
        if 'ready-for-human' not in labels:
            continue
        
        # Filter to those that do NOT have any member of TEAM in requested_reviewers
        reviewers = [r.get('login') for r in pr.get('requested_reviewers', []) or [] if r.get('login') in TEAM]
        if not reviewers:
            candidate_prs.append(pr)

    # Sort candidate PRs oldest first (PR number ascending)
    candidate_prs.sort(key=lambda x: x.get('number', 0))

    print(f"\nFound {len(candidate_prs)} unassigned 'ready-for-human' candidate PRs:")
    for pr in candidate_prs:
        print(f"  #{pr.get('number')}: {pr.get('title')}")

    # 3. Assignment Algorithm
    assignments = []  # list of (pr_number, reviewer, reason)

    for pr in candidate_prs:
        # Check termination condition: are all members at 10?
        if all(workload[m] >= 10 for m in TEAM):
            print("\nHard ceiling of 10 reviews reached for all team members. Stopping assignments.")
            break

        pr_number = pr.get('number')
        pr_title = pr.get('title', '')
        pr_body = pr.get('body') or ""
        
        # Extract workflow issues
        pr_issues = extract_issues(pr_title) | extract_issues(pr_body)
        
        selected_reviewer = None
        selection_reason = ""

        # Priority 1: Workflow Affinity
        # "If ANY extracted issue ID matches a key in tracking_issue_to_reviewer, and that reviewer's count < 10, select that reviewer."
        matching_reviewers = []
        for issue_id in pr_issues:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    matching_reviewers.append(rev)
        
        if matching_reviewers:
            # Select the reviewer with workflow affinity. If there are multiple, choose the one with less workload.
            selected_reviewer = min(matching_reviewers, key=lambda m: workload[m])
            selection_reason = f"Workflow affinity (matched issues {pr_issues})"

        # Priority 2: Underloaded Balancing (workload < 5)
        # "Otherwise, if any team members have workload < 5, select the member with the lowest workload."
        if not selected_reviewer:
            underloaded = [m for m in TEAM if workload[m] < 5]
            if underloaded:
                selected_reviewer = min(underloaded, key=lambda m: (workload[m], m))
                selection_reason = f"Underloaded balancing (workload < 5)"

        # Priority 3: Capacity Absorption (5 <= workload < 10)
        # "Otherwise, if all team members have reached workload >= 5 but unassigned candidate PRs remain, select the member with the lowest workload who has workload < 10."
        if not selected_reviewer:
            available = [m for m in TEAM if workload[m] < 10]
            if available:
                selected_reviewer = min(available, key=lambda m: (workload[m], m))
                selection_reason = f"Capacity absorption (5 <= workload < 10)"

        if selected_reviewer:
            # Update state
            workload[selected_reviewer] += 1
            # Add all referenced issue IDs to tracking_issue_to_reviewer
            for issue_id in pr_issues:
                tracking_issue_to_reviewer[issue_id] = selected_reviewer
            
            assignments.append((pr_number, selected_reviewer, selection_reason))
            print(f"Queued PR #{pr_number} -> {selected_reviewer} ({selection_reason})")
        else:
            print(f"PR #{pr_number} could not be assigned (all eligible reviewers at capacity limit)")

    # 4. Execute Assignments
    if not assignments:
        print("\nNo new assignments to make.")
    else:
        print("\n--- Executing Assignments ---")
        for pr_number, reviewer, reason in assignments:
            if dry_run:
                print(f"[DRY-RUN] Would assign PR #{pr_number} to {reviewer} ({reason})")
            else:
                print(f"Assigning PR #{pr_number} to {reviewer}...")
                cmd = ["gh", "pr", "edit", str(pr_number), "--repo", "GoogleCloudPlatform/k8s-config-connector", "--add-reviewer", reviewer]
                try:
                    run_command(cmd)
                    print(f"  Successfully assigned PR #{pr_number} to {reviewer}")
                except Exception as e:
                    print(f"  Failed to assign PR #{pr_number} to {reviewer}: {e}", file=sys.stderr)

    print("\n--- Final Team Workloads ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned reviews")

if __name__ == "__main__":
    main()
