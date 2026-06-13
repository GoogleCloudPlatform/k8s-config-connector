# ComputeURLMap Greenfield KRM Types Journal

## Observations
- `ComputeURLMap` is a complex resource that heavily references other GCP resources (`ComputeBackendService` and `ComputeBackendBucket`) at multiple levels of nesting.
- Because `ComputeBackendBucket` was not previously defined as a direct reference type in Config Connector, we implemented `ComputeBackendBucketRef` under `apis/compute/v1beta1/backendbucket_reference.go` following the canonical structure of `ComputeBackendServiceRef`.
- Rather than relying solely on auto-generated mappers (which would struggle with type mismatches between protobuf's basic `string` pointer fields and KRM's custom reference structs), we implemented a comprehensive hand-written mapper `urlmap_mapper.go` in `pkg/controller/direct/compute/`.
- This approach cleanly overrides the top-level spec mapping using the SKIPPING mechanism of the mapper generator, while avoiding compiling issues and ensuring robust, fully type-safe translation of URLs and References.
