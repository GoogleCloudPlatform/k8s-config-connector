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

with open("apis/cloudbuild/v1alpha1/connection_types.go", "r") as f:
    lines = f.readlines()

new_lines = []
i = 0
while i < len(lines):
    line = lines[i]
    if line.strip().startswith("// +kcc:proto:field="):
        # The previous line might be a Required/Optional comment
        j = i - 1
        is_required = False
        while j >= 0 and lines[j].strip().startswith("//"):
            if "Required." in lines[j]:
                is_required = True
                break
            j -= 1
        
        # Check if there's already a kubebuilder validation comment
        has_validation = False
        j = i - 1
        while j >= 0 and lines[j].strip().startswith("//"):
            if "+kubebuilder:validation" in lines[j]:
                has_validation = True
                break
            j -= 1
            
        if not has_validation:
            if is_required:
                new_lines.insert(len(new_lines), "        // +kubebuilder:validation:Required\n")
            else:
                new_lines.insert(len(new_lines), "        // +kubebuilder:validation:Optional\n")
    new_lines.append(line)
    i += 1

with open("apis/cloudbuild/v1alpha1/connection_types.go", "w") as f:
    f.writelines(new_lines)
