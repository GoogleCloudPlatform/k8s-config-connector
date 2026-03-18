#!/usr/bin/env python3
import os
import re
import yaml
import sys

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

def get_implemented_refs():
    apis_dir = "apis"
    implemented = set()
    regex = re.compile(r"type\s+([A-Za-z0-9_]+Ref)\s+(struct|interface)")
    for root, _, files in os.walk(apis_dir):
        for file in files:
            if file.endswith(".go"):
                with open(os.path.join(root, file), 'r', encoding='utf-8') as f:
                    content = f.read()
                    matches = regex.findall(content)
                    for match in matches:
                        implemented.add(match[0])
    return implemented

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
    target_refs = set()
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
                                    target_refs.update(find_refs(version["schema"]["openAPIV3Schema"]))
            except Exception as e:
                pass

    if not found_crd:
        print(f"Error: Kind '{target_kind}' not found in CRDs.")
        sys.exit(1)

    # Clean up common generic field
    target_refs.discard("externalRef")

    if not target_refs:
        print(f"Resource '{target_kind}' has no reference fields.")
        sys.exit(0)

    implemented_types = get_implemented_refs()
    impl_lower_to_orig = {t.lower(): t for t in implemented_types}

    missing = []
    found = []

    for field in sorted(target_refs):
        if field.lower() in impl_lower_to_orig:
            found.append((field, impl_lower_to_orig[field.lower()]))
            continue
        
        # Heuristic matching: Try to match with group prefix or exact suffix
        # ComputeInstance -> networkRef -> ComputeNetworkRef
        # Prefer refs under apis/refs/v1beta1/ if possible
        matched = False
        candidates = []
        for t in implemented_types:
            if t.lower().endswith(field.lower()):
                candidates.append(t)
        
        if candidates:
            # Sort candidates to prefer those that are shorter (more exact) or start with the group prefix
            # This is a simple heuristic, we just take the first one after sorting
            candidates.sort(key=lambda x: len(x))
            found.append((field, candidates[0]))
            matched = True
            
        if not matched:
            missing.append(field)

    print(f"=== Reference Status for {target_kind} ===")
    print(f"\nImplemented ({len(found)}):")
    for field, t in found:
        print(f"  [x] {field} -> {t}")

    print(f"\nNot Implemented ({len(missing)}):")
    for field in missing:
        print(f"  [ ] {field}")

if __name__ == "__main__":
    main()
