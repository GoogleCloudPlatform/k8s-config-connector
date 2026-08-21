# BigtableGCPolicy Identity & Reference Journal

## Observations

1. **Logical Sub-resource Structure**:
   `BigtableGCPolicy` is mapped as a sub-resource under a Bigtable Column Family in a Table. In the GCP API, GC policies are not standalone endpoints with distinct URIs. However, to cleanly represent the identity for Config Connector resources, we used a logical structure matching:
   `"projects/{project}/instances/{instance}/tables/{table}/columnFamilies/{columnFamily}"`

2. **No Status Cross-Check**:
   Since the `BigtableGCPolicy` resource schema has no `status.externalRef` or `status.name` fields, we bypassed any status cross-checks in `GetIdentity`. It returns the spec-derived identity directly.

3. **CAI Validation Exemption**:
   As `BigtableGCPolicy` is not a top-level resource indexed by Cloud Asset Inventory (CAI), we registered its normalized template in `pkg/gcpurls/registry_test.go` under `ignoredTemplates` to prevent the `TestRegisteredTemplatesMatchCAI` validation test from failing.
