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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/deletiondefender"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apikeys"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/iam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/resourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/gsakeysecretgenerator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/unmanageddetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/kccfeatureflags"
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

const controllerName = "registration-controller"
const serviceAccountKeyAPIGroup = "iam.cnrm.cloud.google.com"
const serviceAccountKeyKind = "IAMServiceAccountKey"

var logger = crlog.Log.WithName(controllerName)

// Add creates a new registration Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager, p *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, dclConfig *dcl.Config, dclConverter *conversion.Converter, regFunc registrationFunc, defaulters []k8s.Defaulter) error {
	r := &ReconcileRegistration{
		Client:           mgr.GetClient(),
		provider:         p,
		smLoader:         smLoader,
		dclConfig:        dclConfig,
		dclConverter:     dclConverter,
		mgr:              mgr,
		controllers:      make(map[string]map[string]controllerContext),
		registrationFunc: regFunc,
		defaulters:       defaulters,
	}
	c, err := crcontroller.New(controllerName, mgr,
		crcontroller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
		})
	if err != nil {
		return err
	}
	return c.Watch(source.Kind(mgr.GetCache(), &apiextensions.CustomResourceDefinition{}), &handler.EnqueueRequestForObject{}, ManagedByKCCPredicate{})
}

var _ reconcile.Reconciler = &ReconcileRegistration{}

// ReconcileRegistration reconciles a CRD owned by KCC
type ReconcileRegistration struct {
	client.Client
	provider         *tfschema.Provider
	smLoader         *servicemappingloader.ServiceMappingLoader
	dclConfig        *dcl.Config
	dclConverter     *conversion.Converter
	mgr              manager.Manager
	controllers      map[string]map[string]controllerContext
	registrationFunc registrationFunc
	defaulters       []k8s.Defaulter
	mu               sync.Mutex
}

type controllerContext struct {
	registered    bool
	schemaUpdater k8s.SchemaReferenceUpdater
}

// registrationFunc is the function that handles the registration of a controller for the given CRD and returns an interface to update its schema reference.
type registrationFunc func(*ReconcileRegistration, *apiextensions.CustomResourceDefinition, schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error)

func (r *ReconcileRegistration) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
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

func RegisterDefaultController(config *controller.Config) registrationFunc { //nolint:revive
	return func(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
		return registerDefaultController(r, config, crd, gvk)
	}
}

func registerDefaultController(r *ReconcileRegistration, config *controller.Config, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}

	var schemaUpdater k8s.SchemaReferenceUpdater

	if kccfeatureflags.UseDirectReconciler(gvk.GroupKind()) {
		switch gvk.GroupKind() {
		case schema.GroupKind{Group: "apikeys.cnrm.cloud.google.com", Kind: "APIKeysKey"}:
			if err := apikeys.AddKeyReconciler(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:
			if err := iam.AddServiceAccountController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetwork"}:
			if err := compute.AddNetworkController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSubnetwork"}:
			if err := compute.AddSubnetworkController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagKey"}:
			if err := resourcemanager.AddTagKeyController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagValue"}:
			if err := resourcemanager.AddTagValueController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagBinding"}:
			if err := resourcemanager.AddTagBindingController(r.mgr, config); err != nil {
				return nil, err
			}
			return schemaUpdater, nil

		default:
			return nil, fmt.Errorf("requested direct reconciler for %v, but it is not supported", gvk.GroupKind())
		}
	}

	// Depending on which resource it is, we need to register a different controller.
	switch gvk.Kind {
	case "IAMPolicy":
		if err := policy.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, r.defaulters); err != nil {
			return nil, err
		}
	case "IAMPartialPolicy":
		if err := partialpolicy.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, r.defaulters); err != nil {
			return nil, err
		}
	case "IAMPolicyMember":
		if err := policymember.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, r.defaulters); err != nil {
			return nil, err
		}
	case "IAMAuditConfig":
		if err := auditconfig.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig, r.defaulters); err != nil {
			return nil, err
		}

	default:
		// register controllers for dcl-based CRDs
		if val, ok := crd.Labels[k8s.DCL2CRDLabel]; ok && val == "true" {
			su, err := dclcontroller.Add(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader, r.defaulters)
			if err != nil {
				return nil, fmt.Errorf("error adding dcl controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
			}
			return su, nil
		}
		// register controllers for tf-based CRDs
		if val, ok := crd.Labels[crdgeneration.TF2CRDLabel]; !ok || val != "true" {
			logger.Info("unrecognized CRD; skipping controller registration", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
			return nil, nil
		}
		su, err := tf.Add(r.mgr, crd, r.provider, r.smLoader, r.defaulters)
		if err != nil {
			return nil, fmt.Errorf("error adding terraform controller for %v to a manager: %w", crd.Spec.Names.Kind, err)
		}
		schemaUpdater = su
		// register the controller to automatically create secrets for GSA keys
		if isServiceAccountKeyCRD(crd) {
			logger.Info("registering the GSA-Key-to-Secret generation controller")
			if err := gsakeysecretgenerator.Add(r.mgr, crd); err != nil {
				return nil, fmt.Errorf("error adding the gsa-to-secret generator for %v to a manager: %w", crd.Spec.Names.Kind, err)
			}
		}
	}
	return schemaUpdater, nil
}

func RegisterDeletionDefenderController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, _ schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}
	if err := deletiondefender.Add(r.mgr, crd); err != nil {
		return nil, fmt.Errorf("error registering deletion defender controller for '%v': %w", crd.GetName(), err)
	}
	return nil, nil
}

func RegisterUnmanagedDetectorController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, _ schema.GroupVersionKind) (k8s.SchemaReferenceUpdater, error) {
	if _, ok := k8s.IgnoredKindList[crd.Spec.Names.Kind]; ok {
		return nil, nil
	}
	if err := unmanageddetector.Add(r.mgr, crd); err != nil {
		return nil, fmt.Errorf("error registering unmanaged detector controller for '%v': %w", crd.GetName(), err)
	}
	return nil, nil
}
