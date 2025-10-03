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
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"k8s.io/klog"
	"sigs.k8s.io/yaml"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type GitOptions struct {
	BaseRepo      string
	CopyRepo      string
	CopyRemote    string
	SSHKeyPath    string
	CommitMessage string
}

func (o *GitOptions) DefaultGitBranch(r *ResourceMetadata) string {
	return "resource-" + r.Service + "-" + r.Proto
}

type ResourceMetadata struct {
	Service   string `yaml:"service"`
	Package   string `yaml:"package"`
	Proto     string `yaml:"proto"`
	Kind      string `yaml:"kind"`
	ProtoPath string `yaml:"protopath"`
	Validated *bool  `yaml:"validated"`
}

func (r *ResourceMetadata) String() string {
	return fmt.Sprintf("%s:%s:%s", r.Service, r.Package, r.Proto)
}

func run(ctx context.Context) error {
	var o GitOptions

	klog.InitFlags(nil)

	flag.StringVar(&o.BaseRepo, "base-repo", o.BaseRepo, "the base repo where your LLM codebot lives.")
	flag.StringVar(&o.CopyRepo, "copy-repo", o.CopyRepo, "another KCC repo to allow LLM to edit files. Need to be a clean environment. ")
	flag.StringVar(&o.CopyRemote, "copy-remote", o.CopyRemote, "your local git remote name in the copy repo, default to origin")
	flag.StringVar(&o.SSHKeyPath, "ssh-key", o.SSHKeyPath, "path to your ssh key. default to $HOME/.ssh/id_rsa")
	flag.StringVar(&o.CommitMessage, "commit-msg", o.CommitMessage, "the git commit description. default to `add resource xxx`")
	flag.Parse()
	if o.BaseRepo == "" {
		o.BaseRepo = os.Getenv("REPO_ROOT")
	}

	totalResources, err := ListResource(ctx, o.BaseRepo)
	if err != nil {
		klog.Infof("list resource failed: %v", err)
		return err
	}
	klog.Infof("Total number is %d", len(totalResources))

	serviceResources := map[string][]*ResourceMetadata{}
	for _, r := range totalResources {
		if r.Service == "" {
			continue
		}
		resources, ok := serviceResources[r.Service]
		if !ok {
			resources = []*ResourceMetadata{}
		}
		resources = append(resources, r)
		serviceResources[r.Service] = resources
	}

	// TODO: LLM decides what to do with this resource

	exception, err := os.Create("exception-resource.txt")
	if err != nil {
		klog.Infof("creating exception file %s: %v", "exception-resource.txt", err)
		return err
	}
	defer exception.Close()

	for _, r := range totalResources {
		err := func() error {

			if err := CheckoutBranch(ctx, &o, r); err != nil {
				klog.Infof("checkout branch failed: %v", err)
				return err
			}

			if err := GenerateResource(ctx, &o, r); err != nil {
				klog.Infof("generate resource failed: %v", err)
				return err
			}
			if err := CommitBranch(ctx, &o, r); err != nil {
				klog.Infof("commit changes failed: %v", err)
				return err
			}

			if err := PushToRemote(ctx, &o, r); err != nil {
				klog.Infof("push to remote branch failed: %v", err)
				return err
			}
			return nil
		}()
		if err != nil {
			_, wErr := exception.WriteString("resource: " + r.Kind)
			if wErr != nil {
				klog.Infof("write exception failed 1: %s: %v", r.Kind, wErr)
				continue
			}
			_, wErr = exception.WriteString(err.Error())
			if wErr != nil {
				klog.Infof("write exception failed 2: %s: %v", r.Kind, wErr)
				continue
			}
			_, wErr = exception.WriteString("\n")
			if wErr != nil {
				klog.Infof("write exception failed 3: %s: %v", r.Kind, wErr)
				continue
			}
		}
	}
	return nil
}

func ListResource(ctx context.Context, repoRoot string) ([]*ResourceMetadata, error) {
	// python3 mockgcp/dev/tools/print-gcloud-resource-commands >> gcloud-list.txt
	// p := filepath.Join(repo_root, "gcloud-list.txt")
	p := filepath.Join(repoRoot, "proto-list-2.yaml")
	fmt.Println("proto file", p)
	fileContents, err := os.ReadFile(p)

	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	// Create an instance of the struct
	var totalResources []*ResourceMetadata

	// Unmarshal the YAML data into the struct
	resourceBytes := strings.Split(string(fileContents), "\n---\n")
	for _, b := range resourceBytes {
		var r ResourceMetadata

		err = yaml.Unmarshal([]byte(b), &r)

		existingPath1 := filepath.Join(repoRoot, "apis", r.Service, "v1beta1", strings.ToLower(r.Kind)+"_types.go")
		if _, err := os.Stat(existingPath1); err == nil {
			klog.Infof("resource %s already exists, skip", r.Kind)
			continue
		}
		existingPath2 := filepath.Join(repoRoot, "apis", r.Service, "v1beta1", strings.ToLower(r.Proto)+"_types.go")
		if _, err := os.Stat(existingPath2); err == nil {
			klog.Infof("resource %s already exists, skip", r.Kind)
			continue
		}

		totalResources = append(totalResources, &r)
	}
	if err != nil {
		klog.Infof("error unmarshaling YAML: %v\n", err)
		return nil, err
	}
	return totalResources, nil
}
func copyFile(o *GitOptions, name string) error {
	src := filepath.Join(o.BaseRepo, name)
	dest := filepath.Join(o.CopyRepo, name)

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return err
		}
	}

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func GenerateResource(ctx context.Context, o *GitOptions, r *ResourceMetadata) error {
	if o.CopyRepo == "" {
		return fmt.Errorf("to generate a new resource, you need to specify the copy-repo flag")
	}

	ignoredFiles := []string{
		".build/googleapis.pb",
	}
	for _, ignoredFile := range ignoredFiles {
		if err := copyFile(o, ignoredFile); err != nil {
			return fmt.Errorf("copy file %s: %w", ignoredFile, err)
		}
		klog.Infof("copied git ignored file: %s", ignoredFile)
	}

	args := []string{
		"run",
		"./dev/tools/controllerbuilder/main.go",
		"generate-direct-reconciler",
		"--service=" + r.Package,
		"--kind=" + r.Kind,
		"--api-version=" + r.Service + ".cnrm.cloud.google.com/v1alpha1",
		"--proto-resource=" + r.Proto,
	}
	cmd := exec.CommandContext(ctx, "go", args...)

	cmd.Dir = o.CopyRepo

	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		klog.Infof(string(output))
	}

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command failed with exit code %d: %s", exitError.ExitCode(), string(exitError.Stderr))
		}
		return fmt.Errorf("command failed: %w", err)
	}
	klog.Infof("generated code for resource: %s", r.String())
	return nil
}

func CheckoutBranch(ctx context.Context, o *GitOptions, r *ResourceMetadata) error {
	if o.CopyRepo == "" {
		return fmt.Errorf("you need to specify the copy-repo")
	}
	branchName := o.DefaultGitBranch(r)

	repo, err := git.PlainOpen(o.CopyRepo)
	if err != nil {
		return fmt.Errorf("opening repository %s: %w", o.CopyRepo, err)
	}

	clean, err := isGitClean(repo)
	if err != nil {
		return fmt.Errorf("check git repo cleanness failed: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("get worktree: %w", err)
	}

	exist, err := DoesBranchExist(repo, branchName)
	if !exist {
		baseRef := plumbing.ReferenceName("refs/heads/master")
		/*
			masterRef, err := repo.Reference(baseRef, true)
			if err != nil {
				klog.Infof("master ref: %v", err)
			}*/

		if !clean {
			klog.Infof("the copy-repo %s has uncommitted files. Please stash or commit your change first", branchName)
		}

		if err = worktree.Checkout(&git.CheckoutOptions{
			Branch: baseRef,
			Create: false,
			Force:  true,
		}); err != nil {
			return err
		}

		newBranchRefName := plumbing.ReferenceName("refs/heads/" + branchName)
		/*
			err = repo.Storer.SetReference(plumbing.NewHashReference(newBranchRefName, masterRef.Hash()))
			if err != nil {
				log.Fatal(err)
			}*/

		if err = worktree.Checkout(&git.CheckoutOptions{
			Branch: newBranchRefName,
			Create: true,
		}); err != nil {
			return err
		}
		if _, err := repo.Head(); err != nil {
			return err
		}
	} else {
		head, err := repo.Head()
		if err != nil {
			return fmt.Errorf("failed to get current HEAD: %w", err)
		}
		currentBranch := head.Name().Short()
		// only check out if not in current branch, otherwise git ignored changes like googleapis.pb could be wiped out.
		if currentBranch == branchName {
			klog.Infof("already on branch %s, skipping checkout", branchName)
			return nil
		}
		// todo: git rebase
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + branchName),
			Create: false,
			Force:  true,
		})
		if err != nil {
			return fmt.Errorf("checkout existing branch %s: %w", branchName, err)
		}

	}

	return nil
}

func CommitBranch(ctx context.Context, o *GitOptions, r *ResourceMetadata) error {
	// We should run the command against copy repo, not base repo
	repo, err := git.PlainOpen(o.CopyRepo)
	if err != nil {
		return fmt.Errorf("error opening repository: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	stagedfilePatterns := []string{
		"apis/" + r.Service + "/v1alpha1/" + "*.go",
		"config/crds/resources/*.yaml",
		"pkg/controller/direct/" + r.Service + "/*.go",
		"pkg/controller/direct/register/register.go",
		"mockgcp/" + "mock" + r.Service + "/*.go",
		"mockgcp/mock_http_roundtrip.go",
		"config/tests/samples/create/harness.go",
		"pkg/test/resourcefixture/testdata/basic/**/**",
	}
	for _, p := range stagedfilePatterns {
		if err := worktree.AddGlob(p); err != nil {
			if !strings.Contains(err.Error(), "glob pattern did not match any files") {
				return fmt.Errorf("git add %s failed: %w", p, err)
			}
		}

	}
	clean, err := isGitClean(repo)
	if err != nil {
		return fmt.Errorf("check git repo cleanness failed: %w", err)
	}
	if clean {
		klog.Infof("no unstaged or staged changes. skipping git-commit")
		return nil
	}

	// TODO: AI generate message
	msg := ""
	if o.CommitMessage == "" {
		msg = "feat: add resource " + r.Kind
		klog.Infof("`commit-msg` unset, use default `feat: add resource %s`", r.Kind)
	} else {
		msg = o.CommitMessage
	}

	commit, err := worktree.Commit(msg, &git.CommitOptions{
		Author: authorFromDefaultConfig(),
	})
	if err != nil {
		return err
	}
	klog.Infof("git commit: %s", commit)

	return nil
}
func PushToRemote(ctx context.Context, o *GitOptions, r *ResourceMetadata) error {
	remoteName := ""
	if o.CopyRemote == "" {
		klog.Infof("`remote` not specified, use `origin`")
		remoteName = "origin"
	} else {
		remoteName = o.CopyRemote
	}

	branchName := o.DefaultGitBranch(r)
	repo, err := git.PlainOpen(o.CopyRepo)
	if err != nil {
		return fmt.Errorf("opening repository %s: %w", o.CopyRepo, err)
	}
	remote, err := repo.Remote(remoteName)
	if err != nil {
		return fmt.Errorf("getting remote %s: %w", remoteName, err)
	}

	sshKeyPath := ""
	if o.SSHKeyPath == "" {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("Could not get current user: %w", err)
		}
		homeDir := usr.HomeDir
		sshKeyPath = filepath.Join(homeDir, "/.ssh/id_rsa")
		klog.Infof("`ssh-key` not specified, use %s", sshKeyPath)
	} else {
		sshKeyPath = o.SSHKeyPath
	}
	sshAuth, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
	if err != nil {
		klog.Infof("Error configuring SSH authentication: %v", err)
	}
	pushOptions := &git.PushOptions{
		RemoteName: remoteName,
		RefSpecs: []config.RefSpec{config.RefSpec(
			fmt.Sprintf("refs/heads/%s:refs/heads/%s", branchName, branchName),
		)},
		Auth: sshAuth,
	}

	err = remote.Push(pushOptions)

	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			klog.Infof("No changes to push")
		} else {
			klog.Infof("pushing to remote: %v", err)
		}
	}
	klog.Infof("Push change to remote: %s/%s", remoteName, o.DefaultGitBranch(r))
	return nil
}

func isGitClean(repo *git.Repository) (bool, error) {
	worktree, err := repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("error getting worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return false, fmt.Errorf("error getting status: %w", err)
	}
	return status.IsClean(), nil
}

func DoesBranchExist(r *git.Repository, branchName string) (bool, error) {
	branches, err := r.Branches()
	if err != nil {
		return false, fmt.Errorf("failed to get branches: %w", err)
	}

	branchExists := false
	if err = branches.ForEach(func(ref *plumbing.Reference) error {
		if strings.TrimPrefix(ref.Name().String(), "refs/heads/") == branchName {
			branchExists = true
			return nil
		}
		return nil
	}); err != nil {
		return false, fmt.Errorf("failed to iterate branches: %w", err)
	}
	return branchExists, nil
}

func authorFromDefaultConfig() *object.Signature {
	a := &object.Signature{}

	cfg, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		klog.Infof("Error loading global config: %v", err)
	}

	if cfg.User.Name == "" || cfg.User.Email == "" {
		cfg, err = config.LoadConfig(config.SystemScope)
		if err != nil {
			klog.Infof("Error loading system config: %v", err)
		}
	}

	if cfg.User.Name == "" || cfg.User.Email == "" {
		cfg, err = config.LoadConfig(config.LocalScope)
		if err != nil {
			klog.Infof("Error loading local config: %v", err)
		}
	}
	a.Name = cfg.User.Name
	a.Email = cfg.User.Email
	a.When = time.Now()
	return a
}
