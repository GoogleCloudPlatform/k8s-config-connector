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
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

Some hints:

* You should use the CreateFile method to create the script.yaml file in the appropriate directory.  You can use ListFilesInWorkspace to make sure that you are creating a test in a new directory.

* You can run the help command to see the available subcommands, for example you might run <TICK>gcloud pubsub topics --help<TICK>.
If you want to see the flags for any individual commands, you can run the help for them also, for example you might run <TICK>gcloud pubsub topics create --help<TICK>.

* You should run the help command for each command you output, to verify the flags and arguments of the commands.

Please create a test case for the gcloud commands under <TICK><GCLOUD_COMMAND><TICK>
Please create the test case in the file <TICK>mock<GROUP>/testdata/<RESOURCE>/crud/script.yaml<TICK>

When you have completed, please output the name of the test script you have created, in a JSON format like this:

{ "path_to_created_test": "mock<GROUP>/testdata/<RESOURCE>/crud/script.yaml" }`

func createScriptYaml(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		log.Printf("SKIPPING %s, no gcloud command", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Check to see if the script file already exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)
	if _, err := os.Stat(scriptFullPath); !errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s already exists", branch.Name, scriptFullPath)
		return
	}

	// Delete then write the prompt file.
	promptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "prompt.txt")
	writeTemplateToFile(branch, promptPath, SCRIPT_YAML_PROMPT)

	// Run the LLM to generate the file.
	cfg := CommandConfig{
		Name:    "Generate script",
		Cmd:     "codebot",
		Args:    []string{"--ui-type=prompt", "--prompt=prompt.txt"},
		WorkDir: workDir,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Check to see if the script file was created
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created", branch.Name, scriptFullPath)
		return
	}

	// Add the new file to the current branch.
	gitAdd(workDir, &out, scriptFile)

	// Commit the change to the current branch.
	gitCommit(workDir, &out, fmt.Sprintf("Adding LLM/gcloud generated test script.yaml for %s", branch.Name))
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
			Name:    fmt.Sprintf("Enable API %s", api),
			Cmd:     "gcloud",
			Args:    []string{"services", "enable", api},
			WorkDir: workDir,
		}
		_, _, err := executeCommand(opts, cfg)
		if err != nil {
			log.Printf("[Enable APIs] Failed to enable API %s: %v", api, err)
		}
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

func captureHttpLog(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

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
	if _, err := os.Stat(logFullPath); !errors.Is(err, os.ErrNotExist) {
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
		},
		WorkDir: workDir,
		Env:     map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "real"},
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
	gitAdd(workDir, &out, logFullPath)

	// Commit the change to the current branch.
	gitCommit(workDir, &out, fmt.Sprintf("Adding mockgcptests generated _http.log for %s", branch.Name))
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
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	var hasChange = false
	checkoutBranch(branch, workDir, &out)

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing %s", branch.Name, logFullPath)
		return
	}

	// Delete then write the service go prompt file.
	servicePromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "service_prompt.txt")
	writeTemplateToFile(branch, servicePromptPath, MOCK_SERVICE_GO_GEN)

	// Delete then write the resource go prompt file.
	resourcePromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "resource_prompt.txt")
	writeTemplateToFile(branch, resourcePromptPath, MOCK_RESOURCE_GO_GEN)

	// Run the controller builder to generate the service go file.
	serviceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), "service.go")
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
		cfg := CommandConfig{
			Name: "Generate service mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", opts.branchRepoDir,
				"--proto-dir", fmt.Sprintf("%s/.build/third_party/googleapis/", opts.branchRepoDir),
				"--input-file", "service_prompt.txt",
			},
			WorkDir: workDir,
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
		gitAdd(workDir, &out, serviceFile)
	} else {
		log.Printf("SKIPPING generating service mock go, %s already exists", serviceFile)
	}

	// Run the controller builder to generate the resource go file.
	resourceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), fmt.Sprintf("%s.go", strings.ToLower(branch.Resource)))
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
		cfg := CommandConfig{
			Name: "Generate resource mock",
			Cmd:  "controllerbuilder",
			Args: []string{
				"prompt",
				"--src-dir", opts.branchRepoDir,
				"--proto-dir", fmt.Sprintf("%s/.build/third_party/googleapis/", opts.branchRepoDir),
				"--input-file", "resource_prompt.txt",
			},
			WorkDir: workDir,
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
		gitAdd(workDir, &out, resourceFile)
	} else {
		log.Printf("SKIPPING generating resource mock go, %s already exists", resourceFile)
	}

	// Commit the change to the current branch.
	if hasChange {
		gitCommit(workDir, &out, fmt.Sprintf("Adding mock service and resource for %s", branch.Name))
	} else {
		log.Printf("SKIPPING git commit, no new changes for %s", branch.Name)
	}
}

const ADD_SERVICE_TO_ROUNDTRIP string = `Please add the services in <TICK>mock<SERVICE><TICK> to <TICK>mock_http_roundtrip.go<TICK>

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mock<SERVICE> into the list of services.
* Please keep the list of services in alphabetical order.
* Don't forget to import the package!`

func addServiceToRoundTrip(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

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
		Name:    "Add service to roundtrip",
		Cmd:     "codebot",
		Args:    []string{"--ui-type=prompt", "--prompt=roundtrip_prompt.txt"},
		WorkDir: workDir,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Add the new files to the current branch.
	gitAdd(workDir, &out, "mock_http_roundtrip.go")

	// Commit the change to the current branch.
	gitCommit(workDir, &out, fmt.Sprintf("Adding service to mock_http_roundtrip.go for %s", branch.Name))
}

const ADD_PROTO_TO_MAKEFILE string = `Please add the generation for <TICK><PROTO_PACKAGE><TICK> to the <TICK>generate-grpc-for-google-protos<TICK> target in <TICK>Makefile<TICK>.

Hints:

* Use the ReadFile command to read the contents of the file.

* Use the EditFile command to insert the appropriate third_party directory into the list of paths.

* The generate-grpc-for-google-protos command contains a long protoc command, split across multiple lines.  There should be a backslash character (\) on all lines but the last.  Make sure there is a space before the backslash.`

func addProtoToMakfile(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

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
		Name:    "Add proto to makefile",
		Cmd:     "codebot",
		Args:    []string{"--ui-type=prompt", "--prompt=makefile_prompt.txt"},
		WorkDir: workDir,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Add the new files to the current branch.
	gitAdd(workDir, &out, "Makefile")

	// Commit the change to the current branch.
	gitCommit(workDir, &out, fmt.Sprintf("Adding proto to Makefile for %s", branch.Name))
}

func runMockgcpTests(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)

	// Run the test against the generated mocks to determine quality
	cfg := CommandConfig{
		Name: "Run mock tests",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests", "-v",
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource),
		},
		WorkDir: workDir,
		Env:     map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "mock"},
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
