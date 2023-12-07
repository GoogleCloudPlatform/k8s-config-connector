// Copyright 2022 Google LLC
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

package mocktests

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	yamlserializer "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
)

type Harness struct {
	*testing.T

	k8s        *mockkubeapiserver.MockKubeAPIServer
	restConfig *rest.Config
	Scheme     *runtime.Scheme
	Ctx        context.Context
	Client     client.Client
}

func (h *Harness) RESTConfig() *rest.Config {
	if h.restConfig == nil {
		h.Fatalf("cannot call RESTConfig before Start")
	}
	return h.restConfig
}

func (h *Harness) NewClient(config *rest.Config, options client.Options) (client.Client, error) {
	if h.Client == nil {
		h.Fatalf("WithObjects must be called before NewClient")
	}
	return h.Client, nil
}

func NewHarness(t *testing.T) *Harness {
	h := &Harness{
		T:      t,
		Scheme: runtime.NewScheme(),
		Ctx:    context.Background(),
	}
	corev1.AddToScheme(h.Scheme)

	iamv1beta1.SchemeBuilder.AddToScheme(h.Scheme)

	t.Cleanup(h.Stop)
	return h
}

func (h *Harness) ParseObjects(y string) []*unstructured.Unstructured {
	t := h.T

	var objects []*unstructured.Unstructured

	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(y)), 100)
	for {
		var rawObj runtime.RawExtension
		if err := decoder.Decode(&rawObj); err != nil {
			if err != io.EOF {
				t.Fatalf("error decoding yaml: %v", err)
			}
			break
		}

		m, _, err := yamlserializer.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			t.Fatalf("error decoding yaml: %v", err)
		}

		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(m)
		if err != nil {
			t.Fatalf("error parsing object: %v", err)
		}
		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		objects = append(objects, unstructuredObj)
	}

	return objects
}

func (h *Harness) WithObjects(initObjs ...*unstructured.Unstructured) {
	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		h.Fatalf("error building mock kube-apiserver: %v", err)
	}

	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}, "namespaces", meta.RESTScopeRoot)
	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}, "secrets", meta.RESTScopeNamespace)

	smLoader, err := servicemappingloader.New()
	if err != nil {
		h.Fatalf("error getting new service mapping loader: %v", err)
	}
	supportedGVKs := supportedgvks.All(smLoader, dclmetadata.New())
	for _, gvk := range supportedGVKs {
		var resource string
		switch gvk.Kind {
		// TODO: Any special pluralization cases go here (unless we can get them from supportedgvks)
		default:
			resource = strings.ToLower(gvk.Kind) + "s"
		}
		k8s.RegisterType(gvk, resource, meta.RESTScopeNamespace)
	}

	for _, obj := range initObjs {
		if err := k8s.AddObject(obj); err != nil {
			h.Errorf("error adding object %v: %v", obj, err)
		}
	}
	addr, err := k8s.StartServing()
	if err != nil {
		h.Errorf("error starting mock kube-apiserver: %v", err)
	}

	h.restConfig = &rest.Config{
		Host: addr.String(),
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
	}

	client, err := client.New(h.restConfig, client.Options{})
	if err != nil {
		h.Fatalf("error building client: %v", err)
	}

	h.Client = client
}

func (h *Harness) Stop() {
	if h.k8s != nil {
		if err := h.k8s.Stop(); err != nil {
			h.Logf("error closing mock kube-apiserver: %v", err)
			h.Errorf("error closing mock kube-apiserver: %v", err)
		}
	}
}

// MustReadFile returns the contents of the file - as a string
// It fails the test if the file cannot be read
func (h *Harness) MustReadFile(p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		h.Fatalf("error reading file %q: %v", p, err)
	}
	return string(b)
}
