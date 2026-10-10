---
name: promote-resource-to-beta
description: Guides through qualifying alpha resources and promoting a Config Connector resource from v1alpha1 to v1beta1 across APIs, controllers, and test fixtures.
---

# Promote Resource to Beta

This skill guides an automated agent through qualifying an alpha resource and promoting a Config Connector resource from `v1alpha1` to `v1beta1`.

## Phase 1: Qualify an Alpha Resource for Beta Promotion

Before promoting a resource to beta, ensure it has full API test coverage.

1. **Check Candidates**:
   - Check `experiments/promoter/results/candidates.json` to see if `apiCoverage` is `false`.
2. **Add Full Test Suite**:
   - Create comprehensive `create.yaml` and `update.yaml` test fixtures covering all spec fields.
   - Run tests against mock or real GCP to generate and verify golden HTTP logs (`_http.log`).
3. **Update API Coverage**:
   - Set `"apiCoverage": true` in `experiments/promoter/results/candidates.json`.

---

## Phase 2: Promote the API

1. **Create `v1beta1` Directory**:
   - Create `apis/<service>/v1beta1/`.
   - Copy `generate.sh` from `v1alpha1/` to `v1beta1/` (or copy from `apis/documentai/v1beta1/generate.sh` as template).
   - In `v1beta1/generate.sh`, update `--api-version` to `v1beta1` and specify the promoted resource kind.

2. **Copy and Update Core API Files**:
   - Copy `_types.go` (and any hand-written `_identity.go` / `_reference.go` overrides) to `v1beta1/`. If the resource uses `controllerbuilder generate-identity` in `apis/<service>/generate.sh`, add a `generate-identity --api-version <service>.cnrm.cloud.google.com/v1beta1 --resource <Kind>:<Proto>` invocation so `<kind>_identity.generated.go`, `<kind>_reference.generated.go`, and `<kind>_identity_generated_test.go` are regenerated in `v1beta1/`.
   - Update package name to `package v1beta1`.
   - Update import paths within the files if needed.

3. **Generate and Validate Code**:
   - Run `./generate.sh` in the `v1beta1` directory to produce `types.generated.go`.
   - Run `make generate-crds` from the root of the repository.

4. **Annotations & Backward Compatibility**:
   - On the main resource struct in `v1beta1/<kind>_types.go`, add:
     - `// +kubebuilder:storageversion`
     - `// +kubebuilder:metadata:labels: "internal.cloud.google.com/additional-versions=v1alpha1"`
   - Remove the old `_types.go`, `_identity.go`, and `_reference.go` from `v1alpha1/` if the resource is fully moved to `v1beta1`.

5. **Resolving Cross-Version Dependencies**:
   - If a `v1beta1` resource references another resource still in `v1alpha1`, import the `v1alpha1` package with alias `krmv1alpha1`:
     ```go
     import krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/<service>/v1alpha1"
     ```
   - If an existing `v1alpha1` resource references the promoted `v1beta1` resource, add an aliased import for `v1beta1` in the `v1alpha1` file.

---

## Phase 3: Promote the Controller

1. **Import Alias Convention**:
   - When a controller file needs types from both versions, keep `krm` as the alias for `v1alpha1` and use `krmv1beta1` for `v1beta1`:
     ```go
     import (
         krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/<service>/v1alpha1"
         krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/<service>/v1beta1"
     )
     ```
   - If the controller only references the newly promoted `v1beta1` types, update the import path directly to `apis/<service>/v1beta1` with `krm`.

2. **Update Type Usages**:
   - Update references to promoted types to use `krmv1beta1.` or `krm.`.
   - Update mappers (`mapper.go`) and fuzzers (`*_fuzzer.go`) accordingly.

3. **Verify Build**:
   - Run `go test ./pkg/controller/direct/<service>/...` or `go build ./...` from the controller package.

---

## Phase 4: Promote Test Fixtures

1. **Move Test Directories**:
   - Move fixture directories from `pkg/test/resourcefixture/testdata/basic/<service>/v1alpha1/<kind>/` to `.../v1beta1/<kind>/`.

2. **Update YAML Fixtures**:
   - In `create.yaml` and `update.yaml`, update `apiVersion` to `<service>.cnrm.cloud.google.com/v1beta1`.
   - Keep dependency resource apiVersions in `dependencies.yaml` unchanged unless the dependency itself was also promoted.

3. **Update Dependent Test Fixtures**:
   - Search across `pkg/test/resourcefixture/testdata/basic/` for any other `dependencies.yaml` files referencing the promoted resource and update their `apiVersion` to `v1beta1`.

4. **Verify with Mock Tests**:
   - Run `WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<service>/v1beta1/<kind>/<testname>/`.
   - Verify golden files and HTTP traffic logs.
