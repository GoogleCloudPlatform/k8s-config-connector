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
import subprocess
import re
import sys

# Define the team members as specified in the team definition
TEAM_MEMBERS = {
    'acpana',
    'anfernee',
    'anhdle-sso',
    'barney-s',
    'gemmahou',
    'maqiuyujoyce'
}

def fetch_open_prs():
    prs = []
    page = 1
    while True:
        cmd = ["gh", "api", f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            print(f"Error fetching open PRs page {page}: {result.stderr}", file=sys.stderr)
            break
        page_prs = json.loads(result.stdout)
        if not page_prs:
            break
        prs.extend(page_prs)
        if len(page_prs) < 100:
            break
        page += 1
    return prs

def main():
    print("Fetching open PRs from GoogleCloudPlatform/k8s-config-connector...")
    prs = fetch_open_prs()
    print(f"Total open PRs: {len(prs)}")

    # Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping
    workload = {member: 0 for member in TEAM_MEMBERS}
    tracking_issue_to_reviewer = {}
    candidates = []

    issue_pattern = re.compile(r'#(\d+)')

    for pr in prs:
        number = pr.get('number')
        title = pr.get('title', '')
        body = pr.get('body', '') or ''
        
        # Check requested reviewers
        req_reviewers = pr.get('requested_reviewers', []) or []
        req_logins = [user.get('login') for user in req_reviewers if user]
        
        # Find if any team member is currently a requested reviewer
        assigned_team_members = [login for login in req_logins if login in TEAM_MEMBERS]
        
        # Update workload
        for member in assigned_team_members:
            workload[member] += 1
            
        # Extract all referenced issue IDs from title and body
        issues = set(issue_pattern.findall(title + " " + body))
        
        # If assigned, map all extracted issues to the first assigned reviewer
        if assigned_team_members:
            for issue in issues:
                tracking_issue_to_reviewer[issue] = assigned_team_members[0]
                
        # Check if labeled 'ready-for-human'
        labels = [l.get('name') for l in pr.get('labels', []) if l]
        if 'ready-for-human' in labels:
            if not assigned_team_members:
                candidates.append({
                    'number': number,
                    'title': title,
                    'issues': issues
                })

    print("\n--- Current Workloads ---")
    for member in sorted(TEAM_MEMBERS):
        print(f"  {member}: {workload[member]} assigned open reviews")

    print(f"\nTotal unassigned 'ready-for-human' candidate PRs: {len(candidates)}")
    for cand in candidates:
        print(f"  - #{cand['number']}: {cand['title']} (Issues: {cand['issues']})")

    # Step 3: Assignment Algorithm
    queued_assignments = []

    for cand in candidates:
        pr_num = cand['number']
        issues = cand['issues']
        
        selected_user = None
        
        # Priority 1: Workflow Affinity
        # If any extracted issue ID matches a key in tracking_issue_to_reviewer and that reviewer has < 10 workload
        for issue in sorted(issues):
            if issue in tracking_issue_to_reviewer:
                reviewer = tracking_issue_to_reviewer[issue]
                if workload[reviewer] < 10:
                    selected_user = reviewer
                    print(f"PR #{pr_num}: Priority 1 (Workflow Affinity) selected '{selected_user}' via Issue #{issue}")
                    break
                    
        # Priority 2: Underloaded Balancing (workload < 5)
        if not selected_user:
            underloaded = [m for m in TEAM_MEMBERS if workload[m] < 5]
            if underloaded:
                # Select the member with the lowest workload, using username as tie-breaker
                underloaded.sort(key=lambda m: (workload[m], m))
                selected_user = underloaded[0]
                print(f"PR #{pr_num}: Priority 2 (Underloaded Balancing) selected '{selected_user}' with workload {workload[selected_user]}")
                
        # Priority 3: Capacity Absorption (5 <= workload < 10)
        if not selected_user:
            capacity_members = [m for m in TEAM_MEMBERS if workload[m] < 10]
            if capacity_members:
                # Select the member with the lowest workload, using username as tie-breaker
                capacity_members.sort(key=lambda m: (workload[m], m))
                selected_user = capacity_members[0]
                print(f"PR #{pr_num}: Priority 3 (Capacity Absorption) selected '{selected_user}' with workload {workload[selected_user]}")
                
        if selected_user:
            # Update state & Queue Assignment
            workload[selected_user] += 1
            for issue in issues:
                tracking_issue_to_reviewer[issue] = selected_user
            queued_assignments.append((pr_num, selected_user))
        else:
            print(f"PR #{pr_num}: Skip assignment - All team members are at maximum capacity (10 reviews)")

    # Step 4: Execute Assignments & Exit
    print("\n--- Executing Assignments ---")
    if not queued_assignments:
        print("No new assignments to make.")
    else:
        for pr_num, user in queued_assignments:
            print(f"Assigning PR #{pr_num} to reviewer '{user}'...")
            cmd = ["gh", "pr", "edit", str(pr_num), "--repo", "GoogleCloudPlatform/k8s-config-connector", "--add-reviewer", user]
            result = subprocess.run(cmd, capture_output=True, text=True)
            if result.returncode != 0:
                print(f"Failed to assign PR #{pr_num} to '{user}': {result.stderr}", file=sys.stderr)
            else:
                print(f"Successfully assigned PR #{pr_num} to '{user}'.")

    print("\n--- Final Workloads ---")
    for member in sorted(TEAM_MEMBERS):
        print(f"  {member}: {workload[member]} assigned open reviews")

if __name__ == '__main__':
    main()
