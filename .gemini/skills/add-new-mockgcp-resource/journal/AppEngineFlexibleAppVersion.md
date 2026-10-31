# AppEngineFlexibleAppVersion MockGCP Alignment Journal

- **Date:** October 31, 2026
- **Resource Kind:** `AppEngineFlexibleAppVersion`
- **GCP Service:** `appengine.googleapis.com` (`google.appengine.v1.Versions`)

## Overview

The `mockappengine` package implements the `Versions` service (`pb.VersionsServer`) for `AppEngineStandardAppVersion`. For `AppEngineFlexibleAppVersion`, we extended `mockappengine` to support Flexible App Engine versions (`env: "flexible"`), container deployments, and matching real GCP API response behaviors.

## Real GCP Recording & Findings

1. **API Enablement & Service Setup**:
   - `appengineflex.googleapis.com` must be enabled on the GCP project to deploy flexible app versions.
   - Recording was executed successfully against real GCP project `cnrm-barni-4` using `hack/record-gcp "fixtures/^appengineflexibleappversionbasic$"`.

2. **KRM `serviceRef` vs Terraform Provider**:
   - KCC's `serviceRef` reference model uses `apps/${projectId}/services/default`.
   - The Terraform provider for `google_app_engine_flexible_app_version` expects `service = "default"`.
   - Added `GetAppEngineFlexibleAppVersionResourceOverrides` in `pkg/resourceoverrides/appengine_standardappversion.go` (and registered in `pkg/resourceoverrides/overrides.go`) to trim `apps/{app}/services/{service}` to `{service}` during pre-actuation and restore it post-actuation.

3. **Container Deployment & Runtimes**:
   - For flexible app versions, using `runtime: custom` with `deployment.container.image` (e.g., `gcr.io/google-samples/hello-app:1.0`) avoids slow Cloud Build source builds and allows clean container deployment.
   - Real GCP resolves the container image tag (`:1.0`) to a full digest (`@sha256:...`).

4. **REST API Behaviors & Differences from Standard Versions**:
   - `env`: Terraform sends `env: "flex"`, but real GCP stores and returns `env: "flexible"`.
   - Flexible app versions do NOT return default `instanceClass: "F1"` (omitted unless explicitly set).
   - Flexible app versions with container images do NOT return default `handlers` or `diskUsageBytes`.
   - `runtimeChannel: "default"` is omitted for custom flexible runtimes.
   - Real GCP returns `betaSettings` containing `no_appserver_affinity: "true"`, `use_deployment_manager: "true"`, and a dynamic internal version ID `minor_version_id_ccfe`.

5. **Normalization**:
   - `minor_version_id_ccfe` is a server-generated random number that changes on every deployment/update. It was normalized in `mockappengine/normalize.go` and normalized in `pkg/test/resourcefixture/golden_alignment_test.go` (`normalizeRepresentation`).

6. **Terraform Container Image Diff & Ratcheting**:
   - Real GCP resolves container image tags to digests (`@sha256:649188051906c9df7f87fb2ef04a21acbdb009c0877fa67b3132b83230411f6e`).
   - The legacy Terraform provider `google_app_engine_flexible_app_version` lacks a `DiffSuppressFunc` between image tags and digests. Consequently, during re-reconciliation, Terraform detects a diff on `deployment.0.container.0.image` and triggers an update.
   - Handled in `tests/e2e/ratcheting.go` to disable re-reconciliation testing with an explanatory comment, matching the existing behavior for `AppEngineStandardAppVersion`.

## Validation

- Successfully recorded real GCP golden HTTP log (`_http.log`) and golden object file using `hack/record-gcp`.
- Verified with `hack/compare-mock "fixtures/appengine"`: all App Engine fixtures (`appenginefirewallrulebasic`, `appengineflexibleappversionbasic`, `appengineservicesplittrafficbasic`, `appenginestandardappversionbasic`) pass with 0 errors.
- Verified with `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/appengine"`: all App Engine fixtures pass golden log alignment with 0 errors.
- Verified with `go test -v ./tests/apichecks`: all CRD and API checks pass cleanly.
