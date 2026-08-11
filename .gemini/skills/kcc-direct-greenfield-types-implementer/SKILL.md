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
- **Field Validation**: Manually add or verify kubebuilder tags:
  - Use `// +kubebuilder:validation:Required` for fields that are mandatory in the GCP API.
  - Use `// +kubebuilder:validation:Optional` for all other fields.
- **Enums**: 
  - Use `*string` for the Go type of proto enum fields (do NOT use custom wrapped string types).
  - Use `// +kubebuilder:validation:Enum=VALUE1;VALUE2` to provide validation in the CRD while keeping the Go type simple.

### 4. Reference Fields (do not skip)
Fields that point at another GCP resource **must** be implemented as KCC reference fields
(`Ref` suffix, e.g. `refsv1beta1.ProjectRef`, `pubsubv1beta1.PubSubTopicRef`), per
`.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`.

**You MUST NOT add entries to `tests/apichecks/testdata/exceptions/missingrefs.txt`.**
That file is a shrink-only ratchet: new entries fail CI and cannot be regenerated away,
including under `WRITE_GOLDEN_OUTPUT=1`. If the check flags a field, implement the
reference - do not suppress it.

**Primary signal - check the proto, not the description.** GCP protos annotate reference
fields with `(google.api.resource_reference)`, which names the exact target type:

```proto
optional string service_account = 16 [
  (google.api.field_behavior) = REQUIRED,
  (google.api.resource_reference) = { type: "iam.googleapis.com/ServiceAccount" }
];
```

Inspect the source `.proto` for this annotation on every string field before deciding it is
a plain string. It is authoritative where present, but coverage is uneven (~15% of string
fields overall, 0% in compute), so its absence proves nothing.

**Secondary signals** (used when the annotation is missing):
- Description contains a path template: `projects/`, `locations/{`, `folders/{`, `organizations/{`
- A Cloud Storage **bucket** name -> `StorageBucketRef`
- KMS keys, service accounts, networks/subnetworks, Pub/Sub topics, Secret Manager secrets

**Not every URI is a reference.** Cloud Storage *object paths* (`gs://bucket/path/*`) cannot
be modeled as refs today: `StorageBucketIdentity.FromExternal` accepts only the bare
`gs://<bucket>` form. `bq://` is not a GCP resource name at all. These are recorded in
`tests/apichecks/testdata/exceptions/refs_not_representable.txt` with a reason, and should
be left as strings rather than forced into a ref.

### 5. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
