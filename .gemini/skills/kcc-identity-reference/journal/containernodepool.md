# Journal: ContainerNodePool Identity & Reference Migration

Migrating `ContainerNodePool` to the `identity` and `refs` pattern with `gcpurls.Template`.

## Observations & Learnings

1. **Zonal vs. Regional Variations**:
   - Similar to `ContainerCluster`, `ContainerNodePool` has two primary formats: a regional path (`projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}`) and a zonal path (`projects/{project}/zones/{zone}/clusters/{cluster}/nodePools/{nodePool}`).
   - Rather than storing both regional locations and zonal zones into a single `Location` field and using hyphen-counting string parsing, we split them into separate `Zone` and `Location` fields in the `ContainerNodePoolIdentity` struct (similar to the `ComputeDisk` identity pattern).
   - This keeps the fields matching GKE's templates more cleanly and avoids string parsing in `String()` and `ParentString()`.

2. **Parent Path Reconstitution (`ParentString()`)**:
   - Since regional and zonal paths are cleanly split into `Location` and `Zone` fields respectively, we implement `ParentString()` to construct the parent GKE cluster's GCP URI without any hyphen-counting.
   - We check if the `Zone` field is set to determine the parent structure (e.g. `projects/{project}/zones/{zone}/clusters/{cluster}` vs `projects/{project}/locations/{location}/clusters/{cluster}`).

3. **Status Cross-Check**:
   - `ContainerNodePool` does not contain `status.externalRef` or `status.name` in its current schema. According to KCC's strict guidelines to never modify the API schema during this phase, `GetIdentity` skips status cross-checks and delegates directly to the spec identity resolver.
