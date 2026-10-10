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
`apis/<service>/judgement_queue.yaml`. Pass `--prepopulate-spec` and
`--emit-reference-hints` together, so every field `TestMissingRefs` would flag
also gets an entry.

Each entry names a resource (`kind` and `group`) or a shared proto message
(`protoMessage`), a `field`, a `reason` and a `detail`. New entries start as
`status: open`. This is from `apis/chronicle`, with the second entry already
resolved:

```yaml
entries:
  - kind: ChronicleWatchlist
    group: chronicle.cnrm.cloud.google.com
    reason: untriaged-bulk-generation
    detail: spec was generated from proto definition; verify refs, omissions, and KRM conventions
    status: open
  - kind: ChronicleWatchlist
    group: chronicle.cnrm.cloud.google.com
    reason: parent-ref-not-modelled
    detail: the parent is projects/{project}/locations/{location}/instances/{instance}, and no InstanceRef type exists to point at; add the parent resource first, then model this as a reference
    status: resolved
    resolution: deferred
    note: kept location and projectRef until ChronicleInstance exists
```

Review each open entry. Then set `status: resolved` and a `resolution`:

| Resolution | Use it when | Note |
|---|---|---|
| `accepted` | The generated output is right as it is. | Optional |
| `edited` | You changed the type by hand. | Required: say what you changed |
| `deferred` | The finding is right but is handled elsewhere, such as `refs_deferred.txt`. | Required: say where |
| `not-applicable` | The finding is wrong for this resource. | Required: say why |

**Never delete an entry.** The file is the record of what was decided. An
`edited` entry and its note are how a hand edit gets reported. Regenerating keeps
your status, resolution and note, and only refreshes `detail`.

**An open reference entry suppresses one `[refs]` finding.** `TestMissingRefs`
skips a finding only when an open entry has the same kind, group and field and a
`possible-reference*` reason. Every other finding on the resource is still checked.
Once you resolve the entry, the check applies to that field again. So if you
resolve a reference entry without adding the reference, also add the field to
`tests/apichecks/testdata/exceptions/refs_deferred.txt` with a `reason=`, or
`TestMissingRefs` fails.

Other reasons, including `untriaged-bulk-generation`, do not suppress anything.
They are there so a person looks at them.

Note that a reference guess can be confidently wrong. Shared reference types are
matched by name suffix, so a proto whose pattern ends in `instances/{instance}` can
pick up an unrelated `Ref` type from another service. The queue is where you catch
that.

### 6. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
