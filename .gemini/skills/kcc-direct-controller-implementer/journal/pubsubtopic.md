# PubSubTopic Direct Controller Implementation Journal

## Observations
1. **Client Name Discrepancy**: Unlike what one might expect based on typical services (e.g. `PublisherClient`), the low-level client in GCP Pub/Sub Go SDK (`cloud.google.com/go/pubsub/v2/apiv1`) used for managing topics is actually `TopicAdminClient`.
2. **Mock GCP Integration**: The mock GCP implementation for Pub/Sub topic update (`mockpubsub/topic.go`) initially only supported updating `labels`, `schema_settings`, and `message_retention_duration`. To prevent potential test failures when testing `Update` reconciliation paths with other fields, we extended `mockpubsub/topic.go` to support `message_storage_policy` and `kms_key_name`.
3. **No External Status Fields**: KRM `PubSubTopicStatus` is very simple and only has `observedGeneration` and `conditions`, so no external identifier or complex observed fields needed to be mapped from the proto structure.
