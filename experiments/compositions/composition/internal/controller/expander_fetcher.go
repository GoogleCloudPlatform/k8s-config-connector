package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	compositionv1 "google.com/composition/api/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Fetcher struct {
	InputCR  *unstructured.Unstructured
	Expander *compositionv1.Expander
	PlanCR   *compositionv1.Plan
	logger   logr.Logger
	ctx      context.Context
	client   client.Client
	values   map[string]interface{}
}

func NewFetcher(ctx context.Context, logger logr.Logger,
	r *ExpanderReconciler, plan *compositionv1.Plan,
	cr *unstructured.Unstructured, expander *compositionv1.Expander) *Fetcher {
	return &Fetcher{
		InputCR:  cr,
		PlanCR:   plan,
		Expander: expander,
		logger:   logger,
		ctx:      ctx,
		values:   make(map[string]interface{}),
		client:   r.Client,
	}
}

func (f *Fetcher) updateValues(obj *unstructured.Unstructured, vf *compositionv1.ValuesFrom) error {
	for index := range vf.FieldRef {
		fr := &vf.FieldRef[index]
		path := strings.Split(strings.TrimLeft(fr.Path, "."), ".")
		v, ok, err := unstructured.NestedFieldCopy(obj.Object, path...)
		if err != nil {
			f.logger.Error(err, "Failed to get traverse field path",
				"gvk", obj.GroupVersionKind(), "namespace", obj.GetNamespace(),
				"name", obj.GetName(), "fieldPath", fr.Path)
			return err
		}
		if ok {
			if f.values[vf.Name] == nil {
				f.values[vf.Name] = map[string]interface{}{}
			}
			f.values[vf.Name].(map[string]interface{})[fr.As] = v
		} else {
			f.logger.Error(err, "field path not present in object yet",
				"gvk", obj.GroupVersionKind(), "namespace", obj.GetNamespace(),
				"name", obj.GetName(), "fieldPath", fr.Path)
			return fmt.Errorf("Waiting for field: %s, in object: %s/%s/%s", fr.Path,
				obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
		}
	}

	return nil
}

// TODO(barni@): This is generic enough to be a util function. Move it into a util package.
// possible use in composition reconciler as well.
func (f *Fetcher) getObject(vf *compositionv1.ValuesFrom) (*unstructured.Unstructured, error) {
	obj := unstructured.Unstructured{}
	gvk := schema.GroupVersionKind{
		Group:   vf.ResourceRef.Group,
		Version: vf.ResourceRef.Version,
		Kind:    vf.ResourceRef.Kind,
	}
	nn := types.NamespacedName{Name: vf.ResourceRef.Name, Namespace: f.InputCR.GetNamespace()}
	obj.SetGroupVersionKind(gvk)
	if err := f.client.Get(f.ctx, nn, &obj); err != nil {
		f.logger.Info("Failed to get dependent object", "gvk", gvk, "name", nn)
		return nil, err
	}
	return &obj, nil
}

func (f *Fetcher) Fetch() error {
	for index := range f.Expander.ValuesFrom {
		vf := &f.Expander.ValuesFrom[index]
		obj, err := f.getObject(vf)
		if err != nil {
			return err
		}
		err = f.updateValues(obj, vf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Fetcher) UpdatePlanCR() error {
	var stage compositionv1.Stage
	stage, ok := f.PlanCR.Spec.Stages[f.Expander.Name]
	if !ok {
		stage = compositionv1.Stage{}
	}

	// marshall values
	values, err := json.Marshal(f.values)
	if err != nil {
		return fmt.Errorf("Failed to marshal values when updating InputCR status: %v", err)
	}

	stage.Values = string(values)
	f.PlanCR.Spec.Stages[f.Expander.Name] = stage

	err = f.client.Update(f.ctx, f.PlanCR)
	if err != nil {
		f.logger.Error(err, "Error updating PlanCR updating fetched values")
		return fmt.Errorf("Failed updating Plan CR while setting fetched values: %v", err)
	}
	return nil
}
