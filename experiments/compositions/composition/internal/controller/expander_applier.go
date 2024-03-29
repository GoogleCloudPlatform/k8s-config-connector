package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	compositionv1 "google.com/composition/api/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	PlanCR          *compositionv1.Plan
	Dryrun          bool
	Resource        string
	CompositionName string
	ExpanderName    string
	Name            string
	logger          logr.Logger
	ctx             context.Context
	client          client.Client
	objects         []applyset.ApplyableObject
}

func NewApplier(ctx context.Context, logger logr.Logger,
	r *ExpanderReconciler, plan *compositionv1.Plan,
	cr *unstructured.Unstructured, expanderName string) *Applier {
	return &Applier{
		RESTMapper:      r.RESTMapper,
		Config:          r.Config,
		Dynamic:         r.Dynamic,
		InputCR:         cr,
		PlanCR:          plan,
		Resource:        r.Resource,
		CompositionName: r.Composition.Name,
		ExpanderName:    expanderName,
		Name:            r.Composition.Name + "-" + cr.GetName(),
		Dryrun:          false,
		logger:          logger,
		ctx:             ctx,
		client:          r.Client,
	}
}

func (a *Applier) Load() error {
	var stage compositionv1.Stage
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

	a.objects = []applyset.ApplyableObject{}
	// loop over objects and extract unstructured
	for _, item := range objects.Items {

		//
		// TODO(barni@): requires more investigation
		// If a resource does not have a namespace dont set it.
		// Some Use cases for supporting cluster scoped resources:
		//   - Charlie onboarding Alice namespace
		//   - Charlie setting cluster rbac rules or bindings
		//

		//if item.GetNamespace() != "" {
		//	// Force set the namespace to the Input APIs (CRD_V) namespace
		//	item.SetNamespace(a.InputCR.GetNamespace())
		//} else {
		//	a.logger.Info("Namespace not reset for resource since it did not have namespace set",
		//		"gvk", item.GroupVersionKind(), "name", item.GetName())
		//}

		a.objects = append(a.objects, item.UnstructuredObject())
	}

	return nil
}

func (a *Applier) getApplyOptions(prune bool) (applyset.Options, error) {
	var options applyset.Options
	force := true
	patchOptions := metav1.PatchOptions{
		FieldManager: a.Resource + "-controller",
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

func (a *Applier) Apply(oldAppliers []*Applier, prune bool) error {
	options, err := a.getApplyOptions(prune)
	if err != nil {
		return err
	}
	applySet, err := applyset.New(options)
	if err != nil {
		return err
	}

	appliers := append(oldAppliers, a)
	objects := flattenObjects(appliers...)

	if err = applySet.SetDesiredObjects(objects); err != nil {
		return err
	}
	results, err := applySet.ApplyOnce(a.ctx)
	if err != nil {
		return err
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
	return err
}

func (a *Applier) Wait() (bool, error) {
	// TODO(barni@): Do we have standard status fields in KCC, ARC, ...
	// If so we can wait here. Else it is not feasible to implement a reliable wait
	return true, nil
}
