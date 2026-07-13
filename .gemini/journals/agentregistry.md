### [2026-07-06] AgentRegistry Google APIs Pin Update
- **Context**: Implementing the initial KRM types and IdentityV2 for `AgentRegistryBinding`.
- **Problem**: The original googleapis commit pin in `apis/git.versions` did not contain the `agentregistry` proto files, as the service was launched in 2026 whereas the old pin was from 2025.
- **Solution**: Updated `apis/git.versions` to use `2b625c91510a2e8320a778bc88af8b65bc4a19a2` (the Google APIs commit from July 6, 2026) to pull in the `agentregistry` proto files, allowing the code generator to successfully run.
- **Impact**: All future code generators and validation scripts will now have access to the `agentregistry` proto files.
