# Notes: Adding Missing Fields

- When modifying a hand-coded type or mapper, always ensure the struct tags (`json`, etc.) match Kubernetes API conventions.
- If a missing field is part of a larger nested object, consider whether to add the entire nested object or just the required fields.
- Some fields are output-only (Status/ObservedState) vs input (Spec). Fuzzer lists are usually indicative of the GCP API payload. Use your judgement and the API docs if available.
- `f.Unimplemented_NotYetTriaged` is a good temporary place for fields we discover but are out-of-scope for the current PR.
