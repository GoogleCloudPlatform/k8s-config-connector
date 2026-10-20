# Config Connector Practical Complex Examples

This guide provides practical, production-grade, and heavily commented YAML manifests for Config Connector (KCC). These examples demonstrate **advanced configurations**, **optional field interplay**, **inter-resource dependencies**, and **enterprise operational patterns** across major Google Cloud services.

---

## Overview and General Conventions

When adapting these examples for your clusters and environments:
* **Project Scoping**: Replace `${PROJECT_ID?}` with your Google Cloud Project ID, or configure the `cnrm.cloud.google.com/project-id` annotation on your namespace.
* **Region & Zone**: Adjust regions (`us-central1`, `us-east4`) and zones to match your target topology.
* **Resource References**: References use the standard `*Ref` contract (`name` for intra-namespace, `name` + `namespace` for cross-namespace, or `external` for unmanaged resources).
* **GitOps Hygiene**: The examples include `cnrm.cloud.google.com/state-into-spec: absent` to prevent diff conflicts with GitOps continuous delivery tools like ArgoCD and Flux.

---

## 1. Cloud SQL

### 1.1 Enterprise Regional High-Availability PostgreSQL with Private IP & Hardened Security

This example configures a production-grade Cloud SQL instance with:
* **High Availability**: Regional availability (`REGIONAL`) with automatic failover.
* **Private Networking**: Integrated with a VPC via Private Service Access (Service Networking).
* **Point-in-Time Recovery (PITR)**: Automated transaction log retention and WAL archiving.
* **Maintenance Management**: Configured maintenance window with update deny periods for peak business seasons.
* **Security Hardening**: Database flags enforcing SSL and auditing, password validation policies, and disk autoresize.

```yaml
# ------------------------------------------------------------------------------
# 1. Cloud SQL High-Availability PostgreSQL Instance
# ------------------------------------------------------------------------------
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: pg-prod-instance
  annotations:
    # Protect mission-critical database from accidental deletion if KRM object is deleted
    cnrm.cloud.google.com/deletion-policy: abandon
    # Maintain clean GitOps status by preventing GCP defaults from writing back into spec
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  databaseVersion: POSTGRES_15
  region: us-central1
  settings:
    # Dedicated machine type (e.g. 4 vCPUs, 16 GB RAM)
    tier: db-custom-4-16384
    # Multi-zone regional failover
    availabilityType: REGIONAL

    # Storage configuration with auto-expansion enabled
    diskType: PD_SSD
    diskSize: 100
    diskAutoresize: true
    diskAutoresizeLimit: 1000

    # Private IP configuration via Private Services Access (requires ServiceNetworking connection)
    ipConfiguration:
      ipv4Enabled: false
      privateNetworkRef:
        name: production-vpc
      # Require encrypted connections
      requireSsl: true

    # Comprehensive backup and Point-in-Time Recovery (PITR)
    backupConfiguration:
      enabled: true
      startTime: "02:00"
      pointInTimeRecoveryEnabled: true
      backupRetentionSettings:
        retainedBackups: 30
        retentionUnit: COUNT
      transactionLogRetentionDays: 7

    # Maintenance windows and deny periods
    maintenanceWindow:
      day: 7 # Sunday
      hour: 4 # 04:00 UTC
      updateTrack: stable
    denyMaintenancePeriod:
      startDate: "2026-11-20"
      endDate: "2026-12-05"
      time: "00:00:00"

    # Database flags for security and performance
    databaseFlags:
      - name: log_connections
        value: "on"
      - name: log_disconnections
        value: "on"
      - name: log_min_duration_statement
        value: "1000" # Log slow queries > 1s
      - name: autovacuum
        value: "on"

    # Password validation policy
    passwordValidationPolicy:
      minComplexity: COMPLEXITY_DEFAULT
      minLength: 16
      disallowUsernameSubstring: true
      enablePasswordPolicy: true

    # Active connection limits
    insightsConfig:
      queryInsightsEnabled: true
      queryStringLength: 1024
      recordApplicationTags: true
      recordClientAddress: false
---
# ------------------------------------------------------------------------------
# 2. Database Definition
# ------------------------------------------------------------------------------
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLDatabase
metadata:
  name: app-orders-db
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  charset: UTF8
  collation: en_US.UTF8
  # Intra-namespace reference to parent SQLInstance
  instanceRef:
    name: pg-prod-instance
---
# ------------------------------------------------------------------------------
# 3. Application User with Secret-Sourced Password
# ------------------------------------------------------------------------------
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLUser
metadata:
  name: app-backend-user
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  instanceRef:
    name: pg-prod-instance
  # Credentials sourced securely from a Kubernetes Secret
  password:
    valueFrom:
      secretKeyRef:
        name: db-user-secrets
        key: app-password
  # Cloud SQL password policy enforcement for this specific user
  passwordPolicy:
    allowedFailedAttempts: 5
    enableFailedAttemptsCheck: true
    enablePasswordVerification: true
```

---

### 1.2 Cross-Region Advanced Disaster Recovery (DR) Pair

Cloud SQL cross-region disaster recovery requires establishing a Primary instance and a cross-region Read/Failover Replica. This scenario leverages `cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"` to:
1. Allow simultaneous declarative creation of both instances without circular dependency deadlocks.
2. Prevent destructive role reversions when an out-of-band switchover or emergency failover is initiated in GCP.

```yaml
# ------------------------------------------------------------------------------
# Primary SQL Instance in us-central1
# ------------------------------------------------------------------------------
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: sql-dr-primary
  annotations:
    # Enable Advanced Disaster Recovery drift suppression & pair setup
    cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"
    cnrm.cloud.google.com/deletion-policy: abandon
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  databaseVersion: MYSQL_8_0
  region: us-central1
  settings:
    tier: db-custom-4-16384
    availabilityType: REGIONAL
    backupConfiguration:
      enabled: true
      binaryLogEnabled: true # Mandatory for MySQL cross-region replication
---
# ------------------------------------------------------------------------------
# Disaster Recovery (DR) Replica in us-east4
# ------------------------------------------------------------------------------
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: sql-dr-replica
  annotations:
    # Enable Advanced Disaster Recovery drift suppression & pair setup
    cnrm.cloud.google.com/sqlinstance-advanced-dr: "enabled"
    cnrm.cloud.google.com/deletion-policy: abandon
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  databaseVersion: MYSQL_8_0
  region: us-east4
  # Points to the primary instance as the replication source
  masterInstanceRef:
    name: sql-dr-primary
  settings:
    tier: db-custom-4-16384
    availabilityType: REGIONAL
```

---

## 2. Google Kubernetes Engine (GKE)

### 2.1 Production Private Cluster with Workload Identity & CMEK

This example demonstrates a hardened GKE cluster decoupled from its node pools:
* **Node Pool Decoupling**: Uses `cnrm.cloud.google.com/remove-default-node-pool: "true"` so the cluster is created without an unmanaged default node pool. Nodes are managed separately via `ContainerNodePool`.
* **Private Cluster Configuration**: Public endpoint disabled, control plane accessible only via authorized networks or internal VPC routes.
* **Workload Identity**: Fully enabled for secure, keyless authentication to Google Cloud APIs.
* **DatapathProvider (GKE Datapath v2)**: Leverages eBPF for network routing and network policy enforcement.
* **Database Encryption (CMEK)**: Encrypts Kubernetes secrets in etcd using a Cloud KMS key.

```yaml
# ------------------------------------------------------------------------------
# 1. Production GKE Cluster Control Plane
# ------------------------------------------------------------------------------
apiVersion: container.cnrm.cloud.google.com/v1beta1
kind: ContainerCluster
metadata:
  name: prod-cluster-01
  annotations:
    # Instructs KCC to remove the initial default node pool immediately after cluster creation
    cnrm.cloud.google.com/remove-default-node-pool: "true"
    cnrm.cloud.google.com/deletion-policy: abandon
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  location: us-central1
  # Release Channel for automatic control plane upgrades
  releaseChannel:
    channel: REGULAR

  # VPC-Native Networking (IP Aliasing)
  networkingMode: VPC_NATIVE
  networkRef:
    name: production-vpc
  subnetworkRef:
    name: gke-subnet
  ipAllocationPolicy:
    clusterSecondaryRangeName: pods
    servicesSecondaryRangeName: services

  # Network Datapath v2 (eBPF based network policy)
  networkPolicy:
    enabled: true
    provider: CALICO

  # Hardened Private Cluster Settings
  privateClusterConfig:
    enablePrivateNodes: true
    enablePrivateEndpoint: false
    masterIpv4CidrBlock: 172.16.0.0/28
  masterAuthorizedNetworksConfig:
    cidrBlocks:
      - cidrBlock: 10.100.0.0/16
        displayName: Bastion and Internal Corporate Net

  # Security & Workload Identity
  workloadIdentityConfig:
    workloadPool: ${PROJECT_ID?}.svc.id.goog

  # CMEK Encryption for etcd Kubernetes Secrets
  databaseEncryption:
    state: ENCRYPTED
    keyNameRef:
      name: gke-etcd-crypto-key

  # Maintenance Window Policy
  maintenancePolicy:
    recurringWindow:
      recurrence: "FREQ=WEEKLY;BYDAY=SA,SU"
      window:
        startTime: "2026-01-01T00:00:00Z"
        endTime: "2026-01-01T08:00:00Z"
---
# ------------------------------------------------------------------------------
# 2. Independent Application Node Pool with Spot VMs & Taints
# ------------------------------------------------------------------------------
apiVersion: container.cnrm.cloud.google.com/v1beta1
kind: ContainerNodePool
metadata:
  name: app-workload-pool
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  location: us-central1
  clusterRef:
    name: prod-cluster-01

  # Autoscaling configuration per zone
  autoscaling:
    minNodeCount: 2
    maxNodeCount: 10

  # Node upgrade rollout strategy
  upgradeSettings:
    maxSurge: 1
    maxUnavailable: 0
    strategy: SURGE

  # Node configuration
  nodeConfig:
    machineType: e2-standard-4
    diskType: pd-ssd
    diskSizeGb: 100
    imageType: COS_CONTAINERD

    # Custom dedicated Google Service Account for nodes (avoids default Compute engine SA)
    serviceAccountRef:
      name: gke-node-sa

    # Cost-optimization with Spot instances
    spot: true

    # Hardening: Shielded Nodes
    shieldedInstanceConfig:
      enableSecureBoot: true
      enableIntegrityMonitoring: true

    # Workload Identity metadata configuration on node
    workloadMetadataConfig:
      mode: GKE_METADATA

    # Node Labels
    labels:
      workload-tier: application
      lifecycle: spot

    # Node Taints to isolate general pods from this specialized pool
    taint:
      - key: dedicated
        value: application
        effect: NO_SCHEDULE
```

---

## 3. Cloud Storage

### 3.1 Enterprise Hardened Bucket with CMEK, Autoclass, Retention & Lifecycle Rules

This configuration demonstrates complex field interplay:
* **Autoclass & Lifecycle Policies**: Autoclass automatically transitions objects across access tiers (Standard -> Nearline -> Archive). Custom lifecycle rules clean up incomplete multipart uploads and purge noncurrent versions.
* **Customer-Managed Encryption Keys (CMEK)**: References a Cloud KMS key.
* **Object Retention / Bucket Lock**: Enforces WORM (Write Once, Read Many) compliance with locked retention duration.
* **Soft Delete Policy**: Protects against accidental deletion with a 14-day recovery window.
* **Dual-Region Placement**: Replicates data across paired regions (`US-CENTRAL1` and `US-EAST4`).

```yaml
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  # Bucket names must be globally unique
  name: corp-secure-archive-${PROJECT_ID?}
  annotations:
    # Safety guard: Prevent accidental deletion of non-empty buckets
    cnrm.cloud.google.com/force-destroy: "false"
    # Unlink KRM object upon deletion without destroying physical bucket and data
    cnrm.cloud.google.com/deletion-policy: abandon
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  # Custom Dual-Region Placement for multi-region resilience
  customPlacementConfig:
    dataLocations:
      - US-CENTRAL1
      - US-EAST4

  # Enforce IAM-only access (no object-level ACLs)
  uniformBucketLevelAccess: true

  # Customer-Managed Encryption Key (CMEK)
  encryption:
    kmsKeyNameRef:
      name: storage-crypto-key

  # Intelligent lifecycle cost optimization
  autoclass:
    enabled: true
    terminalStorageClass: ARCHIVE

  # Versioning enabled (required for versioned lifecycle conditions)
  versioning:
    enabled: true

  # Soft delete policy (14-day retention recovery window)
  softDeletePolicy:
    retentionDurationSeconds: 1209600

  # Object Retention Policy (Bucket Lock for Compliance)
  retentionPolicy:
    retentionPeriod: 7776000 # 90 days in seconds
    isLocked: true

  # Multi-Action Lifecycle Rules
  lifecycleRule:
    # Rule 1: Abort incomplete multipart uploads after 3 days to avoid stranded storage costs
    - action:
        type: AbortIncompleteMultipartUpload
      condition:
        age: 3
    # Rule 2: Expire and delete noncurrent versions after 60 days
    - action:
        type: Delete
      condition:
        numNewerVersions: 3
        daysSinceNoncurrentTime: 60
        withState: ARCHIVED

  # Cross-Origin Resource Sharing (CORS)
  cors:
    - origin: ["https://console.example.com"]
      method: ["GET", "HEAD"]
      responseHeader: ["Content-Type", "ETag"]
      maxAgeSeconds: 3600
```

---

## 4. Pub/Sub

### 4.1 Resilient Ingestion Pipeline with Schema Validation & Dead-Letter Queue (DLQ)

This architecture builds an enterprise Pub/Sub ingestion pipeline:
* **PubSubSchema**: Defines an Avro schema enforcing structured message payloads.
* **PubSubTopic**: Validates ingested messages against the schema and encrypts messages using CMEK.
* **Dead-Letter Queue (DLQ)**: Failed deliveries are diverted to a designated dead-letter topic after 5 unsuccessful delivery attempts.
* **PubSubSubscription**: Implements exponential backoff retry policies and exactly-once delivery.

```yaml
# ------------------------------------------------------------------------------
# 1. Avro Schema Definition
# ------------------------------------------------------------------------------
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubSchema
metadata:
  name: order-event-schema
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  type: AVRO
  definition: |
    {
      "type": "record",
      "name": "OrderEvent",
      "fields": [
        {"name": "orderId", "type": "string"},
        {"name": "userId", "type": "string"},
        {"name": "amount", "type": "double"},
        {"name": "timestamp", "type": "long"}
      ]
    }
---
# ------------------------------------------------------------------------------
# 2. Dead-Letter Topic (DLQ)
# ------------------------------------------------------------------------------
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: order-events-dlq
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  messageRetentionDuration: 1209600s # Retain failed messages for 14 days
---
# ------------------------------------------------------------------------------
# 3. Ingestion Topic with Schema Validation & CMEK
# ------------------------------------------------------------------------------
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: order-events-topic
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  # Enforce Avro message payload validation
  schemaSettings:
    schemaRef:
      name: order-event-schema
    encoding: JSON
  # Encryption via Cloud KMS
  kmsKeyRef:
    name: pubsub-crypto-key
---
# ------------------------------------------------------------------------------
# 4. Processing Subscription with Dead-Letter Policy & Retry Exponential Backoff
# ------------------------------------------------------------------------------
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubSubscription
metadata:
  name: order-processor-sub
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  topicRef:
    name: order-events-topic
  ackDeadlineSeconds: 60
  enableExactlyOnceDelivery: true
  retainAckedMessages: false
  messageRetentionDuration: 604800s # 7 days

  # Dead-Letter Policy diverts unprocessable messages after 5 failed deliveries
  deadLetterPolicy:
    deadLetterTopicRef:
      name: order-events-dlq
    maxDeliveryAttempts: 5

  # Exponential backoff retry policy for transient subscriber failures
  retryPolicy:
    minimumBackoff: 10s
    maximumBackoff: 600s

  # Optional message filtering (processes only events with matching attributes)
  filter: 'attributes.environment = "production"'
```

---

## 5. Compute Engine & Networking

### 5.1 Hardened Compute Instance in Custom VPC with Shielded VM & CMEK Disks

Demonstrates provisioning compute workloads securely:
* **Custom Networking**: Interlinks `ComputeNetwork`, `ComputeSubnetwork`, and `ComputeFirewall`.
* **Shielded VM**: Enforces Secure Boot, vTPM, and integrity monitoring.
* **CMEK Encryption**: Boot disk and auxiliary data disks encrypted using customer KMS keys.
* **Identity**: Attaches a custom dedicated Google Service Account with minimal OAuth scopes.

```yaml
# ------------------------------------------------------------------------------
# 1. Custom VPC Network
# ------------------------------------------------------------------------------
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: enterprise-vpc
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  autoCreateSubnetworks: false
  routingConfig:
    routingMode: REGIONAL
---
# ------------------------------------------------------------------------------
# 2. Subnetwork with Private Google Access
# ------------------------------------------------------------------------------
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeSubnetwork
metadata:
  name: backend-subnet-uscentral1
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  region: us-central1
  networkRef:
    name: enterprise-vpc
  ipCidrRange: 10.20.0.0/24
  privateIpGoogleAccess: true
---
# ------------------------------------------------------------------------------
# 3. Ingress Firewall Policy for Application Tag
# ------------------------------------------------------------------------------
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeFirewall
metadata:
  name: allow-internal-backend
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  networkRef:
    name: enterprise-vpc
  direction: INGRESS
  priority: 1000
  targetTags:
    - backend-server
  sourceRanges:
    - 10.20.0.0/16
  allow:
    - protocol: tcp
      ports:
        - "8080"
        - "8443"
---
# ------------------------------------------------------------------------------
# 4. Hardened Compute Instance
# ------------------------------------------------------------------------------
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeInstance
metadata:
  name: app-backend-vm-01
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  zone: us-central1-a
  machineType: e2-standard-4

  # Network attachment to custom subnetwork (no external public IP)
  networkInterface:
    - subnetworkRef:
        name: backend-subnet-uscentral1
      # Omitting accessConfig ensures the VM receives NO external public IP address

  # Boot Disk with CMEK
  bootDisk:
    initializeParams:
      size: 50
      type: pd-ssd
      sourceImageRef:
        external: debian-cloud/debian-12
    kmsKeyRef:
      name: compute-disk-crypto-key

  # Shielded VM Configuration
  shieldedInstanceConfig:
    enableSecureBoot: true
    enableVtpm: true
    enableIntegrityMonitoring: true

  # Tags for firewall targeting
  tags:
    - backend-server

  # Dedicated Service Account
  serviceAccount:
    serviceAccountRef:
      name: backend-vm-sa
    scopes:
      - https://www.googleapis.com/auth/cloud-platform

  # Startup Script and Metadata
  metadata:
    - key: enable-oslogin
      value: "TRUE"
    - key: startup-script
      value: |
        #!/bin/bash
        echo "Starting application initialization..."
```

---

## 6. IAM & Security Hardening

### 6.1 Workload Identity Integration via `IAMPartialPolicy`

Using `IAMPolicy` can inadvertently overwrite pre-existing cloud bindings managed by Google Cloud or other teams. In contrast, `IAMPartialPolicy` manages only specific declarative bindings without clobbering out-of-band roles:

```yaml
# ------------------------------------------------------------------------------
# 1. Google IAM Service Account
# ------------------------------------------------------------------------------
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  name: payment-service-gsa
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  displayName: "Payment Service Workload Identity GSA"
---
# ------------------------------------------------------------------------------
# 2. Grant Workload Identity User Role to Kubernetes Service Account
# ------------------------------------------------------------------------------
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPartialPolicy
metadata:
  name: payment-gsa-workload-identity
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  resourceRef:
    apiVersion: iam.cnrm.cloud.google.com/v1beta1
    kind: IAMServiceAccount
    name: payment-service-gsa
  bindings:
    - role: roles/iam.workloadIdentityUser
      members:
        # Binds Kubernetes Service Account 'payment-ksa' in namespace 'payments'
        - member: serviceAccount:${PROJECT_ID?}.svc.id.goog[payments/payment-ksa]
---
# ------------------------------------------------------------------------------
# 3. Grant Secret Accessor to the GSA at the Project Level
# ------------------------------------------------------------------------------
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPartialPolicy
metadata:
  name: payment-gsa-secret-accessor
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  resourceRef:
    apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
    kind: Project
    external: ${PROJECT_ID?}
  bindings:
    - role: roles/secretmanager.secretAccessor
      members:
        - member: serviceAccount:payment-service-gsa@${PROJECT_ID?}.iam.gserviceaccount.com
```

---

### 6.2 CMEK-Encrypted Secret Manager Secret with Automatic Rotation

```yaml
apiVersion: secretmanager.cnrm.cloud.google.com/v1beta1
kind: SecretManagerSecret
metadata:
  name: payment-gateway-api-token
  annotations:
    cnrm.cloud.google.com/state-into-spec: absent
spec:
  # Customer-Managed Encryption Key (CMEK)
  customerManagedEncryption:
    kmsKeyRef:
      name: secret-crypto-key

  # Automatic multi-region replication
  replication:
    automatic: true

  # Automated rotation schedule: triggers notification topic every 30 days
  rotation:
    rotationPeriod: 2592000s # 30 days
    nextRotationTime: "2026-12-01T00:00:00Z"
  topics:
    - topicRef:
        name: secret-rotation-notifications-topic
```

---

## 7. Related Guides

* **[Config Connector Core Principles](./core-principles.md)**: Conceptual guide covering the reconciliation loop, annotations catalog, drift strategies, and dependency resolution.
* **[Managing Config Connector with GitOps (ArgoCD)](./gitops-argocd.md)**: Deploying and upgrading Config Connector via GitOps pipelines.
