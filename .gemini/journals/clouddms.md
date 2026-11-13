# Cloud DMS Journal

### [2026-11-06] CloudDMSConnectionProfile Greenfield Scaffolding
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `CloudDMSConnectionProfile` under `clouddms.cnrm.cloud.google.com/v1alpha1`.
- **Legacy Identity & Reference Cleanup**: `apis/clouddms/v1alpha1` previously had handwritten `connectionprofile_identity.go` and `connectionprofile_reference.go` files (added for references from `CloudDMSMigrationJob`) which lacked the `clouddms` prefix (flagged in `naming_violations.txt`) and did not implement `IdentityV2`. Removing these allowed `controllerbuilder generate-identity` to deterministically generate standard `IdentityV2` and `refs.Ref` in `clouddmsconnectionprofile_identity.generated.go` and `clouddmsconnectionprofile_reference.generated.go` without redeclaration conflicts, while resolving the naming violations.
- **Empty Oneof Option**: The `oracle.staticServiceIPConnectivity` field corresponds to an empty oneof message in protobuf. Added `clouddmsconnectionprofiles.clouddms.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` under the `// GCP known pattern: Empty oneof option (spec.oracle.staticServiceIPConnectivity)` category.
- **Blocked Kinds**: Added `CloudDMSConnectionProfile` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.

### [2026-11-09] CloudDMSConnectionProfile Judgement Queue Triage & Type Finalization
- **Context**: Triaging judgement queue entries and finalizing direct KRM types for `CloudDMSConnectionProfile`.
- **Parent**: Converted `projectRef` and `location` to inlined `*parent.ProjectAndLocationRef`.
- **Labels & Lifecycle State**: Removed `spec.labels` (reconciled via `metadata.labels`) and moved `spec.state` to `status.observedState.state` since connection profile state is server-managed lifecycle state.
- **Resource References**:
  - `alloydb.settings.encryptionConfig.kmsKeyRef` (*kmsv1beta1.KMSCryptoKeyRef)
  - `alloydb.settings.vpcNetworkRef` (*computerefs.ComputeNetworkRef)
  - `cloudsql.settings.cmekKeyNameRef` (*kmsv1beta1.KMSCryptoKeyRef)
  - `cloudsql.settings.sourceRef` (*CloudDMSConnectionProfileRef)
  - `cloudsql.settings.ipConfig.privateNetworkRef` (*computerefs.ComputeNetworkRef)
  - `oracle.privateConnectivity.privateConnectionRef` (*PrivateConnectionRef)
  - `postgresql.privateServiceConnectConnectivity.serviceAttachmentRef` (*refsv1beta1.ComputeServiceAttachmentRef)
- **Non-References**: `alloydb.clusterID`, `mysql.cloudSQLID`, and `postgresql.cloudSQLID` are plain string identifiers, not full GCP resource references.
- **Sensitive Fields**: Modeled sensitive password fields (.spec.alloydb.settings.initialUser.password, .spec.cloudsql.settings.rootPassword, .spec.mysql.password, .spec.oracle.forwardSSHConnectivity.password, .spec.oracle.password, and .spec.postgresql.password) using `*refsv1beta1secret.Legacy` following the KCC secret ref style, implemented custom proto mappers, and resolved the corresponding `sensitive-field` entries in `judgement_queue.yaml`.

