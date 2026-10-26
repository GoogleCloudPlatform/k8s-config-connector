# Greenfield Bulk Acceleration: PR & Workstream Index

This document indexes the Pull Requests, branches, design documents, and empirical validation artifacts for the **Greenfield Bulk Acceleration** initiative.

---

## 1. Documentation & Empirical Reports

| PR | Branch | File / Scope | Summary |
| :--- | :--- | :--- | :--- |
| **[#14416](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14416)** | `docs-greenfield-bulk-acceleration-plan` | `docs/ai/greenfield-bulk-acceleration-plan.md` | **Primary Report & Acceleration Plan:** End-to-end empirical validation across 34 greenfield resources (Batch 1 + Batch 2), root-cause analysis of generator/runtime bugs, parallel E2E test execution benchmarks (`2.6x–7.3x` speedup), `AutoResolveJudgement` evaluation across 275 existing resources (`91.6%` exact-path reproduction, `-45.7%` queued fields, `1.65%` false-positive rate) and 31 new resources (`-81.4%` targeted queue reduction), and next-step team recommendations. |
| **[#14253](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14253)** | `docs-controller-generation-evaluation` | `docs/designs/controller-generation-evaluation.md` | Architectural and empirical comparison of deterministic controller generation (`--prepopulate-resource`). |
| **[#14130](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14130)** | `docs-declarative-resource-criteria` | `docs/ai/declarative-resource-criteria.md` | Criteria for classifying declarative vs. non-declarative GCP resources during candidate sourcing. |

---

## 2. Core Generator, Runtime & Test Runner PRs (Ready for Review)

These standalone PRs isolate the reusable `controllerbuilder`, runtime reconciler, `refs.Ref`, and E2E test runner improvements discovered during the 34-resource validation:

| PR | Branch | Area | Summary |
| :--- | :--- | :--- | :--- |
| **[#14415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14415)** | `exp-judgement-autoresolve` | `dev/tools/controllerbuilder` | **Judgement Queue Auto-Resolution (`AutoResolveJudgement`):** Automatically resolves 5 deterministic `judgement_queue.yaml` patterns during `generate-types --prepopulate-resource` (Rule 1: `*parent.ProjectAndLocationRef`; Rule 2: child `<Parent>Ref` when `Normalize()` exists, including cross-version sibling scan; Rule 3: omit top-level `.spec.labels`; Rule 4: route `etag` & output-only comment fields to `ObservedState`; Rule 5: authoritative `google.api.resource_reference` $\rightarrow$ `*Ref` in top-level and nested structs). |
| **[#14414](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14414)** | `feat-e2e-parallel` | `tests/e2e`, `pkg/test` | **Parallel E2E Fixture Execution (`E2E_PARALLEL=true`):** Opt-in `t.Parallel()` execution for `TestAllInSeries` with per-subtest HTTP/gRPC sink routing via User-Agent context (`WithUserAgentSink`), per-subtest temp output directories, and mutex-guarded global registration (`2.6x–3.7x` RealGCP speedup; `6.2x–7.3x` MockGCP speedup). |
| **[#14413](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14413)** | `fix-greenfield-slice-elem-merge` | `pkg/controller/direct/common` | **Recursive Slice Element Server-Default Merge:** Updates `MergeGreenfieldServerDefaults` (`greenfield.go`) to recursively merge unset server-defaulted sub-fields on matching struct slice elements, preventing post-Create/Update drift loops on repeated message fields. |
| **[#14412](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14412)** | `fix-gcpwatch-fetcher-rest-options` | `pkg/gcpwatch` | **REST Client Options Fix for `NewIAMClient`:** Replaces `config.GRPCClientOptions()` with `config.RESTClientOptions()` in `pkg/gcpwatch/fetcher.go` so `--E2E_GCP_TARGET=real` does not fail with `WithHTTPClient is incompatible with gRPC dial options`. |
| **[#14411](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14411)** | `feat-refs-v1beta1-normalize` | `apis/refs/v1beta1`, `pkg/controller/direct` | **Shared `v1beta1` `refs.Ref` Normalization & Slice Helpers:** Implements `Normalize(ctx, reader, defaultNamespace)` on shared `v1beta1` reference types (`ProjectRef`, `FolderRef`, `OrganizationRef`, `BillingAccountRef`, `StorageBucketRef`, `PubSubTopicRef`, `SecretManagerSecretRef`, `SecretManagerSecretVersionRef`, `SQLInstanceRef`) and adds `direct.ToStringSlice` / `direct.ToInt64Slice` in `maputils.go`. |
| **[#14317](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14317)** | `fix-proto3-optional-enum-mapper-fuzzer` | `dev/tools/controllerbuilder` | **Proto3 Optional Enums in Mappers & Fuzzers:** Fixes `generate-mapper` and `generate-fuzzer` for `proto3_optional` enum fields (`*pb.Enum` pointer wrappers). |

---

## 3. Empirical Validation & Resource Onboarding PRs (Batch 1, Batch 2 & Experiment 3.b)

| PR | Branch | Cohort | Summary |
| :--- | :--- | :--- | :--- |
| **[#14410](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14410)** | `feat-sample-16-judgement-triage` | **Batch 1 (16 resources)** | Step 1b `judgement_queue.yaml` triage + regenerated types, mappers, controllers, fuzzers, and CRDs for the 16 Batch 1 sampled resources. |
| **[#14440](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14440)** | `exp-realgcp-bulk-validation` | **Batch 1 + 3 Parents (19 resources)** | E2E fixtures (`create.yaml`, `update.yaml`, `dependencies.yaml`), CRDs, and recorded RealGCP `_http.log` golden logs for the 16 Batch 1 resources + 3 prerequisite parent resources (`BigQueryAnalyticsHubDataExchange`, `ConfigDeliveryFleetPackage`, `MigrationCenterGroup`). |
| **[#14441](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14441)** | `exp-realgcp-batch2-15` | **Batch 2 (15 resources)** | Step 1b `judgement_queue.yaml` triage, Step 2 Phase A controllers & fuzzers, CRDs, E2E fixtures, and recorded RealGCP `_http.log` golden logs for the 15 Batch 2 resources across 10 services (`artifactregistry`, `bigqueryanalyticshub`, `configdelivery`, `dataform`, `dataplex`, `dialogflow`, `kms`, `migrationcenter`, `translate`, `vmmigration`). |
| **[#14442](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14442)** | `exp-3b-batch34-autoresolve` | **Experiment 3.b (31 resources)** | Regenerates all 31 Batch 1 + Batch 2 resources across 18 services from scratch using `AutoResolveJudgement` ([#14415](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14415)). All 18 API packages compile with zero manual edits; targeted queue items drop by **-81.4%** (`102 -> 19`) and **15 / 31 resources (48.4%)** require zero code edits in Step 1b. |

---

## 4. Prior Bulk Controller Generation PRs

| PR | Branch | Cohort | Summary |
| :--- | :--- | :--- | :--- |
| **[#14284](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14284)** | `feat-merged-bulk-controllers` | 85 resources | Deterministic direct controllers, mappers, and round-trip fuzzers for 85 greenfield resources. |
| **[#14229](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14229)** | `feat-14066-sample-controllers` | 30 resources | Direct controllers and round-trip fuzzers for 30 greenfield resources from [#14066](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14066). |
| **[#14309](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/14309)** | `feat-remaining-16-controllers` | 13 resources | Direct controllers, mappers, and round-trip fuzzers for 13 higher-touch greenfield resources. |

---

## 5. Summary of Key Empirical Metrics

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
