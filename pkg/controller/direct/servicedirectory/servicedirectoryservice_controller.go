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

package servicedirectory

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/servicedirectory/apiv1beta1"
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
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
	registry.RegisterModel(krm.ServiceDirectoryServiceGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.RegistrationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegistrationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServiceDirectory registration client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ServiceDirectoryService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ServiceDirectoryServiceIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := ServiceDirectoryServiceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Support labels on Metadata
	desired.Metadata = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ServiceDirectoryServiceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ServiceDirectoryServiceAdapter struct {
	id        *krm.ServiceDirectoryServiceIdentity
	gcpClient *gcp.RegistrationClient
	desired   *pb.Service
	actual    *pb.Service
}

var _ directbase.Adapter = &ServiceDirectoryServiceAdapter{}

func (a *ServiceDirectoryServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting ServiceDirectoryService", "name", fqn)

	req := &pb.GetServiceRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServiceDirectoryService %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *ServiceDirectoryServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating ServiceDirectoryService", "name", fqn)

	req := &pb.CreateServiceRequest{
		Parent:    parent,
		ServiceId: a.id.Service,
		Service:   a.desired,
	}
	created, err := a.gcpClient.CreateService(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceDirectoryService %s: %w", a.id.Service, err)
	}
	log.V(2).Info("successfully created ServiceDirectoryService", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ServiceDirectoryServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating ServiceDirectoryService", "name", fqn)

	diffs, updateMask, err := compareService(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.Service)
		desired.Name = fqn

		req := &pb.UpdateServiceRequest{
			Service:    desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateService(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ServiceDirectoryService %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ServiceDirectoryServiceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Service) error {
	mapCtx := &direct.MapContext{}
	status := ServiceDirectoryServiceStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ServiceDirectoryServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ServiceDirectoryService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ServiceDirectoryServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ServiceDirectoryServiceGVK)
	return u, nil
}

func (a *ServiceDirectoryServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting ServiceDirectoryService", "name", fqn)

	req := &pb.DeleteServiceRequest{Name: fqn}
	err := a.gcpClient.DeleteService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ServiceDirectoryService, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting ServiceDirectoryService %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted ServiceDirectoryService", "name", fqn)
	return true, nil
}

func compareService(ctx context.Context, actual, desired *pb.Service) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ServiceDirectoryServiceSpec_FromProto, ServiceDirectoryServiceSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Metadata = actual.Metadata
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
