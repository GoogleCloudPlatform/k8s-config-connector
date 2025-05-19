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
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	GenerativeCommandRetryBackoff = 60 * time.Second
)

var commandMap = map[int64]string{
	cmdEnableGCPAPIs:                "enablegcpapis",
	cmdCreateScriptYaml:             "createscriptyaml",
	cmdCaptureHttpLog:               "capturehttplog",
	cmdGenerateMockGo:               "generatemockgo",
	cmdAddServiceRoundTrip:          "addserviceroundtrip",
	cmdAddProtoMakefile:             "addprotomakefile",
	cmdCaptureMockOutput:            "capturemockoutput",
	cmdRunAndFixMockTests:           "runandfixmocktests",
	cmdGenerateTypes:                "generatetypes",
	cmdGenerateCRD:                  "generatecrd",
	cmdGenerateFuzzer:               "generatefuzzer",
	cmdBuildProto:                   "buildproto",
	cmdAdjustTypes:                  "adjusttypes",
	cmdGenerateMapper:               "generatemapper",
	cmdRunAndFixFuzzTests:           "runandfixfuzztests",
	cmdRunAndFixAPIChecks:           "runandfixapichecks",
	cmdControllerClient:             "controllerclient",
	cmdGenerateController:           "generatecontroller",
	cmdBuildAndFixController:        "buildandfixcontroller",
	cmdCreateIdentity:               "createidentity",
	cmdControllerCreateTest:         "controllercreatetest",
	cmdCaptureGoldenRealGCPOutput:   "capturegoldenrealgcpoutput",
	cmdRunAndFixGoldenRealGCPOutput: "runandfixgoldenrealgcpoutput",
	cmdCaptureGoldenMockOutput:      "capturegoldenmockoutput",
	cmdRunAndFixGoldenMockOutput:    "runandfixgoldenmockoutput",
	cmdMoveExistingTest:             "moveexistingtest",
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

func checkAndCommitLocalChanges(ctx context.Context, branch Branch, workDir string) {
	checkUncommittedChanges(ctx, branch, workDir, true)
}

func checkUncommittedChanges(ctx context.Context, branch Branch, workDir string, commit bool) {
	log.Print("Checking uncommitted changes: git status --porcelain")
	statusCmd := exec.CommandContext(ctx, "git", "status", "--porcelain")
	statusCmd.Dir = workDir

	results, err := execCommand(statusCmd)
	if err != nil {
		log.Fatal(err)
	}
	if results.Stdout != "" {
		currentBranchCmd := exec.CommandContext(ctx, "git", "rev-parse", "--abbrev-ref", "HEAD")
		currentBranchCmd.Dir = workDir

		currentBranchResult, err := execCommand(currentBranchCmd)
		if err != nil {
			log.Fatal(err)
		}
		if commit {
			if err := gitAdd(ctx, workDir, "."); err != nil {
				log.Fatalf("Tried to add changes at branch %q before committing but failed: %v", strings.TrimSuffix(currentBranchResult.Stdout, "\n"), err)
			}
			if err := gitCommit(ctx, workDir, "[Warning] Unfinished changes detected by conductor"); err != nil {
				log.Fatalf("Tried to commit changes at branch %q before checking out to branch %q but failed: %v", strings.TrimSuffix(currentBranchResult.Stdout, "\n"), branch.Local, err)
			}
			log.Printf("Successfully committed changes at branch %q before checking out to branch %q:\n%s\n", strings.TrimSuffix(currentBranchResult.Stdout, "\n"), branch.Local, results.Stdout)
		} else {
			log.Fatalf("Found uncommitted changes at branch %q before checking out to branch %q:\n%s\n", strings.TrimSuffix(currentBranchResult.Stdout, "\n"), branch.Local, results.Stdout)
		}
	}

	log.Print("The branches are ready for running the command.")
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

func gitAdd(ctx context.Context, workDir string, files ...string) error {
	args := []string{"git", "add"}
	for _, file := range files {
		args = append(args, file)
	}
	gitadd := exec.CommandContext(ctx, args[0], args[1:]...)
	gitadd.Dir = workDir

	results, err := execCommand(gitadd)
	if err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}
	log.Printf("BRANCH ADD: %v\n", formatCommandOutput(results.Stdout))
	return nil
}

func gitCommit(ctx context.Context, workDir string, msg string) error {
	authorName := "kcc-conductor-bot"
	authorEmail := "kcc-conductor-bot@google.com"
	authorFlag := fmt.Sprintf("%s <%s>", authorName, authorEmail)

	log.Printf("COMMAND: git commit -m %q --author=%q", msg, authorFlag)
	gitcommit := exec.CommandContext(ctx, "git", "commit", "-m", msg, "--author", authorFlag)
	gitcommit.Dir = workDir

	results, err := execCommand(gitcommit)
	if err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}
	log.Printf("BRANCH COMMIT: %v\n", formatCommandOutput(results.Stdout))
	return nil
}

func gitStatusCheck(workDir string, filePath string) (bool, error) {
	log.Printf("COMMAND: git status -- %s", filePath)
	args := []string{"status", "-s", filePath}
	gitstatus := exec.Command("git", args...)
	gitstatus.Dir = workDir
	var out bytes.Buffer
	gitstatus.Stdout = &out
	if err := gitstatus.Run(); err != nil {
		log.Printf("Git status on file %s/%s error: %q\n", workDir, filePath, err)
		return false, err
	}
	log.Printf("Git status on file %s/%s: %s", workDir, filePath, out.String())
	return len(strings.TrimSpace(out.String())) > 0, nil
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

func gitRevert(ctx context.Context, workDir string, filePath string) error {
	log.Printf("COMMAND: git checkout -- %s", filePath)
	args := []string{"checkout", "--", filePath}
	gitcheckout := exec.CommandContext(ctx, "git", args...)
	gitcheckout.Dir = workDir

	results, err := execCommand(gitcheckout)
	if err != nil {
		log.Printf("Git checkout on file %s/%s error: %v", workDir, filePath, err)
		return err
	}
	if results.Stdout != "" {
		log.Printf("Git checkout output: %s", formatCommandOutput(results.Stdout))
	}
	return nil
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
	if opts.verbose {
		// In verbose mode, write to both stderr and the log file
		log.SetOutput(io.MultiWriter(os.Stderr, out))
	} else {
		// In non-verbose mode, write only to the log file
		log.SetOutput(out)
	}

	return func() {
		// Initially force out to stdout in case we hit an error we don't
		// want to pollute a different runs logs with our logs.
		log.SetOutput(os.Stdout)
		if err := out.Close(); err != nil {
			log.Printf("Failed to close logging file %s, :%v", logFile, err)
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
	MaxAttempts  int               // Maximum number of attempts allowed for this command to run
	RetryBackoff time.Duration     // Time to wait between retries (default 1s)
}

// Helper function to execute a command with timing and logging
func executeCommand(opts *RunnerOptions, cfg CommandConfig) (ExecResults, error) {
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
	if cfg.MaxAttempts > 0 && cfg.MaxAttempts < maxRetries {
		maxRetries = cfg.MaxAttempts
	}

	var lastErr error
	var output, errOutput string
	var exitCode int
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

		if opts.verbose {
			cmd.Stdout = io.MultiWriter(os.Stdout, cmd.Stdout)
			cmd.Stderr = io.MultiWriter(os.Stderr, cmd.Stderr)
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
		cancel()

		output = outBuf.String()
		errOutput = errBuf.String()
		exitCode = cmd.ProcessState.ExitCode()

		if err == nil {
			log.Printf("[%s] SUCCESS (%v)", cfg.Name, diff)
			printCommandOutput(output)
			if errBuf.Len() > 0 {
				log.Printf("[%s] stderr output:", cfg.Name)
				printCommandOutput(errOutput)
			}
			return ExecResults{Stdout: output, Stderr: errOutput, ExitCode: exitCode}, nil
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

	return ExecResults{Stdout: output, Stderr: errOutput, ExitCode: exitCode}, fmt.Errorf("command failed after %d attempts: %w", maxRetries, lastErr)
}

func runLinters(opts *RunnerOptions) error {
	log.Printf("Running linters")

	cfg := CommandConfig{
		Name:        "Go Fmt",
		Cmd:         "go",
		Args:        []string{"fmt", "./..."},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}
	op, err := executeCommand(opts, cfg)
	if err != nil {
		return fmt.Errorf("Error running go fmt: %w", err)
	}
	if op.ExitCode != 0 {
		return fmt.Errorf("linting failed with exit code %d", op.ExitCode)
	}

	return nil
}

func checkMakeReadyPR(opts *RunnerOptions) error {
	log.Printf("Running make ready-pr verification")

	cfg := CommandConfig{
		Name:        "Make Ready-PR",
		Cmd:         "make",
		Args:        []string{"ready-pr"},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}

	if _, err := executeCommand(opts, cfg); err != nil {
		return fmt.Errorf("make ready-pr failed: %w", err)
	}

	return nil
}

func stageChanges(ctx context.Context, opts *RunnerOptions, paths []string, commitMsg string) (bool, error) {
	// First check for changes in specified paths
	changesCommitted := false
	var changedFiles []string
	for _, path := range paths {
		log.Printf("Checking for changes in %s", path)
		changed, err := gitStatusCheck(opts.branchRepoDir, path)
		if err != nil {
			return changesCommitted, fmt.Errorf("git status failed for %s: %w", path, err)
		}
		if changed {
			changedFiles = append(changedFiles, path)
		}
	}

	// If we have changes in specified paths, add and commit them
	if len(changedFiles) > 0 {
		for _, file := range changedFiles {
			if err := gitAdd(ctx, opts.branchRepoDir, file); err != nil {
				log.Printf("[ERROR GIT ADD]failed to stage %s: %v", file, err)
				// best effort to continue
			}
		}
		if err := gitCommit(ctx, opts.branchRepoDir, commitMsg); err != nil {
			log.Printf("failed to commit changes: %v", err)
		}
		changesCommitted = true
	}

	// Check for any other uncommitted changes
	changed, err := gitStatusCheck(opts.branchRepoDir, ".")
	if err != nil {
		return changesCommitted, fmt.Errorf("git status failed: %w", err)
	}

	if changed {
		log.Printf("WARNING: Found unexpected changes. Trying to stage and commit them")
		if err := gitAdd(ctx, opts.branchRepoDir, "."); err != nil {
			return changesCommitted, fmt.Errorf("failed to stage extra files: %w", err)
		}
		if err := gitCommit(ctx, opts.branchRepoDir, fmt.Sprintf("%s. WARN: extra files found, committing them. ", commitMsg)); err != nil {
			return changesCommitted, fmt.Errorf("failed to commit extra files: %w", err)
		}
		changesCommitted = true
		log.Printf("WARNING: Committed unexpected changes")
	}

	return changesCommitted, nil
}

type BranchProcessorFn func(ctx context.Context, opts *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error)
type SkipProcessorOnMsgFn func(msg string) bool

type BranchProcessor struct {
	CommitMsgTemplate  string
	Fn                 BranchProcessorFn
	AttemptsOnNoChange int // Number of attempts to run the processor if no changes are detected
	VerifyFn           BranchProcessorFn
	VerifyAttempts     int
	AttemptsOnChanges  int // Number of attempts to run the processor if changes are detected
	SkipProcessorOnMsg SkipProcessorOnMsgFn
	CommitOptional     bool // Whether a commit is optional or required; default to be required
}

func (b *BranchProcessor) CommitMsg(branch Branch) string {
	s := b.CommitMsgTemplate
	s = strings.ReplaceAll(s, "{{command}}", branch.Command)
	s = strings.ReplaceAll(s, "{{group}}", branch.Group)
	s = strings.ReplaceAll(s, "{{resource}}", branch.Resource)
	s = strings.ReplaceAll(s, "{{kind}}", branch.Kind)
	return s
}

// runBranchFnWithRetries runs the processor multiple times until changes are detected and commits them
func runBranchFnWithRetriesAndCommit(ctx context.Context, opts *RunnerOptions, branch Branch, description string, processor BranchProcessor, commitMsg string, inExecResults *ExecResults) (bool, *ExecResults, error) {
	// Try running processor up to N times until we see changes in all affected paths
	var execResults *ExecResults
	changesCommitted := false
	if commitMsg == "" {
		commitMsg = processor.CommitMsg(branch)
	}
	maxAttempts := processor.AttemptsOnNoChange
	if maxAttempts == 0 {
		maxAttempts = 3
	}
	attempt := 0
	var affectedPaths []string
	lintFailure := false
	goCmdsFailure := false
	for attempt < maxAttempts {
		var err error
		attempt++
		// Process the branch
		log.Printf("Running BranchFn for branch %s: %s, processor: %s (attempt: %d/%d)", branch.Name, description, commitMsg, attempt, maxAttempts)
		affectedPaths, execResults, err = processor.Fn(ctx, opts, branch, inExecResults)
		if err != nil {
			log.Printf("failed to process branch %s on attempt %d: %v", branch.Name, attempt, err)
			continue
		}

		// Check if any files changed
		someChanged := false
		for _, path := range affectedPaths {
			changed, err := gitStatusCheck(opts.branchRepoDir, path)
			if err != nil {
				log.Printf("failed to check status of %s: %v", path, err)
				continue
			}
			if changed {
				someChanged = true
				break
			}
		}

		// If commitMsg is not set, we don't need to commit the changes,
		// we assume there are no changes to commit.
		// If CommitOptional is explicitly set to true and there is no change,
		// we assume there are no changes to commit.
		if commitMsg == "" || processor.CommitOptional && !someChanged {
			return changesCommitted, execResults, nil
		}

		lintFailure = false
		goCmdsFailure = false
		// Run go mod tidy
		if err := runGoCmds(opts, affectedPaths); err != nil {
			log.Printf("go cmds failed for branch %s: %v", branch.Name, err)
			goCmdsFailure = true
		}
		// Run basic linting
		if err := runLinters(opts); err != nil {
			log.Printf("linting failed for branch %s: %v", branch.Name, err)
			lintFailure = true
		}
		// Exit loop if we saw changes or hit max attempts
		if someChanged || attempt >= maxAttempts {
			break
		}

		log.Printf("No changes detected in attempt %d, retrying...", attempt)
		time.Sleep(10 * time.Second)
	}

	//if err := checkMakeReadyPR(opts); err != nil {
	//	log.Printf("verification failed for branch %s: %v", branch.Name, err)
	//}

	// Stage and commit changes
	if lintFailure {
		commitMsg = fmt.Sprintf("%s. WARN: linting failed.", commitMsg)
	}
	if goCmdsFailure {
		commitMsg = fmt.Sprintf("%s. WARN: go cmds failed.", commitMsg)
	}
	var err error
	changesCommitted, err = stageChanges(ctx, opts, affectedPaths, commitMsg)
	if err != nil {
		return changesCommitted, execResults, fmt.Errorf("failed to commit changes for branch %s: %w", branch.Name, err)
	}

	return changesCommitted, execResults, nil
}

// runProcessOnBranch handles the processing of a single branch
func runProcessOnBranch(ctx context.Context, opts *RunnerOptions, branch Branch, description string, processor BranchProcessor) error {
	log.Printf("Processing branch %s: %s, processor: %s", branch.Name, description, processor.CommitMsg(branch))

	if processor.VerifyFn == nil || processor.VerifyAttempts == 0 {
		attempts := processor.AttemptsOnChanges
		if attempts == 0 {
			attempts = 1
		}

		// Run atmost attempts times if we see changes. This is for cases like make ready-pr
		// where we want to run it multiple times if we see changes.
		for i := 0; i < attempts; i++ {
			changesCommitted, _, err := runBranchFnWithRetriesAndCommit(ctx, opts, branch, description, processor, "", nil)
			if err != nil {
				return err
			}
			if !changesCommitted {
				break
			}
			if attempts > 1 {
				log.Printf("Changes committed. Running again: %d/%d", i+1, attempts)
			}
		}
		return nil
	}

	// If processor has a VerifyFn and max attempts, run verification loop
	for i := 0; i < processor.VerifyAttempts; i++ {
		var err error
		paths, execResults, err := processor.VerifyFn(ctx, opts, branch, nil)
		if execResults == nil {
			log.Printf("execResults is nil for branch %s, commit message: %s", branch.Name, processor.CommitMsg(branch))
			continue
		}
		if execResults.ExitCode == 0 {
			log.Printf("Verification attempt %d succeeded for branch %s, commit message: %s", i+1, branch.Name, processor.CommitMsg(branch))
			break
		}
		if err != nil {
			log.Printf("Verification attempt %d Error: %v for branch %s, commit message: %s", i+1, err, branch.Name, processor.CommitMsg(branch))
		} else {
			log.Printf("Verification attempt %d failed for branch %s, commit message: %s", i+1, branch.Name, processor.CommitMsg(branch))
		}
		if len(paths) != 0 {
			// Revert the changes for the failed verification
			for _, path := range paths {
				if err := gitRevert(ctx, opts.branchRepoDir, path); err != nil {
					log.Printf("Failed to revert changes for branch %s: %v", branch.Name, err)
					continue
				}
			}
		}
		log.Printf("Verification attempt %d failed for branch %s, commit message: %s", i+1, branch.Name, processor.CommitMsg(branch))
		commitMsg := fmt.Sprintf("Autofix attempt %d. %s", i+1, processor.CommitMsg(branch))
		_, _, err = runBranchFnWithRetriesAndCommit(ctx, opts, branch, description, processor, commitMsg, execResults)
		if err != nil {
			log.Printf("Failed to run branch %s: %v", branch.Name, err)
			continue
		}
	}

	return nil
}

// processBranches applies the given processors to each branch
func processBranches(ctx context.Context, opts *RunnerOptions, skipOnGitMessage *regexp.Regexp, branches []Branch, description string, processors []BranchProcessor) {
	// If processors filter is provided, filter the processors
	if opts.processors != "" {
		selectedProcessors := strings.Split(opts.processors, ",")
		// Create a map for faster lookup
		processorSet := make(map[string]bool)
		for _, p := range selectedProcessors {
			processorSet[strings.TrimSpace(p)] = true
		}

		// Filter processors
		var filteredProcessors []BranchProcessor
		for _, processor := range processors {
			// Get the function name using reflection
			fnName := getFunctionName(processor.Fn)
			if processorSet[fnName] {
				filteredProcessors = append(filteredProcessors, processor)
				log.Printf("Including processor: %s", fnName)
			} else {
				log.Printf("Skipping processor: %s (not in --processors list)", fnName)
			}
		}

		// If no processors matched, log a warning
		if len(filteredProcessors) == 0 {
			log.Printf("WARNING: No processors matched the filter: %s", opts.processors)
			log.Printf("Available processors: %s", getAllProcessorNames(processors))
			return
		}

		processors = filteredProcessors
	}

	for _, branch := range branches {
		err := processBranch(ctx, opts, branch, skipOnGitMessage, description, processors)
		if err != nil {
			log.Printf("Error processing branch %s: %v", branch.Name, err)
			// Not breaking because we want to process the next branch
		}
	}
}

func processBranch(ctx context.Context, opts *RunnerOptions, branch Branch, skipOnGitMessage *regexp.Regexp, description string, processors []BranchProcessor) error {
	// Skip if branch should be skipped
	if branch.Skip {
		log.Printf("Skipping branch %s: marked as Skip", branch.Name)
		return nil
	}

	close := setLoggingWriter(opts, branch)
	defer close()

	checkAndCommitLocalChanges(ctx, branch, opts.branchRepoDir)
	checkoutBranch(ctx, branch, opts.branchRepoDir)

	// Run git diff command
	message, found := getLatestGitMessage(opts.branchRepoDir, opts, branch)
	if skipOnGitMessage != nil {
		if found && skipOnGitMessage.MatchString(message) {
			log.Printf("Skipping branch %s: git commit %s", branch.Name, message)
			return nil
		}
	}

	for _, processor := range processors {
		if processor.SkipProcessorOnMsg != nil && found && processor.SkipProcessorOnMsg(message) {
			log.Printf("Skipping branch(processor) %s(%s): git commit %s", branch.Name, processor.CommitMsgTemplate, message)
			continue
		}
		err := runProcessOnBranch(ctx, opts, branch, description, processor)
		if err != nil {
			return err
		}
	}
	return nil
}

// getFunctionName returns the name of a function
func getFunctionName(fn BranchProcessorFn) string {
	// Get the function value using reflection
	fnValue := reflect.ValueOf(fn)

	// Get the function pointer
	fnPointer := fnValue.Pointer()

	// Get the full function name including package path
	fullName := runtime.FuncForPC(fnPointer).Name()

	// Extract just the function name without package path
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

// getAllProcessorNames returns a comma-separated list of all processor function names
func getAllProcessorNames(processors []BranchProcessor) string {
	var names []string
	for _, processor := range processors {
		names = append(names, getFunctionName(processor.Fn))
	}
	return strings.Join(names, ", ")
}

func runGoCmds(opts *RunnerOptions, affectedPaths []string) error {
	log.Printf("Running go mod tidy")
	errStrings := []string{}

	// Filter for paths containing Go files
	goFilePaths := filterGoFilePaths(opts.branchRepoDir, affectedPaths)

	// Only run goimports if we found Go files
	if len(goFilePaths) > 0 {
		for _, path := range goFilePaths {
			cfg := CommandConfig{
				Name:        "Go Imports",
				Cmd:         "goimports",
				Args:        []string{"-w", path},
				WorkDir:     opts.branchRepoDir,
				MaxAttempts: 1,
			}
			op, err := executeCommand(opts, cfg)
			if err != nil {
				errStrings = append(errStrings, fmt.Sprintf("Error running goimports for %s: %v", path, err))
			}
			if op.ExitCode != 0 {
				errStrings = append(errStrings, fmt.Sprintf("goimports failed with exit code %d for %s", op.ExitCode, path))
			}
		}
	} else {
		log.Printf("No Go files found in affected paths, skipping goimports")
	}

	// Always run go mod tidy regardless of file types
	cfg := CommandConfig{
		Name:        "Go Mod Tidy",
		Cmd:         "go",
		Args:        []string{"mod", "tidy"},
		WorkDir:     opts.branchRepoDir,
		MaxAttempts: 1,
	}
	op, err := executeCommand(opts, cfg)
	if err != nil {
		errStrings = append(errStrings, fmt.Sprintf("Error running go mod tidy: %v", err))
	}
	if op.ExitCode != 0 {
		errStrings = append(errStrings, fmt.Sprintf("go mod tidy failed with exit code %d", op.ExitCode))
	}

	if len(errStrings) > 0 {
		log.Printf("Error running go cmds: %s", strings.Join(errStrings, "\n"))
		return fmt.Errorf("Error running go cmds: %s", strings.Join(errStrings, "\n"))
	}

	return nil
}

// hasGoFiles checks if the given path (file or directory) contains any Go files
// If path is a file, it checks if it's a Go file
// If path is a directory, it walks the directory to find any Go files
func hasGoFiles(basePath, relativePath string) bool {
	fullPath := filepath.Join(basePath, relativePath)
	info, err := os.Stat(fullPath)
	if err != nil {
		log.Printf("Warning: Could not stat path %s: %v", relativePath, err)
		return false
	}

	// If it's a file, check if it's a Go file
	if !info.IsDir() {
		return strings.HasSuffix(strings.ToLower(relativePath), ".go")
	}

	// If it's a directory, walk it to find Go files
	foundGoFile := false
	err = filepath.Walk(fullPath, func(walkPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(walkPath), ".go") {
			foundGoFile = true
			return filepath.SkipAll // Stop walking once we find a Go file
		}
		return nil
	})

	if err != nil {
		log.Printf("Warning: Error walking directory %s: %v", relativePath, err)
		return false
	}

	return foundGoFile
}

// filterGoFilePaths returns a slice containing only paths that have Go files
// Handles both absolute and relative paths
func filterGoFilePaths(repoRootPath string, inputPaths []string) []string {
	var goFilePaths []string
	for _, path := range inputPaths {
		// Check if this is an absolute path within our repository
		var checkPath string
		if filepath.IsAbs(path) && strings.HasPrefix(path, repoRootPath) {
			// It's an absolute path within our repo, convert to relative for checking
			relPath, err := filepath.Rel(repoRootPath, path)
			if err != nil {
				log.Printf("Warning: Could not convert absolute path %s to relative: %v", path, err)
				continue
			}
			checkPath = relPath
		} else if filepath.IsAbs(path) {
			// It's an absolute path outside our repo, skip it
			log.Printf("Warning: Path %s is outside repository, skipping", path)
			continue
		} else {
			// It's already a relative path
			checkPath = path
		}

		// Now check if this path contains Go files
		if hasGoFiles(repoRootPath, checkPath) {
			goFilePaths = append(goFilePaths, path) // Keep the original path format
		}
	}
	return goFilePaths
}

const COPYRIGHT_HEADER string = `# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

`

type branchAscending []Branch

func (v branchAscending) Len() int {
	return len(v)
}

func (v branchAscending) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v branchAscending) Less(i, j int) bool {
	if v[i].Group != v[j].Group {
		return v[i].Group < v[j].Group
	}
	if v[i].Kind != v[j].Kind {
		return v[i].Kind < v[j].Kind
	}
	return v[i].Name < v[j].Name
}

func writeBranchesStableOrder(branches Branches, fileName string) {
	sort.Sort(branchAscending(branches.Branches))
	data := []byte(COPYRIGHT_HEADER)
	yamlData, err := yaml.Marshal(branches)
	if err != nil {
		log.Fatal(err)
	}
	data = append(data, yamlData...)
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
