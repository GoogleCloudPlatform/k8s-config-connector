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

package runner

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// generateControllerClient creates a controller client for the branch
// This implements the logic from 01-create-controller-client.sh
func generateControllerClient(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller client for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.ProtoSvc == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoSvc field", branch.Name)
	}

	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Create the controller client
	outputPath := filepath.Join("pkg", "controller", "direct", branch.Group, "client.go")
	protoDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:controller-client\n// proto.service: %s\n", branch.ProtoSvc)

	// Ensure the directory exists
	clientDir := filepath.Join(opts.branchRepoDir, "pkg", "controller", "direct", branch.Group)
	if err := os.MkdirAll(clientDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create directory %s: %w", clientDir, err)
	}

	// Run controllerbuilder command
	cfg := CommandConfig{
		Name:    "Controller Builder",
		Cmd:     "controllerbuilder",
		Args:    []string{"prompt", "--src-dir", opts.branchRepoDir, "--proto-dir", protoDir},
		WorkDir: opts.branchRepoDir,
		Stdin:   strings.NewReader(prompt),
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run controllerbuilder: %w", err)
	}

	// Write the output to the client.go file
	clientFile := filepath.Join(clientDir, "client.go")

	// Write the client.go file
	if err := os.WriteFile(clientFile, []byte(output.Stdout), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write client.go: %w", err)
	}

	log.Printf("Successfully generated controller client for %s at %s", branch.Name, clientFile)

	return []string{outputPath}, &output, nil
}

// generateController creates a controller for the branch
// This implements the logic from 02-create-controller.sh
func generateController(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.ProtoSvc == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoSvc field", branch.Name)
	}

	if branch.ProtoMsg == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoMsg field", branch.Name)
	}

	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Create the controller
	controllerFileName := strings.ToLower(branch.Proto) + "_controller.go"
	outputPath := filepath.Join("pkg", "controller", "direct", branch.Group, controllerFileName)

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:controller\n// proto.service: %s\n// proto.message: %s\n// crd.type: %s\n// crd.version: %s\n",
		branch.ProtoSvc, branch.ProtoMsg, branch.Kind, "v1alpha1") // Using v1alpha1 as default CRD version

	// Ensure the directory exists
	controllerDir := filepath.Join(opts.branchRepoDir, "pkg", "controller", "direct", branch.Group)
	protoDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")
	if err := os.MkdirAll(controllerDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create directory %s: %w", controllerDir, err)
	}

	// Run controllerbuilder command
	cfg := CommandConfig{
		Name:    "Controller Builder",
		Cmd:     "controllerbuilder",
		Args:    []string{"prompt", "--src-dir", opts.branchRepoDir, "--proto-dir", protoDir},
		WorkDir: opts.branchRepoDir,
		Stdin:   strings.NewReader(prompt),
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run controllerbuilder: %w", err)
	}

	// Write the output to the controller file
	controllerFile := filepath.Join(controllerDir, controllerFileName)

	// Write the controller file
	if err := os.WriteFile(controllerFile, []byte(output.Stdout), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write controller file: %w", err)
	}

	log.Printf("Successfully generated controller for %s at %s", branch.Name, controllerFile)

	return []string{outputPath}, &output, nil
}

// generateControllerIdentity creates an identity file for the branch
// This implements the first part of 03-create-identity.sh
func generateControllerIdentity(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller identity for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.ProtoSvc == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoSvc field", branch.Name)
	}

	if branch.ProtoMsg == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoMsg field", branch.Name)
	}

	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Default CRD version
	crdVersion := "v1alpha1"

	// Create the identity file
	identityFileName := strings.ToLower(branch.Kind) + "_identity.go"
	outputPath := filepath.Join("apis", branch.Group, strings.ToLower(crdVersion), identityFileName)
	protoDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:krm-identity\n// proto.service: %s\n// proto.message: %s\n// crd.type: %s\n// crd.version: %s\n",
		branch.ProtoSvc, branch.ProtoMsg, branch.Kind, crdVersion)

	// Ensure the directory exists
	identityDir := filepath.Join(opts.branchRepoDir, "apis", branch.Group, strings.ToLower(crdVersion))
	if err := os.MkdirAll(identityDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create directory %s: %w", identityDir, err)
	}

	// Run controllerbuilder command
	cfg := CommandConfig{
		Name:    "Controller Builder Identity",
		Cmd:     "controllerbuilder",
		Args:    []string{"prompt", "--src-dir", opts.branchRepoDir, "--proto-dir", protoDir},
		WorkDir: opts.branchRepoDir,
		Stdin:   strings.NewReader(prompt),
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run controllerbuilder for identity: %w", err)
	}

	// Write the output to the identity file
	identityFile := filepath.Join(identityDir, identityFileName)

	// Write the identity file
	if err := os.WriteFile(identityFile, []byte(output.Stdout), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write identity file: %w", err)
	}

	log.Printf("Successfully generated controller identity for %s at %s", branch.Name, identityFile)

	return []string{outputPath}, &output, nil
}

// generateControllerReference creates a reference file for the branch
// This implements the second part of 03-create-identity.sh
func generateControllerReference(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller reference for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.ProtoSvc == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoSvc field", branch.Name)
	}

	if branch.ProtoMsg == "" {
		return nil, nil, fmt.Errorf("branch %s is missing ProtoMsg field", branch.Name)
	}

	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Default CRD version
	crdVersion := "v1alpha1"

	// Create the reference file
	referenceFileName := strings.ToLower(branch.Kind) + "_reference.go"
	outputPath := filepath.Join("apis", branch.Group, strings.ToLower(crdVersion), referenceFileName)

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:krm-reference\n// proto.service: %s\n// proto.message: %s\n// crd.type: %s\n// crd.version: %s\n",
		branch.ProtoSvc, branch.ProtoMsg, branch.Kind, crdVersion)
	protoDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")

	// Ensure the directory exists
	referenceDir := filepath.Join(opts.branchRepoDir, "apis", branch.Group, strings.ToLower(crdVersion))
	if err := os.MkdirAll(referenceDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create directory %s: %w", referenceDir, err)
	}

	// Run controllerbuilder command
	cfg := CommandConfig{
		Name:    "Controller Builder Reference",
		Cmd:     "controllerbuilder",
		Args:    []string{"prompt", "--src-dir", opts.branchRepoDir, "--proto-dir", protoDir},
		WorkDir: opts.branchRepoDir,
		Stdin:   strings.NewReader(prompt),
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run controllerbuilder for reference: %w", err)
	}

	// Write the output to the reference file
	referenceFile := filepath.Join(referenceDir, referenceFileName)

	// Write the reference file
	if err := os.WriteFile(referenceFile, []byte(output.Stdout), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write reference file: %w", err)
	}

	log.Printf("Successfully generated controller reference for %s at %s", branch.Name, referenceFile)

	return []string{outputPath}, &output, nil
}

// createControllerTest creates test files for the branch
// This implements the first part of 04-create-test.sh
func createControllerTest(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Creating controller test for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Default CRD version and group
	crdVersion := "v1alpha1"
	crdGroup := fmt.Sprintf("%s.cnrm.cloud.google.com", branch.Group)

	// Create the test directory
	testDir := filepath.Join(
		"pkg", "test", "resourcefixture", "testdata", "basic",
		branch.Group, strings.ToLower(crdVersion),
		strings.ToLower(branch.Kind),
		fmt.Sprintf("%s-minimal", strings.ToLower(branch.Kind)),
	)

	fullTestDir := filepath.Join(opts.branchRepoDir, testDir)
	if err := os.MkdirAll(fullTestDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create test directory %s: %w", fullTestDir, err)
	}

	currentYear := time.Now().Year()
	yamlCopyright := fmt.Sprintf(`# Copyright %d Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
`, currentYear)

	// Create create.yaml
	createYaml := fmt.Sprintf(`%s
apiVersion: %s/%s
kind: %s
metadata:
  name: %s-minimal-${uniqueId}
spec:
  projectRef:
    external: ${projectId}
  location: us-central1
  description: "Initial description"
`, yamlCopyright, crdGroup, crdVersion, branch.Kind, strings.ToLower(branch.Kind))

	createYamlPath := filepath.Join(fullTestDir, "create.yaml")
	if err := os.WriteFile(createYamlPath, []byte(createYaml), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write create.yaml: %w", err)
	}

	// Create update.yaml
	updateYaml := fmt.Sprintf(`%s
apiVersion: %s/%s
kind: %s
metadata:
  name: %s-minimal-${uniqueId}
spec:
  projectRef:
    external: ${projectId}
  location: us-central1
  description: "Updated description"
`, yamlCopyright, crdGroup, crdVersion, branch.Kind, strings.ToLower(branch.Kind))

	updateYamlPath := filepath.Join(fullTestDir, "update.yaml")
	if err := os.WriteFile(updateYamlPath, []byte(updateYaml), 0644); err != nil {
		return nil, nil, fmt.Errorf("failed to write update.yaml: %w", err)
	}

	log.Printf("Successfully created test files for %s in %s", branch.Name, testDir)

	return []string{
		filepath.Join(testDir, "create.yaml"),
		filepath.Join(testDir, "update.yaml"),
	}, nil, nil
}

// updateTestHarness updates the test harness to support the new resource
// This implements the second part of 04-create-test.sh
func updateTestHarness(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Updating test harness for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Default CRD group
	crdGroup := fmt.Sprintf("%s.cnrm.cloud.google.com", branch.Group)

	// Run codebot to update the test harness
	prompt := fmt.Sprintf(`Please add a case statement for Group "%s" and Kind "%s" to the switch statement in MaybeSkip,
in the file config/tests/samples/create/harness.go

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to case statement into the list of cases.
* Try to insert it in sorted order first by group, and then by kind.
* If the case already exists, do not make any changes.
`, crdGroup, branch.Kind)

	cfg := CommandConfig{
		Name:         "Codebot Update Harness",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		WorkDir:      opts.branchRepoDir,
		Stdin:        strings.NewReader(prompt),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run codebot to update test harness: %w", err)
	}

	log.Printf("Successfully updated test harness for %s", branch.Name)

	return []string{"config/tests/samples/create/harness.go"}, &output, nil
}

// captureGoldenTestOutput captures golden test output for the branch
// This implements 05-capture-golden-test-output.sh
func captureGoldenTestOutput(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Capturing golden test output for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Run the compare-mock script
	fixtureArg := fmt.Sprintf("fixtures/%s-minimal", strings.ToLower(branch.Kind))

	cfg := CommandConfig{
		Name:        "Compare Mock",
		Cmd:         "hack/compare-mock",
		Args:        []string{fixtureArg},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}

	// We don't care about the error here, as the script might return non-zero
	// The script is run with "|| true" in the original shell script
	output, _ := executeCommand(opts, cfg)

	log.Printf("Compare-mock output: %s", output.Stdout)
	if output.Stderr != "" {
		log.Printf("Compare-mock stderr: %s", output.Stderr)
	}

	// The golden output files will be in the fixtures directory
	// We'll return this path so that the changes can be committed
	fixturesPath := filepath.Join("fixtures", fmt.Sprintf("%s-minimal", strings.ToLower(branch.Kind)))

	log.Printf("Successfully captured golden test output for %s", branch.Name)

	return []string{fixturesPath}, &output, nil
}
