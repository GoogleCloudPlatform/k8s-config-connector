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

const GENERATOR_SCRIPT_TEMPLATE = `#!/bin/bash
# Copyright 2025 Google LLC
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

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-types \
    --service <PROTO_PACKAGE> \
    --api-version <CRD_GROUP>/<CRD_VERSION> \
    --resource <CRD_KIND>:<PROTO_RESOURCE>

go run . generate-mapper \
    --service <PROTO_PACKAGE> \
    --api-version <CRD_GROUP>/<CRD_VERSION>

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/<SERVICE>/
`

const UPDATE_GENERATE_SCRIPT_PROMPT = `
Please update the apis/<SERVICE>/v1alpha1/generate.sh script for the <SERVICE> API to generate the CRD for the <CRD_KIND> resource.

The generate.sh script is located at apis/<SERVICE>/v1alpha1/generate.sh.

Add the parameter <TICK> --resource <CRD_KIND>:<PROTO_RESOURCE><TICK> to the <TICK>go run . generate-types --api-version <CRD_GROUP>/<CRD_VERSION>  <TICK> command.

At the end of the script, ensure the following lines are present:

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/<SERVICE>/
`

func generateCRDScripts(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir)

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Create the apis/<service>/<version> directory
	serviceDir := filepath.Join(workDir, "apis", branch.Group, "v1alpha1")
	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Create or update generate.sh
	scriptPath := filepath.Join(serviceDir, "generate.sh")
	// Check if generate.sh already exists.
	if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
		// File doesn't exist, use template approach
		log.Printf("Creating new generate.sh at %s", scriptPath)

		// Replace template markers with actual values and write to file
		writeTemplateToFile(branch, scriptPath, GENERATOR_SCRIPT_TEMPLATE)
	} else {
		// File exists, use codebot to update it
		log.Printf("Updating existing generate.sh at %s", scriptPath)

		// Delete then write the prompt file.
		promptPath := filepath.Join(workDir, "mockgcp", "crdgen_prompt.txt")
		writeTemplateToFile(branch, promptPath, UPDATE_GENERATE_SCRIPT_PROMPT)

		// Run codebot
		start := time.Now()
		log.Println("COMMAND: codebot --ui-type=prompt --prompt=crdgen_prompt.txt")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		codebot := exec.CommandContext(ctx, "codebot", "--ui-type=prompt", "--prompt=mockgcp/crdgen_prompt.txt")
		codebot.Dir = workDir
		codebot.Stdout = &out
		codebot.Stderr = &out
		if err := codebot.Run(); err != nil {
			stop := time.Now()
			diff := stop.Sub(start)
			log.Printf("CODEBOT GENERATE ERROR (%v): \n", diff)
			printCommandOutput(out.String())
			out.Reset()
			log.Fatal(err)
		}
		stop := time.Now()
		diff := stop.Sub(start)
		log.Printf("CODEBOT GENERATE output in %v: \n", diff)
		printCommandOutput(out.String())
		out.Reset()
	}

	// Add and commit changes
	scriptRelativePath := fmt.Sprintf("apis/%s/v1alpha1/generate.sh", branch.Group)
	gitAdd(workDir, &out, scriptRelativePath)
	gitCommit(workDir, &out, fmt.Sprintf("add/update crd generation script for %s", branch.Group))

	// Run the generator script
	log.Printf("Running generator script %s", scriptPath)
	cmd := exec.Command(scriptPath)
	cmd.Dir = workDir
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Generator script error: \n")
		printCommandOutput(out.String())
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("Generator script output: \n")
	printCommandOutput(out.String())
	out.Reset()
	// Add and commit generated files
	gitAdd(workDir, &out,
		fmt.Sprintf("apis/%s/v1alpha1/", branch.Group),
		fmt.Sprintf("pkg/controller/direct/%s/", branch.Group),
		"config/crds/resources/")
	gitCommit(workDir, &out, fmt.Sprintf("autogenerated types and CRDs using %s", scriptRelativePath))
}

func generateSpecStatus(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir)

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Run controllerbuilder to generate spec and status
	log.Printf("Generating spec and status for %s", branch.Name)
	cmd := exec.Command("controllerbuilder", "prompt",
		"--src-dir", workDir,
		"--proto-dir", filepath.Join(workDir, ".build/third_party/googleapis/"))
	cmd.Dir = workDir
	stdinInput := fmt.Sprintf("// +kcc:proto=%s.%s\n", branch.ProtoSvc, branch.Proto)
	cmd.Stdin = strings.NewReader(stdinInput)
	cmd.Stdout = &out
	cmd.Stderr = &out

	log.Printf("Running command: controllerbuilder prompt --src-dir %s --proto-dir %s",
		workDir, filepath.Join(workDir, ".build/third_party/googleapis/"))
	log.Printf("With stdin input: %s", stdinInput)

	if err := cmd.Run(); err != nil {
		log.Printf("Spec/Status generation error:\n")
		printCommandOutput(out.String())
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("Spec/Status generation output: \n")
	printCommandOutput(out.String())
	out.Reset()

	// Add and commit changes
	gitAdd(workDir, &out, fmt.Sprintf("apis/%s/v1alpha1/%s_types.go", branch.Group, strings.ToLower(branch.Resource)))
	gitCommit(workDir, &out, fmt.Sprintf("%s: Update types from generated", branch.Kind))
}

func generateFuzzer(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := filepath.Join(opts.branchRepoDir)

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Generate fuzzer file
	fuzzerDir := filepath.Join(workDir, "pkg/controller/direct", branch.Group)
	if err := os.MkdirAll(fuzzerDir, 0755); err != nil {
		log.Fatal(err)
	}

	fuzzerPath := filepath.Join(fuzzerDir, fmt.Sprintf("%s_fuzzer.go", strings.ToLower(branch.Resource)))
	cmd := exec.Command("controllerbuilder", "prompt",
		"--src-dir", workDir,
		"--proto-dir", filepath.Join(workDir, ".build/third_party/googleapis/"))
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(fmt.Sprintf(`// +tool:fuzz-gen
// proto.message: %s
`, branch.ProtoMsg))
	var fuzzerOut strings.Builder
	cmd.Stdout = &fuzzerOut
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Fuzzer generation error: \n")
		printCommandOutput(out.String())
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("Fuzzer generation output: \n")
	printCommandOutput(fuzzerOut.String())
	fuzzerOut.Reset()

	if err := os.WriteFile(fuzzerPath, []byte(fuzzerOut.String()), 0644); err != nil {
		log.Fatal(err)
	}

	// Update register.go to import the new package
	registerPath := filepath.Join(workDir, "pkg/controller/direct/register/register.go")
	importLine := fmt.Sprintf(`_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/%s"`, branch.Group)

	// Use codebot to add the import
	cmd = exec.Command("codebot", "--prompt=/dev/stdin")
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(fmt.Sprintf("Add an unnamed (_) go import for %s to the imports in %s", importLine, registerPath))
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Import addition error: \n")
		printCommandOutput(out.String())
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("Import addition output: \n")
	printCommandOutput(out.String())
	out.Reset()

	// Add and commit changes
	gitAdd(workDir, &out,
		fmt.Sprintf("pkg/controller/direct/%s/%s_fuzzer.go", branch.Group, strings.ToLower(branch.Resource)),
		"pkg/controller/direct/register/register.go")
	gitCommit(workDir, &out, fmt.Sprintf("%s: Create fuzz test", branch.Kind))
}
