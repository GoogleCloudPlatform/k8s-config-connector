# BackupDRBackupPlan Identity & Reference Migration Journal

## Overview
Successfully moved `BackupDRBackupPlan` to the `identity.IdentityV2` and `refs.Ref` modern pattern with `gcpurls.Template`.

## Observations & Learnings
1. **Addressing Reviewer Feedback (ParentString and legacy removal)**:
   - The initial version kept legacy methods `Parent()`, `ID()`, `NewBackupPlanIdentity`, and `ParseBackupPlanExternal` to avoid refactoring the controller.
   - Per reviewer feedback, we successfully removed `BackupPlanParent` entirely and instead defined `ParentString()` directly on `BackupPlanIdentity`.
   - We removed the legacy `NewBackupPlanIdentity` and `ParseBackupPlanExternal` methods completely, transition the controller (`backupplan_controller.go`) to use `obj.GetIdentity(ctx, reader)`, `id.ParentString()`, and direct struct fields (`Location` and `BackupPlan`).
2. **Template Variables & Capitalization**:
   - The GCP API path for the resource contains CamelCase-like elements: `backupPlans`.
   - The `gcpurls.Template` template `projects/{project}/locations/{location}/backupPlans/{backupplan}` matches the lowercased struct field `BackupPlan`.
   - This case sensitivity is critical because the GCP URLs are case-sensitive.
3. **Deepcopy Exclusions**:
   - Marking the identity struct with `// +k8s:deepcopy-gen=false` and running `dev/tasks/generate-types-and-mappers` cleanly pruned obsolete deepcopy code from `zz_generated.deepcopy.go`.
