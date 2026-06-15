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
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	listFile       = flag.String("list-file", "", "File containing the list of tests")
	j              = flag.Int("j", 4, "Number of parallel processes")
	baseTest       = flag.String("base-test", "", "Base test name, e.g. TestE2EScript/scenarios")
	checkUnchanged = flag.Bool("check-unchanged", false, "Fail if the list file is modified (i.e. missing tests found)")
	timeout        = flag.Duration("timeout", 5*time.Minute, "Timeout for each test")
)

func main() {
	flag.Parse()
	if *listFile == "" || *baseTest == "" || len(flag.Args()) == 0 {
		log.Fatalf("Usage: paralleltestrunner -list-file=<file> -base-test=<base> <cmd> [args...]")
	}

	tests := readList(*listFile)

	success := true
	var wg sync.WaitGroup
	sem := make(chan struct{}, *j)
	var mu sync.Mutex

	if len(tests) > 0 {
		for _, test := range tests {
			wg.Add(1)
			sem <- struct{}{}
			go func(t string) {
				defer wg.Done()
				defer func() { <-sem }()

				parts := strings.Split(*baseTest, "/")
				parts = append(parts, strings.Split(t, "/")...)
				exactRun := "^" + strings.Join(parts, "$/^") + "$"

				ctx, cancel := context.WithTimeout(context.Background(), *timeout)
				defer cancel()

				cmd := exec.CommandContext(ctx, flag.Args()[0], flag.Args()[1:]...)
				cmd.Env = os.Environ()
				cmd.Env = append(cmd.Env, "RUN_TESTS="+exactRun)
				
				// Configure process group setup to ensure that any orphaned subprocesses 
				// started by the subtest execution (like shell scripts or test binaries) 
				// are cleanly terminated if the test times out. Without this, orphaned 
				// children can hold stdout/stderr pipes open, causing cmd.CombinedOutput() 
				// to hang indefinitely and eventually time out the entire CI job.
				setProcessGroup(cmd)

				startTime := time.Now()
				output, err := cmd.CombinedOutput()
				duration := time.Since(startTime)

				artifactsDir := os.Getenv("ARTIFACTS")
				if artifactsDir == "" {
					repoRoot := os.Getenv("REPO_ROOT")
					if repoRoot == "" {
						repoRoot = "."
					}
					artifactsDir = filepath.Join(repoRoot, ".build")
				}
				logDir := filepath.Join(artifactsDir, "test-output", t)
				if mkdirErr := os.MkdirAll(logDir, 0755); mkdirErr != nil {
					log.Printf("Failed to create log dir %s: %v", logDir, mkdirErr)
				}
				logPath := filepath.Join(logDir, "test.log")
				if writeErr := os.WriteFile(logPath, output, 0644); writeErr != nil {
					log.Printf("Failed to write log file %s: %v", logPath, writeErr)
				}

				mu.Lock()
				defer mu.Unlock()
				if err != nil {
					if ctx.Err() == context.DeadlineExceeded {
						fmt.Printf("FAIL:TIMEOUT: %s (log: %s)\n", t, logPath)
					} else {
						fmt.Printf("FAIL: %s (log: %s)\n", t, logPath)
					}
					fmt.Printf("--- LOGS FOR FAILING TEST %s ---\n", t)
					os.Stdout.Write(output)
					if !bytes.HasSuffix(output, []byte("\n")) {
						fmt.Println()
					}
					fmt.Printf("--- END LOGS FOR FAILING TEST %s ---\n", t)
					success = false
				} else {
					fmt.Printf("PASS: %s (%v)\n", t, duration.Round(time.Millisecond))
				}
			}(test)
		}
		wg.Wait()
	}

	fmt.Printf("Running catch-all to find any new tests...\n")

	cmd := exec.Command(flag.Args()[0], flag.Args()[1:]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "RUN_TESTS="+*baseTest)
	cmd.Env = append(cmd.Env, "SKIP_ALL=1")

	catchAllStart := time.Now()
	output, err := cmd.CombinedOutput()
	catchAllDuration := time.Since(catchAllStart)
	newTests := parseNewTests(output, *baseTest, tests)

	if len(newTests) > 0 {
		fmt.Printf("Found %d new tests:\n", len(newTests))
		for _, t := range newTests {
			fmt.Printf("  %s\n", t)
		}

		allTests := append(tests, newTests...)
		sort.Strings(allTests)
		writeList(*listFile, allTests)

		if *checkUnchanged {
			fmt.Printf("ERROR: test list file %s was modified because new tests were found. Please commit the updated file.\n", *listFile)
			if err != nil {
				fmt.Printf("Catch-all also failed:\n%s\n", string(output))
			}
			os.Exit(1)
		}
	}

	if err != nil && len(newTests) > 0 {
		fmt.Printf("Catch-all failed (new tests likely failed):\n%s\n", string(output))
		success = false
	} else if err != nil && len(newTests) == 0 {
		fmt.Printf("Catch-all failed with no new tests found:\n%s\n", string(output))
		success = false
	} else {
		fmt.Printf("Catch-all succeeded in %v\n", catchAllDuration.Round(time.Millisecond))
	}

	if !success {
		os.Exit(1)
	}
}

func readList(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Fatalf("error reading list file: %v", err)
	}
	defer f.Close()
	var list []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t != "" && !strings.HasPrefix(t, "#") {
			list = append(list, t)
		}
	}
	return list
}

func writeList(path string, list []string) {
	var buf bytes.Buffer
	for _, t := range list {
		buf.WriteString(t + "\n")
	}
	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		log.Fatalf("error writing list file: %v", err)
	}
}

func parseNewTests(output []byte, baseTest string, knownTests []string) []string {
	prefix := "=== RUN   " + baseTest + "/"
	var newTests []string

	known := make(map[string]bool)
	for _, t := range knownTests {
		known[t] = true
	}

	seen := make(map[string]bool)
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			t := strings.TrimSpace(strings.TrimPrefix(line, prefix))
			if !known[t] && !seen[t] {
				seen[t] = true
				newTests = append(newTests, t)
			}
		}
	}
	return newTests
}
