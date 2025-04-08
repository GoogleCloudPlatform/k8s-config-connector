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

// +tool:controller
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.EntryGroup
// crd.type: DataCatalogEntryGroup
// crd.version: v1alpha1

package datacatalog

import (
	"context"
	"fmt"
	"reflect"

	api "cloud.google.com/go/datacatalog/apiv1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DataCatalogEntryGroupGVK, NewEntryGroupModel)
}

func NewEntryGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &entryGroupModel{config: *config}, nil
}

var _ directbase.Model = &entryGroupModel{}

type entryGroupModel struct {
	config config.ControllerConfig
}

func (m *entryGroupModel) client(ctx context.Context, projectID string) (*api.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog entrygroup client: %w", err)
	}

	return gcpClient, nil
}

func (m *entryGroupModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataCatalogEntryGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEntryGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &entryGroupAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *entryGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type entryGroupAdapter struct {
	gcpClient *api.Client
	id        *krm.EntryGroupIdentity
	desired   *krm.DataCatalogEntryGroup
	actual    *pb.EntryGroup
}

var _ directbase.Adapter = &entryGroupAdapter{}

func (a *entryGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting datacatalog entrygroup", "name", a.id)

	req := &pb.GetEntryGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEntryGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting datacatalog entrygroup %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *entryGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating datacatalog entrygroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogEntryGroupSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEntryGroupRequest{
		Parent:       a.id.Parent().String(),
		EntryGroupId: a.id.ID(),
		EntryGroup:   desired,
	}
	created, err := a.gcpClient.CreateEntryGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating datacatalog entrygroup %s: %w", a.id.String(), err)
	}
	log.Info("successfully created datacatalog entrygroup in gcp", "name", a.id)

	status := &krm.DataCatalogEntryGroupStatus{}
	status.ObservedState = DataCatalogEntryGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *entryGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating datacatalog entrygroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogEntryGroupSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	_, err := common.CompareProtoMessage(desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updateMask := &fieldmaskpb.FieldMask{}
	// Only description and display_name are updatable.
	// transferred_to_dataplex is immutable after being set to true.
	if !reflect.DeepEqual(desired.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	var updated *pb.EntryGroup
	if len(updateMask.Paths) == 0 {
		log.Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		req := &pb.UpdateEntryGroupRequest{
			EntryGroup: desired,
			UpdateMask: updateMask,
		}
		updated, err = a.gcpClient.UpdateEntryGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating datacatalog entrygroup %s: %w", a.id.String(), err)
		}
		log.Info("successfully updated datacatalog entrygroup in gcp", "name", a.id)
	}

	status := &krm.DataCatalogEntryGroupStatus{}
	status.ObservedState = DataCatalogEntryGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *entryGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataCatalogEntryGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataCatalogEntryGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.DataCatalogEntryGroupGVK)

	u.Object = uObj
	return u, nil
}

func (a *entryGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting datacatalog entrygroup", "name", a.id)

	req := &pb.DeleteEntryGroupRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteEntryGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.Info("skipping delete for non-existent datacatalog entrygroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting datacatalog entrygroup %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted datacatalog entrygroup", "name", a.id)

	return true, nil
}
