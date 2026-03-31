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

import json
import re

def parse_data(md_file_path, config_file_path):
    resources = {}
    
    # Parse static_config.go first to establish the base list
    with open(config_file_path, 'r') as f:
        config_lines = f.readlines()
        
    for line in config_lines:
        line = line.strip()
        # Look for lines like: {Group: "...", Kind: "..."}: {DefaultController: ..., SupportedControllers: ...}
        if not line.startswith('{Group: '):
            continue
            
        group_match = re.search(r'Group:\s*"([^"]+)"', line)
        kind_match = re.search(r'Kind:\s*"([^"]+)"', line)
        default_ctrl_match = re.search(r'DefaultController:\s*k8s\.ReconcilerType([A-Za-z]+)', line)
        supported_ctrls_match = re.search(r'SupportedControllers:\s*\[\]k8s\.ReconcilerType\{(.*?)\}', line)
        
        if group_match and kind_match:
            group_full = group_match.group(1)
            # extract the prefix (e.g., 'accesscontextmanager' from 'accesscontextmanager.cnrm.cloud.google.com')
            group = group_full.split('.')[0]
            kind = kind_match.group(1)
            
            resources[kind] = create_default_resource(kind, group)
            
            if default_ctrl_match:
                # E.g. ReconcilerTypeTerraform -> Terraform
                resources[kind]['defaultController'] = default_ctrl_match.group(1)
                # Override the old controllerType field for UI backwards compatibility if needed, 
                # or just use defaultController
                resources[kind]['controllerType'] = default_ctrl_match.group(1)
                
            if supported_ctrls_match:
                ctrls_raw = supported_ctrls_match.group(1)
                # find all occurrences of k8s.ReconcilerTypeXXX
                supported = re.findall(r'k8s\.ReconcilerType([A-Za-z]+)', ctrls_raw)
                resources[kind]['supportedControllers'] = supported
                if 'Direct' in supported:
                    resources[kind]['state'] = 'Completed'
                    resources[kind]['steps'] = {
                        "gen-types": True,
                        "identity-reference": True,
                        "mapper-fuzzer": True,
                        "mocks": True,
                        "controller": True,
                        "tests": True
                    }

    # Now parse MD file to enrich dependencies and notes
    with open(md_file_path, 'r') as f:
        md_lines = f.readlines()
        
    for line in md_lines:
        line = line.strip()
        if not line.startswith('- '):
            continue
            
        if ' depends on: ' in line:
            parts = line[2:].split(' depends on: ')
            kind = parts[0].strip()
            dep = parts[1].strip()
            
            # Only add if it exists in static_config.go
            if kind in resources:
                if dep not in resources[kind]['dependencies']:
                    resources[kind]['dependencies'].append(dep)
        else:
            kind_raw = line[2:]
            kind = kind_raw.split('(')[0].strip()
            
            # Only enrich if it exists in static_config.go
            if kind in resources:
                if 'missing _reference.go only' in line:
                    resources[kind]['notes'] = 'Missing _reference.go'
                    resources[kind]['steps']['identity-reference'] = False

    return list(resources.values())

def create_default_resource(kind, group="unknown"):
    return {
        "group": group,
        "kind": kind,
        "version": "v1beta1",
        "controllerType": "Unknown",
        "defaultController": "Unknown",
        "supportedControllers": [],
        "dependencies": [],
        "state": "Not Started",
        "steps": {
            "gen-types": False,
            "identity-reference": False,
            "mapper-fuzzer": False,
            "mocks": False,
            "controller": False,
            "tests": False
        },
        "mocksLastRefreshed": "Never",
        "notes": ""
    }

if __name__ == "__main__":
    data = parse_data('../../hack/resource-dependencies.md', '../../pkg/controller/resourceconfig/static_config.go')
    data = sorted(data, key=lambda x: x['kind'])
    with open('data.json', 'w') as f:
        json.dump(data, f, indent=2)
    print(f"Generated data.json with {len(data)} resources.")
