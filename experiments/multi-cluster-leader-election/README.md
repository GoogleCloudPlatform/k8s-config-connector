# multi-cluster-leader-election
This is a feature that extends the built-in leader election supported by controller runtime.


## Description
Leader election can be enabled in Controller-runtime managers. Currently, it only supports leader election (aka. Resource lock) between managers(and their controllers) inside the same cluster.

Multi-cluster leader election aims to extend this capability to support leader election for manages across multiple clusters.

### High level architecture
Multi-cluster leader election implements a CRD (type lease) to respresent the resource lock and implements its interfaces.

A lease controller runs per cluster and creates a lease object per namespace. 

The lease CRD is defined as below:

```go
type LeaseSpec struct {
	HolderIdentity string
	LeaseDurationSeconds int32 
}

type LeaseStatus struct {
	IsLeader bool 
	ObservedHolderIdentity string 
	ObservedAcquireTime metav1.MicroTime 
	ObservedRenewTime metav1.MicroTime 
	LeaseTransitions int32
}
```

The lease controller watches namespaces on the cluster, creating a new lock (lease) for each namespace. 

The controller is responsible for reconciling the lease object periodically based on leaseDuration then updates the lease status.

The backend behinds the lease object can be in many forms, such as a database, GCS bucket or any kind of storage that supports lock and multi-tenency. 

In this implementation, lease is backed by Google cloud storage. The lease controller has a storage client that talks with a GCS bucket and updates the lease status accordingly.

Leader election is per namespace and happens on each reconciliation where the operator is to reconcile the resource, so that the operator checks if it is currently the leader for the namespace before taking further actions.

In the example below, the guest-book controller checks the lease object in its namespace and decides whether it should do resource reconciling :

```go
func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

    if !r.checkResourceLockBeforeActuation(ctx, req.NamespacedName.Namespace) {
		logger.Info("Operator is currently not the leader, do nothing")
		return ctrl.Result{}, nil
	}

	guestBook := &webappv1.Guestbook{}
	if err := r.Get(ctx, req.NamespacedName, guestBook); err != nil {
		logger.Error(err, "Unable to fetch guest book in namespace", "ns", req.NamespacedName.Namespace)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Operator is the leader: Print guest book info", "book", guestBook.Spec.Book)

	return ctrl.Result{}, nil
}

func (r *GuestbookReconciler) checkResourceLockBeforeActuation(ctx context.Context, namespaceName string) bool {
	logger.Info("Checking who is the current leader before resource acutation for namespace", "ns", namespaceName)

	lease := &leaderelectionv1.Lease{}
	err := r.Get(ctx, client.ObjectKey{Namespace: namespaceName, Name: leaseName}, lease)
	if err != nil {
		logger.Error(err, "unbale to fetch the resource lock for namespace", "ns", namespaceName)
		return false
	}

	// Check lease object
	if reflect.ValueOf(lease.Status).IsZero() {
		return false
	}

	return lease.Status.IsLeader
}

```

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### GKE Workload Identify

**Run these commands once per project**
```sh
gcloud iam service-accounts create lease-controller-manager

gcloud projects add-iam-policy-binding PROJECT_ID --member "serviceAccount:lease-controller-manager@PROJECT_ID.iam.gserviceaccount.com"  --role "roles/storage.admin"

gcloud iam service-accounts add-iam-policy-binding  lease-controller-manager@PROJECT_ID.iam.gserviceaccount.com    --member="serviceAccount:PROJECT_ID.svc.id.goog[multi-cluster-leader-election-system/leader-election-controller-manager]"  --role="roles/iam.workloadIdentityUser"
```
**Run this command once per cluster**
```sh
kubectl annotate serviceaccount KSA_NAME \
    --namespace NAMESPACE \
    iam.gke.io/gcp-service-account=lease-controller-manager@GSA_PROJECT.iam.gserviceaccount.com
```
### To Deploy on the cluster
**Build and push your image to gcr.io:**

```sh
make docker-build docker-push
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make  make deploy-controller && kubectl delete pods -n multi-cluster-leader-election-system --all
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

