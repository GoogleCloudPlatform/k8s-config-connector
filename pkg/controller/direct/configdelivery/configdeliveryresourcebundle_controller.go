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

package configdelivery

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
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
	registry.RegisterModel(krm.ConfigDeliveryResourceBundleGVK, NewResourceBundleModel)
}

func NewResourceBundleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &resourceBundleModel{config: config}, nil
}

var _ directbase.Model = &resourceBundleModel{}

type resourceBundleModel struct {
	config *config.ControllerConfig
}

func (m *resourceBundleModel) client(ctx context.Context, project string) (*gcp.Client, error) {
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

func (m *resourceBundleModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeliveryResourceBundle{}
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
	id, ok := idVal.(*krm.ConfigDeliveryResourceBundleIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := ConfigDeliveryResourceBundleSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = obj.GetLabels()
	desired.Name = id.String()

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &ConfigDeliveryResourceBundleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *resourceBundleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ConfigDeliveryResourceBundleIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &ConfigDeliveryResourceBundleAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ConfigDeliveryResourceBundleAdapter struct {
	id        *krm.ConfigDeliveryResourceBundleIdentity
	gcpClient *gcp.Client
	desired   *pb.ResourceBundle
	actual    *pb.ResourceBundle
}

var _ directbase.Adapter = &ConfigDeliveryResourceBundleAdapter{}

func (a *ConfigDeliveryResourceBundleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting ConfigDeliveryResourceBundle", "name", fqn)

	req := &pb.GetResourceBundleRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetResourceBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConfigDeliveryResourceBundle %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *ConfigDeliveryResourceBundleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parentPath := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating ConfigDeliveryResourceBundle", "name", fqn)

	req := &pb.CreateResourceBundleRequest{
		Parent:           parentPath,
		ResourceBundleId: a.id.ResourceBundle,
		ResourceBundle:   a.desired,
	}
	op, err := a.gcpClient.CreateResourceBundle(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConfigDeliveryResourceBundle %s: %w", a.id.ResourceBundle, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ConfigDeliveryResourceBundle creation: %w", err)
	}

	log.V(2).Info("successfully created ConfigDeliveryResourceBundle", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ConfigDeliveryResourceBundleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating ConfigDeliveryResourceBundle", "name", fqn)

	diffs, updateMask, err := compareResourceBundle(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.ResourceBundle)
		desired.Name = fqn

		req := &pb.UpdateResourceBundleRequest{
			ResourceBundle: desired,
			UpdateMask:     updateMask,
		}

		op, err := a.gcpClient.UpdateResourceBundle(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ConfigDeliveryResourceBundle %s: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for ConfigDeliveryResourceBundle update: %w", err)
		}

		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ConfigDeliveryResourceBundleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ResourceBundle) error {
	mapCtx := &direct.MapContext{}
	status := krm.ConfigDeliveryResourceBundleStatus{}
	status.ObservedState = ConfigDeliveryResourceBundleObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *ConfigDeliveryResourceBundleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeliveryResourceBundle{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeliveryResourceBundleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ResourceBundle)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ResourceBundle)
	u.SetGroupVersionKind(krm.ConfigDeliveryResourceBundleGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *ConfigDeliveryResourceBundleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting ConfigDeliveryResourceBundle", "name", fqn)

	req := &pb.DeleteResourceBundleRequest{
		Name:  fqn,
		Force: true,
	}
	op, err := a.gcpClient.DeleteResourceBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ConfigDeliveryResourceBundle, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConfigDeliveryResourceBundle %s: %w", fqn, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for ConfigDeliveryResourceBundle deletion: %w", err)
	}

	log.V(2).Info("successfully deleted ConfigDeliveryResourceBundle", "name", fqn)
	return true, nil
}

func compareResourceBundle(ctx context.Context, actual, desired *pb.ResourceBundle) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ConfigDeliveryResourceBundleSpec_FromProto, ConfigDeliveryResourceBundleSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Labels = actual.Labels
	maskedActual.Name = desired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
