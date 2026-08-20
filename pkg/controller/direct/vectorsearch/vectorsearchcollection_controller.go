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

package vectorsearch

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/vectorsearch/apiv1"
	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vectorsearch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.VectorSearchCollectionGVK, NewCollectionModel)
}

func NewCollectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCollection{config: *config}, nil
}

var _ directbase.Model = &modelCollection{}

type modelCollection struct {
	config config.ControllerConfig
}

func (m *modelCollection) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Collection client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCollection) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VectorSearchCollection{}
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
	id := identity.(*krm.VectorSearchCollectionIdentity)

	// Get vectorsearch GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := VectorSearchCollectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredpb.Name = id.String()

	return &CollectionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredpb,
	}, nil
}

func (m *modelCollection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CollectionAdapter struct {
	id        *krm.VectorSearchCollectionIdentity
	gcpClient *gcp.Client
	desired   *pb.Collection
	actual    *pb.Collection
}

var _ directbase.Adapter = &CollectionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CollectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Collection", "name", a.id)

	req := &pb.GetCollectionRequest{Name: a.id.String()}
	collectionpb, err := a.gcpClient.GetCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Collection %q: %w", a.id, err)
	}

	a.actual = collectionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CollectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Collection", "name", a.id)

	req := &pb.CreateCollectionRequest{
		Parent:       a.id.ParentString(),
		Collection:   a.desired,
		CollectionId: a.id.Collection,
	}
	op, err := a.gcpClient.CreateCollection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Collection %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Collection %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Collection", "name", a.id)

	// Fetch the fully populated resource from GCP after creation to ensure read-only fields are populated
	getReq := &pb.GetCollectionRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetCollection(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching Collection after creation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func compareCollection(ctx context.Context, actual, desired *pb.Collection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VectorSearchCollectionSpec_FromProto, VectorSearchCollectionSpec_ToProto)
	if err != nil {
		return nil, nil, fmt.Errorf("masking actual Collection: %w", err)
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Collection)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CollectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Collection", "name", a.id)

	diffs, updateMask, err := compareCollection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	for _, path := range updateMask.Paths {
		if path == "vector_schema" {
			return fmt.Errorf("VectorSchema is immutable and cannot be updated")
		}
		if path == "data_schema" {
			return fmt.Errorf("DataSchema is immutable and cannot be updated")
		}
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateCollectionRequest{
		UpdateMask: updateMask,
		Collection: a.desired,
	}
	op, err := a.gcpClient.UpdateCollection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Collection %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Collection %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Collection", "name", a.id)

	// Fetch the latest state of the resource from GCP to get fully populated read-only fields
	getReq := &pb.GetCollectionRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetCollection(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching Collection after update %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CollectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Collection) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VectorSearchCollectionStatus{}
	status.ObservedState = VectorSearchCollectionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CollectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VectorSearchCollection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VectorSearchCollectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Collection)
	u.SetGroupVersionKind(krm.VectorSearchCollectionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CollectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Collection", "name", a.id)

	req := &pb.DeleteCollectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Collection, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Collection %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Collection", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Collection %s: %w", a.id, err)
	}
	return true, nil
}
