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

package clouddms

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudDMSPrivateConnectionGVK, NewCloudDMSPrivateConnectionModel)
}

func NewCloudDMSPrivateConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCloudDMSPrivateConnection{config: *config}, nil
}

var _ directbase.Model = &modelCloudDMSPrivateConnection{}

type modelCloudDMSPrivateConnection struct {
	config config.ControllerConfig
}

func (m *modelCloudDMSPrivateConnection) client(ctx context.Context) (*gcp.DataMigrationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataMigrationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataMigration client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCloudDMSPrivateConnection) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDMSPrivateConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idObj.(*krm.CloudDMSPrivateConnectionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idObj)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CloudDMSPrivateConnectionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelCloudDMSPrivateConnection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type CloudDMSPrivateConnectionAdapter struct {
	id        *krm.CloudDMSPrivateConnectionIdentity
	gcpClient *gcp.DataMigrationClient
	desired   *krm.CloudDMSPrivateConnection
	actual    *pb.PrivateConnection
	reader    client.Reader
}

var _ directbase.Adapter = &CloudDMSPrivateConnectionAdapter{}

func (a *CloudDMSPrivateConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PrivateConnection", "name", a.id)

	req := &pb.GetPrivateConnectionRequest{Name: a.id.String()}
	pc, err := a.gcpClient.GetPrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PrivateConnection %q: %w", a.id, err)
	}

	a.actual = pc
	return true, nil
}

func (a *CloudDMSPrivateConnectionAdapter) normalizeReference(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.VpcPeeringConfig != nil && obj.Spec.VpcPeeringConfig.VpcNameRef != nil {
		if err := obj.Spec.VpcPeeringConfig.VpcNameRef.Normalize(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}

func (a *CloudDMSPrivateConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivateConnection", "name", a.id)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := CloudDMSPrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreatePrivateConnectionRequest{
		Parent:              parent,
		PrivateConnection:   resource,
		PrivateConnectionId: a.id.PrivateConnection,
	}

	op, err := a.gcpClient.CreatePrivateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PrivateConnection %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PrivateConnection %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PrivateConnection", "name", a.id)

	status := &krm.CloudDMSPrivateConnectionStatus{}
	status.ObservedState = CloudDMSPrivateConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CloudDMSPrivateConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivateConnection", "name", a.id)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := CloudDMSPrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	paths := []string{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		report.AddField("display_name", a.actual.DisplayName, resource.DisplayName)
		paths = append(paths, "display_name")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, resource.Labels)
		paths = append(paths, "labels")
	}
	if desired.Spec.VpcPeeringConfig != nil && !reflect.DeepEqual(resource.GetVpcPeeringConfig(), a.actual.GetVpcPeeringConfig()) {
		report.AddField("vpc_peering_config", a.actual.GetVpcPeeringConfig(), resource.GetVpcPeeringConfig())
		paths = append(paths, "vpc_peering_config")
	}

	if len(paths) != 0 {
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating PrivateConnection is not supported, fields: %v", paths)
	}

	status := &krm.CloudDMSPrivateConnectionStatus{}
	status.ObservedState = CloudDMSPrivateConnectionObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *CloudDMSPrivateConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDMSPrivateConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudDMSPrivateConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CloudDMSPrivateConnectionGVK)

	u.Object = uObj
	return u, nil
}

func (a *CloudDMSPrivateConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivateConnection", "name", a.id)

	req := &pb.DeletePrivateConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent PrivateConnection, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PrivateConnection %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PrivateConnection", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateConnection %s: %w", a.id, err)
	}
	return true, nil
}
