# GKEMulticloud Journal

### [2026-06-15] Template Mapping with Underscores in IdentityV2
- **Context**: Implementing IdentityV2 for `GKEMulticloudAttachedCluster` using `gcpurls.Template`.
- **Problem**: The GCP resource URL template has `{attached_cluster}`, but using `AttachedCluster string` in the identity struct caused a panic at startup: `panic: field "attached_cluster" not found in struct v1alpha1.GKEMulticloudAttachedClusterIdentity`.
- **Solution**: The `gcpurls.Template` parser matches field names case-insensitively but literally. Field names with camel-casing (like `AttachedCluster`) convert to lowercase `attachedcluster`, which does not match the lowercase snake_case placeholder `attached_cluster`. To solve this, the struct field must be named `Attached_cluster string` (retaining the underscore so that converting to lowercase matches `"attached_cluster"` literally).
- **Impact**: Any direct controller identity struct that maps a template placeholder with an underscore must name its field with the same underscore (e.g. `Attached_cluster`) to avoid panics.

### [2026-06-15] Unreachable Types Scaffolding Pattern
- **Context**: Scaffolding GKEMulticloudAttachedCluster direct types.
- **Problem**: Running `generate.sh` initially commented out nested types in `types.generated.go` because the initial scaffold Spec and Status fields did not reference them.
- **Solution**: To make the generator uncomment and compile these types, they should be referenced explicitly in `GKEMulticloudAttachedClusterSpec` and `GKEMulticloudAttachedClusterObservedState`. For custom reference fields (such as `Fleet.Project`), define them manually in `types.go` (and set `// +kcc:proto=...`) so that the generator bypasses generating the default proto string mapping.
- **Impact**: Speeds up scaffolding and avoids compiling errors by maintaining the correct reachable types set.
