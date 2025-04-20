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
// proto.service: google.cloud.metastore.v1.DataprocMetastoreFederation
// proto.message: google.cloud.metastore.v1.Federation
// crd.type: MetastoreFederation
// crd.version: v1alpha1

package metastore

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/metastore/apiv1"
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MetastoreFederationGVK, NewMetastoreFederationModel)
}

func NewMetastoreFederationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &MetastoreFederationModel{config: *config}, nil
}

var _ directbase.Model = &MetastoreFederationModel{}

type MetastoreFederationModel struct {
	config config.ControllerConfig
}

func (m *MetastoreFederationModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MetastoreFederation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFederationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	metastoreClient, err := gcp.NewDataprocMetastoreFederationClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MetastoreFederationAdapter{
		gcpClient: metastoreClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *MetastoreFederationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MetastoreFederationAdapter struct {
	gcpClient *gcp.DataprocMetastoreFederationClient
	id        *krm.FederationIdentity
	desired   *krm.MetastoreFederation
	actual    *pb.Federation
	reader    client.Reader
}

var _ directbase.Adapter = &MetastoreFederationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MetastoreFederationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MetastoreFederation", "name", a.id)

	req := &pb.GetFederationRequest{Name: a.id.String()}
	metastorefederationpb, err := a.gcpClient.GetFederation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MetastoreFederation %q: %w", a.id, err)
	}

	a.actual = metastorefederationpb
	return true, nil
}

func (a *MetastoreFederationAdapter) resolveReferences(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.BackendMetastores != nil {
		// NOTE: The Spec.BackendMetastores field is a map[string]BackendMetastore,
		// but the loop iterates using an integer 'rank'. This seems incorrect.
		// Assuming iteration over the map is intended, the key should be used.
		// However, sticking to the original structure for now, just fixing the field name.
		// Consider refactoring this loop if 'rank' is not the correct way to index.
		for rank, backend := range obj.Spec.BackendMetastores { // Potential issue: iterating map with integer index 'rank'
			if backend.ServiceRef != nil {
				// Resolve the reference using the correct field and original resolver attempt
				_, err := backend.ServiceRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
				if err != nil {
					// Use the correct field name in the error message
					return fmt.Errorf("resolving backendMetastores[%s].serviceRef: %w", rank, err)
				}
				// Do not reassign resolved value back to the .ref field
			}
			// Update to the map is likely unnecessary if 'backend' is not modified.
			obj.Spec.BackendMetastores[rank] = backend
		}
	}

	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetastoreFederationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MetastoreFederation", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MetastoreFederationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateFederationRequest{
		Parent:       a.id.Parent().String(),
		FederationId: a.id.ID(),
		Federation:   resource,
	}
	op, err := a.gcpClient.CreateFederation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MetastoreFederation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetastoreFederation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MetastoreFederation", "name", a.id)

	status := &krm.MetastoreFederationStatus{}
	status.ObservedState = MetastoreFederationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetastoreFederationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MetastoreFederation", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MetastoreFederationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.BackendMetastores != nil && !reflect.DeepEqual(resource.BackendMetastores, a.actual.BackendMetastores) {
		paths = append(paths, "backend_metastores")
	}

	var updated *pb.Federation
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateFederationRequest{
			Federation: resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateFederation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MetastoreFederation %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("MetastoreFederation %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated MetastoreFederation", "name", a.id)
	}

	status := &krm.MetastoreFederationStatus{}
	status.ObservedState = MetastoreFederationObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MetastoreFederationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MetastoreFederation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MetastoreFederationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Parent.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Parent.Location = a.id.Parent().Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MetastoreFederationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MetastoreFederationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MetastoreFederation", "name", a.id)

	req := &pb.DeleteFederationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteFederation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MetastoreFederation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MetastoreFederation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MetastoreFederation", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MetastoreFederation %s: %w", a.id, err)
	}
	return true, nil
}
