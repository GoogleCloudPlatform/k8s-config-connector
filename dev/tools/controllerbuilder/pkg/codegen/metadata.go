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

type ServiceMetadata struct {
	Service        string             `yaml:"service"`
	APIVersion     string             `yaml:"apiVersion"`
	GenerateMapper bool               `yaml:"generateMapper"` // TODO: remove this field and always generate types and mappers together
	Resources      []ResourceMetadata `yaml:"resources"`
}

type ResourceMetadata struct {
	Kind              string           `yaml:"kind"`
	ProtoName         string           `yaml:"protoName"`
	SkipScaffoldFiles bool             `yaml:"skipScaffoldFiles"`
	IgnoredFields     []string         `yaml:"ignoredFields,omitempty"`
	ReferenceFields   []ReferenceField `yaml:"referenceFields,omitempty"`
}

type ReferenceField struct {
	// ProtoField is the fully qualified proto field name
	// e.g., "google.cloud.managedkafka.v1.NetworkConfig.subnet"
	ProtoField string `yaml:"protoField"`
	// GoName is the name of the reference field in Go
	// e.g., "SubnetworkRef"
	GoName string `yaml:"goName"`
	// RefType is the referenced KRM resource type
	// e.g., "ComputeSubnetworkRef"
	RefType string `yaml:"refType"`
	// Optional. Description is the description of the reference field.
	// If provided, this will override the original field comment.
	Description string `yaml:"description,omitempty"`
}

// LoadAllServiceMetadata loads all service metadata from a given path.
// If the path points to a directory, it recursively loads all .yaml files in that directory.
// Returns a map of service names to their metadata.
func LoadAllServiceMetadata(metadataPath string) (map[string]*ServiceMetadata, error) {
	if metadataPath == "" {
		return nil, nil
	}

	metadata := make(map[string]*ServiceMetadata)

	fileInfo, err := os.Stat(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat metadata path: %v", err)
	}

	if fileInfo.IsDir() {
		// Walk through all files in the directory and its subdirectories
		err := filepath.WalkDir(metadataPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Skip directories and non-yaml files
			if d.IsDir() || (!strings.HasSuffix(d.Name(), ".yaml") && !strings.HasSuffix(d.Name(), ".yml")) {
				return nil
			}

			serviceMetadata, err := LoadServiceMetadata(path)
			if err != nil {
				return fmt.Errorf("loading metadata from %s: %w", path, err)
			}
			if metadata != nil {
				metadata[serviceMetadata.Service] = serviceMetadata
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("walking metadata directory: %v", err)
		}
	} else {
		// Load single service metadata file
		serviceMetadata, err := LoadServiceMetadata(metadataPath)
		if err != nil {
			return nil, err
		}
		if serviceMetadata != nil {
			metadata[serviceMetadata.Service] = serviceMetadata
		}
	}

	return metadata, nil
}

// LoadServiceMetadata loads a single service metadata file
func LoadServiceMetadata(metadataPath string) (*ServiceMetadata, error) {
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %v", err)
	}

	var serviceMetadata ServiceMetadata
	if err := yaml.Unmarshal(data, &serviceMetadata); err != nil {
		return nil, fmt.Errorf("failed to parse metadata file: %v", err)
	}

	if serviceMetadata.Service == "" {
		return nil, fmt.Errorf("service name is required in metadata file %s", metadataPath)
	}

	return &serviceMetadata, nil
}

// LoadIgnoredFields returns a set of fully qualified field names that should be ignored
// For example: "google.cloud.bigquery.datatransfer.v1.TransferConfig.error.details"
func LoadIgnoredFields(metadataPath string) (sets.String, error) {
	metadata, err := LoadAllServiceMetadata(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("loading metadata: %w", err)
	}

	ignoredFields := sets.NewString()
	if metadata == nil {
		return ignoredFields, nil
	}

	for _, serviceMetadata := range metadata {
		for _, resource := range serviceMetadata.Resources {
			resourceFullName := fmt.Sprintf("%s.%s", serviceMetadata.Service, resource.ProtoName)
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
