---
name: Greenfield MockGCP Alignment
description: Identifies greenfield resources that have completed Phase 2 (controller and E2E fixtures) but lack a high fidelity MockGCP implementation or alignment, and creates tracking issues for them.
schedule: "never"
skipPR: true
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
You are a software development assistant for the Kubernetes Config Connector project.
You have access to the GitHub CLI (`gh`) and bash tools.

# Task

Your goal is to identify greenfield resources that have completed Phase 2 (Controller & E2E) but are missing high fidelity MockGCP implementations or alignment (Phase 3), and create issues to track this work.

## 1. Identify Target Resources
Consult the dashboard at `hack/tools/greenfield/RESOURCE_STATUS.md` and identify resources where `Current Phase` is `2`, the `State` is `MERGED`, and the `Phase 3 PR` column contains `-` (indicating MockGCP alignment is not yet completed).

## 2. Check Issue Limits
Check how many issues related to direct mockgcp alignment have been opened today. To avoid overwhelming the team, **do not open more than 5 issues per day**. If the limit is reached, stop creating new issues.

## 3. Avoid Duplicates
Before creating an issue for a resource, search existing open and closed issues using `gh search issues` to ensure an issue for this resource hasn't already been created. Search for the exact resource name in the title, e.g., `Implement direct mockgcp for <ResourceKind>`.

## 4. Create Minimal Issues
For each identified resource, create a new issue using `gh`.

## 5. Audit
- If you have opened an issue or multiple issues, comment the links, in a single GitHub comment, to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 3 (MockGCP Alignment)` chore.
- If you have determined that you don't need to open an issue, make a single comment to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439 as to why you are not opening an issue this run. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 3 (MockGCP Alignment)` chore.

### Issue Title
`Implement direct mockgcp for <ResourceKind>`

*(Replace `<ResourceKind>` with the actual Kind of the resource, e.g., `VertexAIDataset`)*

### Issue Labels
Add the following labels to the created issue:
- `overseer`
- `area/direct`
- `priority/medium`
- `step/mockgcp`
- `greenfield`
- `chore/ai`

### Issue Body
Use the following exact template for the issue body, replacing the bracketed placeholders (`<ResourceKind>`, `<service_name>`, `<api_version>`, `<proto_package>`, `<kind_lowercase>`, `<group>`) with the actual values for the resource:

------------ BEGIN ISSUE BODY TEMPLATE ------------
# Role
You are a software developer for the Config Connector (KCC) project.
Your task is to implement Phase 3 (MockGCP and Alignment) for the `<ResourceKind>` resource using the "direct" approach.

# Implementation Instructions

1. **Implement and Align MockGCP**: Use skill `.gemini/skills/kcc-direct-mockgcp-implementer/SKILL.md` with:
   - ResourceKind: <ResourceKind>
   - service_name: <service_name>
   - group: <group>
   - api_version: <api_version>
   - kind_lowercase: <kind_lowercase>
   - testname: The specific test folder name (e.g. `basic`) under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/`

2. **Journal Findings**: Use skill `.gemini/skills/kcc-agentic-journaler/SKILL.md` to record service-specific findings or quirks in `.gemini/journals/<service_name>.md`.

3. **Create PR**:
   - Create a Pull Request with the mockgcp implementation and aligned golden logs.
   - Use the `send-pr` skill.
   - Pass the issue labels to the `send-pr.sh` script using the `--labels` flag (e.g. `--labels "overseer,area/direct,priority/medium,step/mockgcp,greenfield,chore/ai"`).
   - Apply the same labels to the PR as are on this issue.
   - Include a link to the chore file (`.agents/greenfield-mockgcp.md`) in the PR description.
   - Include `Fixes #<issue-number>` in the PR description.
------------ END ISSUE BODY TEMPLATE ------------
