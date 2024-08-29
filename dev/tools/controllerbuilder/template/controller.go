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

package template

type ControllerArgs struct {
	Service     string
	Version     string
	Kind        string
	KindToLower string
}

const ControllerTemplate = `
package {{.Service}}

import (
	"context"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/{{.Service}}/{{.Version}}"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/{{.Service}}/apiv1"
	// TODO(user): Update the import with the google cloud client api protobuf
	{{.Service}}pb "cloud.google.com/go/{{.Service}}/v1/{{.KindToLower}}pb"
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
	ctrlName = "{{.Service}}-controller"
    // TODO(user): Confirm service domain
	serviceDomain = "//{{.Service}}.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.GroupVersionKind, NewModel)
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
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building {{.Service}} client: %w", err)
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

	// TODO(user): Use the proper function to validate and resolve dependent KCC resources. 
	// i.e. ResolveProject, ResolveNetwork. etc  
	// TODO(kcc): ops.WithProjectRef, ops.WithNetworkRef 
	projectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	var id *{{.Kind}}Identity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != projectID {
			return nil, fmt.Errorf("{{.Kind}} %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, projectID)
		}
		if id.location != location {
			return nil, fmt.Errorf("{{.Kind}} %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, location)
		}
		if id.{{.KindToLower}} != resourceID {
			return nil, fmt.Errorf("{{.Kind}}  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.{{.KindToLower}}, resourceID)
		}
	}

	// TODO(kcc): GetGCPClient as interface method.
	// Get {{.Service}} GCP client
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
	resourceID string
	projectID  string
	gcpClient  *gcp.Client
	desired    *krm.{{.Kind}}
	actual     *{{.Service}}pb.{{.Kind}}
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting {{.Kind}}", "name", a.id.fullyQualifiedName())

	if a.resourceID == "" {
		return false, nil
	}

	// TODO(user): write the gcp "GET" operation.
	req := &{{.Service}}pb.Get{{.Kind}}Request{Name: a.id.fullyQualifiedName()}
	{{.KindToLower}}pb, err := a.gcpClient.Get{{.Kind}}(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.Kind}} %q: %w", a.id.fullyQualifiedName(), err)
	}

	a.actual = {{.KindToLower}}pb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating {{.Kind}}", "name", a.id.fullyQualifiedName())
	mapCtx := &direct.MapContext{}

	projectID := a.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	desired := a.desired.DeepCopy()
	// TODO(user): Please add the Spec_ToProto mappers under the same package
	resource := {{.Kind}}Spec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "CREATE" or "INSERT" request with required fields.
	req := &{{.Service}}pb.Create{{.Kind}}Request{
		Resource:               resource,
		Project:                a.ProjectID,
	}
	op, err := a.gcpClient.Create{{.Kind}}(ctx, req)
	if err != nil {
		return fmt.Errorf("creating {{.Kind}} %s: %w", a.id.fullyQualifiedName(), err)
	}
	// TODO(user): Adjust the response, depending on the LRO or not.
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s waiting creation: %w", a.id.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created {{.Kind}}", "name", a.id.fullyQualifiedName())

	status := &krm.{{.Kind}}Status{}
	// TODO(user): (Optional) Please add the StatusObservedState_FromProto mappers under the same package
	status := {{.Kind}}StatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating {{.Kind}}", "name", a.id.fullyQualifiedName())
	mapCtx := &direct.MapContext{}

	// TODO(user): (Optional) Add GCP mutable fields.
	// TODO(kcc): Autogen "func immutable()" for each field
	// TODO(kcc): autogen updateMastk.path for mutable gcp fields. 
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	desired := a.desired.DeepCopy()
	resource := {{.Kind}}Spec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "UPDATE" or "PATCH" request with required fields.
	req := &{{.Service}}pb.Update{{.Kind}}Request{
		Resource:               resource,
		Project:                a.ProjectID,
	}
	op, err := a.gcpClient.Update{{.Kind}}(ctx, req)
	if err != nil {
		return fmt.Errorf("updating {{.Kind}} %s: %w", a.id.fullyQualifiedName(), err)
	}
	// TODO(user): Adjust the response, depending on the LRO or not.
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s waiting update: %w", a.id.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated {{.Kind}}", "name", a.id.fullyQualifiedName())

	status := &krm.{{.Kind}}Status{}
	// TODO(user): (Optional) Please add the StatusObservedState_FromProto mappers under the same package
	status := {{.Kind}}StatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO(kcc) 
	return nil, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting {{.Kind}}", "name", a.id.fullyQualifiedName())

	if a.resourceID == "" {
		return false, nil
	}
	req := &{{.Service}}pb.Delete{{.Kind}}Request{Name: a.id.fullyQualifiedName()}
	op, err := a.gcpClient.Delete{{.Kind}}(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting {{.Kind}} %s: %w", a.id.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted {{.Kind}}", "name", a.id.fullyQualifiedName())

	// TODO(user): Adjust the response, depending on the LRO or not.
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete {{.Kind}} %s: %w", a.id.fullyQualifiedName(), err)
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
