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
	"testing"
)

func TestExecCommandRunnerDryRun(t *testing.T) {
	var stdout, stderr bytes.Buffer
	runner := newExecCommandRunner(&stdout, &stderr, true)

	err := runner.run(context.Background(), "nonexistent-binary-for-dry-run", "storage", "buckets", "create", "gs://my-bucket")

	if err != nil {
		t.Fatalf("unexpected error on dry run: %v", err)
	}
	wantOut := "[dry-run] would run: nonexistent-binary-for-dry-run storage buckets create gs://my-bucket\n"
	if stdout.String() != wantOut {
		t.Fatalf("got stdout %q, want %q", stdout.String(), wantOut)
	}
}

func TestExecCommandRunnerOutput(t *testing.T) {
	var stdout, stderr bytes.Buffer
	runner := newExecCommandRunner(&stdout, &stderr, true)

	out, err := runner.output(context.Background(), "echo", "hello world")
	if err != nil {
		t.Fatalf("unexpected error running output: %v", err)
	}
	if string(out) != "hello world\n" {
		t.Fatalf("got output %q, want %q", string(out), "hello world\n")
	}
	// Output should print nothing to stdout/stderr
	if stdout.Len() != 0 {
		t.Fatalf("expected stdout to be empty, got %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected stderr to be empty, got %q", stderr.String())
	}
}
