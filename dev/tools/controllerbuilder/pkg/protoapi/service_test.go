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

package protoapi

import (
	"errors"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestResourceAPIFor_StandardCRUDAndLRO(t *testing.T) {
	file := testFile("google/devtools/artifactregistry/v1/service.proto", "google.devtools.artifactregistry.v1")
	addAIPResource(file, "Repository", "artifactregistry.googleapis.com/Repository",
		"projects/{project}/locations/{location}/repositories/{repository}")
	addLabels(file.MessageType[0])
	addAIPMessages(file, "Repository")
	addAIPMethods(file, "ArtifactRegistry", "artifactregistry.googleapis.com", "Repository", true)

	p := newTestProto(t, file)
	api := mustResourceAPI(t, p, "google.devtools.artifactregistry.v1.Repository")

	if api.Service == nil || api.Service.Name() != "ArtifactRegistry" {
		t.Fatalf("Service = %v, want ArtifactRegistry", api.Service)
	}
	if api.DefaultHost != "artifactregistry.googleapis.com" {
		t.Errorf("DefaultHost = %q, want artifactregistry.googleapis.com", api.DefaultHost)
	}
	if !api.HasLabels {
		t.Errorf("HasLabels = false, want true")
	}

	if api.Get == nil || api.Get.NameField != "name" {
		t.Fatalf("Get = %+v, want NameField=name", api.Get)
	}
	if api.Create == nil {
		t.Fatalf("Create = nil")
	}
	if api.Create.ParentField != "parent" || api.Create.ResourceField != "repository" || api.Create.IDField != "repository_id" {
		t.Errorf("Create fields = (%q, %q, %q), want (parent, repository, repository_id)",
			api.Create.ParentField, api.Create.ResourceField, api.Create.IDField)
	}
	if api.Create.LRO == nil || api.Create.LRO.ResponseType != "google.devtools.artifactregistry.v1.Repository" {
		t.Errorf("Create.LRO = %+v, want ResponseType=google.devtools.artifactregistry.v1.Repository", api.Create.LRO)
	}

	if api.Update == nil || api.Update.ResourceField != "repository" || api.Update.UpdateMaskField != "update_mask" {
		t.Fatalf("Update = %+v, want ResourceField=repository, UpdateMaskField=update_mask", api.Update)
	}
	if api.Delete == nil || api.Delete.NameField != "name" {
		t.Fatalf("Delete = %+v, want NameField=name", api.Delete)
	}
	if api.Delete.LRO == nil || api.Delete.LRO.ResponseType != "google.protobuf.Empty" {
		t.Errorf("Delete.LRO = %+v, want ResponseType=google.protobuf.Empty", api.Delete.LRO)
	}
}

func TestResourceAPIFor_ClientIDFieldSuffixFallback(t *testing.T) {
	file := testFile("google/devtools/cloudbuild/v1/cloudbuild.proto", "google.devtools.cloudbuild.v1")
	addAIPResource(file, "BuildTrigger", "cloudbuild.googleapis.com/BuildTrigger",
		"projects/{project}/locations/{location}/triggers/{trigger}")
	file.MessageType = append(file.MessageType,
		message("CreateBuildTriggerRequest",
			stringField("parent", 1),
			messageField("trigger", 2, ".google.devtools.cloudbuild.v1.BuildTrigger"),
			stringField("trigger_id", 3),
		),
	)
	file.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name: proto.String("CloudBuild"),
		Method: []*descriptorpb.MethodDescriptorProto{
			method("CreateBuildTrigger",
				".google.devtools.cloudbuild.v1.CreateBuildTriggerRequest",
				".google.devtools.cloudbuild.v1.BuildTrigger", nil),
		},
	}}

	p := newTestProto(t, file)
	api := mustResourceAPI(t, p, "google.devtools.cloudbuild.v1.BuildTrigger")
	if api.Create == nil || api.Create.IDField != "trigger_id" {
		t.Fatalf("Create.IDField = %q, want trigger_id", api.Create.IDField)
	}
}

func TestResourceAPIFor_UpdateMaskFieldFallback(t *testing.T) {
	file := testFile("google/spanner/admin/instance/v1/spanner_instance_admin.proto", "google.spanner.admin.instance.v1")
	addAIPResource(file, "Instance", "spanner.googleapis.com/Instance",
		"projects/{project}/instances/{instance}")
	file.MessageType = append(file.MessageType,
		message("UpdateInstanceRequest",
			messageField("instance", 1, ".google.spanner.admin.instance.v1.Instance"),
			messageField("response_read_mask", 2, ".google.protobuf.FieldMask"),
			messageField("field_mask", 3, ".google.protobuf.FieldMask"),
		),
	)
	file.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name: proto.String("InstanceAdmin"),
		Method: []*descriptorpb.MethodDescriptorProto{
			method("UpdateInstance",
				".google.spanner.admin.instance.v1.UpdateInstanceRequest",
				".google.spanner.admin.instance.v1.Instance", nil),
		},
	}}

	p := newTestProto(t, file)
	api := mustResourceAPI(t, p, "google.spanner.admin.instance.v1.Instance")
	if api.Update == nil || api.Update.UpdateMaskField != "field_mask" {
		t.Fatalf("Update.UpdateMaskField = %q, want field_mask", api.Update.UpdateMaskField)
	}
}

func TestResourceAPIByName_Errors(t *testing.T) {
	file := testFile("acme/v1/widget.proto", "acme.v1")
	file.EnumType = []*descriptorpb.EnumDescriptorProto{{
		Name:  proto.String("State"),
		Value: []*descriptorpb.EnumValueDescriptorProto{{Name: proto.String("STATE_UNSPECIFIED"), Number: proto.Int32(0)}},
	}}
	p := newTestProto(t, file)

	if _, err := p.ResourceAPIByName("acme.v1.NoSuchMessage"); !errors.Is(err, protoregistry.NotFound) {
		t.Errorf("ResourceAPIByName(missing) err = %v, want NotFound", err)
	}
	if _, err := p.ResourceAPIByName("acme.v1.State"); !errors.Is(err, protoregistry.NotFound) {
		t.Errorf("ResourceAPIByName(enum) err = %v, want NotFound", err)
	}
}

func TestSnakeCase(t *testing.T) {
	cases := map[string]string{
		"Repository":          "repository",
		"AppConnection":       "app_connection",
		"TlsInspectionPolicy": "tls_inspection_policy",
		"SACRealm":            "sac_realm",
		"VPCSCConfig":         "vpcsc_config",
	}
	for in, want := range cases {
		if got := snakeCase(in); got != want {
			t.Errorf("snakeCase(%q) = %q, want %q", in, got, want)
		}
	}
}

func addAIPResource(file *descriptorpb.FileDescriptorProto, name, resourceType, pattern string) {
	msg := message(name, stringField("name", 1))
	msg.Options = &descriptorpb.MessageOptions{}
	proto.SetExtension(msg.Options, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    resourceType,
		Pattern: []string{pattern},
	})
	file.MessageType = append(file.MessageType, msg)
}

func addLabels(msg *descriptorpb.DescriptorProto) {
	entry := &descriptorpb.DescriptorProto{
		Name:    proto.String("LabelsEntry"),
		Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
		Field: []*descriptorpb.FieldDescriptorProto{
			stringField("key", 1),
			stringField("value", 2),
		},
	}
	msg.NestedType = append(msg.NestedType, entry)
	msg.Field = append(msg.Field, &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("labels"),
		Number:   proto.Int32(int32(len(msg.Field) + 1)),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String("LabelsEntry"),
	})
}

func addAIPMessages(file *descriptorpb.FileDescriptorProto, name string) {
	pkg := file.GetPackage()
	resType := "." + pkg + "." + name
	snake := snakeCase(name)
	file.MessageType = append(file.MessageType,
		message("Get"+name+"Request", stringField("name", 1)),
		message("Create"+name+"Request",
			stringField("parent", 1),
			messageField(snake, 2, resType),
			stringField(snake+"_id", 3),
		),
		message("Update"+name+"Request",
			messageField(snake, 1, resType),
			messageField("update_mask", 2, ".google.protobuf.FieldMask"),
		),
		message("Delete"+name+"Request", stringField("name", 1)),
	)
}

func addAIPMethods(file *descriptorpb.FileDescriptorProto, serviceName, host, name string, lro bool) {
	pkg := file.GetPackage()
	resType := "." + pkg + "." + name
	createOut, deleteOut := resType, ".google.protobuf.Empty"
	var createOpts, deleteOpts *descriptorpb.MethodOptions
	if lro {
		createOut = ".google.longrunning.Operation"
		deleteOut = ".google.longrunning.Operation"
		createOpts = withOperationInfo(name)
		deleteOpts = withOperationInfo("google.protobuf.Empty")
	}

	svcOpts := &descriptorpb.ServiceOptions{}
	if host != "" {
		proto.SetExtension(svcOpts, annotations.E_DefaultHost, host)
	}
	file.Service = append(file.Service, &descriptorpb.ServiceDescriptorProto{
		Name:    proto.String(serviceName),
		Options: svcOpts,
		Method: []*descriptorpb.MethodDescriptorProto{
			method("Get"+name, "."+pkg+".Get"+name+"Request", resType, nil),
			method("Create"+name, "."+pkg+".Create"+name+"Request", createOut, createOpts),
			method("Update"+name, "."+pkg+".Update"+name+"Request", resType, nil),
			method("Delete"+name, "."+pkg+".Delete"+name+"Request", deleteOut, deleteOpts),
		},
	})
}

func newTestProto(t *testing.T, files ...*descriptorpb.FileDescriptorProto) *Proto {
	t.Helper()
	std := []*descriptorpb.FileDescriptorProto{
		{
			Name:    proto.String("google/protobuf/empty.proto"),
			Package: proto.String("google.protobuf"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{Name: proto.String("Empty")},
			},
		},
		{
			Name:    proto.String("google/protobuf/field_mask.proto"),
			Package: proto.String("google.protobuf"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{Name: proto.String("FieldMask")},
			},
		},
		{
			Name:    proto.String("google/longrunning/operations.proto"),
			Package: proto.String("google.longrunning"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{Name: proto.String("Operation")},
			},
		},
	}
	fileset, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: append(std, files...),
	})
	if err != nil {
		t.Fatalf("protodesc.NewFiles: %v", err)
	}
	return &Proto{files: fileset}
}

func mustResourceAPI(t *testing.T, p *Proto, fullName protoreflect.FullName) *ResourceAPI {
	t.Helper()
	api, err := p.ResourceAPIByName(fullName)
	if err != nil {
		t.Fatalf("ResourceAPIByName(%q): %v", fullName, err)
	}
	return api
}

func testFile(name, pkg string) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:    proto.String(name),
		Package: proto.String(pkg),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"google/protobuf/empty.proto",
			"google/protobuf/field_mask.proto",
			"google/longrunning/operations.proto",
		},
	}
}

func method(name, input, output string, opts *descriptorpb.MethodOptions) *descriptorpb.MethodDescriptorProto {
	return &descriptorpb.MethodDescriptorProto{
		Name:       proto.String(name),
		InputType:  proto.String(input),
		OutputType: proto.String(output),
		Options:    opts,
	}
}

func withOperationInfo(responseType string) *descriptorpb.MethodOptions {
	var info []byte
	info = protowire.AppendTag(info, operationInfoResponseType, protowire.BytesType)
	info = protowire.AppendString(info, responseType)

	var raw []byte
	raw = protowire.AppendTag(raw, operationInfoField, protowire.BytesType)
	raw = protowire.AppendBytes(raw, info)

	opts := &descriptorpb.MethodOptions{}
	if err := proto.Unmarshal(raw, opts); err != nil {
		panic(err)
	}
	return opts
}

func message(name string, fields ...*descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
	return &descriptorpb.DescriptorProto{
		Name:  proto.String(name),
		Field: fields,
	}
}

func messageField(name string, num int32, typeName string) *descriptorpb.FieldDescriptorProto {
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(num),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String(typeName),
	}
}

func stringField(name string, num int32) *descriptorpb.FieldDescriptorProto {
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(num),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
	}
}
