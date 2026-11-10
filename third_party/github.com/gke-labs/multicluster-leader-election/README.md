# Multi-Cluster Leader Election

This project provides a robust, decentralized, and Kubernetes-native leader election mechanism that allows a single replica of a controller to be elected as a leader from a pool of candidates running across multiple Kubernetes clusters.

## How It Works

The system consists of two primary components:

1.  **A `multiclenterlease-controller`:** A controller that runs in each participating cluster. It watches local `MultiClusterLease` custom resources and contends for a global lock (e.g., a GCS object) on behalf of local candidates.
2.  **A Client Library (`resourcelock`):** A Go library that implements `client-go`'s standard `resourcelock.Interface`. Client controllers import this library to participate in the election.

The `MultiClusterLease` CRD acts as a communication bridge between the client controllers and the election controller, ensuring that clients remain completely decoupled from the global backend.

## Usage

To use this system in a controller built with `controller-runtime`, you provide an instance of the custom `MultiClusterLeaseLock` directly to the manager's options. The manager will then use this custom lock for its leader election process.

This is typically done in your `main.go`:

```go
// In your controller's cmd/manager/main.go

import (
	"context"
	"os"
	"time"

	// ... other imports
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// Import the custom lock library
	multiclusterleaselock "github.com/gke-labs/multicluster-leader-election/pkg/client"
)

func main() {
    // ... standard flag parsing and setup ...

    // The identity should be a unique name for a candidate pod
    podIdentity, err := os.Hostname() + "_" + string(uuid.NewUUID())
    // ... handle error ...

    // 1. Create an instance of the custom MultiClusterLeaseLock.
    // The manager will use this object to contend for leadership.
    myGlobalLock := multiclusterleaselock.New(
        mgr.GetClient(),
        "my-global-leader-lock",           // The name for the lock object.
        "my-global-leader-lock-namespace", // The namespace for the lock object.
        podIdentity,
        15*time.Second,                    // The retry period.
    )

    // 2. Create the Manager, enabling leader election and providing the custom lock.
    mgr, err := manager.New(cfg, manager.Options{
        // ... other options ...
        LeaderElection:                     true,
        LeaderElectionResourceLockInterface: myGlobalLock, // <-- Provide the custom lock here!
    })
    // ... handle error ...

    // ... register your controllers with the manager ...

    // 3. Start the manager as usual.
    // The manager will now handle the entire leader election lifecycle internally
    // using our custom multi-cluster lock.
    if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
        // ... handle error ...
    }
}
```

## Local End-to-End (E2E) Testing

These instructions explain how to run the end-to-end tests locally. The tests use live Google Cloud Storage (GCS) as a backend for the resource lock and `kind` to create local Kubernetes clusters.

### Prerequisites

1.  **gcloud CLI**: Ensure you have the `gcloud` command-line tool installed and authenticated.
2.  **A Google Cloud Project**: Have a GCP project where you can create service accounts and GCS buckets.
3.  **kind**: Install `kind` to run local Kubernetes clusters.
4.  **make**: The `make` command should be available in your shell.

### Step 1: Configure GCP Service Account and GCS Bucket

The E2E tests require a GCP Service Account with permissions to access a GCS bucket, which will be used for the leader election lock.

First, set up your environment variables:

```bash
export PROJECT_ID=$(gcloud config get-value project)
export SA_NAME="multiclusterlease-e2e-tester"
export BUCKET_NAME="multiclusterlease-test" # Bucket names must be globally unique. Choose a name that does not already exist.

Now, run the following commands to create the bucket, the service account, grant permissions, and download the key.

```bash
# Create the GCS bucket (if it doesn't exist)
gcloud storage buckets create gs://${BUCKET_NAME} --project=${PROJECT_ID}

# Create the service account
gcloud iam service-accounts create ${SA_NAME} \
  --project=${PROJECT_ID} \
  --display-name="MultiClusterLease E2E Tester"

# Grant the service account permissions to manage objects in the bucket
gcloud storage buckets add-iam-policy-binding gs://${BUCKET_NAME} \
  --member="serviceAccount:${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com" \
  --role="roles/storage.admin"

# Download the service account key
gcloud iam service-accounts keys create ./keyfile.json \
  --iam-account=${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com

# Export the path to the key file for the tests
export GCP_SA_KEY_PATH="$(pwd)/keyfile.json"
```

Important: The keyfile.json contains sensitive credentials. Ensure it is included in your .gitignore file and is never committed to your repository.

### Step 2: Run the E2E tests

Before running the tests, ensure the `--gcs-bucket` argument in `config/manager/manager.yaml` matches the actual bucket name in use.

```
make test-e2e

make test-e2e-multi
```

The test will verify leader election, lock renewal, and leadership failover scenarios.