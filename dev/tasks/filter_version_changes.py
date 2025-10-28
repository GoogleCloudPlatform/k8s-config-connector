#!/usr/bin/env python3
# Copyright 2025 Google LLC
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

import re
import sys

def filter_version_upgrades(input_file, f_out):
    with open(input_file, 'r') as f_in:
        diff_block = []
        in_diff_block = False
        for line in f_in:
            if line.startswith('diff --git'):
                if in_diff_block:
                    process_diff_block(diff_block, f_out)
                diff_block = [line]
                in_diff_block = True
            elif in_diff_block:
                diff_block.append(line)
        if in_diff_block:
            process_diff_block(diff_block, f_out)

def process_diff_block(diff_block, f_out):
    # Heuristic: if the only changes are to version and copyright, remove the block
    other_changes = []
    for line in diff_block:
        if (line.startswith('+') or line.startswith('-')) and not (line.startswith('+++') or line.startswith('---')):
            if not re.search(r'cnrm\.cloud\.google\.com/version:|Copyright', line):
                other_changes.append(line)

    if len(other_changes) > 0:
        f_out.writelines(diff_block)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: filter_version_changes.py <input_file>", file=sys.stderr)
        sys.exit(1)
    input_file = sys.argv[1]
    filter_version_upgrades(input_file, sys.stdout)