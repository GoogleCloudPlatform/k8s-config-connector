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

package ha

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// commandRunner executes external CLI commands (such as `gcloud`) or logs them during dry runs.
type commandRunner interface {
	run(ctx context.Context, name string, args ...string) error
}

// commandOutputRunner executes external CLI commands and returns their captured stdout.
type commandOutputRunner interface {
	output(ctx context.Context, name string, args ...string) (string, error)
}

type execCommandRunner struct {
	stdout io.Writer
	stderr io.Writer
	dryRun bool
}

func newExecCommandRunner(stdout, stderr io.Writer, dryRun bool) *execCommandRunner {
	if stdout == nil {
		stdout = io.Discard
	}

	if stderr == nil {
		stderr = io.Discard
	}

	return &execCommandRunner{
		stdout: stdout,
		stderr: stderr,
		dryRun: dryRun,
	}
}

func (r *execCommandRunner) run(ctx context.Context, name string, args ...string) error {
	cmdLine := formatCommandLine(name, args)
	if r.dryRun {
		fmt.Fprintf(r.stdout, "[dry-run] would run: %s\n", cmdLine)
		return nil
	}

	fmt.Fprintf(r.stdout, "Running: %s\n", cmdLine)

	var errBuf bytes.Buffer
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = r.stdout
	cmd.Stderr = io.MultiWriter(r.stderr, &errBuf)

	if err := cmd.Run(); err != nil {
		stderrMsg := strings.TrimSpace(errBuf.String())
		if stderrMsg != "" {
			return fmt.Errorf("running `%s`: %w: %s", cmdLine, err, stderrMsg)
		}

		return fmt.Errorf("running `%s`: %w", cmdLine, err)
	}

	return nil
}

func (r *execCommandRunner) output(ctx context.Context, name string, args ...string) (string, error) {
	cmdLine := formatCommandLine(name, args)
	var outBuf, errBuf bytes.Buffer
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		stderrMsg := strings.TrimSpace(errBuf.String())
		if stderrMsg != "" {
			return "", fmt.Errorf("running `%s`: %w: %s", cmdLine, err, stderrMsg)
		}

		return "", fmt.Errorf("running `%s`: %w", cmdLine, err)
	}

	return outBuf.String(), nil
}

func formatCommandLine(name string, args []string) string {
	if len(args) == 0 {
		return name
	}

	return name + " " + strings.Join(args, " ")
}
