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
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// KRMPackageInfo captures structural facts inspected from the KRM API package
// (apis/<service>/<version>) and existing direct controller package
// (pkg/controller/direct/<service>) to specialize the generated controller.
type KRMPackageInfo struct {
	// Identity facts
	IdentityTypeName     string // e.g., "BeyondCorpAppConnectionIdentity" or "AppConnectionIdentity"
	HasGetIdentity       bool   // true if (*<Kind>) GetIdentity(ctx, reader) exists
	HasIdentitySpecified bool   // true if (*<IdentityType>) HasIdentitySpecified() exists
	HasFromExternal      bool   // true if (*<IdentityType>) FromExternal(string) exists
	ParentStringExpr     string // e.g., "a.id.ParentString()" or "a.id.Parent().String()"
	IDExpr               string // e.g., "a.id.ID()" or "a.id.Connection"

	// Spec & Status facts
	HasSpecLabels     bool   // true if <Kind>Spec has a Labels field
	HasSpecReferences bool   // true if <Kind>Spec or nested structs have *Ref / *Refs fields
	StatusInspected   bool   // true if <Kind>Status struct was found
	HasObservedState  bool   // true if <Kind>Status has an ObservedState field
	LocationExpr      string // e.g., "direct.ValueOf(obj.Spec.Location)" or "obj.Spec.Location"

	// Controller package facts
	MapperVersionSuffix      string // "" or "_v1alpha1" / "_v1beta1"
	DetectedRegionalEndpoint string // "" or "prefix" / "rep" inferred from sibling controllers
	UseGRPCClient            bool   // true if sibling controllers in pkg/controller/direct/<service> use GRPCClientOptions()
}

// ApplyDeterministicPasses runs the 5-pass deterministic specialization pipeline
// over a ControllerArgs struct using the proto ResourceAPI and KRM package AST info.
func ApplyDeterministicPasses(
	args *ccTemplate.ControllerArgs,
	api *protoapi.ResourceAPI,
	krmInfo *KRMPackageInfo,
	regionalEndpoint string,
) {
	args.EnsureDefaults()
	args.PassesApplied = true

	ApplyMethodPruningPass(args, api)
	ApplyIDStrategyPass(args, api, krmInfo)
	ApplyUpdateMaskPass(args, api)
	ApplyLROAndGetPass(args, api)
	ApplyProtoAndServiceTraitsPass(args, api, krmInfo, regionalEndpoint)
}

// ApplyMethodPruningPass (Pass 1) specializes CRUD method presence and request/method
// names based on the resource's proto service methods (Archetypes A, B, D, G).
func ApplyMethodPruningPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if api == nil {
		return
	}

	if api.Get != nil {
		if api.Get.Descriptor != nil {
			args.GetMethodName = string(api.Get.Descriptor.Name())
			args.GetRequestType = string(api.Get.Descriptor.Input().Name())
		}
		args.GetNameField = protoFieldNameToGoName(api.Get.NameField, "Name")
	}

	args.HasCreate = api.Create != nil
	args.IsSingleton = api.Get != nil && !args.HasCreate && api.Update != nil && api.Delete == nil

	if api.Create != nil {
		if api.Create.Descriptor != nil {
			args.CreateMethodName = string(api.Create.Descriptor.Name())
			args.CreateRequestType = string(api.Create.Descriptor.Input().Name())
		}
		args.CreateParentField = protoFieldNameToGoName(api.Create.ParentField, "")
		args.CreateBodyField = protoFieldNameToGoName(api.Create.ResourceField, args.ProtoResource)
		args.CreateIDField = protoFieldNameToGoName(api.Create.IDField, "")
	} else {
		args.CreateMethodName = ""
		args.CreateRequestType = ""
		args.CreateParentField = ""
		args.CreateBodyField = ""
		args.CreateIDField = ""
		args.CreateIsLRO = false
	}

	args.HasUpdate = api.Update != nil
	if api.Update != nil {
		if api.Update.Descriptor != nil {
			args.UpdateMethodName = string(api.Update.Descriptor.Name())
			args.UpdateRequestType = string(api.Update.Descriptor.Input().Name())
		}
		args.UpdateBodyField = protoFieldNameToGoName(api.Update.ResourceField, args.ProtoResource)
		args.UpdateHasTopLevelName = api.Update.NameField != "" && api.Update.ResourceField != ""
	} else {
		args.UpdateMethodName = ""
		args.UpdateRequestType = ""
		args.UpdateBodyField = ""
		args.UpdateMaskField = ""
		args.UpdateHasTopLevelName = false
		args.UpdateIsLRO = false
	}

	args.HasDelete = api.Delete != nil
	if api.Delete != nil {
		if api.Delete.Descriptor != nil {
			args.DeleteMethodName = string(api.Delete.Descriptor.Name())
			args.DeleteRequestType = string(api.Delete.Descriptor.Input().Name())
		}
		args.DeleteNameField = protoFieldNameToGoName(api.Delete.NameField, "Name")
		args.DeleteForceField = protoFieldNameToGoName(api.Delete.ForceField, "")
	} else {
		args.DeleteMethodName = ""
		args.DeleteRequestType = ""
		args.DeleteNameField = ""
		args.DeleteForceField = ""
		args.DeleteIsLRO = false
	}
}

// ApplyIDStrategyPass (Pass 2) specializes how the resource ID is passed on Create
// and how the Identity type is constructed (Standard <resource>_id vs. Archetype C
// Variant 3A resource.Name vs. Variant 3B server-generated ID).
func ApplyIDStrategyPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI, krmInfo *KRMPackageInfo) {
	if api != nil && api.Message != nil {
		if nameFd := api.Message.Fields().ByName("name"); nameFd != nil {
			args.ResourceNameField = protoCamelCase(string(nameFd.Name()))
		} else {
			args.ResourceNameField = ""
		}
	}

	if krmInfo != nil {
		if krmInfo.IdentityTypeName != "" {
			args.IdentityType = "krm." + krmInfo.IdentityTypeName
		}
		args.UseGetIdentity = krmInfo.HasGetIdentity
		if krmInfo.ParentStringExpr != "" {
			args.ParentStringExpr = krmInfo.ParentStringExpr
		}
		if krmInfo.IDExpr != "" {
			args.IDExpr = krmInfo.IDExpr
		}
		args.HasIdentityFromExternal = krmInfo.HasFromExternal
	}

	if !args.HasCreate {
		args.IsServerGeneratedID = false
		args.SetResourceNameInCreate = false
		return
	}

	if args.CreateIDField != "" {
		// Case 1: Standard <resource>_id field on CreateRequest.
		args.IsServerGeneratedID = false
		args.SetResourceNameInCreate = false
		return
	}

	// Case 2 (Archetype C): CreateRequest has no <resource>_id field.
	if krmInfo != nil && krmInfo.HasIdentitySpecified {
		// Variant 3B: Server-generated ID.
		args.IsServerGeneratedID = true
		args.SetResourceNameInCreate = false
	} else {
		// Variant 3A: Client-specified resource name in body.
		args.IsServerGeneratedID = false
		args.SetResourceNameInCreate = args.ResourceNameField != ""
	}
}

// ApplyUpdateMaskPass (Pass 3) specializes UpdateRequest construction for
// standard UpdateMask vs. full-replacement Update without UpdateMask (Archetype F).
func ApplyUpdateMaskPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if !args.HasUpdate || api == nil || api.Update == nil {
		args.UpdateMaskField = ""
		return
	}
	if api.Update.UpdateMaskField != "" {
		args.UpdateMaskField = protoFieldNameToGoName(api.Update.UpdateMaskField, "UpdateMask")
	} else {
		args.UpdateMaskField = ""
	}
}

// ApplyLROAndGetPass (Pass 4) specializes Create, Update, and Delete for
// synchronous vs. Long-Running Operation (LRO) execution and Empty responses.
func ApplyLROAndGetPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if api == nil {
		return
	}

	if api.Create != nil {
		args.CreateIsLRO = api.Create.LRO != nil
		if api.Create.LRO != nil {
			args.CreateLROIsEmpty = isEmptyProtoType(api.Create.LRO.ResponseType)
			args.CreateSyncIsEmpty = false
		} else {
			args.CreateLROIsEmpty = false
			args.CreateSyncIsEmpty = isEmptyProtoType(api.Create.Response)
		}
	}

	if api.Update != nil {
		args.UpdateIsLRO = api.Update.LRO != nil
		if api.Update.LRO != nil {
			args.UpdateLROIsEmpty = isEmptyProtoType(api.Update.LRO.ResponseType)
			args.UpdateSyncIsEmpty = false
		} else {
			args.UpdateLROIsEmpty = false
			args.UpdateSyncIsEmpty = isEmptyProtoType(api.Update.Response)
		}
	}

	if api.Delete != nil {
		args.DeleteIsLRO = api.Delete.LRO != nil
		if api.Delete.LRO != nil {
			args.DeleteLROIsEmpty = isEmptyProtoType(api.Delete.LRO.ResponseType)
			args.DeleteSyncIsEmpty = false
		} else {
			args.DeleteLROIsEmpty = false
			args.DeleteSyncIsEmpty = isEmptyProtoType(api.Delete.Response)
		}
	}
}

// ApplyProtoAndServiceTraitsPass (Pass 5) specializes GAPIC client/protobuf imports,
// client constructor, labels diffing, reference normalization, ObservedState,
// and regional endpoint handling.
func ApplyProtoAndServiceTraitsPass(
	args *ccTemplate.ControllerArgs,
	api *protoapi.ResourceAPI,
	krmInfo *KRMPackageInfo,
	regionalEndpoint string,
) {
	if api != nil {
		if api.Service != nil {
			deriveGAPICClientAndPBImports(args, api.Service)
		}
		if api.DefaultHost != "" {
			args.ServiceHost = api.DefaultHost
		}
	}

	hasProtoLabels := api != nil && (api.HasLabels || (api.Message != nil && api.Message.Fields().ByName("labels") != nil))
	hasKRMLabels := krmInfo != nil && krmInfo.HasSpecLabels
	// Use CompareGreenfieldSpecAndLabels (metadata.labels) when the proto has labels
	// that are NOT already mapped via Spec.Labels.
	args.HasLabels = hasProtoLabels && !hasKRMLabels

	if krmInfo != nil {
		args.HasReferences = krmInfo.HasSpecReferences
		if krmInfo.StatusInspected {
			args.HasObservedState = krmInfo.HasObservedState
		}
		if krmInfo.MapperVersionSuffix != "" {
			args.MapperVersionSuffix = krmInfo.MapperVersionSuffix
		}
		if krmInfo.LocationExpr != "" {
			args.LocationExpr = krmInfo.LocationExpr
		}
		if krmInfo.UseGRPCClient {
			args.UseGRPCClient = true
			args.GCPClientConstructor = strings.Replace(args.GCPClientConstructor, "RESTClient", "Client", 1)
		}
	}

	if regionalEndpoint != "" {
		args.RegionalEndpoint = regionalEndpoint
	} else if krmInfo != nil && krmInfo.DetectedRegionalEndpoint != "" {
		args.RegionalEndpoint = krmInfo.DetectedRegionalEndpoint
	}
}

// deriveGAPICClientAndPBImports populates GCPClientImportPath, GCPClientType,
// GCPClientConstructor, PBImportPath, and PBPackageAlias from a ServiceDescriptor.
func deriveGAPICClientAndPBImports(args *ccTemplate.ControllerArgs, svc protoreflect.ServiceDescriptor) {
	fileOpts, ok := svc.ParentFile().Options().(*descriptorpb.FileOptions)
	if !ok || fileOpts == nil {
		return
	}
	goPkg := fileOpts.GetGoPackage()
	if goPkg == "" {
		return
	}
	parts := strings.Split(goPkg, ";")
	pbImportPath := strings.TrimSpace(parts[0])
	if pbImportPath == "" {
		return
	}
	pbAlias := path.Base(pbImportPath)
	if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
		pbAlias = strings.TrimSpace(parts[1])
	}
	gcpClientImportPath := path.Dir(pbImportPath)
	goPkgShortName := path.Base(path.Dir(gcpClientImportPath))

	svcName := string(svc.Name())
	svcPrefix := strings.TrimSuffix(svcName, "Service")
	if strings.EqualFold(svcPrefix, goPkgShortName) {
		svcPrefix = ""
	}

	args.PBImportPath = pbImportPath
	args.PBPackageAlias = pbAlias
	args.GCPClientImportPath = gcpClientImportPath
	args.GCPClientType = "*gcp." + svcPrefix + "Client"
	args.GCPClientConstructor = "gcp.New" + svcPrefix + "RESTClient"
}

// InspectKRMPackageForController parses the Go files in apis/<service>/<version>
// and pkg/controller/direct/<service> to extract Identity, Spec, Status, Mapper,
// and RegionalEndpoint facts for deterministic controller generation.
func InspectKRMPackageForController(rootPath, serviceName, version, kind, protoResource string) *KRMPackageInfo {
	info := &KRMPackageInfo{}
	apiDir := filepath.Join(rootPath, "apis", serviceName, version)
	directDir := filepath.Join(rootPath, "pkg", "controller", "direct", serviceName)

	inspectAPIDir(apiDir, kind, protoResource, info)
	inspectDirectDir(directDir, kind, version, info)

	// Also check generate.sh for --multiversion if MapperVersionSuffix was not yet found
	if info.MapperVersionSuffix == "" {
		for _, genShPath := range []string{
			filepath.Join(apiDir, "generate.sh"),
			filepath.Join(directDir, "generate.sh"),
		} {
			if data, err := os.ReadFile(genShPath); err == nil {
				if strings.Contains(string(data), "--multiversion") {
					info.MapperVersionSuffix = "_" + version
					break
				}
			}
		}
	}

	return info
}

func inspectAPIDir(apiDir, kind, protoResource string, info *KRMPackageInfo) {
	entries, err := os.ReadDir(apiDir)
	if err != nil {
		return
	}

	fset := token.NewFileSet()
	var files []*ast.File
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(apiDir, entry.Name()), nil, 0)
		if err == nil {
			files = append(files, f)
		}
	}

	kindIdentityName := kind + "Identity"
	protoIdentityName := protoResource + "Identity"

	structDecls := make(map[string]*ast.StructType)
	methodsByRecv := make(map[string]map[string]*ast.FuncDecl)

	for _, f := range files {
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				if d.Tok != token.TYPE {
					continue
				}
				for _, spec := range d.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if st, ok := ts.Type.(*ast.StructType); ok {
						structDecls[ts.Name.Name] = st
					}
				}
			case *ast.FuncDecl:
				if d.Recv == nil || len(d.Recv.List) == 0 {
					continue
				}
				recvName := receiverTypeName(d.Recv.List[0].Type)
				if recvName == "" {
					continue
				}
				if methodsByRecv[recvName] == nil {
					methodsByRecv[recvName] = make(map[string]*ast.FuncDecl)
				}
				methodsByRecv[recvName][d.Name.Name] = d
			}
		}
	}

	// 1. Check GetIdentity on Kind
	if kindMethods := methodsByRecv[kind]; kindMethods != nil {
		if _, ok := kindMethods["GetIdentity"]; ok {
			info.HasGetIdentity = true
		}
	}

	// 2. Inspect Identity struct & methods (prefer <Kind>Identity if present, else <ProtoResource>Identity)
	chosenIdentity := ""
	if _, ok := structDecls[kindIdentityName]; ok {
		chosenIdentity = kindIdentityName
	} else if _, ok := structDecls[protoIdentityName]; ok {
		chosenIdentity = protoIdentityName
	} else if _, ok := methodsByRecv[kindIdentityName]; ok {
		chosenIdentity = kindIdentityName
	} else if _, ok := methodsByRecv[protoIdentityName]; ok {
		chosenIdentity = protoIdentityName
	}

	if chosenIdentity != "" {
		info.IdentityTypeName = chosenIdentity
		idMethods := methodsByRecv[chosenIdentity]
		if idMethods != nil {
			if _, ok := idMethods["HasIdentitySpecified"]; ok {
				info.HasIdentitySpecified = true
			}
			if _, ok := idMethods["FromExternal"]; ok {
				info.HasFromExternal = true
			}
			if _, ok := idMethods["ParentString"]; ok {
				info.ParentStringExpr = "a.id.ParentString()"
			} else if _, ok := idMethods["Parent"]; ok {
				info.ParentStringExpr = "a.id.Parent().String()"
			}
			if _, ok := idMethods["ID"]; ok {
				info.IDExpr = "a.id.ID()"
			}
		}

		// If ID() method does not exist on chosenIdentity, check exported string fields on the Identity struct
		if info.IDExpr == "" {
			if st := structDecls[chosenIdentity]; st != nil && st.Fields != nil {
				var lastExportedStringField string
				for _, field := range st.Fields.List {
					if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "string" {
						for _, name := range field.Names {
							if name.IsExported() {
								lastExportedStringField = name.Name
							}
						}
					}
				}
				if lastExportedStringField != "" {
					info.IDExpr = "a.id." + lastExportedStringField
				}
			}
		}
	}

	// 3. Inspect <Kind>Spec for Labels, Location, and Reference fields
	specName := kind + "Spec"
	if specStruct := structDecls[specName]; specStruct != nil && specStruct.Fields != nil {
		for _, field := range specStruct.Fields.List {
			for _, name := range field.Names {
				if name.Name == "Labels" {
					info.HasSpecLabels = true
				}
				if name.Name == "Location" {
					if _, isStar := field.Type.(*ast.StarExpr); isStar {
						info.LocationExpr = "direct.ValueOf(obj.Spec.Location)"
					} else {
						info.LocationExpr = "obj.Spec.Location"
					}
				}
			}
		}
		visited := make(map[string]bool)
		info.HasSpecReferences = structHasResourceReferences(specName, structDecls, visited)
	}

	// 4. Inspect <Kind>Status for ObservedState
	statusName := kind + "Status"
	if statusStruct := structDecls[statusName]; statusStruct != nil && statusStruct.Fields != nil {
		info.StatusInspected = true
		for _, field := range statusStruct.Fields.List {
			for _, name := range field.Names {
				if name.Name == "ObservedState" {
					info.HasObservedState = true
				}
			}
		}
	}
}

func structHasResourceReferences(structName string, structDecls map[string]*ast.StructType, visited map[string]bool) bool {
	if visited[structName] {
		return false
	}
	visited[structName] = true
	st := structDecls[structName]
	if st == nil || st.Fields == nil {
		return false
	}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			if isNonRootRefFieldName(name.Name) {
				return true
			}
		}
		if elemName := baseIdentTypeName(field.Type); elemName != "" {
			if structHasResourceReferences(elemName, structDecls, visited) {
				return true
			}
		}
	}
	return false
}

func isNonRootRefFieldName(name string) bool {
	switch name {
	case "ProjectRef", "FolderRef", "OrganizationRef", "BillingAccountRef":
		return false
	}
	return strings.HasSuffix(name, "Ref") || strings.HasSuffix(name, "Refs")
}

func inspectDirectDir(directDir, kind, version string, info *KRMPackageInfo) {
	versionedToProto := kind + "Spec_" + version + "_ToProto"
	hasGRPCClient := false
	hasRESTClient := false
	_ = filepath.WalkDir(directDir, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return nil
		}
		content := string(data)
		if strings.Contains(content, versionedToProto) {
			info.MapperVersionSuffix = "_" + version
		}
		if strings.Contains(content, "GRPCClientOptions()") {
			hasGRPCClient = true
		}
		if strings.Contains(content, "RESTClientOptions()") {
			hasRESTClient = true
		}
		if strings.Contains(content, ".rep.googleapis.com") {
			info.DetectedRegionalEndpoint = "rep"
		} else if info.DetectedRegionalEndpoint == "" {
			if (strings.Contains(content, "option.WithEndpoint(") && strings.Contains(content, "-"+strings.ToLower(filepath.Base(directDir)))) ||
				strings.Contains(content, "option.WithEndpoint(fmt.Sprintf(\"%s-") {
				info.DetectedRegionalEndpoint = "prefix"
			}
		}
		return nil
	})
	if hasGRPCClient && !hasRESTClient {
		info.UseGRPCClient = true
	}
}

func receiverTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

func baseIdentTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return baseIdentTypeName(t.X)
	case *ast.ArrayType:
		return baseIdentTypeName(t.Elt)
	}
	return ""
}

func isEmptyProtoType(typeName string) bool {
	trimmed := strings.TrimPrefix(typeName, ".")
	return trimmed == "google.protobuf.Empty"
}

func protoFieldNameToGoName(fieldName string, fallback string) string {
	if fieldName == "" {
		return fallback
	}
	return protoCamelCase(fieldName)
}

// protoCamelCase converts a proto snake_case identifier into the Go camel-case
// struct field name produced by protoc-gen-go (strs.GoCamelCase).
func protoCamelCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '_' && i+1 < len(s) && isASCIILower(s[i+1]):
			// Skip '_' and uppercase next character in the next iteration.
		case c == '_':
			// Skip '_'
		case isASCIILower(c) && (i == 0 || s[i-1] == '_' || isASCIIDigit(s[i-1])):
			b = append(b, c-('a'-'A'))
		default:
			b = append(b, c)
		}
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
