// Copyright 2024 Google LLC
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

package keyhandle

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/kms"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/kms/apiv1"

	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName = "kms-keyhandle-controller"
)

func init() {
	registry.RegisterModel(krm.KMSKeyHandleGVK, NewModel, registry.CannotBeDeleted())
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.AutokeyClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAutokeyRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building KeyHandle client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSKeyHandle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.KMSKeyHandleIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := kms.KMSKeyHandleSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.KMSKeyHandleIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
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
	id        *krm.KMSKeyHandleIdentity
	gcpClient *gcp.AutokeyClient
	desired   *kmspb.KeyHandle
	actual    *kmspb.KeyHandle
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting KeyHandle", "name", a.id.String())

	// Check whether Config Connector knows the resource identity.
	// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
	// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
	if a.id.KeyHandle == "" {
		return false, nil
	}

	req := &kmspb.GetKeyHandleRequest{Name: a.id.String()}
	keyhandlepb, err := a.gcpClient.GetKeyHandle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting KeyHandle %q: %w", a.id.String(), err)
	}

	a.actual = keyhandlepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating KeyHandle")

	req := &kmspb.CreateKeyHandleRequest{}
	if a.id.KeyHandle != "" {
		// Optional. Id of the [KeyHandle][google.cloud.kms.v1.KeyHandle]. Must be
		// unique to the resource project and location. If not provided by the caller,
		// a new UUID is used.
		a.desired.Name = a.id.String()
		req.KeyHandleId = a.id.KeyHandle
	}
	req.Parent = a.id.ParentString()
	req.KeyHandle = a.desired

	op, err := a.gcpClient.CreateKeyHandle(ctx, req)
	if err != nil {
		return fmt.Errorf("creating KeyHandle %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("KeyHandle %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created KeyHandle", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KeyHandle", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, _, err := compareKeyHandle(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("KMSKeyHandle is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.KMSKeyHandle{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(kms.KMSKeyHandleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.KeyHandle)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.KeyHandle)
	u.SetGroupVersionKind(krm.KMSKeyHandleGVK)
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Delete operation not supported for KeyHandle, so this operation is a no-op.
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for KeyHandle", "name", a.id.String())
	return true, nil
}

func compareKeyHandle(ctx context.Context, actual, desired *kmspb.KeyHandle) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, kms.KMSKeyHandleSpec_FromProto, kms.KMSKeyHandleSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *kmspb.KeyHandle) {
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

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *kmspb.KeyHandle) error {
	mapCtx := &direct.MapContext{}
	status := &krm.KMSKeyHandleStatus{}
	status.ObservedState = kms.KMSKeyHandleObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := latest.Name
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}
