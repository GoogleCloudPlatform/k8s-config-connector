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

func checkoutBranch(branch Branch, workDir string, out *strings.Builder) {
	log.Printf("COMMAND: git checkout %s", branch.Local)
	checkout := exec.Command("git", "checkout", branch.Local)
	checkout.Dir = workDir
	checkout.Stdout = out
	if err := checkout.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("BRANCH CHECKOUT: %q\n", formatCommandOutput(out.String()))
	out.Reset()
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

func gitAdd(workDir string, out *strings.Builder, files ...string) {
	params := ""
	first := true
	for _, file := range files {
		if first {
			first = false
		} else {
			params += " "
		}
		params += file
	}
	log.Printf("COMMAND: git add %s", params)
	args := []string{"add"}
	args = append(args, files...)
	gitadd := exec.Command("git", args...)
	gitadd.Dir = workDir
	gitadd.Stdout = out
	gitadd.Stderr = out
	if err := gitadd.Run(); err != nil {
		log.Printf("GIT add error: %q\n", formatCommandOutput(out.String()))
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("BRANCH ADD: %q\n", formatCommandOutput(out.String()))
	out.Reset()
}

func gitCommit(workDir string, out *strings.Builder, msg string) {
	log.Printf("COMMAND: git commit -m %q", msg)
	gitcommit := exec.Command("git", "commit", "-m", fmt.Sprintf("conductor: %q", msg))
	gitcommit.Dir = workDir
	gitcommit.Stdout = out
	gitcommit.Stderr = out
	if err := gitcommit.Run(); err != nil {
		log.Printf("GIT commit error: %q\n", formatCommandOutput(out.String()))
		out.Reset()
		log.Fatal(err)
	}
	log.Printf("BRANCH COMMIT: %q\n", formatCommandOutput(out.String()))
	out.Reset()
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
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Printf("Error creating logging dir %s, :%v", logDir, err)
			return noOp
		}
	}

	var out *os.File
	var err error
	logFile := filepath.Join(logDir, "out.log")
	if out, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		log.Printf("Error opening logging file %s, :%v", logFile, err)
		return noOp
	}
	log.SetOutput(out)

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
	Name    string            // Name of the command for logging
	Cmd     string            // The command to run
	Args    []string          // Command arguments
	WorkDir string            // Working directory
	Stdin   io.Reader         // Optional stdin
	Env     map[string]string // Optional environment variables
	Timeout time.Duration     // Timeout duration (default 5m)
}

// Helper function to execute a command with timing and logging
func executeCommand(cfg CommandConfig, out *strings.Builder, stderr *strings.Builder) error {
	if cfg.Timeout == 0 {
		cfg.Timeout = 5 * time.Minute
	}

	log.Printf("Starting command step: %s", cfg.Name)
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
	defer cancel()

	cmd := exec.CommandContext(ctx, cfg.Cmd, cfg.Args...)
	cmd.Dir = cfg.WorkDir
	cmd.Stdout = out
	if stderr != nil {
		cmd.Stderr = stderr
	} else {
		cmd.Stderr = out
	}
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
	if err != nil {
		log.Printf("[%s] ERROR (%v): \n", cfg.Name, diff)
		log.Printf("[%s] err: %q\n", cfg.Name, err)
	} else {
		log.Printf("[%s] SUCCESS (%v): \n", cfg.Name, diff)
	}
	printCommandOutput(out.String())
	if stderr != nil && stderr.Len() > 0 {
		log.Printf("[%s] stderr output:\n", cfg.Name)
		printCommandOutput(stderr.String())
		stderr.Reset()
	}
	out.Reset()
	return err
}

func repoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	repoRoot := strings.TrimSpace(string(output))
	return repoRoot, nil
}
