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

with open("apis/vertexai/v1alpha1/customjob_types.go", "r") as f:
    lines = f.readlines()

out = []
doc_lines = []

for line in lines:
    stripped = line.strip()
    if stripped.startswith("//"):
        doc_lines.append(line)
    else:
        if stripped != "" and ("json:" in line or "type " in line):
            # It's a field or type definition
            # Analyze doc_lines to see if it says Required or Optional
            doc_text = " ".join(doc_lines)
            if "json:" in line:
                is_required = "Required." in doc_text or "Required " in doc_text
                is_optional = "Optional." in doc_text or "Optional " in doc_text
                
                # Check if it already has tags
                has_required_tag = "+required" in doc_text or "+kubebuilder:validation:Required" in doc_text
                has_optional_tag = "+optional" in doc_text

                if is_required and not has_required_tag:
                    doc_lines.append("\t// +kubebuilder:validation:Required\n")
                    if ",omitempty" in line:
                        line = line.replace(",omitempty", "")
                if is_optional and not has_optional_tag:
                    doc_lines.append("\t// +optional\n")
                    # ensure it has omitempty
                    if ",omitempty" not in line and '"' in line:
                        line = line.replace('"', ',omitempty"', 1) # This is hacky, but wait, JSON tag usually has omitempty if optional

                # Add Enum tag for AcceleratorType
                if "AcceleratorType *string" in line:
                    enum_values = [
                        "ACCELERATOR_TYPE_UNSPECIFIED", "NVIDIA_TESLA_K80", "NVIDIA_TESLA_P100", 
                        "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_A100", 
                        "NVIDIA_A100_80GB", "NVIDIA_L4", "NVIDIA_H100_80GB", "NVIDIA_H100_MEGA_80GB", 
                        "TPU_V2", "TPU_V3", "TPU_V4_POD", "TPU_V5_LITEPOD"
                    ]
                    enum_str = ";".join(enum_values)
                    doc_lines.append(f"\t// +kubebuilder:validation:Enum={enum_str}\n")
            
            out.extend(doc_lines)
            out.append(line)
            doc_lines = []
        elif stripped == "" or "{" in line or "}" in line or "package " in line or "import " in line or "var " in line:
            out.extend(doc_lines)
            out.append(line)
            doc_lines = []
        else:
            doc_lines.append(line)

with open("apis/vertexai/v1alpha1/customjob_types.go", "w") as f:
    f.writelines(out)
