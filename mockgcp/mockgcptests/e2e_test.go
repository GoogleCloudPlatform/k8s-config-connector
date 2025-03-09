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
	ProjectID               string
	ProjectNumber           int64
	UniqueID                string
	BillingAccountID        string
	IAMTestOrganizationID   string
	IAMTestBillingAccountID string
	User                    string
}

func TestScripts(t *testing.T) {
	baseDir, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("cannot find base dir for mockgcp: %v", err)
	}
	scriptPaths := findScripts(t, baseDir)

	for _, scriptPath := range scriptPaths {
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
				ProjectID:               project.ProjectID,
				ProjectNumber:           project.ProjectNumber,
				UniqueID:                uniqueID,
				BillingAccountID:        testgcp.TestBillingAccountID.Get(),
				IAMTestBillingAccountID: testgcp.IAMIntegrationTestsBillingAccountID.Get(),
				IAMTestOrganizationID:   testgcp.IAMIntegrationTestsOrganizationID.Get(),
				User:                    GetDefaultAccount(t),
			}
			script := loadScript(t, testDir, placeholders)

			for _, step := range script.SetupSteps {
				if step.Exec != "" {
					cmd := exec.CommandContext(ctx, "bash", "-c", step.Exec)
					var stdout bytes.Buffer
					cmd.Stdout = &stdout
					var stderr bytes.Buffer
					cmd.Stderr = &stderr

					cmd.Env = append(cmd.Env, os.Environ()...)

					if h.gcpAccessToken != "" {
						cmd.Env = append(cmd.Env, fmt.Sprintf("CLOUDSDK_AUTH_ACCESS_TOKEN=%v", h.gcpAccessToken))
					}
					cmd.Env = append(cmd.Env, "CLOUDSDK_CORE_PROJECT="+h.Project.ProjectID)
					cmd.Dir = testDir

					t.Logf("executing setup step command %q", step.Exec)
					if err := cmd.Run(); err != nil {
						t.Logf("stdout: %v", stdout.String())
						t.Logf("stderr: %v", stderr.String())

						t.Errorf("error running command %q: %v", step.Exec, err)
					}
				}
			}

			h.StartProxy(ctx)

			for _, step := range script.Steps {
				stepCmd := ""
				stepType := ""
				needProxy := false
				if step.Pre != "" {
					stepCmd = step.Pre
					stepType = "pre"
				} else if step.Exec != "" {
					stepCmd = step.Exec
					stepType = "exec"
					needProxy = true
				} else if step.Post != "" {
					stepCmd = step.Post
					stepType = "post"
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
					if needProxy {
						gcloudConfig := h.proxy.BuildGcloudConfig(h.ProxyEndpoint, h.MockGCP)
						cmd.Env = append(cmd.Env, gcloudConfig.EnvVars...)
					}
					cmd.Dir = testDir

					t.Logf("executing step type: %s  cmd: %q", stepType, stepCmd)
					if err := cmd.Run(); err != nil {
						t.Logf("stdout: %v", stdout.String())
						t.Logf("stderr: %v", stderr.String())

						t.Errorf("error running step type: %s  cmd: %q: %v", stepType, stepCmd, err)
					}
				}
			}

			{
				httpEvents := h.Events.HTTPEvents

				for _, httpEvent := range httpEvents {
					// gcloud includes a UUID in the user-agent, along with a lot of other client info (e.g. kernel version, python version)
					// Just remove it from the golden output.
					httpEvent.Request.RemoveHeader("user-agent")

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

			for _, step := range script.TeardownSteps {
				if step.Exec != "" {
					cmd := exec.CommandContext(ctx, "bash", "-c", step.Exec)
					var stdout bytes.Buffer
					cmd.Stdout = &stdout
					var stderr bytes.Buffer
					cmd.Stderr = &stderr

					cmd.Env = append(cmd.Env, os.Environ()...)

					if h.gcpAccessToken != "" {
						cmd.Env = append(cmd.Env, fmt.Sprintf("CLOUDSDK_AUTH_ACCESS_TOKEN=%v", h.gcpAccessToken))
					}
					cmd.Env = append(cmd.Env, "CLOUDSDK_CORE_PROJECT="+h.Project.ProjectID)
					cmd.Dir = testDir

					t.Logf("executing teardown step command %q", step.Exec)
					if err := cmd.Run(); err != nil {
						t.Logf("stdout: %v", stdout.String())
						t.Logf("stderr: %v", stderr.String())

						t.Errorf("error running command %q: %v", step.Exec, err)
					}
				}
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

	SetupSteps    []*Step
	TeardownSteps []*Step
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
	s.Steps = loadFile(t, dir, "script.yaml", placeholders, true)
	s.SetupSteps = loadFile(t, dir, "setup.yaml", placeholders, false)
	s.TeardownSteps = loadFile(t, dir, "teardown.yaml", placeholders, false)
	return s
}

func loadFile(t *testing.T, dir, fileName string, placeholders Placeholders, mustRead bool) []*Step {
	var b []byte
	var err error
	if mustRead {
		b = test.MustReadFile(t, filepath.Join(dir, fileName))
	} else {
		b, err = os.ReadFile(filepath.Join(dir, fileName))
		if err != nil {
			return nil
		}
	}
	if len(b) == 0 {
		return nil
	}

	b = ReplaceTestVars(t, b, placeholders)

	var steps []*Step
	if err := yaml.Unmarshal(b, &steps); err != nil {
		t.Fatalf("error unmarshalling steps in %s: %v", fileName, err)
	}
	return steps
}

// ReplaceTestVars replaces all occurrences of placeholder strings e.g. ${uniqueId} in a given byte slice.
func ReplaceTestVars(t *testing.T, b []byte, placeholders Placeholders) []byte {
	s := string(b)
	s = strings.Replace(s, "${uniqueId}", placeholders.UniqueID, -1)
	s = strings.Replace(s, "${projectId}", placeholders.ProjectID, -1)
	s = strings.Replace(s, "${projectNumber}", strconv.FormatInt(placeholders.ProjectNumber, 10), -1)
	s = strings.Replace(s, "${BILLING_ACCOUNT_ID}", placeholders.BillingAccountID, -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IAMIntegrationTestsOrganizationID.Key), placeholders.IAMTestOrganizationID, -1)
	s = strings.Replace(s, fmt.Sprintf("${%s}", testgcp.IAMIntegrationTestsBillingAccountID.Key), placeholders.IAMTestBillingAccountID, -1)
	s = strings.Replace(s, "${user}", placeholders.User, -1)
	return []byte(s)
}
