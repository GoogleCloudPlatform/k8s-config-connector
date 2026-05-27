### [2026-05-27] NetworkSecuritySACRealm Hallucination
- **Context**: Implementing initial types for NetworkSecuritySACRealm (Issue #8736).
- **Problem**: The task requested implementation for `NetworkSecuritySACRealm:SACRealm`, but the protobuf message `SACRealm` and the API endpoint `sacRealms` do not exist in the pinned `googleapis` repository (SHA `731d7f2ab6`) under `google.cloud.networksecurity.v1` or any other version. The script `dev/ci/presubmits/tests-e2e-fixtures-networksecuritysacrealm` also does not exist.
- **Solution**: The issue was identified as an AI hallucination by the greenfield issue creator agent. I aborted the resource implementation, documented this finding, and removed the scaffolded `generate.sh`. This PR serves to capture the finding and close the hallucinated issue.
- **Impact**: Agents should verify the existence of a protobuf message in `googleapis.pb` before attempting to scaffold direct types, as the issue queue might contain hallucinated targets.
