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
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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

func createScriptYamlBash(opts *RunnerOptions, branch Branch) {
	if branch.Command == "" {
		log.Printf("SKIPPING %s, no gcloud command\r\n", branch.Name)
		return
	}

	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDirBash(opts, "mockgcp", stdin, stdout)

	// Change to the checkout branch
	log.Printf("COMMAND: git checkout %s and echo done\r\n", branch.Local)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git checkout %s && echo done\n", branch.Local))); err != nil {
		log.Fatal(err)
	}
	done := false
	outBuffer := make([]byte, 1000)
	var msg string
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("BRANCH CHECKOUT %s\r\n", msg)

	// Check to see if the script file already exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)
	if _, err := os.Stat(scriptFullPath); !errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s already exists\r\n", branch.Name, scriptFullPath)
		return
	}

	tmp := strings.ReplaceAll(SCRIPT_YAML_PROMPT, "<TICK>", "`")
	tmp = strings.ReplaceAll(tmp, "<GCLOUD_COMMAND>", branch.Command)
	tmp = strings.ReplaceAll(tmp, "<GROUP>", branch.Group)
	prompt := strings.ReplaceAll(tmp, "<RESOURCE>", strings.ToLower(branch.Resource))
	log.Printf("CODEBOT PROMPT %s\r\n", prompt)

	// Delete then write the prompt file.
	promptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "prompt.txt")
	if _, err := os.Stat(promptPath); !errors.Is(err, os.ErrNotExist) {
		log.Println("COMMAND: cleaning up old prompt.txt")
		err = os.Remove(promptPath)
		if err != nil {
			log.Printf("Attempt to clean up prompt.txt failed with %v", err)
		}
	}
	log.Println("COMMAND: writing new prompt.txt")
	err = os.WriteFile(promptPath, []byte(prompt), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Run the LLM to generate the file.
	log.Println("COMMAND: codebot --ui-type=prompt --prompt=prompt.txt and echo done")
	if _, err = stdin.Write([]byte("codebot --ui-type=prompt --prompt=prompt.txt && echo done\n")); err != nil {
		log.Fatal(err)
	}
	done = false
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("CODEBOT GENERATE %s\r\n", msg)

	// Check to see if the script file was created
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created\r\n", branch.Name, scriptFullPath)
		return
	}

	// Add the new file to the current branch.
	log.Printf("COMMAND: git add %s and echo done\r\n", scriptFile)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git add %s && echo done\n", scriptFile))); err != nil {
		log.Fatal(err)
	}
	done = false
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("BRANCH ADD %s\r\n", msg)

	// Commit the change to the current branch.
	log.Printf("COMMAND: git commit -m \"Adding LLM/gcloud generated test script.yaml for %s\" and echo done\r\n", branch.Name)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git commit -m \"Adding LLM/gcloud generated test script.yaml for %s\" && echo done\n", branch.Name))); err != nil {
		log.Fatal(err)
	}
	done = false
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("BRANCH COMMIT %s\r\n", msg)
}

func createScriptYaml(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		//stdout.
		log.Printf("SKIPPING %s, no gcloud command\r\n", branch.Name)
		return
	}

	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, out)

	// Check to see if the script file already exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(opts.branchRepoDir, "mockgcp", scriptFile)
	if _, err := os.Stat(scriptFullPath); !errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s already exists\r\n", branch.Name, scriptFullPath)
		return
	}

	tmp := strings.ReplaceAll(SCRIPT_YAML_PROMPT, "<TICK>", "`")
	tmp = strings.ReplaceAll(tmp, "<GCLOUD_COMMAND>", branch.Command)
	tmp = strings.ReplaceAll(tmp, "<GROUP>", branch.Group)
	prompt := strings.ReplaceAll(tmp, "<RESOURCE>", strings.ToLower(branch.Resource))
	log.Printf("CODEBOT PROMPT %s\r\n", prompt)

	// Delete then write the prompt file.
	promptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "prompt.txt")
	if _, err := os.Stat(promptPath); !errors.Is(err, os.ErrNotExist) {
		log.Println("COMMAND: cleaning up old prompt.txt")
		err = os.Remove(promptPath)
		if err != nil {
			log.Printf("Attempt to clean up prompt.txt failed with %v", err)
		}
	}
	log.Println("COMMAND: writing new prompt.txt")
	if err := os.WriteFile(promptPath, []byte(prompt), 0644); err != nil {
		log.Fatal(err)
	}

	// Run the LLM to generate the file.
	start := time.Now()
	log.Println("COMMAND: codebot --ui-type=prompt --prompt=prompt.txt")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	codebot := exec.CommandContext(ctx, "codebot", "--ui-type=prompt", "--prompt=prompt.txt")
	// codebot := exec.Command("codebot", "--ui-type=prompt", "--prompt=prompt.txt")
	codebot.Dir = workDir
	codebot.Stdout = &out
	if err := codebot.Run(); err != nil {
		stop := time.Now()
		diff := stop.Sub(start)
		log.Printf("CODEBOT GENERATE ERROR (%v): %q\n", diff, out.String())
		log.Fatal(err)
	}
	stop := time.Now()
	diff := stop.Sub(start)
	log.Printf("CODEBOT GENERATE (%v): %q\n", diff, out.String())

	// Check to see if the script file was created
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created\r\n", branch.Name, scriptFullPath)
		return
	}

	// Add the new file to the current branch.
	log.Printf("COMMAND: git add %s\r\n", scriptFile)
	gitadd := exec.Command("git", "add", scriptFile)
	gitadd.Dir = workDir
	gitadd.Stdout = &out
	if err := gitadd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH ADD: %q\n", out.String())

	// Commit the change to the current branch.
	log.Printf("COMMAND: git commit -m \"Adding LLM/gcloud generated test script.yaml for %s\"\r\n", branch.Name)
	gitcommit := exec.Command("git", "commit", "-m", fmt.Sprintf("\"Adding LLM/gcloud generated test script.yaml for %s\"", branch.Name))
	gitcommit.Dir = workDir
	gitcommit.Stdout = &out
	if err := gitcommit.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH COMMIT: %q\n", out.String())
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
	checkoutBranch(branch, workDir, out)

	// Check to see if the script file exists
	scriptFile := fmt.Sprintf("mock%s/testdata/%s/crud/script.yaml", branch.Group, branch.Resource)
	scriptFullPath := filepath.Join(workDir, scriptFile)
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing script %s\r\n", branch.Name, scriptFullPath)
		return
	}

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)
	if _, err := os.Stat(logFullPath); !errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s already exists\r\n", branch.Name, logFullPath)
		return
	}

	// Current HTTP Log generation is determenistic not ML generated.

	// Run the test to generate the log.
	start := time.Now()
	log.Printf("COMMAND: go test ./mockgcptests -run TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	test := exec.CommandContext(ctx, "go", "test", "./mockgcptests", "-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", branch.Group, branch.Resource))
	test.Dir = workDir
	test.Env = append(os.Environ(), "WRITE_GOLDEN_OUTPUT=1", "E2E_GCP_TARGET=real")
	test.Stdout = &out
	if err := test.Run(); err != nil {
		log.Printf("TEST GENERATE error: %q\n", out.String())
		// Currently ignoring error and just basing on if the _http.log was generated.
		// log.Fatal(err)
	}
	stop := time.Now()
	diff := stop.Sub(start)
	log.Printf("TEST GENERATE (%v): %q\n", diff, out.String())

	// Check to see if the script file was created
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s was not created\r\n", branch.Name, logFullPath)
		return
	}

	// Add the new file to the current branch.
	log.Printf("COMMAND: git add %s\r\n", logFullPath)
	gitadd := exec.Command("git", "add", logFullPath)
	gitadd.Dir = workDir
	gitadd.Stdout = &out
	if err := gitadd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH ADD: %q\n", out.String())

	// Commit the change to the current branch.
	log.Printf("COMMAND: git commit -m \"Adding mockgcptests generated _http.log for %s\"\r\n", branch.Name)
	gitcommit := exec.Command("git", "commit", "-m", fmt.Sprintf("\"Adding mockgcptests generated _http.log for %s\"", branch.Name))
	gitcommit.Dir = workDir
	gitcommit.Stdout = &out
	if err := gitcommit.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH COMMIT: %q\n", out.String())
}

const MOCK_SERVICE_GO_GEN string = `mock<SERVICE>/service.go
// +tool:mockgcp-service
// http.host: <HTTP_HOST>
// proto.service: <PROTO_SERVICE>`

const MOCK_RESOURCE_GO_GEN string = `mock<SERVICE>/<RESOURCE>.go
// +tool:mockgcp-support
// proto.service: <PROTO_SERVICE>
// proto.message: <PROTO_MESSAGE>`

func generateMockGo(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir, "mockgcp")

	var out strings.Builder
	checkoutBranch(branch, workDir, out)

	// Check to see if the http log file already exists
	logFile := fmt.Sprintf("mock%s/testdata/%s/crud/_http.log", branch.Group, branch.Resource)
	logFullPath := filepath.Join(workDir, logFile)
	if _, err := os.Stat(logFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, missing %s\r\n", branch.Name, logFullPath)
		return
	}

	// Check to see if the script file exists
	serviceGoFile := fmt.Sprintf("mock%s/service.go", branch.Group)
	serviceGoFullPath := filepath.Join(workDir, serviceGoFile)
	if _, err := os.Stat(serviceGoFullPath); errors.Is(err, os.ErrNotExist) {
		log.Printf("SKIPPING %s, %s already exists\r\n", branch.Name, serviceGoFullPath)
		return
	}

	tmp := strings.ReplaceAll(MOCK_SERVICE_GO_GEN, "<TICK>", "`")
	tmp = strings.ReplaceAll(tmp, "<SERVICE>", branch.Group)
	tmp = strings.ReplaceAll(tmp, "<HTTP_HOST>", branch.HostName)
	service_prompt := strings.ReplaceAll(tmp, "<PROTO_SERVICE>", branch.ProtoSvc)
	log.Printf("MOCK SERVICE GEN PROMPT %s\r\n", service_prompt)

	tmp = strings.ReplaceAll(MOCK_RESOURCE_GO_GEN, "<TICK>", "`")
	tmp = strings.ReplaceAll(tmp, "<SERVICE>", branch.Group)
	tmp = strings.ReplaceAll(tmp, "<RESOURCE>", strings.ToLower(branch.Resource))
	tmp = strings.ReplaceAll(tmp, "<PROTO_SERVICE>", branch.ProtoSvc)
	resource_prompt := strings.ReplaceAll(tmp, "<PROTO_MESSAGE>", branch.ProtoPath)
	log.Printf("MOCK RESOURCE GEN PROMPT %s\r\n", resource_prompt)

	// Delete then write the service prompt file.
	servicePromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "service_prompt.txt")
	if _, err := os.Stat(servicePromptPath); !errors.Is(err, os.ErrNotExist) {
		log.Println("COMMAND: cleaning up old service_prompt.txt")
		err = os.Remove(servicePromptPath)
		if err != nil {
			log.Printf("Attempt to clean up service_prompt.txt failed with %v", err)
		}
	}
	log.Println("COMMAND: writing new service_prompt.txt")
	if err := os.WriteFile(servicePromptPath, []byte(service_prompt), 0644); err != nil {
		log.Fatal(err)
	}

	// Delete then write the resource prompt file.
	resourcePromptPath := filepath.Join(opts.branchRepoDir, "mockgcp", "resource_prompt.txt")
	if _, err := os.Stat(resourcePromptPath); !errors.Is(err, os.ErrNotExist) {
		log.Println("COMMAND: cleaning up old resource_prompt.txt")
		err = os.Remove(resourcePromptPath)
		if err != nil {
			log.Printf("Attempt to clean up resource_prompt.txt failed with %v", err)
		}
	}
	log.Println("COMMAND: writing new resource_prompt.txt")
	if err := os.WriteFile(resourcePromptPath, []byte(resource_prompt), 0644); err != nil {
		log.Fatal(err)
	}

	// Run the controller builder to generate the service go file.
	serviceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), "service.go")
	if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
		start := time.Now()
		log.Printf("COMMAND: controllerbuilder prompt --src-dir %s --proto-dir %s/.build/third_party/googleapis/ --input-file=service_prompt.txt", opts.branchRepoDir, opts.branchRepoDir)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		service_go := exec.CommandContext(ctx, "controllerbuilder", "prompt",
			"--src-dir", opts.branchRepoDir,
			"--proto-dir", fmt.Sprintf("%s/.build/third_party/googleapis/", opts.branchRepoDir),
			"--input-file", "service_prompt.txt")
		service_go.Dir = workDir
		var serviceOut strings.Builder
		service_go.Stdout = &serviceOut
		if err := service_go.Run(); err != nil {
			log.Printf("MOCK SERVICE GENERATE error: %q\n", err)
			// Currently ignoring error and just basing on if the _http.log was generated.
			// log.Fatal(err)
		}
		stop := time.Now()
		diff := stop.Sub(start)
		if err := os.WriteFile(serviceFile, []byte(serviceOut.String()), 0755); err != nil {
			log.Printf("WRITE MOCK SERVICE %s error: %q\n", serviceFile, err)
		}
		log.Printf("MOCK SERVICE GENERATE (%v): %q\n", diff, serviceOut.String())

		// Check to see if the service go file was created
		if _, err := os.Stat(serviceFile); errors.Is(err, os.ErrNotExist) {
			log.Printf("SKIPPING %s, %s was not created\r\n", branch.Name, serviceFile)
			return
		}
	} else {
		log.Printf("SKIPPING generating service mock go, %s already exists\r\n", serviceFile)
	}

	// Run the controller builder to generate the resource go file.
	resourceFile := filepath.Join(workDir, fmt.Sprintf("mock%s", branch.Group), fmt.Sprintf("%s.go", strings.ToLower(branch.Resource)))
	if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
		start := time.Now()
		log.Printf("COMMAND: controllerbuilder prompt --src-dir %s --proto-dir %s/.build/third_party/googleapis/ --input-file=resource_prompt.txt", opts.branchRepoDir, opts.branchRepoDir)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		resource_go := exec.CommandContext(ctx, "controllerbuilder", "prompt",
			"--src-dir", opts.branchRepoDir,
			"--proto-dir", fmt.Sprintf("%s/.build/third_party/googleapis/", opts.branchRepoDir),
			"--input-file", "resource_prompt.txt")
		resource_go.Dir = workDir
		var resourceOut strings.Builder
		resource_go.Stdout = &resourceOut
		if err := resource_go.Run(); err != nil {
			log.Printf("MOCK RESOURCE GENERATE error: %q\n", err)
			// Currently ignoring error and just basing on if the _http.log was generated.
			// log.Fatal(err)
		}
		stop := time.Now()
		diff := stop.Sub(start)
		if err := os.WriteFile(resourceFile, []byte(resourceOut.String()), 0755); err != nil {
			log.Printf("WRITE MOCK RESOURCE %s error: %q\n", resourceFile, err)
		}
		log.Printf("MOCK RESOURCE GENERATE (%v): %q\n", diff, resourceOut.String())

		// Check to see if the service go file was created
		if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
			log.Printf("SKIPPING %s, %s was not created\r\n", branch.Name, resourceFile)
			return
		}
	} else {
		log.Printf("SKIPPING generating resource mock go, %s already exists\r\n", resourceFile)
	}

	// Add the new files to the current branch.
	log.Printf("COMMAND: git add %s %s\r\n", serviceFile, resourceFile)
	gitadd := exec.Command("git", "add", serviceFile, resourceFile)
	gitadd.Dir = workDir
	gitadd.Stdout = &out
	if err := gitadd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH ADD: %q\n", out.String())

	// Commit the change to the current branch.
	log.Printf("COMMAND: git commit -m \"Adding mock service and resource for %s\"\r\n", branch.Name)
	gitcommit := exec.Command("git", "commit", "-m", fmt.Sprintf("\"Adding mock service and resource for %s\"", branch.Name))
	gitcommit.Dir = workDir
	gitcommit.Stdout = &out
	if err := gitcommit.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH COMMIT: %q\n", out.String())
}
