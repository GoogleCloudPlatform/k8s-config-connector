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
	"errors"
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

	// Use codebot to generate the initial create.yaml content
	createPrompt := fmt.Sprintf(`Generate a %s/create.yaml file for testing a Kubernetes controller.

First, read the CRD file at config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<pluralized-kind>.%s.cnrm.cloud.google.com.yaml to understand the schema.
Replace <pluralized-kind> with the pluralized version of the kind: %s in the filename.

The file should follow these requirements:
- Add an Apache 2.0 license header with Copyright %d Google LLC
- Use apiVersion: %s/%s 
- Use kind: %s
- Include metadata.name: %s-minimal-${uniqueId}
- If the CRD has a projectRef field, use projectRef.external: ${projectId}
- If the CRD has a location field, use location: us-central1
- Keep the YAML simple but valid for initial creation
- Follow the schema defined in the CRD file
- Only include required fields from the CRD

Use ReadFile to read the CRD file.
Use CreateFile to write the YAML content to the %s/create.yaml file.
Respond only with the YAML content, no explanations.`,
		testDir, branch.Group, branch.Kind, currentYear, crdGroup, crdVersion, branch.Kind, strings.ToLower(branch.Kind), testDir)

	cfg := CommandConfig{
		Name:         "Generate Create YAML",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(createPrompt),
		WorkDir:      opts.branchRepoDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	_, err := executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate create.yaml: %w", err)
	}

	// Use codebot to generate the update.yaml content by modifying create.yaml
	updatePrompt := fmt.Sprintf(`Generate an %s/update.yaml file for testing a Kubernetes controller by modifying the create.yaml file.

First, read the %s/create.yaml file that was just generated.
Then modify 1-2 fields to create a valid update scenario.

The file should follow these requirements:
- Keep the same license header, apiVersion, kind and metadata.name
- Modify 1-2 fields in the spec section to test updates
- Ensure the changes are valid according to the CRD schema
- Keep other fields unchanged from create.yaml

Use ReadFile to read the %s/create.yaml file.
Use CreateFile to write the YAML content to the %s/update.yaml file.
Respond only with the YAML content, no explanations.`,
		testDir, testDir, testDir, testDir)

	cfg = CommandConfig{
		Name:         "Generate Update YAML",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(updatePrompt),
		WorkDir:      opts.branchRepoDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	_, err = executeCommand(opts, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate update.yaml: %w", err)
	}

	return []string{
		filepath.Join(testDir, "create.yaml"),
		filepath.Join(testDir, "update.yaml"),
	}, nil, nil
}

func moveTestToSubDir(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Moving test to sub directory for %s", branch.Name)

	// Check if we have the required fields
	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Default CRD version and group
	crdVersion := "v1alpha1"

	if branch.Controller == "terraform-v1beta1" {
		crdVersion = "v1beta1"
		log.Printf("This is a TF-based v1beta1 resource")
	}

	// Move the existing test data to subdirectory if exists.
	relativeParentDir := filepath.Join(
		"pkg", "test", "resourcefixture", "testdata", "basic",
		branch.Group, strings.ToLower(crdVersion),
		strings.ToLower(branch.Kind),
	)
	parentDir := filepath.Join(opts.branchRepoDir, relativeParentDir)
	subDir := filepath.Join(parentDir, strings.ToLower(branch.Kind))
	_, err := os.Stat(subDir)
	if err == nil {
		log.Printf("Sub directory %s for %s already exists", subDir, branch.Name)
		return nil, nil, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, nil, fmt.Errorf("error checking whether sub directory %s exists: %w", subDir, err)
	}

	createFilePath := filepath.Join(parentDir, "create.yaml")
	_, err = os.Stat(createFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("No test data under directory %s for %s", parentDir, branch.Name)
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("error checking whether test data %s exists: %w", createFilePath, err)
	}

	if err := os.MkdirAll(subDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create sub directory %s: %w", subDir, err)
	}

	files, err := os.ReadDir(parentDir)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading directory %s: %w", parentDir, err)
	}
	// Iterate through the files and move them to the subdirectory
	changedPaths := make([]string, 0)
	for _, file := range files {
		if file.Name() == "_vcr_cassettes" || !file.IsDir() { // Check if it's a file (not a directory)
			oldPath := filepath.Join(parentDir, file.Name())
			newPath := filepath.Join(subDir, file.Name())
			err := os.Rename(oldPath, newPath) // Moves the file
			if err != nil {
				return nil, nil, fmt.Errorf("error moving file %s: %w", oldPath, err)
			}
			changedPaths = append(changedPaths, newPath)
		}
	}
	if len(changedPaths) > 0 {
		changedPaths = []string{parentDir}
	} else {
		changedPaths = nil
	}

	return changedPaths, nil, nil
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

// buildController attempts to build the controller code for a branch
func buildController(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Building controller code for branch %s", branch.Name)

	// Collect all Go files in the controller directory
	controllerDir := filepath.Join(opts.branchRepoDir, "pkg", "controller", "direct", branch.Group)
	goFiles := []string{}

	err := filepath.Walk(controllerDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			// Convert to relative path
			relPath, err := filepath.Rel(opts.branchRepoDir, path)
			if err != nil {
				return err
			}
			goFiles = append(goFiles, relPath)
		}
		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("error walking controller directory: %w", err)
	}

	if len(goFiles) == 0 {
		return nil, nil, fmt.Errorf("no Go files found in %s", controllerDir)
	}
	// Define controller file paths that should exist
	resourceLower := strings.ToLower(branch.Resource)
	controllerPath := filepath.Join("pkg", "controller", "direct", branch.Group, fmt.Sprintf("%s_controller.go", resourceLower))
	clientPath := filepath.Join("pkg", "controller", "direct", branch.Group, "client.go")

	// Check if controller files exist
	fullControllerPath := filepath.Join(opts.branchRepoDir, controllerPath)
	fullClientPath := filepath.Join(opts.branchRepoDir, clientPath)

	if _, err := os.Stat(fullControllerPath); err != nil {
		log.Printf("Controller file not found at %s, skipping build", fullControllerPath)
		return []string{}, nil, fmt.Errorf("controller file not found: %s", fullControllerPath)
	}

	if _, err := os.Stat(fullClientPath); err != nil {
		log.Printf("Client file not found at %s, skipping build", fullClientPath)
		return []string{}, nil, fmt.Errorf("client file not found: %s", fullClientPath)
	}

	args := []string{
		"build",
		"-o", "/dev/null", // Don't create output binary
	}
	for _, file := range goFiles {
		args = append(args, file)
	}

	// Run go build command
	cfg := CommandConfig{
		Name:        "Build Controller",
		Cmd:         "go",
		Args:        args,
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}

	results, err := executeCommand(opts, cfg)
	return []string{}, &results, err
}

const FIX_CONTROLLER_BUILD string = `I need your help fixing Go build errors in controller code for Config Connector.

The controller build is failing with errors that need to be fixed. Here are the details of the errors:

${BUILD_OUTPUT}

The controller code is located in the following files:
- ${CONTROLLER_PATH}
- ${CLIENT_PATH}

Please examine the error messages and suggest fixes for the code. Focus on:

1. Missing imports
2. Type mismatches
3. Undefined variables or functions
4. Interface implementation issues
5. Syntax errors

Here are the current contents of the controller files:

CONTROLLER FILE:
${CONTROLLER_FILE_CONTENTS}

CLIENT FILE:
${CLIENT_FILE_CONTENTS}

Please provide specific changes to fix the build errors. After making the changes, I will attempt to build the code again.
`

// fixControllerBuild attempts to fix controller build errors
func fixControllerBuild(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Fixing controller build errors for branch %s", branch.Name)

	// Define controller file paths that should exist
	resourceLower := strings.ToLower(branch.Resource)
	controllerPath := filepath.Join("pkg", "controller", "direct", branch.Group, fmt.Sprintf("%s_controller.go", resourceLower))
	clientPath := filepath.Join("pkg", "controller", "direct", branch.Group, "client.go")
	// If build passed, no need to fix anything
	if execResults.ExitCode == 0 {
		log.Printf("Controller build passed for branch %s, no fixes needed", branch.Name)
		return []string{}, nil, nil
	}

	// Read file contents
	controllerContent, err := os.ReadFile(filepath.Join(opts.branchRepoDir, controllerPath))
	if err != nil {
		log.Printf("Error reading controller file: %v", err)
		controllerContent = []byte("File not found")
	}

	clientContent, err := os.ReadFile(filepath.Join(opts.branchRepoDir, clientPath))
	if err != nil {
		log.Printf("Error reading client file: %v", err)
		clientContent = []byte("File not found")
	}

	// Create build output string combining stdout and stderr
	buildOutput := execResults.Stdout
	if execResults.Stderr != "" {
		if buildOutput != "" {
			buildOutput += "\n\n"
		}
		buildOutput += execResults.Stderr
	}

	// Create prompt with file contents and build output
	prompt := strings.ReplaceAll(FIX_CONTROLLER_BUILD, "${BUILD_OUTPUT}", buildOutput)
	prompt = strings.ReplaceAll(prompt, "${CONTROLLER_PATH}", controllerPath)
	prompt = strings.ReplaceAll(prompt, "${CLIENT_PATH}", clientPath)
	prompt = strings.ReplaceAll(prompt, "${CONTROLLER_FILE_CONTENTS}", string(controllerContent))
	prompt = strings.ReplaceAll(prompt, "${CLIENT_FILE_CONTENTS}", string(clientContent))

	// Run codebot to fix the issues
	cfg := CommandConfig{
		Name:         "Fix Controller Build Errors",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(prompt),
		WorkDir:      opts.branchRepoDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	results, err := executeCommand(opts, cfg)
	return []string{controllerPath, clientPath}, &results, err
}

// runGoldenMockTests ruVns the golden mock tests for the branch
// It executes the hack/compare-mock command to validate the controller against fixtures
func runGoldenMockTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Running golden mock tests for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Create the fixture name based on the kind
	fixtureArg := fmt.Sprintf("fixtures/%s-minimal", strings.ToLower(branch.Kind))

	// Run the compare-mock command
	cfg := CommandConfig{
		Name:        "Compare Mock Tests",
		Cmd:         "hack/compare-mock",
		Args:        []string{fixtureArg},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
		Env: map[string]string{
			"E2E_GCP_TARGET":  "mock",
			"E2E_KUBE_TARGET": "envtest",
		},
	}

	output, _ := executeCommand(opts, cfg)

	// The golden output files will be in the fixtures directory
	// We'll return this path so that the changes can be reverted
	fixturesPath := filepath.Join("pkg", "test", "resourcefixture")

	return []string{fixturesPath}, &output, nil
}

// fixGoldenTests fixes issues in the golden tests for the branch
// It uses codebot to fix test errors in YAML files and controller code
func fixGoldenTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Fixing golden tests for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	if branch.Group == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Group field", branch.Name)
	}

	// Fix YAML files in fixtures directory
	artifactzLogFile := filepath.Join("artifactz", "realgcp", "http-logs", "unknown", "requests.log")
	// Read the http log file for detailed transaction analysis
	reqLogContent, err := os.ReadFile(filepath.Join(opts.branchRepoDir, artifactzLogFile))
	if err != nil {
		log.Printf("Warning: Could not read http log file %s: %v", artifactzLogFile, err)
	}
	fixtureDir := filepath.Join("pkg", "test", "resourcefixture", "testdata", "basic", branch.Group, "v1alpha1", strings.ToLower(branch.Kind), strings.ToLower(branch.Kind)+"-minimal")
	createYaml := filepath.Join(fixtureDir, "create.yaml")
	updateYaml := filepath.Join(fixtureDir, "update.yaml")
	normalizeFile := filepath.Join("tests", "e2e", "normalize.go")
	// Fix controller file if needed
	controllerFileName := strings.ToLower(branch.Proto) + "_controller.go"
	controllerPath := filepath.Join("pkg", "controller", "direct", branch.Group, controllerFileName)
	//controllerFullPath := filepath.Join(opts.branchRepoDir, controllerPath)

	// Check if execResults contains any useful error information
	var testErrors string
	if execResults != nil {
		if execResults.Stderr != "" {
			testErrors = execResults.Stderr
		} else if execResults.Stdout != "" {
			testErrors = execResults.Stdout
		}
	}

	// If there are still issues, use codebot to analyze and provide more detailed fixes
	log.Printf("Performing detailed error analysis for branch %s", branch.Name)

	// Use codebot to analyze the test errors and provide more specific fixes
	prompt := fmt.Sprintf(`Analyze the following test errors for a Kubernetes controller:

%s

The http transactions captured in the log file are:

%s

--------------------------------

These errors are for the %s kind in group %s.

Please identify specific issues that need to be fixed:
1. Are there any KRM object creation or update problems?
2. Are there issues with the controller implementation?
3. Are there any schema validation errors?
4. Are there any missing fields or incorrect field types?
5. Test diff due to time differences can be fixed by normalizing the timestamps in normalize.go

We expect the fixes to be in these files:
1. %s
2. %s
3. %s
4. %s

Provide come up with recommendations along with the file names for fixing these issues.

Based on the recommendations please fix the issues in the files.

Use ReadFile, EditFile to make the changes.
`,
		testErrors, reqLogContent, branch.Kind, branch.Group, createYaml, controllerPath, updateYaml, normalizeFile)

	cfg := CommandConfig{
		Name:         "Fix Golden Tests",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		WorkDir:      opts.branchRepoDir,
		Stdin:        strings.NewReader(prompt),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	output, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("Failed to fix golden tests: %v", err)
	}

	return []string{createYaml, updateYaml, controllerPath}, &output, nil
}

// runGoldenRealGCPTests runs the golden real GCP tests for the branch
// This implements 05-capture-golden-test-output.sh
func runGoldenRealGCPTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Running golden real GCP tests for branch %s", branch.Name)

	// Check if we have the required fields
	if branch.Kind == "" {
		return nil, nil, fmt.Errorf("branch %s is missing Kind field", branch.Name)
	}

	// Run the compare-mock script
	fixtureArg := fmt.Sprintf("fixtures/%s-minimal", strings.ToLower(branch.Kind))

	env := map[string]string{
		"E2E_GCP_TARGET":  "real",
		"E2E_KUBE_TARGET": "envtest",
		"ARTIFACTS":       "artifactz/real",
	}
	if opts.timeout > 0 {
		env["E2E_TEST_TIMEOUT"] = opts.timeout.String()
	}
	cfg := CommandConfig{
		Name:        "Record GCP Output",
		Cmd:         "hack/record-gcp",
		Args:        []string{fixtureArg},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 2, // First attempt is to generate the golden output and the second attempt is to verify the golden output
		Env:         env,
	}

	// We don't care about the error here, as the script might return non-zero
	// The script is run with "|| true" in the original shell script
	output, _ := executeCommand(opts, cfg)

	log.Printf("record-gcp output: %s", output.Stdout)
	if output.Stderr != "" {
		log.Printf("record-gcp stderr: %s", output.Stderr)
	}

	// The golden output files will be in the fixtures directory
	// We'll return this path so that the changes can be committed
	fixturesPath := filepath.Join("pkg", "test", "resourcefixture")

	log.Printf("Captured golden real GCP output for %s", branch.Name)

	return []string{fixturesPath}, &output, nil
}

const FIX_MOCKGCP_FOR_GOLDEN_FAILURES string = `Please fix the mock resource file: ${MOCK_RESOURCE_FILE}  based on the failures seen in the mock test run.

The mock resource file implements the http endpoints for the resource.
The test output contains the diff in the expected and the actual http responses.
Use that to determine what changes need to be made to the mock resource file.
Make sure that all changes are applied.

Start with the errors from the top of the output and work your way down.
If you see a change that needs to be made, make the change and run the mock test again.
Fixing from top to bottom is important and will help you make progress.

Once the changes are applied, use the RunShellCommand command to run the mock test to make sure the changes are valid.
Use the output of the command to determine if the code changes are valid.

RunShellCommand:
hack/record-gcp fixtures/${KIND_LOWER}-minimal

Hints:

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to edit the file.
* Use the WriteFile command to write the file.	
* Use the RunShellCommand command to run the mock test after making the changes.
* The proto files are in the PROTO FILE CONTENTS section.

The results of the mock test run are:

Original HTTP Log File: ${ORIGINAL_HTTP_LOG_FILE}

EXITCODE: ${TEST_OUTPUT_EXITCODE}

STDERR:
${TEST_OUTPUT_STDERR}

STDOUT:
${TEST_OUTPUT_STDOUT}

PROTO FILE CONTENTS:
Imported Proto go files. The list of files and its contents:

${PROTO_FILE_CONTENTS}
`

func fixMockGcpForGoldenTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string

	mockfolder := fmt.Sprintf("mock%s", branch.Group)
	resourceName := strings.ToLower(branch.Resource)
	if resourceName == "service" {
		// Special case for service, it's actually a resource.
		log.Printf("WARNING: Special case for resource with names 'service', setting file name to 'resourceservice.go'")
		resourceName = "resourceservice"
	}
	resourceFileRelativePath := filepath.Join(mockfolder, fmt.Sprintf("%s.go", resourceName))
	resourceFile := filepath.Join(opts.branchRepoDir, "mockgcp", resourceFileRelativePath)
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing resource file %s", resourceFile)
	}
	affectedPaths = append(affectedPaths, resourceFile)

	// Extract the proto go package path (e.g. cloud/eventarc/v1) from the full proto path
	// ./mockgcp/generated/mockgcp + google/cloud/eventarc/v1/eventarc.proto =>  ./mockgcp/generated/mockgcp/cloud/eventarc/v1
	protoPath := filepath.Join(strings.Split(filepath.Dir(branch.ProtoPath), "/")[1:]...)
	protoDirRelative := filepath.Join("mockgcp", "generated", "mockgcp", protoPath)
	protoDirAbsolute := filepath.Join(opts.branchRepoDir, protoDirRelative)

	protoFiles, err := os.ReadDir(protoDirAbsolute)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read proto directory: %s %w", protoDirAbsolute, err)
	}
	var protoContents strings.Builder
	for _, file := range protoFiles {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			filePathRelative := filepath.Join(protoDirRelative, file.Name())
			content, err := os.ReadFile(filepath.Join(opts.branchRepoDir, filePathRelative))
			if err != nil {
				return nil, nil, fmt.Errorf("failed to read proto file %s: %w", filePathRelative, err)
			}
			protoContents.WriteString("File: " + filePathRelative + "\nContent:\n")
			protoContents.Write(content)
			protoContents.WriteString("\n-----------------------------\n")
		}
	}

	// pkg/test/resourcefixture/testdata/basic/asset/v1alpha1/assetfeed/assetfeed-minimal/_http.log
	fixtureDir := filepath.Join("pkg", "test", "resourcefixture", "testdata", "basic", branch.Group, "v1alpha1", strings.ToLower(branch.Kind), strings.ToLower(branch.Kind)+"-minimal")
	logFileRelativePath := filepath.Join(fixtureDir, "_http.log")
	logFilePath := filepath.Join(opts.branchRepoDir, logFileRelativePath)
	// Create prompt with file contents
	prompt := strings.ReplaceAll(FIX_MOCKGCP_FOR_GOLDEN_FAILURES, "${MOCK_RESOURCE_FILE}", resourceFileRelativePath)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_EXITCODE}", fmt.Sprintf("%d", execResults.ExitCode))
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDERR}", execResults.Stderr)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDOUT}", execResults.Stdout)
	prompt = strings.ReplaceAll(prompt, "${GROUP}", branch.Group)
	prompt = strings.ReplaceAll(prompt, "${RESOURCE}", resourceName)
	prompt = strings.ReplaceAll(prompt, "${PROTO_FILE_CONTENTS}", protoContents.String())
	prompt = strings.ReplaceAll(prompt, "${ORIGINAL_HTTP_LOG_FILE}", logFilePath)
	prompt = strings.ReplaceAll(prompt, "${KIND_LOWER}", strings.ToLower(branch.Kind))
	// Run the LLM to generate the file.
	cfg := CommandConfig{
		Name:         "Fix mockgcp for golden failures",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(prompt),
		WorkDir:      filepath.Join(opts.branchRepoDir, "mockgcp"),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	results, err := executeCommand(opts, cfg)
	return affectedPaths, &results, err
}
