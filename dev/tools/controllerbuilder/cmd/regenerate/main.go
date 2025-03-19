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
	"path/filepath"
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
}

func (o *Options) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.SourceDir = root
	return nil
}

func (o *Options) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.SourceDir, "src", o.SourceDir, "path to source code")
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

	// Walk srcdir, visiting all go files
	if err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		// fmt.Printf("checking %s\n", path)
		if err := checkFile(ctx, path); err != nil {
			return fmt.Errorf("checking file %q: %w", path, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory %s: %w", srcDir, err)
	}
	return nil
}

func checkFile(ctx context.Context, p string) error {
	src, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}

	markers := []string{"+generated:"}
	annotations, err := annotations.FindFileAnnotations(src, markers)
	if err != nil {
		return err
	}
	for _, annotation := range annotations {
		fmt.Printf("found annotation in %q: %+v\n", p, annotation)

		switch annotation.Key {
		case "+generated:mapper":
			service := strings.Join(annotation.Attributes["service"], ",")
			apiVersion := strings.Join(annotation.Attributes["krm.group"], ",") + "/" + strings.Join(annotation.Attributes["krm.version"], ",")
			args := []string{
				"generate-mapper",
				"--service", service,
				"--api-version", apiVersion,
			}
			fmt.Printf("command is: %v\n", strings.Join(args, " "))
		}
	}

	return nil

}
