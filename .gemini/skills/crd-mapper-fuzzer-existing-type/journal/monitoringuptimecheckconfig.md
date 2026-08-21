# Journal: MonitoringUptimeCheckConfig

## Learnings & Observations

1. **Direct Controller Proto Annotation (`+kcc:spec:proto=...`)**:
   - Including `+kcc:spec:proto=google.monitoring.v3.UptimeCheckConfig` in the KRM types of a DCL-based/TF-based resource causes the generator to attempt to generate direct controller mapper files under `pkg/controller/direct/<service>/`.
   - If the nested structures are not generated or marked unreachable in the API package (e.g. because of hand-coding or skipping), these generated mappers will reference undefined functions, causing compile errors.
   - **Resolution**: Since this is a DCL-reconciled resource (and we are only scaffolding the KRM types), we can simply omit the `+kcc:spec:proto` annotation on the struct to prevent direct mapper generation.

2. **Legacy Password Constraints (`value` / `valueFrom`)**:
   - The password field signature `"value,valueFrom"` matches the legacy KCC sensitive fields rule.
   - However, `scripts/add-validation-to-crds/parse-crds.go` restricts this rule to specific kinds (e.g. `AlloyDBUser`, `ContainerCluster`).
   - **Resolution**: Added `MonitoringUptimeCheckConfig` to the check in `parse-crds.go` so that the OpenAPI `oneOf` constraint blocks are correctly and automatically generated for its password field, resulting in a 100% schema match.

3. **Global Resources**:
   - `MonitoringUptimeCheckConfig` is a global resource (scoped to project, not any specific location). Therefore, we removed the scaffolded `location` field from the Spec struct to match the baseline schema perfectly.
