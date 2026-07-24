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
TEAM = {
    "acpana",
    "anfernee",
    "anhdle-sso",
    "barney-s",
    "gemmahou",
    "maqiuyujoyce"
}

def run_command(cmd):
    """Runs a shell command and returns stdout."""
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    if result.returncode != 0:
        print(f"Error running command {' '.join(cmd)}: {result.stderr}", file=sys.stderr)
        return ""
    return result.stdout

def fetch_open_prs():
    """Fetches all open PRs using paginated GitHub REST API calls."""
    prs = []
    page = 1
    while True:
        # Fetch open PRs in pages of 100
        cmd = [
            "gh", "api",
            f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"
        ]
        output = run_command(cmd)
        if not output:
            break
        try:
            data = json.loads(output)
        except json.JSONDecodeError as e:
            print(f"JSON decode error on page {page}: {e}", file=sys.stderr)
            break
        if not data:
            break
        prs.extend(data)
        page += 1
    return prs

def extract_issue_ids(title, body):
    """Extracts all referenced issue IDs (numbers) from the title and body."""
    content = f"{title or ''}\n{body or ''}"
    # Match #<NUMBER>
    matches = re.findall(r'#(\d+)', content)
    # Return sorted unique issue numbers as strings
    return sorted(list(set(matches)))

def balance_workloads(all_prs, team):
    """Calculates existing workloads and assigns eligible unassigned 'ready-for-human' PRs.
    
    Returns:
        (assignments, workload, tracking_issue_to_reviewer)
        where:
        - assignments is a list of tuples: (pr_number, selected_reviewer, reason)
        - workload is the updated workload dictionary
        - tracking_issue_to_reviewer is the updated issue mapping dictionary
    """
    workload = {member: 0 for member in team}
    tracking_issue_to_reviewer = {}

    for pr in all_prs:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "") or ""
        
        # Extract requested reviewers
        requested_reviewers = [r.get("login") for r in pr.get("requested_reviewers", []) if r]
        
        # Intersect with team
        team_reviewers = [r for r in requested_reviewers if r in team]
        
        # Each team reviewer on an open PR counts towards their workload
        for reviewer in team_reviewers:
            workload[reviewer] += 1
            
        # If there are team reviewers assigned, map the referenced issues to them
        if team_reviewers:
            issues = extract_issue_ids(title, body)
            for issue in issues:
                for reviewer in team_reviewers:
                    tracking_issue_to_reviewer[issue] = reviewer

    # Identify unassigned candidate PRs: must have "ready-for-human" label and NO team member requested.
    candidates = []
    for pr in all_prs:
        labels = [l.get("name") for l in pr.get("labels", []) if l]
        if "ready-for-human" not in labels:
            continue
            
        requested_reviewers = [r.get("login") for r in pr.get("requested_reviewers", []) if r]
        team_reviewers = [r for r in requested_reviewers if r in team]
        
        # Filter to those that do NOT currently have any member of k8s-config-connector-team listed in reviewRequests
        if not team_reviewers:
            candidates.append(pr)

    # Sort candidates by PR number (ascending) to be deterministic and process oldest first.
    candidates.sort(key=lambda x: x.get("number", 0))

    assignments = []
    for pr in candidates:
        pr_number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "") or ""
        
        # Extract all issue IDs from the candidate
        issues = extract_issue_ids(title, body)
        
        # Identify matched reviewers based on workflow affinity
        matched_reviewers = []
        for issue in issues:
            if issue in tracking_issue_to_reviewer:
                reviewer = tracking_issue_to_reviewer[issue]
                if reviewer not in matched_reviewers:
                    matched_reviewers.append(reviewer)
                    
        # Filter matched reviewers to those who are under the hard limit of 10
        valid_affinity_reviewers = [r for r in matched_reviewers if workload[r] < 10]
        
        selected_reviewer = None
        reason = ""
        
        # Priority 1: Workflow Affinity
        if valid_affinity_reviewers:
            # Select the reviewer with the lowest workload. In case of ties, select alphabetically.
            valid_affinity_reviewers.sort(key=lambda x: (workload[x], x))
            selected_reviewer = valid_affinity_reviewers[0]
            reason = f"Workflow Affinity (matched issues: {[i for i in issues if tracking_issue_to_reviewer.get(i) == selected_reviewer]})"
        else:
            # Priority 2: Underloaded Balancing (C_user < 5)
            underloaded = [m for m in team if workload[m] < 5]
            if underloaded:
                underloaded.sort(key=lambda x: (workload[x], x))
                selected_reviewer = underloaded[0]
                reason = f"Underloaded Balancing (current count: {workload[selected_reviewer]} < 5)"
            else:
                # Priority 3: Capacity Absorption (5 <= C_user < 10)
                absorbers = [m for m in team if 5 <= workload[m] < 10]
                if absorbers:
                    absorbers.sort(key=lambda x: (workload[x], x))
                    selected_reviewer = absorbers[0]
                    reason = f"Capacity Absorption (current count: {workload[selected_reviewer]} >= 5 and < 10)"
                else:
                    reason = "All team members reached the hard ceiling of 10"
                    
        if selected_reviewer:
            # Update State
            workload[selected_reviewer] += 1
            for issue in issues:
                tracking_issue_to_reviewer[issue] = selected_reviewer
            assignments.append((pr_number, selected_reviewer, reason))
            
    return assignments, workload, tracking_issue_to_reviewer

def main():
    print("Fetching open pull requests...")
    all_prs = fetch_open_prs()
    print(f"Found {len(all_prs)} open pull requests.")

    # Calculate assignments
    assignments, workload, tracking_issue_to_reviewer = balance_workloads(all_prs, TEAM)

    print("\nCurrent workload:")
    for member, count in sorted(workload.items()):
        print(f"  {member}: {count}")

    print(f"\nTracking issue mappings found: {len(tracking_issue_to_reviewer)}")
    print(f"\nQueueing {len(assignments)} assignments...")

    # Execute Assignments
    if not assignments:
        print("\nNo assignments to make.")
        return

    print("\nExecuting assignments...")
    for pr_number, reviewer, reason in assignments:
        print(f"Assigning PR #{pr_number} to {reviewer} (Reason: {reason})...")
        cmd = [
            "gh", "pr", "edit", str(pr_number),
            "--repo", "GoogleCloudPlatform/k8s-config-connector",
            "--add-reviewer", reviewer
        ]
        result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        if result.returncode != 0:
            print(f"Failed to assign PR #{pr_number} to {reviewer}: {result.stderr}", file=sys.stderr)
        else:
            print(f"Successfully assigned PR #{pr_number} to {reviewer}.")

    print("\nSummary of Assignments Made:")
    for pr_number, reviewer, reason in assignments:
        print(f"  PR #{pr_number} -> {reviewer} ({reason})")

if __name__ == "__main__":
    main()
