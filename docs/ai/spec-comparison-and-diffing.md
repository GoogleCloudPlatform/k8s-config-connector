# KRM-Roundtrip Spec Comparison: `CompareBrownfieldSpec` Review & `CompareGreenfieldSpec` Design

## 1. Background & Motivation

Across the 370 existing direct controllers in `pkg/controller/direct/`, the single largest source of hand-written controller logic (**64 / 370 controllers, 17.3%**) is custom `populateDefaults(desired, actual)` boilerplate in `Update()`.

When GCP populates server-side default values on a resource (`actual`) for fields that the user left unset (`nil`) in KRM `spec` (`desired`), a naive proto comparison reports a permanent false-positive diff on every reconcile loop.

To solve this systematically, [`pkg/controller/direct/common/compare.go`](../../pkg/controller/direct/common/compare.go) introduced a 5-step KRM-roundtrip comparison pipeline (`CompareBrownfieldSpec` and `CompareBrownfieldSpecAndLabels`):

1. **Map Live State to KRM (`specFromProto`):** Convert `actualProto` into `actualKRM`.
2. **Adopt Unset Fields (`MergeUnsetFields`):** Recursively copy non-nil pointer/slice/map fields from `actualKRM` into a cloned copy of `desiredKRM` wherever `desiredKRM` is `nil`.
3. **Map Back to Proto (`specToProto`):** Convert both `clonedDesiredKRM` and `actualKRM` into `desiredProtoMasked` and `actualProtoMasked` (which strips `ObservedState`-only fields not present in `Spec`).
4. **Normalize (`normalize`):** Apply optional resource-specific canonicalization.
5. **Compute Top-Level Diff & `UpdateMask` (`DiffForTopLevelFields`):** Compare top-level proto fields between `desiredProtoMasked` and `actualProtoMasked`.

---

## 2. Findings in `CompareBrownfieldSpec`

During the design review for deterministic controller generation (`KCC-DETGEN` Phase 2), we audited `CompareBrownfieldSpec` and its callers (e.g., `eventarcmessagebus_controller.go`, `filestoreinstance_controller.go`) and identified 5 issues to revisit with the team:

### Finding 1: Discarding `desiredProtoMasked` + Returning a Top-Level `UpdateMask` Causes Nested Data Loss
- **Current behavior:** `CompareBrownfieldSpec` returns `(*structuredreporting.Diff, *fieldmaskpb.FieldMask, error)` and discards `desiredProtoMasked`. Furthermore, `DiffForTopLevelFields` emits **top-level field names** in `UpdateMask.Paths` (e.g., `["encryption_config"]` rather than `["encryption_config.kms_key"]`).
- **Why this is dangerous:** Callers pass their *unmerged* `desired` proto into `Update<Resource>Request` along with the top-level `UpdateMask`. If a nested struct (e.g., `spec.encryptionConfig`) has two sub-fields—`fieldA` (set and updated by the user) and `fieldB` (unset by the user, defaulted by GCP)—sending unmerged `desired` with `UpdateMask: ["encryption_config"]` **overwrites `fieldB` with its zero value on the server**.
- **Impact on `F-update-without-mask` controllers:** For the ~18 GCP resources whose `Update` RPC does not take an `update_mask`, sending unmerged `desired` wipes out *top-level* server-defaulted fields as well.
- **Fix in `CompareGreenfieldSpec`:** Return `mergedDesiredProto` (`(ProtoT, *structuredreporting.Diff, *fieldmaskpb.FieldMask, error)`) cloned prior to diff-only normalization hooks so callers send `mergedDesiredProto` in `Update<Resource>Request`.

### Finding 2: Unconditional Global `SortRepeatedFields` Silently Drops Order-Sensitive List Updates
- **Current behavior:** Step 4 of `CompareBrownfieldSpec` unconditionally calls `SortRepeatedFields` on both `desiredProtoMasked` and `actualProtoMasked`, sorting **every** proto `repeated` field (`[]string`, `[]SubMessage`, etc.) at every nesting depth.
- **Why this is dangerous:** Many GCP repeated fields are order-sensitive (e.g., container `command`/`args`, ordered firewall/routing rules, priority lists). If a user reorders elements in `spec`, `SortRepeatedFields` sorts both `desired` and `actual` into the same order and **reports zero diff, silently dropping the update**.
- **Fix in `CompareGreenfieldSpec`:** Remove global `SortRepeatedFields` from default comparison. When a specific repeated field is an unordered set in GCP, slice sorting should be opted into explicitly inside the resource's `normalize` callback.

### Finding 3: Dropped `OUTPUT_ONLY` and `IMMUTABLE` Proto `field_behavior` Checks
- **Current behavior:** The legacy `BasicDiff` helper checked `google.api.field_behavior` for `OUTPUT_ONLY` (skipping the field) and `IMMUTABLE` (returning an error when modified). `DiffForTopLevelFields` dropped both checks, relying entirely on `Spec_FromProto`/`Spec_ToProto` omitting `OUTPUT_ONLY` fields.
- **Why this matters:**
  1. If a mapper or `normalize` callback touches a proto field marked `OUTPUT_ONLY`, `DiffForTopLevelFields` will include it in `UpdateMask`, causing the GCP API to reject the `Update` call.
  2. If a user modifies an `IMMUTABLE` field in `spec`, `DiffForTopLevelFields` includes it in `UpdateMask` instead of surfacing a clear `change to immutable field <name>` error (note: unset `IMMUTABLE` fields in `desiredKRM` are safely populated from `actualKRM` in Step 2 and therefore do not trigger an error).
- **Fix in `CompareGreenfieldSpec`:** `DiffForTopLevelFieldsGreenfield` restores both `FieldBehavior_OUTPUT_ONLY` and `FieldBehavior_IMMUTABLE` checks.

### Finding 4: Populating Labels onto `mergedDesiredProto`
- **Current behavior:** `CompareBrownfieldSpecAndLabels` compares `label.GCPLabels(u)` against `actualProto` and appends `labelsFieldName` to `updateMask.Paths`, but leaves the caller responsible for remembering to copy `label.GCPLabels(u)` onto the proto sent in `Update<Resource>Request`.
- **Fix in `CompareGreenfieldSpecAndLabels`:** Automatically populates `desiredLabels` onto `mergedDesiredProto`'s `labelsFieldName` map field so `mergedDesiredProto` is completely ready to send in the `Update` request.

### Finding 5: Open Question for Team Review — Clearing Previously Set Optional Slices/Maps/Pointers
- **Current behavior in `MergeUnsetFields`:** Because `MergeUnsetFields` only inspects `desiredKRM` (current `spec`) and `actualKRM` (live GCP state), any `nil` pointer, slice, or map in `desiredKRM` adopts `actualKRM`'s value.
- **Trade-off to discuss with the team:** While this prevents false diffs on server-defaulted fields, it also means that if a user previously set an optional field in `spec` and later deleted that line from their YAML (making it `nil` in `desiredKRM`), `MergeUnsetFields` treats it as "unspecified (adopt live state)" rather than "explicitly cleared by user". Distinguishing "never set by user" from "removed by user" requires consulting Kubernetes field ownership (`metadata.managedFields` / `cnrm.cloud.google.com/last-applied-configuration`) rather than `spec` + `actual` alone.

---

## 3. Why We Branched `CompareGreenfieldSpec` & `CompareGreenfieldSpecAndLabels`

To ensure zero regressions for existing brownfield controllers that already call `CompareBrownfieldSpec` / `CompareBrownfieldSpecAndLabels` (and may rely on global slice sorting or the 3-value return signature), we:
1. Kept `CompareBrownfieldSpec` and `CompareBrownfieldSpecAndLabels` unchanged.
2. Added `CompareGreenfieldSpec`, `CompareGreenfieldSpecAndLabels`, and `DiffForTopLevelFieldsGreenfield` in [`pkg/controller/direct/common/greenfield.go`](../../pkg/controller/direct/common/greenfield.go) for generated and greenfield direct controllers.
