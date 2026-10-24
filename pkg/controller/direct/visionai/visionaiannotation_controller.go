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

// +generated:controller
// krm.group: visionai.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: VisionAIAnnotation

package visionai

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/visionai/apiv1"
	visionaipb "cloud.google.com/go/visionai/apiv1/visionaipb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VisionAIAnnotationGVK, NewAnnotationModel)
}

func NewAnnotationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAnnotation{config: *config}, nil
}

var _ directbase.Model = &modelAnnotation{}

type modelAnnotation struct {
	config config.ControllerConfig
}

func (m *modelAnnotation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VisionAIAnnotation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.VisionAIAnnotationIdentity)

	// Get visionai GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AnnotationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAnnotation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//visionai.googleapis.com/") {
		return nil, nil
	}
	id := &krm.VisionAIAnnotationIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AnnotationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelAnnotation) client(ctx context.Context) (*gcp.WarehouseClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewWarehouseRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Annotation client: %w", err)
	}
	return gcpClient, err
}

type AnnotationAdapter struct {
	id        *krm.VisionAIAnnotationIdentity
	gcpClient *gcp.WarehouseClient
	desired   *krm.VisionAIAnnotation
	actual    *visionaipb.Annotation
}

var _ directbase.Adapter = &AnnotationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AnnotationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Annotation", "name", a.id)

	req := &visionaipb.GetAnnotationRequest{
		Name: a.id.String(),
	}
	annotationpb, err := a.gcpClient.GetAnnotation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Annotation %q: %w", a.id, err)
	}

	a.actual = annotationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AnnotationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Annotation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VisionAIAnnotationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &visionaipb.CreateAnnotationRequest{
		Parent:       a.id.ParentString(),
		AnnotationId: direct.LazyPtr(a.id.Annotation),
		Annotation:   resource,
	}
	_, err := a.gcpClient.CreateAnnotation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Annotation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Annotation", "name", a.id)

	status := &krm.VisionAIAnnotationStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AnnotationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Annotation", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		VisionAIAnnotationSpec_FromProto,
		VisionAIAnnotationSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &visionaipb.UpdateAnnotationRequest{
			UpdateMask: updateMask,
			Annotation: resource,
		}
		_, err = a.gcpClient.UpdateAnnotation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Annotation %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Annotation", "name", a.id)
	}

	status := &krm.VisionAIAnnotationStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AnnotationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VisionAIAnnotation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VisionAIAnnotationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.VisionAIAnnotationGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AnnotationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Annotation", "name", a.id)

	req := &visionaipb.DeleteAnnotationRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAnnotation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Annotation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Annotation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Annotation", "name", a.id)
	return true, nil
}
