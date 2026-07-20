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

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityAuthzPolicyGVK, NewAuthzPolicyModel)
}

func NewAuthzPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &authzPolicyModel{config: *config}, nil
}

var _ directbase.Model = &authzPolicyModel{}

type authzPolicyModel struct {
	config config.ControllerConfig
}

func (m *authzPolicyModel) client(ctx context.Context) (pb.NetworkSecurityClient, longrunningpb.OperationsClient, error) {
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

func (m *authzPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityAuthzPolicy{}
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
	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityAuthzPolicySpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	// Propagating KRM Metadata Labels
	desired.Labels = obj.GetLabels()

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &authzPolicyAdapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.NetworkSecurityAuthzPolicyIdentity),
		desired:          desired,
	}, nil
}

func (m *authzPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type authzPolicyAdapter struct {
	gcpClient        pb.NetworkSecurityClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.NetworkSecurityAuthzPolicyIdentity
	desired          *pb.AuthzPolicy
	actual           *pb.AuthzPolicy
}

var _ directbase.Adapter = &authzPolicyAdapter{}

func (a *authzPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity authz policy", "name", a.id)

	req := &pb.GetAuthzPolicyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAuthzPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity authz policy %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *authzPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity authz policy", "name", a.id)

	parent := a.id.ParentString()
	req := &pb.CreateAuthzPolicyRequest{
		Parent:        parent,
		AuthzPolicyId: a.id.AuthzPolicy,
		AuthzPolicy:   a.desired,
	}
	op, err := a.gcpClient.CreateAuthzPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity authz policy %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("networksecurity authz policy %s waiting for creation: %w", a.id.String(), err)
	}

	actual, err := a.gcpClient.GetAuthzPolicy(ctx, &pb.GetAuthzPolicyRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting networksecurity authz policy after creation: %w", err)
	}

	log.V(2).Info("successfully created networksecurity authz policy", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *authzPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity authz policy", "name", a.id)

	diffs, updateMask, err := compareAuthzPolicy(ctx, a.actual, a.desired)
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

	req := &pb.UpdateAuthzPolicyRequest{
		UpdateMask:  updateMask,
		AuthzPolicy: a.desired,
	}
	req.AuthzPolicy.Name = a.id.String()

	op, err := a.gcpClient.UpdateAuthzPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networksecurity authz policy %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("networksecurity authz policy %s waiting update: %w", a.id, err)
	}

	// Fetch fully-populated resource after update
	latest, err := a.gcpClient.GetAuthzPolicy(ctx, &pb.GetAuthzPolicyRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting networksecurity authz policy after update: %w", err)
	}

	log.V(2).Info("successfully updated networksecurity authz policy", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func compareAuthzPolicy(ctx context.Context, actual, desired *pb.AuthzPolicy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityAuthzPolicySpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecurityAuthzPolicySpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	// policy_profile is immutable in GCP, but check if user attempted to change it
	if actual.PolicyProfile != desired.PolicyProfile {
		return nil, nil, fmt.Errorf("field policyProfile is immutable and cannot be updated")
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *authzPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.AuthzPolicy) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityAuthzPolicyStatus{}
	status.ObservedState = NetworkSecurityAuthzPolicyObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *authzPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *authzPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity authz policy", "name", a.id)

	req := &pb.DeleteAuthzPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAuthzPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity authz policy %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("networksecurity authz policy %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *authzPolicyAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
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
