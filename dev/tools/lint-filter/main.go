// Copyright 2026 Google LLC
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// LineRange represents a range of changed lines
type LineRange struct {
	Start int
	Count int
}

// FileChanges maps filenames to their changed line ranges
type FileChanges map[string][]LineRange

func main() {
	changes, err := getChanges()
	if err != nil {
		// If git fails (e.g. no git repo, detached head issues), we warn but pass everything through?
		// Or fail? Let's fail safe: pass everything through if we can't detect changes (e.g. initial commit or CI weirdness),
		// OR simpler: just print error and exit.
		// For now, let's print error and output nothing (assume filtering failed implies safe to ignore or we want to fix env).
		fmt.Fprintf(os.Stderr, "Warning: could not determine changed lines: %v. Linting all files.\n", err)
		// Fallback: copy stdin to stdout
		if _, err := os.Stdout.ReadFrom(os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "Error copying stdin: %v\n", err)
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	foundIssues := false
	for scanner.Scan() {
		line := scanner.Text()
		if isRelevant(line, changes) {
			fmt.Println(line)
			foundIssues = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	if foundIssues {
		os.Exit(1)
	}
}

func getChanges() (FileChanges, error) {
	// Try to find a merge base or common ancestor.
	// Allow overriding via environment variable
	base := os.Getenv("LINT_FILTER_GIT_BASE")

	if base == "" {
		// Default to origin/master if available
		base = "origin/master"
		// Check if origin/master exists
		cmdCheck := exec.Command("git", "rev-parse", "--verify", "origin/master")
		if err := cmdCheck.Run(); err != nil {
			// Fallback to master
			base = "master"
			cmdCheck = exec.Command("git", "rev-parse", "--verify", "master")
			if err := cmdCheck.Run(); err != nil {
				// Fallback to HEAD~1 (assuming standard commit workflow)
				base = "HEAD~1"
			}
		}
	}
	cmd := exec.Command("git", "diff", "--unified=0", base)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start git diff command: %w", err)
	}

	changes := make(FileChanges)
	var currentFile string

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "+++ b/") {
			currentFile = strings.TrimPrefix(l, "+++ b/")
		} else if strings.HasPrefix(l, "@@") {
			// @@ -oldStart,oldLen +newStart,newLen @@
			parts := strings.Split(l, " ")
			if len(parts) < 3 {
				continue
			}
			newRange := parts[2] // +newStart,newLen

			// Handle cases where +newStart is missing the comma (count=1)
			if !strings.HasPrefix(newRange, "+") {
				continue // Should start with +
			}
			newRange = strings.TrimPrefix(newRange, "+")

			start := 0
			count := 1

			var parseErr error
			if strings.Contains(newRange, ",") {
				sub := strings.Split(newRange, ",")
				start, parseErr = strconv.Atoi(sub[0])
				if parseErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to parse start line %q in git diff: %v\n", sub[0], parseErr)
					continue
				}
				count, parseErr = strconv.Atoi(sub[1])
				if parseErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to parse count %q in git diff: %v\n", sub[1], parseErr)
					continue
				}
			} else {
				start, parseErr = strconv.Atoi(newRange)
				if parseErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to parse start line %q in git diff: %v\n", newRange, parseErr)
					continue
				}
			}

			if count > 0 {
				changes[currentFile] = append(changes[currentFile], LineRange{Start: start, Count: count})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading git diff output: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("git diff command failed: %w", err)
	}
	return changes, nil
}

func isRelevant(line string, changes FileChanges) bool {
	// "go run" might print "exit status N" to stderr, which is part of the input to lint-filter.
	// We need to ignore "exit status 3", as it means violations found.
	if line == "exit status 3" {
		return false
	}

	parts := strings.SplitN(line, ":", 4)
	if len(parts) < 3 {
		return true // Pass through lines that don't look like file:line:col (e.g. potential panic or other error)
	}

	file := parts[0]

	// Normalize file path: if it's absolute, make it relative to the current working directory.
	// Assumes the current working directory is the repository root.
	if strings.HasPrefix(file, "/") {
		wd, err := os.Getwd()
		if err == nil && strings.HasPrefix(file, wd+string(os.PathSeparator)) {
			file = strings.TrimPrefix(file, wd+string(os.PathSeparator))
		}
	}

	// Attempt to parse line number
	lineNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return true // Not a diagnostic line, pass through
	}

	// Check if the normalized file path and line number are relevant.
	if checkFile(file, lineNum, changes) {
		return true
	}

	return false
}

func checkFile(file string, lineNum int, changes FileChanges) bool {
	ranges, ok := changes[file]
	if !ok {
		return false
	}
	for _, r := range ranges {
		if lineNum >= r.Start && lineNum < r.Start+r.Count {
			return true
		}
	}
	return false
}
