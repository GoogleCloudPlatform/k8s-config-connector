# Journal: ContainerNodePool Identity & Reference Migration

Migrating `ContainerNodePool` to the `identity` and `refs` pattern with `gcpurls.Template`.

## Observations & Learnings

1. **Zonal vs. Regional Variations**:
   - Similar to `ContainerCluster`, `ContainerNodePool` has two primary formats: a regional path (`projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}`) and a zonal path (`projects/{project}/zones/{location}/clusters/{cluster}/nodePools/{nodePool}`).
   - GKE templates parse regional and zonal locations into a single `Location` field on the identity struct.

2. **Parent Path Reconstitution (`ParentString()`)**:
   - Since both zonal and regional paths map to a single `Location` field, we implement `ParentString()` to construct the parent GKE cluster's GCP URI based on the format of the location.
   - We determine whether a location is zonal or regional using a simple, robust hyphen-count check (`strings.Count(Location, "-") >= 2`), which matches standard GCP region vs. zone naming patterns (e.g. `us-central1` vs `us-central1-a`).

3. **Status Cross-Check**:
   - `ContainerNodePool` does not contain `status.externalRef` or `status.name` in its current schema. According to KCC's strict guidelines to never modify the API schema during this phase, `GetIdentity` skips status cross-checks and delegates directly to the spec identity resolver.
