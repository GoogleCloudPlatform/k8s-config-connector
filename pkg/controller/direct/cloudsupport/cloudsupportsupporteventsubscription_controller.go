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

package cloudsupport

import (
	"context"
	"fmt"

	api "cloud.google.com/go/support/apiv2"
	pb "cloud.google.com/go/support/apiv2/supportpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsupport/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSupportSupportEventSubscriptionGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	id      *krm.CloudSupportSupportEventSubscriptionIdentity
	desired *pb.SupportEventSubscription
	actual  *pb.SupportEventSubscription
	gcp     *api.SupportEventSubscriptionClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*api.SupportEventSubscriptionClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewSupportEventSubscriptionRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SupportEventSubscription REST client: %w", err)
	}
	return gcpClient, nil
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.CloudSupportSupportEventSubscription{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := CloudSupportSupportEventSubscriptionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		id:      id.(*krm.CloudSupportSupportEventSubscriptionIdentity),
		desired: desired,
		gcp:     gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CloudSupportSupportEventSubscriptionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		id:  id,
		gcp: gcp,
	}, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding CloudSupportSupportEventSubscription", "name", a.id.String())

	req := &pb.GetSupportEventSubscriptionRequest{
		Name: a.id.String(),
	}
	resp, err := a.gcp.GetSupportEventSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("CloudSupportSupportEventSubscription not found", "name", a.id.String())
			return false, nil
		}
		return false, fmt.Errorf("getting CloudSupportSupportEventSubscription %q: %w", a.id.String(), err)
	}

	a.actual = resp
	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudSupportSupportEventSubscription", "name", a.id.String())

	req := &pb.CreateSupportEventSubscriptionRequest{
		Parent:                   a.id.ParentString(),
		SupportEventSubscription: a.desired,
	}

	resp, err := a.gcp.CreateSupportEventSubscription(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudSupportSupportEventSubscription %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created CloudSupportSupportEventSubscription", "name", a.id.String())
	a.actual = resp

	return a.updateStatus(ctx, createOp, resp)
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudSupportSupportEventSubscription", "name", a.id.String())

	diffs, updateMask, err := compareSupportEventSubscription(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected, skipping update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	if len(updateMask.Paths) == 0 {
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, &structuredreporting.Diff{
		Object: updateOp.GetUnstructured(),
		Fields: diffs.Fields,
	})

	updatePb := proto.Clone(a.desired).(*pb.SupportEventSubscription)
	updatePb.Name = a.id.String()
	req := &pb.UpdateSupportEventSubscriptionRequest{
		SupportEventSubscription: updatePb,
		UpdateMask:               updateMask,
	}

	updated, err := a.gcp.UpdateSupportEventSubscription(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CloudSupportSupportEventSubscription %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareSupportEventSubscription(ctx context.Context, actual, desired *pb.SupportEventSubscription) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CloudSupportSupportEventSubscriptionSpec_FromProto, CloudSupportSupportEventSubscriptionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudSupportSupportEventSubscription", "name", a.id.String())

	req := &pb.DeleteSupportEventSubscriptionRequest{
		Name: a.id.String(),
	}
	_, err := a.gcp.DeleteSupportEventSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CloudSupportSupportEventSubscription, assuming it was already deleted", "name", a.id.String())
			return false, nil
		}
		return false, fmt.Errorf("deleting CloudSupportSupportEventSubscription %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted CloudSupportSupportEventSubscription", "name", a.id.String())
	return true, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SupportEventSubscription) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CloudSupportSupportEventSubscriptionStatus{}
	status.ObservedState = CloudSupportSupportEventSubscriptionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSupportSupportEventSubscription{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudSupportSupportEventSubscriptionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.SupportEventSubscription)

	if a.id.Organization != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.ParentString()}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.SupportEventSubscription)
	u.SetGroupVersionKind(krm.CloudSupportSupportEventSubscriptionGVK)

	return u, nil
}
