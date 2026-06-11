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
import subprocess
import re
import sys
import json
from datetime import datetime

def run_shell(cmd, cwd=None):
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd=cwd)
    return result

def get_gcp_resources(googleapis_dir):
    resources = {} 
    service_rpcs = {} # Map service to its RPCs
    
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
                
                # Extract RPCs in this file
                rpc_matches = re.findall(r'rpc\s+([A-Za-z0-9]+)', content)
                service_rpcs[service_pkg].update(rpc_matches)

                # Split by possible resource markers to isolate blocks
                blocks = re.split(r'google\.api\.resource', content)
                for block in blocks:
                    type_match = re.search(r'type:\s*"([^"]+)"', block)
                    if type_match:
                        rtype = type_match.group(1)
                        if "/" not in rtype: continue
                        
                        if rtype not in resources:
                            service_name, name = rtype.split('/')
                            resources[rtype] = {
                                'service': service_name.split('.')[0],
                                'pkg': service_pkg,
                                'name': name,
                                'ops': set(),
                                'patterns': []
                            }
                        
                        p_matches = re.findall(r'pattern:\s*"([^"]+)"', block)
                        resources[rtype]['patterns'].extend(p_matches)

    # 3. Match ops to resources within their service package
    for rtype, info in resources.items():
        name = info['name']
        pkg = info['pkg']
        if pkg not in service_rpcs: continue
        
        all_rpcs = service_rpcs[pkg]
        create_variants = [f"Create{name}", f"Upsert{name}", f"BatchCreate{name}"]
        for v in create_variants:
            if v in all_rpcs:
                info['ops'].add('CREATE')
                break
        delete_variants = [f"Delete{name}", f"Finish{name}", f"Abort{name}", f"Cancel{name}", f"Terminate{name}", f"Destroy{name}"]
        for v in delete_variants:
            if v in all_rpcs:
                info['ops'].add('DELETE')
                break
                
    return resources

def get_kcc_resources(kcc_dir):
    resources = []
    crd_dir = os.path.join(kcc_dir, "config/crds/resources")
    if not os.path.exists(crd_dir):
        return []
    for file in os.listdir(crd_dir):
        if file.endswith(".yaml"):
            path = os.path.join(crd_dir, file)
            with open(path, 'r') as f:
                content = f.read()
                group_match = re.search(r'group:\s*([^\s]+)', content)
                kind_section = re.search(r'names:.*?\s+kind:\s*([^\s]+)', content, re.DOTALL)
                if group_match and kind_section:
                    group = group_match.group(1).strip('"\'')
                    kind = kind_section.group(1).strip('"\'')
                    resources.append({'group': group, 'kind': kind})
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
    }

    for gcp_type, info in gcp_resources.items():
        gcp_service_base = info['service']
        gcp_name = info['name']
        matched_kcc_service = None
        if gcp_service_base in kcc_map:
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

def prepare_repo(repo_url, target_dir, sha):
    if not os.path.exists(target_dir):
        run_shell(f"git clone --depth 100 {repo_url} {target_dir}")
    run_shell("git fetch origin master --depth 100", cwd=target_dir)
    run_shell(f"git checkout -f {sha}", cwd=target_dir)

def canonicalize_resource_name(name):
    """Strips hierarchical prefixes (Global, Regional, Zonal) from resource names."""
    prefixes = ["Global", "Regional", "Zonal"]
    for p in prefixes:
        if name.startswith(p) and len(name) > len(p) and name[len(p)].isupper():
            return name[len(p):]
    return name

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
        
    return unified

def main():
    if len(sys.argv) < 3:
        print("Usage: python calculate_coverage.py <googleapis_sha> <kcc_sha> [k]")
        sys.exit(1)
        
    googleapis_sha = sys.argv[1]
    kcc_sha = sys.argv[2]
    k = int(sys.argv[3]) if len(sys.argv) > 3 else 10
    
    temp_dir = "/usr/local/google/home/acpana/.gemini/tmp/k8s-config-connector-11/kcc_coverage"
    os.makedirs(temp_dir, exist_ok=True)
    
    googleapis_dir = os.path.join(temp_dir, "googleapis")
    prepare_repo("https://github.com/googleapis/googleapis.git", googleapis_dir, googleapis_sha)

    kcc_dir = "." if kcc_sha.upper() == "LOCAL" else os.path.join(temp_dir, "kcc")
    if kcc_sha.upper() != "LOCAL":
        prepare_repo("https://github.com/GoogleCloudPlatform/k8s-config-connector.git", kcc_dir, kcc_sha)

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
    }

    for key, info in gcp_resources.items():
        gcp_service_base = info['service']
        gcp_name = info['name']
        matched_kcc_service = None
        if gcp_service_base in kcc_map:
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
                covered.add(key)
                break
            found_prefix_match = False
            for p in prefixes:
                if kind_norm.startswith(p) and len(kind_norm) > len(p):
                    if kind_norm[len(p):] == name_norm:
                        covered.add(key)
                        found_prefix_match = True
                        break
            if found_prefix_match: break
            if kind_norm.endswith('s') and kind_norm[:-1] == name_norm:
                covered.add(key)
                break
    
    # Categorization
    all_gcp_keys = set(gcp_resources.keys())
    covered = {key for key in covered if key in all_gcp_keys}
    missing = all_gcp_keys - covered
    
    # Calculate how many raw rtypes were unified
    total_raw_rtypes = sum(len(info['rtypes']) for info in gcp_resources.values())
    unification_count = total_raw_rtypes - len(gcp_resources)

    # Manageable = Has Create OR Delete
    manageable_gcp = {key for key, info in gcp_resources.items() if 'CREATE' in info['ops'] or 'DELETE' in info['ops']}
    missing_manageable = manageable_gcp - covered

    # Fully Manageable = Has Create AND Delete
    fully_manageable_gcp = {key for key, info in gcp_resources.items() if 'CREATE' in info['ops'] and 'DELETE' in info['ops']}
    missing_fully_manageable = fully_manageable_gcp - covered

    # Easy = Fully Manageable AND Leaf pattern
    missing_easy = {key for key in missing_fully_manageable if is_leaf(gcp_resources[key]['patterns'])}
    
    # Next Layer = Fully Manageable AND 1 parent deep
    missing_next_layer = {key for key in missing_fully_manageable if is_next_layer(gcp_resources[key]['patterns'])}
    
    # Next Next Layer = Fully Manageable AND 2 parents deep
    missing_next_next_layer = {key for key in missing_fully_manageable if is_next_next_layer(gcp_resources[key]['patterns'])}
    
    # Generate Gap Analysis Table for tracking
    gap_file = os.path.join(os.path.dirname(__file__), "gap_analysis.txt")
    now = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    
    analysis_lines = [
        f"Gap Analysis Snapshot - {now}",
        f"GoogleAPIs SHA: {googleapis_sha}",
        f"KCC SHA:        {kcc_sha}",
        "-" * 55,
        f"{'Metric':<30} | {'Value':<10}",
        "-" * 55,
        f"{'Total GCP Resources (Raw)':<30} | {all_gcp_raw_count:<10}",
        f"{'Unified (Hierarchical)':<30} | {unification_count:<10}",
        f"{'Processed Resources (Unified)':<30} | {len(all_gcp_keys):<10}",
        f"{'Skipped (Policy)':<30} | {skipped_count:<10}",
        f"{'Implemented in KCC':<30} | {len(covered):<10}",
        f"{'Missing from KCC':<30} | {len(missing):<10}",
        "-" * 55,
        f"{'Missing Manageable':<30} | {len(missing_manageable):<10}",
        f"{'Missing Fully Manageable':<30} | {len(missing_fully_manageable):<10}",
        f"{'Missing Next Layer':<30} | {len(missing_next_layer):<10}",
        f"{'Missing Next Next Layer':<30} | {len(missing_next_next_layer):<10}",
        "-" * 55,
        f"{'Current Coverage':<30} | {len(covered)/max(1, len(all_gcp_keys)):.2%}",
        ""
    ]
    
    with open(gap_file, 'w') as f:
        f.write("\n".join(analysis_lines))

    print("\n--- Coverage Summary ---")
    print(f"Total GCP Resources (Raw): {all_gcp_raw_count}")
    print(f"  - Unified:              {unification_count}")
    print(f"  - Processed (Unified):  {len(all_gcp_keys)}")
    print(f"  - Skipped (Policy):     {skipped_count}")
    print(f"  - Implemented in KCC:   {len(covered)}")
    print(f"  - Missing from KCC:     {len(missing)}")
    print(f"  - Coverage:             {len(covered)/max(1, len(all_gcp_keys)):.2%}")
    
    print("\n--- Gap Breakdown (Missing Resources) ---")
    print(f"Total Missing:            {len(missing)}")
    print(f"  - Manageable:           {len(missing_manageable)} (Has Create OR Delete)")
    print(f"  - Fully Manageable:     {len(missing_fully_manageable)} (Has Create AND Delete)")
    print(f"  - Next Next Layer:      {len(missing_next_next_layer)} (Fully Manageable + 2 Parents)")
    print(f"  - Next Layer Targets:   {len(missing_next_layer)} (Fully Manageable + 1 Parent)")
    print(f"  - Easy Targets:         {len(missing_easy)} (Fully Manageable + Leaf Pattern)")
    print(f"\n[SAVED] Gap analysis snapshot written to {gap_file}")

    print(f"\n--- Next {k} Easiest Resources to Implement ---")
    print("(Criteria: Easy Targets)")
    for m in sorted(list(missing_easy))[:k]:
        patterns = ", ".join(gcp_resources[m]['patterns'])
        print(f"  - {m}")
        print(f"    Patterns: {patterns}")

if __name__ == "__main__":
    main()
