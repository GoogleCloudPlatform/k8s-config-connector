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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/kms/apiv1"

	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "kms-keyhandle-controller"
)

func init() {
	registry.RegisterModel(krm.KMSKeyHandleGVK, NewModel)
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

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.KMSKeyHandle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewKMSKeyHandleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.KMSKeyHandleIdentity
	gcpClient *gcp.AutokeyClient
	desired   *krm.KMSKeyHandle
	actual    *kmspb.KeyHandle
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting KeyHandle", "name", a.id.String())

	// Check whether Config Connector knows the resource identity.
	// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
	// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
	if a.id.ID() == "" {
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
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating KeyHandle")
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := KMSKeyHandleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()

	req := &kmspb.CreateKeyHandleRequest{}
	if a.id.ID() != "" {
		// Optional. Id of the [KeyHandle][google.cloud.kms.v1.KeyHandle]. Must be
		// unique to the resource project and location. If not provided by the caller,
		// a new UUID is used.
		resource.Name = a.id.String()
		req.KeyHandleId = a.id.ID()
	}
	req.Parent = parent.String()
	req.KeyHandle = resource

	op, err := a.gcpClient.CreateKeyHandle(ctx, req)
	if err != nil {
		return fmt.Errorf("creating KeyHandle %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("KeyHandle %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created KeyHandle", "name", a.id.String())

	status := &krm.KMSKeyHandleStatus{}
	status.ObservedState = KMSKeyHandleStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := created.Name
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update operation not supported for KeyHandle.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KeyHandle", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := KMSKeyHandleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.KMSKeyHandleStatus{}
		status.ObservedState = KMSKeyHandleStatusObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		externalRef := a.actual.Name
		status.ExternalRef = &externalRef
		return updateOp.UpdateStatus(ctx, status, nil)
	} else {
		return fmt.Errorf("update operation not supported for resource %v %v, field(s) changed: %v",
			a.desired.GroupVersionKind(), k8s.GetNamespacedName(a.desired), paths)
	}
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.KMSKeyHandle{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSKeyHandleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.Parent()
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: parent.ProjectID}
	obj.Spec.Location = direct.LazyPtr(parent.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Delete operation not supported for KeyHandle, so this operation is a no-op.
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for KeyHandle", "name", a.id.String())
	return true, nil
}
