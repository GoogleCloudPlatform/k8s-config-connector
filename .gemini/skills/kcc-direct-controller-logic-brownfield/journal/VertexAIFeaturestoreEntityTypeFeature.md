# VertexAIFeaturestoreEntityTypeFeature Direct Migration Journal

## Overview
Successfully implemented the direct controller for `VertexAIFeaturestoreEntityTypeFeature` (`vertexai.cnrm.cloud.google.com/v1alpha1`) and verified dual-controller reconciliation against the legacy Terraform controller, MockGCP, and live GCP.

## Key Learnings & Observations

### 1. GAPIC Client and LRO Semantics
- The controller utilizes the GAPIC REST client `gcp.NewFeaturestoreRESTClient` with regional endpoint `https://{location}-aiplatform.googleapis.com`.
- `CreateFeature` returns an asynchronous Long-Running Operation (`*gcp.CreateFeatureOperation`). As noted in direct controller standards, LRO completion does not fully populate server-generated timestamps or read-only fields on the LRO response. We execute a synchronous `GetFeature` immediately after `op.Wait(ctx)` to fetch the complete object before updating KRM status.
- `UpdateFeature` in the Vertex AI Featurestore API is synchronous and returns `(*pb.Feature, error)` directly (no LRO wait needed).
- `DeleteFeature` is an asynchronous LRO returning `*gcp.DeleteFeatureOperation`.

### 2. Brownfield Label Reconciliation and Structured Diffs
- Standard brownfield label management was implemented using `label.GCPLabels(u)` in `buildDesired`, correctly mapping Kubernetes `metadata.labels` onto the GCP `labels` field.
- In `Update`, `common.CompareBrownfieldSpecAndLabels` compares both `Spec` fields and `metadata.labels`, generating the exact `updateMask` (`description`, `labels`).

### 3. Golden Log Alignment & Format Differences (GAPIC REST vs Terraform Discovery)
- The legacy Terraform controller used the Google API Go discovery client (`?alt=json`), sending and receiving string-encoded enums (`"valueType": "INT64"`).
- The direct controller uses the official GAPIC REST client (`?%24alt=json%3Benum-encoding%3Dint`), sending and receiving integer-encoded protobuf enums (`"valueType": 9`).
- Recording real GCP HTTP logs (`_http.log`) with the direct controller enabled aligns both real GCP and MockGCP logs cleanly, passing `TestGoldenLogAlignment` and in-place `TestMigrationToDirect` tests.
