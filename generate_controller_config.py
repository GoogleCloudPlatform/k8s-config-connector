# Copyright 2025 Google LLC
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
import yaml

def get_group_and_kind(crd_file):
    with open(crd_file, 'r') as f:
        crd = yaml.safe_load(f)
        return crd['spec']['group'], crd['spec']['names']['kind']

def get_default_controller(crd_file):
    with open(crd_file, 'r') as f:
        crd = yaml.safe_load(f)
        labels = crd.get('metadata', {}).get('labels', {})
        if labels.get('cnrm.cloud.google.com/tf2crd') == 'true':
            return 'terraform'
        if labels.get('cnrm.cloud.google.com/dcl2crd') == 'true':
            return 'dcl'
    return 'direct'

def get_supported_controllers(kind, crd_file_path):
    supported = []
    # Check for direct controller
    direct_controller_path = os.path.join('pkg', 'controller', 'direct', kind.lower() + '_controller.go')
    if os.path.exists(direct_controller_path):
        supported.append('direct')

    # This is a simplified check. A more robust check would involve parsing the servicemappings
    # For now, we assume if a resource has a default controller, it's supported.
    # A more accurate implementation would be needed for a real-world scenario.
    if get_default_controller(crd_file_path) == 'terraform':
        if 'terraform' not in supported:
            supported.append('terraform')
    if get_default_controller(crd_file_path) == 'dcl':
        if 'dcl' not in supported:
            supported.append('dcl')
    if not supported:
        supported.append(get_default_controller(crd_file_path))
    return list(set(supported))

crd_files = [f for f in os.listdir('crds') if f.endswith('.yaml')]

resources = []
for crd_file in crd_files:
    crd_path = os.path.join('crds', crd_file)
    group, kind = get_group_and_kind(crd_path)
    default_controller = get_default_controller(crd_path)
    supported_controllers = get_supported_controllers(kind, crd_path)
    resources.append({
        'group': group,
        'kind': kind,
        'defaultController': default_controller,
        'supportedControllers': supported_controllers,
    })

with open('config/controller_config.yaml', 'w') as f:
    yaml.dump({'resources': resources}, f, default_flow_style=False)

print("Successfully generated config/controller_config.yaml")