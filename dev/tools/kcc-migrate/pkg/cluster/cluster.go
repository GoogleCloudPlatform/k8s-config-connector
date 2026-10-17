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

package cluster

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// KubectlClient wraps kubectl calls.
type KubectlClient struct {
	Kubeconfig string
	Context    string
}

func NewKubectlClient(kubeconfig, context string) *KubectlClient {
	return &KubectlClient{
		Kubeconfig: kubeconfig,
		Context:    context,
	}
}

func (k *KubectlClient) baseArgs() []string {
	var args []string
	if k.Kubeconfig != "" {
		args = append(args, "--kubeconfig", k.Kubeconfig)
	}
	if k.Context != "" {
		args = append(args, "--context", k.Context)
	}
	return args
}

// RunCommand runs a kubectl command and returns output.
func (k *KubectlClient) RunCommand(extraArgs ...string) (string, error) {
	args := append(k.baseArgs(), extraArgs...)
	cmd := exec.Command("kubectl", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("kubectl %s failed: %v, stderr: %s", strings.Join(extraArgs, " "), err, stderr.String())
	}
	return stdout.String(), nil
}

// GetKCCCRDs returns all CRDs in *.cnrm.cloud.google.com.
func (k *KubectlClient) GetKCCCRDs() ([]string, error) {
	out, err := k.RunCommand("get", "crds", "-o", "custom-columns=NAME:.metadata.name", "--no-headers")
	if err != nil {
		return nil, err
	}

	var kccCRDs []string
	for _, line := range strings.Split(out, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, ".cnrm.cloud.google.com") {
			kccCRDs = append(kccCRDs, trimmed)
		}
	}
	return kccCRDs, nil
}

// ExportResources exports all resources of specified CRDs in the given namespaces to a directory concurrently.
func (k *KubectlClient) ExportResources(crds []string, namespaces []string, outputDir string) (int, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return 0, fmt.Errorf("creating output dir: %w", err)
	}

	type task struct {
		crd string
		ns  string
	}

	taskList := make([]task, 0, len(crds)*len(namespaces))
	for _, crd := range crds {
		for _, ns := range namespaces {
			taskList = append(taskList, task{crd: crd, ns: ns})
		}
	}

	tasks := make(chan task, len(taskList))
	for _, t := range taskList {
		tasks <- t
	}
	close(tasks)

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalExported := 0
	numWorkers := 20
	if len(taskList) < numWorkers {
		numWorkers = len(taskList)
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				kind := strings.Split(t.crd, ".")[0]
				args := []string{"get", t.crd, "-n", t.ns, "-o", "yaml"}
				out, err := k.RunCommand(args...)
				if err != nil {
					continue
				}

				if strings.Contains(out, "items: []") || strings.TrimSpace(out) == "" {
					continue
				}

				fileName := fmt.Sprintf("%s_%s.yaml", t.ns, kind)
				filePath := filepath.Join(outputDir, fileName)
				if err := os.WriteFile(filePath, []byte(out), 0644); err == nil {
					mu.Lock()
					totalExported++
					mu.Unlock()
					fmt.Printf("    Exported %s/%s -> %s\n", t.ns, kind, fileName)
				}
			}
		}()
	}

	wg.Wait()
	return totalExported, nil
}

// ApplyDirectory applies manifests in a directory to the target cluster.
func (k *KubectlClient) ApplyDirectory(dirPath string, dryRun bool) (string, error) {
	args := []string{"apply", "-f", dirPath}
	if dryRun {
		args = append(args, "--dry-run=server")
	}
	return k.RunCommand(args...)
}

// ApplyFile applies a single manifest file to the target cluster.
func (k *KubectlClient) ApplyFile(filePath string, dryRun bool) (string, error) {
	args := []string{"apply", "-f", filePath}
	if dryRun {
		args = append(args, "--dry-run=server")
	}
	return k.RunCommand(args...)
}
