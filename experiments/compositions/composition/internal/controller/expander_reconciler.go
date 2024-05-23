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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"golang.org/x/time/rate"
	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	"google.com/composition/pkg/containerexecutor/jobcontainerexecutor"
	pb "google.com/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ExpanderReconciler reconciles a expander object
type ExpanderReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	Recorder    record.EventRecorder
	RESTMapper  meta.RESTMapper
	Config      *rest.Config
	Dynamic     *dynamic.DynamicClient
	InputGVK    schema.GroupVersionKind
	InputGVR    schema.GroupVersionResource
	Composition types.NamespacedName
}

var planGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1alpha1",
	Kind:    "Plan",
}

var contextGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1alpha1",
	Kind:    "Context",
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

		value, ev, reason, err := r.getExpanderValue(ctx, expander.Version, expander.Type)
		if err != nil {
			logger.Error(err, "Error getting expander version", "expander", expander.Type,
				"version", expander.Version, "reason", reason)
			newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
			return ctrl.Result{}, err
		}

		logger.Info("Got valid expander value", "value", value)

		success := false
		planUpdated := false
		if ev.Spec.Type == compositionv1alpha1.ExpanderTypeJob {
			reason, err := r.runJob(ctx, logger, &inputcr, expander.Name, planNN.Name, value, ev.Spec.ImageRegistry)
			if err != nil {
				newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
				return ctrl.Result{}, err
			}
		} else {
			updated, reason, err := r.evaluateAndSavePlan(ctx, logger, &inputcr, expander, planNN, value)
			if err != nil {
				newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
				return ctrl.Result{}, err
			}
			planUpdated = updated
		}

		// ------------------- APPLIER SECTION -----------------------

		// Create Applier and wait for the Applier to complete
		logger = loggerCR.WithName(expander.Name).WithName("Apply")

		oldGeneration := plancr.GetGeneration()
		// Re-read the Plan CR to load the expanded manifests
		if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
			logger.Error(err, "unable to read Plan CR")
			return ctrl.Result{}, err
		}
		if planUpdated && oldGeneration == plancr.GetGeneration() {
			logger.Error(err, "Did not get the latest planCR. Will retry.")
			return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
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
	newStatus.InputGeneration = inputcr.GetGeneration()
	newStatus.Generation = plancr.GetGeneration()
	return ctrl.Result{}, nil
}

func (r *ExpanderReconciler) getExpanderValue(
	ctx context.Context, inputExpanderVersion string, expanderType string,
) (string, *compositionv1alpha1.ExpanderVersion, string, error) {
	logger := log.FromContext(ctx)

	value := ""
	var ev compositionv1alpha1.ExpanderVersion
	err := r.Client.Get(ctx,
		types.NamespacedName{
			Name:      "composition-" + expanderType,
			Namespace: "composition-system"},
		&ev)

	if err != nil {
		// The CR should be created before the specified expander can be used.
		logger.Error(err, "Failed to get the ExpanderVersionCR")
		if apierrors.IsNotFound(err) {
			return value, nil, "MissingExpanderCR", err
		} else {
			return value, nil, "ErrorGettingExpanderVersionCR", err
		}
	}

	if ev.Status.VersionMap == nil {
		return value, nil, "ErrorEmptyVersionMap", fmt.Errorf("ExpanderVersion .status.versionMap is empty")
	}

	logger.Info("input expander version", "current", inputExpanderVersion)
	value, ok := ev.Status.VersionMap[inputExpanderVersion]
	if !ok {
		return value, nil, "VersionNotFound", fmt.Errorf("%s version not found", inputExpanderVersion)
	}
	return value, &ev, "", nil
}

func (r *ExpanderReconciler) runJob(ctx context.Context, logger logr.Logger,
	cr *unstructured.Unstructured, expanderName, planName, value, registry string) (string, error) {
	jf := jobcontainerexecutor.NewJobFactory(ctx, logger, r.Client, r.InputGVK, r.InputGVR,
		r.Composition.Name, r.Composition.Namespace,
		cr, expanderName, value, planName, registry)

	// Create Expander Job and wait for the Job to complete
	logger.Info("Creating expander job")
	err := jf.Create()
	defer jf.CleanUp()
	if err != nil {
		// Reference for event API: Event(object runtime.Object, eventtype, reason, message string)
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Error creating job for name: %s", expanderName))
		logger.Error(err, "Unable to create expander job")
		return "ExpanderJobCreationFailed", err
	}
	r.Recorder.Event(cr, "Normal", "ExpansionStarted", fmt.Sprintf("Job Created for name: %s", expanderName))
	logger.Info("Successfully created expander job")

	success, err := jf.Wait()
	if err != nil {
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Error waiting for Job for name: %s", expanderName))
		logger.Error(err, "Failed waiting for expander job")
		return "WaitingForExpanderFailed", err
	}

	if !success {
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Job failed for name: %s", expanderName))
		logger.Info("Expander job completed but Failed")
		return "ExpansionFailed", fmt.Errorf("Expander Job Failed")
	}
	r.Recorder.Event(cr, "Normal", "ExpansionSucceded", fmt.Sprintf("Job succeded for name: %s", expanderName))
	logger.Info("Expander job Completed successfully")
	return "", nil
}

func (r *ExpanderReconciler) evaluateAndSavePlan(ctx context.Context, logger logr.Logger,
	cr *unstructured.Unstructured, expander compositionv1alpha1.Expander,
	planNN types.NamespacedName, grpcService string) (bool, string, error) {
	// Set up a connection to the server.
	updated := false

	conn, err := grpc.Dial(grpcService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err, "grpc dial failed: "+grpcService)
		return updated, "GRPCConnError", err
	}

	// read context in cr.namespace
	contextcr := unstructured.Unstructured{}
	contextcr.SetGroupVersionKind(contextGVK)
	contextNN := types.NamespacedName{Namespace: cr.GetNamespace(), Name: "context"}
	if err := r.Get(ctx, contextNN, &contextcr); err != nil {
		logger.Error(err, "unable to fetch Context CR", "context", contextNN)
		return updated, "GetContextFailed", err
	}
	contextBytes, err := json.Marshal(contextcr.Object)
	if err != nil {
		logger.Error(err, "failed to marshal Context Object")
		return updated, "MarshallContextFailed", err
	}

	// marshall facade cr
	facadeBytes, err := json.Marshal(cr.Object)
	if err != nil {
		logger.Error(err, "failed to marshall Facade Object")
		return updated, "MarshallFacadeFailed", err
	}

	// marshall expande config
	configBytes, err := json.Marshal(expander.Template)
	if err != nil {
		logger.Error(err, "failed to marshall Expander Config")
		return updated, "MarshallExpanderConfigFailed", err
	}

	expanderClient := pb.NewExpanderClient(conn)
	result, err := expanderClient.Evaluate(ctx,
		// TODO (barney-s) value (pass as parameter)
		&pb.EvaluateRequest{
			Config:   configBytes,
			Context:  contextBytes,
			Facade:   facadeBytes,
			Resource: r.InputGVR.Resource,
			Value:    []byte{},
		})
	if err != nil {
		logger.Error(err, "expander.Evaluate() Failed", "expander", expander.Name)
		return updated, "EvaluateError", err
	}
	if result.Status != pb.Status_SUCCESS {
		logger.Error(nil, "expander.Evaluate() Status is not Success", "expander", expander.Name, "status", result.Status)
		err = fmt.Errorf("Evaluate Failed: %s", result.Error.Message)
		return updated, "EvaluateStatusFailed", err
	}

	// Write to Plan object
	// Re-read the Plan CR to load the expanded manifests
	plancr := compositionv1alpha1.Plan{}
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		logger.Error(err, "unable to read Plan CR", "plan", planNN)
		return updated, "GetPlanFailed", err
	}

	if plancr.Spec.Stages == nil {
		plancr.Spec.Stages = map[string]compositionv1alpha1.Stage{}
		updated = true
	}
	oldValue := ""
	if _, ok := plancr.Spec.Stages[expander.Name]; !ok {
		updated = true
	} else {
		oldValue = plancr.Spec.Stages[expander.Name].Manifest
	}

	if result.Type == pb.ResultType_MANIFESTS {
		s, err := strconv.Unquote(string(result.Manifests))
		if err != nil {
			logger.Error(err, "unable to unquote grpc response")
			return updated, "UnquoteResponseFailed", err
		}
		plancr.Spec.Stages[expander.Name] = compositionv1alpha1.Stage{
			Manifest: s,
		}
	} else {
		s, err := strconv.Unquote(string(result.Values))
		if err != nil {
			logger.Error(err, "unable to unquote grpc response")
			return updated, "UnquoteResponseFailed", err
		}
		plancr.Spec.Stages[expander.Name] = compositionv1alpha1.Stage{
			Manifest: s,
		}
	}

	if plancr.Spec.Stages[expander.Name].Manifest != oldValue {
		updated = true
	}

	err = r.Client.Update(ctx, &plancr)
	if err != nil {
		logger.Error(err, "error updating plan", "plan", planNN)
		return updated, "UpdatePlanFailed", err
	}

	return updated, "", nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExpanderReconciler) SetupWithManager(mgr ctrl.Manager, cr *unstructured.Unstructured) error {
	var err error
	// TODO(barni@): Can we setup dynamic controller at main.go for CompositionReconciler instead of 1 per ExpanderReconciler
	r.Dynamic, err = dynamic.NewForConfig(r.Config)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	ratelimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond, 120*time.Second),
		// 40 qps, 400 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(40), 400)},
	)

	return ctrl.NewControllerManagedBy(mgr).
		For(cr).
		WithOptions(controller.Options{RateLimiter: ratelimiter}).
		Complete(r)
}
