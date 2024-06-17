// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	multiclusterv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/leaderelection"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler/pauser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/resourcelock/bucketlease"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/util"
)

func New(c MultiClusterLeaseReconcilerConfig) *MultiClusterLeaseReconciler {
	return &MultiClusterLeaseReconciler{
		client:         c.Client,
		scheme:         c.Scheme,
		leaderElectors: make(map[string]runningLeaderElector),
		log:            c.Log,
		eventRecorder:  c.EventRecorder,
	}
}

type MultiClusterLeaseReconcilerConfig struct {
	Client        client.Client
	Scheme        *runtime.Scheme
	Log           logr.Logger
	EventRecorder record.EventRecorder
}

// MultiClusterLeaseReconciler reconciles a MultiClusterLease object
type MultiClusterLeaseReconciler struct {
	client        client.Client
	scheme        *runtime.Scheme
	log           logr.Logger
	eventRecorder record.EventRecorder

	leaderElectors      map[string]runningLeaderElector // a map of running leader electors keyed by `MultiClusterUID` and `Identity`.
	leaderElectorsMutex sync.Mutex
}

type runningLeaderElector struct {
	le *leaderelection.LeaderElector
	lh lifecyclehandler.LifecycleHandler // lifecycleHandlers to be called during the lifecycle of leader election.
}

// +kubebuilder:rbac:groups=multicluster.core.cnrm.cloud.google.com,resources=multiclusterleases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=multicluster.core.cnrm.cloud.google.com,resources=multiclusterleases/status,verbs=get;update;patch

func (r *MultiClusterLeaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	mcl, err := util.GetMultiClusterLease(ctx, r.client, req.NamespacedName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("MultiClusterLease not found in API server; skipping the reconciliation", "namespaced name", req.NamespacedName)
			return reconcile.Result{}, nil
		}
	}
	// TODO: validate mcl

	if !mcl.GetDeletionTimestamp().IsZero() { // handle deletion
		if !r.hasFinalizer(mcl) {
			// operator finalizer already removed. just waiting for MCL to be garbage collected.
			return reconcile.Result{}, nil
		}
		if err := r.stopLeaderElection(mcl); err != nil {
			return reconcile.Result{}, err
		}
		if err := r.removeFinalizer(ctx, mcl); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if err := r.ensureFinalizer(ctx, mcl); err != nil {
		return reconcile.Result{}, err
	}

	v := r.getOrCreateLeaderElector(ctx, mcl)
	le := v.le
	lh := v.lh
	if le.IsLeader() { // leader path
		if err := le.Renew(ctx); err != nil {
			return r.handleRenewFailed(ctx, mcl, le, lh)
		}
		return r.handleRenewSucceeded(ctx, mcl, le, lh)
	}
	if err := le.Acquire(ctx); err != nil { // non-leader path
		return r.handleAcquireFailed(ctx, mcl, le, lh)
	}
	return r.handleAcquireSucceeded(ctx, mcl, le, lh)
}

func (r *MultiClusterLeaseReconciler) getOrCreateLeaderElector(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease) runningLeaderElector {
	log := r.log.WithValues("multiClusterUID", mcl.Spec.MultiClusterUID, "identity", mcl.Spec.Identity)
	r.leaderElectorsMutex.Lock()
	defer r.leaderElectorsMutex.Unlock()

	if le, found := r.leaderElectors[mapKey(mcl)]; found {
		return le
	}

	bl := bucketlease.New(ctx, bucketlease.Config{
		Identity:      mcl.Spec.Identity,
		BucketName:    mcl.Spec.Bucket,
		LeaseName:     mcl.Spec.MultiClusterUID,
		Log:           log,
		EventRecorder: r.eventRecorder,
	})

	le := leaderelection.New(leaderelection.Config{
		ElectionID:    mcl.Spec.MultiClusterUID,
		Identity:      mcl.Spec.Identity,
		Lock:          bl,
		LeaseDuration: time.Duration(*mcl.Spec.LeaseDurationSeconds) * time.Second,
		RenewDeadline: time.Duration(*mcl.Spec.RenewDeadlineSeconds) * time.Second,
		RetryPeriod:   time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second,
		Log:           log,
		EventRecorder: r.eventRecorder,
	})

	lh := pauser.New(pauser.PauserConfig{
		Client:    r.client,
		Log:       r.log,
		Namespace: mcl.Namespace,
		Identity:  mcl.Spec.Identity,
	})
	runningLeaderElector := runningLeaderElector{
		le: le,
		lh: lh,
	}
	r.leaderElectors[mapKey(mcl)] = runningLeaderElector
	return runningLeaderElector
}

func (r *MultiClusterLeaseReconciler) handleRenewSucceeded(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, le *leaderelection.LeaderElector, _ lifecyclehandler.LifecycleHandler) (ctrl.Result, error) {
	if err := r.updateStatus(ctx, mcl, le.LeaderIdentity(), le.IsLeader()); err != nil {
		r.log.Error(err, "error updating MultiClusterLease status")
	}
	return ctrl.Result{RequeueAfter: time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second}, nil
	// TODO: add jitter
}

func (r *MultiClusterLeaseReconciler) handleRenewFailed(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, le *leaderelection.LeaderElector, lh lifecyclehandler.LifecycleHandler) (ctrl.Result, error) {
	if err := lh.OnStoppedLeading(ctx); err != nil {
		r.log.Error(err, "error calling stopped leading function")
	}
	if err := r.updateStatus(ctx, mcl, le.LeaderIdentity(), le.IsLeader()); err != nil {
		r.log.Error(err, "error updating MultiClusterLease status")
	}
	return ctrl.Result{RequeueAfter: time.Duration(*mcl.Spec.LeaseDurationSeconds) * time.Second}, nil
	// TODO: add jitter
}

func (r *MultiClusterLeaseReconciler) handleAcquireSucceeded(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, le *leaderelection.LeaderElector, lh lifecyclehandler.LifecycleHandler) (ctrl.Result, error) {
	if err := lh.OnStartedLeading(ctx); err != nil {
		r.log.Error(err, "error calling started leading function")
	}
	if err := r.updateStatus(ctx, mcl, le.LeaderIdentity(), le.IsLeader()); err != nil {
		r.log.Error(err, "error updating MultiClusterLease status")
	}
	return ctrl.Result{RequeueAfter: time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second}, nil
	// TODO: add jitter
}

func (r *MultiClusterLeaseReconciler) handleAcquireFailed(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, le *leaderelection.LeaderElector, lh lifecyclehandler.LifecycleHandler) (ctrl.Result, error) {
	if mcl.Status.LeaderIdentity != le.LeaderIdentity() { // observed a new leader
		if err := lh.OnNewLeader(ctx, le.LeaderIdentity()); err != nil {
			r.log.Error(err, "error calling new leader function")
		}
	}
	if err := r.updateStatus(ctx, mcl, le.LeaderIdentity(), le.IsLeader()); err != nil {
		r.log.Error(err, "error updating MultiClusterLease status")
	}
	return ctrl.Result{RequeueAfter: time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second}, nil
	// TODO: add jitter
}

func (r *MultiClusterLeaseReconciler) updateStatus(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, leaderIdentity string, isLeader bool) error {
	mcl, err := util.GetMultiClusterLease(ctx, r.client, types.NamespacedName{Name: mcl.Name, Namespace: mcl.Namespace})
	if err != nil {
		return fmt.Errorf("error getting MultiClusterLease: %w", err)
	}
	mcl.Status = multiclusterv1alpha1.MultiClusterLeaseStatus{
		IsLeader:         isLeader,
		LeaderIdentity:   leaderIdentity,
		LastObservedTime: time.Now().Format(time.RFC3339),
	}

	// TODO: retry in case there is a conflict
	err = r.client.Status().Update(ctx, mcl)
	if err != nil {
		return fmt.Errorf("error updating MultiClusterLease status: %w", err)
	}
	return nil
}

func (r *MultiClusterLeaseReconciler) stopLeaderElection(mcl *multiclusterv1alpha1.MultiClusterLease) error {
	r.leaderElectorsMutex.Lock()
	defer r.leaderElectorsMutex.Unlock()
	v, found := r.leaderElectors[mapKey(mcl)]
	if !found {
		r.log.Info("Leader election already stoppped", "multiClusterUID", mcl.Spec.MultiClusterUID, "identity", mcl.Spec.Identity)
		return nil
	}

	if err := v.lh.OnStopping(); err != nil {
		return err
	}
	delete(r.leaderElectors, mapKey(mcl))
	return nil
}

func (r *MultiClusterLeaseReconciler) StopAndWait() {
	r.log.Info("stopping all leader election loops")
	defer r.log.Info("all leader election loops stopped")

	r.leaderElectorsMutex.Lock()
	defer r.leaderElectorsMutex.Unlock()
	for _, v := range r.leaderElectors {
		if err := v.lh.OnStopping(); err != nil {
			r.log.Error(err, "error stopping leader election", "multiClusterID", v.le.ElectionID(), "identity", v.le.Identity())
		}
	}
}

func (r *MultiClusterLeaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&multiclusterv1alpha1.MultiClusterLease{}).
		Complete(r)
}
