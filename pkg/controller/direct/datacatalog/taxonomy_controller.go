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
// proto.service: google.cloud.datacatalog.v1.PolicyTagManager
// proto.message: google.cloud.datacatalog.v1.Taxonomy
// crd.type: DataCatalogTaxonomy
// crd.version: v1beta1

package datacatalog

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/datacatalog/apiv1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	rg := &TaxonomyReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.DataCatalogTaxonomyGVK, NewTaxonomyModel, rg)
}

type TaxonomyReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &TaxonomyReconcileGate{}

func (r *TaxonomyReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	return r.optIn.ShouldReconcile(o)
}

func NewTaxonomyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &taxonomyModel{config: *config}, nil
}

var _ directbase.Model = &taxonomyModel{}

type taxonomyModel struct {
	config config.ControllerConfig
}

func (m *taxonomyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataCatalogTaxonomy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTaxonomyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get DataCatalog GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	policyTagClient, err := gcpClient.newPolicyTagManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	return &taxonomyAdapter{
		gcpClient: policyTagClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *taxonomyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type taxonomyAdapter struct {
	gcpClient *gcp.PolicyTagManagerClient
	id        *krm.TaxonomyIdentity
	desired   *krm.DataCatalogTaxonomy
	actual    *pb.Taxonomy
}

var _ directbase.Adapter = &taxonomyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *taxonomyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataCatalogTaxonomy", "name", a.id)

	if a.id.ID() == "" {
		log.V(2).Info("no resource ID in get indicates the create intention", "name", a.id)
		return false, nil
	}

	req := &pb.GetTaxonomyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTaxonomy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataCatalogTaxonomy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *taxonomyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataCatalogTaxonomy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataCatalogTaxonomySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	req := &pb.CreateTaxonomyRequest{
		Parent:   a.id.Parent().String(),
		Taxonomy: resource,
	}
	created, err := a.gcpClient.CreateTaxonomy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataCatalogTaxonomy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataCatalogTaxonomy", "name", a.id)

	status := &krm.DataCatalogTaxonomyStatus{}
	status.ObservedState = DataCatalogTaxonomyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	_, actualResourceID, err := krm.ParseTaxonomyExternal(created.Name)
	if err != nil {
		return fmt.Errorf("parsing the resource name in the response of CreateFirewallPolicy: %w", err)
	}
	status.ExternalRef = direct.LazyPtr(fmt.Sprintf("%s/taxonomies/%s", a.id.Parent(), actualResourceID))
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *taxonomyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataCatalogTaxonomy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataCatalogTaxonomySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	// TODO: activated_policy_types

	var updated *pb.Taxonomy
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateTaxonomyRequest{
			Taxonomy:   resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		var err error
		updated, err = a.gcpClient.UpdateTaxonomy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataCatalogTaxonomy %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataCatalogTaxonomy", "name", a.id)
	}

	status := &krm.DataCatalogTaxonomyStatus{}
	status.ObservedState = DataCatalogTaxonomyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *taxonomyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataCatalogTaxonomy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataCatalogTaxonomySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Region = a.id.Parent().Region
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DataCatalogTaxonomyGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *taxonomyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataCatalogTaxonomy", "name", a.id)

	req := &pb.DeleteTaxonomyRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTaxonomy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataCatalogTaxonomy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataCatalogTaxonomy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataCatalogTaxonomy", "name", a.id)

	return true, nil
}
