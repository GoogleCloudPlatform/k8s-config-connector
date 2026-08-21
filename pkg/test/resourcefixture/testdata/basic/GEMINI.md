# Basic Test Fixtures

This directory contains the "basic" test fixtures for Config Connector. These are golden tests that verify the end-to-end reconciliation of GCP resources.

## Directory Structure

Tests are organized by GCP service, version, and resource kind:
`pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<test-name>/`

Example: `pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket/storagebucketbasic/`

## Relationship to CRDs and Controllers

Each fixture directory (e.g., `.../storage/v1beta1/storagebucket/`) directly corresponds to a CRD found in `config/crds/resources`.

- **Coverage**: Every CRD should have at least one test fixture in this directory (the "basic" test) to verify core CRUD operations.
- **Controller Validation**: These tests are the primary way to verify the underlying controller for a CRD, whether it is **Direct**, **Terraform (TF)**, or **DCL**.
- **Regressions**: When migrating a CRD from TF/DCL to **Direct**, these fixtures ensure that the behavior and the resulting GCP resources remain identical.
- **Status Verification**: The `_generated_object_<test-name>.golden.yaml` file captures the `status` field, ensuring that `Ready` conditions and other status fields are correctly populated.

## Test Files

Each test directory contains:

- `create.yaml` (Required): Initial Kubernetes manifest for the primary resource.
- `update.yaml` (Optional): Modified manifest to test resource updates.
- `dependencies.yaml` (Optional): Additional Kubernetes resources that must exist before the primary resource.
- `_http.log`: Golden file containing the recorded HTTP/gRPC traffic between Config Connector and GCP APIs. This can be recorded from real GCP (standard) or MockGCP (for mock-only resources or when freezing mock behavior).
- `_generated_object_<test-name>.golden.yaml`: Golden file containing the final state of the Kubernetes object.

## Running Tests

The primary test runner is `TestAllInSeries` in `tests/e2e/unified_test.go`.

### Using Hack Scripts (Recommended)

The `hack` scripts streamline recording and comparing test output:

```bash
# Record against real GCP (updates _http.log and object golden files)
# Provide the full path starting with pkg/
hack/record-gcp pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket/storagebucketbasic

# Verify against MockGCP
# Provide the full path starting with pkg/
hack/compare-mock pkg/test/resourcefixture/testdata/basic/storage/v1beta1/storagebucket/storagebucketbasic
```

### Manual Usage

```bash
# Run a specific fixture against MockGCP
RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/storagebucketbasic

# Run all fixtures for a service (matches any fixture with "storage" in its name)
RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/storage

# Record against real GCP manually (not recommended over hack/record-gcp)
RUN_E2E=1 E2E_GCP_TARGET=real E2E_KUBE_TARGET=envtest WRITE_GOLDEN_OUTPUT=1 GOLDEN_OBJECT_CHECKS=1 GOLDEN_REQUEST_CHECKS=1 go test -v ./tests/e2e -run TestAllInSeries/fixtures/storagebucketbasic
```

## Creating a New Test

1. **Scaffold**: Create the directory structure `.../basic/<service>/<version>/<kind>/<test-name>/`.
2. **Manifests**: Add `create.yaml` and optionally `dependencies.yaml` or `update.yaml`.
   - Ensure resources follow the order of their dependencies (dependencies first).
3. **Record**: Run `hack/record-gcp <path>` to generate golden files. This requires real GCP credentials and an active project.
4. **Mock**: Run `hack/compare-mock <path>` to ensure the test passes against MockGCP. If it fails, you may need to improve the MockGCP implementation.

## Best Practices

- **Uniqueness**: The test framework replaces `${uniqueId}` with a unique ID in your YAML files. It does *not* automatically append a suffix to resource names. You should explicitly use `${uniqueId}` in your resource names (e.g., `name: storagebucket-${uniqueId}`) to avoid collisions when running against real GCP.
- **Normalization**: If tests are flaky due to volatile fields (e.g., server-generated IDs or timestamps), add normalization rules in the corresponding MockGCP service directory (e.g., `mockgcp/mockstorage/normalize.go`). Use `Previsit` or `ConfigureVisitor` to define replacement rules, ensuring they are scoped to the correct service URL.
- **Dependencies**: Keep `dependencies.yaml` minimal. Only include resources directly required by the primary resource. Resources in `dependencies.yaml` should follow the order of their dependencies.
