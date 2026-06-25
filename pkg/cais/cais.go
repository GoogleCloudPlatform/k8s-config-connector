// Copyright 2026 Google LLC
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

package cais

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/objects"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CAISIdentityResult struct {
	Group     string `json:"group"`
	Version   string `json:"version"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	CAISURL   string `json:"caisURL"`
	Error     string `json:"error,omitempty"`
}

func NewScheme() *runtime.Scheme {
	return runtime.NewScheme()
}

var unsupportedKinds = map[schema.GroupKind]bool{
	{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringNotificationChannel"}:        true,
	{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Folder"}:                          true,
	{Group: "bigqueryconnection.cnrm.cloud.google.com", Kind: "BigQueryConnectionConnection"}: true,
	{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreBackupSchedule"}:               true,
	{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIDataset"}:                        true,
	{Group: "workflowexecutions.cnrm.cloud.google.com", Kind: "WorkflowsExecution"}:           true,
	{Group: "workstations.cnrm.cloud.google.com", Kind: "WorkstationConfig"}:                  true,
	{Group: "cloudbuild.cnrm.cloud.google.com", Kind: "CloudBuildWorkerPool"}:                 true,
	{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreIndex"}:                        true,
	{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringUptimeCheckConfig"}:          true,
	{Group: "container.cnrm.cloud.google.com", Kind: "ContainerCluster"}:                      true,
	{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogBucket"}:                        true,
}

func GetCAISIdentities(ctx context.Context, scheme *runtime.Scheme, reader client.Reader, objectsList []*unstructured.Unstructured) ([]CAISIdentityResult, error) {
	var results []CAISIdentityResult

	for _, u := range objectsList {
		gvk := u.GroupVersionKind()

		res := CAISIdentityResult{
			Group:     gvk.Group,
			Version:   gvk.Version,
			Kind:      gvk.Kind,
			Namespace: u.GetNamespace(),
			Name:      u.GetName(),
			CAISURL:   "unknown",
		}

		gk := gvk.GroupKind()
		if unsupportedKinds[gk] {
			res.Error = "not yet supported"
			results = append(results, res)
			continue
		}

		var obj runtime.Object
		var err error
		if obj, err = kccscheme.NewObject(gk); err != nil {
			results = append(results, res)
			continue
		}

		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			res.Error = fmt.Sprintf("failed to convert unstructured object to %T: %s", obj, err.Error())
			results = append(results, res)
			continue
		}

		if resource, ok := obj.(identity.Resource); ok {
			id, err := resource.GetIdentity(ctx, reader)
			if err != nil {
				res.Error = err.Error()
				results = append(results, res)
				continue
			}
			if id == nil {
				continue
			}

			hasIdentity := true
			if sgId, ok := id.(identity.ServerGeneratedIdentity); ok {
				hasIdentity = sgId.HasIdentitySpecified()
			}

			if hasIdentity {
				if idV2, ok := id.(identity.IdentityV2); ok {
					host := idV2.Host()
					res.CAISURL = fmt.Sprintf("//%s/%s", host, id.String())
				}
			}
			results = append(results, res)
		} else {
			results = append(results, res)
		}
	}

	return results, nil
}

type InMemoryReader struct {
	scheme  *runtime.Scheme
	objects []*unstructured.Unstructured
}

func NewInMemoryReader(scheme *runtime.Scheme, objects []*unstructured.Unstructured) *InMemoryReader {
	return &InMemoryReader{
		scheme:  scheme,
		objects: objects,
	}
}

func (r *InMemoryReader) getGVK(obj runtime.Object) (schema.GroupVersionKind, error) {
	gvks, err := kccscheme.ObjectKinds(obj)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	if len(gvks) == 0 {
		return schema.GroupVersionKind{}, fmt.Errorf("could not determine GVK for %T", obj)
	}
	return gvks[0], nil
}

func (r *InMemoryReader) find(gvk schema.GroupVersionKind, namespace, name string) *unstructured.Unstructured {
	for _, obj := range r.objects {
		ogvk := obj.GroupVersionKind()
		if ogvk.Kind == gvk.Kind && ogvk.Group == gvk.Group {
			if obj.GetNamespace() == namespace && obj.GetName() == name {
				return obj
			}
		}
	}
	return nil
}

func (r *InMemoryReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	gvk, err := r.getGVK(obj)
	if err != nil {
		return err
	}

	found := r.find(gvk, key.Namespace, key.Name)
	if found == nil {
		return apierrors.NewNotFound(schema.GroupResource{Group: gvk.Group, Resource: gvk.Kind}, key.Name)
	}

	if u, ok := obj.(*unstructured.Unstructured); ok {
		u.Object = runtime.DeepCopyJSON(found.Object)
		return nil
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(found.Object, obj)
	if err != nil {
		return fmt.Errorf("failed to convert unstructured object to %T: %w", obj, err)
	}
	return nil
}

func (r *InMemoryReader) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return fmt.Errorf("InMemoryReader.List not implemented")
}

func ReadObjects(stdin bool, file string, directory string) ([]*unstructured.Unstructured, error) {
	var objectsList []*unstructured.Unstructured

	if stdin {
		objs, err := objects.ParseObjectsFromStream(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to read from stdin: %w", err)
		}
		objectsList = append(objectsList, objs...)
	}

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %q: %w", file, err)
		}
		defer f.Close()
		objs, err := objects.ParseObjectsFromStream(f)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file %q: %w", file, err)
		}
		objectsList = append(objectsList, objs...)
	}

	if directory != "" {
		err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".yaml" || ext == ".yml" {
				f, err := os.Open(path)
				if err != nil {
					return fmt.Errorf("failed to open file %q: %w", path, err)
				}
				defer f.Close()
				objs, err := objects.ParseObjectsFromStream(f)
				if err != nil {
					return fmt.Errorf("failed to parse file %q: %w", path, err)
				}
				objectsList = append(objectsList, objs...)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to walk directory %q: %w", directory, err)
		}
	}

	return objectsList, nil
}
