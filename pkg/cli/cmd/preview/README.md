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
