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
import re
import os
import sys
import yaml
from collections import defaultdict

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))

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

def get_implemented_controllers(direct_dir=None):
    if direct_dir is None:
        direct_dir = os.path.join(SCRIPT_DIR, "../../pkg/controller/direct")
    implemented_controllers = set()
    if not os.path.exists(direct_dir):
        return implemented_controllers
    for root, _, files in os.walk(direct_dir):
        for file in files:
            if file.endswith("_controller.go"):
                implemented_controllers.add(file)
    return implemented_controllers

def get_test_fixtures_and_http_logs(basic_dir=None):
    kinds_with_fixtures = set()
    # kinds_with_migration_fixtures tracks resources with fixtures that do NOT explicitly set
    # the 'reconciler: direct' annotation in create.yaml. These fixtures are used by the E2E
    # migration test suite (TestMigrationToDirect) to verify transitioning a resource from a legacy
    # controller (Terraform/DCL) to the Direct controller.
    kinds_with_migration_fixtures = set()
    kinds_with_mock_log = set()
    if basic_dir is None:
        basic_dir = os.path.join(SCRIPT_DIR, "../../pkg/test/resourcefixture/testdata/basic")
    if not os.path.exists(basic_dir):
        return kinds_with_fixtures, kinds_with_migration_fixtures, kinds_with_mock_log
    for root, _, files in os.walk(basic_dir):
        if 'create.yaml' in files:
            create_path = os.path.join(root, 'create.yaml')
            try:
                with open(create_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                    is_direct_first = ('alpha.cnrm.cloud.google.com/reconciler: direct' in content or 
                                       'cnrm.cloud.google.com/reconciler: direct' in content)
                    for doc in yaml.safe_load_all(content):
                        if doc and isinstance(doc, dict) and 'kind' in doc:
                            k = doc['kind']
                            kinds_with_fixtures.add(k)
                            if not is_direct_first:
                                kinds_with_migration_fixtures.add(k)
                            if '_http_mock.log' in files:
                                kinds_with_mock_log.add(k)
            except Exception:
                pass
    return kinds_with_fixtures, kinds_with_migration_fixtures, kinds_with_mock_log

def get_implemented_direct_go_files(direct_dir=None):
    if direct_dir is None:
        direct_dir = os.path.normpath(os.path.join(SCRIPT_DIR, "../../pkg/controller/direct"))
    group_direct_files = defaultdict(list)
    if not os.path.exists(direct_dir):
        return group_direct_files
    for root, _, files in os.walk(direct_dir):
        rel = os.path.relpath(root, direct_dir).lower().split(os.sep)
        group = rel[0] if rel and rel[0] != '.' else ''
        if not group:
            continue
        for f in files:
            f_lower = f.lower()
            if f_lower.endswith('.go'):
                path = os.path.join(root, f)
                try:
                    with open(path, 'r', encoding='utf-8') as file_obj:
                        content = file_obj.read()
                        group_direct_files[group].append((f_lower, content))
                except Exception:
                    pass
    return group_direct_files

def build_dependency_graph(crds_dir=None):
    if crds_dir is None:
        crds_dir = os.path.join(SCRIPT_DIR, "../../config/crds/resources")
    known_kinds = {}
    dependencies = defaultdict(set)
    crd_docs = []
    
    crd_versions = {}
    if not os.path.exists(crds_dir):
        return dependencies, known_kinds, crd_versions
        
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
                        crd_versions[kind] = [v["name"] for v in doc["spec"]["versions"]]
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

    return dependencies, known_kinds, crd_versions

def parse_data(config_file_path, apis_dir, crds_dir, direct_dir=None):
    if direct_dir is None:
        direct_dir = os.path.join(SCRIPT_DIR, "../../pkg/controller/direct")
    resources = {}
    
    dependencies, known_kinds, crd_versions = build_dependency_graph(crds_dir)

    # Load existing data to preserve manual updates
    existing_resources = {}
    data_json_path = os.path.join(SCRIPT_DIR, 'data.json')
    if os.path.exists(data_json_path):
        try:
            with open(data_json_path, 'r') as f:
                existing_list = json.load(f)
                existing_resources = {res['kind']: res for res in existing_list}
        except Exception as e:
            print(f"Warning: Could not load existing data.json: {e}", file=sys.stderr)
            
    # Overrides for brownfield resources that were originally generated via tf2crd or dcl2crd (supporting v1beta1),
    # but had their legacy controller entries cleaned up from static_config.go after 100% direct completion.
    # These must be tracked in data.json as completed brownfield migrations so we do not lose track of completed work.
    COMPLETED_BROWNFIELD_OVERRIDES = {
        "AlloyDBInstance": "alloydb",
        "CertificateManagerDNSAuthorization": "certificatemanager",
        "CloudIdentityGroup": "cloudidentity",
        "CloudIdentityMembership": "cloudidentity",
        "ComputeFirewallPolicyRule": "compute",
        "GKEHubFeatureMembership": "gkehub",
        "LoggingLogMetric": "logging",
        "MonitoringDashboard": "monitoring",
        "SQLInstance": "sql",
    }

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
            
            supported = []
            if supported_ctrls_match:
                ctrls_raw = supported_ctrls_match.group(1)
                supported = re.findall(r'k8s\.ReconcilerType([A-Za-z]+)', ctrls_raw)
            
            # Skip if it only supports Direct (born direct greenfield resources), UNLESS it is a known fully-migrated brownfield resource override
            if len(supported) == 1 and supported[0] == 'Direct' and kind not in COMPLETED_BROWNFIELD_OVERRIDES:
                continue

            # Skip if it is v1alpha1-only (only has v1alpha1 version in CRD)
            if kind in crd_versions and all(ver == 'v1alpha1' for ver in crd_versions[kind]) and kind not in COMPLETED_BROWNFIELD_OVERRIDES:
                continue
                
            resources[kind] = create_default_resource(kind, group)
            resources[kind]['supportedControllers'] = supported
            
            # Restore existing manual fields if they exist
            if kind in existing_resources:
                existing = existing_resources[kind]
                resources[kind]['notes'] = existing.get('notes', resources[kind]['notes'])
                resources[kind]['mocksLastRefreshed'] = existing.get('mocksLastRefreshed', resources[kind]['mocksLastRefreshed'])
                if 'edgeCases' in existing:
                    resources[kind]['edgeCases'] = existing['edgeCases']
            
            if default_ctrl_match:
                resources[kind]['defaultController'] = default_ctrl_match.group(1)
                resources[kind]['controllerType'] = default_ctrl_match.group(1)
                
            if resources[kind]['defaultController'] == 'Direct' or kind in COMPLETED_BROWNFIELD_OVERRIDES:
                resources[kind]['state'] = 'Completed'
                resources[kind]['steps'] = {
                    "gen-types": True,
                    "identity-reference": True,
                    "mapper-fuzzer": True,
                    "mocks": True,
                    "controller": True,
                    "tests": True
                }
            else:
                if resources[kind]['state'] == 'Completed':
                    resources[kind]['state'] = 'Not Started'

    # Ensure all COMPLETED_BROWNFIELD_OVERRIDES exist in resources even if missing from static_config.go
    for kind, group in COMPLETED_BROWNFIELD_OVERRIDES.items():
        if kind not in resources:
            resources[kind] = create_default_resource(kind, group)
            resources[kind]['supportedControllers'] = ['Direct']
            resources[kind]['defaultController'] = 'Direct'
            resources[kind]['controllerType'] = 'Direct'
            resources[kind]['state'] = 'Completed'
            resources[kind]['steps'] = {
                "gen-types": True,
                "identity-reference": True,
                "mapper-fuzzer": True,
                "mocks": True,
                "controller": True,
                "tests": True
            }

    dependencies, known_kinds, _ = build_dependency_graph(crds_dir)
    implemented_types = get_implemented_types(apis_dir)
    implemented_controllers = get_implemented_controllers(direct_dir)
    kinds_with_fixtures, kinds_with_migration_fixtures, kinds_with_mock_log = get_test_fixtures_and_http_logs()
    group_direct_files = get_implemented_direct_go_files(direct_dir)

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

    for node in sorted(nodes):
        if node not in topo_order:
            topo_order[node] = order_index
            order_index += 1

    for kind, res in resources.items():
        res['sortOrder'] = topo_order.get(kind, 9999)
        res['downstreamCount'] = downstream_counts.get(kind, 0)
        group_lower = res['group'].lower()
        kind_lower = kind.lower()

        if kind in dependencies:
            valid_deps = [dep for dep in dependencies[kind] if dep in resources]
            res['dependencies'] = sorted(valid_deps)

        has_std_identity = False
        has_std_reference = False
        non_std_id_ref_files = []

        if kind in implemented_types:
            res['steps']['gen-types'] = True

            std_id_filename = f"{kind_lower}_identity.go"
            std_ref_filename = f"{kind_lower}_reference.go"

            for filepath in implemented_types[kind]:
                dirpath = os.path.dirname(filepath)
                filename = os.path.basename(filepath)
                prefix = filename.replace("_types.go", "")

                if os.path.exists(os.path.join(dirpath, std_id_filename)):
                    has_std_identity = True

                if os.path.exists(os.path.join(dirpath, std_ref_filename)):
                    has_std_reference = True

                # Check for non-standard / shared identity or reference files in apis/<group>/<version>/
                possible_alt_names = [f"{prefix}_identity.go", "identity.go", f"{prefix}_reference.go", "reference.go"]
                for alt_name in possible_alt_names:
                    if alt_name not in (std_id_filename, std_ref_filename):
                        alt_path = os.path.join(dirpath, alt_name)
                        if os.path.exists(alt_path):
                            non_std_id_ref_files.append(alt_name)

            if has_std_identity and has_std_reference:
                res['steps']['identity-reference'] = True
                if res.get('notes') in ('Missing _reference.go', 'Missing _reference.go or _identity.go'):
                    res['notes'] = ''
            else:
                res['notes'] = 'Missing _reference.go or _identity.go'
                res['steps']['identity-reference'] = False

        std_mapper_file = os.path.join(SCRIPT_DIR, "../../pkg/controller/direct", group_lower, f"{kind_lower}_mapper.go")
        std_fuzzer_file = os.path.join(SCRIPT_DIR, "../../pkg/controller/direct", group_lower, f"{kind_lower}_fuzzer.go")

        has_mapper = os.path.exists(std_mapper_file)
        has_fuzzer = os.path.exists(std_fuzzer_file)

        if has_mapper and has_fuzzer:
            res['steps']['mapper-fuzzer'] = True
        else:
            res['steps']['mapper-fuzzer'] = False

        # Compute artifact metadata for architectural gap analysis
        non_std_files = []
        if group_lower in group_direct_files:
            mapper_regex = re.compile(rf'\b({re.escape(kind)}(Spec|Status|ObservedState)?(_)?(ToProto|FromProto|ToKRM|FromKRM|Mapper))\b')
            fuzzer_regex = re.compile(rf'\b(Fuzz{re.escape(kind)}|{re.escape(kind)}Fuzzer)\b')
            for filename, content in group_direct_files[group_lower]:
                if filename not in (f'{kind_lower}_mapper.go', f'{kind_lower}_fuzzer.go'):
                    if mapper_regex.search(content) or fuzzer_regex.search(content):
                        non_std_files.append(filename)

        res['artifacts'] = {
            'has_std_identity': has_std_identity,
            'has_std_reference': has_std_reference,
            'has_std_mapper': has_mapper,
            'has_std_fuzzer': has_fuzzer,
            'non_std_id_ref_files': sorted(list(set(non_std_id_ref_files))),
            'non_std_mapper_fuzzer_files': sorted(list(set(non_std_files))),
        }

        possible_controller_names = [
            f"{kind.lower()}_controller.go",
            f"{res['group'].lower()}{kind.lower()}_controller.go",
        ]
        if kind in implemented_types:
            for filepath in implemented_types[kind]:
                filename = os.path.basename(filepath)
                prefix = filename.replace("_types.go", "")
                possible_controller_names.append(f"{prefix}_controller.go")
        
        for ctrl_name in possible_controller_names:
            if ctrl_name in implemented_controllers:
                res['steps']['controller'] = True
                # Controller completion implies types, identity/ref, and mapper/fuzzer prerequisites
                res['steps']['gen-types'] = True
                res['steps']['identity-reference'] = True
                res['steps']['mapper-fuzzer'] = True
                break

        if kind in kinds_with_mock_log:
            res['steps']['mocks'] = True

        if 'Direct' in res.get('supportedControllers', []) and kind in kinds_with_migration_fixtures:
            res['steps']['tests'] = True

        if res['state'] != 'Completed' and any(res['steps'].values()):
            res['state'] = 'In Progress'

    # Sequentially index all resources consecutively to remove gaps in sortOrder,
    # and re-order sortOrder based on deprecation to place deprecated resources last
    active_kinds = []
    deprecated_kinds = []
    for kind, res in resources.items():
        is_deprecated = res.get('edgeCases', {}).get('gcpAPIDeprecated') == True
        if is_deprecated:
            deprecated_kinds.append(kind)
        else:
            active_kinds.append(kind)

    active_kinds.sort(key=lambda k: resources[k].get('sortOrder', 9999))
    deprecated_kinds.sort(key=lambda k: resources[k].get('sortOrder', 9999))

    current_order = 1
    for kind in active_kinds:
        resources[kind]['sortOrder'] = current_order
        current_order += 1
    for kind in deprecated_kinds:
        resources[kind]['sortOrder'] = current_order
        current_order += 1

    # Report any resources removed from data.json after this run
    if existing_resources:
        old_kinds = set(existing_resources.keys())
        new_kinds = set(resources.keys())
        removed_kinds = old_kinds - new_kinds
        if removed_kinds:
            print(f"\n[NOTICE] {len(removed_kinds)} resource(s) removed from tracked resource list after this run:", file=sys.stderr)
            for k in sorted(removed_kinds):
                grp = existing_resources[k].get('group', 'unknown')
                print(f"   - {k} (Group: {grp})", file=sys.stderr)
            print("", file=sys.stderr)

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
    config_path = os.path.join(SCRIPT_DIR, '../../pkg/controller/resourceconfig/static_config.go')
    apis_dir = os.path.join(SCRIPT_DIR, '../../apis')
    crds_dir = os.path.join(SCRIPT_DIR, '../../config/crds/resources')
    data = parse_data(config_path, apis_dir, crds_dir)
    data = sorted(data, key=lambda x: x['kind'])
    output_path = os.path.join(SCRIPT_DIR, 'data.json')
    with open(output_path, 'w') as f:
        json.dump(data, f, indent=2)
    print(f"Generated data.json at {output_path} with {len(data)} resources.")
