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
	"sort"
	"strings"

	semver "github.com/Masterminds/semver/v3"
	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	"google.com/composition/pkg/containerexecutor/jobcontainerexecutor"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
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
	InputGVR      schema.GroupVersionResource
	Composition   types.NamespacedName
}

var planGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1alpha1",
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
	var compositionCR compositionv1alpha1.Composition
	if err := r.Client.Get(ctx, r.Composition, &compositionCR); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Composition Object")
			return ctrl.Result{}, err
		}
	}

	// Associate a plan object with this input CR
	var plancr compositionv1alpha1.Plan
	planNN := types.NamespacedName{Name: r.InputGVR.Resource + "-" + inputcr.GetName(), Namespace: inputcr.GetNamespace()}
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Plan Object")
			return ctrl.Result{}, err
		}
		// create a plan object
		plancr = compositionv1alpha1.Plan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      planNN.Name,
				Namespace: planNN.Namespace,
			},
			Spec: compositionv1alpha1.PlanSpec{
				Stages: map[string]compositionv1alpha1.Stage{},
			},
			Status: compositionv1alpha1.PlanStatus{
				Stages: map[string]compositionv1alpha1.StageStatus{},
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

	// Create a new status for comparison later
	newStatus := compositionv1alpha1.PlanStatus{
		Stages: map[string]compositionv1alpha1.StageStatus{},
	}

	// Try updating status before returning
	defer func() {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			nn := types.NamespacedName{Namespace: plancr.Namespace, Name: plancr.Name}
			err := r.Client.Get(ctx, nn, &plancr)
			if err != nil {
				return err
			}
			plancr.Status = *newStatus.DeepCopy()
			return r.Client.Status().Update(ctx, &plancr)
		})
		if err != nil {
			logger.Error(err, "unable to update Plan status")
		}
	}()

	expandersProcessed := []string{}
	for index, expander := range compositionCR.Spec.Expanders {
		// -------------------- FETCH VALUES ---------------------------
		// TODO(barni@): identify dependend variables and path where values need to be read from
		// Update the CRD_V's status to reflect these values
		if len(expander.ValuesFrom) != 0 {
			// Clear plan.WAITING Condition
			logger = loggerCR.WithName(expander.Name).WithName("Fetcher")
			logger.Info("Fetching Values")
			ff := NewFetcher(ctx, logger, r, &plancr, &inputcr, &expander)
			err := ff.Fetch()
			if err != nil {
				r.Recorder.Event(&inputcr, "Warning", "FetcherFailed", fmt.Sprintf("Error getting values for expander: %s", expander.Name))
				logger.Error(err, "Unable to fetch dependent valuesFrom ")
				// Inject plan.WAITING Condition
				newStatus.AppendWaitingCondition(expander.Name, err.Error(), "FetchValuesFromFailed")
				return ctrl.Result{}, err
			}
			err = ff.UpdatePlanCR()
			if err != nil {
				r.Recorder.Event(&inputcr, "Warning", "FetcherFailed", fmt.Sprintf("Error updating fetched values for expander: %s", expander.Name))
				// Inject plan.ERROR Condition
				newStatus.AppendErrorCondition(expander.Name, err.Error(), "UpdatingPlanFailed")
				return ctrl.Result{}, err
			}
			logger.Info("Fetched Values")
		}

		// ------------------- EXPANSION SECTION -----------------------
		logger = loggerCR.WithName(expander.Name).WithName("Expand")

		expanderVersion, expanderRegistry, err := r.getExpanderVersion(ctx, expander.Version, expander.Type)
		if err != nil {
			reason := "FailedGettingValidExpanderVersion"
			if apierrors.IsNotFound(err) {
				// The CR should be created before the specified expander can be used.
				reason = "MissingExpanderCR"
				logger.Error(err, "Failed to get the ExpanderVersionCR", "expander", expander.Type, "reason", reason)
			}
			if strings.Contains(err.Error(), "invalidExpanderVersion") {
				reason = "InvalidExpanderVersion"
				logger.Error(err, "Expander Version Invalid for", "expander", expander.Type, "version", expander.Version, "reason", reason)
			}
			newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
			return ctrl.Result{}, err
		}
		logger.Info("Get valid expander version", "expanderVersion", expanderVersion)

		jf := jobcontainerexecutor.NewJobFactory(ctx, logger, r.Client, r.InputGVK, r.InputGVR,
			r.Composition.Name, r.Composition.Namespace,
			&inputcr, expander.Name, expanderVersion, expander.Type, planNN.Name, expanderRegistry)

		// Create Expander Job and wait for the Job to complete
		logger.Info("Creating expander job")
		err = jf.Create()
		defer jf.CleanUp()
		if err != nil {
			// Reference for event API: Event(object runtime.Object, eventtype, reason, message string)
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Error creating job for name: %s", expander.Name))
			logger.Error(err, "Unable to create expander job")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, err.Error(), "ExpanderJobCreationFailed")
			return ctrl.Result{}, err
		}
		r.Recorder.Event(&inputcr, "Normal", "ExpansionStarted", fmt.Sprintf("Job Created for name: %s", expander.Name))
		logger.Info("Successfully created expander job")

		success, err := jf.Wait()
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Error waiting for Job for name: %s", expander.Name))
			logger.Error(err, "Failed waiting for expander job")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, err.Error(), "WaitingForExpanderFailed")
			return ctrl.Result{}, err
		}

		if !success {
			r.Recorder.Event(&inputcr, "Warning", "ExpansionFailed", fmt.Sprintf("Job failed for name: %s", expander.Name))
			logger.Info("Expander job completed but Failed")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, "Expander job completed but Failed", "ExpansionFailed")
			return ctrl.Result{}, fmt.Errorf("Expander Job Failed")
		}
		r.Recorder.Event(&inputcr, "Normal", "ExpansionSucceded", fmt.Sprintf("Job succeded for name: %s", expander.Name))
		logger.Info("Expander job Completed successfully")

		// ------------------- APPLIER SECTION -----------------------

		// Create Applier and wait for the Applier to complete
		logger = loggerCR.WithName(expander.Name).WithName("Apply")

		// Re-read the Plan CR to load the expanded manifests
		if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
			logger.Error(err, "unable to read Plan CR")
			return ctrl.Result{}, err
		}

		applier := NewApplier(ctx, logger, r, &plancr, &inputcr, &compositionCR, expander.Name)

		err = applier.Load() // Load Manifests
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed", fmt.Sprintf("error loading manifests for expander, name: %s", expander.Name))
			logger.Error(err, "Unable to Load manifests for applying")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, err.Error(), "FailedLoadingManifestsFromPlan")
			return ctrl.Result{}, err
		}
		logger.Info("Successfully loaded manifests for applying")
		if newStatus.Stages == nil {
			newStatus.Stages = map[string]compositionv1alpha1.StageStatus{}
		}
		newStatus.Stages[expander.Name] = compositionv1alpha1.StageStatus{ResourceCount: applier.Count()}

		prune := false
		if index == len(compositionCR.Spec.Expanders)-1 {
			prune = true
		}

		_, err = applier.Apply(oldAppliers, prune) // Apply Manifests
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed", fmt.Sprintf("error applying manifests for expander, name: %s", expander.Name))
			logger.Error(err, "Unable to apply manifests")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, err.Error(), "FailedApplyingManifests")
			return ctrl.Result{}, err
		}

		logger.Info("Successfully applied manifests")
		r.Recorder.Event(&inputcr, "Normal", "ResourcesApplied", fmt.Sprintf("All expanded resources were applied. name: %s", expander.Name))

		success, err = applier.Wait()
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Failed waiting for resources to be reconciled. name: %s", expander.Name))
			logger.Error(err, "Failed waiting for applied resources to reconcile")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander.Name, err.Error(), "FailedWaitingForAppliedResources")
			return ctrl.Result{}, err
		}

		oldAppliers = append(oldAppliers, applier)

		if !success {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Some resources are not healthy. name: %s", expander.Name))
			logger.Info("Applied succesfully but some resources did not become healthy")
			// Inject plan.Waiting Condition
			newStatus.AppendWaitingCondition(expander.Name, "Not all resources are healthy", "WaitingForAppliedResources")
			return ctrl.Result{}, fmt.Errorf("Some applied resources are not healthy")
		}

		expandersProcessed = append(expandersProcessed, expander.Name)
		r.Recorder.Event(&inputcr, "Normal", "ResourcesReconciled", fmt.Sprintf("All applied resources were reconciled. name: %s", expander.Name))
		logger.Info("Applied resources successfully.")
		// Inject plan.Ready Condition with list of expanders
		newStatus.ClearCondition(compositionv1alpha1.Ready)
		message := fmt.Sprintf("Processed stages: %s", strings.Join(expandersProcessed, ", "))
		newStatus.AppendCondition(compositionv1alpha1.Ready, metav1.ConditionFalse, message, "PendingStages")
	}

	// Inject plan.Ready Condition with list of expanders
	newStatus.ClearCondition(compositionv1alpha1.Ready)
	message := fmt.Sprintf("Processed stages: %s", strings.Join(expandersProcessed, ", "))
	newStatus.AppendCondition(compositionv1alpha1.Ready, metav1.ConditionTrue, message, "ProcessedAllStages")
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

func (r *ExpanderReconciler) getExpanderVersion(ctx context.Context, inputExpanderVersion string, expanderType string) (string, string, error) {
	logger := log.FromContext(ctx)

	var expanderVersionCR compositionv1alpha1.ExpanderVersion
	var expanderVersion, expanderRegistry string
	var availableVersions []string
	err := r.Client.Get(ctx, types.NamespacedName{Name: "composition-" + expanderType, Namespace: "composition-system"}, &expanderVersionCR)
	if err != nil {
		// The CR should be created before the specified expander can be used.
		return "", "", err
	}
	expanderRegistry = expanderVersionCR.Spec.ImageRegistry
	availableVersions = expanderVersionCR.Spec.ValidVersions
	SemVerVersions := make([]*semver.Version, len(availableVersions))
	for i, r := range availableVersions {
		v, err := semver.NewVersion(r)
		if err != nil {
			logger.Info("Error parsing version: %s", err)
		}
		SemVerVersions[i] = v
	}
	sort.Sort(semver.Collection(SemVerVersions))
	logger.Info("inputexpanderverions", "current", inputExpanderVersion)
	// Currently the version will default to latest if not set.
	if inputExpanderVersion != "latest" {
		// Verify if the input version is valid.
		semVersion, err := semver.NewVersion(inputExpanderVersion)
		if err != nil {
			return "", "", fmt.Errorf("invalidExpanderVersion")
		}
		if !isValidVersion(SemVerVersions, semVersion) {
			return "", "", fmt.Errorf("invalidExpanderVersion")
		}
		expanderVersion = inputExpanderVersion
	} else {
		// The existing semver package sort only supports pure numbers.
		expanderVersion = "v" + SemVerVersions[len(availableVersions)-1].String() + ".alpha"
	}
	return expanderVersion, expanderRegistry, nil
}

func isValidVersion(availableVersions []*semver.Version, version *semver.Version) bool {
	var low, high, mid int
	low = 0
	high = len(availableVersions) - 1
	for low <= high {
		mid = low + (high-low)/2

		// Check the version is present at mid
		if availableVersions[mid].Compare(version) == 0 {
			return true
		}
		// If the version is greater, ignore left half
		if availableVersions[mid].LessThan(version) {
			low = mid + 1
			// If the version is smaller, ignore right half
		} else {
			high = mid - 1
		}
	}
	return false
}
