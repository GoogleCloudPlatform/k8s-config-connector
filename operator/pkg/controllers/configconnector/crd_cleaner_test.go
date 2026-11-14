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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/environment"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
)

func TestApplyUniverseDomainAndFilterManifestCRDs(t *testing.T) {
	ctx := t.Context()
	m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())

	env := environment.Config{
		Type:           environment.SovereignCloud,
		UniverseDomain: "apis-eu0.sovereign.goog",
		SupportedServices: map[string]bool{
			"compute": true,
			"storage": true,
		},
	}

	if err := controllers.ApplyUniverseDomain(m, env.UniverseDomain); err != nil {
		t.Fatalf("ApplyUniverseDomain failed: %v", err)
	}
	FilterManifestCRDs(m, env)

	foundWorkloadWithEnv := false
	for _, item := range m.Items {
		if item.Kind == "StatefulSet" || item.Kind == "Deployment" {
			val := getEnvVarValue(t, item.UnstructuredObject(), "GOOGLE_CLOUD_UNIVERSE_DOMAIN")
			if val != env.UniverseDomain {
				t.Errorf("%s/%s GOOGLE_CLOUD_UNIVERSE_DOMAIN = %q, want %q", item.Kind, item.GetName(), val, env.UniverseDomain)
			}
			foundWorkloadWithEnv = true
		}
		if item.Kind == "CustomResourceDefinition" && isKCCCRD(item) {
			if !env.IsCRDGroupSupported(item.UnstructuredObject().GetName()) {
				t.Errorf("unexpected unsupported CRD remaining in manifest: %s", item.UnstructuredObject().GetName())
			}
		}
	}
	if !foundWorkloadWithEnv {
		t.Fatalf("expected at least one StatefulSet or Deployment with GOOGLE_CLOUD_UNIVERSE_DOMAIN injected")
	}
}
