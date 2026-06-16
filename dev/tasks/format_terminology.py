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

"""
Scans Go source files in apis/ and pkg/apis/ and updates standalone occurrences
of 'GCP' to 'Google Cloud' within Go comments and docstrings.
"""

import os
import re

def process_file(filepath):
    with open(filepath, 'r') as f:
        lines = f.readlines()
    
    modified = False
    new_lines = []
    
    for line in lines:
        # Skip compiler/generator directives
        if line.lstrip().startswith('// +') or line.lstrip().startswith('//go:'):
            new_lines.append(line)
            continue
            
        new_line = line
        # If line is entirely a comment
        if line.lstrip().startswith('//'):
            new_line = re.sub(r'\bGCP\b', 'Google Cloud', line)
        # Or if line has an end-of-line comment (not containing string literals/URLs)
        elif '//' in line and not '"' in line and not '`' in line:
            parts = line.split('//', 1)
            new_comment = re.sub(r'\bGCP\b', 'Google Cloud', '//' + parts[1])
            new_line = parts[0] + new_comment
            
        if new_line != line:
            modified = True
        new_lines.append(new_line)
        
    if modified:
        print(f"Formatted terminology in: {filepath}")
        with open(filepath, 'w') as f:
            f.writelines(new_lines)

def main():
    repo_root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
    dirs_to_scan = [os.path.join(repo_root, 'apis'), os.path.join(repo_root, 'pkg/apis')]
    for root_dir in dirs_to_scan:
        for folder, _, files in os.walk(root_dir):
            for file in files:
                if file.endswith('.go'):
                    process_file(os.path.join(folder, file))

if __name__ == '__main__':
    main()
