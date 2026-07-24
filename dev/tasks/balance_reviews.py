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
import argparse

# The k8s-config-connector-team composition.
# Modify this single list to update the composition as membership changes.
TEAM = [
    'acpana',
    'anfernee',
    'anhdle-sso',
    'barney-s',
    'gemmahou',
    'maqiuyujoyce'
]

def fetch_open_prs():
    """Fetches all open PRs with required fields using paginated REST API."""
    cmd = [
        "gh", "api",
        "repos/GoogleCloudPlatform/k8s-config-connector/pulls?per_page=100",
        "--paginate",
        "--jq",
        '.[] | {number: .number, title: .title, body: .body, labels: [.labels[].name], requested_reviewers: [.requested_reviewers[].login]}'
    ]
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, check=True)
    
    prs = []
    for line in result.stdout.strip().split('\n'):
        if line.strip():
            try:
                prs.append(json.loads(line))
            except json.JSONDecodeError:
                pass
    return prs

def audit_workloads(prs, team):
    """
    Builds data structures from the current open PRs:
    1. workload: Maps team member to open review request count.
    2. tracking_issue_to_reviewer: Maps issue IDs to the team member assigned.
    """
    workload = {member: 0 for member in team}
    tracking_issue_to_reviewer = {}
    issue_pattern = re.compile(r'#(\d+)')

    for pr in prs:
        reviewers = pr.get('requested_reviewers', [])
        team_reviewers = [r for r in reviewers if r in team]
        
        if team_reviewers:
            for r in team_reviewers:
                workload[r] += 1
            
            title = pr.get('title', '') or ''
            body = pr.get('body', '') or ''
            issue_ids = issue_pattern.findall(f"{title} {body}")
            for issue_id in issue_ids:
                # Map to the first team reviewer for consistency
                tracking_issue_to_reviewer[issue_id] = team_reviewers[0]

    return workload, tracking_issue_to_reviewer

def get_candidate_prs(prs, team):
    """Filters PRs that are 'ready-for-human' and have no team reviewer requested."""
    candidates = []
    for pr in prs:
        labels = pr.get('labels', [])
        if 'ready-for-human' in labels:
            reviewers = pr.get('requested_reviewers', [])
            has_team_reviewer = any(r in team for r in reviewers)
            if not has_team_reviewer:
                candidates.append(pr)
    
    # Sort candidates by PR number ascending to process oldest PRs first
    candidates.sort(key=lambda x: x['number'])
    return candidates

def assign_reviews(candidates, workload, tracking_issue_to_reviewer, team):
    """
    Assigns candidate PRs to team members using priority-based rules:
    - Priority 1: Workflow affinity (extracted issue ID matches existing reviewer, count < 10)
    - Priority 2: Underloaded balancing (count < 5, select lowest count)
    - Priority 3: Capacity absorption (5 <= count < 10, select lowest count)
    """
    assignments = []
    issue_pattern = re.compile(r'#(\d+)')

    for pr in candidates:
        pr_number = pr['number']
        title = pr.get('title', '') or ''
        body = pr.get('body', '') or ''
        
        issue_ids = issue_pattern.findall(f"{title} {body}")
        unique_issue_ids = list(dict.fromkeys(issue_ids))
        
        selected_user = None
        priority_reason = ""

        # Priority 1: Workflow Affinity
        potential_reviewers = []
        for issue_id in unique_issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    potential_reviewers.append(rev)
        
        if potential_reviewers:
            # Deterministic lowest count, alphabetical tie-break
            selected_user = min(set(potential_reviewers), key=lambda r: (workload[r], r))
            priority_reason = "Priority 1 (Workflow Affinity)"

        # Priority 2: Underloaded Balancing
        if not selected_user:
            underloaded = [m for m in team if workload[m] < 5]
            if underloaded:
                selected_user = min(underloaded, key=lambda m: (workload[m], m))
                priority_reason = "Priority 2 (Underloaded Balancing)"

        # Priority 3: Capacity Absorption
        if not selected_user:
            capacity_members = [m for m in team if workload[m] < 10]
            if capacity_members:
                selected_user = min(capacity_members, key=lambda m: (workload[m], m))
                priority_reason = "Priority 3 (Capacity Absorption)"

        if selected_user:
            assignments.append({
                'pr_number': pr_number,
                'title': title,
                'reviewer': selected_user,
                'reason': priority_reason,
                'issue_ids': unique_issue_ids
            })
            workload[selected_user] += 1
            for issue_id in unique_issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_user
        else:
            # All team members have workload = 10, termination condition met
            break

    return assignments

def execute_assignments(assignments, dry_run=False):
    """Executes gh pr edit commands to assign reviewers to PRs."""
    if not assignments:
        print("No new assignments to execute.")
        return

    print(f"\nExecuting {len(assignments)} assignments:")
    for assoc in assignments:
        pr_number = assoc['pr_number']
        reviewer = assoc['reviewer']
        reason = assoc['reason']
        title = assoc['title']
        
        print(f"Assigning PR #{pr_number} ({title[:50]}...) to {reviewer} via {reason}")
        if not dry_run:
            cmd = [
                "gh", "pr", "edit", str(pr_number),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", reviewer
            ]
            result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            if result.returncode != 0:
                print(f"Error assigning PR #{pr_number}: {result.stderr.strip()}", file=sys.stderr)
            else:
                print(f"Successfully assigned PR #{pr_number} to {reviewer}")

def main():
    parser = argparse.ArgumentParser(description="PR Review Workload Balancer")
    parser.add_argument("--dry-run", action="store_true", help="Perform a dry-run and do not make changes")
    args = parser.parse_args()

    print("Auditing current team workloads...")
    try:
        prs = fetch_open_prs()
    except subprocess.CalledProcessError as e:
        print(f"Error fetching open PRs: {e.stderr}", file=sys.stderr)
        sys.exit(1)

    workload, tracking_issue_to_reviewer = audit_workloads(prs, TEAM)
    
    print("\nInitial Workloads:")
    for member in TEAM:
        print(f"  {member}: {workload[member]} assigned open reviews")

    candidates = get_candidate_prs(prs, TEAM)
    print(f"\nFound {len(candidates)} unassigned candidate 'ready-for-human' PRs.")

    assignments = assign_reviews(candidates, workload, tracking_issue_to_reviewer, TEAM)
    
    execute_assignments(assignments, dry_run=args.dry_run)

    print("\nUpdated Workloads:")
    for member in TEAM:
        print(f"  {member}: {workload[member]} assigned open reviews")

if __name__ == "__main__":
    main()
