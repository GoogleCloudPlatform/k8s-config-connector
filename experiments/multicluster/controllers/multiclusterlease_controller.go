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
	"sync"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	multiclusterv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler/pauser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/lifecyclehandler/statusupdater"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/resourcelock/bucketlease"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/util"
)

func New(c MultiClusterLeaseReconcilerConfig) *MultiClusterLeaseReconciler {
	return &MultiClusterLeaseReconciler{
		client:          c.Client,
		scheme:          c.Scheme,
		electionRunning: make(map[string]context.CancelFunc),
		log:             c.Log,
		eventRecorder:   c.EventRecorder,
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

	electionRunning      map[string]context.CancelFunc
	electionRunningMutex sync.Mutex
	wg                   sync.WaitGroup
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
		r.stopLeaderElection(mcl)
		if err := r.removeFinalizer(ctx, mcl); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if err := r.ensureFinalizer(ctx, mcl); err != nil {
		return reconcile.Result{}, err
	}

	if err := r.ensureLeaderElectionIsRunning(ctx, mcl, req.NamespacedName); err != nil {
		return reconcile.Result{}, err
	}

	// TODO: add jitter so that controllers for multiple namepaces don't reconcile at the same time.
	// TODO: figure out requeue interval.
	return reconcile.Result{RequeueAfter: 10 * time.Minute}, nil
}

func (r *MultiClusterLeaseReconciler) ensureLeaderElectionIsRunning(ctx context.Context, mcl *multiclusterv1alpha1.MultiClusterLease, nn types.NamespacedName) error {
	log := r.log.WithValues("multiClusterUID", mcl.Spec.MultiClusterUID, "identity", mcl.Spec.Identity, "name", nn.Name, "namespace", nn.Namespace)

	// check if leader election is running
	r.electionRunningMutex.Lock()
	if _, found := r.electionRunning[mapKey(mcl)]; found {
		log.Info("MultiClusterLease is already running")
		r.electionRunningMutex.Unlock()
		return nil
	}
	internalCtx, cancel := context.WithCancel(ctx)
	r.electionRunning[mapKey(mcl)] = cancel
	r.electionRunningMutex.Unlock()

	// launch leader election loop
	log.Info("launching leader election loop")
	go func() {
		defer func() {
			log.Info("leader election loop stopped")
			r.electionRunningMutex.Lock()
			r.wg.Done()
			delete(r.electionRunning, mapKey(mcl))
			r.electionRunningMutex.Unlock()
		}()

		r.wg.Add(1)
		p := pauser.New(pauser.PauserConfig{
			Client:    r.client,
			Log:       r.log,
			Namespace: nn.Namespace,
			Identity:  mcl.Spec.Identity,
		})
		s := statusupdater.New(statusupdater.Config{
			Client:   r.client,
			Identity: mcl.Spec.Identity,
			NN:       nn,
		})
		blc := bucketlease.Config{
			Identity:      mcl.Spec.Identity,
			BucketName:    mcl.Spec.Bucket,
			LeaseName:     mcl.Spec.MultiClusterUID,
			Log:           r.log,
			EventRecorder: r.eventRecorder,
		}
		le, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
			Lock:          bucketlease.New(ctx, blc),
			LeaseDuration: time.Duration(*mcl.Spec.LeaseDurationSeconds) * time.Second, // TODO: handle nil value
			RenewDeadline: time.Duration(*mcl.Spec.RenewDeadlineSeconds) * time.Second,
			RetryPeriod:   time.Duration(*mcl.Spec.RetryPeriodSeconds) * time.Second,
			Callbacks:     lifecyclehandler.ChainCallbacks(p.Callbacks(ctx), s.Callbacks(ctx)),
		})
		if err != nil {
			log.Error(err, "failed to create leader elector")
			return
		}

		le.Run(internalCtx)
	}()

	return nil
}

func (r *MultiClusterLeaseReconciler) stopLeaderElection(mcl *multiclusterv1alpha1.MultiClusterLease) {
	r.electionRunningMutex.Lock()
	if cancel, found := r.electionRunning[mapKey(mcl)]; found {
		cancel()
	}
	r.electionRunningMutex.Unlock()
}

func (r *MultiClusterLeaseReconciler) StopAndWait() {
	r.log.Info("stopping all leader election loops")
	defer r.log.Info("all leader election loops stopped")

	r.electionRunningMutex.Lock()
	for _, cancel := range r.electionRunning {
		cancel()
	}
	r.electionRunningMutex.Unlock()
	r.wg.Wait()
}

func (r *MultiClusterLeaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&multiclusterv1alpha1.MultiClusterLease{}).
		Complete(r)
}
