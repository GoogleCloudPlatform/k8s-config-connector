#!/usr/bin/env python3
import os
import yaml
import sys
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

def main():
    target_kind = sys.argv[1] if len(sys.argv) > 1 else None
    
    crds_dir = "config/crds/resources"
    if not os.path.exists(crds_dir):
        print(f"Error: {crds_dir} not found.")
        sys.exit(1)

    # All known kinds from CRDs: kind -> group
    known_kinds = {}
    
    # Dependencies: kind -> set of kinds
    dependencies = defaultdict(set)
    raw_refs = defaultdict(set)

    # 1. First pass: Collect all known kinds
    crd_docs = []
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
            except Exception as e:
                print(f"Error parsing {filepath}: {e}", file=sys.stderr)

    # 2. Second pass: Extract refs and map to kinds
    for doc in crd_docs:
        kind = doc["spec"]["names"]["kind"]
        group = doc["spec"]["group"]
        
        for version in doc["spec"]["versions"]:
            if "schema" in version and "openAPIV3Schema" in version["schema"]:
                refs = find_refs(version["schema"]["openAPIV3Schema"])
                raw_refs[kind].update(refs)
                
                # Attempt to map each ref field to a known kind
                for ref in refs:
                    if ref == "externalRef": # 'external' is a common inner field of refs, not a separate reference type
                        continue
                    
                    ref_name = ref[:-3] # remove 'Ref' suffix
                    matched_kind = None
                    group_prefix = group.split('.')[0]
                    
                    for k in known_kinds.keys():
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

    # 3. Output dependencies
    if target_kind:
        if target_kind not in known_kinds:
            print(f"Error: Kind '{target_kind}' not found in CRDs.")
            sys.exit(1)
            
        print(f"Dependency tree for {target_kind}:")
        visited = set()
        
        def print_tree(k, depth=0):
            if k in visited:
                print(f"{'  ' * depth}- {k} (cycle detected/already visited)")
                return
            visited.add(k)
            deps = sorted(dependencies.get(k, []))
            if deps:
                for d in deps:
                    print(f"{'  ' * depth}- {d}")
                    print_tree(d, depth + 1)
        
        # Start the tree
        if dependencies.get(target_kind):
            for d in sorted(dependencies[target_kind]):
                print(f"- {d}")
                print_tree(d, 1)
        else:
            print("(No dependencies)")

    else:
        print("Usage: python dependency-graph.py <Kind>")
        print("Example: python dependency-graph.py ComputeInstance\n")
        print("All Resource Dependencies:")
        for k, deps in sorted(dependencies.items()):
            if deps:
                print(f"{k} -> {', '.join(sorted(deps))}")

if __name__ == "__main__":
    main()
