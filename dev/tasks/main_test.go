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

package main

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

func TestTagExists(t *testing.T) {
	ctx := context.Background()
	tmpDir, err := os.MkdirTemp("", "git-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	runGit := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = tmpDir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v failed: %v\n%s", args, err, out)
		}
	}

	runGit("init")
	runGit("config", "user.email", "test@example.com")
	runGit("config", "user.name", "Test User")
	runGit("commit", "--allow-empty", "-m", "initial commit")

	w := &gitWorktree{dir: tmpDir}

	tagName := "v1.0.0"

	// Check before tag exists
	exists, err := w.TagExists(ctx, tagName)
	if err != nil {
		t.Errorf("TagExists failed: %v", err)
	}
	if exists {
		t.Errorf("expected tag %s to not exist", tagName)
	}

	// Create tag
	runGit("tag", tagName)

	// Check after tag exists
	exists, err = w.TagExists(ctx, tagName)
	if err != nil {
		t.Errorf("TagExists failed: %v", err)
	}
	if !exists {
		t.Errorf("expected tag %s to exist", tagName)
	}
}

func TestRevParse(t *testing.T) {
	ctx := context.Background()
	tmpDir, err := os.MkdirTemp("", "git-test-revparse")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	runGit := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = tmpDir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v failed: %v\n%s", args, err, out)
		}
	}

	runGit("init", "-b", "master")
	runGit("config", "user.email", "test@example.com")
	runGit("config", "user.name", "Test User")
	runGit("commit", "--allow-empty", "-m", "initial commit")

	w := &gitWorktree{dir: tmpDir}

	sha, err := w.RevParse(ctx, "master")
	if err != nil {
		t.Errorf("RevParse failed: %v", err)
	}
	if len(sha) != 40 {
		t.Errorf("expected 40 chars SHA, got %q", sha)
	}
}
