---
name: kcc-bulk-resource-generator-generate
description: Phase 1 builder workflow for bulk-generating KRM types, identities, references, mappers, CRDs, and judgement queues across multiple greenfield GCP resources and splitting the resulting services into Category 1, Category 2, and Category 3 branches.
---

# KCC Bulk Resource Generator — Phase 1: Generate & Split

Use this skill to execute bulk `--prepopulate-resource` scaffolding (`<kind>_types.go`, `types.generated.go`, `<kind>_identity.generated.go`, `<kind>_reference.generated.go`, `mapper.generated.go`, CRDs, and `apis/<service>/judgement_queue.yaml`) across **multiple greenfield resources and services** and partition the output into three PR categories.

---

## 0. Sourcing Candidate GCP Resources

When the user asks to source missing GCP resources (or does not provide an explicit list of `<service>:<Kind>:<ProtoMessage>` targets), derive the candidate batch deterministically using `hack/tools/greenfield`:

1. **Generate the Unified Coverage List against Local `googleapis` and KCC:**
   ```bash
   python3 hack/tools/greenfield/calculate_coverage.py LOCAL LOCAL 50 --resources-list
   ```
   This reads `.build/third_party/googleapis` (pinned by `apis/git.versions`), applies `hack/tools/greenfield/coverage_skip.json`, `hack/tools/greenfield/deprecated_services.json`, and `hack/tools/greenfield/services_acronyms.json`, scans `pkg/controller/resourceconfig/static_config.go`, `config/crds/resources/*.yaml`, and `apis/*/*/*_types.go`, and writes `hack/tools/greenfield/resources_list.json`.
   > **Remember:** Revert or clean `hack/tools/greenfield/resources_list.json` (`git checkout -- hack/tools/greenfield/resources_list.json`) before committing your branch unless explicitly asked to update the snapshot.

2. **Filter `resources_list.json` to Actionable Bulk-Generation Candidates:**
   Select entries from `hack/tools/greenfield/resources_list.json` that satisfy **all** of the following:
   - `status == "missing"`
   - **Manageable lifecycle (`CREATE` + `READ` + `DELETE`):**
     - **Group 1 (Full CRUD):** `missing_ops == []` (`CREATE`, `READ`, `UPDATE`, `DELETE`)
     - **Group 2 (Immutable Declarative Infrastructure):** `missing_ops == ["UPDATE"]` (`CREATE`, `READ`, `DELETE`)
   - **GA Proto Package Only:** `pkg` is non-empty (skip REST-only/Discovery APIs with `pkg == ""`) and contains **no** `alpha` or `beta` version segment (enforced by `controllerbuilder`'s default `--allow-non-ga-proto=false`).
   - **Compiled in `.build/googleapis.pb`:** Verify that `proto_message` (`<pkg>.<ProtoName>`) exists as a message descriptor in `.build/googleapis.pb`.
   - **Declarative Resource Criteria:** Exclude non-declarative resources (ephemeral jobs, evaluations, previews, imports/exports, queries, conversational sessions, and other data-plane/imperative artifacts) enforced by `controllerbuilder`'s non-declarative filter (`dev/tools/controllerbuilder/pkg/commands/generatetypes/nondeclarative.go`) and [`docs/ai/declarative-resource-criteria.md`](../../../../docs/ai/declarative-resource-criteria.md).

---

## 1. Core Guardrails

1. **GA Protos Only (Fail-Fast on Non-GA Protos; Never Pass `--allow-non-ga-proto`):**
   - `controllerbuilder generate-types` rejects non-GA proto packages (`v1alpha*`, `v1beta*`, `v1p1beta1`, etc.) by default.
   - **Never** add `--allow-non-ga-proto` to `apis/<service>/generate.sh`.
   - **Fail-Fast Rule:** If `controllerbuilder generate-types` exits with `non-GA proto package ... is not allowed without --allow-non-ga-proto`, immediately run `git checkout -- apis/<service> && git clean -fd apis/<service>` and drop that resource from the batch. Do **not** retry with flags or hunt for alpha/beta workarounds.
2. **Always Generate Mappers alongside Types and Identities (in Strict Execution Order):**
   - Every service in the batch **must** run all three `controllerbuilder` commands in `apis/<service>/generate.sh` in the exact order below (because `generate-mapper` inspects the Go AST produced by `generate-types` and `generate-identity`, **all** `generate-types` and `generate-identity` calls across all API versions must precede `generate-mapper`; never append new `generate-types` or `generate-identity` blocks below an existing `generate-mapper` call):
     1. `generate-types --prepopulate-resource`
     2. `generate-identity`
     3. `generate-mapper --prepopulate-resource` (plus `--multiversion` if the service also generates `v1beta1` types).
3. **Pure Mechanical Generation (Leave `judgement_queue.yaml` Open):**
   - Do **not** resolve or delete entries in `apis/<service>/judgement_queue.yaml` during bulk generation. Each newly scaffolded `<Kind>` must keep `reason: untriaged-bulk-generation` with `status: open` for later Step 2 triage.

---

## 2. Standard `apis/<service>/generate.sh` Template

For each target service `<service_short>`, update or create `apis/<service_short>/generate.sh` so that `generate-types`, `generate-identity`, and `generate-mapper` all run with `--prepopulate-resource`.

> **Legacy `generate.sh` Upgrade Rule:** Many older `apis/<service>/generate.sh` scripts invoke `go run . generate-types` directly instead of `${CONTROLLERBUILDER}`. Always replace `go run .` with the `${CONTROLLERBUILDER}` block below so batch execution uses the prebuilt `bin/controllerbuilder` binary instead of recompiling on every command.

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
  --service <proto_service_package> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --resource <Kind1>:<ProtoMessage1> \
  --resource <Kind2>:<ProtoMessage2> \
  --prepopulate-resource

${CONTROLLERBUILDER} generate-identity \
  --service <proto_service_package> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --resource <Kind1>:<ProtoMessage1> \
  --resource <Kind2>:<ProtoMessage2>

${CONTROLLERBUILDER} generate-mapper \
  --service <proto_service_package> \
  --api-version <group>.cnrm.cloud.google.com/v1alpha1 \
  --include-skipped-output \
  --prepopulate-resource

cd ${REPO_ROOT}
dev/tasks/generate-crds

if [ -d "${REPO_ROOT}/pkg/controller/direct/<service_short>" ]; then
  go run -mod=readonly golang.org/x/tools/cmd/goimports@${GOLANG_X_TOOLS_VERSION} -w pkg/controller/direct/<service_short>/
fi
```

### Multi-Version Services (`v1alpha1` + `v1beta1`)
If `apis/<service_short>/generate.sh` already has a `v1beta1` `generate-types` block:
- Keep the existing `v1beta1` `generate-types` block unchanged.
- Insert the new `v1alpha1` `generate-types` and `generate-identity` blocks **before** the `generate-mapper` block (never append them at the bottom of the script after `generate-mapper`).
- Pass `--multiversion` and `--prepopulate-resource` to the single `generate-mapper` invocation (including both proto packages in `--service` if `v1alpha1` and `v1beta1` use different proto packages).

### KCC API Directory Aliases & Pre-Existing Unprefixed `*_reference.go` Stubs
1. **Service-to-API-Directory Mapping:** Some GCP proto service names map to a different canonical KCC directory under `apis/` (for example: `analyticshub` $\rightarrow$ `apis/bigqueryanalyticshub`, `cloudkms` $\rightarrow$ `apis/kms`, `config` $\rightarrow$ `apis/configdeployment`). Always check existing `apis/` directories and `Kind` prefixes before creating a new `apis/<service>` folder.
2. **Pre-Existing Unprefixed `<resource>_reference.go` Stubs:** Occasionally an `apis/<service>/v1alpha1/` directory already contains a handwritten stub reference file named `<resource>_reference.go` (without the `<service>` filename prefix, e.g., `connectionprofile_reference.go` or `dataattribute_reference.go`) created so another resource could reference it before `<service><resource>_types.go` existed:
   - Rename it with `git mv <resource>_reference.go <service><resource>_reference.go` so `TestDirectResourceFileNaming` in `tests/apichecks` passes once `<service><resource>_types.go` is generated.
   - If the stub `*_reference.go` declared `var <Kind>GVK = GroupVersion.WithKind("<Kind>")`, move that `var <Kind>GVK` declaration into `<service><resource>_types.go`.

---

## 3. Fast Local Batch Execution

When generating across many services locally, avoid recompiling `controllerbuilder` or regenerating all CRDs on every service iteration (`generate-proto.sh` automatically caches `.build/googleapis-<SHA>.pb`):

1. **Build `bin/controllerbuilder` and extract protos once:**
   ```bash
   (cd dev/tools/controllerbuilder && go build -o ../../../bin/controllerbuilder . && ./generate-proto.sh)
   ```
2. **Run each service's `generate.sh` with `SKIP_GENERATE_CRDS=1` (run twice to verify idempotence):**
   ```bash
   for svc in <service_1> <service_2> ...; do
     chmod +x "apis/${svc}/generate.sh"
     SKIP_GENERATE_CRDS=1 ./apis/${svc}/generate.sh
     SKIP_GENERATE_CRDS=1 ./apis/${svc}/generate.sh
   done
   ```
3. **Add all newly generated `<Kind>`s to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`, generate CRDs, and format once at the end:**
   - Add each newly generated `<Kind>` (sorted alphabetically) to the `BLOCKED_KINDS=( ... )` array in `dev/tasks/build-release-bundle` so unverified Greenfield CRDs are excluded from `operator/channels/packages/configconnector/*/crds.yaml` until graduated at the end of the pipeline.
   ```bash
   ./dev/tasks/generate-crds
   make fmt
   ```

---

## 4. Deterministic Service Classification & Splitting into 3 PR Categories

For each generated service `<svc>`, run `go build ./apis/<svc>/... ./pkg/controller/direct/<svc>/...` and `go test ./tests/apichecks/...`, then classify `<svc>` using the following **deterministic decision tree**:

1. **Check for Category 3 (`Major Surgery Needed to Get PR Green`) first:**
   - Did `apis/<svc>/generate.sh` require multi-stage `types.generated.go` file renaming across multiple proto packages (e.g., `discoveryengine`, `networksecurity`)? $\rightarrow$ **Category 3**.
   - Did `go build ./pkg/controller/direct/<svc>/...` fail with missing `*Request`/`*Metadata` mapper functions (`undefined: direct.<RequestOrMetadata>_FromProto`) or broken existing `v1beta1` controller callers after enabling `--multiversion` / plural acronym casing (`URIs`)? $\rightarrow$ Fix in `pkg/controller/direct/<svc>/` and classify as **Category 3**.
   - Did `go test ./tests/apichecks/...` fail on `<svc>`'s CRD with `TestCRDObjectTypes` (`field ... is an object with no properties` from an empty parameterless proto `oneof` branch struct, e.g., `message LlmAgent {}`)? $\rightarrow$ Add the CRD to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with the comment `// GCP known pattern: Empty oneof option (<field path>)` and classify as **Category 3**.
2. **Otherwise, check for Category 2 (`Initial Setup — go.mod & Service Scaffolding`):**
   - Did `<svc>` require creating a brand-new `apis/<svc>/v1alpha1/doc.go` or `groupversion_info.go` package? $\rightarrow$ **Category 2**.
   - Did `go build ./pkg/controller/direct/<svc>/...` fail with `no required module provides package cloud.google.com/go/...`? $\rightarrow$ Run `go get <pkg>@latest && go mod tidy` and classify as **Category 2**.
   - Did `go build ./pkg/controller/direct/<svc>/...` fail with `undefined: direct.<SharedType>_FromProto` (e.g., `Money`, `LatLng`, `Expr`)? $\rightarrow$ Append `,google.type` to `generate-mapper --service` in `apis/<svc>/generate.sh`, re-run `SKIP_GENERATE_CRDS=1 ./apis/<svc>/generate.sh`, and classify as **Category 2**.
3. **Otherwise, classify as Category 1 (`Zero-Touch`):**
   - Existing service where adding `--resource <Kind>:<Proto>` and `--prepopulate-resource` to `generate-types`, `generate-identity`, and `generate-mapper` compiled and passed `apichecks` (with `WRITE_GOLDEN_OUTPUT=1`) out of the box—zero `go.mod` changes, zero new package boilerplate, and zero handwritten code changes.

| Category | Summary | PR Strategy |
| :--- | :--- | :--- |
| **Category 1: Zero-Touch** | Existing services where types, identities, and mappers compile and pass `apichecks` with no `go.mod`, package scaffolding, or manual code edits. | **1 bulk PR** for all Category 1 services. |
| **Category 2: Initial Setup** | Services needing only initial `go.mod`/`go.sum` dependencies, new `apis/<svc>/v1alpha1/{doc.go,groupversion_info.go}` boilerplate, or `,google.type` in `generate-mapper --service`. | **1 bulk setup PR** for all Category 2 services. |
| **Category 3: Major Surgery** | Services needing custom mapper stubs, existing controller caller updates, `TestCRDObjectTypes` empty `oneof` fixes, or multi-stage `generate.sh` scripts. | **Dedicated PR(s)** per service or failure mode. |

### Leak-Proof Recipe for Splitting the Scratch Branch into Category Branches

To prevent untracked files from one category leaking into another category's branch:

1. **Commit the entire generated state onto a temporary scratch branch (`bulk-scratch`) first** so `git status --porcelain` is completely clean.
2. **For each target category branch (`<category-branch>`):**
   ```bash
   git checkout -B <category-branch> upstream/master
   git clean -fd

   # Check out only the services belonging to this category from bulk-scratch
   for svc in <cat_svc_1> <cat_svc_2> ...; do
     git checkout bulk-scratch -- "apis/${svc}" "pkg/controller/direct/${svc}"
   done

   # For Category 2 (or Category 3 if needed), also check out go.mod/go.sum and run go mod tidy:
   # git checkout bulk-scratch -- go.mod go.sum && go mod tidy

   # Add ONLY this category's newly generated Kinds (sorted alphabetically) to BLOCKED_KINDS in dev/tasks/build-release-bundle,
   # then regenerate CRDs, supportedgvks, and apichecks goldens for ONLY this category's services:
   ./dev/tasks/generate-crds
   # If generate-crds added a brand-new API group to config/installbundle/components/clusterroles/cnrm_{admin,viewer}.yaml,
   # also run ./dev/tasks/generate-channels so 0-cnrm-system.yaml stays in sync.
   WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...
   git add -A && git commit -m "<category commit message>"
   ```
3. Proceed to **Phase 2 (Isolated Subagent Verification)** in `.gemini/skills/kcc-bulk-resource-generator/SKILL.md`.

- **Declarative Resource Filtering:** Before bulk-generating types or controllers, check candidate resources against [`docs/ai/declarative-resource-criteria.md`](../../../docs/ai/declarative-resource-criteria.md) (or `../../../../docs/ai/declarative-resource-criteria.md`) to separate read-only/ephemeral/data-plane resources from declarative control-plane resources.
