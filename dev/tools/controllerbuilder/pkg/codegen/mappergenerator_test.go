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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protodesc"
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

// TestMapValueConverters checks which converters a map of messages calls. The
// name follows the KRM value type, so a strict struct that holds
// map[string]TargetMessageRequired calls TargetMessageRequired's converters.
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
		{name: "the Required copy", krmFieldType: "map[string]TargetMessageRequired", wantOK: true, wantName: "TargetMessageRequired"},
		{name: "a pointer to the Required copy", krmFieldType: "map[string]*TargetMessageRequired", wantOK: true, wantName: "TargetMessageRequired"},
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
