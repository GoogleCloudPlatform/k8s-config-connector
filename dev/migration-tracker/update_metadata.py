#!/usr/bin/env python3
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
import os

def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    data_path = os.path.join(script_dir, 'data.json')
    repo_root = os.path.abspath(os.path.join(script_dir, '../../'))
    
    with open(data_path, 'r') as f:
        data = json.load(f)
        
    for resource in data:
        kind_lower = resource['kind'].lower()
        group_lower = resource['group'].lower()
        
        # Determine gcpTestLocation
        gcp_test_dir = ""
        base_test_dir = os.path.join(repo_root, "pkg/test/resourcefixture/testdata/basic")
        
        # Preferred structure: group/version/kind
        version = resource.get('version', 'v1beta1')
        preferred_path = os.path.join(base_test_dir, group_lower, version, kind_lower)
        if os.path.isdir(preferred_path):
            gcp_test_dir = os.path.relpath(preferred_path, repo_root)
        else:
            # Fallback search
            if os.path.exists(base_test_dir):
                for root, dirs, files in os.walk(base_test_dir):
                    if os.path.basename(root) == kind_lower:
                        gcp_test_dir = os.path.relpath(root, repo_root)
                        break
                    
        if gcp_test_dir:
            resource['gcpTestLocation'] = gcp_test_dir
        else:
            resource['gcpTestLocation'] = ""
            
        # Determine mockgcpLocation
        mock_dir = ""
        expected_mock_dir = os.path.join(repo_root, "mockgcp", f"mock{group_lower}")
        if os.path.isdir(expected_mock_dir):
            mock_dir = os.path.relpath(expected_mock_dir, repo_root)
            
        if mock_dir:
            resource['mockgcpLocation'] = mock_dir
        else:
            resource['mockgcpLocation'] = ""

    with open(data_path, 'w') as f:
        json.dump(data, f, indent=2)

    print("Successfully updated metadata in data.json")

if __name__ == "__main__":
    main()
