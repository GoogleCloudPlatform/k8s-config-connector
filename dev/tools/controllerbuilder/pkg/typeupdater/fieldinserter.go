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

package typeupdater

import (
	"bytes"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type InsertFieldOptions struct {
	ParentMessageFullName string
	FieldToInsert         string
	ProtoSourcePath       string
	APIDirectory          string
	GoPackagePath         string
	MetadataDir           string
}

type FieldInserter struct {
	opts *InsertFieldOptions
	// newField is the internal representation of the field to be inserted
	newField newField
	// dependentMessages is a map nested messages under the new field to be inserted.
	// key: fully qualified name of proto message
	// value: internal representation of the messages to be inserted
	dependentMessages map[string]newMessage
	// ignoredFields is a set of fields that are ignored in the metadata files
	ignoredFields sets.String
}

type newField struct {
	proto            protoreflect.FieldDescriptor
	parent           protoreflect.MessageDescriptor
	generatedContent []byte // the content of the generated Go type corresponding to this field
}

type newMessage struct {
	proto            protoreflect.MessageDescriptor
	generatedContent []byte // the content of the generated Go type corresponding to this message
}

func NewFieldInserter(opts *InsertFieldOptions) *FieldInserter {
	ignoredFields, err := codegen.LoadIgnoredFields(opts.MetadataDir)
	if err != nil {
		klog.Fatalf("failed to load metadata files in directory %s: %v", opts.MetadataDir, err)
	}

	return &FieldInserter{
		opts:          opts,
		ignoredFields: ignoredFields,
	}
}

func (u *FieldInserter) Run() error {
	// 1. find new field and its dependent proto messages that needs to be generated
	if err := u.analyze(); err != nil {
		return err
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

// anaylze finds the new field, its parent message, and all dependent messages that need to be generated.
func (u *FieldInserter) analyze() error {
	// find the new proto field to be inserted
	newProtoField, parent, err := findNewField(u.opts.ProtoSourcePath, u.opts.ParentMessageFullName, u.opts.FieldToInsert)
	if err != nil {
		return fmt.Errorf("failed to find field to insert in proto: %w", err)
	}
	u.newField = newField{
		proto:  newProtoField,
		parent: parent,
	}

	// find the dependent proto messags of this new field
	msgs, err := findDependentMsgs(newProtoField, u.ignoredFields)
	if err != nil {
		return err
	}
	codegen.RemoveNotMappedToGoStruct(msgs)
	if err := removeAlreadyGenerated(u.opts.GoPackagePath, u.opts.APIDirectory, msgs); err != nil {
		return err
	}
	u.dependentMessages = make(map[string]newMessage, len(msgs))
	for _, msg := range msgs {
		u.dependentMessages[string(msg.FullName())] = newMessage{
			proto: msg,
		}
	}
	return nil
}

// findNewField locates the parent message and the new field in the proto file
func findNewField(pbSourcePath, parentMsgFullName, newFieldName string) (protoreflect.FieldDescriptor, protoreflect.MessageDescriptor, error) {
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

	return fieldDesc, msgDesc, nil
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

func (u *FieldInserter) generate() error {
	var buf bytes.Buffer
	klog.Infof("generate Go code for field %s", u.newField.proto.Name())
	codegen.WriteField(&buf, u.newField.proto, u.newField.parent, 0, false) // TODO: add support for transitive output fields
	u.newField.generatedContent = buf.Bytes()

	for key, msg := range u.dependentMessages {
		var buf bytes.Buffer
		klog.Infof("generate Go code for messge %s", msg.proto.FullName())
		codegen.WriteMessage(&buf, msg.proto)
		msg.generatedContent = buf.Bytes()
		u.dependentMessages[key] = msg
	}
	return nil
}
