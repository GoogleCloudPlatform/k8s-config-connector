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

import os
import subprocess
import sys
import tempfile
import unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import changed_kinds as ck  # noqa: E402

INDEX = {"foowidget": "FooWidget", "foogadget": "FooGadget"}


class ClassifyFileTest(unittest.TestCase):
    def test_roles(self):
        cases = {
            "pkg/controller/direct/foo/foowidget_controller.go": ("foowidget", "controller"),
            "pkg/controller/direct/foo/foowidget_mapper.go": ("foowidget", "mapper"),
            "pkg/controller/direct/foo/foowidget_fuzzer.generated.go": ("foowidget", "tests"),
            "pkg/controller/direct/foo/mapper.generated.go": (None, "generated"),
            "pkg/controller/direct/foo/foogadget/adapter.go": ("foogadget", "controller"),
            "apis/foo/v1alpha1/foowidget_types.go": ("foowidget", "types"),
            "apis/foo/v1alpha1/foowidget_identity.go": ("foowidget", "identity"),
            "apis/foo/v1alpha1/zz_generated.deepcopy.go": (None, "other"),
            "pkg/test/resourcefixture/testdata/basic/foo/v1alpha1/foowidget/foowidget-minimal/create.yaml":
                ("foowidget", "tests"),
            "mockgcp/mockfoo/widget.go": (None, "other"),
        }
        for path, want in cases.items():
            self.assertEqual(ck.classify_file(path, INDEX), want, path)


TYPES = '''package v1alpha1

type FooWidgetSpec struct {{
	Name *string `json:"name,omitempty"`
{extra}}}

type FooWidgetStatus struct {{
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
{status}}}

type FooWidget struct {{
	Spec   FooWidgetSpec   `json:"spec,omitempty"`
	Status FooWidgetStatus `json:"status,omitempty"`
}}
'''

STATIC = '''package resourceconfig
var X = map[x]y{{
{entries}}}
'''


def static(**defaults):
    rows = []
    for kind, ctrl in defaults.items():
        rows.append(f'\t{{Group: "foo.cnrm.cloud.google.com", Kind: "{kind}"}}: '
                    f'{{DefaultController: k8s.ReconcilerType{ctrl}, SupportedControllers: '
                    f'[]k8s.ReconcilerType{{k8s.ReconcilerType{ctrl}}}}},\n')
    return STATIC.format(entries="".join(rows))


class ScoutTest(unittest.TestCase):
    def setUp(self):
        self.repo = tempfile.mkdtemp()
        self.git("init", "-q", "-b", "master")
        self.git("config", "user.email", "t@example.com")
        self.git("config", "user.name", "t")

    def git(self, *args):
        return subprocess.run(["git", "-C", self.repo] + list(args), check=True,
                              capture_output=True, text=True).stdout.strip()

    def write(self, path, body):
        full = os.path.join(self.repo, path)
        os.makedirs(os.path.dirname(full), exist_ok=True)
        with open(full, "w") as f:
            f.write(body)

    def commit(self, msg):
        self.git("add", "-A")
        self.git("commit", "-qm", msg)
        return self.git("rev-parse", "HEAD")

    def test_classes(self):
        # Base: FooWidget has types and a direct controller; FooGadget has types only.
        self.write("apis/foo/v1alpha1/foowidget_types.go", TYPES.format(extra="", status=""))
        self.write("apis/foo/v1alpha1/foogadget_types.go", "type FooGadget struct {}\n")
        self.write("pkg/controller/direct/foo/foowidget_controller.go", "package foo\n")
        self.write(ck.STATIC_CONFIG, static(FooWidget="Direct"))
        base = self.commit("base")

        # A spec field and a status field added to FooWidget.
        self.write("apis/foo/v1alpha1/foowidget_types.go", TYPES.format(
            extra='\tSize *int64 `json:"size,omitempty"`\n',
            status='\tState *string `json:"state,omitempty"`\n'))
        self.commit("FooWidget: add size (#11)")
        # FooGadget gets a direct controller.
        self.write("pkg/controller/direct/foo/foogadget_controller.go", "package foo\n")
        self.write(ck.STATIC_CONFIG, static(FooWidget="Direct", FooGadget="Direct"))
        self.commit("FooGadget: direct controller (#12)")
        head = self.git("rev-parse", "HEAD")

        report = ck.scout(self.repo, base, head, mechanical_threshold=8, sample=1)
        got = {r["kind"]: r for r in report["kinds"]}
        self.assertEqual(got["FooGadget"]["class"], "new")
        self.assertEqual(got["FooGadget"]["prs"], [12])
        self.assertEqual(got["FooWidget"]["class"], "fields")
        self.assertEqual(got["FooWidget"]["fields"], {"added": ["size"]})  # status field ignored
        self.assertTrue(all(r["selected"] for r in report["kinds"]))

    def test_mechanical_and_types(self):
        kinds = [f"Foo{c}" for c in "ABCDEFGHIJ"]
        for k in kinds:
            low = k.lower()
            self.write(f"apis/foo/v1alpha1/{low}_types.go", f"type {k} struct {{}}\n")
            self.write(f"pkg/controller/direct/foo/{low}_controller.go", "package foo\n")
            self.write(f"pkg/test/resourcefixture/testdata/basic/foo/v1alpha1/{low}/{low}-minimal/create.yaml", "x\n")
        self.write("apis/foo/v1alpha1/foogadget_types.go", "type FooGadget struct {}\n")
        self.write(ck.STATIC_CONFIG, static(**{k: "Direct" for k in kinds}))
        base = self.commit("base")

        # One refactor touching all 10 controllers, and a types-only change to FooGadget.
        for k in kinds:
            self.write(f"pkg/controller/direct/foo/{k.lower()}_controller.go", "package foo\n// refactor\n")
        self.commit("refactor all controllers (#20)")
        self.write("apis/foo/v1alpha1/foogadget_types.go", "type FooGadget struct {}\n// more\n")
        head = self.commit("FooGadget types (#21)")

        report = ck.scout(self.repo, base, head, mechanical_threshold=8, sample=2)
        got = {r["kind"]: r for r in report["kinds"]}
        self.assertEqual(len(report["mechanical_commits"]), 1)
        mech = [r for r in report["kinds"] if r["class"] == "mechanical"]
        self.assertEqual(len(mech), 10)
        self.assertEqual(sum(r["selected"] for r in mech), 2)  # the sample
        self.assertEqual(got["FooGadget"]["class"], "types")
        self.assertFalse(got["FooGadget"]["selected"])


if __name__ == "__main__":
    unittest.main()
