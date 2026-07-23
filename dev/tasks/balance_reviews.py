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
        print(f"Fetching open PRs page {page}...", file=sys.stderr)
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
        prs.extend(data)
        if len(data) < 100:
            break
        page += 1
    return prs

def extract_issue_ids(title, body):
    issues = set()
    for text in [title, body]:
        if not text:
            continue
        # Find #1234
        for m in re.findall(r'#(\d+)', text):
            issues.add(int(m))
        # Find URLs
        for m in re.findall(r'github\.com/GoogleCloudPlatform/k8s-config-connector/(?:issues|pull)/(\d+)', text):
            issues.add(int(m))
    return issues

def main():
    print("Auditing current team workloads and workflow mappings...", file=sys.stderr)
    prs = fetch_open_prs()
    print(f"Total open PRs retrieved: {len(prs)}", file=sys.stderr)
    
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}
    candidates = []
    
    # Step 1 & 2: Build workload map and workflow affinity map, filter candidate PRs
    for pr in prs:
        num = pr["number"]
        title = pr["title"] or ""
        body = pr["body"] or ""
        labels = {lbl["name"] for lbl in pr.get("labels", []) if "name" in lbl}
        req_reviewers = {rev["login"] for rev in pr.get("requested_reviewers", []) if "login" in rev}
        
        # Check if any team member is already requested as a reviewer
        team_reviewers = req_reviewers.intersection(TEAM)
        
        if team_reviewers:
            for r in team_reviewers:
                workload[r] += 1
            
            # Build tracking issue map
            referenced_issues = extract_issue_ids(title, body)
            referenced_issues.discard(num) # Discard own PR number
            
            for r in team_reviewers:
                for issue_id in referenced_issues:
                    tracking_issue_to_reviewer[issue_id] = r
        else:
            # Candidate for assignment
            if "ready-for-human" in labels:
                candidates.append(pr)

    # Sort candidates by number ascending (oldest first)
    candidates.sort(key=lambda x: x["number"])
    
    print("\nInitial Workloads:", file=sys.stderr)
    for m in sorted(TEAM):
        print(f"  {m}: {workload[m]} assigned", file=sys.stderr)
    print(f"Found {len(candidates)} unassigned 'ready-for-human' candidate PRs.", file=sys.stderr)
    
    assignments = []
    
    # Step 3: Assignment Algorithm
    for pr in candidates:
        num = pr["number"]
        title = pr["title"] or ""
        body = pr["body"] or ""
        
        # 1. Extract Workflow Tracking Issues
        referenced_issues = extract_issue_ids(title, body)
        referenced_issues.discard(num)
        
        selected_user = None
        reason = ""
        
        # 2. Select Reviewer by Priority
        # Priority 1 (Workflow Affinity):
        affinity_reviewers = []
        for issue_id in referenced_issues:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    affinity_reviewers.append(rev)
        
        if affinity_reviewers:
            # Select the one with lowest workload among affinity reviewers
            affinity_reviewers.sort(key=lambda r: workload[r])
            selected_user = affinity_reviewers[0]
            reason = f"Workflow Affinity (matched referenced issue in tracking issues, workload {workload[selected_user]})"
            
        # Priority 2 (Underloaded Balancing, C_user < 5):
        if not selected_user:
            underloaded = [m for m in TEAM if workload[m] < 5]
            if underloaded:
                underloaded.sort(key=lambda m: workload[m])
                selected_user = underloaded[0]
                reason = f"Underloaded Balancing (workload {workload[selected_user]} < 5)"
                
        # Priority 3 (Capacity Absorption, 5 <= C_user < 10):
        if not selected_user:
            available = [m for m in TEAM if workload[m] < 10]
            if available:
                available.sort(key=lambda m: workload[m])
                selected_user = available[0]
                reason = f"Capacity Absorption (workload {workload[selected_user]} < 10)"
                
        # Ceiling Reached:
        if not selected_user:
            print(f"Skipping PR #{num}: All team members have reached the ceiling of 10 assigned reviews.", file=sys.stderr)
            continue
            
        # 3. Update State & Queue Assignment
        workload[selected_user] += 1
        for issue_id in referenced_issues:
            tracking_issue_to_reviewer[issue_id] = selected_user
            
        assignments.append((num, selected_user, reason))
        
        # 4. Termination Condition:
        if all(workload[m] >= 10 for m in TEAM):
            print("Stopping assignments: All team members have reached the maximum hard ceiling of 10 open reviews.", file=sys.stderr)
            break

    # Step 4: Execute Assignments & Print Summary
    print("\n--- Queueing and executing assignments ---", file=sys.stderr)
    successful_assignments = []
    for num, selected_user, reason in assignments:
        cmd = [
            "gh", "pr", "edit", str(num),
            "--repo", "GoogleCloudPlatform/k8s-config-connector",
            "--add-reviewer", selected_user
        ]
        print(f"Assigning PR #{num} to {selected_user} ...", file=sys.stderr)
        res = subprocess.run(cmd, capture_output=True, text=True)
        if res.returncode == 0:
            print(f"Successfully assigned PR #{num} to {selected_user}.", file=sys.stderr)
            successful_assignments.append((num, selected_user, reason))
        else:
            print(f"Error assigning PR #{num} to {selected_user}: {res.stderr.strip()}", file=sys.stderr)

    print("\n==========================================")
    print("PR Review Assignment Balancer Summary")
    print("==========================================")
    if not successful_assignments:
        print("No new PR assignments were made.")
    else:
        for num, user, reason in successful_assignments:
            print(f"- Assigned PR #{num} to {user} ({reason})")
            
    print("\nUpdated Team Workloads:")
    for m in sorted(TEAM):
        print(f"  {m}: {workload[m]} assigned reviews")
    print("==========================================")

if __name__ == "__main__":
    main()
