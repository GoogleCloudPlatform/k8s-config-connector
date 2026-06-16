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
Scans Go source files across apis/, pkg/apis/, and pkg/clients/ and updates
standalone occurrences of 'GCP' to 'Google Cloud' within Go comments and
docstrings (handling both // and /* ... */ block comments).
"""

import os
import re

def process_file(filepath):
    with open(filepath, 'r') as f:
        content = f.read()
    
    def replace_in_comment(match):
        comment_text = match.group(0)
        # Skip compiler/generator directives
        if comment_text.startswith('// +') or comment_text.startswith('//go:'):
            return comment_text
        # Replace standalone GCP
        return re.sub(r'\bGCP\b', 'Google Cloud', comment_text)
        
    # Match either /* ... */ or // ...
    new_content = re.sub(r'/\*[\s\S]*?\*/|//.*', replace_in_comment, content)
    
    if new_content != content:
        print(f"Formatted terminology in: {filepath}")
        with open(filepath, 'w') as f:
            f.write(new_content)

def main():
    repo_root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
    dirs_to_scan = [
        os.path.join(repo_root, 'apis'),
        os.path.join(repo_root, 'pkg/apis'),
        os.path.join(repo_root, 'pkg/clients/generated/apis'),
    ]
    for root_dir in dirs_to_scan:
        for folder, _, files in os.walk(root_dir):
            for file in files:
                if file.endswith('.go'):
                    process_file(os.path.join(folder, file))

if __name__ == '__main__':
    main()
