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

crd_dir = "config/crds/resources"
for filename in os.listdir(crd_dir):
    if not filename.endswith(".yaml"):
        continue
    filepath = os.path.join(crd_dir, filename)
    with open(filepath, 'r') as f:
        content = f.read()
    
    # Check for dcl2crd label
    if 'cnrm.cloud.google.com/dcl2crd: "true"' not in content:
        continue
        
    # Check for v1beta1 version
    if 'name: v1beta1' not in content:
        continue
        
    # Extract group
    group_match = re.search(r'^  group: (.*?)$', content, re.MULTILINE)
    if not group_match:
        continue
    group_full = group_match.group(1).strip()
    group_short = group_full.split('.')[0]
    
    # Extract kind
    kind_match = re.search(r'^    kind: (.*?)$', content, re.MULTILINE)
    if not kind_match:
        continue
    kind = kind_match.group(1).strip()
    
    # Check if types.go exists
    # The types.go is usually named lowercased kind + _types.go, but can vary.
    # Let's just check if any _types.go exists in the directory apis/<group_short>/v1beta1/
    types_dir = os.path.join("apis", group_short, "v1beta1")
    types_exist = False
    if os.path.isdir(types_dir):
        for f in os.listdir(types_dir):
            if f.endswith("_types.go"):
                types_exist = True
                break
                
    if not types_exist:
        print(f"Candidate: Group={group_short} Kind={kind}")

