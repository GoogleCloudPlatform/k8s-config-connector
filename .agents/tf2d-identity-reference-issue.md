---
name: Issue for Identity and Reference
description: Identify resources missing the identity and reference pattern and create minimal issues for them.
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
Identify resources that need to be migrated to the identity and reference pattern and create minimal tracking issues.

## 1. Identify Target Resources
- Run `./hack/directmigration/find-missing-identity-reference.py --output candidates.txt`.
- Run `./hack/directmigration/topological-sort.py candidates.txt` to get the implementation order.
- Iterate through the sorted list, focusing only on the resources that were identified in `candidates.txt`.

## 2. Constraints
- **Limit:** Max 20 issues per day. Count open issues labeled with `overseer` and `step/identity-reference`. If the count is 20 or more, STOP.
- **Avoid Duplicates:** Search `gh search issues` for the resource Kind in the title (open or closed) before creating.

## 3. Create Issues
For each identified resource, create a new issue:

- **Title:** `Move <Kind> to identity and refs pattern`
- **Labels:** `overseer`, `area/direct`, `direct-migration`, `priority/medium`, `step/identity-reference`
- **Body:** Use the following template, replacing `<Kind>` with the resource Kind:

------------ BEGIN ISSUE BODY TEMPLATE ------------
Please follow the skill .gemini/skills/kcc-identity-reference/SKILL.md for <Kind>

If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-identity-reference/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
------------ END ISSUE BODY TEMPLATE ------------
