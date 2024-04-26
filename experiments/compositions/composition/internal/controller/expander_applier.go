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

	"github.com/go-logr/logr"
	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/applylib/applyset"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

type Applier struct {
	RESTMapper      meta.RESTMapper
	Config          *rest.Config
	Dynamic         *dynamic.DynamicClient
	InputCR         *unstructured.Unstructured
	PlanCR          *compositionv1alpha1.Plan
	Dryrun          bool
	InputGVR        schema.GroupVersionResource
	CompositionName string
	NamespaceMode   compositionv1alpha1.NamespaceMode
	ExpanderName    string
	Name            string
	logger          logr.Logger
	ctx             context.Context
	client          client.Client
	objects         []applyset.ApplyableObject
}

func NewApplier(ctx context.Context, logger logr.Logger,
	r *ExpanderReconciler, plan *compositionv1alpha1.Plan,
	cr *unstructured.Unstructured, c *compositionv1alpha1.Composition,
	expanderName string) *Applier {
	return &Applier{
		RESTMapper:      r.RESTMapper,
		Config:          r.Config,
		Dynamic:         r.Dynamic,
		InputCR:         cr,
		PlanCR:          plan,
		InputGVR:        r.InputGVR,
		CompositionName: r.Composition.Name,
		NamespaceMode:   c.Spec.NamespaceMode,
		ExpanderName:    expanderName,
		Name:            r.Composition.Name + "-" + cr.GetName(),
		Dryrun:          false,
		logger:          logger,
		ctx:             ctx,
		client:          r.Client,
	}
}

func (a *Applier) Count() int {
	return len(a.objects)
}

func (a *Applier) Load() error {
	var stage compositionv1alpha1.Stage
	stage, ok := a.PlanCR.Spec.Stages[a.ExpanderName]
	if !ok {
		a.logger.Info(".spec.stages did not have a matching expander name")
		return fmt.Errorf(".spec.stages did not have a matching expander name")
	}
	if stage.Manifest == "" {
		a.logger.Info(".spec.stages[name] has empty manifests. Nothing to apply")
		return fmt.Errorf(".spec.stages[name] has empty manifests. Nothing to apply.")
	}

	objects, err := manifest.ParseObjects(a.ctx, stage.Manifest)
	if err != nil {
		a.logger.Error(err, "Error parsing manifest")
		return err
	}

	err = a.injectOwnerRef(objects)
	if err != nil {
		a.logger.Error(err, "Error injecting ownerRefs")
		return err
	}

	a.objects = []applyset.ApplyableObject{}
	// loop over objects and extract unstructured
	for _, item := range objects.Items {
		// If namespaceMode is "" or "implicit", inherit the namespace
		if a.NamespaceMode == compositionv1alpha1.NamespaceModeNone ||
			a.NamespaceMode == compositionv1alpha1.NamespaceModeInherit {
			// Force set the namespace to the Input APIs (CRD_V) namespace
			item.SetNamespace(a.InputCR.GetNamespace())
		}
		a.objects = append(a.objects, item.UnstructuredObject())
	}

	return nil
}

func (a *Applier) injectOwnerRef(objects *manifest.Objects) error {
	for _, o := range objects.Items {
		// TODO (barney-s): This would result in some objects not being cleaned up.
		//  objects not in the plan namespace (cross namespace composition) would be skipped
		//  may be it is ok if we also create the namespace.
		if o.GetNamespace() != a.PlanCR.GetNamespace() {
			continue
		}
		gvk := a.PlanCR.GroupVersionKind()

		ownerRefs := []interface{}{
			map[string]interface{}{
				"apiVersion":         gvk.Group + "/" + gvk.Version,
				"blockOwnerDeletion": true,
				"controller":         true,
				"kind":               gvk.Kind,
				"name":               a.PlanCR.GetName(),
				"uid":                string(a.PlanCR.GetUID()),
			},
		}

		if err := o.SetNestedField(ownerRefs, "metadata", "ownerReferences"); err != nil {
			return err
		}
	}
	return nil
}

func (a *Applier) getApplyOptions(prune bool) (applyset.Options, error) {
	var options applyset.Options
	force := true
	patchOptions := metav1.PatchOptions{
		FieldManager: a.InputGVR.Resource + "-controller",
		Force:        &force,
	}
	if a.Dryrun {
		patchOptions.DryRun = []string{metav1.DryRunAll}
	}

	parentGVK := a.PlanCR.GroupVersionKind()
	restMapping, err := a.RESTMapper.RESTMapping(parentGVK.GroupKind(), parentGVK.Version)
	if err != nil {
		return options, err
	}

	parent := applyset.NewParentRef(a.PlanCR, a.PlanCR.GetName(), a.PlanCR.GetNamespace(), restMapping)
	options = applyset.Options{
		RESTMapper:   a.RESTMapper,
		Client:       a.Dynamic,
		Prune:        prune,
		PatchOptions: patchOptions,
		Parent:       parent,
		ParentClient: a.client,
	}
	return options, nil
}

func flattenObjects(appliers ...*Applier) []applyset.ApplyableObject {
	objects := []applyset.ApplyableObject{}
	for _, applier := range appliers {
		objects = append(objects, applier.objects...)
	}
	return objects
}

func (a *Applier) Apply(oldAppliers []*Applier, prune bool) (*applyset.ApplyResults, error) {
	options, err := a.getApplyOptions(prune)
	if err != nil {
		return nil, err
	}
	applySet, err := applyset.New(options)
	if err != nil {
		return nil, err
	}

	appliers := append(oldAppliers, a)
	objects := flattenObjects(appliers...)

	if err = applySet.SetDesiredObjects(objects); err != nil {
		return nil, err
	}
	results, err := applySet.ApplyOnce(a.ctx)
	if err != nil {
		return results, err
	}

	/*
		if containsCRDObject(objects) {
			// Reset the REST Mapper's cache so we can discover any new CRDs immediately without
			// needing to wait for the cache to expire. This clears any cached 'no matches'.
			c.RestMapper.Reset()
		}
	*/

	if !results.AllApplied() {
		err = fmt.Errorf("Unable to apply all objects")
	}
	return results, err
}

func (a *Applier) Wait() (bool, error) {
	// TODO(barni@): Do we have standard status fields in KCC, ARC, ...
	// If so we can wait here. Else it is not feasible to implement a reliable wait
	return true, nil
}
