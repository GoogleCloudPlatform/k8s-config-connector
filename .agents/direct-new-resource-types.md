---
name: Direct New Resource Types
description: Identifies missing GCP resources from protos and creates tasks to implement their initial types, CRDs, and IdentityV2.
schedule: "@daily"
---

<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Role
You are a senior cloud architect for the Kubernetes Config Connector project.
Your goal is to identify GCP resources that are missing from KCC but defined in the `googleapis` protos, and create actionable implementation tasks for their initial KRM types and CRDs (Step 1).

# Scan Trigger: Missing Resource Identification
1.  **Run Coverage Analysis**:
    - Run the coverage script using the pinned Google APIs SHA from `apis/git.versions` and the current KCC HEAD:
      ```bash
      # GoogleAPIs SHA: 731d7f2ab6
      # KCC SHA: dc1dd45d0b
      python3 hack/tools/calculate_coverage.py 731d7f2ab6 dc1dd45d0b 5
      ```

2.  **Identify Candidate**:
    - Look at the "Next 5 Easiest Resources to Implement" section.
    - Pick the first resource that is not already tracked by an issue.
3.  **Verify Necessity**:
    - Check if an implementation issue already exists (open or closed):
      ```bash
      gh issue list --state all --search "ai:chore: Implement direct types for: <Kind>"
      ```

      Do nothing if there are more than 5 outstanding open issues.

4.  **Task**: If no tracking issue exists, create a new issue for the resource types implementation.
    - **Title**: `ai:chore: Implement direct types for: <Kind>`
    - **Labels**: `overseer`, `area/direct`, `priority/medium`, `step/gen-types`
    - **Body**: Use the **TYPES ISSUE BODY TEMPLATE** below. Append a link to this chore file (`.agents/direct-new-resource-types.md`) at the end of the issue body for traceability.

---

## TYPES ISSUE BODY TEMPLATE
# Role
You are a software developer for the Config Connector (KCC) project.
Your task is to implement the initial KRM types, CRD, and IdentityV2 for the `<Kind>` resource using the "direct" approach.

# Context: Implementation Versions
To ensure stability and reproducibility, this task is pinned to the following repository versions:
- **Google APIs SHA**: `731d7f2ab6` (from `apis/git.versions`)
- **KCC Base SHA**: `dc1dd45d0b`

# Implementation Instructions

### Phase 1: Types and CRDs

1.  **Add to generate.sh**:
    Modify `apis/<group>/v1alpha1/generate.sh` (create if missing) to include `<Kind>`.
    Example:
    ```bash
    go run . generate-types \
      --service <proto.package.name> \
      --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
      --resource <Kind>:<ProtoMessageName>
    ```

2.  **Generate Scaffolding**:
    Run `apis/<group>/v1alpha1/generate.sh`. This should create `apis/<group>/v1alpha1/<kind_lowercase>_types.go`.

3.  **Copyright Headers**:
    Ensure new files have the correct header: `// Copyright 2026 Google LLC`.

4.  **Labels**:
    Ensure the following annotations are present in the types file:
    ```go
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
    ```

5.  **References**:
    Map proto fields to KCC references (e.g., `ProjectRef`, `NetworkRef`) using the standard `refs.` package and `+kcc:proto` tags.

6.  **Status**:
    Ensure `status.observedGeneration` is an `*int64`.

### Phase 2: IdentityV2 and Resource References

1.  **Implement IdentityV2**:
    Create `apis/<group>/v1alpha1/<kind_lowercase>_identity.go`.
    - Use the `gcpurls.Template` pattern and the `kcc-identity-reference` skill.
    - Ensure it implements `identity.IdentityV2` and `identity.ExternalIdentifier`.
    - Implement `GetIdentity(ctx, reader)` on the resource struct.

2.  **Identity Unit Tests**:
    Create `apis/<group>/v1alpha1/<kind_lowercase>_identity_test.go`.
    - Verify `FromExternal` with full GCP URLs and interface compliance.

### Phase 3: Final Generation and Verification

1.  **Generate Mappers**:
    Run `dev/tasks/generate-types-and-mappers`.
2.  **Compile**:
    Run `make all-binary` to ensure the generated code compiles. Fix any issues discovered.

---

### Phase 4: Knowledge Capture & Self-Optimization

1. **Categorize Your Findings**:
   Reflect on the implementation. Distinguish between:
   - **General Mechanics**: Patterns applicable to *any* KCC resource (e.g., "Efficient proto-to-KRM field mapping").
   - **Domain-Specific Best Practices**: Reusable steps for a specific implementation area (e.g., "Canonical IdentityV2 pattern").
   - **Service-Specific Tribal Knowledge**: Quirks unique to this API (e.g., "<Kind> requires an extra field-mask check").

2. **Route to the Correct Destination**:
   - **Existing Skills**: If you have a breakthrough in an area covered by an existing skill (e.g. `kcc-identity-reference`), update that skill's `SKILL.md` or append to its `journal.md`.
   - **Service Journals**: For service-specific "gotchas" or unusual field mappings, create or append to `.gemini/journals/<service>.md`. Do **not** put these in a general skill folder.
   - **New Skill Discovery**: If you identified a complex, repeatable recipe that isn't yet documented, create a new skill directory in `.gemini/skills/` and use the `skill-creator` tool to initialize it.

3. **Validation**:
   State in your final summary exactly what new knowledge was captured and which file was modified.
