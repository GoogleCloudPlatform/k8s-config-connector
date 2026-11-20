# Securityposture Service Notes

- `SecurityposturePosture`:
  - Parent URL format is `organizations/{organization}/locations/{location}/postures/{posture}`. Kept `OrganizationRef` and `Location` in `SecurityposturePostureSpec` as no single shared parent reference covers organization+location.
  - Moved `Etag` from `SecurityposturePostureSpec` to `SecurityposturePostureObservedState`.
  - Moved `CustomConstraint` struct into `securitypostureposture_types.go` and converted `CustomConstraint.Name` (`*string`) to `CustomConstraintRef` (`*orgpolicyv1beta1.CustomConstraintRef`) referencing `OrgPolicyCustomConstraint`.
  - Confirmed `cannedConstraintID`, `policyID`, and `policySetID` are string identifiers/values rather than KCC resource references.
  - In direct controller, `UpdatePostureRequest` requires `RevisionId` (`revision_id`) which must be set from `a.actual.GetRevisionId()`.
  - `SecurityposturePosture` is an organization-level resource (`organizations/{organization}/locations/{location}/postures/{posture}`). Recording E2E fixtures against real GCP requires organization-level IAM permissions (`roles/securityposture.admin` or `securityposture.postures.*`) on the target organization; otherwise GCP returns 403 `IAM_PERMISSION_DENIED` (`securityposture.postures.get` / `securityposture.postures.create`).
  - Per GCP API specifications, state updates cannot be clubbed with updates to description or policySets. Maintaining `state: DRAFT` across `create.yaml` and `update.yaml` avoids this restriction during updates.
