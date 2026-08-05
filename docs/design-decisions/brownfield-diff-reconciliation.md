# Design Decision Document: Brownfield KRM Spec Comparison & Reconciliation Pipeline for Direct Controllers

## 1. Problem Statement

In Config Connector (KCC) direct controllers, **Brownfield resources** need to support the legacy "unmanage" behavior for backward compatibility. This means the comparison of the desired and the actual state for a brownfield direct controller needs to "skip" unmanaged fields. The reconciler must compare `desiredKRM` against live `actualProto` to only detect legitimate drift on user-managed fields.

Currently (as of July 31 2026), most of the state comparisons for brownfield resources are hand-written. We should provide a generic solution to avoid the manual work.

One thing to note is that the list field is always managed by KCC once it is set in the desired state. We currently don't distinguish between list elements that are managed by KCC and those that are not. We also don't support unmanaged nested fields in the list element. More details can be found at https://docs.cloud.google.com/config-connector/docs/concepts/managing-fields-externally#behavior_for_list_fields_in_resource_spec.

## 2. Proposed Solution: The 5-Step Pipeline (6 States)

We introduce a single generic function, `common.CompareBrownfieldSpec`, in `pkg/controller/direct/common/compare.go`. It processes `desired` and `actual` through 5 transformation steps across 6 states:

```
A.  actualProto (GCP API Response)
       │
       ▼  1. actualKRM = Spec_FromProto(actualProto)
B.  actualKRM (Strips status, output-only, & unsupported fields)
       │
       ▼  2. MergeUnsetFields(desiredKRM, actualKRM)
C.  desiredKRM (Adopts actualKRM's values for unspecified fields)
       │
       │  3. desiredProtoMasked = Spec_ToProto(desiredKRM)
       ▼     actualProtoMasked  = Spec_ToProto(actualKRM)
D.  [desiredProtoMasked, actualProtoMasked]
       │
       ▼  4. Normalize & SortRepeatedFields on both Protos
E.  [Normalized & Deterministically Sorted Protos]
       │
       ▼  5. DiffForTopLevelFields(desiredProtoMasked, actualProtoMasked)
F.  (structuredreporting.Diff, fieldmaskpb.FieldMask)
```

## 3. Detailed Step Breakdown

### Step 1: `actualProto` -> `actualKRM` (`Spec_FromProto`)
- Converts raw GCP Proto to KRM Spec struct using each resource's own mapper `Spec_FromProto`.
- **Key Benefit**: Inherently masks out all status fields, `etag`s, output-only fields, and unsupported GCP API fields.

### Step 2: Merge Unspecified Fields (`MergeUnsetFields`)
- Recursively traverses `desiredKRM` and `actualKRM` using Go reflection.
- For any pointer/slice field where `desiredKRM` is `nil` (unspecified by the user), it copies the value from `actualKRM`.
- **Key Benefit**: Adopts existing values from actual into `desiredKRM` for both top-level and nested fields.

### Step 3: Convert KRM Specs to Protos (`Spec_ToProto`)
- Converts both `desiredKRM` and `actualKRM` to Proto messages (`desiredProtoMasked` and `actualProtoMasked`).
- **Key Benefit**: Ensures both Proto messages share identical schema boundaries and mapper normalization.

### Step 4: Normalization & Slice Sorting (`normalize` & `SortRepeatedFields`)
- Applies controller-specific link/reference normalizations if provided (e.g. converting GCP project numbers to project IDs in parent links). Normalization is required for resources with reference or link fields to ensure consistent formatting between desired and actual.
- Traverses repeated fields (slices) in both Proto messages to normalize order for unordered lists.
- **Key Benefit**: Eliminates false diffs caused by reference format mismatches or GCP list reordering.

### Step 5: Top-Level Diff & FieldMask Generation (`DiffForTopLevelFields`)
- Compares top-level fields between `desiredProtoMasked` and `actualProtoMasked`.
- **Key Benefit**: Emits structured diffs for logging and returns top-level field paths for `fieldmaskpb.FieldMask`, which is required by GCP Update RPCs.
