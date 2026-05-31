# GKEHubMembershipBinding Identity and Reference Migration

## Observations

- **GKEHubMembershipBinding** belongs to the `v1alpha1` API version under the `gkehub` group.
- The parent resource is `GKEHubMembership`. The reference parsing utilizes the `krmv1beta1` resolver helper `ResolveGKEHubMembershipRef` which retrieves the parent's coordinates (Project, Location, MembershipID).
- The identity struct was upgraded to `GKEHubMembershipBindingIdentity` with the fields:
  - `Project`
  - `Location`
  - `Membership`
  - `MembershipBinding`
- Removed the deprecated/legacy `.Parent()` and `.ID()` helper methods from the identity struct. The mapping logic in the mapper and controller now explicitly maps the fields or constructs parent identity using `krmv1beta1.NewGKEHubMembershipIdentity(Project, Location, Membership)` directly, which keeps the identity definition clean and aligned with the canonical pattern.
- Successfully created `membershipbinding_reference.go` declaring the `GKEHubMembershipBindingRef` struct and implementing the `refs.Ref` interface along with the corresponding `ResolveGKEHubMembershipBindingRef` helper function.
- All unit and registry tests passed successfully.
