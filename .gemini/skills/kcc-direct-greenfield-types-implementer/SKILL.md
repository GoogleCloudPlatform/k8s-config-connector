---
name: kcc-direct-greenfield-types-implementer
description: Automate the initial scaffolding of a KCC "direct" resource, including CRD types and generation scripts. Use this when starting a new "direct" implementation for a GCP resource.
---

# KCC Direct Greenfield Types Implementer

This skill guides the initial scaffolding of *new* (greenfield) KCC "direct" resources, ensuring standardized CRD generation and adherence to project-wide validation patterns.

## Prerequisites
You **must** also apply the standards from the base skill: `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`.

## Inputs
- `service`: The Google API service name (e.g., `google.cloud.aiplatform.v1`).
- `resource`: The mapping of KCC Kind to GCP Resource (e.g., `VertexAIExampleStore:ExampleStore`).
- `api_version`: The KCC API version (default: `v1alpha1`).

## Workflow

### 1. Add to generate.sh
Locate `apis/<service_short>/generate.sh`. If it doesn't exist, create it following the standard KCC template:
```bash
#!/bin/bash
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

CONTROLLERBUILDER="${CONTROLLERBUILDER:-}"
if [[ -z "${CONTROLLERBUILDER}" ]]; then
  if [[ -x "${REPO_ROOT}/bin/controllerbuilder" ]]; then
    CONTROLLERBUILDER="${REPO_ROOT}/bin/controllerbuilder"
  else
    CONTROLLERBUILDER="go run ${REPO_ROOT}/dev/tools/controllerbuilder"
  fi
fi
source "${REPO_ROOT}/dev/tools/goimports.sh"
cd ${REPO_ROOT}/dev/tools/controllerbuilder
# Note: generate-proto.sh reuses cached .build/googleapis-<SHA>.pb files by default.
# Pass --force (or FORCE_GENERATE_PROTOS=1) to force re-compiling proto descriptors when testing proto edits:
./generate-proto.sh

${CONTROLLERBUILDER} generate-types \
  --service <service> \
  --api-version <group>.cnrm.cloud.google.com/<api_version> \
  --resource <resource>
```

### 2. Generate Types
Set executable permissions and run the `generate.sh` script:
```bash
chmod +x apis/<service_short>/generate.sh
./apis/<service_short>/generate.sh
```

### 3. Validate and Enhance Output
Apply the baseline validations from `kcc-direct-base-types-implementer`, plus these greenfield-specific rules:

- **Stability Level**: Add `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"`.
- **1:1 Kind to Proto Mapping**: Enforce a strict 1:1 relationship between resource Kinds and Proto definitions. Ensure no single Proto is shared by multiple Kinds, and no single Kind points to multiple Protos.
- **Field Validation**: Manually add or verify kubebuilder tags:
  - Use `// +kubebuilder:validation:Required` for fields that are mandatory in the GCP API.
  - Use `// +kubebuilder:validation:Optional` for all other fields.
- **Pointers**: Ensure all Go scalar primitive types (e.g., `string`, `bool`, `int`, etc.) are pointers (e.g., `*string`, `*bool`). Pay special attention to the `Location` field, which must also be a pointer (`*string`).
- **Enums**: 
  - Use `*string` for the Go type of proto enum fields (do NOT use custom wrapped string types).
  - Use `// +kubebuilder:validation:Enum=VALUE1;VALUE2` to provide validation in the CRD while keeping the Go type simple.
- **Exception Files**: Do not add exceptions to any exception files (e.g., in `dev/tools/controllerbuilder/`) other than `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### 4. Opt-in generation flags

`generate-types` can derive a lot more from the proto than it does by default. Every
flag below is **off unless you pass it**, because turning one on for a resource that
already exists can rename fields, move them between spec and status, or tighten the
CRD schema — all breaking changes. A greenfield resource has no users yet, so it is
the right place to turn them on.

Add them to the service's `generate.sh` so the choice is recorded and regeneration
is reproducible. Opt in **one service at a time**.

| Flag | What it does |
|---|---|
| `--prepopulate-spec` | Fills the scaffolded Spec from the proto instead of leaving it empty. This is the main one; most of the others only matter alongside it. |
| `--emit-required-from-proto` | Emits `// +required` for fields the proto marks `REQUIRED`. |
| `--emit-parent-refs` | Names the resource's real root, location and parent from the proto's `google.api.resource` pattern, rather than guessing. |
| `--emit-sibling-refs` | Generates a `<Kind>Ref` for a field that points at another resource in the same service being generated in the same run. |
| `--emit-reference-hints` | Reports fields that look like references without guessing at one. |
| `--emit-plural-acronyms` | Cases plural acronyms the way KRM conventions want, so `related_uris` becomes `relatedURIs` rather than `relatedUris`. |
| `--emit-message-maps` | Generates `map<string, Message>` fields as a map of the value's Go type instead of dropping them. `generate-mapper` needs the same flag. |
| `--place-server-set-fields` | Moves an allowlist of server-computed names (`createTime`, `uid`, `selfLink`, `etag`, …) into ObservedState when the proto carries no `field_behavior` anywhere. |
| `--detect-output-only-in-comments` | Reports spec fields whose proto comment says "Output only." but that carry no annotation. Reports only — moving one is a hand edit. |

### 5. Work the judgement queue

Several of the flags above make a call the proto could not make for them. Rather
than hide those calls, the generator writes each one to
`apis/<service>/needs_judgement_call.txt`, and `--detect-output-only-in-comments`
writes its findings to `apis/<service>/detected_output_only_in_comments.txt`.

**The queue is a to-do list, not an output artifact.** Each entry names a field path,
a reason and enough detail to decide. For each one either accept the generator's
choice or hand-edit the type, then delete the entry.

A resource is in exactly one of two states. **While it has queue entries, its
`[refs]` findings in `TestMissingRefs` are suppressed** — the generator has already
admitted it was unsure, so the ratchet does not also shout about it. Once you clear
the entries, the check applies in full.

That is the reason not to leave entries sitting: a stale queue is not a harmless
to-do, it is a check that is switched off. Equally, do not delete an entry you have
not actually looked at.

Note that a reference guess can be confidently wrong. Shared reference types are
matched by name suffix, so a proto whose pattern ends in `instances/{instance}` can
pick up an unrelated `Ref` type from another service. The queue is where you catch
that.

### 6. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
