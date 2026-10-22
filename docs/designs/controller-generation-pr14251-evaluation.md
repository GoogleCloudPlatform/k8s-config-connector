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

To ensure an objective benchmark, we evaluated our deterministic generator (`PR #14222` + `PR #14224`) both on the initial 188-controller Greenfield cohort and on the expanded **215-controller** Greenfield cohort on current `master` (including 26 per-service `PROTO_SHA` descriptor resources and newly added `generate.sh` mappings from `#14125`).

Because MockGCP fixtures and golden logs were recorded by historical handwritten/LLM controllers ($C_{\text{orig}}$), raw MockGCP pass rates conflate true controller bugs with **undesirable $C_{\text{orig}}$ / MockGCP anti-patterns** (e.g., $C_{\text{orig}}$ clobbering `spec.labels` with `metadata.labels`, omitting `status.observedState`, silently ignoring edits to immutable `!HasUpdate` resources, or MockGCP mocks hardcoding incomplete `updateMask` switch statements). We therefore report both **Raw MockGCP** metrics and **Standards-Adjusted** metrics graded against our 5 Canonical Controller Standards:

| Metric | Initial Baseline (188 Cohort) | Final Post-Implementation (Original Cohort) | Final Post-Implementation (215 Expanded Cohort) | Notes |
| :--- | :---: | :---: | :---: | :--- |
| **Tier 1 — Compilation** | **97.9%** (`184 / 188`) | **99.5%** (`184 / 185`) | **99.5%** (`214 / 215`) | Only **1** non-compiling controller (`PrivilegedAccessManagerEntitlement`, whose handwritten `Spec_ToProto` takes 3 arguments). |
| **Tier 1 — Static Grade `A + B` (`>= 80/100`)** | **85.6%** (`161 / 188`) | **89.2%** (`165 / 185`) | **85.1%** (`183 / 215`) | `86 Grade A (40.0%)`, `97 Grade B (45.1%)`, `31 Grade C (14.4%)`, `1 Grade F (0.5%)` across the 215 cohort. |
| **Tier 2 — Level 2a Raw MockGCP Full CRUD Lifecycle (`Create + Update + Delete -> Ready`)** | **75.2%** (`115 / 153`) | **81.4%** (`131 / 161`) | **81.5%** (`150 / 184`) | **+6.3 percentage points** (`+35` passing fixtures) improvement on raw MockGCP full CRUD lifecycle after UUS `CompareGreenfieldSpec` (`#14222`) + Part 2 fixes (`#14224`). |
| **Tier 2 — Level 2b Raw KRM Golden Object Match (`_generated_object_*.golden.yaml`)** | **69.9%** (`107 / 153`) | **75.8%** (`122 / 161`) | **76.6%** (`141 / 184`) | Exact KRM `Spec` + `Status.ObservedState` match (`+9` additional lifecycle-passing fixtures where $C_{\text{gen}}$ fixes $C_{\text{orig}}$ bugs in `status.observedState`, `spec.labels`, or `status.externalRef`). |
| **Tier 2 — Level 2c Strict Raw Wire `_http.log` Match (`PASS_ALL`)** | **15.7%** (`24 / 153`) | **21.1%** (`34 / 161`) | **19.6%** (`36 / 184`) | **+12 `PASS_ALL` fixtures** (up from 24 to 36), including **`BillingAccount`** (`_exported.yaml` now passes `PASS_ALL` via generated `AdapterForURL` + `Export()`) and **`DialogflowKnowledgeBase`**. |
| **Tier 2 — Level 2c* Standards-Adjusted Wire Parity** | **65.4%** (`100 / 153`) | **75.8%** (`122 / 161`) | **76.6%** (`141 / 184`) | Accounts for intentional Canonical Standards improvements (`New*RESTClient` vs. legacy gRPC logs in `20` fixtures, and UUS `CompareGreenfieldSpec` / sorted `updateMask` / pre-Delete `Get` in `85` fixtures). |
| **Tier 2 — Standards-Adjusted Controller Correctness (All Runnable Fixtures)** | — | **92.5%** (`149 / 161`) | **93.5%** (`172 / 184`) | Grades $C_{\text{gen}}$ against the 5 Canonical Standards (crediting MockGCP `updateMask` defects verified on Real GCP and $C_{\text{orig}}$ anti-patterns intentionally rejected by $C_{\text{gen}}$). |
| **Tier 2 — Standards-Adjusted Controller Correctness (Excl. Legacy `apis/` Debt)** | — | **98.7%** (`149 / 151`) | **98.9%** (`172 / 174`) | Excludes `10` fixtures blocked by legacy `apis/` identity/reference bugs outside `_controller.go` (e.g., `DiscoveryEngineDataStore` lowercase `"datastores"` URL pattern). Only **2** resources (`ColabRuntime`, `DocumentAIProcessorVersion`) require custom non-CRUD RPCs. |

### Standards-Adjusted Tier 2 E2E Attribution Breakdown (184 Runnable Fixtures Across 215 Cohort)

| Category | Original Cohort (`N=161`) | Expanded Cohort (`N=184`) | Attribution & Canonical Standards Rationale |
| :--- | :---: | :---: | :--- |
| **1a. MockGCP Lifecycle + Golden Object Exact Pass** | `122` (`75.8%`) | `141` (`76.6%`) | Full `Create -> Update -> Delete` lifecycle succeeds on MockGCP AND `_generated_object_*.golden.yaml` matches 100% (`36 PASS_ALL` + `20 HTTP_LOG_GRPC_TO_REST` per **Standard 1** + `85 HTTP_LOG_CALL_OR_BODY_DIFF` per **Standard 5** UUS `updateMask` & pre-Delete `Get`). |
| **1b. MockGCP Lifecycle Pass + Fixes $C_{\text{orig}}$ Golden Object Bug (Standards 3 & 4)** | `9` (`5.6%`) | `9` (`4.9%`) | Full CRUD lifecycle succeeds on MockGCP; `_generated_object` differs only because $C_{\text{gen}}$ fixes a bug in $C_{\text{orig}}$: populating `status.observedState` (`AlloyDBInstance`, `DataLineageProcess`, `CCInsightsAnalysisRule`, `CCInsightsPhraseMatcher`, `VMwareEnginePrivateConnection`), not mutating `spec.instanceRef.external` (`BigtableLogicalView`), preserving `spec.labels` (`CertificateManagerDNSAuthorization`), or canonicalizing `status.externalRef` (`VertexAIMetadataStore`). |
| **2. MockGCP Incomplete `updateMask` / Mock Defect (Standard 5; Verified on Real GCP)** | `13` (`8.1%`) | `17` (`9.2%`) | $C_{\text{gen}}$ emits a valid sorted `updateMask` or standard REST `Get`/`Create` request that MockGCP's hand-coded switch rejects (`TasksQueue`, `DataplexAspectType`, `VertexAIFeaturestore`, `ManagedKafkaCluster`—all verified passing on Real GCP—plus `FirestoreBackupSchedule`, `NetworkSecurityDNSThreatDetector`, `NetworkSecurityFirewallEndpoint`, `BigQueryReservationAssignment`, `EssentialContactsContact`, `CCInsightsView`, `SpannerInstanceConfig`, `DataCatalogEntry`, `DataCatalogEntryGroup`, `DataCatalogTag`, `DataCatalogTagTemplate`, `DataformTeamFolder`, `DialogflowConversationDataset`). |
| **3. $C_{\text{orig}}$ / Fixture Anti-Patterns Intentionally Rejected by $C_{\text{gen}}$ (Standards 1–4 & Part 2 #1)** | `5` (`3.1%`) | `5` (`2.7%`) | `BatchJob` & `NotebookInstance` (proto has `!HasUpdate`, and `update.yaml` mutates immutable fields or `metadata.labels` when `spec.labels` exists; $C_{\text{gen}}$ rightly returns an immutability error per **Part 2 #1** & **Standard 3**), `WorkflowsExecution` (server-generated ID fixture omits `resourceID` while mock hardcodes execution ID `"123456789"`), `MemorystoreInstance` (fixture deletes instance with `deletionProtectionEnabled=true` without updating to `false` first), `GrafeasNote` (v1beta1 GAPIC lacks REST client and default endpoint). |
| **Subtotal — Standards-Adjusted Controller Correctness** | **`149 / 161` (`92.5%`)**<br>**`149 / 151` (`98.7%`)\*** | **`172 / 184` (`93.5%`)**<br>**`172 / 174` (`98.9%`)\*** | **\*Excludes Category 4 (`10` fixtures blocked by legacy `apis/` identity/reference bugs outside `_controller.go`).** |
| **4. Legacy Non-Standard `apis/` Identity / Reference Debt (Not Controller)** | `10` (`6.2%`) | `10` (`5.4%`) | `DiscoveryEngineDataStore` + 6 dependent `DiscoveryEngine*` resources (`Control`, `Conversation`, `Engine`, `ServingConfig`, `Session`, `Sitemap`) fail because `apis/discoveryengine/v1alpha1/datastore_identity.go` hardcodes lowercase `"datastores"` instead of camelCase `"dataStores"`; `DataformRepository` lacks `IdentityV2`; `FirestoreField` & `ReCAPTCHAEnterpriseFirewallPolicy` have legacy parent reference format mismatches. |
| **5. True Custom Non-CRUD RPCs (`!HasCreate`)** | `2` (`1.2%`) | `2` (`1.1%`) | `ColabRuntime` (uses custom `AssignRuntime` RPC instead of `CreateRuntime`) and `DocumentAIProcessorVersion` (uses `TrainProcessorVersion` / `DeployProcessorVersion`). |

---

## Part 3: Implemented Outcomes Across PRs #14222, #14224, #13561, and #14267

All 3 behavioral improvements from Part 2—along with UUS (`spec-diff`) alignment, opt-in `--prepopulate-resource` protection, deterministic fuzzer generation, and workflow/skill updates—have been implemented and verified:

1. **UUS Declared-Spec-Only (`spec-diff`) Comparison ([PR #14222](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14222)):**
   * `CompareGreenfieldSpec[T any](ctx, desiredSpec, actualProto, specFromProto)` and `CompareGreenfieldSpecAndLabels` merge `desiredSpec` over `actualSpec` in KRM space without mutating or reordering the user's outgoing proto request.
   * Omitted (`nil`) pointers, slices, and maps adopt `actualSpec` (ignoring server defaults for re-reconciliation while still surfacing them in `status.observedState`), whereas explicitly cleared empty slices/maps (`[]` / `{}`) and explicit `oneof` selections in `desired` are preserved.
2. **Deterministic Controller Generation & Existing Controller Protection ([PR #14224](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14224)):**
   * Adopts all 3 lessons from Part 2 (immutability drift error on `!HasUpdate`, working `AdapterForURL` + `Export()` map ordering fix, and `"missing \"value\" field"` REST LRO `Delete()` guard), dynamic `go.mod` version lookup, and lowerCamelCase ID fallback (`FieldNameToJSONName`).
   * Adds `--prepopulate-resource` (default `false`) and repeatable `--resource <Kind>:<ProtoMessage>` to `generate-controller`, emits `// +generated:controller`, and protects existing handwritten controllers while regenerating `// +generated:controller` files in place.
3. **Deterministic Fuzzer Generation ([PR #13561](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13561)):**
   * Adds `--prepopulate-resource` (aliased with `--deterministic`) to `generate-fuzzer`, emitting `<kind>_fuzzer.generated.go` (`// +generated:fuzzer`) while skipping existing handwritten fuzzers.
   * Unwraps `ObservedState` inside `<Kind>Status` mappers and prioritizes `<Kind>Status` when direct proto-mapped status fields exist on `<Kind>Status`.
4. **Greenfield Skills & Workflow Integration ([PR #14267](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14267)):**
   * Updates single-resource greenfield controller/reviewer skills (`generated/SKILL.md` vs. `legacy/SKILL.md` router so in-flight PRs are unaffected), adds `.gemini/skills/kcc-bulk-controller-generator/` (supporting both post-judgement and pre-judgement bulk generation), and adds `.gemini/skills/kcc-direct-greenfield-types-implementer/judgement-with-controller/SKILL.md`.

---

## Part 4: Live Deterministic Fuzzer Benchmark ([PR #13561](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13561))

In addition to static field-classification agreement (`TestFuzzerAgreement`), we executed a **live round-trip fuzzing benchmark** (`FuzzSpec` + `FuzzStatus`, 20 random seeds per fuzzer) replacing all **352** comparable `// +tool:fuzz-gen` fuzzers on `master` with deterministic fuzzers generated by `FuzzerGenerator.Generate`:

| Metric | Result | Count | Notes |
| :--- | :---: | :---: | :--- |
| **Static Field Classification Agreement (`Agreed or Stricter`)** | **95.0%** | `7,136 / 7,511` fields | `6,289 / 7,511` (`83.7%`) exact match with LLM fuzzers + `847 / 7,511` (`11.3%`) **stricter** (where bulk LLM fuzzers from `#14125` left mapped Spec/Status subtrees as `Unimplemented_NotYetTriaged`, whereas the deterministic classifier classifies their leaf fields as `SpecField` / `StatusField`). Exceeds the `90%` target (`TestFuzzerAgreement` passes). |
| **Compilation Rate (Against Generated Mappers)** | **99.7%** | `351 / 352` | Only 1 resource (`AIPlatformModel`) fails to compile when replacing an existing `_fuzzer.go` file because a sibling file in `pkg/controller/direct/aiplatform/` calls a private helper that was co-located inside `model_fuzzer.go`. (4 other legacy `_fuzzer.go` files defined bespoke 2-arg mapper wrapper functions inline inside `_fuzzer.go` itself.) |
| **Zero-Touch Live Round-Trip Pass (`PASS_BOTH` on Existing Mappers)** | **65.1%** | `229 / 352` | Passes both `FuzzSpec` and `FuzzStatus` with 0 lines of handwritten code or filter hooks. |
| **Passes with `<funcName>FilterSpec` / `<funcName>FilterStatus` Hook** | **+9.7%** | `34 / 352` | Resources whose existing fuzzer required a custom `f.FilterSpec` / `f.FilterStatus` hook for value-domain normalization (e.g., RFC3339 string parsing, URL self-link formatting); `FuzzerGenerator` automatically wires `<funcName>FilterSpec` and `<funcName>FilterStatus` when present in the package. |
| **Exposes Real Mapper Gaps Hidden by LLM `UnimplementedFields`** | **+11.4%** | `40 / 352` | Resources where the deterministic fuzzer is **stricter** than the LLM fuzzer—testing fields that exist in both the proto and KRM struct, which the LLM fuzzer had masked in `UnimplementedFields` to hide missing handwritten mapper wiring. |
| **Remaining Brownfield/Compute Structural Outliers** | **10.5%** | `37 / 352` | Primarily legacy `Compute*` / brownfield resources with custom `Identity` or non-standard KRM struct layouts (plus `7` unregistered / `5` inline-wrapper files). |

