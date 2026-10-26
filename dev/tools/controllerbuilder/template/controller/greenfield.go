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

package controller

import (
	"fmt"
	"strings"
)

// NewBaseControllerArgs returns a ControllerArgs initialized with the Pass 1
// standard placeholder defaults (full CRUD, LRO, UpdateMask, standard ID).
func NewBaseControllerArgs(kccService, kccVersion, kind, protoResource, protoVersion string) *ControllerArgs {
	args := &ControllerArgs{
		KCCService:    kccService,
		KCCVersion:    kccVersion,
		Kind:          kind,
		ProtoResource: protoResource,
		ProtoVersion:  protoVersion,
	}
	args.EnsureDefaults()
	return args
}

// EnsureDefaults populates Pass 1 base placeholder defaults on a ControllerArgs
// if it has not yet been initialized by NewBaseControllerArgs or deterministic passes.
func (a *ControllerArgs) EnsureDefaults() {
	if a.PackageName == "" {
		a.PackageName = a.KCCService
	}
	if a.GCPClientImportPath == "" {
		a.GCPClientImportPath = fmt.Sprintf("cloud.google.com/go/%s/apiv1", a.KCCService)
	}
	if a.GCPClientType == "" {
		a.GCPClientType = "*gcp.Client"
	}
	if a.GCPClientConstructor == "" {
		a.GCPClientConstructor = "gcp.NewRESTClient"
	}
	if a.PBImportPath == "" {
		a.PBImportPath = fmt.Sprintf("cloud.google.com/go/%s/%s/%spb", a.KCCService, a.ProtoVersion, a.KCCService)
	}
	if a.PBPackageAlias == "" {
		a.PBPackageAlias = a.KCCService + "pb"
	}
	if !a.PassesApplied && !a.HasGet && !a.FindUsesList {
		a.HasGet = true
	}
	if a.GetMethodName == "" {
		a.GetMethodName = "Get" + a.ProtoResource
	}
	if a.GetRequestType == "" {
		a.GetRequestType = "Get" + a.ProtoResource + "Request"
	}
	if a.GetNameField == "" {
		a.GetNameField = "Name"
	}
	if a.IdentityType == "" {
		a.IdentityType = "krm." + a.ProtoResource + "Identity"
	}
	if a.IdentityConstructorExpr == "" {
		a.IdentityConstructorExpr = fmt.Sprintf("krm.New%sIdentity(ctx, reader, obj)", a.ProtoResource)
	}
	if a.ParentStringExpr == "" {
		a.ParentStringExpr = "a.id.Parent().String()"
	}
	if a.IDExpr == "" {
		a.IDExpr = "a.id.ID()"
	}
	if a.LocationExpr == "" {
		a.LocationExpr = "obj.Spec.Location"
	}
	if !a.PassesApplied && !a.HasCreate && !a.HasUpdate && !a.HasDelete && !a.IsSingleton {
		a.HasCreate = true
		a.CreateMethodName = "Create" + a.ProtoResource
		a.CreateRequestType = "Create" + a.ProtoResource + "Request"
		a.CreateParentField = "Parent"
		a.CreateBodyField = a.ProtoResource
		a.CreateIDField = a.ProtoResource + "Id"
		a.CreateIsLRO = true

		a.HasUpdate = true
		a.UpdateMethodName = "Update" + a.ProtoResource
		a.UpdateRequestType = "Update" + a.ProtoResource + "Request"
		a.UpdateBodyField = a.ProtoResource
		a.UpdateMaskField = "UpdateMask"
		a.UpdateIsLRO = true

		a.HasDelete = true
		a.DeleteMethodName = "Delete" + a.ProtoResource
		a.DeleteRequestType = "Delete" + a.ProtoResource + "Request"
		a.DeleteNameField = "Name"
		a.DeleteIsLRO = true

		a.ResourceNameField = "Name"
		a.HasObservedState = true
	}
}

// ServiceShortHost returns ServiceHost without the ".googleapis.com" suffix,
// used when constructing REP regional endpoints.
func (a *ControllerArgs) ServiceShortHost() string {
	return strings.TrimSuffix(a.ServiceHost, ".googleapis.com")
}

// NeedsCreatedVar reports whether Create() uses the returned resource proto.
func (a *ControllerArgs) NeedsCreatedVar() bool {
	return (a.HasObservedState || (a.IsServerGeneratedID && a.ResourceNameField != "")) &&
		!a.CreateLROIsEmpty && !a.CreateSyncIsEmpty
}

// NeedsSingletonCreatedVar reports whether singleton Create() uses the returned resource proto.
func (a *ControllerArgs) NeedsSingletonCreatedVar() bool {
	return a.HasObservedState && !a.UpdateLROIsEmpty && !a.UpdateSyncIsEmpty
}

// NeedsUpdatedVar reports whether Update() uses the returned resource proto.
func (a *ControllerArgs) NeedsUpdatedVar() bool {
	return a.HasObservedState && !a.UpdateLROIsEmpty && !a.UpdateSyncIsEmpty
}

// ClientOptionsMethod returns the ControllerConfig method used to construct client options.
func (a *ControllerArgs) ClientOptionsMethod() string {
	if a.UseGRPCClient {
		return "GRPCClientOptions"
	}
	return "RESTClientOptions"
}

// UpdateCompareLHS returns the left-hand side variable tuple for CompareGreenfieldSpec in Update().
func (a *ControllerArgs) UpdateCompareLHS() string {
	if !a.HasUpdate {
		return "_, report, updateMask, err"
	}
	resVar := "_"
	if a.UpdateBodyField != "" || a.ResourceNameField != "" {
		resVar = "resource"
	}
	maskVar := "_"
	if a.UpdateMaskField != "" {
		maskVar = "updateMask"
	}
	return fmt.Sprintf("%s, report, %s, err", resVar, maskVar)
}

// ObservedStateSuffix returns the version suffix for <Kind>ObservedState*_FromProto.
func (a *ControllerArgs) ObservedStateSuffix() string {
	if a.HasExplicitObservedStateSuffix {
		return a.ObservedStateVersionSuffix
	}
	return a.MapperVersionSuffix
}

// ServerIDFieldAssignable reports whether a server-generated ID can be assigned
// back onto a.id.<Field> via FromExternal(created.<ResourceNameField>) so that
// a.id.String() preserves canonical projectId while adopting the server-generated ID.
func (a *ControllerArgs) ServerIDFieldAssignable() bool {
	return a.IsServerGeneratedID &&
		a.HasIdentityFromExternal &&
		strings.HasPrefix(a.IDExpr, "a.id.") &&
		!strings.HasSuffix(a.IDExpr, "()")
}

// IdentityStructFieldName returns the struct field name on a.id (e.g. "Connection" from "a.id.Connection").
func (a *ControllerArgs) IdentityStructFieldName() string {
	return strings.TrimPrefix(a.IDExpr, "a.id.")
}

// CanGenerateAdapterForURL reports whether AdapterForURL can be generated using
// id.FromExternal(url) and m.client(ctx, ...).
func (a *ControllerArgs) CanGenerateAdapterForURL() bool {
	if !a.HasIdentityFromExternal {
		return false
	}
	if a.RegionalEndpoint != "" && a.ServiceHost != "" {
		return strings.HasPrefix(a.LocationExpr, "id.")
	}
	return true
}

const GreenfieldControllerTemplate = `// Copyright 2026 Google LLC
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
{{- if .PassesApplied }}

// +generated:controller
// krm.group: {{.KCCService}}.cnrm.cloud.google.com
// krm.version: {{.KCCVersion}}
// krm.kind: {{.Kind}}
{{- end }}

package {{.PackageName}}

import (
	"context"
{{- if .FindUsesList }}
	"errors"
{{- end }}
	"fmt"
{{- if or (and .HasDelete .DeleteIsLRO) (and .CanGenerateAdapterForURL .ServiceHost) }}
	"strings"
{{- end }}

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/{{.KCCService}}/{{.KCCVersion}}"
{{- if .IdentityImportPath }}
	"{{.IdentityImportPath}}"
{{- end }}
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
{{- if .MapperPackageImport }}
	"{{.MapperPackageImport}}"
{{- end }}
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "{{.GCPClientImportPath}}"
	{{.PBPackageAlias}} "{{.PBImportPath}}"
{{- if .FindUsesList }}
	"google.golang.org/api/iterator"
{{- end }}
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.{{.Kind}}GVK, New{{.ProtoResource}}Model)
}

func New{{.ProtoResource}}Model(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{{.ProtoResource}}{config: *config}, nil
}

var _ directbase.Model = &model{{.ProtoResource}}{}

type model{{.ProtoResource}} struct {
	config config.ControllerConfig
}

func (m *model{{.ProtoResource}}) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.{{.Kind}}{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

{{- if .UseGetIdentity }}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*{{.IdentityType}})
{{- else }}
	id, err := {{.IdentityConstructorExpr}}
	if err != nil {
		return nil, err
	}
{{- end }}
{{- if .HasReferences }}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
{{- end }}

	// Get {{.KCCService}} GCP client
{{- if and .RegionalEndpoint .ServiceHost }}
	gcpClient, err := m.client(ctx, {{.LocationExpr}})
{{- else }}
	gcpClient, err := m.client(ctx)
{{- end }}
	if err != nil {
		return nil, err
	}
	return &{{.ProtoResource}}Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model{{.ProtoResource}}) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
{{- if .CanGenerateAdapterForURL }}
{{- if .ServiceHost }}
	if !strings.HasPrefix(url, "//{{.ServiceHost}}/") {
		return nil, nil
	}
{{- end }}
	id := &{{.IdentityType}}{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
{{- if and .RegionalEndpoint .ServiceHost }}
	gcpClient, err := m.client(ctx, {{.LocationExpr}})
{{- else }}
	gcpClient, err := m.client(ctx)
{{- end }}
	if err != nil {
		return nil, err
	}
	return &{{.ProtoResource}}Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
{{- else }}
	// TODO: Support URLs
	return nil, nil
{{- end }}
}

{{- if and (eq .RegionalEndpoint "prefix") .ServiceHost }}
func (m *model{{.ProtoResource}}) client(ctx context.Context, location string) ({{.GCPClientType}}, error) {
	var opts []option.ClientOption
	opts, err := m.config.{{.ClientOptionsMethod}}()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-{{.ServiceHost}}:443", location)))
	}
	gcpClient, err := {{.GCPClientConstructor}}(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.ProtoResource}} client: %w", err)
	}
	return gcpClient, err
}
{{- else if and (eq .RegionalEndpoint "rep") .ServiceHost }}
func (m *model{{.ProtoResource}}) client(ctx context.Context, location string) ({{.GCPClientType}}, error) {
	var opts []option.ClientOption
	opts, err := m.config.{{.ClientOptionsMethod}}()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://{{.ServiceShortHost}}.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := {{.GCPClientConstructor}}(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.ProtoResource}} client: %w", err)
	}
	return gcpClient, err
}
{{- else }}
func (m *model{{.ProtoResource}}) client(ctx context.Context) ({{.GCPClientType}}, error) {
	var opts []option.ClientOption
	opts, err := m.config.{{.ClientOptionsMethod}}()
	if err != nil {
		return nil, err
	}
	gcpClient, err := {{.GCPClientConstructor}}(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.ProtoResource}} client: %w", err)
	}
	return gcpClient, err
}
{{- end }}

type {{.ProtoResource}}Adapter struct {
	id        *{{.IdentityType}}
	gcpClient {{.GCPClientType}}
	desired   *krm.{{.Kind}}
	actual    *{{.PBPackageAlias}}.{{.ProtoResource}}
}

var _ directbase.Adapter = &{{.ProtoResource}}Adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter ` + "`" + `Update` + "`" + ` call.
// Return false means the object is not found. This triggers Adapter ` + "`" + `Create` + "`" + ` call.
// Return a non-nil error requeues the requests.
func (a *{{.ProtoResource}}Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting {{.ProtoResource}}", "name", a.id)

{{- if .IsServerGeneratedID }}
{{- if .HasIdentitySpecifiedMethod }}
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}
{{- else if and .ServerIDUsesEmptyIDGuard .IDExpr }}
	if {{.IDExpr}} == "" {
		return false, nil
	}
{{- end }}
{{- end }}
{{- if .HasGet }}

	req := &{{.PBPackageAlias}}.{{.GetRequestType}}{
		{{.GetNameField}}: a.id.String(),
	}
	{{.ProtoResource | ToLower}}pb, err := a.gcpClient.{{.GetMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.ProtoResource}} %q: %w", a.id, err)
	}

	a.actual = {{.ProtoResource | ToLower}}pb
	return true, nil
{{- else if .FindUsesList }}

	req := &{{.PBPackageAlias}}.{{.ListRequestType}}{
		{{.ListParentField}}: {{.ParentStringExpr}},
	}
	it := a.gcpClient.{{.ListMethodName}}(ctx, req)
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("listing {{.ProtoResource}} in %q: %w", {{.ParentStringExpr}}, err)
		}
{{- if .ResourceNameField }}
		if item.{{.ResourceNameField}} == a.id.String() {
			a.actual = item
			return true, nil
		}
{{- end }}
	}
	return false, nil
{{- else }}
	return false, fmt.Errorf("find not supported for %s", krm.{{.Kind}}GVK)
{{- end }}
}

// Create creates the resource in GCP based on ` + "`" + `spec` + "`" + ` and update the Config Connector object ` + "`" + `status` + "`" + ` based on the GCP response.
func (a *{{.ProtoResource}}Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
{{- if .HasCreate }}
	log := klog.FromContext(ctx)
	log.V(2).Info("creating {{.ProtoResource}}", "name", a.id)
{{- if or .CreateBodyField (and .HasObservedState (not .CreateLROIsEmpty) (not .CreateSyncIsEmpty)) }}
	mapCtx := &direct.MapContext{}
{{- end }}
{{- if .CreateBodyField }}

	desired := a.desired.DeepCopy()
	resource := {{.MapperPrefix}}{{.Kind}}Spec{{.MapperVersionSuffix}}_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- if and .SetResourceNameInCreate .ResourceNameField }}
	resource.{{.ResourceNameField}} = a.id.String()
{{- end }}
{{- end }}

	req := &{{.PBPackageAlias}}.{{.CreateRequestType}}{
{{- if .CreateParentField }}
		{{.CreateParentField}}: {{.ParentStringExpr}},
{{- end }}
{{- if .CreateNameField }}
		{{.CreateNameField}}: a.id.String(),
{{- end }}
{{- if .CreateIDField }}
{{- if .CreateIDFieldIsPointer }}
		{{.CreateIDField}}: direct.LazyPtr({{.IDExpr}}),
{{- else }}
		{{.CreateIDField}}: {{.IDExpr}},
{{- end }}
{{- end }}
{{- if .CreateBodyField }}
		{{.CreateBodyField}}: resource,
{{- end }}
	}
{{- if .CreateIsLRO }}
	op, err := a.gcpClient.{{.CreateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- if .CreateLROIsEmpty }}
	err = op.Wait(ctx)
{{- else if .NeedsCreatedVar }}
	created, err := op.Wait(ctx)
{{- else }}
	_, err = op.Wait(ctx)
{{- end }}
	if err != nil {
		return fmt.Errorf("{{.ProtoResource}} %s waiting creation: %w", a.id, err)
	}
{{- else if .CreateSyncIsEmpty }}
	err := a.gcpClient.{{.CreateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- else if .NeedsCreatedVar }}
	created, err := a.gcpClient.{{.CreateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- else }}
	_, err := a.gcpClient.{{.CreateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- end }}
	log.V(2).Info("successfully created {{.ProtoResource}}", "name", a.id)

	status := &krm.{{.Kind}}Status{}
{{- if and .HasObservedState (not .CreateLROIsEmpty) (not .CreateSyncIsEmpty) }}
	status.ObservedState = {{.MapperPrefix}}{{.Kind}}ObservedState{{.ObservedStateSuffix}}_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}
{{- if and .IsServerGeneratedID .ResourceNameField (not .CreateLROIsEmpty) (not .CreateSyncIsEmpty) }}
{{- if .ServerIDFieldAssignable }}
	parsedID := &{{.IdentityType}}{}
	if err := parsedID.FromExternal(created.{{.ResourceNameField}}); err != nil {
		return err
	}
	{{.IDExpr}} = parsedID.{{.IdentityStructFieldName}}
	status.ExternalRef = direct.LazyPtr(a.id.String())
{{- else }}
	status.ExternalRef = direct.LazyPtr(created.{{.ResourceNameField}})
{{- end }}
{{- else }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
{{- end }}
	return createOp.UpdateStatus(ctx, status, nil)
{{- else if .IsSingleton }}
	log := klog.FromContext(ctx)
	log.V(2).Info("creating (updating singleton) {{.ProtoResource}}", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := {{.MapperPrefix}}{{.Kind}}Spec{{.MapperVersionSuffix}}_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- if .ResourceNameField }}
	resource.{{.ResourceNameField}} = a.id.String()
{{- end }}

	req := &{{.PBPackageAlias}}.{{.UpdateRequestType}}{
{{- if .UpdateBodyField }}
		{{.UpdateBodyField}}: resource,
{{- end }}
	}
{{- if .UpdateIsLRO }}
	op, err := a.gcpClient.{{.UpdateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- if .UpdateLROIsEmpty }}
	err = op.Wait(ctx)
{{- else if .NeedsSingletonCreatedVar }}
	created, err := op.Wait(ctx)
{{- else }}
	_, err = op.Wait(ctx)
{{- end }}
	if err != nil {
		return fmt.Errorf("{{.ProtoResource}} %s waiting creation: %w", a.id, err)
	}
{{- else if .UpdateSyncIsEmpty }}
	err := a.gcpClient.{{.UpdateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- else if .NeedsSingletonCreatedVar }}
	created, err := a.gcpClient.{{.UpdateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- else }}
	_, err := a.gcpClient.{{.UpdateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- end }}
	log.V(2).Info("successfully created {{.ProtoResource}}", "name", a.id)

	status := &krm.{{.Kind}}Status{}
{{- if .NeedsSingletonCreatedVar }}
	status.ObservedState = {{.MapperPrefix}}{{.Kind}}ObservedState{{.ObservedStateSuffix}}_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
{{- else }}
	return fmt.Errorf("create not supported for %s", krm.{{.Kind}}GVK)
{{- end }}
}

// Update updates the resource in GCP based on ` + "`" + `spec` + "`" + ` and update the Config Connector object ` + "`" + `status` + "`" + ` based on the GCP response.
func (a *{{.ProtoResource}}Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating {{.ProtoResource}}", "name", a.id)

	{{.UpdateCompareLHS}} := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		{{.MapperPrefix}}{{.Kind}}Spec{{.MapperVersionSuffix}}_FromProto,
		{{.MapperPrefix}}{{.Kind}}Spec{{.MapperVersionSuffix}}_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

{{- if .HasUpdate }}
{{- if .NeedsUpdatedVar }}
	updated := a.actual
{{- end }}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)

{{- if .ResourceNameField }}
		resource.{{.ResourceNameField}} = a.id.String()
{{- end }}
		req := &{{.PBPackageAlias}}.{{.UpdateRequestType}}{
{{- if .UpdateHasTopLevelName }}
			Name: a.id.String(),
{{- end }}
{{- if .UpdateMaskField }}
			{{.UpdateMaskField}}: updateMask,
{{- end }}
{{- if .UpdateBodyField }}
			{{.UpdateBodyField}}: resource,
{{- end }}
		}
{{- if .UpdateIsLRO }}
		op, err := a.gcpClient.{{.UpdateMethodName}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- if .UpdateLROIsEmpty }}
		err = op.Wait(ctx)
{{- else if .NeedsUpdatedVar }}
		updated, err = op.Wait(ctx)
{{- else }}
		_, err = op.Wait(ctx)
{{- end }}
		if err != nil {
			return fmt.Errorf("{{.ProtoResource}} %s waiting update: %w", a.id, err)
		}
{{- else if .UpdateSyncIsEmpty }}
		err = a.gcpClient.{{.UpdateMethodName}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- else if .NeedsUpdatedVar }}
		updated, err = a.gcpClient.{{.UpdateMethodName}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- else }}
		_, err = a.gcpClient.{{.UpdateMethodName}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- end }}
		log.V(2).Info("successfully updated {{.ProtoResource}}", "name", a.id)
	}

	status := &krm.{{.Kind}}Status{}
{{- if .NeedsUpdatedVar }}
	mapCtx := &direct.MapContext{}
	status.ObservedState = {{.MapperPrefix}}{{.Kind}}ObservedState{{.ObservedStateSuffix}}_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
{{- else }}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating {{.Kind}} is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.{{.Kind}}Status{}
{{- if .HasObservedState }}
	mapCtx := &direct.MapContext{}
	status.ObservedState = {{.MapperPrefix}}{{.Kind}}ObservedState{{.ObservedStateSuffix}}_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
{{- end }}
}

// Export maps the GCP object to a Config Connector resource ` + "`" + `spec` + "`" + `.
func (a *{{.ProtoResource}}Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.{{.Kind}}{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf({{.MapperPrefix}}{{.Kind}}Spec{{.MapperVersionSuffix}}_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.{{.Kind}}GVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *{{.ProtoResource}}Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
{{- if .HasDelete }}
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting {{.ProtoResource}}", "name", a.id)

	req := &{{.PBPackageAlias}}.{{.DeleteRequestType}}{
		{{.DeleteNameField}}: a.id.String(),
	}
{{- if .DeleteIsLRO }}
	op, err := a.gcpClient.{{.DeleteMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent {{.ProtoResource}}, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)

{{- if .DeleteLROIsEmpty }}
	err = op.Wait(ctx)
{{- else }}
	_, err = op.Wait(ctx)
{{- end }}
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete {{.ProtoResource}} %s: %w", a.id, err)
	}
	return true, nil
{{- else if .DeleteSyncIsEmpty }}
	err := a.gcpClient.{{.DeleteMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent {{.ProtoResource}}, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)
	return true, nil
{{- else }}
	_, err := a.gcpClient.{{.DeleteMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent {{.ProtoResource}}, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)
	return true, nil
{{- end }}
{{- else }}
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for {{.ProtoResource}}, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
{{- end }}
}
`
