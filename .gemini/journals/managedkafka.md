# ManagedKafka Journal

### [2026-05-22] ManagedKafkaConsumerGroup Implementation
- **Context**: Implementing direct controller for ManagedKafkaConsumerGroup (Phase 2).
- **Problem**: ManagedKafkaConsumerGroup in GCP is typically created automatically by Kafka consumers. There is no `CreateConsumerGroup` method in the REST API (only `Get`, `Update`, `Delete`, `List`). However, `Update` (patch) can be used to manage offsets.
- **Solution**: 
    - Moved `topics` field from `ObservedState` to `Spec` in KRM to allow users to manage offsets via Config Connector.
    - Implemented `UpdateConsumerGroup` as an "upsert" in MockGCP to facilitate testing.
    - Added normalization for `operations` in `cloudresourcemanager` to avoid golden file diffs when project dependencies are created.
    - Migrated `ManagedKafkaConsumerGroup`, `Cluster`, and `Topic` to `IdentityV2` and `refs.Ref` pattern for consistency.
- **Impact**: Enables management of Kafka consumer group offsets via KCC. Future implementations of similar resources should check if `patch` acts as upsert or if a separate creation mechanism is needed.
