---
name: kcc-direct-proto-comparison
description: Standards and rules for comparing desired and actual protobuf messages in KCC direct controllers without triggering infinite update loops from output-only or unmapped fields.
---

# KCC Direct Controller Protobuf Comparison Rules

This skill defines the mandatory rules for comparing desired state (`desiredPb` mapped via `*_ToProto`) against actual GCP state (`a.actual`) in `Adapter.Update()` across KCC direct controllers.

## Background & Root Cause of Infinite Update Loops

When a direct controller reconciles an existing GCP resource in `Update()`:
1. `desiredPb` is constructed from the KRM `Spec` using `<Kind>Spec_ToProto(mapCtx, &desired.Spec)` (or sub-message `*_ToProto` mappers). Any proto fields that are output-only (`[(google.api.field_behavior) = OUTPUT_ONLY]`), server-computed (`etag`, `update_time`, `uid`), or mapped only to `status.observedState` (`// MISSING:` in `*_ToProto`) remain unpopulated (zero/nil) in `desiredPb`.
2. `a.actual` (and nested sub-messages like `a.actual.GetSpec()`, `a.actual.GetQuotaConfig()`, `a.actual.GetBackupSchedule()`, `a.actual.GetNetworkConfig()`, `a.actual.Fields`, `a.actual.GetGitRemoteSettings()`) is returned directly by the GCP API and has those output-only / server-populated fields set.

### The Anti-Pattern (Causes Endless Update Loops)
Directly comparing raw `a.actual` (or a nested sub-message/map/slice of `a.actual`) against `desiredPb` using `proto.Equal` or `reflect.DeepEqual`:

```go
// ANTI-PATTERN 1: proto.Equal on raw actual sub-message containing output-only fields (e.g. etag, update_time)
if proto.Equal(a.actual.GetSpec(), desiredPb.GetSpec()) { ... }

// ANTI-PATTERN 2: reflect.DeepEqual on protobuf messages or maps of protobuf messages
// reflect.DeepEqual compares both output-only fields AND unexported protobuf internal fields (sizeCache, state).
if !reflect.DeepEqual(resource.QuotaConfig, a.actual.QuotaConfig) { ... }
```

Because `a.actual` contains output-only fields that `*_ToProto` never populates on `desiredPb`, the equality check **always returns `false`**, causing the controller to issue an `Update`/`PATCH` API call on every single reconciliation loop (violating KCC's zero-write steady-state invariant).

---

## Mandatory Rules for Comparing Desired vs. Actual State in `Update()`

### Rule 1: Prefer Canonical Whole-Spec Comparison Helpers
Whenever possible, use KCC's built-in spec comparison helpers, which automatically round-trip `a.actual` through `<Kind>Spec_FromProto` and `<Kind>Spec_ToProto` (stripping unmapped `// MISSING:` fields) and clear `OUTPUT_ONLY` / `INPUT_ONLY` proto fields:
- **Greenfield controllers:** Use `common.CompareGreenfieldSpec` (or `common.CompareGreenfieldSpecAndLabels`).
- **Brownfield controllers:** Use `common.CompareBrownfieldSpec` (or `common.CompareBrownfieldSpecAndLabels`).

### Rule 2: Never Use `reflect.DeepEqual` on Protobuf Messages
- Never pass `proto.Message` structs (or slices/maps of `proto.Message`) to `reflect.DeepEqual`. Protobuf structs contain unexported internal state (`state`, `sizeCache`, `unknownFields`) that can differ between freshly constructed protos and deserialized API responses even when all proto fields match.
- Always use `proto.Equal` (after masking/stripping output-only fields) or `common.CompareProtoMessage`.

### Rule 3: Mask or Strip Output-Only / `// MISSING:` Fields Before `proto.Equal`
When comparing individual sub-messages, maps, or slices manually in `Update()`:
1. **Inspect the `*_ToProto` mapper** for that message (and all nested child messages) in `mapper.generated.go` or manual mapper files. Check for any `// MISSING:` comments or proto fields annotated with `Output only` in the Go protobuf struct definition.
2. **Option A — Round-trip `actual` through `*_FromProto` -> `*_ToProto` (Preferred when mappers exist):**
   ```go
   func quotaConfigEqual(mapCtx *direct.MapContext, desired, actual *pb.QuotaConfig) bool {
       if desired == nil && actual == nil {
           return true
       }
       if desired == nil || actual == nil {
           return false
       }
       maskedActual := QuotaConfig_ToProto(mapCtx, QuotaConfig_FromProto(mapCtx, actual))
       return proto.Equal(desired, maskedActual)
   }
   ```
3. **Option B — Clone `actual` and clear output-only fields before `proto.Equal`:**
   ```go
   func policySpecEqual(actual, desired *orgpolicypb.PolicySpec) bool {
       if actual == nil && desired == nil {
           return true
       }
       if actual == nil || desired == nil {
           return false
       }
       actualCopy := proto.Clone(actual).(*orgpolicypb.PolicySpec)
       actualCopy.Etag = ""
       actualCopy.UpdateTime = nil
       return proto.Equal(actualCopy, desired)
   }
   ```

### Rule 4: Unit-Test Custom Equality Helpers with Populated Output-Only Fields
Whenever a controller defines a custom equality helper (such as `policySpecEqual`, `quotaConfigEqual`, `backupSchedulesEqual`, `networkConfigEqual`, `tagFieldsEqual`, `gitRemoteSettingsEqual`), add a unit test in `<resource>_controller_test.go` verifying that the helper returns `true` when `actual` has its output-only fields populated and all spec-mapped fields match `desired`.
