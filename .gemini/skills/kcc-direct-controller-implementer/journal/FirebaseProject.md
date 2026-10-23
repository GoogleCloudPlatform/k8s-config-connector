# FirebaseProject Journal

### 2026-10-22 FirebaseProject Direct Controller Implementation
- **Context**: Brownfield migration of `FirebaseProject` under `firebase.cnrm.cloud.google.com/v1alpha1` to the Direct controller model.
- **Discovery Client Usage**: Firebase Management API uses `google.golang.org/api/firebase/v1beta1`. Client is instantiated via `m.config.RESTClientOptions()` with `api.NewService(ctx, opts...)`.
- **Operations & Immutability**:
  - `Create`: calls `Projects.AddFirebase(name, &api.AddFirebaseRequest{})` returning an LRO, which is awaited via `Operations.Get`. Then `Projects.Get` is called to fetch the full object.
  - `Update`: FirebaseProject in KRM Spec has no mutable fields (`projectRef`, `resourceID`). Reconciles status with `actual`.
  - `Delete`: Firebase Projects cannot be deleted from Google Cloud. Model is registered with `registry.CannotBeDeleted()`, and Delete returns `(true, nil)` to allow clean removal from Kubernetes.
  - `Find`: `Projects.Get(name)`. Handled with `direct.IsNotFound(err)` for 404s.
  - `Export`: Converts `api.FirebaseProject` to KRM format via `FirebaseProjectSpec_FromAPI`.
- **Dual-Controller Testing**:
  - Added `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `FirebaseProject` in `pkg/controller/resourceconfig/static_config.go`.
  - Verified with `hack/compare-mock pkg/test/resourcefixture/testdata/basic/firebase/v1alpha1/firebaseproject/firebaseprojectautogen`.
