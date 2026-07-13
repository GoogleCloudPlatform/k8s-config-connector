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

package networksecurity

import (
	"context"
	"fmt"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityFirewallEndpointGVK, NewFirewallEndpointModel)
}

func NewFirewallEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firewallEndpointModel{config: *config}, nil
}

var _ directbase.Model = &firewallEndpointModel{}

type firewallEndpointModel struct {
	config config.ControllerConfig
}

func (m *firewallEndpointModel) client(ctx context.Context) (pb.FirewallActivationClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))
	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity service: %w", err)
	}
	return pb.NewFirewallActivationClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *firewallEndpointModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityFirewallEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityFirewallEndpointSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.(*krm.NetworkSecurityFirewallEndpointIdentity).String()

	return &firewallEndpointAdapter{
		id:               id.(*krm.NetworkSecurityFirewallEndpointIdentity),
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		desired:          desired,
	}, nil
}

func (m *firewallEndpointModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type firewallEndpointAdapter struct {
	id               *krm.NetworkSecurityFirewallEndpointIdentity
	gcpClient        pb.FirewallActivationClient
	operationsClient longrunningpb.OperationsClient
	desired          *pb.FirewallEndpoint
	actual           *pb.FirewallEndpoint
}

var _ directbase.Adapter = &firewallEndpointAdapter{}

func (a *firewallEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirewallEndpoint", "name", a.id)

	req := &pb.GetFirewallEndpointRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetProjectFirewallEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirewallEndpoint %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *firewallEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirewallEndpoint", "name", a.id)

	req := &pb.CreateFirewallEndpointRequest{
		Parent:             a.id.ParentString(),
		FirewallEndpointId: a.id.FirewallEndpoint,
		FirewallEndpoint:   a.desired,
	}
	op, err := a.gcpClient.CreateProjectFirewallEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirewallEndpoint %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("FirewallEndpoint %s waiting creation: %w", a.id, err)
	}

	latest, err := a.gcpClient.GetProjectFirewallEndpoint(ctx, &pb.GetFirewallEndpointRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting FirewallEndpoint after creation: %w", err)
	}

	log.V(2).Info("successfully created FirewallEndpoint", "name", a.id)

	return a.updateStatus(ctx, createOp, latest)
}

func (a *firewallEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirewallEndpoint", "name", a.id)

	diffs, updateMask, err := compareFirewallEndpoint(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	log.V(2).Info("fields need update", "name", a.id, "updateMask", updateMask)

	req := &pb.UpdateFirewallEndpointRequest{
		UpdateMask:       updateMask,
		FirewallEndpoint: a.desired,
	}
	req.FirewallEndpoint.Name = a.id.String()

	op, err := a.gcpClient.UpdateProjectFirewallEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FirewallEndpoint %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("FirewallEndpoint %s waiting update: %w", a.id, err)
	}

	latest, err := a.gcpClient.GetProjectFirewallEndpoint(ctx, &pb.GetFirewallEndpointRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting FirewallEndpoint after update: %w", err)
	}

	log.V(2).Info("successfully updated FirewallEndpoint", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func compareFirewallEndpoint(ctx context.Context, actual, desired *pb.FirewallEndpoint) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityFirewallEndpointSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecurityFirewallEndpointSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *firewallEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityFirewallEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityFirewallEndpointSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.FirewallEndpoint)
	u.SetGroupVersionKind(krm.NetworkSecurityFirewallEndpointGVK)

	u.Object = uObj
	return u, nil
}

func (a *firewallEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirewallEndpoint", "name", a.id)

	req := &pb.DeleteFirewallEndpointRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteProjectFirewallEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FirewallEndpoint, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirewallEndpoint %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting delete FirewallEndpoint %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted FirewallEndpoint", "name", a.id)
	return true, nil
}

func (a *firewallEndpointAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FirewallEndpoint) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityFirewallEndpointStatus{}
	status.ObservedState = NetworkSecurityFirewallEndpointObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *firewallEndpointAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	for {
		req := &longrunningpb.GetOperationRequest{
			Name: op.GetName(),
		}
		current, err := a.operationsClient.GetOperation(ctx, req)
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.GetName(), err)
		}
		if current.GetDone() {
			if current.GetError() != nil {
				return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
		}
	}
}
