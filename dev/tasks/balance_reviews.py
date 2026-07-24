#!/usr/bin/env python3
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
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
TEAM = [
    'acpana',
    'anfernee',
    'anhdle-sso',
    'barney-s',
    'gemmahou',
    'maqiuyujoyce'
]

REPO = "GoogleCloudPlatform/k8s-config-connector"

def run_cmd(args):
    """Runs a shell command and returns stdout. Exits on failure."""
    try:
        result = subprocess.run(args, capture_output=True, text=True, check=True)
        return result.stdout
    except subprocess.CalledProcessError as e:
        print(f"Error running command {' '.join(args)}:")
        print(f"Exit code: {e.returncode}")
        print(f"Stdout: {e.stdout}")
        print(f"Stderr: {e.stderr}")
        sys.exit(e.returncode)

def extract_issue_ids(title, body):
    """Extracts all issue/PR numbers from title and body."""
    ids = set()
    for text in [title, body]:
        if text:
            # Matches pattern of # followed by digits (e.g. #1234)
            matches = re.findall(r'#(\d+)', text)
            for m in matches:
                ids.add(int(m))
    return ids

def main():
    print("=== Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping ===")
    
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}
    assigned_pr_numbers = set()
    
    for member in TEAM:
        print(f"Auditing assignments for {member}...")
        args = [
            "gh", "search", "prs",
            "--repo", REPO,
            "--state", "open",
            "--review-requested", member,
            "--json", "number,title,body",
            "--limit", "100"
        ]
        output = run_cmd(args)
        prs = json.loads(output)
        
        workload[member] = len(prs)
        print(f"  {member} has {len(prs)} open assigned reviews.")
        
        for pr in prs:
            num = pr['number']
            title = pr.get('title', '')
            body = pr.get('body', '')
            assigned_pr_numbers.add(num)
            
            # Extract workflow tracking issues
            issue_ids = extract_issue_ids(title, body)
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = member
                
    print("\nInitial Workloads:")
    for member, count in workload.items():
        print(f"  {member}: {count}")
    
    print("\n=== Step 2: Fetching Unassigned Candidate PRs ===")
    
    args = [
        "gh", "pr", "list",
        "--repo", REPO,
        "--label", "ready-for-human",
        "--state", "open",
        "--limit", "200",
        "--json", "number,title,body"
    ]
    output = run_cmd(args)
    candidate_prs = json.loads(output)
    
    # Filter candidate PRs that do not currently have any team member listed in reviewRequests
    unassigned_candidates = []
    for pr in candidate_prs:
        num = pr['number']
        if num not in assigned_pr_numbers:
            unassigned_candidates.append(pr)
            
    print(f"Found {len(candidate_prs)} total 'ready-for-human' open PRs.")
    print(f"Found {len(unassigned_candidates)} unassigned 'ready-for-human' PRs.")
    
    print("\n=== Step 3: Assignment Algorithm ===")
    
    assignments_queue = []
    
    for pr in unassigned_candidates:
        num = pr['number']
        title = pr.get('title', '')
        body = pr.get('body', '')
        
        issue_ids = extract_issue_ids(title, body)
        
        # Select reviewer based on priority
        selected_reviewer = None
        priority_reason = ""
        
        # Priority 1: Workflow Affinity
        affinity_reviewers = set()
        for issue_id in issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                reviewer = tracking_issue_to_reviewer[issue_id]
                if workload[reviewer] < 10:
                    affinity_reviewers.add(reviewer)
                    
        if affinity_reviewers:
            # Select the affinity reviewer with the lowest workload
            selected_reviewer = min(affinity_reviewers, key=lambda r: workload[r])
            priority_reason = f"Workflow Affinity (matching issue(s): {', '.join(f'#{i}' for i in issue_ids if tracking_issue_to_reviewer.get(i) == selected_reviewer)})"
        else:
            # Priority 2: Underloaded Balancing (C_user < 5)
            underloaded_members = [m for m in TEAM if workload[m] < 5]
            if underloaded_members:
                selected_reviewer = min(underloaded_members, key=lambda r: workload[r])
                priority_reason = "Underloaded Balancing (< 5 reviews)"
            else:
                # Priority 3: Capacity Absorption (5 <= C_user < 10)
                capacity_members = [m for m in TEAM if workload[m] < 10]
                if capacity_members:
                    selected_reviewer = min(capacity_members, key=lambda r: workload[r])
                    priority_reason = "Capacity Absorption (5-9 reviews)"
                    
        if selected_reviewer:
            # Update State
            workload[selected_reviewer] += 1
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_reviewer
            # Also include the PR's own number as mapped to the selected reviewer
            tracking_issue_to_reviewer[num] = selected_reviewer
            
            assignments_queue.append((num, title, selected_reviewer, priority_reason))
            print(f"Queued PR #{num} ('{title}') -> {selected_reviewer} via {priority_reason}")
        else:
            print(f"Skipped PR #{num} ('{title}') -> All team members have reached the ceiling of 10 assigned reviews.")
            
    print("\n=== Step 4: Execute Assignments & Exit ===")
    
    if not assignments_queue:
        print("No new assignments needed. Current team workloads are balanced.")
        return
        
    print(f"Executing {len(assignments_queue)} review assignments...")
    for num, title, reviewer, reason in assignments_queue:
        print(f"Assigning PR #{num} to {reviewer}...")
        edit_args = [
            "gh", "pr", "edit", str(num),
            "--repo", REPO,
            "--add-reviewer", reviewer
        ]
        run_cmd(edit_args)
        
    print("\nSummary of Assignments Made:")
    for num, title, reviewer, reason in assignments_queue:
        print(f"  - PR #{num} ('{title}') -> {reviewer} [{reason}]")
        
    print("\nFinal Workloads:")
    for member, count in workload.items():
        print(f"  {member}: {count}")

if __name__ == "__main__":
    main()
