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

package netapp

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/netapp/apiv1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetAppActiveDirectoryGVK, NewActiveDirectoryModel)
}

func NewActiveDirectoryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelActiveDirectory{config: *config}, nil
}

var _ directbase.Model = &modelActiveDirectory{}

type modelActiveDirectory struct {
	config config.ControllerConfig
}

func (m *modelActiveDirectory) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building netapp client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelActiveDirectory) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetAppActiveDirectory{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.NetAppActiveDirectoryIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	if obj.Spec.Password != nil {
		if err := obj.Spec.Password.NormalizeSecret(ctx, "spec.password", obj.GetNamespace(), reader); err != nil {
			return nil, fmt.Errorf("normalizing spec.password: %w", err)
		}
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &activeDirectoryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelActiveDirectory) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.NetAppActiveDirectoryIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &activeDirectoryAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type activeDirectoryAdapter struct {
	id        *krm.NetAppActiveDirectoryIdentity
	gcpClient *gcp.Client
	desired   *krm.NetAppActiveDirectory
	actual    *pb.ActiveDirectory
}

var _ directbase.Adapter = &activeDirectoryAdapter{}

func (a *activeDirectoryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ActiveDirectory", "name", a.id.String())

	req := &pb.GetActiveDirectoryRequest{Name: a.id.String()}
	activedirectorypb, err := a.gcpClient.GetActiveDirectory(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ActiveDirectory %q: %w", a.id.String(), err)
	}

	a.actual = activedirectorypb
	return true, nil
}

func (a *activeDirectoryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ActiveDirectory", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := NetAppActiveDirectorySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	req := &pb.CreateActiveDirectoryRequest{
		Parent:            a.id.ParentString(),
		ActiveDirectoryId: a.id.ActiveDirectory,
		ActiveDirectory:   desired,
	}
	op, err := a.gcpClient.CreateActiveDirectory(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ActiveDirectory %s: %w", a.id.String(), err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("ActiveDirectory %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created ActiveDirectory", "name", a.id.String())

	latest, err := a.gcpClient.GetActiveDirectory(ctx, &pb.GetActiveDirectoryRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting ActiveDirectory %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *activeDirectoryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ActiveDirectory", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := NetAppActiveDirectorySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	diffs, updateMask, err := a.compareActiveDirectory(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateActiveDirectoryRequest{
		UpdateMask:      updateMask,
		ActiveDirectory: desired,
	}
	req.ActiveDirectory.Name = a.id.String()

	op, err := a.gcpClient.UpdateActiveDirectory(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ActiveDirectory %s: %w", a.id.String(), err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("ActiveDirectory %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated ActiveDirectory", "name", a.id.String())

	latest, err := a.gcpClient.GetActiveDirectory(ctx, &pb.GetActiveDirectoryRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting ActiveDirectory %s after update: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *activeDirectoryAdapter) compareActiveDirectory(ctx context.Context, actual, desired *pb.ActiveDirectory) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetAppActiveDirectorySpec_FromProto, NetAppActiveDirectorySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	// Password is write-only / unreadable from GCP API. Copy desired password to avoid false diff.
	maskedActual.Password = desired.Password

	clonedDesired := proto.Clone(desired).(*pb.ActiveDirectory)

	populateDefaults := func(obj *pb.ActiveDirectory) {
		if obj.OrganizationalUnit == "" {
			obj.OrganizationalUnit = "CN=Computers"
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *activeDirectoryAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ActiveDirectory) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetAppActiveDirectoryStatus{}
	status.ObservedState = NetAppActiveDirectoryObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *activeDirectoryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetAppActiveDirectory{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetAppActiveDirectorySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.ActiveDirectory
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ActiveDirectory)
	u.SetGroupVersionKind(krm.NetAppActiveDirectoryGVK)

	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

func (a *activeDirectoryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ActiveDirectory", "name", a.id.String())

	req := &pb.DeleteActiveDirectoryRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteActiveDirectory(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ActiveDirectory, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ActiveDirectory %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ActiveDirectory %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted ActiveDirectory", "name", a.id.String())
	return true, nil
}
