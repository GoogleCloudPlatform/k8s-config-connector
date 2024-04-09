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

package controller

import (
	"context"
	"fmt"

	compositionv1 "google.com/composition/api/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ExpanderReconciler reconciles a expander object
type ExpanderReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Recorder      record.EventRecorder
	RESTMapper    meta.RESTMapper
	Config        *rest.Config
	ImageRegistry string
	Dynamic       *dynamic.DynamicClient
	InputGVK      schema.GroupVersionKind
	Resource      string
	Composition   types.NamespacedName
}

var planGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1",
	Kind:    "Plan",
}

func (r *ExpanderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger = logger.WithName(r.Composition.Name).WithName(r.InputGVK.Group)

	logger.Info("Got Input API for expansion", "request", req)

	inputcr := unstructured.Unstructured{}
	inputcr.SetGroupVersionKind(r.InputGVK)
	if err := r.Get(ctx, req.NamespacedName, &inputcr); err != nil {
		logger.Error(err, "unable to fetch Input API Object")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	loggerCR := logger.WithName(inputcr.GetName())

	// Grab the latest composition
	// TODO(barni@) - Decide how we want the latest composition changes are to be applied.
	var compositionCR compositionv1.Composition
	if err := r.Client.Get(ctx, r.Composition, &compositionCR); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Composition Object")
			return ctrl.Result{}, err
		}
	}

	// Associate a plan object with this input CR
	var plancr compositionv1.Plan
	planNN := types.NamespacedName{Name: r.Resource + "-" + inputcr.GetName(), Namespace: inputcr.GetNamespace()}
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Plan Object")
			return ctrl.Result{}, err
		}
		// create a plan object
		plancr = compositionv1.Plan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      planNN.Name,
				Namespace: planNN.Namespace,
			},
			Spec: compositionv1.PlanSpec{
				Stages: map[string]compositionv1.Stage{},
			},
		}
		if err := ctrl.SetControllerReference(&inputcr, &plancr, r.Scheme); err != nil {
			logger.Error(err, "Unable to set controller reference for Plan Object")
			return ctrl.Result{}, err
		}
		if err := r.Client.Create(ctx, &plancr); err != nil {
			logger.Error(err, "Unable to create Plan Object")
			return ctrl.Result{}, err
		}

		// Read after create to get the UID
		if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
			logger.Error(err, "Unable to fetch Plan Object")
			return ctrl.Result{}, err
		}
	}
	// TODO(barni@): Handle existing jobs for this CR

	// Since applylib doesnt support multiple apply batches we are
	// tracking all old objects and accumulating them each apply.
	oldAppliers := []*Applier{}

	for index, expander := range compositionCR.Spec.Expanders {
		// -------------------- FETCH VALUES ---------------------------
		// TODO(barni@): identify dependend variables and path where values need to be read from
		// Update the CRD_V's status to reflect these values
		if len(expander.ValuesFrom) != 0 {
			logger = loggerCR.WithName(expander.Name).WithName("Fetcher")
			ff := NewFetcher(ctx, logger, r, &plancr, &inputcr, &expander)
			err := ff.Fetch()
			if err != nil {
				r.Recorder.Event(&inputcr, "Warning", "FetcherFailed", fmt.Sprintf("Error getting values for expander: %s", expander.Name))
				logger.Error(err, "Unable to fetch dependent valuesFrom ")
				return ctrl.Result{}, err
			}
			err = ff.UpdatePlanCR()
			if err != nil {
				r.Recorder.Event(&inputcr, "Warning", "FetcherFailed", fmt.Sprintf("Error updating fetched values for expander: %s", expander.Name))
				return ctrl.Result{}, err
			}
		}

		// ------------------- EXPANSION SECTION -----------------------
		logger = loggerCR.WithName(expander.Name).WithName("Expand")

		jf := NewJobFactory(ctx, logger, r, &inputcr, expander.Name, planNN.Name, r.ImageRegistry)

		// Create Expander Job and wait for the Job to complete
		logger.Info("Creating expander job")
		err := jf.Create()
		defer jf.CleanUp()
		if err != nil {
			// Reference for event API: Event(object runtime.Object, eventtype, reason, message string)
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Error creating job for name: %s", expander.Name))
			logger.Error(err, "Unable to create expander job")
			return ctrl.Result{}, err
		}
		r.Recorder.Event(&inputcr, "Normal", "ExpansionStarted", fmt.Sprintf("Job Created for name: %s", expander.Name))
		logger.Info("Successfully created expander job")

		success, err := jf.Wait()
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Error waiting for Job for name: %s", expander.Name))
			logger.Error(err, "Failed waiting for expander job")
			return ctrl.Result{}, err
		}

		if !success {
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Job failed for name: %s", expander.Name))
			logger.Info("Expander job completed but Failed")
			return ctrl.Result{}, fmt.Errorf("Expander Job Failed")
		}
		r.Recorder.Event(&inputcr, "Normal", "ExpansionSucceded", fmt.Sprintf("Job succeded for name: %s", expander.Name))
		logger.Info("Expander job Completed successfully")

		// ------------------- APPLIER SECTION -----------------------

		// Create Applier and wait for the Applier to complete
		// TODO(barni@): create CRD_P (plan) or treat CRD_V (facade/cloudsql) as plan and apply the expanded resources
		logger = loggerCR.WithName(expander.Name).WithName("Apply")

		// Re-read the Plan CR to load the expanded manifests
		if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
			logger.Error(err, "unable to read Plan CR")
			return ctrl.Result{}, err
		}

		applier := NewApplier(ctx, logger, r, &plancr, &inputcr, expander.Name)

		err = applier.Load() // Load Manifests
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed", fmt.Sprintf("error loading manifests for expander, name: %s", expander.Name))
			logger.Error(err, "Unable to Load manifests for applying")
			return ctrl.Result{}, err
		}
		logger.Info("Successfully loaded manifests for applying")

		prune := false
		if index == len(compositionCR.Spec.Expanders)-1 {
			prune = true
		}

		err = applier.Apply(oldAppliers, prune) // Apply Manifests
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed", fmt.Sprintf("error applying manifests for expander, name: %s", expander.Name))
			logger.Error(err, "Unable to apply manifests")
			return ctrl.Result{}, err
		}

		logger.Info("Successfully applied manifests")
		r.Recorder.Event(&inputcr, "Normal", "ResourcesApplied", fmt.Sprintf("All expanded resources were applied. name: %s", expander.Name))

		success, err = applier.Wait()
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Failed waiting for resources to be reconciled. name: %s", expander.Name))
			logger.Error(err, "Failed waiting for applied resources to reconcile")
			return ctrl.Result{}, err
		}

		oldAppliers = append(oldAppliers, applier)

		if !success {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Some resources are not healthy. name: %s", expander.Name))
			logger.Info("Applied succesfully but some resources did not become healthy")
			return ctrl.Result{}, fmt.Errorf("Some applied resources are not healthy")
		}

		r.Recorder.Event(&inputcr, "Normal", "ResourcesReconciled", fmt.Sprintf("All applied resources were reconciled. name: %s", expander.Name))
		logger.Info("Applied resources successfully.")

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExpanderReconciler) SetupWithManager(mgr ctrl.Manager, cr *unstructured.Unstructured) error {
	var err error
	// TODO(barni@): Can we setup dynamic controller at main.go for CompositionReconciler instead of 1 per ExpanderReconciler
	r.Dynamic, err = dynamic.NewForConfig(r.Config)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(cr).
		Complete(r)
}
