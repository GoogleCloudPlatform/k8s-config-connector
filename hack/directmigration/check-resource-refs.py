#!/usr/bin/env python3
import os
import re
import yaml
import sys

def find_refs(schema, path="spec"):
    # Returns a list of (path, field_name)
    refs = []
    if isinstance(schema, dict):
        if 'properties' in schema:
            for prop, prop_schema in schema['properties'].items():
                current_path = f"{path}.{prop}" if path else prop
                if prop.endswith("Ref"):
                    refs.append((current_path, prop))
                else:
                    refs.extend(find_refs(prop_schema, current_path))
        elif 'items' in schema:
            refs.extend(find_refs(schema['items'], f"{path}[]"))
    return refs

def get_implemented_refs():
    dirs_to_search = ["apis", "pkg/clients/generated/apis"]
    implemented = {}
    regex = re.compile(r"type\s+([A-Za-z0-9_]+Ref)\s+(struct|interface)")
    for d in dirs_to_search:
        if not os.path.exists(d):
            continue
        for root, _, files in os.walk(d):
            for file in files:
                if file.endswith(".go") and not file.endswith("_types.go"):
                    filepath = os.path.join(root, file)
                    with open(filepath, 'r', encoding='utf-8') as f:
                        content = f.read()
                        matches = regex.findall(content)
                        for match in matches:
                            name = match[0]
                            lower = name.lower()
                            if lower not in implemented:
                                implemented[lower] = []
                            implemented[lower].append((name, filepath))
    return implemented

def get_kind_definitions():
    # Scans apis/ for GroupVersion.WithKind("<Kind>") to build authoritative map
    # Kind -> list of (package_path, group_name)
    kind_defs = {}
    dirs_to_search = ["apis", "pkg/clients/generated/apis"]
    regex = re.compile(r'GroupVersion\.WithKind\("([^"]+)"\)')
    regex_alias = re.compile(r'GroupVersionKind\s*\{\s*Group:\s*[^,]+,\s*Version:\s*[^,]+,\s*Kind:\s*"([^"]+)"\s*\}')
    
    for d in dirs_to_search:
        if not os.path.exists(d):
            continue
        for root, _, files in os.walk(d):
            for file in files:
                if file.endswith(".go"):
                    filepath = os.path.join(root, file)
                    with open(filepath, 'r', encoding='utf-8') as f:
                        content = f.read()
                        for match in list(regex.finditer(content)) + list(regex_alias.finditer(content)):
                            kind = match.group(1)
                            # Extract group name from path
                            parts = filepath.split('/')
                            try:
                                apis_idx = parts.index("apis")
                                group_name = parts[apis_idx + 1]
                            except:
                                group_name = ""
                            
                            if kind not in kind_defs:
                                kind_defs[kind] = []
                            kind_defs[kind].append((filepath, group_name))
    return kind_defs

def get_tf_reference_mapping(target_kind):
    # Parses config/servicemappings/*.yaml
    # Returns a dict of { field_name: target_kind }
    mappings = {}
    crds_dir = "config/servicemappings"
    if not os.path.exists(crds_dir):
        return mappings

    for filename in os.listdir(crds_dir):
        if not filename.endswith(".yaml"):
            continue
        filepath = os.path.join(crds_dir, filename)
        with open(filepath, 'r') as f:
            try:
                for doc in yaml.safe_load_all(f):
                    if doc and doc.get("kind") == "ServiceMapping":
                        for resource in doc.get("spec", {}).get("resources", []):
                            if resource.get("kind") == target_kind:
                                refs = resource.get("resourceReferences", [])
                                for ref in refs:
                                    # tf field references can define the expected KCC Kind directly
                                    key = ref.get("key")
                                    if key and "gvk" in ref and "kind" in ref["gvk"]:
                                        mappings[key] = ref["gvk"]["kind"]
                                    
                                    # Or they can be in a 'types' list
                                    if "types" in ref:
                                        for t in ref["types"]:
                                            t_key = t.get("key")
                                            if t_key and "gvk" in t and "kind" in t["gvk"]:
                                                mappings[t_key] = t["gvk"]["kind"]
            except Exception as e:
                pass
    return mappings

def build_dcl_to_kcc_kind_map():
    # parses pkg/dcl/metadata/metadata.go to map "Service/Resource" to KCC Kind
    dcl_to_kcc = {}
    metadata_file = "pkg/dcl/metadata/metadata.go"
    if not os.path.exists(metadata_file):
        return dcl_to_kcc
        
    with open(metadata_file, "r") as f:
        content = f.read()

    # Regex findall ServiceBlocks
    service_pattern = re.compile(r'Name:\s+"([^"]+)",(.*?)(?=Name:\s+"|$)', re.DOTALL)
    for match in service_pattern.finditer(content):
        service_name = match.group(1)
        block = match.group(2)
        
        dcl_service = service_name
        sm_match = re.search(r'ServiceNameUsedByDCL:\s+"([^"]+)",', block)
        if sm_match:
            dcl_service = sm_match.group(1).capitalize()
            
        resources = re.findall(r'Kind:\s+"([^"]+)"(.*?)(?=}?,)', block, re.DOTALL)
        for kind, rest in resources:
            dcl_type = None
            dcl_type_match = re.search(r'DCLType:\s+"([^"]+)"', rest)
            if dcl_type_match:
                dcl_type = dcl_type_match.group(1)
            else:
                if kind.startswith(service_name):
                    dcl_type = kind[len(service_name):]
                else:
                    dcl_type = kind
            
            dcl_resource = f"{dcl_service}/{dcl_type}".lower()
            dcl_to_kcc[dcl_resource] = kind
            
    # Hardcode some common defaults if missing from metadata
    defaults = {
        "cloudresourcemanager/project": "Project",
        "iam/serviceaccount": "IAMServiceAccount",
        "storage/bucket": "StorageBucket",
        "compute/network": "ComputeNetwork",
        "compute/subnetwork": "ComputeSubnetwork",
        "compute/image": "ComputeImage",
        "compute/instance": "ComputeInstance",
        "compute/nodegroup": "ComputeNodeGroup",
        "container/cluster": "ContainerCluster",
        "container/nodepool": "ContainerNodePool",
        "metastore/service": "MetastoreService",
        "cloudkms/cryptokey": "KMSCryptoKey"
    }
    for k, v in defaults.items():
        if k not in dcl_to_kcc:
            dcl_to_kcc[k] = v
            
    return dcl_to_kcc

def find_dcl_refs(schema, dcl_to_kcc, path=""):
    refs = {}
    if 'properties' in schema:
        for prop, prop_schema in schema['properties'].items():
            if 'x-dcl-references' in prop_schema:
                ref_name = prop
                if ref_name.endswith("Name"):
                    ref_name = ref_name[:-4]
                if not ref_name.endswith("Ref"):
                    ref_name += "Ref"
                
                target_dcl_resource = prop_schema['x-dcl-references'][0]['resource']
                target_kind = dcl_to_kcc.get(target_dcl_resource.lower(), target_dcl_resource.split('/')[-1])
                refs[ref_name] = target_kind
            else:
                refs.update(find_dcl_refs(prop_schema, dcl_to_kcc, path + "." + prop))
    elif 'items' in schema:
        refs.update(find_dcl_refs(schema['items'], dcl_to_kcc, path + "[]"))
    return refs

def get_dcl_reference_mapping(target_kind):
    mappings = {}
    dcl_to_kcc = build_dcl_to_kcc_kind_map()
    
    # reverse map to find DCL service/resource
    target_dcl_resource = None
    for k, v in dcl_to_kcc.items():
        if v == target_kind:
            target_dcl_resource = k
            break
            
    if not target_dcl_resource:
        return mappings
        
    service, resource = target_dcl_resource.split('/', 1)
    
    # search for the DCL yaml file
    dcl_dir = "third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google"
    if not os.path.exists(dcl_dir):
        return mappings

    for root, _, files in os.walk(dcl_dir):
        for file in files:
            if file.endswith(".yaml"):
                filepath = os.path.join(root, file)
                with open(filepath, 'r') as f:
                    try:
                        dcl = yaml.safe_load(f)
                        if not dcl or 'components' not in dcl or 'schemas' not in dcl['components']:
                            continue
                            
                        # find schema
                        schema_name = None
                        for s_name in dcl['components']['schemas']:
                            if s_name.lower() == resource:
                                schema_name = s_name
                                break
                        
                        if schema_name:
                            schema = dcl['components']['schemas'][schema_name]
                            refs = find_dcl_refs(schema, dcl_to_kcc)
                            mappings.update(refs)
                            
                    except Exception as e:
                        pass
    return mappings

def get_kind_to_service_map():
    mapping = {
        "KMSCryptoKey": "kms",
        "IAMServiceAccount": "iam",
        "Project": "resourcemanager",
        "Folder": "resourcemanager",
        "Organization": "resourcemanager",
        "ComputeNodeGroup": "compute",
        "ContainerNodePool": "container"
    }
    crds_dir = "config/crds/resources"
    if not os.path.exists(crds_dir):
        return mapping
    for filename in os.listdir(crds_dir):
        if not filename.endswith(".yaml"):
            continue
        filepath = os.path.join(crds_dir, filename)
        with open(filepath, 'r') as f:
            try:
                for doc in yaml.safe_load_all(f):
                    if doc and doc.get("kind") == "CustomResourceDefinition":
                        kind = doc["spec"]["names"]["kind"]
                        group = doc["spec"]["group"]
                        service = group.split(".")[0]
                        mapping[kind] = service
            except Exception as e:
                pass
    return mapping

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 hack/check-resource-refs.py <Kind>")
        sys.exit(1)
        
    target_kind = sys.argv[1]
    
    crds_dir = "config/crds/resources"
    if not os.path.exists(crds_dir):
        print(f"Error: {crds_dir} not found.")
        sys.exit(1)

    # Find the target CRD and extract its refs
    target_refs = []
    found_crd = False
    
    for filename in os.listdir(crds_dir):
        if not filename.endswith(".yaml"):
            continue
        filepath = os.path.join(crds_dir, filename)
        with open(filepath, 'r') as f:
            try:
                for doc in yaml.safe_load_all(f):
                    if doc and doc.get("kind") == "CustomResourceDefinition":
                        kind = doc["spec"]["names"]["kind"]
                        if kind == target_kind:
                            found_crd = True
                            for version in doc["spec"]["versions"]:
                                if "schema" in version and "openAPIV3Schema" in version["schema"]:
                                    target_refs.extend(find_refs(version["schema"]["openAPIV3Schema"], ""))
            except Exception as e:
                pass

    if not found_crd:
        print(f"Error: Kind '{target_kind}' not found in CRDs.")
        sys.exit(1)

    # Clean up common generic field
    target_refs = [r for r in target_refs if r[1] != "externalRef"]

    if not target_refs:
        print(f"Resource '{target_kind}' has no reference fields.")
        sys.exit(0)

    # Get expected target kinds from TF or DCL mappings
    tf_mappings = get_tf_reference_mapping(target_kind)
    dcl_mappings = get_dcl_reference_mapping(target_kind)
    
    # Merge mappings, prioritizing TF if there's an overlap (unlikely)
    expected_targets = {**dcl_mappings, **tf_mappings}

    implemented_refs = get_implemented_refs()
    kind_defs = get_kind_definitions()
    kind_to_service = get_kind_to_service_map()

    missing = []
    found = []

    for path, field in sorted(target_refs):
        # 1. Check if we know the expected Target Kind
        target_kind_name = expected_targets.get(field)
        
        matched = False
        candidates = [] # list of (name, filepath)
        
        if target_kind_name:
            if target_kind_name not in kind_defs and target_kind_name not in kind_to_service:
                missing.append(f"{field} ({path}) - Kind '{target_kind_name}' doesn't seem to be supported in KCC at all.")
                continue

            # Kind is supported! Find all valid API groups for this Kind
            valid_groups = {g[1] for g in kind_defs.get(target_kind_name, [])}
            expected_service = kind_to_service.get(target_kind_name)
            if expected_service:
                valid_groups.add(expected_service)
            
            expected_ref_type = f"{target_kind_name}Ref".lower()
            possible = []
            
            if expected_ref_type in implemented_refs:
                possible.extend(implemented_refs[expected_ref_type])
            else:
                # Some types omit the service prefix
                for t_lower, t_list in implemented_refs.items():
                    if expected_ref_type.endswith(t_lower):
                        possible.extend(t_list)
            
            # STRICT FILTERING: The reference type MUST exist in the Kind's API group folder, refs, or common.
            if possible:
                valid_prefixes = [f"apis/{g}/" for g in valid_groups if g]
                valid_prefixes.extend(["apis/refs/", "apis/common/", "pkg/clients/generated/apis/"])
                valid_prefixes = tuple(valid_prefixes)
                
                possible = [c for c in possible if c[1].startswith(valid_prefixes)]

            if possible:
                # Sort by length (longer match is better) and then by being exported (starts with uppercase)
                possible.sort(key=lambda x: (len(x[0]), x[0][0].isupper()), reverse=True)
                best = possible[0]
                # Filter for those with the same length and capitalization status
                candidates = [c for c in possible if len(c[0]) == len(best[0]) and c[0][0].isupper() == best[0][0].isupper()]
                matched = True
        else:
            # Fallback to the old heuristic matching
            possible = []
            if field.lower() in implemented_refs:
                possible.extend(implemented_refs[field.lower()])
            else:
                for t_lower, t_list in implemented_refs.items():
                    if t_lower.endswith(field.lower()):
                        possible.extend(t_list)
            
            if possible:
                possible.sort(key=lambda x: (len(x[0]), not x[0][0].isupper()))
                best = possible[0]
                candidates = [c for c in possible if len(c[0]) == len(best[0]) and c[0][0].isupper() == best[0][0].isupper()]
                matched = True
            
        if matched:
            found.append((path, field, candidates))
        else:
            if target_kind_name:
                missing.append(f"{field} ({path}) (expected {target_kind_name}Ref)")
            else:
                missing.append(f"{field} ({path})")

    print(f"=== Reference Status for {target_kind} ===")
    print(f"\nImplemented ({len(found)}):")
    for path, field, types in found:
        type_strs = [f"{t} ({f})" for t, f in types]
        print(f"  [x] {field} ({path}) -> {', '.join(type_strs)}")

    print(f"\nNot Implemented ({len(missing)}):")
    for msg in missing:
        print(f"  [ ] {msg}")

if __name__ == "__main__":
    main()
