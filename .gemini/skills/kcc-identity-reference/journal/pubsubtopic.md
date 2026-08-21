# PubSubTopic Identity and Reference Transition Journal

## Observations

1. **Controller Type**: `PubSubTopic` is managed by the legacy/Terraform-based controller, so it does not have a direct Go controller. However, its identity and reference structures can still be updated to the canonical `IdentityV2` / `refs.Ref` patterns, facilitating clean reference resolution for other resources.
2. **Missing Status Fields**: `PubSubTopicStatus` lacks `externalRef` or `name` fields (it only contains `conditions` and `observedGeneration`). Therefore, the `GetIdentity` implementation does not and must not perform any status cross-checks.
3. **Generation Cleanup**: The old `TopicIdentity` generated deepcopy code in `zz_generated.deepcopy.go` was obsolete once we renamed the struct to `PubSubTopicIdentity`. Running `dev/tasks/generate-types-and-mappers` cleanly removed the old deepcopy functions and updated all generated code.
4. **No CRD Schema Changes**: Since we did not change any API types or schema definitions, running `dev/tasks/diff-crds` confirmed that the Custom Resource Definition schemas remained completely unchanged.
