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

import re
import os
import subprocess

def get_repo_root():
    try:
        return subprocess.check_output(["git", "rev-parse", "--show-toplevel"]).strip().decode("utf-8")
    except Exception:
        return os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

def parse_types(content):
    # Find all type declarations and their names
    # Returns a list of dicts: {"name": str, "start": int, "end": int}
    lines = content.split('\n')
    types_found = []
    
    i = 0
    while i < len(lines):
        line = lines[i].strip()
        # Match type declarations
        match = re.match(r'^type\s+(\w+)\s+(struct|string|int|int32|int64|float64|bool|\[\]byte)', line)
        if match:
            type_name = match.group(1)
            is_struct = match.group(2) == 'struct'
            
            # Find preceding comments/annotations
            start_idx = i
            while start_idx > 0 and lines[start_idx - 1].strip().startswith('//'):
                start_idx -= 1
                
            end_idx = i
            if is_struct:
                # Find matching closing brace
                brace_count = 0
                started = False
                for j in range(i, len(lines)):
                    l = lines[j]
                    if '{' in l:
                        brace_count += l.count('{')
                        started = True
                    if '}' in l:
                        brace_count -= l.count('}')
                    if started and brace_count == 0:
                        end_idx = j
                        break
            else:
                end_idx = i
                
            types_found.append({
                "name": type_name,
                "start": start_idx,
                "end": end_idx + 1 # exclusive
            })
            i = end_idx + 1
        else:
            i += 1
            
    return types_found

def main():
    root = get_repo_root()
    types_file = os.path.join(root, "apis/vertexai/v1alpha1/types.generated.go")
    index_file = os.path.join(root, "apis/vertexai/v1alpha1/vertexaiindexendpoint_types.generated.go")
    
    if not os.path.exists(types_file) or not os.path.exists(index_file):
        print(f"Files not found: {types_file} or {index_file}")
        return
        
    with open(types_file, 'r') as f:
        types_content = f.read()
        
    with open(index_file, 'r') as f:
        index_content = f.read()
        
    declared_in_types = {t["name"] for t in parse_types(types_content)}
    
    index_types = parse_types(index_content)
    index_lines = index_content.split('\n')
    
    # We will mark lines to be deleted
    to_delete = set()
    for t in index_types:
        if t["name"] in declared_in_types:
            print(f"Removing duplicate type {t['name']} from vertexaiindexendpoint_types.generated.go")
            for idx in range(t["start"], t["end"]):
                to_delete.add(idx)
                
    # Reconstruct the file excluding deleted lines
    new_lines = [line for idx, line in enumerate(index_lines) if idx not in to_delete]
    
    # Also clean up duplicate blank lines left over
    final_lines = []
    prev_blank = False
    for line in new_lines:
        is_blank = line.strip() == ""
        if is_blank and prev_blank:
            continue
        final_lines.append(line)
        prev_blank = is_blank
        
    with open(index_file, 'w') as f:
        f.write('\n'.join(final_lines))

if __name__ == "__main__":
    main()
