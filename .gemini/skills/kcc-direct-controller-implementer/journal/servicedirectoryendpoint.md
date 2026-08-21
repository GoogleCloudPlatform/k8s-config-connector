# ServiceDirectoryEndpoint Direct Controller Journal

## Overview
Implemented the direct controller for `ServiceDirectoryEndpoint` at `pkg/controller/direct/servicedirectory/servicedirectoryendpoint_controller.go` and verified the `servicedirectoryendpoint` E2E fixtures against `mockgcp`.

## Observations & Implementation Details
- **Mock GCP Endpoint Methods**: Added the mock Endpoint REST methods (`GetEndpoint`, `CreateEndpoint`, `UpdateEndpoint`, `DeleteEndpoint`) on `RegistrationServiceV1` in `mockgcp/mockservicedirectory/endpointv1.go` to support local E2E verification.
- **Reference Normalization**: Successfully resolved KRM references to parent services, address references, and network references using `common.NormalizeReferences`.
- **ComputeNetwork Format Update**: Discovered that the older legacy controller allowed `locations/global/networks` structure, but the direct controller expects the canonical `projects/{project}/global/networks/{network}` pattern for ComputeNetwork references. Updated the test fixture YAML files to reflect this.
- **Dynamic Controller Testing**: In line with current practices, added the direct controller as a supported reconciler under `static_config.go` and registered the GVK in `config/tests/samples/create/harness.go` to automatically run and verify both old and new controller reconcile paths.
