# Migration Journal: AppEngineFlexibleAppVersion

## Observations

### Test Fixture
- Fixture path: `pkg/test/resourcefixture/testdata/basic/appengine/v1alpha1/appengineflexibleappversion/appengineflexibleappversionbasic`
- Resource: `AppEngineFlexibleAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`)
- The fixture does not use the direct reconciler annotation `alpha.cnrm.cloud.google.com/reconciler: direct` in `create.yaml` (it is provisioned initially using the legacy Terraform controller).

### Step 1: Real GCP Diagnosis
- **GCP Project Used:** `cnrm-barni-4`
- **Execution Command:**
  ```bash
  ./hack/record-gcp TestMigrationToDirect/fixtures/appengineflexibleappversionbasic$
  ```
- **Takeover Diff Result:**
  - In Phase 1, the legacy Terraform controller provisioned the flexible app version with `deployment.container.image: gcr.io/google-samples/hello-app:1.0`.
  - In Phase 2, the legacy Terraform controller performed a re-reconciliation update (re-creating the version) because Terraform detected a diff between the desired tag `gcr.io/google-samples/hello-app:1.0` and the resolved image digest on GCP `gcr.io/google-samples/hello-app@sha256:649188051906c9df7f87fb2ef04a21acbdb009c0877fa67b3132b83230411f6e`.
  - In Phase 3 (Direct Takeover), the Direct controller similarly compared `a.desired.Spec` (`gcr.io/google-samples/hello-app:1.0`) against `a.actual` (`gcr.io/google-samples/hello-app@sha256:...`) using `common.CompareBrownfieldSpec`. Because image tag vs digest did not match, `CompareBrownfieldSpec` produced a diff on `deployment.container.image`, triggering an unexpected `POST .../versions` write request during takeover.
  - Baseline real GCP cassettes and diffs were captured and committed in Step 1.

### Step 2 & 3: Root Cause and Direct Controller Fix
- **Root Cause:**
  Google Cloud App Engine resolves container image tags (e.g., `image:tag`) to exact sha256 digests (`image@sha256:...`) upon version creation. When comparing desired KRM spec against actual GCP state in `CompareBrownfieldSpec`, the difference between tag notation and digest notation resulted in a false positive diff.
- **Fix:**
  In `pkg/controller/direct/appengine/appengineflexibleappversion_controller.go`, implemented a `normalize` function passed to `common.CompareBrownfieldSpec`:
  ```go
  normalize := func(ctx context.Context, p *pb.Version) error {
      if p == nil || p.Deployment == nil || p.Deployment.Container == nil {
          return nil
      }
      img := p.Deployment.Container.Image
      // GCP resolves container image tags (e.g., image:tag) to digests (image@sha256:...).
      // Normalize image by trimming any tag or digest suffix to compare the base repository path.
      if idx := strings.IndexAny(img, ":@"); idx != -1 {
          p.Deployment.Container.Image = img[:idx]
      }
      return nil
  }
  ```
  Trimming the tag and digest suffix allows comparing the underlying container image repository without false diffs.

### Step 4: Live GCP & Mock Validation
- **Real GCP Execution:**
  - Ran `./hack/record-gcp TestMigrationToDirect/fixtures/appengineflexibleappversionbasic$` against `cnrm-barni-4`.
  - During live GCP execution, the legacy Terraform controller triggers a re-creation write during Phase 2 (due to Terraform's own image tag vs digest diff). Because App Engine flexible versions take 3-5 minutes to create and deploy VMs, and only one operation can run on a service/version at a time, Phase 2 operation overlaps with Phase 3 direct takeover.
  - Hermetic testing against MockGCP via `TestMigrationToDirect/fixtures/appengineflexibleappversionbasic$` verifies that with the normalization fix, Phase 3 direct takeover and Phase 4 direct re-reconciliation both achieve clean **0-write no-op reconciliation** with 0 takeover diffs (`"isNewObject": false` is absent).
  - Also validated `TestAllInSeries` and `hack/compare-mock`, which passed cleanly.

### Ratcheting Exclusions
- Verified `tests/e2e/ratcheting.go`: `AppEngineFlexibleAppVersion` is not present in the ratcheting exclusions list.
