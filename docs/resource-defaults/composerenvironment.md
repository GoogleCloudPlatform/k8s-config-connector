# Cloud Composer Server Defaults and Mutability Drift Analysis

This document analyzes the GCP HTTP responses captured during real GCP recording of `ComposerEnvironment` (`composer.googleapis.com/v1`). It contrasts historical recordings with current live API responses to identify server-side default behaviors, default value drift over time, and mutability boundaries.

---

## 1. Major Discovery: Default Version Upgrade

When creating a minimal `ComposerEnvironment` without specifying `config.softwareConfig.imageVersion` or `config.environmentSize`, Google Cloud Composer has shifted its default environment generation from **Cloud Composer 2** to **Cloud Composer 3**.

### Comparison: Historical Recording vs. Live GCP API

| Characteristic | Historical Recording (`_http_previous.log`) | Live GCP API (`_http.log`) | Stability Type |
| :--- | :--- | :--- | :--- |
| **`imageVersion`** | `"composer-2.11.3-airflow-2.10.2"` | `"composer-3-airflow-2.11.1-build.14"` | **Dynamic** (Drifts across releases) |
| **Architecture Model** | Customer-visible GKE cluster & Cloud SQL | Managed serverless execution environment | **Dynamic** |
| **`config.gkeCluster`** | Populated (`projects/.../clusters/...`) | *Omitted* (Hidden infrastructure) | **Dynamic** |
| **`config.databaseConfig`** | Populated (`db-custom-2-7680`) | *Omitted* (Managed metadata backend) | **Dynamic** |
| **Worker Storage (`storageGb`)** | `1` GB | `10` GB | **Dynamic** |
| **Web Server Memory (`memoryGb`)** | `2` GB | `4` GB | **Dynamic** |
| **`dagProcessor` & `triggerer`** | *Not present by default* | `{ count: 1, cpu: 1, ... }` | **Dynamic** |
| **`dataRetentionConfig`** | `{ retentionMode: DISABLED }` | `{ retentionDays: 60, retentionMode: ENABLED }` | **Dynamic** |

> [!NOTE]
> This version shift proves that server-side defaults for Cloud Composer are **dynamic**: omitting fields like `imageVersion` or `workloadsConfig` causes live GCP state to drift as Google rolls out new Composer baselines.

---

## 2. Detailed Breakdown of Defaulted and Generated Fields

Fields returned by the GCP API fall into the following hierarchy based on their configurability, server default values, and mutability:

---

### A. Configurable Fields

#### (1) Has Non-Zero Server-Generated Value
When omitted in `spec`, GCP assigns a specific non-zero / non-empty server default:

##### a. Mutable (In-Place Updatable via `fieldUpdaters`)

| Field Path | Stability Type | Default Value / Behavior (when omitted in `spec`) |
| :--- | :--- | :--- |
| **`config.environmentSize`** | **Static** | Defaults to `ENVIRONMENT_SIZE_SMALL` (enum value `1`). |
| **`config.workloadsConfig.scheduler`** | **Static** | `{ count: 1, cpu: 0.5, memoryGb: 2, storageGb: 1 }`. |
| **`config.webServerNetworkAccessControl`** | **Static** | Open CIDR access: `[{ value: "0.0.0.0/0" }, { value: "::0/0" }]`. |
| **`config.maintenanceWindow`** | **Static** | Weekly recurring maintenance schedule (`FREQ=WEEKLY;BYDAY=FR,SA,SU`). |
| **`config.softwareConfig.imageVersion`** | **Dynamic** | Defaults to latest Airflow/Composer image (e.g. `"composer-3-airflow-2.11.1-build.14"`). Drifts across releases. |
| **`config.workloadsConfig.dagProcessor`** | **Dynamic** | `{ count: 1, cpu: 1, memoryGb: 4, storageGb: 1 }` (Composer 3). |
| **`config.workloadsConfig.triggerer`** | **Dynamic** | `{ count: 1, cpu: 1, memoryGb: 2 }` (Composer 3). |
| **`config.workloadsConfig.webServer`** | **Dynamic** | `{ cpu: 1, memoryGb: 4, storageGb: 1 }` (Composer 3). |
| **`config.workloadsConfig.worker`** | **Dynamic** | `{ min: 1, max: 3, cpu: 0.5, memoryGb: 2, storageGb: 10 }` (Composer 3). |
| **`config.dataRetentionConfig`** | **Dynamic** | `{ airflowMetadataRetentionConfig: { retentionDays: 60, retentionMode: RETENTION_MODE_ENABLED } }` (Composer 3). |
| **`config.softwareConfig.schedulerCount`** | **Dynamic** | `1` in Composer 1 & 2; omitted in Composer 3. |
| **`config.databaseConfig.machineType`** | **Dynamic** | `"db-custom-2-7680"` in Composer 1 & 2; omitted in Composer 3. |

##### b. Immutable (Create-Only — Must be set at creation time if customized)

| Field Path | Stability Type | Default Value / Behavior (when omitted in `spec`) |
| :--- | :--- | :--- |
| **`storageConfig.bucket` / `bucketRef`** | **Dynamic** | GCP auto-provisions a dedicated GCS bucket (e.g. `"us-central1-composerenviron-<hash>-bucket"`). |
| **`config.nodeConfig.composerInternalIpv4CidrBlock`** | **Dynamic** | GCP auto-allocates an internal management subnet (e.g. `"100.64.128.0/20"`). |
| **`config.nodeConfig.composerNetworkAttachment`** | **Dynamic** | GCP auto-allocates managed network attachment in Composer 3. |
| **`config.nodeConfig.networkRef` / `subnetworkRef`** | **Dynamic** | Defaults to the project's `"default"` VPC network in Composer 1 & 2; omitted in Composer 3 unless configured. |
| **`config.nodeConfig.machineType`** | **Dynamic** | Defaults to standard VM machine type (e.g. `"n1-standard-1"` in Composer 1 & 2); omitted in Composer 3. |
| **`config.nodeConfig.diskSizeGb`** | **Dynamic** | Defaults to 100 GB in Composer 1 & 2; omitted in Composer 3. |
| **`config.softwareConfig.pythonVersion`** | **Dynamic** | Defaults to `"3"` in Composer 1 & 2; omitted in Composer 3. |
| **`config.softwareConfig.webServerPluginsMode`** | **Dynamic** | Defaults to `PLUGINS_ENABLED` (enum value `2`) in Composer 1 & 2. |
| **`config.privateEnvironmentConfig.networkingConfig`** | **Static** | Defaults to standard private PSC networking configuration. |

---

#### (2) Defaulted to the Type's Zero Value
When omitted in `spec`, these fields remain empty, nil, or false:

| Field Path | Zero Value Type | Default Behavior |
| :--- | :--- | :--- |
| **`labels`** | Map | Empty map `{}`. |
| **`config.softwareConfig.airflowConfigOverrides`** | Map | Empty map `{}` (`airflow.cfg` defaults). |
| **`config.softwareConfig.pypiPackages`** | Map | Empty map `{}` (standard Airflow bundle). |
| **`config.softwareConfig.envVariables`** | Map | Empty map `{}`. |
| **`config.softwareConfig.cloudDataLineageIntegration`** | Object | Disabled `{}`. |
| **`config.encryptionConfig`** | Object | Empty `{}` (Google-managed default encryption). |
| **`config.recoveryConfig.scheduledSnapshotsConfig`** | Object | Disabled `{ enabled: false }` or empty `{}`. |
| **`config.masterAuthorizedNetworksConfig`** | Object | Disabled `{}`. |
| **`config.resilienceMode`** | Enum | Standard resilience mode (enum `0` / empty). |
| **`config.nodeConfig.oauthScopes`** | List | Empty list `[]`. |
| **`config.nodeConfig.tags`** | List | Empty list `[]`. |
| **`config.nodeConfig.enableIpMasqAgent`** | Boolean | `false`. |
| **`config.nodeConfig.bootDiskType`** | String | Empty string `""`. |

---

### B. Output-Only / Status Fields
These fields are runtime metadata computed and assigned exclusively by GCP. They are mapped into `status.observedState` and **must never be compared against `spec`**:

| Field Path | Type / Format | Description |
| :--- | :--- | :--- |
| **`name`** | String (URI) | Full resource URI (`projects/{project}/locations/{location}/environments/{id}`). |
| **`uuid`** | String (UUID) | Server-generated unique identifier per environment instance. |
| **`state`** | Enum | Lifecycle state enum (`CREATING` $\rightarrow$ `RUNNING` $\rightarrow$ `UPDATING`). |
| **`createTime`** / **`updateTime`** | Timestamp | UTC timestamps of resource creation and latest modification. |
| **`config.airflowUri`** | String (URL) | Generated URL for the Airflow Web UI. |
| **`config.airflowByoidUri`** | String (URL) | Generated URL for Bring Your Own Identity (BYOID) Web UI. |
| **`config.dagGcsPrefix`** | String (GCS URI) | Generated GCS storage path for Airflow DAGs (e.g. `gs://<bucket>/dags`). |
| **`config.gkeCluster`** | String (Resource URI) | Reference to the underlying GKE cluster (Composer 1 & 2 only). |

---

## 3. Reconciliation & Diffing Guidelines for KCC Direct Controller

1. **Declarative Intent for Omitted Fields**:
   When a configurable field is omitted in `desired.Spec`, the declared intent is that the resource should adopt its default state. The controller must understand whether the actual GCP state currently matches the intended default value:
   - If `actual` is already at the default value $\rightarrow$ no-op (avoid redundant patch calls).
   - If `actual` has drifted away from the default value (e.g. previously modified or overridden) $\rightarrow$ the controller should reconcile and patch the field back to its default value where supported.

2. **Patch Isolation**:
   Cloud Composer does not accept multi-field update masks. Each mutable change must be issued as an isolated patch request with a single-field `updateMask` via `fieldUpdaters`.

3. **Validation Isolation**:
   `validateUpdatableFields` must only validate diffs on fields explicitly specified in `desired.Spec`, ignoring output-only fields and omitted server defaults.
