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

// +tool:controller
// proto.service: google.identity.accesscontextmanager.v1.AccessContextManager
// crd.type: AccessContextManagerGCPUserAccessBinding
// crd.version: v1alpha1

package accesscontextmanager

import (
	"context"
	"fmt"

	api "cloud.google.com/go/accesscontextmanager/apiv1"
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AccessContextManagerGCPUserAccessBindingGVK, NewAccessContextManagerGCPUserAccessBindingModel)
}

func NewAccessContextManagerGCPUserAccessBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &accessContextManagerGCPUserAccessBindingModel{config: *config}, nil
}

var _ directbase.Model = &accessContextManagerGCPUserAccessBindingModel{}

type accessContextManagerGCPUserAccessBindingModel struct {
	config config.ControllerConfig
}

func (m *accessContextManagerGCPUserAccessBindingModel) client(ctx context.Context) (*api.Client, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building accesscontextmanager client: %w", err)
	}

	return gcpClient, nil
}

func (m *accessContextManagerGCPUserAccessBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AccessContextManagerGCPUserAccessBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resolvedID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedID.(*krm.AccessContextManagerGCPUserAccessBindingIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &accessContextManagerGCPUserAccessBindingAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *accessContextManagerGCPUserAccessBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AccessContextManagerGCPUserAccessBindingIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &accessContextManagerGCPUserAccessBindingAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type accessContextManagerGCPUserAccessBindingAdapter struct {
	gcpClient *api.Client
	id        *krm.AccessContextManagerGCPUserAccessBindingIdentity
	desired   *krm.AccessContextManagerGCPUserAccessBinding
	actual    *pb.GcpUserAccessBinding
	reader    client.Reader
}

var _ directbase.Adapter = &accessContextManagerGCPUserAccessBindingAdapter{}

func (a *accessContextManagerGCPUserAccessBindingAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard
	if a.id.GcpUserAccessBinding == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("finding AccessContextManagerGCPUserAccessBinding", "name", a.id.String())

	req := &pb.GetGcpUserAccessBindingRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetGcpUserAccessBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AccessContextManagerGCPUserAccessBinding %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *accessContextManagerGCPUserAccessBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AccessContextManagerGCPUserAccessBinding", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredpb := AccessContextManagerGCPUserAccessBindingSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Rule 3: Create() implementation MUST NEVER set the desiredpb.Name in the GCP request payload
	desiredpb.Name = ""

	req := &pb.CreateGcpUserAccessBindingRequest{
		Parent:               a.id.ParentString(),
		GcpUserAccessBinding: desiredpb,
	}

	opLRO, err := a.gcpClient.CreateGcpUserAccessBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AccessContextManagerGCPUserAccessBinding: %w", err)
	}

	created, err := opLRO.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AccessContextManagerGCPUserAccessBinding creation: %w", err)
	}
	log.V(2).Info("successfully created AccessContextManagerGCPUserAccessBinding", "name", created.GetName())

	// Since the ID is server-generated, write it back to spec.resourceID if not set.
	createdID := &krm.AccessContextManagerGCPUserAccessBindingIdentity{}
	if err := createdID.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created resource name %q: %w", created.GetName(), err)
	}
	resourceID := createdID.GcpUserAccessBinding
	if a.desired.Spec.ResourceID == nil || *a.desired.Spec.ResourceID == "" {
		if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
			return err
		}
	}
	a.id.GcpUserAccessBinding = resourceID
	a.actual = created

	return a.updateStatus(ctx, createOp, created)
}

func (a *accessContextManagerGCPUserAccessBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AccessContextManagerGCPUserAccessBinding", "name", a.id.String())

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AccessContextManagerGCPUserAccessBindingSpec_FromProto,
		AccessContextManagerGCPUserAccessBindingSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for AccessContextManagerGCPUserAccessBinding", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if a.actual.GroupKey != "" && a.desired.Spec.GroupKey != a.actual.GroupKey {
		return fmt.Errorf("groupKey is immutable and cannot be updated")
	}

	mapCtx := &direct.MapContext{}
	desiredpb := AccessContextManagerGCPUserAccessBindingSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredpb.Name = a.id.String()

	if updateMask == nil || len(updateMask.Paths) == 0 {
		updateMask = &fieldmaskpb.FieldMask{Paths: []string{"access_levels"}}
	}

	req := &pb.UpdateGcpUserAccessBindingRequest{
		GcpUserAccessBinding: desiredpb,
		UpdateMask:           updateMask,
	}

	opLRO, err := a.gcpClient.UpdateGcpUserAccessBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AccessContextManagerGCPUserAccessBinding %s: %w", a.id.String(), err)
	}

	latest, err := opLRO.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AccessContextManagerGCPUserAccessBinding %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated AccessContextManagerGCPUserAccessBinding", "name", a.id.String())

	a.actual = latest
	return a.updateStatus(ctx, updateOp, latest)
}

func (a *accessContextManagerGCPUserAccessBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AccessContextManagerGCPUserAccessBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AccessContextManagerGCPUserAccessBindingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &krm.OrganizationRef{
		External: "organizations/" + a.id.Organization,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.GcpUserAccessBinding)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.GcpUserAccessBinding)
	u.SetGroupVersionKind(krm.AccessContextManagerGCPUserAccessBindingGVK)
	return u, nil
}

func (a *accessContextManagerGCPUserAccessBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AccessContextManagerGCPUserAccessBinding", "name", a.id.String())

	if a.id.GcpUserAccessBinding == "" {
		return true, nil
	}

	req := &pb.DeleteGcpUserAccessBindingRequest{
		Name: a.id.String(),
	}
	opLRO, err := a.gcpClient.DeleteGcpUserAccessBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AccessContextManagerGCPUserAccessBinding %s: %w", a.id.String(), err)
	}

	if err := opLRO.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for AccessContextManagerGCPUserAccessBinding %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *accessContextManagerGCPUserAccessBindingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.GcpUserAccessBinding) error {
	status := &krm.AccessContextManagerGCPUserAccessBindingStatus{}
	status.Name = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}
