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
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/autogen/pkg/git"
	"k8s.io/klog/v2"
)

func main() {
	err := run(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	log := klog.FromContext(ctx)

	repoRoot := ""
	flag.StringVar(&repoRoot, "repo-root", repoRoot, "The root of the repository")
	flag.Parse()

	repo, err := git.NewRepo(ctx, repoRoot)
	if err != nil {
		return fmt.Errorf("failed to create repo: %w", err)
	}

	autofixDir := filepath.Join(repoRoot, "dev", "autofix")
	var autofixFiles []string

	// Walk the autofix directory and collect all the files
	if err := filepath.WalkDir(autofixDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		autofixFiles = append(autofixFiles, path)
		return nil
	}); err != nil {
		return fmt.Errorf("failed to walk autofix directory: %w", err)
	}

	sort.Strings(autofixFiles)

	log.Info("found autofix commands", "autofixFiles", autofixFiles)

	_, commits, err := repo.GetBranchCommits(ctx)
	if err != nil {
		return fmt.Errorf("failed to get base commit: %w", err)
	}

	worktreeDir, err := os.MkdirTemp("", "autofix")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(worktreeDir)

	worktree, err := repo.CreateWorktree(ctx, worktreeDir)
	if err != nil {
		return fmt.Errorf("failed to create worktree: %w", err)
	}
	defer func() {
		if err := repo.DeleteWorktree(ctx, worktreeDir); err != nil {
			log.Error(err, "failed to delete worktree")
		}
	}()

	var fixups []string

	for i, commit := range commits {

		log.Info("commit", "commit", commit)
		if i == 0 {
			if err := worktree.Reset(ctx, commit.Hash); err != nil {
				return fmt.Errorf("failed to checkout commit %s: %w", commit.Hash, err)
			}
		} else {
			if err := worktree.CherryPick(ctx, commit.Hash); err != nil {
				return fmt.Errorf("failed to cherry-pick commit %s: %w", commit.Hash, err)
			}
		}

		for _, file := range autofixFiles {
			// Run the autogen command
			log.Info("running autofix command", "file", file)
			cmd := exec.CommandContext(ctx, file)
			cmd.Dir = worktree.Root
			out, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("failed to run autogen command %s: %w", file, err)
			}
			log.Info("autofix command output", "command", file, "output", string(out))
		}

		// Get a list of changed files
		changedFiles, err := worktree.GetChangedFiles(ctx)
		if err != nil {
			return fmt.Errorf("failed to get changed files: %w", err)
		}
		if len(changedFiles) == 0 {
			log.Info("no changes detected, skipping commit")
			continue
		}
		log.Info("changed files", "changedFiles", changedFiles)

		// Stage the changes
		log.Info("staging changes")
		if err := worktree.StageChanges(ctx, changedFiles...); err != nil {
			return fmt.Errorf("failed to stage changes: %w", err)
		}

		if err := worktree.CommitFixup(ctx, commit.Hash); err != nil {
			return fmt.Errorf("failed to commit fixup: %w", err)
		}

		commitHash, err := worktree.RevParse(ctx, "HEAD")
		if err != nil {
			return fmt.Errorf("failed to get commit: %w", err)
		}
		log.Info("committed fixup", "commit", commitHash)
		fixups = append(fixups, commitHash)
	}

	if len(fixups) > 0 {
		log.Info("fixups", "fixups", fixups)
		for _, fixup := range fixups {
			if err := repo.CherryPick(ctx, fixup); err != nil {
				return fmt.Errorf("failed to cherry-pick fixup %s: %w", fixup, err)
			}
		}
	}

	return nil
}
