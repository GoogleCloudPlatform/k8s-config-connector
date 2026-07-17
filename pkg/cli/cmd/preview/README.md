# Config Connector Preview Command

The `config-connector preview` command allows you to preview the reconciliation behavior of Config Connector resources. It is particularly useful for validating how resources will behave when migrated from one controller type to another (e.g., from Terraform-based controllers to the "direct" approach).

## Overview

The `preview` command performs two separate reconciliation runs:

1.  **Default Run (Baseline):** Reconciles resources using their default configured controllers. This establishes a baseline of the current "healthy" state of your resources.
2.  **Alternative Run:** Reconciles the same resources but forces them to use an alternative controller type (if supported for that resource kind). This is typically used to test the "direct" controller implementation for resources currently managed by Terraform or DCL.

After both runs complete, the tool generates a summary report comparing the results, highlighting any differences in resource state (diffs) or reconciliation health.

## Requirements

-   The `config-connector` CLI must be installed. See the [installation guide](../../../../docs/cli/README.md) for details on installing via `gcloud`.
-   The tool requires credentials to access the Kubernetes cluster (via `kubeconfig`).
-   The tool requires GCP credentials (ADC) to perform read-only checks against GCP resources to verify state.

## How it Works

1.  **Resource Discovery:** The tool connects to your Kubernetes cluster and lists all Config Connector resources (either in a specific namespace or across all namespaces).
2.  **In-Memory Reconciliation:** It uses an internal manager that mimics the standard Config Connector controller but captures all interactions (Kube API calls, GCP API calls, and object diffs) without actually applying changes to the cluster or GCP (it uses interceptors).
3.  **Analysis:** It tracks whether resources reached a "Healthy" state (successfully reconciled without unexpected GCP side effects) or an "Unhealthy" state.
4.  **Reporting:** It outputs a summary report and detailed logs of any discrepancies found during the alternative run.

## Usage

### Basic Usage

Preview all resources in the current context:

```bash
config-connector preview
```

Preview resources in a specific namespace:

```bash
config-connector preview --namespace my-namespace
```

### Advanced Usage

Enable full reports with detailed object events:

```bash
config-connector preview --full-report
```

Adjust GCP API rate limiting:

```bash
config-connector preview --gcp-qps 10 --gcp-burst 10
```

## Flags

| Flag | Description | Default |
| :--- | :--- | :--- |
| `--kubeconfig` | Path to the kubeconfig file. | Uses default rules |
| `--namespace`, `-n` | Namespace to preview. If not specified, all namespaces are previewed. | All namespaces |
| `--timeout` | Timeout for each reconciliation run in minutes. | 15 |
| `--report-prefix` | Prefix for the generated report files. A timestamp is appended. | `preview-report` |
| `--full-report`, `-f` | Write a full report with all captured objects and events. | `false` |
| `--gcp-qps`, `-q` | Maximum QPS for GCP API requests per service. Set to 0 to disable. | 5.0 |
| `--gcp-burst`, `-b` | Maximum burst for GCP API requests per service. | 5 |
| `--in-cluster` | Run the tool within a GKE cluster (uses in-cluster config). | `false` |
| `--verbose`, `-v` | Log verbosity level (klog). | 0 |

## Reports

The tool generates several files:

-   `preview-report-<timestamp>`: A summary table showing the Group, Kind, Name, and the reconciliation status/diffs for both the default and alternative runs.
-   `preview-report-<timestamp>-detail`: (Generated if errors occur) A JSON file containing detailed information about "Unhealthy" resources.
-   `preview-report-<timestamp>-full-default`: (Generated if `--full-report` is set) Detailed event log for the default run.
-   `preview-report-<timestamp>-full-alternative`: (Generated if `--full-report` is set) Detailed event log for the alternative run.

### Understanding the Summary Report

The summary report contains the following columns:

-   **GROUP/KIND/NAME**: Identifies the resource.
-   **DEFAULT-CONTROLLER**: The controller used in the baseline run (e.g., `Terraform`, `Direct`).
-   **DEFAULT-RESULT**: `HEALTHY` or `UNHEALTHY`.
-   **DEFAULT-DIFFS**: Fields that had diffs during reconciliation.
-   **ALTERNATIVE-CONTROLLER**: The alternative controller being tested.
-   **ALTERNATIVE-RESULT**: `HEALTHY`, `UNHEALTHY`, or `Missing` (if the alternative controller failed to pick up the resource).
-   **ALTERNATIVE-DIFFS**: Fields that had diffs during the alternative run.

---

## User Instructions & Execution Methods

### Method 1: Build & Run from Local Source

Recommended for developers testing against a specific local or release branch.

1. **Authenticate with GCP and Kubernetes**:
   ```bash
   gcloud auth application-default login
   gcloud container clusters get-credentials <CLUSTER_NAME> --region <REGION> --project <PROJECT_ID>
   ```

2. **Check Out Target Version & Build**:
   ```bash
   git clone https://github.com/GoogleCloudPlatform/k8s-config-connector.git
   cd k8s-config-connector
   git checkout tags/v1.153.0   # Or your target release branch
   go build -o config-connector ./cmd/config-connector
   ```

3. **Run Preview**:
   ```bash
   mkdir -p reports
   ./config-connector preview --kubeconfig ~/.kube/config --report-prefix reports/preview-report --full-report
   ```

---

### Method 2: Run via Docker Image Locally

Recommended for administrators who want to run the pre-built release container without a local Go setup.

1. **Authenticate on Host**:
   ```bash
   gcloud auth application-default login
   gcloud container clusters get-credentials <CLUSTER_NAME> --region <REGION> --project <PROJECT_ID>
   ```

2. **Run Container Image**:
   ```bash
   mkdir -p reports
   docker run --rm -it \
     --net=host \
     -u $(id -u):$(id -g) \
     -v ~/.kube/config:/configconnector/.kube/config:ro \
     -v ~/.config/gcloud:/configconnector/.config/gcloud:ro \
     -v /usr/bin/gke-gcloud-auth-plugin:/usr/bin/gke-gcloud-auth-plugin:ro \
     -v /usr/bin/gcloud:/usr/bin/gcloud:ro \
     -v /usr/lib/google-cloud-sdk:/usr/lib/google-cloud-sdk:ro \
     -v $(pwd)/reports:/configconnector/reports \
     -e KUBECONFIG=/configconnector/.kube/config \
     -e GOOGLE_APPLICATION_CREDENTIALS=/configconnector/.config/gcloud/application_default_credentials.json \
     gcr.io/gke-release/cnrm/config-connector-cli:1.153.0 \
     preview --kubeconfig=/configconnector/.kube/config --report-prefix=/configconnector/reports/preview-report
   ```

---

### Method 3: Run In-Cluster as a Kubernetes Job

Recommended for automated CI/CD pipelines or cluster-level health audits using GKE Workload Identity. When running in-cluster, specify the `--in-cluster` flag.

#### Namespace and Service Account Configuration

The Job is deployed in the `cnrm-system` namespace where Config Connector controllers run:
- **Cluster Mode**: Launch a single Job using the controller manager Service Account (`cnrm-controller-manager-config-control` or cluster SA).
- **Namespace Mode**: Launch a separate Job for each enabled namespace with:
  1. `--namespace <TARGET_NAMESPACE>` in the container command arguments.
  2. `serviceAccount` and `serviceAccountName` set to `cnrm-controller-manager-<TARGET_NAMESPACE>`.

#### 1. Create the Job Manifest (`preview-job.yaml`)

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: preview-job-config-control
  namespace: cnrm-system
spec:
  backoffLimit: 0
  activeDeadlineSeconds: 900        # 15 minute timeout
  ttlSecondsAfterFinished: 86400    # Retain completed Job for 24 hours
  template:
    metadata:
      labels:
        cnrm.cloud.google.com/component: cnrm-controller-manager
    spec:
      serviceAccount: cnrm-controller-manager-config-control       # GSA for target namespace
      serviceAccountName: cnrm-controller-manager-config-control   # KSA for target namespace
      restartPolicy: Never
      terminationGracePeriodSeconds: 10
      volumes:
        - name: preview-volume
          emptyDir: {}
      containers:
        - name: preview
          image: gcr.io/gke-release/cnrm/config-connector-cli:1.153.0
          imagePullPolicy: Always
          workingDir: /workspace
          command: ["/bin/sh", "-c"]
          args:
            - |
              config-connector preview --in-cluster --namespace config-control --timeout 15 --report-prefix /workspace/preview-report
              echo ""
              echo "==================== PREVIEW SUMMARY REPORT ===================="
              cat /workspace/preview-report-*
          volumeMounts:
            - name: preview-volume
              mountPath: /workspace
          env:
            - name: GOMEMLIMIT
              value: "1840MiB"
          resources:
            requests:
              cpu: "500m"
              memory: "2048Mi"
            limits:
              memory: "2048Mi"
          securityContext:
            privileged: false
            runAsUser: 1000
            runAsNonRoot: true
            allowPrivilegeEscalation: false
```

#### 2. Deploy the Job & View Output

##### Option A: View the Summary Report directly in terminal
```bash
# 1. Apply the preview Job
kubectl apply -f preview-job.yaml

# 2. Wait for completion
kubectl wait --for=condition=complete job/preview-job-config-control -n cnrm-system --timeout=15m

# 3. View the generated summary table directly in your terminal
kubectl logs job/preview-job-config-control -n cnrm-system
```

##### Option B: Copy Report Files Locally (`kubectl cp`)
```bash
# 1. Get the pod name created by the Job
POD_NAME=$(kubectl get pods -n cnrm-system -l job-name=preview-job-config-control -o jsonpath='{.items[0].metadata.name}')

# 2. Copy the report files from the container workspace to your local directory
mkdir -p reports
kubectl cp cnrm-system/${POD_NAME}:workspace/ reports/
```
