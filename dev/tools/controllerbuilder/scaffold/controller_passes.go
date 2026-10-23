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
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

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

	applyFindMethod(args, api)
	applyCreateMethod(args, api.Create)
	applyUpdateMethod(args, api.Update)
	applyDeleteMethod(args, api.Delete)

	args.IsSingleton = (args.HasGet || args.FindUsesList) && !args.HasCreate && api.Update != nil && api.Delete == nil
}

// ApplyIDStrategyPass (Pass 2) specializes how the resource ID is passed on Create
// and how the Identity type is constructed (Standard <resource>_id vs. Archetype C
// Variant 3A resource.Name vs. Variant 3B server-generated ID).
func ApplyIDStrategyPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI, krmInfo *KRMPackageInfo) {
	nameIsOutputOnly := inspectProtoResourceNameField(args, api)
	applyKRMIdentityInfo(args, krmInfo)

	if !args.HasCreate {
		args.IsServerGeneratedID = false
		args.SetResourceNameInCreate = false
		return
	}

	if krmInfo != nil && krmInfo.HasIdentitySpecified {
		args.IsServerGeneratedID = true
		args.HasIdentitySpecifiedMethod = true
		args.SetResourceNameInCreate = false
		return
	}

	if args.CreateIDField != "" || args.CreateNameField != "" {
		args.IsServerGeneratedID = false
		args.SetResourceNameInCreate = false
		return
	}

	if krmInfo != nil && krmInfo.IdentityInspected {
		if krmInfo.IdentityIsClientSpecified {
			args.IsServerGeneratedID = false
			args.SetResourceNameInCreate = args.ResourceNameField != ""
		} else {
			args.IsServerGeneratedID = true
			args.SetResourceNameInCreate = false
			args.ServerIDUsesEmptyIDGuard = !krmInfo.IdentityUsesMetadataName
		}
		return
	}

	args.IsServerGeneratedID = nameIsOutputOnly
	args.SetResourceNameInCreate = !nameIsOutputOnly && args.ResourceNameField != ""
}

// ApplyUpdateMaskPass (Pass 3) specializes UpdateRequest construction for
// standard UpdateMask vs. full-replacement Update without UpdateMask (Archetype F).
func ApplyUpdateMaskPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if !args.HasUpdate || api == nil || api.Update == nil || api.Update.UpdateMaskField == "" {
		args.UpdateMaskField = ""
		return
	}
	args.UpdateMaskField = protoFieldNameToGoName(api.Update.UpdateMaskField, "UpdateMask")
}

// ApplyLROAndGetPass (Pass 4) specializes Create, Update, and Delete for
// synchronous vs. Long-Running Operation (LRO) execution and Empty responses.
func ApplyLROAndGetPass(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if api == nil {
		return
	}
	if api.Create != nil {
		args.CreateIsLRO, args.CreateLROIsEmpty, args.CreateSyncIsEmpty = classifyMethodExecution(api.Create)
	}
	if api.Update != nil {
		args.UpdateIsLRO, args.UpdateLROIsEmpty, args.UpdateSyncIsEmpty = classifyMethodExecution(api.Update)
	}
	if api.Delete != nil {
		args.DeleteIsLRO, args.DeleteLROIsEmpty, args.DeleteSyncIsEmpty = classifyMethodExecution(api.Delete)
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

	applyKRMServiceTraits(args, krmInfo)
	resolveGAPICClientFromModCache(args, krmInfo)

	if regionalEndpoint != "" {
		args.RegionalEndpoint = regionalEndpoint
	} else if krmInfo != nil && krmInfo.DetectedRegionalEndpoint != "" && krmInfo.LocationExpr != "" {
		args.RegionalEndpoint = krmInfo.DetectedRegionalEndpoint
	}
}

// --- Pass 1 Helpers (Footnotes) ---

func applyFindMethod(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) {
	if api.Get != nil {
		args.HasGet = true
		args.FindUsesList = false
		if api.Get.Descriptor != nil {
			args.GetMethodName = string(api.Get.Descriptor.Name())
			args.GetRequestType = string(api.Get.Descriptor.Input().Name())
		}
		args.GetNameField = protoFieldNameToGoName(api.Get.NameField, "Name")
		return
	}
	if api.List != nil {
		args.HasGet = false
		args.FindUsesList = true
		if api.List.Descriptor != nil {
			args.ListMethodName = string(api.List.Descriptor.Name())
			args.ListRequestType = string(api.List.Descriptor.Input().Name())
		}
		args.ListParentField = protoFieldNameToGoName(api.List.ParentField, "Parent")
		return
	}
	args.HasGet = false
	args.FindUsesList = false
}

func applyCreateMethod(args *ccTemplate.ControllerArgs, create *protoapi.StandardMethod) {
	args.HasCreate = create != nil
	if create == nil {
		args.CreateMethodName = ""
		args.CreateRequestType = ""
		args.CreateParentField = ""
		args.CreateNameField = ""
		args.CreateBodyField = ""
		args.CreateIDField = ""
		args.CreateIsLRO = false
		return
	}

	bodyFallback := args.ProtoResource
	if create.Descriptor != nil {
		args.CreateMethodName = string(create.Descriptor.Name())
		args.CreateRequestType = string(create.Descriptor.Input().Name())
		bodyFallback = ""
	}
	args.CreateParentField = protoFieldNameToGoName(create.ParentField, "")
	args.CreateBodyField = protoFieldNameToGoName(create.ResourceField, bodyFallback)
	if args.CreateBodyField == "" && create.NameField != "" {
		args.CreateNameField = protoFieldNameToGoName(create.NameField, "")
	} else {
		args.CreateNameField = ""
	}
	args.CreateIDField = protoFieldNameToGoName(create.IDField, "")
	args.CreateIDFieldIsPointer = false
	if create.Descriptor != nil && create.IDField != "" {
		if fd := create.Descriptor.Input().Fields().ByName(protoreflect.Name(create.IDField)); fd != nil && fd.HasOptionalKeyword() {
			args.CreateIDFieldIsPointer = true
		}
	}
}

func applyUpdateMethod(args *ccTemplate.ControllerArgs, update *protoapi.StandardMethod) {
	args.HasUpdate = update != nil
	if update == nil {
		args.UpdateMethodName = ""
		args.UpdateRequestType = ""
		args.UpdateBodyField = ""
		args.UpdateMaskField = ""
		args.UpdateHasTopLevelName = false
		args.UpdateIsLRO = false
		return
	}
	if update.Descriptor != nil {
		args.UpdateMethodName = string(update.Descriptor.Name())
		args.UpdateRequestType = string(update.Descriptor.Input().Name())
	}
	args.UpdateBodyField = protoFieldNameToGoName(update.ResourceField, args.ProtoResource)
	args.UpdateHasTopLevelName = update.NameField != "" && update.ResourceField != ""
}

func applyDeleteMethod(args *ccTemplate.ControllerArgs, del *protoapi.StandardMethod) {
	args.HasDelete = del != nil
	if del == nil {
		args.DeleteMethodName = ""
		args.DeleteRequestType = ""
		args.DeleteNameField = ""
		args.DeleteIsLRO = false
		return
	}
	if del.Descriptor != nil {
		args.DeleteMethodName = string(del.Descriptor.Name())
		args.DeleteRequestType = string(del.Descriptor.Input().Name())
	}
	args.DeleteNameField = protoFieldNameToGoName(del.NameField, "Name")
}

// --- Pass 2 Helpers (Footnotes) ---

func inspectProtoResourceNameField(args *ccTemplate.ControllerArgs, api *protoapi.ResourceAPI) (nameIsOutputOnly bool) {
	if api == nil || api.Message == nil {
		return false
	}
	nameFd := api.Message.Fields().ByName("name")
	if nameFd == nil {
		args.ResourceNameField = ""
		return false
	}
	args.ResourceNameField = protoCamelCase(string(nameFd.Name()))
	return hasProtoFieldBehavior(nameFd, annotations.FieldBehavior_OUTPUT_ONLY)
}

func applyKRMIdentityInfo(args *ccTemplate.ControllerArgs, krmInfo *KRMPackageInfo) {
	if krmInfo == nil {
		return
	}
	if krmInfo.IdentityImportPath != "" {
		args.IdentityImportPath = krmInfo.IdentityImportPath
	}
	if krmInfo.IdentityTypeName != "" {
		prefix := krmInfo.IdentityPkgPrefix
		if prefix == "" {
			prefix = "krm."
		}
		args.IdentityType = prefix + krmInfo.IdentityTypeName
	}
	if krmInfo.IdentityConstructorExpr != "" {
		args.IdentityConstructorExpr = krmInfo.IdentityConstructorExpr
	}
	args.UseGetIdentity = krmInfo.HasGetIdentity
	if krmInfo.HasNoParent {
		args.CreateParentField = ""
	} else if krmInfo.ParentStringExpr != "" {
		args.ParentStringExpr = krmInfo.ParentStringExpr
	}
	if krmInfo.IDExpr != "" {
		args.IDExpr = krmInfo.IDExpr
	}
	args.HasIdentityFromExternal = krmInfo.HasFromExternal
	args.HasIdentitySpecifiedMethod = krmInfo.HasIdentitySpecified
}

func hasProtoFieldBehavior(fd protoreflect.FieldDescriptor, want annotations.FieldBehavior) bool {
	if fd == nil || fd.Options() == nil {
		return false
	}
	ext := proto.GetExtension(fd.Options(), annotations.E_FieldBehavior)
	behaviors, ok := ext.([]annotations.FieldBehavior)
	if !ok {
		return false
	}
	for _, b := range behaviors {
		if b == want {
			return true
		}
	}
	return false
}

// --- Pass 4 Helpers (Footnotes) ---

func classifyMethodExecution(m *protoapi.StandardMethod) (isLRO, lroIsEmpty, syncIsEmpty bool) {
	if m.LRO != nil {
		return true, isEmptyProtoType(m.LRO.ResponseType), false
	}
	return false, false, isEmptyProtoType(m.Response)
}

func isEmptyProtoType(typeName string) bool {
	return strings.TrimPrefix(typeName, ".") == "google.protobuf.Empty"
}

// --- Pass 5 Helpers (Footnotes) ---

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
	gcpClientImportPath := deriveGAPICImportFromPBImport(pbImportPath)
	goPkgShortName := path.Base(path.Dir(gcpClientImportPath))

	svcName := string(svc.Name())
	svcPrefix := strings.TrimSuffix(svcName, "ServiceV2")
	svcPrefix = strings.TrimSuffix(svcPrefix, "Service")
	if strings.EqualFold(svcPrefix, goPkgShortName) {
		svcPrefix = ""
	}

	args.PBImportPath = pbImportPath
	args.PBPackageAlias = pbAlias
	args.GCPClientImportPath = gcpClientImportPath
	args.GCPClientType = "*gcp." + svcPrefix + "Client"
	args.GCPClientConstructor = "gcp.New" + svcPrefix + "RESTClient"
}

func applyKRMServiceTraits(args *ccTemplate.ControllerArgs, krmInfo *KRMPackageInfo) {
	if krmInfo == nil {
		return
	}
	if krmInfo.PackageName != "" {
		args.PackageName = krmInfo.PackageName
	}
	if krmInfo.MapperPackageImport != "" {
		args.MapperPackageImport = krmInfo.MapperPackageImport
		args.MapperPrefix = krmInfo.MapperPrefix
	}
	if krmInfo.DetectedPBImportPath != "" && krmInfo.DetectedPBImportPath != args.PBImportPath {
		args.PBImportPath = krmInfo.DetectedPBImportPath
		if krmInfo.DetectedPBAlias != "" {
			args.PBPackageAlias = krmInfo.DetectedPBAlias
		}
		if gcpPath := deriveGAPICImportFromPBImport(krmInfo.DetectedPBImportPath); gcpPath != "" {
			args.GCPClientImportPath = gcpPath
		}
	}

	args.HasReferences = krmInfo.HasSpecReferences
	if krmInfo.StatusInspected {
		args.HasObservedState = krmInfo.HasObservedState
	}
	if krmInfo.MapperVersionSuffix != "" {
		args.MapperVersionSuffix = krmInfo.MapperVersionSuffix
	}
	if krmInfo.HasExplicitObsSuffix {
		args.ObservedStateVersionSuffix = krmInfo.ObservedStateVersionSuffix
		args.HasExplicitObservedStateSuffix = true
	}
	if krmInfo.LocationExpr != "" {
		args.LocationExpr = krmInfo.LocationExpr
	}
	if krmInfo.UseGRPCClient {
		args.UseGRPCClient = true
		args.GCPClientConstructor = strings.Replace(args.GCPClientConstructor, "RESTClient", "Client", 1)
	}
}

func deriveGAPICImportFromPBImport(pbImportPath string) string {
	if strings.HasPrefix(pbImportPath, "cloud.google.com/go/") {
		return path.Dir(pbImportPath)
	}
	if strings.HasPrefix(pbImportPath, "google.golang.org/genproto/googleapis/") {
		rel := strings.TrimPrefix(pbImportPath, "google.golang.org/genproto/googleapis/")
		dir, ver := path.Dir(rel), path.Base(rel)
		if strings.HasPrefix(ver, "v") {
			return path.Join("cloud.google.com/go", dir, "api"+ver)
		}
	}
	return path.Dir(pbImportPath)
}

func resolveGAPICClientFromModCache(args *ccTemplate.ControllerArgs, krmInfo *KRMPackageInfo) {
	if krmInfo == nil || krmInfo.RootPath == "" || args.GCPClientImportPath == "" {
		return
	}
	pkgDir := findModCachePackageDir(krmInfo.RootPath, args.GCPClientImportPath)
	if pkgDir == "" {
		return
	}
	funcs, methodRecv := scanGAPICPackageFuncsAndMethods(pkgDir)
	if len(methodRecv) == 0 {
		return
	}

	var matchedClientStruct string
	for _, m := range targetCRUDMethodNames(args) {
		if recv, ok := methodRecv[m]; ok {
			matchedClientStruct = recv
			break
		}
	}
	if matchedClientStruct == "" {
		return
	}

	args.GCPClientType = "*gcp." + matchedClientStruct
	prefix := strings.TrimSuffix(matchedClientStruct, "Client")
	restCtor := "New" + prefix + "RESTClient"
	grpcCtor := "New" + prefix + "Client"

	// Standard 1: Always prefer New*RESTClient whenever available in the GAPIC SDK.
	if funcs[restCtor] {
		args.GCPClientConstructor = "gcp." + restCtor
		args.UseGRPCClient = false
	} else if funcs[grpcCtor] {
		args.GCPClientConstructor = "gcp." + grpcCtor
		args.UseGRPCClient = true
	}
}

func targetCRUDMethodNames(args *ccTemplate.ControllerArgs) []string {
	var methods []string
	if args.HasGet && args.GetMethodName != "" {
		methods = append(methods, args.GetMethodName)
	}
	if args.FindUsesList && args.ListMethodName != "" {
		methods = append(methods, args.ListMethodName)
	}
	if args.HasCreate && args.CreateMethodName != "" {
		methods = append(methods, args.CreateMethodName)
	}
	if args.HasUpdate && args.UpdateMethodName != "" {
		methods = append(methods, args.UpdateMethodName)
	}
	if args.HasDelete && args.DeleteMethodName != "" {
		methods = append(methods, args.DeleteMethodName)
	}
	return methods
}

func scanGAPICPackageFuncsAndMethods(pkgDir string) (funcs map[string]bool, methodRecv map[string]string) {
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, nil
	}
	funcs = make(map[string]bool)
	methodRecv = make(map[string]string)
	fset := token.NewFileSet()

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(pkgDir, entry.Name()), nil, 0)
		if err != nil {
			continue
		}
		for _, decl := range f.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name == nil {
				continue
			}
			if fn.Recv == nil || len(fn.Recv.List) == 0 {
				funcs[fn.Name.Name] = true
				continue
			}
			recv := receiverTypeName(fn.Recv.List[0].Type)
			if ast.IsExported(recv) && strings.HasSuffix(recv, "Client") {
				methodRecv[fn.Name.Name] = recv
			}
		}
	}
	return funcs, methodRecv
}

var (
	modVersionsMu     sync.Mutex
	modVersionsByRoot = make(map[string]map[string]string)
)

func findModCachePackageDir(rootPath, importPath string) string {
	modVersions := loadModVersionsForRoot(rootPath)
	if len(modVersions) == 0 {
		return ""
	}

	bestMod := ""
	bestVer := ""
	for mod, ver := range modVersions {
		if importPath == mod || strings.HasPrefix(importPath, mod+"/") {
			if len(mod) > len(bestMod) {
				bestMod = mod
				bestVer = ver
			}
		}
	}
	if bestMod == "" {
		return ""
	}
	modCache := goModCacheDir()
	if modCache == "" {
		return ""
	}
	relSub := strings.TrimPrefix(strings.TrimPrefix(importPath, bestMod), "/")
	dir := filepath.Join(modCache, bestMod+"@"+bestVer, filepath.FromSlash(relSub))
	if st, err := os.Stat(dir); err == nil && st.IsDir() {
		return dir
	}
	return ""
}

func loadModVersionsForRoot(rootPath string) map[string]string {
	cleanRoot := filepath.Clean(rootPath)
	modVersionsMu.Lock()
	defer modVersionsMu.Unlock()

	if cached, ok := modVersionsByRoot[cleanRoot]; ok {
		return cached
	}
	data, err := os.ReadFile(filepath.Join(cleanRoot, "go.mod"))
	if err != nil {
		return nil
	}
	versions := make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if idx := strings.Index(line, "//"); idx >= 0 {
			line = strings.TrimSpace(line[:idx])
		}
		fields := strings.Fields(line)
		switch {
		case len(fields) == 2 && strings.HasPrefix(fields[1], "v"):
			versions[fields[0]] = fields[1]
		case len(fields) == 3 && fields[0] == "require" && strings.HasPrefix(fields[2], "v"):
			versions[fields[1]] = fields[2]
		}
	}
	modVersionsByRoot[cleanRoot] = versions
	return versions
}

func goModCacheDir() string {
	if modCache := os.Getenv("GOMODCACHE"); modCache != "" {
		return modCache
	}
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return filepath.Join(strings.Split(gopath, string(os.PathListSeparator))[0], "pkg", "mod")
	}
	if home, err := os.UserHomeDir(); err == nil {
		return filepath.Join(home, "go", "pkg", "mod")
	}
	return ""
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
