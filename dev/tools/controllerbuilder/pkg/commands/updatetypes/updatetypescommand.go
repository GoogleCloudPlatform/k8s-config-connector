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

package updatetypes

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

const kccProtoPrefix = "+kcc:proto="

type UpdateTypeOptions struct {
	*options.GenerateOptions

	parentMessageFullName string
	newField              string
	ignoredFields         string // TODO: could be part of GenerateOptions
	apiDirectory          string
	goPackagePath         string
}

func (o *UpdateTypeOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return nil
	}
	o.apiDirectory = root + "/apis/"
	o.goPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	return nil
}

func (o *UpdateTypeOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.parentMessageFullName, "parent-message-full-name", o.parentMessageFullName, "Fully qualified name of the proto message holding the new field")
	cmd.Flags().StringVar(&o.newField, "new-field", o.newField, "Name of the new field")
	cmd.Flags().StringVar(&o.ignoredFields, "ignored-fields", o.ignoredFields, "Comma-separated list of fields to ignore")
	cmd.Flags().StringVar(&o.apiDirectory, "api-dir", o.apiDirectory, "Base directory for APIs")
	cmd.Flags().StringVar(&o.goPackagePath, "api-go-package-path", o.goPackagePath, "Package path")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &UpdateTypeOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "update-types",
		Short: "update KRM types for a proto service",
		RunE: func(cmd *cobra.Command, args []string) error {
			updater := NewTypeUpdater(opt)
			if err := updater.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

type TypeUpdater struct {
	opts               *UpdateTypeOptions
	newField           newProtoField
	dependentMessages  map[string]protoreflect.MessageDescriptor // key: fully qualified name of proto message
	generatedGoField   generatedGoField                          // TODO: support multiple new fields
	generatedGoStructs []generatedGoStruct
}

type newProtoField struct {
	field         protoreflect.FieldDescriptor
	parentMessage protoreflect.MessageDescriptor
}

type generatedGoField struct {
	parentMessage string // fully qualified name of the parent proto message of this field
	content       []byte // the content of the generated Go field
}

type generatedGoStruct struct {
	name    string // fully qualified name of the proto message
	content []byte // the content of the generated Go struct
}

func NewTypeUpdater(opts *UpdateTypeOptions) *TypeUpdater {
	return &TypeUpdater{
		opts: opts,
	}
}

func (u *TypeUpdater) Run() error {
	// 1. find new field and its dependent proto messages that needs to be generated
	if err := u.analyze(); err != nil {
		return nil
	}

	// 2. generate Go types for the new field and its dependent proto messages
	if err := u.generate(); err != nil {
		return err
	}

	// 3. insert the generated Go code back to files
	if err := u.insertGoField(); err != nil {
		return err
	}
	if err := u.insertGoMessages(); err != nil {
		return err
	}

	return nil
}

// analyze finds the new field, its parent message, and all dependent messages that need to be generated.
func (u *TypeUpdater) analyze() error {
	parentMessage, newField, err := findNewField(u.opts.ProtoSourcePath, u.opts.parentMessageFullName, u.opts.newField)
	if err != nil {
		return err
	}
	u.newField = newProtoField{
		field:         newField,
		parentMessage: parentMessage,
	}

	msgs, err := findDependentMsgs(newField, sets.NewString(strings.Split(u.opts.ignoredFields, ",")...))
	if err != nil {
		return err
	}

	codegen.RemoveNotMappedToGoStruct(msgs)

	if err := removeAlreadyGenerated(u.opts.goPackagePath, u.opts.apiDirectory, msgs); err != nil {
		return err
	}
	u.dependentMessages = msgs
	return nil
}

// findNewField locates the parent message and the new field in the proto file
func findNewField(pbSourcePath, parentMsgFullName, newFieldName string) (protoreflect.MessageDescriptor, protoreflect.FieldDescriptor, error) {
	fileData, err := os.ReadFile(pbSourcePath)
	if err != nil {
		return nil, nil, fmt.Errorf("reading %q: %w", pbSourcePath, err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(fileData, fds); err != nil {
		return nil, nil, fmt.Errorf("unmarshalling %q: %w", pbSourcePath, err)
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return nil, nil, fmt.Errorf("building file description: %w", err)
	}

	// Find the parent message
	messageDesc, err := files.FindDescriptorByName(protoreflect.FullName(parentMsgFullName))
	if err != nil {
		return nil, nil, err
	}
	msgDesc, ok := messageDesc.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected descriptor type: %T", msgDesc)
	}

	// Find the new field in parent message
	fieldDesc := msgDesc.Fields().ByName(protoreflect.Name(newFieldName))
	if fieldDesc == nil {
		return nil, nil, fmt.Errorf("field not found in message")
	}

	return msgDesc, fieldDesc, nil
}

// findDependentMsgs finds all dependent proto messages for the given field, ignoring specified fields
func findDependentMsgs(field protoreflect.FieldDescriptor, ignoredProtoFields sets.String) (map[string]protoreflect.MessageDescriptor, error) {
	deps := make(map[string]protoreflect.MessageDescriptor)
	codegen.FindDependenciesForField(field, deps, ignoredProtoFields)
	return deps, nil
}

// removeAlreadyGenerated removes proto messages that have already been generated (including manually edited)
func removeAlreadyGenerated(goPackagePath, outputAPIDirectory string, targets map[string]protoreflect.MessageDescriptor) error {
	packages, err := gocode.LoadPackageTree(goPackagePath, outputAPIDirectory)
	if err != nil {
		return err
	}
	for _, p := range packages {
		for _, s := range p.Structs {
			if annotation := s.GetAnnotation("+kcc:proto"); annotation != "" {
				delete(targets, annotation)
			}
		}
	}
	return nil
}

func (u *TypeUpdater) generate() error {
	var buf bytes.Buffer
	klog.Infof("generate Go code for field %s", u.newField.field.Name())
	codegen.WriteField(&buf, u.newField.field, u.newField.parentMessage, 0)
	u.generatedGoField = generatedGoField{
		parentMessage: string(u.newField.parentMessage.FullName()),
		content:       buf.Bytes(),
	}

	for _, msg := range u.dependentMessages {
		var buf bytes.Buffer
		klog.Infof("generate Go code for messge %s", msg.FullName())
		codegen.WriteMessage(&buf, msg)
		u.generatedGoStructs = append(u.generatedGoStructs,
			generatedGoStruct{
				name:    string(msg.FullName()),
				content: buf.Bytes(),
			})
	}
	return nil
}
