// Copyright 2026 Google LLC
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
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ghodss/yaml" //nolint:depguard
)

type ResourceConfig struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type ServiceMapping struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Name      string           `json:"name"`
		Version   string           `json:"version"`
		Resources []ResourceConfig `json:"resources"`
	} `json:"spec"`
}

func main() {
	content, err := ioutil.ReadFile("scripts/resource-autogen/allowlist/allowlist.go")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	inAlpha := false
	for _, line := range lines {
		if strings.Contains(line, "alphaAllowlist = []string{") {
			inAlpha = true
			continue
		}
		if inAlpha && strings.Contains(line, "}") {
			inAlpha = false
			continue
		}
		if inAlpha {
			trimmed := strings.Trim(line, "\t \",")
			if trimmed == "" || strings.HasPrefix(trimmed, "//") {
				continue
			}
			parts := strings.Split(trimmed, "/")
			if len(parts) != 2 {
				continue
			}
			service := strings.Replace(parts[0], "_", "", -1)

			smFile := "scripts/resource-autogen/generated/servicemappings/" + service + ".yaml"
			smContent, err := ioutil.ReadFile(smFile)
			if err != nil {
				// Try with underscores removed from service name in a more sophisticated way if needed,
				// but let's try this first.
				// Wait, the service name in scripts/resource-autogen/generated/servicemappings/
				// seems to match the parts[0] in many cases.
				smFile = "scripts/resource-autogen/generated/servicemappings/" + parts[0] + ".yaml"
				smContent, err = ioutil.ReadFile(smFile)
				if err != nil {
					continue
				}
			}

			var sm ServiceMapping
			if err := yaml.Unmarshal(smContent, &sm); err != nil {
				continue
			}

			if sm.Spec.Version == "v1beta1" {
				fmt.Printf("%s\n", trimmed)
			}
		}
	}
}
