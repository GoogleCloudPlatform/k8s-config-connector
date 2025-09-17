#!/usr/bin/env python3
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
import glob
import yaml
import subprocess
import re

def get_repo_root():
    return subprocess.check_output(['git', 'rev-parse', '--show-toplevel']).strip().decode('utf-8')

def find_direct_controller_kinds(repo_root):
    direct_controller_kinds = set()
    controller_dir = os.path.join(repo_root, 'pkg', 'controller', 'direct')
    for filepath in glob.glob(os.path.join(controller_dir, '**', '*_controller.go'), recursive=True):
        with open(filepath, 'r') as f:
            content = f.read()
            # Look for RegisterModel(krm.<Kind>GVK, ...)
            matches = re.findall(r'RegisterModel\((?:krm|krmv1beta1)\.(\w+)GVK', content)
            for match in matches:
                direct_controller_kinds.add(match)
    return direct_controller_kinds

def main():
    repo_root = get_repo_root()
    crd_dir = os.path.join(repo_root, 'config', 'crds', 'resources')
    
    direct_controller_kinds = find_direct_controller_kinds(repo_root)
    
    resources = []
    
    for crd_file in sorted(glob.glob(os.path.join(crd_dir, '*.yaml'))):
        with open(crd_file, 'r') as f:
            try:
                crd = next(yaml.safe_load_all(f), None)
            except yaml.YAMLError as e:
                print(f"Skipping file {crd_file} due to YAML error: {e}")
                continue

        if not crd or 'spec' not in crd or crd.get('apiVersion') != 'apiextensions.k8s.io/v1' or crd.get('kind') != 'CustomResourceDefinition':
            continue

        spec = crd.get('spec', {})
        names = spec.get('names', {})
        group = spec.get('group')
        kind = names.get('kind')
        
        if not group or not kind:
            continue
            
        labels = crd.get('metadata', {}).get('labels', {})
        
        has_direct_controller = kind in direct_controller_kinds
        is_tf_resource = labels.get('cnrm.cloud.google.com/tf2crd') == 'true'
        is_dcl_resource = labels.get('cnrm.cloud.google.com/dcl2crd') == 'true'

        # Determine supported controllers
        supported_controllers = set()
        if has_direct_controller:
            supported_controllers.add('k8s.ReconcilerTypeDirect')
        if is_dcl_resource:
            supported_controllers.add('k8s.ReconcilerTypeDCL')

        # Add Terraform support if not a DCL-only resource.
        if not is_dcl_resource or has_direct_controller or is_tf_resource:
            if kind not in ('IAMPolicy', 'IAMPartialPolicy', 'IAMPolicyMember', 'IAMAuditConfig'):
                supported_controllers.add('k8s.ReconcilerTypeTerraform')

        # Determine default controller
        if kind in ('IAMPolicy', 'IAMPartialPolicy', 'IAMPolicyMember', 'IAMAuditConfig'):
            default_controller = f'k8s.ReconcilerType{kind}'
            supported_controllers = {default_controller}
        elif is_dcl_resource:
            default_controller = 'k8s.ReconcilerTypeDCL'
        elif is_tf_resource:
            default_controller = 'k8s.ReconcilerTypeTerraform'
        elif has_direct_controller:
            default_controller = 'k8s.ReconcilerTypeDirect'
        else:
            default_controller = 'k8s.ReconcilerTypeTerraform' # Fallback
        
        if not supported_controllers:
            supported_controllers.add(default_controller)

        resources.append({
            'Group': group,
            'Kind': kind,
            'DefaultController': default_controller,
            'SupportedControllers': sorted(list(supported_controllers))
        })

    # Generate the go file content
    lines = []
    lines.append('// Copyright 2025 Google LLC')
    lines.append('//')
    lines.append('// Licensed under the Apache License, Version 2.0 (the "License");')
    lines.append('// you may not use this file except in compliance with the License.')
    lines.append('// You may obtain a copy of the License at')
    lines.append('//')
    lines.append('//      http://www.apache.org/licenses/LICENSE-2.0')
    lines.append('//')
    lines.append('// Unless required by applicable law or agreed to in writing, software')
    lines.append('// distributed under the License is distributed on an "AS IS" BASIS,')
    lines.append('// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.')
    lines.append('// See the License for the specific language governing permissions and')
    lines.append('// limitations under the License.')
    lines.append('//')
    lines.append('// *** AUTO-GENERATED BY dev/tasks/generate_static_config.py ***')
    lines.append('// *** DO NOT EDIT MANUALLY ***')
    lines.append('')
    lines.append('package resourceconfig')
    lines.append('')
    lines.append('import "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"')
    lines.append('')
    lines.append('// ControllerConfigStatic is the static controller configuration for all resources.')
    lines.append('// It is ordered by GroupKind alphabetically.')
    lines.append('var ControllerConfigStatic = &ResourcesControllerMap{')

    # Find longest Group+Kind for alignment
    max_len = 0
    for res in resources:
        l = len('    {Group: "%s", Kind: "%s"}:' % (res['Group'], res['Kind']))
        if l > max_len:
            max_len = l

    for res in sorted(resources, key=lambda x: (x['Group'], x['Kind'])):
        key_part = '    {Group: "%s", Kind: "%s"}:' % (res['Group'], res['Kind'])
        padding = ' ' * (max_len - len(key_part) + 1)
        
        supported_str = ", ".join(res['SupportedControllers'])
        
        value_part = '{DefaultController: %s, SupportedControllers: []k8s.ReconcilerType{%s}},' % (res['DefaultController'], supported_str)
        
        lines.append("%s%s%s" % (key_part, padding, value_part))

    lines.append("}")
    
    go_file_content = "\n".join(lines)

    output_path = os.path.join(repo_root, 'pkg', 'controller', 'resourceconfig', 'static_config.go')
    with open(output_path, 'w') as f:
        f.write(go_file_content)

    print(f"Generated {output_path}")

if __name__ == '__main__':
    main()
