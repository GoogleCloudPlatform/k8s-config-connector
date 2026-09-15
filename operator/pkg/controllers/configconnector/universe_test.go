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

package configconnector

import (
	"context"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// managerArgs returns the args of the manager container of the controller
// manager StatefulSet in m.
func managerArgs(t *testing.T, m *manifest.Objects) []string {
	t.Helper()

	var found []string
	statefulSets := 0
	for _, item := range m.Items {
		if !IsControllerManagerStatefulSet(item) {
			continue
		}
		statefulSets++
		if err := item.MutateContainers(func(container map[string]interface{}) error {
			name, _, _ := unstructured.NestedString(container, "name")
			if name != "manager" {
				return nil
			}
			args, _, _ := unstructured.NestedStringSlice(container, "args")
			found = args
			return nil
		}); err != nil {
			t.Fatalf("MutateContainers failed: %v", err)
		}
	}
	if statefulSets == 0 {
		t.Fatal("no controller manager StatefulSet found in manifest")
	}
	return found
}

func containsArg(args []string, want string) bool {
	for _, a := range args {
		if a == want {
			return true
		}
	}
	return false
}

func hasArgWithPrefix(args []string, prefix string) bool {
	for _, a := range args {
		if len(a) >= len(prefix) && a[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

func TestTransformForUniverse(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		universe     *corev1beta1.UniverseSpec
		wantDomain   string
		wantPrefix   string
		wantAnyFlags bool
	}{
		{
			name:         "unset_noUniverseFlags",
			universe:     nil,
			wantAnyFlags: false,
		},
		{
			name: "set_bothFlagsApplied",
			universe: &corev1beta1.UniverseSpec{
				Domain: "example-apis.test",
				Prefix: "zzz",
			},
			wantDomain:   "--universe-domain=example-apis.test",
			wantPrefix:   "--universe-prefix=zzz",
			wantAnyFlags: true,
		},
		{
			// The domain and the prefix are independent values. Using a pair
			// with no shared stem means an implementation that derives one from
			// the other fails here.
			name: "domainAndPrefixAreIndependent",
			universe: &corev1beta1.UniverseSpec{
				Domain: "sovereign.example",
				Prefix: "abc",
			},
			wantDomain:   "--universe-domain=sovereign.example",
			wantPrefix:   "--universe-prefix=abc",
			wantAnyFlags: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := testcontroller.ParseObjects(ctx, t, testcontroller.ClusterModeComponents)
			cc := &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{Universe: tc.universe},
			}

			r := &Reconciler{}
			if err := r.transformForUniverse()(ctx, cc, m); err != nil {
				t.Fatalf("transformForUniverse failed: %v", err)
			}

			args := managerArgs(t, m)

			if !tc.wantAnyFlags {
				if hasArgWithPrefix(args, "--universe-domain") {
					t.Errorf("did not expect --universe-domain, got args %v", args)
				}
				if hasArgWithPrefix(args, "--universe-prefix") {
					t.Errorf("did not expect --universe-prefix, got args %v", args)
				}
				return
			}

			if !containsArg(args, tc.wantDomain) {
				t.Errorf("expected arg %q, got args %v", tc.wantDomain, args)
			}
			if !containsArg(args, tc.wantPrefix) {
				t.Errorf("expected arg %q, got args %v", tc.wantPrefix, args)
			}
		})
	}
}
