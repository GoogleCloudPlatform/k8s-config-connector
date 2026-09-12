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

package configconnectorcontext

import (
	"strings"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

// statefulSetWithManager returns a minimal controller manager StatefulSet, with
// the same shape the per-namespace manifest has.
func statefulSetWithManager(t *testing.T) *unstructured.Unstructured {
	t.Helper()

	const manifest = `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: cnrm-controller-manager
  namespace: cnrm-system
spec:
  template:
    spec:
      containers:
      - name: prom-to-sd
        image: prom-to-sd
      - name: manager
        image: manager
        args: ["--scoped-namespace=foo", "--stderrthreshold=INFO"]
`
	obj := map[string]interface{}{}
	if err := yaml.Unmarshal([]byte(manifest), &obj); err != nil {
		t.Fatalf("error parsing manifest: %v", err)
	}
	return &unstructured.Unstructured{Object: obj}
}

func managerArgs(t *testing.T, u *unstructured.Unstructured) []string {
	t.Helper()

	containers, found, err := unstructured.NestedSlice(u.Object, "spec", "template", "spec", "containers")
	if err != nil || !found {
		t.Fatalf("error resolving containers: %v", err)
	}
	for _, c := range containers {
		container, ok := c.(map[string]interface{})
		if !ok {
			t.Fatalf("container was not a map: %v", c)
		}
		name, _, _ := unstructured.NestedString(container, "name")
		if name != "manager" {
			continue
		}
		args, _, _ := unstructured.NestedStringSlice(container, "args")
		return args
	}
	t.Fatal("no manager container found")
	return nil
}

func hasArgWithPrefix(args []string, prefix string) bool {
	for _, a := range args {
		if strings.HasPrefix(a, prefix) {
			return true
		}
	}
	return false
}

func containsArg(args []string, want string) bool {
	for _, a := range args {
		if a == want {
			return true
		}
	}
	return false
}

func TestApplyUniverseToManagerContainer(t *testing.T) {
	tests := []struct {
		name       string
		universe   *corev1beta1.UniverseSpec
		wantDomain string
		wantPrefix string
	}{
		{
			name:     "unset_noUniverseFlags",
			universe: nil,
		},
		{
			name: "set_bothFlagsApplied",
			universe: &corev1beta1.UniverseSpec{
				Domain: "example-apis.test",
				Prefix: "zzz",
			},
			wantDomain: "--universe-domain=example-apis.test",
			wantPrefix: "--universe-prefix=zzz",
		},
		{
			// The domain and the prefix are independent values. A pair with no
			// shared stem catches an implementation that derives one from the
			// other.
			name: "domainAndPrefixAreIndependent",
			universe: &corev1beta1.UniverseSpec{
				Domain: "sovereign.example",
				Prefix: "abc",
			},
			wantDomain: "--universe-domain=sovereign.example",
			wantPrefix: "--universe-prefix=abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := statefulSetWithManager(t)
			cc := &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{Universe: tc.universe},
			}

			if err := applyUniverseToManagerContainer(u, cc); err != nil {
				t.Fatalf("applyUniverseToManagerContainer failed: %v", err)
			}

			args := managerArgs(t, u)

			// Pre-existing args must survive in every case.
			if !containsArg(args, "--scoped-namespace=foo") {
				t.Errorf("pre-existing arg was lost, got args %v", args)
			}

			if tc.universe == nil {
				if hasArgWithPrefix(args, "--universe-") {
					t.Errorf("did not expect universe flags, got args %v", args)
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

// TestApplyUniverseToManagerContainer_isIdempotent guards against the flags
// accumulating. The operator re-runs this transform on every reconcile, and
// setFlagForManagerContainer is expected to replace rather than append.
func TestApplyUniverseToManagerContainer_isIdempotent(t *testing.T) {
	u := statefulSetWithManager(t)
	cc := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Universe: &corev1beta1.UniverseSpec{Domain: "example-apis.test", Prefix: "zzz"},
		},
	}

	for i := 0; i < 3; i++ {
		if err := applyUniverseToManagerContainer(u, cc); err != nil {
			t.Fatalf("applyUniverseToManagerContainer failed on pass %d: %v", i, err)
		}
	}

	args := managerArgs(t, u)
	domainCount, prefixCount := 0, 0
	for _, a := range args {
		if strings.HasPrefix(a, "--universe-domain") {
			domainCount++
		}
		if strings.HasPrefix(a, "--universe-prefix") {
			prefixCount++
		}
	}
	if domainCount != 1 {
		t.Errorf("--universe-domain appears %d times after 3 applications, want 1 (args %v)", domainCount, args)
	}
	if prefixCount != 1 {
		t.Errorf("--universe-prefix appears %d times after 3 applications, want 1 (args %v)", prefixCount, args)
	}
}
