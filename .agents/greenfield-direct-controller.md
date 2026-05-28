---
name: Greenfield Direct Controller
description: Identifies greenfield resources that have completed Phase 1 (types and identity) but lack a controller, and creates minimal issues for them.
schedule: "@daily"
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

Your goal is to identify greenfield resources that have their KRM types and IdentityV2 implemented (Phase 1 Complete), but are missing their direct controller and E2E fixtures (Phase 2), and create minimal issues to track this work.

## 1. Identify Target Resources
Identify resources that have an `_identity.go` file in their `apis/<group>/<version>/` directory but do not yet have an E2E fixture directory in `pkg/test/resourcefixture/testdata/basic/<group>/<version>/<kind_lowercase>/`. This ensures we catch resources that might have a controller stubbed out but lack the required test coverage.

Consult the dashboard at `hack/tools/greenfield/RESOURCE_STATUS.md` and look for resources where the `Current Phase` is `1` and the `State` is `MERGED`. Prioritize resources that are ready for Phase 2 from this dashboard.

## 2. Check Issue Limits
Check how many issues related to direct controllers have been opened today. To avoid overwhelming the team, **do not open more than 5 issues per day**. If the limit is reached, stop creating new issues.

## 3. Avoid Duplicates
Before creating an issue for a resource (e.g., `VertexAIDataset`), search existing open and closed issues using `gh search issues` to ensure an issue for this resource hasn't already been created. Search for the exact resource name in the title, e.g., `Implement direct controller for VertexAIDataset`.

## 4. Create Minimal Issues
For each identified resource, create a new issue using `gh`.

## 5. Audit
- If you have opened an issue or multiple issues, comment the links, in a single GitHub comment, to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 2 (Controller & E2E)` chore.
- If you have determined that you don't need to open an issue, make a single comment to the tracking issue https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8439 as to why you are not opening an issue this run. **IMPORTANT**: Start your comment by stating you are acting on behalf of the `Greenfield Phase 2 (Controller & E2E)` chore.

### Issue Title
`Implement direct controller for <ResourceKind>`

*(Replace `<ResourceKind>` with the actual Kind of the resource, e.g., `VertexAIDataset`)*

### Issue Labels
Add the following labels to the created issue:
- `overseer`
- `area/direct`
- `priority/medium`
- `step/controller`
- `greenfield`
- `chore/ai`

### Issue Body
Use the following exact template for the issue body, replacing the bracketed placeholders (`<ResourceKind>`, `<service_name>`, `<api_version>`, `<proto_package>`) with the actual values for the resource:

------------ BEGIN ISSUE BODY TEMPLATE ------------
Please implement the Phase 2 (Controller and E2E Tests) for `<ResourceKind>` by following these two skills:

1. **Controller Scaffolding & Mappers**: Follow `.gemini/skills/kcc-direct-controller-implementer/SKILL.md`
   - `resource_kind`: `<ResourceKind>`
   - `package_path`: `pkg/controller/direct/<service_name>/<kind_lowercase>/`
   - `proto_package`: `<proto_package>`

2. **Controller Logic & E2E Fixtures**: Follow `.gemini/skills/kcc-direct-controller-logic-implementer/SKILL.md`
   - `resource_kind`: `<ResourceKind>`
   - `service_name`: `<service_name>`
   - `api_version`: `<api_version>`

3. **Verify/Add Identity Tests**: 
   - Check if `apis/<group>/<api_version>/<kind_lowercase>_identity_test.go` exists.
   - If it is missing (or incomplete), implement it now to ensure full coverage of the IdentityV2 interface.

If you find any shortcomings in the skills, you may update their respective `SKILL.md` files. For service-specific observations, MockGCP alignment issues, or GCP SDK quirks, use `.gemini/journals/<service_name>.md` as described in the `kcc-agentic-journaler` skill to avoid merge conflicts.

4. **Create PR**:
   - create a Pull Request with your changes using the `send-pr` skill.
   - pass the issue labels to the `send-pr.sh` script using the `--labels` flag (e.g. `--labels "overseer,area/direct,priority/medium,step/gen-types,greenfield,chore/ai"`).
   - apply the same labels to the PR as are on this issue.
   - include a link to the chore file (`.agents/greenfield-direct-controller.md`) in the PR description.
   - include `Fixes #<issue-number>` in the PR description.
------------ END ISSUE BODY TEMPLATE ------------
