# BeyondCorpAppConnector Identity and Reference Journal

- **Resource**: `BeyondCorpAppConnector`
- **Group**: `beyondcorp.cnrm.cloud.google.com`
- **Version**: `v1alpha1`
- **Kind**: `BeyondCorpAppConnector`
- **Controller Type**: Legacy (Terraform-based).

## Observations and Learnings

1. **Deterministic Generation**:
   Added `${CONTROLLERBUILDER} generate-identity` to `apis/beyondcorp/generate.sh` for `BeyondCorpAppConnector:AppConnector` with proto service `google.cloud.beyondcorp.appconnectors.v1`.
   This deterministically generates:
   - `beyondcorpappconnector_identity.generated.go`
   - `beyondcorpappconnector_reference.generated.go`
   - `beyondcorpappconnector_identity_generated_test.go` (conformance testing)

2. **Declaration-by-Declaration Override ("Hand-written wins")**:
   `BeyondCorpAppConnector` is currently managed by the Terraform controller and does not yet have `status.externalRef` in its CRD status schema. Therefore, per KCC identity standards, `Normalize` must use `refs.NormalizeWithFallback` rather than delegating directly to `refs.Normalize`.
   We implemented `Normalize` in `beyondcorpappconnector_reference.go`. The generator cleanly detects `Normalize` and suppresses emitting `Normalize` in `beyondcorpappconnector_reference.generated.go`.

3. **Fallback Readiness Check**:
   The fallback function in `NormalizeWithFallback` inspects `obj.Status.Conditions` for `Ready` == `"True"`. If not ready, it returns `""` which causes `refs.NormalizeWithFallback` to return `k8s.NewReferenceNotReadyError`. Only when ready does it extract the identity via `getIdentityFromBeyondCorpAppConnectorSpec`.

4. **Schema Parity**:
   Ran `dev/tasks/diff-crds` to ensure zero schema changes were introduced.
