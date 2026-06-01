---
name: Greenfield New Resource Types
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
      python3 hack/tools/greenfield/calculate_coverage.py 731d7f2ab6 dc1dd45d0b 5
      ```

2.  **Identify Candidate**:
    - Look at the "Next 5 Easiest Resources to Implement" section.
    - Pick the first resource that is not already tracked by an issue.
3.  **Verify Necessity**:
    - Check if an implementation issue already exists (open or closed):
      ```bash
      gh issue list --state all --search "ai:chore: Implement direct types for: <Kind>"
      ```

      Do nothing if there are more than 5 outstanding open issues AND if you have opened an issue today.

4.  **Task**: If no tracking issue exists, create a new issue for the resource types implementation.
    - **Title**: `ai:chore: Implement direct types for: <Kind>`
    - **Labels**: `overseer`, `area/direct`, `priority/medium`, `step/gen-types`, `greenfield`, `chore/ai`
    - **Body**: Use the **TYPES ISSUE BODY TEMPLATE** below. Append a link to this chore file (`.agents/greenfield-direct-new-resource-types.md`) at the end of the issue body for traceability.
<<<<<<< ours

5. **Audit**:
   - If you have opened an issue or multiple issues, comment the links, in a single GitHub comment, to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 1 (Types & Identity)` chore.
   - If you have determined that you don't need to open an issue, make a single comment to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439 as to why you are not opening an issue this run. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 1 (Types & Identity)` chore.

---

## TYPES ISSUE BODY TEMPLATE
# Role
You are a software developer for the Config Connector (KCC) project.
Your task is to implement the initial KRM types, CRD, and IdentityV2 for the `<Kind>` resource using the "direct" approach.

# Context: Implementation Versions
To ensure stability and reproducibility, this task is pinned to the following repository versions:
<!-- TODO: Dynamically determine these SHAs during task generation (see PR #7946 feedback). -->
- **Google APIs SHA**: `731d7f2ab6` (from `apis/git.versions`)
- **KCC Base SHA**: `dc1dd45d0b`

# Implementation Instructions

1. **Scaffold types**: Use skill `.gemini/skills/kcc-direct-resource-scaffolder/SKILL.md` with:
   - service: <proto.package.name>
   - resource: <Kind>:<ProtoMessageName>
   - api_version: <group>.cnrm.cloud.google.com/v1alpha1

2. **Implement Identity**: Use skill `.gemini/skills/kcc-direct-identity-implementer/SKILL.md` with:
   - resource_kind: <Kind>
   - template: <GCP_URL_Template> (e.g. projects/{project}/locations/{location}/<plural>/{<resource>})

3. **Run local validations**:
   - run `scripts/validate-prereqs.sh`, if the script fails the output will contain agent hints for fixing the errors.
   - run `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`

4. **Run local validations**:
   - run `scripts/validate-prereqs.sh`, if the script fails the output will contain agent hints for fixing the errors.
   - run `./dev/ci/presubmits/tests-e2e-fixtures-<kind_lowercase>`

5. **Journal Findings**: Use skill `.gemini/skills/kcc-agentic-journaler/SKILL.md` to capture quirks and update knowledge.

6. **Create PR**:
   - create a Pull Request with your changes.
   - apply the same labels to the PR as are on this issue.
   - include a link to the chore file (`.agents/greenfield-direct-new-resource-types.md`) in the PR description.
   - include `Fixes #<issue-number>` in the PR description.

5. **Create PR**:
   - create a Pull Request with your changes using the `send-pr` skill.
   - make sure the Pull Request does not contain the controller or mapper steps.
   - pass the issue labels to the `send-pr.sh` script using the `--labels` flag (e.g. `--labels "overseer,area/direct,priority/medium,step/gen-types,greenfield,chore/ai"`).
   - include a link to the chore file (`.agents/greenfield-direct-new-resource-types.md`) in the PR description.
   - include `Fixes #<issue-number>` in the PR description.

---
