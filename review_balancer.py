import subprocess
import json
import re
import sys

# Define the team members
TEAM = ["acpana", "anfernee", "anhdle-sso", "barney-s", "gemmahou", "maqiuyujoyce"]

def fetch_open_prs():
    prs = []
    page = 1
    while True:
        print(f"Fetching page {page} of open PRs via REST API...")
        cmd = [
            "gh", "api", "-X", "GET", 
            "repos/GoogleCloudPlatform/k8s-config-connector/pulls",
            "-f", "state=open",
            "-f", "per_page=100",
            "-f", f"page={page}"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
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
    return prs

def main():
    print("=== Step 1: Audit Current Team Workloads & Multilevel Workflow Mapping ===")
    prs = fetch_open_prs()
    print(f"Total open PRs retrieved: {len(prs)}")
    
    # Initialize workload tracking structures
    workload = {member: 0 for member in TEAM}
    tracking_issue_to_reviewer = {}
    
    issue_regex = re.compile(r'#(\d+)')
    
    for pr in prs:
        number = pr['number']
        title = pr['title']
        body = pr['body'] or ""
        labels = [l['name'] for l in pr.get('labels', [])]
        
        # Extract all referenced issue IDs
        issues = sorted(list(set(map(int, issue_regex.findall(title + " " + body)))))
        
        # Get requested reviewers
        requested_reviewers = [r['login'] for r in pr.get('requested_reviewers', [])]
        
        # Check if any team member is requested
        team_reviewers = [r for r in requested_reviewers if r in workload]
        
        if team_reviewers:
            for r in team_reviewers:
                workload[r] += 1
                for issue in issues:
                    tracking_issue_to_reviewer[issue] = r
                    
    print("\n--- Initial Team Workloads ---")
    for member, count in workload.items():
        print(f"  {member}: {count} open review requests")
        
    print("\n=== Step 2: Fetch Unassigned Candidate PRs ===")
    candidates = []
    for pr in prs:
        number = pr['number']
        title = pr['title']
        body = pr['body'] or ""
        labels = [l['name'] for l in pr.get('labels', [])]
        
        if "ready-for-human" in labels:
            requested_reviewers = [r['login'] for r in pr.get('requested_reviewers', [])]
            team_reviewers = [r for r in requested_reviewers if r in workload]
            
            if not team_reviewers:
                issues = sorted(list(set(map(int, issue_regex.findall(title + " " + body)))))
                candidates.append({
                    "number": number,
                    "title": title,
                    "issues": issues,
                    "labels": labels
                })
                
    # Sort candidates by PR number ascending (oldest first)
    candidates.sort(key=lambda x: x['number'])
    print(f"Found {len(candidates)} unassigned candidate PRs labeled 'ready-for-human'.")
    for c in candidates:
        print(f"  PR #{c['number']}: '{c['title']}' (Issues: {c['issues']})")
        
    print("\n=== Step 3: Assignment Algorithm ===")
    assignments = []
    
    for c in candidates:
        number = c['number']
        title = c['title']
        issues = c['issues']
        
        selected_user = None
        
        # Priority 1: Workflow Affinity
        for issue in issues:
            if issue in tracking_issue_to_reviewer:
                potential_reviewer = tracking_issue_to_reviewer[issue]
                if workload[potential_reviewer] < 10:
                    selected_user = potential_reviewer
                    print(f"  PR #{number}: Workflow affinity match with issue #{issue}. Selected: {selected_user} (current workload: {workload[selected_user]})")
                    break
        
        # Priority 2: Underloaded Balancing (C_user < 5)
        if not selected_user:
            underloaded = [m for m in TEAM if workload[m] < 5]
            if underloaded:
                # Sort underloaded members by workload (ascending), then alphabetically by name to break ties
                underloaded.sort(key=lambda m: (workload[m], m))
                selected_user = underloaded[0]
                print(f"  PR #{number}: Underloaded balancing (< 5). Selected: {selected_user} (current workload: {workload[selected_user]})")
                
        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        if not selected_user:
            available = [m for m in TEAM if workload[m] < 10]
            if available:
                # Sort by workload (ascending), then alphabetically
                available.sort(key=lambda m: (workload[m], m))
                selected_user = available[0]
                print(f"  PR #{number}: Capacity absorption (5 <= count < 10). Selected: {selected_user} (current workload: {workload[selected_user]})")
                
        if selected_user:
            # Update State
            workload[selected_user] += 1
            for issue in issues:
                tracking_issue_to_reviewer[issue] = selected_user
            
            assignments.append((number, selected_user))
        else:
            print(f"  PR #{number}: Skipped. All team members have reached the hard ceiling of 10.")
            
    print("\n=== Step 4: Execute Assignments & Exit ===")
    if not assignments:
        print("No new reviewer assignments to execute.")
    else:
        for number, user in assignments:
            print(f"Executing: Assigning {user} to PR #{number}...")
            cmd = [
                "gh", "pr", "edit", str(number),
                "--repo", "GoogleCloudPlatform/k8s-config-connector",
                "--add-reviewer", user
            ]
            result = subprocess.run(cmd, capture_output=True, text=True)
            if result.returncode != 0:
                print(f"  FAILED to assign {user} to PR #{number}: {result.stderr}", file=sys.stderr)
            else:
                print(f"  Successfully assigned {user} to PR #{number}.")
                
    print("\n--- Final Workload Projection ---")
    for member, count in workload.items():
        print(f"  {member}: {count} open review requests")
        
    print("\nSummary of assignments made:")
    if not assignments:
        print("  None")
    for number, user in assignments:
        print(f"  PR #{number} -> {user}")

if __name__ == "__main__":
    main()
