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

package controllers

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// restClientPathApplier builds a REST client for its object the same way
// k8s.io/cli-runtime/pkg/resource does, rewriting the rest.Config it is given,
// and records the path a GET of the object would be sent to.
type restClientPathApplier struct {
	barrier sync.WaitGroup

	mu    sync.Mutex
	paths map[string]string
}

func (f *restClientPathApplier) Apply(ctx context.Context, opt applier.ApplierOptions) error {
	obj := opt.Objects[0]
	gvk := obj.GroupVersionKind()
	gv := gvk.GroupVersion()

	cfg := opt.RESTConfig
	cfg.GroupVersion = &gv
	if gv.Group == "" {
		cfg.APIPath = "/api"
	} else {
		cfg.APIPath = "/apis"
	}

	// Let every concurrent Apply set up its config before any of them builds a client.
	f.barrier.Done()
	f.barrier.Wait()

	client, err := rest.RESTClientFor(cfg)
	if err != nil {
		return err
	}
	plural, _ := meta.UnsafeGuessKindToResource(gvk)
	path := client.Get().
		NamespaceIfScoped(obj.GetNamespace(), obj.GetNamespace() != "").
		Resource(plural.Resource).
		Name(obj.GetName()).
		URL().Path

	f.mu.Lock()
	defer f.mu.Unlock()
	f.paths[obj.GetName()] = path
	return nil
}

func TestApplierDoesNotShareRESTConfig(t *testing.T) {
	objects := []*manifest.Object{
		newTestObject(t, "v1", "ServiceAccount", "cnrm-system", "cnrm-controller-manager"),
		newTestObject(t, "apiextensions.k8s.io/v1", "CustomResourceDefinition", "", "folders.resourcemanager.cnrm.cloud.google.com"),
	}
	want := map[string]string{
		"cnrm-controller-manager":                       "/api/v1/namespaces/cnrm-system/serviceaccounts/cnrm-controller-manager",
		"folders.resourcemanager.cnrm.cloud.google.com": "/apis/apiextensions.k8s.io/v1/customresourcedefinitions/folders.resourcemanager.cnrm.cloud.google.com",
	}

	cfg := &rest.Config{Host: "https://kubernetes.example.com"}
	cfg.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	inner := &restClientPathApplier{paths: map[string]string{}}
	inner.barrier.Add(len(objects))
	a := &copyConfigApplier{inner: inner}

	var wg sync.WaitGroup
	errs := make(chan error, len(objects))
	for _, obj := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs <- a.Apply(context.Background(), applier.ApplierOptions{
				RESTConfig: cfg,
				Objects:    []*manifest.Object{obj},
			})
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("Apply() error: %v", err)
		}
	}

	if !reflect.DeepEqual(inner.paths, want) {
		t.Errorf("GET paths = %v, want %v", inner.paths, want)
	}
	if cfg.GroupVersion != nil || cfg.APIPath != "" {
		t.Errorf("shared rest.Config was modified: GroupVersion=%v, APIPath=%q", cfg.GroupVersion, cfg.APIPath)
	}
}

func newTestObject(t *testing.T, apiVersion, kind, namespace, name string) *manifest.Object {
	t.Helper()
	u := &unstructured.Unstructured{}
	u.SetAPIVersion(apiVersion)
	u.SetKind(kind)
	u.SetNamespace(namespace)
	u.SetName(name)
	obj, err := manifest.NewObject(u)
	if err != nil {
		t.Fatalf("NewObject(%s %s/%s): %v", kind, namespace, name, err)
	}
	return obj
}
