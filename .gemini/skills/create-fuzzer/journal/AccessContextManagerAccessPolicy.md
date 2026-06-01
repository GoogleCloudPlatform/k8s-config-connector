# Journal: AccessContextManagerAccessPolicy Fuzzer Integration

This journal documents the specific details and learnings from creating and registering the fuzzer for `AccessContextManagerAccessPolicy`.

## Context & Structure
- **Resource:** `AccessContextManagerAccessPolicy`
- **Location of Mappers:** `pkg/controller/direct/accesscontextmanager/`
- **Location of Fuzzer:** `pkg/controller/direct/accesscontextmanager/accesscontextmanageraccesspolicy_fuzzer.go`
- **GCP Proto Type:** `google.identity.accesscontextmanager.v1.AccessPolicy` (represented in Go as `*pb.AccessPolicy`)

## Implementation Details

1. **Mapper & Structure Alignment**
   - Both Spec and Status ObservedState mappers were already present in `mapper.generated.go`.
   - `AccessContextManagerAccessPolicyObservedState` is currently empty because fields like `CreateTime`, `UpdateTime`, and `Name` are directly under the KRM `Status` instead of being nested under `ObservedState`.

2. **Field Configuration in Fuzzer**
   - `.title` is mapped under KRM Spec, registered via `f.SpecField(".title")`.
   - `.name` and `.parent` represent GCP resource identities, so they are marked via `f.Unimplemented_Identity(".name")` and `f.Unimplemented_Identity(".parent")` respectively.
   - `.scopes`, `.create_time`, and `.update_time` are not yet fully mapped or triaged, so they are marked via `f.Unimplemented_NotYetTriaged(...)`.
   - `.etag` is marked via `f.Unimplemented_Etag()`.

3. **Global Package Initialization Registration**
   - Registered the `accesscontextmanager` package in `pkg/controller/direct/register/register.go` to ensure its fuzzer runs in the generic fuzz tests.

## Validation & Verification
- Created a deterministic unit test in `pkg/controller/direct/accesscontextmanager/accesscontextmanageraccesspolicy_fuzzer_test.go` to fuzz 1000 iterations over both Spec and Status round-trips.
- Running `go test -count=1 -v ./pkg/controller/direct/accesscontextmanager/` successfully passes.
