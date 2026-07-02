# Fuzzer Creation Journal

This journal documents less general or edge-case observations made while creating fuzzers for KCC resources.

## 2026-05-06: BigQueryDataset Fuzzer and krmgen.go Fixes

*   **NoProto Hand-written Clients:** BigQueryDataset uses the hand-written cloud.google.com/go/bigquery client rather than a standard Protobuf GRPC API. This necessitated using NewKRMTypedFuzzer_NoProto.
*   **Precision Loss in Time fields:** time.Duration fields in the BigQuery client (DefaultTableExpiration, MaxTimeTravel) had precision loss when mapped to KRM (e.g., truncating to hours or milliseconds). The FilterSpec had to enforce this truncation on the generated API struct to ensure roundtripping succeeded.
*   **OneOf fields in NoProto:** The BigQuery AccessEntry struct had multiple Entity fields, but the mapper relied on the EntityType field to dictate which one was valid. The randomly generated struct needed a FilterSpec to enforce valid EntityType enums and clear out invalid sibling fields, because the fuzzer randomly populates all of them.
*   **krmgen.go Pointer/Unexported Bug:** Found and fixed a panic in KCC's pkg/test/fuzz/krmgen.go. The fuzzer recursively visits fields. If an API struct contains an unexported pointer field (like .c *Client inside bigquery.Table), krmgen.go would blindly instantiate the pointer field.Set(reflect.New(...)) before checking if the field was overriden/ignored or if it CanSet(). We patched krmgen.go to respect field.CanSet() before initializing pointers and clearing non-proto fields.
*   **Slice Element Field Paths:** Confirmed that krmgen.go deliberately clears the fieldName string when iterating into slice or map elements. This means any field path overrides or ignores inside a slice element must be specified relative to the element type, not the root type. For instance, f.Unimplemented_NotYetTriaged(".Condition") instead of f.Unimplemented_NotYetTriaged(".Access.Condition").

*   **Real Bug Found by Fuzzer (strings.Trim vs strings.TrimPrefix):** During the BigQueryDataset fuzzer implementation, it was discovered that BigQueryDatasetStatus_ToProto used strings.Trim(..., "https://...") instead of strings.TrimPrefix. This caused the SelfLink string to be heavily mangled (stripping characters from both ends), which broke the parsing logic for restoring FullID. This proves the value of the NoProto fuzzer.
