# Issue 4: Dynamic Service Usage Discovery & Regional CRD Gating for Sovereign Partitions

**Status:** Open  
**Epic:** KCC Support in Sovereign Cloud  
**Components:** `operator/pkg/discovery`, `operator/config/crd`, `pkg/crd`

---

## 1. Problem Description

Currently, the KCC operator unconditionally installs all 295 Custom Resource Definitions (CRDs) during bootstrap. However, in Sovereign Cloud environments and specialized edge regions, only a small fraction of Google Cloud services are enabled or deployed.

Empirical scanning via `serviceusage.googleapis.com` in Google Sovereign Cloud (`u-region-1`) revealed that only **33 services** exist in the universe, including:
- Core compute and networking (`compute`, `container`, `dns`, `networksecurity`, `networkconnectivity`)
- Storage & Data (`storage`, `bigquery`, `dataproc`, `sqladmin`, `pubsub`)
- Identity & Governance (`iam`, `accesscontextmanager`, `orgpolicy`, `serviceusage`, `cloudkms`, `logging`, `monitoring`)

Over 260 services (such as `spanner`, `alloydb`, `aiplatform`, `looker`, `clouddeploy`, `apigee`, `vertexai`) are completely absent from the sovereign boundary.

Installing 295 CRDs in this environment has severe drawbacks:
1. **Memory & etcd Bloat:** Kubernetes API server and etcd store megabytes of unused OpenAPI schemas, increasing API server memory consumption by over 400MB.
2. **Controller Resource Exhaustion:** KCC spawns reconciler workers for all 295 resource types. Unsupported controllers enter perpetual polling/reconcile failures.
3. **Webhook Timeouts:** Admission webhooks process resources across hundreds of inactive schemas.

## 2. Solution: Dynamic CRD Discovery & Gating Engine

Instead of a monolithic, static installation, KCC Operator must support an automated discovery mode:

1. **Service Usage Probing:**  
   During bootstrap (or on a periodic configurable sync loop), the operator queries the target universe's Service Usage API (`serviceusage.<UNIVERSE_DOMAIN>/v1/projects/<PROJECT>/services?filter=state:ENABLED`) or reads a declared service allowlist.
2. **CRD-to-Service Mapping Catalog:**  
   Maintain a mapping between GCP API service names and KCC CRD Group/Kinds:
   ```
   compute.googleapis.com       -> compute.*.cnrm.cloud.google.com
   storage.googleapis.com       -> storage.*.cnrm.cloud.google.com
   pubsub.googleapis.com        -> pubsub.*.cnrm.cloud.google.com
   sqladmin.googleapis.com      -> sql.*.cnrm.cloud.google.com
   ```
3. **Selective CRD Pruning & Registration:**  
   The operator only applies CRDs matching the available services in the sovereign universe. If a customer subsequently enables an API (e.g., `dataproc`), KCC discovers the service and dynamically registers the corresponding CRDs and starts their reconcilers.
4. **Declarative Mode via ConfigConnector Spec:**  
   Allow platform engineers to declare a profile in the `ConfigConnector` CR:
   ```yaml
   apiVersion: core.cnrm.cloud.google.com/v1beta1
   kind: ConfigConnector
   metadata:
     name: configconnector.core.cnrm.cloud.google.com
   spec:
     mode: cluster
     googleServiceAccount: "kcc-system@sample-project.partition.iam.gserviceaccount.com"
     crdManagement:
       discovery: Auto # Options: All, Auto, Manual
       universeDomain: "custom.universe.goog"
   ```

## 3. Benefits

- Reduces installed CRDs from 295 down to the exact 33 supported in the partition (~88% reduction in KCC CRD footprint).
- Saves ~400MB API server / etcd memory.
- Prevents spurious reconciliation loops for nonexistent regional endpoints.

## 4. Acceptance Criteria

- [ ] Operator correctly maps enabled GCP services to KCC CRD subsets.
- [ ] In Sovereign Germany (`custom.universe.goog`), only the ~33 supported service CRDs are registered when `crdManagement.discovery: Auto` is enabled.
- [ ] Enabling a new service in GCP automatically triggers CRD registration and controller activation.
