# Journal: CertificateManagerCertificateMap Transition to Direct KRM Types

## Observations

1. **Schema Integrity and Versioning:**
   - Prior to transition, the `CertificateManagerCertificateMap` resource was defined in both `v1alpha1` and `v1beta1`, with the `v1alpha1` version serving as the storage version (as indicated by the Terraform service mapping metadata).
   - Upon transition to the direct controller KRM types, the `v1beta1` version has been promoted to be the storage version by adding the `// +kubebuilder:storageversion` annotation on its struct definition.
   - Switched ProjectRef to use `refs.ProjectRef` to avoid introducing any `kind` field.
   - Removed `externalRef` and `observedState` under status to maintain 100% strict schema compatibility with the baseline CRD, keeping flat status fields and achieving completely clean `dev/tasks/diff-crds` output.

2. **Acronym Check & Field Renaming:**
   - The field `description` was manually added to `CertificateManagerCertificateMapSpec` because the baseline `v1beta1` CRD contains it, ensuring strict schema parity (as verified by `dev/tasks/diff-crds`).
   - Mapped `.status.gclbTargets[].targetHttpsProxy` and `.status.gclbTargets[].targetSslProxy` acronym exceptions in `testdata/exceptions/acronyms.txt` because the baseline CRD properties had lowercase acronym casing.

3. **Round-trip Fuzzer:**
   - We created `certificatemap_fuzzer.go` under `pkg/controller/direct/certificatemanager/`.
   - It registers a spec-only round-trip fuzzer using the generated `v1beta1` FromProto/ToProto conversion methods for `Spec` with `NewKRMTypedSpecFuzzer`.
