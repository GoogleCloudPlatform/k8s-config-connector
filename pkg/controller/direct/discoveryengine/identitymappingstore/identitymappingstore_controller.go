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

package identitymappingstore

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/discoveryengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineIdentityMappingStoreGVK, NewIdentityMappingStoreModel)
}

func NewIdentityMappingStoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIdentityMappingStore{config: *config}, nil
}

var _ directbase.Model = &modelIdentityMappingStore{}

type modelIdentityMappingStore struct {
	config config.ControllerConfig
}

func (m *modelIdentityMappingStore) client(ctx context.Context, projectID string) (*gcp.IdentityMappingStoreClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.clientWithOpts(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine identitymappingstore client: %w", err)
	}

	return gcpClient, err
}

func (m *modelIdentityMappingStore) clientWithOpts(ctx context.Context, opts []option.ClientOption) (*gcp.IdentityMappingStoreClient, error) {
	return gcp.NewIdentityMappingStoreRESTClient(ctx, opts...)
}

func (m *modelIdentityMappingStore) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineIdentityMappingStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.DiscoveryEngineIdentityMappingStoreIdentity)

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &IdentityMappingStoreAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelIdentityMappingStore) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type IdentityMappingStoreAdapter struct {
	id        *krm.DiscoveryEngineIdentityMappingStoreIdentity
	gcpClient *gcp.IdentityMappingStoreClient
	desired   *krm.DiscoveryEngineIdentityMappingStore
	actual    *pb.IdentityMappingStore
}

var _ directbase.Adapter = &IdentityMappingStoreAdapter{}

func (a *IdentityMappingStoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DiscoveryEngineIdentityMappingStore", "name", a.id)

	req := &pb.GetIdentityMappingStoreRequest{Name: a.id.String()}
	identitymappingstorepb, err := a.gcpClient.GetIdentityMappingStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DiscoveryEngineIdentityMappingStore %q: %w", a.id, err)
	}

	a.actual = identitymappingstorepb
	return true, nil
}

func (a *IdentityMappingStoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DiscoveryEngineIdentityMappingStore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := discoveryengine.DiscoveryEngineIdentityMappingStoreSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateIdentityMappingStoreRequest{
		Parent:                 parent,
		IdentityMappingStore:   resource,
		IdentityMappingStoreId: a.id.IdentityMappingStore,
	}
	created, err := a.gcpClient.CreateIdentityMappingStore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DiscoveryEngineIdentityMappingStore %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DiscoveryEngineIdentityMappingStore", "name", a.id)

	status := &krm.DiscoveryEngineIdentityMappingStoreStatus{}
	status.ObservedState = discoveryengine.DiscoveryEngineIdentityMappingStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *IdentityMappingStoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DiscoveryEngineIdentityMappingStore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := discoveryengine.DiscoveryEngineIdentityMappingStoreSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mappersErr(mapCtx)
	}

	desiredPb.Name = a.actual.Name
	desiredPb.KmsKeyName = a.actual.KmsKeyName

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.DiscoveryEngineIdentityMappingStoreStatus{}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	return fmt.Errorf("DiscoveryEngineIdentityMappingStore is immutable and cannot be updated")
}

func mappersErr(mapCtx *direct.MapContext) error {
	return mapCtx.Err()
}

func (a *IdentityMappingStoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DiscoveryEngineIdentityMappingStore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(discoveryengine.DiscoveryEngineIdentityMappingStoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DiscoveryEngineIdentityMappingStoreGVK)

	u.Object = uObj
	return u, nil
}

func (a *IdentityMappingStoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DiscoveryEngineIdentityMappingStore", "name", a.id)

	req := &pb.DeleteIdentityMappingStoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteIdentityMappingStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DiscoveryEngineIdentityMappingStore, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting DiscoveryEngineIdentityMappingStore %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DiscoveryEngineIdentityMappingStore", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete DiscoveryEngineIdentityMappingStore %s: %w", a.id, err)
	}
	return true, nil
}
