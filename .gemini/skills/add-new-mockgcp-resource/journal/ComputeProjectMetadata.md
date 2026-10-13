# ComputeProjectMetadata MockGCP Alignment Journal

- **Date:** October 13, 2026
- **Resource Kind:** `ComputeProjectMetadata`
- **GCP Resource:** `compute, projects`

## Overview

A mock implementation for GCP compute `projects` resource (`ProjectsServer`) did not exist under `mockgcp/mockcompute`.
`ComputeProjectMetadata` manages project-wide common instance metadata via the Compute v1 API:
- `Get`: `GET /compute/v1/projects/{project}`
- `SetCommonInstanceMetadata`: `POST /compute/v1/projects/{project}/setCommonInstanceMetadata`

## Discrepancies and Findings

1. **Missing Projects Service:** `mockcompute` lacked implementation and registration of `ProjectsServer` and `ProjectsHandler`.
2. **Terraform Request Formatting:** The legacy Terraform provider's `expandComputeMetadata` helper allocates a slice of `MetadataItems` with `make([]*compute.MetadataItems, len(m))` before appending items, causing the JSON payload sent to `setCommonInstanceMetadata` to include `null` elements in `items` (e.g., `{"items": [null, null, {"key": "foo", "value": "bar"}]}`). In MockGCP, grpc-gateway's protojson unmarshaler fails with `proto: syntax error: unexpected token null` when encountering null values in repeated message fields. In `mockcompute/service.go` HTTP mux, we strip null elements from the `items` array for `/setCommonInstanceMetadata` requests before handing them off to grpc-gateway.
3. **Target Link Normalization:** In `mockcompute/normalize.go`, LRO operations with `targetLink` matching `https://compute.googleapis.com/compute/v1/projects/{project}` identified the kind as `projects` and attempted to replace `targetId` (the numeric project number) with `${projectID}`, causing duplicate string replacement conflicts with `${projectNumber}`. Normalization was updated to map `projects` target IDs to `${projectNumber}` and names to `${projectId}`.
4. **Real GCP Recording Status:** In this test environment, running `hack/record-gcp` against real GCP failed because the active sandbox service account (`cnrm-barni-1.svc.id.goog`) lacks access to `cnrm-barni-1` or any GCP project (`ERROR: (gcloud.projects.describe) [cnrm-barni-1.svc.id.goog] does not have permission to access projects instance [cnrm-barni-1]`). Therefore, real GCP HTTP logs could not be recorded in this session. `_http_mock.log` and golden objects were generated via `hack/compare-mock`.

## Solution

1. Implemented `ProjectsV1` in `mockgcp/mockcompute/projectsv1.go` providing `Get` and `SetCommonInstanceMetadata`.
2. Registered `ProjectsServer` and `ProjectsHandler` in `mockgcp/mockcompute/service.go`.
3. Added null-filtering in HTTP request body preprocessing for `setCommonInstanceMetadata` in `mockgcp/mockcompute/service.go`.
4. Updated `mockgcp/mockcompute/normalize.go` to handle `projects` resource URL kind mappings without duplicate replacement errors.
5. Enabled `ComputeProjectMetadata` in `config/tests/samples/create/harness.go`.
6. Removed `ComputeProjectMetadata` from ratcheting exclusion in `tests/e2e/ratcheting.go`.
7. Added unit tests in `mockgcp/mockcompute/projectsv1_test.go`.

## Validation

- Unit tests pass: `go test -v ./mockgcp/mockcompute -run TestProjectMetadataLifecycle`
- Golden fixture mock tests pass: `hack/compare-mock "fixtures/computeprojectmetadata"`
