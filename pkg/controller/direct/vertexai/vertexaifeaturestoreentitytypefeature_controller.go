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
	registry.RegisterModel(krm.VertexAIFeaturestoreEntityTypeFeatureGVK, NewFeaturestoreEntityTypeFeatureModel)
}

func NewFeaturestoreEntityTypeFeatureModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFeaturestoreEntityTypeFeature{config: *config}, nil
}

var _ directbase.Model = &modelFeaturestoreEntityTypeFeature{}

type modelFeaturestoreEntityTypeFeature struct {
	config config.ControllerConfig
}

func (m *modelFeaturestoreEntityTypeFeature) client(ctx context.Context, location string) (*gcp.FeaturestoreClient, error) {
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

func (m *modelFeaturestoreEntityTypeFeature) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIFeaturestoreEntityTypeFeature{}
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
	featureID, ok := id.(*krm.VertexAIFeaturestoreEntityTypeFeatureIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	gcpClient, err := m.client(ctx, featureID.Location)
	if err != nil {
		return nil, err
	}
	return &FeaturestoreEntityTypeFeatureAdapter{
		id:        featureID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelFeaturestoreEntityTypeFeature) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.VertexAIFeaturestoreEntityTypeFeatureIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &FeaturestoreEntityTypeFeatureAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type FeaturestoreEntityTypeFeatureAdapter struct {
	id        *krm.VertexAIFeaturestoreEntityTypeFeatureIdentity
	gcpClient *gcp.FeaturestoreClient
	desired   *krm.VertexAIFeaturestoreEntityTypeFeature
	actual    *pb.Feature
}

var _ directbase.Adapter = &FeaturestoreEntityTypeFeatureAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FeaturestoreEntityTypeFeatureAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FeaturestoreEntityTypeFeature", "name", a.id)

	req := &pb.GetFeatureRequest{Name: a.id.String()}
	featurepb, err := a.gcpClient.GetFeature(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FeaturestoreEntityTypeFeature %q: %w", a.id, err)
	}

	a.actual = featurepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *FeaturestoreEntityTypeFeatureAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FeaturestoreEntityTypeFeature", "name", a.id)

	u := createOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	req := &pb.CreateFeatureRequest{
		Parent:    a.id.ParentString(),
		FeatureId: a.id.Feature,
		Feature:   desired,
	}
	op, err := a.gcpClient.CreateFeature(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FeaturestoreEntityTypeFeature %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting creation FeaturestoreEntityTypeFeature %s: %w", a.id, err)
	}

	getReq := &pb.GetFeatureRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetFeature(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting FeaturestoreEntityTypeFeature after create %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created FeaturestoreEntityTypeFeature", "name", a.id)
	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *FeaturestoreEntityTypeFeatureAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FeaturestoreEntityTypeFeature", "name", a.id)

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
		VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return fmt.Errorf("comparing actual and desired FeaturestoreEntityTypeFeature %s: %w", a.id.String(), err)
	}

	if updateMask == nil || len(updateMask.Paths) == 0 {
		log.V(2).Info("no diff detected for FeaturestoreEntityTypeFeature", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diff.Object = u
	structuredreporting.ReportDiff(ctx, diff)

	desired.Name = a.id.String()
	req := &pb.UpdateFeatureRequest{
		Feature:    desired,
		UpdateMask: updateMask,
	}

	log.V(2).Info("making VertexAI UpdateFeature call", "request", req)

	updated, err := a.gcpClient.UpdateFeature(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FeaturestoreEntityTypeFeature %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FeaturestoreEntityTypeFeatureAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIFeaturestoreEntityTypeFeature{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Entitytype = a.id.ParentString()
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Feature)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Feature)
	u.SetGroupVersionKind(krm.VertexAIFeaturestoreEntityTypeFeatureGVK)
	if len(a.actual.Labels) > 0 {
		u.SetLabels(a.actual.Labels)
	}

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FeaturestoreEntityTypeFeatureAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FeaturestoreEntityTypeFeature", "name", a.id)

	req := &pb.DeleteFeatureRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteFeature(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FeaturestoreEntityTypeFeature, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting FeaturestoreEntityTypeFeature %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete FeaturestoreEntityTypeFeature %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FeaturestoreEntityTypeFeature", "name", a.id)
	return true, nil
}

func (a *FeaturestoreEntityTypeFeatureAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.Feature, error) {
	mapCtx := &direct.MapContext{}
	desired := VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)
	return desired, nil
}

func (a *FeaturestoreEntityTypeFeatureAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Feature) error {
	mapCtx := &direct.MapContext{}
	status := VertexAIFeaturestoreEntityTypeFeatureStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if status == nil {
		status = &krm.VertexAIFeaturestoreEntityTypeFeatureStatus{}
	}
	status.Region = direct.LazyPtr(a.id.Location)
	return op.UpdateStatus(ctx, status, nil)
}
