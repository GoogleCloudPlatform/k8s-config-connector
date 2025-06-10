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

	"gopkg.in/yaml.v3"
)

// BranchScript represents script data for a branch
type BranchScript struct {
	Name        string              `yaml:"name"`           // Branch name
	LocalBranch string              `yaml:"localbranch"`    // Local branch name
	Content     []map[string]string `yaml:"script-content"` // Raw script content with exec commands
	Warnings    []string            `yaml:"warnings"`       // List of warnings about the script content
}

// ScriptData represents the collection of scripts
type ScriptData struct {
	Branches []BranchScript `yaml:"branches"`
}

func enableAPIs(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.ApisEnabled == nil || len(branch.ApisEnabled) == 0 {
		log.Printf("[Enable APIs] SKIPPING %s, no APIs to enable", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")
	for _, api := range branch.ApisEnabled {
		log.Printf("[Enable APIs] Enabling API %s", api)
		cfg := CommandConfig{
			Name:         fmt.Sprintf("Enable API %s", api),
			Cmd:          "gcloud",
			Args:         []string{"services", "enable", api},
			WorkDir:      workDir,
			MaxAttempts:  3,
			RetryBackoff: 10 * time.Second,
		}
		_, err := executeCommand(opts, cfg)
		if err != nil {
			log.Printf("[Enable APIs] Failed to enable API %s: %v", api, err)
		}
	}
}

func readScriptYaml(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		log.Printf("SKIPPING %s, no gcloud command", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")
	checkoutBranch(ctx, branch, workDir)

	// Check to see if the script file already exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s doesn't exists", branch.Name, scriptFullPath)
		return
	}

	// Read and parse the script file
	data, err := os.ReadFile(scriptFullPath)
	if err != nil {
		log.Printf("Error reading file %s: %v", scriptFullPath, err)
		return
	}

	// Parse YAML content
	var scriptContent []map[string]string
	err = yaml.Unmarshal(data, &scriptContent)
	if err != nil {
		log.Printf("Error parsing YAML from %s: %v", scriptFullPath, err)
		return
	}

	// Check commands and collect warnings
	var warnings []string
	hasCreate := false
	hasDelete := false
	hasList := false

	for _, line := range scriptContent {
		if cmd, ok := line["exec"]; ok && strings.Contains(cmd, "gcloud") {
			if strings.Contains(cmd, " create ") {
				hasCreate = true
			}
			if strings.Contains(cmd, " delete ") {
				hasDelete = true
			}
			if strings.Contains(cmd, " list") {
				hasList = true
			}
		}
	}

	if !hasCreate {
		warning := "No Create found"
		warnings = append(warnings, warning)
		log.Printf("WARNING: %s in %s", warning, scriptFile)
	}
	if !hasDelete {
		warning := "No Delete found"
		warnings = append(warnings, warning)
		log.Printf("WARNING: %s in %s", warning, scriptFile)
	}
	if hasList {
		warning := "List found"
		warnings = append(warnings, warning)
		log.Printf("WARNING: %s in %s", warning, scriptFile)
	}

	// Read existing all_scripts.yaml or create new ScriptData
	allScriptsPath := filepath.Join(opts.loggingDir, "all_scripts.yaml")
	var scriptData ScriptData

	if data, err := os.ReadFile(allScriptsPath); err == nil {
		err = yaml.Unmarshal(data, &scriptData)
		if err != nil {
			log.Printf("Error parsing existing all_scripts.yaml: %v", err)
			return
		}
	}

	// Find and update existing branch or add new one
	newBranchScript := BranchScript{
		Name:        branch.Name,
		LocalBranch: branch.Local,
		Content:     scriptContent,
		Warnings:    warnings,
	}

	found := false
	for i, bs := range scriptData.Branches {
		if bs.Name == branch.Name {
			scriptData.Branches[i] = newBranchScript
			found = true
			break
		}
	}
	if !found {
		scriptData.Branches = append(scriptData.Branches, newBranchScript)
	}

	// Write updated data back to all_scripts.yaml
	yamlData, err := yaml.Marshal(scriptData)
	if err != nil {
		log.Printf("Error marshaling script data: %v", err)
		return
	}

	err = os.MkdirAll(opts.loggingDir, 0755)
	if err != nil {
		log.Printf("Error creating logs directory: %v", err)
		return
	}

	err = os.WriteFile(allScriptsPath, yamlData, 0644)
	if err != nil {
		log.Printf("Error writing all_scripts.yaml: %v", err)
		return
	}
}

func writeScriptYaml(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		log.Printf("SKIPPING %s, no gcloud command", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")
	checkoutBranch(ctx, branch, workDir)

	// Read all_scripts.yaml
	allScriptsPath := filepath.Join(opts.loggingDir, "all_scripts.yaml")
	data, err := os.ReadFile(allScriptsPath)
	if err != nil {
		log.Printf("Error reading all_scripts.yaml: %v", err)
		return
	}

	var scriptData ScriptData
	err = yaml.Unmarshal(data, &scriptData)
	if err != nil {
		log.Printf("Error parsing all_scripts.yaml: %v", err)
		return
	}

	// Find the branch's script data
	var branchScript *BranchScript
	for i := range scriptData.Branches {
		if scriptData.Branches[i].Name == branch.Name {
			branchScript = &scriptData.Branches[i]
			break
		}
	}

	if branchScript == nil {
		log.Printf("No script found for branch %s", branch.Name)
		return
	}

	// Create script file path
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)

	// Create directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(scriptFullPath), 0755)
	if err != nil {
		log.Printf("Error creating script directory: %v", err)
		return
	}

	// Write the script content directly
	yamlData, err := yaml.Marshal(branchScript.Content)
	if err != nil {
		log.Printf("Error marshaling script content: %v", err)
		return
	}

	err = os.WriteFile(scriptFullPath, yamlData, 0644)
	if err != nil {
		log.Printf("Error writing script file: %v", err)
		return
	}

	// Add and commit the changes
	if !gitFileHasChange(workDir, scriptFile) {
		log.Printf("SKIPPING %s, no changes to %s", branch.Name, scriptFile)
		return
	}
	err = gitAdd(ctx, workDir, scriptFile)
	if err != nil {
		log.Printf("ERROR adding %s: %v", scriptFile, err)
	}
	err = gitCommit(ctx, workDir, fmt.Sprintf("Manually updated script.yaml for %s", branch.Name))
	if err != nil {
		log.Printf("ERROR committing %s: %v", scriptFile, err)
	}
}

const CAPTURE_HTTP_LOG string = `I need to capture the logs from GCP for running a mockgcp test that I just created.  I then need to create a git commit.

For example, if I just created a script mockpubsub/testdata/topic/crud/script.yaml, then I should run

<TICK>WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockpubsub/testdata/topic/crud<TICK>

I would then run <TICK>git add mockpubsub/testdata/topic/crud/_http.log<TICK>, then <TICK>git commit<TICK> that with a commit message like "mockgcp: Capture golden output for mockpubsub/testdata/topic/crud"

For example, if I just created a script mockstorage/testdata/topic/bucket/script.yaml, then I should run

<TICK>WRITE_GOLDEN_OUTPUT=1 E2E_GCP_TARGET=real go test ./mockgcptests -run TestScripts/mockstorage/testdata/bucket/crud<TICK>

I would then run <TICK>git add mockstorage/testdata/bucket/crud/_http.log<TICK>, then <TICK>git commit<TICK> that with a commit message like "mockgcp: Capture golden output for mockstorage/testdata/bucket/crud"

Please capture the logs for the script I just created, called <TICK>mockworkflows/testdata/workflow/crud/script.yaml<TICK>.

When you are done, please output a JSON result like this:

{ "status": "success" }


If you have problems, please output a JSON result like this:

{ "status": "failure", "reason": "Fill in any information on why you could not complete the task" }`

func captureHttpLog(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string

	// Check to see if the script file exists
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", fmt.Sprintf("mock%s", branch.Group), "testdata", branch.Resource, "crud", "script.yaml")
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing script %s", scriptFullPath)
	}

	// Check to see if the http log file already exists
	logFileRelativePath := filepath.Join("mockgcp", fmt.Sprintf("mock%s", branch.Group), "testdata", branch.Resource, "crud", "_http.log")
	logFilePath := filepath.Join(opts.branchRepoDir, logFileRelativePath)
	if _, err := os.Stat(logFilePath); !errors.Is(err, os.ErrNotExist) && !opts.force {
		return affectedPaths, nil, fmt.Errorf("http log %s already exists", logFilePath)
	}

	// Current HTTP Log generation is determenistic not ML generated.

	// Run the test to generate the log.
	cfg := CommandConfig{
		Name: "Generate HTTP log",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests",
			"-v", // verbose
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource),
			"-timeout", fmt.Sprintf("%s", opts.timeout),
		},
		WorkDir:     filepath.Join(opts.branchRepoDir, "mockgcp"),
		Env:         map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "real"},
		MaxAttempts: 1,
	}
	results, err := executeCommand(opts, cfg)
	affectedPaths = append(affectedPaths, logFileRelativePath)
	return affectedPaths, &results, err
}

func readHttpLog(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s does not exists", branch.Name, logFullPath)
		return
	}
	data, err := os.ReadFile(logFullPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}

const MOCK_SERVICE_GO_GEN string = `// +tool:mockgcp-service
// http.host: <HTTP_HOST>
// proto.service: <PROTO_SERVICE>
`

const MOCK_RESOURCE_GO_GEN string = `// +tool:mockgcp-support
// proto.service: <PROTO_SERVICE>
// proto.message: <PROTO_MESSAGE>
`

func generateMockGo(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string
	workDir := opts.branchRepoDir

	// Check to see if the http log file already exists
	mockfolder := fmt.Sprintf("mock%s", branch.Group)

	// Create the directory if it doesn't exist
	mockfolderPath := filepath.Join(opts.branchRepoDir, "mockgcp", mockfolder)
	if err := os.MkdirAll(mockfolderPath, 0755); err != nil {
		return affectedPaths, nil, fmt.Errorf("failed to create directory %s: %w", mockfolderPath, err)
	}

	// Run the controller builder to generate the service go file.
	serviceFileRelativePath := filepath.Join("mockgcp", mockfolder, "service.go")
	serviceFile := filepath.Join(opts.branchRepoDir, serviceFileRelativePath)
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) || opts.force {
		// Delete then write the service go prompt file.
		promptPath := filepath.Join(opts.loggingDir, branch.Name, "generate-mocks-service-prompt.txt")
		writeTemplateToFile(branch, promptPath, MOCK_SERVICE_GO_GEN)
		cfg := CommandConfig{
			Name: "Generate service mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", "./mockgcp",
				"--proto-dir", fmt.Sprintf(".build/third_party/googleapis/"),
				"--input-file", promptPath,
			},
			WorkDir:      opts.branchRepoDir,
			RetryBackoff: GenerativeCommandRetryBackoff,
		}
		output, err := executeCommand(opts, cfg)
		if err != nil {
			return affectedPaths, nil, fmt.Errorf("MOCK SERVICE GENERATE error: %w\n", err)
		} else {
			if err := os.WriteFile(serviceFile, []byte(output.Stdout), 0755); err != nil {
				return affectedPaths, nil, fmt.Errorf("WRITE MOCK SERVICE %s error: %w\n", serviceFile, err)
			}
			affectedPaths = append(affectedPaths, serviceFileRelativePath)
			log.Printf("MOCK SERVICE GENERATE: %q\n", formatCommandOutput(output.Stdout))
		}
	} else {
		log.Printf("SKIPPING generating service mock go, %s already exists", serviceFile)
	}

	// Run the controller builder to generate the resource go file.
	resourceName := strings.ToLower(branch.Resource)
	if resourceName == "service" {
		// Special case for service, it's actually a resource.
		log.Printf("WARNING: Special case for resource with names 'service', setting file name to 'resourceservice.go'")
		resourceName = "resourceservice"
	}
	resourceFileRelativePath := filepath.Join("mockgcp", mockfolder, fmt.Sprintf("%s.go", resourceName))
	resourceFile := filepath.Join(opts.branchRepoDir, resourceFileRelativePath)
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) || opts.force {
		// Delete then write the resource go prompt file.
		promptPath := filepath.Join(opts.loggingDir, branch.Name, "generate-mocks-resource-prompt.txt")
		writeTemplateToFile(branch, promptPath, MOCK_RESOURCE_GO_GEN)
		cfg := CommandConfig{
			Name: "Generate resource mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", "./mockgcp",
				"--proto-dir", ".build/third_party/googleapis/",
				"--input-file", promptPath,
			},
			WorkDir:      workDir,
			RetryBackoff: GenerativeCommandRetryBackoff,
		}
		output, err := executeCommand(opts, cfg)
		if err != nil {
			return affectedPaths, nil, fmt.Errorf("MOCK RESOURCE GENERATE error: %w\n", err)
		} else {
			if err := os.WriteFile(resourceFile, []byte(output.Stdout), 0755); err != nil {
				return affectedPaths, nil, fmt.Errorf("WRITE MOCK RESOURCE %s error: %w\n", resourceFile, err)
			}
			affectedPaths = append(affectedPaths, resourceFileRelativePath)
			log.Printf("MOCK RESOURCE GENERATE: %q\n", formatCommandOutput(output.Stdout))
		}
	} else {
		log.Printf("SKIPPING generating resource mock go, %s already exists", resourceFile)
	}

	return affectedPaths, nil, nil
}

func readMockGo(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	serviceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), "service.go")
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING reading service mock go, %s does not exist", serviceFile)
	}
	data, err := os.ReadFile(serviceFile)
	if err != nil {
		fmt.Println("Error reading service file:", err)
		return
	}
	fmt.Println(string(data))

	resourceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), fmt.Sprintf("%s.go", strings.ToLower(branch.Resource)))
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING reading resource mock go, %s does not exist", resourceFile)
	}
	data, err = os.ReadFile(serviceFile)
	if err != nil {
		fmt.Println("Error reading resource file:", err)
		return
	}
	fmt.Println(string(data))
}

const ADD_SERVICE_TO_ROUNDTRIP string = `Please add the services in <TICK>mock<SERVICE><TICK> to <TICK>mock_http_roundtrip.go<TICK>

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mock<SERVICE> into the list of services.
* Please keep the list of services in alphabetical order.
* Don't forget to import the package!`

func addServiceToRoundTrip(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string

	mockfolder := fmt.Sprintf("mock%s", branch.Group)
	serviceFilePath := filepath.Join(opts.branchRepoDir, "mockgcp", mockfolder, "service.go")
	if _, err := os.Stat(serviceFilePath); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing %s", serviceFilePath)
	}

	// Delete then write the add service to roundtrip prompt file.
	promptPath := filepath.Join(opts.loggingDir, branch.Name, "add-service-to-roundtrip-prompt.txt")
	writeTemplateToFile(branch, promptPath, ADD_SERVICE_TO_ROUNDTRIP)

	// Run the LLM to add the service to roundtrip file.
	cfg := CommandConfig{
		Name:         "Add service to roundtrip",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", fmt.Sprintf("--prompt=%s", promptPath)},
		WorkDir:      filepath.Join(opts.branchRepoDir, "mockgcp"),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	results, err := executeCommand(opts, cfg)

	affectedPaths = append(affectedPaths, filepath.Join("mockgcp", "mock_http_roundtrip.go"))

	return affectedPaths, &results, err
}

const ADD_PROTO_TO_MAKEFILE string = `Please add the generation for <TICK><PROTO_PACKAGE><TICK> to the <TICK>gen-proto-no-fixup<TICK> target in <TICK>Makefile<TICK>.

Hints:

* Use the ReadFile command to read the contents of the file.

* Use the EditFile command to insert the appropriate third_party directory into the list of paths.

* The gen-proto-no-fixup command contains a long protoc command, split across multiple lines.  There should be a backslash character (\) on all lines but the last.  Make sure there is a space before the backslash.

* The generation path being added  should begin with <TICK>./third_party/googleapis/mockgcp/<RESOURCE><TICK> and should not contain google after mockgcp.

* This is not a correct path: <TICK>./third_party/googleapis/mockgcp/google/cloud/metastore/...<TICK>

* This is the correct path: <TICK>./third_party/googleapis/mockgcp/cloud/metastore/...<TICK>
`

func addProtoToMakefile(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string

	// TODO: why should we check this?
	mockfolder := fmt.Sprintf("mock%s", branch.Group)
	serviceFilePath := filepath.Join(opts.branchRepoDir, "mockgcp", mockfolder, "service.go")
	if _, err := os.Stat(serviceFilePath); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing %s", serviceFilePath)
	}

	// TODO: Populate the ProtoPath in branches-all.yaml.
	// Maybe populate it with the actual filepath?
	protoFile := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis", branch.ProtoPath)
	if _, err := os.Stat(protoFile); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing %s", protoFile)
	}

	// Delete then write the add service to roundtrip prompt file.
	// TODO: review later. PROTO_PACKAGE = branch.ProtoPath or branch.ProtoPackage?
	template := strings.ReplaceAll(ADD_PROTO_TO_MAKEFILE, "<PROTO_PACKAGE>", branch.ProtoPath)
	promptPath := filepath.Join(opts.loggingDir, branch.Name, "add-proto-to-makefile-prompt.txt")
	writeTemplateToFile(branch, promptPath, template)

	// Run the LLM to add the service to roundtrip file.
	cfg := CommandConfig{
		Name:         "Add proto to makefile",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", fmt.Sprintf("--prompt=%s", promptPath)},
		WorkDir:      filepath.Join(opts.branchRepoDir, "mockgcp"),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	results, err := executeCommand(opts, cfg)
	affectedPaths = append(affectedPaths, filepath.Join("mockgcp", "Makefile"))
	return affectedPaths, &results, err
}

func runMockgcpTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {

	var affectedPaths []string

	mockfolder := fmt.Sprintf("mock%s", branch.Group)

	// Running the tests will affect the http log file. We need to revert it. So return the path.
	logFileFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", mockfolder, "testdata", branch.Resource, "crud", "_http.log")
	affectedPaths = append(affectedPaths, logFileFullPath)
	if _, err := os.Stat(logFileFullPath); errors.Is(err, os.ErrNotExist) {
		return affectedPaths, nil, fmt.Errorf("missing http log %s", logFileFullPath)
	}

	// Run the test against the generated mocks to determine quality
	cfg := CommandConfig{
		Name: "Run mock tests",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests",
			"-v", // verbose
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource),
			"-timeout", fmt.Sprintf("%s", opts.timeout),
		},
		WorkDir:     filepath.Join(opts.branchRepoDir, "mockgcp"),
		Env:         map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "mock"},
		MaxAttempts: 1,
	}
	results, err := executeCommand(opts, cfg)
	return affectedPaths, &results, err
}

const FIX_MOCKGCP_FAILURES string = `Please fix the mock resource file: ${MOCK_RESOURCE_FILE}  based on the failures seen in the mock test run.

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
E2E_GCP_TARGET=mock go test ./mockgcptests -run TestScripts/mock${GROUP}/testdata/${RESOURCE}/crud

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

func fixMockgcpFailures(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
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

	logFileRelativePath := filepath.Join("mockgcp", fmt.Sprintf("mock%s", branch.Group), "testdata", branch.Resource, "crud", "_http.log")
	logFilePath := filepath.Join(opts.branchRepoDir, logFileRelativePath)
	// Create prompt with file contents
	prompt := strings.ReplaceAll(FIX_MOCKGCP_FAILURES, "${MOCK_RESOURCE_FILE}", resourceFileRelativePath)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_EXITCODE}", fmt.Sprintf("%d", execResults.ExitCode))
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDERR}", execResults.Stderr)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDOUT}", execResults.Stdout)
	prompt = strings.ReplaceAll(prompt, "${GROUP}", branch.Group)
	prompt = strings.ReplaceAll(prompt, "${RESOURCE}", resourceName)
	prompt = strings.ReplaceAll(prompt, "${PROTO_FILE_CONTENTS}", protoContents.String())
	prompt = strings.ReplaceAll(prompt, "${ORIGINAL_HTTP_LOG_FILE}", logFilePath)
	// Run the LLM to generate the file.
	cfg := CommandConfig{
		Name:         "Fix mockgcp failures",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(prompt),
		WorkDir:      filepath.Join(opts.branchRepoDir, "mockgcp"),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	results, err := executeCommand(opts, cfg)
	return affectedPaths, &results, err
}

func buildProtoFiles(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var affectedPaths []string

	// Run make gen-proto command
	cfg := CommandConfig{
		Name:        "Generate proto files",
		Cmd:         "make",
		Args:        []string{"gen-proto"},
		WorkDir:     filepath.Join(opts.branchRepoDir, "mockgcp"),
		MaxAttempts: 1,
	}

	results, err := executeCommand(opts, cfg)
	affectedPaths = append(affectedPaths, filepath.Join("mockgcp", "generated"))
	if err != nil {
		return affectedPaths, nil, fmt.Errorf("Proto generation error: %w", err)
	}

	return affectedPaths, &results, err
}
