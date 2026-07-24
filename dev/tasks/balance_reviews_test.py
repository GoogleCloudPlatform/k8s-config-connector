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
from balance_reviews import balance_workloads, extract_issue_ids

class TestBalanceReviews(unittest.TestCase):

    def setUp(self):
        self.team = {
            "acpana",
            "anfernee",
            "anhdle-sso",
            "barney-s",
            "gemmahou",
            "maqiuyujoyce"
        }

    def test_extract_issue_ids(self):
        title = "Fixes #10276 and #10976"
        body = "This fixes the issues in Workflow: #10976. Child issue is #10276."
        issues = extract_issue_ids(title, body)
        self.assertEqual(issues, ["10276", "10976"])

    def test_workflow_affinity_priority(self):
        # PR 1 is assigned to anfernee, references issue 10276
        # PR 2 is ready-for-human and unassigned, also references issue 10276
        prs = [
            {
                "number": 1,
                "title": "PR 1 referencing #10276",
                "body": "",
                "requested_reviewers": [{"login": "anfernee"}],
                "labels": []
            },
            {
                "number": 2,
                "title": "PR 2 referencing #10276",
                "body": "",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        assignments, workload, _ = balance_workloads(prs, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0][0], 2)
        self.assertEqual(assignments[0][1], "anfernee")
        self.assertIn("Workflow Affinity", assignments[0][2])

    def test_underloaded_balancing_priority(self):
        # No workflow affinity. Let's set up existing workloads:
        # acpana has 4, anfernee has 2, others have 0.
        # Deterministic alphabetical selection when workloads are tied.
        prs = [
            {
                "number": 1,
                "requested_reviewers": [{"login": "acpana"}],
            },
            {
                "number": 2,
                "requested_reviewers": [{"login": "acpana"}],
            },
            {
                "number": 3,
                "requested_reviewers": [{"login": "acpana"}],
            },
            {
                "number": 4,
                "requested_reviewers": [{"login": "acpana"}],
            },
            {
                "number": 5,
                "requested_reviewers": [{"login": "anfernee"}],
            },
            {
                "number": 6,
                "requested_reviewers": [{"login": "anfernee"}],
            },
            # Candidate PR
            {
                "number": 100,
                "title": "New candidate",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        assignments, workload, _ = balance_workloads(prs, self.team)
        # Members with workload 0 are: anhdle-sso, barney-s, gemmahou, maqiuyujoyce
        # Alphabetically, anhdle-sso is the first.
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0][1], "anhdle-sso")
        self.assertIn("Underloaded Balancing", assignments[0][2])

    def test_capacity_absorption_priority(self):
        # Set all team workloads to >= 5.
        # acpana has 6, anfernee has 5, others have 7.
        prs = []
        # acpana: 6
        for i in range(6):
            prs.append({"number": i, "requested_reviewers": [{"login": "acpana"}]})
        # anfernee: 5
        for i in range(5):
            prs.append({"number": 10 + i, "requested_reviewers": [{"login": "anfernee"}]})
        # all others: 7 each
        for member in self.team:
            if member not in ["acpana", "anfernee"]:
                for i in range(7):
                    prs.append({"number": 100 + i, "requested_reviewers": [{"login": member}]})
                    
        # Candidate PR
        prs.append({
            "number": 1000,
            "title": "Unassigned PR",
            "requested_reviewers": [],
            "labels": [{"name": "ready-for-human"}]
        })
        
        assignments, workload, _ = balance_workloads(prs, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0][1], "anfernee") # lowest workload which is 5
        self.assertIn("Capacity Absorption", assignments[0][2])

    def test_ceiling_reached(self):
        # Set all team workloads to 10.
        prs = []
        for member in self.team:
            for i in range(10):
                prs.append({"number": i, "requested_reviewers": [{"login": member}]})
                
        # Candidate PR
        prs.append({
            "number": 1000,
            "title": "Unassigned PR",
            "requested_reviewers": [],
            "labels": [{"name": "ready-for-human"}]
        })
        
        assignments, workload, _ = balance_workloads(prs, self.team)
        self.assertEqual(len(assignments), 0)

    def test_multilevel_workflow_mapping(self):
        # Root issue: #10976, sub-issue: #10276
        # Open PR #10992 references both #10276 and #10976, is reviewed by anfernee
        # Unassigned candidate PR #11000 references root issue #10976.
        # It should match anfernee via workflow affinity.
        prs = [
            {
                "number": 10992,
                "title": "Sub-issue PR for #10276 and root #10976",
                "body": "",
                "requested_reviewers": [{"login": "anfernee"}],
                "labels": []
            },
            {
                "number": 11000,
                "title": "Root issue PR referencing #10976",
                "body": "",
                "requested_reviewers": [],
                "labels": [{"name": "ready-for-human"}]
            }
        ]
        assignments, workload, tracking_issue_to_reviewer = balance_workloads(prs, self.team)
        self.assertEqual(len(assignments), 1)
        self.assertEqual(assignments[0][1], "anfernee")
        self.assertIn("Workflow Affinity", assignments[0][2])
        self.assertEqual(tracking_issue_to_reviewer["10976"], "anfernee")
        self.assertEqual(tracking_issue_to_reviewer["10276"], "anfernee")

if __name__ == '__main__':
    unittest.main()
