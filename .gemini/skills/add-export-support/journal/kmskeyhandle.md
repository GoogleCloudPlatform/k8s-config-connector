# Export Support for KMSKeyHandle

## Observations and Learnings

1. **Server-Generated Resources without `spec.resourceID`:**
   - Resources with server-assigned IDs (e.g., `KMSKeyHandle`) do not require `spec.resourceID` on creation.
   - When exporting resources during E2E tests, the CAIS utility `cais.GetCAISIdentities` previously skipped identity resolution for `ServerGeneratedIdentity` types if `spec.resourceID` was empty.
   - To support export testing for these, we updated `pkg/cais/cais.go` to also allow identity resolution if `status.externalRef` is populated, since a reconciled resource already has its GCP-assigned identifier stored in its status.

2. **Ref Scheme registration for Direct GVKs:**
   - `KMSKeyHandle` GVK was not registered in `kccscheme` because `keyhandle_reference.go` only registered the reference type via `refs.Register(&KMSKeyHandleRef{})`.
   - We updated it to `refs.Register(&KMSKeyHandleRef{}, &KMSKeyHandle{})` so that `kccscheme.RegisterType` is invoked for the resource GVK itself. This is critical for CAIS identity lookups to recognize the GVK.

3. **Dynamic UUID Normalization in `_identities.yaml`:**
   - Because MockGCP assigns a dynamic random UUID to the KeyHandle on every test execution, the `caisURL` in `_identities.yaml` would change on every run.
   - We added a normalization rule in `pkg/cais/caistesting/testing.go` within `NormalizeDynamicIDs` to replace `/keyHandles/<uuid>` with the stable placeholder `/keyHandles/${keyHandleID}`. This aligns it perfectly with `_exported.yaml` normalization.
