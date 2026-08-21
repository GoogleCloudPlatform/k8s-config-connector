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
from balance_reviews import extract_issue_ids, audit_workloads, balance_reviews

class TestBalanceReviews(unittest.TestCase):
    
    def test_extract_issue_ids(self):
        # Normal references
        self.assertEqual(extract_issue_ids("Title with #123", "Body with #456 and #123"), [123, 456])
        # Empty title/body
        self.assertEqual(extract_issue_ids("", None), [])
        # No issues
        self.assertEqual(extract_issue_ids("No issues here", "Some text"), [])
        # Multi-digit numbers
        self.assertEqual(extract_issue_ids("Fixes #10976 and #10276", ""), [10276, 10976])

    def test_audit_workloads(self):
        team = ["alice", "bob", "charlie"]
        
        prs = [
            {
                "number": 1,
                "title": "PR 1",
                "body": "Fixes #100",
                "requested_reviewers": [{"login": "alice"}],
                "labels": []
            },
            {
                "number": 2,
                "title": "PR 2",
                "body": "Workflow: #200",
                "requested_reviewers": [{"login": "bob"}, {"login": "outsider"}],
                "labels": []
            },
            {
                "number": 3,
                "title": "Candidate 1",
                "body": "Fixes #100",
                "requested_reviewers": [{"login": "outsider"}],
                "labels": [{"name": "ready-for-human"}]
            },
            {
                "number": 4,
                "title": "Candidate 2",
                "body": "No tracking issue",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            },
            {
                "number": 5,
                "title": "Not ready",
                "body": "",
                "requested_reviewers": [],
                "labels": []
            }
        ]
        
        workloads, tracking_map, candidates = audit_workloads(prs, team)
        
        # Check workloads
        self.assertEqual(workloads, {"alice": 1, "bob": 1, "charlie": 0})
        
        # Check tracking issue maps
        self.assertEqual(tracking_map, {100: {"alice"}, 200: {"bob"}})
        
        # Check candidate extraction (only #3 and #4 should be candidates)
        candidate_numbers = [pr["number"] for pr in candidates]
        self.assertEqual(candidate_numbers, [3, 4])

    def test_balance_reviews_priority_1_workflow_affinity(self):
        team = ["alice", "bob"]
        
        # Alice is already reviewing issue 100 on another PR
        workloads = {"alice": 1, "bob": 0}
        tracking_map = {100: {"alice"}}
        
        candidates = [
            {
                "number": 3,
                "title": "PR 3",
                "body": "Fixes #100",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        
        assignments = balance_reviews(candidates, team, workloads, tracking_map)
        
        # Even though Bob has lower workload (0 < 1), Alice should be selected due to Workflow Affinity
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]["selected_user"], "alice")
        self.assertIn("Workflow Affinity", assignments[0]["reason"])
        
        # Verify state was updated
        self.assertEqual(workloads["alice"], 2)
        self.assertEqual(tracking_map[100], {"alice"})

    def test_balance_reviews_priority_2_underloaded_balancing(self):
        team = ["alice", "bob"]
        
        # Nobody is reviewing anything related to issue 300
        workloads = {"alice": 2, "bob": 1}
        tracking_map = {}
        
        candidates = [
            {
                "number": 3,
                "title": "PR 3",
                "body": "Fixes #300",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        
        assignments = balance_reviews(candidates, team, workloads, tracking_map)
        
        # Bob has lower workload and workload < 5, so Bob is selected
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]["selected_user"], "bob")
        self.assertIn("Underloaded Balancing", assignments[0]["reason"])

    def test_balance_reviews_priority_3_capacity_absorption(self):
        team = ["alice", "bob"]
        
        # All members are at or above 5 but below 10
        workloads = {"alice": 6, "bob": 5}
        tracking_map = {}
        
        candidates = [
            {
                "number": 3,
                "title": "PR 3",
                "body": "Fixes #300",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        
        assignments = balance_reviews(candidates, team, workloads, tracking_map)
        
        # Bob has lower workload, Bob is selected via Capacity Absorption
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0]["selected_user"], "bob")
        self.assertIn("Capacity Absorption", assignments[0]["reason"])

    def test_balance_reviews_hard_ceiling(self):
        team = ["alice", "bob"]
        
        # Both are at ceiling or above
        workloads = {"alice": 10, "bob": 10}
        tracking_map = {}
        
        candidates = [
            {
                "number": 3,
                "title": "PR 3",
                "body": "Fixes #300",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        
        assignments = balance_reviews(candidates, team, workloads, tracking_map)
        
        # Should skip assignment due to ceiling
        self.assertEqual(len(assignments), 0)

if __name__ == "__main__":
    unittest.main()
