---
name: kcc-brownfield-triage
description: Triage a KCC brownfield migration issue - filter out skipped legacy resources (shut down, deprecated, superseded, or duplicated) and check whether an existing Terraform/DCL resource kind is ready to migrate now via the `kcc-brownfield` graph (API access and proto/SDK support, quota, dependencies and fixtures, lifecycle, state on master, `factory-graph` start block), deciding ready, blocked, or unsure.
---

# KCC brownfield triage

Triage a KCC brownfield migration issue (Terraform/DCL-to-direct) and decide whether it is ready to start, should be skipped/stopped, or needs a person to check.

Read-only triage: **do not edit code, open PRs, or post comments on the issue.** Report your findings and labels at the end.

In the commands below, replace `<issue_number>` with the issue number and `<gcp_project>` with the test GCP project.

## Why

The `kcc-brownfield` graph runs six steps (`types` → `identity` → `fuzzer` → `mockgcp` → `controller` → `validate-promotion`) and requires real-GCP recording (`hack/record-gcp`) in `mockgcp` and `validate-promotion`. Starting a migration for a GCP service or resource that has been shut down, scheduled for shutdown, deprecated in favor of an in-product replacement, or already duplicated in KCC wastes agent runs and fails real-GCP verification. Filtering these out up front keeps the migration queue focused on viable resources.

## Checks

### 1. Find the kind, service, and legacy configuration

Read the issue:
```bash
gh issue view <issue_number> --comments
```
Look for the `factory-graph` fenced block (`params.kind`) or the kind in the issue title (for example `Migrate ComputeURLMap to direct` or `Implement direct controller for ...`).

Confirm it is an existing brownfield KCC resource and note its `group`, `version`, and current controller configuration:
- Check `pkg/controller/resourceconfig/static_config.go` for `{Group: ..., Kind: "<Kind>"}`.
- Check the CRD under `config/crds/resources/` (file name `apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<service>.cnrm.cloud.google.com.yaml`).
- Check legacy metadata in `legacy_resources_classified.json` (if present) and service mappings under `pkg/servicemapping/servicemappingloader/` or `pkg/dcl/` to see the underlying GCP API/Terraform/DCL resource type.

### 2. Check Skip & Access-Gate Patterns (filter out non-viable, redundant, or allowlist-gated migrations)

Before spending time on deep implementation checks or live quota probes, evaluate whether `<Kind>` falls into any of the five patterns below. If any pattern clearly applies, the issue is **blocked** (do not start the `kcc-brownfield` graph).

#### Pattern A: Fully Shut Down or Sunset Services

The underlying GCP product, API, or architecture has been shut down, turned down, or permanently retired.

How to identify:
- **GCP service / CLI probe**: Check if the API domain can still be discovered or enabled (`gcloud services list --available`, `gcloud services enable <api>.googleapis.com --project <gcp_project>`), or if API / `gcloud` calls fail with retirement, shutdown, or non-existent service errors.
- **API & docs check**: Search `.build/third_party/googleapis` and GCP product documentation / release notes for retirement, end-of-life, or turndown dates.
- **Fixture behavior**: Existing E2E tests under `pkg/test/resourcefixture/testdata/basic/` cannot succeed against real GCP because the service no longer accepts calls or new provisioning.

#### Pattern B: Deprecated Services & Scheduled Shutdowns

The GCP product, API version, or hosted control-plane offering is officially deprecated, closed to new customers (sunset), or has an announced shutdown date in favor of a newer GCP product (for example, legacy ML/TPU/feature-store/source-repo/deployment services or hosted wrappers replaced by standalone tooling).

How to identify:
- **Proto / API metadata**: Inspect `.build/third_party/googleapis` for `option deprecated = true;` on services, RPCs, or messages, and read comments for deprecation or shutdown timelines.
- **GCP docs & CLI notices**: Check GCP product deprecation announcements, `gcloud` deprecation warnings, and whether the service is restricted from new projects/customers.
- **Terraform provider notices**: Check resource docs and schemas (note: vendored files under `third_party/github.com/hashicorp/terraform-provider-google-beta/` may lag upstream, so verify against current GCP docs/API status rather than assuming silence in the vendored fork means active support).

#### Pattern C: Deprecated Resources with Direct In-Product Replacements

The GCP service is still active, but this specific resource or API model is legacy, deprecated, or superseded by a newer resource within the same product or ecosystem (for example: legacy health checks or load-balancing targets replaced by unified resources, legacy org/resource ACLs or security policies superseded by unified firewall policies or IAM/UBLA-compatible models, or older metadata/config/branding APIs replaced by newer product APIs).

How to identify:
- **Successor notices**: Inspect GCP REST/gRPC API docs, proto comments (`deprecated = true` / *"use ... instead"*), `gcloud` docs, and upstream Terraform provider docs for *"deprecated in favor of"*, *"superseded by"*, or migration guides pointing to a replacement resource.
- **Modern default incompatibilities**: Check if the legacy resource conflicts with modern GCP project defaults (such as legacy object/bucket ACLs that fail when Uniform Bucket-Level Access is enforced).
- **Replacement coverage in KCC**: Check `pkg/controller/resourceconfig/static_config.go` and `config/crds/resources/` to see whether the replacement Kind already exists in KCC (or already has a Direct controller).

#### Pattern D: Duplicated Resources in KCC

Another KCC CRD — particularly an existing `v1beta1` or `Direct` resource — already manages the same underlying GCP resource or capability.

How to identify:
- **Resolve the underlying GCP target**: Check `<Kind>`'s CRD (`config/crds/resources/`), Terraform mapping (`pkg/servicemapping/servicemappingloader/`), or DCL mapping (`pkg/dcl/`) to determine the exact GCP API endpoint, URL template, or resource type.
- **Search sibling Kinds**: Search `pkg/controller/resourceconfig/static_config.go` and `config/crds/resources/` across the same and related `*.cnrm.cloud.google.com` groups for another Kind targeting the same GCP resource (for example, multiple `v1alpha1` scope-specific Kinds such as Project/Folder/Organization variants duplicated by a single `v1beta1` unified Kind).

#### Pattern E: Early Access, Private Preview, or Allowlist-Only APIs / Features

The GCP service, API version, or specific resource feature is not generally available and requires project allowlisting, enrollment, or special bind permissions (for example, `financialservices.googleapis.com` / Anti-Money Laundering AI or Early Access partner APIs). Because `hack/record-gcp` must run against a live test project in `mockgcp` and `validate-promotion`, migrations cannot complete unless `<gcp_project>` is allowlisted.

How to identify:
- **Service visibility & enablement probe**: Check whether the service appears in `gcloud services list --available --project <gcp_project>` and whether `gcloud services enable <api>.googleapis.com --project <gcp_project>` succeeds. Allowlisted/private APIs typically omit the service from public availability or fail enablement—even for project Owners—with `PERMISSION_DENIED` requiring `servicemanagement.services.bind` (and subsequent API calls return `403 SERVICE_DISABLED`).
- **Resource/feature-level allowlist probe**: Even when the parent API is enabled (for example, Compute Engine), specific resource types or fields may return `403 PERMISSION_DENIED` / `FAILED_PRECONDITION` errors stating the project is not allowlisted or enrolled in Preview/Early Access.
- **Docs, protos & repo history**: Check GCP documentation for *"Early Access"*, *"Private Preview"*, *"allowlist"*, or *"contact your account team to enable"*; check `.build/third_party/googleapis` for restricted visibility/preview notices; and check `.gemini/journals/<service>.md` plus open/closed GitHub issues (`allowlist <service>`) for prior `testing-blocked` or `servicemanagement.services.bind` findings.

### 3. Readiness checks (when not skipped)

Check all of the following for viable resources:

- **API access and service usability**: Not just documented — usable in `<gcp_project>`.
  Check the service API can be enabled and called in `<gcp_project>` (`gcloud services enable ...`, `gcloud ... list`, or `gcloud asset`/`curl` check) without hitting an allowlist gate, Preview/private-API block, or product retirement error. Both `mockgcp` and `validate-promotion` require `hack/record-gcp` to succeed against real GCP.

- **Proto / Go SDK availability**: Verify a direct client surface exists.
  Check `.build/third_party/googleapis` (or `google.golang.org/api/...` / `cloud.google.com/go/...` in `go.mod`) for the proto messages, gRPC/REST client, and CRUD methods corresponding to the CRD. If the legacy TF/DCL resource relies on an undocumented, legacy-only, or removed API surface with no usable Go SDK/proto, note the gap.

- **Quota**: Enough quota in `<gcp_project>` for the resource and its test fixtures to create, update, and re-reconcile during `TestAllInSeries` and `TestMigrationToDirect`.

- **Existing E2E fixtures and dependencies**: Look under `pkg/test/resourcefixture/testdata/basic/` (and `pkg/test/resourcefixture/testdata/basic/<service>/...`) for existing `<Kind>` test cases (`create.yaml`, `update.yaml`, `dependencies.yaml`, `_http.log`).
  - Check whether every `dependencies.yaml` resource (and `*Ref` in the CRD) works and isn't itself a non-viable/skipped legacy resource.
  - Check whether the fixture requires organization-level, folder-level, billing, or manual out-of-band setup that cannot run in the standard E2E project without special configuration.

- **Lifecycle**: Create, update (if spec has mutable fields), and delete work end-to-end in `<gcp_project>`, and the resource is cleanly deletable during test cleanup.

- **State on master**: Check what already exists on `master` (and check open PRs with `gh pr list --search "<Kind>"`):
  - `pkg/controller/resourceconfig/static_config.go`:
    - If `DefaultController` is already `k8s.ReconcilerTypeDirect`, migration is already complete → **blocked** (nothing left to migrate; recommend closing / `overseer/stop`).
    - If `SupportedControllers` includes `k8s.ReconcilerTypeDirect` while `DefaultController` is still `Terraform`/`DCL`, note this (only `validate-promotion` / promotion remains, or controller already exists).
  - Existing direct artifacts on `master`:
    - Direct types + `generate.sh`: `apis/<service>/<version>/<kind>_types.go` and `apis/<service>/<version>/generate.sh`
    - Identity & refs: `apis/<service>/<version>/<kind>_identity.go` and `_reference.go`
    - Fuzzer: `pkg/controller/direct/<service>/<kind>_fuzzer.go` + `pkg/controller/direct/register/register.go`
    - MockGCP: `mockgcp/mock<service>/`
    - Direct controller: `pkg/controller/direct/<service>/<kind>_controller.go`
  - Note: `.agents/graphs/kcc-brownfield.yaml` handles already-present `types` and `identity` steps by having those steps close as *not planned* (skipped), so existing Step 1/Step 2 files do **not** block starting `kcc-brownfield` unless the resource is already migrated (`DefaultController: Direct`) or an active PR is already in flight.

- **Start block**: The issue body has a `factory-graph` fenced block with `graph: kcc-brownfield` and `params.kind: <Kind>`:
  ````markdown
  ```factory-graph
  graph: kcc-brownfield
  params:
    kind: <Kind>
  ```
  ````
  If missing or malformed, it is **blocked**.

### 4. Decide

- **ready** — Does not match any skip or access-gate pattern (not shut down, deprecated/sunset, superseded by an in-product replacement, duplicated in KCC, or allowlist-restricted/Early Access), not already `DefaultController: Direct` on `master`, API is usable with quota, proto/SDK and dependencies/fixtures are viable, lifecycle works, and the `kcc-brownfield` `factory-graph` start block is present.
- **blocked** — Matches one of the skip/access-gate patterns (Pattern A/B/C/D for permanent skip, or Pattern E for allowlist/Early Access block), OR already migrated to `DefaultController: Direct`, OR API is unusable, OR missing quota/parent/lifecycle/proto support or `factory-graph` start block.
- **unsure** — Cannot be confirmed with read-only checks (for example, ambiguous deprecation/replacement status, unverified Early Access/allowlist or quota status without live project access, or special org/billing setup requirements).

## Report

First line:
- `"**Ready to implement.**"` when **ready**
- `"**Blocked: recommend overseer/stop.**"` when **blocked** (including skipped or allowlist-gated resources)
- `"**Needs a person to check.**"` when **unsure**

Then one bullet per check, stating what you found and how you verified it (skip & access-gate pattern checks, API access, proto/SDK, quota, dependencies/fixtures, lifecycle, state on master, `factory-graph` start block).

If **blocked**, add `"To unblock:"` stating clearly whether it is a permanent skip or an access blocker:
- For **Patterns A–D**: `"Resource matches skip pattern (<Pattern A/B/C/D>): <evidence and replacement/status> — close issue as won't-do / do not migrate."`
- For **Pattern E**: `"Resource uses an Early Access / allowlist-only API (<service/feature>): <error or doc evidence> — grant allowlisting on <gcp_project> (e.g. `servicemanagement.services.bind` / feature enrollment) before running `kcc-brownfield`, or defer/close if real-GCP testing is unavailable."`
- For other blockers (missing quota, missing `factory-graph` block, already `DefaultController: Direct`, etc.): state the concrete action needed.
If **unsure**, add `"To check:"` with the exact questions for a person to verify.
If **ready**, end with `"Starting its graph: labelled `overseer/graph`."`.

## Labels

Apply **only** the single overseer routing label indicated by the decision (do not add or remove any other labels):
- `overseer/graph` when **ready**
- `overseer/stop` when **blocked** or **unsure**
