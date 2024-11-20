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
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	corev1 "k8s.io/api/core/v1"
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
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var FacadeControllers sync.Map
var errDuplicate = fmt.Errorf("duplicate composition")

// CompositionReconciler reconciles a Composition object
type CompositionReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	Recorder        record.EventRecorder
	mgr             ctrl.Manager
	handoffChannels map[schema.GroupVersionKind]chan event.GenericEvent
}

// TODO: To simplify preview for customers, grant superuser to the composition controller. This should be revisited going forward.
//+kubebuilder:rbac:groups=*,resources=*,verbs=*
//+kubebuilder:rbac:groups=composition.google.com,resources=compositions;contexts;expanderversions;facades;getterconfigurations;plans,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=composition.google.com,resources=compositions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=composition.google.com,resources=compositions/finalizers,verbs=update
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=create;get;list;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=facade.facade,resources=*,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;create;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;create;patch;delete
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=create;get;patch;list;delete
//+kubebuilder:rbac:groups="batch",resources=jobs,verbs=create;get;patch;list;delete
//+kubebuilder:rbac:groups=facade.compositions.google.com,resources=*,verbs=get;list;patch;update;watch;create;delete
//+kubebuilder:rbac:groups=facade.compositions.google.com,resources=*/status,verbs=get;update

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

	composition.Status.ClearCondition(compositionv1alpha1.Error)
	logger.Info("Validating Composition object")
	if !composition.Validate() {
		logger.Info("Validation Failed")
		return ctrl.Result{}, fmt.Errorf("Validation failed")
	}

	logger.Info("Validating expander configs")
	if err := r.validateExpanders(ctx, logger, &composition); err != nil {
		logger.Info("expander config validation failed")
		return ctrl.Result{}, err
	}

	logger.Info("Processing Composition object")
	if err := r.processComposition(ctx, &composition, logger); err != nil {
		// Don't return an error to avoid requeuing a request that we know can't succeed.
		if errors.Is(err, errDuplicate) {
			return ctrl.Result{}, nil
		}
		logger.Info("Error processing Composition")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CompositionReconciler) validateExpanders(
	ctx context.Context, logger logr.Logger, c *compositionv1alpha1.Composition,
) error {
	errorStages := []string{}
	if c.Status.Stages == nil {
		c.Status.Stages = make(map[string]compositionv1alpha1.StageValidationStatus)
	}
	for _, expander := range c.Spec.Expanders {
		uri, ev, reason, err := r.getExpanderValue(ctx, expander.Version, expander.Type)
		if err != nil {
			logger.Error(err, "Error getting ExpanderVersion", "type", expander.Type, "version", expander.Version)
			errorStages = append(errorStages, expander.Name)
			c.Status.Stages[expander.Name] = compositionv1alpha1.StageValidationStatus{
				ValidationStatus: compositionv1alpha1.ValidationStatusError,
				Reason:           reason,
				Message:          err.Error(),
			}
			// Try the next expander
			continue
		}
		logger.Info("Got valid expander uri", "uri", uri)

		// We dont have validate for Job type expander
		if ev.Spec.Type != compositionv1alpha1.ExpanderTypeGRPC {
			c.Status.Stages[expander.Name] = compositionv1alpha1.StageValidationStatus{
				ValidationStatus: compositionv1alpha1.ValidationStatusUnknown,
				Message:          "expander type does not implement validation",
				Reason:           "NoValidationSupport",
			}
		} else {
			reason, err := r.validateExpanderConfig(ctx, logger, expander, ev, uri)
			if err != nil {
				logger.Error(err, "Validating config failed", "type", expander.Type, "version", expander.Version)
				errorStages = append(errorStages, expander.Name)
				c.Status.Stages[expander.Name] = compositionv1alpha1.StageValidationStatus{
					ValidationStatus: compositionv1alpha1.ValidationStatusFailed,
					Reason:           reason,
					Message:          err.Error(),
				}
				// Try the next expander
				continue
			}

			c.Status.Stages[expander.Name] = compositionv1alpha1.StageValidationStatus{
				ValidationStatus: compositionv1alpha1.ValidationStatusSuccess,
				Reason:           "ValidationPassed",
				Message:          "",
			}
		}
	}

	if len(errorStages) != 0 {
		message := fmt.Sprintf("Validating failed for stages: %s", strings.Join(errorStages, ", "))
		c.Status.Conditions = append(c.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            message,
			Reason:             "ValidationFailed",
			Type:               string(compositionv1alpha1.Error),
			Status:             metav1.ConditionTrue,
		})
		r.Recorder.Event(c, "Warning", "ValidationFailed", message)
		return fmt.Errorf("Validation Failed")
	}

	return nil
}
func (r *CompositionReconciler) validateExpanderConfig(ctx context.Context, logger logr.Logger,
	expander compositionv1alpha1.Expander, ev *compositionv1alpha1.ExpanderVersion, grpcService string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(grpcService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err, "grpc dial failed: "+grpcService)
		return "GRPCConnError", err
	}

	// marshall expander config
	// read Expander config from  in cr.namespace
	var configBytes []byte
	if expander.Reference != nil {
		expanderconfigcr := unstructured.Unstructured{}
		expanderconfigcr.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   ev.Spec.Config.Group,
			Version: ev.Spec.Config.Version,
			Kind:    ev.Spec.Config.Kind,
		})
		expanderconfigNN := types.NamespacedName{Namespace: expander.Reference.Namespace, Name: expander.Reference.Name}
		if err := r.Get(ctx, expanderconfigNN, &expanderconfigcr); err != nil {
			logger.Error(err, "unable to fetch ExpanderConfig CR", "expander config", expanderconfigNN)
			return "GetExpanderConfigFailed", err
		}
		configBytes, err = json.Marshal(expanderconfigcr.Object)
		if err != nil {
			logger.Error(err, "failed to marshal ExpanderConfig Object")
			return "MarshallExpanderConfigFailed", err
		}
	} else {
		// TODO check if json.Marshall is escaping quotes
		// Also causes > to be replaced unicode 'if loop.index \u003e 1'
		err = nil
		//configBytes, err = json.Marshal(expander.Template)
		configBytes = []byte(expander.Template)
		if err != nil {
			logger.Error(err, "failed to marshall Expander template")
			return "MarshallExpanderTemplateFailed", err
		}
	}

	expanderClient := pb.NewExpanderClient(conn)
	result, err := expanderClient.Validate(ctx,
		&pb.ValidateRequest{
			Config: configBytes,
		})
	if err != nil {
		logger.Error(err, "expander.Validate() Failed", "expander", expander.Name)
		return "ValidateError", err
	}
	if result.Status != pb.Status_SUCCESS {
		logger.Error(nil, "expander.Validate() Status is not Success", "expander",
			expander.Name, "status", result.Status, "message", result.Error.Message)
		err = fmt.Errorf(result.Error.Message)
		return "ValidateStatusFailed", err
	}
	return "", nil
}

func (r *CompositionReconciler) getExpanderValue(
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
	cNN := types.NamespacedName{Namespace: c.Namespace, Name: c.Name}
	nn, loaded := FacadeControllers.LoadOrStore(gvk, cNN)
	if loaded {
		existingNN := nn.(types.NamespacedName)
		if existingNN.Namespace != c.Namespace || existingNN.Name != c.Name {
			msg := fmt.Sprintf("Failed to apply composition %v over existing composition named %v for GVK %v", cNN, nn, gvk)
			logger.Error(errDuplicate, msg)
			c.Status.Conditions = append(c.Status.Conditions, metav1.Condition{
				LastTransitionTime: metav1.Now(),
				Message:            msg,
				Reason:             "DuplicateForGVK",
				Type:               string(compositionv1alpha1.Error),
				Status:             metav1.ConditionTrue,
			})
			return errDuplicate
		}
		logger.Info("Sending event to handoff channel")
		r.handoffChannels[gvk] <- event.GenericEvent{
			Object: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      c.Name,
					Namespace: c.Namespace,
				},
			},
		}

		// Reconciler already exists nothing to be done
		logger.Info("Reconciler already exists for InputAPI CRD")
		return nil
	}

	logger.Info("Starting Reconciler for InputAPI CRD")
	r.handoffChannels[gvk] = make(chan event.GenericEvent)
	expanderController := &ExpanderReconciler{
		Client:                    r.Client,
		Recorder:                  r.mgr.GetEventRecorderFor(crd.Spec.Names.Plural + "-expander"),
		Scheme:                    r.Scheme,
		InputGVK:                  gvk,
		Composition:               types.NamespacedName{Name: c.Name, Namespace: c.Namespace},
		InputGVR:                  gvk.GroupVersion().WithResource(crd.Spec.Names.Plural),
		RESTMapper:                r.mgr.GetRESTMapper(),
		Config:                    r.mgr.GetConfig(),
		CompositionChangedWatcher: r.handoffChannels[gvk],
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
	r.handoffChannels = make(map[schema.GroupVersionKind]chan event.GenericEvent)
	return ctrl.NewControllerManagedBy(mgr).
		For(&compositionv1alpha1.Composition{}).
		Complete(r)
}
