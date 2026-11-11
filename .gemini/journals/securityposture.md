# Securityposture Service Notes

- `SecurityposturePosture`:
  - Parent URL format is `organizations/{organization}/locations/{location}/postures/{posture}`. Kept `OrganizationRef` and `Location` in `SecurityposturePostureSpec` as no single shared parent reference covers organization+location.
  - Moved `Etag` from `SecurityposturePostureSpec` to `SecurityposturePostureObservedState`.
  - Moved `CustomConstraint` struct into `securitypostureposture_types.go` and converted `CustomConstraint.Name` (`*string`) to `CustomConstraintRef` (`*orgpolicyv1beta1.CustomConstraintRef`) referencing `OrgPolicyCustomConstraint`.
  - Confirmed `cannedConstraintID`, `policyID`, and `policySetID` are string identifiers/values rather than KCC resource references.
