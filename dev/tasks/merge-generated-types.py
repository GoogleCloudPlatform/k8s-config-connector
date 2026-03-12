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

import sys
import re

def merge_generated_files(file1, file2, outfile):
    with open(file1, 'r') as f:
        content1 = f.read()
    with open(file2, 'r') as f:
        content2 = f.read()

    # Keep header from file1, but merge resource list
    resources1 = re.findall(r'// resource: (.*)', content1)
    resources2 = re.findall(r'// resource: (.*)', content2)
    all_resources = sorted(list(set(resources1 + resources2)))

    pkg_decl = '\npackage v1beta1\n'
    pkg_idx1 = content1.find(pkg_decl)
    if pkg_idx1 == -1:
        pkg_idx1 = content1.find('\npackage ')
    header_end1 = content1.find('\n', pkg_idx1 + 1) + 1
    
    header = content1[:header_end1]
    # Update resource list in header
    header = re.sub(r'// resource: .*\n', '', header)
    # Re-insert the +generated:types if it was removed
    if '+generated:types' not in header:
        header = re.sub(r'package v1beta1', '// +generated:types\npackage v1beta1', header)

    for res in all_resources:
        header = header.replace('package v1beta1', f'// resource: {res}\npackage v1beta1')
    
    body1 = content1[header_end1:]

    pkg_idx2 = content2.find(pkg_decl)
    if pkg_idx2 == -1:
        pkg_idx2 = content2.find('\npackage ')
    header_end2 = content2.find('\n', pkg_idx2 + 1) + 1
    body2 = content2[header_end2:]

    # Use a set to keep track of seen types to avoid duplicates
    seen_types = set()
    merged_body = body1
    
    # Identify types in body1
    for match in re.finditer(r'type (\w+) struct', body1):
        seen_types.add(match.group(1))

    # Append types from body2 that are not in body1
    # Types usually start with // +kcc:proto=
    # We'll split by // +kcc:proto=
    types2 = re.split(r'(?=// \+kcc:proto=)', body2)
    for t in types2:
        if not t.strip():
            continue
        match = re.search(r'type (\w+) struct', t)
        if match:
            typename = match.group(1)
            if typename not in seen_types:
                merged_body += '\n' + t.strip() + '\n'
                seen_types.add(typename)

    with open(outfile, 'w') as f:
        f.write(header + merged_body)

if __name__ == "__main__":
    merge_generated_files(sys.argv[1], sys.argv[2], sys.argv[3])
