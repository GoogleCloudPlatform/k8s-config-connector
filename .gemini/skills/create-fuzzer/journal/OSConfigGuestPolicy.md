# Fuzzer Journal: OSConfigGuestPolicy

## Observations and Learnings

- **Top-Level Status Fields**: Unlike many other KRM schemas where `CreateTime`, `UpdateTime`, and `Etag` are nested inside `status.observedState`, `OSConfigGuestPolicyStatus` places these fields directly at the top level of the status object.
- **Manual Status Mapping**: To facilitate round-trip validation of status fields, we implemented manual mapping functions:
  - `OSConfigGuestPolicyStatus_FromProto`
  - `OSConfigGuestPolicyStatus_ToProto`
- **Proto Compatibility**: The corresponding proto structure `osconfigpb.GuestPolicy` has native fields for `create_time`, `update_time`, and `etag`, which makes direct mapping extremely straightforward and robust.
- **Fuzzer Verification**: By registering the Spec & Status mappers with the typed KRM fuzzer framework, we successfully round-tripped and proved the lossless mapping of all `OSConfigGuestPolicy` spec and status fields.
