#!/usr/bin/env python3
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

def get_open_prs():
    print("Fetching all open PRs from GoogleCloudPlatform/k8s-config-connector...")
    prs = []
    page = 1
    while True:
        print(f"Fetching page {page}...")
        cmd = [
            "gh", "api",
            f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"
        ]
        result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        if result.returncode != 0:
            print(f"Error fetching page {page}: {result.stderr}", file=sys.stderr)
            break
        page_prs = json.loads(result.stdout)
        if not page_prs:
            break
        prs.extend(page_prs)
        if len(page_prs) < 100:
            break
        page += 1
    print(f"Successfully fetched {len(prs)} open PRs.")
    return prs

def extract_issue_ids(title, body):
    text = f"{title or ''} {body or ''}"
    # Match #<NUMBER> (e.g. #10976, #10276)
    # Using raw string pattern
    ids = re.findall(r'#(\d+)', text)
    return set(ids)

def main():
    # Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping
    open_prs = get_open_prs()

    # Initialize workload tracking
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}

    print("\n--- Auditing current open PRs for team reviews ---")
    for pr in open_prs:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "")
        requested_reviewers = pr.get("requested_reviewers", [])

        assigned_team_reviewers = []
        for rr in requested_reviewers:
            login = rr.get("login")
            if login in TEAM:
                assigned_team_reviewers.append(login)

        if assigned_team_reviewers:
            issue_ids = extract_issue_ids(title, body)
            for r in assigned_team_reviewers:
                workload[r] += 1
                for issue_id in issue_ids:
                    # In case of multiple, any can map, but let's associate them
                    tracking_issue_to_reviewer[issue_id] = r
            print(f"PR #{number} ({title[:50]}...): assigned to {assigned_team_reviewers}, issues: {sorted(list(issue_ids))}")

    print("\n--- Initial Workload Summary ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned reviews")

    print(f"\nTracking Issue to Reviewer Map ({len(tracking_issue_to_reviewer)} mappings):")
    for issue_id in sorted(tracking_issue_to_reviewer.keys(), key=int):
        print(f"  #{issue_id} -> {tracking_issue_to_reviewer[issue_id]}")

    # Step 2: Fetch Unassigned Candidate PRs Labeled 'ready-for-human'
    print("\n--- Filtering Candidate ready-for-human PRs ---")
    candidates = []
    for pr in open_prs:
        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "")
        labels = [l.get("name") for l in pr.get("labels", []) if l.get("name")]
        
        if "ready-for-human" not in labels:
            continue

        requested_reviewers = [rr.get("login") for rr in pr.get("requested_reviewers", []) if rr.get("login")]
        has_team_reviewer = any(login in TEAM for login in requested_reviewers)
        if has_team_reviewer:
            # Already assigned to a team member
            continue

        candidates.append(pr)
        issue_ids = extract_issue_ids(title, body)
        print(f"Candidate PR #{number}: '{title}' | issues: {sorted(list(issue_ids))}")

    print(f"\nFound {len(candidates)} unassigned candidate 'ready-for-human' PRs.")

    # Step 3: Assignment Algorithm
    assignments = []  # List of tuples (pr_number, selected_user, reason, title)

    for pr in candidates:
        # Check termination condition: every team member has reached hard ceiling of 10
        if all(workload[m] >= 10 for m in TEAM):
            print("\nHard ceiling of 10 assigned reviews reached for all team members. Stopping assignments.")
            break

        number = pr.get("number")
        title = pr.get("title", "")
        body = pr.get("body", "")
        issue_ids = extract_issue_ids(title, body)

        selected_user = None
        reason = ""

        # Priority 1: Workflow Affinity
        # If ANY extracted issue ID matches a key in tracking_issue_to_reviewer, and that reviewer's count < 10, select that reviewer.
        affinity_candidates = []
        for issue_id in sorted(issue_ids):
            if issue_id in tracking_issue_to_reviewer:
                rev = tracking_issue_to_reviewer[issue_id]
                if workload[rev] < 10:
                    affinity_candidates.append(rev)
        
        if affinity_candidates:
            # Prefer the one with lower workload to balance, else alphabetically
            affinity_candidates.sort(key=lambda m: (workload[m], m))
            selected_user = affinity_candidates[0]
            reason = f"Workflow Affinity (via issue tracking, mapping to {selected_user})"

        # Priority 2: Underloaded Balancing (C_user < 5)
        if not selected_user:
            underloaded_members = [m for m in TEAM if workload[m] < 5]
            if underloaded_members:
                # Select the member with the lowest C_user (alphabetical for tie-break)
                underloaded_members.sort(key=lambda m: (workload[m], m))
                selected_user = underloaded_members[0]
                reason = f"Underloaded Balancing (C_user < 5, currently {workload[selected_user]})"

        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        if not selected_user:
            capacity_members = [m for m in TEAM if 5 <= workload[m] < 10]
            if capacity_members:
                capacity_members.sort(key=lambda m: (workload[m], m))
                selected_user = capacity_members[0]
                reason = f"Capacity Absorption (5 <= C_user < 10, currently {workload[selected_user]})"

        if selected_user:
            # Assign PR
            assignments.append((number, selected_user, reason, title))
            
            # Update state
            workload[selected_user] += 1
            for issue_id in issue_ids:
                tracking_issue_to_reviewer[issue_id] = selected_user
            
            print(f"Selected {selected_user} for PR #{number} due to: {reason}")
        else:
            print(f"Skipping PR #{number} because no team members can absorb more workload.")

    # Step 4: Execute Assignments & Exit
    print("\n--- Summary of Assignments to Execute ---")
    if not assignments:
        print("No new assignments needed or possible.")
    else:
        for number, selected_user, reason, title in assignments:
            print(f"PR #{number} -> {selected_user} ({reason}) | Title: '{title}'")

        print("\nExecuting assignments on GitHub...")
        for number, selected_user, reason, title in assignments:
            print(f"Assigning {selected_user} to PR #{number}...")
            cmd = [
                "gh", "pr", "edit", str(number),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", selected_user
            ]
            res = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            if res.returncode == 0:
                print(f"Successfully assigned {selected_user} to PR #{number}.")
            else:
                print(f"Failed to assign {selected_user} to PR #{number}: {res.stderr}", file=sys.stderr)

    print("\n--- Final Workload Summary ---")
    for member in sorted(TEAM):
        print(f"  {member}: {workload[member]} assigned reviews")

if __name__ == "__main__":
    main()
