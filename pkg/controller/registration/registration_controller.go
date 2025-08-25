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
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	tf "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/unmanageddetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crcontroller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const serviceAccountKeyAPIGroup = "iam.cnrm.cloud.google.com"
const serviceAccountKeyKind = "IAMServiceAccountKey"

type RegistrationControllerOptions struct {
	ControllerName            string
	ControllerConfig          *config.ControllerConfig
	ResourcesControllerConfig *resourceconfig.ResourcesControllerMap
}

// AddDefaultControllers creates the registration controller with the default controller factory,
// this will dynamically create the default controllers for each CRD.
func AddDefaultControllers(ctx context.Context, mgr manager.Manager, rd *controller.Deps, opts RegistrationControllerOptions) error {
	if err := add(mgr, rd,
		registerDefaultControllers(ctx, opts.ControllerConfig), opts); err != nil {
		return fmt.Errorf("error adding default registration controller: %w", err)
	}
	return nil
}

// AddDeletionDefender creates the registration controller with the deletion-defender factory,
// this will dynamically create the deletion-defender controller bound to each CRD.
func AddDeletionDefender(mgr manager.Manager, rd *controller.Deps) error {
	opt := RegistrationControllerOptions{
		ControllerName: "deletion-defender-registration-controller",
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
		ControllerName: "unmanaged-detector-registration-controller",
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
		Client:            mgr.GetClient(),
		provider:          rd.TFProvider,
		smLoader:          rd.TFLoader,
		dclConfig:         rd.DCLConfig,
		dclConverter:      rd.DCLConverter,
		mgr:               mgr,
		controllers:       make(map[string]map[string]controllerContext),
		registrationFunc:  regFunc,
		defaulters:        rd.Defaulters,
		jitterGenerator:   rd.JitterGen,
		dependencyTracker: rd.DependencyTracker,
	}

	r.controllerSelector = resourceconfig.NewControllerSelector(opts.ResourcesControllerConfig)
	c, err := crcontroller.New(opts.ControllerName, mgr,
		crcontroller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
		})
	if err != nil {
		return err
	}
	// return c.Watch(source.Kind(mgr.GetCache(), &apiextensions.CustomResourceDefinition{},
	// 	&handler.TypedEnqueueRequestForObject[*apiextensions.CustomResourceDefinition]{}, ManagedByKCCPredicate{}))
	return c.Watch(source.Kind(mgr.GetCache(), &apiextensions.CustomResourceDefinition{},
		&handler.TypedEnqueueRequestForObject[*apiextensions.CustomResourceDefinition]{}))
}

var _ reconcile.Reconciler = &ReconcileRegistration{}

// ReconcileRegistration reconciles a CRD owned by KCC
type ReconcileRegistration struct {
	client.Client
	provider           *tfschema.Provider
	smLoader           *servicemappingloader.ServiceMappingLoader
	dclConfig          *dcl.Config
	dclConverter       *conversion.Converter
	mgr                manager.Manager
	controllers        map[string]map[string]controllerContext
	registrationFunc   registrationFunc
	defaulters         []k8s.Defaulter
	jitterGenerator    jitter.Generator
	dependencyTracker  *gcpwatch.DependencyTracker
	controllerSelector *resourceconfig.ControllerSelector

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

	cds.DependencyTracker = r.dependencyTracker

	var schemaUpdater k8s.SchemaReferenceUpdater
	ccc := &v1beta1.ConfigConnectorContext{}
	if err := r.Client.Get(ctx, client.ObjectKey{Name: v1beta1.ConfigConnectorContextAllowedName}, ccc); err != nil {
		logger.Info("error getting configconnectorcontext: %w", err)
	}
	controllerType := r.controllerSelector.SelectController(gvk, ccc)
	switch controllerType {
	case k8s.ReconcilerTypeDirect:
		groupKind := gvk.GroupKind()

		model, err := registry.GetModel(groupKind)
		if err != nil {
			return nil, fmt.Errorf("error getting model for %v: %w", groupKind, err)
		}

		if err := directbase.AddController(r.mgr, gvk, model,
			directbase.Deps{
				JitterGenerator: r.jitterGenerator,
				Defaulters:      r.defaulters,
				// for iam controllers
				IAMAdapterDeps: &directbase.IAMAdapterDeps{
					KubeClient: r.Client,
					ControllerDeps: &controller.Deps{
						TFProvider:   r.provider,
						TFLoader:     r.smLoader,
						DCLConfig:    r.dclConfig,
						DCLConverter: r.dclConverter,
					},
				},
			}); err != nil {
			return nil, fmt.Errorf("error adding direct controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
		}
	case k8s.ReconcilerTypeTerraform:
		var err error
		schemaUpdater, err = tf.Add(r.mgr, crd, r.provider, r.smLoader, r.defaulters, r.jitterGenerator, nil)
		if err != nil {
			return nil, fmt.Errorf("error adding terraform controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
		}
	case k8s.ReconcilerTypeDCL:
		var err error
		schemaUpdater, err = dclcontroller.Add(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, r.defaulters, r.jitterGenerator, nil)
		if err != nil {
			return nil, fmt.Errorf("error adding dcl controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
		}
	default:
		logger.Info("no controller configured for resource", "group", gvk.Group, "kind", gvk.Kind)
	}

	// Depending on which resource it is, we need to register a different controller.
	switch gvk.Kind {
	case "IAMPolicy":
		if err := policy.Add(r.mgr, &cds); err != nil {
			return nil, err
		}
	case "IAMPartialPolicy":
		if err := partialpolicy.Add(r.mgr, &cds); err != nil {
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

func registerUnmanagedDetectorController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, _ schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	ctx := context.TODO()

	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}
	if err := unmanageddetector.Add(ctx, r.mgr, crd); err != nil {
		return nil, fmt.Errorf("error registering unmanaged detector controller for '%v': %w", crd.GetName(), err)
	}
	return nil, nil
}
