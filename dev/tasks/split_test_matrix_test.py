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
import os
import tempfile
import shutil
import subprocess
import json

class TestSplitTestMatrix(unittest.TestCase):

    def setUp(self):
        self.test_dir = tempfile.mkdtemp()
        self.repo_root = self.test_dir

        # Create mock config/crds structure
        self.crds_dir = os.path.join(self.repo_root, "config", "crds")
        os.makedirs(self.crds_dir, exist_ok=True)

        mock_services = ["compute", "storage", "alloydb", "bigquery", "kms"]
        for svc in mock_services:
            crd_file = os.path.join(self.crds_dir, f"{svc}.yaml")
            with open(crd_file, "w") as f:
                f.write(f"group: {svc}.cnrm.cloud.google.com\n")

        # Copy split-test-matrix script to temp test location
        self.script_path = os.path.join(self.repo_root, "dev", "tasks", "split-test-matrix")
        os.makedirs(os.path.dirname(self.script_path), exist_ok=True)

        actual_script = os.path.join(os.path.dirname(__file__), "split-test-matrix")
        shutil.copy(actual_script, self.script_path)
        os.chmod(self.script_path, 0o755)

    def tearDown(self):
        shutil.rmtree(self.test_dir)

    def run_split_script(self, changed_files):
        """Runs split-test-matrix with mocked CHANGED_FILES git output."""
        env = os.environ.copy()
        # Mock git command by writing a wrapper script or passing mock diff
        git_wrapper_dir = os.path.join(self.test_dir, "bin")
        os.makedirs(git_wrapper_dir, exist_ok=True)

        git_mock = os.path.join(git_wrapper_dir, "git")
        with open(git_mock, "w") as f:
            f.write(f'#!/bin/sh\nif [ "$1" = "rev-parse" ]; then echo "{self.repo_root}"; else echo "{changed_files}"; fi\n')
        os.chmod(git_mock, 0o755)

        env["PATH"] = f"{git_wrapper_dir}:{env['PATH']}"

        res = subprocess.run([self.script_path], cwd=self.repo_root, capture_output=True, text=True, env=env)
        self.assertEqual(res.returncode, 0, f"Script failed with error:\n{res.stderr}")

        output_vars = {}
        for line in res.stdout.splitlines():
            if "=" in line:
                k, v = line.split("=", 1)
                output_vars[k] = v
        return output_vars

    def test_priority_placement_with_subset(self):
        """Tests that changed services are correctly placed in priority list and unchanged in remaining list."""
        changed_files = (
            "apis/compute/v1beta1/computeinstance_types.go\n"
            "pkg/controller/direct/storage/storagebucket_controller.go\n"
            "pkg/test/resourcefixture/testdata/basic/alloydb/v1beta1/alloydbinstance/create.yaml"
        )

        outputs = self.run_split_script(changed_files)

        self.assertEqual(outputs["has_priority_tests"], "true")

        priority_fixtures = json.loads(outputs["priority_fixtures_json"])
        remaining_fixtures = json.loads(outputs["remaining_fixtures_json"])
        samples = json.loads(outputs["samples_json"])
        scenarios = json.loads(outputs["scenarios_json"])

        # Changed services MUST be in priority fixtures
        self.assertIn("compute", priority_fixtures)
        self.assertIn("storage", priority_fixtures)
        self.assertIn("alloydb", priority_fixtures)
        self.assertEqual(len(priority_fixtures), 3)

        # Unchanged services MUST be in remaining fixtures
        self.assertIn("bigquery", remaining_fixtures)
        self.assertIn("kms", remaining_fixtures)
        self.assertNotIn("compute", remaining_fixtures)
        self.assertNotIn("storage", remaining_fixtures)
        self.assertNotIn("alloydb", remaining_fixtures)
        self.assertEqual(len(remaining_fixtures), 4) # 2 unchanged + 2 static (reconcileintervalannotations, stateabsentinspec)

        # Total fixtures set equals all 5 services plus 2 static services
        all_fixtures = sorted(priority_fixtures + remaining_fixtures)
        self.assertEqual(all_fixtures, sorted(["alloydb", "bigquery", "compute", "kms", "storage", "reconcileintervalannotations", "stateabsentinspec"]))

        # 23 supported sample services present in samples
        expected_samples = ["alloydb", "apigateway", "apigee", "bigquery", "bigqueryconnection", "bigquerydatapolicy", "bigquerydatatransfer", "cloudbuild", "cloudidentity", "compute", "dataflow", "dataproc", "gkehub", "iam", "kms", "orgpolicy", "redis", "secretmanager", "spanner", "sql", "storage", "unclassified", "workstations"]
        self.assertEqual(sorted(samples), sorted(expected_samples))

        # Scenario suites present
        self.assertEqual(scenarios, ["acquisition", "gkehubfeaturemembership", "powertool", "unclassified"])

    def test_no_changes_placement(self):
        """Tests that when no files are changed, priority is false and all services are in remaining list."""
        changed_files = ""

        outputs = self.run_split_script(changed_files)

        self.assertEqual(outputs["has_priority_tests"], "false")

        priority_fixtures = json.loads(outputs["priority_fixtures_json"])
        remaining_fixtures = json.loads(outputs["remaining_fixtures_json"])

        # Dummy 'none' passed for priority when empty
        self.assertEqual(priority_fixtures, ["none"])

        # All 5 services plus 2 static services in remaining fixtures
        self.assertEqual(sorted(remaining_fixtures), sorted(["alloydb", "bigquery", "compute", "kms", "storage", "reconcileintervalannotations", "stateabsentinspec"]))

    def test_sample_list_independence_from_all_crd_services(self):
        """Tests that samples list is NOT naively populated with all CRD services (catches phantom sample job regression)."""
        # Add 10 additional fake CRD services to config/crds
        for idx in range(10):
            crd_file = os.path.join(self.crds_dir, f"fakecrdservice{idx}.yaml")
            with open(crd_file, "w") as f:
                f.write(f"group: fakecrdservice{idx}.cnrm.cloud.google.com\n")

        outputs = self.run_split_script("")

        remaining_fixtures = json.loads(outputs["remaining_fixtures_json"])
        samples = json.loads(outputs["samples_json"])

        # Fixture services MUST include the 10 fake CRD services (5 original + 10 fake + 2 static = 17 total)
        self.assertEqual(len(remaining_fixtures), 17)
        for idx in range(10):
            self.assertIn(f"fakecrdservice{idx}", remaining_fixtures)

        # Samples list MUST NOT contain any of the fake CRD services and MUST stay at the exact 23 supported sample services
        expected_samples = ["alloydb", "apigateway", "apigee", "bigquery", "bigqueryconnection", "bigquerydatapolicy", "bigquerydatatransfer", "cloudbuild", "cloudidentity", "compute", "dataflow", "dataproc", "gkehub", "iam", "kms", "orgpolicy", "redis", "secretmanager", "spanner", "sql", "storage", "unclassified", "workstations"]
        self.assertEqual(len(samples), 23)
        self.assertEqual(sorted(samples), sorted(expected_samples))
        self.assertNotEqual(sorted(samples), sorted(remaining_fixtures))

if __name__ == "__main__":
    unittest.main()
