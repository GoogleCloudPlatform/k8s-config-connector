# Fuzzer Journal: PubSubTopic

## Observations
- `PubSubTopic` does not currently map `status` (Observed State) fields. For resources like this, `NewKRMTypedSpecFuzzer` and `RegisterKRMSpecFuzzer` are ideal as they only fuzz `Spec` mapping without requiring empty status mappings.
- The `google.pubsub.v1.Topic` protobuf message has several service-managed fields (like `.tags`, `.schema_settings.first_revision_id`, `.schema_settings.last_revision_id`, `.message_storage_policy.enforce_in_transit`, `.satisfies_pzs`, `.state`, `.ingestion_data_source_settings`, and `.message_transforms`) that are not yet exposed or triaged by KCC.
- We marked these fields as unimplemented using `f.Unimplemented_NotYetTriaged` to satisfy the round-trip fuzz test.
