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
import argparse

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
                        if kind.endswith("Spec") or kind.endswith("Status") or kind.endswith("List") or kind.endswith("ObservedState"):
                            continue
                            
                        if kind not in implemented_kinds:
                            implemented_kinds[kind] = []
                        implemented_kinds[kind].append(filepath)
    
    return implemented_kinds

def get_resource_classifications(config_path="pkg/controller/resourceconfig/static_config.go"):
    """
    Parses static_config.go to classify resources as 'greenfield', 'migration', or 'completed_migration'.
    """
    classifications = {}
    if not os.path.exists(config_path):
        print(f"Warning: {config_path} not found. Cannot classify resources.", file=sys.stderr)
        return classifications
    
    with open(config_path, "r", encoding='utf-8') as f:
        for line in f:
            match = re.search(r'Kind:\s*"([^"]+)".*?SupportedControllers:\s*\[\]k8s\.ReconcilerType\{(.*?)\}', line)
            if match:
                kind = match.group(1)
                controllers = match.group(2)
                has_tf = "k8s.ReconcilerTypeTerraform" in controllers
                has_dcl = "k8s.ReconcilerTypeDCL" in controllers
                has_direct = "k8s.ReconcilerTypeDirect" in controllers
                
                if (has_tf or has_dcl) and not has_direct:
                    classifications[kind] = "migration"
                elif has_direct and not has_tf and not has_dcl:
                    classifications[kind] = "greenfield"
                elif has_direct and (has_tf or has_dcl):
                    classifications[kind] = "completed_migration"
                else:
                    classifications[kind] = "other"
    return classifications

def main():
    parser = argparse.ArgumentParser(description="Find migration resources with _types.go that are missing _identity.go or _reference.go")
    parser.add_argument("-o", "--output", help="Output file to write missing Kind names to (one per line).")
    args = parser.parse_args()

    crd_resources = get_crd_resources()
    implemented_types = get_implemented_types()
    classifications = get_resource_classifications()

    missing_files = []
    fully_implemented = []

    for kind, group, filename in sorted(crd_resources, key=lambda x: x[0]):
        classification = classifications.get(kind, "unknown")
        
        # Only process resources identified as "migration" (supports TF/DCL, does NOT yet support Direct)
        if classification != "migration":
            continue

        has_types = kind in implemented_types
        
        # Skip resources where types haven't been implemented yet (they are Step 1)
        if not has_types:
            continue

        has_reference = False
        has_identity = False
        filepaths = implemented_types[kind]
        
        for filepath in filepaths:
            dir_path = os.path.dirname(filepath)
            
            # Pattern 1: Derived from types filename
            ref_pattern1 = filepath.replace("_types.go", "_reference.go")
            ident_pattern1 = filepath.replace("_types.go", "_identity.go")
            
            # Pattern 2: Derived from Kind name
            ref_pattern2 = os.path.join(dir_path, f"{kind.lower()}_reference.go")
            ident_pattern2 = os.path.join(dir_path, f"{kind.lower()}_identity.go")

            if os.path.exists(ref_pattern1) or os.path.exists(ref_pattern2):
                has_reference = True
            if os.path.exists(ident_pattern1) or os.path.exists(ident_pattern2):
                has_identity = True
        
        if has_reference and has_identity:
            fully_implemented.append((kind, group))
        else:
            missing_files.append((kind, group, filepaths, has_identity, has_reference))

    if args.output:
        with open(args.output, 'w') as f:
            for kind, _, _, _, _ in missing_files:
                f.write(f"{kind}:missing_identity_or_reference\n")
        print(f"Successfully wrote {len(missing_files)} kinds to {args.output}")
    else:
        print(f"Resources fully implemented (_types.go, _identity.go, _reference.go): {len(fully_implemented)}")
        print(f"Migration resources WITH _types.go but MISSING _identity.go or _reference.go: {len(missing_files)}")
        
        print("-" * 50)
        print("Missing identity or reference files:")
        print("-" * 50)
        for kind, group, filepaths, has_ident, has_ref in missing_files:
            missing_parts = []
            if not has_ident: missing_parts.append("_identity.go")
            if not has_ref: missing_parts.append("_reference.go")
            print(f"Kind: {kind}")
            print(f"  Group: {group}")
            print(f"  Missing: {', '.join(missing_parts)}")
            print(f"  Type File(s): {', '.join(filepaths)}")
            print()

if __name__ == "__main__":
    main()
