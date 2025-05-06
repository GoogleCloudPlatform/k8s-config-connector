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

package kccconfig

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
)

func init() {
	s := scheme.Scheme
	if err := corev1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering core kcc operator scheme: %v", err)
	}
}

func TestFetchLiveKCCState(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		cc             *corev1beta1.ConfigConnector
		ccc            *corev1beta1.ConfigConnectorContext
		namespacedName types.NamespacedName
		expectError    bool
	}{
		{
			name: "fetch live KCC state in cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "cluster",
				},
			},
		},
		{
			name: "fetch live KCC state in namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
			},
			namespacedName: types.NamespacedName{Namespace: "foo-ns"},
		},
		{
			name: "ccc not found in namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			namespacedName: types.NamespacedName{Namespace: "foo-ns"},
			expectError:    true,
		},
		{
			name:        "no cc or ccc",
			expectError: false,
		},
		{
			// This should not be a valid use case in prod env. The behavior is
			// identical to the "no cc or ccc" case.
			name: "no cc",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
			},
			namespacedName: types.NamespacedName{Namespace: "foo-ns"},
			expectError:    false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			client := mgr.GetClient()
			if tc.cc != nil {
				if err := client.Create(ctx, tc.cc); err != nil {
					t.Fatalf("error creating %+v: %v", tc.cc.GroupVersionKind(), err)
				}
			}
			if tc.ccc != nil {
				ns := &corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: tc.ccc.GetNamespace(),
					},
				}
				if err := client.Create(ctx, ns); err != nil {
					t.Fatalf("error creating %+v: %v", ns.GroupVersionKind(), err)
				}
				if err := client.Create(ctx, tc.ccc); err != nil {
					t.Fatalf("error creating %+v: %v", tc.ccc.GroupVersionKind(), err)
				}
			}
			gotCC, gotCCC, err := FetchLiveKCCState(ctx, client, tc.namespacedName)
			if tc.expectError {
				if err == nil {
					t.Fatalf("got nil, but expect an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.cc != nil {
				if !reflect.DeepEqual(gotCC.Spec, tc.cc.Spec) {
					t.Errorf("unexpected diff: %v", cmp.Diff(gotCC.Spec, tc.cc.Spec))
				}
			}
			if tc.ccc != nil {
				if !reflect.DeepEqual(gotCCC.Spec, tc.ccc.Spec) {
					t.Errorf("unexpected diff: %v", cmp.Diff(gotCCC.Spec, tc.ccc.Spec))
				}
			}
		})
	}
}
