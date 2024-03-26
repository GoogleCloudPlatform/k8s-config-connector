/*
Copyright 2019 The Kubernetes Authors.

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

package declarative

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

type ManifestLoaderFunc func() (ManifestController, error)

// DefaultManifestLoader is the manifest loader we use when a manifest loader is not otherwise configured
var DefaultManifestLoader ManifestLoaderFunc

// ReconcilerOption implements the options pattern for reconcilers
type ReconcilerOption func(params reconcilerParams) reconcilerParams

// Options are a set of reconcilerOptions applied to all controllers
var Options struct {
	// Begin options are applied before evaluating controller specific options
	Begin []ReconcilerOption
	// End options are applied after evaluating controller specific options
	End []ReconcilerOption
}

type reconcilerParams struct {
	rawManifestOperations []ManifestOperation
	objectTransformations []ObjectTransform
	manifestController    ManifestController

	applier applier.Applier

	cascadingStrategy metav1.DeletionPropagation
	prune             bool
	preserveNamespace bool
	kustomize         bool
	validate          bool
	metrics           bool

	sink       Sink
	ownerFn    OwnerSelector
	labelMaker LabelMaker
	status     Status

	// hooks allow for interception of events during the reconciliation lifecycle
	hooks []Hook
}

type ManifestController interface {
	// ResolveManifest returns a raw manifest as a map[string]string for a given CR object
	ResolveManifest(ctx context.Context, object runtime.Object) (map[string]string, error)
}

type Sink interface {
	// Notify tells the Sink that all objs have been created
	Notify(ctx context.Context, dest DeclarativeObject, objs *manifest.Objects) error
}

// ManifestOperation is an operation that transforms raw string manifests before applying it
type ManifestOperation = func(context.Context, DeclarativeObject, string) (string, error)

// ObjectTransform is an operation that transforms the manifest objects before applying it
type ObjectTransform = func(context.Context, DeclarativeObject, *manifest.Objects) error

// OwnerSelector selects a runtime.Object to be the owner of a given manifest.Object
type OwnerSelector = func(context.Context, DeclarativeObject, manifest.Object, manifest.Objects) (DeclarativeObject, error)

// LabelMaker returns a fixed set of labels for a given DeclarativeObject
type LabelMaker = func(context.Context, DeclarativeObject) map[string]string

// WithRawManifestOperation adds the specific ManifestOperations to the chain of manifest changes
func WithRawManifestOperation(operations ...ManifestOperation) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.rawManifestOperations = append(p.rawManifestOperations, operations...)
		return p
	}
}

// WithObjectTransform adds the specified ObjectTransforms to the chain of manifest changes
func WithObjectTransform(operations ...ObjectTransform) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.objectTransformations = append(p.objectTransformations, operations...)
		return p
	}
}

// WithManifestController overrides the default source for loading manifests
func WithManifestController(mc ManifestController) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.manifestController = mc
		return p
	}
}

// WithApplyPrune turns on the --prune behavior of kubectl apply. This behavior deletes any
// objects that exist in the API server that are not deployed by the current version of the manifest
// which match a label specific to the addon instance.
//
// This option requires WithLabels to be used
func WithApplyPrune() ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.prune = true
		return p
	}
}

// WithOwner sets an owner ref on each deployed object by the OwnerSelector
func WithOwner(ownerFn OwnerSelector) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.ownerFn = ownerFn
		return p
	}
}

// WithLabels sets a fixed set of labels configured provided by a LabelMaker
// to all deployment objecs for a given DeclarativeObject
func WithLabels(labelMaker LabelMaker) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.labelMaker = labelMaker
		return p
	}
}

// WithStatus provides a Status interface that will be used during Reconcile
func WithStatus(status Status) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.status = status
		return p
	}
}

// WithPreserveNamespace preserves the namespaces defined in the deployment manifest
// instead of matching the namespace of the DeclarativeObject
func WithPreserveNamespace() ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.preserveNamespace = true
		return p
	}
}

// WithApplyKustomize run kustomize build to create final manifest
func WithApplyKustomize() ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.kustomize = true
		return p
	}
}

// WithManagedApplication is a transform that will modify the Application object
// in the deployment to match the configuration of the rest of the deployment.
func WithManagedApplication(labelMaker LabelMaker) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.objectTransformations = append(p.objectTransformations, func(ctx context.Context, instance DeclarativeObject, objects *manifest.Objects) error {
			return transformApplication(ctx, instance, objects, labelMaker)
		})
		return p
	}
}

// WithApplyValidation enables validation with kubectl apply
func WithApplyValidation() ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.validate = true
		return p
	}
}

// WithApplier allows us to select a different applier strategy
func WithApplier(applier applier.Applier) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.applier = applier
		return p
	}
}

// WithCascadingStrategy allows us to select a different CascadingStrategy, which ultimately sets the PropagationPolicy
func WithCascadingStrategy(cs metav1.DeletionPropagation) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.cascadingStrategy = cs
		return p
	}
}

// WithReconcileMetrics enables metrics of declarative reconciler.
// If metricsDuration is positive, metrics will be removed from
// Prometheus registry when metricsDuration times reconciliation
// has happened since k8s object related to that metrics is deleted
// and that k8s object hasn't been handled in those reconciliations.
// If metricsDuration is less than or equal to 0, metrics won't be
// removed.
//
// Argument ot specifies which ObjectTracker manages metrics of k8s
// objects. This enables specifying different metricsDuration for
// each set of k8s objects managed by different controllers, but
// all metrics are registered against manager's Prometheus registry.
// If ot is nil, package-scoped internal variable is used.
//
// If WithReconcileMetrics is called multiple times with same ot
// argument, largest metricsDuration is set against that ot.
func WithReconcileMetrics(metricsDuration int, ot *ObjectTracker) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		var err error

		p.metrics = true
		metricsRegisterOnce.Do(func() {
			for _, m := range metricsList {
				err = metrics.Registry.Register(m)
				if err != nil {
					break
				}
			}
		})

		if err != nil {
			panic(err)
		}

		if ot == nil {
			globalObjectTracker.setMetricsDurationInternal(metricsDuration)
		} else {
			ot.setMetricsDurationInternal(metricsDuration)
		}

		return p
	}
}

// WithHook allows for us to intercept and inject behaviours at various points in the lifecycle
func WithHook(hook Hook) ReconcilerOption {
	return func(p reconcilerParams) reconcilerParams {
		p.hooks = append(p.hooks, hook)
		return p
	}
}
