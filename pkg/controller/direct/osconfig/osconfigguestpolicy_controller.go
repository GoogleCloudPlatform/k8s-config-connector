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
// proto.service: google.cloud.osconfig.v1beta
// proto.message: google.cloud.osconfig.v1beta.GuestPolicy
// crd.type: OSConfigGuestPolicy
// crd.version: v1beta1

package osconfig

import (
	"context"
	"fmt"

	osconfig "cloud.google.com/go/osconfig/apiv1beta"
	osconfigpb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func init() {
	registry.RegisterModel(krm.OSConfigGuestPolicyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*osconfig.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := osconfig.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building OSConfig client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.OSConfigGuestPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idIdentity.(*krm.OSConfigGuestPolicyIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.OSConfigGuestPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type adapter struct {
	gcpClient *osconfig.Client
	id        *krm.OSConfigGuestPolicyIdentity
	desired   *krm.OSConfigGuestPolicy
	actual    *osconfigpb.GuestPolicy
}

var _ directbase.Adapter = &adapter{}

// Find retrieves the GCP resource.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting OSConfigGuestPolicy", "name", a.id)
	fqn := a.id.String()
	req := &osconfigpb.GetGuestPolicyRequest{Name: fqn}
	actual, err := a.gcpClient.GetGuestPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting OSConfigGuestPolicy %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating OSConfigGuestPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := OSConfigGuestPolicySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.ParentString()
	guestPolicyID := a.id.OSConfigGuestPolicy

	req := &osconfigpb.CreateGuestPolicyRequest{
		Parent:        parent,
		GuestPolicyId: guestPolicyID,
		GuestPolicy:   desired,
	}

	created, err := a.gcpClient.CreateGuestPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating OSConfigGuestPolicy %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created OSConfigGuestPolicy in gcp", "name", a.id)

	if err := unstructured.SetNestedField(u.Object, guestPolicyID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating OSConfigGuestPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := OSConfigGuestPolicySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, updateMask, err := compareGuestPolicy(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	desired.Name = a.id.String()
	// Etag must match if provided on update
	if a.actual.Etag != "" {
		desired.Etag = a.actual.Etag
	}

	req := &osconfigpb.UpdateGuestPolicyRequest{
		GuestPolicy: desired,
		UpdateMask:  updateMask,
	}

	updated, err := a.gcpClient.UpdateGuestPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating OSConfigGuestPolicy %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated OSConfigGuestPolicy", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OSConfigGuestPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OSConfigGuestPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.OSConfigGuestPolicy)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.OSConfigGuestPolicy)
	u.SetGroupVersionKind(krm.OSConfigGuestPolicyGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting OSConfigGuestPolicy", "name", a.id)
	fqn := a.id.String()
	req := &osconfigpb.DeleteGuestPolicyRequest{Name: fqn}
	err := a.gcpClient.DeleteGuestPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent OSConfigGuestPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting OSConfigGuestPolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted OSConfigGuestPolicy", "name", a.id)
	return true, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *osconfigpb.GuestPolicy) error {
	mapCtx := &direct.MapContext{}
	status := OSConfigGuestPolicyStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareGuestPolicy(ctx context.Context, actual, desired *osconfigpb.GuestPolicy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, OSConfigGuestPolicySpec_FromProto, OSConfigGuestPolicySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *osconfigpb.GuestPolicy) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
