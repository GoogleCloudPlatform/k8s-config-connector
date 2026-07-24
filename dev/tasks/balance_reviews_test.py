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
from dev.tasks.balance_reviews import (
    extract_issue_ids,
    audit_workloads_and_workflows,
    filter_candidates,
    run_assignment,
    TEAM_MEMBERS,
)


class TestBalanceReviews(unittest.TestCase):

    def test_extract_issue_ids(self):
        title = "Migrate KRM types to direct for ComputeTargetVPNGateway (#11257)"
        body = "Fixes #11257\nWorkflow: #10976 and #10276."
        ids = extract_issue_ids(title, body)
        self.assertEqual(ids, ["10276", "10976", "11257"])

    def test_audit_workloads_and_workflows(self):
        prs = [
            {
                "number": 1,
                "title": "PR 1 #101",
                "body": "Some body",
                "requested_reviewers": [{"login": "barney-s"}]
            },
            {
                "number": 2,
                "title": "PR 2",
                "body": "Fixes #102",
                "requested_reviewers": [{"login": "anfernee"}, {"login": "external-user"}]
            }
        ]
        workload, tracking_map = audit_workloads_and_workflows(prs)
        self.assertEqual(workload["barney-s"], 1)
        self.assertEqual(workload["anfernee"], 1)
        self.assertEqual(workload["acpana"], 0)
        self.assertEqual(tracking_map["101"], "barney-s")
        self.assertEqual(tracking_map["102"], "anfernee")

    def test_filter_candidates(self):
        prs = [
            {
                "number": 1,
                "labels": [{"name": "ready-for-human"}],
                "requested_reviewers": []
            },
            {
                "number": 2,
                "labels": [{"name": "ready-for-human"}],
                "requested_reviewers": [{"login": "acpana"}]
            },
            {
                "number": 3,
                "labels": [{"name": "not-ready"}],
                "requested_reviewers": []
            },
            {
                "number": 4,
                "draft": True,
                "labels": [{"name": "ready-for-human"}],
                "requested_reviewers": []
            }
        ]
        candidates = filter_candidates(prs)
        self.assertEqual(len(candidates), 1)
        self.assertEqual(candidates[0]["number"], 1)

    def test_run_assignment_priority_workflow_affinity(self):
        candidates = [
            {
                "number": 10,
                "title": "Candidate referencing #100",
                "body": ""
            }
        ]
        tracking_issue_to_reviewer = {"100": "gemmahou"}
        workload = {m: 5 for m in TEAM_MEMBERS}
        workload["gemmahou"] = 3

        assignments = run_assignment(candidates, workload, tracking_issue_to_reviewer)
        self.assertEqual(assignments, [(10, "gemmahou")])
        self.assertEqual(workload["gemmahou"], 4)

    def test_run_assignment_priority_underloaded(self):
        candidates = [
            {
                "number": 10,
                "title": "Candidate with no refs",
                "body": ""
            }
        ]
        tracking_issue_to_reviewer = {}
        workload = {m: 6 for m in TEAM_MEMBERS}
        workload["anhdle-sso"] = 3

        assignments = run_assignment(candidates, workload, tracking_issue_to_reviewer)
        self.assertEqual(assignments, [(10, "anhdle-sso")])
        self.assertEqual(workload["anhdle-sso"], 4)

    def test_run_assignment_priority_capacity_absorption(self):
        candidates = [
            {
                "number": 10,
                "title": "Candidate with no refs",
                "body": ""
            }
        ]
        tracking_issue_to_reviewer = {}
        workload = {m: 6 for m in TEAM_MEMBERS}
        workload["acpana"] = 5

        assignments = run_assignment(candidates, workload, tracking_issue_to_reviewer)
        self.assertEqual(assignments, [(10, "acpana")])
        self.assertEqual(workload["acpana"], 6)

    def test_run_assignment_ceiling_reached(self):
        candidates = [
            {
                "number": 10,
                "title": "Candidate with no refs",
                "body": ""
            }
        ]
        tracking_issue_to_reviewer = {}
        workload = {m: 10 for m in TEAM_MEMBERS}

        assignments = run_assignment(candidates, workload, tracking_issue_to_reviewer)
        self.assertEqual(assignments, [])


if __name__ == "__main__":
    unittest.main()
