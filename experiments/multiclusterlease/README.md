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
	multiclusterleaselock "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/client"
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


