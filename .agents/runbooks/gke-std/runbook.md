# Runbook: Deploy Config Connector from Source to GKE (gke1)

## What this needs

### Infrastructure Requirement
This run **cannot** run purely within a standalone pod and **requires real cloud infrastructure**. Config Connector (KCC) is a suite of Kubernetes controllers that directly manage Google Cloud Platform resources. It requires:
1. A **Google Kubernetes Engine (GKE) cluster** with **Workload Identity** enabled (`<PROJECT_ID>.svc.id.goog`) so that in-cluster Kubernetes Service Accounts can securely authenticate to GCP APIs without static long-lived credentials.
2. A **Google Service Account (GSA)** with project-level IAM roles (e.g. `roles/owner` or `roles/editor`) to manage GCP resources.
3. An **IAM Policy Binding** linking the Kubernetes Service Account (`cnrm-controller-manager` in namespace `cnrm-system`) to the GSA (`roles/iam.workloadIdentityUser`).
4. An **Artifact Registry repository** to host the container images built from source (`controller`, `recorder`, `webhook`, `deletiondefender`, `unmanageddetector`, `operator`).
5. **Custom Resource Definitions (CRDs)** installed into the Kubernetes API server for all supported GCP services.

### Feasibility & Prerequisites Checklist
- [ ] **Tool: gcloud CLI** — Present (`Google Cloud SDK 586.0.0`).
- [ ] **Tool: kubectl** — Present (`v1.35.8-dispatcher` with built-in Kustomize engine `v5.7.1`).
- [ ] **Tool: go** — Present (`go1.26.4`).
- [ ] **Tool: git** — Present (`git version 2.43.0`).
- [ ] **Tool: docker / container builder** — Local docker daemon is missing in this environment. Container builds are delegated to **Cloud Build** via `gcloud builds submit`, which runs natively on GCP.
- [ ] **API: Kubernetes Engine (`container.googleapis.com`)** — Enabled.
- [ ] **API: Cloud Build (`cloudbuild.googleapis.com`)** — Enabled.
- [ ] **API: Artifact Registry (`artifactregistry.googleapis.com`)** — Enabled.
- [ ] **API: IAM (`iam.googleapis.com`)** — Enabled.
- [ ] **API: Pub/Sub (`pubsub.googleapis.com`)** — Enabled for verification testing.
- [ ] **IAM Project Role: `roles/owner` / `roles/editor`** — Granted to the active deployer identity (`cnrm-barni-1.svc.id.goog`).
- [ ] **IAM Service Account Management: `iam.serviceAccounts.create`, `iam.serviceAccounts.setIamPolicy`** — Granted.
- [ ] **Cluster Management: `container.clusters.create`, `container.clusters.get`** — Granted.

---

## Preconditions

1. Target GCP project is determined and active in `gcloud`.
2. Target compute region and zone are set.
3. Parameters are defined in `params.env` with exact literal values matching the assigned `${RESOURCE_PREFIX}`.
4. Repository working tree is at the target commit to build from source.

---

## Steps

### Step 1: Create Artifact Registry Repository
Create an Artifact Registry Docker repository to store container images built from source for this run. Tag it with the `repo-agent-instance` label for traceability.

```bash
gcloud artifacts repositories create "${ARTIFACT_REPO}" \
    --repository-format=docker \
    --location="${REGION}" \
    --description="KCC source build images for ${RESOURCE_PREFIX}" \
    --labels="repo-agent-instance=${RESOURCE_PREFIX}" \
    --project="${PROJECT}" || true
```

*Why:* Config Connector components (`controller`, `recorder`, `webhook`, `deletiondefender`, `unmanageddetector`) must be built from the local source tree and hosted in a registry accessible to the GKE cluster.

---

### Step 2: Create Google Service Account (GSA) and Assign Project IAM Permissions
Create a dedicated Google Service Account for Config Connector and grant it the project-level role required to manage GCP resources. Also ensure Cloud Build's default service account has the requisite permissions (`storage.admin`, `artifactregistry.writer`, `logging.logWriter`) to read source archives and publish container images.

```bash
gcloud iam service-accounts create "${GSA_NAME}" \
    --display-name="${GSA_NAME}" \
    --description="GSA for Config Connector in ${CLUSTER_NAME}" \
    --project="${PROJECT}" || true

gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/owner" \
    --quiet

PROJECT_NUMBER="$(gcloud projects describe "${PROJECT}" --format="value(projectNumber)")"
CB_SA="${PROJECT_NUMBER}-compute@developer.gserviceaccount.com"
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/storage.admin" \
    --quiet
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/artifactregistry.writer" \
    --quiet
gcloud projects add-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/logging.logWriter" \
    --quiet

gcloud storage buckets add-iam-policy-binding "gs://${PROJECT}_cloudbuild" \
    --member="serviceAccount:${CB_SA}" \
    --role="roles/storage.admin" \
    --quiet 2>/dev/null || true
```

*Why:* Config Connector reconcilers act on behalf of this Google Service Account when creating, updating, and deleting cloud infrastructure in the project. Cloud Build execution also requires bucket read access on the upload bucket and Artifact Registry write access to push built images.

---

### Step 3: Create GKE Cluster with Workload Identity Enabled
Provision a GKE cluster with Workload Identity configured to use the project's workload pool `${PROJECT}.svc.id.goog`.

```bash
gcloud container clusters create "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}" \
    --workload-pool="${PROJECT}.svc.id.goog" \
    --num-nodes=3 \
    --machine-type=e2-standard-4 \
    --labels="repo-agent-instance=${RESOURCE_PREFIX}"

gcloud container clusters get-credentials "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}"
```

*Why:* Workload Identity is the standard, secure mechanism for pods in GKE to assume Google Service Account identities without mounting private keys.

---

### Step 4: Configure Workload Identity Binding for the KCC Kubernetes Service Account
Bind the GSA to the Kubernetes Service Account `cnrm-controller-manager` in the `cnrm-system` namespace.

```bash
gcloud iam service-accounts add-iam-policy-binding "${GSA_EMAIL}" \
    --member="serviceAccount:${PROJECT}.svc.id.goog[cnrm-system/cnrm-controller-manager]" \
    --role="roles/iam.workloadIdentityUser" \
    --project="${PROJECT}"
```

*Why:* When the `cnrm-controller-manager` pod runs with the `iam.gke.io/gcp-service-account` annotation, GKE's metadata server checks this binding to issue short-lived Google OAuth tokens.

---

### Step 5: Build Container Images from Source using Cloud Build
Submit a Cloud Build job to build all Config Connector component binaries and container images from source (with BuildKit enabled for cache mounts), and push them to the Artifact Registry repository.

```bash
if ! gcloud artifacts docker images list "${IMAGE_PREFIX%/}" --include-tags --filter="TAGS:${IMAGE_TAG}" --format="value(TAGS)" 2>/dev/null | grep -q "${IMAGE_TAG}"; then
    gcloud builds submit \
        --project="${PROJECT}" \
        --config="${RUN_DIR}/cloudbuild.yaml" \
        --substitutions="_IMAGE_PREFIX=${IMAGE_PREFIX},_IMAGE_TAG=${IMAGE_TAG}" \
        .
fi
```

*Why:* Because a local Docker daemon is not running in this environment, Cloud Build provides a hermetic build environment. The build requires `DOCKER_BUILDKIT=1` because `build/builder/Dockerfile` leverages BuildKit `--mount=type=cache` options. Re-runs skip the build when images are already present.

---

### Step 6: Generate Kustomize Image Patches
Generate the component image patch files from the repository's patch templates, pointing each component to the newly built image in Artifact Registry and sizing memory limits appropriately.

```bash
cp config/installbundle/components/manager/base/manager_image_patch_template.yaml config/installbundle/components/manager/base/manager_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}controller:${IMAGE_TAG}@" config/installbundle/components/manager/base/manager_image_patch.yaml

cat <<EOF > config/installbundle/components/recorder/recorder_image_patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resource-stats-recorder
spec:
  template:
    spec:
      containers:
      - image: ${IMAGE_PREFIX}recorder:${IMAGE_TAG}
        name: recorder
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 512Mi
EOF

cat <<EOF > config/installbundle/components/webhook/webhook_image_patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webhook-manager
spec:
  template:
    spec:
      containers:
      - image: ${IMAGE_PREFIX}webhook:${IMAGE_TAG}
        name: webhook
        resources:
          limits:
            memory: 512Mi
          requests:
            cpu: 250m
            memory: 512Mi
        env:
        - name: GOMEMLIMIT
          value: 460MiB
EOF

cp config/installbundle/components/deletiondefender/deletiondefender_image_patch_template.yaml config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}deletiondefender:${IMAGE_TAG}@" config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml

cp config/installbundle/components/unmanageddetector/unmanageddetector_image_patch_template.yaml config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml
sed -i -e "s@image: .*@image: ${IMAGE_PREFIX}unmanageddetector:${IMAGE_TAG}@" config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml
```

*Why:* The repo's Kustomize base expects concrete `*_image_patch.yaml` files generated from `*_image_patch_template.yaml` to substitute custom container images into the deployment manifests. In addition, `webhook` and `recorder` memory requests/limits must be raised to 512Mi (`GOMEMLIMIT=460MiB`) to prevent container OOMKilled crashes during dynamic validation across the 600+ installed CRD schemas.

---

### Step 7: Install Config Connector CRDs
Apply all Custom Resource Definitions (both core operator CRDs and resource CRDs) into the GKE cluster.

```bash
kubectl apply -f operator/config/crd/bases/
kubectl apply -f config/crds/resources/
```

*Why:* `cnrm-controller-manager` requires core CRDs (`ConfigConnector`, `ConfigConnectorContext`) to initialize its registration controllers without failing on missing API group `core.cnrm.cloud.google.com/v1beta1`. Resource CRDs are required before user manifests can be submitted.

---

### Step 8: Deploy Config Connector Controller and Components
Render the Kustomize release manifests for cluster-scoped Workload Identity mode, substitute the GSA email and project ID, and apply them to the cluster.

```bash
kubectl kustomize config/installbundle/releases/scopes/cluster/withworkloadidentity | \
    sed -e "s/cnrm-system@\${PROJECT_ID?}\.iam\.gserviceaccount\.com/${GSA_EMAIL}/g" | \
    sed -e "s/\${PROJECT_ID?}/${PROJECT}/g" | \
    kubectl apply -f -
```

*Why:* This creates the `cnrm-system` namespace, RBAC roles/bindings, CRD webhooks, deletion defenders, unmanaged resource detectors, and the main `cnrm-controller-manager` deployment configured with Workload Identity.

---

### Step 9: Configure Managed Test Namespace
Create a dedicated test namespace and annotate it with the target GCP project ID.

```bash
kubectl create namespace "${TEST_NAMESPACE}" --dry-run=client -o yaml | kubectl apply -f -
kubectl annotate namespace "${TEST_NAMESPACE}" "cnrm.cloud.google.com/project-id=${PROJECT}" --overwrite
```

*Why:* By default, Config Connector maps Kubernetes namespaces to GCP projects using the `cnrm.cloud.google.com/project-id` annotation.

---

## Verify

### Step 1: Verify Pods in `cnrm-system` Namespace
Ensure that all Config Connector pods are scheduled, running, and pass their readiness checks.

```bash
kubectl wait -n cnrm-system --for=condition=Ready pod --all --timeout=300s
kubectl get pods -n cnrm-system
```

---

### Step 2: Create a Test GCP Resource via KCC CRD
Submit a `PubSubTopic` Custom Resource in the test namespace.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: "${TEST_TOPIC_NAME}"
  namespace: "${TEST_NAMESPACE}"
  labels:
    repo-agent-instance: "${RESOURCE_PREFIX}"
EOF
```

---

### Step 3: Verify Resource Reconciliation Status
Wait for the Config Connector controller to reconcile the resource and mark its status as `Ready: True` with reason `UpToDate`.

```bash
kubectl wait -n "${TEST_NAMESPACE}" --for=condition=Ready "pubsubtopic/${TEST_TOPIC_NAME}" --timeout=180s
kubectl get pubsubtopic "${TEST_TOPIC_NAME}" -n "${TEST_NAMESPACE}" -o yaml
```

---

### Step 4: Verify Resource Exists on Real GCP
Use `gcloud` to verify that the Pub/Sub topic actually exists in the GCP project.

```bash
gcloud pubsub topics describe "${TEST_TOPIC_NAME}" --project="${PROJECT}"
```

---

## Teardown

### Step 1: Delete Test KRM Resource and Namespace
Delete the test resource so Config Connector cleans up the underlying GCP resource before cluster deletion.

```bash
kubectl delete pubsubtopic "${TEST_TOPIC_NAME}" -n "${TEST_NAMESPACE}" --ignore-not-found=true --timeout=120s
kubectl delete namespace "${TEST_NAMESPACE}" --ignore-not-found=true
```

---

### Step 2: Delete Config Connector Manifests and CRDs
Remove all Config Connector controllers and CRDs from the cluster.

```bash
kubectl kustomize config/installbundle/releases/scopes/cluster/withworkloadidentity | \
    sed -e "s/cnrm-system@\${PROJECT_ID?}\.iam\.gserviceaccount\.com/${GSA_EMAIL}/g" | \
    sed -e "s/\${PROJECT_ID?}/${PROJECT}/g" | \
    kubectl delete -f - --ignore-not-found=true

kubectl delete -f operator/config/crd/bases/ --ignore-not-found=true
kubectl delete -f config/crds/resources/ --ignore-not-found=true
```

---

### Step 3: Delete GKE Cluster
Delete the GKE cluster.

```bash
gcloud container clusters delete "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}" \
    --quiet
```

---

### Step 4: Remove IAM Bindings and Delete GSA
Remove the project-level IAM policy binding and delete the Google Service Account.

```bash
gcloud projects remove-iam-policy-binding "${PROJECT}" \
    --member="serviceAccount:${GSA_EMAIL}" \
    --role="roles/owner" \
    --quiet || true

gcloud iam service-accounts delete "${GSA_EMAIL}" \
    --project="${PROJECT}" \
    --quiet || true
```

---

### Step 5: Delete Artifact Registry Repository
Delete the Artifact Registry repository created for this run.

```bash
gcloud artifacts repositories delete "${ARTIFACT_REPO}" \
    --location="${REGION}" \
    --project="${PROJECT}" \
    --quiet || true
```

---

### Step 6: Clean Up Local Generated Patch Files
Remove generated patch files from the workspace.

```bash
rm -f config/installbundle/components/manager/base/manager_image_patch.yaml
rm -f config/installbundle/components/recorder/recorder_image_patch.yaml
rm -f config/installbundle/components/webhook/webhook_image_patch.yaml
rm -f config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml
rm -f config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml
```
