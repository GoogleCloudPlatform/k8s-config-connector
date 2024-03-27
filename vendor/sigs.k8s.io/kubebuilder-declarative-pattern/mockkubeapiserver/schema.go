package mockkubeapiserver

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
	"sigs.k8s.io/structured-merge-diff/v4/merge"
	"sigs.k8s.io/structured-merge-diff/v4/typed"

	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver/forked"
)

type mockSchema struct {
	builtin   *Schema
	resources []*ResourceInfo
}

func (s *mockSchema) Init() error {
	schema, err := KubernetesBuiltInSchema()
	if err != nil {
		return err
	}
	s.builtin = schema
	return nil
}

type typeInfo struct {
	ParserType typed.ParseableType
}

func (r *ResourceInfo) DoServerSideApply(ctx context.Context, live *unstructured.Unstructured, patchYAML []byte, options metav1.PatchOptions) (*unstructured.Unstructured, bool, error) {
	if r.TypeInfo == nil {
		return nil, false, fmt.Errorf("no type info for %v", r.GVK)
	}

	updater := merge.Updater{}

	liveObject, err := r.TypeInfo.ParserType.FromUnstructured(live.Object)
	if err != nil {
		return nil, false, fmt.Errorf("error parsing live object: %w", err)
	}

	configObject, err := r.TypeInfo.ParserType.FromYAML(typed.YAMLObject(patchYAML))
	if err != nil {
		return nil, false, fmt.Errorf("error parsing patch object: %w", err)
	}
	force := false
	if options.Force != nil {
		force = *options.Force
	}
	var managers fieldpath.ManagedFields
	manager := metav1.ManagedFieldsEntry{
		Manager: options.FieldManager,
	}
	// TODO: This is surprising ... the manager key is not the manager key, but rather the json-encoded form of the object
	// (at least Decode assumes this)
	managerJSON, err := json.Marshal(manager)
	if err != nil {
		return nil, false, fmt.Errorf("error encoding manager: %w", err)
	}

	apiVersion := fieldpath.APIVersion(r.GVK.GroupVersion().String())
	mergedObject, newManagers, err := updater.Apply(liveObject, configObject, apiVersion, managers, string(managerJSON), force)
	if err != nil {
		return nil, false, fmt.Errorf("error applying patch: %w", err)
	}
	if mergedObject == nil {
		// This indicates that the object was unchanged
		return nil, false, nil
	}

	if mergedObject == nil {
		return nil, false, fmt.Errorf("merged object was nil: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.Object = mergedObject.AsValue().Unstructured().(map[string]interface{})

	times := make(map[string]*metav1.Time)
	if err := forked.EncodeObjectManagedFields(u, newManagers, times); err != nil {
		return nil, false, fmt.Errorf("error from EncodeObjectManagedFields: %w", err)
	}

	return u, true, nil
}
