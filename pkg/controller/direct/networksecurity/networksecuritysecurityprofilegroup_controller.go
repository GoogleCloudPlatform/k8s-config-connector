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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.NetworkSecuritySecurityProfileGroupGVK, NewSecurityProfileGroupModel)
}

func NewSecurityProfileGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &securityProfileGroupModel{config: *config}, nil
}

var _ directbase.Model = &securityProfileGroupModel{}

type securityProfileGroupModel struct {
	config config.ControllerConfig
}

func (m *securityProfileGroupModel) client(ctx context.Context) (pb.SecurityProfileGroupServiceClient, longrunningpb.OperationsClient, error) {
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
	return pb.NewSecurityProfileGroupServiceClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *securityProfileGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecuritySecurityProfileGroup{}
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
	desired := NetworkSecuritySecurityProfileGroupSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.(*krm.NetworkSecuritySecurityProfileGroupIdentity).String()

	return &securityProfileGroupAdapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.NetworkSecuritySecurityProfileGroupIdentity),
		desired:          desired,
	}, nil
}

func (m *securityProfileGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type securityProfileGroupAdapter struct {
	gcpClient        pb.SecurityProfileGroupServiceClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.NetworkSecuritySecurityProfileGroupIdentity
	desired          *pb.SecurityProfileGroup
	actual           *pb.SecurityProfileGroup
}

var _ directbase.Adapter = &securityProfileGroupAdapter{}

func (a *securityProfileGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityProfileGroup", "name", a.id)

	req := &pb.GetSecurityProfileGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSecurityProfileGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityProfileGroup %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *securityProfileGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityProfileGroup", "name", a.id)

	req := &pb.CreateSecurityProfileGroupRequest{
		Parent:                 a.id.ParentString(),
		SecurityProfileGroupId: a.id.Securityprofilegroup,
		SecurityProfileGroup:   a.desired,
	}
	op, err := a.gcpClient.CreateSecurityProfileGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecurityProfileGroup %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("SecurityProfileGroup %s waiting creation: %w", a.id, err)
	}

	latest, err := a.gcpClient.GetSecurityProfileGroup(ctx, &pb.GetSecurityProfileGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting SecurityProfileGroup after creation: %w", err)
	}

	log.V(2).Info("successfully created SecurityProfileGroup", "name", a.id)

	return a.updateStatus(ctx, createOp, latest)
}

func (a *securityProfileGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityProfileGroup", "name", a.id)

	diffs, updateMask, err := compareSecurityProfileGroup(ctx, a.actual, a.desired)
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

	req := &pb.UpdateSecurityProfileGroupRequest{
		UpdateMask:           updateMask,
		SecurityProfileGroup: a.desired,
	}
	req.SecurityProfileGroup.Name = a.id.String()

	op, err := a.gcpClient.UpdateSecurityProfileGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating SecurityProfileGroup %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("SecurityProfileGroup %s waiting update: %w", a.id, err)
	}

	latest, err := a.gcpClient.GetSecurityProfileGroup(ctx, &pb.GetSecurityProfileGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting SecurityProfileGroup after update: %w", err)
	}

	log.V(2).Info("successfully updated SecurityProfileGroup", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func compareSecurityProfileGroup(ctx context.Context, actual, desired *pb.SecurityProfileGroup) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecuritySecurityProfileGroupSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecuritySecurityProfileGroupSpec_v1alpha1_ToProto(mapCtx, spec)
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

func (a *securityProfileGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecuritySecurityProfileGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecuritySecurityProfileGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	if a.actual.ThreatPreventionProfile != "" {
		obj.Spec.ThreatPreventionProfileRef = &krm.NetworkSecuritySecurityProfileRef{External: a.actual.ThreatPreventionProfile}
	}
	if a.actual.CustomMirroringProfile != "" {
		obj.Spec.CustomMirroringProfileRef = &krm.NetworkSecuritySecurityProfileRef{External: a.actual.CustomMirroringProfile}
	}
	if a.actual.CustomInterceptProfile != "" {
		obj.Spec.CustomInterceptProfileRef = &krm.NetworkSecuritySecurityProfileRef{External: a.actual.CustomInterceptProfile}
	}
	if a.actual.UrlFilteringProfile != "" {
		obj.Spec.URLFilteringProfileRef = &krm.NetworkSecuritySecurityProfileRef{External: a.actual.UrlFilteringProfile}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Securityprofilegroup)
	u.SetGroupVersionKind(krm.NetworkSecuritySecurityProfileGroupGVK)

	u.Object = uObj
	return u, nil
}

func (a *securityProfileGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityProfileGroup", "name", a.id)

	req := &pb.DeleteSecurityProfileGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteSecurityProfileGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SecurityProfileGroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting SecurityProfileGroup %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting delete SecurityProfileGroup %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted SecurityProfileGroup", "name", a.id)
	return true, nil
}

func (a *securityProfileGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SecurityProfileGroup) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecuritySecurityProfileGroupStatus{}
	status.ObservedState = NetworkSecuritySecurityProfileGroupObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *securityProfileGroupAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
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
