As part of moving resources from terraform controllers to direct controllers (Epic #5954), we need to create the Go types for `IAPBrand`.

Currently, `IAPBrand` is managed by the DCL controller (marked with `dcl2crd=true`). The goal is to create the Go types in `apis/iap/v1beta1/` so that we can eventually migrate the controller implementation to the "direct" approach.

### Instructions

1.  **Create a generate.sh**:
    Create `apis/iap/v1beta1/generate.sh` which includes `IAPBrand`.
    It likely maps to something like `google.cloud.iap.v1`.
    Example:
    ```bash
    go run . generate-types \
      --service google.cloud.iap.v1 \
      --api-version iap.cnrm.cloud.google.com/v1beta1 \
      --resource IAPBrand:Brand \
      --include-skipped-output

    go run . generate-mapper \
      --service google.cloud.iap.v1 \
      --api-version iap.cnrm.cloud.google.com/v1beta1 \
      --include-skipped-output
    ```

2.  Set the write permission on the new `apis/iap/v1beta1/generate.sh` file. You should do this by running both `chmod +x apis/iap/v1beta1/generate.sh` and `git add --chmod=+x apis/iap/v1beta1/generate.sh`.

3.  **Generate Scaffolding**:
    Run `apis/iap/v1beta1/generate.sh`. This should create `apis/iap/v1beta1/iapbrand_types.go`.

4.  **Iterate on Types**:
    Compare the generated CRD with the existing one using `dev/tasks/diff-crds`.
    Modify `apis/iap/v1beta1/iapbrand_types.go` until the CRD matches the existing one at `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_iapbrands.iap.cnrm.cloud.google.com.yaml`.

    **Acceptance Criteria:**
    - Running `dev/tasks/diff-crds` should not show differences (or minimal acceptable ones like descriptions).
    - Ensure that running the check_crd_equivalence MCP on the CRD should return EQUIVALENT.
    - Changes to the schema (fields added/removed) are NOT acceptable.

5.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    ```go
    // Copyright 2026 Google LLC
    ```
    Please do not change the copyright on existing files.

6.  **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    ```go
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
    ```
    The goal is to maintain these annotations, not add an annotation if it is missing.

7.  **Status**:
    `status.observedGeneration` should be an `*int64`.

8. **Generate Mappers**:
   - Running `dev/tasks/generate-types-and-mappers` will generate the mapper code once the `apis/iap/v1beta1/iapbrand_types.go` file is generating an equivalent CRD.
   - Run `make all-binary` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

This issue is part of Epic #5954.
