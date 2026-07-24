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
from dev.tasks.balance_reviews import extract_issue_ids, compute_assignments

class TestBalanceReviews(unittest.TestCase):
    def test_extract_issue_ids(self):
        title = "Greenfield: Implement controller for resource Fixes #10276"
        body = "Workflow: #10976. Some other text with #notanumber and #12345."
        result = extract_issue_ids(title, body)
        self.assertEqual(result, {10276, 10976, 12345})

    def test_compute_assignments_priority_1_affinity(self):
        # Setup:
        # User A has 1 review assigned, which references issue #100.
        # User B has 0 reviews assigned.
        # A new candidate PR references issue #100.
        # Priority 1: Should assign to User A due to affinity, even though User B is more underloaded.
        team = ["userA", "userB"]
        prs = [
            {
                "number": 1,
                "title": "Fixes #100",
                "body": "",
                "labels": ["some-label"],
                "requested_reviewers": ["userA"]
            },
            {
                "number": 2,
                "title": "Candidate PR referencing #100",
                "body": "Fixes #100",
                "labels": ["ready-for-human"],
                "requested_reviewers": []
            }
        ]
        workload, tracking_issue_to_reviewer, candidates, assignments = compute_assignments(prs, team)
        
        self.assertEqual(len(candidates), 1)
        self.assertEqual(candidates[0]["number"], 2)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0], (2, "userA", "Candidate PR referencing #100"))
        self.assertEqual(workload["userA"], 2)
        self.assertEqual(workload["userB"], 0)

    def test_compute_assignments_priority_2_underloaded(self):
        # Setup:
        # User A has 3 reviews assigned.
        # User B has 1 review assigned.
        # User C has 5 reviews assigned (fully loaded base, but below ceiling).
        # A new candidate PR has no affinity.
        # Priority 2: Should assign to the most underloaded member (< 5 workload), which is User B.
        team = ["userA", "userB", "userC"]
        prs = [
            {"number": 1, "title": "A1", "body": "", "labels": [], "requested_reviewers": ["userA"]},
            {"number": 2, "title": "A2", "body": "", "labels": [], "requested_reviewers": ["userA"]},
            {"number": 3, "title": "A3", "body": "", "labels": [], "requested_reviewers": ["userA"]},
            {"number": 4, "title": "B1", "body": "", "labels": [], "requested_reviewers": ["userB"]},
            {"number": 5, "title": "C1", "body": "", "labels": [], "requested_reviewers": ["userC"]},
            {"number": 6, "title": "C2", "body": "", "labels": [], "requested_reviewers": ["userC"]},
            {"number": 7, "title": "C3", "body": "", "labels": [], "requested_reviewers": ["userC"]},
            {"number": 8, "title": "C4", "body": "", "labels": [], "requested_reviewers": ["userC"]},
            {"number": 9, "title": "C5", "body": "", "labels": [], "requested_reviewers": ["userC"]},
            {
                "number": 10,
                "title": "Candidate PR",
                "body": "",
                "labels": ["ready-for-human"],
                "requested_reviewers": []
            }
        ]
        workload, tracking_issue_to_reviewer, candidates, assignments = compute_assignments(prs, team)
        
        self.assertEqual(len(candidates), 1)
        self.assertEqual(assignments[0], (10, "userB", "Candidate PR"))
        self.assertEqual(workload["userB"], 2)

    def test_compute_assignments_priority_3_capacity(self):
        # Setup:
        # User A has 5 reviews assigned.
        # User B has 6 reviews assigned.
        # Every member is >= 5 workload (so no underloaded members).
        # A new candidate PR has no affinity.
        # Priority 3: Should assign to the member with the lowest workload under 10, which is User A.
        team = ["userA", "userB"]
        prs = []
        for i in range(5):
            prs.append({"number": i, "title": f"A{i}", "body": "", "labels": [], "requested_reviewers": ["userA"]})
        for i in range(6):
            prs.append({"number": 100+i, "title": f"B{i}", "body": "", "labels": [], "requested_reviewers": ["userB"]})
        prs.append({
            "number": 200,
            "title": "Candidate PR",
            "body": "",
            "labels": ["ready-for-human"],
            "requested_reviewers": []
        })
        workload, tracking_issue_to_reviewer, candidates, assignments = compute_assignments(prs, team)
        
        self.assertEqual(len(candidates), 1)
        self.assertEqual(assignments[0], (200, "userA", "Candidate PR"))
        self.assertEqual(workload["userA"], 6)

    def test_compute_assignments_hard_ceiling(self):
        # Setup:
        # User A has 10 reviews assigned (at ceiling).
        # User B has 9 reviews assigned.
        # A new candidate PR has affinity with User A.
        # But User A is at ceiling (10 reviews).
        # So User A must NOT be assigned additional PRs.
        # It should assign to User B (Priority 3).
        team = ["userA", "userB"]
        prs = []
        for i in range(10):
            prs.append({"number": i, "title": f"A{i} Fixes #100", "body": "", "labels": [], "requested_reviewers": ["userA"]})
        for i in range(9):
            prs.append({"number": 100+i, "title": f"B{i}", "body": "", "labels": [], "requested_reviewers": ["userB"]})
        prs.append({
            "number": 200,
            "title": "Candidate PR with Affinity for A",
            "body": "Fixes #100",
            "labels": ["ready-for-human"],
            "requested_reviewers": []
        })
        workload, tracking_issue_to_reviewer, candidates, assignments = compute_assignments(prs, team)
        
        self.assertEqual(len(candidates), 1)
        self.assertEqual(assignments[0], (200, "userB", "Candidate PR with Affinity for A"))
        self.assertEqual(workload["userA"], 10)
        self.assertEqual(workload["userB"], 10)

if __name__ == "__main__":
    unittest.main()
