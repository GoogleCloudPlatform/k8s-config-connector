#!/usr/bin/env python3
import json
import subprocess
import re
import sys

# Define team members
TEAM = {"acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"}

# Issue regex
ISSUE_RE = re.compile(r'#(\d+)')

def run_command(cmd):
    result = subprocess.run(cmd, capture_output=True, text=True)
    if result.returncode != 0:
        print(f"Command failed: {' '.join(cmd)}\nError: {result.stderr}", file=sys.stderr)
        sys.exit(1)
    return result.stdout

def get_open_prs():
    print("Fetching open pull requests...", flush=True)
    prs = []
    page = 1
    while True:
        url = f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"
        stdout = run_command(["gh", "api", url])
        page_prs = json.loads(stdout)
        if not page_prs:
            break
        prs.extend(page_prs)
        print(f"  Fetched page {page} ({len(page_prs)} PRs, total so far: {len(prs)})", flush=True)
        if len(page_prs) < 100:
            break
        page += 1
    return prs

def main():
    prs = get_open_prs()
    print(f"Total open PRs: {len(prs)}", flush=True)

    # Step 1: Initialize states
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {} # Maps issue_id -> set of team member usernames

    candidates = []

    # Parse all open PRs
    for pr in prs:
        number = pr['number']
        title = pr.get('title') or ""
        body = pr.get('body') or ""

        # Extract issue references
        issues = set(ISSUE_RE.findall(title) + ISSUE_RE.findall(body))

        # Check reviewers
        requested_reviewers = pr.get('requested_reviewers', [])
        assigned_team_reviewers = [r['login'] for r in requested_reviewers if r.get('login') in TEAM]

        # Update workload
        for r in assigned_team_reviewers:
            workload[r] += 1
            # Update tracking mapping
            for issue_id in issues:
                if issue_id not in tracking_issue_to_reviewer:
                    tracking_issue_to_reviewer[issue_id] = set()
                tracking_issue_to_reviewer[issue_id].add(r)

        # Determine if it's a candidate for assignment
        # STRICT Eligibility: Open, Labeled 'ready-for-human', and NOT assigned to any TEAM member
        labels = [l['name'] for l in pr.get('labels', [])]
        if 'ready-for-human' in labels:
            if not assigned_team_reviewers:
                candidates.append((number, title, body, issues))

    # Print current workloads
    print("\n--- Current Team Workloads ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned open reviews")

    print(f"\nFound {len(candidates)} unassigned 'ready-for-human' PR(s).")
    for number, title, _, _ in candidates:
        print(f"  Candidate: #{number} - {title}")

    # Step 3: Assignment Algorithm
    assignments = []
    for number, title, body, pr_issues in candidates:
        # Check if all team members reached 10
        if all(workload[m] >= 10 for m in TEAM):
            print("All team members are at or above the hard ceiling of 10. Skipping remaining assignments.", flush=True)
            break

        selected_reviewer = None

        # Priority 1: Workflow Affinity
        # Find any team member already reviewing a tracking issue/sub-issue of this PR
        affinity_reviewers = set()
        for issue_id in pr_issues:
            if issue_id in tracking_issue_to_reviewer:
                affinity_reviewers.update(tracking_issue_to_reviewer[issue_id])

        # Filter to those who have workload < 10
        eligible_affinity = [r for r in affinity_reviewers if workload[r] < 10]
        if eligible_affinity:
            # Select the one with the lowest workload
            selected_reviewer = min(eligible_affinity, key=lambda r: workload[r])
            reason = f"Workflow Affinity (lowest workload among: {eligible_affinity})"
        else:
            # Priority 2: Underloaded Balancing (< 5)
            underloaded = [r for r in TEAM if workload[r] < 5]
            if underloaded:
                selected_reviewer = min(underloaded, key=lambda r: workload[r])
                reason = "Underloaded Balancing (< 5)"
            else:
                # Priority 3: Capacity Absorption (5 <= workload < 10)
                available = [r for r in TEAM if workload[r] < 10]
                if available:
                    selected_reviewer = min(available, key=lambda r: workload[r])
                    reason = "Capacity Absorption"

        if selected_reviewer:
            # Record assignment
            assignments.append((number, selected_reviewer, reason))
            # Update in-memory state
            workload[selected_reviewer] += 1
            for issue_id in pr_issues:
                if issue_id not in tracking_issue_to_reviewer:
                    tracking_issue_to_reviewer[issue_id] = set()
                tracking_issue_to_reviewer[issue_id].add(selected_reviewer)
        else:
            print(f"Could not assign PR #{number} - No eligible reviewer under the ceiling.", flush=True)

    # Step 4: Execute Assignments
    if not assignments:
        print("\nNo new assignments to make.")
        return

    print("\n--- Proposed Assignments ---")
    for number, user, reason in assignments:
        print(f"  PR #{number} -> {user} ({reason})")

    print("\nExecuting assignments...", flush=True)
    for number, user, reason in assignments:
        cmd = ["gh", "pr", "edit", str(number), "--repo", "GoogleCloudPlatform/k8s-config-connector", "--add-reviewer", user]
        print(f"  Assigning PR #{number} to {user}...", flush=True)
        run_command(cmd)

    print("\nSuccessfully updated review assignments!", flush=True)

    # Print ending workloads
    print("\n--- Final Team Workloads ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned open reviews")

if __name__ == "__main__":
    main()
