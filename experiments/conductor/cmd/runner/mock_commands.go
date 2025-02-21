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

When you have completed, please output the name of the test script you have created, in a JSON format like this:

{ "path_to_created_test": "mock<GROUP>/testdata/<RESOURCE>/crud/script.yaml" }`

func createScriptYaml(opts *RunnerOptions, branch Branch) {
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

	cdRepoBranchDir(opts, "mockgcp", stdin, stdout)

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
