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
import os
import sys
import glob
import yaml
from collections import defaultdict

def find_refs(schema, path=""):
    refs = set()
    if isinstance(schema, dict):
        if 'properties' in schema:
            for prop, prop_schema in schema['properties'].items():
                if prop.endswith("Ref"):
                    refs.add(prop)
                else:
                    refs.update(find_refs(prop_schema, path + "." + prop))
        elif 'items' in schema:
            refs.update(find_refs(schema['items'], path + "[]"))
    return refs

def get_implemented_types(apis_dir="../../apis"):
    implemented_kinds = {}
    struct_regex = re.compile(r"type\s+([A-Za-z0-9_]+)\s+struct\s*\{")
    if not os.path.exists(apis_dir):
        return implemented_kinds

    for root, _, files in os.walk(apis_dir):
        for file in files:
            if file.endswith("_types.go"):
                filepath = os.path.join(root, file)
                with open(filepath, 'r', encoding='utf-8') as f:
                    content = f.read()
                    matches = struct_regex.findall(content)
                    for kind in matches:
                        if kind.endswith("Spec") or kind.endswith("Status") or kind.endswith("List") or kind.endswith("ObservedState"):
                            continue
                        if kind not in implemented_kinds:
                            implemented_kinds[kind] = []
                        implemented_kinds[kind].append(filepath)
    return implemented_kinds

def find_direct_controllers(repo_root):
    controllers = {}
    fuzzers = {}
    controller_dir = os.path.join(repo_root, 'pkg', 'controller', 'direct')
    if not os.path.exists(controller_dir):
        return controllers, fuzzers
    for filepath in glob.glob(os.path.join(controller_dir, '**', '*_controller.go'), recursive=True):
        with open(filepath, 'r') as f:
            content = f.read()
        matches = re.findall(r':=\s*&krm\.(\w+)\{', content)
        if not matches:
            matches = re.findall(r':=\s*&krmv1beta1\.(\w+)\{', content)
        if not matches:
            matches = re.findall(r'RegisterModel\((?:krm|krmv1beta1)\.(\w+)GVK', content)
        for kind in matches:
            controllers[kind] = filepath
            
    for filepath in glob.glob(os.path.join(controller_dir, '**', '*_fuzzer.go'), recursive=True):
        with open(filepath, 'r') as f:
            content = f.read()
        matches = re.findall(r"fuzztesting\.NewKRMTyped(?:Spec)?Fuzzer\(\s*&pb\.[A-Za-z0-9_.]+\{\},\s*([A-Za-z0-9_]+)Spec_FromProto", content)
        for kind in matches:
            fuzzers[kind] = filepath
            
    return controllers, fuzzers

def find_test_fixtures(repo_root):
    test_kinds = set()
    basic_dir = os.path.join(repo_root, 'pkg', 'test', 'resourcefixture', 'testdata', 'basic')
    if not os.path.exists(basic_dir):
        return test_kinds
    for root, dirs, files in os.walk(basic_dir):
        if 'create.yaml' in files:
            grandparent = os.path.basename(os.path.dirname(root))
            test_kinds.add(grandparent.lower())
    return test_kinds

def build_dependency_graph(crds_dir="../../config/crds/resources"):
    known_kinds = {}
    dependencies = defaultdict(set)
    crd_docs = []
    
    if not os.path.exists(crds_dir):
        return dependencies, known_kinds
        
    for filename in os.listdir(crds_dir):
        if not filename.endswith(".yaml"):
            continue
        filepath = os.path.join(crds_dir, filename)
        with open(filepath, 'r') as f:
            try:
                for doc in yaml.safe_load_all(f):
                    if doc and doc.get("kind") == "CustomResourceDefinition":
                        crd_docs.append(doc)
                        kind = doc["spec"]["names"]["kind"]
                        group = doc["spec"]["group"]
                        known_kinds[kind] = group
            except Exception:
                pass

    for doc in crd_docs:
        kind = doc["spec"]["names"]["kind"]
        group = doc["spec"]["group"]
        
        for version in doc["spec"]["versions"]:
            if "schema" in version and "openAPIV3Schema" in version["schema"]:
                refs = find_refs(version["schema"]["openAPIV3Schema"])
                
                for ref in refs:
                    if ref == "externalRef":
                        continue
                    
                    ref_name = ref[:-3]
                    matched_kind = None
                    group_prefix = group.split('.')[0]
                    
                    for k in sorted(list(known_kinds.keys())):
                        if k.lower() == ref_name.lower() or k.lower() == (group_prefix + ref_name).lower():
                            matched_kind = k
                            break
                    
                    if not matched_kind:
                        if "project" in ref_name.lower(): matched_kind = "Project"
                        elif "folder" in ref_name.lower(): matched_kind = "Folder"
                        elif "organization" in ref_name.lower(): matched_kind = "Organization"
                        elif "billingaccount" in ref_name.lower(): matched_kind = "BillingAccount"
                        elif "network" in ref_name.lower(): matched_kind = "ComputeNetwork"
                        elif "bucket" in ref_name.lower(): matched_kind = "StorageBucket"
                        elif "kmskey" in ref_name.lower() or "cryptokey" in ref_name.lower(): matched_kind = "KMSCryptoKey"
                        elif "serviceaccount" in ref_name.lower(): matched_kind = "IAMServiceAccount"

                    if matched_kind and matched_kind in known_kinds and matched_kind != kind:
                        dependencies[kind].add(matched_kind)

    return dependencies, known_kinds

def parse_data(config_file_path, apis_dir, crds_dir, repo_root):
    resources = {}
    
    # Load existing data.json to preserve manual state
    existing_resources = {}
    data_json_path = os.path.join(repo_root, 'dev', 'migration-tracker', 'data.json')
    if os.path.exists(data_json_path):
        try:
            with open(data_json_path, 'r') as f:
                existing_list = json.load(f)
                for item in existing_list:
                    existing_resources[item['kind']] = item
        except Exception as e:
            print(f"Error loading existing data.json: {e}")

    with open(config_file_path, 'r') as f:
        config_lines = f.readlines()
        
    # Pre-fetch codebase state
    dependencies, known_kinds = build_dependency_graph(crds_dir)
    implemented_types = get_implemented_types(apis_dir)
    direct_controllers, fuzzers = find_direct_controllers(repo_root)
    test_fixtures = find_test_fixtures(repo_root)

    for line in config_lines:
        line = line.strip()
        if not line.startswith('{Group: '):
            continue
            
        group_match = re.search(r'Group:\s*"([^"]+)"', line)
        kind_match = re.search(r'Kind:\s*"([^"]+)"', line)
        default_ctrl_match = re.search(r'DefaultController:\s*k8s\.ReconcilerType([A-Za-z]+)', line)
        supported_ctrls_match = re.search(r'SupportedControllers:\s*\[\]k8s\.ReconcilerType\{(.*?)\}', line)
        
        if group_match and kind_match:
            group_full = group_match.group(1)
            group = group_full.split('.')[0]
            kind = kind_match.group(1)
            
            # 1. Start with existing resource data if available, or create default
            if kind in existing_resources:
                res = existing_resources[kind]
                # Update basic fields that might have changed structurally
                res['group'] = group
                res['version'] = "v1beta1"
            else:
                res = create_default_resource(kind, group)

            # 2. Update controller configuration from static_config.go
            if default_ctrl_match:
                res['defaultController'] = default_ctrl_match.group(1)
                res['controllerType'] = default_ctrl_match.group(1)
                
            supported = []
            if supported_ctrls_match:
                ctrls_raw = supported_ctrls_match.group(1)
                supported = re.findall(r'k8s\.ReconcilerType([A-Za-z]+)', ctrls_raw)
                if len(supported) == 1 and supported[0] == 'Direct':
                    continue
                res['supportedControllers'] = supported

            # 3. Auto-detect steps and states
            has_types = kind in implemented_types
            has_reference = False
            if has_types:
                for filepath in implemented_types[kind]:
                    ref_filepath = filepath.replace("_types.go", "_reference.go")
                    if os.path.exists(ref_filepath):
                        has_reference = True
                        break

            has_controller = kind in direct_controllers
            has_fuzzer = kind in fuzzers
            has_tests = kind.lower() in test_fixtures

            # Update steps (merging with existing)
            res['steps']['gen-types'] = res['steps'].get('gen-types', False) or has_types
            res['steps']['identity-reference'] = res['steps'].get('identity-reference', False) or has_reference
            res['steps']['controller'] = res['steps'].get('controller', False) or has_controller
            res['steps']['mapper-fuzzer'] = res['steps'].get('mapper-fuzzer', False) or has_fuzzer
            res['steps']['tests'] = res['steps'].get('tests', False)
            
            if 'Direct' in supported:
                res['steps']['mocks'] = res['steps'].get('mocks', True)
                res['steps']['tests'] = res['steps'].get('tests', True) or has_tests
                # A completed resource should have these true
                res['state'] = 'Completed'
            else:
                res['steps']['mocks'] = res['steps'].get('mocks', False)
                
                # Determine state
                any_step = (
                    res['steps'].get('gen-types', False) or
                    res['steps'].get('identity-reference', False) or
                    res['steps'].get('mapper-fuzzer', False) or
                    res['steps'].get('controller', False) or
                    res['steps'].get('tests', False) or
                    res['steps'].get('mocks', False)
                )
                if any_step:
                    res['state'] = 'In Progress'
                else:
                    res['state'] = 'Not Started'

            # Missing reference notes
            if has_types and not has_reference:
                res['notes'] = 'Missing _reference.go'
            elif res.get('notes') == 'Missing _reference.go':
                res['notes'] = ''

            resources[kind] = res

    # Calculate topological sort order and downstream count
    nodes = set(known_kinds.keys())
    local_graph = {node: set() for node in nodes}
    for node in nodes:
        if node in dependencies:
            local_graph[node] = {dep for dep in dependencies[node] if dep in nodes}

    in_degree = {node: len(local_graph[node]) for node in nodes}
    rev_graph = {node: set() for node in nodes}
    for node in nodes:
        for dep in local_graph[node]:
            rev_graph[dep].add(node)

    def get_all_downstream(n):
        visited = set()
        queue = [n]
        while queue:
            curr = queue.pop(0)
            for d in rev_graph[curr]:
                if d not in visited:
                    visited.add(d)
                    queue.append(d)
        return len(visited)

    downstream_counts = {node: get_all_downstream(node) for node in nodes}

    queue = [node for node in nodes if in_degree[node] == 0]
    queue.sort(key=lambda x: (-downstream_counts.get(x, 0), x))
    
    order_index = 1
    topo_order = {}

    while queue:
        current = queue.pop(0)
        topo_order[current] = order_index
        order_index += 1
        
        for dependent in sorted(list(rev_graph[current])):
            in_degree[dependent] -= 1
            if in_degree[dependent] == 0:
                queue.append(dependent)
        queue.sort(key=lambda x: (-downstream_counts.get(x, 0), x))

    for node in nodes:
        if node not in topo_order:
            topo_order[node] = order_index
            order_index += 1

    for kind, res in resources.items():
        res['sortOrder'] = topo_order.get(kind, 9999)
        res['downstreamCount'] = downstream_counts.get(kind, 0)
        
        if kind in dependencies:
            valid_deps = [dep for dep in dependencies[kind] if dep in resources]
            res['dependencies'] = sorted(valid_deps)

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
    repo_root = '../..'
    config_path = os.path.join(repo_root, 'pkg/controller/resourceconfig/static_config.go')
    apis_dir = os.path.join(repo_root, 'apis')
    crds_dir = os.path.join(repo_root, 'config/crds/resources')
    data = parse_data(config_path, apis_dir, crds_dir, repo_root)
    data = sorted(data, key=lambda x: x['kind'])
    with open('data.json', 'w') as f:
        json.dump(data, f, indent=2)
    print(f"Generated data.json with {len(data)} resources.")
