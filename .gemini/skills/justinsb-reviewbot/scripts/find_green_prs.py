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
import subprocess
import os
import sys

def run_cmd(cmd, cwd=None):
    res = subprocess.run(cmd, capture_output=True, text=True, cwd=cwd)
    return res

def main():
    kcc_root = os.path.expanduser("~/projects/kcc/reviewprs/k8s-config-connector")
    print("Fetching last 50 assigned PRs (metadata only)...")
    
    # Query 50 PRs without status checks to avoid GraphQL 504 timeouts
    cmd = ["gh", "pr", "list", "--assignee", "justinsb", "--json", "number,title,mergeable", "--limit", "50"]
    res = run_cmd(cmd, cwd=kcc_root)
    if res.returncode != 0:
        print("Error fetching PRs list:", res.stderr, file=sys.stderr)
        sys.exit(1)

    prs = json.loads(res.stdout)
    mergeable_prs = [pr for pr in prs if pr.get("mergeable") == "MERGEABLE"]
    
    print(f"Found {len(mergeable_prs)} mergeable PRs. Fetching status checks for each...")
    green_prs = []
    
    for pr in mergeable_prs:
        num = pr["number"]
        title = pr["title"]
        
        # Fetch status check rollup for this specific PR (fast, single-PR query)
        view_cmd = ["gh", "pr", "view", str(num), "--json", "statusCheckRollup"]
        view_res = run_cmd(view_cmd, cwd=kcc_root)
        if view_res.returncode != 0:
            print(f"Warning: Failed to fetch checks for PR #{num}: {view_res.stderr}", file=sys.stderr)
            continue
            
        data = json.loads(view_res.stdout)
        checks = data.get("statusCheckRollup") or []
        
        has_failure = False
        has_pending = False
        total_checks = 0
        for c in checks:
            conclusion = c.get("conclusion")
            state = c.get("state")
            status = c.get("status")
            
            if conclusion == "FAILURE" or state == "FAILURE":
                has_failure = True
                break
            elif status in ["IN_PROGRESS", "QUEUED", "REQUESTED", "WAITING"] or state in ["PENDING"]:
                has_pending = True
            
            total_checks += 1
            
        if not has_failure and total_checks > 0:
            green_prs.append((num, title, has_pending, total_checks))

    print("\n--- Green / Candidate PRs ---")
    if not green_prs:
        print("No completely green PRs found in the latest 50.")
        return
        
    for num, title, pending, count in green_prs:
        status_str = "⏳ Pending checks" if pending else "✅ All checks passed"
        print(f"#{num} - {title} ({status_str}, {count} total checks)")

if __name__ == "__main__":
    main()
