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
// See the License for the_identity.go specific language governing permissions and
// limitations under the License.

// +tool:controller
// proto.service: google.cloud.compute.v1.BackendServices
// proto.message: google.cloud.compute.v1.SignedUrlKey
// crd.type: ComputeBackendServiceSignedURLKey
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeBackendServiceSignedURLKeyGVK, NewBackendServiceSignedURLKeyModel)
}

func NewBackendServiceSignedURLKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backendServiceSignedURLKeyModel{config: config}, nil
}

var _ directbase.Model = &backendServiceSignedURLKeyModel{}

type backendServiceSignedURLKeyModel struct {
	config *config.ControllerConfig
}

func (m *backendServiceSignedURLKeyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeBackendServiceSignedURLKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	backendServicesClient, err := gcpClient.newBackendServicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &BackendServiceSignedURLKeyAdapter{
		gcpClient: backendServicesClient,
		id:        id.(*krm.ComputeBackendServiceSignedURLKeyIdentity),
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *backendServiceSignedURLKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackendServiceSignedURLKeyAdapter struct {
	gcpClient *compute.BackendServicesClient
	id        *krm.ComputeBackendServiceSignedURLKeyIdentity
	desired   *krm.ComputeBackendServiceSignedURLKey
	reader    client.Reader
}

var _ directbase.Adapter = &BackendServiceSignedURLKeyAdapter{}

// Find retrieves the GCP resource.
func (a *BackendServiceSignedURLKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting parent BackendService for SignedURLKey", "name", a.id)

	backendService, err := a.gcpClient.Get(ctx, &computepb.GetBackendServiceRequest{
		Project:        a.id.Project,
		BackendService: a.id.BackendService,
	})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackendService %q: %w", a.id.BackendService, err)
	}

	if backendService.CdnPolicy != nil {
		for _, keyName := range backendService.CdnPolicy.SignedUrlKeyNames {
			if keyName == a.id.Name {
				return true, nil
			}
		}
	}

	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackendServiceSignedURLKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SignedURLKey", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve/normalize the sensitive KeyValue secret field
	if err := desired.Spec.KeyValue.NormalizeSecret(ctx, "spec.keyValue", desired.Namespace, a.reader); err != nil {
		return err
	}

	resource := ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.KeyName = direct.LazyPtr(a.id.Name)

	req := &computepb.AddSignedUrlKeyBackendServiceRequest{
		Project:              a.id.Project,
		BackendService:       a.id.BackendService,
		SignedUrlKeyResource: resource,
	}
	op, err := a.gcpClient.AddSignedUrlKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeBackendServiceSignedURLKey %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ComputeBackendServiceSignedURLKey %s creation: %w", a.id.String(), err)
	}
	log.Info("successfully created ComputeBackendServiceSignedURLKey in gcp", "name", a.id)

	status := &krm.ComputeBackendServiceSignedURLKeyStatus{}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *BackendServiceSignedURLKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// Signed URL keys are immutable and do not support updates.
	return nil
}

func (a *BackendServiceSignedURLKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("export is not supported for ComputeBackendServiceSignedURLKey")
}

func (a *BackendServiceSignedURLKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SignedURLKey", "name", a.id)

	req := &computepb.DeleteSignedUrlKeyBackendServiceRequest{
		Project:        a.id.Project,
		BackendService: a.id.BackendService,
		KeyName:        a.id.Name,
	}
	op, err := a.gcpClient.DeleteSignedUrlKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeBackendServiceSignedURLKey %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for ComputeBackendServiceSignedURLKey %s deletion: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeBackendServiceSignedURLKey in gcp", "name", a.id)

	return true, nil
}
