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

def prune(file_path):
    if not os.path.exists(file_path):
        return
    
    with open(file_path, 'r') as f:
        lines = f.readlines()
    
    new_lines = []
    i = 0
    
    # We want to keep mappers for CustomJob and shared types that use 'krm.'
    # We want to discard mappers for Model, DataLabelingJob, etc.
    
    # Pattern to match function declarations
    func_pattern = re.compile(r'func (\w+)\((.*?)\) (.*?) \{')
    
    while i < len(lines):
        line = lines[i]
        match = func_pattern.match(line)
        if match:
            func_name = match.group(1)
            args = match.group(2)
            ret_type = match.group(3)
            
            # Find end of function
            start_i = i
            depth = 1
            i += 1
            while i < len(lines) and depth > 0:
                depth += lines[i].count('{')
                depth -= lines[i].count('}')
                i += 1
            
            func_body = lines[start_i:i]
            
            # Decision logic:
            keep = False
            
            # Keep if it's a CustomJob-specific mapper
            if "CustomJob" in func_name:
                keep = True
            
            # Keep if it's a shared mapper but uses 'krm.' (not krmaiplatformv1alpha1.)
            if ("krm." in args or "krm." in ret_type) and not "krmaiplatformv1alpha1" in args and not "krmaiplatformv1alpha1" in ret_type:
                keep = True
            
            # Special case for basic shared types that might not have 'krm.' prefix if they are local?
            # But in this package, 'krm' is the alias for vertexai/v1alpha1.
            
            # Discard if it's for other resources
            if "DataLabelingJob" in func_name or "AIPlatformModel" in func_name or "Notebook" in func_name or "Colab" in func_name:
                keep = False
            
            if keep:
                new_lines.extend(func_body)
            else:
                for b in func_body:
                    new_lines.append("// " + b)
        else:
            new_lines.append(line)
            i += 1

    with open(file_path, 'w') as f:
        f.writelines(new_lines)

if __name__ == "__main__":
    prune('pkg/controller/direct/vertexai/customjob/mapper.generated.go')
