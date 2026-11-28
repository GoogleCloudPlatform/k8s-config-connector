# Monitoring Config Connector with Prometheus and Grafana

This guide explains how to scrape Prometheus metrics from Config Connector controllers and import the Grafana dashboard model ([`grafana-metrics-dashboard.json`](grafana-metrics-dashboard.json)).

---

## Metrics Endpoints and Services

Config Connector exposes Prometheus metrics in the `cnrm-system` namespace on port `8888`. Depending on your installation method, the controller manager Service is named `cnrm-controller-manager-service` or `cnrm-manager` (or `cnrm-manager-<namespace>` in namespaced mode), and the resource stats recorder Service is named `cnrm-resource-stats-recorder-service`. Both Services carry the label `cnrm.cloud.google.com/monitored: "true"` and annotations `prometheus.io/scrape: "true"` and `prometheus.io/port: "8888"`.

The controller manager HTTP server (`--prometheus-scrape-endpoint=:8888`) registers up to three paths:

| Component | HTTP Path | Exposed Metrics |
| :--- | :--- | :--- |
| Controller Manager | `/metrics` | OpenCensus reconciliation metrics (`configconnector_reconcile_requests_total`, `configconnector_reconcile_request_duration_seconds`, `configconnector_reconcile_occupied_workers_total`, `configconnector_reconcile_workers_total`, `configconnector_internal_errors_total`) |
| Controller Manager | `/prom-metrics` | Go runtime, process, and HTTP transport metrics (`process_resident_memory_bytes`, `process_cpu_seconds_total`, `go_goroutines`, `configconnector_gcp_api_requests_total`, `configconnector_gcp_api_request_duration_seconds`) |
| Controller Manager | `/metrics.v2` | Standard `controller-runtime` metrics (`workqueue_depth`, `workqueue_adds_total`, `workqueue_queue_duration_seconds_bucket`) when `METRICS_VERSION=v2` is set on the manager container (or on `--metrics-bind-address`, default `:8080`) |
| Resource Stats Recorder | `/metrics` | Applied custom resource counts (`configconnector_applied_resources_total` / `configconnector_applied_objects_total`, `configconnector_build_info`) |

---

## Exposed Metrics Reference

| Metric Name | Type | Labels | Description |
| :--- | :--- | :--- | :--- |
| `configconnector_reconcile_requests_total` | Counter | `group_version_kind`, `namespace`, `status` (`OK` \| `ERROR`), optional `name` | Total number of reconciliation attempts recorded by controller workers. |
| `configconnector_reconcile_request_duration_seconds` | Histogram | `group_version_kind`, `namespace`, `status`, optional `name` | Latency distribution of reconciliation requests in seconds (`0s` to `3600s` buckets). |
| `configconnector_reconcile_occupied_workers_total` | Gauge | `group_version_kind` | Number of active reconciliation workers currently processing resources of a given kind. |
| `configconnector_reconcile_workers_total` | Gauge | `group_version_kind` | Total configured concurrent reconciliation workers for a given kind. |
| `configconnector_internal_errors_total` | Counter | `group_version_kind`, `namespace` | Total number of internal controller errors encountered during reconciliation. |
| `configconnector_applied_resources_total` (`configconnector_applied_objects_total`) | Gauge | `group_version_kind`, `namespace`, `status` | Count of Config Connector custom resources in the cluster grouped by `Ready` condition reason (for example, `UpToDate`, `Updating`, `UpdateFailed`, `NoCondition`). |
| `workqueue_depth` | Gauge | `name` | Current depth of the `controller-runtime` workqueue per controller. |
| `process_resident_memory_bytes` | Gauge | Pod / instance labels | Resident set size (RSS) memory usage of the controller process in bytes. |
| `process_cpu_seconds_total` | Counter | Pod / instance labels | Total user and system CPU time spent by the controller process in seconds. |
| `go_goroutines` | Gauge | Pod / instance labels | Number of active Go goroutines in the controller binary. |

> [!WARNING]
> **Cardinality of `configconnector_reconcile_requests_total` and `configconnector_reconcile_request_duration_seconds`**  
> Both metrics include `namespace` and `group_version_kind` labels by default. Passing the `--resource-name-label=true` flag to the controller manager adds the per-object `name` label (`ResourceNameTag`), which creates unbounded label cardinality proportional to the total number of managed resources across the cluster. Do not enable `--resource-name-label` in production clusters or export per-resource series to managed cloud metrics backends unless you configure metric relabeling rules (`metric_relabel_configs`) to drop the `name` label before ingestion.

---

## GKE Autopilot Compatibility

GKE Autopilot restricts direct scraping of host node kubelet and cAdvisor endpoints (`:10250/metrics/cadvisor`). Dashboards that rely on `container_memory_working_set_bytes` or `container_cpu_usage_seconds_total` often show empty panels in Autopilot environments when scraping solely via in-cluster Prometheus collectors.

To work identically on both GKE Standard and GKE Autopilot without host-level privileges, [`grafana-metrics-dashboard.json`](grafana-metrics-dashboard.json) queries the Go runtime and process collectors exposed directly by the Config Connector binary on `/prom-metrics` (`process_resident_memory_bytes`, `rate(process_cpu_seconds_total[5m])`, and `go_goroutines`).

---

## Step 1: Configure Prometheus Scraping

Configure your Prometheus server, Google Cloud Managed Service for Prometheus (`PodMonitoring`), or Prometheus Operator (`ServiceMonitor`) to scrape port `8888` in the `cnrm-system` namespace.

### Option A: Standard `prometheus.yml` Scrape Config

```yaml
scrape_configs:
  - job_name: 'config-connector-opencensus'
    metrics_path: /metrics
    kubernetes_sd_configs:
      - role: endpoints
        namespaces:
          names:
            - cnrm-system
    relabel_configs:
      - source_labels: [__meta_kubernetes_service_label_cnrm_cloud_google_com_monitored]
        action: keep
        regex: "true"
      - source_labels: [__meta_kubernetes_endpoint_port_name]
        action: keep
        regex: metrics
      - source_labels: [__meta_kubernetes_namespace]
        target_label: kubernetes_namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod

  - job_name: 'config-connector-runtime'
    metrics_path: /prom-metrics
    kubernetes_sd_configs:
      - role: endpoints
        namespaces:
          names:
            - cnrm-system
    relabel_configs:
      - source_labels: [__meta_kubernetes_service_name]
        action: keep
        regex: (cnrm-controller-manager-service|cnrm-manager.*)
      - source_labels: [__meta_kubernetes_endpoint_port_name]
        action: keep
        regex: metrics
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
```

> [!NOTE]
> If you enable `METRICS_VERSION=v2` on the manager StatefulSet to expose `workqueue_depth` on port `8888`, add a third scrape job with `metrics_path: /metrics.v2` targeting `(cnrm-controller-manager-service|cnrm-manager.*)`.

### Option B: Prometheus Operator `ServiceMonitor`

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: config-connector-metrics
  namespace: cnrm-system
spec:
  selector:
    matchLabels:
      cnrm.cloud.google.com/monitored: "true"
  namespaceSelector:
    matchNames:
      - cnrm-system
  endpoints:
    - port: metrics
      path: /metrics
      interval: 30s
    - port: metrics
      path: /prom-metrics
      interval: 30s
```

---

## Step 2: Import the Grafana Dashboard

1. Open your Grafana instance and navigate to **Dashboards** -> **New** -> **Import**.
2. Upload [`docs/guides/grafana-metrics-dashboard.json`](grafana-metrics-dashboard.json) or paste its raw JSON contents into the **Import via panel json** text box, then click **Load**.
3. Select your Prometheus data source in the `Prometheus` dropdown prompt and click **Import**.
4. Use the **Namespace** and **GroupVersionKind** template variables at the top of the dashboard to filter reconciliation rate, latency percentiles (`p95`/`p99`), worker utilization, and applied resource inventory.
