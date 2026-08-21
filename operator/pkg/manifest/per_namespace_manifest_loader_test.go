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

package manifest_test

import (
	"context"
	"path"
	"reflect"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestManifestLoader_ResolveNamespacedComponents(t *testing.T) {
	t.Parallel()
	baseDir, repo := newTestNewLocalRepository()
	manifestPath := path.Join(baseDir, "packages/configconnector/0.0.0-test/namespaced/per-namespace-components.yaml")
	ml := manifest.NewPerNamespaceManifestLoader(repo)
	tests := []struct {
		name   string
		ccc    *corev1beta1.ConfigConnectorContext
		result map[string]string
	}{
		{
			name: "resolve namespaced component",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorContextAllowedName,
				},
			},
			result: map[string]string{manifestPath: namespacedComponentsOnly},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			m, err := ml.ResolveManifest(ctx, tc.ccc)
			if err != nil {
				t.Fatalf("unexpected error while loadding the manifest for namespaced components: %v", err)
			}

			if !reflect.DeepEqual(m, tc.result) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(m, tc.result))
			}
		})
	}
}
