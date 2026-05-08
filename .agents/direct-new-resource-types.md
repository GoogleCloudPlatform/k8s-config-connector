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
    - **Labels**: `overseer`, `area/direct`, `priority/medium`, `step/gen-types`, `greenfield`
    - **Body**: Use the **TYPES ISSUE BODY TEMPLATE** below. Append a link to this chore file (`.agents/direct-new-resource-types.md`) at the end of the issue body for traceability.

---

## TYPES ISSUE BODY TEMPLATE
# Role
You are a software developer for the Config Connector (KCC) project.
Your task is to implement the initial KRM types, CRD, and IdentityV2 for the `<Kind>` resource using the "direct" approach.

# Context: Implementation Versions
To ensure stability and reproducibility, this task is pinned to the following repository versions:
# TODO: Dynamically determine these SHAs during task generation (see PR #7946 feedback).
- **Google APIs SHA**: `731d7f2ab6` (from `apis/git.versions`)
- **KCC Base SHA**: `dc1dd45d0b`

# Implementation Instructions

### Phase 0: Tooling & Skill Activation

1.  **Activate Skills**: This project uses specialized skills to ensure architectural consistency. Before starting, identify and activate relevant skills (e.g., `kcc-direct-types-implementer`, `kcc-identity-reference`, `kcc-agentic-journaler`) using the `activate_skill` tool to receive expert procedural guidance.
2.  **Consult Knowledge Base**: Check the `.gemini/journals/` directory for any existing service-specific "tribal knowledge" that might apply to this resource.
3.  **Record Versions**: Record the Google APIs and KCC SHAs (listed above) in your initial implementation notes to ensure a stable baseline.

### Phase 1: Types and CRDs
Implement the initial KRM types for `<Kind>` using the "direct" approach.

**Tooling**: Activate the `kcc-direct-types-implementer` skill. Follow its guidance to ensure your `_types.go` file meets KCC's metadata and field-mapping standards.

### Phase 2: IdentityV2 and Resource References
Implement `apis/<group>/v1alpha1/<kind_lowercase>_identity.go`.

**Tooling**: Activate the `kcc-identity-reference` skill. It provides the canonical `gcpurls.Template` pattern and ensures interface compliance for `IdentityV2` and `ExternalIdentifier`.

### Phase 3: Final Generation and Verification
1.  **Generate Mappers**: Run `dev/tasks/generate-types-and-mappers`.
2.  **Compile**: Run `make all-binary` to ensure the generated code compiles. Fix any issues discovered.

---

### Phase 4: Knowledge Capture & Self-Optimization
Reflect on your implementation and record any breakthroughs or service-specific quirks.

**Tooling**: Activate the `kcc-agentic-journaler` skill. Use its routing logic to ensure your learnings are stored in the correct destination (either a general skill or a service-specific journal).
