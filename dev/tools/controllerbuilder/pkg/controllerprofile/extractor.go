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
	"path"
	"strings"
	"unicode"
)

// ProtoMethodSpec describes a standard RPC method on a GCP service.
type ProtoMethodSpec struct {
	Name            string
	InputType       string
	ParentField     string
	ResourceField   string
	IDField         string
	NameField       string
	UpdateMaskField string
	ForceField      string
	IsLRO           bool
	ResponseType    string // Full name of synchronous response or LRO result
}

// ProtoServiceSpec captures the protobuf descriptor facts needed to construct a profile.
type ProtoServiceSpec struct {
	ServiceName  string // e.g., "FeatureOnlineStoreAdminService"
	ServiceHost  string // e.g., "aiplatform.googleapis.com"
	GoPackage    string // e.g., "cloud.google.com/go/aiplatform/apiv1/aiplatformpb;aiplatformpb"
	MessageName  string // e.g., "FeatureView"
	HasNameField bool
	HasLabels    bool

	Get    *ProtoMethodSpec
	Create *ProtoMethodSpec
	Update *ProtoMethodSpec
	Delete *ProtoMethodSpec
}

// KRMSpecInfo captures AST facts discovered from the KRM API and direct controller directories.
type KRMSpecInfo struct {
	IdentityTypeName     string // e.g., "AIPlatformFeatureViewIdentity"
	HasGetIdentity       bool   // (*Kind) GetIdentity(ctx, reader) exists
	HasIdentitySpecified bool   // (*Identity) HasIdentitySpecified() exists
	HasFromExternal      bool   // (*Identity) FromExternal(string) exists
	ParentStringExpr     string // e.g., "a.id.ParentString()"
	IDExpr               string // e.g., "a.id.ID()"

	HasSpecLabels     bool   // Spec has Labels map
	HasSpecReferences bool   // Spec has *Ref fields
	HasObservedState  bool   // Status has ObservedState
	LocationExpr      string // e.g., "direct.ValueOf(obj.Spec.Location)"

	MapperVersionSuffix string // "" or "_v1alpha1"
	RegionalEndpoint    string // "" | "prefix" | "rep"
	UseGRPCClient       bool   // true if package uses GRPCClientOptions()
}

// BuildProfile synthesizes a complete ControllerProfile from protobuf service descriptors and KRM AST facts.
func BuildProfile(
	kccService, kccVersion, kind, protoResource, protoVersion string,
	protoSpec *ProtoServiceSpec,
	krmSpec *KRMSpecInfo,
) *ControllerProfile {
	profile := &ControllerProfile{
		KCCService:    kccService,
		KCCVersion:    kccVersion,
		Kind:          kind,
		ProtoResource: protoResource,
		ProtoVersion:  protoVersion,
	}

	profile.Lifecycle = extractLifecycleTrait(protoSpec, protoResource)
	profile.Identity = extractIdentityTrait(protoSpec, krmSpec, &profile.Lifecycle)
	profile.Execution = extractExecutionTrait(protoSpec)
	profile.Transport = extractTransportTrait(protoSpec, krmSpec)

	return profile
}

func extractLifecycleTrait(spec *ProtoServiceSpec, protoResource string) LifecycleTrait {
	var trait LifecycleTrait
	if spec == nil {
		return trait
	}

	trait.HasGet = spec.Get != nil
	if spec.Get != nil {
		trait.GetMethodName = spec.Get.Name
		trait.GetRequestType = spec.Get.InputType
		trait.GetNameField = protoToGoName(spec.Get.NameField, "Name")
	}

	trait.CanCreate = spec.Create != nil
	trait.IsSingleton = spec.Get != nil && spec.Create == nil && spec.Update != nil && spec.Delete == nil

	if spec.Create != nil {
		trait.CreateMethodName = spec.Create.Name
		trait.CreateRequestType = spec.Create.InputType
		trait.CreateParentField = protoToGoName(spec.Create.ParentField, "Parent")
		trait.CreateBodyField = protoToGoName(spec.Create.ResourceField, protoResource)
		trait.CreateIDField = protoToGoName(spec.Create.IDField, "")
	}

	if spec.Update != nil {
		trait.UpdateMethodName = spec.Update.Name
		trait.UpdateRequestType = spec.Update.InputType
		trait.UpdateBodyField = protoToGoName(spec.Update.ResourceField, protoResource)
		trait.UpdateHasTopLevelName = spec.Update.NameField != "" && spec.Update.ResourceField != ""
		if spec.Update.UpdateMaskField != "" {
			trait.UpdateStyle = UpdateStyleMask
			trait.UpdateMaskField = protoToGoName(spec.Update.UpdateMaskField, "UpdateMask")
		} else {
			trait.UpdateStyle = UpdateStyleFullReplacement
		}
	} else {
		trait.UpdateStyle = UpdateStyleUnsupported
	}

	if spec.Delete != nil {
		trait.DeleteStyle = DeleteStyleStandard
		trait.DeleteMethodName = spec.Delete.Name
		trait.DeleteRequestType = spec.Delete.InputType
		trait.DeleteNameField = protoToGoName(spec.Delete.NameField, "Name")
		trait.DeleteForceField = protoToGoName(spec.Delete.ForceField, "")
	} else {
		trait.DeleteStyle = DeleteStyleUnsupported
	}

	return trait
}

func extractIdentityTrait(protoSpec *ProtoServiceSpec, krmSpec *KRMSpecInfo, lifecycle *LifecycleTrait) IdentityTrait {
	var trait IdentityTrait
	if protoSpec != nil && protoSpec.HasNameField {
		trait.NameField = "Name"
	}

	if krmSpec != nil {
		if krmSpec.IdentityTypeName != "" {
			trait.TypeName = "krm." + krmSpec.IdentityTypeName
		}
		trait.UseGetIdentity = krmSpec.HasGetIdentity
		trait.ParentStringExpr = krmSpec.ParentStringExpr
		trait.IDExpr = krmSpec.IDExpr
		trait.HasFromExternal = krmSpec.HasFromExternal
	}

	if lifecycle.IsSingleton {
		trait.Strategy = IdentitySingleton
		return trait
	}

	if !lifecycle.CanCreate {
		trait.Strategy = IdentityClientField
		return trait
	}

	if lifecycle.CreateIDField != "" {
		trait.Strategy = IdentityClientField
		return trait
	}

	// CreateRequest has no designated ID field
	if krmSpec != nil && krmSpec.HasIdentitySpecified {
		trait.Strategy = IdentityServerGenerated
	} else {
		trait.Strategy = IdentityClientBody
	}

	return trait
}

func extractExecutionTrait(spec *ProtoServiceSpec) ExecutionTrait {
	var trait ExecutionTrait
	if spec == nil {
		return trait
	}

	if spec.Create != nil {
		trait.Create = OpExecution{
			IsLRO:   spec.Create.IsLRO,
			IsEmpty: isEmptyResponse(spec.Create.ResponseType),
		}
	}
	if spec.Update != nil {
		trait.Update = OpExecution{
			IsLRO:   spec.Update.IsLRO,
			IsEmpty: isEmptyResponse(spec.Update.ResponseType),
		}
	}
	if spec.Delete != nil {
		trait.Delete = OpExecution{
			IsLRO:   spec.Delete.IsLRO,
			IsEmpty: isEmptyResponse(spec.Delete.ResponseType),
		}
	}

	return trait
}

func extractTransportTrait(protoSpec *ProtoServiceSpec, krmSpec *KRMSpecInfo) TransportTrait {
	var trait TransportTrait
	if protoSpec != nil {
		trait.ServiceHost = protoSpec.ServiceHost
		deriveClientAndPBImports(&trait, protoSpec)
	}

	hasProtoLabels := protoSpec != nil && protoSpec.HasLabels
	hasKRMLabels := krmSpec != nil && krmSpec.HasSpecLabels

	if hasProtoLabels && !hasKRMLabels {
		trait.LabelDiffing = LabelsMetadata
	} else if hasKRMLabels {
		trait.LabelDiffing = LabelsSpec
	} else {
		trait.LabelDiffing = LabelsNone
	}

	if krmSpec != nil {
		trait.HasReferences = krmSpec.HasSpecReferences
		trait.HasObservedState = krmSpec.HasObservedState
		trait.MapperVersionSuffix = krmSpec.MapperVersionSuffix
		trait.LocationExpr = krmSpec.LocationExpr
		if krmSpec.UseGRPCClient {
			trait.UseGRPCClient = true
			trait.GCPClientConstructor = strings.Replace(trait.GCPClientConstructor, "RESTClient", "Client", 1)
		}
		if krmSpec.RegionalEndpoint != "" {
			trait.RegionalEndpoint = krmSpec.RegionalEndpoint
		}
	}

	return trait
}

func deriveClientAndPBImports(trait *TransportTrait, spec *ProtoServiceSpec) {
	if spec.GoPackage == "" {
		return
	}
	parts := strings.Split(spec.GoPackage, ";")
	pbImport := strings.TrimSpace(parts[0])
	pbAlias := path.Base(pbImport)
	if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
		pbAlias = strings.TrimSpace(parts[1])
	}

	gcpClientImport := path.Dir(pbImport)
	goPkgShortName := path.Base(path.Dir(gcpClientImport))

	svcPrefix := strings.TrimSuffix(spec.ServiceName, "Service")
	if strings.EqualFold(svcPrefix, goPkgShortName) {
		svcPrefix = ""
	}

	trait.PBImportPath = pbImport
	trait.PBPackageAlias = pbAlias
	trait.GCPClientImportPath = gcpClientImport
	trait.GCPClientType = "*gcp." + svcPrefix + "Client"
	trait.GCPClientConstructor = "gcp.New" + svcPrefix + "RESTClient"
}

func isEmptyResponse(typeName string) bool {
	trimmed := strings.TrimPrefix(typeName, ".")
	return trimmed == "google.protobuf.Empty" || trimmed == "Empty"
}

func protoToGoName(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return snakeToCamel(s)
}

func snakeToCamel(s string) string {
	var result strings.Builder
	capitalize := true
	for _, r := range s {
		if r == '_' {
			capitalize = true
			continue
		}
		if capitalize {
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
