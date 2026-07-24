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

import argparse
import json
import re
import subprocess
import sys

# Team members definition
TEAM = ["acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"]

def fetch_all_open_prs():
    """Fetches all open PRs in GoogleCloudPlatform/k8s-config-connector."""
    prs = []
    page = 1
    while True:
        cmd = [
            "gh", "api",
            f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"
        ]
        res = subprocess.run(cmd, capture_output=True, text=True)
        if res.returncode != 0:
            print(f"Error fetching page {page}: {res.stderr}", file=sys.stderr)
            break
        try:
            page_prs = json.loads(res.stdout)
        except json.JSONDecodeError as e:
            print(f"Error parsing JSON on page {page}: {e}", file=sys.stderr)
            break
        if not page_prs:
            break
        prs.extend(page_prs)
        if len(page_prs) < 100:
            break
        page += 1
    return prs

def extract_issue_ids(title, body):
    """Extracts all referenced issue IDs matching #<NUMBER> from title and body."""
    text = f"{title or ''} {body or ''}"
    # Find all matches of # followed by digits
    matches = re.findall(r'#(\d+)', text)
    # Return unique issue IDs as strings
    return sorted(list(set(matches)))

def main():
    parser = argparse.ArgumentParser(description="PR Review Workload Balancer")
    parser.add_argument("--dry-run", action="store_true", help="Print assignment plan without modifying anything")
    args = parser.parse_args()

    print("Fetching open pull requests...")
    all_prs = fetch_all_open_prs()
    print(f"Successfully fetched {len(all_prs)} open PRs.")

    # Step 1: Audit Current Workloads & Workflow Mapping
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}

    for pr in all_prs:
        requested_reviewers = pr.get("requested_reviewers", [])
        assigned_team_reviewers = [
            r["login"] for r in requested_reviewers if r.get("login") in TEAM
        ]
        if assigned_team_reviewers:
            # Increment workloads and map referenced issues
            for rev in assigned_team_reviewers:
                workload[rev] += 1
                issue_ids = extract_issue_ids(pr.get("title"), pr.get("body"))
                for issue_id in issue_ids:
                    tracking_issue_to_reviewer[issue_id] = rev

    print("\nInitial Workloads:")
    for member, count in workload.items():
        print(f"  {member}: {count} assigned reviews")

    # Step 2: Fetch / Filter Unassigned Candidate PRs Labeled 'ready-for-human'
    candidates = []
    for pr in all_prs:
        # Check if labeled 'ready-for-human'
        labels = [l.get("name") for l in pr.get("labels", []) if l.get("name")]
        if "ready-for-human" not in labels:
            continue

        # Filter out those that already have a team member requested
        requested_reviewers = pr.get("requested_reviewers", [])
        has_team_reviewer = any(r.get("login") in TEAM for r in requested_reviewers)
        if has_team_reviewer:
            continue

        candidates.append(pr)

    print(f"\nFound {len(candidates)} unassigned 'ready-for-human' PR candidates.")

    # Step 3: Assignment Algorithm
    assignments = []  # List of tuples: (pr_number, pr_title, reviewer, reason)

    # Sort candidates by PR number to have a deterministic order (e.g., oldest first)
    candidates.sort(key=lambda pr: pr.get("number", 0))

    for pr in candidates:
        pr_number = pr.get("number")
        title = pr.get("title")
        body = pr.get("body")

        # Extract issue IDs
        extracted_issues = extract_issue_ids(title, body)

        # Check if all team members reached the hard limit of 10
        if all(workload[u] >= 10 for u in TEAM):
            print("All team members have reached the hard ceiling of 10 reviews. Stopping further assignments.")
            break

        selected_reviewer = None
        reason = ""

        # Priority 1: Workflow Affinity
        for issue_id in extracted_issues:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    selected_reviewer = rev
                    reason = f"Workflow Affinity (Issue #{issue_id})"
                    break

        # Priority 2: Underloaded Balancing (C_user < 5)
        if not selected_reviewer:
            underloaded = [u for u in TEAM if workload[u] < 5]
            if underloaded:
                selected_reviewer = min(underloaded, key=lambda u: workload[u])
                reason = "Underloaded Balancing (< 5 reviews)"

        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        if not selected_reviewer:
            absorbers = [u for u in TEAM if workload[u] < 10]
            if absorbers:
                selected_reviewer = min(absorbers, key=lambda u: workload[u])
                reason = "Capacity Absorption (workload balancer)"

        if selected_reviewer:
            # Update state
            workload[selected_reviewer] += 1
            for issue_id in extracted_issues:
                tracking_issue_to_reviewer[issue_id] = selected_reviewer
            
            assignments.append((pr_number, title, selected_reviewer, reason))
        else:
            print(f"Could not assign PR #{pr_number}: '{title}' (all potential eligible reviewers are at capacity).")

    print("\nProposed Assignments:")
    if not assignments:
        print("  No assignments to be made.")
    for pr_number, title, reviewer, reason in assignments:
        print(f"  PR #{pr_number}: Assigning to {reviewer} ({reason})")

    print("\nFinal Projected Workloads:")
    for member, count in workload.items():
        print(f"  {member}: {count} assigned reviews")

    # Step 4: Execute Assignments
    if args.dry_run:
        print("\n[Dry Run] Skipping executing assignments.")
    else:
        print("\nExecuting assignments...")
        for pr_number, title, reviewer, reason in assignments:
            print(f"Assigning PR #{pr_number} to {reviewer}...")
            cmd = [
                "gh", "pr", "edit", str(pr_number),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", reviewer
            ]
            res = subprocess.run(cmd, capture_output=True, text=True)
            if res.returncode == 0:
                print(f"Successfully assigned PR #{pr_number} to {reviewer}.")
            else:
                print(f"Failed to assign PR #{pr_number} to {reviewer}: {res.stderr}", file=sys.stderr)

if __name__ == "__main__":
    main()
