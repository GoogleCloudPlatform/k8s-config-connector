// Copyright 2023 Google LLC
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

package controllers

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/engines/manifestengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/watchset"
)

var _ reconcile.Reconciler = &instanceReconciler{}

// instanceReconciler reconciles a CompositeDefinition object
type instanceReconciler struct {
	client     client.Client
	restMapper meta.RESTMapper
	config     *rest.Config
	scheme     *runtime.Scheme

	dynamicClient dynamic.Interface

	fileName   string
	engine     string
	definition string

	gvk             schema.GroupVersionKind
	watchsetManager *watchset.ControllerManager
}

func (r *instanceReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log := klog.FromContext(ctx)

	id := req.NamespacedName
	subject := &unstructured.Unstructured{}
	subject.SetAPIVersion(r.gvk.GroupVersion().Identifier())
	subject.SetKind(r.gvk.Kind)

	gvk, err := apiutil.GVKForObject(subject, r.scheme)
	if err != nil {
		return reconcile.Result{}, err
	}

	restMapping, err := r.restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return reconcile.Result{}, err
	}
	gvr := restMapping.Resource
	watches := r.watchsetManager.ReconcileStart(ctx, id)

	if err := r.client.Get(ctx, id, subject); err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	watches.DependencySet.WatchObject(gvr, id, subject.GetResourceVersion())

	log.Info("reconcile request for object", "id", id)

	result, err := r.reconcileExists(ctx, watches.DependencySet, id, subject)
	if err != nil {
		watches.ReconcileFailed()
		return reconcile.Result{}, err
	}

	watches.ReconcileSuccess()

	// TODO: Why do we have to pass String() ?
	log.Info("result", "result", result, "dependencies", watches.DependencySet.String())
	return reconcile.Result{}, err
}

type Engine interface {
	BuildObjects(ctx context.Context, fileName string, script string, subject *unstructured.Unstructured) ([]*unstructured.Unstructured, error)
}

func (r *instanceReconciler) BuildDeploymentObjects(ctx context.Context, name types.NamespacedName, subject *unstructured.Unstructured) (*manifest.Objects, error) {
	var engine Engine
	switch r.engine {
	case "yaml":
		engine = manifestengine.NewEngine(r.restMapper, r.dynamicClient)
	default:
		return nil, fmt.Errorf("engine %q not known", r.engine)
	}

	objects, err := engine.BuildObjects(ctx, r.fileName, r.definition, subject)
	if err != nil {
		return nil, err
	}
	out := &manifest.Objects{}
	for _, u := range objects {
		o, err := manifest.NewObject(u)
		if err != nil {
			return nil, err
		}
		out.Items = append(out.Items, o)
	}
	return out, nil
}

func (r *instanceReconciler) reconcileExists(ctx context.Context, dependencies *watchset.DependencySet, name types.NamespacedName, instance *unstructured.Unstructured) (*declarative.StatusInfo, error) {
	log := log.FromContext(ctx)
	log.WithValues("object", name.String()).Info("reconciling")

	statusInfo := &declarative.StatusInfo{}
	statusInfo.Subject = instance

	dynamicClient := dependencies.TrackingDynamicClient(r.dynamicClient)

	objects, err := r.BuildDeploymentObjects(ctx, name, instance)
	if err != nil {
		log.Error(err, "building deployment objects")
		return statusInfo, fmt.Errorf("error building deployment objects: %v", err)
	}

	log.WithValues("objects", fmt.Sprintf("%d", len(objects.Items))).Info("built deployment objects")
	statusInfo.Manifest = objects

	var newItems []*manifest.Object
	for _, obj := range objects.Items {
		newItems = append(newItems, obj)
	}
	objects.Items = newItems

	extraArgs := []string{}

	ns := name.Namespace

	parentRef, err := applier.NewParentRef(r.restMapper, instance, instance.GroupVersionKind(), instance.GetName(), instance.GetNamespace())
	if err != nil {
		return statusInfo, err
	}
	applierOpt := applier.ApplierOptions{
		RESTConfig:        r.config,
		RESTMapper:        r.restMapper,
		Namespace:         ns,
		ParentRef:         parentRef,
		Objects:           objects.GetItems(),
		DynamicClient:     dynamicClient,
		Validate:          false, //r.options.validate,
		ExtraArgs:         extraArgs,
		Force:             true,
		CascadingStrategy: "Foreground", //r.options.cascadingStrategy,
		Client:            r.client,
	}

	// TODO: Don't prune until objects are healthy
	applierOpt.Prune = true

	patchOptions := metav1.PatchOptions{FieldManager: "kdp-test"}

	applier := applier.NewApplySetApplier(patchOptions, metav1.DeleteOptions{}, applier.ApplysetOptions{})

	if err := applier.Apply(ctx, applierOpt); err != nil {
		log.Error(err, "applying manifest")
		statusInfo.KnownError = declarative.KnownErrorApplyFailed
		return statusInfo, fmt.Errorf("error applying manifest: %v", err)
	}

	statusInfo.LiveObjects = func(ctx context.Context, gvk schema.GroupVersionKind, nn types.NamespacedName) (*unstructured.Unstructured, error) {
		// TODO: Applier should return the objects in their post-apply state, so we don't have to requery

		mapping, err := r.restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return nil, fmt.Errorf("unable to get mapping for resource %v: %w", gvk, err)
		}

		var resource dynamic.ResourceInterface
		switch mapping.Scope {
		case meta.RESTScopeNamespace:
			resource = r.dynamicClient.Resource(mapping.Resource).Namespace(nn.Namespace)
		case meta.RESTScopeRoot:
			resource = r.dynamicClient.Resource(mapping.Resource)
		default:
			return nil, fmt.Errorf("unknown scope %v", mapping.Scope)
		}
		u, err := resource.Get(ctx, nn.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("error getting object: %w", err)
		}
		return u, nil
	}

	return statusInfo, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *instanceReconciler) init(mgr ctrl.Manager, watchsets *watchset.Manager, subject *v1alpha1.CompositeDefinition) error {
	r.client = mgr.GetClient()
	r.restMapper = mgr.GetRESTMapper()
	r.config = mgr.GetConfig()
	r.scheme = mgr.GetScheme()

	d, err := dynamic.NewForConfig(r.config)
	if err != nil {
		return err
	}
	r.dynamicClient = d

	r.definition = subject.Spec.Definition
	r.engine = subject.Spec.Engine
	r.fileName = subject.GetName()

	r.gvk = schema.FromAPIVersionAndKind(subject.Spec.ReconcilerFor.APIVersion, subject.Spec.ReconcilerFor.Kind)

	return nil
}

type instanceReconcilerRunner struct {
	controller controller.Controller
	reconciler *instanceReconciler
	ctx        context.Context
	cancel     func()
	result     *Future[error]
}

func newInstanceReconcilerRunner(mgr ctrl.Manager, watchsets *watchset.Manager, subject *v1alpha1.CompositeDefinition) (*instanceReconcilerRunner, error) {
	r := &instanceReconciler{}
	if err := r.init(mgr, watchsets, subject); err != nil {
		return nil, err
	}

	c, err := controller.NewUnmanaged("instance-reconciler", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return nil, err
	}

	watchsetManager, err := watchsets.NewControllerManager(c)
	if err != nil {
		return nil, err
	}
	r.watchsetManager = watchsetManager

	actsOn := &unstructured.Unstructured{}
	actsOn.SetAPIVersion(r.gvk.GroupVersion().Identifier())
	actsOn.SetKind(r.gvk.Kind)

	// Watch for changes to CompositeDefinition
	err = c.Watch(source.TypedKind(mgr.GetCache(), actsOn, &handler.TypedEnqueueRequestForObject[*unstructured.Unstructured]{}))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	runner := &instanceReconcilerRunner{
		reconciler: r,
		controller: c,
		ctx:        ctx,
		cancel:     cancel,
		result:     newFuture[error](),
	}

	return runner, nil
}

func (r *instanceReconcilerRunner) stop() error {
	r.reconciler.watchsetManager.Stop()

	r.cancel()

	return r.result.Wait()
}

func (r *instanceReconcilerRunner) start() {
	go func() {
		err := r.controller.Start(r.ctx)
		if err != nil {
			klog.Warningf("error from instance-reconciler controller: %v", err)
		}
		r.result.Set(err)
	}()
}
