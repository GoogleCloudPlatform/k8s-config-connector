# Direct Migration Hack Scripts

This directory contains utility scripts to help analyze and implement the direct migration.

## Use Cases

### Generating a Migration Roadmap

When migrating legacy (Terraform or DCL-backed) resources to the "Direct" controller approach, it is often helpful to work in an order that respects resource dependencies (e.g., migrating `ComputeNetwork` before `ComputeSubnetwork`).

You can generate a prioritized list of resources to migrate using the following workflow:

```bash
# 1. Identify all resources missing a Direct implementation and save to a file
./hack/directmigration/find-missing-types.py --output missing-kinds.txt

# 2. Sort the missing resources topologically to determine the implementation order
./hack/directmigration/topological-sort.py missing-kinds.txt
```

This will produce a list where independent resources appear first, followed by resources that depend on them, helping ensure that foundational resources are migrated before their dependents.

### Auditing Resource References

Before migrating a resource, it is important to understand which other resources it references. You can use `dependency-graph.py` to see a recursive tree of dependencies for a specific Kind.

```bash
# View the recursive dependency tree for a specific resource
./hack/directmigration/dependency-graph.py ComputeInstance
```

If you want to verify if the reference fields in the CRD are already implemented in Go types, use `find-implemented-refs.py`.

```bash
# Check which reference fields in CRDs have corresponding Go types
./hack/directmigration/find-implemented-refs.py
```

### Visualizing Dependencies for a Subset of Kinds

If you are working on a specific set of resources (e.g., a specific service like `Storage`), you can use `topological-sort.py` to understand their internal ordering.

```bash
# Sort a specific list of kinds
./hack/directmigration/topological-sort.py StorageBucket StorageBucketAccessControl StorageNotification
```

## Python Scripts

### Resource Dependencies & Ordering

#### `dependency-graph.py`
Analyzes the `config/crds/resources/` directory to build a dependency graph between KCC resources. A resource is considered dependent on another if its `spec` contains a reference field (e.g., `networkRef` implies a dependency on `ComputeNetwork`).
*   **Usage:** `./hack/directmigration/dependency-graph.py [Kind]`
    *   If a `Kind` is provided, it prints a recursive dependency tree for that specific resource.
    *   If no argument is provided, it prints all known resource dependencies across the entire project.

#### `topological-sort.py`
Takes a list of KCC Resource Kinds and outputs them in topologically sorted order based on their dependencies. Resources without dependencies are printed first, followed by resources that depend on them.
*   **Usage:** 
    *   With arguments: `./hack/directmigration/topological-sort.py ComputeInstance ComputeNetwork ComputeSubnetwork`
    *   With a file: `./hack/directmigration/topological-sort.py missing-kinds.txt`
*   **Features:**
    *   Outputs resources with **no dependencies** clearly.
    *   Detects and warns about **dependency cycles** (e.g., `ComputeDisk` <-> `ComputeImage`).
    *   Accepts input as space-separated arguments, comma-separated arguments, or a file containing kinds on separate lines.

### Codebase Analysis

#### `find-missing-types.py`
Compares the list of CustomResourceDefinitions (CRDs) in `config/crds/resources/` against the Go types defined in the `apis/` directory. It helps identify legacy "TF/DCL-based" controllers that haven't been migrated to the newer "Direct" controller approach.
*   **Usage:** `./hack/directmigration/find-missing-types.py [--output missing-kinds.txt]`
    *   Prints a summary and lists all CRDs that do not have a corresponding `_types.go` file.
    *   The `--output` (`-o`) flag saves the missing kinds to a file, which can be directly fed into `topological-sort.py`.

#### `find-implemented-refs.py`
Analyzes reference fields (`*Ref`) defined in the CRDs' OpenAPI schemas and checks if there is a corresponding Go type implementation in the `apis/` directory.
*   **Usage:** `./hack/directmigration/find-implemented-refs.py`

#### `check-resource-refs.py`
A more comprehensive script that parses `*Ref` fields from schemas, finds Go implementations, and cross-references them against Terraform reference mappings (`config/servicemappings/*.yaml`). Useful for auditing whether KCC references correctly map to the underlying Terraform types.
*   **Usage:** `./hack/directmigration/check-resource-refs.py`
