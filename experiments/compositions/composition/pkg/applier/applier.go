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

package applier

import (
	"context"
	"fmt"
	"strings"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/pkg/cel"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/applylib/applyset"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const StageLabel = "compositions.google.com/applier-stage"

type ApplierClient struct {
	RESTMapper meta.RESTMapper
	Dynamic    *dynamic.DynamicClient
	Client     client.Client
}

type Applier struct {
	client    ApplierClient
	planCR    *compositionv1alpha1.Plan
	Dryrun    bool
	stageName string
	namespace string
	resource  string
	logger    logr.Logger
	ctx       context.Context
	objects   []applyset.ApplyableObject
	results   *applyset.ApplyResults
	readiness []compositionv1alpha1.ReadyOn
}

func NewApplier(
	ctx context.Context, logger logr.Logger,
	ac ApplierClient,
	stage string, namespace string, resource string,
	plan *compositionv1alpha1.Plan,
	readiness []compositionv1alpha1.ReadyOn,
) *Applier {
	return &Applier{
		client:    ac,
		resource:  resource,
		planCR:    plan,
		namespace: namespace,
		stageName: stage,
		Dryrun:    false,
		logger:    logger,
		ctx:       ctx,
		readiness: readiness,
	}
}

func (a *Applier) Count() int {
	return len(a.objects)
}

func (a *Applier) UpdatePruneStatus(status *compositionv1alpha1.PlanStatus) {
	for _, resultObj := range a.results.Objects {
		if !resultObj.Apply.IsPruned {
			continue
		}

		rs := compositionv1alpha1.ResourceStatus{
			Group:     resultObj.GVK.Group,
			Version:   resultObj.GVK.Version,
			Kind:      resultObj.GVK.Kind,
			Namespace: resultObj.NameNamespace.Namespace,
			Name:      resultObj.NameNamespace.Name,
			Status:    "Pruned",
			Health:    compositionv1alpha1.HEALTHY, // Is it ?
		}
		if resultObj.Apply.Error != nil {
			rs.Status = fmt.Sprintf("Prune Error: %s", resultObj.Apply.Error)
		}
		status.LastPruned = append(status.LastPruned, rs)
	}
}

func (a *Applier) UpdateStageStatus(status *compositionv1alpha1.PlanStatus) {
	applyCount := 0
	if status.Stages[a.stageName] == nil {
		status.Stages[a.stageName] = &compositionv1alpha1.StageStatus{}
	}
	status.Stages[a.stageName].LastApplied = []compositionv1alpha1.ResourceStatus{}

	for _, resultObj := range a.results.Objects {
		// Match objects from this applier only.
		fromCurrentApplier := false
		for _, applierObj := range a.objects {
			if applierObj.GroupVersionKind() == resultObj.GVK &&
				applierObj.GetNamespace() == resultObj.NameNamespace.Namespace &&
				applierObj.GetName() == resultObj.NameNamespace.Name {
				fromCurrentApplier = true
			}
		}
		if fromCurrentApplier {
			if status.Stages == nil {
				status.Stages = map[string]*compositionv1alpha1.StageStatus{}
			}
			rs := compositionv1alpha1.ResourceStatus{
				Group:     resultObj.GVK.Group,
				Version:   resultObj.GVK.Version,
				Kind:      resultObj.GVK.Kind,
				Namespace: resultObj.NameNamespace.Namespace,
				Name:      resultObj.NameNamespace.Name,
				Status:    "",
				Health:    compositionv1alpha1.UNHEALTHY,
			}
			if resultObj.Apply.IsPruned {
				rs.Status = "Unexpected Prune"
			} else {
				if resultObj.Apply.Error != nil {
					rs.Status = fmt.Sprintf("Apply Error: %s", resultObj.Apply.Error)
				} else {
					applyCount++
					rs.Status = resultObj.Apply.Message
				}
				if resultObj.Health.IsHealthy {
					rs.Health = compositionv1alpha1.HEALTHY
				}
			}
			status.Stages[a.stageName].LastApplied = append(status.Stages[a.stageName].LastApplied, rs)
		}
	}
	status.Stages[a.stageName].AppliedCount = applyCount
}

func (a *Applier) Load() error {
	var stage compositionv1alpha1.Stage
	stage, ok := a.planCR.Spec.Stages[a.stageName]
	if !ok {
		a.logger.Info(".spec.stages did not have a matching expander name")
		return fmt.Errorf(".spec.stages did not have a matching expander name")
	}

	a.objects = []applyset.ApplyableObject{}

	// We dont error out on empty manifests
	if stage.Manifest == "" {
		a.logger.Info(".spec.stages[name] has empty manifests. Nothing to apply")
		return nil
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
	a.addStageLabel(objects)

	// loop over objects and extract unstructured
	for _, item := range objects.Items {

		// If namespace is passed it is namespace mode composition
		if a.namespace != "" {
			err := item.SetNamespace(a.namespace)
			if err != nil {
				a.logger.Error(err, "Error setting namespace")
				return err
			}
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
		if o.GetNamespace() != a.planCR.GetNamespace() {
			continue
		}
		gvk := a.planCR.GroupVersionKind()

		ownerRefs := []interface{}{
			map[string]interface{}{
				"apiVersion":         gvk.Group + "/" + gvk.Version,
				"blockOwnerDeletion": true,
				"controller":         true,
				"kind":               gvk.Kind,
				"name":               a.planCR.GetName(),
				"uid":                string(a.planCR.GetUID()),
			},
		}

		if err := o.SetNestedField(ownerRefs, "metadata", "ownerReferences"); err != nil {
			return err
		}
	}
	return nil
}

func (a *Applier) addStageLabel(objects *manifest.Objects) {
	labels := map[string]string{StageLabel: a.stageName}
	for _, o := range objects.Items {
		o.AddLabels(labels)
	}
}

func (a *Applier) getApplyOptions(prune bool) (applyset.Options, error) {
	var options applyset.Options
	force := true
	patchOptions := metav1.PatchOptions{
		FieldManager: a.resource + "-controller",
		Force:        &force,
	}
	if a.Dryrun {
		patchOptions.DryRun = []string{metav1.DryRunAll}
	}

	parentGVK := a.planCR.GroupVersionKind()
	restMapping, err := a.client.RESTMapper.RESTMapping(parentGVK.GroupKind(), parentGVK.Version)
	if err != nil {
		return options, err
	}

	parent := applyset.NewParentRef(a.planCR, a.planCR.GetName(), a.planCR.GetNamespace(), restMapping)
	options = applyset.Options{
		RESTMapper:    a.client.RESTMapper,
		Client:        a.client.Dynamic,
		Prune:         prune,
		PatchOptions:  patchOptions,
		Parent:        parent,
		ParentClient:  a.client.Client,
		Tooling:       "compositions",
		ComputeHealth: a.isObjectReady,
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
	a.results, err = applySet.ApplyOnce(a.ctx)
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

	if !a.results.AllApplied() {
		err = fmt.Errorf("Unable to apply all objects")
	}
	return err
}

func (a *Applier) getReadinessRule(u *unstructured.Unstructured) string {
	gvk := u.GetObjectKind().GroupVersionKind()
	match := false
	for i := range a.readiness {
		if a.readiness[i].Group == gvk.Group &&
			a.readiness[i].Version == gvk.Version &&
			a.readiness[i].Kind == gvk.Kind {
			match = true
		}
		if !match {
			continue
		}
		if a.readiness[i].Name != "" {
			match = a.readiness[i].Name == u.GetName()
		}
		if !match {
			continue
		}
		if a.readiness[i].Namespace != "" {
			match = a.readiness[i].Namespace == u.GetNamespace()
		}
		if !match {
			continue
		}
		return a.readiness[i].Ready
	}
	return ""
}

// isObjectReady - is the object ready
func (a *Applier) isObjectReady(u *unstructured.Unstructured) (bool, string, error) {
	ready := false
	var err error
	message := ""

	rule := a.getReadinessRule(u)
	if rule != "" {
		// Create a CEL engine
		celEngine, err := cel.NewEngine(u)
		if err != nil {
			return ready, message, fmt.Errorf("error creating CEL engine: %w", err)
		}
		result, err := celEngine.Eval(rule)
		if err != nil {
			return ready, message, fmt.Errorf("error Evaluating expression: %s, %w", rule, err)
		}
		resultBool, ok := result.Value().(bool)
		if !ok {
			message := fmt.Sprintf("Readiness Rule [%s] not evaluating to bool. result: %v", rule, result.Value())
			return ready, message, fmt.Errorf("%s, %w", message, err)
		}
		message = fmt.Sprintf("Readiness Rule evaluated to %v", resultBool)
		ready = resultBool
	} else {
		// Fall back to kstatus
		result, err := status.Compute(u)
		if err == nil {
			ready = result.Status == status.CurrentStatus
			message = result.Message
		}
	}
	return ready, message, err
}

func (a *Applier) AreResourcesReady() (bool, error) {
	// Check for readiness or progress CEL rules
	allReady := true
	for i, resultObj := range a.results.Objects {
		// Match objects from this applier only.
		fromCurrentApplier := false
		for _, applierObj := range a.objects {
			if applierObj.GroupVersionKind() == resultObj.GVK &&
				applierObj.GetNamespace() == resultObj.NameNamespace.Namespace &&
				applierObj.GetName() == resultObj.NameNamespace.Name {
				fromCurrentApplier = true
			}
		}
		if fromCurrentApplier {
			if resultObj.Apply.IsPruned {
				continue
			}
			a.results.Objects[i].Health.Message = resultObj.Health.Message
			a.results.Objects[i].Health.IsHealthy = resultObj.Health.IsHealthy
			if !resultObj.Health.IsHealthy {
				allReady = false
			}
		}
	}

	return allReady, nil
}

func (a *Applier) AddAppliedObjectsIntoValues(values map[string]interface{}) map[string]interface{} {
	for _, resultObj := range a.results.Objects {
		if resultObj.Apply.IsPruned {
			continue
		}
		obj := resultObj.LastApplied
		name := obj.GetName()
		gvk := obj.GroupVersionKind()
		kind := strings.ToLower(gvk.Kind)

		// short path: values.<kind>.<name>. May clash
		_, ok := values[kind]
		if !ok {
			values[kind] = map[string]interface{}{}
		}
		ref := values[kind].(map[string]interface{})

		_, ok = ref[name]
		if ok {
			// Clash !! We will ignore
			a.logger.Info("Clash when adding applied objects to values.", "kind", kind, "name", name)
		} else {
			ref[name] = obj.Object
		}

		// long path: values.<group>.<kind>.<namespace>.<name> will not clash
		// Long path may not be practical since the namespace is not part of the composition and
		// most templating languages dont support nested templatable variable.
		// ex: {{ values.deployment.teampage.status.something  }} will work in jinja2
		// but this {{ values.apps.deployment.{{teampage.metadata.namespace}}.status.something }} wont work
		// So leaving this code commented for now.
		/*
			group := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "_")
			namespace := obj.GetNamespace()

			if group == "" {
				group = "core"
			}
			_, ok = values[group]
			if !ok {
				values[group] = map[string]interface{}{}
			}
			ref = values[group].(map[string]interface{})

			_, ok = ref[kind]
			if !ok {
				ref[kind] = map[string]interface{}{}
			}
			ref = ref[kind].(map[string]interface{})

			_, ok = ref[namespace]
			if !ok {
				ref[namespace] = map[string]interface{}{}
			}
			ref = ref[namespace].(map[string]interface{})

			_, ok = ref[name]
			if !ok {
				ref[name] = map[string]interface{}{}
			}
			ref[name] = obj.Object
		*/

	}
	return values
}
