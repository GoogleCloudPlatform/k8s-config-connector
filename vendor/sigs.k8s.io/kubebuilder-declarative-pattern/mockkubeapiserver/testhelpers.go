/*
Copyright 2022 The Kubernetes Authors.

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

package mockkubeapiserver

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

// RegisterType registers a type with the schema for the mock kubeapiserver
func (s *MockKubeAPIServer) RegisterType(gvk schema.GroupVersionKind, resource string, scope meta.RESTScope) {
	s.storage.RegisterType(gvk, resource, scope)
}

// AddObject pre-creates an object
func (s *MockKubeAPIServer) AddObject(obj *unstructured.Unstructured) error {
	klog.Infof("precreating %s object %s/%s", obj.GroupVersionKind().Kind, obj.GetNamespace(), obj.GetName())
	return s.storage.AddObject(obj)
}

// AddObjectsFromManifest pre-creates the objects in the manifest
func (s *MockKubeAPIServer) AddObjectsFromManifest(y string) error {
	for _, obj := range strings.Split(y, "\n---\n") {
		u := &unstructured.Unstructured{}
		if err := yaml.Unmarshal([]byte(obj), &u.Object); err != nil {
			return fmt.Errorf("failed to unmarshal object %q: %w", obj, err)
		}
		if err := s.AddObject(u); err != nil {
			return err
		}
	}
	return nil
}
