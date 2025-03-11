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
)

// isExecutable checks if the given file path is executable.
func isExecutable(info os.FileInfo) bool {
	// Check if it's a regular file and has at least one execute permission bit set
	mode := info.Mode()
	return mode.IsRegular() && (mode.Perm()&0111 != 0)
}

// executeScript runs the script located at scriptPath
func executeScript(opts *RunnerOptions, branch Branch) error {
	ctx := context.TODO()
	workDir := opts.branchRepoDir

	if opts.scriptFile == "" {
		return fmt.Errorf("no script file specified")
	}

	scriptPath := opts.scriptFile
	checkoutBranch(ctx, branch, workDir)

	// Resolve absolute path for the script
	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to resolve script path: %w", err)
	}

	// Check if script exists
	info, err := os.Stat(absPath); os.IsNotExist(err) 
	if err != nil {
		return fmt.Errorf("script does not exist: %s", absPath)
	}

	if !isExecutable(info) {
		return fmt.Errorf("script is not executable: %s", absPath)
	}

	cfg := CommandConfig{
		Name: "Execute script",
		Cmd:  "/bin/bash",
		Args: []string{absPath},
		WorkDir: workDir,
		Env: map[string]string{
			"PATH": os.Getenv("PATH"),
		},
		MaxRetries:   1,
		RetryBackoff: 0, // No retry logic for script execution
	}

	_, _, err = executeCommand(opts, cfg)
	if err != nil {
		log.Printf("SCRIPT EXECUTION FAILED: %v\n", err)
		return err
	}

	log.Println("SCRIPT EXECUTION SUCCESS")
	return nil
}

