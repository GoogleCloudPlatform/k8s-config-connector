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

import yaml
import glob
import os

candidate_files = []
# Just iterate over all files in config/crds/resources
for root, _, files in os.walk('config/crds/resources'):
    for file in files:
        if file.endswith('.yaml'):
            candidate_files.append(os.path.join(root, file))

for path in candidate_files:
    try:
        with open(path, 'r') as f:
            content = f.read()
            if 'cnrm.cloud.google.com/dcl2crd: "true"' not in content:
                continue
            
            # parse yaml
            docs = list(yaml.safe_load_all(content))
            for doc in docs:
                if not doc:
                    continue
                metadata = doc.get('metadata', {})
                labels = metadata.get('labels', {})
                if labels.get('cnrm.cloud.google.com/dcl2crd') != "true":
                    continue
                
                spec = doc.get('spec', {})
                group = spec.get('group', '').split('.')[0]
                kind = spec.get('names', {}).get('kind', '')
                versions = spec.get('versions', [])
                
                has_beta = any('v1beta1' in v.get('name', '') for v in versions)
                if has_beta:
                    types_pattern = f"apis/{group}/v1beta1/*_types.go"
                    if not glob.glob(types_pattern):
                        print(f"Candidate: {group} {kind}")
    except Exception as e:
        pass
