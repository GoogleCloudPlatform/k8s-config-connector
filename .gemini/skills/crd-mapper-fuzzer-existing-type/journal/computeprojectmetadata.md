# Journal: ComputeProjectMetadata Direct KRM types transition

## Observations and Learnings

- **Strict Schema Compatibility**: The baseline `ComputeProjectMetadata` CRD contains only a single `metadata` map of string-to-string in its `spec`, with no other standard fields like parent references, locations, or resource IDs.
- **Handcoded Types Over Generated**: The GCE GRPc/Protobuf `Metadata` resource represents metadata as a slice of `Items` (`repeated Items items`). Automatically generating KRM structures from this proto message would result in an incompatible schema (`spec.items` slice of objects) compared to the original CRD's `spec.metadata` map. Thus, we hand-coded the types in `metadata_types.go` to maintain 100% strict schema compatibility.
- **Pluralization Correction via kubebuilder Marker**: When controller-gen compiled `Metadata`, it generated the CRD resource path/plural form as `computeprojectmetadata` (singular). This mismatched the original CRD which used `computeprojectmetadatas`. We resolved this by explicitly appending `path=computeprojectmetadatas` to the `// +kubebuilder:resource:` marker.
- **Required Spec Field**: The original CRD marked `spec` as a required field at the root of the OpenAPI v3 schema. To match this exactly, we removed the `omitempty` tag from the `Spec` field on the `ComputeProjectMetadata` struct.
- **Stable and Deterministic Map Serialization**: Since map iterations in Go are non-deterministic, converting from `map[string]string` in KRM `Spec` to `repeated Items` in Protobuf `Metadata` can result in unstable proto payloads. We implemented deterministic sorting by sorting the items slice alphabetically by key before returning.
