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

package dataform

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/direct/common"

	gcp "cloud.google.com/go/dataform/apiv1beta1"
	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "dataform-controller"

func init() {
	registry.RegisterModel(krm.DataformRepositoryGVK, NewModel)
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
		return nil, fmt.Errorf("building dataform client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured, handlers ...common.CommonHandler) (directbase.Adapter, error) {
	obj := &krm.DataformRepository{}
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
	location := obj.Spec.Region
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		resourceID: resourceID,
		projectID:  projectID,
		gcpClient:  gcpClient,
		location:   location,
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
	location   string
	gcpClient  *gcp.Client
	desired    *krm.DataformRepository
	actual     *dataformpb.Repository
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &dataformpb.GetRepositoryRequest{Name: a.fullyQualifiedName()}
	actual, err := a.gcpClient.GetRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataformRepository %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = actual
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
	mapCtx := &direct.MapContext{}
	resource := DataformRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepository spec to api: %w", mapCtx.Err())
	}

	req := &dataformpb.CreateRepositoryRequest{
		Parent:       a.getParent(),
		Repository:   resource,
		RepositoryId: a.resourceID,
	}
	_, err := a.gcpClient.CreateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepository %s creation failed: %w", resource.Name, err)
	}

	status := &krm.DataformRepositoryStatus{}

	// TODO(acpana): add observed state
	// status.ObservedState.CreateTime = ToOpenAPIDateTime(updated.GetCreateTime())
	// status.ObservedState.UpdateTime = ToOpenAPIDateTime(updated.GetUpdateTime())
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, u *unstructured.Unstructured) error {

	updateMask := &fieldmaskpb.FieldMask{}

	if a.desired.Spec.GitRemoteSettings != nil {
		if !reflect.DeepEqual(a.desired.Spec.GitRemoteSettings, a.actual.GitRemoteSettings) {
			updateMask.Paths = append(updateMask.Paths, "git_remote_settings")
		}
	}
	if a.desired.Spec.WorkspaceCompilationOverrides != nil {
		if !reflect.DeepEqual(a.desired.Spec.WorkspaceCompilationOverrides, a.actual.WorkspaceCompilationOverrides) {
			updateMask.Paths = append(updateMask.Paths, "workspace_compilation_overrides")
		}
	}

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	resource := DataformRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepository spec to api: %w", mapCtx.Err())
	}

	resource.Name = a.fullyQualifiedName()
	req := &dataformpb.UpdateRepositoryRequest{UpdateMask: updateMask, Repository: resource}
	_, err := a.gcpClient.UpdateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepository %s update failed: %w", resource.Name, err)
	}

	status := &krm.DataformRepositoryStatus{}

	// TODO(acpana): add observed state
	// status.ObservedState.CreateTime = ToOpenAPIDateTime(updated.GetCreateTime())
	// status.ObservedState.UpdateTime = ToOpenAPIDateTime(updated.GetUpdateTime())
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

	req := &dataformpb.DeleteRepositoryRequest{Name: a.fullyQualifiedName()}
	if err := a.gcpClient.DeleteRepository(ctx, req); err != nil {
		return false, fmt.Errorf("deleting DataformRepository %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

func (a *Adapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", a.projectID, a.location, a.resourceID)
}

func (a *Adapter) getParent() string {
	return fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)
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
