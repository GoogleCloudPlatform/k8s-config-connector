# Fuzzing PubSubSubscription

We successfully implemented and verified the round-trip fuzzer for `PubSubSubscription` under the central fuzzer suite.

## Observations & Key Learnings

### 1. Handling Near Miss Casing Mismatches
When generating mappers via the `generate-mapper` tool, fields with acronyms or slight casing mismatches (such as `TTL` vs `Ttl` on `ExpirationPolicy`, or `OIDCToken` vs `OidcToken` on `PushConfig`) may not be automatically matched and will be listed as `MISSING` with a `(near miss):` comment:
- `// (near miss): "TTL" vs "Ttl"`
- `// (near miss): "OIDCToken" vs "OidcToken"`

To resolve this, we:
- Hand-coded the mappings for both `ExpirationPolicy` and `PushConfig` in `pkg/controller/direct/pubsub/subscription_mapper.go`.
- Ran the localized generator script (`apis/pubsub/v1beta1/generate.sh`) to automatically pick up these manual overrides. The generator successfully skipped generating those functions in `mapper.generated.go` and commented out the old ones.

### 2. Spec-only Fuzzers
Since `PubSubSubscription` does not have any direct observed state fields (other than standard conditions/observedGeneration), it only requires spec fuzzing. The fuzzer is registered via `fuzztesting.RegisterKRMSpecFuzzer` and `fuzztesting.NewKRMTypedSpecFuzzer` in `pkg/controller/direct/pubsub/subscription_fuzzer.go` (renamed from the non-idiomatic `pubsubsubscription_fuzzer.go`).
