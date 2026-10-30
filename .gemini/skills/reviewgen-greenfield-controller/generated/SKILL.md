---
name: reviewgen-greenfield-controller-generated
description: Review criteria for KCC Greenfield controller PRs using deterministic controller (`generate-controller --prepopulate-resource`) and fuzzer (`generate-fuzzer --prepopulate-resource`) generation.
---

# Review Guide for KCC Greenfield Controller (Deterministic Controller & Fuzzer)

Use this review guide when reviewing a **Greenfield Step 2** PR for `<Kind>`. Enforce all of the following criteria and invariants.

---

## 1. Post-Judgement Prerequisite, Deterministic `generate.sh`, & `static_config.go`
*   **Post-Judgement Check (`apis/<service>/judgement_queue.yaml`)**: Verify that `<Kind>`'s `untriaged-bulk-generation` and all actionable entries in `apis/<service>/judgement_queue.yaml` are `status: resolved` (and top-level `Labels` is not left in `<Kind>Spec`).
*   Verify that `apis/<service>/generate.sh` invokes both:
    *   `${CONTROLLERBUILDER} generate-controller` (or `generate-direct-reconciler`) `--service <proto.service> --api-version <group>/<version> --prepopulate-resource --resource <Kind>:<ProtoMessage>`
    *   `${CONTROLLERBUILDER} generate-fuzzer --prepopulate-resource --service <proto.service> --api-version <group>/<version> --resource <Kind>:<ProtoMessage>`
*   Verify that `pkg/controller/direct/<service>/<proto_lower>_controller.go` carries the `// +generated:controller` header marker, `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` carries the `// +generated:fuzzer` header marker, and `pkg/controller/resourceconfig/static_config.go` has been updated via `python3 dev/tasks/generate_static_config.py`.
*   Verify `TestGenerateScripts`, `TestDirectResourceFileNaming`, `TestRegisteredTemplatesMatchCAI`, and `TestVerifyLabelsTestCasesExist` pass.

## 2. Canonical Greenfield Controller Invariants (The 5 Standards)
1.  **GAPIC REST Client Preference**:
    *   Verify `m.client(ctx)` uses `gcp.New<Service>RESTClient(ctx, opts...)` with `m.config.RESTClientOptions()` (falling back to `New<Service>Client` with `m.config.GRPCClientOptions()` only when REST is unsupported or broken on real GCP, with the reason stated in the PR).
2.  **`common.NormalizeGreenfieldReferences` in `AdapterForObject`**:
    *   Verify `common.NormalizeGreenfieldReferences(ctx, reader, obj)` is called in `AdapterForObject` before `GetIdentity` and before `<Kind>Spec_ToProto`.
3.  **`IdentityV2` / `ServerGeneratedIdentity` Handling**:
    *   Verify `AdapterForObject` calls `obj.GetIdentity(ctx, reader)` and asserts `*krm.<Kind>Identity`.
    *   Verify the parent in GCP requests comes from the identity (`a.id.ParentString()`), never from spec fields.
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
*   Verify `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind_lower>/` contains both `<kind_lower>-minimal` and `<kind_lower>-maximal` fixtures with `create.yaml`, `update.yaml` (if mutable), `_http.log`, and `_audit_probe.log` recorded against **real GCP** (never mockgcp). The maximal fixture covers the supported spec fields.
*   Verify the fixtures set the parent the way `<Kind>Spec` models it: a child resource sets `spec.parentRef` (to the parent created in `dependencies.yaml`, or `external:` for an external-only parent), and no `update.yaml` changes the parent fields (`parentRef`, or `projectRef`/`location`).
*   Verify `<Kind>` is not excluded in `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go`.
*   Verify the PR description states the GCP project and PDT start/end times of the `hack/record-gcp` run, and that the `gcloud logging read` audit-log output is posted as a PR comment.

# Review Comment Template

```markdown
### KCC Greenfield Controller Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Post-Judgement, Deterministic `generate.sh`, & `static_config.go`**: [Pass/Fail] - (Confirm `judgement_queue.yaml` is resolved, `generate.sh` invokes `generate-controller` & `generate-fuzzer` with `--prepopulate-resource`, files have `+generated:controller` / `+generated:fuzzer`, and `static_config.go` is updated)
* **Client Creation & Reference Normalization**: [Pass/Fail] - (Confirm GAPIC REST client and `common.NormalizeGreenfieldReferences` in `AdapterForObject`)
* **CompareGreenfieldSpec & Structured Reporting**: [Pass/Fail] - (Confirm `common.CompareGreenfieldSpec` / `CompareGreenfieldSpecAndLabels` and `structuredreporting.ReportDiff` are used in `Update`)
* **LRO Wait & Mandatory Status Updates**: [Pass/Fail] - (Confirm `op.Wait(ctx)` and `setStatus` on both no-op and mutating paths)
* **Fuzzer & Real-GCP Test Coverage**: [Pass/Fail] - (Confirm `<kind>_fuzzer.generated.go`, minimal/maximal fixtures with real-GCP `_http.log` & `_audit_probe.log`, and ratcheting enabled)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```

