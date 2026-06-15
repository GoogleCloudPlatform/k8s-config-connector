# ComputeSnapshot Identity and Reference Journal

## Observations & Learnings

- **Global and Regional Resource Integration:** `ComputeSnapshot` can represent both global and regional snapshots in GCP. While snapshots are typically global, they can be regional as well.
- **Dual Template Pattern via Struct Mapping:** Similar to `ComputeDisk`, we define templates for both global and regional snapshots to handle parsing and formatting from the unified `ComputeSnapshotIdentity` struct.
- **Handling No Location in Spec:** The `ComputeSnapshot` KRM spec does not specify a `region` or `location` field. Therefore, all snapshots constructed via the KRM spec are global by default, while external snapshots referenced via the `.external` field can be either global or regional.
- **No Schema Changes:** Complying strictly with safety guidelines, no fields were added or modified in the CRD schema. The identity is successfully matched and cross-checked against the `status.selfLink` field, which was already present.
