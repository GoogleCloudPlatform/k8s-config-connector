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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/securesourcemanager/apiv1"

	// TODO(user): Update the import with the google cloud client api protobuf
	securesourcemanagerpb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecureSourceManagerRepositoryGVK, NewSecureSourceManagerRepositoryModel)
}

func NewSecureSourceManagerRepositoryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSecureSourceManagerRepository{config: *config}, nil
}

var _ directbase.Model = &modelSecureSourceManagerRepository{}

type modelSecureSourceManagerRepository struct {
	config config.ControllerConfig
}

func (m *modelSecureSourceManagerRepository) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Repository client: %w", err)
	}
	return gcpClient, err
}

func (m *modelSecureSourceManagerRepository) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecureSourceManagerRepository{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve SecureSourceManagerInstanceRef
	instanceRef := obj.Spec.InstanceRef
	normalizedExternal, err := instanceRef.NormalizedExternal(ctx, reader, instanceRef.Namespace)
	if err != nil {
		return nil, err
	}
	obj.Spec.InstanceRef.External = normalizedExternal

	id, err := krm.NewSecureSourceManagerRepositoryRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get securesourcemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SecureSourceManagerRepositoryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSecureSourceManagerRepository) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type SecureSourceManagerRepositoryAdapter struct {
	id        *krm.SecureSourceManagerRepositoryRef
	gcpClient *gcp.Client
	desired   *krm.SecureSourceManagerRepository
	actual    *securesourcemanagerpb.Repository
}

var _ directbase.Adapter = &SecureSourceManagerRepositoryAdapter{}

func (a *SecureSourceManagerRepositoryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecureSourceManagerRepository", "name", a.id.External)

	req := &securesourcemanagerpb.GetRepositoryRequest{Name: a.id.External}
	repositorypb, err := a.gcpClient.GetRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecureSourceManagerRepository %q: %w", a.id.External, err)
	}

	a.actual = repositorypb
	return true, nil
}

func (a *SecureSourceManagerRepositoryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Repository", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecureSourceManagerRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	repositoryID, err := a.id.ResourceID()
	if err != nil {
		return err
	}

	req := &securesourcemanagerpb.CreateRepositoryRequest{
		Parent:       parent.String(),
		Repository:   resource,
		RepositoryId: repositoryID,
	}
	op, err := a.gcpClient.CreateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Repository %s: %w", a.id.External, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Repository %s waiting creation: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created Repository", "name", a.id.External)

	status := &krm.SecureSourceManagerRepositoryStatus{}
	status.ObservedState = SecureSourceManagerRepositoryObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *SecureSourceManagerRepositoryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("update of SecureSourceManagerRepository not supported")
	return nil
}

func (a *SecureSourceManagerRepositoryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecureSourceManagerRepository{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecureSourceManagerRepositorySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	repositoryID, err := a.id.ResourceID()
	if err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: parent.ProjectID}
	obj.Spec.Location = parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(repositoryID)
	u.SetGroupVersionKind(krm.SecureSourceManagerRepositoryGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *SecureSourceManagerRepositoryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Repository", "name", a.id.External)

	req := &securesourcemanagerpb.DeleteRepositoryRequest{Name: a.id.External}
	_, err := a.gcpClient.DeleteRepository(ctx, req)
	// TODO - remove after the Go protobuf fix is in. https://github.com/golang/protobuf/issues/1620#issuecomment-2402608919
	// Handles the LRO parsing error.
	if err != nil {
		if !strings.Contains(err.Error(), "(line 14:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting Repository %s: %w", a.id.External, err)
		}
	}
	return true, nil
}
