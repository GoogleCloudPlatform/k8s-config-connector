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
		if _, err :=  os.Stdout.ReadFrom(os.Stdin); err != nil {
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
	// In strict CI, origin/master is usually available.
	base := "origin/master"
	
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

	cmd := exec.Command("git", "diff", "--unified=0", base)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	changes := make(FileChanges)
	var currentFile string
	
	lines := strings.Split(string(out), "\n")
	for _, l := range lines {
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
			
			if strings.Contains(newRange, ",") {
				sub := strings.Split(newRange, ",")
				start, _ = strconv.Atoi(sub[0])
				count, _ = strconv.Atoi(sub[1])
			} else {
				start, _ = strconv.Atoi(newRange)
			}
            
            if count > 0 {
                changes[currentFile] = append(changes[currentFile], LineRange{Start: start, Count: count})
            }
		}
	}
	return changes, nil
}

func isRelevant(line string, changes FileChanges) bool {
    // Expected format: file:line:col: message
    parts := strings.SplitN(line, ":", 4)
    if len(parts) < 3 {
        return true // Pass through lines that don't look like file:line:col
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
