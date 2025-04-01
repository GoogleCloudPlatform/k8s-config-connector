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
)

// runFuzzerTests runs the fuzzer tests for a branch
func runFuzzerTests(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Running fuzzer tests for branch %s", branch.Name)

	// Create the path to the fuzzer file
	fuzzerPath := filepath.Join("pkg", "controller", "direct", branch.Group, branch.Resource+"_fuzzer.go")
	fullFuzzerPath := filepath.Join(opts.branchRepoDir, fuzzerPath)

	// Create the path to the mapper file
	mapperPath := filepath.Join("pkg", "controller", "direct", branch.Group, "mapper.generated.go")
	fullMapperPath := filepath.Join(opts.branchRepoDir, mapperPath)

	// Check if fuzzer file exists
	if _, err := os.Stat(fullFuzzerPath); err != nil {
		log.Printf("Fuzzer file not found at %s, skipping fuzzer tests", fullFuzzerPath)
		return []string{}, nil, fmt.Errorf("fuzzer file not found: %s", fullFuzzerPath)
	}

	// Check if mapper file exists
	if _, err := os.Stat(fullMapperPath); err != nil {
		log.Printf("Mapper file not found at %s, skipping fuzzer tests", fullMapperPath)
		return []string{}, nil, fmt.Errorf("mapper file not found: %s", fullMapperPath)
	}

	// Run the fuzzer tests
	cfg := CommandConfig{
		Name: "Run Fuzzer Tests",
		Cmd:  "go",
		Args: []string{
			"test", "-v",
			"./pkg/fuzztesting/fuzztests/",
			"-fuzz=FuzzAllMappers",
			"-fuzztime", "60s",
		},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}

	results, err := executeCommand(opts, cfg)
	return []string{fuzzerPath}, &results, err
}

const FIX_FUZZER_FAILURES string = `Please fix the fuzzer file at: ${FUZZER_FILE}  and mapper file at: ${MAPPER_FILE} based on the failures seen in the fuzzer test run.

The fuzzer test is failing with the error output shown below. Your task is to fix the issues in the fuzzer file
to make the test pass.

Start with the errors from the top of the output and work your way down.
If you see a change that needs to be made, make the change and run the fuzzer test again.
Fixing from top to bottom is important and will help you make progress.

Once the changes are applied, use the RunShellCommand command to run the fuzzer test again to make sure the changes are valid.
Use the output of the command to determine if the code changes are valid.

RunShellCommand:
go test -v ./pkg/fuzztesting/fuzztests/ -fuzz=FuzzAllMappers -fuzztime 10s

Hints:

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to edit the file.
* Look for issues with defaults, required fields, or field constraints in the fuzzer.
* Pay attention to differences between fields in the mapper and the fuzzer.
* Make sure the fuzzer handles all field types correctly.
* Check any list, map, or pointer fields that might need special handling.

The results of the fuzzer test run are:

EXITCODE: ${TEST_OUTPUT_EXITCODE}

STDERR:
${TEST_OUTPUT_STDERR}

STDOUT:
${TEST_OUTPUT_STDOUT}
`

// fixFuzzerFailures attempts to fix issues in the fuzzer file based on test failures
func fixFuzzerFailures(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Fixing fuzzer failures for branch %s", branch.Name)

	// If tests passed, no need to fix anything
	if execResults.ExitCode == 0 {
		log.Printf("Fuzzer tests pass for branch %s, no fixes needed", branch.Name)
		return []string{}, nil, nil
	}

	// Create the path to the fuzzer file
	fuzzerPath := filepath.Join("pkg", "controller", "direct", branch.Group, branch.Resource+"_fuzzer.go")
	fullFuzzerPath := filepath.Join(opts.branchRepoDir, fuzzerPath)

	// Create the path to the mapper file
	mapperPath := filepath.Join("pkg", "controller", "direct", branch.Group, "mapper.generated.go")
	fullMapperPath := filepath.Join(opts.branchRepoDir, mapperPath)

	// Check if fuzzer file exists
	if _, err := os.Stat(fullFuzzerPath); err != nil {
		log.Printf("Fuzzer file not found at %s, cannot fix", fullFuzzerPath)
		return []string{}, nil, fmt.Errorf("fuzzer file not found: %s", fullFuzzerPath)
	}

	// Check if mapper file exists
	if _, err := os.Stat(fullMapperPath); err != nil {
		log.Printf("Mapper file not found at %s, cannot fix", fullMapperPath)
		return []string{}, nil, fmt.Errorf("mapper file not found: %s", fullMapperPath)
	}

	// Create prompt with file contents and test output
	prompt := strings.ReplaceAll(FIX_FUZZER_FAILURES, "${FUZZER_FILE}", fuzzerPath)
	prompt = strings.ReplaceAll(prompt, "${MAPPER_FILE}", mapperPath)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_EXITCODE}", fmt.Sprintf("%d", execResults.ExitCode))
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDERR}", execResults.Stderr)
	prompt = strings.ReplaceAll(prompt, "${TEST_OUTPUT_STDOUT}", execResults.Stdout)

	// Run codebot to fix the issues
	cfg := CommandConfig{
		Name:         "Fix Fuzzer Failures",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		Stdin:        strings.NewReader(prompt),
		WorkDir:      opts.branchRepoDir,
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	results, err := executeCommand(opts, cfg)
	return []string{fuzzerPath}, &results, err
}
