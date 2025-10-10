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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	ProtoSourcePath string
	APIVersion      string
	ConfigFilePath  string
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
	cmd.PersistentFlags().StringVarP(&o.APIVersion, "api-version", "v", o.APIVersion, "the KRM API version. used to import the KRM API")
	cmd.PersistentFlags().StringVar(&o.ConfigFilePath, "config", "", "path to service config file, the config file will override other flags")
}

func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	repoRoot := strings.TrimSpace(string(output))
	return repoRoot, nil
}
