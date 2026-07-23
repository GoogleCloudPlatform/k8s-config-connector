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

def run_command(cmd):
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    if result.returncode != 0:
        print(f"Error running command {' '.join(cmd)}: {result.stderr}", file=sys.stderr)
        sys.exit(result.returncode)
    return result.stdout

def fetch_open_prs():
    print("Fetching open PRs from GoogleCloudPlatform/k8s-config-connector via REST API...")
    # Use gh api --paginate to fetch all open PRs in pages of 100
    cmd = ["gh", "api", "--paginate", "repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100"]
    stdout = run_command(cmd)
    
    prs = []
    decoder = json.JSONDecoder()
    pos = 0
    content = stdout.strip()
    while pos < len(content):
        # Skip leading whitespace
        while pos < len(content) and content[pos].isspace():
            pos += 1
        if pos >= len(content):
            break
        try:
            obj, idx = decoder.raw_decode(content, pos)
            if isinstance(obj, list):
                prs.extend(obj)
            else:
                prs.append(obj)
            pos = idx
        except json.JSONDecodeError as e:
            # Fallback to direct json.loads in case it is formatted as a single array
            try:
                prs = json.loads(content)
                break
            except Exception:
                print(f"Error decoding JSON: {e}", file=sys.stderr)
                sys.exit(1)
            
    print(f"Successfully fetched {len(prs)} open PRs.")
    return prs

def extract_issue_ids(title, body):
    title = title or ""
    body = body or ""
    # Extract hash matches like #1234
    hash_matches = re.findall(r'#(\d+)', title + "\n" + body)
    # Extract URL matches like /issues/1234 or /pull/1234 or /pulls/1234
    url_matches = re.findall(r'/(?:issues|pull|pulls)/(\d+)', title + "\n" + body)
    return {int(num) for num in hash_matches + url_matches}

def main():
    prs = fetch_open_prs()
    
    # Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}
    
    # Track which PRs are assigned to who for log outputs
    current_assignments = {member: [] for member in TEAM}
    
    for pr in prs:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "")
        requested_reviewers = pr.get("requested_reviewers", [])
        
        # Check if any team member is a requested reviewer
        assigned_team_reviewers = []
        for r in requested_reviewers:
            login = r.get("login")
            if login in TEAM:
                assigned_team_reviewers.append(login)
                
        if assigned_team_reviewers:
            # Extract all referenced issue IDs from title and body
            issue_ids = extract_issue_ids(title, body)
            if number:
                issue_ids.add(number)
                
            for reviewer in assigned_team_reviewers:
                workload[reviewer] += 1
                current_assignments[reviewer].append(number)
                for issue_id in issue_ids:
                    tracking_issue_to_reviewer[issue_id] = reviewer

    print("\nCurrent workload audit:")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} open assigned reviews (PRs: {sorted(current_assignments[member])})")
        
    # Step 2: Fetch Unassigned Candidate PRs (labeled "ready-for-human" and not reviewed by team)
    candidates = []
    for pr in prs:
        # Check for label "ready-for-human"
        labels = pr.get("labels", [])
        has_ready_label = any(l.get("name") == "ready-for-human" for l in labels)
        if not has_ready_label:
            continue
            
        # Check if already reviewed by any team member
        requested_reviewers = pr.get("requested_reviewers", [])
        has_team_reviewer = any(r.get("login") in TEAM for r in requested_reviewers)
        if has_team_reviewer:
            continue
            
        candidates.append(pr)
        
    # Sort candidates by PR number ascending (oldest first)
    candidates.sort(key=lambda pr: pr.get("number", 0))
    
    print(f"\nFound {len(candidates)} unassigned candidate 'ready-for-human' PRs:")
    for c in candidates:
        print(f"  #{c.get('number')}: {c.get('title')}")
        
    # Step 3: Assignment Algorithm
    assignments_to_make = []
    
    for pr in candidates:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "")
        
        # Check if all team members are at the hard ceiling of 10
        if all(workload[m] >= 10 for m in TEAM):
            print(f"Hard ceiling (10) reached for all team members. Skipping assignment for PR #{number}.")
            continue
            
        # 1. Extract Workflow Tracking Issues
        extracted_issue_ids = extract_issue_ids(title, body)
        if number:
            extracted_issue_ids.add(number)
            
        # 2. Select Reviewer by Priority
        selected_reviewer = None
        priority_used = ""
        
        # Priority 1 (Workflow Affinity): Matches existing issue being reviewed by a member who is under hard ceiling (< 10)
        affinity_candidates = []
        for issue_id in extracted_issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                reviewer = tracking_issue_to_reviewer[issue_id]
                if workload[reviewer] < 10:
                    affinity_candidates.append(reviewer)
                    
        if affinity_candidates:
            # Select the reviewer with the lowest workload, then tie-break alphabetically
            selected_reviewer = min(affinity_candidates, key=lambda r: (workload[r], r))
            priority_used = "Priority 1 (Workflow Affinity)"
            
        # Priority 2 (Underloaded Balancing): If any team members have workload < 5, select lowest workload, tie-break alphabetically
        if not selected_reviewer:
            underloaded = [r for r in TEAM if workload[r] < 5]
            if underloaded:
                selected_reviewer = min(underloaded, key=lambda r: (workload[r], r))
                priority_used = "Priority 2 (Underloaded Balancing)"
                
        # Priority 3 (Capacity Absorption): Select member with lowest workload < 10, tie-break alphabetically
        if not selected_reviewer:
            capable = [r for r in TEAM if workload[r] < 10]
            if capable:
                selected_reviewer = min(capable, key=lambda r: (workload[r], r))
                priority_used = "Priority 3 (Capacity Absorption)"
                
        if selected_reviewer:
            # 3. Update State & Queue Assignment
            workload[selected_reviewer] += 1
            for issue_id in extracted_issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_reviewer
                
            assignments_to_make.append((number, selected_reviewer, priority_used, title))
            
    # Step 4: Execute Assignments & Exit
    if not assignments_to_make:
        print("\nNo new assignments to make.")
        return
        
    print("\nQueued assignments to execute:")
    for num, reviewer, priority, title in assignments_to_make:
        print(f"  Assign PR #{num} to {reviewer} via {priority} - '{title}'")
        
    print("\nExecuting assignments on GitHub...")
    for num, reviewer, _, _ in assignments_to_make:
        edit_cmd = [
            "gh", "pr", "edit", str(num),
            "--repo", "GoogleCloudPlatform/k8s-config-connector",
            "--add-reviewer", reviewer
        ]
        print(f"Running: {' '.join(edit_cmd)}")
        run_command(edit_cmd)
        
    print("\nAll assignments successfully executed!")
    print("\nPost-assignment workload summary:")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} open assigned reviews")

if __name__ == "__main__":
    main()
