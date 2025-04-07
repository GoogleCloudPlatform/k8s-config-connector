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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.EntryType
// crd.type: DataplexEntryType
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexEntryTypeGVK, NewEntryTypeModel)
}

func NewEntryTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &entryTypeModel{config: config}, nil
}

var _ directbase.Model = &entryTypeModel{}

type entryTypeModel struct {
	config *config.ControllerConfig
}

func (m *entryTypeModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexEntryType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEntryTypeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	adapter := &entryTypeAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	catalogClient, err := gcpClient.catalogClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = catalogClient

	return adapter, nil
}

func (m *entryTypeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type entryTypeAdapter struct {
	gcpClient *gcp.CatalogClient
	id        *krm.EntryTypeIdentity
	desired   *krm.DataplexEntryType
	actual    *pb.EntryType
	reader    client.Reader
}

var _ directbase.Adapter = &entryTypeAdapter{}

func (a *entryTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex entrytype", "name", a.id)

	req := &pb.GetEntryTypeRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEntryType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex entrytype %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *entryTypeAdapter) normalizeReferences(ctx context.Context) error {
	if a.desired.Spec.RequiredAspects != nil {
		for i := range a.desired.Spec.RequiredAspects {
			if a.desired.Spec.RequiredAspects[i].AspectTypeRef != nil {
				_, err := a.desired.Spec.RequiredAspects[i].AspectTypeRef.NormalizedExternal(ctx, a.reader, a.desired.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *entryTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex entrytype", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := DataplexEntryTypeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEntryTypeRequest{
		Parent:      a.id.Parent().String(),
		EntryTypeId: a.id.ID(),
		EntryType:   resource,
	}
	op, err := a.gcpClient.CreateEntryType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex entrytype %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex entrytype %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex entrytype in gcp", "name", a.id)

	status := &krm.DataplexEntryTypeStatus{}
	status.ObservedState = DataplexEntryTypeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *entryTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex entrytype", "name", a.id)

	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := DataplexEntryTypeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if desired.Spec.TypeAliases != nil && !reflect.DeepEqual(resource.TypeAliases, a.actual.TypeAliases) {
		updateMask.Paths = append(updateMask.Paths, "type_aliases")
	}
	if desired.Spec.Platform != nil && !reflect.DeepEqual(resource.Platform, a.actual.Platform) {
		updateMask.Paths = append(updateMask.Paths, "platform")
	}
	if desired.Spec.System != nil && !reflect.DeepEqual(resource.System, a.actual.System) {
		updateMask.Paths = append(updateMask.Paths, "system")
	}
	if desired.Spec.RequiredAspects != nil && !reflect.DeepEqual(resource.RequiredAspects, a.actual.RequiredAspects) {
		updateMask.Paths = append(updateMask.Paths, "required_aspects")
	}
	// Authorization is immutable

	var updated *pb.EntryType
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		req := &pb.UpdateEntryTypeRequest{
			UpdateMask: updateMask,
			EntryType:  resource,
		}
		op, err := a.gcpClient.UpdateEntryType(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex entrytype %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex entrytype %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex entrytype", "name", a.id)
	}

	status := &krm.DataplexEntryTypeStatus{}
	status.ObservedState = DataplexEntryTypeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *entryTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexEntryType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexEntryTypeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.DataplexEntryTypeGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *entryTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex entrytype", "name", a.id)

	req := &pb.DeleteEntryTypeRequest{
		Name: a.id.String(),
		Etag: direct.ValueOf(a.desired.Spec.Etag),
	}
	op, err := a.gcpClient.DeleteEntryType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent dataplex entrytype, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex entrytype %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex entrytype", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex entrytype %s: %w", a.id.String(), err)
	}
	return true, nil
}
