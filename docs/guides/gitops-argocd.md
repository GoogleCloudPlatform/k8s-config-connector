# Managing Config Connector with GitOps (ArgoCD)

This guide provides a production-tested workflow for installing, managing, and upgrading Config Connector (KCC) using GitOps tools such as **ArgoCD** or **Flux**.

---

## Architecture and deployment options

Config Connector can be deployed and managed via GitOps using two primary patterns:

1. **Operator-based GitOps (Recommended)**: ArgoCD deploys the Config Connector Operator via pre-rendered release bundles. The Operator then manages CRD lifecycle and controller deployments automatically based on declarative `ConfigConnector` custom resources.
2. **Pre-rendered manifest GitOps**: ArgoCD directly manages the rendered release bundle manifests (`0-cnrm-system.yaml` and resource CRDs) rendered from Kustomize templates.

---

## Prerequisites and key considerations

Before setting up ArgoCD applications for Config Connector, ensure you meet the following prerequisites and review key operational constraints:

### Prerequisites
* A running Kubernetes cluster (e.g., GKE Standard or GKE Autopilot) with Workload Identity enabled.
* ArgoCD installed and running in your cluster.
* A Google Cloud IAM Service Account created with the required GCP roles for the resources you intend to manage.

### Key considerations

> [!NOTE]
> **Raw upstream repository vs. rendered manifests**  
> The raw source repository path (`operator/config/default`) contains release patch templates (`manager_image_patch_template.yaml`) that require build-time rendering. Attempting to point ArgoCD directly at raw upstream source paths without generating `manager_image_patch.yaml` results in Kustomize build errors.  
> **Solution**: Use official published release bundles from Google Cloud Storage, or render the Kustomize targets using `./dev/tasks/build-release-bundle` in your CI pipeline before committing manifests to your GitOps repository.

> [!IMPORTANT]
> **Mandatory Server-Side Apply for CRDs and AI workloads**  
> Client-side `kubectl apply` stores the serialized configuration in the `kubectl.kubernetes.io/last-applied-configuration` annotation. Because the Kubernetes API server enforces a 256KB (262,144 bytes) limit on `metadata.annotations`, applying large manifests client-side will fail with `Too long: may not be more than 262144 bytes` in two primary scenarios:
> 1. **Config Connector CRD bundles**: KCC includes over 250 CRDs with extensive OpenAPI validation schemas.
> 2. **AI hypercomputing workload schemas**: Modern AI/ML orchestration stacks (such as Kueue `workloads.kueue.x-k8s.io` and `clusterqueues.kueue.x-k8s.io`, or JobSet) deployed alongside KCC contain massive schemas that exceed this annotation limit.
> 
> **Solution**: You **must enable Server-Side Apply (SSA)** in your GitOps sync configurations (such as ArgoCD `ServerSideApply=true`). SSA transfers field management to the API server, bypassing client-side annotation size limits. Config Sync applies resources server-side by default.

> [!CAUTION]
> **CRD upgrades and safety rules**  
> * **In-place updates**: CRDs update **in-place** when applying new versions. etcd object creation timestamps and live GCP resources remain untouched.  
> * **Never delete CRDs**: Running `kubectl delete crd` cascade-deletes all Custom Resource (CR) instances across all namespaces. Because Config Connector resources have finalizers attached by default, deleting CR instances instructs Config Connector to **delete the underlying production GCP infrastructure** in Google Cloud.

---

## Step-by-step installation guide

### Step 1: Obtain or render the installation manifests

Obtain the required pre-rendered Config Connector installation manifests using either of the following methods:

#### **Method A: Official Google Cloud published release bundles (Recommended)**
If you are installing Config Connector without building from source, download the latest official release bundle tarball from Google Cloud Storage:

```bash
# Download the official published release bundle
curl -O https://storage.googleapis.com/config-connector-operator/latest/release-bundle.tar.gz

# Extract the release bundle files
tar -xvf release-bundle.tar.gz
```

This tarball contains pre-rendered, production-ready manifests:
* `0-cnrm-system.yaml`: Contains the Operator deployment, RBAC roles, ServiceAccounts, and `configconnector-operator-system` namespace.
* `crds.yaml`: Contains the complete set of Config Connector CustomResourceDefinitions.

#### **Method B: Building from GitHub repository source**
If you are tracking the open-source GitHub repository source code:

```bash
# Clone the upstream repository
git clone https://github.com/GoogleCloudPlatform/k8s-config-connector.git
cd k8s-config-connector

# Render standard cluster bundle manifests
./dev/tasks/build-release-bundle
```
This populates rendered manifests in `config/installbundle/releases/` ready for GitOps use.

#### **Recommended GitOps repository structure**
Organize your team's GitOps repository (for example, `https://github.com/<your-org>/kcc-gitops.git`) as follows:

```text
kcc-gitops/
├── operator/
│   └── 0-cnrm-system.yaml    # Operator deployment & RBAC
└── crds/
    └── crds.yaml              # Complete KCC CRD bundle
```

---

### Step 2: Configure ArgoCD Application for KCC Operator

Create an ArgoCD `Application` targeting your rendered Config Connector manifests directory in your GitOps repository. Ensure `ServerSideApply=true` is set under `syncOptions`:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: config-connector-operator
  namespace: argocd
spec:
  project: default
  source:
    repoURL: 'https://github.com/<your-org>/kcc-gitops.git'
    targetRevision: main
    path: operator                     # Path to 0-cnrm-system.yaml and crds.yaml
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: configconnector-operator-system
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
      - ServerSideApply=true           # CRITICAL: Bypasses 262KB annotation limit on CRDs
```

---

### Step 3: Declaratively configure Config Connector

Once ArgoCD deploys the Operator, declaratively configure Config Connector by committing a `ConfigConnector` custom resource to your GitOps repository.

#### **Cluster mode configuration**:
Use `mode: cluster` when managing GCP resources across the entire cluster using a single Google Service Account:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster                         # Cluster-wide reconciliation mode
  googleServiceAccount: "kcc-system@<YOUR_PROJECT_ID>.iam.gserviceaccount.com"
```

#### **Namespace mode configuration**:
Use `mode: namespaced` and deploy a `ConfigConnectorContext` resource in each managed namespace to bind dedicated GCP Service Accounts per namespace:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: config-control             # Target namespace to enable
spec:
  googleServiceAccount: "kcc-namespace-sa@<YOUR_PROJECT_ID>.iam.gserviceaccount.com"
```

---

## AI workloads and hypercomputing considerations (Kueue & JobSet)

When using Config Connector as the foundation for AI clusters (e.g., orchestrating Cloud TPU/GPU node pools with `ContainerNodePool`), users frequently deploy modern AI orchestration frameworks—such as [Kueue](https://kueue.sigs.k8s.io/) and [JobSet](https://jobset.sigs.k8s.io/)—directly onto the provisioned infrastructure using GitOps.

### The 256KB annotation limit
Standard client-side `kubectl apply` records the applied object state in the `kubectl.kubernetes.io/last-applied-configuration` annotation. Because the Kubernetes API server validates and limits `metadata.annotations` to 256KB (262,144 bytes), deploying massive CRDs and custom resources with rich validation schemas fails with:

```text
Error from server (Invalid): CustomResourceDefinition.apiextensions.k8s.io "...": 
CustomResourceDefinition.apiextensions.k8s.io "workloads.kueue.x-k8s.io" is invalid: 
metadata.annotations: Too long: may not be more than 262144 bytes
```

This error commonly affects:
* Kueue workload definitions: `workloads.kueue.x-k8s.io`
* Kueue cluster queue definitions: `clusterqueues.kueue.x-k8s.io`
* JobSet custom resource definitions: `jobsets.jobset.sigs.k8s.io`

### Enabling Server-Side Apply (SSA)
Server-Side Apply transfers field management and diff calculations to the Kubernetes API server, bypassing client-side annotation storage limits.

#### ArgoCD Application configuration
Ensure `ServerSideApply=true` is enabled in `syncOptions` for your AI workload Applications:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: kueue-system
  namespace: argocd
spec:
  project: default
  source:
    repoURL: 'https://github.com/<your-org>/ai-workloads.git'
    targetRevision: main
    path: kueue
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: kueue-system
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
      - ServerSideApply=true           # CRITICAL: Bypasses 262KB annotation limit on Kueue/JobSet CRDs
```

#### Config Sync
[Config Sync](https://cloud.google.com/kubernetes-engine/enterprise/config-sync/docs/overview) applies resources using Server-Side Apply by default, so no additional configuration is required to avoid client-side annotation overflow when managing large schemas.

---

## Safe upgrade workflow

When upgrading Config Connector to a newer release:

1. **Update Manifests**: Download the new release tarball (Method A) or pull upstream updates and re-run `./dev/tasks/build-release-bundle` (Method B).
2. **Commit to GitOps Repo**: Push the updated `crds.yaml` and `0-cnrm-system.yaml` to your GitOps repository branch.
3. **ArgoCD In-Place Sync**: ArgoCD applies the updated CRD OpenAPI schemas via Server-Side Apply.
4. **Verification**:
   - Verify CRDs updated in-place: `kubectl get crd containerclusters.container.cnrm.cloud.google.com` (creation timestamp should remain unchanged).
   - Check operator status: `kubectl get configconnector -o yaml`.
   - Existing GCP resources and running workloads remain completely uninterrupted.
