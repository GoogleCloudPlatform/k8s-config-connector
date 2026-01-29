// Copyright 2025 Google LLC
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

package clouddeploy

import (
	"context"
	"fmt"

	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/deploy/apiv1"
	clouddeploypb "cloud.google.com/go/deploy/apiv1/deploypb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krmv1beta1.CloudDeployDeliveryPipelineGVK, NewDeliveryPipelineModel)
}

func NewDeliveryPipelineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDeliveryPipeline{config: *config}, nil
}

var _ directbase.Model = &modelDeliveryPipeline{}

type modelDeliveryPipeline struct {
	config config.ControllerConfig
}

func (m *modelDeliveryPipeline) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DeliveryPipeline client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDeliveryPipeline) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krmv1beta1.CloudDeployDeliveryPipeline{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krmv1beta1.NewDeliveryPipelineIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get clouddeploy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeliveryPipelineAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDeliveryPipeline) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DeliveryPipelineAdapter struct {
	id        *krmv1beta1.DeliveryPipelineIdentity
	gcpClient *gcp.CloudDeployClient
	desired   *krmv1beta1.CloudDeployDeliveryPipeline
	actual    *clouddeploypb.DeliveryPipeline
}

var _ directbase.Adapter = &DeliveryPipelineAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DeliveryPipelineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeliveryPipeline", "name", a.id)

	req := &clouddeploypb.GetDeliveryPipelineRequest{Name: a.id.String()}
	deliverypipelinepb, err := a.gcpClient.GetDeliveryPipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeliveryPipeline %q: %w", a.id, err)
	}

	a.actual = deliverypipelinepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeliveryPipelineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeliveryPipeline", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeliveryPipelineSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &clouddeploypb.CreateDeliveryPipelineRequest{
		Parent:             a.id.Parent().String(),
		DeliveryPipeline:   resource,
		DeliveryPipelineId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateDeliveryPipeline(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeliveryPipeline %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeliveryPipeline %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeliveryPipeline", "name", a.id)

	status := &krmv1beta1.DeliveryPipelineStatus{}
	status.ObservedState = DeliveryPipelineObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeliveryPipelineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeliveryPipeline", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := DeliveryPipelineSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	desiredPb.Name = a.id.String()
	req := &clouddeploypb.UpdateDeliveryPipelineRequest{
		UpdateMask:       updateMask,
		DeliveryPipeline: desiredPb,
	}
	op, err := a.gcpClient.UpdateDeliveryPipeline(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DeliveryPipeline %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeliveryPipeline %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated DeliveryPipeline", "name", a.id)

	status := &krmv1beta1.DeliveryPipelineStatus{}
	status.ObservedState = DeliveryPipelineObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DeliveryPipelineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krmv1beta1.CloudDeployDeliveryPipeline{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DeliveryPipelineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = &a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krmv1beta1.CloudDeployDeliveryPipelineGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DeliveryPipelineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeliveryPipeline", "name", a.id)

	req := &clouddeploypb.DeleteDeliveryPipelineRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDeliveryPipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DeliveryPipeline, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DeliveryPipeline %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DeliveryPipeline", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete DeliveryPipeline %s: %w", a.id, err)
	}
	return true, nil
}
