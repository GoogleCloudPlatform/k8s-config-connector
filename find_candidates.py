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
import glob
import sys
import yaml

crd_dir = os.path.join(os.path.dirname(__file__) or ".", "config", "crds", "resources")
if not os.path.exists(crd_dir):
    print(f"Error: Directory {crd_dir} not found.")
    sys.exit(1)

for file_path in glob.glob(os.path.join(crd_dir, "**", "*.yaml"), recursive=True):
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            data = yaml.safe_load(f)
            
        metadata = data.get('metadata') or {}
        labels = metadata.get('labels') or {}
        if labels.get('cnrm.cloud.google.com/dcl2crd') != 'true':
            continue
            
        spec = data.get('spec') or {}
        versions = spec.get('versions') or []
        has_beta = any((v.get('name') == 'v1beta1') for v in versions)
        if not has_beta:
            continue
            
        group = (spec.get('group') or '').split('.')[0]
        names = spec.get('names') or {}
        kind = names.get('kind', '')
        
        # Check if types.go exists
        types_pattern = os.path.join(os.path.dirname(__file__) or ".", 'apis', group, 'v1beta1', '*_types.go')
        if not glob.glob(types_pattern):
            print(f"Candidate: {group} {kind}")
    except yaml.YAMLError as e:
        print(f"YAML parsing error in {file_path}: {e}")
    except IOError as e:
        print(f"IO error reading {file_path}: {e}")
    except Exception as e:
        print(f"Unexpected error processing {file_path}: {e}")
