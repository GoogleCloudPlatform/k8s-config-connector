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

package codegen

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Service        string           `yaml:"service"`
	APIVersion     string           `yaml:"apiVersion"`
	GenerateMapper bool             `yaml:"generateMapper"`
	Resources      []ResourceConfig `yaml:"resources"`
}

type ResourceConfig struct {
	Kind              string `yaml:"kind"`
	ProtoName         string `yaml:"protoName"`
	SkipScaffoldFiles bool   `yaml:"skipScaffoldFiles"`
}

func LoadConfig(configPath string) (*ServiceConfig, error) {
	if configPath == "" {
		return nil, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config ServiceConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
