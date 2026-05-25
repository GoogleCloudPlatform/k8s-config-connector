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
import sys
import os

def comment_out_functions(file_path):
    if not os.path.exists(file_path):
        print(f"File {file_path} not found, skipping.")
        return

    print(f"Applying fix_vertexai_mappers to {file_path}")
    with open(file_path, 'r') as f:
        content = f.read()

    # Packages to comment out
    bad_packages = ['krmcolabv1alpha1', 'krmvertexaiv1beta1', 'krmcomputev1beta1', 'krmaiplatformv1alpha1']
    
    # Pattern to match function declarations
    # Matches: func Name(args) ReturnType {
    func_pattern = re.compile(r'func (\w+)\((.*?)\) (.*?) \{')
    
    lines = content.split('\n')
    new_lines = []
    i = 0
    
    active_functions = {} # name -> list of decl info

    while i < len(lines):
        line = lines[i]
        match = func_pattern.match(line)
        if match:
            func_name = match.group(1)
            args = match.group(2)
            return_type = match.group(3)
            
            # Check if it uses any of the bad packages
            uses_bad_pkg = any(pkg in args or pkg in return_type for pkg in bad_packages)
            
            # Find end of function
            start_i = i
            depth = 1
            i += 1
            while i < len(lines) and depth > 0:
                depth += lines[i].count('{')
                depth -= lines[i].count('}')
                i += 1
            
            func_body = lines[start_i:i]
            
            if uses_bad_pkg:
                # Comment out
                for j in range(len(func_body)):
                    new_lines.append("// " + func_body[j])
            else:
                # Potential keep, but check for duplicates later
                if func_name not in active_functions:
                    active_functions[func_name] = []
                active_functions[func_name].append({
                    'start': len(new_lines),
                    'body': func_body,
                    'return_type': return_type,
                    'args': args
                })
                # Add placeholders
                for j in range(len(func_body)):
                    new_lines.append(func_body[j])
        else:
            new_lines.append(line)
            i += 1

    # Now handle duplicates for uncommented functions
    for func_name, decls in active_functions.items():
        if len(decls) > 1:
            to_keep = None
            for d in decls:
                # Prefer the one that uses 'krm.' (our primary package)
                if 'krm.' in d['return_type'] or 'krm.' in d['args']:
                    to_keep = d
                    break
            
            if not to_keep:
                # If none use krm., keep the first one
                to_keep = decls[0]
            
            for d in decls:
                if d != to_keep:
                    # Comment out this duplicate
                    for j in range(len(d['body'])):
                        new_lines[d['start'] + j] = "// " + d['body'][j]

    with open(file_path, 'w') as f:
        f.write('\n'.join(new_lines))

if __name__ == "__main__":
    if len(sys.argv) > 1:
        for path in sys.argv[1:]:
            comment_out_functions(path)
    else:
        comment_out_functions('pkg/controller/direct/vertexai/mapper.generated.go')
