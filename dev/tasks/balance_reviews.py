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

TEAM = ["acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"]

def fetch_open_prs(repo="GoogleCloudPlatform/k8s-config-connector"):
    """Fetch all open PRs from GitHub REST API (handles pagination up to 10 pages)."""
    prs = []
    for page in range(1, 11):
        cmd = ["gh", "api", f"/repos/{repo}/pulls?state=open&per_page=100&page={page}"]
        res = subprocess.run(cmd, capture_output=True, text=True)
        if res.returncode != 0:
            print(f"Error fetching open PRs (page {page}): {res.stderr}", file=sys.stderr)
            break
        page_prs = json.loads(res.stdout)
        if not page_prs:
            break
        prs.extend(page_prs)
        if len(page_prs) < 100:
            break
    return prs

def extract_issue_ids(title, body):
    """Extract all referenced issue/PR ID integers from title and body text."""
    if not title:
        title = ""
    if not body:
        body = ""
    text = f"{title}\n{body}"
    # Extract numbers following '#', e.g. #10976, Fixes #10276, etc.
    matches = re.findall(r"#(\d+)", text)
    return sorted(list(set(int(num) for num in matches)))

def audit_workloads(prs, team):
    """Audit current team workloads and build multi-level workflow maps.
    
    Returns:
        workloads: dict of team member -> open review count.
        tracking_issue_to_reviewer: dict of issue_id -> set of team member usernames currently reviewing it.
        unassigned_candidates: list of PRs that are candidate ready-for-human PRs.
    """
    workloads = {member: 0 for member in team}
    tracking_issue_to_reviewer = {}
    unassigned_candidates = []

    for pr in prs:
        labels = [l["name"] for l in pr.get("labels", [])]
        reviewers = [r["login"] for r in pr.get("requested_reviewers", [])]
        assigned_team_reviewers = [r for r in reviewers if r in team]

        title = pr.get("title") or ""
        body = pr.get("body") or ""
        issue_ids = extract_issue_ids(title, body)

        if assigned_team_reviewers:
            for reviewer in assigned_team_reviewers:
                workloads[reviewer] += 1
            for issue_id in issue_ids:
                if issue_id not in tracking_issue_to_reviewer:
                    tracking_issue_to_reviewer[issue_id] = set()
                for reviewer in assigned_team_reviewers:
                    tracking_issue_to_reviewer[issue_id].add(reviewer)
        else:
            if "ready-for-human" in labels:
                unassigned_candidates.append(pr)

    return workloads, tracking_issue_to_reviewer, unassigned_candidates

def balance_reviews(unassigned_candidates, team, workloads, tracking_issue_to_reviewer):
    """Perform the review assignment algorithm on unassigned ready-for-human candidates."""
    assignments = []
    
    # Sort candidates by number to ensure deterministic ordering
    sorted_candidates = sorted(unassigned_candidates, key=lambda pr: pr.get("number", 0))

    for pr in sorted_candidates:
        pr_number = pr.get("number")
        title = pr.get("title") or ""
        body = pr.get("body") or ""
        issue_ids = extract_issue_ids(title, body)

        # Check termination condition: If all team members are at or above the ceiling of 10, stop.
        if all(workloads[m] >= 10 for m in team):
            print("Hard ceiling of 10 reviews reached for all team members. Skipping further assignments.")
            break

        selected_user = None

        # Priority 1: Workflow Affinity
        # If any extracted issue ID matches a reviewer who has < 10 workload, prefer them.
        affinity_candidates = set()
        for issue_id in issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                for r in tracking_issue_to_reviewer[issue_id]:
                    if workloads[r] < 10:
                        affinity_candidates.add(r)
        
        if affinity_candidates:
            # Prefer the one with the lowest current workload among affinity candidates
            selected_user = min(affinity_candidates, key=lambda u: (workloads[u], u))
            reason = f"Workflow Affinity (matches issue(s) {issue_ids})"
        else:
            # Priority 2: Underloaded Balancing (workload < 5)
            underloaded = [u for u in team if workloads[u] < 5]
            if underloaded:
                selected_user = min(underloaded, key=lambda u: (workloads[u], u))
                reason = "Underloaded Balancing (< 5 reviews)"
            else:
                # Priority 3: Capacity Absorption (5 <= workload < 10)
                absorbers = [u for u in team if workloads[u] < 10]
                if absorbers:
                    selected_user = min(absorbers, key=lambda u: (workloads[u], u))
                    reason = "Capacity Absorption"

        if selected_user:
            # Update state immediately
            workloads[selected_user] += 1
            for issue_id in issue_ids:
                if issue_id not in tracking_issue_to_reviewer:
                    tracking_issue_to_reviewer[issue_id] = set()
                tracking_issue_to_reviewer[issue_id].add(selected_user)
            
            assignments.append({
                "pr_number": pr_number,
                "title": title,
                "selected_user": selected_user,
                "reason": reason
            })

    return assignments

def execute_assignments(assignments, repo="GoogleCloudPlatform/k8s-config-connector"):
    """Execute gh pr edit commands to assign reviewers."""
    executed = []
    for assign in assignments:
        pr_number = assign["pr_number"]
        user = assign["selected_user"]
        print(f"Assigning review for PR #{pr_number} to {user} (Reason: {assign['reason']})")
        cmd = ["gh", "pr", "edit", str(pr_number), "--repo", repo, "--add-reviewer", user]
        res = subprocess.run(cmd, capture_output=True, text=True)
        if res.returncode == 0:
            executed.append(assign)
        else:
            print(f"Failed to assign PR #{pr_number} to {user}: {res.stderr}", file=sys.stderr)
    return executed

def main():
    print("=== PR Review Workload Balancer ===")
    prs = fetch_open_prs()
    print(f"Successfully fetched {len(prs)} open PRs.")

    workloads, tracking_maps, candidates = audit_workloads(prs, TEAM)
    
    print("\n--- Current Workloads ---")
    for member in TEAM:
        print(f"  {member}: {workloads[member]} assigned reviews")

    print(f"\nFound {len(candidates)} unassigned ready-for-human PRs.")

    assignments = balance_reviews(candidates, TEAM, workloads, tracking_maps)

    if not assignments:
        print("\nNo review assignments needed.")
        sys.exit(0)

    print(f"\nProposed {len(assignments)} assignments:")
    for a in assignments:
        print(f"  PR #{a['pr_number']} -> {a['selected_user']} ({a['reason']})")

    print("\nExecuting assignments...")
    executed = execute_assignments(assignments)
    print(f"\nSuccessfully executed {len(executed)} of {len(assignments)} assignments.")

if __name__ == "__main__":
    main()
