# Journal Entry - PubSubTopic MessageStoragePolicy.enforceInTransit

## Overview
Added support for the `enforceInTransit` parameter under `messageStoragePolicy` in `PubSubTopic` to both legacy (Terraform) and direct controllers.

## Details
- Patched the local copy of the Terraform Google Beta Provider under `third_party/` to expose `enforce_in_transit` under `message_storage_policy` in `google_pubsub_topic` and implemented the flattener and expander.
- Updated KRM types in `apis/pubsub/v1beta1/topic_types.go` to add `EnforceInTransit *bool` to the `MessageStoragePolicy` struct.
- Updated direct controller mapping in `pkg/controller/direct/pubsub/topic_mapper.go` to proxy the `EnforceInTransit` value to the GCP Pub/Sub API.
- Updated fuzz test configuration `pkg/controller/direct/pubsub/topic_fuzzer.go` to triaging the field.
- Regenerated CRDs, client code, and deepcopy implementations.
- Expanded the E2E test fixture under `pkg/test/resourcefixture/testdata/basic/pubsub/v1beta1/pubsubtopic/pubsubtopic/` to verify setting and changing `enforceInTransit` during create and update.
