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
	"strings"

	"github.com/go-logr/logr"
	"github.com/gobuffalo/flect"
	"google.com/composition/pkg/crds"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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

	compositionv1alpha1 "google.com/composition/api/v1alpha1"
)

// FacadeBindingReconciler reconciles a FacadeBinding object
type FacadeBindingReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Recorder      record.EventRecorder
	ImageRegistry string
	mgr           ctrl.Manager
}

//+kubebuilder:rbac:groups=composition.google.com,resources=facadebindings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=composition.google.com,resources=facadebindings/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=composition.google.com,resources=facadebindings/finalizers,verbs=update
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=facade.facade,resources=*,verbs=get;list;watch;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *FacadeBindingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Got a new request!", "request", req)

	var fb compositionv1alpha1.FacadeBinding
	if err := r.Client.Get(ctx, req.NamespacedName, &fb); err != nil {
		logger.Error(err, "unable to fetch FacadeBinding")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Grab status for comparison later
	oldStatus := fb.Status.DeepCopy()

	// Try updating status before returning
	defer func() {
		if !reflect.DeepEqual(oldStatus, fb.Status) {
			newStatus := fb.Status.DeepCopy()
			err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				nn := types.NamespacedName{Namespace: fb.Namespace, Name: fb.Name}
				err := r.Client.Get(ctx, nn, &fb)
				if err != nil {
					return err
				}
				fb.Status = *newStatus.DeepCopy()
				return r.Client.Status().Update(ctx, &fb)
			})
			if err != nil {
				logger.Error(err, "unable to update FaceBinding status")
			}
		}
	}()

	logger = logger.WithName(fb.Name).WithName(fmt.Sprintf("%d", fb.Generation))

	logger.Info("Validating FacadeBinding object")
	if !fb.Validate() {
		logger.Info("Validation Failed")
		return ctrl.Result{}, fmt.Errorf("Validation failed")
	}

	fb.Status.ClearCondition(compositionv1alpha1.Error)
	//fb.Status.ClearCondition(compositionv1alpha1.Ready)
	if fb.Spec.OpenAPIV3Schema != nil {
		if err := r.createCRD(ctx, &fb, logger); err != nil {
			logger.Info("Error creating CRD")
			return ctrl.Result{}, err
		}
	}
	logger.Info("Processing FacadeBinding object")
	if err := r.processBinding(ctx, &fb, logger); err != nil {
		logger.Info("Error processing FaceBinding")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *FacadeBindingReconciler) createCRD(ctx context.Context,
	fb *compositionv1alpha1.FacadeBinding, logger logr.Logger) error {
	var crd extv1.CustomResourceDefinition
	logger = logger.WithName(fb.Spec.FacadeKind)

	plural := strings.ToLower(flect.Pluralize(fb.Spec.FacadeKind))
	crdName := fmt.Sprintf("%s.%s", plural, crds.FacadeGroup)

	logger.Info("Checking if CRD exists", "crd", crdName)
	err := r.Client.Get(ctx, types.NamespacedName{Name: crdName, Namespace: ""}, &crd)
	// CRD exists. Nothing to be done.
	if err == nil {
		logger.Info("CRD exists. Not creating.", "crd", crdName)
		return nil
	}

	// If we are unable to get it for some reason other than not found return
	if !apierrors.IsNotFound(err) {
		fb.Status.Conditions = append(fb.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             "FailedGettingFacadeCRD",
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		logger.Error(err, "failed to get an Facade CRD object")
		r.Recorder.Event(fb, "Warning", "MissingFacadeCRD", fmt.Sprintf("Failed to get Facade CRD: %s", fb.Spec.FacadeAPI))
		return err
	}

	// Construct Facade CRD from the openAPI Schema
	crdInfo := crds.NewFacadeCRDInfo(fb.Spec.FacadeKind, plural, nil, "v1", nil, nil)
	err = crdInfo.SetSpec(fb.Spec.OpenAPIV3Schema)
	if err == nil {
		unversionedFacadeCRD, err := crdInfo.CRD()
		if err == nil {
			facadeCRD := &apiextensionsv1.CustomResourceDefinition{}
			err = r.Scheme.Convert(unversionedFacadeCRD, facadeCRD, nil)
			if err == nil {
				err = r.Client.Create(ctx, facadeCRD)
				if err != nil {
					logger.Error(err, "failed to Create Facade CRD")
				}
			} else {
				logger.Error(err, "CRD conversion error")
			}
		} else {
			logger.Error(err, "Error getting unversioned CRD")
		}
	} else {
		logger.Error(err, "Unable to set CRD Spec from Schema")
	}

	if err != nil {
		fb.Status.Conditions = append(fb.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             "CreateFacadeCRDFailed",
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		r.Recorder.Event(fb, "Warning", "CreateFacadeCRDFailed", fmt.Sprintf("Failed to Create Facade CRD: %s", fb.Spec.FacadeKind))
	}
	logger.Info("Created Facade CRD", "crd", crdName)
	return err
}

func (r *FacadeBindingReconciler) processBinding(
	ctx context.Context, fb *compositionv1alpha1.FacadeBinding, logger logr.Logger,
) error {
	var crd extv1.CustomResourceDefinition
	logger = logger.WithName(fb.Spec.FacadeAPI)

	crdName := fb.Spec.FacadeAPI
	if fb.Spec.OpenAPIV3Schema != nil {
		plural := strings.ToLower(flect.Pluralize(fb.Spec.FacadeKind))
		crdName = fmt.Sprintf("%s.%s", plural, crds.FacadeGroup)
	}

	err := r.Client.Get(ctx, types.NamespacedName{Name: crdName, Namespace: ""}, &crd)
	if err != nil {
		reason := "FailedGettingFacadeCRD"
		if apierrors.IsNotFound(err) {
			reason = "MissingFacadeCRD"
		}
		fb.Status.Conditions = append(fb.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             reason,
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		logger.Error(err, "failed to get an Facade CRD object")
		r.Recorder.Event(fb, "Warning", "MissingFacadeCRD", fmt.Sprintf("Failed to get Facade CRD: %s", fb.Spec.FacadeAPI))
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
		logger.Info("Reconciler already exists for Facade CRD")
		return nil
	}

	logger.Info("Starting Reconciler for InputAPI CRD")
	expanderController := &ExpanderReconciler{
		Client:        r.Client,
		Recorder:      r.mgr.GetEventRecorderFor(crd.Spec.Names.Plural + "-expander"),
		Scheme:        r.Scheme,
		InputGVK:      gvk,
		ImageRegistry: r.ImageRegistry,
		Composition:   types.NamespacedName{Name: fb.Spec.CompositionName, Namespace: fb.Spec.CompositionNamespace},
		InputGVR:      gvk.GroupVersion().WithResource(crd.Spec.Names.Plural),
		RESTMapper:    r.mgr.GetRESTMapper(),
		Config:        r.mgr.GetConfig(),
	}

	if err := expanderController.SetupWithManager(r.mgr, cr); err != nil {
		fb.Status.Conditions = append(fb.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            err.Error(),
			Reason:             "InternalError",
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		logger.Error(err, "Failed to start reconciler for Facade CRD")
		return err
	}
	r.Recorder.Event(fb, "Normal", "FacadeReconcilerStarted", fmt.Sprintf("Reconciler started for Facade CR: %s", fb.Spec.FacadeAPI))

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FacadeBindingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.mgr = mgr
	return ctrl.NewControllerManagedBy(mgr).
		For(&compositionv1alpha1.FacadeBinding{}).
		Complete(r)
}
