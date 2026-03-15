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
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestCommonWebhookManifest(t *testing.T) {
	webhookManifest, err := BuildCommonWebhooks()
	if err != nil {
		t.Fatalf("error building common webhooks: %v", err)
	}

	got, err := webhookManifest.ToYAML()
	if err != nil {
		t.Fatalf("error converting manifest to yaml: %v", err)
	}

	p := filepath.Join("testdata", "webhooks", "common", "manifest.yaml")
	test.CompareGoldenFile(t, p, got)
}

func TestAbandonOnUninstallWebhookManifest(t *testing.T) {
	webhookManifest, err := BuildAbandonOnUninstallWebhookManifest()
	if err != nil {
		t.Fatalf("error building abandon on uninstall webhook manifest: %v", err)
	}

	got, err := webhookManifest.ToYAML()
	if err != nil {
		t.Fatalf("error converting manifest to yaml: %v", err)
	}
	p := filepath.Join("testdata", "webhooks", "abandon-on-uninstall", "manifest.yaml")
	test.CompareGoldenFile(t, p, got)
}
