# HealthcareConsentStore Journal

### 2026-10-23 HealthcareConsentStore Direct Controller Implementation
- **Context**: Brownfield migration of `HealthcareConsentStore` under `healthcare.cnrm.cloud.google.com/v1alpha1` to the Direct controller model.
- **Client Library**: Healthcare API uses the Google Cloud discovery client `google.golang.org/api/healthcare/v1`. The client is instantiated via `m.config.RESTClientOptions()` with `api.NewService(ctx, opts...)`.
- **Operations & Mutability**:
  - `Create`: synchronous call `Projects.Locations.Datasets.ConsentStores.Create(parent, desired).ConsentStoreId(consentStoreId)`.
  - `Update`: mutable fields (`defaultConsentTtl`, `enableConsentCreateOnUpdate`, `labels`) updated via `Projects.Locations.Datasets.ConsentStores.Patch(name, patch).UpdateMask(updateMask)`. Top-level diffs are computed using `diffs.GoogleAPI.Diff`.
  - `Delete`: `Projects.Locations.Datasets.ConsentStores.Delete(name)`. Handled with `direct.IsNotFound(err)` for idempotency.
  - `Find`: `Projects.Locations.Datasets.ConsentStores.Get(name)`. Handled with `direct.IsNotFound(err)`.
- **Labels**:
  - GCP labels are populated from KRM `metadata.labels` using canonical `label.GCPLabels(u)` in `buildDesired`.
- **Dual-Controller Configuration & Verification**:
  - Registered `k8s.ReconcilerTypeDirect` in `SupportedControllers` in `pkg/controller/resourceconfig/static_config.go`.
  - Verified roundtrip fuzzing with `TestFocusedMappers`.
  - Verified E2E test `TestAllInSeries/fixtures/healthcareconsentstorebasic` passes against mock GCP with zero re-reconciliation diffs.
