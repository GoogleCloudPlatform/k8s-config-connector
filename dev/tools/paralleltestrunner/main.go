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
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	listFile       = flag.String("list-file", "", "File containing the list of tests")
	j              = flag.Int("j", 0, "Number of parallel processes (0 = auto-detect)")
	baseTest       = flag.String("base-test", "", "Base test name, e.g. TestE2EScript/scenarios")
	checkUnchanged = flag.Bool("check-unchanged", false, "Fail if the list file is modified (i.e. missing tests found)")
	timeout        = flag.Duration("timeout", 5*time.Minute, "Timeout for each test")
	shardIndex     = flag.Int("shard-index", -1, "0-based shard index for partitioning tests")
	totalShards    = flag.Int("total-shards", 0, "Total number of shards for partitioning tests")
)

func main() {
	flag.Parse()
	if *listFile == "" || *baseTest == "" || len(flag.Args()) == 0 {
		log.Fatalf("Usage: paralleltestrunner -list-file=<file> -base-test=<base> <cmd> [args...]")
	}

	restoreFD := configureFileDescriptors()
	defer restoreFD()

	tests := readList(*listFile)
	if len(tests) == 0 {
		tests = discoverTestsDynamically(*listFile, *baseTest, flag.Args())
	}

	if *totalShards > 1 && *shardIndex >= 0 {
		var sharded []string
		for idx, t := range tests {
			if idx%*totalShards == *shardIndex {
				sharded = append(sharded, t)
			}
		}
		log.Printf("paralleltestrunner: sharding tests (shard %d of %d): selected %d / %d tests", *shardIndex+1, *totalShards, len(sharded), len(tests))
		tests = sharded
	}

	numJobs := calculateOptimalJobs(*j, len(tests))

	success := true
	var wg sync.WaitGroup
	sem := make(chan struct{}, numJobs)
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

		updatedAllTests := append(tests, newTests...)
		sort.Strings(updatedAllTests)
		writeList(*listFile, updatedAllTests)

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
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning list file %s: %v", path, err)
	}
	return list
}

func configureFileDescriptors() func() {
	var originalRLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &originalRLimit); err != nil {
		return func() {}
	}

	// Dynamic Core Ratio Formula: Every batch of 4 CPU cores requires 1024 file descriptors
	requiredFDs := uint64((runtime.NumCPU()+3)/4) * 1024

	// Apply safety ceiling cap of 12,288 (12K FDs) as a guardrail against resource exhaustion
	maxFDTarget := uint64(12288)
	targetFDs := min(originalRLimit.Max, min(requiredFDs, maxFDTarget))

	// Elevate soft limit if lower than required (capped by system hard limit)
	if originalRLimit.Cur < targetFDs {
		newRLimit := originalRLimit
		newRLimit.Cur = targetFDs
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &newRLimit); err == nil {
			log.Printf("paralleltestrunner: dynamically elevated ulimit -n soft limit from %d to %d (hard limit: %d)", originalRLimit.Cur, newRLimit.Cur, originalRLimit.Max)
			return func() {
				_ = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &originalRLimit)
				log.Printf("paralleltestrunner: restored original ulimit -n soft limit to %d", originalRLimit.Cur)
			}
		}
	}

	return func() {}
}

func calculateOptimalJobs(userJ int, numTests int) int {
	// Short-circuit: no tests need to run
	if numTests < 1 {
		log.Printf("paralleltestrunner: 0 tests to run")
		return 0
	}

	// If user explicitly passed -j > 0, honor user preference (capped by numTests)
	if userJ > 0 {
		jobs := min(userJ, numTests)
		log.Printf("paralleltestrunner: using user-specified parallelism -j=%d (selected %d worker jobs for %d tests)", userJ, jobs, numTests)
		return jobs
	}

	// Check current file descriptor ulimit
	var rLimit syscall.Rlimit
	fdLimit := uint64(1024)
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err == nil {
		fdLimit = rLimit.Cur
	}

	// Each envtest process consumes ~250 file descriptors (etcd + kube-apiserver + controller-manager).
	safeFDJobs := int(fdLimit / 250)

	// Target a 2x CPU multiplier for I/O-bound envtest workers, bounded between 4 and 8 jobs:
	// - Floor of 4: Ensures 2-vCPU CI runners (e.g. GitHub Actions) utilize 4 parallel workers efficiently.
	// - 2x Multiplier: Maximizes throughput for I/O-heavy envtest setups (etcd DB operations, IPC, and HTTP requests).
	// - Ceiling of 8: Avoids CPU thread scheduling contention and webhook startup 10s timeouts observed at >8 workers.
	targetJobs := max(4, min(runtime.NumCPU()*2, 8))
	jobs := min(targetJobs, safeFDJobs)

	// Cap by total test count if total tests < jobs
	jobs = min(jobs, numTests)
	log.Printf("paralleltestrunner: auto-scaled parallelism -j=%d (CPU cores: %d, FD limit: %d, total tests: %d)", jobs, runtime.NumCPU(), fdLimit, numTests)
	return jobs
}

func writeList(path string, list []string) {
	var buf bytes.Buffer
	for _, t := range list {
		buf.WriteString(t)
		buf.WriteByte('\n')
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
	if err := scanner.Err(); err != nil {
		log.Printf("error scanning output in parseNewTests: %v", err)
	}
	return newTests
}

func discoverTestsDynamically(listFile string, baseTest string, args []string) []string {
	log.Printf("paralleltestrunner: list file %q missing or empty, discovering tests dynamically...", listFile)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "RUN_TESTS="+baseTest)
	cmd.Env = append(cmd.Env, "SKIP_ALL=1")
	output, _ := cmd.CombinedOutput()
	tests := parseNewTests(output, baseTest, nil)
	log.Printf("paralleltestrunner: dynamically discovered %d tests", len(tests))
	return tests
}
