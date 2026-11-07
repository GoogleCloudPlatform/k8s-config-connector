# Cloud DMS Journal

### [2026-11-06] CloudDMSConnectionProfile Greenfield Scaffolding
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `CloudDMSConnectionProfile` under `clouddms.cnrm.cloud.google.com/v1alpha1`.
- **Legacy Identity & Reference Cleanup**: `apis/clouddms/v1alpha1` previously had handwritten `connectionprofile_identity.go` and `connectionprofile_reference.go` files (added for references from `CloudDMSMigrationJob`) which lacked the `clouddms` prefix (flagged in `naming_violations.txt`) and did not implement `IdentityV2`. Removing these allowed `controllerbuilder generate-identity` to deterministically generate standard `IdentityV2` and `refs.Ref` in `clouddmsconnectionprofile_identity.generated.go` and `clouddmsconnectionprofile_reference.generated.go` without redeclaration conflicts, while resolving the naming violations.
- **Empty Oneof Option**: The `oracle.staticServiceIPConnectivity` field corresponds to an empty oneof message in protobuf. Added `clouddmsconnectionprofiles.clouddms.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` under the `// GCP known pattern: Empty oneof option (spec.oracle.staticServiceIPConnectivity)` category.
- **Blocked Kinds**: Added `CloudDMSConnectionProfile` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
