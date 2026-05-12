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

package kms

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/kms/apiv1"
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName = "kms-ekmconnection-controller"
)

func init() {
	registry.RegisterModel(krm.KMSEKMConnectionGVK, NewEKMConnectionModel)
}

func NewEKMConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &ekmConnectionModel{config: *config}, nil
}

var _ directbase.Model = &ekmConnectionModel{}

type ekmConnectionModel struct {
	config config.ControllerConfig
}

func (m *ekmConnectionModel) client(ctx context.Context) (*gcp.EkmClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEkmRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EkmClient: %w", err)
	}
	return gcpClient, err
}

func (m *ekmConnectionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSEKMConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID := id.(*krm.KMSEKMConnectionIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EKMConnectionAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *ekmConnectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EKMConnectionAdapter struct {
	id        *krm.KMSEKMConnectionIdentity
	gcpClient *gcp.EkmClient
	desired   *krm.KMSEKMConnection
	actual    *kmspb.EkmConnection
}

var _ directbase.Adapter = &EKMConnectionAdapter{}

func (a *EKMConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting KMSEKMConnection", "name", a.id)

	req := &kmspb.GetEkmConnectionRequest{Name: a.id.String()}
	ekmconnectionpb, err := a.gcpClient.GetEkmConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting KMSEKMConnection %q: %w", a.id, err)
	}

	a.actual = ekmconnectionpb
	return true, nil
}

func (a *EKMConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating KMSEKMConnection", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := KMSEKMConnectionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &kmspb.CreateEkmConnectionRequest{
		Parent:          parent,
		EkmConnectionId: a.id.EkmConnection,
		EkmConnection:   resource,
	}
	created, err := a.gcpClient.CreateEkmConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating KMSEKMConnection %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created KMSEKMConnection", "name", a.id)

	status := &krm.KMSEKMConnectionStatus{}
	status.ObservedState = KMSEKMConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EKMConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSEKMConnection", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := KMSEKMConnectionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}

	// EkmConnection fields that can be updated:
	if !reflect.DeepEqual(resource.ServiceResolvers, a.actual.ServiceResolvers) {
		updateMask.Paths = append(updateMask.Paths, "service_resolvers")
	}
	if resource.Etag != "" && resource.Etag != a.actual.Etag {
		updateMask.Paths = append(updateMask.Paths, "etag")
	}
	if resource.KeyManagementMode != a.actual.KeyManagementMode {
		updateMask.Paths = append(updateMask.Paths, "key_management_mode")
	}
	if resource.CryptoSpacePath != "" && resource.CryptoSpacePath != a.actual.CryptoSpacePath {
		updateMask.Paths = append(updateMask.Paths, "crypto_space_path")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	req := &kmspb.UpdateEkmConnectionRequest{
		UpdateMask:    updateMask,
		EkmConnection: resource,
	}
	updated, err := a.gcpClient.UpdateEkmConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating KMSEKMConnection %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated KMSEKMConnection", "name", a.id)

	status := &krm.KMSEKMConnectionStatus{}
	status.ObservedState = KMSEKMConnectionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *EKMConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.KMSEKMConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *KMSEKMConnectionSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// TODO: set ProjectRef, Location etc.
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *EKMConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Kcc currently handles deletion manually? Or doesn't support deletion?
	// EkmConnection does not have a delete method in the proto! Wait, let me check the proto.
	// Oh! `google.cloud.kms.v1.EkmService` has `List`, `Get`, `Create`, `Update`, `GetEkmConfig`, `UpdateEkmConfig`, `VerifyConnectivity`.
	// IT HAS NO DELETE METHOD!
	// Is it possible to delete an EkmConnection?
	return false, fmt.Errorf("Delete operation not supported for KMSEKMConnection resource")
}
