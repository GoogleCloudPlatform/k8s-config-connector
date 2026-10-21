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
	"bytes"
	"fmt"
	"go/format"
	"text/template"
)

var controllerTemplate = `// Copyright 2026 Google LLC
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

package {{.KCCService}}

import (
	"context"
	"fmt"
	"strings"

{{- if .Identity.HasProjectRef }}
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
{{- end }}
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/{{.KCCService}}/{{.KCCVersion}}"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
{{- if .Transport.HasLabels }}
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
{{- end }}
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "{{.Transport.GCPClientImportPath}}"
	{{.Transport.PBPackageAlias}} "{{.Transport.PBImportPath}}"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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

{{- if .Identity.UseGetIdentity }}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*{{.Identity.TypeName}})
{{- else }}
	id := &{{.Identity.TypeName}}{}
{{- end }}

{{- if .Transport.HasReferences }}
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
{{- end }}

	gcpClient, err := m.client(ctx{{if .Transport.LocationExpr}}, {{.Transport.LocationExpr}}{{end}})
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
{{- if .Identity.HasFromExternal }}
	id := &{{.Identity.TypeName}}{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx{{if .Transport.LocationExpr}}, id.Parent().Location{{end}})
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

func (m *model{{.ProtoResource}}) client(ctx context.Context{{if .Transport.LocationExpr}}, location string{{end}}) ({{.Transport.GCPClientType}}, error) {
	var opts []option.ClientOption
{{- if .Transport.UseGRPCClient }}
	opts, err := m.config.GRPCClientOptions()
{{- else }}
	opts, err := m.config.RESTClientOptions()
{{- end }}
	if err != nil {
		return nil, err
	}

{{- if eq .Transport.RegionalEndpoint "prefix" }}
	if location != "" {
		endpoint := fmt.Sprintf("%s-{{.Transport.ServiceHost}}", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}
{{- else if eq .Transport.RegionalEndpoint "rep" }}
	if location != "" {
		endpoint := fmt.Sprintf("{{.KCCService}}.%s.rep.googleapis.com", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}
{{- end }}

	gcpClient, err := {{.Transport.GCPClientConstructor}}(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.ProtoResource}} client: %w", err)
	}
	return gcpClient, err
}

type {{.ProtoResource}}Adapter struct {
	id        *{{.Identity.TypeName}}
	gcpClient {{.Transport.GCPClientType}}
	desired   *krm.{{.Kind}}
	actual    *{{.Transport.PBPackageAlias}}.{{.ProtoResource}}
}

var _ directbase.Adapter = &{{.ProtoResource}}Adapter{}

func (a *{{.ProtoResource}}Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting {{.ProtoResource}}", "name", a.id)

{{- if eq .Identity.Strategy "server_generated" }}
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}
{{- end }}

	req := &{{.Transport.PBPackageAlias}}.{{.Lifecycle.GetRequestType}}{
		{{.Lifecycle.GetNameField}}: a.id.String(),
	}
	resp, err := a.gcpClient.{{.Lifecycle.GetMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.ProtoResource}} %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

func (a *{{.ProtoResource}}Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
{{- if .Lifecycle.IsSingleton }}
	return a.Update(ctx, directbase.NewUpdateOperation(createOp.GetUnstructured(), createOp.Reader))
{{- else if not .Lifecycle.CanCreate }}
	return fmt.Errorf("creating {{.Kind}} is not supported by GCP API")
{{- else }}
	log := klog.FromContext(ctx)
	log.V(2).Info("creating {{.ProtoResource}}", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := {{.Kind}}Spec{{.Transport.MapperVersionSuffix}}_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

{{- if eq .Identity.Strategy "client_body" }}
	resource.{{.Identity.NameField}} = a.id.String()
{{- end }}

	req := &{{.Transport.PBPackageAlias}}.{{.Lifecycle.CreateRequestType}}{
		{{.Lifecycle.CreateParentField}}: a.id.{{.Identity.ParentStringExpr}},
{{- if .Lifecycle.CreateIDField }}
		{{.Lifecycle.CreateIDField}}: a.id.{{.Identity.IDExpr}},
{{- end }}
		{{.Lifecycle.CreateBodyField}}: resource,
	}

{{- if .Execution.Create.IsLRO }}
	op, err := a.gcpClient.{{.Lifecycle.CreateMethodName}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
	{{- if .NeedsCreatedVar }}
	created, err := op.Wait(ctx)
	{{- else }}
	err = op.Wait(ctx)
	{{- end }}
	if err != nil {
		return fmt.Errorf("waiting create {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- else }}
	{{- if .NeedsCreatedVar }}
	created, err := a.gcpClient.{{.Lifecycle.CreateMethodName}}(ctx, req)
	{{- else }}
	_, err := a.gcpClient.{{.Lifecycle.CreateMethodName}}(ctx, req)
	{{- end }}
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
{{- end }}
	log.V(2).Info("successfully created {{.ProtoResource}}", "name", a.id)

	status := &krm.{{.Kind}}Status{}
{{- if .Transport.HasObservedState }}
	{{- if .NeedsCreatedVar }}
	status.ObservedState = {{.Kind}}ObservedState{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, created)
	{{- else }}
	status.ObservedState = {{.Kind}}ObservedState{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, resource)
	{{- end }}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}

{{- if eq .Identity.Strategy "server_generated" }}
	status.ExternalRef = direct.LazyPtr(created.GetName())
{{- else }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
{{- end }}
	return createOp.UpdateStatus(ctx, status, nil)
{{- end }}
}

func (a *{{.ProtoResource}}Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating {{.ProtoResource}}", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := {{.Kind}}Spec{{.Transport.MapperVersionSuffix}}_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, updateMask, err := compare{{.ProtoResource}}(ctx, a.actual, desired)
	if err != nil {
		return err
	}

{{- if eq .Lifecycle.UpdateStyle "unsupported" }}
	_ = updateMask
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("{{.Kind}} resource is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
	}

	status := &krm.{{.Kind}}Status{}
{{- if .Transport.HasObservedState }}
	mapCtx = &direct.MapContext{}
	status.ObservedState = {{.Kind}}ObservedState{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
{{- end }}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
{{- else }}
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &{{.Transport.PBPackageAlias}}.{{.Lifecycle.UpdateRequestType}}{
{{- if .Lifecycle.UpdateHasTopLevelName }}
			{{.Lifecycle.GetNameField}}: a.id.String(),
{{- end }}
			{{.Lifecycle.UpdateBodyField}}: desired,
{{- if eq .Lifecycle.UpdateStyle "mask" }}
			{{.Lifecycle.UpdateMaskField}}: updateMask,
{{- end }}
		}

{{- if .Execution.Update.IsLRO }}
		op, err := a.gcpClient.{{.Lifecycle.UpdateMethodName}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
		{{- if .NeedsUpdatedVar }}
		updated, err := op.Wait(ctx)
		{{- else }}
		err = op.Wait(ctx)
		{{- end }}
		if err != nil {
			return fmt.Errorf("waiting update {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- else }}
		{{- if .NeedsUpdatedVar }}
		updated, err := a.gcpClient.{{.Lifecycle.UpdateMethodName}}(ctx, req)
		{{- else }}
		_, err := a.gcpClient.{{.Lifecycle.UpdateMethodName}}(ctx, req)
		{{- end }}
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
{{- end }}
		log.V(2).Info("successfully updated {{.ProtoResource}}", "name", a.id)

		status := &krm.{{.Kind}}Status{}
{{- if .Transport.HasObservedState }}
		mapCtx = &direct.MapContext{}
		{{- if .NeedsUpdatedVar }}
		status.ObservedState = {{.Kind}}ObservedState{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, updated)
		{{- else }}
		status.ObservedState = {{.Kind}}ObservedState{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, desired)
		{{- end }}
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
{{- end }}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	return nil
{{- end }}
}

func compare{{.ProtoResource}}(ctx context.Context, actual, desired *{{.Transport.PBPackageAlias}}.{{.ProtoResource}}) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, {{.Kind}}Spec{{.Transport.MapperVersionSuffix}}_FromProto, {{.Kind}}Spec{{.Transport.MapperVersionSuffix}}_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *{{.ProtoResource}}Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.{{.Kind}}{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf({{.Kind}}Spec{{.Transport.MapperVersionSuffix}}_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

{{- if .Identity.HasProjectRef }}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
{{- end }}
{{- if .Identity.HasLocation }}
	obj.Spec.Location = a.id.Location
{{- end }}
{{- if .Identity.HasResourceID }}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.{{.Identity.IDExpr}})
{{- end }}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.{{.Identity.IDExpr}})
	u.SetGroupVersionKind(krm.{{.Kind}}GVK)
{{- if .Transport.HasLabels }}
	export.SetLabels(u, a.actual.Labels)
{{- end }}
	return u, nil
}

func (a *{{.ProtoResource}}Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
{{- if or .Lifecycle.IsSingleton (eq .Lifecycle.DeleteStyle "unsupported") }}
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for {{.ProtoResource}}, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
{{- else }}
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting {{.ProtoResource}}", "name", a.id)

	req := &{{.Transport.PBPackageAlias}}.{{.Lifecycle.DeleteRequestType}}{
		{{.Lifecycle.DeleteNameField}}: a.id.String(),
{{- if .Lifecycle.DeleteForceField }}
		{{.Lifecycle.DeleteForceField}}: true,
{{- end }}
	}

{{- if .Execution.Delete.IsLRO }}
	op, err := a.gcpClient.{{.Lifecycle.DeleteMethodName}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)

	{{- if .Execution.Delete.IsEmpty }}
	err = op.Wait(ctx)
	{{- else }}
	_, err = op.Wait(ctx)
	{{- end }}
	if err != nil {
		if !direct.IsNotFound(err) && !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("waiting delete {{.ProtoResource}} %s: %w", a.id, err)
		}
	}
	return true, nil
{{- else }}
	{{- if .Execution.Delete.IsEmpty }}
	err := a.gcpClient.{{.Lifecycle.DeleteMethodName}}(ctx, req)
	{{- else }}
	_, err := a.gcpClient.{{.Lifecycle.DeleteMethodName}}(ctx, req)
	{{- end }}
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)
	return true, nil
{{- end }}
{{- end }}
}
`

// RenderController executes the controller template with the given profile and formats the output Go code.
func RenderController(profile *ControllerProfile) ([]byte, error) {
	tmpl, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse controller template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, profile); err != nil {
		return nil, fmt.Errorf("execute controller template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Return unformatted buffer with error so debugging syntax issues is possible
		return buf.Bytes(), fmt.Errorf("format generated Go source: %w", err)
	}
	return formatted, nil
}
