---
name: kcc-direct-greenfield-types-prepopulate
description: Step 1 of the KCC Greenfield pipeline. Mechanically scaffolds a new greenfield resource (or batch of resources) using `generate-types --prepopulate-resource`, `generate-identity`, and `generate-mapper --prepopulate-resource`, emitting `apis/<service>/judgement_queue.yaml` for Step 2 triage.
---

# KCC Direct Greenfield Types Implementer — Step 1: Prepopulate Workflow

Use this workflow for **Greenfield Step 1** when `apis/<service>/<version>/<kind>_types.go` does **not** exist yet. This step mechanically generates the initial KRM types, identity, reference, mapper, CRDs, and `apis/<service>/judgement_queue.yaml` using `controllerbuilder`.

> [!IMPORTANT]
> **Strict Mechanical Scope for Step 1**:
> - **Do NOT hand-edit** `<kind>_types.go`, `types.generated.go`, `<kind>_identity.generated.go`, `<kind>_reference.generated.go`, or `apis/<service>/judgement_queue.yaml`.
> - **Do NOT resolve** any entries in `apis/<service>/judgement_queue.yaml` during Step 1; `untriaged-bulk-generation` for `<Kind>` **must** remain `status: open` so Step 2 (`judgement/SKILL.md`) can triage it.
> - **Do NOT pass `--allow-non-ga-proto`** unless explicitly instructed by a human.

---

## Inputs
- `service`: The Google API proto package name (e.g., `google.cloud.apihub.v1`).
- `resource`: The mapping of KCC Kind to GCP Proto message (e.g., `APIHubPlugin:Plugin`). The Kind must follow the service-prefix rule in Phase 1, step 3.
- `api_version`: The KCC API version (must be `v1alpha1` for new greenfield resources).

---

## Workflow

### Phase 1: Proto Discovery & GA Verification
1. Ensure `.build/third_party/googleapis` is populated (run `./dev/tools/controllerbuilder/generate-proto.sh` if missing).
2. Locate the target proto message in `.build/third_party/googleapis/google/...` and confirm it belongs to a **GA** proto package (e.g., `v1`, `v2`, `v4` — not `v1alpha*` or `v1beta*`).
   - Do **not** pass `--allow-non-ga-proto` unless explicitly instructed.
3. **Determine the Kind** from the service-prefix files in `dev/tools/controllerbuilder/serviceprefixes/`. `generate-types --prepopulate-resource` enforces this and writes nothing if it does not hold.
   - The **API** is the proto service's `google.api.default_host` without `.googleapis.com` (e.g. `sqladmin`), or, if the package has no service, the host of the message's `google.api.resource` type.
   - The **Kind** is the API's `servicePrefix` followed by the `google.api.resource` type in the singular: `logging` + `logging.googleapis.com/Settings` gives `LoggingSetting`. Do not use the proto message name if it differs from the resource type, and do not keep a plural.
   - If the API is in `service-prefixes.yaml`, use its `servicePrefix`.
   - If the API is only in `service-prefix-exceptions.yaml`, its prefix is **pending approval**: stop and ask a human to approve it (move the entry to `service-prefixes.yaml`, keeping or changing `servicePrefix`, with the decision in `reason:`). Do **not** approve it yourself.
   - If the API is in **neither** file, add an entry to `service-prefix-exceptions.yaml` with `api`, `title` (the service config title), the proposed `servicePrefix`, and a `reason:` saying why it needs a decision (e.g. `new API with no KCC Kinds yet`), then stop and ask a human to approve it. Propose the API name kept verbatim (including `Cloud` and `Admin`), with word boundaries and casing from the title; if KCC already has Kinds for the API, propose the prefix they use.

### Phase 2: Configure `apis/<service_short>/generate.sh`
Locate `apis/<service_short>/generate.sh` (or create it if the service is brand new) and configure all three `controllerbuilder` commands (`generate-types`, `generate-identity`, and `generate-mapper`) with `--prepopulate-resource`:

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
./generate-proto.sh

${CONTROLLERBUILDER} generate-types \
  --service <service> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource <Kind>:<ProtoMessage> \
  --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
  --service <service> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --resource <Kind>:<ProtoMessage>

${CONTROLLERBUILDER} generate-mapper \
  --service <service> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/<service_short>" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/<service_short>/
fi
```

*Notes on `apis/<service_short>/generate.sh`*:
- **Strict Command Ordering**: Always place all `generate-types` and `generate-identity` invocations **before** `generate-mapper`, followed by `cd ${REPO_ROOT}`, `dev/tasks/generate-crds`, and `goimports`. When adding a new resource to an existing `generate.sh` that already has `generate-mapper`, insert the new `generate-types` / `generate-identity` blocks **above** `generate-mapper` (never below `generate-mapper`, otherwise `generate-mapper` runs before `types.generated.go` is updated).
- If the service already has `v1beta1` resources in `apis/<service_short>/generate.sh`, pass `--multiversion` and `--prepopulate-resource` to the single `generate-mapper` invocation targeting `<group>.cnrm.cloud.google.com/v1beta1`.
- `--prepopulate-resource` automatically enables all bulk-generation features (including `--emit-plural-acronyms` and `--emit-message-maps`) in both `generate-types` and `generate-mapper`.

### Phase 3: Run Generation, Blocklist Registration, & Minimal Mapper Compilation Fixes
1. Make `apis/<service_short>/generate.sh` executable and run it:
   ```bash
   chmod +x apis/<service_short>/generate.sh
   ./apis/<service_short>/generate.sh
   ```
   - **If `generate-types` fails with `rejected N resource(s) whose Kind does not match the service-prefix table`**, nothing was generated for that `generate-types` invocation (and, because `generate.sh` uses `set -o errexit`, the rest of the script did not run). Each listed resource says what to fix:
     - `API <api> has service prefix <P>, so the Kind should be <Kind>`: change `--resource <Kind>:<Message>` in `apis/<service_short>/generate.sh` (and every other place that names the Kind, such as `generate-identity` and `BLOCKED_KINDS`) to the Kind given, and re-run.
     - `the service prefix of API <api> (<P>) is pending approval in service-prefix-exceptions.yaml`: stop and ask a human to approve the prefix (Phase 1, step 3).
     - `API <api> has no service prefix in service-prefixes.yaml or service-prefix-exceptions.yaml`: propose one in `service-prefix-exceptions.yaml` and ask a human to approve it (Phase 1, step 3).
     - There is no override flag; do not work around the check.
2. **Add `<Kind>` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`**:
   - Add `<Kind>` formatted with 4-space indentation and double quotes (`    "<Kind>"`), sorted alphabetically into the existing `BLOCKED_KINDS=( ... )` array in `dev/tasks/build-release-bundle` (without duplicating any pre-existing entry). This prevents unverified Greenfield CRDs from being packaged into the release bundle (`operator/channels/packages/configconnector/*/crds.yaml`) until a graduation agent removes `<Kind>` at the end of the generation pipeline.
   - If `dev/tasks/generate-crds` added a brand-new API group to `config/installbundle/components/clusterroles/cnrm_admin.yaml` or `cnrm_viewer.yaml`, run `./dev/tasks/generate-channels` and `WRITE_GOLDEN_OUTPUT=1 go test ./operator/pkg/controllers/configconnector/...` so `0-cnrm-system.yaml` and `operator/pkg/controllers/configconnector/testdata/golden/simple/_expected.yaml` stay in sync.
3. Verify that `apis/<service_short>/judgement_queue.yaml` was generated and contains `reason: untriaged-bulk-generation` with `status: open` for `<Kind>`.
4. Test compilation of both `apis/<service_short>/...` and `pkg/controller/direct/<service_short>/...`:
   ```bash
   go build ./apis/<service_short>/... ./pkg/controller/direct/<service_short>/...
   ```
5. **If `pkg/controller/direct/<service_short>/...` fails to compile**, apply **only** the minimal fixes required for `mapper.generated.go` to build cleanly (without hand-editing `<kind>_types.go` or `judgement_queue.yaml`):
   - **Missing Go client library in `go.mod`**: Run `go get cloud.google.com/go/<service>/apiv1/<service>pb` and `go mod tidy`.
   - **Missing external proto mappers (e.g. `Money_FromProto`, `LatLng_FromProto`, `Expr_FromProto`)**: Append `,google.type` (or the relevant external proto package) to `--service` on `generate-mapper` in `apis/<service_short>/generate.sh` and re-run `./apis/<service_short>/generate.sh`.
   - **Missing mapper helpers for `*Request` / `*Metadata` messages or wrapper types**: Add a minimal hand-written `<kind>_mapper.go` (using a valid `<Kind>` prefix so `TestDirectResourceFileNaming` does not flag `naming_violations.txt`, or `mapper.go`) in `pkg/controller/direct/<service_short>/` implementing only the missing `*_FromProto` / `*_ToProto` helper functions.
   - **Existing `v1beta1` controller callers when adding `--multiversion`**: Update existing callers in `pkg/controller/direct/<service_short>/` to call the versioned `*_v1beta1_FromProto` / `*_v1beta1_ToProto` functions (or updated `URIs` acronym casing).

### Phase 4: Idempotency & `apichecks` Verification
1. Re-run `./apis/<service_short>/generate.sh` and verify that re-running the script produces zero uncommitted diffs.
2. Run unit tests and `apichecks`:
   ```bash
   go test -v ./apis/<service_short>/... ./pkg/controller/direct/<service_short>/...
   go test -v ./tests/apichecks/...
   ```
   - Because `<Kind>` has `untriaged-bulk-generation` with `status: open` in `apis/<service_short>/judgement_queue.yaml`, `TestMissingRefs`, `TestDescriptions`, `TestFieldNames`, `TestRequiredFields`, and `TestCRDBackwardCompatibility` automatically skip `<Kind>` until Step 2 (`judgement/SKILL.md`).
   - In Step 1 (`--prepopulate-resource`), running `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` may legitimately record untriaged items in `alpha-missingfields.txt`, `refs_not_representable.txt`, `spec_dislike_etag.txt`, `sensitive.txt`, `comments.txt`, and `recursivetypes.txt` (revert any spurious `tests/apichecks/testdata/exceptions/multi_version_crd_diff/` formatting churn).
   - If `recursivetypes.txt` or `envtest` (`pkg/cli/preview`) flags a recursive proto message cycle where `controller-gen` emits an empty `{}` object schema, add `+kubebuilder:validation:Schemaless`, `+kubebuilder:validation:XPreserveUnknownFields` (or `+kubebuilder:validation:items:XPreserveUnknownFields` for slice fields), and `+kubebuilder:validation:Type=object` on the cycle-edge field in `apis/<service_short>/v1alpha1/<kind>_types.go` and re-run `./dev/tasks/generate-crds`.
   - If `TestCRDObjectTypes` fails with `field ... is an object with no properties` on a genuine parameterless proto `oneof` variant struct (e.g., `message LlmAgent {}`), add the CRD plural to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with the comment `// GCP known pattern: Empty oneof option (<field path>)`.
   - **Never** modify `tests/apichecks/testdata/exceptions/missingrefs.txt`, `naming_violations.txt`, or `tests/apichecks/testdata/exceptions/multi_version_crd_diff/`.
