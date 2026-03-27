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

def build_dependency_graph(crds_dir="config/crds/resources"):
    if not os.path.exists(crds_dir):
        print(f"Error: {crds_dir} not found.", file=sys.stderr)
        sys.exit(1)

    known_kinds = {}
    dependencies = defaultdict(set)
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

    return dependencies, known_kinds

def topological_sort(target_kinds, dependencies):
    nodes = set(target_kinds)
    local_graph = {node: set() for node in nodes}
    for node in nodes:
        if node in dependencies:
            # A node depends on its dependencies (dependencies must come first)
            local_graph[node] = {dep for dep in dependencies[node] if dep in nodes}

    # Track nodes without any dependencies initially
    no_dependencies = [node for node in nodes if len(local_graph[node]) == 0]

    # in_degree represents how many unfulfilled dependencies a node has
    in_degree = {node: len(local_graph[node]) for node in nodes}
    
    # rev_graph maps from a dependency to the nodes that depend on it
    rev_graph = {node: set() for node in nodes}
    for node in nodes:
        for dep in local_graph[node]:
            rev_graph[dep].add(node)

    # queue contains nodes with no unfulfilled dependencies
    queue = [node for node in nodes if in_degree[node] == 0]
    queue.sort()
    
    result = []
    
    while queue:
        current = queue.pop(0)
        result.append(current)
        
        # When 'current' is processed, fulfill the dependency for nodes waiting on it
        for dependent in sorted(list(rev_graph[current])):
            in_degree[dependent] -= 1
            if in_degree[dependent] == 0:
                queue.append(dependent)
                queue.sort()

    cycle_nodes = []
    if len(result) != len(nodes):
        # Nodes with in_degree > 0 are part of a cycle or depend on a cycle
        remaining = [node for node in nodes if in_degree[node] > 0]
        cycle_nodes = remaining
        result.extend(sorted(remaining))
        
    return result, cycle_nodes, no_dependencies

def main():
    if len(sys.argv) < 2:
        print("Usage: python hack/directmigration/topological-sort.py <Kind1> [<Kind2> ...] OR python hack/directmigration/topological-sort.py <file_path>")
        print("Example: python hack/directmigration/topological-sort.py ComputeInstance ComputeNetwork ComputeSubnetwork")
        print("Example: python hack/directmigration/topological-sort.py missing-kinds.txt")
        sys.exit(1)

    input_args = sys.argv[1:]
    target_kinds = []
    annotations = {}

    def add_kind(part):
        part = part.strip()
        if not part:
            return
        if ':' in part:
            kind, annotation = part.split(':', 1)
            target_kinds.append(kind)
            annotations[kind] = annotation
        else:
            target_kinds.append(part)

    # If the first argument is a file, read from it
    if len(input_args) == 1 and os.path.isfile(input_args[0]):
        with open(input_args[0], 'r') as f:
            for line in f:
                # Support both comma-separated and one-per-line
                parts = line.replace(',', ' ').split()
                for p in parts:
                    add_kind(p)
    else:
        # Handle comma-separated list if passed as strings (e.g., "Kind1, Kind2")
        for arg in input_args:
            parts = arg.replace(',', ' ').split()
            for p in parts:
                add_kind(p)

    if not target_kinds:
        print("Error: No kinds provided.", file=sys.stderr)
        sys.exit(1)

    dependencies, known_kinds = build_dependency_graph()

    # Validate inputs against known CRDs
    missing_kinds = [k for k in target_kinds if k not in known_kinds]
    if missing_kinds:
        print(f"Warning: The following kinds are not found in CRDs and will be treated as independent: {', '.join(missing_kinds)}\n", file=sys.stderr)

    sorted_kinds, cycle_nodes, no_dependencies = topological_sort(target_kinds, dependencies)

    print("=== Resources with NO dependencies ===")
    if no_dependencies:
        for kind in sorted(no_dependencies):
            suffix = ""
            if kind in annotations:
                if annotations[kind] == "missing_reference":
                    suffix = " (missing _reference.go only)"
                else:
                    suffix = f" ({annotations[kind]})"
            print(f"- {kind}{suffix}")
    else:
        print("(None)")
    print()

    if cycle_nodes:
        print("=== Warning: Dependency Cycle Detected! ===")
        print("The following resources are involved in a cycle or depend on one:")
        for kind in sorted(cycle_nodes):
            deps = [d for d in dependencies.get(kind, []) if d in cycle_nodes]
            print(f"- {kind} depends on: {', '.join(deps) if deps else '(unknown)'}")
        print()

    print("=== Topologically Sorted Resources ===")
    for kind in sorted_kinds:
        suffix = ""
        if kind in annotations:
            if annotations[kind] == "missing_reference":
                suffix = " (missing _reference.go only)"
            else:
                suffix = f" ({annotations[kind]})"
        print(f"{kind}{suffix}")

if __name__ == "__main__":
    main()