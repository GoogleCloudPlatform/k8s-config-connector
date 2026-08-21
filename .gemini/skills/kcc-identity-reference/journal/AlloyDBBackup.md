### [2026-05-29] AlloyDBBackup Identity & Refs
- **Context**: Moving AlloyDBBackup to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. AlloyDBBackup was using the older `identity.Identity` with regexp parsing.
    2. `backup_reference.go` contained a copy-pasted `BackupAccessRef` struct that was unused and had comments referencing "AccessLevel resource" and fields as pointers.
- **Solution**:
    1. Reimplemented `apis/alloydb/v1beta1/backup_identity.go` to use `identity.IdentityV2` backed by `gcpurls.Template`.
    2. Completely replaced `BackupAccessRef` with a standard `AlloyDBBackupRef` struct implementing `refs.Ref` and registered with `refs.Register`.
    3. Rewrote unit tests in `backup_identity_test.go` to verify standard template-based URL parsing and error cases.
- **Impact**: AlloyDBBackup matches the modern direct controller identity and references conventions. Remnants of copy-pasted unused structures are removed.
