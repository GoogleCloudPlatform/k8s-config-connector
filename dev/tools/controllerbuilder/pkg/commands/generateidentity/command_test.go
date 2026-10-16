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

package generateidentity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
)

func TestLoadAndApplyConfigPreservesExplicitFlags(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "service.yaml")
	cfgYAML := `service: mockgcp.cloud.widget.v1
apiVersion: widget.cnrm.cloud.google.com/v1alpha1
resources:
  - kind: WidgetAlpha
    protoName: Alpha
  - kind: WidgetBeta
    protoName: Beta
`
	if err := os.WriteFile(cfgPath, []byte(cfgYAML), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	opt := &GenerateIdentityOptions{
		GenerateOptions: &options.GenerateOptions{
			ConfigFilePath:  cfgPath,
			ProtoSourcePath: "some.pb",
		},
		ServiceName: "google.cloud.widget.v1",
		Resources: options.ResourceList{
			{Kind: "WidgetAlpha", ProtoName: "Alpha"},
		},
	}
	if err := opt.loadAndApplyConfig(); err != nil {
		t.Fatalf("loadAndApplyConfig: %v", err)
	}
	if opt.ServiceName != "google.cloud.widget.v1" {
		t.Errorf("ServiceName = %q, want CLI flag value %q", opt.ServiceName, "google.cloud.widget.v1")
	}
	if opt.APIVersion != "widget.cnrm.cloud.google.com/v1alpha1" {
		t.Errorf("APIVersion = %q, want config fallback %q", opt.APIVersion, "widget.cnrm.cloud.google.com/v1alpha1")
	}
	if len(opt.Resources) != 1 || opt.Resources[0].Kind != "WidgetAlpha" {
		t.Errorf("Resources = %+v, want only explicit CLI resource WidgetAlpha", opt.Resources)
	}
}

func TestValidateRejectsGroupLessAPIVersion(t *testing.T) {
	opt := &GenerateIdentityOptions{
		GenerateOptions: &options.GenerateOptions{
			APIVersion:      "v1alpha1",
			ProtoSourcePath: "some.pb",
		},
		ServiceName: "google.cloud.widget.v1",
		Resources: options.ResourceList{
			{Kind: "Widget", ProtoName: "Widget"},
		},
	}
	err := opt.validate()
	if err == nil || !strings.Contains(err.Error(), "Group/Version") {
		t.Fatalf("validate() error = %v, want Group/Version error", err)
	}
}
