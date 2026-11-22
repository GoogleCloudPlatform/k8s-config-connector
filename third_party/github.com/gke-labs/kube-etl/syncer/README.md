# KRMSyncer

The KRMSyncer is a Kubernetes-native tool designed for multi-cluster synchronization of KRM resources.

## Features

- **Push & Pull Models:** Support both pushing from local to remote and pulling from remote to local clusters.
- **Dynamic Watching:** Dynamically registers watches for resources specified in the configuration.
- **Resource Syncing:** Syncs standard resources (e.g., ConfigMaps, Secrets).
- **Status Syncing:** Optionally syncs the status subresource.
- **Suspension:** Supports pausing sync operations via a `suspend` field.
- **Namespace Mapping:** Supports syncing to a specific destination namespace.

## Overview

The operator manages the `KRMSyncer` Custom Resource to coordinate resource replication:

1.  **Reconciling (Active cluster)**:
    *   Watches specific Kubernetes resources defined in rules.
    *   Continuously syncs their state directly to the destination.
    *   The remote GKE cluster is referenced directly in `spec.remote.gkeCluster`; the controller authenticates to it with Workload Identity.
    *   Default mode is `pull`.

2.  **Suspended (Passive cluster)**:
    *   Acts as the receiver.
    *   The controller in this mode remains idle regarding synchronization, waiting for updates from the other cluster.

## Configuration (KRMSyncer CRD)

The `KRMSyncer` resource allows you to define what to sync and where to sync it.

```yaml
apiVersion: syncer.gkelabs.io/v1alpha1
kind: KRMSyncer
metadata:
  name: resource-sync
spec:
  suspend: false
  mode: pull # Defaults to 'pull'.
  rules:
    - group: "pubsub.cnrm.cloud.google.com"
      version: "v1beta1"
      kind: "PubsubTopic"
      namespaces: ["default"] # Only sync PubsubTopic resources in the 'default' namespace
  remote:
    gkeCluster:
      project: my-project
      location: us-central1
      name: source-cluster
      endpoint: Default # Optional. Default | DNS | PrivateIP
```

`spec.remote.gkeCluster.endpoint` selects how the controller reaches the remote cluster:
- `Default` (default): the cluster's default endpoint, the same one `gcloud container clusters get-credentials` uses.
- `DNS`: the cluster's DNS-based endpoint. It must be enabled for external traffic (`gcloud container clusters update --enable-dns-access`); needs no VPC connectivity, and is unaffected by credential rotation.
- `PrivateIP`: the cluster's private endpoint. The controller must be able to reach the cluster's VPC.

## Retries

Failed syncs are retried per object with exponential backoff (controller-runtime's default): 5ms after the first failure, doubling after each failure, up to 1000s (~16.7 minutes). The backoff resets when the sync succeeds.

| Consecutive failures | Delay before next retry |
|---|---|
| 1 | 5 ms |
| 10 | ~2.6 s |
| 14 | ~41 s |
| 16 | ~2.7 min |
| 18 | ~10.9 min |
| 19+ | ~16.7 min (cap) |

## Run Integration test

```bash
# Build the manager binary
cd syncer
make test-integration
```

To build and push the controller image without `deploy.sh`, set `DEST_LOCATION` to the destination cluster's region or zone. The image goes to `<dest-region>-docker.pkg.dev/<project>/krmsyncer/controller:latest`, the same default `deploy.sh` uses, so the `krmsyncer` repository must already exist. Set `IMAGE_LOCATION` to use another Artifact Registry region or multi-region, or `IMG` to set the image directly:
```bash
make docker-build docker-push DEST_LOCATION=us-central1-a
make docker-build docker-push IMAGE_LOCATION=us
```

## Getting Started

### Prerequisites

Before running [`deploy.sh`](deploy.sh), make sure you have the following.

**Local tools**

- `kubectl`, `gcloud`, and `docker`.
- [`gke-gcloud-auth-plugin`](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl#install_plugin), so your local `kubectl` context can authenticate to the destination cluster.
- Push access to the image registry (e.g. `gcloud auth configure-docker <image-location>-docker.pkg.dev`). With the default image, the script creates the `krmsyncer` Artifact Registry repository if needed, which requires the Artifact Registry API to be enabled in the project.

**Clusters**

- The destination cluster must have [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) enabled.
- You need permission to create service accounts and grant IAM roles in the project.
- You need permission to grant IAM roles in the source cluster's project, and to describe the source cluster.
- You need a kubeconfig context for the destination cluster, with permission to deploy on it. Create it with the following command, or pass an existing context with `--dest-context`:
  ```bash
  gcloud container clusters get-credentials <DEST_CLUSTER_NAME> --location <DEST_CLUSTER_LOCATION> --project <GCP_PROJECT_ID>
  ```

- The source cluster must expose the endpoint you select with `--source-endpoint`, and the destination cluster must be able to reach it (see [Network access to the source cluster](#network-access-to-the-source-cluster)).
- To use the sample `KRMSyncer` CR, destination cluster needs the Config Connector CRDs installed for all the resources in the source cluster.

> [!WARNING]
> **The destination cluster must be passive.** The sample `KRMSyncer` CR syncs the `status` (and, with `--sync-mode full`, the `spec`) of every Config Connector resource from the source cluster into the destination cluster. If Config Connector is running in the destination cluster, and not just installed as CRDs, both clusters reconcile the same GCP resources and fight over them.

### Network access to the source cluster

The controller runs in the destination cluster and calls the source cluster's control plane directly, so the destination cluster's Pods need a network path to the endpoint selected by `--source-endpoint`. Configure this before deploying; otherwise the controller only logs connection timeouts.

| `--source-endpoint` | Network requirements |
|---|---|
| `DNS` | None between the clusters: traffic goes over Google's network and access is controlled by IAM (`container.clusters.connect`, included in `roles/container.viewer`). The DNS endpoint must accept external traffic: `gcloud container clusters update <SOURCE_CLUSTER_NAME> --location <SOURCE_CLUSTER_LOCATION> --enable-dns-access`. |
| `PrivateIP` | The destination cluster must be able to route to the source's private endpoint: same VPC or Shared VPC, or a directly connected VPC (VPC peering, Cloud VPN, Interconnect; peering is not transitive). If the destination is in a different region than the source, the source must have [control plane global access](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters#cp-global-access) enabled (`--enable-master-global-access`); without it, cross-region connections time out even on the same VPC. Firewall rules must allow the destination's node and Pod ranges to reach it on TCP 443. If the source enforces authorized networks on the private endpoint (`privateEndpointEnforcementEnabled`), add those ranges to its authorized networks. |
| `Default` | The source's IP endpoint must accept the destination's traffic. If it's the public endpoint, add the destination's egress IPs (e.g. its Cloud NAT addresses) to the source's [authorized networks](https://cloud.google.com/kubernetes-engine/docs/how-to/authorized-networks). If the cluster only exposes its private endpoint, the `PrivateIP` requirements apply. |

> [!NOTE]
> **Config Controller as the source.** Config Controller clusters (`krmapihost-*`) are private clusters that, by default, also have a public endpoint, control plane global access off, and DNS external traffic off. In practice:
> - Use `--source-endpoint Default` (the public endpoint). It works from any region and VPC.
> - Use `--source-endpoint PrivateIP` only if the destination cluster is in the same region and on the same VPC as the source cluster. Cross-region connections to the private endpoint time out.
> - `--source-endpoint DNS` requires enabling DNS access on the managed cluster, which may not be permitted.
>
> If the source was created with `--use-private-endpoint`, it has no public endpoint; see [Private-only Config Controller](#private-only-config-controller).

#### Private-only Config Controller

A Config Controller cluster created with `--use-private-endpoint` is reachable only at its private endpoint, an internal IP in the source's VPC subnet. Its network settings are fixed when it is created, so the destination cluster has to be created to fit them. Use `--source-endpoint PrivateIP`.

First, read the source's settings:
```bash
gcloud container clusters describe <SOURCE_CLUSTER_NAME> --location <SOURCE_CLUSTER_LOCATION> --project <SOURCE_PROJECT_ID> \
  --format="yaml(location, network, privateClusterConfig.privateEndpoint, masterAuthorizedNetworksConfig.cidrBlocks)"
```

Then create the destination cluster so that:

1. **Region:** it is in the source's `location`. The source has control plane global access off, so connections from other regions time out, even over the same VPC, VPC peering, Cloud VPN or Interconnect.
2. **Network:** it is on the source's `network`. For a destination in another project, use [Shared VPC](https://cloud.google.com/vpc/docs/shared-vpc) and pass the host project's network and subnet.
3. **IP ranges:** its node subnet range and Pod range fall inside the source's `masterAuthorizedNetworksConfig.cidrBlocks`. Authorized networks are enforced on the private endpoint, and the control plane sees the Pods' own IPs (GKE does not SNAT Pod traffic to internal ranges). With the default `0.0.0.0/0`, any range works; otherwise pick unused ranges inside the listed blocks. If no free range fits, the destination can't reach the source.
4. **Workload Identity:** it is enabled (`--workload-pool`), which the controller needs to authenticate.

For example:
```bash
gcloud container clusters create <DEST_CLUSTER_NAME> --project <DEST_PROJECT_ID> \
  --location <SOURCE_CLUSTER_LOCATION> \
  --network <SOURCE_NETWORK> --subnetwork <SUBNET_IN_SOURCE_REGION> \
  --enable-ip-alias --cluster-ipv4-cidr <POD_RANGE_INSIDE_AUTHORIZED_BLOCKS> \
  --workload-pool <DEST_PROJECT_ID>.svc.id.goog
```

Also check:
- **Firewall:** egress within a VPC is allowed by default. If your organization adds egress deny rules, allow TCP 443 from the destination's node and Pod ranges to the private endpoint.
- **Google APIs:** the controller still calls `container.googleapis.com` and the Workload Identity token endpoints. If the destination's nodes have no internet egress (private nodes without Cloud NAT), enable [Private Google Access](https://cloud.google.com/vpc/docs/private-google-access) on their subnet:
  ```bash
  gcloud compute networks subnets update <DEST_SUBNET> --region <SOURCE_CLUSTER_LOCATION> --enable-private-ip-google-access
  ```

If the destination must run in a different region, it can't reach a private-only source directly.

#### Inspecting endpoints and testing reachability

Inspect the source cluster's endpoints and access settings with:
```bash
gcloud container clusters describe <SOURCE_CLUSTER_NAME> --location <SOURCE_CLUSTER_LOCATION> --project <SOURCE_PROJECT_ID> \
  --format="yaml(network, endpoint, privateClusterConfig.privateEndpoint, privateClusterConfig.masterGlobalAccessConfig, controlPlaneEndpointsConfig, masterAuthorizedNetworksConfig)"
```

To confirm the destination cluster can reach the endpoint, open a TCP connection to port 443 from inside it (replace the address with the selected endpoint). `open` means there is a network path; a timeout means traffic is dropped (routing, firewall or authorized networks):
```bash
kubectl --context=<DEST_CONTEXT> run nettest --rm -i --restart=Never --image=busybox -- \
  nc -zv -w 10 <SOURCE_ENDPOINT> 443
```

### 1. Deploy KRMSyncer to the Destination Cluster

Use the [`deploy.sh`](deploy.sh) script.

```bash
cd syncer

./deploy.sh \
  --source-cluster <SOURCE_CLUSTER_NAME> \
  --source-location <SOURCE_CLUSTER_LOCATION> \
  --dest-cluster <DEST_CLUSTER_NAME> \
  --dest-location <DEST_CLUSTER_LOCATION> \
  --project <GCP_PROJECT_ID> \
  [--dest-context <DEST_CONTEXT>] \
  [--source-project <SOURCE_PROJECT_ID>] \
  [--source-endpoint Default|DNS|PrivateIP] \
  [-n <NAMESPACE>] \
  [--sync-mode status|full] \
  [-i <IMAGE> | --image-location <LOCATION>] \
  [--skip-build]
```

`--source-project` sets the source cluster's project (default: `--project`).

`--dest-context` sets the kubectl context of the destination cluster. It defaults to `gke_<project>_<dest-location>_<dest-cluster>`, the name `gcloud container clusters get-credentials` creates; set it if you renamed the context. The script warns if the context's server is not one of the destination cluster's endpoints.

`--source-endpoint` selects the source cluster's control plane endpoint: `Default`, `DNS` or `PrivateIP` (default: `Default`).

`-n` sets the namespace for the `KRMSyncer` CR (default: `krmsyncer-system`, the same namespace as the controller).

`--sync-mode` selects what the `KRMSyncer` CR replicates from Config Connector resources (default: `status`):
- `status`: only `status`, plus `spec.resourceID` for the GVKs listed in [`config/templates/service-generated-resource-id-gvks.yaml`](config/templates/service-generated-resource-id-gvks.yaml) (resources whose ID is generated by the GCP service, which KCC writes back into the spec).
- `full`: `spec` and `status` of all KCC resources.

`-i` sets the controller image to deploy. Defaults to `<image-location>-docker.pkg.dev/<project>/krmsyncer/controller:latest` (Container Registry, `gcr.io`, is shut down).

`--image-location` sets the Artifact Registry location of the default image: a region (e.g. `us-central1`) or a multi-region (`us`, `europe`, `asia`). Defaults to the region of `--dest-location`. Ignored if `-i` is set.

`--skip-build` deploys the image as is, without building and pushing it.

This command:
1. Runs read-only preflight checks, so a misconfiguration fails before anything is created: Workload Identity is enabled on the destination cluster, the GSA (if it already exists) was created for this destination cluster, the source cluster exists and exposes the selected endpoint, the destination cluster is reachable and you can create CRDs and cluster RBAC on it, the manifests render, and (unless `--skip-build` is set) Docker is running.
2. Configures Workload Identity. It creates a Google Service Account (GSA) for the destination cluster, `krmsyncer-<dest-cluster>-<hash>@<project>.iam.gserviceaccount.com` (up to 12 characters of the cluster name plus a hash of its location and name, to stay within the 30-character limit on GSA names). It grants the GSA `roles/container.viewer` on the source cluster's project, and lets the controller's `krmsyncer-system/krmsyncer-controller-manager` Kubernetes ServiceAccount impersonate it. `roles/container.viewer` lets the GSA look up the source cluster, connect to it (including over the DNS endpoint) and read its Kubernetes objects.
3. Builds and pushes the controller image (unless `--skip-build` is set). With the default image, it first creates the `krmsyncer` Docker repository in Artifact Registry in `--image-location` if it doesn't exist. It then grants the destination cluster's node service accounts (those of its node pools and the node auto-provisioning/Autopilot default) `roles/artifactregistry.reader` on that repository, so nodes can pull the image even when they run as a minimal custom service account; this also runs with `--skip-build`.
4. Deploys the RBAC, KRMSyncer CRD and controller into the `krmsyncer-system` namespace of the destination cluster, and annotates the controller's ServiceAccount with the GSA.
5. Applies the `kcc-resource-syncer` `KRMSyncer` CR from [`config/templates/krmsyncer.yaml`](config/templates/krmsyncer.yaml). It references the source cluster in `spec.remote.gkeCluster` and syncs all Config Connector resources from it: with `--sync-mode status` (default), only `status`, plus `spec.resourceID` for the GVKs listed in [`config/templates/service-generated-resource-id-gvks.yaml`](config/templates/service-generated-resource-id-gvks.yaml); with `--sync-mode full`, `spec` and `status`. The destination cluster needs the matching KCC CRDs pre-installed.

The script only adds IAM bindings. If you re-run it with a different `--source-project`, the GSA keeps its access to the previous source project until you revoke it with [`cleanup.sh`](cleanup.sh) (see [Cleaning up IAM bindings](#cleaning-up-iam-bindings)).

> [!WARNING]
> All `KRMSyncer` objects share the controller's Google identity. Anyone who can create a `KRMSyncer` can sync with any cluster that identity can access (with `roles/container.viewer`, every cluster in the source project and in any project it was previously granted), so restrict who can create `KRMSyncer` objects.

View the controller logs with:
```bash
kubectl --context=<DEST_CONTEXT> -n krmsyncer-system logs deploy/krmsyncer-controller-manager -f
```

### Cleaning up IAM bindings

`deploy.sh` grants the destination cluster's GSA `roles/container.viewer` on the source project, but never removes grants. After `deploy.sh` has switched a destination cluster to a different `--source-project`, revoke the old grant with [`cleanup.sh`](cleanup.sh):

```bash
cd syncer

./cleanup.sh \
  --dest-cluster <DEST_CLUSTER_NAME> \
  --dest-location <DEST_CLUSTER_LOCATION> \
  --project <GCP_PROJECT_ID> \
  [--dest-context <DEST_CONTEXT>] \
  --source-project <OLD_SOURCE_PROJECT_ID> \
  [--dry-run]
```

`--source-project` names a previous source project to revoke (required). Repeat it for several projects.

`--dest-context` sets the kubectl context of the destination cluster.

`--dry-run` prints the grants that would be revoked without changing anything.

The script only revokes what is safe to revoke:

- It removes only the GSA's own unconditional `roles/container.viewer` binding. Other members of the role keep their access.
- It never revokes a project that a `KRMSyncer` CR in the destination cluster still references in `spec.remote.gkeCluster.project`, in any namespace.
- It refuses a GSA that `deploy.sh` did not create for this destination cluster (checked through the GSA's description), since other clusters may depend on its grants.

To remove KRMSyncer from a destination cluster entirely, delete the `KRMSyncer` CRs first. Then run `cleanup.sh` with the remaining source project and delete the GSA:
```bash
gcloud iam service-accounts delete <GSA_EMAIL> --project <GCP_PROJECT_ID>
```
Revoke the project grants **before** deleting the GSA. A deleted GSA has no access, but its bindings stay in the policies as `deleted:serviceAccount:...` members until IAM purges them.

### 2. Verify the Results
1. Create a test resource in the Source cluster:
   ```bash
   kubectl --context=<source-cluster-context> apply -f - <<EOF
   apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
   kind: PubSubTopic
   metadata:
     name: test-topic
     namespace: default
   EOF
   ```

1. Check the Destination cluster:
   Switch your kubectl context to the Destination cluster and verify the PubsubTopic has appeared:
   ```bash
   kubectl --context=<dest-cluster-context> get pubsubtopic test-topic
   ```
1.  Expected Result:
- The `test-topic` PubsubTopic created in the Source cluster should automatically appear in the Destination cluster within seconds.
- If you update the PubsubTopic in the Source cluster, the changes should reflect in the Destination cluster.
- If you delete it from the Source cluster, it should be removed from the Destination cluster.
