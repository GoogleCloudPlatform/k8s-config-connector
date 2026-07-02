# Skill: Create Fuzzer

## Description

This skill provides expert guidance for generating fuzzers for Config Connector's "direct" controllers. Fuzzing helps ensure that resources round-trip correctly between their KRM representations and the GCP API representation.

## Context & Usage

You should activate this skill whenever you are tasked with creating a `*_fuzzer.go` for a new or existing resource in KCC, or when troubleshooting an issue with an existing fuzzer.

Fuzzers validate that `FromProto` and `ToProto` mappers do not lose or mutate data unintentionally.

## Implementation Guide

### 1. File Structure

Create a file named `<kind>_fuzzer.go` in the same directory as the controller. Add the `// +tool:fuzz-gen` tag if the file was auto-generated, otherwise regular copyright headers apply.

```go
package <package_name>

import (
	pb "path/to/gcp/api/package"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(myKindFuzzer())
	// or fuzztesting.RegisterKRMFuzzer_NoProto(myKindFuzzer())
}

func myKindFuzzer() fuzztesting.KRMFuzzer {
    // ...
}
```

### 2. Protobuf vs NoProto

GCP APIs use either Protobuf-based clients or older handwritten API clients.
*   **Protobuf Clients:** Use `fuzztesting.NewKRMTypedFuzzer`. You must pass an empty struct pointer of the proto message (e.g., `&pb.MyMessage{}`).
*   **Non-Protobuf Clients (NoProto):** Use `fuzztesting.NewKRMTypedFuzzer_NoProto`. This behaves similarly, but handles plain Go structs instead of `proto.Message`.

### 3. Field Mappings

The fuzz testing framework randomly fills the API struct, maps it to KRM, and maps it back. You must explicitly declare every field in the API struct.

*   `f.SpecField(".FieldName")`: The field is fully supported in the `Spec`.
*   `f.StatusField(".FieldName")`: The field is fully supported in the `Status`.
*   `f.IdentityField(".Name")`: The field represents the resource identity.
*   `f.Unimplemented_NotYetTriaged(".FieldName")`: The field is intentionally ignored or not yet implemented.

**Crucial detail for Slice Elements:** The fuzzing engine resets the field path prefix when evaluating elements of a slice or values in a map. For example, if you have `Access []AccessEntry` and `AccessEntry` has a `Condition` field that is not implemented, the correct path is `.Condition`, **not** `.Access.Condition`.

### 4. Custom Filters (FilterSpec / FilterStatus)

Sometimes, perfectly random data cannot survive a round-trip without custom logic. Provide a `FilterSpec` function to constrain or normalize the randomly generated `APIType` before the framework converts it.

Common use-cases:
*   **Enums:** Constrain an `int32` to a valid enum list.
*   **Unions / OneOfs:** If an API uses a struct with multiple fields acting as a "OneOf" (but not natively a Protobuf `oneof`), you must randomly select one and `nil` the others.
*   **Precision Loss:** Truncate `time.Duration` or `time.Time` fields if the KRM mapping only supports milliseconds or hours.

```go
	f.FilterSpec = func(in *pb.DatasetMetadata) {
		in.DefaultPartitionExpiration = in.DefaultPartitionExpiration.Truncate(time.Millisecond)
	}
```
