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

package codegen

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestResolveKRMFieldName(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("acronyms.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: protoPtr("TestMessage"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   protoPtr("related_uris"),
				Number: protoPtr(int32(1)),
				Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
				Label:  labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED),
			}},
		}},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	field := fd.Messages().ByName("TestMessage").Fields().ByName("related_uris")

	for _, tc := range []struct {
		name           string
		goFields       []string
		pluralAcronyms bool
		want           string
	}{
		{
			name:     "an exact match is used with the flag off",
			goFields: []string{"RelatedUris"},
			want:     "RelatedUris",
		},
		{
			name:           "an exact match is used with the flag on",
			goFields:       []string{"RelatedUris", "RelatedURIs"},
			pluralAcronyms: true,
			want:           "RelatedUris",
		},
		{
			name:     "the plural acronym spelling is not tried with the flag off",
			goFields: []string{"RelatedURIs"},
			want:     "RelatedUris",
		},
		{
			name:           "the plural acronym spelling is found with the flag on",
			goFields:       []string{"RelatedURIs"},
			pluralAcronyms: true,
			want:           "RelatedURIs",
		},
		{
			name:           "neither spelling exists",
			goFields:       []string{"Other"},
			pluralAcronyms: true,
			want:           "RelatedUris",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			goFields := map[string]*gocode.StructField{}
			for _, name := range tc.goFields {
				goFields[name] = &gocode.StructField{}
			}

			// Act
			got := resolveKRMFieldName(field, "RelatedUris", goFields, tc.pluralAcronyms)

			// Assert
			if got != tc.want {
				t.Errorf("resolveKRMFieldName() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestKRMMapValueType(t *testing.T) {
	for _, tc := range []struct {
		name          string
		elemType      string
		krmFieldType  string
		wantGoType    string
		wantIsPointer bool
	}{
		{
			name:         "a generated struct takes the krm alias",
			elemType:     "TargetMessage",
			krmFieldType: "map[string]TargetMessage",
			wantGoType:   "krm.TargetMessage",
		},
		{
			name:          "a pointer to a generated struct",
			elemType:      "TargetMessage",
			krmFieldType:  "map[string]*TargetMessage",
			wantGoType:    "*krm.TargetMessage",
			wantIsPointer: true,
		},
		{
			name:         "a qualified type keeps its own qualifier",
			elemType:     "apiextensionsv1.JSON",
			krmFieldType: "map[string]apiextensionsv1.JSON",
			wantGoType:   "apiextensionsv1.JSON",
		},
		{
			name:         "a built-in type has no qualifier",
			elemType:     "string",
			krmFieldType: "map[string]string",
			wantGoType:   "string",
		},
		{
			name:          "a pointer to a built-in type",
			elemType:      "int64",
			krmFieldType:  "map[string]*int64",
			wantGoType:    "*int64",
			wantIsPointer: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			goType, isPointer := krmMapValueType(tc.elemType, tc.krmFieldType, "krm")

			// Assert
			if goType != tc.wantGoType || isPointer != tc.wantIsPointer {
				t.Errorf("krmMapValueType(%q, %q) = (%q, %v), want (%q, %v)",
					tc.elemType, tc.krmFieldType, goType, isPointer, tc.wantGoType, tc.wantIsPointer)
			}
		})
	}
}

// TestMapValueConverters checks which converters a map of messages uses. They
// follow the KRM value type, so map[string]TargetMessageWithRequired uses
// TargetMessageWithRequired's converters.
func TestMapValueConverters(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("maps.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("TargetMessage")},
			{
				Name: protoPtr("TestMessage"),
				NestedType: []*descriptorpb.DescriptorProto{{
					Name: protoPtr("TargetsEntry"),
					Field: []*descriptorpb.FieldDescriptorProto{
						{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
						{Name: protoPtr("value"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: protoPtr(".google.cloud.test.v1.TargetMessage")},
					},
					Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
				}},
				Field: []*descriptorpb.FieldDescriptorProto{{
					Name: protoPtr("targets"), Number: protoPtr(int32(1)),
					Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
					TypeName: protoPtr(".google.cloud.test.v1.TestMessage.TargetsEntry"),
					Label:    labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED),
				}},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	field := fd.Messages().ByName("TestMessage").Fields().ByName("targets")

	for _, tc := range []struct {
		name         string
		krmFieldType string
		wantOK       bool
		wantName     string
	}{
		{name: "the plain struct", krmFieldType: "map[string]TargetMessage", wantOK: true, wantName: "TargetMessage"},
		{name: "the Required copy", krmFieldType: "map[string]TargetMessageWithRequired", wantOK: true, wantName: "TargetMessageWithRequired"},
		{name: "a pointer to the Required copy", krmFieldType: "map[string]*TargetMessageWithRequired", wantOK: true, wantName: "TargetMessageWithRequired"},
		{name: "a type from another package falls back to the proto name", krmFieldType: "map[string]other.Thing", wantOK: true, wantName: "TargetMessage"},
		{name: "a slice keeps the custom-method path", krmFieldType: "[]TargetMessage", wantOK: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			fromProto, toProto, elemType, ok := mapValueConverters(field, tc.krmFieldType, "_v1alpha1")

			// Assert
			if ok != tc.wantOK {
				t.Fatalf("mapValueConverters(%q) ok = %v, want %v", tc.krmFieldType, ok, tc.wantOK)
			}
			if !ok {
				return
			}
			wantFrom, wantTo := tc.wantName+"_v1alpha1_FromProto", tc.wantName+"_v1alpha1_ToProto"
			if fromProto != wantFrom || toProto != wantTo || elemType != tc.wantName {
				t.Errorf("mapValueConverters(%q) = (%q, %q, %q), want (%q, %q, %q)",
					tc.krmFieldType, fromProto, toProto, elemType, wantFrom, wantTo, tc.wantName)
			}
		})
	}
}

func TestVisitMessageRecursesIntoSkippedContainer(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("serving_config.proto"),
		Package: protoPtr("google.cloud.retail.v2"),
		Options: &descriptorpb.FileOptions{
			GoPackage: protoPtr("cloud.google.com/go/retail/apiv2/retailpb;retailpb"),
		},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("SearchRequest"),
				NestedType: []*descriptorpb.DescriptorProto{
					{Name: protoPtr("DynamicFacetSpec")},
					{Name: protoPtr("UnusedNested")},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}

	g := NewMapperGenerator(func(msg protoreflect.MessageDescriptor) (string, bool) {
		if string(msg.FullName()) == "google.cloud.retail.v2.SearchRequest" {
			return "", false
		}
		return "retail", true
	}, t.TempDir(), nil, false)
	g.precomputedMappings["google.cloud.retail.v2.SearchRequest.DynamicFacetSpec"] = []*gocode.GoStruct{
		{Name: "SearchRequest_DynamicFacetSpec", GoPackage: "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"},
	}

	g.visitFile(fd)

	if len(g.typePairs) != 1 {
		t.Fatalf("len(typePairs) = %d, want 1", len(g.typePairs))
	}
	if got := g.typePairs[0].KRMType.Name; got != "SearchRequest_DynamicFacetSpec" {
		t.Errorf("typePairs[0].KRMType.Name = %q, want SearchRequest_DynamicFacetSpec", got)
	}
}

func TestMapperGeneratorLoadExistingImportsAndDisambiguate(t *testing.T) {
	baseDir := t.TempDir()
	pkgDir := filepath.Join(baseDir, "merchantapi")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	existingMapper := `// +generated:mapper
// krm.group: merchantapi.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.shopping.merchant.accounts.v1

package merchantapi

import (
	pb "cloud.google.com/go/shopping/merchant/accounts/apiv1/accountspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/merchantapi/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccountType_FromProto(mapCtx *direct.MapContext, in *pb.AccountType) *krm.AccountType {
	return nil
}
`
	if err := os.WriteFile(filepath.Join(pkgDir, "mapper.generated.go"), []byte(existingMapper), 0644); err != nil {
		t.Fatalf("writing mapper.generated.go: %v", err)
	}

	mg := NewMapperGenerator(nil, baseDir, &annotations.FileAnnotation{
		Key: "+generated:mapper",
		Attributes: map[string][]string{
			"krm.group":     {"merchantapi.cnrm.cloud.google.com"},
			"krm.version":   {"v1alpha1"},
			"proto.service": {"google.shopping.merchant.datasources.v1"},
		},
	}, false)
	mg.LoadExistingImports("merchantapi")

	// A second proto package requesting "pb" must receive a non-colliding alias.
	gotAlias := mg.AddGoImportAlias("cloud.google.com/go/shopping/merchant/datasources/apiv1/datasourcespb", "pb")
	if gotAlias != "datasourcespb" {
		t.Errorf("AddGoImportAlias(datasourcespb, \"pb\") = %q, want \"datasourcespb\"", gotAlias)
	}

	// Re-registering the original package keeps its existing "pb" alias.
	origAlias := mg.AddGoImportAlias("cloud.google.com/go/shopping/merchant/accounts/apiv1/accountspb", "pb")
	if origAlias != "pb" {
		t.Errorf("AddGoImportAlias(accountspb, \"pb\") = %q, want \"pb\"", origAlias)
	}
}

func TestMapperGeneratorProto3OptionalEnum(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Syntax:  protoPtr("proto3"),
		Name:    protoPtr("optional_enum.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		Options: &descriptorpb.FileOptions{
			GoPackage: protoPtr("cloud.google.com/go/test/apiv1/testpb;testpb"),
		},
		EnumType: []*descriptorpb.EnumDescriptorProto{{
			Name: protoPtr("Mode"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: protoPtr("MODE_UNSPECIFIED"), Number: protoPtr(int32(0))},
				{Name: protoPtr("MODE_ACTIVE"), Number: protoPtr(int32(1))},
			},
		}},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: protoPtr("Widget"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:           protoPtr("mode"),
				JsonName:       protoPtr("mode"),
				Number:         protoPtr(int32(1)),
				Type:           typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_ENUM),
				TypeName:       protoPtr(".google.cloud.test.v1.Mode"),
				Label:          labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL),
				Proto3Optional: protoPtr(true),
				OneofIndex:     protoPtr(int32(0)),
			}},
			OneofDecl: []*descriptorpb.OneofDescriptorProto{{
				Name: protoPtr("_mode"),
			}},
		}},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}
	msg := fd.Messages().ByName("Widget")

	mg := NewMapperGenerator(nil, t.TempDir(), &annotations.FileAnnotation{
		Key: "+generated:mapper",
		Attributes: map[string][]string{
			"krm.group":     {"test.cnrm.cloud.google.com"},
			"krm.version":   {"v1alpha1"},
			"proto.service": {"google.cloud.test.v1"},
		},
	}, false)

	pair := &typePair{
		KRMType: &gocode.GoStruct{
			Name:      "Widget",
			GoPackage: "github.com/GoogleCloudPlatform/k8s-config-connector/apis/test/v1alpha1",
			Fields: []*gocode.StructField{{
				Name: "Mode",
				Type: "*string",
			}},
		},
		Proto: msg,
	}
	var buf bytes.Buffer
	mg.writeMapFunctionsForPair(&buf, "test", pair)
	got := buf.String()

	wantToProto := "out.Mode = direct.EnumPtr_ToProto[testpb.Mode](mapCtx, in.Mode)"
	if !strings.Contains(got, wantToProto) {
		t.Errorf("generated mapper missing %q; got:\n%s", wantToProto, got)
	}
	unwantedOneof := "Widget_Mode_ToProto"
	if strings.Contains(got, unwantedOneof) {
		t.Errorf("generated mapper unexpectedly emitted synthetic oneof helper %q; got:\n%s", unwantedOneof, got)
	}
}

