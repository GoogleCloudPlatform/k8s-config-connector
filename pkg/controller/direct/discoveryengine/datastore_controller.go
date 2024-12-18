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

package discoveryengine

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineDataStoreGVK, NewDataStoreModel)
}

func NewDataStoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataStoreModel{config: *config}, nil
}

var _ directbase.Model = &dataStoreModel{}

type dataStoreModel struct {
	config config.ControllerConfig
}

func (m *dataStoreModel) client(ctx context.Context, projectID string) (*gcp.DataStoreClient, error) {
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

	gcpClient, err := gcp.NewDataStoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine datastore client: %w", err)
	}

	return gcpClient, err
}

func (m *dataStoreModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DiscoveryEngineDataStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewDiscoveryEngineDataStoreIDFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineDataStoreSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.ProjectID)
	if err != nil {
		return nil, err
	}

	return &dataStoreAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *dataStoreModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		id, err := krm.ParseDiscoveryEngineDataStoreExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineDataStore format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, id.ProjectID)
			if err != nil {
				return nil, err
			}
			return &dataStoreAdapter{
				gcpClient: gcpClient,
				id:        id,
			}, nil
		}
	}
	return nil, nil
}

type dataStoreAdapter struct {
	gcpClient *gcp.DataStoreClient
	id        *krm.DiscoveryEngineDataStoreID
	desired   *pb.DataStore
	actual    *pb.DataStore
}

var _ directbase.Adapter = &dataStoreAdapter{}

func (a *dataStoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine datastore", "name", a.id)

	req := &pb.GetDataStoreRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDataStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine datastore %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataStoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine datastore", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	req := &pb.CreateDataStoreRequest{
		Parent:      a.id.CollectionLink.String(),
		DataStore:   desired,
		DataStoreId: a.id.DataStore,
	}
	op, err := a.gcpClient.CreateDataStore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine datastore %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine datastore %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created discoveryengine datastore in gcp", "name", a.id)

	status := &krm.DiscoveryEngineDataStoreStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine datastore", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	req := &pb.UpdateDataStoreRequest{
		UpdateMask: updateMask,
		DataStore:  desired,
	}
	updated, err := a.gcpClient.UpdateDataStore(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine datastore %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated discoveryengine datastore", "name", a.id)

	status := &krm.DiscoveryEngineDataStoreStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineDataStore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineDataStoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	obj.Spec.Location = a.id.Location
	obj.Spec.Collection = a.id.Collection
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.DataStore)
	u.SetGroupVersionKind(krm.DiscoveryEngineDataStoreGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *dataStoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine datastore", "name", a.id)

	req := &pb.DeleteDataStoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataStore(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting discoveryengine datastore %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine datastore", "name", a.id)

	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine datastore %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
