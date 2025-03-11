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

const SCRIPT_YAML_PROMPT string = `I am trying to create a test case for mockgcp.

A good test case for mockgcp does the basic operations on a GCP resource by using gcloud to create, list, describe and delete the resource.  It can also do a simple update.

For example, if asked to create a mockgcp test for the gcloud commands under <TICK>gcloud pubsub topics<TICK>, we create the file mockpubsub/testdata/topic/crud/script.yaml with the following contents:

<TICK><TICK><TICK>script.yaml
- exec: gcloud pubsub topics create test-${uniqueId}
- exec: gcloud pubsub topics describe test-${uniqueId}
- exec: gcloud pubsub topics delete test-${uniqueId}
<TICK><TICK><TICK>

Or to create mockgcp test for the gcloud commands under <TICK>gcloud storage buckets<TICK> we create the file mockstorage/testdata/bucket/crud/script.yaml with the following contents:

<TICK><TICK><TICK>script.yaml
- exec: gcloud storage buckets create gs://test-${uniqueId}
- exec: gcloud storage buckets describe gs://test-${uniqueId}
- exec: gcloud storage buckets delete gs://test-${uniqueId}
<TICK><TICK><TICK>

Depended resources must be created first and prepended with -pre and cleaned up at the end prepended with -post.  Example pre and post usage:
<TICK><TICK><TICK>script.yaml
- pre: gcloud pubsub topics create test-topic-${uniqueId} --project=${projectId}
- exec: gcloud asset feeds create test-${uniqueId} --pubsub-topic=projects/${projectId}/topics/test-topic-${uniqueId} --project=${projectId}
- exec: gcloud asset feeds describe test-${uniqueId} --project=${projectId}
- exec: gcloud asset feeds update test-${uniqueId} --project=${projectId} --content-type=resource
- exec: gcloud asset feeds delete test-${uniqueId} --project=${projectId}
- post: gcloud pubsub topics delete test-topic-${uniqueId} --project=${projectId}
<TICK><TICK><TICK>

Some hints:

* You should use the CreateFile method to create the script.yaml file in the appropriate directory.  You can use ListFilesInWorkspace to make sure that you are creating a test in a new directory.

* You can run the help command to see the available subcommands, for example you might run <TICK>gcloud pubsub topics --help<TICK>.
If you want to see the flags for any individual commands, you can run the help for them also, for example you might run <TICK>gcloud pubsub topics create --help<TICK>.

* You should run the help command for each command you output, to verify the flags and arguments of the commands.

* If you must specify a project, use the --project flag with this variable ${projectId}, for example <TICK>gcloud pubsub topics create test-${uniqueId} --project=${projectId}<TICK>.

* If you must use project in a resource path, use this variable ${projectId}, for example <TICK>gcloud data-catalog tags create --entry=projects/${projectId}/locations/us-central1/entryGroups/test-entry-group/entries/test-entry-${uniqueId} --tag=test-tag-${uniqueId}<TICK>

* The allowed variables are:
  * ${projectId} - The project ID to use for the test.
  * ${uniqueId} - A unique ID to use for the test.
  * ${BILLING_ACCOUNT_ID} - The billing account ID to use for the test.
  * ${organizationId} - The organization ID if mandatory in the command.
  * ${projectNumber} - The project number to use for the test.

* If the resource requires dependent resources, you should create them in the same script.yaml file.

* Depended resources must be created first and prepended with -pre

* Depended resources must be cleaned up at the end prepended with -post

* Most importantly make sure that all required parameters and flags are included in the commands

Please create a test case for the gcloud commands under <TICK><GCLOUD_COMMAND><TICK>
Please create the test case in the file <TICK>mock<GROUP>/testdata/<RESOURCE>/crud/script.yaml<TICK>

When you have completed, please output the name of the test script you have created, in a JSON format like this:

{ "path_to_created_test": "mock<GROUP>/testdata/<RESOURCE>/crud/script.yaml" }`

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

func createScriptYaml(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		log.Printf("SKIPPING %s, no gcloud command", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")
	logDir := filepath.Join(opts.loggingDir, branch.Name)

	checkoutBranch(ctx, branch, workDir)

	// Check to see if the script file already exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)

	// Check if file exists and handle force flag
	if _, err := os.Stat(scriptFullPath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("ERROR checking file %s: %v", scriptFullPath, err)
			return
		}
	} else if opts.force {
		if err := os.Remove(scriptFullPath); err != nil {
			log.Printf("ERROR deleting existing file %s: %v", scriptFullPath, err)
			return
		}
	} else {
		log.Printf("SKIPPING %s, %s already exists", branch.Name, scriptFullPath)
		return
	}

	// Delete then write the prompt file.
	promptPath := filepath.Join(logDir, "create-script-prompt.txt")
	writeTemplateToFile(branch, promptPath, SCRIPT_YAML_PROMPT)

	// Run the LLM to generate the file.
	cfg := CommandConfig{
		Name:         "Generate script",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", "--prompt=" + promptPath},
		WorkDir:      workDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("SCRIPT FILE GENERATE error: %q\n", err)
		// Currently ignoring error and just basing on if the script.yaml was generated.
		// log.Fatal(err)
	}

	// Check to see if the script file was created
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created", branch.Name, scriptFullPath)
		return
	}

	// Add the new file to the current branch.
	gitAdd(ctx, workDir, scriptFile)

	// Commit the change to the current branch.
	gitCommit(ctx, workDir, fmt.Sprintf("Adding LLM/gcloud generated test script.yaml for %s", branch.Name))
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
			MaxRetries:   3,
			RetryBackoff: 10 * time.Second,
		}
		_, _, err := executeCommand(opts, cfg)
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
	gitAdd(ctx, workDir, scriptFile)
	gitCommit(ctx, workDir, fmt.Sprintf("Manually updated script.yaml for %s", branch.Name))
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

func captureHttpLog(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	// Check to see if the script file exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(workDir, scriptFile)
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing script %s", branch.Name, scriptFullPath)
		return
	}

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)
	if _, err := os.Stat(logFullPath); !errors.Is(err, os.ErrNotExist) && !opts.force {
		log.Printf("SKIPPING %s, %s already exists", branch.Name, logFullPath)
		return
	}

	// Current HTTP Log generation is determenistic not ML generated.

	// Run the test to generate the log.
	cfg := CommandConfig{
		Name: "Generate HTTP log",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests",
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource),
			"-timeout", fmt.Sprintf("%s", opts.timeout),
		},
		WorkDir:    workDir,
		Env:        map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "real"},
		MaxRetries: 1,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("TEST GENERATE error: %q\n", err)
		// Currently ignoring error and just basing on if the _http.log was generated.
		// log.Fatal(err)
	}

	// Check to see if the script file was created
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created", branch.Name, logFullPath)
		return
	}

	// Add the new file to the current branch.
	gitAdd(ctx, workDir, logFullPath)

	// Commit the change to the current branch.
	gitCommit(ctx, workDir, fmt.Sprintf("Adding mockgcptests generated _http.log for %s", branch.Name))
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

func generateMockGo(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	var hasChange = false
	checkoutBranch(ctx, branch, workDir)

	// Check to see if the http log file already exists
	mockfolder := fmt.Sprintf("mock%s", branch.Group)

	logFullPath := filepath.Join(workDir, "mockgcp", mockfolder, "testdata", branch.Resource, "crud", "_http.log")
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing %s", branch.Name, logFullPath)
		return
	}

	// Delete then write the service go prompt file.
	servicePromptPath := filepath.Join(workDir, "service_prompt.txt")
	writeTemplateToFile(branch, servicePromptPath, MOCK_SERVICE_GO_GEN)

	// Delete then write the resource go prompt file.
	resourcePromptPath := filepath.Join(workDir, "resource_prompt.txt")
	writeTemplateToFile(branch, resourcePromptPath, MOCK_RESOURCE_GO_GEN)

	// Run the controller builder to generate the service go file.
	serviceFile := filepath.Join(workDir, "mockgcp", mockfolder, "service.go")
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) || opts.force {
		cfg := CommandConfig{
			Name: "Generate service mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", "./mockgcp",
				"--proto-dir", fmt.Sprintf(".build/third_party/googleapis/"),
				"--input-file", "service_prompt.txt",
			},
			WorkDir:      workDir,
			RetryBackoff: GenerativeCommandRetryBackoff,
		}
		output, _, err := executeCommand(opts, cfg)
		if err != nil {
			log.Printf("MOCK SERVICE GENERATE error: %q\n", err)
			// Currently ignoring error and just basing on if the _http.log was generated.
			// log.Fatal(err)
		} else {
			if err := os.WriteFile(serviceFile, []byte(output), 0755); err != nil {
				log.Printf("WRITE MOCK SERVICE %s error: %q\n", serviceFile, err)
			}
			log.Printf("MOCK SERVICE GENERATE: %q\n", formatCommandOutput(output))
		}

		// Check to see if the service go file was created
		if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
			log.Printf("SKIPPING %s, %s was not created", branch.Name, serviceFile)
			return
		}

		// Add the new files to the current branch.
		hasChange = true
		gitAdd(ctx, workDir, serviceFile)
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
	resourceFile := filepath.Join(workDir, "mockgcp", mockfolder, fmt.Sprintf("%s.go", resourceName))
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) || opts.force {
		cfg := CommandConfig{
			Name: "Generate resource mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", "./mockgcp",
				"--proto-dir", ".build/third_party/googleapis/",
				"--input-file", "resource_prompt.txt",
			},
			WorkDir:      workDir,
			RetryBackoff: GenerativeCommandRetryBackoff,
		}
		output, _, err := executeCommand(opts, cfg)
		if err != nil {
			log.Printf("MOCK RESOURCE GENERATE error: %q\n", err)
			// Currently ignoring error and just basing on if the _http.log was generated.
			// log.Fatal(err)
		} else {
			if err := os.WriteFile(resourceFile, []byte(output), 0755); err != nil {
				log.Printf("WRITE MOCK RESOURCE %s error: %q\n", resourceFile, err)
			}
			log.Printf("MOCK RESOURCE GENERATE: %q\n", formatCommandOutput(output))
		}

		// Check to see if the service go file was created
		if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
			log.Printf("SKIPPING %s, %s was not created", branch.Name, resourceFile)
			return
		}

		// Add the new files to the current branch.
		hasChange = true
		gitAdd(ctx, workDir, resourceFile)
	} else {
		log.Printf("SKIPPING generating resource mock go, %s already exists", resourceFile)
	}

	// Commit the change to the current branch.
	if hasChange {
		gitCommit(ctx, workDir, fmt.Sprintf("Adding mock service and resource for %s", branch.Name))
	} else {
		log.Printf("SKIPPING git commit, no new changes for %s", branch.Name)
	}
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

func addServiceToRoundTrip(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	serviceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), "service.go")
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing %s", branch.Name, serviceFile)
		return
	}

	// Delete then write the add service to roundtrip prompt file.
	roundtripPromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "roundtrip_prompt.txt")
	writeTemplateToFile(branch, roundtripPromptPath, ADD_SERVICE_TO_ROUNDTRIP)

	// Run the LLM to add the service to roundtrip file.
	cfg := CommandConfig{
		Name:         "Add service to roundtrip",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", "--prompt=roundtrip_prompt.txt"},
		WorkDir:      workDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("addServiceToRoundTrip error: %q\n", err)
		// Currently ignoring error and just basing on if the mock_http_roundtrip.go was diff.
		// log.Fatal(err)
	}
	if !gitFileHasChange(workDir, "mock_http_roundtrip.go") {
		return
	}

	// Add the new files to the current branch.
	gitAdd(ctx, workDir, "mock_http_roundtrip.go")

	// Commit the change to the current branch.
	gitCommit(ctx, workDir, fmt.Sprintf("Adding service to mock_http_roundtrip.go for %s", branch.Name))
}

const ADD_PROTO_TO_MAKEFILE string = `Please add the generation for <TICK><PROTO_PACKAGE><TICK> to the <TICK>gen-proto-no-fixup<TICK> target in <TICK>Makefile<TICK>.

Hints:

* Use the ReadFile command to read the contents of the file.

* Use the EditFile command to insert the appropriate third_party directory into the list of paths.

* The gen-proto-no-fixup command contains a long protoc command, split across multiple lines.  There should be a backslash character (\) on all lines but the last.  Make sure there is a space before the backslash.

* The generation path being added  should begin with <TICK>./third_party/googleapis/mockgcp/<RESOURCE><TICK> and should not contain google after mockgcp.

* This is not a correct path: <TICK>./third_party/googleapis/mockgcp/google/cloud/metastore/...<TICK>

* This is the correct path: <TICK>./third_party/googleapis/cloud/mockgcp/metastore/...<TICK>
`

func addProtoToMakefile(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	// TODO: Populate the ProtoPath in branches-all.yaml.
	// Maybe populate it with the actual filepath?
	apisDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")
	protoFile := filepath.Join(apisDir, branch.ProtoPath)
	if _, err := os.Stat(protoFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing %s", branch.Name, protoFile)
		return
	}

	// Delete then write the add service to roundtrip prompt file.
	roundtripPromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "makefile_prompt.txt")
	template := strings.ReplaceAll(ADD_PROTO_TO_MAKEFILE, "<PROTO_PACKAGE>", branch.ProtoPath)
	writeTemplateToFile(branch, roundtripPromptPath, template)

	// Run the LLM to add the service to roundtrip file.
	cfg := CommandConfig{
		Name:         "Add proto to makefile",
		Cmd:          "codebot",
		Args:         []string{"--ui-type=prompt", "--prompt=makefile_prompt.txt"},
		WorkDir:      workDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		// log.Fatal(err)
		log.Printf("updating proto Makefile error: %q\n", err)
	}

	hasChange := false

	makefile := filepath.Join("mockgcp", "Makefile")
	if gitFileHasChange(opts.branchRepoDir, makefile) {
		hasChange = true
		gitAdd(ctx, opts.branchRepoDir, makefile)
	}

	fixup := filepath.Join("mockgcp", "fixup-third-party.sh")
	if gitFileHasChange(opts.branchRepoDir, fixup) {
		hasChange = true
		gitAdd(ctx, opts.branchRepoDir, fixup)
	}

	// Add the new files to the current branch.
	if hasChange {
		// Commit the change to the current branch.
		gitCommit(ctx, opts.branchRepoDir, fmt.Sprintf("Adding proto to Makefile for %s", branch.Name))
	}
}

func runMockgcpTests(opts *RunnerOptions, branch Branch) {
    if opts.defaultRetries > 0 {
        log.Printf("Command does not support retries")
    }
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)

	// Run the test against the generated mocks to determine quality
	cfg := CommandConfig{
		Name: "Run mock tests",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests",
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource),
			"-timeout", fmt.Sprintf("%s", opts.timeout),
		},
		WorkDir:    workDir,
		Env:        map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "mock"},
        MaxRetries: -1, // no retries, even if user defined
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("TEST RUN error: %q\n", err)
		// Currently ignoring error and just basing on if the _http.log was generated.
		// log.Fatal(err)
	}

	// Check to see if the script file was created
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created", branch.Name, logFullPath)
		return
	}
}

func buildProtoFiles(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	checkoutBranch(ctx, branch, workDir)

	// Run make gen-proto command
	cfg := CommandConfig{
		Name:       "Generate proto files",
		Cmd:        "make",
		Args:       []string{"gen-proto"},
		WorkDir:    workDir,
		MaxRetries: 1,
	}

	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("Proto generation error: %v", err)
		return
	}

	// Check if the generated directory exists and has changes
	generatedDir := filepath.Join(workDir, "generated")
	if _, err := os.Stat(generatedDir); err != nil {
		log.Printf("Generated directory not found: %v", err)
		return
	}

	if !gitStatusCheck(workDir, "generated") && !gitFileHasChange(workDir, "generated") {
		log.Printf("SKIPPING %s, no changes to generated directory", branch.Name)
		return
	}

	// Add the generated directory to git
	gitAdd(ctx, workDir, "generated")

	// Commit the changes
	gitCommit(ctx, workDir, fmt.Sprintf("Generated proto files for %s", branch.Name))
}
