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

// commandRunner executes external CLI commands (such as `gcloud`) or prints them during dry runs.
type commandRunner interface {
	run(ctx context.Context, name string, args ...string) error
}

type execCommandRunner struct {
	stdout io.Writer
	stderr io.Writer
	dryRun bool
}

func newExecCommandRunner(stdout, stderr io.Writer, dryRun bool) commandRunner {
	return &execCommandRunner{
		stdout: stdout,
		stderr: stderr,
		dryRun: dryRun,
	}
}

func (r *execCommandRunner) run(ctx context.Context, name string, args ...string) error {
	if r.stdout != nil {
		fmt.Fprintf(r.stdout, "$ %s %s\n", name, strings.Join(args, " "))
	}
	if r.dryRun {
		return nil
	}

	var errBuf bytes.Buffer
	cmd := exec.CommandContext(ctx, name, args...)
	if r.stdout != nil {
		cmd.Stdout = r.stdout
	}
	if r.stderr != nil {
		cmd.Stderr = io.MultiWriter(r.stderr, &errBuf)
	} else {
		cmd.Stderr = &errBuf
	}

	if err := cmd.Run(); err != nil {
		stderrMsg := strings.TrimSpace(errBuf.String())
		if stderrMsg != "" {
			return fmt.Errorf("running `%s %s`: %w: %s", name, strings.Join(args, " "), err, stderrMsg)
		}
		return fmt.Errorf("running `%s %s`: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

// gkeClusterInfo holds the GCP project, location, and cluster name parsed from a GKE context.
type gkeClusterInfo struct {
	contextName string
	identity    string
	project     string
	location    string
	clusterName string
}

func (g gkeClusterInfo) mclGSAEmail() string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", mclServiceAccountName, g.project)
}

func (g gkeClusterInfo) syncerGSAEmail() string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", syncerServiceAccountName, g.project)
}

// parseGKEContext parses a standard gcloud kubeconfig context name of the form
// "gke_<project>_<location>_<cluster-name>" and derives its DNS-1123 cluster identity.
func parseGKEContext(contextName string) (gkeClusterInfo, error) {
	trimmed := strings.TrimSpace(contextName)
	parts := strings.Split(trimmed, "_")
	if len(parts) != 4 || parts[0] != "gke" || parts[1] == "" || parts[2] == "" || parts[3] == "" {
		return gkeClusterInfo{}, fmt.Errorf(
			"context %q does not match expected GKE context format \"gke_<project>_<location>_<cluster>\"; run `gcloud container clusters get-credentials <cluster> --location=<location> --project=<project>`",
			contextName,
		)
	}
	identity, err := sanitizeClusterIdentity(trimmed)
	if err != nil {
		return gkeClusterInfo{}, err
	}
	return gkeClusterInfo{
		contextName: trimmed,
		identity:    identity,
		project:     parts[1],
		location:    parts[2],
		clusterName: parts[3],
	}, nil
}
