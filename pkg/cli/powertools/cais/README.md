# CAIS Identity Export Powertool

The `cais` powertool calculates and exports the Cloud Asset Inventory (CAIS) identity/URL for Config Connector (KCC) resources. It is useful for understanding how KCC resources map to GCP Asset Inventory paths and determining the canonical Google Cloud Resource URL for one or more resources.

> **WARNING:** This is an **experimental** powertool and is subject to change or removal in future releases without notice.

## Overview

Cloud Asset Inventory (CAI) uses specific URL structures to identify resources across Google Cloud (e.g., `//dns.googleapis.com/projects/my-project-id/managedZones/my-zone`). Config Connector maps its Kubernetes resources to these CAIS URLs using internal identity-mapping logic.

The powertool supports two main modes of operation:
1. **Local Files / Directories Mode:** Parses local YAML manifests and calculates CAIS identities offline.
2. **Live Cluster Mode:** Queries resources in a running Kubernetes cluster and resolves CAIS identities using live object state.

---

## How to Run

Since this tool is part of the experimental `powertools` suite, you can run it directly using `go run` from the root of the Config Connector repository:

```bash
go run cmd/config-connector/main.go powertools cais [flags]
```

Alternatively, if the `config-connector` binary is compiled and installed on your system, you can run:

```bash
config-connector powertools cais [flags]
```

---

## Input Modes

### 1. Static Files or Directories (`--file`, `--dir`, `--stdin`)

You can run the powertool against local YAML files or directory trees containing Kubernetes manifests.

```bash
# Read a single manifest file
go run cmd/config-connector/main.go powertools cais -f config/samples/resources/project/project-in-folder/resourcemanager_v1beta1_project.yaml

# Read a directory recursively (shows both strengths and weaknesses well)
go run cmd/config-connector/main.go powertools cais -d config/samples/resources/pubsubsubscription/bigquery-pubsub-subscription

# Read from standard input
cat config/samples/resources/project/project-in-folder/resourcemanager_v1beta1_project.yaml | go run cmd/config-connector/main.go powertools cais --stdin
```

> ⚠️ **Important Limitation:** When parsing local files, any server-generated IDs (such as Google Cloud project numbers, folder IDs, or randomly assigned resource hashes) that are only known after successful GCP reconciliation will be listed as `unknown`. This is because the local YAML spec does not contain the live status fields where these IDs are stored.

### 2. Live Kubernetes Cluster Mode

If no local file, directory, or standard input flag is provided, the tool defaults to **Live Cluster Mode**. It connects to your active Kubernetes cluster (using the default kubeconfig context or the path specified via `--kubeconfig`) and reads live resources.

Because the tool can inspect live Kubernetes resources, it reads the status blocks which contain GCP-assigned properties (like project numbers, folder IDs, etc.). This allows the powertool to resolve **full, exact CAIS URLs** with no `unknown` parts.

```bash
# Get CAIS identities for all KCC resources in a specific namespace in the cluster
go run cmd/config-connector/main.go powertools cais --match-namespace my-namespace

# Get CAIS identity for a specific resource by kind and name
go run cmd/config-connector/main.go powertools cais --match-kind DNSManagedZone --match-name my-zone --match-namespace my-namespace
```

---

## Flags and Filtering

The powertool supports various flags to control input sources, target specific resources, or format the output.

### Target and Filtering Flags

| Flag | Description |
| :--- | :--- |
| `--stdin` | Read YAML manifests from standard input. |
| `-f`, `--file` | Path to a file containing a YAML manifest. |
| `-d`, `--dir` | Path to a directory containing YAML manifests (scanned recursively). |
| `--match-kind` | Filter results to match only the specified Kubernetes Kind (e.g., `StorageBucket`). |
| `--match-name` | Filter results to match only the specified resource Name. |
| `--match-namespace` | Filter results to match only the specified Kubernetes Namespace. |

### Cluster Configuration Flags (Live Cluster Mode)

| Flag | Description |
| :--- | :--- |
| `--kubeconfig` | Path to the kubeconfig file to use for CLI requests. |
| `--as` | Username to impersonate for the operation. |
| `--as-group` | Group(s) to impersonate for the operation (can be repeated). |

### Output Formatting

The tool can output results in multiple formats using the `--format` flag:

| Format | Description |
| :--- | :--- |
| `text` (Default) | A human-readable, tab-separated table. |
| `yaml` or `yml` | Structured YAML array of CAIS identity results. |
| `json` | Structured JSON array of CAIS identity results. |

Example command requesting JSON output:
```bash
go run cmd/config-connector/main.go powertools cais -f config/samples/resources/project/project-in-folder/resourcemanager_v1beta1_project.yaml --format json
```

---

## Structured Output

The output includes metadata about the resource and the calculated CAIS URL:

```yaml
- caisURL: //dns.googleapis.com/projects/my-project-id/managedZones/test-zone
  group: dns.cnrm.cloud.google.com
  version: v1beta1
  kind: DNSManagedZone
  name: test-zone
  namespace: test-namespace
```

If the resource kind is not supported, the `error` field in the output will indicate the reason (e.g., `not yet supported`).

> 📝 **Note:** As this is an experimental powertool, the structured output format and schema are subject to change as we continue its development.
