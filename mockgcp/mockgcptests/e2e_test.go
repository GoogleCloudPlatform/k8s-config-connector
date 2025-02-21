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

package mockgcptests

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/tests/e2e"
	"sigs.k8s.io/yaml"
)

func TestScripts(t *testing.T) {
	baseDir, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("cannot find base dir for mockgcp: %v", err)
	}

	scriptPaths := findScripts(t, baseDir)

	for _, scriptPath := range scriptPaths {
		t.Run(scriptPath, func(t *testing.T) {
			ctx := context.TODO()
			ctx, closeContext := context.WithCancel(ctx)
			t.Cleanup(closeContext)

			uniqueID := fmt.Sprintf("%x", time.Now().UnixNano())

			h := NewHarness(t)
			h.Init()

			project := h.Project
			testDir := filepath.Join(baseDir, scriptPath)

			script := loadScript(t, testDir, uniqueID, project)

			h.StartProxy()

			for _, step := range script.Steps {
				if step.Exec != "" {
					args := strings.Fields(step.Exec)

					cmd := exec.CommandContext(ctx, args[0], args[1:]...)
					var stdout bytes.Buffer
					cmd.Stdout = &stdout
					var stderr bytes.Buffer
					cmd.Stderr = &stderr

					cmd.Env = append(cmd.Env, os.Environ()...)

					if h.gcpAccessToken != "" {
						cmd.Env = append(cmd.Env, fmt.Sprintf("CLOUDSDK_AUTH_ACCESS_TOKEN=%v", h.gcpAccessToken))
					}
					gcloudConfig := h.proxy.BuildGcloudConfig(h.ProxyEndpoint, h.MockGCP)
					cmd.Env = append(cmd.Env, "CLOUDSDK_CORE_PROJECT="+h.Project.ProjectID)
					cmd.Env = append(cmd.Env, gcloudConfig.EnvVars...)
					cmd.Dir = testDir

					t.Logf("executing step command %q", step.Exec)
					if err := cmd.Run(); err != nil {
						t.Logf("stdout: %v", stdout.String())
						t.Logf("stderr: %v", stderr.String())

						t.Errorf("error running command %q: %v", step.Exec, err)
					}
				}
			}

			{
				httpEvents := h.Events.HTTPEvents

				for _, httpEvent := range httpEvents {
					// gcloud includes a UUID in the user-agent, along with a lot of other client info (e.g. kernel version, python version)
					// Just remove it from the golden output.
					httpEvent.Request.RemoveHeader("user-agent")
				}

				folderID := ""
				organizationID := ""

				e2e.NormalizeHTTPLog(t, httpEvents, h.RegisteredServices(), testgcp.GCPProject{ProjectID: h.Project.ProjectID, ProjectNumber: h.Project.ProjectNumber}, uniqueID, folderID, organizationID)

				x := e2e.NewNormalizer(uniqueID, testgcp.GCPProject{ProjectID: h.Project.ProjectID, ProjectNumber: h.Project.ProjectNumber})

				x.Preprocess(httpEvents)

				expectedPath := filepath.Join(script.SourceDir, "_http.log")
				got := x.Render(httpEvents)
				h.CompareGoldenFile(expectedPath, got)
			}

		})
	}
}

func findScripts(t *testing.T, rootDir string) []string {
	var relPaths []string
	if err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "script.yaml" {
			relPath, err := filepath.Rel(rootDir, filepath.Dir(path))
			if err != nil {
				return fmt.Errorf("getting relative path during directory walk: %w", err)
			}
			relPaths = append(relPaths, relPath)
		}
		return nil
	}); err != nil {
		t.Fatalf("error walking directory %q: %v", rootDir, err)
	}
	return relPaths
}

type Script struct {
	Name      string
	SourceDir string
	Steps     []*Step
}

type Step struct {
	Exec string `json:"exec"`
}

func loadScript(t *testing.T, dir string, uniqueID string, project GCPProject) *Script {
	s := &Script{
		Name:      dir,
		SourceDir: dir,
	}
	b := test.MustReadFile(t, filepath.Join(dir, "script.yaml"))

	b = ReplaceTestVars(t, b, uniqueID, project)

	var steps []*Step
	if err := yaml.Unmarshal(b, &steps); err != nil {
		t.Fatalf("error unmarshalling steps: %v", err)
	}

	s.Steps = steps

	return s
}

// ReplaceTestVars replaces all occurrences of placeholder strings e.g. ${uniqueId} in a given byte slice.
func ReplaceTestVars(t *testing.T, b []byte, uniqueID string, project GCPProject) []byte {
	s := string(b)
	s = strings.Replace(s, "${uniqueId}", uniqueID, -1)
	s = strings.Replace(s, "${projectId}", project.ProjectID, -1)
	return []byte(s)
}
