# BackupDRBackupPlan Identity & Reference Migration Journal

## Overview
Successfully moved `BackupDRBackupPlan` to the `identity.IdentityV2` and `refs.Ref` modern pattern with `gcpurls.Template`.

## Observations & Learnings
1. **Preserving Backward Compatibility**: The existing direct controller (`pkg/controller/direct/backupdr/backupplan_controller.go`) and other resource types (such as `BackupDRBackupPlanAssociation`) refer to legacy identity helpers like `Parent()`, `ID()`, and reference normalizer methods like `NormalizedExternal(...)`.
   - Rather than refactoring the controller logic and potentially introducing risks of regressions, we implemented the canonical modern interfaces (`identity.IdentityV2`, `refs.Ref`) while keeping legacy methods (`Parent()`, `ID()`, `NormalizedExternal(...)`) on the structs.
   - This keeps the transition 100% safe, fast, and compilation-ready without complex multi-file edits.
2. **Template Variables & Capitalization**:
   - The GCP API path for the resource contains CamelCase-like elements: `backupPlans`.
   - The `gcpurls.Template` template `projects/{project}/locations/{location}/backupPlans/{backupplan}` matches the lowercased struct field `BackupPlan`.
   - This case sensitivity is critical because the GCP URLs are case-sensitive.
3. **Deepcopy Exclusions**:
   - Marking the identity struct with `// +k8s:deepcopy-gen=false` and running `dev/tasks/generate-types-and-mappers` cleanly pruned obsolete deepcopy code from `zz_generated.deepcopy.go`.
