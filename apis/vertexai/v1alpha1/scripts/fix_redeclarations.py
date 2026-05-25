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
import re
import sys

def fix_file(path):
    if not os.path.exists(path):
        print(f"File {path} not found, skipping.")
        return
    
    print(f"Fixing {path}")
    with open(path, "r") as f:
        lines = f.readlines()

    # Pre-scan for functions and their return types
    func_info = {} # name -> list of (line_index, return_type)
    
    for i, line in enumerate(lines):
        match = re.match(r"func (\w+)\(.*?\) ([\*\w\.]+)", line)
        if match:
            name = match.group(1)
            ret_type = match.group(2)
            if name not in func_info:
                func_info[name] = []
            func_info[name].append((i, ret_type))

    # Determine which ones to keep
    keep_indices = set()
    for name, infos in func_info.items():
        if len(infos) == 1:
            keep_indices.add(infos[0][0])
            continue
        
        best_index = infos[0][0]
        found_preferred = False
        for idx, ret_type in infos:
            if ret_type.startswith("*krm.") or ret_type == "krm.":
                best_index = idx
                found_preferred = True
                break
            if not "." in ret_type: # Local type
                best_index = idx
                found_preferred = True
        
        if not found_preferred:
             best_index = infos[0][0]
             
        keep_indices.add(best_index)

    # Check manual mapper files in the same directory (and subdirectories) for existing functions
    dir_path = os.path.dirname(path)
    manual_funcs = set()
    for root, _, files in os.walk(dir_path):
        for entry in files:
            if entry.endswith(".go") and entry != os.path.basename(path):
                with open(os.path.join(root, entry), "r") as f_manual:
                    manual_content = f_manual.read()
                    funcs = re.findall(r"func (\w+)\(", manual_content)
                    for f in funcs:
                        manual_funcs.add(f)

    new_lines = []
    skip_until_next_func = False
    for i, line in enumerate(lines):
        match = re.match(r"func (\w+)\(", line)
        if match:
            name = match.group(1)
            if name in manual_funcs or i not in keep_indices:
                skip_until_next_func = True
            else:
                skip_until_next_func = False
        
        if skip_until_next_func:
            new_lines.append("// " + line)
        else:
            new_lines.append(line)

    with open(path, "w") as f:
        f.writelines(new_lines)

if __name__ == "__main__":
    if len(sys.argv) > 1:
        for path in sys.argv[1:]:
            fix_file(path)
    else:
        fix_file("pkg/controller/direct/vertexai/mapper.generated.go")
