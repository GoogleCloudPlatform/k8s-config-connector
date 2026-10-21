// Copyright 2025 Google LLC
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

package scaffold

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestRenderControllerPlaceholder(t *testing.T) {
	args := ccTemplate.NewBaseControllerArgs(
		"beyondcorp",
		"v1alpha1",
		"BeyondCorpAppConnection",
		"AppConnection",
		"v1",
	)
	out, err := RenderController(args)
	if err != nil {
		t.Fatalf("RenderController failed: %v", err)
	}
	mustParseGo(t, out)

	code := string(out)
	for _, want := range []string{
		"resource, report, updateMask, err := common.CompareGreenfieldSpec(",
		"AppConnectionId: a.id.ID()",
		"UpdateMask:    updateMask",
		"created, err := op.Wait(ctx)",
		"updated, err = op.Wait(ctx)",
		"_, err = op.Wait(ctx)",
	} {
		if !strings.Contains(code, want) {
			t.Errorf("placeholder controller missing expected snippet %q\n--- source ---\n%s", want, code)
		}
	}
}

func TestApplyDeterministicPasses_Archetypes(t *testing.T) {
	tests := []struct {
		name             string
		args             *ccTemplate.ControllerArgs
		api              *protoapi.ResourceAPI
		krmInfo          *KRMPackageInfo
		regionalEndpoint string
		wantContains     []string
		wantNotContains  []string
	}{
		{
			name: "Archetype A: standard LRO with labels, references, and IdentityV2",
			args: ccTemplate.NewBaseControllerArgs("beyondcorp", "v1alpha1", "BeyondCorpAppConnection", "AppConnection", "v1"),
			api: &protoapi.ResourceAPI{
				DefaultHost: "beyondcorp.googleapis.com",
				HasLabels:   true,
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "app_connection",
					IDField:       "app_connection_id",
					LRO:           &protoapi.LROInfo{ResponseType: "google.cloud.beyondcorp.appconnections.v1.AppConnection"},
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "app_connection",
					UpdateMaskField: "update_mask",
					LRO:             &protoapi.LROInfo{ResponseType: "google.cloud.beyondcorp.appconnections.v1.AppConnection"},
				},
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					LRO:       &protoapi.LROInfo{ResponseType: "google.protobuf.Empty"},
				},
			},
			krmInfo: &KRMPackageInfo{
				IdentityTypeName:  "BeyondCorpAppConnectionIdentity",
				HasGetIdentity:    true,
				ParentStringExpr:  "a.id.ParentString()",
				IDExpr:            "a.id.ID()",
				HasSpecLabels:     false,
				HasSpecReferences: true,
				StatusInspected:   true,
				HasObservedState:  true,
			},
			wantContains: []string{
				"i, err := obj.GetIdentity(ctx, reader)",
				"id := i.(*krm.BeyondCorpAppConnectionIdentity)",
				"common.NormalizeReferences(ctx, reader, obj, nil)",
				"Parent:          a.id.ParentString()",
				"AppConnectionId: a.id.ID()",
				"resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(",
				"UpdateMask:    updateMask",
				"err = op.Wait(ctx)",
			},
			wantNotContains: []string{
				"common.CompareGreenfieldSpec(",
				"resource.Name = a.id.String()\n\n\treq := &beyondcorppb.Create",
			},
		},
		{
			name: "Archetype A: synchronous CRUD with Empty Delete response",
			args: ccTemplate.NewBaseControllerArgs("dataform", "v1beta1", "DataformRepository", "Repository", "v1beta1"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "repository",
					IDField:       "repository_id",
					Response:      "google.cloud.dataform.v1beta1.Repository",
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "repository",
					UpdateMaskField: "update_mask",
					Response:        "google.cloud.dataform.v1beta1.Repository",
				},
				Delete: &protoapi.StandardMethod{
					NameField:  "name",
					ForceField: "force",
					Response:   "google.protobuf.Empty",
				},
			},
			wantContains: []string{
				"created, err := a.gcpClient.CreateRepository(ctx, req)",
				"updated, err = a.gcpClient.UpdateRepository(ctx, req)",
				"err := a.gcpClient.DeleteRepository(ctx, req)",
				"Force: true",
			},
			wantNotContains: []string{
				"op.Wait(ctx)",
			},
		},
		{
			name: "Archetype B: no-update (immutable resource)",
			args: ccTemplate.NewBaseControllerArgs("gkebackup", "v1alpha1", "GKEBackupRestore", "Restore", "v1"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "restore",
					IDField:       "restore_id",
					LRO:           &protoapi.LROInfo{ResponseType: "google.cloud.gkebackup.v1.Restore"},
				},
				Update: nil,
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					LRO:       &protoapi.LROInfo{ResponseType: "google.protobuf.Empty"},
				},
			},
			wantContains: []string{
				`return fmt.Errorf("updating GKEBackupRestore is not supported, fields modified: %v", updateMask.Paths)`,
			},
			wantNotContains: []string{
				"UpdateRestoreRequest",
				"a.gcpClient.UpdateRestore",
			},
		},
		{
			name: "Archetype C Variant 3A: client-specified resource name in body",
			args: ccTemplate.NewBaseControllerArgs("servicedirectory", "v1alpha1", "ServiceDirectoryNamespace", "Namespace", "v1"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "namespace",
					IDField:       "", // no <resource>_id field on CreateRequest
					Response:      "google.cloud.servicedirectory.v1.Namespace",
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "namespace",
					UpdateMaskField: "update_mask",
					Response:        "google.cloud.servicedirectory.v1.Namespace",
				},
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					Response:  "google.protobuf.Empty",
				},
			},
			krmInfo: &KRMPackageInfo{
				IdentityTypeName:     "ServiceDirectoryNamespaceIdentity",
				HasGetIdentity:       true,
				HasIdentitySpecified: false,
			},
			wantContains: []string{
				"resource.Name = a.id.String()",
				"Parent:    a.id.Parent().String()",
				"Namespace: resource",
			},
			wantNotContains: []string{
				"NamespaceId:",
				"HasIdentitySpecified()",
			},
		},
		{
			name: "Archetype C Variant 3B: server-generated ID",
			args: ccTemplate.NewBaseControllerArgs("cloudtalentsolution", "v1alpha1", "CloudTalentSolutionCompany", "Company", "v4"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "company",
					IDField:       "",
					Response:      "google.cloud.talent.v4.Company",
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "company",
					UpdateMaskField: "update_mask",
					Response:        "google.cloud.talent.v4.Company",
				},
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					Response:  "google.protobuf.Empty",
				},
			},
			krmInfo: &KRMPackageInfo{
				IdentityTypeName:     "CloudTalentSolutionCompanyIdentity",
				HasGetIdentity:       true,
				HasIdentitySpecified: true,
			},
			wantContains: []string{
				"if !a.id.HasIdentitySpecified() {\n\t\treturn false, nil\n\t}",
				"status.ExternalRef = direct.LazyPtr(created.Name)",
			},
			wantNotContains: []string{
				"CompanyId:",
			},
		},
		{
			name: "Archetype D: singleton (Get + Update only, no Create, no Delete)",
			args: ccTemplate.NewBaseControllerArgs("artifactregistry", "v1alpha1", "ArtifactRegistryVPCSCConfig", "VPCSCConfig", "v1"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: nil,
				Update: &protoapi.StandardMethod{
					ResourceField:   "vpcsc_config",
					UpdateMaskField: "update_mask",
					Response:        "google.devtools.artifactregistry.v1.VPCSCConfig",
				},
				Delete: nil,
			},
			wantContains: []string{
				"creating (updating singleton) VPCSCConfig",
				"req := &artifactregistrypb.UpdateVPCSCConfigRequest{",
				"VpcscConfig: resource,",
				"deletion not supported for VPCSCConfig, removing finalizer without deleting GCP resource",
			},
			wantNotContains: []string{
				"CreateVPCSCConfigRequest",
				"DeleteVPCSCConfigRequest",
			},
		},
		{
			name: "Archetype F: update without UpdateMask",
			args: ccTemplate.NewBaseControllerArgs("cloudiot", "v1alpha1", "CloudIoTDeviceRegistry", "DeviceRegistry", "v1"),
			api: &protoapi.ResourceAPI{
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "device_registry",
					IDField:       "device_registry_id",
					Response:      "google.cloud.iot.v1.DeviceRegistry",
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "device_registry",
					UpdateMaskField: "", // no update_mask field
					Response:        "google.cloud.iot.v1.DeviceRegistry",
				},
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					Response:  "google.protobuf.Empty",
				},
			},
			wantContains: []string{
				"req := &cloudiotpb.UpdateDeviceRegistryRequest{\n\t\t\tDeviceRegistry: resource,\n\t\t}",
			},
			wantNotContains: []string{
				"UpdateMask:",
				"fieldmaskpb.FieldMask",
			},
		},
		{
			name: "Combined Archetype B + G: immutable and undeletable (e.g. KMSKeyHandle)",
			args: ccTemplate.NewBaseControllerArgs("kms", "v1beta1", "KMSKeyHandle", "KeyHandle", "v1"),
			api: &protoapi.ResourceAPI{
				DefaultHost: "cloudkms.googleapis.com",
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "key_handle",
					IDField:       "key_handle_id",
					LRO:           &protoapi.LROInfo{ResponseType: "google.cloud.kms.v1.KeyHandle"},
				},
				Update: nil,
				Delete: nil,
			},
			krmInfo: &KRMPackageInfo{
				LocationExpr: "direct.ValueOf(obj.Spec.Location)",
			},
			regionalEndpoint: "rep",
			wantContains: []string{
				`opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://cloudkms.%s.rep.googleapis.com", location)))`,
				"gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))",
				`return fmt.Errorf("updating KMSKeyHandle is not supported, fields modified: %v", updateMask.Paths)`,
				"deletion not supported for KeyHandle, removing finalizer without deleting GCP resource",
			},
			wantNotContains: []string{
				"UpdateKeyHandle",
				"DeleteKeyHandle",
			},
		},
		{
			name: "Regional endpoint prefix pattern ({loc}-{service}:443)",
			args: ccTemplate.NewBaseControllerArgs("cloudbuild", "v1beta1", "CloudBuildWorkerPool", "WorkerPool", "v1"),
			api: &protoapi.ResourceAPI{
				DefaultHost: "cloudbuild.googleapis.com",
				Get: &protoapi.StandardMethod{
					NameField: "name",
				},
				Create: &protoapi.StandardMethod{
					ParentField:   "parent",
					ResourceField: "worker_pool",
					IDField:       "worker_pool_id",
					LRO:           &protoapi.LROInfo{ResponseType: "google.devtools.cloudbuild.v1.WorkerPool"},
				},
				Update: &protoapi.StandardMethod{
					ResourceField:   "worker_pool",
					UpdateMaskField: "update_mask",
					LRO:             &protoapi.LROInfo{ResponseType: "google.devtools.cloudbuild.v1.WorkerPool"},
				},
				Delete: &protoapi.StandardMethod{
					NameField: "name",
					LRO:       &protoapi.LROInfo{ResponseType: "google.protobuf.Empty"},
				},
			},
			krmInfo: &KRMPackageInfo{
				LocationExpr:        "direct.ValueOf(obj.Spec.Location)",
				MapperVersionSuffix: "_v1beta1",
			},
			regionalEndpoint: "prefix",
			wantContains: []string{
				`opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-cloudbuild.googleapis.com:443", location)))`,
				"CloudBuildWorkerPoolSpec_v1beta1_ToProto",
				"CloudBuildWorkerPoolObservedState_v1beta1_FromProto",
				"CloudBuildWorkerPoolSpec_v1beta1_FromProto",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ApplyDeterministicPasses(tc.args, tc.api, tc.krmInfo, tc.regionalEndpoint)
			out, err := RenderController(tc.args)
			if err != nil {
				t.Fatalf("RenderController failed: %v", err)
			}
			mustParseGo(t, out)

			code := string(out)
			for _, want := range tc.wantContains {
				if !strings.Contains(code, want) {
					t.Errorf("expected generated code to contain %q, but it did not.\n--- source ---\n%s", want, code)
				}
			}
			for _, notWant := range tc.wantNotContains {
				if strings.Contains(code, notWant) {
					t.Errorf("expected generated code NOT to contain %q, but it did.\n--- source ---\n%s", notWant, code)
				}
			}
		})
	}
}

func TestInspectKRMPackageForController(t *testing.T) {
	tmpRoot := t.TempDir()
	apiDir := filepath.Join(tmpRoot, "apis", "widgets", "v1alpha1")
	directDir := filepath.Join(tmpRoot, "pkg", "controller", "direct", "widgets")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(directDir, 0755); err != nil {
		t.Fatal(err)
	}

	typesSrc := `package v1alpha1

type WidgetsGadgetSpec struct {
	Location *string
	Labels   map[string]string
	NetworkRef *NetworkRef
}

type WidgetsGadgetStatus struct {
	ObservedState *WidgetsGadgetObservedState
}

type WidgetsGadget struct {
	Spec   WidgetsGadgetSpec
	Status WidgetsGadgetStatus
}
`
	identitySrc := `package v1alpha1

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WidgetsGadgetIdentity struct {
	Project  string
	Location string
	Gadget   string
}

func (i *WidgetsGadgetIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func (i *WidgetsGadgetIdentity) HasIdentitySpecified() bool {
	return i.Gadget != ""
}

func (i *WidgetsGadgetIdentity) FromExternal(ref string) error {
	return nil
}

func (obj *WidgetsGadget) GetIdentity(ctx context.Context, reader client.Reader) (interface{}, error) {
	return &WidgetsGadgetIdentity{}, nil
}
`
	mapperSrc := `package widgets

func WidgetsGadgetSpec_v1alpha1_ToProto() {}
`
	siblingControllerSrc := `package widgets

import (
	"fmt"
	"google.golang.org/api/option"
)

func client(location string) {
	_ = option.WithEndpoint(fmt.Sprintf("%s-widgets.googleapis.com:443", location))
}
`

	if err := os.WriteFile(filepath.Join(apiDir, "gadget_types.go"), []byte(typesSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(apiDir, "gadget_identity.go"), []byte(identitySrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directDir, "mapper.generated.go"), []byte(mapperSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directDir, "other_controller.go"), []byte(siblingControllerSrc), 0644); err != nil {
		t.Fatal(err)
	}

	info := InspectKRMPackageForController(tmpRoot, "widgets", "v1alpha1", "WidgetsGadget", "Gadget")
	if info.IdentityTypeName != "WidgetsGadgetIdentity" {
		t.Errorf("IdentityTypeName = %q, want WidgetsGadgetIdentity", info.IdentityTypeName)
	}
	if !info.HasGetIdentity {
		t.Errorf("HasGetIdentity = false, want true")
	}
	if !info.HasIdentitySpecified {
		t.Errorf("HasIdentitySpecified = false, want true")
	}
	if !info.HasFromExternal {
		t.Errorf("HasFromExternal = false, want true")
	}
	if info.ParentStringExpr != "a.id.ParentString()" {
		t.Errorf("ParentStringExpr = %q, want a.id.ParentString()", info.ParentStringExpr)
	}
	if info.IDExpr != "a.id.Gadget" {
		t.Errorf("IDExpr = %q, want a.id.Gadget", info.IDExpr)
	}
	if !info.HasSpecLabels {
		t.Errorf("HasSpecLabels = false, want true")
	}
	if !info.HasSpecReferences {
		t.Errorf("HasSpecReferences = false, want true")
	}
	if !info.HasObservedState {
		t.Errorf("HasObservedState = false, want true")
	}
	if info.LocationExpr != "direct.ValueOf(obj.Spec.Location)" {
		t.Errorf("LocationExpr = %q, want direct.ValueOf(obj.Spec.Location)", info.LocationExpr)
	}
	if info.MapperVersionSuffix != "_v1alpha1" {
		t.Errorf("MapperVersionSuffix = %q, want _v1alpha1", info.MapperVersionSuffix)
	}
	if info.DetectedRegionalEndpoint != "prefix" {
		t.Errorf("DetectedRegionalEndpoint = %q, want prefix", info.DetectedRegionalEndpoint)
	}
}

func TestApplyDeterministicPasses_RealGoogleapis(t *testing.T) {
	pbPath := filepath.Join("..", "..", "..", "..", ".build", "googleapis.pb")
	if _, err := os.Stat(pbPath); err != nil {
		t.Skipf("skipping real googleapis.pb test: %v", err)
	}
	p, err := protoapi.LoadProto(pbPath, "")
	if err != nil {
		t.Fatalf("LoadProto(%q) failed: %v", pbPath, err)
	}

	cases := []struct {
		service      string
		version      string
		kind         string
		protoRes     string
		protoVer     string
		fullProto    string
		wantContains []string
	}{
		{
			service:   "beyondcorp",
			version:   "v1alpha1",
			kind:      "BeyondCorpAppConnection",
			protoRes:  "AppConnection",
			protoVer:  "v1",
			fullProto: "google.cloud.beyondcorp.appconnections.v1.AppConnection",
			wantContains: []string{
				`gcp "cloud.google.com/go/beyondcorp/appconnections/apiv1"`,
				`appconnectionspb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"`,
				`gcpClient, err := gcp.NewRESTClient(ctx, opts...)`,
				`resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(`,
			},
		},
		{
			service:   "artifactregistry",
			version:   "v1alpha1",
			kind:      "ArtifactRegistryVPCSCConfig",
			protoRes:  "VPCSCConfig",
			protoVer:  "v1",
			fullProto: "google.devtools.artifactregistry.v1.VPCSCConfig",
			wantContains: []string{
				`creating (updating singleton) VPCSCConfig`,
				`deletion not supported for VPCSCConfig, removing finalizer without deleting GCP resource`,
			},
		},
		{
			service:   "kms",
			version:   "v1beta1",
			kind:      "KMSKeyHandle",
			protoRes:  "KeyHandle",
			protoVer:  "v1",
			fullProto: "google.cloud.kms.v1.KeyHandle",
			wantContains: []string{
				`gcp "cloud.google.com/go/kms/apiv1"`,
				`kmspb "cloud.google.com/go/kms/apiv1/kmspb"`,
				`gcpClient, err := gcp.NewAutokeyRESTClient(ctx, opts...)`,
				`return fmt.Errorf("updating KMSKeyHandle is not supported, fields modified: %v", updateMask.Paths)`,
				`deletion not supported for KeyHandle, removing finalizer without deleting GCP resource`,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.kind, func(t *testing.T) {
			api, err := p.ResourceAPIByName(protoreflect.FullName(tc.fullProto))
			if err != nil {
				t.Fatalf("ResourceAPIByName(%q) failed: %v", tc.fullProto, err)
			}
			args := ccTemplate.NewBaseControllerArgs(tc.service, tc.version, tc.kind, tc.protoRes, tc.protoVer)
			ApplyDeterministicPasses(args, api, nil, "")
			out, err := RenderController(args)
			if err != nil {
				t.Fatalf("RenderController(%s) failed: %v", tc.kind, err)
			}
			mustParseGo(t, out)
			code := string(out)
			for _, want := range tc.wantContains {
				if !strings.Contains(code, want) {
					t.Errorf("expected generated %s controller to contain %q\n--- source ---\n%s", tc.kind, want, code)
				}
			}
		})
	}
}

func mustParseGo(t *testing.T, src []byte) {
	t.Helper()
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "controller.go", src, parser.AllErrors); err != nil {
		t.Fatalf("generated controller failed to parse as valid Go:\n%v\n--- source ---\n%s", err, string(src))
	}
}

