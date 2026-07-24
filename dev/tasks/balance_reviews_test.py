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

import unittest
from dev.tasks.balance_reviews import audit_workloads, get_candidate_prs, assign_reviews

class TestPRReviewWorkloadBalancer(unittest.TestCase):
    def setUp(self):
        self.team = ['acpana', 'anfernee', 'anhdle-sso', 'barney-s', 'gemmahou', 'maqiuyujoyce']

    def test_audit_workloads(self):
        prs = [
            {
                'number': 100,
                'title': 'Test 1 Fixes #10976',
                'body': 'Child of #10276',
                'requested_reviewers': ['acpana']
            },
            {
                'number': 101,
                'title': 'Test 2',
                'body': 'Just a test',
                'requested_reviewers': ['anfernee']
            }
        ]
        workload, tracking = audit_workloads(prs, self.team)
        
        # Verify workload counts
        self.assertEqual(workload['acpana'], 1)
        self.assertEqual(workload['anfernee'], 1)
        self.assertEqual(workload['anhdle-sso'], 0)
        
        # Verify issue tracking maps
        self.assertEqual(tracking['10976'], 'acpana')
        self.assertEqual(tracking['10276'], 'acpana')

    def test_get_candidate_prs_only_ready_for_human(self):
        prs = [
            {
                'number': 100,
                'labels': ['ready-for-human'],
                'requested_reviewers': []
            },
            {
                'number': 101,
                'labels': [],
                'requested_reviewers': []
            },
            {
                'number': 102,
                'labels': ['ready-for-human'],
                'requested_reviewers': ['acpana'] # Already has a team member
            }
        ]
        candidates = get_candidate_prs(prs, self.team)
        self.assertEqual(len(candidates), 1)
        self.assertEqual(candidates[0]['number'], 100)

    def test_priority1_workflow_affinity(self):
        # acpana is reviewing something related to issue #1234
        workload = {m: 0 for m in self.team}
        workload['acpana'] = 4
        tracking = {'1234': 'acpana'}
        
        candidates = [
            {
                'number': 200,
                'title': 'Fixes #1234',
                'body': '',
                'labels': ['ready-for-human']
            }
        ]
        
        assignments = assign_reviews(candidates, workload, tracking, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]['reviewer'], 'acpana')
        self.assertEqual(assignments[0]['reason'], 'Priority 1 (Workflow Affinity)')
        self.assertEqual(workload['acpana'], 5)

    def test_priority2_underloaded_balancing(self):
        # Everyone starts at workload of 3, but anfernee is at 2, and barney-s is at 1
        workload = {m: 3 for m in self.team}
        workload['anfernee'] = 2
        workload['barney-s'] = 1
        
        tracking = {}
        
        candidates = [
            {
                'number': 200,
                'title': 'No issue references',
                'body': '',
                'labels': ['ready-for-human']
            }
        ]
        
        # Barney-s has the lowest count (1), so they should get selected
        assignments = assign_reviews(candidates, workload, tracking, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]['reviewer'], 'barney-s')
        self.assertEqual(assignments[0]['reason'], 'Priority 2 (Underloaded Balancing)')
        self.assertEqual(workload['barney-s'], 2)

    def test_priority2_underloaded_balancing_tie_break(self):
        # Both acpana and anfernee are at workload 2, others at 4
        workload = {m: 4 for m in self.team}
        workload['acpana'] = 2
        workload['anfernee'] = 2
        
        tracking = {}
        
        candidates = [
            {
                'number': 200,
                'title': 'No issue references',
                'body': '',
                'labels': ['ready-for-human']
            }
        ]
        
        # acpana and anfernee are tied, but acpana is alphabetically first
        assignments = assign_reviews(candidates, workload, tracking, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]['reviewer'], 'acpana')

    def test_priority3_capacity_absorption(self):
        # Everyone is at 5, except acpana at 6, and barney-s is at 5, but let's make some at 5, some at 7
        workload = {m: 7 for m in self.team}
        workload['anhdle-sso'] = 6
        workload['barney-s'] = 5
        
        tracking = {}
        
        candidates = [
            {
                'number': 200,
                'title': 'No issue references',
                'body': '',
                'labels': ['ready-for-human']
            }
        ]
        
        # Everyone >= 5, so Priority 2 is skipped. Barney-s is lowest at 5, so they get selected.
        assignments = assign_reviews(candidates, workload, tracking, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]['reviewer'], 'barney-s')
        self.assertEqual(assignments[0]['reason'], 'Priority 3 (Capacity Absorption)')

    def test_hard_ceiling(self):
        # Everyone is at 10 (hard ceiling)
        workload = {m: 10 for m in self.team}
        tracking = {}
        
        candidates = [
            {
                'number': 200,
                'title': 'No issue references',
                'body': '',
                'labels': ['ready-for-human']
            }
        ]
        
        # No assignment because everyone is at 10
        assignments = assign_reviews(candidates, workload, tracking, self.team)
        self.assertEqual(len(assignments), 0)

if __name__ == '__main__':
    unittest.main()
