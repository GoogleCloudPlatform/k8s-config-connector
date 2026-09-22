// Copyright 2024 Google LLC
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

package options

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	ProtoSourcePath      string
	ProtoOverlayPath     string
	APIVersion           string
	ConfigFilePath       string
	IncludeSkippedOutput bool
}

func (o *GenerateOptions) InitDefaults() error {
	root, err := RepoRoot()
	if err != nil {
		return err
	}
	o.ProtoSourcePath = root + "/.build/googleapis.pb"
	return nil
}

func (o *GenerateOptions) BindPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&o.ProtoSourcePath, "proto-source-path", o.ProtoSourcePath, "path to (compiled) proto for APIs")
	cmd.PersistentFlags().StringVar(&o.ProtoOverlayPath, "overlay", o.ProtoOverlayPath, "path to a proto overlay file")
	cmd.PersistentFlags().StringVarP(&o.APIVersion, "api-version", "v", o.APIVersion, "the KRM API version. used to import the KRM API")
	cmd.PersistentFlags().StringVar(&o.ConfigFilePath, "config", "", "path to service config file, the config file will override other flags")
	cmd.PersistentFlags().BoolVar(&o.IncludeSkippedOutput, "include-skipped-output", false, "output skipped types and mappers as commented-out code")
}
func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	if output, err := cmd.Output(); err == nil {
		return strings.TrimSpace(string(output)), nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("failed to locate repo root (go.mod not found in parents)")
}
