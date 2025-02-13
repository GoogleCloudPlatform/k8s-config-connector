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
	"os/exec"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

type MessageDiff struct { // This structure tracks the difference in fields for a given proto message
	MessageName   string
	NewFields     []string               // fields added in the new version
	RemovedFields []string               // fields removed in the new version
	ChangedFields map[string]FieldChange // fields changed in the new version
}

type FieldChange struct {
	OldType    protoreflect.Kind
	NewType    protoreflect.Kind
	IsRepeated bool
}

type DetectorOptions struct {
	TargetMessages sets.String
	ConfigDir      string
}

type FieldDetector struct {
	opts          *DetectorOptions
	oldFiles      *protoregistry.Files // proto files that are generated with the pinned version
	newFiles      *protoregistry.Files // proto files that are generated from the remote HEAD
	ignoredFields sets.String
}

// NewFieldDetector detects any proto field changes between the current pinned and the HEAD.
func NewFieldDetector(opts *DetectorOptions) (*FieldDetector, error) {
	repoRoot, err := options.RepoRoot()
	if err != nil {
		return nil, fmt.Errorf("finding repo root: %w", err)
	}
	pinnedProtoPath := filepath.Join(repoRoot, ".build", "googleapis.pb")
	headProtoPath := filepath.Join(repoRoot, ".build", "googleapis_head.pb")

	// generate pinned version proto. The default version is recorded in "mockgcp/git.versions".
	if err := generateProto(repoRoot, pinnedProtoPath, ""); err != nil {
		return nil, fmt.Errorf("generating pinned proto: %w", err)
	}
	old, err := protoapi.LoadProto(pinnedProtoPath)
	if err != nil {
		return nil, fmt.Errorf("loading old proto: %w", err)
	}
	// generate HEAD version proto
	if err := generateProto(repoRoot, headProtoPath, "HEAD"); err != nil {
		return nil, fmt.Errorf("generating HEAD proto: %w", err)
	}
	new, err := protoapi.LoadProto(headProtoPath)
	if err != nil {
		return nil, fmt.Errorf("loading new proto: %w", err)
	}

	// load ignored fields
	ignoredFields, err := codegen.LoadIgnoredFields(opts.ConfigDir)
	if err != nil {
		return nil, fmt.Errorf("loading ignored fields: %w", err)
	}

	return &FieldDetector{
		opts:          opts,
		oldFiles:      old.Files(),
		newFiles:      new.Files(),
		ignoredFields: ignoredFields,
	}, nil
}

// use the script at dev/tools/controllerbuilder/generate-proto.sh
func generateProto(repoRoot, outputPath, version string) error {
	scriptPath := filepath.Join(repoRoot, "dev", "tools", "controllerbuilder", "generate-proto.sh")

	cmd := exec.Command("bash", scriptPath, version, outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (d *FieldDetector) DetectNewFields() ([]MessageDiff, error) {
	targets := d.opts.TargetMessages
	// auto populate target messages if not specified by user
	if len(targets) == 0 {
		var err error
		targets, err = d.defaultTargetMessages()
		if err != nil {
			return nil, fmt.Errorf("failed to getdefault target messages: %w", err)
		}
	}
	if len(targets) == 0 {
		return nil, fmt.Errorf("no target messages specified")
	}

	var diffs []MessageDiff
	for fqn := range targets {
		diff, err := d.compareMessage(d.oldFiles, d.newFiles, fqn)
		if err != nil {
			return nil, fmt.Errorf("error when comparing message %s: %w", fqn, err)
		}
		if hasChanges(diff) {
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func hasChanges(diff MessageDiff) bool {
	return len(diff.NewFields) > 0 || len(diff.RemovedFields) > 0 || len(diff.ChangedFields) > 0
}

func (d *FieldDetector) compareMessage(oldFiles, newFiles *protoregistry.Files, messageName string) (MessageDiff, error) {
	diff := MessageDiff{
		MessageName:   messageName,
		ChangedFields: make(map[string]FieldChange),
	}

	oldMsg := findMessage(oldFiles, messageName)
	newMsg := findMessage(newFiles, messageName)

	if oldMsg == nil && newMsg == nil {
		return diff, fmt.Errorf("message %s not found in either file", messageName)
	}

	// case 1. new message added
	if oldMsg == nil {
		newFields := getMessageFields(newMsg)
		for fieldName := range newFields {
			if codegen.IsFieldIgnored(d.ignoredFields, messageName, fieldName) {
				klog.Warningf("new field %s in message %s is ignored", fieldName, messageName)
				continue
			}
			diff.NewFields = append(diff.NewFields, fieldName)
		}
		return diff, nil
	}

	// case 2. message removed
	if newMsg == nil {
		oldFields := getMessageFields(oldMsg)
		for fieldName := range oldFields {
			if codegen.IsFieldIgnored(d.ignoredFields, messageName, fieldName) {
				klog.Warningf("removed field %s in message %s is ignored", fieldName, messageName)
				continue
			}
			diff.RemovedFields = append(diff.RemovedFields, fieldName)
		}
		return diff, nil
	}

	// case 3. message exists in both old and new proto
	oldFields := getMessageFields(oldMsg)
	newFields := getMessageFields(newMsg)

	// 3.1 Find new and changed fields
	for fieldName, newField := range newFields {
		if codegen.IsFieldIgnored(d.ignoredFields, messageName, fieldName) {
			klog.Warningf("new field %s in message %s is ignored", fieldName, messageName)
			continue
		}

		oldField, exists := oldFields[fieldName]
		if !exists {
			diff.NewFields = append(diff.NewFields, fieldName)
			continue
		}

		if oldField.Kind() != newField.Kind() || oldField.IsList() != newField.IsList() {
			diff.ChangedFields[fieldName] = FieldChange{
				OldType:    oldField.Kind(),
				NewType:    newField.Kind(),
				IsRepeated: newField.IsList(),
			}
		}
	}

	// 3.2 Find removed fields
	for fieldName := range oldFields {
		if codegen.IsFieldIgnored(d.ignoredFields, messageName, fieldName) {
			klog.Warningf("removed field %s in message %s is ignored", fieldName, messageName)
			continue
		}

		if _, exists := newFields[fieldName]; !exists {
			diff.RemovedFields = append(diff.RemovedFields, fieldName)
		}
	}

	return diff, nil
}

func findMessage(files *protoregistry.Files, name string) protoreflect.MessageDescriptor {
	desc, err := files.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil
	}
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	if !ok {
		return nil
	}
	return msgDesc
}

func getMessageFields(msg protoreflect.MessageDescriptor) map[string]protoreflect.FieldDescriptor {
	fields := make(map[string]protoreflect.FieldDescriptor)
	fieldDescriptors := msg.Fields()
	for i := 0; i < fieldDescriptors.Len(); i++ {
		field := fieldDescriptors.Get(i)
		fields[string(field.Name())] = field
	}
	return fields
}
