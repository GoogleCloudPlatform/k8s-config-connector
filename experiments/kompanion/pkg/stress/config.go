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

package stress

import (
	"fmt"
	"os"
	"time"

	"sigs.k8s.io/yaml"
)

type StressConfig struct {
	NumNamespaces     int                `json:"numNamespaces"`
	TypeOfStress      string             `json:"typeOfStress"`
	IAMReference      IAMReference       `json:"iamReference"`
	ResourcesToStress []ResourceToStress `json:"resourcesToStress"`
	Duration          string             `json:"duration"`
}

type IAMReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ResourceToStress struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
	Role  string `json:"role"`
}

func (c *StressConfig) Validate() error {
	if c.NumNamespaces <= 0 {
		return fmt.Errorf("numNamespaces must be greater than 0")
	}
	if c.TypeOfStress == "" {
		return fmt.Errorf("typeOfStress must be specified")
	}
	if c.IAMReference.Kind == "" || c.IAMReference.Name == "" || c.IAMReference.Namespace == "" {
		return fmt.Errorf("iamReference must have kind, name, and namespace specified")
	}
	if len(c.ResourcesToStress) == 0 {
		return fmt.Errorf("resourcesToStress must contain at least one resource")
	}
	for _, r := range c.ResourcesToStress {
		if r.Kind == "" || r.Count <= 0 || r.Role == "" {
			return fmt.Errorf("each resourceToStress must have kind, count, and role specified")
		}
	}
	if _, err := time.ParseDuration(c.Duration); err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	return nil
}

func LoadConfig(filePath string) (*StressConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config StressConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}
