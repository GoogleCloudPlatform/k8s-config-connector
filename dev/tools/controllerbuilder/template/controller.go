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

const ctrlName = "{{.Service}}-controller"

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
	resourceID := ValueOf(obj.Spec.ResourceID)
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

	// TODO(kcc): GetGCPClient as interface method.
	// Get {{.Service}} GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		resourceID: resourceID,
		projectID:  projectID,
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
	if a.resourceID == "" {
		return false, nil
	}

	// TODO(user): write the gcp "GET" operation.
	req := &{{.Service}}pb.Get{{.Kind}}Request{Name: a.fullyQualifiedName()}
	{{.KindToLower}}pb, err := a.gcpClient.Get{{.Kind}}(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting {{.Kind}} %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = {{.KindToLower}}pb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	projectID := a.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	desired := a.desired.DeepCopy()
	resource := &{{.Service}}pb.{{.Kind}}{
		Name: a.fullyQualifiedName(),
	}
	// TODO(user): Please add the krm to proto mapping file under apis/{{.Service}}/{{.Version}}
	err := krm.Convert_{{.Kind}}_KRM_TO_API_v1(desired, resource)
	if err != nil {
		return fmt.Errorf("converting {{.Kind}} spec to api: %w", err)
	}

	// TODO(user): write the gcp "CREATE" operation.
	req := &{{.Service}}pb.Create{{.Kind}}Request{}
	op, err := a.gcpClient.Create{{.Kind}}(ctx, req)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s creating failed: %w", resource.Name, err)
	}
	// TODO(user): Adjust the response, depending on the LRO or not.
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s waiting creation failed: %w", resource.Name, err)
	}

	status := &krm.{{.Kind}}Status{}
	// TODO(user): Please add the proto to krm mapping file under apis/{{.Service}}/{{.Version}}
	if err := krm.Convert_{{.Kind}}_API_v1_To_KRM_status(created, status); err != nil {
		return fmt.Errorf("update {{.Kind}} status %w", err)
	}
	status.ObservedState.CreateTime = ToOpenAPIDateTime(created.GetCreateTime())
	status.ObservedState.UpdateTime = ToOpenAPIDateTime(created.GetUpdateTime())
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, u *unstructured.Unstructured) error {

	updateMask := &fieldmaskpb.FieldMask{}

	// TODO(user): Add GCP mutable fields.
	// TODO(kcc): Autogen "func immutable()" for each field
	// TODO(kcc): autogen updateMastk.path for mutable gcp fields. 
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	resource := &{{.Service}}pb.{{.Kind}}{
		Name: a.fullyQualifiedName(),
	}
	desired := a.desired.DeepCopy()
	err := krm.Convert_{{.Kind}}_KRM_To_API_v1(desired, resource)
	if err != nil {
		return fmt.Errorf("converting {{.Kind}} spec to api: %w", err)
	}

	// TODO(user): write the gcp "UPDATE" operation.
	req := &{{.Service}}pb.Updat{{.Kind}}Request{}
	op, err := a.gcpClient.Update{{.Kind}}(ctx, req)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s updating failed: %w", resource.Name, err)
	}
	// TODO(user): Adjust the response, depending on the LRO or not.
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("{{.Kind}} %s waiting update failed: %w", resource.Name, err)
	}
	status := &krm.{{.Kind}}Status{}
	if err := krm.Convert_{{.Kind}}_API_v1_To_KRM_status(updated, status); err != nil {
		return fmt.Errorf("update {{.Kind}} status %w", err)
	}
	status.ObservedState.CreateTime = ToOpenAPIDateTime(updated.GetCreateTime())
	status.ObservedState.UpdateTime = ToOpenAPIDateTime(updated.GetUpdateTime())
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO(kcc) 
	return nil, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}
	req := &{{.Service}}pb.Delete{{.Kind}}Request{Name: a.fullyQualifiedName()}
	op, err := a.gcpClient.Delete{{.Kind}}(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting {{.Kind}} %s: %w", a.fullyQualifiedName(), err)
	}
	// TODO(user): Adjust the response, depending on the LRO or not.
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete {{.Kind}} %s: %w", a.fullyQualifiedName(), err)
	}
	return true, nil
}

// TODO(kcc): interface method
func (a *Adapter) fullyQualifiedName() string {
	// TODO(user): Write the GCP URI for your resource
	return fmt.Sprintf("projects/%s/{{.Kind}}s/%s", a.projectID, a.resourceID)
}

// TODO(kcc): ops.WithParent
func (a *Adapter) getParent() string {
	// TODO(user): Write the GCP URI parent for your resource
	return fmt.Sprintf("projects/%s", a.projectID)
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
	}

	u.Object["status"] = status

	return nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		if apiError.HTTPCode() == code {
			return true
		}
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

// LazyPtr returns a pointer to v, unless it is the empty value, in which case it returns nil.
// It is essentially the inverse of ValueOf, though it is lossy
// because we can't tell nil and empty apart without a pointer.
func LazyPtr[T comparable](v T) *T {
	var defaultValue T
	if v == defaultValue {
		return nil
	}
	return &v
}

func ToOpenAPIDateTime(ts *timestamppb.Timestamp) *string {
	formatted := ts.AsTime().Format(time.RFC3339)
	return &formatted
}

`
