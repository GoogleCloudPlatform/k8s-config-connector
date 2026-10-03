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

package financialservices

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/financialservices/apiv1"
	pb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/financialservices/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.FinancialServicesInstanceGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.AMLClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAMLRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FinancialServices AML REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FinancialServicesInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	instanceID := id.(*krm.FinancialServicesInstanceIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := FinancialServicesInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        instanceID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//financialservices.googleapis.com/") {
		return nil, nil
	}

	trimmed := strings.TrimPrefix(url, "//")
	id := &krm.FinancialServicesInstanceIdentity{}
	if err := id.FromExternal(trimmed); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type Adapter struct {
	id        *krm.FinancialServicesInstanceIdentity
	gcpClient *gcp.AMLClient
	desired   *pb.Instance
	actual    *pb.Instance
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FinancialServicesInstance", "name", a.id.String())

	req := &pb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FinancialServicesInstance %q: %w", a.id.String(), err)
	}

	a.actual = instancepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FinancialServicesInstance", "name", a.id.String())

	instance := proto.Clone(a.desired).(*pb.Instance)

	req := &pb.CreateInstanceRequest{
		Parent:     a.id.ParentString(),
		InstanceId: a.id.Instance,
		Instance:   instance,
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created FinancialServicesInstance", "name", a.id.String())

	latest, err := a.gcpClient.GetInstance(ctx, &pb.GetInstanceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting FinancialServicesInstance %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FinancialServicesInstance", "name", a.id.String())

	diffs, updateMask, err := compareFinancialServicesInstance(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs == nil || !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateInstanceRequest{
		Instance:   proto.Clone(a.desired).(*pb.Instance),
		UpdateMask: updateMask,
	}
	req.Instance.Name = a.id.String()

	op, err := a.gcpClient.UpdateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated FinancialServicesInstance", "name", a.id.String())

	latest, err := a.gcpClient.GetInstance(ctx, &pb.GetInstanceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting FinancialServicesInstance %s after update: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareFinancialServicesInstance(ctx context.Context, actual, desired *pb.Instance) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, FinancialServicesInstanceSpec_FromProto, FinancialServicesInstanceSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Instance)

	populateDefaults := func(obj *pb.Instance) {
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := &krm.FinancialServicesInstanceStatus{}
	status.ObservedState = FinancialServicesInstanceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FinancialServicesInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FinancialServicesInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.PtrTo(a.id.Location)
	obj.Spec.ResourceID = direct.PtrTo(a.id.Instance)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.Instance)
	u.SetGroupVersionKind(krm.FinancialServicesInstanceGVK)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FinancialServicesInstance", "name", a.id.String())

	req := &pb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FinancialServicesInstance, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of FinancialServicesInstance %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted FinancialServicesInstance", "name", a.id.String())
	return true, nil
}
