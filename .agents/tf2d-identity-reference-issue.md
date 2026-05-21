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

# Filter Criteria
Find resources that meet the following criteria:
1. **Prioritization**: Priority is given to resources with no dependencies or those that are dependencies for many other unmigrated resources (Topological Order).
2. **Absence of Task**: No GitHub issue (open or closed) should already exist for creating the identity/reference for this specific Group and Kind.

# Task

Your goal is to identify resources that need to be migrated to the identity and reference pattern, and to create minimal issues to track this work.

## 1. Throttle Check
   - Count open issues labeled with `overseer` and `step/identity-reference`. If the count is 10 or more, STOP and log: "Throttle limit reached (10). Skipping issue generation."
   - Check how many issues related to identity and reference have been opened today. To avoid overwhelming the team, **do not open more than 10 issues per day**. If the limit is reached, stop creating new issues.

## 2. Identify Candidate
   - Run `./hack/directmigration/find-missing-identity-reference.py --output candidates.txt`.
   - Run `./hack/directmigration/topological-sort.py candidates.txt` to get the implementation order.
   - Iterate through the sorted list, focusing only on the resources that were identified in `candidates.txt`.

## 3. Verify and De-duplicate
- For each prioritized candidate:
    - Search existing open and closed issues using `gh search issues` to ensure an issue for this resource hasn't already been created. Search for the exact resource name in the title. The issue title could be one of "Create Identity and Reference files for <group> <Kind>" and "Move <Kind> to identity and refs pattern".
        - If an **Open** issue exists: Skip to the next candidate.
        - If a **Closed** issue exists: Check its close reason. If closed as `not planned`, skip to the next candidate. If closed as `completed` (and the file is missing), proceed to recreate the issue.
    - Pick the first candidate that passes both checks.

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
Use the following exact template for the issue body, replacing `<ResourceKind>` with the actual Kind of the resource, and apply the same labels to the PR as are on this issue:

------------ BEGIN ISSUE BODY TEMPLATE ------------
Please follow the skill .gemini/skills/kcc-identity-reference/SKILL.md for <ResourceKind>

If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-identity-reference/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
------------ END ISSUE BODY TEMPLATE ------------
