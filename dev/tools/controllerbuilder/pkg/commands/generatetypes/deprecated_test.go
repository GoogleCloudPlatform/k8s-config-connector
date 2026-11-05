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

package generatetypes

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestDeprecatedResourceReason(t *testing.T) {
	for _, tc := range []struct {
		name              string
		pkg               string
		serviceDeprecated bool
		rpcDeprecated     []bool
		wantDeprecated    bool
		wantSubstr        string
	}{
		{
			name:              "active service with no deprecated RPCs",
			pkg:               "google.cloud.example.v1",
			serviceDeprecated: false,
			rpcDeprecated:     []bool{false, false, false, false},
			wantDeprecated:    false,
		},
		{
			name:              "active service with only some deprecated RPCs is not flagged",
			pkg:               "google.cloud.example.v1",
			serviceDeprecated: false,
			rpcDeprecated:     []bool{true, false, false, false},
			wantDeprecated:    false,
		},
		{
			name:              "service with option deprecated = true",
			pkg:               "google.cloud.example.v1",
			serviceDeprecated: true,
			rpcDeprecated:     []bool{false, false, false, false},
			wantDeprecated:    true,
			wantSubstr:        "service google.cloud.example.v1.WidgetService is marked deprecated in proto",
		},
		{
			name:              "service where all RPCs have option deprecated = true",
			pkg:               "google.cloud.example.v1",
			serviceDeprecated: false,
			rpcDeprecated:     []bool{true, true, true, true},
			wantDeprecated:    true,
			wantSubstr:        "all 4 RPCs in service google.cloud.example.v1.WidgetService are marked deprecated in proto",
		},
		{
			name:              "curated deprecated package (automl) without proto deprecated options",
			pkg:               "google.cloud.automl.v1",
			serviceDeprecated: false,
			rpcDeprecated:     []bool{false, false, false, false},
			wantDeprecated:    true,
			wantSubstr:        "Cloud AutoML (google.cloud.automl.v1) is deprecated in favor of Vertex AI",
		},
		{
			name:              "curated deprecated package (iot) without proto deprecated options",
			pkg:               "google.cloud.iot.v1",
			serviceDeprecated: false,
			rpcDeprecated:     []bool{false, false, false, false},
			wantDeprecated:    true,
			wantSubstr:        "Cloud IoT Core (google.cloud.iot.v1) is deprecated",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pbPath := writeDeprecatedTestProto(t, t.TempDir(), tc.pkg, tc.serviceDeprecated, tc.rpcDeprecated)
			api, err := protoapi.LoadProto(pbPath, "")
			if err != nil {
				t.Fatalf("LoadProto: %v", err)
			}
			d, err := api.Files().FindDescriptorByName(protoreflect.FullName(tc.pkg + ".Widget"))
			if err != nil {
				t.Fatalf("FindDescriptorByName: %v", err)
			}
			msg := d.(protoreflect.MessageDescriptor)
			reason, got := deprecatedResourceReason(api, msg)
			if got != tc.wantDeprecated {
				t.Fatalf("deprecatedResourceReason() = (%q, %v), want deprecated=%v", reason, got, tc.wantDeprecated)
			}
			if tc.wantDeprecated && !strings.Contains(reason, tc.wantSubstr) {
				t.Errorf("deprecatedResourceReason() reason = %q, want substring %q", reason, tc.wantSubstr)
			}
		})
	}
}

func TestDeprecatedResourceReasonMixedPackageServices(t *testing.T) {
	// Simulate a proto package like google.cloud.dataplex.v1 or
	// google.cloud.datacatalog.v1 where one service (LegacyService) is
	// deprecated while another service (ActiveService) in the same package is
	// active.
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	legacyMsgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(legacyMsgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/LegacyTaxonomy",
		Pattern: []string{"projects/{project}/locations/{location}/legacyTaxonomies/{legacy_taxonomy}"},
	})
	activeMsgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(activeMsgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Lake",
		Pattern: []string{"projects/{project}/locations/{location}/lakes/{lake}"},
	})

	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{
		{
			Name:    proto.String("google/cloud/example/v1/legacy.proto"),
			Package: proto.String("google.cloud.example.v1"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{{
				Name:    proto.String("LegacyTaxonomy"),
				Options: legacyMsgOpts,
				Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
			}},
			Service: []*descriptorpb.ServiceDescriptorProto{{
				Name:    proto.String("LegacyTaxonomyService"),
				Options: &descriptorpb.ServiceOptions{Deprecated: proto.Bool(true)},
				Method: []*descriptorpb.MethodDescriptorProto{{
					Name:       proto.String("CreateLegacyTaxonomy"),
					InputType:  proto.String(".google.cloud.example.v1.LegacyTaxonomy"),
					OutputType: proto.String(".google.cloud.example.v1.LegacyTaxonomy"),
				}},
			}},
		},
		{
			Name:    proto.String("google/cloud/example/v1/active.proto"),
			Package: proto.String("google.cloud.example.v1"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{{
				Name:    proto.String("Lake"),
				Options: activeMsgOpts,
				Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
			}},
			Service: []*descriptorpb.ServiceDescriptorProto{{
				Name: proto.String("ActiveService"),
				Method: []*descriptorpb.MethodDescriptorProto{{
					Name:       proto.String("CreateLake"),
					InputType:  proto.String(".google.cloud.example.v1.Lake"),
					OutputType: proto.String(".google.cloud.example.v1.Lake"),
				}},
			}},
		},
	}}

	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(t.TempDir(), "mixed.pb")
	writeFile(t, pbPath, string(b))

	api, err := protoapi.LoadProto(pbPath, "")
	if err != nil {
		t.Fatalf("LoadProto: %v", err)
	}

	legacyDesc, err := api.Files().FindDescriptorByName("google.cloud.example.v1.LegacyTaxonomy")
	if err != nil {
		t.Fatalf("FindDescriptorByName LegacyTaxonomy: %v", err)
	}
	if reason, deprecated := deprecatedResourceReason(api, legacyDesc.(protoreflect.MessageDescriptor)); !deprecated {
		t.Errorf("expected LegacyTaxonomy to be detected as deprecated, got false")
	} else if !strings.Contains(reason, "LegacyTaxonomyService") {
		t.Errorf("unexpected reason for LegacyTaxonomy: %q", reason)
	}

	activeDesc, err := api.Files().FindDescriptorByName("google.cloud.example.v1.Lake")
	if err != nil {
		t.Fatalf("FindDescriptorByName Lake: %v", err)
	}
	if reason, deprecated := deprecatedResourceReason(api, activeDesc.(protoreflect.MessageDescriptor)); deprecated {
		t.Errorf("expected Lake in active service not to be deprecated, got reason: %q", reason)
	}
}

func TestDeprecatedResourceReasonSharedServiceResourceRPCs(t *testing.T) {
	// Simulate a shared active service (like google.cloud.edgenetwork.v1.EdgeNetwork)
	// that manages both an active resource (Network) and a resource (Zone) whose
	// resource-specific RPCs are marked `option deprecated = true;`.
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	zoneOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(zoneOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Zone",
		Pattern: []string{"projects/{project}/locations/{location}/zones/{zone}"},
	})
	networkOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(networkOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Network",
		Pattern: []string{"projects/{project}/locations/{location}/networks/{network}"},
	})

	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/edgenetwork.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name:    proto.String("Zone"),
				Options: zoneOpts,
				Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
			},
			{
				Name:    proto.String("Network"),
				Options: networkOpts,
				Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{{
			Name: proto.String("EdgeNetwork"),
			Method: []*descriptorpb.MethodDescriptorProto{
				{
					Name:       proto.String("GetZone"),
					InputType:  proto.String(".google.cloud.example.v1.Zone"),
					OutputType: proto.String(".google.cloud.example.v1.Zone"),
					Options:    &descriptorpb.MethodOptions{Deprecated: proto.Bool(true)},
				},
				{
					Name:       proto.String("CreateNetwork"),
					InputType:  proto.String(".google.cloud.example.v1.Network"),
					OutputType: proto.String(".google.cloud.example.v1.Network"),
				},
				{
					Name:       proto.String("GetNetwork"),
					InputType:  proto.String(".google.cloud.example.v1.Network"),
					OutputType: proto.String(".google.cloud.example.v1.Network"),
				},
			},
		}},
	}}}

	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(t.TempDir(), "shared_service.pb")
	writeFile(t, pbPath, string(b))

	api, err := protoapi.LoadProto(pbPath, "")
	if err != nil {
		t.Fatalf("LoadProto: %v", err)
	}

	zoneDesc, err := api.Files().FindDescriptorByName("google.cloud.example.v1.Zone")
	if err != nil {
		t.Fatalf("FindDescriptorByName Zone: %v", err)
	}
	if reason, deprecated := deprecatedResourceReason(api, zoneDesc.(protoreflect.MessageDescriptor)); !deprecated {
		t.Errorf("expected Zone with all deprecated RPCs in shared service to be detected as deprecated, got false")
	} else if !strings.Contains(reason, "RPCs managing Zone") {
		t.Errorf("unexpected reason for Zone: %q", reason)
	}

	networkDesc, err := api.Files().FindDescriptorByName("google.cloud.example.v1.Network")
	if err != nil {
		t.Fatalf("FindDescriptorByName Network: %v", err)
	}
	if reason, deprecated := deprecatedResourceReason(api, networkDesc.(protoreflect.MessageDescriptor)); deprecated {
		t.Errorf("expected Network with active RPCs in shared service not to be deprecated, got reason: %q", reason)
	}
}

func TestRunGenerateCRDRejectsNewDeprecatedResourceWithoutFlag(t *testing.T) {
	apiDir := t.TempDir()
	pbPath := writeDeprecatedTestProto(t, t.TempDir(), "google.cloud.example.v1", true, []bool{true, true, true, true})

	baseOpts := func() *GenerateCRDOptions {
		return &GenerateCRDOptions{
			GenerateOptions: &options.GenerateOptions{
				ProtoSourcePath: pbPath,
				APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
			},
			ServiceName:         "google.cloud.example.v1",
			OutputAPIDirectory:  apiDir,
			Resources:           options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
			PrepopulateResource: true,
		}
	}

	// 1. Scaffolding a new resource backed by a deprecated service fails without --allow-deprecated.
	err := RunGenerateCRD(context.Background(), baseOpts())
	if err == nil || !strings.Contains(err.Error(), "--allow-deprecated") {
		t.Fatalf("expected --allow-deprecated error, got: %v", err)
	}
	typeFile := filepath.Join(apiDir, "example/v1alpha1/examplewidget_types.go")
	if _, statErr := os.Stat(typeFile); !os.IsNotExist(statErr) {
		t.Fatalf("expected %s not to be created on failure", typeFile)
	}

	// 2. Passing --allow-deprecated succeeds and creates <kind>_types.go.
	allowed := baseOpts()
	allowed.AllowDeprecated = true
	if err := RunGenerateCRD(context.Background(), allowed); err != nil {
		t.Fatalf("RunGenerateCRD with AllowDeprecated: %v", err)
	}
	if _, statErr := os.Stat(typeFile); statErr != nil {
		t.Fatalf("expected %s to exist after AllowDeprecated: %v", typeFile, statErr)
	}

	// 3. Re-running when <kind>_types.go already exists succeeds even without --allow-deprecated.
	if err := RunGenerateCRD(context.Background(), baseOpts()); err != nil {
		t.Fatalf("RunGenerateCRD on existing resource without AllowDeprecated: %v", err)
	}
}

func writeDeprecatedTestProto(t *testing.T, dir, pkg string, serviceDeprecated bool, rpcDeprecated []bool) string {
	t.Helper()
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	msgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(msgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Widget",
		Pattern: []string{"projects/{project}/locations/{location}/widgets/{widget}"},
	})
	rpcNames := []string{"CreateWidget", "GetWidget", "UpdateWidget", "DeleteWidget"}
	var methods []*descriptorpb.MethodDescriptorProto
	for i, dep := range rpcDeprecated {
		name := rpcNames[i%len(rpcNames)]
		methods = append(methods, &descriptorpb.MethodDescriptorProto{
			Name:       proto.String(name),
			InputType:  proto.String("." + pkg + ".Widget"),
			OutputType: proto.String("." + pkg + ".Widget"),
			Options:    &descriptorpb.MethodOptions{Deprecated: proto.Bool(dep)},
		})
	}
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("widget.proto"),
		Package: proto.String(pkg),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:    proto.String("Widget"),
			Options: msgOpts,
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
				{Name: proto.String("display_name"), Number: proto.Int32(2), Type: str},
			},
		}},
		Service: []*descriptorpb.ServiceDescriptorProto{{
			Name:    proto.String("WidgetService"),
			Options: &descriptorpb.ServiceOptions{Deprecated: proto.Bool(serviceDeprecated)},
			Method:  methods,
		}},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	path := filepath.Join(dir, "widget.pb")
	writeFile(t, path, string(b))
	return path
}
