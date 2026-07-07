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

package configdeliveryfleetpackage

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/configdelivery/apiv1"
	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/configdelivery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
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
	registry.RegisterModel(krm.ConfigDeliveryFleetPackageGVK, NewFleetPackageModel)
}

func NewFleetPackageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &fleetPackageModel{config: config}, nil
}

var _ directbase.Model = &fleetPackageModel{}

type fleetPackageModel struct {
	config *config.ControllerConfig
}

func (m *fleetPackageModel) client(ctx context.Context, project string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions(config.WithDefaultQuotaProject(project))
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ConfigDelivery client: %w", err)
	}
	return gcpClient, nil
}

func (m *fleetPackageModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeliveryFleetPackage{}
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
	id, ok := idVal.(*krm.ConfigDeliveryFleetPackageIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := configdelivery.ConfigDeliveryFleetPackageSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = obj.GetLabels()
	desired.Name = id.String()

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &ConfigDeliveryFleetPackageAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *fleetPackageModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ConfigDeliveryFleetPackageIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &ConfigDeliveryFleetPackageAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ConfigDeliveryFleetPackageAdapter struct {
	id        *krm.ConfigDeliveryFleetPackageIdentity
	gcpClient *gcp.Client
	desired   *pb.FleetPackage
	actual    *pb.FleetPackage
}

var _ directbase.Adapter = &ConfigDeliveryFleetPackageAdapter{}

func (a *ConfigDeliveryFleetPackageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting ConfigDeliveryFleetPackage", "name", fqn)

	req := &pb.GetFleetPackageRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetFleetPackage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConfigDeliveryFleetPackage %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *ConfigDeliveryFleetPackageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parentPath := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating ConfigDeliveryFleetPackage", "name", fqn)

	req := &pb.CreateFleetPackageRequest{
		Parent:         parentPath,
		FleetPackageId: a.id.FleetPackage,
		FleetPackage:   a.desired,
	}
	op, err := a.gcpClient.CreateFleetPackage(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConfigDeliveryFleetPackage %s: %w", a.id.FleetPackage, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ConfigDeliveryFleetPackage creation: %w", err)
	}

	log.V(2).Info("successfully created ConfigDeliveryFleetPackage", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ConfigDeliveryFleetPackageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating ConfigDeliveryFleetPackage", "name", fqn)

	diffs, updateMask, err := compareFleetPackage(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.FleetPackage)
		desired.Name = fqn

		req := &pb.UpdateFleetPackageRequest{
			FleetPackage: desired,
			UpdateMask:   updateMask,
		}

		op, err := a.gcpClient.UpdateFleetPackage(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ConfigDeliveryFleetPackage %s: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for ConfigDeliveryFleetPackage update: %w", err)
		}

		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ConfigDeliveryFleetPackageAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FleetPackage) error {
	mapCtx := &direct.MapContext{}
	status := krm.ConfigDeliveryFleetPackageStatus{}
	status.ObservedState = configdelivery.ConfigDeliveryFleetPackageObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *ConfigDeliveryFleetPackageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeliveryFleetPackage{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(configdelivery.ConfigDeliveryFleetPackageSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.FleetPackage)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.FleetPackage)
	u.SetGroupVersionKind(krm.ConfigDeliveryFleetPackageGVK)

	return u, nil
}

func (a *ConfigDeliveryFleetPackageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting ConfigDeliveryFleetPackage", "name", fqn)

	req := &pb.DeleteFleetPackageRequest{Name: fqn}
	op, err := a.gcpClient.DeleteFleetPackage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ConfigDeliveryFleetPackage, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConfigDeliveryFleetPackage %s: %w", fqn, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for ConfigDeliveryFleetPackage deletion: %w", err)
	}

	log.V(2).Info("successfully deleted ConfigDeliveryFleetPackage", "name", fqn)
	return true, nil
}

func compareFleetPackage(ctx context.Context, actual, desired *pb.FleetPackage) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, configdelivery.ConfigDeliveryFleetPackageSpec_FromProto, configdelivery.ConfigDeliveryFleetPackageSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Labels = actual.Labels
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
