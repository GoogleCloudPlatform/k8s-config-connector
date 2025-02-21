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
	"fmt"
	"log"
	"strings"
)

func createGithubBranch(opts *RunnerOptions, branch Branch) {
	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDir(opts, "", stdin, stdout)

	// Check to see if the branch already exists
	log.Printf("COMMAND: git branch --list %s and echo done\r\n", branch.Local)
	if _, err = stdin.Write([]byte(fmt.Sprintf("git branch --list %s && echo done\n", branch.Local))); err != nil {
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

	cdRepoBranchDir(opts, "", stdin, stdout)

	// Change to the master branch so we're not the branch we are deleting
	log.Printf("COMMAND: git checkout -b %s and echo done\r\n", branch.Local)
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
	log.Printf("BRANCH CHECKOUT %s\r\n", msg)

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
