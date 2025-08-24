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

	autogenDir := filepath.Join(repoRoot, "dev", "autogen")
	var autogenFiles []string

	// Walk the autogen directory and collect all the files
	if err := filepath.WalkDir(autogenDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		autogenFiles = append(autogenFiles, path)
		return nil
	}); err != nil {
		return fmt.Errorf("failed to walk autogen directory: %w", err)
	}

	sort.Strings(autogenFiles)

	log.Info("found autogen commands", "autogenFiles", autogenFiles)

	// Find the base commit
	// Look for a similar commit in the branch
	allRefs, err := repo.ListAllRefs(ctx)
	if err != nil {
		return fmt.Errorf("failed to list all refs: %w", err)
	}
	log.Info("all refs", "refs", allRefs)

	baseCommit, commits, err := repo.GetBranchCommits(ctx)
	if err != nil {
		return fmt.Errorf("failed to get base commit: %w", err)
	}

	createdFixup := false
	for _, file := range autogenFiles {
		// Run the autogen command
		log.Info("running autogen command", "file", file)
		cmd := exec.CommandContext(ctx, file)
		cmd.Dir = repoRoot
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to run autogen command %s: %w", file, err)
		}
		log.Info("autogen command output", "command", file, "output", string(out))

		// Get a list of changed files
		changedFiles, err := repo.GetChangedFiles(ctx)
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
		if err := repo.StageChanges(ctx, changedFiles...); err != nil {
			return fmt.Errorf("failed to stage changes: %w", err)
		}

		commitOptions := git.CommitOptions{
			Message: "autogen: " + filepath.Base(file),
		}

		var match *git.LogEntry
		for _, commit := range commits {
			log.Info("commit", "commit", commit)
			if commit.Message == commitOptions.Message {
				log.Info("found matching commit", "commit", commit)
				if match != nil {
					return fmt.Errorf("found multiple matching commits")
				}
				match = commit
			}
		}
		if match == nil {
			// Commit the change
			log.Info("committing changes")
			if err := repo.Commit(ctx, commitOptions); err != nil {
				return fmt.Errorf("failed to commit changes: %w", err)
			}
		} else {
			log.Info("found matching commit", "commit", match)
			if err := repo.CommitFixup(ctx, match.Hash); err != nil {
				return fmt.Errorf("failed to commit fixup: %w", err)
			}

			createdFixup = true
		}
	}

	if createdFixup {
		if err := repo.Autosquash(ctx, baseCommit.Hash); err != nil {
			return fmt.Errorf("failed to commit fixup: %w", err)
		}
	}

	return nil
}
