// Copyright 2024 Google LLC
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

type ControllerArgs struct {
	// The ConfigConnector Group without cnrm.google.com
	KCCService string
	// The ConfigConnector Version. Only allow v1alpha1 and v1beta1
	KCCVersion string
	// The ConfigConnector Kind
	Kind string
	// The GCP resource name. Normally the same with the `Kind` without KCCService.
	ProtoResource string
	// The GCP API version.
	ProtoVersion string
}

const ControllerTemplate = `
// Copyright 2024 Google LLC
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
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/{{.KCCService}}/{{.KCCVersion}}"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/{{.KCCService}}/apiv1"

	// TODO(user): Update the import with the google cloud client api protobuf
	{{.KCCService}}pb "cloud.google.com/go/{{.KCCService}}/{{.ProtoVersion}}/{{.KCCService}}pb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/googleapis/gax-go/v2/apierror"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "{{.KCCService}}-controller"
    // TODO(user): Confirm service domain
	serviceDomain = "//{{.KCCService}}.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.{{.Kind}}GVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.ProtoResource}} client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.{{.Kind}}{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get location
	location := obj.Spec.Location

	var id *{{.Kind}}Identity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("{{.Kind}} %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("{{.Kind}} %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.{{.ProtoResource}} != resourceID {
			return nil, fmt.Errorf("{{.Kind}}  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.{{.ProtoResource}}, resourceID)
		}
	}

	// TODO(kcc): GetGCPClient as interface method.
	// Get {{.KCCService}} GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient:  gcpClient,
		desired:    obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id         *{{.Kind}}Identity
	gcpClient  *gcp.Client
	desired    *krm.{{.Kind}}
	actual     *{{.KCCService}}pb.{{.ProtoResource}}
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting {{.Kind}}", "name", a.id.FullyQualifiedName())

	// TODO(user): write the gcp "GET" operation.
	req := &{{.KCCService}}pb.Get{{.ProtoResource}}Request{Name: a.id.FullyQualifiedName()}
	{{.ProtoResource | ToLower }}pb, err := a.gcpClient.Get{{.ProtoResource}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.Kind}} %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = {{.ProtoResource | ToLower }}pb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating {{.ProtoResource}}", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := {{.Kind}}Spec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "CREATE" or "INSERT" request with required fields.
	req := &{{.KCCService}}pb.Create{{.ProtoResource}}Request{
		Parent: 						  a.id.Parent.String(),
		{{.ProtoResource}}:               resource,
	}
	op, err := a.gcpClient.Create{{.ProtoResource}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.ProtoResource}} %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created {{.ProtoResource}}", "name", a.id.FullyQualifiedName())

	status := &krm.{{.Kind}}Status{}
	status.ObservedState = {{.Kind}}StatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating {{.ProtoResource}}", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := {{.Kind}}Spec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.FullyQualifiedName())
		return nil
	}
	// TODO(user): Complete the gcp "UPDATE" or "PATCH" request with required fields.
	req := &{{.KCCService}}pb.Update{{.ProtoResource}}Request{
		Name:       			a.id.FullyQualifiedName(),
		UpdateMask:             updateMask,
		{{.ProtoResource}}:     resource,
	}
	op, err := a.gcpClient.Update{{.ProtoResource}}(ctx, req)
	if err != nil {
		return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id.FullyQualifiedName(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.ProtoResource}} %s waiting update: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated {{.ProtoResource}}", "name", a.id.FullyQualifiedName())

	status := &krm.{{.Kind}}Status{}
	status.ObservedState = {{.Kind}}StatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.{{.Kind}}{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf({{.Kind}}Spec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// TODO(user): Update other resource reference 
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Parent.Project}
	obj.Spec.Location = a.id.Parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting {{.ProtoResource}}", "name", a.id.FullyQualifiedName())

	req := &{{.KCCService}}pb.Delete{{.ProtoResource}}Request{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.Delete{{.ProtoResource}}(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id.FullyQualifiedName())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete {{.ProtoResource}} %s: %w", a.id.FullyQualifiedName(), err)
	}
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
`
