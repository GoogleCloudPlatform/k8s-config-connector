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

package securesourcemanager

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/securesourcemanager/apiv1"
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecureSourceManagerInstanceGVK, NewSecureSourceManagerInstanceModel)
}

func NewSecureSourceManagerInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &secureSourceManagerInstanceModel{config: *config}, nil
}

var _ directbase.Model = &secureSourceManagerInstanceModel{}

type secureSourceManagerInstanceModel struct {
	config config.ControllerConfig
}

func (m *secureSourceManagerInstanceModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecureSourceManager Instance client: %w", err)
	}
	return gcpClient, err
}

func (m *secureSourceManagerInstanceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecureSourceManagerInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSecureSourceManagerInstanceRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Spec.KmsKeyRef != nil {
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, u.GetNamespace(), obj.Spec.KmsKeyRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.KmsKeyRef = kmsKeyRef
	}

	if obj.Spec.PrivateConfig != nil {
		caPoolRef, err := refs.ResolvePrivateCACAPoolRef(ctx, reader, u, obj.Spec.PrivateConfig.CaPoolRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.PrivateConfig.CaPoolRef = caPoolRef
	}

	mapCtx := &direct.MapContext{}
	desired := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get securesourcemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &secureSourceManagerInstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *secureSourceManagerInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//securesourcemanager.googleapis.com/") {
		return nil, nil
	}

	id, err := krm.ParseSecureSourceManagerInstanceRef(url)
	if err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &secureSourceManagerInstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type secureSourceManagerInstanceAdapter struct {
	id        *krm.SecureSourceManagerInstanceRef
	gcpClient *gcp.Client
	desired   *pb.Instance
	actual    *pb.Instance
}

var _ directbase.Adapter = &secureSourceManagerInstanceAdapter{}

func (a *secureSourceManagerInstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecureSourceManagerInstance", "name", a.id.External)

	req := &pb.GetInstanceRequest{Name: a.id.External}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecureSourceManagerInstance %q: %w", a.id.External, err)
	}

	a.actual = instancepb
	return true, nil
}

func (a *secureSourceManagerInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Instance", "name", a.id.External)

	instance := direct.ProtoClone(a.desired)

	parent, err := a.id.Parent()
	if err != nil {
		return err
	}

	instanceID, err := a.id.ResourceID()
	if err != nil {
		return err
	}

	req := &pb.CreateInstanceRequest{
		Parent:     parent.String(),
		Instance:   instance,
		InstanceId: instanceID,
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance %q: %w", a.id.External, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of instance %q: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created Instance", "name", a.id.External)

	status := &krm.SecureSourceManagerInstanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SecureSourceManagerInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *secureSourceManagerInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("update of SecureSourceManagerInstance not supported")
	return nil
}

func (a *secureSourceManagerInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecureSourceManagerInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecureSourceManagerInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	instanceID, err := a.id.ResourceID()
	if err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: parent.ProjectID}
	obj.Spec.Location = parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(instanceID)
	u.SetGroupVersionKind(krm.SecureSourceManagerInstanceGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *secureSourceManagerInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Instance", "name", a.id.External)

	req := &pb.DeleteInstanceRequest{Name: a.id.External}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Instance %q: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted Instance", "name", a.id.External)

	err = op.Wait(ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting Instance %s: %w", a.id.External, err)
		}
	}
	return true, nil
}
