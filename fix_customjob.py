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

new_lines = []
for i, line in enumerate(lines):
    # Fix imports
    if 'refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"' in line:
        new_lines.append(line)
        new_lines.append('\tcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"\n')
        continue

    # Fix struct fields
    if 'PersistentResourceRef *refsv1beta1.VertexAIPersistentResourceRef' in line:
        line = line.replace('*refsv1beta1.VertexAIPersistentResourceRef', '*VertexAIPersistentResourceRef')
    elif 'TensorboardRef *refsv1beta1.VertexAITensorboardRef' in line:
        line = line.replace('*refsv1beta1.VertexAITensorboardRef', '*VertexAITensorboardRef')
    elif 'NetworkRef *refsv1beta1.ComputeNetworkRef' in line:
        line = line.replace('*refsv1beta1.ComputeNetworkRef', '*computev1beta1.ComputeNetworkRef')

    # Fix JSON tags
    if 'json:"protectedArtifactLocationId,omitempty"' in line:
        line = line.replace('protectedArtifactLocationId', 'protectedArtifactLocationID')
    if 'json:"reservedIpRanges,omitempty"' in line:
        line = line.replace('reservedIpRanges', 'reservedIPRanges')
    if 'json:"imageUri,omitempty"' in line:
        line = line.replace('imageUri', 'imageURI')
    if 'json:"executorImageUri,omitempty"' in line:
        line = line.replace('executorImageUri', 'executorImageURI')
    if 'json:"packageUris,omitempty"' in line:
        line = line.replace('packageUris', 'packageURIs')
    if 'json:"outputUriPrefix,omitempty"' in line:
        line = line.replace('outputUriPrefix', 'outputURIPrefix')

    # Add required/optional
    if line.strip().startswith('// Required.') or line.strip().startswith('// Required '):
        # We'll just append +required after this comment block later, or right here if the next line is the field
        pass
    
    if line.strip().startswith('// Optional.') or line.strip().startswith('// Optional '):
        pass

    # Actually, let's just insert +required or +optional right before the field definition
    # if we saw "Required" or "Optional" in the docstring.
    
    new_lines.append(line)

with open("apis/vertexai/v1alpha1/customjob_types.go", "w") as f:
    f.writelines(new_lines)
