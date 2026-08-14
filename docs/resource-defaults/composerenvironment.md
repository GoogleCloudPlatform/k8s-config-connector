# Cloud Composer Server Defaults and Mutability Drift Analysis

This document analyzes the GCP HTTP responses captured during real GCP recordings of `ComposerEnvironment` (`composer.googleapis.com/v1`). It contrasts historical recordings across Cloud Composer versions with current live API responses to identify server-side default behaviors, default value drift across environment generations, and mutability boundaries.

---

## 1. Architectural Generations: Composer 1 vs. Composer 2 vs. Composer 3

Google Cloud Composer has evolved across three major architectural generations, with default environment creation shifting from **Composer 1** (GKE node pool model) to **Composer 2** (hybrid GKE + serverless workloads) to **Composer 3** (fully managed serverless execution):

### Comparison Matrix Across Generations

| Characteristic | Cloud Composer 1 | Cloud Composer 2 | Cloud Composer 3 | Stability Type |
| :--- | :--- | :--- | :--- | :--- |
| **`imageVersion`** | `"composer-1.*.*-airflow-*.*.*"` | `"composer-2.11.3-airflow-2.10.2"` | `"composer-3-airflow-2.11.1-build.14"` | **Dynamic** (Drifts across releases) |
| **Architecture Model** | Customer-visible GKE nodes & App Engine web server | GKE cluster + serverless Airflow workloads | Managed serverless execution environment | **Dynamic** |
| **`config.nodeCount`** | Defaults to `3` nodes | *Omitted* (Workload-based scaling) | *Omitted* (Workload-based scaling) | **Dynamic** |
| **`config.environmentSize`** | *Omitted* (`0` / Unspecified) | `ENVIRONMENT_SIZE_SMALL` (enum `1`) | `ENVIRONMENT_SIZE_SMALL` (enum `1`) | **Dynamic** |
| **`config.webServerConfig`** | Populated (`composer-n1-webserver-2`) | *Omitted* | *Omitted* | **Dynamic** |
| **`config.gkeCluster`** | Populated (`projects/.../clusters/...`) | Populated (`projects/.../clusters/...`) | *Omitted* (Hidden infrastructure) | **Dynamic** |
| **`config.databaseConfig`** | Populated (`db-n1-standard-2`) | Populated (`db-custom-2-7680`) | *Omitted* (Managed backend) | **Dynamic** |
| **`config.privateEnvironmentConfig`** | Auto-allocates CIDRs | Auto-allocates CIDRs | Populates `networkingConfig: {}` / `networkingType: PUBLIC` | **Dynamic** |
| **`config.nodeConfig.composerInternalIpv4CidrBlock`** | *Omitted* | `"172.31.251.0/24"` | `"100.64.128.0/20"` | **Dynamic** |
| **`config.nodeConfig.network`** | Defaults to `"default"` VPC | Defaults to `"default"` VPC | *Omitted* unless specified | **Dynamic** |
| **`config.nodeConfig.ipAllocationPolicy`** | `{ useIpAliases: true }` | `{ useIpAliases: true }` | `{}` | **Dynamic** |
| **`config.workloadsConfig`** | *Omitted* | Scheduler, WebServer, Worker | Sched, Web, Work, DAGProc, Triggerer | **Dynamic** |
| **`config.dataRetentionConfig`** | *Omitted* | `{ retentionMode: DISABLED }` | `{ retentionDays: 60, retentionMode: ENABLED }` | **Dynamic** |
| **`config.maintenanceWindow`** | Weekly (`FREQ=WEEKLY...`) | Weekly (`FREQ=WEEKLY...`) | *Omitted* (`nil`) | **Dynamic** |

> [!NOTE]
> Server-side defaults for Cloud Composer are **dynamic**: omitting fields like `imageVersion`, `environmentSize`, `nodeCount`, or `workloadsConfig` causes live GCP state to drift based on the target Composer generation. The direct controller supports all three generations seamlessly during creation and acquisition.

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
| **`config.webServerNetworkAccessControl`** | **Static** | Universal open CIDR access across all environments: `[{ value: "0.0.0.0/0" }, { value: "::0/0" }]`. |
| **`config.nodeCount`** | **Dynamic** | Defaults to `3` in Composer 1; omitted in Composer 2 & 3. |
| **`config.environmentSize`** | **Dynamic** | Omitted (`0`) in Composer 1; defaults to `ENVIRONMENT_SIZE_SMALL` (enum `1`) in Composer 2 & 3. |
| **`config.workloadsConfig.scheduler`** | **Dynamic** | Sizing tied to `environmentSize` (e.g. `{ count: 1, cpu: 0.5, memoryGb: 2, storageGb: 1 }` for Small in Composer 2 & 3); omitted in Composer 1. |
| **`config.maintenanceWindow`** | **Dynamic** | Defaulted to weekly schedule in Composer 1 & 2 (`FREQ=WEEKLY;BYDAY=FR,SA,SU`); omitted (`nil`) in Composer 3. |
| **`config.softwareConfig.imageVersion`** | **Dynamic** | Defaults to latest generation-specific Airflow/Composer image. Drifts across releases. |
| **`config.workloadsConfig.dagProcessor`** | **Dynamic** | `{ count: 1, cpu: 1, memoryGb: 4, storageGb: 1 }` (Composer 3); omitted in Composer 1 & 2. |
| **`config.workloadsConfig.triggerer`** | **Dynamic** | `{ count: 1, cpu: 1, memoryGb: 2 }` (Composer 3); omitted in Composer 1 & 2. |
| **`config.workloadsConfig.webServer`** | **Dynamic** | `{ cpu: 0.5, memoryGb: 2, storageGb: 1 }` (Composer 2) or `{ cpu: 1, memoryGb: 4, storageGb: 1 }` (Composer 3); omitted in Composer 1. |
| **`config.workloadsConfig.worker`** | **Dynamic** | `{ min: 1, max: 3, cpu: 0.5, memoryGb: 2, storageGb: 1 }` (Composer 2) or `storageGb: 10` (Composer 3); omitted in Composer 1. |
| **`config.dataRetentionConfig`** | **Dynamic** | `{ retentionMode: DISABLED, storageMode: CLOUD_LOGGING_ONLY }` (Composer 2) or `{ retentionDays: 60, retentionMode: ENABLED }` (Composer 3); omitted in Composer 1. |
| **`config.softwareConfig.schedulerCount`** | **Dynamic** | `1` in Composer 1 & 2; omitted in Composer 3. |
| **`config.databaseConfig.machineType`** | **Dynamic** | Server-allocated DB machine type in Composer 1 & 2; omitted in Composer 3. |
| **`config.databaseConfig.zone`** | **Dynamic** | Server-allocated zone in Composer 1 & 2; omitted in Composer 3. |
| **`config.webServerConfig.machineType`** | **Dynamic** | Server-allocated App Engine instance size in Composer 1; omitted in Composer 2 & 3. |

##### b. Immutable (Create-Only — Must be set at creation time if customized)

| Field Path | Stability Type | Default Value / Behavior (when omitted in `spec`) |
| :--- | :--- | :--- |
| **`storageConfig.bucket` / `bucketRef`** | **Dynamic** | GCP auto-provisions a dedicated GCS bucket (e.g. `"us-central1-test-<hash>-bucket"`). |
| **`config.nodeConfig.composerInternalIpv4CidrBlock`** | **Dynamic** | Auto-allocated subnet (`"172.31.251.0/24"` in Composer 2; `"100.64.128.0/20"` in Composer 3; omitted in Composer 1). |
| **`config.nodeConfig.composerNetworkAttachment`** | **Dynamic** | GCP auto-allocates managed network attachment in Composer 3. |
| **`config.nodeConfig.networkRef` / `subnetworkRef`** | **Dynamic** | Defaults to `"default"` VPC network in Composer 1 & 2; omitted in Composer 3 unless configured. |
| **`config.nodeConfig.ipAllocationPolicy`** | **Dynamic** | `{ useIpAliases: true }` in Composer 1 & 2; `{}` in Composer 3. |
| **`config.nodeConfig.machineType`** | **Dynamic** | Defaults to VM machine type (e.g. `"n1-standard-1"` in Composer 1 & 2); omitted in Composer 3. |
| **`config.nodeConfig.diskSizeGb`** | **Dynamic** | Defaults to 100 GB in Composer 1 & 2; omitted in Composer 3. |
| **`config.softwareConfig.pythonVersion`** | **Dynamic** | Defaults to `"3"` in Composer 1 & 2; omitted in Composer 3. |
| **`config.softwareConfig.webServerPluginsMode`** | **Dynamic** | Defaults to `PLUGINS_ENABLED` (enum `2`) in Composer 2 & 3. |
| **`config.privateEnvironmentConfig.cloudComposerNetworkIpv4CidrBlock`** | **Dynamic** | Auto-allocated tenant CIDR in Composer 1 & 2 (`"172.31.245.0/24"`). |
| **`config.privateEnvironmentConfig.cloudSqlIpv4CidrBlock`** | **Dynamic** | Auto-allocated Cloud SQL CIDR in Composer 1 & 2 (`"10.0.0.0/12"`). |
| **`config.privateEnvironmentConfig.privateClusterConfig`** | **Dynamic** | Auto-configured cluster config in Composer 1 & 2 (`{}`). |
| **`config.privateEnvironmentConfig.webServerIpv4CidrBlock`** | **Dynamic** | Auto-allocated web server CIDR in Composer 1 & 2. |
| **`config.privateEnvironmentConfig.cloudComposerConnectionSubnetwork`** | **Dynamic** | Server-assigned PSC connection subnetwork. |
| **`config.privateEnvironmentConfig.networkingConfig`** | **Dynamic** | Omitted (`nil`) in Composer 1 & 2; defaulted to `{}` in Composer 3. |

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

1. **Two-Phase Default Population Pattern**:
   - **Phase 1: Deterministic Static Defaults (`populateDesiredWithDefaults`)**: Applied to initialize universal static defaults (`webServerNetworkAccessControl`).
   - **Phase 2: Dynamic Server-Computed Field Inheritance (`populateDesiredWithActualIfComputed`)**: Applied during reconciliation and update to copy runtime server-assigned values (`storageConfig.bucket`, `privateEnvironmentConfig` CIDRs/configs/networkingConfig, `nodeConfig` networks/subnets/CIDRs, `softwareConfig` dynamic versions, `databaseConfig`, `nodeCount`, `webServerConfig`, `environmentSize`, `maintenanceWindow`, and dynamic `workloadsConfig`/`dataRetentionConfig` sizing) from `actualPb` into `desiredPb` whenever the user omits them in `desired.Spec`.

2. **Patch Isolation**:
   Cloud Composer does not accept multi-field update masks. Each mutable change must be issued as an isolated patch request with a single-field `updateMask` via `fieldUpdaters`.

3. **Validation Isolation**:
   `validateUpdatableFields` validates diffs between `desiredPb` (with defaults and server-assigned values populated) and `maskedActual`, ensuring immutable server-computed fields are never flagged as illegal updates across Composer 1, Composer 2, and Composer 3.

### Pros and Cons Analysis

The **Two-Phase Default Population Pattern** (Universal Static Defaults + Dynamic Server-Computed Field Inheritance) is evaluated below against alternative design strategies:

#### Comparison of Design Approaches

| Feature / Scenario | Approach 1: Static Defaults for All Fields | Approach 2: Pure KRM Spec-Only Diffing | Approach 3: Two-Phase Default Population (Selected) |
| :--- | :--- | :--- | :--- |
| **Multi-Generation Support (Composer 1, 2, 3)** | ❌ **Fails**: Hardcoding Composer 2 defaults (e.g., `environmentSize: SMALL`, `workloadsConfig`) breaks Composer 1 and Composer 3. | ⚠️ **Partial**: Does not account for proto-level mapping and output-only vs. spec boundaries. | ✅ **Excellent**: Dynamic inheritance adapts to live architecture without spurious diffs. |
| **Pre-existing Resource Acquisition** | ❌ **Fails**: Minimal YAML acquisition triggers `updating field(s) ... is not supported` for omitted fields. | ⚠️ **Fragile**: Requires users to write exhaustive YAML matching 100% of live state. | ✅ **Seamless**: Minimal KRM YAML acquires any valid environment generation cleanly. |
| **Declarative Enforcement for Unset Fields** | ✅ **Strict**: Forces all omitted fields to hardcoded KCC default values. | ❌ **None**: Unset fields are unmanaged. | ⚠️ **By Design**: Unset fields inherit live server defaults. Users must declare explicit spec fields to enforce strict values. |

---

#### Pros of the Two-Phase Pattern

1. **Generation-Resilient Architecture**:
   - Seamlessly handles Composer 1 (GKE node pool model), Composer 2 (workloads model), and Composer 3 (managed serverless execution) without requiring generation-specific CRD splits.
2. **Robust Resource Acquisition & Adoption**:
   - Users can acquire pre-existing GCP Composer environments with minimal YAML without false-positive mutability errors.
3. **Eliminates Spurious Diffing & Update Loops**:
   - Dynamic fields assigned or generated by GCP are recognized and preserved rather than flagged as unsupported mutations.
4. **Isolated & Safe Single-Mask Patching**:
   - Retains strict validation via `fieldUpdaters` and `validateUpdatableFields`: when a user explicitly declares a mutable field in `spec`, KCC detects the diff, validates that the field is updatable, and issues an isolated single-mask patch request.

---

#### Cons & Trade-offs

1. **Permissive Out-of-Band Drift for Omitted Fields**:
   - If a mutable field is omitted in `spec` and subsequently modified out-of-band (e.g. via GCP Console or `gcloud`), KCC will inherit the new actual value during reconciliation rather than reverting it back to a baseline default.
   - *Mitigation*: Users seeking strict drift correction must explicitly define the target value in `spec`.
2. **Controller Maintenance Overhead**:
   - `populateDesiredWithActualIfComputed` requires explicit handling for each dynamic field hierarchy. When GCP introduces new environment sub-structures, they must be registered in the dynamic inheritance mapper.
