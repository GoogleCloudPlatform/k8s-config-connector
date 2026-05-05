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

def run_shell(cmd, cwd=None):
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd=cwd)
    return result

def get_gcp_resources(googleapis_dir):
    resources = {} 
    
    # Even more robust regex to find type and pattern
    # We'll search for 'type: "' and 'pattern: "' inside the same file
    # and try to group them by the resource block.
    
    for root, _, files in os.walk(googleapis_dir):
        if "third_party" in root: continue
        for file in files:
            if not file.endswith(".proto"): continue
            path = os.path.join(root, file)
            with open(path, 'r', errors='ignore') as f:
                content = f.read()
                
                # Split by possible resource markers to isolate blocks
                blocks = re.split(r'google\.api\.resource', content)
                for block in blocks:
                    type_match = re.search(r'type:\s*"([^"]+)"', block)
                    if type_match:
                        rtype = type_match.group(1)
                        if "/" not in rtype: continue
                        
                        if rtype not in resources:
                            service, name = rtype.split('/')
                            resources[rtype] = {
                                'service': service.split('.')[0],
                                'name': name,
                                'ops': set(),
                                'patterns': []
                            }
                        
                        p_matches = re.findall(r'pattern:\s*"([^"]+)"', block)
                        resources[rtype]['patterns'].extend(p_matches)

    # 2. Get all RPCs
    cmd_rpcs = f"grep -r '^[[:space:]]*rpc ' {googleapis_dir} --include='*.proto'"
    res_rpcs = run_shell(cmd_rpcs)
    all_rpcs = set()
    for line in res_rpcs.stdout.splitlines():
        rpc_match = re.search(r'rpc\s+([A-Za-z0-9]+)', line)
        if rpc_match:
            all_rpcs.add(rpc_match.group(1))

    # 3. Match ops to resources
    for rtype, info in resources.items():
        name = info['name']
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
            for kcc_s, aliases in service_aliases.items():
                if gcp_service_base in aliases or gcp_service_base == kcc_s:
                    if kcc_s in kcc_map:
                        matched_kcc_service = kcc_s
                        break
        
        if not matched_kcc_service: continue
        name_norm = gcp_name.lower()
        for kcc_kind in kcc_map[matched_kcc_service]:
            kind_norm = kcc_kind.lower()
            prefixes = [matched_kcc_service.replace("-", "").lower(), "gcp", "google", "cloud", "bigquery", "api"]
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

def prepare_repo(repo_url, target_dir, sha):
    if not os.path.exists(target_dir):
        run_shell(f"git clone --depth 100 {repo_url} {target_dir}")
    run_shell("git fetch origin master --depth 100", cwd=target_dir)
    run_shell(f"git checkout -f {sha}", cwd=target_dir)

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

    gcp_resources = get_gcp_resources(googleapis_dir)
    kcc_resources = get_kcc_resources(kcc_dir)
    covered = match_resources(gcp_resources, kcc_resources)
    
    all_gcp = set(gcp_resources.keys())
    missing = all_gcp - covered
    
    create_delete_gcp = {rtype for rtype, info in gcp_resources.items() if 'CREATE' in info['ops'] and 'DELETE' in info['ops']}
    missing_lifecycle = create_delete_gcp - covered
    missing_leaves = {rtype for rtype in missing_lifecycle if is_leaf(gcp_resources[rtype]['patterns'])}
    
    print("\n--- Coverage Summary ---")
    print(f"Total GCP Resources:      {len(all_gcp)}")
    print(f"  Implemented in KCC:     {len(covered)}")
    print(f"  Missing from KCC:       {len(all_gcp - covered)}")
    print(f"  Coverage:               {len(covered)/max(1, len(all_gcp)):.2%}")
    
    print("\n--- Full Lifecycle Missing ---")
    print(f"Total Manageable Missing: {len(missing_lifecycle)}")
    print(f"  Leaf (Easy) Missing:    {len(missing_leaves)}")

    print(f"\n--- Next {k} Easiest Resources to Implement ---")
    print("(Criteria: Full Lifecycle + Project/Folder/Org/Location parent)")
    for m in sorted(list(missing_leaves))[:k]:
        patterns = ", ".join(gcp_resources[m]['patterns'])
        print(f"  - {m}")
        print(f"    Patterns: {patterns}")

if __name__ == "__main__":
    main()
