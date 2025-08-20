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

//  go run . --remote upstream  --branch upstream/master --version-file version/VERSION  --source ${REPO} -v=2 --yes=true
import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	// Support GCP auth
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	log := klog.FromContext(ctx)

	remote := ""
	flag.StringVar(&remote, "remote", remote, "name of remote")
	sourceCheckout := ""
	flag.StringVar(&sourceCheckout, "source", sourceCheckout, "location of source code checkout")
	branch := ""
	flag.StringVar(&branch, "branch", branch, "branch for creation of tags")
	versionFile := ""
	flag.StringVar(&versionFile, "version-file", versionFile, "path to version marker file")
	pushOptions := ""
	flag.StringVar(&pushOptions, "push-options", pushOptions, "when pushing tags to git, include these options")
	yes := false
	flag.BoolVar(&yes, "yes", yes, "actually do it")
	addVPrefix := false
	flag.BoolVar(&addVPrefix, "add-v-prefix", addVPrefix, "prefix tag with 'v'")
	klog.InitFlags(nil)
	flag.Parse()

	if branch == "" {
		return fmt.Errorf("must specify --branch")
	}
	branchRegex := regexp.MustCompile(`^release_\d+\.\d+$`)
	if !branchRegex.MatchString(branch) {
		return fmt.Errorf("branch %q does not match expected format 'release_X.Y'", branch)
	}

	if remote == "" {
		return fmt.Errorf("must specify --remote")
	}

	if sourceCheckout == "" {
		return fmt.Errorf("must specify --source")
	}

	if versionFile == "" {
		return fmt.Errorf("must specify --version-file")
	}
	fullVersionFilePath := filepath.Join(sourceCheckout, versionFile)
	if _, err := os.Stat(fullVersionFilePath); err != nil {
		return fmt.Errorf("could not access version file %q: %w", fullVersionFilePath, err)
	}

	homeDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}
	sourceCheckout = strings.ReplaceAll(sourceCheckout, "~", homeDir)

	tmpDir, err := os.MkdirTemp("", "worktree")
	if err != nil {
		return fmt.Errorf("error creating temp directory: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			klog.Warningf("error removing temp directory: %v", err)
		}
	}()

	originalWorktree := gitWorktree{dir: sourceCheckout}

	currentSHA, err := originalWorktree.RevParse(ctx, branch)
	if err != nil {
		return err
	}
	currentVersionBytes, err := originalWorktree.ReadFile(ctx, currentSHA, versionFile)
	if err != nil {
		return err
	}
	lastVersionBytes, err := originalWorktree.ReadFile(ctx, currentSHA+"^", versionFile)
	if err != nil {
		return err
	}

	currentVersion := string(currentVersionBytes)
	currentVersion = strings.TrimSpace(currentVersion)
	lastVersion := string(lastVersionBytes)
	lastVersion = strings.TrimSpace(lastVersion)

	if currentVersion == lastVersion {
		log.Info("no version change; nothing to do", "currentVersion", currentVersion, "lastVersion", lastVersion)
		return nil
	}

	// TODO: Check tag does not exist?
	// originalRefs, err := originalWorktree.ShowRefs(ctx)
	// if err != nil {
	// 	return fmt.Errorf("error showing refs: %w", err)
	// }

	message := "Release " + currentVersion
	tagName := currentVersion
	if addVPrefix {
		tagName = "v" + tagName
	}
	log.Info("creating tag", "tag", tagName, "sha", currentSHA)
	if !yes {
		fmt.Fprintf(os.Stdout, "Would create tag %v pointing to %v\n", tagName, currentSHA)
		fmt.Fprintf(os.Stdout, "Pass --yes to actually do the tagging\n")
		return nil
	}
	if err := originalWorktree.CreateTag(ctx, currentVersion, message, currentSHA); err != nil {
		return err
	}

	if _, err := originalWorktree.runGithubCLI(ctx, "auth", "setup-git"); err != nil {
		return fmt.Errorf("error configuring github auth: %w", err)
	}

	if err := originalWorktree.PushTag(ctx, remote, tagName, pushOptions); err != nil {
		return err
	}

	return nil
}

func runCommand(cmd *exec.Cmd) (*cmdResults, error) {
	results := &cmdResults{}

	if cmd.Stdout == nil {
		cmd.Stdout = &results.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = &results.Stderr
	}
	klog.Infof("running command %q", strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		klog.Warningf("error running command: %v", err)
		klog.Warningf("stdout: %v", results.Stdout.String())
		klog.Warningf("stderr: %v", results.Stderr.String())
		return results, fmt.Errorf("running command %q: %w", strings.Join(cmd.Args, " "), err)
	}
	klog.Infof("command OK")
	klog.Infof("stdout: %v", results.Stdout.String())
	klog.Infof("stderr: %v", results.Stderr.String())
	return results, nil
}

type gitWorktree struct {
	dir string
}

type cmdResults struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}

func (w *gitWorktree) runGit(ctx context.Context, args ...string) (*cmdResults, error) {
	var gitArgs []string
	gitArgs = append(gitArgs, "-C", w.dir)
	gitArgs = append(gitArgs, args...)

	cmd := exec.CommandContext(ctx, "git", gitArgs...)
	return runCommand(cmd)
}

func (w *gitWorktree) runGithubCLI(ctx context.Context, args ...string) (*cmdResults, error) {
	var ghArgs []string
	ghArgs = append(ghArgs, args...)

	cmd := exec.CommandContext(ctx, "gh", ghArgs...)
	cmd.Dir = w.dir
	return runCommand(cmd)
}

func (w *gitWorktree) ShowRefs(ctx context.Context) (map[string]string, error) {
	results, err := w.runGit(ctx, "show-ref")
	if err != nil {
		return nil, err
	}

	refs := make(map[string]string)

	r := bufio.NewReader(&results.Stdout)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading stdout from git show-ref: %w", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			return nil, fmt.Errorf("cannot parse line %q from git show-ref", line)
		}

		refs[tokens[1]] = tokens[0]
	}
	return refs, nil
}

func (w *gitWorktree) ReadFile(ctx context.Context, revision string, path string) ([]byte, error) {
	results, err := w.runGit(ctx, "show", revision+":"+path)
	if err != nil {
		return nil, err
	}

	return results.Stdout.Bytes(), nil
}

func (w *gitWorktree) RevParse(ctx context.Context, revision string) (string, error) {
	results, err := w.runGit(ctx, "rev-parse", revision)
	if err != nil {
		return "", err
	}

	sha := results.Stdout.String()
	sha = strings.TrimSpace(sha)

	// Simple check that this looks like a SHA
	if len(sha) != 40 {
		return "", fmt.Errorf("unexpected SHA, got %q (length %v)", sha, len(sha))
	}
	return sha, nil
}

func (w *gitWorktree) CreateTag(ctx context.Context, tagName string, message string, sha string) error {
	log := klog.FromContext(ctx)

	// TODO: Sign tags?
	results, err := w.runGit(ctx, "tag", "-a", "-m", message, tagName, sha)
	if err != nil {
		return err
	}

	log.V(2).Info("created tag", "results", results)
	return nil
}

func (w *gitWorktree) PushTag(ctx context.Context, remoteName string, tagName string, options string) error {
	log := klog.FromContext(ctx)

	cmd := []string{"push"}
	if options != "" {
		cmd = append(cmd, "-o", options)
	}
	cmd = append(cmd, remoteName, tagName)
	results, err := w.runGit(ctx, cmd...)
	if err != nil {
		return err
	}

	log.V(2).Info("created tag", "results", results)
	return nil
}
