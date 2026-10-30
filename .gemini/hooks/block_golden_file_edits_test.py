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

"""Unit tests for .gemini/hooks/block-golden-file-edits.py."""

import json
import os
import subprocess
import sys
import unittest

HOOK_SCRIPT = os.path.join(os.path.dirname(__file__), "block-golden-file-edits.py")


def run_hook(payload: dict) -> tuple[int, str, str]:
    proc = subprocess.run(
        [sys.executable, HOOK_SCRIPT],
        input=json.dumps(payload),
        text=True,
        capture_output=True,
        check=False,
    )
    return proc.returncode, proc.stdout, proc.stderr


class BlockGoldenFileEditsTest(unittest.TestCase):

    def assertBlocked(self, payload: dict):
        code, stdout, stderr = run_hook(payload)
        self.assertEqual(code, 2, f"Expected exit code 2 for {payload}, got {code} (stdout={stdout}, stderr={stderr})")
        out = json.loads(stdout.strip())
        self.assertEqual(out.get("decision"), "block")
        self.assertIn("BLOCKED: Direct hand-edits to golden files", out.get("reason", ""))

    def assertAllowed(self, payload: dict):
        code, stdout, stderr = run_hook(payload)
        self.assertEqual(code, 0, f"Expected exit code 0 for {payload}, got {code} (stdout={stdout}, stderr={stderr})")

    def test_file_edit_tools_blocked_on_golden_files(self):
        for tool in ["replace_file_content", "multi_replace_file_content", "write_to_file", "edit_file", "Write", "Edit", "MultiEdit"]:
            for suffix in ["_http.log", "_http_mock.log", "_http_report.md", "_generated_object_bar.golden.yaml"]:
                self.assertBlocked({
                    "tool_name": tool,
                    "tool_input": {"TargetFile": "pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/" + suffix},
                })
                self.assertBlocked({
                    "tool_name": tool,
                    "tool_input": {"file_path": "/abs/repo/pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/" + suffix},
                })

    def test_file_edit_tools_allowed_on_create_and_update_yaml(self):
        for fname in ["create.yaml", "update.yaml", "dependencies.yaml"]:
            self.assertAllowed({
                "tool_name": "replace_file_content",
                "tool_input": {"TargetFile": "pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/" + fname},
            })

    def test_shell_commands_blocked_on_hand_edits(self):
        blocked_cmds = [
            "sed -i 's/foo/bar/g' pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http.log",
            "python3 -c 'open(\"pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http.log\", \"w\")'",
            "cat > pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http_report.md <<EOF\nfoo\nEOF",
            "echo hello > pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http_mock.log",
            "tee pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_generated_object_bar.golden.yaml",
        ]
        for cmd in blocked_cmds:
            self.assertBlocked({
                "tool_name": "run_command",
                "tool_input": {"CommandLine": cmd},
            })
            self.assertBlocked({
                "tool_name": "Bash",
                "tool_input": {"command": cmd},
            })

    def test_shell_commands_allowed_for_recording_and_readonly(self):
        allowed_cmds = [
            "hack/record-gcp pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/",
            "hack/compare-mock pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/",
            "go run ./cmd/httplog-report --fixture pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/ --write",
            "go test ./pkg/test/httplogreport/...",
            "git diff pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http.log",
            "git checkout -- pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http.log",
            "cat pkg/test/resourcefixture/testdata/basic/foo/v1beta1/bar/_http.log",
        ]
        for cmd in allowed_cmds:
            self.assertAllowed({
                "tool_name": "run_command",
                "tool_input": {"CommandLine": cmd},
            })


if __name__ == "__main__":
    unittest.main()
