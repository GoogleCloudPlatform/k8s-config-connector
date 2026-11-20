# ComputeInstance issue 293: networkIpRef and sourceImageRef re-reconciliation

Tracking issue: [#293](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/293)

## Summary

Issue 293 reports ComputeInstances that fail or keep re-applying after creation when the spec uses an address reference or an image family.

* `spec.networkInterface[].networkIpRef.external` set to an address selfLink (`https://www.googleapis.com/compute/v1/projects/P/regions/R/addresses/A`) or path (`projects/P/regions/R/addresses/A`). The instance stays in `UpdateFailed` with `cannot make changes to immutable field(s)`, naming `networkInterface.0.NetworkIp`, or its `network_ip` diff is re-applied on every reconcile.
* `spec.bootDisk.initializeParams.sourceImageRef.external` set to an image family (`cos-cloud/cos-81-lts`) reports an immutable `bootDisk.0.InitializeParams.0.Image` error. This is already fixed on master (see below).

**Root cause (network IP).** GCE accepts an address reference as `networkInterfaces[].networkIP` on insert, but it stores the resolved IP. On every later reconcile, the Terraform provider compares the configured reference with the live IP. The outcome depends on how the network is written:

* If the network string matches the live value (the real GCP state shape, where the network is a full URL), `forceNewIfNetworkIPNotUpdatable` sets ForceNew on `network_ip`. KCC reports this as the immutable-field error above.
* If the raw network strings differ (the mock returns `networkInterfaces[].network` as a relative path), there is no ForceNew, and the `network_ip` diff persists on every reconcile.

**Image family.** The image diff is already suppressed by `DiskImageDiffSuppress` (`third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute/resource_compute_disk.go`), which the instance boot disk also uses. The mock can now exercise this case, and the regression tests below guard it. No code change is needed for it.

## Acceptance criteria

| AC | Statement | PR coverage |
| --- | --- | --- |
| AC1 | `networkIpRef.external` as an address selfLink or `projects/P/regions/R/addresses/A` path reconciles with no immutable-field error and no perpetual diff. | Mock stores the resolved IP: [#16171](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16171). Tests: [#16226](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16226). Fix: [#16232](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16232). |
| AC2 | `sourceImageRef` as an image family (`cos-cloud/cos-81-lts`) causes no immutable `bootDisk.0.InitializeParams.0.Image` error. | Already fixed on master by `DiskImageDiffSuppress`. Mock support: [#16171](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16171). Tests: [#16226](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16226). |
| AC3 | No VM restart loop with a hard-coded `networkIp`. | Not reproduced; out of scope. Partial guard: the `computeinstance-293-networkip-ip` fixture and the no-diff re-reconcile check ([#16226](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16226)). |
| AC4 | GKE add-on version lag. | Out of scope. The GKE master upgrade resolved it in the issue thread. |
| AC5 | Documentation of the GKE to Config Connector version mapping. | Out of scope; no in-repo home. External page: [install, upgrade and uninstall](https://docs.cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall) |

## Tests

### Unit tests

`TestComputeInstanceIssue293` in [`pkg/controller/tf/computeinstance_issue293_test.go`](../../../pkg/controller/tf/computeinstance_issue293_test.go) runs the KRM-to-Terraform conversion and the Terraform diff that the controller runs, against live state shaped like real GCP. Each case has two subtests: `unchanged spec` (expects no immutable change and no diff) and `label update` (expects only `labels.env: a -> b`). The control rows instead expect the change to be an immutable change on a specific field.

| Test | What it asserts | Before fix | After fix ([#16232](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16232)) |
| --- | --- | --- | --- |
| `networkIpRef by name` | ComputeAddress name reference, live IP. No immutable change, no diff, label-only update. | PASS | Expected PASS |
| `networkIpRef.external IP` | Literal IP in `networkIpRef.external`. Same assertions. | PASS | Expected PASS |
| `networkIpRef.external address selfLink` | Address selfLink that resolves to the live IP. Same assertions. | FAIL: immutable `networkInterface.0.NetworkIp` error | Expected PASS |
| `networkIpRef.external address path` | Address path that resolves to the live IP. Same assertions. | FAIL: immutable `networkInterface.0.NetworkIp` error | Expected PASS |
| `networkIpRef.external address selfLink, networkRef.external path` | Address selfLink with the network written as a relative path (the mock shape). Same assertions. | FAIL: perpetual `network_interface.0.network_ip` diff | Expected PASS |
| `sourceImageRef COS family shorthand` | `cos-cloud/cos-81-lts` against the live COS image. Same assertions. | PASS | Expected PASS |
| `sourceImageRef COS family path` | `projects/cos-cloud/global/images/family/cos-81-lts` against the live COS image. Same assertions. | PASS | Expected PASS |
| `control: sourceImageRef different COS family` | Spec changes to `cos-cloud/cos-85-lts`. Must require replacement, with `bootDisk.0.InitializeParams.0.Image` as the immutable field. | PASS | Expected PASS |

Go's `-run` pattern replaces spaces in subtest names with underscores, for example `TestComputeInstanceIssue293/networkIpRef.external_address_path`.

### E2E fixtures

The fixtures are in [`pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computeinstance/`](../../../pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computeinstance/). Their goldens describe the fixed behaviour.

| Fixture | What it asserts | Before fix | After fix ([#16232](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16232)) |
| --- | --- | --- | --- |
| `computeinstance-293-networkip-ip` | Create with a literal `networkIpRef.external`, then a label update. Re-reconciliation makes no changes. | PASS | Expected PASS |
| `computeinstance-293-cos-family` | Create with `cos-cloud/cos-121-lts`, then the `projects/cos-cloud/global/images/family/cos-121-lts` path. No immutable error. | PASS | Expected PASS |
| `computeinstance-293-networkip-addresspath` | `networkIpRef.external` as an address path, then a label update. | FAIL: terminal `UpdateFailed` immutable-field error | Expected PASS |
| `computeinstance-293-networkip-addressselflink` | `networkIpRef.external` as an address selfLink, with the no-diff re-reconcile check. | FAIL: `re-reconciliation found changes to apply` on `network_interface.0.network_ip` | Expected PASS |

The two failing fixtures show different symptoms for the same underlying cause, because the mock stores the network differently from real GCP. See [the e2e test pattern notes](../../patterns/e2e_test.md#re-reconciliation-checks) for the checks that produce these failures.

## How to run

```bash
go test ./pkg/controller/tf/ -run TestComputeInstanceIssue293 -count=1
```

All four fixtures against the mock. The mock e2e commands need the platform setup described below.

```bash
hack/compare-mock 'fixtures/computeinstance-293'
```

A single fixture with envtest (see [docs/patterns/e2e_test.md](../../patterns/e2e_test.md#running-a-single-fixture-with-go-test-and-envtest)):

```bash
source dev/tasks/setup-envtest
E2E_KUBE_TARGET=envtest RUN_E2E=1 E2E_GCP_TARGET=mock go test -test.count=1 -timeout 3600s -v ./tests/e2e -run 'TestAllInSeries/fixtures/computeinstance-293-networkip-addressselflink$'
```

The whole compute fixture suite, as CI runs it:

```bash
dev/ci/presubmits/tests-e2e-fixtures-suite compute
```

Harness unit tests for the re-reconciliation check:

```bash
go test ./tests/e2e -run 'TestReconcileEndNotEvaluatedReason|TestDiffRecorderWaitForReconcileAfter'
```

CI's unit-tests job excludes `./tests/e2e` (see `dev/ci/presubmits/unit-tests`), so run these locally when you change the harness.

**Platform setup for the mock.** The mock listens on `127.0.0.2` (`mockgcp/mock_http_roundtrip.go`). Linux routes all of `127.0.0.0/8` to the loopback interface. macOS assigns only `127.0.0.1` to `lo0` by default, so the mock may fail to listen on macOS. Either add the address with `sudo ifconfig lo0 alias 127.0.0.2 up`, or run the mock e2e commands in a Linux container.

## Out of scope

* AC3: the VM restart loop with a hard-coded IP was not reproduced. Only the partial guard above covers it.
* AC4: the GKE add-on version lag was resolved by the GKE master upgrade.
* AC5: the GKE to Config Connector version mapping has no in-repo home. See the external page in the AC table.
* Unrelated harness bug found along the way: `CompareGoldenFile` does not catch a stale `_http_mock.diff`. It is tracked in [#16225](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/16225).

## Status

Fix pending in [#16232](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/16232).
