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

package vertexai

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.VertexAIFeaturestoreEntityTypeGVK, NewFeaturestoreEntityTypeModel)
}

func NewFeaturestoreEntityTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFeaturestoreEntityType{config: *config}, nil
}

var _ directbase.Model = &modelFeaturestoreEntityType{}

type modelFeaturestoreEntityType struct {
	config config.ControllerConfig
}

func (m *modelFeaturestoreEntityType) client(ctx context.Context, location string) (*gcp.FeaturestoreClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewFeaturestoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Featurestore client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelFeaturestoreEntityType) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIFeaturestoreEntityType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	entityTypeID, ok := id.(*krm.VertexAIFeaturestoreEntityTypeIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	gcpClient, err := m.client(ctx, entityTypeID.Location)
	if err != nil {
		return nil, err
	}
	return &FeaturestoreEntityTypeAdapter{
		id:        entityTypeID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFeaturestoreEntityType) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.VertexAIFeaturestoreEntityTypeIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &FeaturestoreEntityTypeAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type FeaturestoreEntityTypeAdapter struct {
	id        *krm.VertexAIFeaturestoreEntityTypeIdentity
	gcpClient *gcp.FeaturestoreClient
	desired   *krm.VertexAIFeaturestoreEntityType
	actual    *pb.EntityType
}

var _ directbase.Adapter = &FeaturestoreEntityTypeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FeaturestoreEntityTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FeaturestoreEntityType", "name", a.id)

	req := &pb.GetEntityTypeRequest{Name: a.id.String()}
	entityTypepb, err := a.gcpClient.GetEntityType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FeaturestoreEntityType %q: %w", a.id, err)
	}

	a.actual = entityTypepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *FeaturestoreEntityTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FeaturestoreEntityType", "name", a.id)

	u := createOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	req := &pb.CreateEntityTypeRequest{
		Parent:       a.id.ParentString(),
		EntityTypeId: a.id.EntityType,
		EntityType:   desired,
	}
	op, err := a.gcpClient.CreateEntityType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FeaturestoreEntityType %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting creation FeaturestoreEntityType %s: %w", a.id, err)
	}

	ret, err := a.gcpClient.GetEntityType(ctx, &pb.GetEntityTypeRequest{
		Name: a.id.String(),
	})
	if err != nil {
		if created != nil {
			ret = created
		} else {
			return fmt.Errorf("getting created FeaturestoreEntityType %s: %w", a.id, err)
		}
	}

	log.V(2).Info("successfully created FeaturestoreEntityType", "name", a.id)
	return a.updateStatus(ctx, createOp, ret)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *FeaturestoreEntityTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FeaturestoreEntityType", "name", a.id)

	u := updateOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	diff, updateMask, err := common.CompareBrownfieldSpecAndLabels(
		ctx,
		u,
		&a.desired.Spec,
		a.actual,
		"labels",
		VertexAIFeaturestoreEntityTypeSpec_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return fmt.Errorf("comparing actual and desired FeaturestoreEntityType %s: %w", a.id.String(), err)
	}

	if updateMask == nil || len(updateMask.Paths) == 0 {
		log.V(2).Info("no diff detected for FeaturestoreEntityType", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diff.Object = u
	structuredreporting.ReportDiff(ctx, diff)

	desired.Name = a.id.String()
	req := &pb.UpdateEntityTypeRequest{
		EntityType: desired,
		UpdateMask: updateMask,
	}

	log.V(2).Info("making VertexAI UpdateEntityType call", "request", req)

	updated, err := a.gcpClient.UpdateEntityType(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FeaturestoreEntityType %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FeaturestoreEntityTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIFeaturestoreEntityType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIFeaturestoreEntityTypeSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Featurestore = a.id.ParentString()
	obj.Spec.ResourceID = direct.LazyPtr(a.id.EntityType)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.EntityType)
	u.SetGroupVersionKind(krm.VertexAIFeaturestoreEntityTypeGVK)
	if len(a.actual.Labels) > 0 {
		u.SetLabels(a.actual.Labels)
	}

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FeaturestoreEntityTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FeaturestoreEntityType", "name", a.id)

	req := &pb.DeleteEntityTypeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEntityType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FeaturestoreEntityType, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FeaturestoreEntityType %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete FeaturestoreEntityType %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FeaturestoreEntityType", "name", a.id)
	return true, nil
}

func (a *FeaturestoreEntityTypeAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.EntityType, error) {
	mapCtx := &direct.MapContext{}
	desired := VertexAIFeaturestoreEntityTypeSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)
	return desired, nil
}

func (a *FeaturestoreEntityTypeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.EntityType) error {
	mapCtx := &direct.MapContext{}
	status := VertexAIFeaturestoreEntityTypeStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if status == nil {
		status = &krm.VertexAIFeaturestoreEntityTypeStatus{}
	}
	status.Region = direct.LazyPtr(a.id.Location)
	return op.UpdateStatus(ctx, status, nil)
}
