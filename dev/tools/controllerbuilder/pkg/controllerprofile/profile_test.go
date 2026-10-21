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

package controllerprofile

import (
	"strings"
	"testing"
)

func TestBuildProfile_StandardLROWithRegionalEndpoint(t *testing.T) {
	protoSpec := &ProtoServiceSpec{
		ServiceName:  "FeatureOnlineStoreAdminService",
		ServiceHost:  "aiplatform.googleapis.com",
		GoPackage:    "cloud.google.com/go/aiplatform/apiv1/aiplatformpb;aiplatformpb",
		MessageName:  "FeatureView",
		HasNameField: true,
		HasLabels:    true,
		Get: &ProtoMethodSpec{
			Name:      "GetFeatureView",
			InputType: "GetFeatureViewRequest",
			NameField: "name",
		},
		Create: &ProtoMethodSpec{
			Name:          "CreateFeatureView",
			InputType:     "CreateFeatureViewRequest",
			ParentField:   "parent",
			ResourceField: "feature_view",
			IDField:       "feature_view_id",
			IsLRO:         true,
			ResponseType:  "FeatureView",
		},
		Update: &ProtoMethodSpec{
			Name:            "UpdateFeatureView",
			InputType:       "UpdateFeatureViewRequest",
			ResourceField:   "feature_view",
			UpdateMaskField: "update_mask",
			IsLRO:           true,
			ResponseType:    "FeatureView",
		},
		Delete: &ProtoMethodSpec{
			Name:         "DeleteFeatureView",
			InputType:    "DeleteFeatureViewRequest",
			NameField:    "name",
			IsLRO:        true,
			ResponseType: "google.protobuf.Empty",
		},
	}

	krmSpec := &KRMSpecInfo{
		IdentityTypeName:  "AIPlatformFeatureViewIdentity",
		HasGetIdentity:    true,
		HasFromExternal:   true,
		ParentStringExpr:  "ParentString()",
		IDExpr:            "ID()",
		HasSpecReferences: true,
		HasObservedState:  true,
		LocationExpr:      "direct.ValueOf(obj.Spec.Location)",
		RegionalEndpoint:  "prefix",
		UseGRPCClient:     true,
	}

	profile := BuildProfile("aiplatform", "v1alpha1", "AIPlatformFeatureView", "FeatureView", "v1", protoSpec, krmSpec)

	// Verify Identity Trait
	if profile.Identity.Strategy != IdentityClientField {
		t.Errorf("got Identity.Strategy = %s, want %s", profile.Identity.Strategy, IdentityClientField)
	}

	// Verify Lifecycle Trait
	if profile.Lifecycle.UpdateStyle != UpdateStyleMask {
		t.Errorf("got Lifecycle.UpdateStyle = %s, want %s", profile.Lifecycle.UpdateStyle, UpdateStyleMask)
	}
	if profile.Lifecycle.DeleteStyle != DeleteStyleStandard {
		t.Errorf("got Lifecycle.DeleteStyle = %s, want %s", profile.Lifecycle.DeleteStyle, DeleteStyleStandard)
	}

	// Verify Execution Trait
	if !profile.Execution.Create.IsLRO || !profile.Execution.Update.IsLRO || !profile.Execution.Delete.IsLRO {
		t.Errorf("expected all CRUD methods to be LRO, got: %+v", profile.Execution)
	}
	if !profile.Execution.Delete.IsEmpty {
		t.Errorf("expected Delete to be Empty response, got false")
	}

	// Verify Transport Trait
	if profile.Transport.RegionalEndpoint != "prefix" {
		t.Errorf("got RegionalEndpoint = %s, want prefix", profile.Transport.RegionalEndpoint)
	}
	if !profile.Transport.UseGRPCClient {
		t.Errorf("expected UseGRPCClient = true")
	}
	if profile.Transport.LabelDiffing != LabelsMetadata {
		t.Errorf("got LabelDiffing = %s, want %s", profile.Transport.LabelDiffing, LabelsMetadata)
	}

	// Render controller and verify valid Go code
	src, err := RenderController(profile)
	if err != nil {
		t.Fatalf("RenderController failed: %v\nCode:\n%s", err, string(src))
	}

	srcStr := string(src)
	mustContain := []string{
		"package aiplatform",
		"gcp.NewFeatureOnlineStoreAdminClient(ctx, opts...)",
		"common.NormalizeReferences(ctx, reader, obj, nil)",
		"common.CompareGreenfieldSpecAndLabels(",
		"op, err := a.gcpClient.CreateFeatureView(ctx, req)",
		"created, err := op.Wait(ctx)",
		"op.Wait(ctx)",
	}
	for _, term := range mustContain {
		if !strings.Contains(srcStr, term) {
			t.Errorf("rendered code missing expected string %q", term)
		}
	}
}

func TestBuildProfile_ServerGeneratedIDSync(t *testing.T) {
	protoSpec := &ProtoServiceSpec{
		ServiceName:  "CloudChannelService",
		ServiceHost:  "cloudchannel.googleapis.com",
		GoPackage:    "cloud.google.com/go/channel/apiv1/channelpb;channelpb",
		MessageName:  "Customer",
		HasNameField: true,
		Get: &ProtoMethodSpec{
			Name:      "GetCustomer",
			InputType: "GetCustomerRequest",
			NameField: "name",
		},
		Create: &ProtoMethodSpec{
			Name:          "CreateCustomer",
			InputType:     "CreateCustomerRequest",
			ParentField:   "parent",
			ResourceField: "customer",
			// No ID field -> ServerGenerated or ClientBody
			IsLRO:        false,
			ResponseType: "Customer",
		},
		Update: &ProtoMethodSpec{
			Name:            "UpdateCustomer",
			InputType:       "UpdateCustomerRequest",
			ResourceField:   "customer",
			UpdateMaskField: "update_mask",
			IsLRO:           false,
			ResponseType:    "Customer",
		},
		Delete: &ProtoMethodSpec{
			Name:         "DeleteCustomer",
			InputType:    "DeleteCustomerRequest",
			NameField:    "name",
			IsLRO:        false,
			ResponseType: "google.protobuf.Empty",
		},
	}

	krmSpec := &KRMSpecInfo{
		IdentityTypeName:     "CloudchannelCustomerIdentity",
		HasGetIdentity:       true,
		HasIdentitySpecified: true, // Signals ServerGenerated
		ParentStringExpr:     "ParentString()",
		IDExpr:               "Customer",
		HasObservedState:     true,
	}

	profile := BuildProfile("cloudchannel", "v1alpha1", "CloudchannelCustomer", "Customer", "v1", protoSpec, krmSpec)

	if profile.Identity.Strategy != IdentityServerGenerated {
		t.Fatalf("got Identity.Strategy = %s, want %s", profile.Identity.Strategy, IdentityServerGenerated)
	}

	src, err := RenderController(profile)
	if err != nil {
		t.Fatalf("RenderController failed: %v\nCode:\n%s", err, string(src))
	}

	srcStr := string(src)
	if !strings.Contains(srcStr, "!a.id.HasIdentitySpecified()") {
		t.Errorf("expected Find() to guard on HasIdentitySpecified()")
	}
	if !strings.Contains(srcStr, "status.ExternalRef = direct.LazyPtr(created.GetName())") {
		t.Errorf("expected Create() to capture externalRef from created.GetName()")
	}
}

func TestBuildProfile_FullReplacementUpdate(t *testing.T) {
	protoSpec := &ProtoServiceSpec{
		ServiceName:  "CloudChannelService",
		ServiceHost:  "cloudchannel.googleapis.com",
		GoPackage:    "cloud.google.com/go/channel/apiv1/channelpb;channelpb",
		MessageName:  "ChannelPartnerRepricingConfig",
		HasNameField: true,
		Get: &ProtoMethodSpec{
			Name:      "GetChannelPartnerRepricingConfig",
			InputType: "GetChannelPartnerRepricingConfigRequest",
			NameField: "name",
		},
		Create: &ProtoMethodSpec{
			Name:          "CreateChannelPartnerRepricingConfig",
			InputType:     "CreateChannelPartnerRepricingConfigRequest",
			ParentField:   "parent",
			ResourceField: "channel_partner_repricing_config",
			IsLRO:         false,
			ResponseType:  "ChannelPartnerRepricingConfig",
		},
		Update: &ProtoMethodSpec{
			Name:          "UpdateChannelPartnerRepricingConfig",
			InputType:     "UpdateChannelPartnerRepricingConfigRequest",
			ResourceField: "channel_partner_repricing_config",
			// No UpdateMaskField!
			IsLRO:        false,
			ResponseType: "ChannelPartnerRepricingConfig",
		},
		Delete: &ProtoMethodSpec{
			Name:         "DeleteChannelPartnerRepricingConfig",
			InputType:    "DeleteChannelPartnerRepricingConfigRequest",
			NameField:    "name",
			IsLRO:        false,
			ResponseType: "google.protobuf.Empty",
		},
	}

	krmSpec := &KRMSpecInfo{
		IdentityTypeName: "CloudchannelChannelPartnerRepricingConfigIdentity",
		HasGetIdentity:   true,
		ParentStringExpr: "ParentString()",
		IDExpr:           "ChannelPartnerRepricingConfig",
		HasObservedState: true,
	}

	profile := BuildProfile("cloudchannel", "v1alpha1", "CloudchannelChannelPartnerRepricingConfig", "ChannelPartnerRepricingConfig", "v1", protoSpec, krmSpec)

	if profile.Lifecycle.UpdateStyle != UpdateStyleFullReplacement {
		t.Fatalf("got UpdateStyle = %s, want %s", profile.Lifecycle.UpdateStyle, UpdateStyleFullReplacement)
	}

	src, err := RenderController(profile)
	if err != nil {
		t.Fatalf("RenderController failed: %v\nCode:\n%s", err, string(src))
	}

	srcStr := string(src)
	if !strings.Contains(srcStr, "resource, report, _, err := common.CompareGreenfieldSpec(") {
		t.Errorf("expected update comparison to omit updateMask receiver, got code: %s", srcStr)
	}
}

func TestBuildProfile_UndeletableResource(t *testing.T) {
	protoSpec := &ProtoServiceSpec{
		ServiceName:  "KeyManagementServiceClient",
		ServiceHost:  "cloudkms.googleapis.com",
		GoPackage:    "cloud.google.com/go/kms/apiv1/kmspb;kmspb",
		MessageName:  "EkmConnection",
		HasNameField: true,
		Get: &ProtoMethodSpec{
			Name:      "GetEkmConnection",
			InputType: "GetEkmConnectionRequest",
			NameField: "name",
		},
		Create: &ProtoMethodSpec{
			Name:          "CreateEkmConnection",
			InputType:     "CreateEkmConnectionRequest",
			ParentField:   "parent",
			ResourceField: "ekm_connection",
			IDField:       "ekm_connection_id",
			IsLRO:         false,
			ResponseType:  "EkmConnection",
		},
		Update: &ProtoMethodSpec{
			Name:            "UpdateEkmConnection",
			InputType:       "UpdateEkmConnectionRequest",
			ResourceField:   "ekm_connection",
			UpdateMaskField: "update_mask",
			IsLRO:           false,
			ResponseType:    "EkmConnection",
		},
		// No Delete method!
	}

	krmSpec := &KRMSpecInfo{
		IdentityTypeName:    "KMSEkmConnectionIdentity",
		HasGetIdentity:      true,
		ParentStringExpr:    "ParentString()",
		IDExpr:              "EkmConnection",
		HasObservedState:    true,
		MapperVersionSuffix: "_v1alpha1",
	}

	profile := BuildProfile("kms", "v1alpha1", "KMSEkmConnection", "EkmConnection", "v1", protoSpec, krmSpec)

	if profile.Lifecycle.DeleteStyle != DeleteStyleUnsupported {
		t.Fatalf("got DeleteStyle = %s, want %s", profile.Lifecycle.DeleteStyle, DeleteStyleUnsupported)
	}

	src, err := RenderController(profile)
	if err != nil {
		t.Fatalf("RenderController failed: %v\nCode:\n%s", err, string(src))
	}

	srcStr := string(src)
	if !strings.Contains(srcStr, "deletion not supported for EkmConnection, removing finalizer without deleting GCP resource") {
		t.Errorf("expected Delete() to log unsupported deletion and return (true, nil)")
	}
}
