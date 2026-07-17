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

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityGatewaySecurityPolicyGVK, NewGatewaySecurityPolicyModel)
}

func NewGatewaySecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gatewaySecurityPolicyModel{config: *config}, nil
}

var _ directbase.Model = &gatewaySecurityPolicyModel{}

type gatewaySecurityPolicyModel struct {
	config config.ControllerConfig
}

func (m *gatewaySecurityPolicyModel) client(ctx context.Context) (pb.NetworkSecurityClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}

	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity service: %w", err)
	}

	return pb.NewNetworkSecurityClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *gatewaySecurityPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityGatewaySecurityPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	if obj.Spec.TLSInspectionPolicyRef != nil {
		if err := obj.Spec.TLSInspectionPolicyRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing TLSInspectionPolicyRef: %w", err)
		}
	}

	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityGatewaySecurityPolicySpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &gatewaySecurityPolicyAdapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.NetworkSecurityGatewaySecurityPolicyIdentity),
		desired:          desired,
	}, nil
}

func (m *gatewaySecurityPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type gatewaySecurityPolicyAdapter struct {
	gcpClient        pb.NetworkSecurityClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.NetworkSecurityGatewaySecurityPolicyIdentity
	desired          *pb.GatewaySecurityPolicy
	actual           *pb.GatewaySecurityPolicy
}

var _ directbase.Adapter = &gatewaySecurityPolicyAdapter{}

func (a *gatewaySecurityPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding NetworkSecurityGatewaySecurityPolicy", "name", a.id)

	req := &pb.GetGatewaySecurityPolicyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetGatewaySecurityPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkSecurityGatewaySecurityPolicy %s: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *gatewaySecurityPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkSecurityGatewaySecurityPolicy", "name", a.id)

	req := &pb.CreateGatewaySecurityPolicyRequest{
		Parent:                  fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		GatewaySecurityPolicyId: a.id.GatewaySecurityPolicy,
		GatewaySecurityPolicy:   a.desired,
	}

	op, err := a.gcpClient.CreateGatewaySecurityPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkSecurityGatewaySecurityPolicy %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("NetworkSecurityGatewaySecurityPolicy %s waiting for creation: %w", a.id, err)
	}

	actual, err := a.gcpClient.GetGatewaySecurityPolicy(ctx, &pb.GetGatewaySecurityPolicyRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting NetworkSecurityGatewaySecurityPolicy after creation: %w", err)
	}

	log.V(2).Info("successfully created NetworkSecurityGatewaySecurityPolicy", "name", a.id)

	return a.updateStatus(ctx, createOp, actual)
}

func (a *gatewaySecurityPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityGatewaySecurityPolicy", "name", a.id)

	diffs, updateMask, err := compareGatewaySecurityPolicy(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.GatewaySecurityPolicy)
		desired.Name = a.id.String()

		req := &pb.UpdateGatewaySecurityPolicyRequest{
			GatewaySecurityPolicy: desired,
			UpdateMask:            updateMask,
		}

		op, err := a.gcpClient.UpdateGatewaySecurityPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkSecurityGatewaySecurityPolicy %s: %w", a.id, err)
		}

		err = a.waitForOperation(ctx, op)
		if err != nil {
			return fmt.Errorf("NetworkSecurityGatewaySecurityPolicy %s waiting for update: %w", a.id, err)
		}

		actual, err := a.gcpClient.GetGatewaySecurityPolicy(ctx, &pb.GetGatewaySecurityPolicyRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting NetworkSecurityGatewaySecurityPolicy after update: %w", err)
		}
		latest = actual
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *gatewaySecurityPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.GatewaySecurityPolicy) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityGatewaySecurityPolicyStatus{}
	status.ObservedState = NetworkSecurityGatewaySecurityPolicyObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *gatewaySecurityPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityGatewaySecurityPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkSecurityGatewaySecurityPolicySpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NetworkSecurityGatewaySecurityPolicyGVK)
	return u, nil
}

func (a *gatewaySecurityPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkSecurityGatewaySecurityPolicy", "name", a.id)

	req := &pb.DeleteGatewaySecurityPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteGatewaySecurityPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkSecurityGatewaySecurityPolicy %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("NetworkSecurityGatewaySecurityPolicy %s waiting for deletion: %w", a.id, err)
	}

	return true, nil
}

func (a *gatewaySecurityPolicyAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %s", op.GetError().GetMessage())
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
			current, err := a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
			if err != nil {
				return fmt.Errorf("getting operation %q: %w", op.Name, err)
			}
			if current.Done {
				if current.GetError() != nil {
					return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
				}
				return nil
			}
		}
	}
}

func compareGatewaySecurityPolicy(ctx context.Context, actual, desired *pb.GatewaySecurityPolicy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkSecurityGatewaySecurityPolicySpec_FromProto, NetworkSecurityGatewaySecurityPolicySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
