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
# See the License for the_license.

import re
import os

def get_deps(code):
    funcs = re.findall(r'^func ([A-Z]\w+)\(', code, re.MULTILINE)
    calls = {}
    for func in funcs:
        body_start = code.find(f'func {func}(')
        body_end = code.find('\n}\n', body_start) + 3
        if body_end < 3: body_end = len(code)
        body = code[body_start:body_end]
        calls[func] = set(re.findall(r'([A-Z]\w+_[A-Za-z0-9_]+)', body))
    return funcs, calls

def prune():
    # Identify handwritten functions in other go files
    handwritten = set()
    dir_path = 'pkg/controller/direct/aiplatform/vertexaihyperparametertuningjob'
    
    # We will gather all dependencies from handwritten code too
    all_handwritten_calls = {}
    for file_name in os.listdir(dir_path):
        if file_name.endswith('.go') and file_name != 'mapper.generated.go':
            with open(os.path.join(dir_path, file_name)) as f:
                content = f.read()
                hw_funcs, hw_calls = get_deps(content)
                handwritten.update(hw_funcs)
                all_handwritten_calls.update(hw_calls)

    print("Handwritten functions to skip in generation:", handwritten)

    # Copy the newly generated main mapper into this package
    with open('pkg/controller/direct/aiplatform/mapper.generated.go') as f:
        code = f.read()

    # Change package name
    code = re.sub(r'package aiplatform', 'package vertexaihyperparametertuningjob', code)

    funcs, calls = get_deps(code)
    
    # Merge the handwritten calls/dependencies so the walker traverses them too
    calls.update(all_handwritten_calls)
    
    needed = set()
    visited = set()
    queue = ['VertexAIHyperparameterTuningJobObservedState_v1alpha1_FromProto', 'VertexAIHyperparameterTuningJobObservedState_v1alpha1_ToProto', 'VertexAIHyperparameterTuningJobSpec_v1alpha1_FromProto', 'VertexAIHyperparameterTuningJobSpec_v1alpha1_ToProto']
    
    # Also support unversioned / fallback starting names if any
    for q in list(queue):
        fallback = q.replace('_v1alpha1', '')
        if fallback in funcs or fallback in handwritten:
            queue.append(fallback)

    while queue:
        curr = queue.pop(0)
        if curr in visited:
            continue
        visited.add(curr)
        
        # Only add to 'needed' for generated output if it's in funcs and not handwritten
        if curr in funcs and curr not in handwritten:
            needed.add(curr)
        
        # Walk calls for both handwritten and generated functions
        if curr in calls:
            for c in calls[curr]:
                if (c in funcs or c in handwritten) and c not in visited:
                    queue.append(c)

    print("Needed functions for mapper.generated.go:", needed)

    out_code = []
    # Safely extract package and import block without any trailing comments
    import_end = code.find('\n)\n')
    if import_end != -1:
        header = code[:import_end + 3]
    else:
        header = code[:code.find('\nfunc ')]
    out_code.append(header)
    
    for func in funcs: # preserve order
        if func in needed and func not in handwritten:
            body_start = code.find(f'\nfunc {func}(')
            body_end = code.find('\n}\n', body_start) + 3
            out_code.append(code[body_start:body_end].strip())

    with open('pkg/controller/direct/aiplatform/vertexaihyperparametertuningjob/mapper.generated.go', 'w') as f:
        f.write('\n\n'.join(out_code) + '\n')

prune()
