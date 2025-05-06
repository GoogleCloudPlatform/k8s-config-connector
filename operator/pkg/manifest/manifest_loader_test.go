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
	"reflect"
	"strings"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestManifestLoader_ResolveManifest(t *testing.T) {
	t.Parallel()
	basedir, repo := newTestNewLocalRepository()
	ml := manifest.NewLoader(repo)
	tests := []struct {
		name   string
		cc     *corev1beta1.ConfigConnector
		result map[string]string
	}{
		{
			name: "cluster mode, workload identity",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "cluster", "workload-identity"}, "/"): clusterModeWIManifest},
		},
		{
			name: "cluster mode, gcp identity",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "cluster", "gcp-identity"}, "/"): clusterModeGcpManifest},
		},
		{
			name: "namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "namespaced",
				},
			},
			result: map[string]string{strings.Join([]string{basedir, "packages", "configconnector", "0.0.0-test", "namespaced"}, "/"): namespacedManifest},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			manifestStr, err := ml.ResolveManifest(context.TODO(), tc.cc)
			if err != nil {
				t.Fatalf("unexpected error while loadding the manifest for namespaced components: %v", err)
			}
			if !reflect.DeepEqual(manifestStr, tc.result) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(manifestStr, tc.result))
			}
		})
	}
}
