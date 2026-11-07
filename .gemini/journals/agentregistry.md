### [2026-07-06] AgentRegistry Google APIs Pin Update
- **Context**: Implementing the initial KRM types and IdentityV2 for `AgentRegistryBinding`.
- **Problem**: The original googleapis commit pin in `apis/git.versions` did not contain the `agentregistry` proto files, as the service was launched in 2026 whereas the old pin was from 2025.
- **Solution**: Updated `apis/git.versions` to use `2b625c91510a2e8320a778bc88af8b65bc4a19a2` (the Google APIs commit from July 6, 2026) to pull in the `agentregistry` proto files, allowing the code generator to successfully run.
- **Impact**: All future code generators and validation scripts will now have access to the `agentregistry` proto files.

### [2026-11-06] Greenfield Step 1a: Prepopulate direct KRM types, identity, and judgement queue for AgentRegistryService
- **Context**: Prepopulating direct KRM types, identity, reference, and `judgement_queue.yaml` for `AgentRegistryService` under `agentregistry.cnrm.cloud.google.com/v1alpha1` (Issue #15175).
- **Details**:
  1. Updated `apis/agentregistry/generate.sh` to add `generate-types --prepopulate-resource`, `generate-identity`, and `generate-mapper --prepopulate-resource` for `AgentRegistryService:Service`.
  2. Executed `./apis/agentregistry/generate.sh` to generate `agentregistryservice_types.go`, `agentregistryservice_identity.generated.go`, `agentregistryservice_reference.generated.go`, `agentregistryservice_identity_generated_test.go`, and `judgement_queue.yaml`.
  3. Added `AgentRegistryService` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Verified all unit tests and `apichecks` pass, updating `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Establishes initial scaffolding and judgement queue for `AgentRegistryService` ready for Step 1b triage.

