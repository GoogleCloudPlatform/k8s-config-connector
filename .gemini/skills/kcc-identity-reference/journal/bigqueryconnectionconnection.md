# BigQueryConnectionConnection Identity and Reference Migration

*   **Custom Helpers**: Preserved `ResolveServiceAccountID` and `BigQueryConnectionServiceAccountRef` in `bigqueryconnectionconnection_reference.go` as they are used by IAM resources.
*   **Version Differences**: `v1alpha1` implementation included UUID validation for `spec.resourceID`, which was preserved in the new `IdentityV2` implementation. `v1beta1` did not have this validation.
*   **Template Format**: The identity format is `projects/{project}/locations/{location}/connections/{connection}`.
*   **Registry Exception**: Added the URL format to `pkg/gcpurls/registry_test.go` as it was not present in CAI metadata.
