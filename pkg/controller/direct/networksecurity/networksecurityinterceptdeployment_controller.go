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

	"google.golang.org/api/option"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	networksecurity "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityInterceptDeploymentGVK, NewInterceptDeploymentModel)
}

func NewInterceptDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &interceptDeploymentModel{config: *config}, nil
}

var _ directbase.Model = &interceptDeploymentModel{}

type interceptDeploymentModel struct {
	config config.ControllerConfig
}

func (m *interceptDeploymentModel) client(ctx context.Context) (*networksecurity.InterceptClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := networksecurity.NewInterceptRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity Intercept REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *interceptDeploymentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityInterceptDeployment{}
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
	desired := NetworkSecurityInterceptDeploymentSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &interceptDeploymentAdapter{
		gcpClient: gcpClient,
		id:        id.(*krm.NetworkSecurityInterceptDeploymentIdentity),
		desired:   desired,
	}, nil
}

func (m *interceptDeploymentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type interceptDeploymentAdapter struct {
	gcpClient *networksecurity.InterceptClient
	id        *krm.NetworkSecurityInterceptDeploymentIdentity
	desired   *pb.InterceptDeployment
	actual    *pb.InterceptDeployment
}

var _ directbase.Adapter = &interceptDeploymentAdapter{}

func (a *interceptDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding NetworkSecurityInterceptDeployment", "name", a.id)

	req := &pb.GetInterceptDeploymentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetInterceptDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkSecurityInterceptDeployment %s: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *interceptDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkSecurityInterceptDeployment", "name", a.id)

	req := &pb.CreateInterceptDeploymentRequest{
		Parent:                a.id.ParentString(),
		InterceptDeploymentId: a.id.InterceptDeployment,
		InterceptDeployment:   a.desired,
	}

	op, err := a.gcpClient.CreateInterceptDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkSecurityInterceptDeployment %s: %w", a.id, err)
	}

	actual, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NetworkSecurityInterceptDeployment %s waiting for creation: %w", a.id, err)
	}

	log.V(2).Info("successfully created NetworkSecurityInterceptDeployment", "name", a.id)

	return a.updateStatus(ctx, createOp, actual)
}

func (a *interceptDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityInterceptDeployment", "name", a.id)

	diffs, updateMask, err := compareInterceptDeployment(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.InterceptDeployment)
		desired.Name = a.id.String()

		req := &pb.UpdateInterceptDeploymentRequest{
			InterceptDeployment: desired,
			UpdateMask:          updateMask,
		}

		op, err := a.gcpClient.UpdateInterceptDeployment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkSecurityInterceptDeployment %s: %w", a.id, err)
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("NetworkSecurityInterceptDeployment %s waiting for update: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *interceptDeploymentAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.InterceptDeployment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityInterceptDeploymentStatus{}
	status.ObservedState = NetworkSecurityInterceptDeploymentObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *interceptDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityInterceptDeployment{}
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityInterceptDeploymentSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec != nil {
		obj.Spec = *spec
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NetworkSecurityInterceptDeploymentGVK)
	return u, nil
}

func (a *interceptDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkSecurityInterceptDeployment", "name", a.id)

	req := &pb.DeleteInterceptDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInterceptDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkSecurityInterceptDeployment %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("NetworkSecurityInterceptDeployment %s waiting for deletion: %w", a.id, err)
	}

	return true, nil
}

func compareInterceptDeployment(ctx context.Context, actual, desired *pb.InterceptDeployment) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkSecurityInterceptDeploymentSpec_v1alpha1_FromProto, NetworkSecurityInterceptDeploymentSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
