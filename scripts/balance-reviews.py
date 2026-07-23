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
import unittest

TEAM = {
    'acpana',
    'anfernee',
    'anhdle-sso',
    'barney-s',
    'gemmahou',
    'maqiuyujoyce'
}

class ReviewBalancer:
    def __init__(self, dry_run=False, verbose=False):
        self.dry_run = dry_run
        self.verbose = verbose
        self.workload = {member: 0 for member in TEAM}
        self.tracking_issue_to_reviewer = {}
        self.queued_assignments = []  # List of tuples: (pr_number, reviewer, reason)

    def extract_issue_ids(self, text):
        if not text:
            return []
        # Find all occurrences of #<number>
        return [int(num) for num in re.findall(r'#(\d+)', text)]

    def fetch_all_open_prs(self):
        if self.verbose:
            print("Fetching all open PRs via GitHub REST API...")
        prs = []
        page = 1
        while True:
            cmd = ["gh", "api", f"repos/GoogleCloudPlatform/k8s-config-connector/pulls?state=open&per_page=100&page={page}"]
            result = subprocess.run(cmd, capture_output=True, text=True)
            if result.returncode != 0:
                print(f"Error fetching open PRs page {page}: {result.stderr}", file=sys.stderr)
                break
            data = json.loads(result.stdout)
            if not data:
                break
            prs.extend(data)
            if len(data) < 100:
                break
            page += 1
        return prs

    def audit_workloads_and_workflow_mapping(self, prs):
        for pr in prs:
            title = pr.get('title', '')
            body = pr.get('body', '') or ''
            issue_ids = self.extract_issue_ids(f"{title} {body}")
            
            reviewers = [r['login'] for r in pr.get('requested_reviewers', [])]
            
            # Map issues to KCC team member requested reviewers
            for r in reviewers:
                if r in TEAM:
                    self.workload[r] += 1
                    for issue_id in issue_ids:
                        self.tracking_issue_to_reviewer[issue_id] = r

        if self.verbose:
            print("\nCurrent workloads:")
            for member, count in sorted(self.workload.items(), key=lambda x: (x[1], x[0])):
                print(f" - {member}: {count} open PRs assigned")

    def run_assignment(self, prs):
        # Filter candidate PRs
        candidates = []
        for pr in prs:
            labels = [l['name'] for l in pr.get('labels', [])]
            if 'ready-for-human' not in labels:
                continue
            
            reviewers = [r['login'] for r in pr.get('requested_reviewers', [])]
            # Must not have any member of the team assigned
            if any(r in TEAM for r in reviewers):
                continue
            
            candidates.append(pr)

        if self.verbose:
            print(f"\nFound {len(candidates)} candidate unassigned 'ready-for-human' PR(s).")

        # Sort candidates by PR number ascending to make assignments deterministic
        candidates.sort(key=lambda x: x['number'])

        for pr in candidates:
            pr_number = pr['number']
            title = pr.get('title', '')
            body = pr.get('body', '') or ''
            issue_ids = self.extract_issue_ids(f"{title} {body}")

            # Check if all team members have reached the ceiling of 10
            if all(self.workload[m] >= 10 for m in TEAM):
                if self.verbose:
                    print("Hard ceiling of 10 reviews reached for all team members. Skipping further assignments.")
                break

            reviewer, reason = self.select_reviewer(issue_ids)
            if reviewer:
                # Update State
                self.workload[reviewer] += 1
                for issue_id in issue_ids:
                    self.tracking_issue_to_reviewer[issue_id] = reviewer
                
                self.queued_assignments.append((pr_number, reviewer, reason))
                if self.verbose:
                    print(f"PR #{pr_number} queueing assignment to {reviewer} (Reason: {reason})")

    def select_reviewer(self, issue_ids):
        # Priority 1: Workflow Affinity
        affinity_reviewers = set()
        for issue_id in issue_ids:
            if issue_id in self.tracking_issue_to_reviewer:
                rev = self.tracking_issue_to_reviewer[issue_id]
                if self.workload[rev] < 10:
                    affinity_reviewers.add(rev)
        
        if affinity_reviewers:
            # Select the reviewer with lowest current workload, then alphabetically
            selected = min(affinity_reviewers, key=lambda x: (self.workload[x], x))
            return selected, "Workflow Affinity"

        # Priority 2: Underloaded Balancing (C_user < 5)
        underloaded = [m for m in TEAM if self.workload[m] < 5]
        if underloaded:
            # Select member with lowest workload, then alphabetically
            selected = min(underloaded, key=lambda x: (self.workload[x], x))
            return selected, "Underloaded Balancing"

        # Priority 3: Capacity Absorption (5 <= C_user < 10)
        capacity = [m for m in TEAM if self.workload[m] < 10]
        if capacity:
            # Select member with lowest workload, then alphabetically
            selected = min(capacity, key=lambda x: (self.workload[x], x))
            return selected, "Capacity Absorption"

        return None, "All Members Saturated"

    def execute_assignments(self):
        if not self.queued_assignments:
            print("No new PR review assignments to make.")
            return

        print(f"\nExecuting {len(self.queued_assignments)} assignments...")
        for pr_number, reviewer, reason in self.queued_assignments:
            if self.dry_run:
                print(f"[DRY-RUN] Would assign PR #{pr_number} to {reviewer} (Reason: {reason})")
            else:
                cmd = ["gh", "pr", "edit", str(pr_number), "--repo", "GoogleCloudPlatform/k8s-config-connector", "--add-reviewer", reviewer]
                result = subprocess.run(cmd, capture_output=True, text=True)
                if result.returncode == 0:
                    print(f"Successfully assigned PR #{pr_number} to {reviewer} (Reason: {reason})")
                else:
                    print(f"Error assigning PR #{pr_number} to {reviewer}: {result.stderr}", file=sys.stderr)


# Self-tests for validating balancer logic
class TestReviewBalancer(unittest.TestCase):
    def setUp(self):
        self.balancer = ReviewBalancer(dry_run=True, verbose=False)

    def test_extract_issue_ids(self):
        text = "Fixes #10976 and references #10276"
        self.assertEqual(self.balancer.extract_issue_ids(text), [10976, 10276])
        self.assertEqual(self.balancer.extract_issue_ids("No issues here"), [])
        self.assertEqual(self.balancer.extract_issue_ids(""), [])

    def test_priority1_workflow_affinity(self):
        # Setup workload where everyone has different workloads
        self.balancer.workload = {
            'acpana': 2,
            'anfernee': 4,
            'anhdle-sso': 6,
            'barney-s': 6,
            'gemmahou': 6,
            'maqiuyujoyce': 6
        }
        # anfernee is already reviewing workflow #12345
        self.balancer.tracking_issue_to_reviewer[12345] = 'anfernee'

        # Candidate references #12345. Even though acpana has lower workload (2 < 4),
        # anfernee should be selected due to Workflow Affinity.
        reviewer, reason = self.balancer.select_reviewer([12345])
        self.assertEqual(reviewer, 'anfernee')
        self.assertEqual(reason, 'Workflow Affinity')

    def test_priority2_underloaded_balancing(self):
        self.balancer.workload = {
            'acpana': 4,
            'anfernee': 3,
            'anhdle-sso': 6,
            'barney-s': 6,
            'gemmahou': 6,
            'maqiuyujoyce': 6
        }
        # Candidate has no workflow affinity.
        # anfernee has the lowest workload under 5, so should be selected.
        reviewer, reason = self.balancer.select_reviewer([99999])
        self.assertEqual(reviewer, 'anfernee')
        self.assertEqual(reason, 'Underloaded Balancing')

    def test_priority2_underloaded_balancing_tiebreaker(self):
        self.balancer.workload = {
            'acpana': 3,
            'anfernee': 3,
            'anhdle-sso': 6,
            'barney-s': 6,
            'gemmahou': 6,
            'maqiuyujoyce': 6
        }
        # Both acpana and anfernee have workload 3. Alphabetical tiebreaker: acpana < anfernee.
        reviewer, reason = self.balancer.select_reviewer([99999])
        self.assertEqual(reviewer, 'acpana')
        self.assertEqual(reason, 'Underloaded Balancing')

    def test_priority3_capacity_absorption(self):
        # Everyone is at 5 or more reviews.
        self.balancer.workload = {
            'acpana': 5,
            'anfernee': 6,
            'anhdle-sso': 7,
            'barney-s': 7,
            'gemmahou': 7,
            'maqiuyujoyce': 7
        }
        # No workflow affinity. acpana has lowest workload < 10, so should absorb.
        reviewer, reason = self.balancer.select_reviewer([99999])
        self.assertEqual(reviewer, 'acpana')
        self.assertEqual(reason, 'Capacity Absorption')

    def test_saturation_limit(self):
        # Everyone is saturated at 10 reviews.
        self.balancer.workload = {member: 10 for member in TEAM}
        reviewer, reason = self.balancer.select_reviewer([99999])
        self.assertIsNone(reviewer)
        self.assertEqual(reason, 'All Members Saturated')


def main():
    parser = argparse.ArgumentParser(description="PR Review Workload Balancer")
    parser.add_argument("--dry-run", action="store_true", help="Print assignment decisions without executing them")
    parser.add_argument("--verbose", action="store_true", help="Print verbose execution and mapping logs")
    parser.add_argument("--run-tests", action="store_true", help="Run self-tests and exit")
    args = parser.parse_args()

    if args.run_tests:
        # Filter unittest argv to avoid conflict with argparse
        unittest.main(argv=[sys.argv[0]])
        return

    balancer = ReviewBalancer(dry_run=args.dry_run, verbose=args.verbose)
    prs = balancer.fetch_all_open_prs()
    balancer.audit_workloads_and_workflow_mapping(prs)
    balancer.run_assignment(prs)
    balancer.execute_assignments()

if __name__ == "__main__":
    main()
