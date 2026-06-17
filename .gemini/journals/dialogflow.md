### [2026-05-26] DialogflowConversationDataset Scaffold
- **Context**: Implementing Greenfield step 1 for DialogflowConversationDataset.
- **Problem**: The proto mapping generated unreachable types due to lack of usages in the scaffolded struct.
- **Solution**: Manually copied fields like `InputConfig` and `ConversationInfo` to the `ConversationDatasetObservedState` before re-running generation.
- **Impact**: Enables properly typing output fields in the CRD schema.

