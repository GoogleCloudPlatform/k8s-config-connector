### [2026-07-06] AgentRegistry Google APIs Pin Update
- **Context**: Implementing the initial KRM types and IdentityV2 for `AgentRegistryBinding`.
- **Problem**: The original googleapis commit pin in `apis/git.versions` did not contain the `agentregistry` proto files, as the service was launched in 2026 whereas the old pin was from 2025.
- **Solution**: Updated `apis/git.versions` to use `2b625c91510a2e8320a778bc88af8b65bc4a19a2` (the Google APIs commit from July 6, 2026) to pull in the `agentregistry` proto files, allowing the code generator to successfully run.
- **Impact**: All future code generators and validation scripts will now have access to the `agentregistry` proto files.

### [2026-09-16] AgentRegistryService Greenfield Types & Identity Scaffold
- **Context**: Greenfield Phase 1 types & identity scaffolding for AgentRegistryService (Service kind in `google.cloud.agentregistry.v1`). Link to issue #11475.
- **Problem**: The default pinned googleapis SHA in KCC's `apis/git.versions` (`1765b559c4`) dates to August 2025 and does not contain `google/cloud/agentregistry`. Hence, running `./apis/agentregistry/generate.sh` with the default version causes protoc/controllerbuilder to fail with: `failed to find the proto message google.cloud.agentregistry.v1.Service: proto: not found`.
- **Solution**: We targeted googleapis commit `5e3f9202b4` (which introduced the v1 Client Libraries for Agent Registry) inside `generate.sh` and executed a full clone of googleapis (bypassing shallow `--depth 1` limits) to successfully checkout the commit and compile the needed proto files. All generated types were manually mapped in `agentregistryservice_types.go`, removing unreachable annotations in `types.generated.go`.
- **Impact**: Future greenfield implementations for newer Google Cloud APIs can utilize this method of targeting specific Google APIs commits via `generate.sh` to ensure compatibility before those services are incorporated in the main pinned `git.versions` SHA of KCC.
