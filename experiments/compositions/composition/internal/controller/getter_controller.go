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
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
)

// GetterConfigurationReconciler reconciles a GetterConfiguration object
type GetterConfigurationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=composition.google.com,resources=getterconfigurations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=composition.google.com,resources=getterconfigurations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=composition.google.com,resources=getterconfigurations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GetterConfiguration object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *GetterConfigurationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Got a new request!", "request", req)

	var obj compositionv1alpha1.GetterConfiguration
	if err := r.Client.Get(ctx, req.NamespacedName, &obj); err != nil {
		logger.Error(err, "unable to fetch GetterConfiguration")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Grab status for comparison later
	oldStatus := obj.Status.DeepCopy()

	// Try updating status before returning
	defer func() {
		if !reflect.DeepEqual(oldStatus, obj.Status) {
			newStatus := obj.Status.DeepCopy()
			err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				nn := types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}
				err := r.Client.Get(ctx, nn, &obj)
				if err != nil {
					return err
				}
				obj.Status = *newStatus.DeepCopy()
				return r.Client.Status().Update(ctx, &obj)
			})
			if err != nil {
				logger.Error(err, "unable to update Composition status")
			}
		}
	}()

	logger = logger.WithName(obj.Name).WithName(fmt.Sprintf("%d", obj.Generation))

	logger.Info("Validating ExpanderVersion object")
	if !obj.Validate() {
		logger.Info("Validation Failed")
		return ctrl.Result{}, fmt.Errorf("Validation failed")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GetterConfigurationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&compositionv1alpha1.GetterConfiguration{}).
		Complete(r)
}
