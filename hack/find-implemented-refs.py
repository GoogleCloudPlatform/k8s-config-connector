#!/usr/bin/env python3
import os
import re
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

def get_expected_refs():
    crds_dir = "config/crds/resources"
    expected = set()
    for filename in os.listdir(crds_dir):
        if not filename.endswith(".yaml"):
            continue
        filepath = os.path.join(crds_dir, filename)
        with open(filepath, 'r') as f:
            try:
                for doc in yaml.safe_load_all(f):
                    if doc and doc.get("kind") == "CustomResourceDefinition":
                        for version in doc["spec"]["versions"]:
                            if "schema" in version and "openAPIV3Schema" in version["schema"]:
                                expected.update(find_refs(version["schema"]["openAPIV3Schema"]))
            except Exception as e:
                pass
    return expected

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
    expected_fields = get_expected_refs()
    implemented_types = get_implemented_refs()
    
    # externalRef is special, it's inside other refs
    expected_fields.discard("externalRef")
    
    impl_lower_to_orig = {t.lower(): t for t in implemented_types}
    
    missing = []
    found = []
    
    for field in sorted(expected_fields):
        if field.lower() in impl_lower_to_orig:
            found.append((field, impl_lower_to_orig[field.lower()]))
            continue
        
        # Try finding a type that ends with this field (e.g., networkRef -> ComputeNetworkRef)
        matched = False
        for t in implemented_types:
            if t.lower().endswith(field.lower()):
                found.append((field, t))
                matched = True
                break
        
        if not matched:
            missing.append(field)
            
    print("=== Implemented Reference Fields ===")
    for field, t in found:
        print(f"{field} -> {t}")
        
    print("\n=== Unimplemented / Unmatched Reference Fields ===")
    for field in missing:
        print(f"{field}")

if __name__ == "__main__":
    main()
