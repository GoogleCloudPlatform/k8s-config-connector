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

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RepoRoot returns the root of the git repository.
func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	repoRoot := strings.TrimSpace(string(output))
	return repoRoot, nil
}

// Run runs the main logic of the webhook manifest generator.
func Run() error {
	repoRoot, err := RepoRoot()
	if err != nil {
		return fmt.Errorf("error getting repo root: %v", err)
	}
	flag.StringVar(&repoRoot, "repo-root", repoRoot, "path to the root of the k8s-config-connector repo")
	flag.Parse()

	var allObjects []*unstructured.Unstructured

	{
		webhookManifest, err := webhook.BuildCommonWebhooks()
		if err != nil {
			return fmt.Errorf("error building common webhooks: %v", err)
		}

		objects, err := webhookManifest.ToUnstructured()
		if err != nil {
			return fmt.Errorf("error converting manifest to unstructured: %v", err)
		}
		allObjects = append(allObjects, objects...)
	}

	{
		webhookManifest, err := webhook.BuildAbandonOnUninstallWebhookManifest()
		if err != nil {
			return fmt.Errorf("error building abandon on uninstall webhook manifest: %v", err)
		}

		objects, err := webhookManifest.ToUnstructured()
		if err != nil {
			return fmt.Errorf("error converting manifest to unstructured: %v", err)
		}
		allObjects = append(allObjects, objects...)
	}

	version := ""
	{
		b, err := os.ReadFile(filepath.Join(repoRoot, "version", "VERSION"))
		if err != nil {
			return fmt.Errorf("error reading version file: %v", err)
		}
		version = string(b)
		version = strings.TrimSpace(version)
	}

	for _, channels := range []string{"channels", "autopilot-channels"} {
		for _, version := range []string{version} {
			p := filepath.Join(repoRoot, "operator", channels, "packages", "configconnector", version, "webhooks.yaml")

			var buffer bytes.Buffer
			for i, obj := range allObjects {
				if i != 0 {
					buffer.WriteString("\n---\n")
				}
				v, err := yaml.Marshal(obj)
				if err != nil {
					return fmt.Errorf("error converting object to yaml: %w", err)
				}
				buffer.WriteString(string(v))
			}

			if err := os.WriteFile(p, buffer.Bytes(), 0o644); err != nil {
				return fmt.Errorf("error writing to %v: %v", p, err)
			}
		}
	}

	return nil
}
