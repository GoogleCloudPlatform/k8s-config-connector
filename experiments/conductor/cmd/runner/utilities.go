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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	GenerativeCommandRetryBackoff = 60 * time.Second
)

var commandMap = map[int64]string{
	cmdEnableGCPAPIs:       "enablegcpapis",
	cmdCreateScriptYaml:    "createscriptyaml",
	cmdCaptureHttpLog:      "capturehttplog",
	cmdGenerateMockGo:      "generatemockgo",
	cmdAddServiceRoundTrip: "addserviceroundtrip",
	cmdAddProtoMakefile:    "addprotomakefile",
	cmdRunMockTests:        "runmocktests",
	cmdGenerateTypes:       "generatetypes",
	cmdGenerateCRD:         "generatecrd",
	cmdGenerateFuzzer:      "generatefuzzer",
}

type exitBash func()

func startBash() (io.WriteCloser, io.ReadCloser, exitBash, error) {
	cmd := exec.Command("bash")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		errBuffer := make([]byte, 1000)
		_, err = stderr.Read(errBuffer)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("BASH ERR %s", string(errBuffer))
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	exit := func() {
		log.Printf("COMMAND: exit")
		if _, err = stdin.Write([]byte("exit\n")); err != nil {
			log.Fatal(err)
		}
		err := cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("BASH DONE")
	}
	return stdin, stdout, exit, err
}

func cdRepoBranchDirBash(opts *RunnerOptions, subdir string, stdin io.WriteCloser, stdout io.ReadCloser) string {
	dir := opts.branchRepoDir
	if subdir != "" {
		dir = filepath.Join(dir, subdir)
	}
	log.Printf("COMMAND: cd %s and echo", dir)
	if _, err := stdin.Write([]byte(fmt.Sprintf("cd %s && echo done\n", dir))); err != nil {
		log.Fatal(err)
	}
	outBuffer := make([]byte, 1000)
	done := false
	var msg string
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("CD OUT %s", msg)
	return msg
}

func checkoutBranch(ctx context.Context, branch Branch, workDir string) {
	checkout := exec.CommandContext(ctx, "git", "checkout", branch.Local)
	checkout.Dir = workDir

	results, err := execCommand(checkout)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH CHECKOUT: %v\n", formatCommandOutput(results.Stdout))
}

type ExecResults struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func execCommand(cmd *exec.Cmd) (ExecResults, error) {
	var stdout strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.Printf("COMMAND: %s", cmd.String())
	err := cmd.Run()
	if err != nil {
		// var exitErr *exec.ExitError
		// if errors.As(err, &exitErr) {
		// 	return ExecResults{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: exitErr.ExitCode()}, nil
		// }

		log.Printf("error running command %v: %v", cmd.String(), err)
		log.Printf("stdout: %s", stdout.String())
		log.Printf("stderr: %s", stderr.String())
		return ExecResults{}, err
	}
	return ExecResults{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: cmd.ProcessState.ExitCode()}, nil
}

func writeTemplateToFile(branch Branch, filePath string, template string) {
	tmp := strings.ReplaceAll(template, "<TICK>", "`")
	tmp = strings.ReplaceAll(tmp, "<GCLOUD_COMMAND>", branch.Command)
	tmp = strings.ReplaceAll(tmp, "<GROUP>", branch.Group)
	tmp = strings.ReplaceAll(tmp, "<SERVICE>", branch.Group)
	tmp = strings.ReplaceAll(tmp, "<HTTP_HOST>", branch.HostName)
	tmp = strings.ReplaceAll(tmp, "<PROTO_SERVICE>", branch.ProtoSvc)
	tmp = strings.ReplaceAll(tmp, "<PROTO_MESSAGE>", branch.ProtoMsg)
	tmp = strings.ReplaceAll(tmp, "<PROTO_PACKAGE>", branch.Package)
	tmp = strings.ReplaceAll(tmp, "<CRD_GROUP>", fmt.Sprintf("%s.cnrm.cloud.google.com", branch.Group))
	tmp = strings.ReplaceAll(tmp, "<CRD_VERSION>", "v1alpha1")
	tmp = strings.ReplaceAll(tmp, "<CRD_KIND>", branch.Kind)
	tmp = strings.ReplaceAll(tmp, "<PROTO_RESOURCE>", branch.Proto)

	contents := strings.ReplaceAll(tmp, "<RESOURCE>", strings.ToLower(branch.Resource))
	log.Printf("TEMPLATE %s %s", filePath, contents)

	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		log.Printf("COMMAND: cleaning up old %s", filePath)
		err = os.Remove(filePath)
		if err != nil {
			log.Printf("Attempt to clean up %s failed with %v", filePath, err)
		}
	}
	log.Printf("COMMAND: writing new %s", filePath)
	if err := os.WriteFile(filePath, []byte(contents), 0644); err != nil {
		log.Fatal(err)
	}
}

func gitAdd(ctx context.Context, workDir string, files ...string) {
	args := []string{"git", "add"}
	for _, file := range files {
		args = append(args, file)
	}
	gitadd := exec.CommandContext(ctx, args[0], args[1:]...)
	gitadd.Dir = workDir

	results, err := execCommand(gitadd)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH ADD: %v\n", formatCommandOutput(results.Stdout))
}

func gitCommit(ctx context.Context, workDir string, msg string) {
	log.Printf("COMMAND: git commit -m %q", msg)
	gitcommit := exec.CommandContext(ctx, "git", "commit", "-m", fmt.Sprintf("conductor: %q", msg))
	gitcommit.Dir = workDir

	results, err := execCommand(gitcommit)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH COMMIT: %v\n", formatCommandOutput(results.Stdout))
}

func gitFileHasChange(workDir string, filePath string) bool {
	log.Printf("COMMAND: git diff -- %s", filePath)
	args := []string{"diff", "--", filePath}
	gitdiff := exec.Command("git", args...)
	gitdiff.Dir = workDir
	var out bytes.Buffer
	gitdiff.Stdout = &out
	if err := gitdiff.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			// Exit code 1 means no changes, not an error
			return false
		}
		log.Printf("Git diff on file %s/%s error: %q\n", workDir, filePath, err)
		return false
	}
	return len(strings.TrimSpace(out.String())) > 0
}

type closer func()

func setLoggingWriter(opts *RunnerOptions, branch Branch) closer {
	// Initially force out to stdout in case we hit an error we don't
	// want to pollute a different runs logs with our logs.
	// TODO: Return a log object so we can run in parrellel.
	log.SetOutput(os.Stdout)
	if opts.loggingDir == "" {
		log.Println("Logging dir not set")
		return noOp
	}
	logDir := filepath.Join(opts.loggingDir, branch.Name)
	log.Printf("Logging dir: %s", logDir)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Printf("Error creating logging dir %s, :%v", logDir, err)
			return noOp
		}
	}

	var out *os.File
	var err error
	logFileName := "output.log"
	commandNumber := opts.command
	commandMessage, ok := commandMap[commandNumber]
	if ok {
		logFileName = fmt.Sprintf("%v_%v.log", commandNumber, commandMessage)
	}
	logFile := filepath.Join(logDir, logFileName)
	if out, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		log.Printf("Error opening logging file %s, :%v", logFile, err)
		return noOp
	}
	log.Printf("Logging to %s", logFile)
	log.SetOutput(io.MultiWriter(os.Stderr, out))

	/*
		var errF *os.File
		errFile := filepath.Join(logDir, "err.log")
		if _, err := os.Stat(errFile); os.IsNotExist(err) {
			errF, err = os.Create(errFile)
			if err != nil {
				return os.Stdout, os.Stderr
			}
		} else {
			errF, err = os.OpenFile(errFile, os.O_APPEND, 0755)
			if err != nil {
				return os.Stdout, os.Stderr
			}
		}
	*/

	return func() {
		// Initially force out to stdout in case we hit an error we don't
		// want to pollute a different runs logs with our logs.
		log.SetOutput(os.Stdout)
		if err := out.Close(); err != nil {
			log.Printf("Failed to clode logging file %s, :%v", logFile, err)
		}
	}
}

func noOp() {
}

func printCommandOutput(output string) {
	// Replace escaped newlines and tabs with their actual characters
	formatted := strings.ReplaceAll(output, "\\n", "\n")
	formatted = strings.ReplaceAll(formatted, "\\t", "\t")

	// Split the string into lines and format each line
	lines := strings.Split(formatted, "\n")
	for _, line := range lines {
		log.Printf("  > %s\n", line)
	}
}

func formatCommandOutput(output string) string {
	// Replace escaped newlines and tabs with their actual characters
	formatted := strings.ReplaceAll(output, "\\n", "\n")
	formatted = strings.ReplaceAll(formatted, "\\t", "\t")
	return formatted
}

// CommandConfig holds the configuration for executing a command
type CommandConfig struct {
	Name         string            // Name of the command for logging
	Cmd          string            // The command to run
	Args         []string          // Command arguments
	WorkDir      string            // Working directory
	Stdin        io.Reader         // Optional stdin
	Env          map[string]string // Optional environment variables
	Timeout      time.Duration     // Timeout duration (default 5m)
	MaxRetries   int               // Maximum number of retries allowed for this command
	RetryBackoff time.Duration     // Time to wait between retries (default 1s)
}

// Helper function to execute a command with timing and logging
func executeCommand(opts *RunnerOptions, cfg CommandConfig) (string, string, error) {
	if cfg.Timeout == 0 {
		if opts != nil && opts.timeout != 0 {
			cfg.Timeout = opts.timeout
		} else {
			cfg.Timeout = 5 * time.Minute
		}
	}

	if cfg.RetryBackoff == 0 {
		cfg.RetryBackoff = time.Second
	}

	maxRetries := opts.defaultRetries
	// If MaxRetries is set in config, cap it at that
	if cfg.MaxRetries > 0 && cfg.MaxRetries < maxRetries {
		maxRetries = cfg.MaxRetries
	}

	var lastErr error
	var output, errOutput string
	retryCount := 1

	for {
		log.Printf("Starting command step: %s (attempt %d/%d)", cfg.Name, retryCount, maxRetries+1)
		log.Printf("[%s] working directory: %s", cfg.Name, cfg.WorkDir)
		log.Printf("[%s] command: %s %s", cfg.Name, cfg.Cmd, strings.Join(cfg.Args, " "))
		if cfg.Stdin != nil {
			log.Printf("[%s] stdin: %s", cfg.Name, cfg.Stdin)
		}
		if len(cfg.Env) > 0 {
			log.Printf("[%s] environment:", cfg.Name)
			for k, v := range cfg.Env {
				log.Printf("  %s=%s", k, v)
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		cmd := exec.CommandContext(ctx, cfg.Cmd, cfg.Args...)
		cmd.Dir = cfg.WorkDir

		var outBuf, errBuf strings.Builder
		cmd.Stdout = &outBuf
		cmd.Stderr = &errBuf

		if cfg.Stdin != nil {
			cmd.Stdin = cfg.Stdin
		}

		// Set up environment variables
		if len(cfg.Env) > 0 {
			cmd.Env = os.Environ() // Start with current environment
			for k, v := range cfg.Env {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
			}
		}

		start := time.Now()
		err := cmd.Run()
		stop := time.Now()
		diff := stop.Sub(start)
		cancel()

		output = outBuf.String()
		errOutput = errBuf.String()

		if err == nil {
			log.Printf("[%s] SUCCESS (%v)", cfg.Name, diff)
			printCommandOutput(output)
			if errBuf.Len() > 0 {
				log.Printf("[%s] stderr output:", cfg.Name)
				printCommandOutput(errOutput)
			}
			return output, errOutput, nil
		}

		lastErr = err
		log.Printf("[%s] ERROR (%v): %v", cfg.Name, diff, err)
		printCommandOutput(output)
		if errBuf.Len() > 0 {
			log.Printf("[%s] stderr output:", cfg.Name)
			printCommandOutput(errOutput)
		}

		retryCount++
		if maxRetries >= 0 && retryCount > maxRetries {
			log.Printf("[%s] Exceeded maximum retries (%d), giving up", cfg.Name, maxRetries)
			break
		}

		log.Printf("[%s] Retrying in %v...", cfg.Name, cfg.RetryBackoff)
		time.Sleep(cfg.RetryBackoff)
	}

	return output, errOutput, fmt.Errorf("command failed after %d attempts: %w", maxRetries, lastErr)
}
