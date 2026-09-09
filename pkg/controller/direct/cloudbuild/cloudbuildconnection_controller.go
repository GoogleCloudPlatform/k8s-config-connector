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

package cloudbuild

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/cloudbuild/apiv2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudBuildConnectionGVK, NewConnectionModel)
}

func NewConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &connectionModel{config: config}, nil
}

var _ directbase.Model = &connectionModel{}

type connectionModel struct {
	config *config.ControllerConfig
}

func (m *connectionModel) client(ctx context.Context) (*gcp.RepositoryManagerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRepositoryManagerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudBuild repository manager REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *connectionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudBuildConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.CloudBuildConnectionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := CloudBuildConnectionSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CloudBuildConnectionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *connectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CloudBuildConnectionIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CloudBuildConnectionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CloudBuildConnectionAdapter struct {
	id        *krm.CloudBuildConnectionIdentity
	gcpClient *gcp.RepositoryManagerClient
	desired   *cloudbuildpb.Connection
	actual    *cloudbuildpb.Connection
}

var _ directbase.Adapter = &CloudBuildConnectionAdapter{}

func (a *CloudBuildConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting CloudBuildConnection", "name", fqn)

	req := &cloudbuildpb.GetConnectionRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CloudBuildConnection %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *CloudBuildConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating CloudBuildConnection", "name", fqn)

	req := &cloudbuildpb.CreateConnectionRequest{
		Parent:       parent,
		ConnectionId: a.id.Connection,
		Connection:   a.desired,
	}
	op, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudBuildConnection %s: %w", a.id.Connection, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of CloudBuildConnection %s: %w", a.id.Connection, err)
	}
	log.V(2).Info("successfully created CloudBuildConnection", "name", created.Name)

	// Fetch fully-populated resource to avoid clearing status fields
	latest, err := a.gcpClient.GetConnection(ctx, &cloudbuildpb.GetConnectionRequest{Name: fqn})
	if err != nil {
		latest = created
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *CloudBuildConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating CloudBuildConnection", "name", fqn)

	diffs, updateMask, err := compareConnection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn
		desired.Etag = a.actual.Etag

		req := &cloudbuildpb.UpdateConnectionRequest{
			Connection: desired,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateConnection(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CloudBuildConnection %s: %w", fqn, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of CloudBuildConnection %s: %w", fqn, err)
		}
		latest = updated
	}

	// Fetch fully-populated resource to avoid clearing status fields
	refetched, err := a.gcpClient.GetConnection(ctx, &cloudbuildpb.GetConnectionRequest{Name: fqn})
	if err == nil {
		latest = refetched
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CloudBuildConnectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *cloudbuildpb.Connection) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CloudBuildConnectionStatus{}
	status.ObservedState = CloudBuildConnectionObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *CloudBuildConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBuildConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBuildConnectionSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Connection)
	obj.Spec.ProjectRef = &refs.ProjectRef{
		External: a.id.Project,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Connection)
	u.SetGroupVersionKind(krm.CloudBuildConnectionGVK)

	return u, nil
}

func (a *CloudBuildConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting CloudBuildConnection", "name", fqn)

	req := &cloudbuildpb.DeleteConnectionRequest{Name: fqn}
	op, err := a.gcpClient.DeleteConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CloudBuildConnection, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting CloudBuildConnection %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CloudBuildConnection, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of CloudBuildConnection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted CloudBuildConnection", "name", fqn)
	return true, nil
}

func compareConnection(ctx context.Context, actual, desired *cloudbuildpb.Connection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CloudBuildConnectionSpec_v1alpha1_FromProto, CloudBuildConnectionSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
