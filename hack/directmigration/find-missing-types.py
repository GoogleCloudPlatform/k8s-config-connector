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

import os
import re
import yaml
import sys

def get_crd_resources(crds_dir="config/crds/resources"):
    """
    Parses all CRDs to find the Kind of each resource.
    Returns a list of tuples: (kind, group, filepath)
    """
    if not os.path.exists(crds_dir):
        print(f"Error: {crds_dir} not found.", file=sys.stderr)
        sys.exit(1)

    resources = []
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
                        resources.append((kind, group, filename))
            except Exception as e:
                print(f"Error parsing {filepath}: {e}", file=sys.stderr)
    return resources

def get_implemented_types(apis_dir="apis"):
    """
    Scans the apis/ directory for *_types.go files.
    Returns a map of Kind -> list of package paths where it's defined.
    """
    if not os.path.exists(apis_dir):
        print(f"Error: {apis_dir} not found.", file=sys.stderr)
        sys.exit(1)

    implemented_kinds = {}
    
    # Regex to find standard struct definitions for KRM resources
    # e.g., type ComputeInstance struct {
    struct_regex = re.compile(r"type\s+([A-Za-z0-9_]+)\s+struct\s*\{")
    
    for root, _, files in os.walk(apis_dir):
        for file in files:
            if file.endswith("_types.go"):
                filepath = os.path.join(root, file)
                with open(filepath, 'r', encoding='utf-8') as f:
                    content = f.read()
                    matches = struct_regex.findall(content)
                    for kind in matches:
                        # Skip typical sub-structs that end in Spec, Status, List, etc.
                        # We are looking for the root resource Kind
                        if kind.endswith("Spec") or kind.endswith("Status") or kind.endswith("List") or kind.endswith("ObservedState"):
                            continue
                            
                        if kind not in implemented_kinds:
                            implemented_kinds[kind] = []
                        implemented_kinds[kind].append(filepath)
    
    return implemented_kinds

def main():
    import argparse
    parser = argparse.ArgumentParser(description="Find CRDs missing _types.go or _reference.go implementations.")
    parser.add_argument("-o", "--output", help="Output file to write missing Kind names to (one per line).")
    args = parser.parse_args()

    crd_resources = get_crd_resources()
    implemented_types = get_implemented_types()

    missing_types = []
    missing_references = []
    found_fully_implemented = []

    for kind, group, filename in sorted(crd_resources, key=lambda x: x[0]):
        if kind in implemented_types:
            # Check if any of the _types.go files have a corresponding _reference.go file
            has_reference = False
            for filepath in implemented_types[kind]:
                ref_filepath = filepath.replace("_types.go", "_reference.go")
                if os.path.exists(ref_filepath):
                    has_reference = True
                    break
            
            if has_reference:
                found_fully_implemented.append((kind, group, implemented_types[kind]))
            else:
                missing_references.append((kind, group, implemented_types[kind]))
        else:
            missing_types.append((kind, group, filename))

    if args.output:
        with open(args.output, 'w') as f:
            for kind, _, _ in missing_types:
                f.write(f"{kind}\n")
            for kind, _, _ in missing_references:
                f.write(f"{kind}:missing_reference\n")
        print(f"Successfully wrote {len(missing_types) + len(missing_references)} kinds to {args.output}")
    else:
        print(f"Total CRDs found: {len(crd_resources)}")
        print(f"Resources fully implemented (_types.go and _reference.go): {len(found_fully_implemented)}")
        print(f"Resources MISSING _types.go: {len(missing_types)}")
        print(f"Resources WITH _types.go but MISSING _reference.go: {len(missing_references)}")
        
        print("-" * 50)
        print("Resources missing _types.go (Likely TF/DCL based):")
        print("-" * 50)
        for kind, group, filename in missing_types:
            print(f"Kind: {kind}")
            print(f"  Group: {group}")
            print(f"  CRD File: {filename}")
            print()

        print("-" * 50)
        print("Resources missing _reference.go (Has _types.go):")
        print("-" * 50)
        for kind, group, filepaths in missing_references:
            print(f"Kind: {kind}")
            print(f"  Group: {group}")
            print(f"  Type File(s): {', '.join(filepaths)}")
            print()

if __name__ == "__main__":
    main()
