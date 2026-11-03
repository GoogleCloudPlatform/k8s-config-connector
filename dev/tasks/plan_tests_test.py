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

"""Tests for dev/tasks/plan-tests. Run: python3 dev/tasks/plan_tests_test.py"""

import importlib.machinery
import importlib.util
import os
import shutil
import sys
import tempfile
import textwrap
import unittest

sys.dont_write_bytecode = True
_HERE = os.path.dirname(os.path.abspath(__file__))
_loader = importlib.machinery.SourceFileLoader("plan_tests", os.path.join(_HERE, "plan-tests"))
_spec = importlib.util.spec_from_loader("plan_tests", _loader)
pt = importlib.util.module_from_spec(_spec)
_loader.exec_module(pt)

M = pt.MODULE

FILES = {
    # Services with their API groups.
    "apis/vectorsearch/v1alpha1/types.go": '''
        // +groupName=vectorsearch.cnrm.cloud.google.com
        package v1alpha1
        import (
            refs "%sapis/refs/v1beta1"
        )
    ''' % M,
    "apis/kms/v1beta1/types.go": '''
        // +groupName=kms.cnrm.cloud.google.com
        package v1beta1
    ''',
    "apis/securitycenter/v1alpha1/types.go": '''
        // +groupName=securitycenter.cnrm.cloud.google.com
        package v1alpha1
    ''',
    "apis/iam/v1beta1/types.go": '''
        // +groupName=iam.cnrm.cloud.google.com
        package v1beta1
    ''',
    "apis/refs/v1beta1/refs.go": "package v1beta1\n",
    # Direct controllers.
    "pkg/controller/direct/directbase/base.go": "package directbase\n",
    "pkg/controller/direct/vectorsearch/controller.go": '''
        package vectorsearch
        import (
            krm "%sapis/vectorsearch/v1alpha1"
            "%spkg/controller/direct/directbase"
        )
    ''' % (M, M),
    "pkg/controller/direct/vectorsearch/fuzzer.go": '''
        package vectorsearch
        func init() { fuzztesting.RegisterKRMFuzzer(nil) }
    ''',
    "pkg/controller/direct/securitycenter/controller.go": '''
        package securitycenter
        import (
            krm "%sapis/securitycenter/v1alpha1"
            kms "%sapis/kms/v1beta1"
            "%spkg/controller/direct/directbase"
        )
    ''' % (M, M, M),
    "pkg/controller/direct/kms/controller.go": '''
        package kms
        import krm "%sapis/kms/v1beta1"
    ''' % M,
    "pkg/controller/direct/kms/fuzzer.go": '''
        package kms
        func init() { fuzztesting.RegisterKRMFuzzer(nil) }
    ''',
    # Aggregators and shared code.
    "pkg/controller/direct/register/register.go": '''
        package register
        import (
            _ "%spkg/controller/direct/vectorsearch"
            _ "%spkg/controller/direct/securitycenter"
            _ "%spkg/controller/direct/kms"
        )
    ''' % (M, M, M),
    "tests/e2e/e2e_test.go": '''
        package e2e
        import (
            _ "%spkg/controller/direct/register"
            _ "%sapis/kms/v1beta1"
        )
    ''' % (M, M),
    "pkg/controller/iam/iam.go": '''
        package iam
        import "%sapis/iam/v1beta1"
    ''' % M,
    "pkg/controller/resourceconfig/static_config.go": "package resourceconfig\n",
    "tests/apichecks/crds_test.go": "package apichecks\n",
    # CRDs (source of the fixture service list).
    "config/crds/resources/vectorsearch.yaml": "spec:\n  group: vectorsearch.cnrm.cloud.google.com\n",
    "config/crds/resources/kms.yaml": "spec:\n  group: kms.cnrm.cloud.google.com\n",
    "config/crds/resources/securitycenter.yaml": "spec:\n  group: securitycenter.cnrm.cloud.google.com\n",
    "config/crds/resources/iam.yaml": "spec:\n  group: iam.cnrm.cloud.google.com\n",
    "config/crds/resources/compute.yaml": "spec:\n  group: compute.cnrm.cloud.google.com\n",
    # Fixtures outside testdata/basic are attributed via create.yaml.
    "pkg/test/resourcefixture/testdata/labels/kmskeyring-labels/create.yaml":
        "apiVersion: kms.cnrm.cloud.google.com/v1beta1\nkind: KMSKeyRing\n",
    "mockgcp/mockkms/service.go": "package mockkms\n",
    "mockgcp/common/common.go": "package common\n",
}


class PlanTestsTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.root = tempfile.mkdtemp()
        for rel, content in FILES.items():
            path = os.path.join(cls.root, rel)
            os.makedirs(os.path.dirname(path), exist_ok=True)
            with open(path, "w") as fh:
                fh.write(textwrap.dedent(content))

    @classmethod
    def tearDownClass(cls):
        shutil.rmtree(cls.root)

    def plan(self, files, event="pull_request", labels=(), diff=None):
        repo = pt.Repo(self.root, diff_lines=(lambda path: diff.get(path)) if diff else None)
        return pt.make_plan(repo, files, event, labels)

    def assertFull(self, plan, reason_part=""):
        self.assertEqual(plan["scope"], "full", plan["reason"])
        self.assertIn(reason_part, plan["reason"])
        self.assertTrue(plan["run_global"])
        self.assertEqual(plan["fuzz_services"], [])          # all fuzzers
        self.assertEqual(plan["unit_packages"], [])          # all packages
        self.assertEqual(len(plan["unit_shards"]), 4)
        self.assertEqual(len(plan["fuzz_shards"]), 8)

    # -- none / lint-only ---------------------------------------------------

    def test_noop_paths(self):
        plan = self.plan([".gemini/skills/x/SKILL.md", ".agents/graphs/g.yaml", "experiments/foo/bar.go"])
        self.assertEqual(plan["scope"], "none")
        self.assertFalse(plan["run_ci"])
        outputs = pt.plan_outputs(plan)
        self.assertEqual(outputs["run_ci"], "false")
        self.assertEqual(outputs["priority_fixtures_json"], '["none"]')

    def test_lint_only_paths(self):
        plan = self.plan(["docs/foo.md", "README.md", "dev/tools/controllerbuilder/main.go"])
        self.assertEqual(plan["scope"], "scoped")
        self.assertTrue(plan["run_ci"])
        self.assertFalse(plan["run_global"])
        self.assertEqual(plan["priority_fixtures"], [])
        self.assertEqual(plan["unit_shards"], [])

    # -- scoped -------------------------------------------------------------

    def test_service_change_is_scoped(self):
        plan = self.plan([
            "pkg/controller/direct/vectorsearch/controller.go",
            "apis/vectorsearch/v1alpha1/types.go",
            "pkg/test/resourcefixture/testdata/basic/vectorsearch/v1alpha1/x/create.yaml",
            ".gemini/journals/vectorsearch.md",
        ])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertFalse(plan["run_global"])
        self.assertEqual(plan["priority_fixtures"], ["vectorsearch"])
        self.assertEqual(plan["remaining_fixtures"], [])
        self.assertEqual(plan["fuzz_services"], ["vectorsearch"])
        self.assertEqual(plan["fuzz_shards"], ["1-of-1"])
        self.assertIn("./pkg/controller/direct/vectorsearch", plan["unit_packages"])
        self.assertIn("./apis/vectorsearch/v1alpha1", plan["unit_packages"])
        self.assertIn(pt.APICHECKS_UNIT, plan["unit_packages"])
        self.assertIn("./pkg/test/resourcefixture/...", plan["unit_packages"])
        self.assertEqual(plan["unit_shards"], ["1-of-1"])

    def test_cross_service_import_is_followed(self):
        plan = self.plan(["apis/kms/v1beta1/types.go"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        # securitycenter imports apis/kms, so its fixtures run too.
        self.assertEqual(plan["priority_fixtures"], ["kms", "securitycenter"])
        # Only kms registers fuzzers in this fake repo.
        self.assertEqual(plan["fuzz_services"], ["kms"])
        self.assertIn("./pkg/controller/direct/securitycenter", plan["unit_packages"])

    def test_aggregators_do_not_escalate(self):
        # register and tests/e2e import every service; reaching them is fine.
        plan = self.plan(["pkg/controller/direct/vectorsearch/controller.go"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])

    def test_samples_and_direct_iam_follow_groups(self):
        plan = self.plan(["pkg/test/resourcefixture/testdata/basic/kms/v1beta1/x/create.yaml"])
        self.assertEqual(plan["samples"], ["kms"])
        self.assertFalse(plan["run_direct_iam"])

    def test_fixture_outside_basic_uses_create_yaml(self):
        plan = self.plan(["pkg/test/resourcefixture/testdata/labels/kmskeyring-labels/create.yaml"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["priority_fixtures"], ["kms"])

    def test_static_config_scoped_by_diff(self):
        f = pt.STATIC_CONFIG
        diff = {f: ['{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyHandle"}: {DefaultController: ...},']}
        plan = self.plan([f], diff=diff)
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["priority_fixtures"], ["kms"])
        self.assertIn("./pkg/controller/resourceconfig/...", plan["unit_packages"])

    def test_static_config_without_diff_is_full(self):
        self.assertFull(self.plan([pt.STATIC_CONFIG]), "cannot read the diff")

    def test_mock_change_is_scoped(self):
        plan = self.plan(["mockgcp/mockkms/service.go"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["mock_packages"], ["./mockkms/..."])
        self.assertTrue(plan["run_mock"])
        self.assertEqual(plan["priority_fixtures"], ["kms"])

    def test_crd_change_is_scoped(self):
        plan = self.plan([
            "config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_kmskeyrings.kms.cnrm.cloud.google.com.yaml"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["priority_fixtures"], ["kms"])
        self.assertIn(pt.APICHECKS_UNIT, plan["unit_packages"])

    def test_leaf_test_package_is_scoped(self):
        plan = self.plan(["tests/apichecks/crds_test.go", "tests/apichecks/testdata/exceptions/x.txt"])
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["unit_packages"], [pt.APICHECKS_UNIT])
        self.assertEqual(plan["priority_fixtures"], [])

    def test_sharded_service_targets(self):
        self.assertEqual(pt.shard_targets("compute"),
                         ["compute-1-of-4", "compute-2-of-4", "compute-3-of-4", "compute-4-of-4"])
        self.assertEqual(pt.shard_targets("kms"), ["kms"])

    def test_register_import_change_is_scoped(self):
        f = "pkg/controller/direct/register/register.go"
        diff = {f: ['\t_ "%spkg/controller/direct/vectorsearch"' % M]}
        plan = self.plan([f], diff=diff)
        self.assertEqual(plan["scope"], "scoped", plan["reason"])
        self.assertEqual(plan["priority_fixtures"], ["vectorsearch"])
        self.assertEqual(plan["fuzz_services"], ["vectorsearch"])

    # -- full ---------------------------------------------------------------

    def test_direct_change_to_aggregator_is_full(self):
        # The walk stops at aggregators it reaches, but editing one directly
        # (e.g. the e2e harness that runs TestPauseInSeries) is full.
        self.assertFull(self.plan(["tests/e2e/e2e_test.go"]), "aggregator package tests/e2e")

    def test_register_non_import_change_is_full(self):
        f = "pkg/controller/direct/register/register.go"
        diff = {f: ["func init() { doSomething() }"]}
        self.assertFull(self.plan([f], diff=diff), "other than service imports")

    def test_register_without_diff_is_full(self):
        self.assertFull(self.plan(["pkg/controller/direct/register/register.go"]), "cannot read the diff")

    def test_shared_package_is_full(self):
        self.assertFull(self.plan(["pkg/controller/direct/directbase/base.go"]), "shared package")

    def test_service_imported_by_shared_code_is_full(self):
        # pkg/controller/iam imports apis/iam.
        self.assertFull(self.plan(["apis/iam/v1beta1/types.go"]), "pkg/controller/iam")

    def test_unknown_path_is_full(self):
        self.assertFull(self.plan([".github/workflows/x.yaml"]), "not attributable")

    def test_shared_mock_is_full(self):
        self.assertFull(self.plan(["mockgcp/common/common.go"]), "shared mockgcp")

    def test_merge_group_is_full(self):
        plan = self.plan(["pkg/controller/direct/vectorsearch/controller.go"], event="merge_group")
        self.assertFull(plan, "merge_group")
        # Changed services still run first.
        self.assertEqual(plan["priority_fixtures"], ["vectorsearch"])
        self.assertNotIn("vectorsearch", plan["remaining_fixtures"])
        self.assertIn("compute-1-of-4", plan["remaining_fixtures"])

    def test_push_is_full(self):
        self.assertFull(self.plan(["README.md"], event="push"), "push")

    def test_full_label(self):
        self.assertFull(self.plan(["pkg/controller/direct/vectorsearch/controller.go"], labels=["ci/full"]),
                        "ci/full")

    def test_unknown_changes_are_full(self):
        self.assertFull(self.plan(None), "unknown")
        self.assertFull(self.plan([]), "no changed files")

    def test_too_many_services_is_full(self):
        old = pt.MAX_SCOPED_SERVICES
        pt.MAX_SCOPED_SERVICES = 1
        try:
            self.assertFull(self.plan(["apis/kms/v1beta1/types.go"]), "services affected")
        finally:
            pt.MAX_SCOPED_SERVICES = old


if __name__ == "__main__":
    unittest.main()
