# AppEngineStandardAppVersion MockGCP Alignment Journal

- **Date:** October 28, 2026
- **Resource Kind:** `AppEngineStandardAppVersion`
- **GCP Service:** `appengine.googleapis.com` (`google.appengine.v1.Versions`)

## Overview

The `mockappengine` package already existed for `AppEngineFirewallRule`, but did not implement the `Versions` service (`pb.VersionsServer`). We implemented the `mockappengine` Versions service using `httptogrpc` backed by `cloud.google.com/go/appengine/apiv1/appenginepb`.

## Real GCP Recording & Findings

1. **App Engine Application & Service Dependency**:
   - Versions reside under an App Engine service: `apps/{app_id}/services/{service_id}/versions/{version_id}`.
   - The GCP project must have an App Engine application initialized, and the default service exists automatically.
   - In KCC KRM, `spec.serviceRef.external` can be passed as `apps/${projectId}/services/default`.
2. **KRM serviceRef vs Terraform Provider**:
   - In KCC, the `serviceRef` identity parser expects `apps/{app}/services/{service}`. However, Terraform provider's `google_app_engine_standard_app_version` resource expects `service = "{service}"` (e.g. `default`).
   - A resource override in `pkg/resourceoverrides/appengine_standardappversion.go` was created to transform `serviceRef.external` from `apps/{app}/services/{service}` to `{service}` during pre-actuation, and preserve the original format for KRM state.
3. **Deployment Files & Staging Bucket**:
   - App Engine Standard versions require deployment code in Cloud Storage (e.g., zip or individual files).
   - In real GCP, App Engine provides an application staging bucket `staging.${projectId}.appspot.com`.
   - `deployment.files` can point directly to uploaded source files in `staging.${projectId}.appspot.com`.
4. **Terraform Deployment Read Bug & Ratcheting**:
   - The Terraform provider for `google_app_engine_standard_app_version` does not persist `deployment` into state during `Read()` (lacks `flattenAppEngineStandardAppVersionDeployment`). As a result, re-reconciliation detects a diff on `deployment.0.files` and re-applies.
   - Because legacy controllers cannot be rewritten here, `AppEngineStandardAppVersion` was handled in `tests/e2e/ratcheting.go` to disable re-reconciliation testing, with a detailed explanatory comment.
5. **REST API Behaviors & Normalization**:
   - In GET/LIST responses from real GCP, `entrypoint` is omitted for Standard App Engine versions.
   - Standard App Engine versions create LRO operations with names like `apps/{project}/operations/{operation_id}` and `OperationMetadataV1`.
   - Timestamps (`insertTime`, `endTime`, `createTime`) in operations and version objects are normalized to standard timestamp placeholders via `mockgcpregistry.SupportsNormalization`.
   - In update calls that submit version deployment files, real GCP appends the updated handler definitions with `AUTH_FAIL_ACTION_REDIRECT`.

## Implementation Details

- **Versions Server**: Implemented `GetVersion`, `CreateVersion`, `UpdateVersion`, `DeleteVersion`, and `ListVersions` in `mockgcp/mockappengine/versions.go`.
- **Operations Path**: Registered `/v1/{prefix=apps/*}/operations/{name}` with `httptogrpc`.
- **Service Mux**: Registered `pb.NewVersionsClient(conn)` and `pb.RegisterVersionsServer(grpcServer, ...)`.
- **Normalization**: Added `ConfigureVisitor` implementing `mockgcpregistry.SupportsNormalization` in `mockgcp/mockappengine/normalize.go`.

## Validation

- Recorded authentic real GCP golden logs (`_http.log`) and golden objects with `hack/record-gcp "fixtures/^appenginestandardappversionbasic$"`.
- Successfully ran `hack/compare-mock "fixtures/^appenginestandardappversionbasic$"` to generate `_http_mock.log` and verify full mock reconciliation.
- Successfully verified both `TestGoldenLogAlignment` and existing `appenginefirewallrulebasic` fixture tests.
