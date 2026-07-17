# Managing Config Connector with GitOps (ArgoCD)

This guide provides a production-tested workflow for installing, managing, and upgrading Config Connector (KCC) using GitOps tooling like **ArgoCD** or **Flux**.

---

## Architecture & Deployment Options

Config Connector can be deployed and managed via GitOps using two primary patterns:

1. **Operator-Based GitOps (Recommended)**: ArgoCD deploys the Config Connector Operator via pre-rendered release bundles. The Operator then manages CRD lifecycle and controller deployments automatically based on declarative `ConfigConnector` custom resources.
2. **Pre-Rendered Manifest GitOps**: ArgoCD directly manages the rendered release bundle manifests (`0-cnrm-system.yaml` and resource CRDs) rendered from Kustomize templates.

---

## Important Gotchas & Prerequisites

Before setting up ArgoCD applications for KCC, note the following critical requirements:

### 1. Raw Upstream Repo vs. Rendered Manifests
The raw source repository (`operator/config/default`) contains release patch templates (`manager_image_patch_template.yaml`) that require build-time rendering. Attempting to point ArgoCD directly at raw upstream source paths without generating `manager_image_patch.yaml` will result in Kustomize build errors.

* **Solution**: Use official published release bundles (downloadable from Google Cloud Storage), or render the Kustomize targets using `dev/tasks/build-release-bundle` in your CI pipeline before committing the manifests to your GitOps repository.

### 2. Mandatory Server-Side Apply for CRDs
KCC includes over 250 CRDs with extensive OpenAPI schemas. Standard client-side `kubectl apply` will fail on large CRD bundles due to Kubernetes' 262KB `kubectl.kubernetes.io/last-applied-configuration` annotation limit in etcd (`metadata.annotations: Too long`).

* **Solution**: You **MUST enable Server-Side Apply** in your ArgoCD Application sync options (`ServerSideApply=true`).

### 3. CRD Upgrades & Safety Rules
* **In-Place Upgrades**: CRDs update **in-place** when applying new versions. etcd object creation timestamps and live GCP resources remain untouched.
* ⚠️ **NEVER DELETE CRDs**: Running `kubectl delete crd` cascade-deletes all Custom Resource (CR) instances across all namespaces. Because KCC resources have finalizers attached by default, deleting CR instances will instruct KCC to **delete the underlying production GCP infrastructure** in Google Cloud.

---

## Step-by-Step Installation Guide

### Step 1: Obtain or Render the Installation Manifests

You can obtain the required pre-rendered KCC installation manifests using either of the following sources:

#### **Method A: Official Google Cloud Published Release Bundles (Recommended)**
If you are installing KCC without building from source, download the latest official release bundle tarball from Google Cloud Storage:
```bash
# Download the official release bundle
curl -O https://storage.googleapis.com/config-connector-operator/latest/release-bundle.tar.gz
tar -xvf release-bundle.tar.gz
```
This tarball contains pre-rendered, production-ready manifests:
* `0-cnrm-system.yaml`: Contains the Operator deployment, RBAC roles, ServiceAccounts, and `configconnector-operator-system` namespace.
* `crds.yaml`: Contains the complete set of KCC CustomResourceDefinitions.

#### **Method B: Building from GitHub Repository Source**
If you are tracking the open-source GitHub repository source code:
```bash
git clone https://github.com/GoogleCloudPlatform/k8s-config-connector.git
cd k8s-config-connector

# Render standard cluster bundle
./dev/tasks/build-release-bundle
```
This populates rendered manifests in `config/installbundle/releases/` ready for GitOps use.

#### **Recommended GitOps Repository Structure**
Organize your team's GitOps repository (e.g., `https://github.com/<your-org>/kcc-gitops.git`) as follows:
```text
kcc-gitops/
├── operator/
│   └── 0-cnrm-system.yaml    # Operator deployment & RBAC
└── crds/
    └── crds.yaml              # Complete KCC CRD bundle
```

---

### Step 2: Configure ArgoCD Application for KCC Operator

Create an ArgoCD `Application` targeting your rendered KCC manifests directory in your GitOps repository. Be sure to include `ServerSideApply=true` in `syncOptions`:

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
    path: operator
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: configconnector-operator-system
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
      - ServerSideApply=true   # CRITICAL: Bypasses 262KB annotation limit on CRDs
```

---

### Step 3: Declaratively Configure Config Connector

Once the Operator is deployed by ArgoCD, declaratively configure Config Connector by committing a `ConfigConnector` custom resource to your GitOps repository.

> 💡 **Prerequisite**: Ensure the specified Google Service Account (`googleServiceAccount`) has been created in GCP IAM and bound to the KCC ServiceAccount via Workload Identity.

#### **Cluster Mode Configuration**:
```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  googleServiceAccount: "kcc-system@<YOUR_PROJECT_ID>.iam.gserviceaccount.com"
```

#### **Namespace Mode Configuration**:
For namespace mode, deploy a `ConfigConnectorContext` resource in each managed namespace to bind dedicated GCP Service Accounts per namespace:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: config-control
spec:
  googleServiceAccount: "kcc-namespace-sa@<YOUR_PROJECT_ID>.iam.gserviceaccount.com"
```

---

## Safe Upgrade Workflow

When upgrading Config Connector to a newer release:

1. **Update Manifests**: Download the new release tarball (Method A) or pull upstream updates and re-run `./dev/tasks/build-release-bundle` (Method B).
2. **Commit to GitOps Repo**: Push the updated `crds.yaml` and `0-cnrm-system.yaml` to your GitOps repository branch.
3. **ArgoCD In-Place Sync**: ArgoCD applies the updated CRD OpenAPI schemas via Server-Side Apply.
4. **Verification**:
   - Verify CRDs updated in-place: `kubectl get crd containerclusters.container.cnrm.cloud.google.com` (creation timestamp should remain unchanged).
   - Check operator status: `kubectl get configconnector -o yaml`.
   - Existing GCP resources and running workloads will remain completely uninterrupted.
