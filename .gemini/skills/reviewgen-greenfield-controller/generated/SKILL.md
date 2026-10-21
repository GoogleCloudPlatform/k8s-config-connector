---
name: reviewgen-greenfield-controller-generated
description: Review criteria for KCC Greenfield controller PRs using deterministic controller (`generate-controller --prepopulate-resource`) and fuzzer (`generate-fuzzer --prepopulate-resource`) generation.
---

# Review Guide for KCC Greenfield Controller (Deterministic Controller & Fuzzer)

Use this review guide when reviewing a **Greenfield Step 2** PR for `<Kind>`. Enforce all of the following criteria and invariants.

---

## 1. Deterministic `generate.sh` Invocation (`generate-controller` & `generate-fuzzer`)
*   Verify that `apis/<service>/generate.sh` invokes both:
    *   `${CONTROLLERBUILDER} generate-controller --service <proto.service> --api-version <group>/<version> --prepopulate-resource --resource <Kind>:<ProtoMessage>`
    *   `${CONTROLLERBUILDER} generate-fuzzer --prepopulate-resource --service <proto.service> --api-version <group>/<version> --resource <Kind>:<ProtoMessage>`
*   Verify that `pkg/controller/direct/<service>/<proto_lower>_controller.go` carries the `// +generated:controller` header marker and `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` carries the `// +generated:fuzzer` header marker.

## 2. Canonical Greenfield Controller Invariants (The 5 Standards)
1.  **GAPIC REST Client Preference**:
    *   Verify `m.client(ctx)` uses `gcp.New<Service>RESTClient(ctx, opts...)` with `m.config.RESTClientOptions()` (falling back to `New<Service>Client` with `m.config.GRPCClientOptions()` only when REST is unsupported or broken on real GCP).
2.  **`common.NormalizeReferences` in `AdapterForObject`**:
    *   Verify `common.NormalizeReferences(ctx, reader, obj, nil)` is called in `AdapterForObject` before `GetIdentity` and before `<Kind>Spec_ToProto`.
3.  **`IdentityV2` / `ServerGeneratedIdentity` Handling**:
    *   Verify `AdapterForObject` calls `obj.GetIdentity(ctx, reader)` and asserts `*krm.<Kind>Identity`.
    *   For `ServerGeneratedIdentity` resources (`Create<Resource>Request` has no `<resource>_id` field), verify `Create` extracts the server-assigned ID via `a.id.FromExternal(created.GetName())` and `setStatus` writes `status.ExternalRef = direct.LazyPtr(a.id.String())`.
4.  **Declared-Spec-Only Diffing (`common.CompareGreenfieldSpec`)**:
    *   **Strict Rule:** Verify `Update` uses `common.CompareGreenfieldSpec(ctx, a.desired, actual, krm.<Kind>Spec_FromProto)` (never `common.CompareProtoMessage` or `common.CompareBrownfieldSpec`).
    *   Verify `structuredreporting.ReportDiff(ctx, report)` is called when `len(paths) != 0` before sending the `Update<Resource>Request`.
5.  **LRO Waiting, Post-LRO `Get<Resource>`, & Mandatory `setStatus`**:
    *   Verify `Create`, `Update`, and `Delete` wait on any returned LRO (`op.Wait(ctx)`).
    *   Verify `setStatus(u, resolved)` is called at the end of `Create` and on **all** `Update` paths (both when `len(paths) == 0` and after an update RPC completes).

## 3. Round-Trip Fuzzer (`<kind_lower>_fuzzer.generated.go`)
*   Verify `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` is present, registers the fuzzer in `init()`, and passes `dev/ci/presubmits/fuzz-roundtrippers`.

## 4. Real-GCP E2E Fixtures & Ratcheting Opt-In
*   Verify `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind_lower>/` contains both `<kind_lower>-minimal` and `<kind_lower>-maximal` fixtures with `create.yaml`, `update.yaml` (if mutable), `_http.log`, and `_audit_probe.log` recorded against **real GCP** (never mockgcp).
*   Verify `<Kind>` is not excluded in `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go`.

# Review Comment Template

```markdown
### KCC Greenfield Controller Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Deterministic `generate.sh` (`generate-controller` & `generate-fuzzer`)**: [Pass/Fail] - (Confirm `generate.sh` invokes both with `--prepopulate-resource` and files have `+generated:controller` / `+generated:fuzzer`)
* **Client Creation & Reference Normalization**: [Pass/Fail] - (Confirm GAPIC REST client and `common.NormalizeReferences` in `AdapterForObject`)
* **CompareGreenfieldSpec & Structured Reporting**: [Pass/Fail] - (Confirm `common.CompareGreenfieldSpec` and `structuredreporting.ReportDiff` are used in `Update`)
* **LRO Wait & Mandatory Status Updates**: [Pass/Fail] - (Confirm `op.Wait(ctx)` and `setStatus` on both no-op and mutating paths)
* **Fuzzer & Real-GCP Test Coverage**: [Pass/Fail] - (Confirm `<kind>_fuzzer.generated.go`, minimal/maximal fixtures with real-GCP `_http.log` & `_audit_probe.log`, and ratcheting enabled)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
