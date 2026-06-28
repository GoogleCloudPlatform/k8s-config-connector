# MockGCP Journal: ComputeRouter

## Context
Compute Engine has a regional Cloud Router resource type `/projects/{project}/regions/{region}/routers/{router}`. 
We needed to implement support for MockGCP to properly emulate and reconcile `ComputeRouter` resources in the `compute.cnrm.cloud.google.com` API group.

## Design Decision
We utilized the modern `grpc-gateway` generated `pb.RoutersServer` interface and `pb.RegisterRoutersHandler` to implement a robust mock:

1. **Service Implementation**:
   - Implemented `RoutersV1` mock under `mockgcp/mockcompute/routersv1.go` that conforms to `pb.RoutersServer`.
   - Handled standard regional lifecycle operations: `Get`, `Insert`, `Patch`, and `Delete`.
   - Utilized regional LRO (Long Running Operation) helpers to simulate asynchronous GCP operations.

2. **Dependency Resolution**:
   - Resolved references like `networkRef` to full canonical Compute Network self-links by reusing existing `parseNetworkSelfLink` utility functions.

3. **Validation & Enablement**:
   - Registered `RoutersServer` and `RoutersHandler` in `mockgcp/mockcompute/service.go`.
   - Added `ComputeRouter` to the allowed `mockgcp` validation list in `config/tests/samples/create/harness.go` to opt-in the resource.

## Benefits
- Perfect fidelity to real GCP responses (as validated by the `_http.log` trace).
- Integrated seamlessly into the existing e2e test runner without requiring any test adjustments.
