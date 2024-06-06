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
	"sync"

	"github.com/go-logr/logr"
	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var FacadeControllers sync.Map

// CompositionReconciler reconciles a Composition object
type CompositionReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
	mgr      ctrl.Manager
}

//+kubebuilder:rbac:groups=composition.google.com,resources=compositions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=composition.google.com,resources=compositions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=composition.google.com,resources=compositions/finalizers,verbs=update
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=facade.facade,resources=*,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;create;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;create;patch;delete
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=create;get;patch;list;delete
//+kubebuilder:rbac:groups="batch",resources=jobs,verbs=create;get;patch;list;delete

// /
// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *CompositionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Got a new request!", "request", req)

	var composition compositionv1alpha1.Composition
	if err := r.Client.Get(ctx, req.NamespacedName, &composition); err != nil {
		logger.Error(err, "unable to fetch Composition")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Grab status for comparison later
	oldStatus := composition.Status.DeepCopy()

	// Try updating status before returning
	defer func() {
		if !reflect.DeepEqual(oldStatus, composition.Status) {
			newStatus := composition.Status.DeepCopy()
			err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				nn := types.NamespacedName{Namespace: composition.Namespace, Name: composition.Name}
				err := r.Client.Get(ctx, nn, &composition)
				if err != nil {
					return err
				}
				composition.Status = *newStatus.DeepCopy()
				return r.Client.Status().Update(ctx, &composition)
			})
			if err != nil {
				logger.Error(err, "unable to update Composition status")
			}
		}
	}()

	logger = logger.WithName(composition.Name).WithName(fmt.Sprintf("%d", composition.Generation))

	logger.Info("Validating Compostion object")
	if !composition.Validate() {
		logger.Info("Validation Failed")
		return ctrl.Result{}, fmt.Errorf("Validation failed")
	}

	composition.Status.ClearCondition(compositionv1alpha1.Error)
	logger.Info("Processing Composition object")
	if err := r.processComposition(ctx, &composition, logger); err != nil {
		logger.Info("Error processing Composition")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CompositionReconciler) processComposition(
	ctx context.Context, c *compositionv1alpha1.Composition, logger logr.Logger,
) error {
	var crd extv1.CustomResourceDefinition
	logger = logger.WithName(c.Spec.InputAPIGroup)

	crdName := c.Spec.InputAPIGroup

	err := r.Client.Get(ctx, types.NamespacedName{Name: crdName, Namespace: ""}, &crd)
	if err != nil {
		reason := "FailedGettingFacadeCRD"
		if apierrors.IsNotFound(err) {
			reason = "MissingFacadeCRD"
		}
		c.Status.Conditions = append(c.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             reason,
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		logger.Error(err, "failed to get an Facade CRD object")
		r.Recorder.Event(c, "Warning", "MissingFacadeCRD",
			fmt.Sprintf("Failed to get Facade CRD: %s", c.Spec.InputAPIGroup))
		return err
	}

	logger.Info("Found InputAPI CRD", "Group", crd.Spec.Group,
		"Version", crd.Spec.Versions[0].Name, "Kind", crd.Spec.Names.Kind)

	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: crd.Spec.Versions[0].Name,
		Kind:    crd.Spec.Names.Kind,
	}
	cr := &unstructured.Unstructured{}
	cr.SetGroupVersionKind(gvk)

	// TODO(barni@) Stop existing reconciler and start a new one
	logger.Info("Checking if Reconciler already exists for InputAPI CRD")
	_, loaded := FacadeControllers.LoadOrStore(gvk, true)
	if loaded {
		// Reconciler already exists nothing to be done
		logger.Info("Reconciler already exists for InputAPI CRD")
		return nil
	}

	logger.Info("Starting Reconciler for InputAPI CRD")
	expanderController := &ExpanderReconciler{
		Client:      r.Client,
		Recorder:    r.mgr.GetEventRecorderFor(crd.Spec.Names.Plural + "-expander"),
		Scheme:      r.Scheme,
		InputGVK:    gvk,
		Composition: types.NamespacedName{Name: c.Name, Namespace: c.Namespace},
		InputGVR:    gvk.GroupVersion().WithResource(crd.Spec.Names.Plural),
		RESTMapper:  r.mgr.GetRESTMapper(),
		Config:      r.mgr.GetConfig(),
	}

	if err := expanderController.SetupWithManager(r.mgr, cr); err != nil {
		c.Status.Conditions = append(c.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             "InternalError",
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		logger.Error(err, "Failed to start reconciler for InputAPI CRD")
		return err
	}
	r.Recorder.Event(c, "Normal", "InputReconcilerStarted",
		fmt.Sprintf("Reconciler started for Facade CR: %s", c.Spec.InputAPIGroup))

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CompositionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.mgr = mgr
	return ctrl.NewControllerManagedBy(mgr).
		For(&compositionv1alpha1.Composition{}).
		Complete(r)
}
