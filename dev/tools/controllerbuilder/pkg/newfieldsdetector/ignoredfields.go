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

package newfieldsdetector

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/util/sets"
)

// IgnoredFieldsConfig represents the structure of the ignored fields YAML file.
//
// Example YAML:
/*
google.cloud.bigquery.connection.v1:
  Connection:
	- salesforceDataCloud
google.api.apikeys.v2:
  Key:
	- createTime
	- updateTime
*/
type IgnoredFieldsConfig struct {
	// key is proto package name (e.g., "google.cloud.compute.v1").
	ProtoPackages map[string]MessageFieldIgnores `yaml:",inline"`
}

type MessageFieldIgnores struct {
	// key is proto message name (e.g. "Instance")
	// value is list of field names to be ignored in the message.
	Messages map[string][]string `yaml:",inline"`
}

// LoadIgnoredFields loads and parses the ignored fields YAML file
func LoadIgnoredFields(configPath string) (sets.String, error) {
	if configPath == "" {
		return sets.NewString(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("reading ignored fields config: %w", err)
	}
	var config IgnoredFieldsConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing ignored fields config: %w", err)
	}
	ignoredFields := sets.NewString()
	// use fully qualified field names in ignoredFields map. e.g. "google.cloud.compute.v1.Instance.id"
	for pkgName, pkgIgnores := range config.ProtoPackages {
		for msgName, fields := range pkgIgnores.Messages {
			for _, fieldName := range fields {
				fullyQualifiedName := fmt.Sprintf("%s.%s.%s", pkgName, msgName, fieldName)
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
