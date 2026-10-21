# Evaluation of Deterministic Controller Generation: PR #14251 vs. PR #14222 / PR #14224

This document provides an unbiased technical comparison between two competing approaches for deterministic Greenfield Direct controller generation in Config Connector:

* **[PR #14251](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14251)** (`barney-s:feature/controller-profile-shadow-eval`): `feat(controllerbuilder): introduce ControllerProfile generation and shadow evaluation`
* **[PR #14222](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14222) + [PR #14224](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14224)** (`ldanielmadariaga:feat-compare-greenfield-spec` & `ldanielmadariaga:feat-cb-deterministic-controller`): `direct/common: add CompareGreenfieldSpec and CompareGreenfieldSpecAndLabels` & `controllerbuilder: implement Pass 1-5 deterministic controller generation pipeline`

---

## Part 1: Objective Comparison — PR #14251 vs. PR #14222 / PR #14224

| Dimension | [PR #14251](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14251) (`controllerprofile`) | [PR #14222](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14222) + [PR #14224](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14224) (`CompareGreenfieldSpec` + Pass 1–5 Pipeline) |
| :--- | :--- | :--- |
| **CLI & Automated Extraction Wiring** | **Not wired.** Adds a standalone package `dev/tools/controllerbuilder/pkg/controllerprofile` (`extractor.go`, `profile.go`, `template.go`). `ProtoServiceSpec` and `KRMSpecInfo` are manually populated structs in unit tests; not connected to `.build/googleapis.pb`, `go/ast`, or `controllerbuilder generate-direct-reconciler` (listed as "Next Steps" in Section 7 of its design doc). | **Fully wired end-to-end.** Connected directly to `controllerbuilder generate-direct-reconciler` via `ExtractResourceAPI` (`.build/googleapis.pb` descriptor inspection), `InspectKRMPackage` (`apis/` + `pkg/controller/direct/` AST inspection), and `InspectGAPICClient` (`go.mod` GAPIC SDK AST inspection). |
| **Empirical Validation ("94.6% Shadow Eval" vs. True E2E)** | **AST method-name check + 1 E2E test.** Its "94.6% across 371 controllers" in `shadow/evaluator.go` does **not** generate or compile 371 controllers—it only parses existing `_controller.go` files to check if 8 boilerplate method names (`HasRegisterModel`, `HasModelConstructor`, `HasClientConstructor`, `HasAdapterForObject`, `HasFind`, `HasCreate`, `HasUpdate`, `HasDelete`) exist in the AST. Only **1 controller** (`SecureSourceManagerInstance`, which is immutable / `UpdateStyleUnsupported`) was actually rendered and tested in MockGCP E2E. | **Full replacement, compilation, and E2E across all 189 Greenfield controllers.** Replaced all 189 Greenfield controllers on disk: **187 / 189 (98.9%)** compile cleanly into the E2E binary, and **148 / 155 (95.5%)** pass MockGCP + Real GCP (`E2E_GCP_TARGET=real`) E2E tests. |
| **Diffing & Update Semantics** | **Broken for mutable resources.** Generates a per-controller `compare<Resource>` helper calling `mappers.OnlySpecFields` + `common.DiffForTopLevelFields`. See critical bugs detailed below. | **Shared, verified `CompareGreenfieldSpec` & `CompareGreenfieldSpecAndLabels`.** Handles server-default merging (`MergeUnsetFields`), `INPUT_ONLY` proto fields, write-only KRM secrets, `metadata.labels` vs. `spec.labels`, and deterministic sorted `updateMask`. |
| **Domain Modeling & Template Representation** | **Leaky enum/trait hierarchy (`ControllerProfile`).** Tries to compress controller variations into mutually exclusive enums (`IdentityStrategy`, `UpdateStyle`, `DeleteStyle`), which breaks down when orthogonal GCP API quirks combine (see analysis below). | **Orthogonal Inspector Facts (`ResourceAPI` + `KRMPackageInfo`) + Flat Template View-Model (`ControllerArgs`).** Separates proto API facts from Go AST/GAPIC facts at the extraction layer and maps them directly to orthogonal template fields without intermediate translation boilerplate. |

### Critical Bugs Found in PR #14251's Generated Code (`template.go`)

Because PR #14251 only ran an E2E test on `SecureSourceManagerInstance` (an immutable resource with `UpdateStyleUnsupported`), its `Update()` path was never exercised against an actual mutation RPC. Inspecting `dev/tools/controllerbuilder/pkg/controllerprofile/template.go` in PR #14251 reveals **5 bugs** that were encountered and solved in PR #14222 / PR #14224:

1. **Missing `desired.Name = a.id.String()` on `Update()` (`400 InvalidArgument` on AIP-134 RPCs):**
   In PR #14251's `template.go`, when `UpdateHasTopLevelName` is `false` (the standard AIP-134 pattern where the resource name is passed inside `req.<Resource>.Name`), `Update()` never sets `desired.Name = a.id.String()`. Every standard AIP-134 `Update<Resource>` call would fail with `400 InvalidArgument: resource.name is required`.
2. **Server-Defaulted Fields Trigger Perpetual False Update Diffs:**
   PR #14251 uses `mappers.OnlySpecFields(actual, Spec_FromProto, Spec_ToProto)` without merging unset (`nil`) KRM pointer fields from `actualKRM` into `desiredKRM` (`MergeUnsetFields`). Any field defaulted by GCP on `Create` remains `nil` in `desired` and non-nil in `maskedActual`, causing `DiffForTopLevelFields` to fire a false update on every reconciliation.
3. **`INPUT_ONLY` Proto Fields & Write-Only KRM Secrets Trigger False Diffs:**
   `OnlySpecFields` + `DiffForTopLevelFields` does not strip `(google.api.field_behavior) = INPUT_ONLY` fields or back-fill write-only KRM secret refs (`*refsv1beta1secret.Legacy`, `*refs.SecretKeyRef`) that GCP omits on `Get`.
4. **`LabelsMetadata` Is Defined in `profile.go` But Never Implemented in `template.go`:**
   While `profile.go` defines `LabelsMetadata` (`metadata.labels` vs. `spec.labels`), `template.go` never references `LabelsStyle` in `Create()` or `Update()`—so resources using `metadata.labels` never send or diff labels.
5. **REST vs. gRPC Selection Is Not Inspected from `go.mod`:**
   PR #14251 defaults `UseGRPCClient` to `false` unless manually overridden, without inspecting the GAPIC `cloud.google.com/go/...` package in `go.mod` to detect whether `New*RESTClient` exists or whether `Delete()` returns an LRO vs. `error`.

### Why PR #14251's `ControllerProfile` Trait/Enum Abstraction Is Counterproductive in Practice

PR #14251 introduces an intermediate `ControllerProfile` struct composed of 4 trait sub-structs (`IdentityTrait`, `LifecycleTrait`, `ExecutionTrait`, `TransportTrait`) and typed enums (`IdentityStrategy`, `UpdateStyle`, `DeleteStyle`, `LabelDiffingStyle`). While this looks tidy on paper for a single happy-path resource, it is less adept at representing the actual matrix of controller structures we already support across 189 resources:

1. **Real GCP Proto/KRM Quirks Are Orthogonal Facts, Not Mutually Exclusive Enums:**
   * PR #14251 defines `IdentityStrategy` as a 4-way mutually exclusive enum (`IdentityStandard`, `IdentityServerGenerated`, `IdentityParentOnly`, `IdentityCreateAssignsName`).
   * Across the 189 Greenfield controllers, these dimensions **cross-cut** each other independently:
     * Whether `Get` exists vs. `List`-based `Find()` (`UseListForFind`) is independent of whether the ID is server-generated (`ServerGeneratedID`) or a singleton (`IsSingleton`).
     * Whether `name` is marked `OUTPUT_ONLY` in the proto (`NameIsOutputOnly`) is independent of whether `Create` takes a separate `<resource>_id` request field (`CreateHasIDField`) and whether `Identity` implements `FromExternal` (`IdentityHasFromExternal`).
   * Even in PR #14251 (which only tested 1 controller), the enum abstraction immediately leaked: in `IdentityTrait` and `LifecycleTrait`, the author still had to bolt on 7 ad-hoc booleans (`CreateAcceptsID`, `CreateSetsName`, `UpdateHasTopLevelName`, `DeleteHasEtag`, `DeleteHasForce`, `DeleteHasAllowMissing`, `HasFromExternal`) alongside the enums because real proto RPC signatures do not fit into mutually exclusive buckets.
2. **Unnecessary Double-Translation Boilerplate:**
   * In PR #14224, our architecture already separates the two real source domains cleanly at the extraction layer:
     1. **Proto API facts:** `protoapi.ResourceAPI` (what RPCs, request fields, and field behaviors exist in `.build/googleapis.pb`).
     2. **KRM & GAPIC AST facts:** `scaffold.KRMPackageInfo` (what KRM structs, mappers, refs, Identity methods, and GAPIC clients exist in Go AST).
   * `controller.ControllerArgs` is purely the flat view-model consumed by `text/template`. Keeping `ControllerArgs` as flat orthogonal fields makes `text/template` conditionals concise and typo-safe (`{{if .CreateHasIDField}}` rather than verbose string comparisons like `{{if eq .Lifecycle.UpdateStyle "Masked"}}`) and avoids touching an extra intermediate translation layer every time a new proto/KRM pattern is added.

---

## Part 2: Patterns & Lessons Learned Worth Adopting from PR #14251

There are **3 concrete behavioral improvements** in PR #14251's template that we should adopt into PR #14224:

### 1. Fail Loudly on Spec Drift for Immutable Resources (`!HasUpdate`)
* **What PR #14251 does (`template.go`):**
  When a resource has no `Update` RPC (`!HasUpdate`), instead of making `Update()` a silent no-op that just refreshes `status.ObservedState`, it still runs the spec comparison (`CompareGreenfieldSpec`) in `Update()`. If `diff.HasDiff()` is true, it reports the diff and returns an explicit immutability error:
  ```go
  if diff.HasDiff() {
      structuredreporting.ReportDiff(ctx, diff)
      return fmt.Errorf("%s resource is immutable and cannot be updated. Field(s) changed: %v", a.id.String(), diff.FieldIDs())
  }
  ```
* **Why we should adopt it:**
  Currently, our `controller.go` template treats `!HasUpdate` as a silent status refresh if a user edits `.spec` on an immutable resource. Running `CompareGreenfieldSpec` and returning an error when `diff.HasDiff()` is true gives users immediate feedback when they attempt to mutate an immutable resource (and matches the `SecureSourceManagerInstance` E2E test expectation).

### 2. Working `AdapterForURL` & Fixing the `Export()` Map Overwrite Bug
* **What PR #14251 does (`template.go`):**
  1. In `Export()`, the legacy `controllerbuilder` template sets `u.SetName(...)` and `u.SetGroupVersionKind(...)` *before* `u.Object = uObj`, which overwrites the `metadata` and `apiVersion`/`kind` keys in `u.Object`! PR #14251 fixes the ordering so `u.Object = uObj` happens first, followed by `u.SetName(a.id.ID())` and `u.SetGroupVersionKind(...)`.
  2. When the identity type has `FromExternal(url)` (which we already detect via `IdentityInfo.HasFromExternal` in `krm_inspector.go`), PR #14251 generates a working `AdapterForURL` instead of `return nil, nil`.
* **Why we should adopt it:**
  In our A/B E2E evaluation across 155 resources, **1 of our 7 remaining MockGCP failures** (`BillingAccount`) failed solely because its test fixture has an `_exported.yaml` assertion and our template's `AdapterForURL` returned `nil, nil`. Adopting this pattern makes `Export()` work out of the box for every `IdentityV2` controller and turns `BillingAccount` green (raising our E2E pass rate from `148 / 155` to `149 / 155 = 96.1%`).

### 3. Guarding LRO `Delete()` `op.Wait(ctx)` Against REST `"missing \"value\" field"` Errors
* **What PR #14251 does (`template.go`):**
  When `Delete` returns an LRO (`DeleteIsLRO = true`) and runs over HTTP/REST (`New*RESTClient`), certain GCP services (such as `SecureSourceManagerInstance`) return a completed LRO response body for `google.protobuf.Empty` that triggers a protobuf JSON unmarshaling error in `op.Wait(ctx)` containing `"missing \"value\" field"`. PR #14251 guards `op.Wait(ctx)` in `Delete()` with:
  ```go
  if err := op.Wait(ctx); err != nil && !strings.Contains(err.Error(), "missing \"value\" field") {
      return false, fmt.Errorf("waiting delete %s %s: %w", "{{.Kind}}", a.id, err)
  }
  ```
* **Why we should adopt it:**
  Since our Canonical Standard 1 mandates REST clients (`New*RESTClient`) by default, guarding LRO `Delete()` against this known `google.protobuf.Empty` REST LRO unmarshaling quirk prevents false `Delete()` failures on services like `SecureSourceManager`.

---

## Appendix: Empirical A/B Evaluation Summary for PR #14222 + PR #14224

To ensure an objective benchmark, we replaced all **189** Greenfield GAPIC proto-backed controllers on disk using our deterministic generator (`PR #14222` + `PR #14224`) and ran the full compilation and E2E test suite (`155` resources with basic E2E fixtures):

| Metric | Pass Rate | Count | Notes |
| :--- | :---: | :---: | :--- |
| **Tier 1 — Compilation** | **98.9%** | `187 / 189` | Only 2 non-standard multi-client/flattened-request outliers (`ApigeeEnvgroupAttachment`, `CloudBuildWorkerPool`). |
| **Tier 2 — Standards-Adjusted E2E Pass Rate (MockGCP + Real GCP)** | **95.5%** | `148 / 155` | Verified on MockGCP (`144 / 155`) plus Real GCP (`E2E_GCP_TARGET=real` on `DataplexAspectType`, `TasksQueue`, `VertexAIFeaturestore`, and `BigQueryConnectionConnection`—all 4 passed 100% on Real GCP, confirming their MockGCP `updateMask` failures were MockGCP simulator limitations). |
| **Projected Tier 2 with `AdapterForURL` + `Export()` Fix (Lesson #2 above)** | **96.1%** | `149 / 155` | Fixes `BillingAccount` (`_exported.yaml`), leaving only **6** genuinely bespoke controllers (`3.9%`) out of 155. |
