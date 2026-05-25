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

def comment_out_functions(file_path, manual_mappers_path, mapper_path):
    with open(file_path, 'r') as f:
        content = f.read()

    with open(manual_mappers_path, 'r') as f:
        manual_content = f.read()

    with open(mapper_path, 'r') as f:
        mapper_content = f.read()

    # Pattern to match function declarations
    func_pattern = re.compile(r'func (\w+)\((.*?)\) (.*?) \{')
    
    # Find all manually implemented functions
    manual_funcs = set(re.findall(r'func (\w+)\(', manual_content))
    manual_funcs.update(re.findall(r'func (\w+)\(', mapper_content))

    lines = content.split('\n')
    new_lines = []
    i = 0
    
    while i < len(lines):
        line = lines[i]
        match = func_pattern.match(line)
        if match:
            func_name = match.group(1)
            
            # Find end of function
            start_i = i
            depth = 1
            i += 1
            while i < len(lines) and depth > 0:
                depth += lines[i].count('{')
                depth -= lines[i].count('}')
                i += 1
            
            func_body = lines[start_i:i]
            
            if func_name in manual_funcs:
                # Comment out
                for j in range(len(func_body)):
                    new_lines.append("// " + func_body[j])
            else:
                for j in range(len(func_body)):
                    new_lines.append(func_body[j])
        else:
            new_lines.append(line)
            i += 1

    with open(file_path, 'w') as f:
        f.write('\n'.join(new_lines))

if __name__ == "__main__":
    comment_out_functions(
        'pkg/controller/direct/vertexai/customjob/mapper.generated.go',
        'pkg/controller/direct/vertexai/customjob/manual_mappers.go',
        'pkg/controller/direct/vertexai/customjob/mapper.go'
    )
