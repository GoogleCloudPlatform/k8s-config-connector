// Copyright 2022 Google LLC
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

package registration

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/deletiondefender"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/gsakeysecretgenerator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/unmanageddetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crcontroller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const serviceAccountKeyAPIGroup = "iam.cnrm.cloud.google.com"
const serviceAccountKeyKind = "IAMServiceAccountKey"

type RegistrationControllerOptions struct {
	ControllerName     string
	SkipNameValidation bool
}

// SkipControllerNameValidation allows skipping the controller name validation
// in controller-runtime. This is useful when running multiple managers in the same process
// (e.g. in preview mode or tests) to avoid "controller name already exists" errors.
var SkipControllerNameValidation bool

// AddDefaultControllers creates the registration controller with the default controller factory,
// this will dynamically create the default controllers for each CRD.
func AddDefaultControllers(ctx context.Context, mgr manager.Manager, rd *controller.Deps, controllerConfig *config.ControllerConfig) error {
	opt := RegistrationControllerOptions{
		ControllerName:     "registration-controller",
		SkipNameValidation: SkipControllerNameValidation,
	}
	if err := add(mgr, rd,
		registerDefaultControllers(ctx, controllerConfig), opt); err != nil {
		return fmt.Errorf("error adding default registration controller: %w", err)
	}
	return nil
}

// AddDeletionDefender creates the registration controller with the deletion-defender factory,
// this will dynamically create the deletion-defender controller bound to each CRD.
func AddDeletionDefender(mgr manager.Manager, rd *controller.Deps) error {
	opt := RegistrationControllerOptions{
		ControllerName:     "deletion-defender-registration-controller",
		SkipNameValidation: SkipControllerNameValidation,
	}

	if err := add(mgr, &controller.Deps{}, registerDeletionDefenderController, opt); err != nil {
		return fmt.Errorf("error adding deletion-defender registration controller: %w", err)
	}
	return nil
}

// AddUnmanagedDetector creates the registration controller with the unmanaged-detector factory,
// this will dynamically create the unmanaged-detector controller bound to each CRD.
func AddUnmanagedDetector(mgr manager.Manager, rd *controller.Deps) error {
	opt := RegistrationControllerOptions{
		ControllerName:     "unmanaged-detector-registration-controller",
		SkipNameValidation: SkipControllerNameValidation,
	}

	registerUnmanagedDetectorController := func(r *ReconcileRegistration, _ *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
		ctx := context.TODO()

		if _, ok := k8s.IgnoredKindList[gvk.Kind]; ok {
			return nil, nil
		}
		if err := unmanageddetector.Add(ctx, r.mgr, gvk); err != nil {
			return nil, fmt.Errorf("error registering unmanaged detector controller for '%v': %w", gvk.Kind, err)
		}
		return nil, nil
	}

	if err := add(mgr, &controller.Deps{}, registerUnmanagedDetectorController, opt); err != nil {
		return fmt.Errorf("error adding unmanaged-detector registration controller: %w", err)
	}
	return nil
}

// add creates a new registration Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func add(mgr manager.Manager, rd *controller.Deps, regFunc registrationFunc, opts RegistrationControllerOptions) error {
	if rd.JitterGen == nil {
		var dclML metadata.ServiceMetadataLoader
		if rd.DCLConverter != nil {
			dclML = rd.DCLConverter.MetadataLoader
		}
		rd.JitterGen = jitter.NewDefaultGenerator(rd.TFLoader, dclML)
	}

	r := &ReconcileRegistration{
		Client:                     mgr.GetClient(),
		provider:                   rd.TFProvider,
		smLoader:                   rd.TFLoader,
		dclConfig:                  rd.DCLConfig,
		dclConverter:               rd.DCLConverter,
		mgr:                        mgr,
		controllers:                make(map[string]map[string]controllerContext),
		registrationFunc:           regFunc,
		defaulters:                 rd.Defaulters,
		jitterGenerator:            rd.JitterGen,
		dependencyTracker:          rd.DependencyTracker,
		reconcilers:                make(map[schema.GroupVersionKind]*parent.Reconcilers),
		immediateReconcileRequests: make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize),
		resourceWatcherRoutines:    semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines),
	}
	c, err := crcontroller.New(opts.ControllerName, mgr,
		crcontroller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
			SkipNameValidation:      &opts.SkipNameValidation,
		})
	if err != nil {
		return err
	}
	// return c.Watch(source.Kind(mgr.GetCache(), &apiextensions.CustomResourceDefinition{},
	// 	&handler.TypedEnqueueRequestForObject[*apiextensions.CustomResourceDefinition]{}, ManagedByKCCPredicate{}))
	return c.Watch(source.Kind(mgr.GetCache(), &apiextensions.CustomResourceDefinition{},
		&handler.TypedEnqueueRequestForObject[*apiextensions.CustomResourceDefinition]{}, ManagedByKCCPredicate[*apiextensions.CustomResourceDefinition]{}))
}

var _ reconcile.Reconciler = &ReconcileRegistration{}

// ReconcileRegistration reconciles a CRD owned by KCC
type ReconcileRegistration struct {
	client.Client
	provider          *tfschema.Provider
	smLoader          *servicemappingloader.ServiceMappingLoader
	dclConfig         *dcl.Config
	dclConverter      *conversion.Converter
	mgr               manager.Manager
	controllers       map[string]map[string]controllerContext
	registrationFunc  registrationFunc
	defaulters        []k8s.Defaulter
	jitterGenerator   jitter.Generator
	dependencyTracker *gcpwatch.DependencyTracker
	reconcilers       map[schema.GroupVersionKind]*parent.Reconcilers

	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies

	mu sync.Mutex
}

type controllerContext struct {
	registered    bool
	schemaUpdater k8s.SchemaReferenceUpdater
}

// registrationFunc is the function that handles the registration of a controller for the given CRD and returns an interface to update its schema reference.
type registrationFunc func(*ReconcileRegistration, *apiextensions.CustomResourceDefinition, schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error)

func (r *ReconcileRegistration) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	logger := crlog.FromContext(ctx)

	// Fetch the TypeProvider tp
	crd := &apiextensions.CustomResourceDefinition{}
	err := r.Get(ctx, request.NamespacedName, crd)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	logger.V(2).Info("Waiting to obtain lock...", "kind", crd.Spec.Names.Kind)
	start := time.Now()
	r.mu.Lock()
	logger.V(2).Info("Obtained lock", "kind", crd.Spec.Names.Kind, "elapsed (μs)", time.Since(start).Microseconds())
	defer func() {
		logger.V(2).Info("Releasing lock...", "kind", crd.Spec.Names.Kind)
		r.mu.Unlock()
	}()
	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: k8s.GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	if kindMapForGroup, exists := r.controllers[gvk.Group]; exists {
		if kindMapForGroup[gvk.Kind].registered {
			logger.Info("controller already registered for kind in API group", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
			if kindMapForGroup[gvk.Kind].schemaUpdater != nil {
				logger.Info("updating schema for controller", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
				if err := kindMapForGroup[gvk.Kind].schemaUpdater.UpdateSchema(crd); err != nil {
					logger.Info("error updating schema for controller", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
				}
			}
			return reconcile.Result{}, nil
		}
	} else {
		r.controllers[gvk.Group] = make(map[string]controllerContext)
	}

	schemaUpdater, err := r.registrationFunc(r, crd, gvk)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error registering controller: %w", err)
	}

	r.controllers[gvk.Group][gvk.Kind] = controllerContext{registered: true, schemaUpdater: schemaUpdater}
	return reconcile.Result{}, nil
}

func isServiceAccountKeyCRD(crd *apiextensions.CustomResourceDefinition) bool {
	return crd.Spec.Group == serviceAccountKeyAPIGroup && crd.Spec.Names.Kind == serviceAccountKeyKind
}

func registerDefaultControllers(ctx context.Context, config *config.ControllerConfig) registrationFunc { //nolint:revive
	return func(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
		return registerDefaultController(ctx, r, config, crd, gvk)
	}
}

func registerDefaultController(ctx context.Context, r *ReconcileRegistration, config *config.ControllerConfig, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	logger := crlog.FromContext(ctx)
	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}
	cds := controller.Deps{
		TFProvider:   r.provider,
		TFLoader:     r.smLoader,
		DCLConfig:    r.dclConfig,
		DCLConverter: r.dclConverter,
		JitterGen:    r.jitterGenerator,
		Defaulters:   r.defaulters,
		//DependencyTracker: r.dependencyTracker,
	}

	// todo acpana house in KCC mgr flag
	v := os.Getenv("KCC_RECONCILE_FLAG_GATE")
	if v == "USE_DEPENDENCY_TRACKER" {
		cds.DependencyTracker = r.dependencyTracker
	}

	var schemaUpdater k8s.SchemaReferenceUpdater

	// Depending on which resource it is, we need to register a different controller.
	switch gvk.Kind {
	case "IAMPolicy":
		if err := policy.Add(r.mgr, &cds); err != nil {
			return nil, err
		}
	case "IAMPolicyMember":
		if err := policymember.Add(r.mgr, &cds); err != nil {
			return nil, err
		}
	case "IAMAuditConfig":
		if err := auditconfig.Add(r.mgr, &cds); err != nil {
			return nil, err
		}

	default:
		// register the controller to automatically create secrets for GSA keys
		if isServiceAccountKeyCRD(crd) {
			logger.Info("registering the GSA-Key-to-Secret generation controller")
			if err := gsakeysecretgenerator.Add(r.mgr, crd, &controller.Deps{JitterGen: r.jitterGenerator}); err != nil {
				return nil, fmt.Errorf("error adding the gsa-to-secret generator for %v to a manager: %w", crd.Spec.Names.Kind, err)
			}
		}

		// register the parent controller for all supported resources.
		if config, err := resourceconfig.LoadConfig().GetControllersForGVK(gvk); err != nil {
			logger.Error(fmt.Errorf("unrecognized CRD: %v", crd.Spec.Names.Kind), "skipping controller registration", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
			return nil, nil
		} else {
			reconcilers := &parent.Reconcilers{}
			var err error
			for _, reconcilerType := range config.SupportedControllers {
				switch reconcilerType {
				case k8s.ReconcilerTypeIAMPartialPolicy:
					reconciler, err := partialpolicy.NewReconciler(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, nil, nil, r.defaulters, r.jitterGenerator, cds.DependencyTracker)
					if err != nil {
						return nil, err
					}
					reconcilers.Custom = &parent.CustomReconciler{
						Type:       k8s.ReconcilerTypeIAMPartialPolicy,
						Reconciler: reconciler,
					}
				case k8s.ReconcilerTypeTerraform:
					reconcilers.TF, err = tf.NewReconciler(r.mgr, crd, r.provider, r.smLoader, nil, nil, r.defaulters, r.jitterGenerator)
					if err != nil {
						return nil, fmt.Errorf("error creating new terraform reconciler: %w", err)
					}
				case k8s.ReconcilerTypeDCL:
					reconcilers.DCL, err = dclcontroller.NewReconciler(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, nil, nil, r.defaulters, r.jitterGenerator)
					if err != nil {
						return nil, fmt.Errorf("error creating new dcl reconciler: %w", err)
					}
				case k8s.ReconcilerTypeDirect:
					model, err := registry.GetModel(gvk.GroupKind())
					if err != nil {
						return nil, fmt.Errorf("error getting model for gvk %v: %w", gvk, err)
					}
					deps := directbase.Deps{
						Defaulters:      r.defaulters,
						JitterGenerator: r.jitterGenerator,
						IAMAdapterDeps: &directbase.IAMAdapterDeps{
							KubeClient: r.Client,
							ControllerDeps: &controller.Deps{
								TFProvider:   r.provider,
								TFLoader:     r.smLoader,
								DCLConfig:    r.dclConfig,
								DCLConverter: r.dclConverter,
							},
						},
					}
					reconcilers.Direct, err = directbase.NewReconciler(r.mgr, nil, nil, gvk, model, deps)
					if err != nil {
						return nil, fmt.Errorf("error creating new direct reconciler: %w", err)
					}
				}
			}
			if SkipControllerNameValidation {
				parent.SkipControllerNameValidation = SkipControllerNameValidation	
			}
			r.reconcilers[gvk] = reconcilers
			if err := parent.Add(r.mgr, gvk, reconcilers); err != nil {
				return nil, fmt.Errorf("error adding parent controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
			}
		}
	}
	return schemaUpdater, nil
}

func registerDeletionDefenderController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, _ schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}
	if err := deletiondefender.Add(r.mgr, crd); err != nil {
		return nil, fmt.Errorf("error registering deletion defender controller for '%v': %w", crd.GetName(), err)
	}
	return nil, nil
}
