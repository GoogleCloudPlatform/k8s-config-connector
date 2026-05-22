# GKEHubFeatureMembership Identity and Reference

## Observations

- `GKEHubFeatureMembership` is a "join" resource in KCC that doesn't correspond to a top-level GCP resource but rather a specific configuration within a `GKEHubFeature`.
- The identity follows the CAI format: `projects/{project}/locations/{location}/memberships/{membership}/features/{feature}`.
- Unlike most resources, `GKEHubFeatureMembership` does not have a `spec.resourceID` field. Its identity is fully determined by its `spec.membershipRef` and `spec.featureRef`.
- `GetIdentity` needs to resolve both references to construct the full GCP identifier.
- The controller was updated to use this new identity, simplifying the `AdapterForObject` logic and ensuring `status.externalRef` is set.

## Implementation Details

- Created `gkehubfeaturemembership_identity.go` and `gkehubfeaturemembership_reference.go`.
- Added `ExternalRef` to `GKEHubFeatureMembershipStatus`.
- Updated `pkg/controller/direct/gkehub/featuremembership_controller.go` to use `GetIdentity()`.
- Removed `pkg/controller/direct/gkehub/references.go` as it was no longer needed.
- Added unit tests for the identity parsing.
