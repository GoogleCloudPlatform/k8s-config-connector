### [2026-07-28] AgentRegistryService Implementation Quirks
- **Context**: Implementing direct types and identity for `AgentRegistryService` (PR #12015).
- **Problem**: 
    1. The pinned Google APIs SHA in the issue (`731d7f2ab6` from Aug 2025) was too old and did not contain the `AgentRegistryService` proto, which was added in June 2026.
    2. `controllerbuilder` prunes types it thinks are unreachable if the Spec/Status structs are not manually populated first.
    3. CAIS identity tests (`TestGoldenIdentitiesYamlFiles`) return `caisURL: unknown` for new resources unless their identity package is imported in `pkg/cli/powertools/cais/cmd_test.go`.
- **Solution**:
    1. Updated `apis/git.versions` to a more recent SHA (`62da249e76` from July 2026).
    2. Manually populated `AgentRegistryServiceSpec` and `AgentRegistryServiceObservedState` with fields from the proto and ran `generate.sh` again.
    3. Added `_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/agentregistry/v1alpha1"` to `pkg/cli/powertools/cais/cmd_test.go` and `pkg/gcpurls/registry_test.go`.
- **Impact**: Future agents working on `AgentRegistry` should ensure they use a sufficiently recent `googleapis` SHA and remember to update CAIS/gcpurls test imports.
