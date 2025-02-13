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
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ServiceConfig struct {
	Service        string           `yaml:"service"`
	APIVersion     string           `yaml:"apiVersion"`
	GenerateMapper bool             `yaml:"generateMapper"`
	Resources      []ResourceConfig `yaml:"resources"`
}

type ResourceConfig struct {
	Kind              string   `yaml:"kind"`
	ProtoName         string   `yaml:"protoName"`
	SkipScaffoldFiles bool     `yaml:"skipScaffoldFiles"`
	IgnoredFields     []string `yaml:"ignoredFields,omitempty"`
}

// LoadAllConfigs loads service configurations from a file or directory path.
// If the path points to a directory, it recursively loads all .yaml files in that directory.
// Returns a map of service names to their configurations.
func LoadAllConfigs(configPath string) (map[string]*ServiceConfig, error) {
	if configPath == "" {
		return nil, nil
	}

	configs := make(map[string]*ServiceConfig)

	fileInfo, err := os.Stat(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat config path: %v", err)
	}

	if fileInfo.IsDir() {
		// Walk through all files in the directory and its subdirectories
		err := filepath.WalkDir(configPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Skip directories and non-yaml files
			if d.IsDir() || (!strings.HasSuffix(d.Name(), ".yaml") && !strings.HasSuffix(d.Name(), ".yml")) {
				return nil
			}

			config, err := LoadConfig(path)
			if err != nil {
				return fmt.Errorf("loading config from %s: %w", path, err)
			}
			if config != nil {
				configs[config.Service] = config
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("walking config directory: %v", err)
		}
	} else {
		// Load single config file
		config, err := LoadConfig(configPath)
		if err != nil {
			return nil, err
		}
		if config != nil {
			configs[config.Service] = config
		}
	}

	return configs, nil
}

// LoadConfig loads a single service configuration file
func LoadConfig(configPath string) (*ServiceConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config ServiceConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	if config.Service == "" {
		return nil, fmt.Errorf("service name is required in config file %s", configPath)
	}

	return &config, nil
}

// LoadIgnoredFields returns a set of fully qualified field names that should be ignored
// For example: "google.cloud.bigquery.datatransfer.v1.TransferConfig.error.details"
func LoadIgnoredFields(configPath string) (sets.String, error) {
	configs, err := LoadAllConfigs(configPath)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	ignoredFields := sets.NewString()
	if configs == nil {
		return ignoredFields, nil
	}

	for _, config := range configs {
		for _, resource := range config.Resources {
			resourceFullName := fmt.Sprintf("%s.%s", config.Service, resource.ProtoName)
			for _, field := range resource.IgnoredFields {
				fullyQualifiedName := fmt.Sprintf("%s.%s", resourceFullName, field)
				ignoredFields.Insert(fullyQualifiedName)
			}
		}
	}
	return ignoredFields, nil
}

// IsFieldIgnored checks if a field should be ignored based on its fully qualified name
func IsFieldIgnored(ignoredFields sets.String, fullyQualifiedMessageName, fieldName string) bool {
	fullyQualifiedFieldName := fmt.Sprintf("%s.%s", fullyQualifiedMessageName, fieldName)
	return ignoredFields.Has(fullyQualifiedFieldName)
}
