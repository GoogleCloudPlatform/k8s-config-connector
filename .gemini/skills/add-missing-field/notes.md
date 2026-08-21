# Notes: Adding Missing Fields

- When modifying a hand-coded type or mapper, always ensure the struct tags (`json`, etc.) match Kubernetes API conventions.
- If a missing field is part of a larger nested object, consider whether to add the entire nested object or just the required fields.
- Some fields are output-only (Status/ObservedState) vs input (Spec). Fuzzer lists are usually indicative of the GCP API payload. Use your judgement and the API docs if available.
- `f.Unimplemented_NotYetTriaged` is a good temporary place for fields we discover but are out-of-scope for the current PR.

## MemorystoreInstance: automatedBackupConfig
- When modifying `_types.go` or `mapper.go`, check if common types like `TimeOfDay` have already been mapped manually (e.g. inside a global `mapper.go`) before trying to map them. You might encounter redeclaration errors during `go build`.
- If a protobuf type only contains input fields, `types.generated.go` will not generate a corresponding `*ObservedState` struct for it. You should only use the input type struct.
- For nested types where `mapper.generated.go` leaves a comment like `// MISSING: <Type>`, look for skipped parent mappers (indicated by `/* found existing non-generated mapping function ... */`) and manually implement `_FromProto` and `_ToProto` for the missing type.