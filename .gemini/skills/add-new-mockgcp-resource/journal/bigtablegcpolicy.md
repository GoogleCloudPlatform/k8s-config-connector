# Journal: Mocking BigtableGCPolicy

## Overview
`BigtableGCPolicy` is a KCC resource that manages GC Policies (GcRules) on column families inside a Bigtable Table. 

## Reconciliation & Mocking Mechanism
Instead of having its own dedicated API endpoint on GCP, Bigtable column families and their GC rules are modified using the `ModifyColumnFamilies` endpoint on a `Table` within the `BigtableTableAdmin` service (`google.bigtable.admin.v2.BigtableTableAdmin`). 

Because KCC delegates the reconciliation of `BigtableGCPolicy` to Terraform, under the hood it sends a series of `ModifyColumnFamilies` requests to update column family schemas:
- Creating a table with column families initialized with empty/default GC rules.
- Modifying column families to specify/update GcRule parameters (e.g. `union` containing rules like `maxAge` or `maxNumVersions`).
- Setting/cleaning up GC rules during deletion.

As a result, `BigtableGCPolicy` is naturally supported in `mockgcp` via the existing `mockbigtable` package (specifically `table.go`).

## Alignment of MockGCP with Real GCP
To match real GCP behavior perfectly, we analyzed `ModifyColumnFamiliesRequest_Modification_Update` and aligned the mock `ModifyColumnFamilies` implementation:
- Previously, when an update modification was received, the mock completely overwrote the existing `ColumnFamily` with the new schema specified in `mod.Update`.
- Real GCP supports an optional `UpdateMask` inside the modification request. If set, only specified fields (`gc_rule` or `value_type`) are updated. If empty/unset, it defaults to updating `gc_rule` only for backward compatibility.
- We updated `ModifyColumnFamilies` in `mockbigtable/table.go` to properly parse and respect this `UpdateMask`. This ensures other properties of the column family (such as `value_type` or other schema configuration) are correctly preserved when updating column family schemas.
