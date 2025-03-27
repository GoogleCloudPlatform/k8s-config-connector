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
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"

	"github.com/spf13/pflag"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	SourceDir string
	DryRun    bool
}

func (o *Options) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.SourceDir = root
	o.DryRun = false
	return nil
}

func (o *Options) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.SourceDir, "src", o.SourceDir, "path to source code")
	flags.BoolVar(&o.DryRun, "dryrun", o.DryRun, "if true, print commands that would be run without executing them")
}

func run(ctx context.Context) error {
	var options Options
	if err := options.InitDefaults(); err != nil {
		return err
	}
	options.BindFlags(pflag.CommandLine)
	pflag.Parse()

	srcDir := options.SourceDir
	if srcDir == "" {
		return fmt.Errorf("src flag is required")
	}

	changedFiles, errors := findChangedFiles(ctx, srcDir, options.DryRun)
	sort.Strings(changedFiles)

	if err := validateChangedFiles(srcDir, changedFiles); err != nil {
		return err
	}

	if len(errors) > 0 {
		errMsg := "errors encountered during processing:\n"
		for _, err := range errors {
			errMsg += fmt.Sprintf("  - %v\n", err)
		}
		return fmt.Errorf(errMsg)
	}

	return nil
}

func findChangedFiles(ctx context.Context, srcDir string, dryRun bool) ([]string, []error) {
	changedFiles := []string{}
	errors := []error{}

	filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			errors = append(errors, err)
			return nil
		}

		hasDiff, err := checkFile(ctx, srcDir, relPath, dryRun)
		if err != nil {
			errors = append(errors, err)
			changedFiles = append(changedFiles, relPath)
			return nil // continue to next file
		}
		if hasDiff {
			changedFiles = append(changedFiles, relPath)
		}
		return nil
	})

	return changedFiles, errors
}

// checkFile returns true if the file has changed, false otherwise.
// It returns an error if there is an error during the process
func checkFile(ctx context.Context, srcDir, relPath string, dryRun bool) (bool, error) {
	fullPath := filepath.Join(srcDir, relPath)
	originalContent, err := os.ReadFile(fullPath)
	if err != nil {
		return false, fmt.Errorf("reading file %q: %w", relPath, err)
	}

	markers := []string{"+generated:"}
	annotations, err := annotations.FindFileAnnotations(originalContent, markers)
	if err != nil {
		return false, err
	}

	// skip files that don't have any annotations
	if len(annotations) == 0 {
		return false, nil
	}

	for _, annotation := range annotations {
		// fmt.Printf("found annotation in %q: %+v\n", relPath, annotation)

		var args []string
		switch annotation.Key {
		case "+generated:mapper":
			service := strings.Join(annotation.Attributes["proto.service"], ",")
			apiVersion := strings.Join(annotation.Attributes["krm.group"], ",") + "/" + strings.Join(annotation.Attributes["krm.version"], ",")
			args = []string{
				"generate-mapper",
				"--service", service,
				"--api-version", apiVersion,
			}
		case "+generated:types":
			service := strings.Join(annotation.Attributes["proto.service"], ",")
			apiVersion := strings.Join(annotation.Attributes["krm.group"], ",") + "/" + strings.Join(annotation.Attributes["krm.version"], ",")
			args = []string{
				"generate-types",
				"--service", service,
				"--api-version", apiVersion,
				"--skip-scaffold-files",
			}
			for _, resource := range annotation.Attributes["resource"] {
				args = append(args, "--resource", resource)
			}
		}

		if err := runControllerbuilderCommand(srcDir, args, dryRun); err != nil {
			return false, err
		}
	}

	// gofmt the file
	if !dryRun {
		cmd := exec.Command(filepath.Join(srcDir, "dev", "tasks", "fix-gofmt"))
		cmd.Dir = srcDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return false, fmt.Errorf("running fix-gofmt: %w", err)
		}
	}

	// Check if the file changed
	newContent, err := os.ReadFile(fullPath)
	if err != nil {
		return false, fmt.Errorf("reading updated file %q for validation: %w", relPath, err)
	}

	return contentChanged(originalContent, newContent), nil
}

func runControllerbuilderCommand(srcDir string, args []string, dryRun bool) error {
	if len(args) == 0 {
		return nil
	}

	if dryRun {
		fmt.Printf("dryrun: %v\n", strings.Join(args, " "))
		return nil
	}

	fmt.Printf("Running command: %v\n", strings.Join(args, " "))
	cmd := exec.Command("go", append([]string{"run", filepath.Join(srcDir, "dev", "tools", "controllerbuilder")}, args...)...)
	cmd.Dir = srcDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running command %v: %w", strings.Join(args, " "), err)
	}

	return nil
}

// contentChanged compares two byte slices, ignoring copyright lines
func contentChanged(original, new []byte) bool {
	originalLines := strings.Split(string(original), "\n")
	newLines := strings.Split(string(new), "\n")

	if len(originalLines) != len(newLines) {
		return true
	}

	for i := range originalLines {
		// Skip copyright lines
		if strings.Contains(originalLines[i], "Copyright") && strings.Contains(originalLines[i], "Google LLC") {
			continue
		}

		if originalLines[i] != newLines[i] {
			return true
		}
	}

	return false
}

func validateChangedFiles(srcDir string, changedFiles []string) error {
	got := strings.Join(changedFiles, "\n")
	goldenFilePath := filepath.Join(srcDir, "tests", "apichecks", "testdata", "exceptions", "fileschanged.txt")

	want, err := os.ReadFile(goldenFilePath)
	if err != nil {
		return fmt.Errorf("reading testdata/fileschanged.txt: %w", err)
	}

	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		if err := os.WriteFile(goldenFilePath, []byte(got), 0644); err != nil {
			return fmt.Errorf("writing testdata/fileschanged.txt: %w", err)
		}
	}

	if got != string(want) {
		return fmt.Errorf("changed files mismatch\n\ngot:\n%s\n\nwant:\n%s", got, want)
	}

	return nil
}
