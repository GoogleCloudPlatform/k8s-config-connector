# ContainerNodePool Migration Journal

## Observation & Design Decisions

- **API Endpoint and Location Normalization:** In the GKE v1 API (`cloud.google.com/go/container/apiv1`), the canonical resource path for ClusterManager API calls is `projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}`, where `{location}` can represent either a region (e.g., `us-central1`) or a zone (e.g., `us-central1-a`). Using the `/locations/` URL format allows the client to interact seamlessly with both regional and zonal clusters without needing legacy `/zones/` REST mappings.
- **Diffing and Server Defaults with `CompareBrownfieldSpec`:** GKE returns extensive server-side defaults (such as default upgrade settings, default management configs, and empty slices vs nil). To achieve zero-write steady-state re-reconciliation, we utilized `common.CompareBrownfieldSpec` to merge server-returned defaults into the desired state before computing diffs and top-level field masks.
- **Handling Multi-endpoint Updates:** Updating a GKE Node Pool requires different GCP endpoints depending on the modified fields:
  - Autoscaling modifications are routed via `ClusterManager.UpdateCluster` using `DesiredNodePoolAutoscaling` and `DesiredNodePoolId`.
  - Node management settings (auto-upgrade / auto-repair) are routed via `ClusterManager.SetNodePoolManagement`.
  - Node pool configuration updates (taints, network tags, labels, linux node config, kubelet config, resource manager tags, containerd config, workload metadata config, logging, resource labels, windows config, upgrade settings, locations, image type, queued provisioning) are routed via `ClusterManager.UpdateNodePool`.
- **Status Fields Alignment:** `instanceGroupUrls` is mapped directly from `NodePool.InstanceGroupUrls`. `managedInstanceGroupUrls` is derived by replacing `/instanceGroupManagers/` with `/instanceGroups/` and `/compute/v1/` with `/compute/v1beta1/` to maintain 100% fidelity with legacy status schemas.
