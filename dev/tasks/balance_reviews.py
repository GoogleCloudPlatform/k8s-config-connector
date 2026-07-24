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

# Define the team members
TEAM_MEMBERS = [
    "acpana",
    "anfernee",
    "anhdle-sso",
    "barney-s",
    "gemmahou",
    "maqiuyujoyce",
]


def extract_issue_ids(title, body):
    """Extracts all issue numbers referenced in the title or body."""
    ids = set()
    pattern = r"#(\d+)"
    if title:
        ids.update(re.findall(pattern, title))
    if body:
        ids.update(re.findall(pattern, body))
    return sorted(list(ids))


def fetch_open_prs():
    """Fetches all open pull requests from GoogleCloudPlatform/k8s-config-connector."""
    cmd = [
        "gh",
        "api",
        "--paginate",
        "repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100",
    ]
    result = subprocess.run(cmd, capture_output=True, text=True, check=True)
    raw = result.stdout.strip()

    # gh api --paginate outputs sequential JSON arrays. We parse them iteratively.
    decoder = json.JSONDecoder()
    pos = 0
    prs = []
    while pos < len(raw):
        # Skip leading whitespace
        while pos < len(raw) and raw[pos].isspace():
            pos += 1
        if pos >= len(raw):
            break
        obj, idx = decoder.raw_decode(raw[pos:])
        if isinstance(obj, list):
            prs.extend(obj)
        pos += idx
    return prs


def audit_workloads_and_workflows(prs):
    """Step 1: Audits team workloads and creates tracking issue to reviewer map."""
    workload = {member: 0 for member in TEAM_MEMBERS}
    tracking_issue_to_reviewer = {}

    for pr in prs:
        # Determine requested reviewers from TEAM_MEMBERS
        requested = [
            r["login"] for r in pr.get("requested_reviewers", []) if r["login"] in TEAM_MEMBERS
        ]
        
        # If any team member is a requested reviewer, increment their workload
        for rev in requested:
            workload[rev] += 1
            
            # Map all extracted issue numbers to this reviewer
            issue_ids = extract_issue_ids(pr.get("title", ""), pr.get("body", ""))
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = rev

    return workload, tracking_issue_to_reviewer


def filter_candidates(prs):
    """Step 2: Filters open candidate PRs with 'ready-for-human' label and no team requested_reviewers."""
    candidates = []
    for pr in prs:
        # Ignore draft PRs
        if pr.get("draft") is True:
            continue

        labels = [l["name"] for l in pr.get("labels", [])]
        if "ready-for-human" not in labels:
            continue

        # Check if any team member is currently listed as requested_reviewer
        requested_team_members = [
            r["login"] for r in pr.get("requested_reviewers", []) if r["login"] in TEAM_MEMBERS
        ]
        if not requested_team_members:
            candidates.append(pr)

    return candidates


def run_assignment(candidates, workload, tracking_issue_to_reviewer):
    """Step 3: Run priority-based assignment algorithm on candidates."""
    assignments = []  # List of tuples: (pr_number, selected_user)

    for pr in candidates:
        pr_number = pr["number"]
        title = pr.get("title", "")
        body = pr.get("body", "")

        # 1. Extract tracking issues
        issue_ids = extract_issue_ids(title, body)

        # 2. Select reviewer by priority
        selected_user = None

        # Priority 1: Workflow Affinity
        affinity_reviewers = []
        for issue_id in issue_ids:
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    affinity_reviewers.append(rev)

        if affinity_reviewers:
            # Sort by workload (lowest first), then alphabetically for determinism
            affinity_reviewers.sort(key=lambda r: (workload[r], r))
            selected_user = affinity_reviewers[0]
            print(f"PR #{pr_number}: Workflow affinity match with {selected_user} (based on referenced issues: {issue_ids})")
        else:
            # Priority 2: Underloaded Balancing (C_user < 5)
            underloaded = [m for m in TEAM_MEMBERS if workload[m] < 5]
            if underloaded:
                underloaded.sort(key=lambda r: (workload[r], r))
                selected_user = underloaded[0]
                print(f"PR #{pr_number}: Underloaded reviewer selected: {selected_user} (workload: {workload[selected_user]} < 5)")
            else:
                # Priority 3: Capacity Absorption (5 <= C_user < 10)
                absorb = [m for m in TEAM_MEMBERS if 5 <= workload[m] < 10]
                if absorb:
                    absorb.sort(key=lambda r: (workload[r], r))
                    selected_user = absorb[0]
                    print(f"PR #{pr_number}: Capacity absorption reviewer selected: {selected_user} (workload: {workload[selected_user]} < 10)")
                else:
                    # Ceiling reached (all team members have workload >= 10)
                    print(f"PR #{pr_number}: Skipped assignment. All team members have reached the maximum hard ceiling of 10.")

        if selected_user:
            # 3. Update state
            workload[selected_user] += 1
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_user

            assignments.append((pr_number, selected_user))

            # Stop assigning if everyone has hit the limit
            if all(workload[m] >= 10 for m in TEAM_MEMBERS):
                print("All team members have reached the maximum limit of 10 reviews. Halting assignments.")
                break

    return assignments


def execute_assignments(assignments, dry_run=False):
    """Step 4: Executes PR assignments using gh CLI."""
    if not assignments:
        print("No new review assignments queued.")
        return

    print(f"\n--- Executing {len(assignments)} Assignments ---")
    for pr_number, user in assignments:
        cmd = [
            "gh",
            "pr",
            "edit",
            str(pr_number),
            "--repo",
            "GoogleCloudPlatform/k8s-config-connector",
            "--add-reviewer",
            user,
        ]
        if dry_run:
            print(f"[DRY-RUN] Would run: {' '.join(cmd)}")
        else:
            print(f"Running: {' '.join(cmd)}")
            subprocess.run(cmd, check=True)
    print("Execution complete.")


def main():
    parser = argparse.ArgumentParser(description="PR Review Workload Balancer")
    parser.add_argument("--dry-run", action="store_true", help="Perform a dry-run without making assignments.")
    args = parser.parse_args()

    print("Auditing open PRs...")
    try:
        prs = fetch_open_prs()
    except Exception as e:
        print(f"Error fetching open PRs: {e}", file=sys.stderr)
        sys.exit(1)

    print(f"Total open PRs fetched: {len(prs)}")

    workload, tracking_issue_to_reviewer = audit_workloads_and_workflows(prs)

    print("\nCurrent Team Workloads:")
    for m in TEAM_MEMBERS:
        print(f" - {m}: {workload[m]}")

    candidates = filter_candidates(prs)
    print(f"\nFound {len(candidates)} unassigned 'ready-for-human' candidates:")
    for pr in candidates:
        print(f" - PR #{pr['number']}: {pr['title']}")

    assignments = run_assignment(candidates, workload, tracking_issue_to_reviewer)

    execute_assignments(assignments, dry_run=args.dry_run)


if __name__ == "__main__":
    main()
