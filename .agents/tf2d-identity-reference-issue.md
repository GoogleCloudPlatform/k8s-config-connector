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
Your goal is to identify resources missing the identity and reference pattern and create minimal issues for backfilling them.

1.  **Identify Candidates**: Find resources in `apis/` that have a `_types.go` file but no corresponding `_identity.go` file. You can also identify them by looking for recently closed issues labelled `area/direct` and `step/gen-types`.
2.  **Check Limits**: Do not open more than 10 issues per day.
3.  **Avoid Duplicates**: Search existing issues (open or closed) to ensure no duplicate exists for the resource.
4.  **Create Issue**: Create a new issue for each identified resource using the template below.

## Issue Template
- **Title**: `Move <ResourceKind> to identity and refs pattern`
- **Labels**: `overseer`, `area/direct`, `priority/medium`, `step/identity-reference`
- **Body**:
------------ BEGIN ISSUE BODY TEMPLATE ------------
Please follow the skill .gemini/skills/kcc-identity-reference/SKILL.md for <ResourceKind>

If you find any shortcomings in the skill (that likely apply to other resources), you may update SKILL.md. Also keep a journal of any less general observations etc. To avoid git merge conflicts, use a file under .gemini/skills/kcc-identity-reference/journal/, named after the kind or a similarly unique name. You may grep journal entries to identify learnings from other resources; if you find an important pattern by doing that you may also update the SKILL.md itself.
------------ END ISSUE BODY TEMPLATE ------------
