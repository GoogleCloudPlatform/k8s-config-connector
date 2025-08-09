// Copyright 2025 Google LLC
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

package test

import (
	"os"
	"path/filepath"
	"testing"

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/golden"
	"sigs.k8s.io/yaml"
)

// Harness is a simple test harness for our operator
type Harness struct {
	T *testing.T
}

// NewHarness constructs a test harness
func NewHarness(t *testing.T) *Harness {
	return &Harness{T: t}
}

// RepoPath is the path to the channels directory
func (h *Harness) RepoPath() string {
	repoPath := "../../../channels"
	return repoPath
}

type Validator interface {
	Validate(reconcilerFactory func(mgr manager.Manager) (*declarative.Reconciler, error))
}

// KDPValidator builds a KDP validator, for use in golden tests.
func (h *Harness) KDPValidator(rewriteObjects func(o *unstructured.Unstructured)) Validator {
	t := h.T

	env := &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths: []string{
				"../../../config/crd/bases",
			},
			ErrorIfPathMissing: true,
		},
	}

	replacements := make(map[string]string)
	rewriteObjectsWithReplacments := func(u *unstructured.Unstructured) {
		rewriteObjects(u)

		ApplyReplacementsToMap(u.Object, replacements)
	}

	goldenOptions := golden.ValidatorOptions{
		RewriteObjects:     rewriteObjectsWithReplacments,
		EnvtestEnvironment: env,
		ManagerOptions: manager.Options{
			Metrics: metricsserver.Options{BindAddress: "0"},
		},
	}
	goldenOptions.WithSchema(v1beta1.AddToScheme, customizev1alpha1.AddToScheme, customizev1beta1.AddToScheme, corev1.AddToScheme, appsv1.AddToScheme)

	v := golden.NewValidator(h.T, goldenOptions)

	// Replace latest value with {{stable.latest}}
	{
		type manifest struct {
			Version string `json:"version"`
		}
		type channel struct {
			Manifests []manifest `json:"manifests"`
		}
		repoPath := h.RepoPath()
		channelPath := filepath.Join(repoPath, "stable")
		b, err := os.ReadFile(channelPath)
		if err != nil {
			t.Fatalf("reading channel %q: %v", channelPath, err)
		}
		var c channel
		if err := yaml.Unmarshal(b, &c); err != nil {
			t.Fatalf("unmarshaling channel %q: %v", channelPath, err)
		}
		if len(c.Manifests) == 0 {
			t.Fatalf("no manifests in channel %q", channelPath)
		}
		if len(c.Manifests) > 1 {
			t.Fatalf("more than one manifest in channel %q", channelPath)
		}
		latestVersion := c.Manifests[0].Version
		replacements[latestVersion] = "{{stable.latest}}"
	}

	return v
}
