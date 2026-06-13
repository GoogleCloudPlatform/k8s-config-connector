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
    with open('pkg/controller/direct/aiplatform/vertexaihyperparametertuningjob/mapper.generated.go') as f:
        code = f.read()

    funcs, calls = get_deps(code)
    
    needed = set()
    queue = ['VertexAIHyperparameterTuningJobObservedState_FromProto', 'VertexAIHyperparameterTuningJobObservedState_ToProto', 'VertexAIHyperparameterTuningJobSpec_FromProto', 'VertexAIHyperparameterTuningJobSpec_ToProto']
    
    while queue:
        curr = queue.pop(0)
        if curr not in needed:
            needed.add(curr)
            if curr in calls:
                for c in calls[curr]:
                    if c in funcs and c not in needed:
                        queue.append(c)

    print("Needed functions:", needed)

    out_code = []
    header = code[:code.find('\nfunc ')]
    out_code.append(header)
    
    for func in funcs: # preserve order
        if func in needed:
            body_start = code.find(f'\nfunc {func}(')
            body_end = code.find('\n}\n', body_start) + 3
            out_code.append(code[body_start:body_end].strip())

    with open('pkg/controller/direct/aiplatform/vertexaihyperparametertuningjob/mapper.generated.go', 'w') as f:
        f.write('\n\n'.join(out_code) + '\n')

prune()
