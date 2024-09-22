/*
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
*/

package controller

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"

	leaderelectionv1 "github.com/600lyy/multi-cluster-leader-election/api/v1"
	"github.com/600lyy/multi-cluster-leader-election/pkg/jitter"
	"github.com/600lyy/multi-cluster-leader-election/pkg/leaderelection"
	"github.com/600lyy/multi-cluster-leader-election/pkg/resourcelock"
)

var (
	leaseOperatorAnnotation = "leaderelection.600lyy.io/lease-operator"
	leaseName               = "lease-gsc"
	controllerName          = "lease-controller"
	logger                  = crlog.Log.WithName(controllerName)
)

// LeaseReconciler reconciles a Lease object
// LeaseReconciler watches namespace objects
type LeaseReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	LeaderElector *leaderelection.LeaderElector
	Identify      string
}

//+kubebuilder:rbac:groups=leaderelection.600lyy.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=leaderelection.600lyy.io,resources=leases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=leaderelection.600lyy.io,resources=leases/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=*

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *LeaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger.Info("reconcling request is from", "namespace", req.NamespacedName.Name)

	namespaceObj := &corev1.Namespace{}
	err := r.Get(ctx, req.NamespacedName, namespaceObj)
	if err != nil {
		logger.Error(err, "unable to fetch namespace", "ns", req.NamespacedName.Name)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	namespaceName := namespaceObj.Name
	lease := &leaderelectionv1.Lease{}
	err = r.Get(ctx, client.ObjectKey{Namespace: namespaceName, Name: leaseName}, lease)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			logger.Error(err, "unable to fetch the lease for namespace", "ns", namespaceName)
			return ctrl.Result{}, err
		}
		// Create a lease object for the namespace if not present
		// Hold off to create a lease in the storage bucket until next reconcile
		ler := &resourcelock.LeaderElectionRecord{
			HolderIdentity:       r.Identify,
			LeaseDurationSeconds: int(r.LeaderElector.Config.LeaseDuration / time.Second),
		}
		lease = &leaderelectionv1.Lease{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespaceName,
				Name:      "lease-gsc",
				Annotations: map[string]string{
					"managed-by": leaseOperatorAnnotation,
				},
			},
			Spec: resourcelock.LeaderElectionRecordToLeaseSpec(ler),
		}

		if err = r.Create(ctx, lease); err != nil {
			logger.Error(err, "failed to create the lease object for namespace", "ns", namespaceName)
			return ctrl.Result{}, err
		}
	} else {
		// reconcile to acquire or renew the lesase in storage bucket
		cctx, cancel := context.WithCancel(ctx)
		defer cancel()

		r.LeaderElector.Config.Lock.(*resourcelock.LeaseLock).LockConfig.Identity = r.Identify
		r.LeaderElector.Config.Lock.(*resourcelock.LeaseLock).LeaseFile = namespaceName + "-" + leaseName
		if err = r.LeaderElector.TryAcquireOrRenew(cctx); err != nil {
			logger.Error(err, "unable to acquire or renew the lease file in storage bucket")
			return ctrl.Result{}, nil
		}

		// update lease status
		lease.Status = resourcelock.LeaderElectionRecordToLeaseStatus(&r.LeaderElector.ObservedRecord)
		lease.Status.IsLeader = r.LeaderElector.IsLeader()
		if err = r.Status().Update(ctx, lease); err != nil {
			logger.Error(err, "unable to update lease status")
			return ctrl.Result{}, err
		}
	}

	// If no error, heading to the next reconcile
	if reconcileReenqueuePeriod, err := jitter.GenerateJitterReenqueuePeriod(lease); err != nil {
		logger.Error(err, "unable to requeue request to trigger next reconcile")
		return ctrl.Result{}, err
	} else {
		logger.Info("successully finished reconcile", "ns", namespaceName, "lease", lease.Name, "time to next reconcile", reconcileReenqueuePeriod)
		return ctrl.Result{RequeueAfter: reconcileReenqueuePeriod}, nil
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *LeaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Named(controllerName).
		Watches(
			&corev1.Namespace{},
			&handler.EnqueueRequestForObject{},
			builder.WithPredicates(ManagedByLeaseControllerPredicate{}),
		).
		Complete(r)
}
