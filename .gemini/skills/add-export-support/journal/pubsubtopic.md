# Export Support Journal for PubSubTopic

## Observations
- Implementing `AdapterForURL` for `PubSubTopic` was extremely straightforward as the `PubSubTopicIdentity` natively supports parsing full CAI-style URLs (including standard host prefixes like `//pubsub.googleapis.com/`).
- `tests/e2e/export.go` handles standard resources natively with CAIS resolution. Adding the `PubSubTopic` mapping using `resolveCAISURI` worked out-of-the-box because `PubSubTopic` correctly implements `identity.Resource` and `PubSubTopicIdentity` implements `identity.IdentityV2`.
- When running the E2E unified tests locally against MockGCP, setting `WRITE_GOLDEN_OUTPUT=1` on subsequent runs might pollute the recorded HTTP logs with dynamic `uniqueId` values. Restoring the HTTP log files and running without `WRITE_GOLDEN_OUTPUT=1` ensures tests pass locally while preserving the golden HTTP assets recorded from real GCP.
