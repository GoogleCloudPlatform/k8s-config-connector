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

package bigtable

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	adminpb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigtableAuthorizedViewGVK, NewAuthorizedViewModel)
}

func NewAuthorizedViewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAuthorizedView{config: *config}, nil
}

var _ directbase.Model = &modelAuthorizedView{}

type modelAuthorizedView struct {
	config config.ControllerConfig
}

func (m *modelAuthorizedView) client(ctx context.Context) (adminpb.BigtableTableAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableAuthorizedView client options: %w", err)
	}
	opts = append(opts, option.WithEndpoint("bigtableadmin.googleapis.com:443"))
	conn, err := gtransport.Dial(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableAuthorizedView connection: %w", err)
	}
	return adminpb.NewBigtableTableAdminClient(conn), nil
}

func (m *modelAuthorizedView) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableAuthorizedView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAuthorizedViewIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating bigtable table admin client: %w", err)
	}
	return &AuthorizedViewAdapter{
		id:        id,
		gcpClient: client,
		desired:   obj,
	}, nil
}

func (m *modelAuthorizedView) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AuthorizedViewAdapter struct {
	id        *krm.AuthorizedViewIdentity
	gcpClient adminpb.BigtableTableAdminClient
	desired   *krm.BigtableAuthorizedView
	actual    *adminpb.AuthorizedView
}

var _ directbase.Adapter = &AuthorizedViewAdapter{}

func (a *AuthorizedViewAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableAuthorizedView", "name", a.id)

	req := &adminpb.GetAuthorizedViewRequest{
		Name: a.id.String(),
	}
	authorizedView, err := a.gcpClient.GetAuthorizedView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableAuthorizedView %q: %w", a.id, err)
	}

	a.actual = authorizedView
	return true, nil
}

func (a *AuthorizedViewAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableAuthorizedView", "name", a.id)
	mapCtx := &direct.MapContext{}

	proto := BigtableAuthorizedViewSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &adminpb.CreateAuthorizedViewRequest{
		Parent:           a.id.Parent().String(),
		AuthorizedViewId: a.id.ID(),
		AuthorizedView:   proto,
	}
	_, err := a.gcpClient.CreateAuthorizedView(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableAuthorizedView", "name", a.id)

	status := &krm.BigtableAuthorizedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	reqGet := &adminpb.GetAuthorizedViewRequest{
		Name: a.id.String(),
	}
	if created, err := a.gcpClient.GetAuthorizedView(ctx, reqGet); err == nil {
		status.ObservedState = BigtableAuthorizedViewObservedState_v1beta1_FromProto(mapCtx, created)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}

	return nil
}

func (a *AuthorizedViewAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableAuthorizedView", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := BigtableAuthorizedViewSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}

	if a.desired.Spec.SubsetView != nil {
		updateMask.Paths = append(updateMask.Paths, "subset_view")
	}
	if a.desired.Spec.DeletionProtection != nil {
		updateMask.Paths = append(updateMask.Paths, "deletion_protection")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	desired.Name = a.id.String()
	req := &adminpb.UpdateAuthorizedViewRequest{
		AuthorizedView: desired,
		UpdateMask:     updateMask,
	}
	_, err := a.gcpClient.UpdateAuthorizedView(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated BigtableAuthorizedView", "name", a.id)

	status := &krm.BigtableAuthorizedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	reqGet := &adminpb.GetAuthorizedViewRequest{
		Name: a.id.String(),
	}
	if updated, err := a.gcpClient.GetAuthorizedView(ctx, reqGet); err == nil {
		status.ObservedState = BigtableAuthorizedViewObservedState_v1beta1_FromProto(mapCtx, updated)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *AuthorizedViewAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	mapCtx := &direct.MapContext{}
	spec := BigtableAuthorizedViewSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Parent references
	spec.InstanceRef = krm.InstanceRef{External: a.id.Parent().Parent.Id}
	spec.TableRef = krm.TableRef{External: a.id.Parent().String()}

	obj := &krm.BigtableAuthorizedView{}
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableAuthorizedViewGVK)
	return u, nil
}

func (a *AuthorizedViewAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableAuthorizedView", "name", a.id)

	req := &adminpb.DeleteAuthorizedViewRequest{
		Name: a.id.String(),
	}
	_, err := a.gcpClient.DeleteAuthorizedView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableAuthorizedView", "name", a.id)
	return true, nil
}
