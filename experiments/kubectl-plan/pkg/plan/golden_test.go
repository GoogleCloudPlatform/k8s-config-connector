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

package plan

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/yaml"
)

func TestPlanner(t *testing.T) {

	dir := "testdata"
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read directory %q: %v", dir, err)
	}
	for _, file := range files {
		p := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			t.Errorf("found non-directory %q", p)
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			p := filepath.Join(dir, file.Name())

			ctx := context.Background()

			testEnv := &envtest.Environment{
				// CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
			}

			restConfig, err := testEnv.Start()
			if err != nil {
				t.Fatalf("error starting testenv: %v", err)
			}

			defer func() {
				if err := testEnv.Stop(); err != nil {
					t.Errorf("error closing mock kube-apiserver: %v", err)
				}
			}()

			httpClient, err := rest.HTTPClientFor(restConfig)
			if err != nil {
				t.Fatalf("error from HTTPClientFor: %v", err)
			}

			kubeClient, err := client.New(restConfig, client.Options{
				HTTPClient: httpClient,
			})
			if err != nil {
				t.Fatalf("building kubeClient: %v", err)
			}

			existingObjects, err := LoadObjectsFromFilesystem(filepath.Join(p, "existing.yaml"))
			if err != nil {
				t.Fatalf("error loading objects: %v", err)
			}
			applyObjects, err := LoadObjectsFromFilesystem(filepath.Join(p, "apply.yaml"))
			if err != nil {
				t.Fatalf("error loading objects: %v", err)
			}

			target, err := NewClusterTarget(restConfig, httpClient)
			if err != nil {
				t.Fatalf("error building target: %v", err)
			}

			ensureNamespace(ctx, t, kubeClient, "default")

			for _, existingObject := range existingObjects {
				if err := kubeClient.Create(ctx, existingObject); err != nil {
					t.Fatalf("error pre-creating existing object: %v", err)
				}
			}

			planner := &Planner{}

			plan, err := planner.BuildPlan(ctx, applyObjects, target)
			if err != nil {
				t.Fatalf("error from BuildPlan: %v", err)
			}

			actual, err := yaml.Marshal(plan)
			if err != nil {
				t.Fatalf("yaml.Marshal failed: %v", err)
			}
			CompareGoldenFile(t, filepath.Join(p, "_plan.yaml"), actual)

			var pretty bytes.Buffer
			if err := printPlan(ctx, plan, &pretty); err != nil {
				t.Fatalf("printPlan failed: %v", err)
			}
			CompareGoldenFile(t, filepath.Join(p, "_plan.txt"), []byte(pretty.String()))
		})
	}
}

func CompareGoldenFile(t *testing.T, p string, got []byte) {
	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		// Short-circuit when the output is correct
		b, err := os.ReadFile(p)
		if err == nil && bytes.Equal(b, got) {
			return
		}

		if err := os.WriteFile(p, got, 0644); err != nil {
			t.Fatalf("failed to write golden output %s: %v", p, err)
		}
		t.Errorf("wrote output to %s", p)
	} else {
		want, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("failed to read file %q: %v", p, err)
		}
		if diff := cmp.Diff(string(want), string(got)); diff != "" {
			t.Errorf("unexpected diff in %s: %s", p, diff)
		}
	}
}

func ensureNamespace(ctx context.Context, t *testing.T, kubeclient client.Client, namespace string) {
	id := types.NamespacedName{Name: namespace}
	gvk := schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}

	existing := &unstructured.Unstructured{}
	existing.SetGroupVersionKind(gvk)
	if err := kubeclient.Get(ctx, id, existing); err != nil {
		if !apierrors.IsNotFound(err) {
			t.Fatalf("checking that namespace exists: %v", err)
		}
	} else {
		// Namespace exists
		return
	}

	ns := &unstructured.Unstructured{}
	ns.SetGroupVersionKind(gvk)
	ns.SetName(id.Name)
	ns.SetNamespace(id.Namespace)

	if err := kubeclient.Create(ctx, ns); err != nil {
		t.Fatalf("creating namespace: %v", err)
	}

}
