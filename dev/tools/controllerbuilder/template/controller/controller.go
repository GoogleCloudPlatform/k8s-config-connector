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

package {{.KCCService}}

import (
	"context"
	"reflect"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/{{.KCCService}}/{{.KCCVersion}}"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/{{.KCCService}}/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	{{.KCCService}}pb "cloud.google.com/go/{{.KCCService}}/{{.ProtoVersion}}/{{.KCCService}}pb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/googleapis/gax-go/v2/apierror"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (m *model{{.ProtoResource}}) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *model{{.ProtoResource}}) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.{{.Kind}}{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.New{{.ProtoResource}}Identity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get {{.KCCService}} GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &{{.ProtoResource}}Adapter{
		id:        id,
		gcpClient:  gcpClient,
		desired:    obj,
	}, nil
}

func (m *model{{.ProtoResource}}) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type {{.ProtoResource}}Adapter struct {
	id         *krm.{{.ProtoResource}}Identity
	gcpClient  *gcp.Client
	desired    *krm.{{.Kind}}
	actual     *{{.KCCService}}pb.{{.ProtoResource}}
}

var _ directbase.Adapter = &{{.ProtoResource}}Adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter ` + "`" + `Update` + "`" + ` call.
// Return false means the object is not found. This triggers Adapter ` + "`" + `Create` + "`" + ` call.
// Return a non-nil error requeues the requests. 
func (a *{{.ProtoResource}}Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting {{.ProtoResource}}", "name", a.id)

	req := &{{.KCCService}}pb.Get{{.ProtoResource}}Request{Name: a.id.String()}
	{{.ProtoResource | ToLower }}pb, err := a.gcpClient.Get{{.ProtoResource}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.ProtoResource}} %q: %w", a.id, err)
	}

	a.actual = {{.ProtoResource | ToLower }}pb
	return true, nil
}

// Create creates the resource in GCP based on ` + "`" + `spec` + "`" + ` and update the Config Connector object ` + "`" + `status` + "`" + ` based on the GCP response.  
func (a *{{.ProtoResource}}Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating {{.ProtoResource}}", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := {{.Kind}}Spec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &{{.KCCService}}pb.Create{{.ProtoResource}}Request{
		Parent: 						  a.id.Parent().String(),
		{{.ProtoResource}}:               resource,
	}
	op, err := a.gcpClient.Create{{.ProtoResource}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.ProtoResource}} %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.ProtoResource}} %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created {{.ProtoResource}}", "name", a.id)

	status := &krm.{{.Kind}}Status{}
	status.ObservedState = {{.Kind}}ObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on ` + "`" + `spec` + "`" + ` and update the Config Connector object ` + "`" + `status` + "`" + ` based on the GCP response.  
func (a *{{.ProtoResource}}Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating {{.ProtoResource}}", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := {{.Kind}}Spec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	// Option 1: This option is good for proto that has ` + "`" + `field_mask` + "`" + ` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	// Option 2: manually add all mutable fields. 
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = paths.Insert("display_name")
		}
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
		req := &{{.KCCService}}pb.Update{{.ProtoResource}}Request{
			Name:       			a.id.String(),
			UpdateMask:             updateMask,
			{{.ProtoResource}}:     desiredPb,
		}
		op, err := a.gcpClient.Update{{.ProtoResource}}(ctx, req)
		if err != nil {
			return fmt.Errorf("updating {{.ProtoResource}} %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("{{.ProtoResource}} %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated {{.ProtoResource}}", "name", a.id)
	}

	status := &krm.{{.Kind}}Status{}
	status.ObservedState = {{.Kind}}ObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource ` + "`" + `spec` + "`" + `. 
func (a *{{.ProtoResource}}Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location =  a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.{{.Kind}}GVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted. 
func (a *{{.ProtoResource}}Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting {{.ProtoResource}}", "name", a.id)

	req := &{{.KCCService}}pb.Delete{{.ProtoResource}}Request{Name: a.id.String()}
	op, err := a.gcpClient.Delete{{.ProtoResource}}(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting {{.ProtoResource}} %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted {{.ProtoResource}}", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete {{.ProtoResource}} %s: %w", a.id, err)
	}
	return true, nil
}
`
