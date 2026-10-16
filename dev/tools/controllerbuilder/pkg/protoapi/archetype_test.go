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
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestClassifyArchetype(t *testing.T) {
	msg := (&emptypb.Empty{}).ProtoReflect().Descriptor()
	get := &StandardMethod{}
	createWithID := &StandardMethod{IDField: "thing_id"}
	createWithoutID := &StandardMethod{}
	updateWithMask := &StandardMethod{UpdateMaskField: "update_mask"}
	updateWithoutMask := &StandardMethod{}
	del := &StandardMethod{}
	list := &StandardMethod{}
	nonstandardGet := &StandardMethod{NonstandardRequest: true}
	nonstandardCreate := &StandardMethod{NonstandardRequest: true}
	nonstandardUpdate := &StandardMethod{UpdateMaskField: "update_mask", NonstandardRequest: true}
	nonstandardDelete := &StandardMethod{NonstandardRequest: true}

	for _, tc := range []struct {
		name string
		api  *ResourceAPI
		want Archetype
	}{
		{
			name: "a failed lookup is unresolved",
			api:  nil,
			want: ArchetypeUnresolved,
		},
		{
			name: "a model without a message is unresolved",
			api:  &ResourceAPI{Style: StyleAIP, Get: get},
			want: ArchetypeUnresolved,
		},
		{
			name: "compute style wins over the shape of the methods",
			api:  &ResourceAPI{Message: msg, Style: StyleCompute, Get: get, Create: createWithoutID, Update: updateWithMask, Delete: del},
			want: ArchetypeComputeStyle,
		},
		{
			name: "compute style without update",
			api:  &ResourceAPI{Message: msg, Style: StyleCompute, Get: get, Create: createWithoutID, Delete: del},
			want: ArchetypeComputeStyle,
		},
		{
			name: "no standard methods",
			api:  &ResourceAPI{Message: msg, Style: StyleNone},
			want: ArchetypeNonstandard,
		},
		{
			name: "no get and no create, even with the other methods",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Update: updateWithMask, Delete: del, List: list},
			want: ArchetypeNonstandard,
		},
		{
			name: "a get with a nonstandard request, the rest standard",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: nonstandardGet, Create: createWithID, Update: updateWithMask, Delete: del, List: list},
			want: ArchetypeNonstandard,
		},
		{
			name: "a create with a nonstandard request is not no-update",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: nonstandardCreate, Delete: del, List: list},
			want: ArchetypeNonstandard,
		},
		{
			name: "an update with a nonstandard request is not singleton",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Update: nonstandardUpdate, List: list},
			want: ArchetypeNonstandard,
		},
		{
			name: "a delete with a nonstandard request is not update-without-mask",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Update: updateWithoutMask, Delete: nonstandardDelete},
			want: ArchetypeNonstandard,
		},
		{
			name: "get and update without create",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Update: updateWithMask},
			want: ArchetypeSingleton,
		},
		{
			name: "singleton beats the other shapes, whatever update and delete look like",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Update: updateWithoutMask, Delete: del, List: list},
			want: ArchetypeSingleton,
		},
		{
			name: "standard",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Update: updateWithMask, Delete: del, List: list},
			want: ArchetypeStandard,
		},
		{
			name: "standard does not need list",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Update: updateWithMask, Delete: del},
			want: ArchetypeStandard,
		},
		{
			name: "create without an id field",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithoutID, Update: updateWithMask, Delete: del, List: list},
			want: ArchetypeServerGeneratedID,
		},
		{
			name: "no update, with an id field",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Delete: del, List: list},
			want: ArchetypeNoUpdate,
		},
		{
			name: "no update, without an id field",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithoutID, Delete: del},
			want: ArchetypeNoUpdate,
		},
		{
			name: "update without a mask",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Update: updateWithoutMask, Delete: del},
			want: ArchetypeUpdateWithoutMask,
		},
		{
			name: "update without a mask, and create without an id field",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithoutID, Update: updateWithoutMask, Delete: del},
			want: ArchetypeUpdateWithoutMask,
		},
		{
			name: "no delete",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID, Update: updateWithMask, List: list},
			want: ArchetypeOther,
		},
		{
			name: "get and create only",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, Create: createWithID},
			want: ArchetypeOther,
		},
		{
			name: "get only",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Get: get, List: list},
			want: ArchetypeOther,
		},
		{
			name: "create, update and delete without get",
			api:  &ResourceAPI{Message: msg, Style: StyleAIP, Create: createWithID, Update: updateWithMask, Delete: del},
			want: ArchetypeOther,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := ClassifyArchetype(tc.api)

			// Assert
			if got != tc.want {
				t.Errorf("ClassifyArchetype() = %s (%s), want %s (%s)", got, got.Name(), tc.want, tc.want.Name())
			}
		})
	}
}

// TestClassifyArchetypeFromDescriptors runs the whole path, from descriptors
// through ResourceAPIFor to ClassifyArchetype, for one resource of each
// archetype, and for one whose Update is flagged NonstandardRequest.
func TestClassifyArchetypeFromDescriptors(t *testing.T) {
	// Arrange
	fd := testFile("things.proto", testPkg)
	svc := &descriptorpb.ServiceDescriptorProto{Name: proto.String("ThingService")}
	fd.Service = []*descriptorpb.ServiceDescriptorProto{svc}
	resources := map[string]crud{
		"Alpha":   {get: true, create: true, createID: true, update: true, updateMask: true, delete: true, list: true},
		"Bravo":   {get: true, create: true, createID: true, delete: true, list: true},
		"Charlie": {get: true, create: true, update: true, updateMask: true, delete: true, list: true},
		"Delta":   {get: true, update: true, updateMask: true, list: true},
		"Echo":    {list: true},
		"Foxtrot": {get: true, create: true, createID: true, update: true, delete: true, list: true},
		"Golf":    {get: true, create: true, createID: true, list: true},
	}
	for _, name := range []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf"} {
		addAIPResource(fd, svc, name, resources[name])
	}
	// EchoUpdate's Update takes the resource itself as its request, as IAM's
	// UpdateServiceAccount does. Read as no Update, it would be a Bravo.
	addAIPResource(fd, svc, "EchoUpdate", crud{get: true, create: true, createID: true, delete: true, list: true})
	svc.Method = append(svc.Method, method("UpdateEchoUpdate", "EchoUpdate", testPkg+".EchoUpdate", nil))
	p := newTestProto(t, fd, computeFile("Hotel"))

	want := map[protoreflect.FullName]Archetype{
		testPkg + ".Alpha":            ArchetypeStandard,
		testPkg + ".Bravo":            ArchetypeNoUpdate,
		testPkg + ".Charlie":          ArchetypeServerGeneratedID,
		testPkg + ".Delta":            ArchetypeSingleton,
		testPkg + ".Echo":             ArchetypeNonstandard,
		testPkg + ".EchoUpdate":       ArchetypeNonstandard,
		testPkg + ".Foxtrot":          ArchetypeUpdateWithoutMask,
		testPkg + ".Golf":             ArchetypeOther,
		testComputePkg + ".Hotel":     ArchetypeComputeStyle,
		testPkg + ".NotInTheDescSet":  ArchetypeUnresolved,
		testPkg + ".ThingService":     ArchetypeUnresolved,
		testComputePkg + ".Operation": ArchetypeNonstandard,
	}
	for name, wantArchetype := range want {
		t.Run(string(name), func(t *testing.T) {
			// Act
			api, _ := p.ResourceAPIByName(name)
			got := ClassifyArchetype(api)

			// Assert
			if got != wantArchetype {
				t.Errorf("ClassifyArchetype(%s) = %s (%s), want %s (%s)", name, got, got.Name(), wantArchetype, wantArchetype.Name())
			}
		})
	}
}

func TestArchetypeCodesAndNames(t *testing.T) {
	// Act
	all := AllArchetypes()

	// Assert
	var codes string
	names := map[string]Archetype{}
	for _, a := range all {
		codes += string(a)
		name := a.Name()
		if name == "" || name == "unknown" {
			t.Errorf("archetype %s has no name", a)
		}
		if other, dup := names[name]; dup {
			t.Errorf("archetypes %s and %s share the name %q", other, a, name)
		}
		names[name] = a
	}
	if codes != "ABCDEFGHI" {
		t.Errorf("AllArchetypes() = %q, want the letters A to I in order", codes)
	}
	if got := Archetype("Z").Name(); got != "unknown" {
		t.Errorf(`Archetype("Z").Name() = %q, want "unknown"`, got)
	}
}

const testComputePkg = "google.cloud.testcompute.v1"

// computeFile returns a file in testComputePkg declaring resource and a
// service that manages it in the google.cloud.compute convention: Get, and an
// Insert whose request carries the resource in a <resource>_resource field.
func computeFile(resource string) *descriptorpb.FileDescriptorProto {
	full := testComputePkg + "." + resource
	fd := testFile("compute.proto", testComputePkg,
		message(resource, stringField("name")),
		message("Operation", stringField("name")),
		message("Get"+resource+"Request", stringField(snakeCase(resource)), stringField("project")),
		message("Insert"+resource+"Request", messageField(snakeCase(resource)+"_resource", full), stringField("project"), stringField("request_id")),
	)
	fd.Service = []*descriptorpb.ServiceDescriptorProto{{
		Name: proto.String(resource + "s"),
		Method: []*descriptorpb.MethodDescriptorProto{
			method("Get", testComputePkg+".Get"+resource+"Request", full, nil),
			method("Insert", testComputePkg+".Insert"+resource+"Request", testComputePkg+".Operation", nil),
		},
	}}
	return fd
}
