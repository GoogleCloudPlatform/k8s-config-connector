---
name: Mock Drift Correction
description: Detects mockgcp services that are still using the obsolete grpc-gateway (httpmux) approach and creates issues to migrate them to httptogrpc.
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
Your goal is to identify mockgcp services that are still using the old, proto-generation-based `grpc-gateway` approach and create GitHub issues to track their migration to the modern `httptogrpc` approach. 

This migration simplifies maintenance because `httptogrpc` uses reflection against the official Google Cloud Go SDK, eliminating the need to manually manage proto files and address "proto drift" through local generation.

# Task
1. Navigate to the `mockgcp` directory: `cd mockgcp`
2. Identify the list of mock services that are NOT using `httptogrpc`. 
   A robust way to check is to find `mock*/service.go` files that do NOT import `httptogrpc`.
   Run this command:
   `grep -L "httptogrpc" mock*/service.go | cut -d/ -f1 | sed 's/^mock//'`
   Identify the corresponding `<service>` names.
3. If the number of identified services to migrate is more than 10, limit your work to the first 10 to avoid overwhelming the team.
4. For each of the identified services:
    - Check if an issue already exists (open or closed) for migrating this service using: `gh issue list --state all --search "Move mockgcp <service> away from grpc-gateway"`.
    - If an issue already exists, skip creating a new one.
    - If no issue exists, create a new GitHub issue using the `gh` tool.
    - The issue title should be: `Move mockgcp <service> away from grpc-gateway`
    - The issue should be labeled with: `overseer`, `priority/medium`, `step/mockgcp`.
    - The issue body MUST contain the exact text from the **ISSUE BODY TEMPLATE** below, replacing `<service>` with the appropriate service name.
    - Append a link to this chore file (`.agents/mock-drift-correction.md`) at the end of the issue body for traceability.

## Issue Title

`Move mockgcp <service> away from grpc-gateway`

## Issue Labels
The issue should be labeled with the following labels:
* `overseer`
* `priority/medium`
* `step/mockgcp`

## ISSUE BODY TEMPLATE

The issue body should contain this text template with the appropriate service filled in.
The body template is treated as markdown. Retain the formatting as is when filling in the service.

------------ BEGIN ISSUE BODY TEMPLATE ------------
Use the skill mockgcp/.gemini/skills/move-away-from-grpc-gateway/skill.md to move mockgcp <service> to http://cloud.google.com/go/<service> and to httptogrpc

If there are additional complexities, please create/update the skill .gemini/skills/move-away-from-grpc-gateway/skill.md . Append to a journal.md alongside (mockgcp/.gemini/skills/move-away-from-grpc-gateway/journal.md) with any less important notes for future aggregation. Do not overwrite any previous notes in the journal, but if you see a pattern you can promote it to the skill.

------------ END ISSUE BODY TEMPLATE ------------
