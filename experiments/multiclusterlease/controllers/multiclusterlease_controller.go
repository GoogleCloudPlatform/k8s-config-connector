// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/leaderelection"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MultiClusterLeaseReconciler reconciles a MultiClusterLease object
type MultiClusterLeaseReconciler struct {
	client.Client
	Log        logr.Logger
	GCSClient  *storage.Client
	BucketName string
	Identity   string

	leaderElectorsMutex sync.Mutex
	leaderElectors      map[string]*leaderelection.LeaderElector // keyed by NamespacedName
}

// NewMultiClusterLeaseReconciler creates a new MultiClusterLeaseReconciler
func NewMultiClusterLeaseReconciler(
	client client.Client,
	log logr.Logger,
	gcsClient *storage.Client,
	bucketName string,
	identity string,
) *MultiClusterLeaseReconciler {
	return &MultiClusterLeaseReconciler{
		Client:         client,
		Log:            log,
		GCSClient:      gcsClient,
		BucketName:     bucketName,
		Identity:       identity,
		leaderElectors: make(map[string]*leaderelection.LeaderElector),
	}
}

// +kubebuilder:rbac:groups=multicluster.core.cnrm.cloud.google.com,resources=multiclusterleases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=multicluster.core.cnrm.cloud.google.com,resources=multiclusterleases/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=multicluster.core.cnrm.cloud.google.com,resources=multiclusterleases/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch

func (r *MultiClusterLeaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("multiclusterlease", req.NamespacedName)
	log.Info("reconciling multiclusterlease")

	// Fetch the MultiClusterLease object
	var mcl v1alpha1.MultiClusterLease
	if err := r.Get(ctx, req.NamespacedName, &mcl); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("MultiClusterLease not found in API server; skipping the reconciliation")
			return ctrl.Result{}, nil
		}
		log.Error(err, "failed to get multiclusterlease")
		return ctrl.Result{}, err
	}

	// TODO: validate MCL

	// Handle deletion
	if !mcl.GetDeletionTimestamp().IsZero() {
		return r.handleDeletion(ctx, &mcl)
	}

	if err := r.ensureFinalizer(ctx, &mcl); err != nil {
		log.Error(err, "failed to ensure finalizer")
		return ctrl.Result{}, err
	}

	// Get or create the LeaderElector for this lease
	leaderElector, err := r.getOrCreateLeaderElector(req.NamespacedName.String(), &mcl)
	if err != nil {
		log.Error(err, "failed to get or create leaderelector")
		return ctrl.Result{}, err
	}

	// Try to acquire or renew the lease in GCS
	leaseInfo, err := leaderElector.AcquireOrRenew(ctx)
	if err != nil {
		log.Error(err, "failed to acquire or renew lease",
			"identity", r.Identity,
			"currentHolder", leaseInfo.HolderIdentity)

		// Update condition to indicate backend is unhealthy
		r.setBackendHealthyCondition(&mcl, false, err)

		// Update status
		if updateErr := r.Status().Update(ctx, &mcl); updateErr != nil {
			log.Error(updateErr, "failed to update status with backend error")
			return ctrl.Result{}, updateErr
		}

		requeueAfter := r.calculateRequeueAfter(&mcl)
		return ctrl.Result{RequeueAfter: requeueAfter}, nil
	}

	// Log lease status
	if leaseInfo.Acquired {
		log.Info("successfully acquired lease",
			"identity", r.Identity,
			"renewTime", leaseInfo.RenewTime)
	} else {
		log.Info("lease held by another identity",
			"holderIdentity", leaseInfo.HolderIdentity,
			"renewTime", leaseInfo.RenewTime)
	}

	// Update status based on lease state
	r.setMCLStatus(&mcl, leaseInfo)
	// Set backend healthy condition
	r.setBackendHealthyCondition(&mcl, true, nil)
	// Set lock acquired condition
	r.setLockAcquiredCondition(&mcl, leaseInfo.HolderIdentity)

	// Update status
	if err := r.Status().Update(ctx, &mcl); err != nil {
		log.Error(err, "failed to update status")
		// Don't return the error to avoid immediate reconciliation
	}

	// Calculate next reconcile time based on lease parameters
	requeueAfter := r.calculateRequeueAfter(&mcl)
	log.Info("completed reconciliation", "nextReconcileIn", requeueAfter)

	return ctrl.Result{RequeueAfter: requeueAfter}, nil
}

func (r *MultiClusterLeaseReconciler) handleDeletion(ctx context.Context, mcl *v1alpha1.MultiClusterLease) (ctrl.Result, error) {
	log := r.Log.WithValues("multiclusterlease", client.ObjectKeyFromObject(mcl))

	// Check if the finalizer is already removed
	if !r.hasFinalizer(mcl) {
		log.Info("finalizer already removed, nothing to do")
		return ctrl.Result{}, nil
	}

	// TODO: get the leader elector and release the lease

	// Clean up the leader elector from our map
	r.leaderElectorsMutex.Lock()
	delete(r.leaderElectors, client.ObjectKeyFromObject(mcl).String())
	r.leaderElectorsMutex.Unlock()

	if err := r.removeFinalizer(ctx, mcl); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("successfully handled deletion")
	return ctrl.Result{}, nil
}

// setMCLStatus updates the status of the MultiClusterLease
func (r *MultiClusterLeaseReconciler) setMCLStatus(mcl *v1alpha1.MultiClusterLease, leaseInfo *leaderelection.LeaseInfo) {
	// Set leading cluster status
	mcl.Status.IsLeadingCluster = leaseInfo.Acquired && *leaseInfo.HolderIdentity == r.Identity

	// Set observed generation
	generation := mcl.Generation
	mcl.Status.ObservedGeneration = &generation

	// Update holder identity
	if leaseInfo.HolderIdentity != nil {
		mcl.Status.GlobalHolderIdentity = leaseInfo.HolderIdentity
	} else {
		mcl.Status.GlobalHolderIdentity = nil
	}

	// Update renew time as a string with second precision
	if leaseInfo.RenewTime != nil {
		timeStr := leaseInfo.RenewTime.Format(time.RFC3339)
		mcl.Status.GlobalRenewTime = &timeStr
	} else {
		mcl.Status.GlobalRenewTime = nil
	}

	// Update lease duration
	if mcl.Spec.LeaseDurationSeconds != nil {
		mcl.Status.GlobalLeaseDurationSeconds = mcl.Spec.LeaseDurationSeconds
	}

	// Update lease transitions
	if leaseInfo.LeaseTransitions != nil {
		mcl.Status.GlobalLeaseTransitions = leaseInfo.LeaseTransitions
	}
}

// setBackendHealthyCondition updates the backend healthy condition
func (r *MultiClusterLeaseReconciler) setBackendHealthyCondition(
	mcl *v1alpha1.MultiClusterLease,
	healthy bool,
	err error,
) {
	condition := metav1.Condition{
		Type:               string(v1alpha1.ConditionTypeBackendHealthy),
		ObservedGeneration: mcl.Generation,
		LastTransitionTime: metav1.Now(),
	}

	if healthy {
		condition.Status = metav1.ConditionTrue
		condition.Reason = "BackendHealthy"
		condition.Message = "successfully communicated with backend"
	} else {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "BackendError"
		condition.Message = fmt.Sprintf("failed to communicate with backend: %v", err)
	}

	meta.SetStatusCondition(&mcl.Status.Conditions, condition)
}

// setLockAcquiredCondition sets the LockAcquiredInBackend condition
func (r *MultiClusterLeaseReconciler) setLockAcquiredCondition(mcl *v1alpha1.MultiClusterLease, holderIdentity *string) {
	var condition metav1.Condition
	condition.Type = string(v1alpha1.ConditionTypeLockAcquiredInBackend)
	condition.ObservedGeneration = mcl.Generation

	if holderIdentity != nil {
		condition.Status = metav1.ConditionTrue
		condition.Reason = "LockAcquired"
		condition.Message = fmt.Sprintf("Lock is held by %s", *holderIdentity)
	} else {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "LockNotAcquired"
		condition.Message = "No cluster currently holds the lock in the backend"
	}

	meta.SetStatusCondition(&mcl.Status.Conditions, condition)
}

// calculateRequeueAfter determines how long to wait before the next reconciliation
func (r *MultiClusterLeaseReconciler) calculateRequeueAfter(mcl *v1alpha1.MultiClusterLease) time.Duration {
	// Default to 10 seconds if not specified
	retryPeriod := 10 * time.Second

	if mcl.Spec.RetryPeriodSeconds != nil {
		retryPeriod = time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second
	}

	// TODO: add jitter

	return retryPeriod
}

// getOrCreateLeaderElector gets an existing LeaderElector for the given lease or creates a new one
func (r *MultiClusterLeaseReconciler) getOrCreateLeaderElector(key string, lease *v1alpha1.MultiClusterLease) (*leaderelection.LeaderElector, error) {
	r.leaderElectorsMutex.Lock()
	defer r.leaderElectorsMutex.Unlock()

	if r.leaderElectors == nil {
		r.leaderElectors = make(map[string]*leaderelection.LeaderElector)
	}

	if le, exists := r.leaderElectors[key]; exists {
		return le, nil
	}

	// Create a new LeaderElector
	le := leaderelection.NewLeaderElector(r.GCSClient, r.BucketName, r.Identity, key, lease)
	r.leaderElectors[key] = le
	return le, nil
}

func (r *MultiClusterLeaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.MultiClusterLease{}).
		Complete(r)

	// TODO: add a predicate to filter out status updates
}
