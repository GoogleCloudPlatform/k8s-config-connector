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
	"regexp"
	"strings"
	"time"
)

const (
	COMMIT_MSG_9A  = "{{kind}}: Normalize api checks"
	COMMIT_MSG_9B  = "make ready-pr"
	COMMIT_MSG_9C  = "push branch"
	COMMIT_MSG_10  = "mockgcp: create test script for {{command}}"
	COMMIT_MSG_11  = "mockgcp: golden output for TestScripts/mock{{group}}/testdata/{{resource}}/crud"
	COMMIT_MSG_12  = "{{kind}}: Add generated mock files"
	COMMIT_MSG_13  = "mockgcp: Add mock{{group}} service to mock_http_roundtrip.go"
	COMMIT_MSG_14  = "{{group}}: Add proto generation to makefile"
	COMMIT_MSG_15  = "chore: Build and add generated proto files"
	COMMIT_MSG_16  = "{{kind}}: Capture mock golden output"
	COMMIT_MSG_17  = "Verify and Fix mock tests"
	COMMIT_MSG_20  = "{{kind}}: Add generated types"
	COMMIT_MSG_21A = "{{kind}}: Add spec and status to generated type"
	COMMIT_MSG_21B = "{{kind}}: Add parent to generated type"
	COMMIT_MSG_21C = "{{kind}}: Adjust identity parent"
	COMMIT_MSG_21D = "{{kind}}: Regenerate types"
	COMMIT_MSG_21E = "{{kind}}: Remove Name Field"
	COMMIT_MSG_21F = "{{kind}}: Move Etag Field"
	COMMIT_MSG_21G = "{{kind}}: Add Required Field Tags"
	COMMIT_MSG_22  = "{{kind}}: Add generated CRD"
	COMMIT_MSG_23  = "{{kind}}: Add generated mapper"
	COMMIT_MSG_24  = "{{kind}}: Add generated fuzzer"
	COMMIT_MSG_25  = "{{kind}}: Verify and Fix fuzzer tests"
	COMMIT_MSG_26  = "{{kind}}: Verify and Fix API checks"
	COMMIT_MSG_40  = "{{kind}}: Add controller client"
	COMMIT_MSG_41  = "{{kind}}: Add controller"
	COMMIT_MSG_42  = "{{kind}}: Build and fix controller"
	COMMIT_MSG_43A = "{{kind}}: Add controller identity"
	COMMIT_MSG_43B = "{{kind}}: Add controller reference"
	COMMIT_MSG_44A = "{{kind}}: Create minimal test"
	COMMIT_MSG_44B = "{{kind}}: Support for testing with mockgcp"
	COMMIT_MSG_45  = "{{kind}}: Record golden logs for real GCP tests"
	COMMIT_MSG_46  = "{{kind}}: Verify and Fix real GCP tests"
	COMMIT_MSG_47  = "{{kind}}: Record golden logs for mock GCP tests"
	COMMIT_MSG_48  = "{{kind}}: Verify and Fix mock GCP tests"
	COMMIT_MSG_50  = "{{kind}}: Move existing test to subdirectory"
)

var REGEX_MSG_9A = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_9A))
var REGEX_MSG_9B = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_9B))
var REGEX_MSG_9C = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_9C))
var REGEX_MSG_10 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_10))
var REGEX_MSG_11 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_11))
var REGEX_MSG_12 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_12))
var REGEX_MSG_13 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_13))
var REGEX_MSG_14 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_14))
var REGEX_MSG_15 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_15))
var REGEX_MSG_16 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_16))
var REGEX_MSG_17 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_17))
var REGEX_MSG_20 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_20))
var REGEX_MSG_21A = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21A))
var REGEX_MSG_21B = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21B))
var REGEX_MSG_21C = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21C))
var REGEX_MSG_21D = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21D))
var REGEX_MSG_21E = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21E))
var REGEX_MSG_21F = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21F))
var REGEX_MSG_21G = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_21G))
var REGEX_MSG_22 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_22))
var REGEX_MSG_23 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_23))
var REGEX_MSG_24 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_24))
var REGEX_MSG_25 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_25))
var REGEX_MSG_26 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_26))
var REGEX_MSG_40 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_40))
var REGEX_MSG_41 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_41))
var REGEX_MSG_42 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_42))
var REGEX_MSG_43A = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_43A))
var REGEX_MSG_43B = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_43B))
var REGEX_MSG_44A = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_44A))
var REGEX_MSG_44B = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_44B))
var REGEX_MSG_45 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_45))
var REGEX_MSG_46 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_46))
var REGEX_MSG_47 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_47))
var REGEX_MSG_48 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_48))
var REGEX_MSG_50 = regexp.MustCompile(convertCommitMsgToRegex(COMMIT_MSG_50))

func skipPost21A(msg string) bool {
	return REGEX_MSG_21A.MatchString(msg)
}

func skipPost21B(msg string) bool {
	return skipPost21A(msg) || REGEX_MSG_21A.MatchString(msg)
}

func skipPost21C(msg string) bool {
	return skipPost21B(msg) || REGEX_MSG_21A.MatchString(msg)
}

func skipPost21D(msg string) bool {
	return skipPost21C(msg) || REGEX_MSG_21A.MatchString(msg)
}

func skipPost21E(msg string) bool {
	return skipPost21D(msg) || REGEX_MSG_21A.MatchString(msg)
}

func skipPost21F(msg string) bool {
	return skipPost21E(msg) || REGEX_MSG_21A.MatchString(msg)
}

func convertCommitMsgToRegex(msg string) string {
	tmp := strings.ReplaceAll(msg, "{{kind}}", "[a-zA-Z]+")
	tmp = strings.ReplaceAll(tmp, "{{command}}", "") // handled by terminating ".*"
	tmp = strings.ReplaceAll(tmp, "{{group}}", "[a-zA-Z]+")
	tmp = strings.ReplaceAll(tmp, "{{resource}}", "[a-zA-Z]+")
	return "^ *" + tmp + ".*$"
}

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
