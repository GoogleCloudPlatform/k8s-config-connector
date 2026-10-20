<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Cloud Monitoring Dashboard Setup Guide for KCC

As a Config Connector user, you likely prefer Infrastructure as Code (IaC) to manual UI configuration. This guide details how to route KCC Prometheus metrics from your GKE cluster directly into Google Cloud Monitoring, and deploy a pre-configured dashboard using a Config Connector `MonitoringDashboard` resource.

## Prerequisites
*   A GKE cluster with Config Connector installed.
*   Permission to manage GKE addons and Workload Identity.
*   Permission to create dashboards in Google Cloud Monitoring.

## Step 1: Enable Google Cloud Managed Service for Prometheus (GMP)

Google Cloud Managed Service for Prometheus (GMP) is the native bridge between in-cluster Prometheus metrics and Google Cloud Monitoring.

If you are using GKE Autopilot, managed collection is enabled by default. If using GKE Standard, enable it on your cluster:

```bash
gcloud container clusters update CLUSTER_NAME \
    --region REGION \
    --enable-managed-prometheus
```

## Step 2: Configure the PodMonitoring Resource

GMP uses a custom resource called `PodMonitoring` to discover and scrape metrics from KCC controller pods.

A pre-configured manifest is provided in [`kcc-podmonitoring.yaml`](./kcc-podmonitoring.yaml), which targets KCC controller pods in `cnrm-system` across internal metrics ports `8888` and `48797`.

Apply the configuration to your cluster:

```bash
kubectl apply -f docs/monitoring/kcc-podmonitoring.yaml
```

Within a few minutes, KCC metrics are ingested into Google Cloud Monitoring under the `prometheus.googleapis.com` metric prefix.

> **Warning (Cost and Cardinality Considerations):**
> Metrics such as `reconcile_requests_total` and `reconcile_request_duration_seconds` carry labels including `group_version_kind`, `namespace`, and `status`. In environments with many managed resources and high reconciliation frequency, these metrics may have unbounded cardinality and can lead to increased Cloud Monitoring metric ingestion costs. If you need to manage ingestion volume, consider using `metricRelabeling` with `action: keep` or `action: drop` in your `PodMonitoring` resource to scrape only the specific metrics or namespaces required.

## Step 3: Deploy the Dashboard

To visualize the metrics, we provide a pre-configured KCC dashboard resource in [`kcc-dashboard.yaml`](./kcc-dashboard.yaml). This dashboard includes panels for:
*   **Total Managed Resources** (view overall resource footprint and growth).
*   **Reconciliation Errors by Kind** (identify failing resources by Kind).
*   **Reconciliation Latency** (monitor P99 reconciliation duration).
*   **System Health Metrics** (track KCC controller Working Set Memory, CPU Core Usage, and Container Restarts in `cnrm-system`).

Before applying, replace `${PROJECT_ID?}` and `${NAMESPACE?}` in [`kcc-dashboard.yaml`](./kcc-dashboard.yaml) with your Google Cloud Project ID and a Config Connector managed namespace. Then, apply it to your cluster:

```bash
kubectl apply -f docs/monitoring/kcc-dashboard.yaml
```

## Step 4: Access the Dashboard
1. Open the [Google Cloud Console](https://console.cloud.google.com).
2. Navigate to **Monitoring** -> **Dashboards**.
3. Locate and click on **Config Connector (KCC) Metrics** to view your live data.
