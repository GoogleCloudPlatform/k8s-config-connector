---
name: reviewgen-greenfield-controller-generated
description: Review criteria for KCC Greenfield controller PRs using deterministic controller (`generate-controller --prepopulate-resource`) and fuzzer (`generate-fuzzer --prepopulate-resource`) generation.
---

# Review Guide for KCC Greenfield Controller (Deterministic Controller & Fuzzer)

Use this review guide when reviewing a **Greenfield Step 2** PR for `<Kind>`. Enforce all of the following criteria and invariants.

---

## 1. Post-Judgement Prerequisite, One-Shot Generation (Not in `generate.sh`), & `static_config.go`
*   **Post-Judgement Check (`apis/<service>/judgement_queue.yaml`)**: Verify that `<Kind>`'s `untriaged-bulk-generation` and all actionable entries in `apis/<service>/judgement_queue.yaml` are `status: resolved`. (Note: in greenfield resources, `Spec.Labels` should be preserved if defined in proto).
*   **No `generate-controller` / `generate-fuzzer` in `generate.sh`**: Verify that `apis/<service>/generate.sh` does **not** invoke `generate-controller` (or `generate-direct-reconciler`) or `generate-fuzzer` with `--resource <Kind>:<ProtoMessage>`. The controller and fuzzer are generated once with direct `controllerbuilder` commands and committed; `generate.sh` runs on every CI build and would overwrite later customizations. **Fail** the PR if it adds such an invocation. If one already existed for `<Kind>`, the PR must remove it.
    *   Do **not** ask the agent to add these invocations back to `generate.sh`, or to "restore deterministic generation" by re-running `generate-controller` over a customized controller: that silently discards the customizations.
*   Verify that `pkg/controller/direct/<service>/<proto_lower>_controller.go` carries the `// +generated:controller` header marker, `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` carries the `// +generated:fuzzer` header marker, and `pkg/controller/resourceconfig/static_config.go` has been updated via `python3 dev/tasks/generate_static_config.py`.
*   Verify `TestGenerateScripts`, `TestDirectResourceFileNaming`, `TestRegisteredTemplatesMatchCAI`, and `TestVerifyLabelsTestCasesExist` pass.
*   Verify `<Kind>` is still in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`: only the pipeline's final Release step removes it.
*   **No Standalone Unit Tests**: Verify that no dedicated `*_test.go` or `*_mappings_test.go` files were added inside `pkg/controller/direct/<service>/`. Direct controllers and mappers are verified centrally via `fuzztests` and E2E fixtures.
*   **Mapper Completeness & Missing References**: Verify that `mapper.generated.go` contains no unmapped `// MISSING:` comments for reference fields that were renamed (e.g. `RepositoryProjectId` -> `RepositoryProjectRef`). All such fields must be explicitly mapped in custom `_FromProto` / `_ToProto` functions in `mapper.go` or `<kind_lower>_mapper.go`.
*   **Dependency Hygiene**: Verify `git diff go.mod mockgcp/go.mod` contains no collateral bumps to unrelated libraries and no `replace` directives.

## 2. Canonical Greenfield Controller Invariants (The 5 Standards)
1.  **GAPIC REST Client Preference**:
    *   Verify `m.client(ctx)` uses `gcp.New<Service>RESTClient(ctx, opts...)` with `m.config.RESTClientOptions()` (falling back to `New<Service>Client` with `m.config.GRPCClientOptions()` only when REST is unsupported or broken on real GCP).
2.  **`common.NormalizeGreenfieldReferences` in `AdapterForObject`**:
    *   Verify `common.NormalizeGreenfieldReferences(ctx, reader, obj)` is called in `AdapterForObject` before `GetIdentity` and before `<Kind>Spec_ToProto`.
3.  **`IdentityV2` / `ServerGeneratedIdentity` Handling**:
    *   Verify `AdapterForObject` calls `obj.GetIdentity(ctx, reader)` and asserts `*krm.<Kind>Identity`.
    *   For `ServerGeneratedIdentity` resources (`Create<Resource>Request` has no `<resource>_id` field), verify `Create` extracts the server-assigned ID via `a.id.FromExternal(created.GetName())` and `setStatus` writes `status.ExternalRef = direct.LazyPtr(a.id.String())`.
4.  **Declared-Spec-Only Diffing (`common.CompareGreenfieldSpec` / `common.CompareGreenfieldSpecAndLabels`)**:
    *   **Strict Rule:** Verify `Update` uses `common.CompareGreenfieldSpec(ctx, a.desired, actual, krm.<Kind>Spec_FromProto)` (or `common.CompareGreenfieldSpecAndLabels(..., a.labels)` when reconciling top-level proto `labels` via `metadata.labels` — never `common.CompareProtoMessage` or `common.CompareBrownfieldSpec`).
    *   Verify `structuredreporting.ReportDiff(ctx, report)` is called when `len(paths) != 0` before sending the `Update<Resource>Request`.
5.  **LRO Waiting, Post-LRO `Get<Resource>`, & Mandatory `setStatus`**:
    *   Verify `Create`, `Update`, and `Delete` wait on any returned LRO (`op.Wait(ctx)`).
    *   Verify `setStatus(u, resolved)` is called at the end of `Create` and on **all** `Update` paths (both when `len(paths) == 0` and after an update RPC completes, or unconditionally on immutable resources without an `Update*` RPC).

## 3. Round-Trip Fuzzer (`<kind_lower>_fuzzer.generated.go`)
*   Verify `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` is present, registers the fuzzer in `init()`, and passes `FOCUS=<Kind> go test ./pkg/fuzztesting/fuzztests -run TestFocusedMappers` and `dev/ci/presubmits/fuzz-roundtrippers`.

## 4. Real-GCP E2E Fixtures & Ratcheting Opt-In
*   Verify `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind_lower>/` contains both `<kind_lower>-minimal` and `<kind_lower>-maximal` fixtures with `create.yaml`, `update.yaml` (if mutable), `_http.log`, and `_audit_probe.log` recorded against **real GCP** (never mockgcp).
*   Verify no `_http.log`, `_audit_probe.log` or golden file was edited by hand: they come from `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`.
*   **Fixture Conventions**: Verify `create.yaml` does not set `spec.resourceID` if its value matches `metadata.name`. Verify `update.yaml` is provided if the resource has any mutable fields, and `dependencies.yaml` is provided if reference fields exist.
*   **Audit Probe Hygiene**: Verify `_audit_probe.log` does not contain ONLY incidental IAM entries for gRPC services where service audit logs are skipped (if so, `_audit_probe.log` must be removed).
*   **Parent fields**: for a child resource, `create.yaml` sets `spec.parentRef` to the parent that `dependencies.yaml` creates (or `external` for an external-only reference type); for a top-level resource, `projectRef` plus `location`. `update.yaml` never changes the parent fields. Project, location and ancestor IDs come from the identity, never from the `cnrm.cloud.google.com/project-id` annotation, `refs.ResolveProjectID` or `refs.GetLocation`.
*   Verify `<Kind>` is not excluded in `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go`.

# Review Comment Template

```markdown
### KCC Greenfield Controller Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Post-Judgement, One-Shot Generation, & `static_config.go`**: [Pass/Fail] - (Confirm `judgement_queue.yaml` is resolved, `generate.sh` does **not** invoke `generate-controller` / `generate-fuzzer` for `<Kind>`, files have `+generated:controller` / `+generated:fuzzer`, `static_config.go` is updated, and `<Kind>` is still in `BLOCKED_KINDS`)
* **Client Creation & Reference Normalization**: [Pass/Fail] - (Confirm GAPIC REST client and `common.NormalizeGreenfieldReferences` in `AdapterForObject`)
* **CompareGreenfieldSpec & Structured Reporting**: [Pass/Fail] - (Confirm `common.CompareGreenfieldSpec` / `CompareGreenfieldSpecAndLabels` and `structuredreporting.ReportDiff` are used in `Update`)
* **LRO Wait & Mandatory Status Updates**: [Pass/Fail] - (Confirm `op.Wait(ctx)` and `setStatus` on both no-op and mutating paths)
* **Fuzzer & Real-GCP Test Coverage**: [Pass/Fail] - (Confirm `<kind>_fuzzer.generated.go`, minimal/maximal fixtures with real-GCP `_http.log` & `_audit_probe.log` (none edited by hand), parent fields set as the spec models them and unchanged in `update.yaml`, and ratcheting enabled)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```

