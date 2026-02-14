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

package git

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"k8s.io/klog/v2"
)

type Repo struct {
	Root string
}

func NewRepo(ctx context.Context, root string) (*Repo, error) {
	if root == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get current working directory: %w", err)
		}

		repoRoot, err := GetGitRepoRoot(ctx, cwd)
		if err != nil {
			return nil, fmt.Errorf("failed to get git repo root: %w", err)
		}
		root = repoRoot
	}

	repoRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for repo-root: %w", err)
	}
	repoRoot = filepath.Clean(repoRoot)

	return &Repo{Root: repoRoot}, nil
}

// GetGitRepoRoot returns the root of the git repository.
// It runs `git rev-parse --show-toplevel` to get the root.
func GetGitRepoRoot(ctx context.Context, cwd string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--show-toplevel")
	cmd.Dir = cwd
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git repo root: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func (r *Repo) GetChangedFiles(ctx context.Context) ([]string, error) {
	gitArgs := []string{"diff", "--name-only"}
	result, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to get changed files: %w", err)
	}
	changedFiles := []string{}
	for _, line := range strings.Split(result.Stdout, "\n") {
		if line == "" {
			continue
		}
		changedFiles = append(changedFiles, line)
	}
	return changedFiles, nil
}

func (r *Repo) StageChanges(ctx context.Context, files ...string) error {
	gitArgs := append([]string{"add"}, files...)
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to stage changes: %w", err)
	}
	return nil
}

type gitResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func (r *Repo) runGit(ctx context.Context, args []string, allowExitCodes ...int) (*gitResult, error) {
	log := klog.FromContext(ctx)

	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = r.Root

	result := &gitResult{}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	log.Info("running git command", "args", strings.Join(args, " "), "cwd", r.Root)

	err := cmd.Run()
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			result.ExitCode = exitErr.ExitCode()
		}
		for _, code := range allowExitCodes {
			if result.ExitCode == code {
				return result, nil
			}
		}
		log.Info("git command failed", "args", strings.Join(args, " "), "exitCode", result.ExitCode)
		log.Info("git command stdout", "stdout", result.Stdout)
		log.Info("git command stderr", "stderr", result.Stderr)
		return result, fmt.Errorf("failed to run git command %s: %w", strings.Join(args, " "), err)
	} else {
		result.ExitCode = cmd.ProcessState.ExitCode()
	}
	return result, nil
}

type CommitOptions struct {
	Message string
}

func (r *Repo) Commit(ctx context.Context, options CommitOptions) error {
	gitArgs := []string{"commit", "-m", options.Message}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	return nil
}

func (r *Repo) CommitFixup(ctx context.Context, hash string) error {
	gitArgs := []string{"commit", "--fixup", hash}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to commit fixup: %w", err)
	}
	return nil
}

type LogEntry struct {
	Message string
	Hash    string
}

// GitLog runs the git log command and calls the callback for each log entry.
// If the callback returns false or an error, the loop will stop.
func (r *Repo) GitLog(ctx context.Context, callback func(log *LogEntry) (bool, error)) error {
	gitArgs := []string{"log", "--format=%H:%s", "-n", "100"}
	result, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to get git log: %w", err)
	}
	for _, line := range strings.Split(result.Stdout, "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		entry := &LogEntry{Hash: parts[0], Message: parts[1]}
		if keepGoing, err := callback(entry); !keepGoing || err != nil {
			return err
		}
	}
	// Easily fixed, but we don't need to do it now.
	return fmt.Errorf("reached 100 commit limit")
}

func (r *Repo) ListAllRefs(ctx context.Context) (GitRefs, error) {
	gitArgs := []string{"for-each-ref", "--format=%(objectname)%00%(refname)%00%(objecttype)"}
	result, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to get git tags: %w", err)
	}
	refs := GitRefs{}
	for _, line := range strings.Split(result.Stdout, "\n") {
		if line == "" {
			continue
		}
		// Split on null character
		parts := strings.SplitN(line, "\000", 3)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid git ref line: %q", line)
		}
		ref := &GitRef{
			Hash:    parts[0],
			RefName: parts[1],
			Type:    parts[2],
		}
		refs[ref.RefName] = ref
	}
	return refs, nil
}

type GitRef struct {
	Hash    string
	RefName string
	Type    string
}

type GitRefs map[string]*GitRef

func (r GitRefs) GetRefsForHash(hash string) []string {
	refs := []string{}
	for _, ref := range r {
		if ref.Hash == hash {
			refs = append(refs, ref.RefName)
		}
	}
	return refs
}

func (r *Repo) Autosquash(ctx context.Context, base string) error {
	// git -c sequence.editor=: rebase -i --autosquash upstream/main
	gitArgs := []string{"-c", "sequence.editor=:", "rebase", "-i", "--autosquash", base}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to autosquash: %w", err)
	}
	return nil
}

func (r *Repo) GetBranchCommits(ctx context.Context) (*LogEntry, []*LogEntry, error) {
	log := klog.FromContext(ctx)

	// Find the base commit
	// Look for a similar commit in the branch
	allRefs, err := r.ListAllRefs(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list all refs: %w", err)
	}
	// log.Info("all refs", "refs", allRefs)

	var baseCommit *LogEntry
	var commits []*LogEntry

	if err := r.GitLog(ctx, func(logEntry *LogEntry) (bool, error) {
		// log.Info("checking commit", "entry", logEntry)
		refs := allRefs.GetRefsForHash(logEntry.Hash)
		for _, ref := range refs {
			log.Info("checking ref", "ref", ref)
			if ref == "refs/heads/master" || ref == "refs/heads/main" || ref == "refs/remotes/upstream/master" || ref == "refs/remotes/upstream/main" || ref == "refs/remotes/upstream/HEAD" {
				log.Info("found base ref", "ref", ref)
				baseCommit = logEntry
				return false, nil
			}
		}
		commits = append(commits, logEntry)
		return true, nil
	}); err != nil {
		return nil, nil, fmt.Errorf("failed to find similar commit: %w", err)
	}

	if baseCommit == nil {
		return nil, nil, fmt.Errorf("no base commit found")
	}

	slices.Reverse(commits)

	log.Info("base commit", "commit", baseCommit)
	log.Info("commits", "commits", commits)

	return baseCommit, commits, nil
}

func (r *Repo) CreateWorktree(ctx context.Context, dir string) (*Repo, error) {
	gitArgs := []string{"worktree", "add", dir}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to create worktree: %w", err)
	}
	return &Repo{Root: dir}, nil
}

func (r *Repo) DeleteWorktree(ctx context.Context, dir string) error {
	gitArgs := []string{"worktree", "remove", dir}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to delete worktree: %w", err)
	}
	return nil
}

func (r *Repo) Reset(ctx context.Context, hash string) error {
	gitArgs := []string{"reset", hash}
	_, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to reset: %w", err)
	}
	return nil
}

func (r *Repo) CherryPick(ctx context.Context, hash string) error {
	log := klog.FromContext(ctx)

	gitArgs := []string{"cherry-pick", hash}
	result, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return fmt.Errorf("failed to cherry-pick: %w", err)
	}
	log.Info("cherry-picked", "result", result)
	return nil
}

func (r *Repo) RevParse(ctx context.Context, hash string) (string, error) {
	gitArgs := []string{"rev-parse", hash}
	result, err := r.runGit(ctx, gitArgs)
	if err != nil {
		return "", fmt.Errorf("failed to rev-parse: %w", err)
	}
	return strings.TrimSpace(result.Stdout), nil
}
