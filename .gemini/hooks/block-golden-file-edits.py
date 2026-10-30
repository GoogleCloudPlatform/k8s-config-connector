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

"""PreToolUse hook blocking direct hand-edits to E2E golden files."""

import json
import re
import sys

BLOCK_REASON = (
    "BLOCKED: Direct hand-edits to golden files "
    "(_http.log, _http_mock.log, _http_report.md, *.golden.yaml) are forbidden. "
    "You must generate golden files by running hack/record-gcp (for RealGCP) "
    "or hack/compare-mock (for MockGCP)."
)

FILE_EDIT_TOOLS = {
    "replace_file_content",
    "multi_replace_file_content",
    "write_to_file",
    "edit_file",
    "Write",
    "Edit",
    "MultiEdit",
}

SHELL_TOOLS = {
    "run_command",
    "Bash",
    "bash",
}

GOLDEN_SUFFIXES = (
    "_http.log",
    "_http_mock.log",
    "_http_report.md",
    ".golden.yaml",
)


def is_protected_golden_path(path: str) -> bool:
    if not path:
        return False
    norm = path.replace("\\", "/")
    if "pkg/test/resourcefixture/testdata/" not in norm:
        return False
    return norm.endswith(GOLDEN_SUFFIXES)


def is_blocked_shell_command(cmd: str) -> bool:
    if not cmd:
        return False
    # Check if command mentions any golden file suffix
    if not any(s in cmd for s in GOLDEN_SUFFIXES):
        return False

    if "pkg/test/resourcefixture/testdata" not in cmd:
        return False

    stripped = cmd.strip()
    allowed_prefixes = (
        "hack/record-gcp",
        "./hack/record-gcp",
        "hack/compare-mock",
        "./hack/compare-mock",
        "go test",
        "go run ./cmd/httplog-report",
        "git ",
    )
    if stripped.startswith(allowed_prefixes):
        return False

    # Check for shell write/mutation patterns targeting golden files
    write_patterns = [
        r"\bsed\b.*\s-[a-zA-Z]*i",
        r"\bperl\b.*\s-[a-zA-Z]*i",
        r"\bpython3?\b.*(?:-c|<<)",
        r"\btee\b",
        r"(?:cat|echo|printf)\b[^;|]*>>?",
        r">\s*[^\s;|&]*(?:_http\.log|_http_mock\.log|_http_report\.md|\.golden\.yaml)",
        r">>\s*[^\s;|&]*(?:_http\.log|_http_mock\.log|_http_report\.md|\.golden\.yaml)",
    ]
    for pat in write_patterns:
        if re.search(pat, cmd, re.DOTALL):
            return True
    return False


def block() -> int:
    print(json.dumps({"decision": "block", "reason": BLOCK_REASON}))
    sys.stderr.write(BLOCK_REASON + "\n")
    return 2


def main() -> int:
    raw = sys.stdin.read()
    if not raw.strip():
        return 0
    try:
        payload = json.loads(raw)
    except Exception:
        return 0

    tool_name = payload.get("tool_name") or payload.get("name") or ""
    tool_input = (
        payload.get("tool_input")
        or payload.get("arguments")
        or payload.get("input")
        or payload
    )
    if not isinstance(tool_input, dict):
        tool_input = {}

    if tool_name in FILE_EDIT_TOOLS:
        target = (
            tool_input.get("TargetFile")
            or tool_input.get("file_path")
            or tool_input.get("path")
            or tool_input.get("target_file")
            or ""
        )
        if is_protected_golden_path(str(target)):
            return block()

    if tool_name in SHELL_TOOLS:
        cmd = (
            tool_input.get("CommandLine")
            or tool_input.get("command")
            or tool_input.get("cmd")
            or ""
        )
        if is_blocked_shell_command(str(cmd)):
            return block()

    return 0


if __name__ == "__main__":
    sys.exit(main())
