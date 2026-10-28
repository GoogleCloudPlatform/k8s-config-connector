# Greenfield Bulk Acceleration: PR & Workstream Index

This document indexes the Pull Requests, branches, design documents, and empirical validation artifacts for the **Greenfield Bulk Acceleration** initiative, along with the exact resource and controller accounting against the **272-resource Code Purple commitment** (`80%` of the `341` total CP Purple resources in [`hack/tools/greenfield/cp_resources_list.json`](../../hack/tools/greenfield/cp_resources_list.json)).

---

## 1. Code Purple Coverage Summary (`272` Target / `341` Full List)

* **Headline Breakdown (`272` Commitment):** Bulk generation covers **~50%** (`48.2%` merged / `51.5%` with active pending PRs) of the 272-resource Code Purple commitment, joining completed baseline (**~20%**), in-flight workflows (**~18%**), and newly merged / validation-pending workflows (**~13%**) for **~100% (`101.5%`)** total coverage.

| Cohort | Types (`<kind>_types.go`) | Controllers (`static_config.go`) | % of `272` Target (Types / Ctrls) | % of `341` Full List (Types / Ctrls) | Notes |
| :--- | ---: | ---: | ---: | ---: | :--- |
| **1. Completed Baseline** | **~55** | **~55** | **20.2% / 20.2%** | **16.1% / 16.1%** | Pre-existing / baseline implemented resources in `static_config.go` |
| **2. Validation-Pending / Newly Merged Workflows** | **35** | **35** | **12.9% / 12.9%** | **10.3% / 10.3%** | Individual `workflow/greenfield` resources with both types and controllers merged on `master` |
| **3. In-Flight `workflow/greenfield` Workflows** | **46** *(23 merged + 23 open)* | **0** *(awaiting Step 2)* | **16.9% / 0%** *(8.5% merged)* | **13.5% / 0%** *(6.7% merged)* | `23` with Step 1 types merged on `master` (`BLOCKED_KINDS`) + `23` with open Step 1 issues/PRs (incl. 19 in [#14072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14072)) |
| **4. Bulk Generation — Merged (`master`)** | **131** | **127** | **48.2% / 46.7%** | **38.4% / 37.2%** | `81` direct `cp_resources_list.json` matches + `50` GA Full-CRUD declarative backfills replacing dropped non-GA / non-declarative entries |
| **5. Bulk Generation — Active Pending** | **+9** *(140 total)* | **+1** *(128 total)* | **+3.3% / +0.4%** *(51.5% / 47.1% total)* | **+2.6% / +0.3%** *(41.1% / 37.5% total)* | [#14271](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14271) (`+5`), [#14132](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14132)–[#14134](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14134) (`+3`), [#14441](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14441)/[#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442) (`+1` `DataplexGlossary`) |
| **TOTAL (Merged on `master` Only)** | **244** | **217** | **89.7% / 79.8%** | **71.6% / 63.6%** | All merged types (incl. `BLOCKED_KINDS`) and merged controllers in `static_config.go` |
| **TOTAL (Merged + Pending + In-Flight)** | **276** | **218** | **101.5% / 80.1%** | **80.9% / 63.9%** | *(`331` types / `121.7%` if also counting held-back [#14072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14072), [#14088](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14088), [#13966](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13966))* |

---

## 2. Documentation & Empirical Reports

| PR | Status | Branch | File / Scope | Summary |
| :--- | :---: | :--- | :--- | :--- |
| **[#14416](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14416)** | OPEN | `docs-greenfield-bulk-acceleration-plan` | `docs/ai/greenfield-bulk-acceleration-plan.md` | **Primary Report & Acceleration Plan:** End-to-end empirical validation across 34 greenfield resources (Batch 1 + Batch 2), root-cause analysis of generator/runtime bugs, parallel E2E test execution benchmarks (`2.6x–7.3x` speedup), `AutoResolveJudgement` evaluation across 275 existing resources (`91.6%` exact-path reproduction, `-45.7%` queued fields, `1.65%` false-positive rate) and 31 new resources (`-81.4%` targeted queue reduction), and next-step team recommendations. |
| **[#14253](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14253)** | OPEN | `docs-controller-generation-evaluation` | `docs/designs/controller-generation-evaluation.md` | Architectural and empirical comparison of deterministic controller generation (`--prepopulate-resource`). |
| **[#14130](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14130)** | OPEN | `docs-declarative-resource-criteria` | `docs/ai/declarative-resource-criteria.md` | Criteria for classifying declarative vs. non-declarative GCP resources during candidate sourcing. |

---

## 3. Core Generator, Runtime & Test Runner PRs

### 3.1 Merged Foundation & Runtime PRs (`master`)

| PR | Status | Merged | Area | Summary |
| :--- | :---: | :---: | :--- | :--- |
| **[#14449](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14449)** | MERGED | 2026-10-26 | `dev/tools/controllerbuilder` | Updates controller template copyright header to 2026. |
| **[#14445](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14445)** | MERGED | 2026-10-26 | `.github/workflows` | Reduces peak disk and `/tmp` quota usage in `golangci-lint` job after bulk generation expanded repository file count. |
| **[#14413](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14413)** | MERGED | 2026-10-26 | `pkg/controller/direct/common` | **Recursive Slice Element Server-Default Merge:** Updates `MergeGreenfieldServerDefaults` (`greenfield.go`) to recursively merge unset server-defaulted sub-fields on matching struct slice elements and compare KRM `*Ref` structs by `.External`. |
| **[#14111](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14111)** | MERGED | 2026-10-26 | `dev/tools/controllerbuilder` | Rejects scaffolding non-declarative / data-plane resources by default (`nondeclarative.go`). |
| **[#14267](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14267)** | MERGED | 2026-10-23 | `.gemini/skills` | Adds greenfield controller/fuzzer generation, bulk generator (`kcc-bulk-controller-generator`), and controller-aware judgement skills. |
| **[#14224](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14224)** | MERGED | 2026-10-23 | `dev/tools/controllerbuilder` | **Deterministic Controller Generation (`generate-controller --prepopulate-resource`):** Proto-driven direct controller codegen supporting gRPC/REST clients, LROs, parent references, and custom ID fields. |
| **[#13561](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13561)** | MERGED | 2026-10-23 | `dev/tools/controllerbuilder` | **Deterministic Fuzzer Generation (`generate-fuzzer --prepopulate-resource`):** Generates round-trip spec/status fuzzers with automatic identity/output-only `UnimplementedFields` classification. |
| **[#14222](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14222)** | MERGED | 2026-10-22 | `pkg/controller/direct/common` | **Greenfield UUS Spec-Diff (`CompareGreenfieldSpec` / `MergeGreenfieldServerDefaults`):** Proto-agnostic spec comparison and server-default merging for generated direct controllers. |
| **[#14191](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14191)** | MERGED | 2026-10-21 | `pkg/controller/direct/networkservices` | Preserves empty nested structs in `EdgeCacheService` `ToProto` mappers (surfaced during round-trip fuzzer generation). |
| **[#14173](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14173)** | MERGED | 2026-10-21 | `dev/tools/controllerbuilder` | Blocks scaffolding deprecated resources without `--allow-deprecated`. |

### 3.2 Open Acceleration & Systemic Fix PRs (Ready for Review)

| PR | Status | Branch | Area | Summary |
| :--- | :---: | :--- | :--- | :--- |
| **[#14415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14415)** | OPEN | `exp-judgement-autoresolve` | `dev/tools/controllerbuilder` | **Judgement Queue Auto-Resolution (`AutoResolveJudgement`):** Automatically resolves 5 deterministic `judgement_queue.yaml` patterns during `generate-types --prepopulate-resource` (`98.35%` precision / `1.65%` FP rate across 262 resources). |
| **[#14414](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14414)** | OPEN | `feat-e2e-parallel` | `tests/e2e`, `pkg/test` | **Parallel E2E Fixture Execution (`E2E_PARALLEL=true`):** Opt-in `t.Parallel()` execution for `TestAllInSeries` with per-subtest HTTP/gRPC sink routing (`2.6x–3.7x` RealGCP speedup; `6.2x–7.3x` MockGCP speedup). |
| **[#14412](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14412)** | OPEN | `fix-gcpwatch-fetcher-rest-options` | `pkg/gcpwatch` | **REST Client Options Fix for `NewIAMClient`:** Replaces `config.GRPCClientOptions()` with `config.RESTClientOptions()` in `pkg/gcpwatch/fetcher.go` so `--E2E_GCP_TARGET=real` works with IAM watch clients. |
| **[#14411](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14411)** | OPEN | `feat-refs-v1beta1-normalize` | `apis/refs/v1beta1`, `pkg/controller/direct` | **Shared `v1beta1` `refs.Ref` Normalization & Slice Helpers:** Implements `Normalize(ctx, reader, defaultNamespace)` on shared `v1beta1` reference types and adds `direct.ToStringSlice` / `direct.ToInt64Slice`. |
| **[#14317](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14317)** | OPEN | `fix-proto3-optional-enum-mapper-fuzzer` | `dev/tools/controllerbuilder` | **Proto3 Optional Enums in Mappers & Fuzzers:** Fixes `generate-mapper` (`direct.EnumPtr_ToProto`) and `generate-fuzzer` for `proto3_optional` enum fields (`*pb.Enum` pointer wrappers). |
| **[#14124](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14124)** | OPEN | `chore-bulk-resource-candidate-sourcing` | `hack/tools/greenfield`, `.gemini/skills` | **Candidate Sourcing Workflow:** Updates `calculate_coverage.py` to prefer GA proto packages over alpha/beta and detect Step-1-scaffolded KCC types. |

---

## 4. Bulk KRM Types, Identities, References & Mappers PRs (`131` Merged / `140` Active / `195` All Open)

| PR | Status | Date | Net-New Types (`<kind>_types.go`) | Summary |
| :--- | :---: | :---: | ---: | :--- |
| **[#13407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13407)** | MERGED | 2026-10-17 | **1** | `ChronicleWatchlist` scaffolded with bulk generation flags. |
| **[#14066](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14066)** | MERGED | 2026-10-20 | **47** | Bulk `v1alpha1` types, identities, references, and mappers for 47 declarative GCP resources across 35 services *(49 `*_types.go` files incl. 2 `recursive_types.go`)*. |
| **[#14125](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14125)** | MERGED | 2026-10-20 | **37** | `[1 of 3 — Category 1: Zero-Touch]` Bulk `v1alpha1` types, identities, references, and mappers (18 services, 37 resources). |
| **[#14126](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14126)** | MERGED | 2026-10-21 | **11** | `[2 of 3 — Category 2: Initial Setup]` Bulk `v1alpha1` types, identities, references, and mappers (4 services, 11 resources). |
| **[#14127](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14127)** | MERGED | 2026-10-21 | **35** | `[3 of 3 — Category 3: Major Surgery]` Bulk `v1alpha1` types, identities, references, and mappers (13 services, 35 resources). |
| **Merged Types Subtotal** | | | **131** | *(`130` from bulk batches + `1` from [#13407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13407))* |
| **[#14271](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14271)** | OPEN | 2026-10-22 | **+5** | Prepopulates 5 declarative Full-CRUD greenfield resources (`apihub`, `apphub`, `networksecurity`) — brings bulk declarative types from `130 -> 135`. |
| **[#14132](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14132)** | OPEN | 2026-10-20 | **+1** | `TelcoAutomationEdgeSLM` `v1alpha1` types, identity, reference, and mapper. |
| **[#14133](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14133)** | OPEN | 2026-10-20 | **+1** | `IAMPolicyBinding` `v1alpha1` types, identity, reference, and mapper. |
| **[#14134](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14134)** | OPEN | 2026-10-20 | **+1** | `AutoMLModel` `v1alpha1` types, identity, reference, and mapper. |
| **[#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442)** | OPEN | 2026-10-26 | **+1** | Adds `DataplexGlossary` types alongside regenerating 30 Batch 1 + Batch 2 types via `AutoResolveJudgement`. |
| **Active Track Subtotal (Merged + Active Pending)** | | | **140** | *(`131` merged + `9` active pending)* |
| **[#14072](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14072)** *(Held)* | OPEN | 2026-10-17 | **+19** | Bulk `v1alpha1` types, identities, and mappers for 19 resources overlapping open `workflow/greenfield` issues. |
| **[#14088](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14088)** *(Held)* | OPEN | 2026-10-19 | **+27** | Bulk `v1alpha1` types for 27 non-declarative / data-plane GCP resources *(31 `*_types.go` files incl. 4 `recursive_types.go`)*. |
| **[#13966](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13966)** *(Demo)* | OPEN | 2026-10-13 | **+9** | Demo bulk-generated types and identities for 9 non-GA resources (`ces`, `discoveryengine`, `gdchardwaremanagement`). |
| **Grand Total (Merged + All Open)** | | | **195** | *(`131` merged + `64` across all open PRs)* |

---

## 5. Bulk Direct Controller & Fuzzer Generation PRs (`127` Merged / `128` Active)

| PR | Status | Merged | Net-New Controllers (`<kind>_controller.go`) | Summary |
| :--- | :---: | :---: | ---: | :--- |
| **[#14229](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14229)** | MERGED | 2026-10-26 | **30** | Direct controllers and round-trip fuzzers for 30 greenfield resources from [#14066](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14066). |
| **[#14284](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14284)** | MERGED | 2026-10-26 | **84** | Deterministic direct controllers, mappers, and round-trip fuzzers for 84 greenfield resources *(85 `static_config.go` registrations)*. |
| **[#14309](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14309)** | MERGED | 2026-10-27 | **13** | Direct controllers, mappers, and round-trip fuzzers for 13 higher-touch greenfield resources. |
| **Merged Controllers Subtotal** | | | **127** | *(Covers 127 of the 131 merged resource types; only `BlockchainNodeEngineBlockchainNode`, `IAMAccessPolicy`, `StorageRapidCache`, and `TPUQueuedResource` lack a controller on `master`)* |
| **[#14441](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14441)** | OPEN | — | **+1** (`DataplexGlossary`) | Adds prerequisite `DataplexGlossary` controller alongside Step 1b + RealGCP E2E updates for 15 Batch 2 controllers. |
| **Total Controllers (Merged + Pending)** | | | **128** | |

---

## 6. Empirical Validation & Resource Onboarding PRs (Batch 1, Batch 2 & Experiment 3.b)

| PR | Status | Branch | Cohort | Summary |
| :--- | :---: | :--- | :--- | :--- |
| **[#14410](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14410)** | OPEN | `feat-sample-16-judgement-triage` | **Batch 1 (16 resources)** | Step 1b `judgement_queue.yaml` triage + regenerated types, mappers, controllers, fuzzers, and CRDs for the 16 Batch 1 sampled resources. |
| **[#14440](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14440)** | OPEN | `exp-realgcp-bulk-validation` | **Batch 1 + 3 Parents (19 resources)** | E2E fixtures (`create.yaml`, `update.yaml`, `dependencies.yaml`), CRDs, and recorded RealGCP `_http.log` golden logs for the 16 Batch 1 resources + 3 prerequisite parent resources (`BigQueryAnalyticsHubDataExchange`, `ConfigDeliveryFleetPackage`, `MigrationCenterGroup`). |
| **[#14441](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14441)** | OPEN | `exp-realgcp-batch2-15` | **Batch 2 (15 resources)** | Step 1b `judgement_queue.yaml` triage, Step 2 Phase A controllers & fuzzers, CRDs, E2E fixtures, and recorded RealGCP `_http.log` golden logs for the 15 Batch 2 resources across 10 services (`artifactregistry`, `bigqueryanalyticshub`, `configdelivery`, `dataform`, `dataplex`, `dialogflow`, `kms`, `migrationcenter`, `translate`, `vmmigration`). |
| **[#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442)** | OPEN | `exp-3b-batch34-autoresolve` | **Experiment 3.b (31 resources)** | Regenerates all 31 Batch 1 + Batch 2 resources across 18 services from scratch using `AutoResolveJudgement` ([#14415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14415)). All 18 API packages compile with zero manual edits; targeted queue items drop by **-81.4%** (`102 -> 19`) and **15 / 31 resources (48.4%)** require zero code edits in Step 1b. |

---

## 7. Summary of Key Empirical Metrics

1. **RealGCP Validation (34 Resources — [#14440](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14440), [#14441](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14441)):**
   - **26 / 34 resources (76.5%)** passed full Create $\rightarrow$ Get $\rightarrow$ Update $\rightarrow$ Delete against live GCP APIs and recorded authentic `_http.log` golden traces.
   - The remaining **8 / 34 (23.5%)** reached live GCP APIs and were blocked only by external project/org prerequisites (org hierarchy, Cloud Armor Enterprise tier, private PSC/VPC peering, or Dialogflow Console ToS), with **0** controller or mapper panics.
2. **Parallel E2E Execution (`E2E_PARALLEL=true` — [#14414](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14414)):**
   - **RealGCP (5 LRO-heavy resources):** `166.96s` sequential $\rightarrow$ **`45.40s` parallel (`3.68x` speedup)**, bounded by the single slowest resource (`43.25s`).
   - **MockGCP (13 Batch 1 resources):** `80.87s` sequential $\rightarrow$ **`13.09s` parallel (`6.18x` speedup)**.
3. **Judgement Queue Auto-Resolution (`AutoResolveJudgement` — [#14415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14415), [#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442)):**
   - **275-Resource Existing Cohort Coverage:**
     - **Overall Coverage:** `97.7%` (`98.0%` excl. unproducible `googleapis` deltas).
     - **1. Reproduced at exact same path:** `9,948 (89.8%)` $\rightarrow$ **`10,152 (91.6%)` (`+204` fields / `+1.8%`**, or `90.1% -> 91.9%` excl. `googleapis` deltas).
     - **2. Logged in review queue:** `394 (3.6%)` $\rightarrow$ **`214 (1.9%)` (`-180` fields / `-45.7%` reduction)**.
     - **3. Generated differently (unlogged):** `488 (4.4%)` $\rightarrow$ **`456 (4.1%)` (`-32` fields)**.
     - **Total False-Positive Rate:** **`1.65%` (`7 / 424` decisions; `98.35%` precision)**.
   - **31-Resource New Cohort (Experiment 3.b — [#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442)):**
     - **Targeted Queue Categories:** `102` $\rightarrow$ **`19` (`-81.4%` reduction)**.
     - **Total Open Queue Items:** `175` $\rightarrow$ **`92` (`-47.4%` reduction)**.
     - **Zero-Touch Step 1b Resources:** **15 / 31 (48.4%)** require zero code edits in Step 1b.

---

## Appendix A: Complete Chronological List of All 61 Merged Bulk-Gen PRs (Sept 23 – Oct 27, 2026)

| # | PR | Created | Merged | Title |
| ---: | :--- | :---: | :---: | :--- |
| 1 | **[#14449](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14449)** | 2026-10-26 | 2026-10-26 | `fix(controllerbuilder): update controller template copyright header to 2026` |
| 2 | **[#14445](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14445)** | 2026-10-26 | 2026-10-26 | `fix(ci): reduce peak disk and /tmp quota usage in golangci-lint job` |
| 3 | **[#14413](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14413)** | 2026-10-25 | 2026-10-26 | `greenfield: merge unset server defaults on matching slice elements` |
| 4 | **[#14309](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14309)** | 2026-10-23 | 2026-10-27 | `feat(direct): generate mappers, controllers, and fuzzers for 13 higher-touch greenfield resources` |
| 5 | **[#14284](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14284)** | 2026-10-22 | 2026-10-26 | `feat(direct): generate direct controllers and fuzzers for 85 greenfield resources` |
| 6 | **[#14267](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14267)** | 2026-10-21 | 2026-10-23 | `skills: add greenfield controller/fuzzer generation, bulk generator, and controller-aware judgement skills` |
| 7 | **[#14229](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14229)** | 2026-10-21 | 2026-10-26 | `feat(direct): generate direct controllers and fuzzers for 30 greenfield resources (#14066)` |
| 8 | **[#14224](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14224)** | 2026-10-21 | 2026-10-23 | `controllerbuilder: implement deterministic controller generation (--prepopulate-resource)` |
| 9 | **[#14222](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14222)** | 2026-10-21 | 2026-10-22 | `direct/common: add CompareGreenfieldSpec and CompareGreenfieldSpecAndLabels (UUS spec-diff)` |
| 10 | **[#14191](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14191)** | 2026-10-20 | 2026-10-21 | `fix(networkservices): preserve empty nested structs in EdgeCacheService ToProto mappers` |
| 11 | **[#14173](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14173)** | 2026-10-20 | 2026-10-21 | `feat(controllerbuilder): block scaffolding new deprecated resources without --allow-deprecated` |
| 12 | **[#14131](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14131)** | 2026-10-20 | 2026-10-21 | `skills: enforce 4-space quoted BLOCKED_KINDS formatting and align Step 1 prepopulate checks` |
| 13 | **[#14127](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14127)** | 2026-10-20 | 2026-10-21 | `[3 of 3 — Category 3: Major Surgery] apis: bulk generate v1alpha1 types, identities, references, and mappers (13 services, 35 resources)` |
| 14 | **[#14126](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14126)** | 2026-10-20 | 2026-10-21 | `[2 of 3 — Category 2: Initial Setup] apis: bulk generate v1alpha1 types, identities, references, and mappers (4 services, 11 resources)` |
| 15 | **[#14125](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14125)** | 2026-10-20 | 2026-10-20 | `[1 of 3 — Category 1: Zero-Touch] apis: bulk generate v1alpha1 types, identities, references, and mappers (18 services, 37 resources)` |
| 16 | **[#14123](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14123)** | 2026-10-20 | 2026-10-20 | `fix(controllerbuilder,apichecks): recover and guard clobbered judgement_queue.yaml entries` |
| 17 | **[#14111](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14111)** | 2026-10-20 | 2026-10-26 | `controllerbuilder: reject scaffolding non-declarative resources by default` |
| 18 | **[#14075](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14075)** | 2026-10-19 | 2026-10-19 | `controllerbuilder: expand restRootOverrides for irregular REST docs roots` |
| 19 | **[#14071](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14071)** | 2026-10-17 | 2026-10-20 | `controllerbuilder: support additive file merging in generatorBase` |
| 20 | **[#14070](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14070)** | 2026-10-17 | 2026-10-20 | `controllerbuilder: preserve multi-service mapper import aliases and reserve scaffolded types before WriteVisitedMessages` |
| 21 | **[#14069](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14069)** | 2026-10-17 | 2026-10-20 | `controllerbuilder: fix proto kinds, multi-location ancestors, CAI discovery names, and authoritative-proto-reference` |
| 22 | **[#14066](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14066)** | 2026-10-17 | 2026-10-20 | `feat(apis): bulk-generate v1alpha1 types, identities, CRDs, and mappers for 47 net-new GCP resources` |
| 23 | **[#14062](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14062)** | 2026-10-17 | 2026-10-19 | `controllerbuilder: recurse into nested messages in generate-mapper and pluralize shortName` |
| 24 | **[#14045](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14045)** | 2026-10-16 | 2026-10-20 | `skills: add Step 1 prepopulate/ and Step 2 judgement/ greenfield skills and update workflow` |
| 25 | **[#14044](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14044)** | 2026-10-16 | 2026-10-16 | `controllerbuilder: emit Location *string when --prepopulate-resource is set` |
| 26 | **[#14028](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14028)** | 2026-10-16 | 2026-10-16 | `Enforce CRD backward compatibility for triaged Kinds and hand-written proto struct invariants` |
| 27 | **[#14027](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14027)** | 2026-10-16 | 2026-10-16 | `controllerbuilder: emit nested sibling, unsupported, and map judgements per Kind` |
| 28 | **[#14020](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14020)** | 2026-10-16 | 2026-10-16 | `controllerbuilder: sort judgement_queue.yaml canonically by (Kind, Field, Reason)` |
| 29 | **[#14002](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14002)** | 2026-10-15 | 2026-10-15 | `controllerbuilder: fix source-link SHA, judgement queue gating, etag, and ref detection` |
| 30 | **[#13981](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13981)** | 2026-10-14 | 2026-10-16 | `controllerbuilder: consolidate bulk-generation flags into --prepopulate-resource` |
| 31 | **[#13965](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13965)** | 2026-10-13 | 2026-10-14 | `controllerbuilder: queue REQUIRED spec fields that the CRD leaves optional` |
| 32 | **[#13964](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13964)** | 2026-10-13 | 2026-10-14 | `controllerbuilder: plan per-Kind *WithRequired structs for --emit-required-from-proto` |
| 33 | **[#13922](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13922)** | 2026-10-12 | 2026-10-16 | `controllerbuilder: handle all-OUTPUT_ONLY messages, expand judgement queue, and gate non-GA protos` |
| 34 | **[#13893](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13893)** | 2026-10-10 | 2026-10-14 | `gcpurls: supplement CAI asset names with Google API Discovery REST formats` |
| 35 | **[#13887](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13887)** | 2026-10-10 | 2026-10-16 | `controllerbuilder: add generate-identity CLI, conformance helper, and tests` |
| 36 | **[#13883](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13883)** | 2026-10-10 | 2026-10-14 | `controllerbuilder: add deterministic IdentityGenerator for IdentityV2 and Reference` |
| 37 | **[#13853](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13853)** | 2026-10-09 | 2026-10-13 | `controllerbuilder: check unreviewed Kinds again` |
| 38 | **[#13850](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13850)** | 2026-10-09 | 2026-10-10 | `controllerbuilder: flag more fields in the judgement queue` |
| 39 | **[#13849](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13849)** | 2026-10-09 | 2026-10-09 | `controllerbuilder: check nested fields for output-only comments` |
| 40 | **[#13847](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13847)** | 2026-10-09 | 2026-10-09 | `controllerbuilder: leave deprecated fields out of new Kinds` |
| 41 | **[#13809](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13809)** | 2026-10-08 | 2026-10-15 | `controllerbuilder: Tweak automatic generation of +required fields when using --emit-required-from-proto` |
| 42 | **[#13774](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13774)** | 2026-10-07 | 2026-10-15 | `controllerbuilder: don't say --emit-source-links always checks links in the same run` |
| 43 | **[#13561](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13561)** | 2026-10-01 | 2026-10-23 | `controllerbuilder: add deterministic fuzzer generator (--prepopulate-resource)` |
| 44 | **[#13556](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13556)** | 2026-09-30 | 2026-10-07 | `controllerbuilder: link a new types file to its proto and docs (--emit-source-links)` |
| 45 | **[#13531](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13531)** | 2026-09-30 | 2026-10-06 | `apichecks: suppress [refs] per field, only for open reference entries` |
| 46 | **[#13530](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13530)** | 2026-09-30 | 2026-10-01 | `controllerbuilder: write the judgement queue as YAML with a status per entry` |
| 47 | **[#13463](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13463)** | 2026-09-25 | 2026-09-26 | `controllerbuilder: avoid duplicate types and prune unused imports in types.generated.go` |
| 48 | **[#13407](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13407)** | 2026-09-23 | 2026-10-17 | `chronicle: add ChronicleWatchlist using bulk generation flags` |
| 49 | **[#13405](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13405)** | 2026-09-23 | 2026-09-29 | `controllerbuilder: flag an ambiguous --resource` |
| 50 | **[#13404](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13404)** | 2026-09-23 | 2026-09-29 | `controllerbuilder: place server-computed fields in ObservedState, behind a flag` |
| 51 | **[#13403](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13403)** | 2026-09-23 | 2026-09-29 | `controllerbuilder: report fields that look like references, without guessing` |
| 52 | **[#13402](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13402)** | 2026-09-23 | 2026-09-28 | `controllerbuilder: Add flags to match KRM acronym casing and deterministically generate map<string, message> maps` |
| 53 | **[#13401](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13401)** | 2026-09-23 | 2026-09-28 | `scaffold: name a resource's real root, location and parent from its proto` |
| 54 | **[#13400](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13400)** | 2026-09-23 | 2026-09-26 | `apichecks: record and read judgement queue to suppress [refs] during scaffolding` |
| 55 | **[#13399](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13399)** | 2026-09-23 | 2026-09-26 | `controllerbuilder: populate scaffolded resource Spec and ObservedState from proto` |
| 56 | **[#13398](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13398)** | 2026-09-23 | 2026-09-25 | `controllerbuilder: derive +required from field_behavior, behind a flag` |
| 57 | **[#13397](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13397)** | 2026-09-23 | 2026-09-25 | `controllerbuilder: read google.api.resource instead of guessing the resource name` |
| 58 | **[#13396](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13396)** | 2026-09-23 | 2026-09-24 | `apichecks: extract reference rules from crds_test.go into controllerbuilder/pkg/refs` |
| 59 | **[#13395](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13395)** | 2026-09-23 | 2026-09-24 | `generate-types-and-mappers: always rebuild controllerbuilder` |
| 60 | **[#13393](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13393)** | 2026-09-23 | 2026-09-24 | `controllerbuilder: drop imports orphaned by --prune-unused-types` |
| 61 | **[#13392](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/13392)** | 2026-09-23 | 2026-09-24 | `apichecks: catch identity collection segments GCP has never seen` |

