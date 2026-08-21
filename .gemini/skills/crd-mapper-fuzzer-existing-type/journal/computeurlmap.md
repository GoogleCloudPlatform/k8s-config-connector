# ComputeURLMap Direct KRM Types Transition Journal

## Overview
During the transition of `ComputeURLMap` (Compute URL Map) to direct KRM types, we ran into several interesting schema quirks and learned valuable lessons about achieving 100% strict schema compatibility.

## Key Observations & Learnings

### 1. Structural Redundancy / Struct Duplication for Schema Matching
Kubernetes CRD schema generation from Go structs uses properties of the types at compile time. In legacy KCC, the same Go struct could have required fields in one part of the CRD schema but optional fields in another part.
In direct KRM Go types, this is mapped 1:1, so a Go struct's requirements are identical everywhere it is used.
To resolve this structural difference and preserve the exact baseline CRD requirements without altering any schemas, we duplicated and created specific required/optional variants of several structs:
- `UrlmapRequestMirrorPolicy` (optional) vs `UrlmapRequestMirrorPolicyRequired` (required)
- `UrlmapWeightedBackendServices` (optional header action) vs `UrlmapWeightedBackendServicesRequired` (required header action)
- `UrlmapCorsPolicy` vs `UrlmapCorsPolicyRequired`
- `UrlmapFaultInjectionPolicy` vs `UrlmapFaultInjectionPolicyRequired` vs `UrlmapFaultInjectionPolicyRouteRules`
- `UrlmapDelay` vs `UrlmapDelayRequired` vs `UrlmapDelayRouteRules`
- `UrlmapFixedDelay` vs `UrlmapFixedDelayRequired`
- `UrlmapTimeout` vs `UrlmapTimeoutRequired`
- `UrlmapRetryPolicy` vs `UrlmapRetryPolicyOptional` vs `UrlmapRetryPolicyRequired`

This precise separation of types allowed us to align perfectly with the baseline CRD schema and have `dev/tasks/diff-crds` return a completely empty diff.

### 2. Custom Reference Types Without `kind` Field
In `ComputeURLMap`, references to other resources (like `BackendBucketRef` and `BackendServiceRef`) did not have a `kind` field in the baseline CRD schema, only `name`, `namespace`, and `external`.
To prevent the CRD generator from automatically adding a `kind` field to these reference structures, we defined a local `UrlmapResourceRef` struct that lacks the `kind` field, and used it instead of the shared `v1alpha1.ResourceRef` which does have a `kind` field in its Go struct definition.

### 3. Comment Commenting Quirks
`generate.sh` runs `prunetypes` which comments out any unreachable types in `types.generated.go`. When adding `ComputeURLMapSpec`, all related types inside `types.generated.go` were uncommented, but any other unused compute types remained commented, keeping the file size and symbol scope minimal.
