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

package manifest

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Objects holds a collection of objects, so that we can filter / sequence them
type Objects struct {
	Items []*Object
	Blobs [][]byte
	Path  string
}

// GetItems is a nil-safe wrapper around
func (o *Objects) GetItems() []*Object {
	if o == nil {
		return nil
	}
	return o.Items
}

type Object struct {
	object *unstructured.Unstructured

	Group     string
	Version   string
	Kind      string
	name      string
	namespace string

	json []byte
}

func ParseJSONToObject(json []byte) (*Object, error) {
	o, gvk, err := unstructured.UnstructuredJSONScheme.Decode(json, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing json into unstructured object: %v", err)
	}

	u, ok := o.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("parsed unexpected type %T", o)
	}

	return &Object{
		object:    u,
		Group:     gvk.Group,
		Version:   gvk.Version,
		Kind:      gvk.Kind,
		name:      u.GetName(),
		namespace: u.GetNamespace(),
		json:      json,
	}, nil
}

func (o *Object) AddLabels(labels map[string]string) {
	merged := make(map[string]string)
	for k, v := range o.object.GetLabels() {
		merged[k] = v
	}

	for k, v := range labels {
		merged[k] = v
	}

	o.object.SetLabels(merged)
	// Invalidate cached json
	o.json = nil
}

func (o *Object) AddAnnotations(annotations map[string]string) {
	merged := make(map[string]string)
	for k, v := range o.object.GetAnnotations() {
		merged[k] = v
	}

	for k, v := range annotations {
		merged[k] = v
	}

	o.object.SetAnnotations(merged)
	// Invalidate cached json
	o.json = nil
}

func (o *Object) GetNamespace() string {
	return o.namespace
}

func (o *Object) SetNamespace(ns string) error {
	if err := o.SetNestedField(ns, "metadata", "namespace"); err != nil {
		return fmt.Errorf("failed to set namespace: %w", err)
	}
	o.namespace = ns
	return nil
}

func (o *Object) GetName() string {
	return o.name
}

func (o *Object) SetName(name string) error {
	if err := o.SetNestedField(name, "metadata", "name"); err != nil {
		return fmt.Errorf("failed to set name: %w", err)
	}
	o.name = name
	return nil
}

func (o *Object) SetNestedStringMap(value map[string]string, fields ...string) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}
	err := unstructured.SetNestedStringMap(o.object.Object, value, fields...)
	// Invalidate cached json
	o.json = nil
	return err
}

func nestedFieldNoCopy(obj map[string]interface{}, fields ...string) (interface{}, bool, error) {
	var val interface{} = obj

	for i, field := range fields {
		if m, ok := val.(map[string]interface{}); ok {
			val, ok = m[field]
			if !ok {
				return nil, false, nil
			}
		} else {
			return nil, false, fmt.Errorf("%v accessor error: %v is of the type %T, expected map[string]interface{}", strings.Join(fields[:i+1], "."), val, val)
		}
	}
	return val, true, nil
}

func (o *Object) MutateContainers(fn func(map[string]interface{}) error) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}

	containers, found, err := nestedFieldNoCopy(o.object.Object, "spec", "template", "spec", "containers")
	if err != nil {
		return fmt.Errorf("error reading containers: %v", err)
	}

	if !found {
		return fmt.Errorf("containers not found")
	}

	containerList, ok := containers.([]interface{})
	if !ok {
		return fmt.Errorf("containers was not a list")
	}

	initContainers, found, err := nestedFieldNoCopy(o.object.Object, "spec", "template", "spec", "initContainers")
	if err != nil {
		return fmt.Errorf("error reading init containers: %v", err)
	}

	if found {
		initContainerList, ok := initContainers.([]interface{})
		if !ok {
			return fmt.Errorf("init containers was not a list")
		}

		containerList = append(containerList, initContainerList...)
	}

	for _, co := range containerList {
		container, ok := co.(map[string]interface{})
		if !ok {
			return fmt.Errorf("container was not an object")
		}

		if err := fn(container); err != nil {
			return err
		}
	}

	// Invalidate cached json
	o.json = nil
	return err
}

func (o *Object) MutatePodSpec(fn func(map[string]interface{}) error) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}

	sp, found, err := nestedFieldNoCopy(o.object.Object, "spec", "template", "spec")
	if err != nil {
		return fmt.Errorf("error reading containers: %v", err)
	}

	if !found {
		return fmt.Errorf("pod spec not found")
	}

	podSpec, ok := sp.(map[string]interface{})
	if !ok {
		return fmt.Errorf("pod spec was not an object")
	}
	if err := fn(podSpec); err != nil {
		return err
	}

	// Invalidate cached json
	o.json = nil
	return err
}

func (o *Object) NestedStringMap(fields ...string) (map[string]string, bool, error) {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}
	return unstructured.NestedStringMap(o.object.Object, fields...)
}

func (o *Object) SetNestedField(value interface{}, fields ...string) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}
	err := unstructured.SetNestedField(o.object.Object, value, fields...)
	// Invalidate cached json
	o.json = nil

	return err
}

func (o *Object) SetNestedSlice(value []interface{}, fields ...string) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}
	err := unstructured.SetNestedSlice(o.object.Object, value, fields...)
	// Invalidate cached json
	o.json = nil

	return err
}

func (o *Object) SetNestedFieldNoCopy(value interface{}, fields ...string) error {
	if o.object.Object == nil {
		o.object.Object = make(map[string]interface{})
	}
	// err := unstructured.SetNestedField(o.object.Object, value, fields...)

	m := o.object.Object

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", fields[:i+1])
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value

	// Invalidate cached json
	o.json = nil

	return nil
}

func (o *Object) JSON() ([]byte, error) {
	if o.json != nil {
		return o.json, nil
	}

	b, err := o.object.MarshalJSON()
	if err != nil {
		return nil, err
	}
	o.json = b
	return b, nil
}

// UnstructuredContent exposes the raw object, primarily for testing
func (o *Object) UnstructuredObject() *unstructured.Unstructured {
	return o.object
}

func (o *Object) GroupKind() schema.GroupKind {
	return o.object.GroupVersionKind().GroupKind()
}

func (o *Object) GroupVersionKind() schema.GroupVersionKind {
	return o.object.GroupVersionKind()
}

// NamespacedName returns the name and namespace of the object in a types.NamespacedName.
// Note that this reflects the state of the object; if the namespace is not yet set,
// it will returned as "" here, even though it would likely be defaulted before apply.
func (o *Object) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: o.GetNamespace(),
		Name:      o.GetName(),
	}
}

func (o *Objects) JSONManifest() (string, error) {
	var b bytes.Buffer

	for i, item := range o.Items {
		if i != 0 {
			b.WriteString("\n\n")
		}
		// We build a JSON manifest because conversion to yaml is harder
		// (and we've lost line numbers anyway if we applied any transforms)
		json, err := item.JSON()
		if err != nil {
			return "", fmt.Errorf("error building json: %v", err)
		}
		b.Write(json)
	}

	return b.String(), nil
}

// Sort will order the items in Objects in order of score, group, kind, name.  The intent is to
// have a deterministic ordering in which Objects are applied.
func (o *Objects) Sort(score func(o *Object) int) {
	sort.Slice(o.Items, func(i, j int) bool {
		iScore := score(o.Items[i])
		jScore := score(o.Items[j])
		return iScore < jScore ||
			(iScore == jScore &&
				o.Items[i].Group < o.Items[j].Group) ||
			(iScore == jScore &&
				o.Items[i].Group == o.Items[j].Group &&
				o.Items[i].Kind < o.Items[j].Kind) ||
			(iScore == jScore &&
				o.Items[i].Group == o.Items[j].Group &&
				o.Items[i].Kind == o.Items[j].Kind &&
				o.Items[i].name < o.Items[j].name)
	})
}

func ParseObjects(ctx context.Context, manifest string) (*Objects, error) {
	log := log.FromContext(ctx)

	objects := &Objects{}
	reader := k8syaml.NewYAMLReader(bufio.NewReader(strings.NewReader(manifest)))
	for {
		raw, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return objects, nil
			}

			return nil, fmt.Errorf("invalid YAML doc: %w", err)
		}

		raw = bytes.TrimSpace(raw)
		out := unstructured.Unstructured{}
		if err := k8syaml.Unmarshal(raw, &out); err != nil {
			log.WithValues("error", err).WithValues("yaml", raw).V(2).Info("Unable to parse into Unstructured, storing as blob")
			objects.Blobs = append(objects.Blobs, append(bytes.TrimPrefix(raw, []byte("---\n")), '\n'))
		}

		if len(raw) == 0 || bytes.Equal(raw, []byte("null")) || len(out.Object) == 0 {
			continue
		}

		o, err := NewObject(&out)
		if err != nil {
			return nil, err
		}
		objects.Items = append(objects.Items, o)
	}
}

func newObject(u *unstructured.Unstructured, json []byte) (*Object, error) {
	o := &Object{
		object: u,
		json:   json,
	}

	gvk := u.GetObjectKind().GroupVersionKind()
	o.Group = gvk.Group
	o.Kind = gvk.Kind
	o.name = u.GetName()
	o.namespace = u.GetNamespace()

	return o, nil
}

func NewObject(u *unstructured.Unstructured) (*Object, error) {
	var json []byte
	return newObject(u, json)
}
