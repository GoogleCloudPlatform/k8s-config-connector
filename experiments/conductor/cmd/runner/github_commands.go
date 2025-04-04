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
	"strings"
	"time"
)

func createGithubBranch(opts *RunnerOptions, branch Branch) {
	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDirBash(opts, "", stdin, stdout)

	// Always create resource branches from master
	log.Printf("COMMAND: git checkout master and echo done\r\n")
	if _, err = stdin.Write([]byte("git checkout master && echo done\n")); err != nil {
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
	log.Printf("CHECKOUT MASTER %s\r\n", msg)

	// Check to see if the branch already exists
	log.Printf("COMMAND: git branch --list %s and echo done\r\n", branch.Local)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git branch --list %s && echo done\n", branch.Local))); err != nil {
		log.Fatal(err)
	}
	done = false
	outBuffer = make([]byte, 1000)
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("CHECK LOCAL BRANCH %s\r\n", msg)
	exists := strings.Contains(msg, branch.Local)

	// Create the actual branch
	// git checkout -b "${RELEASE}"
	if !exists {
		log.Printf("COMMAND: git checkout -b %s and echo done\r\n", branch.Local)
		if _, err = stdin.Write([]byte(fmt.Sprintf("git checkout -b %s && echo done\n", branch.Local))); err != nil {
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
		log.Printf("LOCAL BRANCH CREATE %s\r\n", msg)
	}
	// Currently not bothering with generating main line changes.
	// git push upstream "${RELEASE}"
	// git branch -f ${RELEASE} ${RELEASE}

}

func deleteGithubBranch(opts *RunnerOptions, branch Branch) {
	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDirBash(opts, "", stdin, stdout)

	// Change to the master branch so we're not the branch we are deleting
	log.Printf("COMMAND: git checkout master and echo done\r\n")
	if _, err = stdin.Write([]byte("git checkout master && echo done\n")); err != nil {
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
	log.Printf("CHECKOUT MASTER %s\r\n", msg)

	// Check to see if the branch already exists
	log.Printf("COMMAND: git branch --list %s and echo done\r\n", branch.Local)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git branch --list %s && echo done\n", branch.Local))); err != nil {
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
	log.Printf("CHECK LOCAL BRANCH %s\r\n", msg)
	exists := strings.Contains(msg, branch.Local)

	// Ask for final confirmation
	fmt.Printf("Are you sure you want to delete the branch %s? [y/N]: ", branch.Local)
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		log.Printf("Error reading input: %v", err)
		return
	}
	if response != "y" && response != "Y" {
		log.Printf("Skipping delete for branch %s", branch.Local)
		return
	}
	// Delete the actual branch
	// git checkout -b "${RELEASE}"
	if exists {
		log.Printf("COMMAND: git branch -D %s and echo done\r\n", branch.Local)
		if _, err = stdin.Write([]byte(fmt.Sprintf("git branch -D %s && echo done\n", branch.Local))); err != nil {
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
		log.Printf("LOCAL BRANCH DELETE %s\r\n", msg)
	}
}

func pushBranch(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	// Determine the remote branch name
	remoteBranch := branch.Local
	if opts.branchSuffix != "" {
		remoteBranch = branch.Local + opts.branchSuffix
		log.Printf("Using remote branch name with suffix: %s", remoteBranch)
	} else if branch.Remote != "" && branch.Remote != branch.Local {
		remoteBranch = branch.Remote
		log.Printf("Using configured remote branch name: %s", remoteBranch)
	}

	// Run git push command with force flag
	cfg := CommandConfig{
		Name: "Git push",
		Cmd:  "git",
		Args: []string{
			"push",
			"origin",
			fmt.Sprintf("%s:%s", branch.Local, remoteBranch),
			"--force",
		},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}
	output, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("Git push error for branch %s: %v", branch.Name, err)
	}
	return nil, &output, err
}

func makeReadyPR(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	// Run git push command with force flag
	if opts.skipMakeReadyPR {
		log.Printf("Skipping make ready-pr for branch %s (--skip-makereadypr flag is set)", branch.Name)
		return []string{"."}, nil, nil
	}

	cfg := CommandConfig{
		Name: "Make ready PR",
		Cmd:  "make",
		Args: []string{
			"ready-pr",
		},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
		Timeout:     15 * time.Minute,
	}
	output, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("Make ready PR error for branch %s: %v", branch.Name, err)
	}
	return []string{"."}, &output, err
}
