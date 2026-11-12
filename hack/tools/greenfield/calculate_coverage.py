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

import argparse
import csv
import os
import shutil
import subprocess
import re
import sys
import json
import tarfile
from collections import Counter, defaultdict
from datetime import datetime

def run_shell(cmd, cwd=None):
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd=cwd)
    return result

def get_gcp_resources(googleapis_dir):
    resources = {} 
    service_rpcs = {} # Map service to its RPCs
    service_messages = {} # Map service_pkg to set of message names
    
    for root, _, files in os.walk(googleapis_dir):
        if "third_party" in root: continue
        for file in files:
            if not file.endswith(".proto"): continue
            path = os.path.join(root, file)
            with open(path, 'r', errors='ignore') as f:
                content = f.read()
                
                # Identify service from package
                pkg_match = re.search(r'package\s+([a-z0-9.]+);', content)
                service_pkg = pkg_match.group(1) if pkg_match else "unknown"
                if service_pkg not in service_rpcs:
                    service_rpcs[service_pkg] = set()
                if service_pkg not in service_messages:
                    service_messages[service_pkg] = set()
                
                # Extract RPCs in this file
                rpc_matches = re.findall(r'rpc\s+([A-Za-z0-9]+)', content)
                service_rpcs[service_pkg].update(rpc_matches)

                # Extract messages in this file
                msg_matches = re.findall(r'message\s+([A-Za-z0-9]+)', content)
                service_messages[service_pkg].update(msg_matches)

                # Split by possible resource markers to isolate blocks
                blocks = re.split(r'google\.api\.resource', content)
                for block in blocks:
                    type_match = re.search(r'type:\s*"([^"]+)"', block)
                    if type_match:
                        rtype = type_match.group(1)
                        if "/" not in rtype: continue
                        
                        # Determine if this is a standard message-level google.api.resource definition
                        is_definition = not block.startswith('_reference') and not block.startswith('_container') and not block.startswith('_definition')
                        
                        if is_definition:
                            if rtype not in resources or not resources[rtype].get('is_proto_defined', False):
                                service_name, name = rtype.split('/')
                                resources[rtype] = {
                                    'service': service_name.split('.')[0],
                                    'pkg': service_pkg,
                                    'name': name,
                                    'ops': set(),
                                    'patterns': [],
                                    'is_proto_defined': True
                                }
                            p_matches = re.findall(r'pattern:\s*"([^"]+)"', block)
                            resources[rtype]['patterns'].extend(p_matches)
                        else:
                            if rtype not in resources:
                                service_name, name = rtype.split('/')
                                resources[rtype] = {
                                    'service': service_name.split('.')[0],
                                    'pkg': service_pkg,
                                    'name': name,
                                    'ops': set(),
                                    'patterns': [],
                                    'is_proto_defined': False
                                }
                            p_matches = re.findall(r'pattern:\s*"([^"]+)"', block)
                            resources[rtype]['patterns'].extend(p_matches)

    # 3. Match ops to resources within their service package
    for rtype, info in resources.items():
        name = info['name']
        pkg = info['pkg']
        
        if info.get('is_proto_defined', False):
            # Standard gRPC/Proto resource: match RPCs in its home package and versioned sibling packages
            all_rpcs = set()
            if pkg in service_rpcs:
                all_rpcs.update(service_rpcs[pkg])
                
            pkg_prefix = pkg
            if '.common' in pkg:
                pkg_prefix = pkg.split('.common')[0]
            elif '.v' in pkg:
                m = re.search(r'\.v\d', pkg)
                if m:
                    pkg_prefix = pkg[:m.start()]
                    
            if pkg_prefix != pkg:
                for other_pkg, rpcs in service_rpcs.items():
                    if other_pkg.startswith(pkg_prefix) and other_pkg != pkg:
                        all_rpcs.update(rpcs)
                        
            create_variants = [
                f"Create{name}", f"Insert{name}", f"BatchCreate{name}", f"Upload{name}", f"Upsert{name}"
            ]
            for v in create_variants:
                if v in all_rpcs:
                    info['ops'].add('CREATE')
                    break
            read_variants = [
                f"Get{name}", f"Describe{name}", f"Read{name}", f"BatchGet{name}"
            ]
            for v in read_variants:
                if v in all_rpcs:
                    info['ops'].add('READ')
                    break
            update_variants = [
                f"Update{name}", f"Patch{name}", f"BatchUpdate{name}", f"Upsert{name}"
            ]
            for v in update_variants:
                if v in all_rpcs:
                    info['ops'].add('UPDATE')
                    break
            delete_variants = [
                f"Delete{name}", f"Finish{name}", f"Abort{name}", 
                f"Cancel{name}", f"Terminate{name}", f"Destroy{name}", 
                f"Disable{name}", f"Deactivate{name}", f"Remove{name}"
            ]
            for v in delete_variants:
                if v in all_rpcs:
                    info['ops'].add('DELETE')
                    break
        else:
            # REST-only or legacy resource: look up operations in request message names
            endpoint = info['service']
            
            # Map some service names if needed (e.g. sql to sqladmin/sql)
            service_keywords = [endpoint]
            if endpoint == "sql":
                service_keywords.append("sqladmin")
                
            # Collect all message names from matching packages
            all_msgs = set()
            for service_pkg, msgs in service_messages.items():
                if any(kw in service_pkg for kw in service_keywords):
                    all_msgs.update(msgs)
                    
            create_verbs = ['create', 'insert', 'upload', 'upsert']
            read_verbs = ['get', 'describe', 'read']
            update_verbs = ['update', 'patch', 'upsert']
            delete_verbs = ['delete', 'finish', 'abort', 'cancel', 'terminate', 'destroy', 'disable', 'deactivate', 'remove']
            
            for msg in all_msgs:
                if not msg.endswith('Request'): continue
                if name.lower() in msg.lower():
                    msg_lower = msg.lower()
                    if any(v in msg_lower for v in create_verbs):
                        info['ops'].add('CREATE')
                    if any(v in msg_lower for v in read_verbs):
                        info['ops'].add('READ')
                    if any(v in msg_lower for v in update_verbs):
                        info['ops'].add('UPDATE')
                    if any(v in msg_lower for v in delete_verbs):
                        info['ops'].add('DELETE')
                
    # Clear pkg for non-proto-defined (REST-only) resources as they do not have a true proto package
    for rtype, info in resources.items():
        if not info.get('is_proto_defined', False):
            info['pkg'] = ""
                
    return resources

def get_kcc_resources(kcc_dir):
    resources = []
    static_config_path = os.path.join(kcc_dir, "pkg/controller/resourceconfig/static_config.go")
    if os.path.exists(static_config_path):
        with open(static_config_path, 'r') as f:
            for line in f:
                m = re.search(r'Group:\s*"([a-z0-9.]+)"\s*,\s*Kind:\s*"([A-Za-z0-9]+)"', line)
                if m:
                    resources.append({
                        'group': m.group(1),
                        'kind': m.group(2)
                    })
    return resources

def match_resources(gcp_resources, kcc_resources):
    covered = set()
    kcc_map = {}
    for res in kcc_resources:
        service = res['group'].replace(".cnrm.cloud.google.com", "")
        if service not in kcc_map: kcc_map[service] = []
        kcc_map[service].append(res['kind'])

    service_aliases = {
        "cloudresourcemanager": ["resourcemanager"],
        "analyticshub": ["bigqueryanalyticshub"],
        "biglake": ["bigquerybiglake"],
        "container": ["container", "gke"],
        "sqladmin": ["sql"],
        "cloudquota": ["cloudquotas"],
        "cloudids": ["ids"],
        "file": ["filestore"],
        "bigtableadmin": ["bigtable"],
        "cloudbilling": ["billing"],
        "cloudkms": ["kms"],
        "aiplatform": ["aiplatform", "vertexai"],
        "cloudasset": ["asset"],
        "batch": ["cloudbatch"],
    }

    kcc_kind_aliases = {
        "aiplatform/EntityType": ["VertexAIFeaturestoreEntityType"],
        "aiplatform/Feature": ["VertexAIFeaturestoreEntityTypeFeature"],
        "aiplatform/Index": ["VertexAIIndex"],
        "aiplatform/IndexEndpoint": ["VertexAIIndexEndpoint"],
    }

    for gcp_type, info in gcp_resources.items():
        gcp_service_base = info['service']
        gcp_name = info['name']
        matched_kcc_service = None
        if gcp_type == "iam.googleapis.com/AccessPolicy":
            matched_kcc_service = "accesscontextmanager"
        elif gcp_service_base in kcc_map:
            matched_kcc_service = gcp_service_base
        elif gcp_service_base.endswith('s') and gcp_service_base[:-1] in kcc_map:
            matched_kcc_service = gcp_service_base[:-1]
        elif gcp_service_base + 's' in kcc_map:
            matched_kcc_service = gcp_service_base + 's'
        else:
            for canonical_name, aliases in service_aliases.items():
                all_variants = [canonical_name] + aliases
                if gcp_service_base in all_variants:
                    for variant in all_variants:
                        if variant in kcc_map:
                            matched_kcc_service = variant
                            break
                    if matched_kcc_service:
                        break
        
        if not matched_kcc_service: continue
        name_norm = gcp_name.lower()
        for kcc_kind in kcc_map[matched_kcc_service]:
            kind_norm = kcc_kind.lower()
            prefixes = [matched_kcc_service.replace("-", "").lower(), gcp_service_base.replace("-", "").lower(), "gcp", "google", "cloud", "bigquery", "api"]
            if kind_norm == name_norm:
                covered.add(gcp_type)
                break
            found_prefix_match = False
            for p in prefixes:
                if kind_norm.startswith(p) and len(kind_norm) > len(p):
                    if kind_norm[len(p):] == name_norm:
                        covered.add(gcp_type)
                        found_prefix_match = True
                        break
            if found_prefix_match: break
            if kind_norm.endswith('s') and kind_norm[:-1] == name_norm:
                covered.add(gcp_type)
                break
    return covered

def is_leaf(patterns):
    for p in patterns:
        segments = [s for s in p.split('/') if s]
        if len(segments) == 4:
            if segments[0] in ["projects", "folders", "organizations"]:
                return True
        if len(segments) == 6:
            if segments[0] == "projects" and segments[2] in ["locations", "regions"]:
                return True
    return False

def is_next_layer(patterns):
    for p in patterns:
        segments = [s for s in p.split('/') if s]
        if len(segments) == 6:
            if segments[0] in ["projects", "folders", "organizations"] and segments[2] not in ["locations", "regions"]:
                return True
        if len(segments) == 8:
            if segments[0] == "projects" and segments[2] in ["locations", "regions"]:
                return True
    return False

def is_next_next_layer(patterns):
    for p in patterns:
        segments = [s for s in p.split('/') if s]
        if len(segments) == 8:
            if segments[0] in ["projects", "folders", "organizations"] and segments[2] not in ["locations", "regions"]:
                return True
        if len(segments) == 10:
            if segments[0] == "projects" and segments[2] in ["locations", "regions"]:
                return True
    return False

def get_layer(patterns, ops):
    has_create = 'CREATE' in ops
    has_delete = 'DELETE' in ops

    if has_create and has_delete:
        if is_leaf(patterns):
            return "Easy (Leaf)"
        elif is_next_layer(patterns):
            return "Next Layer (1 Parent)"
        elif is_next_next_layer(patterns):
            return "Next Next Layer (2 Parents)"
        else:
            return "Deeply Nested"

    if len(ops) > 0:
        if not has_create:
            return "Adoptable"
        else:
            return "Partially Manageable"

    return "Not Manageable"

SERVICE_PREFIXES_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                                    "..", "..", "..", "dev", "tools", "controllerbuilder", "serviceprefixes")

def load_service_prefixes():
    """Returns {"prefixes": {api: servicePrefix}, "acronyms": {word: replacement}}.

    Prefixes come from service-prefixes.yaml (approved) and
    service-prefix-pending.yaml (proposed); acronyms from acronyms.yaml. These
    files have a fixed, hand-maintained layout, so they are parsed line by line
    to avoid a PyYAML dependency.
    """
    prefixes = {}
    for name in ("service-prefixes.yaml", "service-prefix-pending.yaml"):
        api = None
        with open(os.path.join(SERVICE_PREFIXES_DIR, name), 'r') as f:
            for line in f:
                m = re.match(r'^  - api:\s*(\S+)\s*$', line)
                if m:
                    api = m.group(1)
                    continue
                m = re.match(r'^    servicePrefix:\s*(\S+)\s*$', line)
                if m and api:
                    prefixes.setdefault(api, m.group(1))

    acronyms = {}
    in_acronyms = False
    with open(os.path.join(SERVICE_PREFIXES_DIR, "acronyms.yaml"), 'r') as f:
        for line in f:
            if line.startswith("acronyms:"):
                in_acronyms = True
                continue
            m = re.match(r'^  (\S+):\s*(\S+)\s*$', line)
            if in_acronyms and m:
                acronyms[m.group(1)] = m.group(2)

    return {"prefixes": prefixes, "acronyms": acronyms}

def load_deprecated_services():
    deprecated_file = os.path.join(os.path.dirname(__file__), "deprecated_services.json")
    deprecated_set = set()
    if os.path.exists(deprecated_file):
        with open(deprecated_file, 'r') as f:
            deprecated_data = json.load(f)
            for item in deprecated_data:
                service_base = item["service"].replace(".googleapis.com", "").lower()
                for kind in item["kinds"]:
                    deprecated_set.add((service_base, kind.lower()))
    return deprecated_set

service_prefixes_config = load_service_prefixes()
service_prefixes = service_prefixes_config.get("prefixes", {})
acronym_corrections = service_prefixes_config.get("acronyms", {})

def get_predicted_kcc_kind(service, gcp_name):
    prefix = service_prefixes.get(service)
    if not prefix:
        parts = re.split(r'[-_]', service)
        prefix = "".join(p.capitalize() for p in parts)

    # Standardize names that start with a lowercase prefix to their proper CamelCase prefix
    if gcp_name.lower().startswith(prefix.lower()):
        kind = prefix + gcp_name[len(prefix):]
    else:
        kind = f"{prefix}{gcp_name}"

    # Apply global case-sensitive acronym and product corrections
    for old_acr, new_acr in acronym_corrections.items():
        kind = re.sub(fr'{old_acr}([A-Z]|$)', fr'{new_acr}\1', kind)

    return kind

def prepare_repo(repo_url, target_dir, sha):
    if not os.path.exists(target_dir):
        run_shell(f"git clone --depth 100 {repo_url} {target_dir}")
    run_shell("git fetch origin master --depth 100", cwd=target_dir)
    run_shell(f"git checkout -f {sha}", cwd=target_dir)

def singularize_name(name):
    """Singularizes a resource name cleanly."""
    if name.endswith('ies'):
        return name[:-3] + 'y'
    if name.endswith('es') and name[:-2].endswith(('s', 'sh', 'ch', 'x', 'z')):
        return name[:-2]
    if name.endswith('s') and not name.endswith(('ss', 'us')):
        return name[:-1]
    return name

def canonicalize_resource_name(name):
    """Strips hierarchical prefixes (Global, Regional, Zonal) from resource names and singularizes them."""
    prefixes = ["Global", "Regional", "Zonal"]
    for p in prefixes:
        if name.startswith(p) and len(name) > len(p) and name[len(p)].isupper():
            name = name[len(p):]
            break
    return singularize_name(name)

def unify_hierarchies(resources):
    """Groups GCP resources by their canonical name to avoid overcounting."""
    unified = {}
    for rtype, info in resources.items():
        service = info['service']
        canonical_name = canonicalize_resource_name(info['name'])
        key = f"{service}/{canonical_name}"
        
        if key not in unified:
            unified[key] = {
                'service': service,
                'pkg': info['pkg'],
                'name': canonical_name,
                'ops': set(),
                'patterns': [],
                'rtypes': [] # Keep track of original rtypes
            }
        
        unified[key]['ops'].update(info['ops'])
        unified[key]['patterns'].extend(info['patterns'])
        unified[key]['rtypes'].append(rtype)

    for key in unified:
        unified[key]['patterns'] = list(dict.fromkeys(unified[key]['patterns']))

    return unified

def run_gap_analysis(googleapis_sha, kcc_sha, k, update_gap, resources_list, googleapis_dir, kcc_dir):
    """Legacy gap analysis: CRUD-based categories, gap_analysis.txt and resources_list.json."""
    gcp_raw = get_gcp_resources(googleapis_dir)
    
    # Apply Skip List before unification
    skip_file = os.path.join(os.path.dirname(__file__), "coverage_skip.json")
    skipped_count = 0
    gcp_filtered = {}
    if os.path.exists(skip_file):
        with open(skip_file, 'r') as f:
            skip_data = json.load(f)
            patterns = [re.compile(s['pattern']) for s in skip_data.get('skips', [])]
            
            for rtype, info in gcp_raw.items():
                is_skipped = False
                for p in patterns:
                    if p.match(rtype):
                        is_skipped = True
                        skipped_count += 1
                        break
                if not is_skipped:
                    gcp_filtered[rtype] = info
    else:
        gcp_filtered = gcp_raw

    # Unify hierarchies (Global/Regional/Zonal -> Base)
    gcp_resources = unify_hierarchies(gcp_filtered)
    all_gcp_raw_count = len(gcp_raw)

    kcc_resources = get_kcc_resources(kcc_dir)

    # match_resources needs to work with unified keys now
    covered = set()
    key_to_kcc_kind = {}
    kcc_map = {}
    for res in kcc_resources:
        service = res['group'].replace(".cnrm.cloud.google.com", "")
        if service not in kcc_map: kcc_map[service] = []
        kcc_map[service].append(res['kind'])

    service_aliases = {
        "cloudresourcemanager": ["resourcemanager"],
        "analyticshub": ["bigqueryanalyticshub"],
        "biglake": ["bigquerybiglake"],
        "container": ["container", "gke"],
        "sqladmin": ["sql"],
        "cloudquota": ["cloudquotas"],
        "cloudids": ["ids"],
        "file": ["filestore"],
        "bigtableadmin": ["bigtable"],
        "cloudbilling": ["billing"],
        "cloudkms": ["kms"],
        "aiplatform": ["aiplatform", "vertexai"],
        "cloudasset": ["asset"],
        "batch": ["cloudbatch"],
    }

    kcc_kind_aliases = {
        "aiplatform/EntityType": ["VertexAIFeaturestoreEntityType"],
        "aiplatform/Feature": ["VertexAIFeaturestoreEntityTypeFeature"],
        "aiplatform/Index": ["VertexAIIndex"],
        "aiplatform/IndexEndpoint": ["VertexAIIndexEndpoint"],
    }

    for key, info in gcp_resources.items():
        if key in kcc_kind_aliases:
            found_alias = False
            for alias_kind in kcc_kind_aliases[key]:
                for service_kinds in kcc_map.values():
                    if alias_kind in service_kinds:
                        covered.add(key)
                        key_to_kcc_kind[key] = alias_kind
                        found_alias = True
                        break
                if found_alias: break
            if found_alias: continue
            
        gcp_service_base = info['service']
        gcp_name = info['name']
        matched_kcc_services = []
        if key == "iam/AccessPolicy":
            matched_kcc_services = ["accesscontextmanager"]
        else:
            variants = [gcp_service_base]
            if gcp_service_base.endswith('s'):
                variants.append(gcp_service_base[:-1])
            variants.append(gcp_service_base + 's')
            
            for canonical_name, aliases in service_aliases.items():
                all_variants = [canonical_name] + aliases
                if gcp_service_base in all_variants:
                    variants.extend(all_variants)
            
            for v in set(variants):
                if v in kcc_map:
                    matched_kcc_services.append(v)
        
        if not matched_kcc_services: continue
        name_norm = gcp_name.lower()
        
        all_service_prefixes = [s.replace("-", "").lower() for s in variants]
        
        for matched_kcc_service in matched_kcc_services:
            for kcc_kind in kcc_map[matched_kcc_service]:
                kind_norm = kcc_kind.lower()
                prefixes = all_service_prefixes + [gcp_service_base.replace("-", "").lower(), "gcp", "google", "cloud", "bigquery", "api"]
                if kind_norm == name_norm:
                    covered.add(key)
                    key_to_kcc_kind[key] = kcc_kind
                    break
                found_prefix_match = False
                for p in prefixes:
                    if kind_norm.startswith(p) and len(kind_norm) > len(p):
                        if kind_norm[len(p):] == name_norm:
                            covered.add(key)
                            key_to_kcc_kind[key] = kcc_kind
                            found_prefix_match = True
                            break
                if found_prefix_match: break
                if kind_norm.endswith('s') and kind_norm[:-1] == name_norm:
                    covered.add(key)
                    key_to_kcc_kind[key] = kcc_kind
                    break
            if key in covered: break
    
    # Categorization
    all_gcp_keys = set(gcp_resources.keys())
    covered = {key for key in covered if key in all_gcp_keys}
    missing = all_gcp_keys - covered

    deprecated_services = load_deprecated_services()
    deprecated_keys = set()
    for key, info in gcp_resources.items():
        if key in key_to_kcc_kind:
            kcc_kind = key_to_kcc_kind[key]
        else:
            kcc_kind = get_predicted_kcc_kind(info['service'], info['name'])
        if (info['service'].lower(), kcc_kind.lower()) in deprecated_services:
            deprecated_keys.add(key)

    implemented_deprecated = covered & deprecated_keys
    missing_deprecated = missing & deprecated_keys

    active_covered = covered - deprecated_keys
    active_missing = missing - deprecated_keys
    
    # Calculate how many raw rtypes were unified
    total_raw_rtypes = sum(len(info['rtypes']) for info in gcp_resources.values())
    unification_count = total_raw_rtypes - len(gcp_resources)

    # Manageable = Has at least one CRUD function (excluding deprecated keys)
    manageable_gcp = {key for key, info in gcp_resources.items() if len(info['ops']) > 0}
    missing_manageable = manageable_gcp - covered - deprecated_keys

    # Fully Manageable = Has both CREATE and DELETE functions (excluding deprecated keys)
    fully_manageable_gcp = {key for key, info in gcp_resources.items() if 'CREATE' in info['ops'] and 'DELETE' in info['ops']}
    missing_fully_manageable = fully_manageable_gcp - covered - deprecated_keys

    # Partially Manageable = Has CREATE but lacks DELETE (excluding deprecated keys)
    partially_manageable_gcp = {key for key, info in gcp_resources.items() if 'CREATE' in info['ops'] and 'DELETE' not in info['ops']}
    missing_partially_manageable = partially_manageable_gcp - covered - deprecated_keys

    # Adoptable = Has no CREATE but has at least one other CRUD function (excluding deprecated keys)
    adoptable_gcp = {key for key, info in gcp_resources.items() if 'CREATE' not in info['ops'] and len(info['ops']) > 0}
    missing_adoptable = adoptable_gcp - covered - deprecated_keys

    # Easy = Fully Manageable AND Leaf pattern
    missing_easy = {key for key in missing_fully_manageable if is_leaf(gcp_resources[key]['patterns'])}
    
    # Next Layer = Fully Manageable AND 1 parent deep
    missing_next_layer = {key for key in missing_fully_manageable if is_next_layer(gcp_resources[key]['patterns'])}
    
    # Next Next Layer = Fully Manageable AND 2 parents deep
    missing_next_next_layer = {key for key in missing_fully_manageable if is_next_next_layer(gcp_resources[key]['patterns'])}
    
    # Generate Gap Analysis Table for tracking
    gap_file = os.path.join(os.path.dirname(__file__), "gap_analysis.txt")
    list_file = os.path.join(os.path.dirname(__file__), "resources_list.json")

    total_manageable = len(covered) + len(missing_manageable)
    manageable_coverage = len(covered) / max(1, total_manageable)
    
    if update_gap:
        now = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

        missing_non_manageable = len(active_missing - missing_manageable)
        missing_deeply_nested = len(missing_fully_manageable - missing_easy - missing_next_layer - missing_next_next_layer)

        analysis_lines = [
            f"Gap Analysis Snapshot - {now}",
            f"GoogleAPIs SHA: {googleapis_sha}",
            f"KCC SHA:        {kcc_sha}",
            "-" * 55,
            f"{'Metric':<30} | {'Value':<10}",
            "-" * 55,
            f"{'Total GCP Resources (Raw)':<30} | {all_gcp_raw_count:<10}",
            f"{'  ├── Processed (Unified)':<30} | {len(all_gcp_keys):<10}",
            f"{'  ├── Deprecated Resources':<30} | {len(deprecated_keys):<10}",
            f"{'  ├── Skipped Resources':<30} | {skipped_count:<10}",
            f"{'  └── Unified (Hierarchical)':<30} | {unification_count:<10}",
            "-" * 55,
            f"{'Implemented in KCC':<30} | {len(active_covered):<10}",
            f"{'Missing from KCC':<30} | {len(active_missing):<10}",
            f"{'  ├── Unmanageable':<30} | {missing_non_manageable:<10}",
            f"{'  └── Manageable':<30} | {len(missing_manageable):<10}",
            f"{'        ├── Adoptable':<30} | {len(missing_adoptable):<10}",
            f"{'        ├── Partially Manageable':<30} | {len(missing_partially_manageable):<10}",
            f"{'        └── Fully Manageable':<30} | {len(missing_fully_manageable):<10}",
            f"{'              ├── Easy (Leaf)':<30} | {len(missing_easy):<10}",
            f"{'              ├── Next Layer':<30} | {len(missing_next_layer):<10}",
            f"{'              ├── Next Next Layer':<30} | {len(missing_next_next_layer):<10}",
            f"{'              └── Deeply Nested':<30} | {missing_deeply_nested:<10}",
            "-" * 55,
            f"{'Total API Coverage':<30} | {len(covered)/max(1, len(all_gcp_keys)):.2%}",
            f"{'Manageable Coverage':<30} | {manageable_coverage:.2%}",
            ""
        ]
        
        with open(gap_file, 'w') as f:
            f.write("\n".join(analysis_lines))

    if resources_list:
        coverage_list = []
        for key, info in sorted(gcp_resources.items()):
            available_ops = sorted(list(info['ops']))
            missing_ops = sorted(list({'CREATE', 'READ', 'UPDATE', 'DELETE'} - info['ops']))
            layer = get_layer(info['patterns'], info['ops'])
            status = "deprecated" if key in deprecated_keys else ("implemented" if key in covered else "missing")

            if key in key_to_kcc_kind:
                kcc_kind = key_to_kcc_kind[key]
            else:
                kcc_kind = get_predicted_kcc_kind(info['service'], info['name'])

            coverage_list.append({
                "resource": key,
                "service": info['service'],
                "kind": kcc_kind,
                "patterns": info['patterns'],
                "layer": layer,
                "available_ops": available_ops,
                "missing_ops": missing_ops,
                "status": status
            })

        with open(list_file, 'w') as f:
            json.dump(coverage_list, f, indent=2)

    if update_gap:
        print(f"\n[SAVED] Gap analysis snapshot written to {gap_file}")
    if resources_list:
        print(f"[SAVED] All resources coverage list written to {list_file}")

    print(f"\n--- Top {k} Missing Manageable Resources ---")
    
    def sort_key(m):
        if m in missing_easy: return 0
        if m in missing_next_layer: return 1
        if m in missing_next_next_layer: return 2
        if m in missing_fully_manageable: return 3
        return 4
        
    for m in sorted(list(missing_manageable), key=lambda x: (sort_key(x), x))[:k]:
        patterns = ", ".join(gcp_resources[m]['patterns'])
        rtypes = ", ".join(gcp_resources[m]['rtypes'])
        layer = "Unknown"
        if m in missing_easy:
            layer = "Easy (Leaf)"
        elif m in missing_next_layer:
            layer = "Next Layer (1 Parent)"
        elif m in missing_next_next_layer:
            layer = "Next Next Layer (2 Parents)"
        elif m in missing_fully_manageable:
            layer = "Deeply Nested"
        elif m in missing_adoptable:
            layer = "Adoptable"
        else:
            layer = "Partially Manageable"

        available_ops = gcp_resources[m]['ops']
        missing_ops = {'CREATE', 'READ', 'UPDATE', 'DELETE'} - available_ops
        print(f"  - {m}")
        print(f"    Patterns:      {patterns}")
        print(f"    ProtoPath:     {rtypes}")
        print(f"    Layer:         {layer}")
        print(f"    Available Ops: {sorted(list(available_ops)) if available_ops else 'None'}")
        print(f"    Missing Ops:   {sorted(list(missing_ops)) if missing_ops else 'None'}")

# =============================================================================
# Manageability mode (--manageability)
#
# Classifies every resource in googleapis as Fully Manageable, Partially
# Manageable (no DELETE / no CREATE), Needs review or Dropped, and every
# v1alpha1 TF/DCL kind without a googleapis match as Brownfield. Each dropped
# row carries the reason and a link (GitHub permalink) to the rule or data that
# dropped it. See README.md ("Manageability Mode") for the rules.
#
# Unlike the legacy gap analysis above, this mode reads a protoc-compiled
# descriptor set (exact RPC names, HTTP bindings, request fields and
# deprecation options) and re-implements the gates that controllerbuilder
# enforces in generate-types (non-GA, deprecated, non-declarative) and the
# server-generated ID check in generate-identity, reading their data straight
# from the KCC checkout so the list stays in sync with the tooling.
# =============================================================================

KCC_GITHUB = "https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/%s/%s"
GOOGLEAPIS_GITHUB = "https://github.com/googleapis/googleapis/blob/%s/%s"
GREENFIELD_REL = "hack/tools/greenfield"
ND_REL = "dev/tools/controllerbuilder/pkg/commands/generatetypes/nondeclarative.go"
DEP_REL = "dev/tools/controllerbuilder/pkg/commands/generatetypes/deprecated.go"
# Mirrors isNonGAProto in dev/tools/controllerbuilder/pkg/commands/generatetypes/generatetypescommand.go.
NONGA_SEGMENT_RE = re.compile(r"(?i)^v[0-9]+[a-z0-9]*(alpha|beta)[a-z0-9]*$")
VERSION_SEGMENT_RE = re.compile(r"^v\d+[a-z0-9]*$")
PB_TYPE_STRING = 9
PB_TYPE_MESSAGE = 11

BUCKET_FM = "Fully Manageable"
BUCKET_PM_NO_DELETE = "Partially Manageable: no DELETE"
BUCKET_PM_NO_CREATE = "Partially Manageable: no CREATE"
BUCKET_REVIEW = "Needs review"
BUCKET_BROWNFIELD = "Brownfield (no googleapis resource)"
BUCKET_DROPPED = "Dropped"
BUCKET_ORDER = [BUCKET_FM, BUCKET_PM_NO_DELETE, BUCKET_PM_NO_CREATE, BUCKET_REVIEW, BUCKET_BROWNFIELD, BUCKET_DROPPED]
STATUS_NET_NEW = "net-new"
STATUS_BROWNFIELD = "brownfield (v1alpha1 TF/DCL)"
STATUS_TYPES_ONLY = "in progress (types only)"
STATUS_DIRECT = "implemented (direct)"
STATUS_V1BETA1 = "v1beta1"
STATUS_ORDER = [STATUS_NET_NEW, STATUS_BROWNFIELD, STATUS_TYPES_ONLY, STATUS_DIRECT, STATUS_V1BETA1]
# Drop reasons in precedence order; policy:<reason> drops (policy_exclusions.json) sort last.
DROP_ORDER = ["skip-list", "example-api", "deprecated", "non-ga-only", "v1beta1-brownfield", "no-crud", "read-only",
              "nondeclarative:pattern", "nondeclarative:bulk-rpc"]

MANAGEABILITY_COLUMNS = [
    "id", "bucket", "kcc_status", "api_group", "api_family", "kind", "resource_type", "service", "api_area", "layer",
    "ops", "flags", "drop_reason", "drop_detail", "drop_link", "all_drop_reasons", "server_generated_id",
    "nondeclarative_go", "kcc_kinds", "kcc_match", "kcc_link", "in_cp_list", "proto_message", "proto_link",
    "other_packages", "patterns", "in_local_googleapis", "first_googleapis_sha", "first_googleapis_link", "rpcs",
    "unannotated", "ga", "predicted_kind",
]

# API groups: Google Cloud is google.cloud.* plus the Cloud APIs whose protos live in other packages.
API_GROUP_CLOUD = "Google Cloud"
API_GROUP_OTHER = "Other Google APIs"
API_GROUP_ORDER = [API_GROUP_CLOUD, API_GROUP_OTHER]
CLOUD_PACKAGE_PREFIXES = (
    "google.cloud.", "google.spanner.", "google.bigtable.", "google.iam.", "google.storage.", "google.storagetransfer.",
    "google.logging.", "google.monitoring.", "google.pubsub.", "google.firestore.", "google.datastore.",
    "google.appengine.", "google.container.", "google.privacy.dlp.", "google.identity.accesscontextmanager.",
    "google.api.",  # API Keys, Cloud Quotas, Service Usage
    "google.devtools.",  # Artifact Registry, Cloud Build, Cloud Trace, Profiler, Error Reporting, ...
    "grafeas.",  # Container Analysis
)
OTHER_API_FAMILIES = (  # everything else is "Misc" (Health, Gemini API, Actions, Home, ChromeOS, ...)
    ("google.apps.", "Workspace"), ("google.chat.", "Workspace"),
    ("google.ads.", "Ads & Marketing"), ("google.marketingplatform.", "Ads & Marketing"),
    ("google.analytics.", "Ads & Marketing"),
    ("google.shopping.", "Shopping"),
    ("google.maps.", "Maps"), ("maps.", "Maps"), ("google.streetview.", "Maps"),
)


def api_group_of(pkg):
    """-> (api_group, api_family). Rows without a proto package are KCC (Google Cloud) kinds."""
    if not pkg or pkg.startswith(CLOUD_PACKAGE_PREFIXES):
        return API_GROUP_CLOUD, ""
    for prefix, family in OTHER_API_FAMILIES:
        if pkg.startswith(prefix):
            return API_GROUP_OTHER, family
    return API_GROUP_OTHER, "Misc"

# googleapis resource type -> KCC kinds whose names don't follow the proto message name (added to other matches).
MANUAL_KIND_ALIASES = {
    "iam.googleapis.com/Role": ["IAMCustomRole"],
    "sqladmin.googleapis.com/DatabaseInstance": ["SQLInstance"],
    "sourcerepo.googleapis.com/Repo": ["SourceRepoRepository"],
    "storagetransfer.googleapis.com/TransferJob": ["StorageTransferJob"],
    "compute.googleapis.com/NetworkEndpointGroup": ["ComputeNetworkEndpointGroup", "ComputeRegionNetworkEndpointGroup",
                                                    "ComputeGlobalNetworkEndpointGroup"],
    "grafeas.io/Note": ["ContainerAnalysisNote"],
    "grafeas.io/Occurrence": ["ContainerAnalysisOccurrence"],
    "cloudasset.googleapis.com/Feed": ["CloudAssetFolderFeed", "CloudAssetOrganizationFeed", "CloudAssetProjectFeed"],
}
# googleapis resource type -> the exact KCC kinds (replaces all other matches).
MANUAL_KIND_OVERRIDES = {
    "cloudfunctions.googleapis.com/CloudFunction": ["CloudFunctionsFunction"],  # Cloud Functions v1
    "cloudfunctions.googleapis.com/Function": ["CloudFunctions2Function"],  # Cloud Functions v2
}
# googleapis service -> KCC API group prefixes, used to find TF/DCL kinds (which have no +kcc:proto annotations).
MATCH_SERVICE_ALIASES = {
    "cloudresourcemanager": ["resourcemanager", "tags"], "analyticshub": ["bigqueryanalyticshub"],
    "biglake": ["bigquerybiglake"], "container": ["container", "gke"], "sqladmin": ["sql"],
    "cloudquota": ["cloudquotas"], "cloudids": ["ids", "cloudids"], "file": ["filestore"],
    "bigtableadmin": ["bigtable"], "cloudbilling": ["billing", "billingbudgets"], "cloudkms": ["kms"],
    "aiplatform": ["aiplatform", "vertexai"], "cloudasset": ["asset"], "batch": ["cloudbatch"],
    "bigquerydatatransfer": ["bigquerydatatransfer"], "cloudscheduler": ["cloudscheduler"],
    "cloudfunctions": ["cloudfunctions", "cloudfunctions2"], "run": ["run"], "iam": ["iam"],
    "storagetransfer": ["storagetransfer"], "cloudbuild": ["cloudbuild"], "cloudidentity": ["cloudidentity"],
    "osconfig": ["osconfig"], "datacatalog": ["datacatalog"], "monitoring": ["monitoring"],
}
# Services skipped when synthesizing resources for unannotated GA APIs (shared types, Ads, LRO, IAM policy).
SYNTH_SKIP_PACKAGE_PREFIXES = ("google.ads.", "google.api", "google.longrunning", "google.iam.v1", "grafeas")
# Parent collections that don't count towards the layer depth.
LAYER_SKIP_PARENTS = {"projects", "folders", "organizations", "billingAccounts", "locations", "regions", "zones",
                      "global"}


# --- Minimal, dependency-free decoder for a protoc FileDescriptorSet ---------
def _pb_varint(b, i):
    r = 0
    s = 0
    while True:
        x = b[i]
        i += 1
        r |= (x & 0x7F) << s
        if x < 0x80:
            return r, i
        s += 7


def _pb_fields(b):
    i = 0
    n = len(b)
    while i < n:
        key, i = _pb_varint(b, i)
        fn, wt = key >> 3, key & 7
        if wt == 0:
            v, i = _pb_varint(b, i)
        elif wt == 2:
            ln, i = _pb_varint(b, i)
            v = b[i:i + ln]
            i += ln
        elif wt == 1:
            v = b[i:i + 8]
            i += 8
        elif wt == 5:
            v = b[i:i + 4]
            i += 4
        else:
            raise ValueError("unsupported protobuf wire type %d" % wt)
        yield fn, wt, v


def _pb_str(v):
    return bytes(v).decode("utf-8", "replace")


def _pb_packed(v):
    out = []
    i = 0
    while i < len(v):
        x, i = _pb_varint(v, i)
        out.append(x)
    return out


def _pb_resource(v):
    r = {"type": "", "patterns": [], "plural": ""}
    for fn, _, x in _pb_fields(v):
        if fn == 1:
            r["type"] = _pb_str(x)
        elif fn == 2:
            r["patterns"].append(_pb_str(x))
        elif fn == 5:
            r["plural"] = _pb_str(x)
    return r


_PB_HTTP_VERBS = {2: "GET", 3: "PUT", 4: "POST", 5: "DELETE", 6: "PATCH"}


def _pb_http(v, out, depth=0):
    for fn, _, x in _pb_fields(v):
        if fn in _PB_HTTP_VERBS:
            out.append((_PB_HTTP_VERBS[fn], _pb_str(x)))
        elif fn == 8:  # custom
            kind = path = ""
            for f2, _, y in _pb_fields(x):
                if f2 == 1:
                    kind = _pb_str(y)
                elif f2 == 2:
                    path = _pb_str(y)
            out.append((kind, path))
        elif fn == 11 and depth == 0:  # additional_bindings
            _pb_http(x, out, depth + 1)


def _pb_field(v):
    name, ftype, type_name, behaviors = "", 0, "", []
    for fn, _, x in _pb_fields(v):
        if fn == 1:
            name = _pb_str(x)
        elif fn == 5:
            ftype = x
        elif fn == 6:
            type_name = _pb_str(x)
        elif fn == 8:
            for f2, w2, y in _pb_fields(x):
                if f2 == 1052:  # google.api.field_behavior
                    behaviors.extend(_pb_packed(y) if w2 == 2 else [y])
    return (name, ftype, type_name, tuple(behaviors))


def _pb_message(v, prefix, fname, out):
    m = {"name": "", "fields": [], "deprecated": False, "resource": None, "file": fname}
    nested = []
    for fn, _, x in _pb_fields(v):
        if fn == 1:
            m["name"] = _pb_str(x)
        elif fn == 2:
            m["fields"].append(_pb_field(x))
        elif fn == 3:
            nested.append(x)
        elif fn == 7:
            for f2, _, y in _pb_fields(x):
                if f2 == 3:
                    m["deprecated"] = bool(y)
                elif f2 == 1053:  # google.api.resource
                    m["resource"] = _pb_resource(y)
    m["full"] = prefix + "." + m["name"]
    out.append(m)
    for x in nested:
        _pb_message(x, m["full"], fname, out)


def _pb_service(v, pkg, fname):
    svc = {"name": "", "methods": [], "deprecated": False, "default_host": "", "file": fname}
    raw = []
    for fn, _, x in _pb_fields(v):
        if fn == 1:
            svc["name"] = _pb_str(x)
        elif fn == 2:
            raw.append(x)
        elif fn == 3:
            for f2, _, y in _pb_fields(x):
                if f2 == 33:
                    svc["deprecated"] = bool(y)
                elif f2 == 1049:  # google.api.default_host
                    svc["default_host"] = _pb_str(y)
    for x in raw:
        mt = {"name": "", "input": "", "output": "", "deprecated": False, "http": []}
        for fn, _, y in _pb_fields(x):
            if fn == 1:
                mt["name"] = _pb_str(y)
            elif fn == 2:
                mt["input"] = _pb_str(y)
            elif fn == 3:
                mt["output"] = _pb_str(y)
            elif fn == 4:
                for f2, _, z in _pb_fields(y):
                    if f2 == 33:
                        mt["deprecated"] = bool(z)
                    elif f2 == 72295728:  # google.api.http
                        _pb_http(z, mt["http"])
        svc["methods"].append(mt)
    svc["full"] = "." + (pkg + "." if pkg else "") + svc["name"]
    return svc


def decode_descriptor_set(pb_path):
    """Returns [{name, package, messages, services}] for every file in a FileDescriptorSet."""
    with open(pb_path, "rb") as f:
        data = memoryview(f.read())
    files = []
    for fn, _, fv in _pb_fields(data):
        if fn != 1:
            continue
        fd = {"name": "", "package": "", "messages": [], "services": []}
        raw_msgs, raw_svcs = [], []
        for f2, _, x in _pb_fields(fv):
            if f2 == 1:
                fd["name"] = _pb_str(x)
            elif f2 == 2:
                fd["package"] = _pb_str(x)
            elif f2 == 4:
                raw_msgs.append(x)
            elif f2 == 6:
                raw_svcs.append(x)
        prefix = "." + fd["package"] if fd["package"] else ""
        for x in raw_msgs:
            _pb_message(x, prefix, fd["name"], fd["messages"])
        for x in raw_svcs:
            fd["services"].append(_pb_service(x, fd["package"], fd["name"]))
        files.append(fd)
    return files


# --- googleapis checkout / compilation helpers --------------------------------
def git_output(repo, *args):
    r = subprocess.run(["git", "-C", repo] + list(args), capture_output=True, text=True)
    return r.stdout.strip() if r.returncode == 0 else ""


def resolve_commit(repo, sha):
    return git_output(repo, "rev-parse", "--verify", "--quiet", sha + "^{commit}")


def export_googleapis_protos(repo, sha, temp_dir):
    """Exports the .proto files of googleapis@sha with `git archive` (the clone's checkout is not touched)."""
    full = resolve_commit(repo, sha)
    if not full:
        return None, None
    dest = os.path.join(temp_dir, "googleapis-" + full)
    if os.path.isdir(dest):
        return dest, full
    tops = [t for t in git_output(repo, "ls-tree", "--name-only", full).split() if t in ("google", "grafeas", "backstory")]
    tmp = dest + ".partial"
    shutil.rmtree(tmp, ignore_errors=True)
    os.makedirs(tmp)
    proc = subprocess.Popen(["git", "-C", repo, "archive", full] + tops, stdout=subprocess.PIPE)
    with tarfile.open(fileobj=proc.stdout, mode="r|") as tf:
        for member in tf:
            if member.isfile() and member.name.endswith(".proto"):
                try:
                    tf.extract(member, tmp, filter="data")
                except TypeError:  # Python < 3.12
                    tf.extract(member, tmp)
    if proc.wait() != 0:
        shutil.rmtree(tmp, ignore_errors=True)
        return None, None
    os.rename(tmp, dest)
    return dest, full


def _protobuf_include_dir():
    candidates = []
    protoc = shutil.which("protoc")
    if protoc:
        candidates.append(os.path.join(os.path.dirname(os.path.realpath(protoc)), "..", "include"))
    candidates += ["/usr/include", "/usr/local/include", "/opt/homebrew/include"]
    for c in candidates:
        if os.path.exists(os.path.join(c, "google", "protobuf", "descriptor.proto")):
            return os.path.abspath(c)
    return None


def compile_descriptor_set(proto_root, out_pb):
    """Compiles every google/ and grafeas/ proto into one descriptor set (cached at out_pb)."""
    if os.path.exists(out_pb):
        return out_pb
    if not shutil.which("protoc"):
        sys.exit("protoc is required for --manageability (apt install protobuf-compiler / brew install protobuf)")
    files = []
    for top in ("google", "grafeas"):
        for root, _, names in os.walk(os.path.join(proto_root, top)):
            files += [os.path.relpath(os.path.join(root, n), proto_root) for n in names if n.endswith(".proto")]
    files.sort()
    include = _protobuf_include_dir()
    excluded = []
    for _ in range(5):  # drop protos with unresolvable imports and retry
        argfile = out_pb + ".args"
        with open(argfile, "w") as f:
            f.write("\n".join(files))
        cmd = ["protoc", "-I", proto_root] + (["-I", include] if include else [])
        cmd += ["--include_imports", "--descriptor_set_out=" + out_pb + ".partial", "@" + argfile]
        r = subprocess.run(cmd, cwd=proto_root, capture_output=True, text=True)
        if r.returncode == 0:
            os.rename(out_pb + ".partial", out_pb)
            os.remove(argfile)
            if excluded:
                print("[WARN] excluded protos that failed to compile: %s" % ", ".join(sorted(excluded)))
            return out_pb
        errors = [l for l in r.stderr.splitlines() if "warning:" not in l]
        bad = {m.group(1) for m in (re.match(r"^([\w./-]+\.proto):\d+:\d+:", l) for l in errors) if m}
        bad &= set(files)
        if not bad:
            sys.exit("protoc failed:\n" + "\n".join(errors[:40]))
        excluded += sorted(bad)
        files = [f for f in files if f not in bad]
    sys.exit("protoc kept failing; excluded so far: %s" % ", ".join(excluded))


def collect_googleapis_pins(kcc_dir):
    """Returns [(sha, users)]: the global pin (apis/git.versions) first, then per-service pins in apis/*/generate.sh."""
    pins = []
    versions = os.path.join(kcc_dir, "apis", "git.versions")
    if os.path.exists(versions):
        with open(versions) as f:
            for line in f:
                m = re.match(r"^https://github.com/googleapis/googleapis\s+([0-9a-f]{40})", line.strip())
                if m:
                    pins.append((m.group(1), ["global (apis/git.versions)"]))
    per_service = defaultdict(set)
    apis_dir = os.path.join(kcc_dir, "apis")
    if os.path.isdir(apis_dir):
        for svc in sorted(os.listdir(apis_dir)):
            path = os.path.join(apis_dir, svc, "generate.sh")
            if os.path.exists(path):
                with open(path, errors="replace") as f:
                    for sha in set(re.findall(r"\b([0-9a-f]{40})\b", f.read())):
                        per_service[sha].add(svc)
    global_shas = {p[0] for p in pins}
    for sha in sorted(per_service):
        if sha not in global_shas:
            pins.append((sha, sorted(per_service[sha])))
    return pins


class GitBlobReader:
    """Reads files at arbitrary commits through one `git cat-file --batch` process."""

    def __init__(self, repo):
        self.proc = subprocess.Popen(["git", "-C", repo, "cat-file", "--batch"], stdin=subprocess.PIPE,
                                     stdout=subprocess.PIPE)
        self.cache = {}

    def get(self, sha, path):
        key = (sha, path)
        if key not in self.cache:
            self.proc.stdin.write(("%s:%s\n" % (sha, path)).encode())
            self.proc.stdin.flush()
            header = self.proc.stdout.readline().decode().split()
            if len(header) != 3:
                self.cache[key] = None
            else:
                data = self.proc.stdout.read(int(header[2]))
                self.proc.stdout.read(1)
                self.cache[key] = data.decode("utf-8", "replace")
        return self.cache[key]


def _norm(s):
    return re.sub(r"[^a-z0-9]", "", s.lower())


def _family(pkg):
    return ".".join(seg for seg in pkg.split(".") if not VERSION_SEGMENT_RE.match(seg))


def _canonical_for_matching(name):
    for p in ("Global", "Regional", "Region", "Zonal"):
        if name.startswith(p) and len(name) > len(p) and name[len(p)].isupper():
            name = name[len(p):]
            break
    return singularize_name(name)


def _is_non_ga(pkg):
    return any(NONGA_SEGMENT_RE.match(seg) for seg in pkg.split("."))


def _major_version(pkg):
    for seg in pkg.split("."):
        m = re.match(r"^v(\d+)", seg)
        if m:
            return int(m.group(1))
    return 0


def _pattern_literals(p):
    """Port of patternLiterals in nondeclarative.go."""
    out = []
    for s in p.strip("/").split("/"):
        if s == "" or s == "*" or s.startswith("{"):
            continue
        if not out and len(s) >= 2 and s[0] == "v" and s[1].isdigit():
            continue
        out.append(s)
    return out


def _matches_resource_pattern(http_path, resource_pattern):
    """Port of matchesResourcePattern in nondeclarative.go."""
    if not http_path or not resource_pattern:
        return False
    base = http_path
    i = base.find(":")
    if i >= 0:
        base = base[:i]
    base = base.replace("{", "").replace("}", "")
    i = base.rfind("=")
    if i >= 0:
        base = base[i + 1:]
    h = _pattern_literals(base)
    return len(h) > 0 and h == _pattern_literals(resource_pattern)


def _http_to_pattern(path):
    """'/compute/v1/projects/{project}/regions/{region}/addresses/{address}' -> resource-style pattern."""
    segs = path.lstrip("/").split("/")
    for i, s in enumerate(segs):
        if VERSION_SEGMENT_RE.match(s):
            segs = segs[i + 1:]
            break
    p = re.sub(r"\{([a-zA-Z0-9_.]+)=[^}]*\}", r"{\1}", "/".join(segs))
    return p if p.startswith(("projects", "organizations", "folders", "billingAccounts")) else None


def _collection_of(pattern):
    lits = [s for s in pattern.split("/") if s and not s.startswith("{")]
    return lits[-1] if lits else ""


def _read_text(path):
    with open(path, encoding="utf-8", errors="replace") as f:
        return f.read()


class ManageabilityClassifier:
    def __init__(self, files, googleapis_sha, proto_root, kcc_dir, kcc_ref, script_dir, pins, blob_reader):
        self.googleapis_sha = googleapis_sha
        self.proto_root = proto_root
        self.kcc_dir = kcc_dir
        self.kcc_ref = kcc_ref
        self.script_dir = script_dir
        self.pins = pins  # [(sha, users, {message full name -> proto file})], global pin first
        self.blobs = blob_reader
        self._line_cache = {}
        self._index(files)
        self._load_rules()
        self._load_kcc_inventory()

    # --- setup ---------------------------------------------------------------
    def _index(self, files):
        self.msgs, self.svcs_by_pkg, self.svcs_by_file = {}, defaultdict(list), defaultdict(list)
        for fd in files:
            for m in fd["messages"]:
                m["pkg"] = fd["package"]
                self.msgs[m["full"]] = m
            for s in fd["services"]:
                self.svcs_by_pkg[fd["package"]].append(s)
                self.svcs_by_file[fd["name"]].append(s)

    def _kcc_text(self, rel):
        path = os.path.join(self.kcc_dir, rel)
        return _read_text(path) if os.path.exists(path) else ""

    def _load_rules(self):
        # nondeclarative.go: collection maps, bulk prefixes and rule line numbers.
        self.ephemeral, self.leaf, self.bulk_prefixes = {}, {}, []
        self.nd_rule_lines = {}
        block = None
        nd = self._kcc_text(ND_REL)
        if not nd:
            print("[WARN] %s not found; non-declarative drops are disabled" % ND_REL)
        for i, ln in enumerate(nd.splitlines(), 1):
            if ln.startswith("var ephemeralCollectionSegments"):
                block = self.ephemeral
            elif ln.startswith("var nonDeclarativeLeafCollections"):
                block = self.leaf
            elif ln.startswith("var bulkDataPlaneRPCPrefixes"):
                block = self.bulk_prefixes
            elif ln.startswith("}"):
                block = None
            elif isinstance(block, dict):
                m = re.match(r'^\s+"(\w+)": "([^"]+)",', ln)
                if m:
                    block[m.group(1)] = (m.group(2), i)
            elif isinstance(block, list):
                m = re.match(r'^\s+"(\w+)",', ln)
                if m:
                    block.append(m.group(1))
            for key, needle in (("1.1", "if !hasCreate && !hasUpdate && (hasGet || hasDelete)"),
                                ("1.2", "if hasCreate && !hasUpdate && (!hasDelete || hasCancel)"),
                                ("3.1", "if len(bulkRPCs) > 0")):
                if needle in ln:
                    self.nd_rule_lines[key] = i
        # deprecated.go: retired GA packages without `option deprecated`.
        self.known_deprecated = {}
        for i, ln in enumerate(self._kcc_text(DEP_REL).splitlines(), 1):
            m = re.match(r'^\s+"(google\.[\w.]+)": "([^"]+)",', ln)
            if m:
                self.known_deprecated[m.group(1)] = (m.group(2), i)
        # JSON configs from the analyzed KCC tree (so output depends only on the SHA inputs), with line numbers
        # for permalinks. Falls back to the copies next to this script for older KCC trees that lack them.
        self.config_fallbacks = []
        self.skips = self._load_patterns("coverage_skip.json", "skips")
        self.policy = self._load_patterns("policy_exclusions.json", "exclusions")
        self.deprecated_kinds = {}
        path = self._config_path("deprecated_services.json")
        if path:
            text = _read_text(path)
            lines = text.splitlines()
            for item in json.loads(text):
                base = item["service"].replace(".googleapis.com", "").lower()
                ln = next((i for i, l in enumerate(lines, 1) if '"service": "%s"' % item["service"] in l), None)
                for kind in item["kinds"]:
                    self.deprecated_kinds[(base, kind.lower())] = (item["name"], item["evidence"], ln)
        self.cp_keys = set()
        path = self._config_path("cp_resources_list.json")
        if path:
            for item in json.loads(_read_text(path)):
                self.cp_keys.add(("res", item.get("resource", "").lower()))
                self.cp_keys.add(("kind", _norm(item.get("kind", ""))))
        path = self._config_path("services_acronyms.json")
        acronyms = json.loads(_read_text(path)) if path else {}
        self.kind_prefixes = acronyms.get("prefixes", {})
        self.kind_acronyms = acronyms.get("acronyms", {})

    def _config_path(self, fname):
        path = os.path.join(self.kcc_dir, GREENFIELD_REL, fname)
        if os.path.exists(path):
            return path
        path = os.path.join(self.script_dir, fname)
        if os.path.exists(path):
            self.config_fallbacks.append(fname)
            return path
        return None

    def _predicted_kind(self, service, name):
        """Same as get_predicted_kcc_kind, but with services_acronyms.json from the analyzed KCC tree."""
        prefix = self.kind_prefixes.get(service) or "".join(p.capitalize() for p in re.split(r"[-_]", service))
        kind = prefix + name[len(prefix):] if name.lower().startswith(prefix.lower()) else prefix + name
        for old, new in self.kind_acronyms.items():
            kind = re.sub(r"%s([A-Z]|$)" % old, r"%s\1" % new, kind)
        return kind

    def _load_patterns(self, fname, key):
        path = self._config_path(fname)
        if not path:
            return []
        text = _read_text(path)
        lines = text.splitlines()
        out = []
        for s in json.loads(text).get(key, []):
            needle = json.dumps(s["pattern"])
            ln = next((i for i, l in enumerate(lines, 1) if needle in l), None)
            out.append((re.compile(s["pattern"]), s["reason"], "%s/%s" % (GREENFIELD_REL, fname), ln))
        return out

    def _load_kcc_inventory(self):
        self.crd = {}
        crd_dir = os.path.join(self.kcc_dir, "config", "crds", "resources")
        for fname in sorted(os.listdir(crd_dir)) if os.path.isdir(crd_dir) else []:
            txt = _read_text(os.path.join(crd_dir, fname))
            g = re.search(r"^  group: (\S+)", txt, re.M)
            k = re.search(r"^    kind: (\S+)", txt, re.M)
            if g and k:
                self.crd[k.group(1)] = {
                    "group": g.group(1), "file": "config/crds/resources/" + fname,
                    "versions": set(re.findall(r"^    name: (v\d+(?:alpha|beta)?\d*)\s*$", txt, re.M)),
                    "tf": 'cnrm.cloud.google.com/tf2crd: "true"' in txt,
                    "dcl": 'cnrm.cloud.google.com/dcl2crd: "true"' in txt}
        self.static = {}
        for m in re.finditer(r'\{Group: "([^"]+)", Kind: "([^"]+)"\}:\s*\{DefaultController: k8s\.ReconcilerType(\w+), '
                             r'SupportedControllers: \[\]k8s\.ReconcilerType\{([^}]*)\}',
                             self._kcc_text("pkg/controller/resourceconfig/static_config.go")):
            self.static[m.group(2)] = (m.group(3), set(re.findall(r"ReconcilerType(\w+)", m.group(4))))
        self.registered = set()
        for root, _, names in os.walk(os.path.join(self.kcc_dir, "pkg", "controller", "direct")):
            for n in names:
                if n.endswith(".go"):
                    for m in re.finditer(r"RegisterModel\((?:[a-z0-9]+\.)?(\w+?)GVK", _read_text(os.path.join(root, n))):
                        self.registered.add(m.group(1))
        # +kcc:proto annotations on <Kind>Spec / <Kind>ObservedState give exact proto -> kind matches.
        self.proto_to_kinds, self.annotated_kinds = defaultdict(set), set()
        for root, _, names in os.walk(os.path.join(self.kcc_dir, "apis")):
            for n in names:
                if not n.endswith("_types.go"):
                    continue
                pending = None
                for ln in _read_text(os.path.join(root, n)).splitlines():
                    s = ln.strip()
                    m = re.match(r"//\s*\+kcc:(?:spec:|observedstate:)?proto=([\w.]+)", s)
                    if m:
                        pending = m.group(1)
                        continue
                    if s.startswith("//"):
                        continue
                    t = re.match(r"type (\w+) struct", s)
                    if t and pending:
                        tn = t.group(1)
                        kind = tn[:-4] if tn.endswith("Spec") else (tn[:-13] if tn.endswith("ObservedState") else None)
                        if kind and kind in self.crd:
                            self.proto_to_kinds[self._proto_key(pending)].add(kind)
                            self.annotated_kinds.add(kind)
                    pending = None
        self.kcc_groups = defaultdict(list)
        for kind, c in self.crd.items():
            self.kcc_groups[c["group"].replace(".cnrm.cloud.google.com", "")].append(kind)

    @staticmethod
    def _proto_key(full_name):
        return ".".join(seg for seg in full_name.split(".") if not VERSION_SEGMENT_RE.match(seg))

    # --- links -----------------------------------------------------------------
    def kcc_link(self, rel, line=None):
        return KCC_GITHUB % (self.kcc_ref, rel) + ("#L%d" % line if line else "")

    def proto_line(self, fname, kind, name):
        key = (fname, kind, name)
        if key not in self._line_cache:
            path = os.path.join(self.proto_root, fname)
            ln = None
            if os.path.exists(path):
                rx = re.compile({"message": r"^\s*message\s+%s\s*\{", "service": r"^\s*service\s+%s\s*\{",
                                 "rpc": r"^\s*rpc\s+%s\s*\("}[kind] % re.escape(name))
                ln = next((i for i, l in enumerate(_read_text(path).splitlines(), 1) if rx.search(l)), None)
            self._line_cache[key] = ln
        return self._line_cache[key]

    def gapi_link(self, fname, kind=None, name=None):
        ln = self.proto_line(fname, kind, name) if kind else None
        return GOOGLEAPIS_GITHUB % (self.googleapis_sha, fname) + ("#L%d" % ln if ln else "")

    def first_pin(self, msg):
        """Earliest KCC googleapis pin (global first) that has msg -> (sha, users, link)."""
        for sha, users, names in self.pins:
            if msg["full"] in names:
                fname = names[msg["full"]]
                text = self.blobs.get(sha, fname) if self.blobs else None
                ln = None
                if text:
                    rx = re.compile(r"^\s*message\s+%s\s*\{" % re.escape(msg["name"]))
                    ln = next((i for i, l in enumerate(text.splitlines(), 1) if rx.search(l)), None)
                return sha, users, GOOGLEAPIS_GITHUB % (sha, fname) + ("#L%d" % ln if ln else "")
        return None, None, None

    # --- controllerbuilder rule ports ----------------------------------------
    def pkg_methods(self, pkg):
        for s in self.svcs_by_pkg.get(pkg, []):
            for m in s["methods"]:
                yield s, m

    def nd_rpc_flags(self, pkg, name, patterns, collection):
        """Port of the flag computation in nonDeclarativeRPCReason."""
        cap = collection[:1].upper() + collection[1:] if collection else ""
        f = {"C": False, "U": False, "D": False, "G": False, "Cancel": False}
        bulk = []
        for _, m in self.pkg_methods(pkg):
            mn = m["name"]
            any_match = False
            for verb, path in m["http"]:
                if any(_matches_resource_pattern(path, p) for p in patterns):
                    any_match = True
                    if verb == "POST" and ":" not in path and mn.startswith("Create"):
                        f["C"] = True
                    elif verb in ("PATCH", "PUT") and mn.startswith("Update"):
                        f["U"] = True
                    elif verb == "DELETE" and mn.startswith("Delete"):
                        f["D"] = True
                    elif verb == "GET" and mn.startswith("Get"):
                        f["G"] = True
            if mn == "Create" + name:
                f["C"] = True
            elif mn == "Update" + name:
                f["U"] = True
            elif mn == "Delete" + name:
                f["D"] = True
            elif mn == "Get" + name:
                f["G"] = True
            elif mn == "Cancel" + name:
                f["Cancel"] = True
            for pre in self.bulk_prefixes:
                if mn == pre + name or (cap and mn == pre + cap) or (any_match and mn.startswith(pre)):
                    bulk.append(mn)
        return f, sorted(set(bulk))

    def nd_pattern_reason(self, patterns):
        """Port of nonDeclarativePatternReason: (segment, category, line, pattern) or None."""
        for p in patterns:
            lits = _pattern_literals(p)
            if not lits:
                continue
            for seg in lits:
                if seg in self.ephemeral:
                    return seg, self.ephemeral[seg][0], self.ephemeral[seg][1], p
            if lits[-1] in self.leaf:
                return lits[-1], self.leaf[lits[-1]][0], self.leaf[lits[-1]][1], p
        return None

    def method_manages(self, m, msg):
        """Port of methodManagesResource in deprecated.go."""
        name = msg["name"]
        if m["name"] in ("Create" + name, "Get" + name, "Update" + name, "Delete" + name):
            return True
        if msg["full"] in (m["input"], m["output"]):
            return True
        inp = self.msgs.get(m["input"])
        return bool(inp) and any(f[1] == PB_TYPE_MESSAGE and f[2] == msg["full"] for f in inp["fields"])

    @staticmethod
    def svc_deprecated(s):
        return s["deprecated"] or (len(s["methods"]) > 0 and all(m["deprecated"] for m in s["methods"]))

    def deprecated_reason(self, msg):
        """Port of deprecatedResourceReason -> (text, link) or None."""
        pkg = msg["pkg"]
        if pkg in self.known_deprecated:
            txt, ln = self.known_deprecated[pkg]
            return txt, self.kcc_link(DEP_REL, ln)
        if msg["deprecated"]:
            return ("message %s is marked deprecated in proto" % msg["full"][1:],
                    self.gapi_link(msg["file"], "message", msg["name"]))
        services = [s for s in self.svcs_by_pkg.get(pkg, []) if any(self.method_manages(m, msg) for m in s["methods"])]
        if not services:
            services = list(self.svcs_by_file.get(msg["file"], []))
        if services and all(self.svc_deprecated(s) for s in services):
            s = services[0]
            return "service %s is marked deprecated in proto" % s["full"][1:], self.gapi_link(s["file"], "service", s["name"])
        matched = [(s, m) for s in self.svcs_by_pkg.get(pkg, []) for m in s["methods"] if self.method_manages(m, msg)]
        if matched and all(s["deprecated"] or m["deprecated"] for s, m in matched):
            s, m = matched[0]
            return ("all %d RPCs managing %s in %s are marked deprecated in proto" % (len(matched), msg["name"], pkg),
                    self.gapi_link(s["file"], "rpc", m["name"]))
        return None

    def server_generated(self, pkg, names):
        """Port of isServerGeneratedCreateInput on Create<R>/Insert<R>; None when there is no such RPC."""
        found = None
        for _, m in self.pkg_methods(pkg):
            if m["name"] in names and m["input"] in self.msgs:
                fields = self.msgs[m["input"]]["fields"]
                res = any(f[0] == "parent" for f in fields) and not any(
                    f[0] != "request_id" and f[0].endswith("_id") and f[1] == PB_TYPE_STRING for f in fields)
                found = res if found is None else (found or res)
        return found

    def ops_for(self, pkg, name, patterns, collection, plural, synth_services):
        """Rubric op detection over non-deprecated RPCs in the resource's own package."""
        cap = collection[:1].upper() + collection[1:] if collection else ""
        plurals = {cap, name + "s", name + "es"}
        if plural:
            plurals.add(plural[:1].upper() + plural[1:])
        if name.endswith("y"):
            plurals.add(name[:-1] + "ies")
        plurals.discard("")
        ops = {k: [] for k in ("C", "Calias", "R", "Ralias", "U", "Ualias", "D", "Dtag", "Dalias", "L")}
        deprecated = []
        for s, m in self.pkg_methods(pkg):
            mn = m["name"]
            hit = []
            if synth_services is not None and s["name"] in synth_services:
                hit += {"Insert": ["C"], "Create": ["C"], "Get": ["R"], "Patch": ["U"], "Update": ["U"],
                        "Delete": ["D"], "List": ["L"]}.get(mn, [])
            for verb, path in m["http"]:
                if patterns and any(_matches_resource_pattern(path, rp) for rp in patterns):
                    if verb == "POST" and ":" not in path and mn.startswith("Create"):
                        hit.append("C")
                    elif verb in ("PATCH", "PUT") and mn.startswith("Update"):
                        hit.append("U")
                    elif verb == "DELETE" and mn.startswith("Delete"):
                        hit.append("D")
                    elif verb == "GET" and mn.startswith("Get"):
                        hit.append("R")
                    elif verb == "GET" and mn.startswith("List"):
                        hit.append("L")
            if mn in ("Create" + name, "Insert" + name):
                hit.append("C")
            if mn in ("BatchCreate" + name, "BatchCreate" + cap, "Upload" + name, "Upsert" + name, "Set" + name):
                hit.append("Calias")
            if mn == "Get" + name:
                hit.append("R")
            if mn in ("BatchGet" + cap, "BatchGet" + name, "Describe" + name, "Read" + name):
                hit.append("Ralias")
            if mn in ("Update" + name, "Patch" + name):
                hit.append("U")
            if mn in ("BatchUpdate" + cap, "BatchUpdate" + name, "Upsert" + name, "Set" + name):
                hit.append("Ualias")
            if mn == "Delete" + name:
                hit.append("D")
            if mn in ("Destroy" + name, "Disable" + name, "Deactivate" + name, "Remove" + name):
                hit.append("Dtag")
            if mn in ("Cancel" + name, "Abort" + name, "Finish" + name, "Terminate" + name):
                hit.append("Dalias")
            if mn in {"List" + p for p in plurals}:
                hit.append("L")
            if not hit:
                continue
            if s["deprecated"] or m["deprecated"]:
                deprecated.append(mn)
                continue
            for h in set(hit):
                ops[h].append(mn)
        return {k: sorted(set(v)) for k, v in ops.items()}, sorted(set(deprecated))

    @staticmethod
    def layer(patterns):
        best = None
        for p in patterns:
            lits = [s for s in p.split("/") if s and not s.startswith("{")]
            depth = len([s for s in lits[:-1] if s not in LAYER_SKIP_PARENTS])
            best = depth if best is None else min(best, depth)
        if best is None:
            return ""
        return ["Easy (Leaf)", "Next Layer (1 Parent)", "Next Next Layer (2 Parents)"][best] if best < 3 else "Deeply Nested"

    def kind_controller(self, kind):
        st = self.static.get(kind)
        if kind in self.registered or (st is not None and "Direct" in st[1]):
            return "direct"
        if st is not None:
            return st[0].lower()  # terraform / dcl
        c = self.crd.get(kind, {})
        return "tf" if c.get("tf") else ("dcl" if c.get("dcl") else "none")

    def heuristic_kinds(self, service, name, pkg):
        variants = {service, service + "s"}
        if _family(pkg).startswith("google.cloud.dialogflow.cx"):
            variants = {"dialogflowcx"}
        if service.endswith("s"):
            variants.add(service[:-1])
        for canon, aliases in MATCH_SERVICE_ALIASES.items():
            if service in [canon] + aliases:
                variants.update([canon] + aliases)
        names = {_norm(name), _norm(_canonical_for_matching(name))}
        prefixes = [_norm(v) for v in variants] + ["gcp", "google", "cloud", "bigquery", "api", "vertexai"]
        predicted = _norm(self._predicted_kind(service, _canonical_for_matching(name)))
        out = set()
        for g in (v for v in variants if v in self.kcc_groups):
            for kind in self.kcc_groups[g]:
                kn = _norm(kind)
                if kn in names or kn == predicted or any(kn.startswith(p) and kn[len(p):] in names for p in prefixes):
                    out.add(kind)
        return out

    # --- universe ----------------------------------------------------------------
    def universe(self):
        by_type = defaultdict(list)  # (type, package family): Dialogflow CX and ES share type strings
        for m in self.msgs.values():
            r = m["resource"]
            if r and r["patterns"] and "/" in r["type"]:
                by_type[(r["type"], _family(m["pkg"]))].append(m)
        families = defaultdict(set)
        for (t, fam) in by_type:
            families[t].add(fam)
        resources = []
        for (rtype, fam), msgs in sorted(by_type.items()):
            ga = sorted((m for m in msgs if not _is_non_ga(m["pkg"])), key=lambda m: (-_major_version(m["pkg"]), m["pkg"]))
            msg = ga[0] if ga else sorted(msgs, key=lambda m: (-_major_version(m["pkg"]), m["pkg"]))[0]
            resources.append({"type": rtype, "id": rtype if len(families[rtype]) == 1 else "%s [%s]" % (rtype, fam),
                              "msg": msg, "is_ga": bool(ga), "patterns": list(dict.fromkeys(msg["resource"]["patterns"])),
                              "plural": msg["resource"]["plural"], "synthesized": None,
                              "other_pkgs": sorted({m["pkg"] for m in msgs if m is not msg})})
        # GA APIs without google.api.resource annotations (Compute, SQL Admin, BigQuery v2, ...).
        annotated = {m["full"] for ms in by_type.values() for m in ms}
        synth = {}
        for pkg, svcs in sorted(self.svcs_by_pkg.items()):
            if _is_non_ga(pkg) or pkg.startswith(SYNTH_SKIP_PACKAGE_PREFIXES):
                continue
            for s in svcs:  # bare-verb services: service Addresses { rpc Get / Insert / ... }
                names = {m["name"]: m for m in s["methods"]}
                if "Get" in names and ("Insert" in names or "Create" in names):
                    out = names["Get"]["output"]
                    if out in annotated or out not in self.msgs:
                        continue
                    e = synth.setdefault(out, {"services": set(), "patterns": [], "host": s["default_host"]})
                    e["services"].add(s["name"])
                    for _, path in names["Get"]["http"]:
                        p = _http_to_pattern(path)
                        if p and p not in e["patterns"]:
                            e["patterns"].append(p)
            all_names = {x["name"] for _, x in self.pkg_methods(pkg)}
            for s in svcs:  # <Verb><Msg>-style: GetDataset + InsertDataset
                for m in s["methods"]:
                    if m["name"].startswith("Get") and m["output"] not in annotated and m["output"] in self.msgs:
                        mm = self.msgs[m["output"]]
                        if mm["pkg"] != pkg or m["name"] != "Get" + mm["name"]:
                            continue
                        if ("Create" + mm["name"]) in all_names or ("Insert" + mm["name"]) in all_names:
                            e = synth.setdefault(mm["full"], {"services": set(), "patterns": [], "host": s["default_host"]})
                            for _, path in m["http"]:
                                p = _http_to_pattern(path)
                                if p and p not in e["patterns"]:
                                    e["patterns"].append(p)
        for full, e in sorted(synth.items()):
            msg = self.msgs[full]
            host = e["host"] or (msg["pkg"].split(".")[2] + ".googleapis.com" if msg["pkg"].count(".") >= 2 else "unknown")
            rtype = "%s/%s" % (host, msg["name"])
            if rtype in families:
                continue
            resources.append({"type": rtype, "id": rtype, "msg": msg, "is_ga": True, "patterns": e["patterns"],
                              "plural": "", "synthesized": e["services"] or set(), "other_pkgs": []})
        return resources

    # --- classification -----------------------------------------------------------
    def classify(self):
        records = []
        global_pin = self.pins[0][0] if self.pins else None
        for R in self.universe():
            records.append(self._classify_one(R, global_pin))
        matched = {k for r in records for k in re.findall(r"(\w+)\(", r["kcc_kinds"])}
        for kind, c in sorted(self.crd.items()):
            if kind in matched or c["versions"] != {"v1alpha1"} or self.kind_controller(kind) not in ("terraform", "dcl", "tf"):
                continue
            rec = {col: "" for col in MANAGEABILITY_COLUMNS}
            rec.update(id="(no googleapis resource) " + kind, service=c["group"].split(".")[0], predicted_kind=kind,
                       kind=kind, kcc_kinds="%s(v1alpha1,%s)" % (kind, self.kind_controller(kind)),
                       kcc_status=STATUS_BROWNFIELD, bucket=BUCKET_BROWNFIELD, kcc_link=self.kcc_link(c["file"]),
                       api_group=API_GROUP_CLOUD, ga=None, flags="brownfield; not-in-googleapis")
            records.append(rec)
        return records

    def _classify_one(self, R, global_pin):
        msg, pkg, name = R["msg"], R["msg"]["pkg"], R["msg"]["name"]
        service = R["type"].split("/")[0].split(".")[0]
        patterns = R["patterns"]
        collection = _collection_of(patterns[0]) if patterns else ""
        if not collection and R["plural"]:
            collection = R["plural"]
        predicted = self._predicted_kind(service, name)
        api_group, api_family = api_group_of(pkg)
        rec = {"id": R["id"], "resource_type": R["type"], "service": service, "api_group": api_group,
               "api_family": api_family, "api_area": ".".join(pkg.split(".")[:2]),
               "proto_message": msg["full"][1:], "proto_link": self.gapi_link(msg["file"], "message", name),
               "ga": R["is_ga"], "unannotated": "yes" if R["synthesized"] is not None else "",
               "other_packages": ";".join(R["other_pkgs"]), "patterns": " | ".join(patterns[:3]),
               "layer": self.layer(patterns), "predicted_kind": predicted}
        drops, flags = [], []

        # KCC kinds: exact +kcc:proto matches, plus heuristics for TF/DCL kinds (never annotated).
        kinds = set(self.proto_to_kinds.get(self._proto_key(msg["full"][1:]), set()))
        match = "proto" if kinds else ""
        extra = {k for k in self.heuristic_kinds(service, name, pkg) if k not in self.annotated_kinds}
        extra |= {k for k in MANUAL_KIND_ALIASES.get(R["type"], []) if k in self.crd}
        if not kinds:
            extra |= self.heuristic_kinds(service, name, pkg)
        if extra - kinds:
            match = (match + "+heuristic") if match else "heuristic"
            kinds |= extra
        if R["type"] in MANUAL_KIND_OVERRIDES:
            kinds = {k for k in MANUAL_KIND_OVERRIDES[R["type"]] if k in self.crd}
            match = "manual"
        rec["kcc_kinds"] = ";".join(sorted("%s(%s,%s)" % (k, "/".join(sorted(self.crd[k]["versions"])),
                                                          self.kind_controller(k)) for k in kinds))
        rec["kcc_match"] = match
        v1beta1 = sorted(k for k in kinds if "v1beta1" in self.crd[k]["versions"])
        if not kinds:
            rec["kcc_status"], rec["kind"], rec["kcc_link"] = STATUS_NET_NEW, predicted, ""
        else:
            controllers = {self.kind_controller(k) for k in kinds}
            rec["kcc_status"] = (STATUS_V1BETA1 if v1beta1 else STATUS_DIRECT if "direct" in controllers else
                                 STATUS_BROWNFIELD if controllers & {"terraform", "dcl", "tf"} else STATUS_TYPES_ONLY)
            rec["kind"] = sorted(kinds)[0]
            rec["kcc_link"] = self.kcc_link(self.crd[(v1beta1 or sorted(kinds))[0]]["file"])
        rec["in_cp_list"] = "yes" if (("res", ("%s/%s" % (service, singularize_name(name))).lower()) in self.cp_keys
                                      or ("kind", _norm(rec["kind"])) in self.cp_keys) else ""

        # D1 skip list / example APIs
        for rx, why, rel, ln in self.skips:
            if rx.match(R["type"]):
                drops.append(("skip-list", why, self.kcc_link(rel, ln)))
                break
        if pkg.startswith("google.example."):
            drops.append(("example-api", "googleapis sample API (%s), not a real service" % pkg, rec["proto_link"]))
        # D2 deprecated
        dep = self.deprecated_reason(msg)
        if dep:
            drops.append(("deprecated", dep[0], dep[1]))
        for k in (rec["kind"], predicted):
            hit = self.deprecated_kinds.get((service, k.lower()))
            if hit:
                drops.append(("deprecated", "%s: %s" % (hit[0], hit[1]),
                              self.kcc_link("%s/deprecated_services.json" % GREENFIELD_REL, hit[2])))
                break
        # D3 non-GA only
        if not R["is_ga"]:
            drops.append(("non-ga-only", "only defined in non-GA package %s (generate-types rejects without "
                                         "--allow-non-ga-proto)" % pkg, rec["proto_link"]))
        # D4 v1beta1 brownfield
        if v1beta1:
            drops.append(("v1beta1-brownfield", "served by v1beta1 KCC kind %s" % ", ".join(v1beta1),
                          self.kcc_link(self.crd[v1beta1[0]]["file"])))

        # Operations
        ops, deprecated_rpcs = self.ops_for(pkg, name, patterns, collection, R["plural"], R["synthesized"])
        if deprecated_rpcs:
            flags.append("deprecated-rpc:" + ",".join(deprecated_rpcs))
        C = bool(ops["C"] or ops["Calias"])
        U = bool(ops["U"] or ops["Ualias"])
        D = bool(ops["D"] or ops["Dtag"] or ops["Dalias"])
        G = bool(ops["R"] or ops["Ralias"])
        L = bool(ops["L"])
        for key, flag in (("Calias", "alias:bulk/upsert"), ("Ualias", "alias:bulk/upsert"), ("Ralias", "alias:read"),
                          ("Dalias", "alias:one-shot")):
            if ops[key]:
                flags.append("%s(%s)" % (flag, ",".join(ops[key])))
        if ops["Dtag"] and not ops["D"]:
            flags.append("delete-via:" + ",".join(ops["Dtag"]))
        rec["ops"] = " ".join(x for x, b in (("C", C), ("R", G), ("U", U), ("D", D), ("L", L)) if b)
        rec["rpcs"] = "; ".join("%s=%s" % (k, ",".join(v)) for k, v in ops.items() if v)

        # READ resolution: List counts as Get only for server-generated IDs.
        sg = self.server_generated(pkg, ["Create" + name, "Insert" + name])
        rec["server_generated_id"] = {True: "yes", False: "no", None: "unknown (no Create RPC)"}[sg]
        read_ok, review = G, ""
        if not G:
            if L and sg:
                read_ok = True
                flags.append("read-via-list")
            else:
                review = "review:list-only" if L else "review:no-read"

        # nondeclarative.go verdict (tool parity for annotated resources)
        nd = ""
        pr = self.nd_pattern_reason(patterns)
        if R["synthesized"] is None:
            f, bulk = self.nd_rpc_flags(pkg, name, patterns, collection)
            one_shot_tool = f["C"] and not f["U"] and (not f["D"] or f["Cancel"])
            if pr:
                nd = "pattern %r has %s collection %r" % (pr[3], pr[1], pr[0])
            elif not f["C"] and not f["U"] and (f["G"] or f["D"]):
                nd = "1.1 read-only/discovered"
            elif one_shot_tool:
                nd = "1.2 one-shot (delete=%s, cancel=%s)" % (f["D"], f["Cancel"])
            elif bulk:
                nd = "3.1 bulk RPCs %s" % ",".join(bulk)
            if one_shot_tool or (C and not U and not (ops["D"] or ops["Dtag"])):
                flags.append("nondeclarative:one-shot")
            if bulk:
                drops.append(("nondeclarative:bulk-rpc", "bulk/purge RPCs %s (Criterion 3.1)" % ",".join(bulk),
                              self.kcc_link(ND_REL, self.nd_rule_lines.get("3.1"))))
        elif C and not U and not D:
            flags.append("nondeclarative:one-shot")
        if pr:
            drops.append(("nondeclarative:pattern", "%s collection %r (Criteria 2/3)" % (pr[1], pr[0]),
                          self.kcc_link(ND_REL, pr[2])))
        rec["nondeclarative_go"] = nd

        # D9 / D5
        if not (C or U or D or G or L):
            drops.append(("no-crud", "no Create/Get/Update/Delete/List RPCs in %s" % pkg, rec["proto_link"]))
        elif not C and not U:
            drops.append(("read-only", "neither CREATE nor UPDATE (nondeclarative.go Criterion 1.1)",
                          self.kcc_link(ND_REL, self.nd_rule_lines.get("1.1"))))
        # Policy exclusions (policy_exclusions.json)
        for rx, why, rel, ln in self.policy:
            if rx.match(R["type"]):
                drops.append(("policy:" + why, "excluded by policy review: %s" % why, self.kcc_link(rel, ln)))
                break

        # Earliest KCC googleapis pin with this message.
        sha, users, link = self.first_pin(msg)
        rec["in_local_googleapis"] = sha[:8] if sha else "master only"
        rec["first_googleapis_sha"] = sha or self.googleapis_sha
        rec["first_googleapis_link"] = link or rec["proto_link"]
        if sha != global_pin:
            flags.append("not-in-kcc-pinned-googleapis")

        # Bucket
        drops.sort(key=lambda d: DROP_ORDER.index(d[0]) if d[0] in DROP_ORDER else len(DROP_ORDER))
        if drops:
            rec["bucket"] = BUCKET_DROPPED
            rec["drop_reason"], rec["drop_detail"], rec["drop_link"] = drops[0]
            rec["all_drop_reasons"] = "; ".join(sorted({d[0] for d in drops}, key=lambda c: DROP_ORDER.index(c)
                                                       if c in DROP_ORDER else len(DROP_ORDER)))
        else:
            rec["drop_reason"] = rec["drop_detail"] = rec["drop_link"] = rec["all_drop_reasons"] = ""
            if review:
                rec["bucket"] = BUCKET_REVIEW
                flags.append(review)
            elif C and read_ok and D:
                rec["bucket"] = BUCKET_FM
            elif C and read_ok:
                rec["bucket"] = BUCKET_PM_NO_DELETE
            elif read_ok and U:
                rec["bucket"] = BUCKET_PM_NO_CREATE
            else:
                rec["bucket"] = BUCKET_REVIEW
                flags.append("review:unclassified")
        if rec["kcc_status"] == STATUS_BROWNFIELD:
            flags.append("brownfield")
        rec["flags"] = "; ".join(flags)
        return rec


def _sorted_rows(records):
    """Total, deterministic order (api group, bucket, KCC status, id) with a fixed column order."""
    rows = sorted(records, key=lambda r: (API_GROUP_ORDER.index(r["api_group"]), BUCKET_ORDER.index(r["bucket"]),
                                          STATUS_ORDER.index(r["kcc_status"]), r["id"]))
    return [{col: r.get(col, "") for col in MANAGEABILITY_COLUMNS} for r in rows]


def manageability_counts(rows):
    counts = {}
    for group in API_GROUP_ORDER:
        c = Counter((r["bucket"], r["kcc_status"]) for r in rows if r["api_group"] == group)
        counts[group] = {b: {s: c[(b, s)] for s in STATUS_ORDER if c[(b, s)]}
                         for b in BUCKET_ORDER if any(c[(b, s)] for s in STATUS_ORDER)}
    return counts


def write_manageability_outputs(records, out_dir, meta):
    """Writes manageability_list.json/.csv. Both are byte-for-byte reproducible for the same SHA inputs."""
    os.makedirs(out_dir, exist_ok=True)
    rows = _sorted_rows(records)
    doc = {"meta": dict(meta, counts=manageability_counts(rows)), "resources": rows}
    json_path = os.path.join(out_dir, "manageability_list.json")
    with open(json_path, "w", newline="\n") as f:
        json.dump(doc, f, indent=1)
        f.write("\n")
    csv_path = os.path.join(out_dir, "manageability_list.csv")
    with open(csv_path, "w", newline="") as f:
        w = csv.DictWriter(f, fieldnames=MANAGEABILITY_COLUMNS, lineterminator="\n")
        w.writeheader()
        w.writerows(rows)
    return json_path, csv_path


def print_manageability_summary(records, k):
    width = max(len(b) for b in BUCKET_ORDER)
    for group in API_GROUP_ORDER:
        counts = Counter((r["bucket"], r["kcc_status"]) for r in records if r["api_group"] == group)
        print("\n--- %s (bucket x KCC status) ---" % group)
        print("%-*s %s  total" % (width, "bucket", " ".join("%13s" % s.split(" ")[0] for s in STATUS_ORDER)))
        for b in BUCKET_ORDER:
            vals = [counts.get((b, s), 0) for s in STATUS_ORDER]
            if sum(vals):
                print("%-*s %s  %5d" % (width, b, " ".join("%13s" % (v or "") for v in vals), sum(vals)))
    print("\n--- Drop reasons (Google Cloud / Other Google APIs) ---")
    dropped = [r for r in records if r["bucket"] == BUCKET_DROPPED]
    for reason, n in Counter(r["drop_reason"] for r in dropped).most_common():
        cloud = sum(1 for r in dropped if r["drop_reason"] == reason and r["api_group"] == API_GROUP_CLOUD)
        print("  %5d %5d  %s" % (cloud, n - cloud, reason))
    order = {BUCKET_FM: 0, BUCKET_PM_NO_DELETE: 1, BUCKET_PM_NO_CREATE: 2, BUCKET_REVIEW: 3}
    layers = {"Easy (Leaf)": 0, "Next Layer (1 Parent)": 1, "Next Next Layer (2 Parents)": 2, "Deeply Nested": 3}
    new = [r for r in records if r["kcc_status"] == STATUS_NET_NEW and r["bucket"] in order]
    cloud_new = sorted((r for r in new if r["api_group"] == API_GROUP_CLOUD),
                       key=lambda r: (order[r["bucket"]], layers.get(r["layer"], 4), r["id"]))
    other_new = Counter(r["api_family"] for r in new if r["api_group"] == API_GROUP_OTHER)
    print("\n--- Top %d net-new Google Cloud candidates (%d total) ---" % (k, len(cloud_new)))
    for r in cloud_new[:k]:
        print("  - %s (%s)" % (r["kind"], r["resource_type"]))
        print("    Bucket: %s | Layer: %s | Ops: %s" % (r["bucket"], r["layer"] or "?", r["ops"]))
        flags = [f for f in r["flags"].split("; ") if f]
        if flags:
            print("    Flags:  %s" % ", ".join(flags))
    print("\nOther Google APIs: %d net-new candidates (%s)" % (
        sum(other_new.values()), ", ".join("%s %d" % kv for kv in sorted(other_new.items()))))


def run_manageability(googleapis_sha, proto_root, googleapis_repo, kcc_dir, kcc_sha, temp_dir, out_dir, k,
                      fetch_pins=True):
    full_sha = resolve_commit(googleapis_repo, googleapis_sha) or googleapis_sha
    pb = compile_descriptor_set(proto_root, os.path.join(temp_dir, "googleapis-%s.pb" % full_sha))
    files = decode_descriptor_set(pb)

    # Descriptor sets for KCC's googleapis pins, compiled from the full upstream tree at each pin (not the
    # curated subset that generate-proto.sh builds into .build/), so "pinned" means "googleapis@pin has it".
    pins = []
    skipped = []
    for sha, users in collect_googleapis_pins(kcc_dir):
        if fetch_pins and not resolve_commit(googleapis_repo, sha):
            run_shell("git fetch --depth 1 origin %s" % sha, cwd=googleapis_repo)
        pin_root, pin_full = export_googleapis_protos(googleapis_repo, sha, temp_dir)
        if not pin_root:
            skipped.append(sha)
            continue
        pin_pb = compile_descriptor_set(pin_root, os.path.join(temp_dir, "googleapis-%s.pb" % pin_full))
        names = {m["full"]: fd["name"] for fd in decode_descriptor_set(pin_pb) for m in fd["messages"]}
        when = git_output(googleapis_repo, "log", "-1", "--format=%ct", sha) or "0"
        pins.append((sha, users, names, int(when)))
    if skipped:
        print("[WARN] googleapis pins not available in %s were skipped; the pin columns are incomplete until they "
              "can be fetched (needs network; don't pass --no-fetch-pins): %s"
              % (googleapis_repo, ", ".join(s[:8] for s in skipped)))
    if pins and "global (apis/git.versions)" in pins[0][1]:
        ordered = [pins[0]] + sorted(pins[1:], key=lambda p: (p[3], p[0]))
    else:
        ordered = sorted(pins, key=lambda p: (p[3], p[0]))
    pins = [(sha, users, names) for sha, users, names, _ in ordered]

    # Only trust HEAD when kcc_dir is the root of its own git checkout (not a snapshot nested in another repo).
    top = git_output(kcc_dir, "rev-parse", "--show-toplevel")
    is_root = bool(top) and os.path.realpath(top) == os.path.realpath(kcc_dir)
    kcc_ref = (git_output(kcc_dir, "rev-parse", "HEAD") if is_root else "") or (
        kcc_sha if kcc_sha.upper() != "LOCAL" else "master")
    script_dir = os.path.dirname(os.path.abspath(__file__))
    classifier = ManageabilityClassifier(files, full_sha, proto_root, kcc_dir, kcc_ref, script_dir, pins,
                                         GitBlobReader(googleapis_repo))
    records = classifier.classify()
    # No timestamps or local paths: the JSON only depends on the inputs it records here.
    meta = {"schema_version": 1, "googleapis_sha": full_sha, "kcc_ref": kcc_ref,
            "googleapis_pins": [{"sha": sha, "used_by": users} for sha, users, _ in pins],
            "skipped_googleapis_pins": skipped, "config_fallbacks": sorted(set(classifier.config_fallbacks))}
    json_path, csv_path = write_manageability_outputs(records, out_dir, meta)
    print_manageability_summary(records, k)
    print("\n[SAVED] Manageability list written to %s and %s" % (json_path, csv_path))


def main():
    parser = argparse.ArgumentParser(
        description="Calculate KCC coverage of googleapis resources.",
        epilog="Example: python3 hack/tools/greenfield/calculate_coverage.py master LOCAL 20 --manageability "
               "--googleapis-repo .build/third_party/googleapis")
    parser.add_argument("googleapis_sha", help="git SHA, branch or tag of googleapis/googleapis")
    parser.add_argument("kcc_sha", help="git SHA of k8s-config-connector, or LOCAL for the current directory")
    parser.add_argument("k", nargs="?", type=int, default=10, help="number of top candidates to print (default 10)")
    parser.add_argument("--update-gap", action="store_true", help="write the gap_analysis.txt snapshot (legacy analysis)")
    parser.add_argument("--resources-list", action="store_true", help="write resources_list.json (legacy analysis)")
    parser.add_argument("--manageability", action="store_true",
                        help="classify every googleapis resource (Fully/Partially Manageable, Needs review, Dropped "
                             "with reason) and write manageability_list.json/.csv; requires protoc")
    parser.add_argument("--googleapis-repo",
                        help="use this local googleapis git clone (e.g. .build/third_party/googleapis) instead of "
                             "cloning; protos are exported with `git archive`, so its checkout is not modified")
    parser.add_argument("--output-dir", help="directory for manageability_list.json/.csv (default: this script's dir)")
    parser.add_argument("--no-fetch-pins", action="store_true",
                        help="don't fetch missing googleapis pin commits (apis/git.versions, apis/*/generate.sh); "
                             "skipped pins are listed in the JSON meta and make the pin columns incomplete")
    args = parser.parse_args()

    temp_dir = os.path.join(os.environ.get("TMPDIR", "/tmp"), "kcc_coverage")
    os.makedirs(temp_dir, exist_ok=True)

    if args.googleapis_repo:
        googleapis_repo = os.path.abspath(args.googleapis_repo)
        googleapis_dir, _ = export_googleapis_protos(googleapis_repo, args.googleapis_sha, temp_dir)
        if not googleapis_dir:
            sys.exit("googleapis commit %s not found in %s (git fetch it first)" % (args.googleapis_sha, googleapis_repo))
    else:
        googleapis_dir = os.path.join(temp_dir, "googleapis")
        prepare_repo("https://github.com/googleapis/googleapis.git", googleapis_dir, args.googleapis_sha)
        googleapis_repo = googleapis_dir

    kcc_dir = "." if args.kcc_sha.upper() == "LOCAL" else os.path.join(temp_dir, "kcc")
    if args.kcc_sha.upper() != "LOCAL":
        prepare_repo("https://github.com/GoogleCloudPlatform/k8s-config-connector.git", kcc_dir, args.kcc_sha)

    if args.manageability:
        out_dir = args.output_dir or os.path.dirname(os.path.abspath(__file__))
        run_manageability(args.googleapis_sha, googleapis_dir, googleapis_repo, kcc_dir, args.kcc_sha, temp_dir,
                          out_dir, args.k, fetch_pins=not args.no_fetch_pins)
        if not (args.update_gap or args.resources_list):
            return

    run_gap_analysis(args.googleapis_sha, args.kcc_sha, args.k, args.update_gap, args.resources_list,
                     googleapis_dir, kcc_dir)


if __name__ == "__main__":
    main()
