# PrivateCACertificateTemplate KRM Transition Journal

## Learnings & Observations

### 1. Verification of Strict Schema Compatibility
We verified that the direct KRM types defined in `apis/privateca/v1beta1/privatecacertificatetemplate_types.go` are 100% strictly compatible with the existing/baseline Custom Resource Definition (CRD) schema. Running `dev/tasks/diff-crds` returned an empty diff, confirming absolutely no drift, schema omissions, or unexpected additions.

### 2. Service-level Consolidation in `generate.sh`
The generation script for `PrivateCACertificateTemplate` is properly integrated into the service-level consolidated script `apis/privateca/generate.sh`:
- Runs `generate-types` for `PrivateCACertificateTemplate` targeting version `v1beta1`.
- Runs `generate-mapper` once at the end of the script to compile mappers for all registered resources under `apis/privateca`.

This conforms to the design principle of service-level directory structure and avoids split version generation files.

### 3. Fuzzer Design & Validation
The KRM fuzzer in `pkg/controller/direct/privateca/privatecacertificatetemplate_fuzzer.go` was executed and validated. Using:
```bash
FOCUS=PrivateCACertificateTemplate go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
```
All fuzz tests passed cleanly, validating that our spec and status schemas round-trip perfectly with the generated mapper logic.
