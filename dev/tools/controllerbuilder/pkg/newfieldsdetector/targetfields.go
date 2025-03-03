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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

// GenerateTypesFlags represents the flags used in "generate-types" command
type GenerateTypesFlags struct {
	Service   string
	Resources []ResourceMapping
}

type ResourceMapping struct {
	KRMType      string
	ProtoMessage string
}

// defaultTargetMessages processes the command-line flags used in "generate.sh" to extract fully qualified proto message names.
// These names represent the top-level messages corresponding to Config Connector KRM resources.
// The function also identifies nested messages of these top-level messages to form a comprehensive list of target messages.
// TODO: Create a structured file to store this information for more reliable parsing. The file can be shared between the type generator and the new field detector.
func (d *FieldDetector) defaultTargetMessages() (sets.String, error) {
	topLevelMessages, err := extractMessagesFromGenerateTypesScript()
	if err != nil {
		return nil, fmt.Errorf("failed to extract messages from generate.sh script: %w", err)
	}

	expandedMessages, err := d.expandToIncludeNestedMessages(topLevelMessages, d.newFiles) // note: using new proto files here
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return expandedMessages, nil
}

func extractMessagesFromGenerateTypesScript() (sets.String, error) {
	repoRoot, err := options.RepoRoot()
	if err != nil {
		return nil, fmt.Errorf("finding repo root: %w", err)
	}
	scriptPath := filepath.Join(repoRoot, "dev", "tools", "controllerbuilder", "generate.sh")

	flags, err := parseGenerateTypesScript(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("parsing generate types script: %w", err)
	}

	targetMessages := sets.NewString()
	for _, flag := range flags {
		for _, resource := range flag.Resources {
			targetMessages.Insert(fmt.Sprintf("%s.%s", flag.Service, resource.ProtoMessage))
		}
	}

	return targetMessages, nil
}

// expandToIncludeNestedMessages expands initial target messages to include all nested messages
func (d *FieldDetector) expandToIncludeNestedMessages(initialTargets sets.String, files *protoregistry.Files) (sets.String, error) {
	allTargets := sets.NewString()
	allTargets.Insert(initialTargets.List()...) // make a copy

	for name := range initialTargets {
		desc, err := files.FindDescriptorByName(protoreflect.FullName(name))
		if err != nil {
			klog.Warningf("Proto message not found %s: %v", name, err)
			continue
		}

		message, ok := desc.(protoreflect.MessageDescriptor)
		if !ok {
			return nil, fmt.Errorf("unexpected descriptor type: %T", desc)
		}

		deps, err := codegen.FindDependenciesForMessage(message, d.ignoredFields)
		if err != nil {
			return nil, fmt.Errorf("failed to find dependencies for message %s: %w", name, err)
		}
		for _, dep := range deps {
			allTargets.Insert(string(dep.FullName()))
		}
	}

	return allTargets, nil
}

func parseGenerateTypesScript(scriptPath string) ([]GenerateTypesFlags, error) {
	file, err := os.Open(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("opening script file: %w", err)
	}
	defer file.Close()

	var flags []GenerateTypesFlags
	var currentFlags *GenerateTypesFlags

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines, comments, and non-flag lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Start of new "generate-types" command
		if strings.Contains(line, "generate-types") {
			if currentFlags != nil {
				flags = append(flags, *currentFlags)
			}
			currentFlags = &GenerateTypesFlags{}
			continue
		}

		if currentFlags == nil {
			continue
		}

		// Parse flags to get proto package and proto message name
		if strings.Contains(line, "--service") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				currentFlags.Service = parts[1]
			}
		} else if strings.Contains(line, "--resource") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				resourceParts := strings.Split(parts[1], ":")
				if len(resourceParts) == 2 {
					currentFlags.Resources = append(currentFlags.Resources, ResourceMapping{
						KRMType:      resourceParts[0],
						ProtoMessage: resourceParts[1],
					})
				}
			}
		}

		// End of a "generate-types" command
		if !strings.HasSuffix(line, "\\") && currentFlags != nil {
			flags = append(flags, *currentFlags)
			currentFlags = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning script file: %w", err)
	}

	return flags, nil
}
