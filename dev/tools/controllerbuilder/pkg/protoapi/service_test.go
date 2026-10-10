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
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const testPkg = "google.cloud.test.v1"

func TestResourceAPIForStandardResource(t *testing.T) {
	// Arrange
	rd := &annotations.ResourceDescriptor{
		Type:    "test.googleapis.com/Repository",
		Pattern: []string{"projects/{project}/locations/{location}/repositories/{repository}"},
	}
	repo := message("Repository", stringField("name"), stringField("etag"))
	addLabels(repo)
	repo.Options = &descriptorpb.MessageOptions{}
	proto.SetExtension(repo.Options, annotations.E_Resource, rd)

	svcOpts := &descriptorpb.ServiceOptions{}
	proto.SetExtension(svcOpts, annotations.E_DefaultHost, "test.googleapis.com")
	fd := testFile("repository.proto", testPkg,
		repo,
		message("OperationMetadata"),
		message("GetRepositoryRequest", stringField("name")),
		message("CreateRepositoryRequest", stringField("parent"), stringField("repository_id"),
			messageField("repository", testPkg+".Repository"), stringField("request_id"), boolField("validate_only")),
		message("UpdateRepositoryRequest", messageField("repository", testPkg+".Repository"),
			messageField("update_mask", "google.protobuf.FieldMask"), boolField("allow_missing")),
		message("DeleteRepositoryRequest", stringField("name"), stringField("etag")),
		message("ListRepositoriesRequest", stringField("parent")),
		message("ListRepositoriesResponse", repeatedField("repositories", testPkg+".Repository"), stringField("next_page_token")),
	)
	fd.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name:    proto.String("RepositoryService"),
		Options: svcOpts,
		Method: []*descriptorpb.MethodDescriptorProto{
			method("GetRepository", "GetRepositoryRequest", testPkg+".Repository",
				httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/{name=projects/*/locations/*/repositories/*}"}}, "name")),
			method("CreateRepository", "CreateRepositoryRequest", "google.longrunning.Operation",
				withOperationInfo(httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Post{Post: "/v1/{parent=projects/*/locations/*}/repositories"}, Body: "repository"}, "parent,repository,repository_id"),
					"Repository", "OperationMetadata")),
			method("UpdateRepository", "UpdateRepositoryRequest", "google.longrunning.Operation",
				withOperationInfo(httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Patch{Patch: "/v1/{repository.name=projects/*/locations/*/repositories/*}"}, Body: "repository"}, "repository,update_mask"),
					"Repository", "OperationMetadata")),
			method("DeleteRepository", "DeleteRepositoryRequest", "google.longrunning.Operation",
				withOperationInfo(httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Delete{Delete: "/v1/{name=projects/*/locations/*/repositories/*}"}}, "name"),
					"google.protobuf.Empty", "OperationMetadata")),
			method("ListRepositories", "ListRepositoriesRequest", testPkg+".ListRepositoriesResponse",
				httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/{parent=projects/*/locations/*}/repositories"}}, "parent")),
		},
	}}
	p := newTestProto(t, fd)

	// Act
	api, err := p.ResourceAPIByName(testPkg + ".Repository")
	if err != nil {
		t.Fatalf("ResourceAPIByName: %v", err)
	}

	// Assert
	if api.Style != StyleAIP {
		t.Errorf("Style = %q, want %q", api.Style, StyleAIP)
	}
	if got := api.Service.FullName(); got != testPkg+".RepositoryService" {
		t.Errorf("Service = %q, want %q", got, testPkg+".RepositoryService")
	}
	if api.DefaultHost != "test.googleapis.com" {
		t.Errorf("DefaultHost = %q, want %q", api.DefaultHost, "test.googleapis.com")
	}
	if api.Metadata == nil || api.Metadata.Pattern != rd.Pattern[0] {
		t.Errorf("Metadata = %+v, want pattern %q", api.Metadata, rd.Pattern[0])
	}
	if !api.HasEtag || !api.HasLabels || api.ServerGeneratedID {
		t.Errorf("HasEtag, HasLabels, ServerGeneratedID = %v, %v, %v, want true, true, false", api.HasEtag, api.HasLabels, api.ServerGeneratedID)
	}
	operationInfo := &LROInfo{ResponseType: testPkg + ".Repository", MetadataType: testPkg + ".OperationMetadata"}
	want := map[string]*methodSummary{
		"get": {
			Name: "GetRepository", HTTPVerb: "GET", HTTPPath: "/v1/{name=projects/*/locations/*/repositories/*}",
			NameField: "name", MethodSignatures: []string{"name"}, Response: testPkg + ".Repository",
		},
		"create": {
			Name: "CreateRepository", HTTPVerb: "POST", HTTPPath: "/v1/{parent=projects/*/locations/*}/repositories",
			ParentField: "parent", ResourceField: "repository", IDField: "repository_id",
			ValidateOnlyField: "validate_only", RequestIDField: "request_id",
			MethodSignatures: []string{"parent,repository,repository_id"},
			LRO:              operationInfo, Response: "google.longrunning.Operation",
		},
		"update": {
			Name: "UpdateRepository", HTTPVerb: "PATCH", HTTPPath: "/v1/{repository.name=projects/*/locations/*/repositories/*}",
			ResourceField: "repository", UpdateMaskField: "update_mask", AllowMissingField: "allow_missing",
			MethodSignatures: []string{"repository,update_mask"},
			LRO:              operationInfo, Response: "google.longrunning.Operation",
		},
		"delete": {
			Name: "DeleteRepository", HTTPVerb: "DELETE", HTTPPath: "/v1/{name=projects/*/locations/*/repositories/*}",
			NameField: "name", EtagField: "etag", MethodSignatures: []string{"name"},
			LRO:      &LROInfo{ResponseType: "google.protobuf.Empty", MetadataType: testPkg + ".OperationMetadata"},
			Response: "google.longrunning.Operation",
		},
		"list": {
			Name: "ListRepositories", HTTPVerb: "GET", HTTPPath: "/v1/{parent=projects/*/locations/*}/repositories",
			ParentField: "parent", MethodSignatures: []string{"parent"}, Response: testPkg + ".ListRepositoriesResponse",
		},
	}
	if diff := cmp.Diff(want, summarizeMethods(api)); diff != "" {
		t.Errorf("standard methods mismatch (-want +got):\n%s", diff)
	}
}

func TestOperationInfo(t *testing.T) {
	for _, tc := range []struct {
		name    string
		output  string
		options *descriptorpb.MethodOptions
		want    *LROInfo
	}{
		{
			name:    "short names are resolved in the method's package",
			output:  "google.longrunning.Operation",
			options: withOperationInfo(&descriptorpb.MethodOptions{}, "Thing", "OperationMetadata"),
			want:    &LROInfo{ResponseType: testPkg + ".Thing", MetadataType: testPkg + ".OperationMetadata"},
		},
		{
			name:    "fully-qualified names are kept",
			output:  "google.longrunning.Operation",
			options: withOperationInfo(&descriptorpb.MethodOptions{}, "google.protobuf.Empty", "google.cloud.common.OperationMetadata"),
			want:    &LROInfo{ResponseType: "google.protobuf.Empty", MetadataType: "google.cloud.common.OperationMetadata"},
		},
		{
			name:    "a leading dot is dropped",
			output:  "google.longrunning.Operation",
			options: withOperationInfo(&descriptorpb.MethodOptions{}, ".google.protobuf.Empty", ""),
			want:    &LROInfo{ResponseType: "google.protobuf.Empty"},
		},
		{
			name:    "an operation without operation_info has an empty LROInfo",
			output:  "google.longrunning.Operation",
			options: &descriptorpb.MethodOptions{},
			want:    &LROInfo{},
		},
		{
			name:    "a method that does not return an operation has no LROInfo",
			output:  testPkg + ".Thing",
			options: withOperationInfo(&descriptorpb.MethodOptions{}, "Thing", "OperationMetadata"),
			want:    nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			fd := testFile("thing.proto", testPkg,
				message("Thing", stringField("name")),
				message("CreateThingRequest", stringField("parent"), messageField("thing", testPkg+".Thing")),
			)
			fd.Service = []*descriptorpb.ServiceDescriptorProto{{
				Name:   proto.String("ThingService"),
				Method: []*descriptorpb.MethodDescriptorProto{method("CreateThing", "CreateThingRequest", tc.output, tc.options)},
			}}
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, testPkg+".Thing")

			// Assert
			if diff := cmp.Diff(tc.want, api.Create.LRO); diff != "" {
				t.Errorf("LRO mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// If some import ever registers google.longrunning.operation_info, the option
// is parsed as an extension field instead of unknown bytes. It must still be
// read.
func TestOperationInfoAsRegisteredExtension(t *testing.T) {
	// Arrange
	lroFile := &descriptorpb.FileDescriptorProto{
		Name:       proto.String("test/operation_info.proto"),
		Package:    proto.String("test.longrunning"),
		Syntax:     proto.String("proto3"),
		Dependency: []string{"google/protobuf/descriptor.proto"},
		MessageType: []*descriptorpb.DescriptorProto{
			message("OperationInfo", stringField("response_type"), stringField("metadata_type")),
		},
		Extension: []*descriptorpb.FieldDescriptorProto{{
			Name:     proto.String("operation_info"),
			Number:   proto.Int32(int32(operationInfoField)),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			TypeName: proto.String(".test.longrunning.OperationInfo"),
			Extendee: proto.String(".google.protobuf.MethodOptions"),
		}},
	}
	extFile, err := protodesc.NewFile(lroFile, protoregistry.GlobalFiles)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}
	info := dynamicpb.NewMessage(extFile.Messages().ByName("OperationInfo"))
	info.Set(info.Descriptor().Fields().ByName("response_type"), protoreflect.ValueOfString("Thing"))
	info.Set(info.Descriptor().Fields().ByName("metadata_type"), protoreflect.ValueOfString("OperationMetadata"))
	opts := &descriptorpb.MethodOptions{}
	proto.SetExtension(opts, dynamicpb.NewExtensionType(extFile.Extensions().Get(0)), info)

	fd := testFile("thing.proto", testPkg,
		message("Thing", stringField("name")),
		message("CreateThingRequest", stringField("parent"), messageField("thing", testPkg+".Thing")),
	)
	fd.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name:   proto.String("ThingService"),
		Method: []*descriptorpb.MethodDescriptorProto{method("CreateThing", "CreateThingRequest", "google.longrunning.Operation", opts)},
	}}
	p := newTestProto(t, fd)

	// Act
	api := mustResourceAPI(t, p, testPkg+".Thing")

	// Assert
	if unknown := api.Create.Descriptor.Options().ProtoReflect().GetUnknown(); len(unknown) != 0 {
		t.Fatalf("the option is held as unknown bytes, so this test does not cover a registered extension")
	}
	want := &LROInfo{ResponseType: testPkg + ".Thing", MetadataType: testPkg + ".OperationMetadata"}
	if diff := cmp.Diff(want, api.Create.LRO); diff != "" {
		t.Errorf("LRO mismatch (-want +got):\n%s", diff)
	}
}

func TestHTTPBinding(t *testing.T) {
	for _, tc := range []struct {
		name     string
		rule     *annotations.HttpRule
		wantVerb string
		wantPath string
	}{
		{"get", &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/v1/{name=things/*}"}}, "GET", "/v1/{name=things/*}"},
		{"post", &annotations.HttpRule{Pattern: &annotations.HttpRule_Post{Post: "/v1/things"}}, "POST", "/v1/things"},
		{"put", &annotations.HttpRule{Pattern: &annotations.HttpRule_Put{Put: "/v1/{thing.name=things/*}"}}, "PUT", "/v1/{thing.name=things/*}"},
		{"patch", &annotations.HttpRule{Pattern: &annotations.HttpRule_Patch{Patch: "/v1/{thing.name=things/*}"}}, "PATCH", "/v1/{thing.name=things/*}"},
		{"delete", &annotations.HttpRule{Pattern: &annotations.HttpRule_Delete{Delete: "/v1/{name=things/*}"}}, "DELETE", "/v1/{name=things/*}"},
		{"custom", &annotations.HttpRule{Pattern: &annotations.HttpRule_Custom{Custom: &annotations.CustomHttpPattern{Kind: "HEAD", Path: "/v1/{name=things/*}"}}}, "CUSTOM", "/v1/{name=things/*}"},
		{"no binding", nil, "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			opts := &descriptorpb.MethodOptions{}
			if tc.rule != nil {
				opts = httpOptions(tc.rule)
			}
			fd := testFile("thing.proto", testPkg,
				message("Thing", stringField("name")),
				message("GetThingRequest", stringField("name")),
			)
			fd.Service = []*descriptorpb.ServiceDescriptorProto{{
				Name:   proto.String("ThingService"),
				Method: []*descriptorpb.MethodDescriptorProto{method("GetThing", "GetThingRequest", testPkg+".Thing", opts)},
			}}
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, testPkg+".Thing")

			// Assert
			if api.Get.HTTPVerb != tc.wantVerb || api.Get.HTTPPath != tc.wantPath {
				t.Errorf("HTTP binding = %q %q, want %q %q", api.Get.HTTPVerb, api.Get.HTTPPath, tc.wantVerb, tc.wantPath)
			}
		})
	}
}

func TestClientIDField(t *testing.T) {
	for _, tc := range []struct {
		name     string
		resource string
		fields   []*descriptorpb.FieldDescriptorProto
		want     string
	}{
		{
			name:     "resource_id",
			resource: "Thing",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("request_id"), stringField("other_id"), stringField("thing_id")},
			want:     "thing_id",
		},
		{
			name:     "multi-word and acronym message names are snake_cased",
			resource: "SACRealm",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("sac_realm_id")},
			want:     "sac_realm_id",
		},
		{
			name:     "falls back to the last words of the resource name",
			resource: "NotificationConfig",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("project_id"), stringField("config_id")},
			want:     "config_id",
		},
		{
			name:     "the longest word-boundary suffix wins",
			resource: "BuildTriggerConfig",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("config_id"), stringField("trigger_config_id")},
			want:     "trigger_config_id",
		},
		{
			name:     "a suffix that is not a whole word does not count",
			resource: "Subnetwork",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("network_id")},
			want:     "",
		},
		{
			name:     "a suffix field must be a string",
			resource: "NotificationConfig",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("parent"), int64Field("config_id")},
			want:     "",
		},
		{
			name:     "other fields ending in _id name something else",
			resource: "Cluster",
			fields:   []*descriptorpb.FieldDescriptorProto{stringField("project_id"), stringField("external_id"), stringField("request_id")},
			want:     "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			fields := append(tc.fields, messageField(snakeCase(tc.resource), testPkg+"."+tc.resource))
			fd := testFile("thing.proto", testPkg,
				message(tc.resource, stringField("name")),
				message("CreateRequest", fields...),
			)
			fd.Service = []*descriptorpb.ServiceDescriptorProto{{
				Name:   proto.String("ThingService"),
				Method: []*descriptorpb.MethodDescriptorProto{method("Create"+tc.resource, "CreateRequest", testPkg+"."+tc.resource, nil)},
			}}
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, protoreflect.FullName(testPkg+"."+tc.resource))

			// Assert
			if api.Create.IDField != tc.want {
				t.Errorf("IDField = %q, want %q", api.Create.IDField, tc.want)
			}
			if api.ServerGeneratedID != (tc.want == "") {
				t.Errorf("ServerGeneratedID = %v, want %v", api.ServerGeneratedID, tc.want == "")
			}
		})
	}
}

func TestUpdateMaskAndEtagDetection(t *testing.T) {
	// Arrange
	fd := testFile("thing.proto", testPkg, message("Thing", stringField("name")))
	svc := &descriptorpb.ServiceDescriptorProto{Name: proto.String("ThingService")}
	fd.Service = []*descriptorpb.ServiceDescriptorProto{svc}
	addAIPResource(fd, svc, "Masked", crud{get: true, create: true, createID: true, update: true, updateMask: true, delete: true, deleteEtag: true})
	addAIPResource(fd, svc, "Unmasked", crud{get: true, create: true, createID: true, update: true, delete: true})
	p := newTestProto(t, fd)

	// Act
	masked := mustResourceAPI(t, p, testPkg+".Masked")
	unmasked := mustResourceAPI(t, p, testPkg+".Unmasked")

	// Assert
	if masked.Update.UpdateMaskField != "update_mask" || masked.Delete.EtagField != "etag" {
		t.Errorf("Masked: UpdateMaskField, EtagField = %q, %q, want update_mask, etag", masked.Update.UpdateMaskField, masked.Delete.EtagField)
	}
	if unmasked.Update.UpdateMaskField != "" || unmasked.Delete.EtagField != "" {
		t.Errorf("Unmasked: UpdateMaskField, EtagField = %q, %q, want empty", unmasked.Update.UpdateMaskField, unmasked.Delete.EtagField)
	}
	if masked.HasEtag || masked.HasLabels {
		t.Errorf("HasEtag, HasLabels = %v, %v, want false, false", masked.HasEtag, masked.HasLabels)
	}
}

func TestServiceChoice(t *testing.T) {
	for _, tc := range []struct {
		name    string
		methods map[string]crud
		want    string
	}{
		{
			name: "the service with the most standard methods wins",
			methods: map[string]crud{
				"AService": {get: true},
				"BService": {get: true, create: true, update: true, delete: true},
			},
			want: testPkg + ".BService",
		},
		{
			name: "a tie goes to the service with the shorter name",
			methods: map[string]crud{
				"RegionThings": {get: true, create: true},
				"Things":       {get: true, delete: true},
			},
			want: testPkg + ".Things",
		},
		{
			name: "a tie between names of the same length goes to the one that sorts first",
			methods: map[string]crud{
				"BService": {get: true, create: true},
				"AService": {get: true, delete: true},
			},
			want: testPkg + ".AService",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			fd := testFile("thing.proto", testPkg)
			for _, name := range slices.Sorted(maps.Keys(tc.methods)) {
				svc := &descriptorpb.ServiceDescriptorProto{Name: proto.String(name)}
				fd.Service = append(fd.Service, svc)
				addAIPMethods(svc, "Thing", tc.methods[name])
			}
			addAIPMessages(fd, "Thing", crud{get: true, create: true, update: true, delete: true, list: true})
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, testPkg+".Thing")

			// Assert
			if got := string(api.Service.FullName()); got != tc.want {
				t.Errorf("Service = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestAIPRequestShape(t *testing.T) {
	for _, tc := range []struct {
		name string
		verb string
		// request holds the request's fields; nil means the method takes the
		// resource itself as its request.
		request         []*descriptorpb.FieldDescriptorProto
		wantNonstandard bool
		// wantServerGeneratedID is only checked for Create.
		wantServerGeneratedID bool
	}{
		{
			name:    "Get takes a name",
			verb:    "Get",
			request: []*descriptorpb.FieldDescriptorProto{stringField("name")},
		},
		{
			name:            "a Get without a name field is flagged",
			verb:            "Get",
			request:         []*descriptorpb.FieldDescriptorProto{stringField("project_id"), stringField("region"), stringField("thing_name")},
			wantNonstandard: true,
		},
		{
			name:    "Delete takes a name",
			verb:    "Delete",
			request: []*descriptorpb.FieldDescriptorProto{stringField("name"), stringField("etag")},
		},
		{
			name:            "a Delete without a name field is flagged",
			verb:            "Delete",
			request:         []*descriptorpb.FieldDescriptorProto{stringField("thing")},
			wantNonstandard: true,
		},
		{
			name:    "Create takes a parent",
			verb:    "Create",
			request: []*descriptorpb.FieldDescriptorProto{stringField("parent"), stringField("thing_id"), messageField("thing", testPkg+".Thing")},
		},
		{
			name:                  "a top-level Create takes the resource and no parent",
			verb:                  "Create",
			request:               []*descriptorpb.FieldDescriptorProto{messageField("thing", testPkg+".Thing"), boolField("validate_only")},
			wantServerGeneratedID: true,
		},
		{
			// It has no ID field either, but being flagged, it says nothing
			// about who names the resource.
			name:            "a Create whose request is the resource itself is flagged",
			verb:            "Create",
			request:         nil,
			wantNonstandard: true,
		},
		{
			name:            "a Create with neither a parent nor the resource is flagged",
			verb:            "Create",
			request:         []*descriptorpb.FieldDescriptorProto{stringField("name"), stringField("source")},
			wantNonstandard: true,
		},
		{
			name:    "Update takes the resource",
			verb:    "Update",
			request: []*descriptorpb.FieldDescriptorProto{messageField("thing", testPkg+".Thing"), messageField("update_mask", "google.protobuf.FieldMask")},
		},
		{
			name:            "an Update without the resource is flagged",
			verb:            "Update",
			request:         []*descriptorpb.FieldDescriptorProto{stringField("name"), messageField("update_mask", "google.protobuf.FieldMask")},
			wantNonstandard: true,
		},
		{
			name:    "a List request is not checked",
			verb:    "List",
			request: []*descriptorpb.FieldDescriptorProto{stringField("project_id")},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			fd := testFile("thing.proto", testPkg, message("Thing", stringField("name")))
			input := "Thing"
			if tc.request != nil {
				fd.MessageType = append(fd.MessageType, message("Request", tc.request...))
				input = "Request"
			}
			fd.Service = []*descriptorpb.ServiceDescriptorProto{{
				Name:   proto.String("ThingService"),
				Method: []*descriptorpb.MethodDescriptorProto{method(tc.verb+"Thing", input, testPkg+".Thing", nil)},
			}}
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, testPkg+".Thing")

			// Assert
			got := summarizeMethods(api)[strings.ToLower(tc.verb)]
			if got == nil {
				t.Fatalf("%sThing is not the AIP %s, want it found by its name whatever its request", tc.verb, tc.verb)
			}
			if got.NonstandardRequest != tc.wantNonstandard {
				t.Errorf("%sThing NonstandardRequest = %v, want %v", tc.verb, got.NonstandardRequest, tc.wantNonstandard)
			}
			if tc.verb == "Create" && api.ServerGeneratedID != tc.wantServerGeneratedID {
				t.Errorf("ServerGeneratedID = %v, want %v", api.ServerGeneratedID, tc.wantServerGeneratedID)
			}
		})
	}
}

// A flagged method still counts toward the service choice: the service is
// where the resource's methods are, even when they are not AIP-shaped.
func TestServiceChoiceCountsFlaggedMethods(t *testing.T) {
	// Arrange
	fd := testFile("thing.proto", testPkg,
		message("Thing", stringField("name")),
		message("GetThingRequest", stringField("name")),
		message("ThingByIDRequest", stringField("project_id"), stringField("thing_id")),
	)
	fd.Service = []*descriptorpb.ServiceDescriptorProto{
		{
			Name:   proto.String("AService"),
			Method: []*descriptorpb.MethodDescriptorProto{method("GetThing", "GetThingRequest", testPkg+".Thing", nil)},
		},
		{
			Name: proto.String("BService"),
			Method: []*descriptorpb.MethodDescriptorProto{
				method("GetThing", "ThingByIDRequest", testPkg+".Thing", nil),
				method("DeleteThing", "ThingByIDRequest", "google.protobuf.Empty", nil),
			},
		},
	}
	p := newTestProto(t, fd)

	// Act
	api := mustResourceAPI(t, p, testPkg+".Thing")

	// Assert
	if got := string(api.Service.FullName()); got != testPkg+".BService" {
		t.Errorf("Service = %q, want %q", got, testPkg+".BService")
	}
	if api.Get == nil || !api.Get.NonstandardRequest || api.Delete == nil || !api.Delete.NonstandardRequest {
		got := summarizeMethods(api)
		t.Errorf("Get, Delete = %+v, %+v, want both flagged NonstandardRequest", got["get"], got["delete"])
	}
}

func TestListMethod(t *testing.T) {
	// Arrange
	fd := testFile("policy.proto", testPkg,
		message("Policy", stringField("name")),
		message("ListPoliciesRequest", stringField("parent")),
		message("ListPoliciesResponse", repeatedField("policies", testPkg+".Policy")),
		message("ListPolicyRevisionsResponse", repeatedField("policies", testPkg+".Policy")),
		message("Widget", stringField("name")),
		message("ListWidgetsSummaryResponse", stringField("summary")),
	)
	fd.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name: proto.String("PolicyService"),
		Method: []*descriptorpb.MethodDescriptorProto{
			method("ListPolicyRevisions", "ListPoliciesRequest", testPkg+".ListPolicyRevisionsResponse", nil),
			method("ListPolicies", "ListPoliciesRequest", testPkg+".ListPoliciesResponse", nil),
			method("ListWidgetsSummary", "ListPoliciesRequest", testPkg+".ListWidgetsSummaryResponse", nil),
		},
	}}
	p := newTestProto(t, fd)

	// Act
	policy := mustResourceAPI(t, p, testPkg+".Policy")
	widget := mustResourceAPI(t, p, testPkg+".Widget")

	// Assert
	// ListPolicies does not start with List + "Policy", but its response
	// holds policies; it beats the longer ListPolicyRevisions.
	if policy.List == nil || policy.List.Descriptor.Name() != "ListPolicies" {
		t.Errorf("Policy List = %v, want ListPolicies", policy.List)
	}
	// With no response match, the name alone is enough.
	if widget.List == nil || widget.List.Descriptor.Name() != "ListWidgetsSummary" {
		t.Errorf("Widget List = %v, want ListWidgetsSummary", widget.List)
	}
}

func TestComputeStyle(t *testing.T) {
	for _, tc := range []struct {
		name       string
		updateName string
		bodyType   string
		wantStyle  APIStyle
		wantUpdate string
	}{
		{name: "Patch is the update", updateName: "Patch", bodyType: "Address", wantStyle: StyleCompute, wantUpdate: "Patch"},
		{name: "Update when there is no Patch", updateName: "Update", bodyType: "Address", wantStyle: StyleCompute, wantUpdate: "Update"},
		// The methods are not found by their AIP names either: Get is not
		// GetAddress, and a bare List is not an AIP-132 List.
		{name: "a body of another type is not compute style", updateName: "Patch", bodyType: "Other", wantStyle: StyleNone},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			const pkg = testComputePkg
			svcOpts := &descriptorpb.ServiceOptions{}
			proto.SetExtension(svcOpts, annotations.E_DefaultHost, "compute.googleapis.com")
			fd := testFile("compute.proto", pkg,
				message("Address", stringField("name"), stringField("region")),
				message("Other", stringField("name")),
				message("Operation", stringField("name")),
				message("AddressList", repeatedField("items", pkg+".Address")),
				message("GetAddressRequest", stringField("address"), stringField("project"), stringField("region")),
				message("InsertAddressRequest", messageField("address_resource", pkg+"."+tc.bodyType), stringField("project"), stringField("region"), stringField("request_id")),
				message("PatchAddressRequest", stringField("address"), messageField("address_resource", pkg+"."+tc.bodyType), stringField("project"), stringField("region"), stringField("request_id")),
				message("DeleteAddressRequest", stringField("address"), stringField("project"), stringField("region"), stringField("request_id")),
				message("ListAddressesRequest", stringField("project"), stringField("region")),
			)
			fd.Service = []*descriptorpb.ServiceDescriptorProto{{
				Name:    proto.String("Addresses"),
				Options: svcOpts,
				Method: []*descriptorpb.MethodDescriptorProto{
					method("Get", pkg+".GetAddressRequest", pkg+".Address", httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/compute/v1/projects/{project}/regions/{region}/addresses/{address}"}})),
					method("Insert", pkg+".InsertAddressRequest", pkg+".Operation", httpOptions(&annotations.HttpRule{Pattern: &annotations.HttpRule_Post{Post: "/compute/v1/projects/{project}/regions/{region}/addresses"}})),
					method(tc.updateName, pkg+".PatchAddressRequest", pkg+".Operation", nil),
					method("Delete", pkg+".DeleteAddressRequest", pkg+".Operation", nil),
					method("List", pkg+".ListAddressesRequest", pkg+".AddressList", nil),
				},
			}}
			p := newTestProto(t, fd)

			// Act
			api := mustResourceAPI(t, p, pkg+".Address")

			// Assert
			if api.Style != tc.wantStyle {
				t.Fatalf("Style = %q, want %q", api.Style, tc.wantStyle)
			}
			if tc.wantStyle != StyleCompute {
				return
			}
			got := summarizeMethods(api)
			want := map[string]string{"get": "Get", "create": "Insert", "update": tc.wantUpdate, "delete": "Delete", "list": "List"}
			for verb, name := range want {
				if got[verb] == nil || got[verb].Name != name {
					t.Errorf("%s = %+v, want %s", verb, got[verb], name)
				}
			}
			if api.Create.ResourceField != "address_resource" || api.Create.RequestIDField != "request_id" || api.Create.HTTPVerb != "POST" {
				t.Errorf("Insert = %+v, want address_resource body, request_id and POST", got["create"])
			}
			if api.Create.IDField != "" || api.ServerGeneratedID || api.Create.LRO != nil {
				t.Errorf("IDField, ServerGeneratedID, LRO = %q, %v, %v, want empty, false, nil", api.Create.IDField, api.ServerGeneratedID, api.Create.LRO)
			}
			if api.DefaultHost != "compute.googleapis.com" {
				t.Errorf("DefaultHost = %q, want compute.googleapis.com", api.DefaultHost)
			}
		})
	}
}

func TestDefaultHostAbsent(t *testing.T) {
	// Arrange
	fd := testFile("thing.proto", testPkg)
	svc := &descriptorpb.ServiceDescriptorProto{Name: proto.String("ThingService")}
	fd.Service = []*descriptorpb.ServiceDescriptorProto{svc}
	addAIPResource(fd, svc, "Thing", crud{get: true})
	p := newTestProto(t, fd)

	// Act
	api := mustResourceAPI(t, p, testPkg+".Thing")

	// Assert
	if api.DefaultHost != "" {
		t.Errorf("DefaultHost = %q, want empty", api.DefaultHost)
	}
}

func TestResourceAPIByNameNotFound(t *testing.T) {
	// Arrange
	fd := testFile("thing.proto", testPkg, message("Thing", stringField("name")))
	fd.Service = []*descriptorpb.ServiceDescriptorProto{{Name: proto.String("ThingService")}}
	p := newTestProto(t, fd)

	for _, name := range []protoreflect.FullName{testPkg + ".Missing", testPkg + ".ThingService"} {
		t.Run(string(name), func(t *testing.T) {
			// Act
			api, err := p.ResourceAPIByName(name)

			// Assert
			if api != nil || !errors.Is(err, protoregistry.NotFound) {
				t.Errorf("ResourceAPIByName(%q) = %v, %v; want nil and an error wrapping protoregistry.NotFound", name, api, err)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{"Repository", "repository"},
		{"TlsInspectionPolicy", "tls_inspection_policy"},
		{"SACRealm", "sac_realm"},
		{"HTTPHealthCheck", "http_health_check"},
		{"Ipv6Range", "ipv6_range"},
		{"URL", "url"},
	} {
		if got := snakeCase(tc.in); got != tc.want {
			t.Errorf("snakeCase(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// methodSummary is the comparable part of a StandardMethod.
type methodSummary struct {
	Name               string
	NonstandardRequest bool
	HTTPVerb           string
	HTTPPath           string
	ParentField        string
	NameField          string
	ResourceField      string
	IDField            string
	UpdateMaskField    string
	EtagField          string
	AllowMissingField  string
	ValidateOnlyField  string
	RequestIDField     string
	MethodSignatures   []string
	LRO                *LROInfo
	Response           string
}

func summarizeMethods(api *ResourceAPI) map[string]*methodSummary {
	out := map[string]*methodSummary{}
	for verb, m := range map[string]*StandardMethod{"get": api.Get, "create": api.Create, "update": api.Update, "delete": api.Delete, "list": api.List} {
		if m == nil {
			continue
		}
		out[verb] = &methodSummary{
			Name:               string(m.Descriptor.Name()),
			NonstandardRequest: m.NonstandardRequest,
			HTTPVerb:           m.HTTPVerb,
			HTTPPath:           m.HTTPPath,
			ParentField:        m.ParentField,
			NameField:          m.NameField,
			ResourceField:      m.ResourceField,
			IDField:            m.IDField,
			UpdateMaskField:    m.UpdateMaskField,
			EtagField:          m.EtagField,
			AllowMissingField:  m.AllowMissingField,
			ValidateOnlyField:  m.ValidateOnlyField,
			RequestIDField:     m.RequestIDField,
			MethodSignatures:   m.MethodSignatures,
			LRO:                m.LRO,
			Response:           m.Response,
		}
	}
	return out
}

// newTestProto builds a Proto from files, plus the well-known files they may
// import: google/protobuf/empty.proto, google/protobuf/field_mask.proto and a
// minimal google/longrunning/operations.proto.
func newTestProto(t *testing.T, files ...*descriptorpb.FileDescriptorProto) *Proto {
	t.Helper()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{
		protodesc.ToFileDescriptorProto(emptypb.File_google_protobuf_empty_proto),
		protodesc.ToFileDescriptorProto(fieldmaskpb.File_google_protobuf_field_mask_proto),
		{
			Name:        proto.String("google/longrunning/operations.proto"),
			Package:     proto.String("google.longrunning"),
			Syntax:      proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{message("Operation", stringField("name"))},
		},
	}}
	fds.File = append(fds.File, files...)
	reg, err := protodesc.NewFiles(fds)
	if err != nil {
		t.Fatalf("protodesc.NewFiles: %v", err)
	}
	return &Proto{files: reg}
}

func mustResourceAPI(t *testing.T, p *Proto, name protoreflect.FullName) *ResourceAPI {
	t.Helper()
	api, err := p.ResourceAPIByName(name)
	if err != nil {
		t.Fatalf("ResourceAPIByName(%q): %v", name, err)
	}
	return api
}

// testFile returns a proto3 file in package pkg holding msgs. It imports the
// files newTestProto provides.
func testFile(name, pkg string, msgs ...*descriptorpb.DescriptorProto) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:        proto.String(name),
		Package:     proto.String(pkg),
		Syntax:      proto.String("proto3"),
		Dependency:  []string{"google/protobuf/empty.proto", "google/protobuf/field_mask.proto", "google/longrunning/operations.proto"},
		MessageType: msgs,
	}
}

// message returns a message whose fields are numbered in order.
func message(name string, fields ...*descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
	for i, f := range fields {
		f.Number = proto.Int32(int32(i + 1))
	}
	return &descriptorpb.DescriptorProto{Name: proto.String(name), Field: fields}
}

func scalarField(name string, typ descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	return &descriptorpb.FieldDescriptorProto{
		Name:  proto.String(name),
		Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Type:  typ.Enum(),
	}
}

func stringField(name string) *descriptorpb.FieldDescriptorProto {
	return scalarField(name, descriptorpb.FieldDescriptorProto_TYPE_STRING)
}

func boolField(name string) *descriptorpb.FieldDescriptorProto {
	return scalarField(name, descriptorpb.FieldDescriptorProto_TYPE_BOOL)
}

func int64Field(name string) *descriptorpb.FieldDescriptorProto {
	return scalarField(name, descriptorpb.FieldDescriptorProto_TYPE_INT64)
}

func messageField(name, fullName string) *descriptorpb.FieldDescriptorProto {
	f := scalarField(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	f.TypeName = proto.String("." + fullName)
	return f
}

func repeatedField(name, fullName string) *descriptorpb.FieldDescriptorProto {
	f := messageField(name, fullName)
	f.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	return f
}

// addLabels adds a map<string, string> labels field to msg.
func addLabels(msg *descriptorpb.DescriptorProto) {
	entry := message("LabelsEntry", stringField("key"), stringField("value"))
	entry.Options = &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)}
	msg.NestedType = append(msg.NestedType, entry)
	labels := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("labels"),
		Number:   proto.Int32(int32(len(msg.Field) + 1)),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String("." + testPkg + "." + msg.GetName() + ".LabelsEntry"),
	}
	msg.Field = append(msg.Field, labels)
}

// method returns a method. input is relative to testPkg unless it contains a
// dot; output must be fully qualified.
func method(name, input, output string, opts *descriptorpb.MethodOptions) *descriptorpb.MethodDescriptorProto {
	return &descriptorpb.MethodDescriptorProto{
		Name:       proto.String(name),
		InputType:  proto.String("." + qualifyTypeName(testPkg, input)),
		OutputType: proto.String("." + output),
		Options:    opts,
	}
}

func httpOptions(rule *annotations.HttpRule, signatures ...string) *descriptorpb.MethodOptions {
	opts := &descriptorpb.MethodOptions{}
	proto.SetExtension(opts, annotations.E_Http, rule)
	if len(signatures) > 0 {
		proto.SetExtension(opts, annotations.E_MethodSignature, signatures)
	}
	return opts
}

// withOperationInfo adds a google.longrunning.operation_info option to opts.
// The extension is not registered in this module, so, like the loader, it is
// stored as unknown bytes.
func withOperationInfo(opts *descriptorpb.MethodOptions, responseType, metadataType string) *descriptorpb.MethodOptions {
	var info []byte
	if responseType != "" {
		info = protowire.AppendTag(info, operationInfoResponseType, protowire.BytesType)
		info = protowire.AppendString(info, responseType)
	}
	if metadataType != "" {
		info = protowire.AppendTag(info, operationInfoMetadataType, protowire.BytesType)
		info = protowire.AppendString(info, metadataType)
	}
	b := opts.ProtoReflect().GetUnknown()
	b = protowire.AppendTag(b, operationInfoField, protowire.BytesType)
	b = protowire.AppendBytes(b, info)
	opts.ProtoReflect().SetUnknown(b)
	return opts
}

// crud selects the standard methods, and request fields, that addAIPResource
// declares for a synthetic resource.
type crud struct {
	get, create, update, delete, list bool
	// createID gives the Create request a <resource>_id field.
	createID bool
	// updateMask gives the Update request an update_mask field.
	updateMask bool
	// deleteEtag gives the Delete request an etag field.
	deleteEtag bool
}

// addAIPResource declares resource and the request messages c needs in fd,
// and the methods c selects in svc, with their AIP names.
func addAIPResource(fd *descriptorpb.FileDescriptorProto, svc *descriptorpb.ServiceDescriptorProto, resource string, c crud) {
	addAIPMessages(fd, resource, c)
	addAIPMethods(svc, resource, c)
}

func addAIPMessages(fd *descriptorpb.FileDescriptorProto, resource string, c crud) {
	full := testPkg + "." + resource
	body := messageField(snakeCase(resource), full)
	create := message("Create"+resource+"Request", stringField("parent"), body)
	if c.createID {
		create = message("Create"+resource+"Request", stringField("parent"), stringField(snakeCase(resource)+"_id"), messageField(snakeCase(resource), full))
	}
	update := message("Update"+resource+"Request", messageField(snakeCase(resource), full))
	if c.updateMask {
		update = message("Update"+resource+"Request", messageField(snakeCase(resource), full), messageField("update_mask", "google.protobuf.FieldMask"))
	}
	del := message("Delete"+resource+"Request", stringField("name"))
	if c.deleteEtag {
		del = message("Delete"+resource+"Request", stringField("name"), stringField("etag"))
	}
	fd.MessageType = append(fd.MessageType,
		message(resource, stringField("name")),
		message("Get"+resource+"Request", stringField("name")),
		create,
		update,
		del,
		message("List"+resource+"sRequest", stringField("parent")),
		message("List"+resource+"sResponse", repeatedField("items", full)),
	)
}

func addAIPMethods(svc *descriptorpb.ServiceDescriptorProto, resource string, c crud) {
	full := testPkg + "." + resource
	add := func(ok bool, name, output string) {
		if ok {
			svc.Method = append(svc.Method, method(name+resource, name+resource+"Request", output, nil))
		}
	}
	add(c.get, "Get", full)
	add(c.create, "Create", full)
	add(c.update, "Update", full)
	add(c.delete, "Delete", "google.protobuf.Empty")
	if c.list {
		svc.Method = append(svc.Method, method("List"+resource+"s", "List"+resource+"sRequest", testPkg+".List"+resource+"sResponse", nil))
	}
}
