# Setting up multi cluster support for KCC

> NOTE: This feature is in alpha stage and is actively being worked on.

Imagine we have two Kubernetes clusters, `mcl-cluster-1` and `mcl-cluster-2`. We want to run KCC in both for high availability, but only one instance should be the active leader at any time.

### Setup

#### 1. Cluster Lock: Powered GCS Bucket

```sh
# This is a one-time setup
gcloud storage buckets create gs://kcc-global-leader-lock
```

#### 2. The `multiclusterlease-controller` (MCL Controller)

Deploy the MCL controller to **both** `mcl-cluster-1` and `mcl-cluster-2`. This controller's job is to watch for `MultiClusterLease` CRs and contend for the lock in the GCS bucket.

```yaml
# Deployed in BOTH mcl-cluster-1 and mcl-cluster-2
apiVersion: apps/v1
kind: Deployment
metadata:
  name: multiclusterlease-controller-manager
  namespace: multiclusterlease-system
spec:
  replicas: 1
  template:
    spec:
      # This controller needs GCP credentials to talk to the bucket.
      # This is best done using Workload Identity, which binds the Kubernetes
      # Service Account (KSA) used by this pod to a Google Service Account (GSA).
      # Ensure the KSA (e.g., 'default' in 'multiclusterlease-system') is
      # configured for Workload Identity.
      serviceAccountName: default
      containers:
      - name: manager
        image: gcr.io/multiclusterlease/controller:latest
        command:
        - /manager
        args:
        - --gcs-bucket=kcc-global-leader-lock
```

At this point, we have an election controller running in each cluster, ready to act on behalf of any client that creates a `MultiClusterLease` resource.

#### 3. KCC integration & E2E Walkthrough

Now, let's set up the lease objects for when the two KCC pods start up in `mcl-cluster-1` and `mcl-cluster-2`.
In each of your clusters, configure your `ConfigConnector` objects such that you specify the name, namespace
where you want the lease objects to be created, and the name of the global lock you set up before. Example:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  experiments:
    leaderElection:
      multiClusterLease:
        leaseName: kcc-mcl-cluster-1-lease # kcc-mcl-cluster-2-lease for your other cluster
        namespace: kcc-system
        globalLockName: kcc-global-leader-lock
```

At this point, you will also need to create the MultiClusterLease CRs for now. In the future, the KCC manager or
the operator may manage the lifecycle of these resources.

```yaml
# In mcl-cluster-1
apiVersion: multicluster.core.cnrm.cloud.google.com/v1alpha1
kind: MultiClusterLease
metadata:
  name:  kcc-mcl-cluster-1-lease
  namespace: kcc-system
spec:
  leaseDurationSeconds: 60
status: {} # Status is empty for now
```

At this point when the KCC manager pods start, they would start with leader election configured. One pod will
be elected leader and the other will exit. 
