# Journal: Direct Migration of FilestoreInstance

Author: overseer,overseer (gemini-3.6-flash)
Date: September 6, 2026

## Learnings & Observations

### 1. Decoupling Reference Normalization from Adapter Initialization
**Problem:** In many direct controllers, `common.NormalizeReferences` is called inside `AdapterForObject` as part of the adapter's construction. During resource deletion (especially during the teardown phase of e2e migration/fixture tests), dependencies (like a `ComputeNetwork` or `Project`) can be deleted before or concurrently with the primary resource. If `NormalizeReferences` is called during deletion and fails to find the referenced dependency in the Kubernetes API server, it blocks the parent controller from constructing the adapter. As a result, the controller gets stuck in a loop, and the resource is never finalized or deleted.

**Solution:** Move `common.NormalizeReferences` and spec-to-proto mapping out of `AdapterForObject` and into the adapter's `Create` and `Update` methods. Since the `Delete` and `Find` operations only require the resource's GCP URL/identity (which does not depend on resolving referenced fields), they are completely decoupled from reference resolution. This allows resource teardown to be robust, resilient, and order-independent in all environments.

### 2. Stripping Volatile Properties of Shared Resources in HTTP Logs
**Problem:** In end-to-end tests running concurrently against a shared live GCP project, resources such as compute networks are often shared or receive dynamic peering connections from other running tests (e.g. dynamic redis or filestore private network peerings). These peerings, along with subnetwork and routing configuration updates, appear in the `compute#network` GET responses returned during the `FilestoreInstance` test, leading to volatile HTTP cassettes and unexpected golden log mismatches.

**Solution:** In `tests/e2e/normalize_legacy.go`, there is a normalizer that strips `peerings`, `routingConfig`, and `subnetworks` from `compute#network` resources, but it was historically limited to tests containing `computeaddress` or `containercluster` in their name. By adding `filestoreinstance` to this name check, the volatile shared properties of the default network are automatically stripped from the captured HTTP traffic, ensuring the golden tests are 100% stable.

### 3. Masking Immutable Fields during Updates
**Problem:** The GCP Filestore API returns a 400 error `unsupported field mask "networks"` if fields that are immutable (such as `networks` or `tier`) are sent inside the `UpdateMask` parameter of the `UpdateInstance` REST call.
**Solution:** Copy the immutable `Networks` and `Tier` fields directly from the `actual` (GCP) protobuf message into the `clonedDesired` protobuf message inside the `Update` method before calling `CompareProtoMessageStructuredDiff`. This eliminates any false diffs on those fields. Additionally, explicitly filter the field mask paths returned by `CompareProtoMessageStructuredDiff` to only allow mutable paths (`description`, `file_shares`, and `labels`).
