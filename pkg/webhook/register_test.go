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

package webhook

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestWebhookManifests(t *testing.T) {
	commonManifest, err := BuildCommonWebhooks()
	if err != nil {
		t.Fatalf("error building common webhooks: %v", err)
	}

	abandonOnUninstallManifest, err := BuildAbandonOnUninstallWebhookManifest()
	if err != nil {
		t.Fatalf("error building abandon on uninstall webhook manifest: %v", err)
	}

	var allObjects []*unstructured.Unstructured
	{
		objects, err := commonManifest.ToUnstructured()
		if err != nil {
			t.Fatalf("error converting common manifest to unstructured: %v", err)
		}
		allObjects = append(allObjects, objects...)
	}
	{
		objects, err := abandonOnUninstallManifest.ToUnstructured()
		if err != nil {
			t.Fatalf("error converting abandon on uninstall manifest to unstructured: %v", err)
		}
		allObjects = append(allObjects, objects...)
	}

	repoRoot := repoRoot(t)

	version := ""
	{
		b, err := os.ReadFile(filepath.Join(repoRoot, "version", "VERSION"))
		if err != nil {
			t.Fatalf("error reading version file: %v", err)
		}
		version = string(b)
		version = strings.TrimSpace(version)
	}

	got, err := ObjectsToYAML(allObjects)
	if err != nil {
		t.Fatalf("error from ObjectsToYAML: %v", err)
	}

	for _, channels := range []string{"channels", "autopilot-channels"} {
		for _, version := range []string{version} {
			p := filepath.Join(repoRoot, "operator", channels, "packages", "configconnector", version, "webhooks.yaml")

			test.CompareGoldenFile(t, p, got)
		}
	}
}

func repoRoot(t *testing.T) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("error getting repo root: %v", err)
	}
	repoRoot := strings.TrimSpace(string(output))
	return repoRoot
}
