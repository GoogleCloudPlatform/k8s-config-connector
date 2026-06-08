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

def build_dependency_graph(crds_dir="../../config/crds/resources"):
    known_kinds = {}
    dependencies = defaultdict(set)
    crd_docs = []
    kind_versions = {}
    
    if not os.path.exists(crds_dir):
        return dependencies, known_kinds, kind_versions
        
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
        
        versions = [v["name"] for v in doc.get("spec", {}).get("versions", [])]
        best_v = versions[0] if versions else "v1beta1"
        has_beta = any("beta" in v for v in versions)
        has_alpha = any("alpha" in v for v in versions)
        if has_beta:
            best_v = next(v for v in versions if "beta" in v)
        elif has_alpha:
            best_v = next(v for v in versions if "alpha" in v)
        kind_versions[kind] = best_v

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

    return dependencies, known_kinds, kind_versions

def parse_data(config_file_path, apis_dir, crds_dir, data_file_path='data.json'):
    existing_data = {}
    if os.path.exists(data_file_path):
        with open(data_file_path, 'r') as f:
            try:
                loaded = json.load(f)
                for r in loaded:
                    existing_data[r['kind']] = r
            except Exception:
                pass

    resources = {}
    
    with open(config_file_path, 'r') as f:
        config_lines = f.readlines()
        
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
            
            if kind in existing_data:
                resources[kind] = existing_data[kind]
                resources[kind]['group'] = group
            else:
                resources[kind] = create_default_resource(kind, group)
            
            if default_ctrl_match:
                resources[kind]['defaultController'] = default_ctrl_match.group(1)
                resources[kind]['controllerType'] = default_ctrl_match.group(1)
                
            if supported_ctrls_match:
                ctrls_raw = supported_ctrls_match.group(1)
                supported = re.findall(r'k8s\.ReconcilerType([A-Za-z]+)', ctrls_raw)
                resources[kind]['supportedControllers'] = supported
                if 'Direct' in supported:
                    resources[kind]['state'] = 'Completed'
                    if 'steps' not in resources[kind]: resources[kind]['steps'] = {}
                    resources[kind]['steps'].update({
                        "gen-types": True,
                        "identity-reference": True,
                        "mapper-fuzzer": True,
                        "mocks": True,
                        "controller": True,
                        "tests": True
                    })

    dependencies, known_kinds, kind_versions = build_dependency_graph(crds_dir)
    implemented_types = get_implemented_types(apis_dir)

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

    repo_root = "../../"
    base_test_dir = os.path.join(repo_root, "pkg/test/resourcefixture/testdata/basic")

    for kind, res in resources.items():
        res['sortOrder'] = topo_order.get(kind, 9999)
        res['downstreamCount'] = downstream_counts.get(kind, 0)
        
        if kind in kind_versions:
            res['version'] = kind_versions[kind]
        
        if kind in dependencies:
            valid_deps = [dep for dep in dependencies[kind] if dep in resources]
            res['dependencies'] = sorted(valid_deps)

        if kind in implemented_types:
            has_reference = False
            for filepath in implemented_types[kind]:
                ref_filepath = filepath.replace("_types.go", "_reference.go")
                if os.path.exists(ref_filepath):
                    has_reference = True
                    break
            
            if not has_reference:
                res['notes'] = 'Missing _reference.go'
                if 'steps' in res:
                    res['steps']['identity-reference'] = False

        kind_lower = kind.lower()
        group_lower = res['group'].lower()
        
        # Determine gcpTestLocation
        gcp_test_dir = ""
        version = res.get('version', 'v1beta1')
        preferred_path = os.path.join(base_test_dir, group_lower, version, kind_lower)
        if os.path.isdir(preferred_path):
            gcp_test_dir = os.path.relpath(preferred_path, repo_root)
        else:
            if os.path.exists(base_test_dir):
                for root, dirs, files in os.walk(base_test_dir):
                    if os.path.basename(root) == kind_lower:
                        gcp_test_dir = os.path.relpath(root, repo_root)
                        break
        res['gcpTestLocation'] = gcp_test_dir

        # Determine mockgcpLocation
        mock_dir = ""
        expected_mock_dir = os.path.join(repo_root, "mockgcp", f"mock{group_lower}")
        if os.path.isdir(expected_mock_dir):
            if kind_lower.startswith(group_lower):
                suffix = kind_lower[len(group_lower):]
            else:
                suffix = kind_lower
            if not suffix:
                suffix = kind_lower
            
            found = False
            for fname in os.listdir(expected_mock_dir):
                if not fname.endswith(".go") or fname in ("service.go", "normalize.go", "register.go", "utils.go", "filter.go", "operations.go"):
                    continue
                name_part = fname[:-3]
                name_part = re.sub(r'v[0-9]+(alpha[0-9]*|beta[0-9]*)?', '', name_part)
                name_part = name_part.replace('_', '')
                
                if suffix == name_part or name_part == suffix + "s" or name_part == suffix + "es" or suffix.endswith(name_part):
                    found = True
                    break
            
            if found:
                mock_dir = os.path.relpath(expected_mock_dir, repo_root)
                
        res['mockgcpLocation'] = mock_dir

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
        "notes": "",
        "gcpTestLocation": "",
        "mockgcpLocation": ""
    }

if __name__ == "__main__":
    config_path = '../../pkg/controller/resourceconfig/static_config.go'
    apis_dir = '../../apis'
    crds_dir = '../../config/crds/resources'
    data_path = 'data.json'
    data = parse_data(config_path, apis_dir, crds_dir, data_path)
    data = sorted(data, key=lambda x: x['kind'])
    with open(data_path, 'w') as f:
        json.dump(data, f, indent=2)
    print(f"Generated data.json with {len(data)} resources.")
