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
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/tests/e2e"
	"sigs.k8s.io/yaml"
)

type Placeholders struct {
	ProjectID        string
	ProjectNumber    int64
	UniqueID         string
	BillingAccountID string
}

func TestScripts(t *testing.T) {
	baseDir, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("cannot find base dir for mockgcp: %v", err)
	}

	scriptPaths := findScripts(t, baseDir)

	for _, scriptPath := range scriptPaths {

		// skip the crud test for vertex AI model for now due to API migration.
		// The gcloud commands still use the legacy REST API (which is said to be deprecated since Jan 31 2025): https://cloud.google.com/ai-platform/prediction/docs/reference/rest
		// But the mock service is implemented based on the new API: https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.models
		if scriptPath == "mockaiplatform/testdata/model/crud" {
			continue
		}
		t.Run(scriptPath, func(t *testing.T) {
			t.Parallel()

			ctx := context.TODO()
			ctx, closeContext := context.WithCancel(ctx)
			t.Cleanup(closeContext)

			uniqueID := fmt.Sprintf("%x", time.Now().UnixNano())

			h := NewHarness(t)
			h.Init()

			project := h.Project
			testDir := filepath.Join(baseDir, scriptPath)
			placeholders := Placeholders{
				ProjectID:        project.ProjectID,
				ProjectNumber:    project.ProjectNumber,
				UniqueID:         uniqueID,
				BillingAccountID: testgcp.TestBillingAccountID.Get(),
			}
			script := loadScript(t, testDir, placeholders)

			h.StartProxy(ctx)

			var httpEvents []*test.LogEntry

			for _, step := range script.Steps {
				stepCmd := ""
				stepType := ""
				captureEvents := true
				if step.Pre != "" {
					stepCmd = step.Pre
					stepType = "pre"
					captureEvents = false
				} else if step.Exec != "" {
					stepCmd = step.Exec
					stepType = "exec"
				} else if step.Post != "" {
					stepCmd = step.Post
					stepType = "post"
					captureEvents = false
				}
				if stepCmd != "" {
					cmd := exec.CommandContext(ctx, "bash", "-c", stepCmd)
					var stdout bytes.Buffer
					cmd.Stdout = &stdout
					var stderr bytes.Buffer
					cmd.Stderr = &stderr

					cmd.Env = append(cmd.Env, os.Environ()...)

					if h.gcpAccessToken != "" {
						cmd.Env = append(cmd.Env, fmt.Sprintf("CLOUDSDK_AUTH_ACCESS_TOKEN=%v", h.gcpAccessToken))
					}
					cmd.Env = append(cmd.Env, "CLOUDSDK_CORE_PROJECT="+h.Project.ProjectID)
					gcloudConfig := h.proxy.BuildGcloudConfig(h.ProxyEndpoint, h.MockGCP)
					cmd.Env = append(cmd.Env, gcloudConfig.EnvVars...)
					cmd.Dir = testDir

					t.Logf("executing step type: %s  cmd: %q", stepType, stepCmd)
					if err := cmd.Run(); err != nil {
						t.Logf("stdout: %v", stdout.String())
						t.Logf("stderr: %v", stderr.String())

						t.Errorf("error running step type: %s  cmd: %q: %v", stepType, stepCmd, err)
					}

					if captureEvents {
						httpEvents = append(httpEvents, h.Events.HTTPEvents...)
					}
					h.Events.HTTPEvents = nil
				}
			}

			{
				for _, httpEvent := range httpEvents {
					// gcloud includes a UUID in the user-agent, along with a lot of other client info (e.g. kernel version, python version)
					// Just remove it from the golden output.
					httpEvent.Request.RemoveHeader("user-agent")

					httpEvent.Request.RemoveHeader("X-Goog-User-Project")

					// The X-Goog-User-Project header is (always) set by gcloud if a quota project is set,
					// so this header reflects configuration not the actual protocol.
					httpEvent.Request.RemoveHeader("X-Goog-User-Project")

					// Remove the Content-Length header, as it changes with dynamic values
					httpEvent.Request.RemoveHeader("Content-Length")
					httpEvent.Response.RemoveHeader("Content-Length")
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
	Pre  string `json:"pre"`
	Post string `json:"post"`
}

func loadScript(t *testing.T, dir string, placeholders Placeholders) *Script {
	s := &Script{
		Name:      dir,
		SourceDir: dir,
	}
	b := test.MustReadFile(t, filepath.Join(dir, "script.yaml"))

	b = ReplaceTestVars(t, b, placeholders)

	var steps []*Step
	if err := yaml.Unmarshal(b, &steps); err != nil {
		t.Fatalf("error unmarshalling steps: %v", err)
	}

	s.Steps = steps

	return s
}

// ReplaceTestVars replaces all occurrences of placeholder strings e.g. ${uniqueId} in a given byte slice.
func ReplaceTestVars(t *testing.T, b []byte, placeholders Placeholders) []byte {
	s := string(b)
	s = strings.Replace(s, "${uniqueId}", placeholders.UniqueID, -1)
	s = strings.Replace(s, "${projectId}", placeholders.ProjectID, -1)
	s = strings.Replace(s, "${projectNumber}", strconv.FormatInt(placeholders.ProjectNumber, 10), -1)
	s = strings.Replace(s, "${BILLING_ACCOUNT_ID}", placeholders.BillingAccountID, -1)
	return []byte(s)
}
