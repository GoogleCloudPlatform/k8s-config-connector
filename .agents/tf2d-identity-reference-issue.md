---
name: Issue for Identity and Reference
description: Identify resources that do not implement the identity and reference pattern and create minimal issues for them.
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

Your goal is to identify resources that need to be migrated to the identity and reference pattern, and to create minimal issues to track this work.

## 1. Identify Target Resources
Identify resources that have a `_types.go` (or `types.go`) file in their `apis/<group>/<version>/` directory but do not yet have an `_identity.go` file. You can find these by inspecting the filesystem or by looking for recently closed issues labelled with `area/direct` and `step/gen-types`.

## 2. Check Issue Limits
Check how many issues related to identity and reference have been opened today. To avoid overwhelming the team, **do not open more than 10 issues per day**. If the limit is reached, stop creating new issues.

## 3. Avoid Duplicates
Before creating an issue for a resource (e.g., `VertexAIDeploymentResourcePool`), search existing open and closed issues using `gh search issues` to ensure an issue for this resource hasn't already been created. Search for the exact resource name in the title.

## 4. Create Minimal Issues
For each identified resource, create a new issue using `gh`.

### Issue Title
`Move <ResourceKind> to identity and refs pattern`

*(Replace `<ResourceKind>` with the actual Kind of the resource, e.g., `VertexAIDeploymentResourcePool`)*

### Issue Labels
Add the following labels to the created issue:
- `overseer`
- `area/direct`
- `priority/medium`
- `step/identity-reference`

### Issue Body
Use the following exact template for the issue body, replacing `<ResourceKind>` with the actual Kind of the resource:

------------ BEGIN ISSUE BODY TEMPLATE ------------
Please follow the skill .gemini/skills/kcc-identity-reference/SKILL.md for <ResourceKind>

If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-identity-reference/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
------------ END ISSUE BODY TEMPLATE ------------
