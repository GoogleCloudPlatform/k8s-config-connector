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
// krm.group: vision.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: VisionReferenceImage

package vision

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vision/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/vision/v2/apiv1"
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VisionReferenceImageGVK, NewReferenceImageModel)
}

func NewReferenceImageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReferenceImage{config: *config}, nil
}

var _ directbase.Model = &modelReferenceImage{}

type modelReferenceImage struct {
	config config.ControllerConfig
}

func (m *modelReferenceImage) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VisionReferenceImage{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.VisionReferenceImageIdentity)

	// Get vision GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReferenceImageAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReferenceImage) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//vision.googleapis.com/") {
		return nil, nil
	}
	id := &krm.VisionReferenceImageIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReferenceImageAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelReferenceImage) client(ctx context.Context) (*gcp.ProductSearchClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewProductSearchRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReferenceImage client: %w", err)
	}
	return gcpClient, err
}

type ReferenceImageAdapter struct {
	id        *krm.VisionReferenceImageIdentity
	gcpClient *gcp.ProductSearchClient
	desired   *krm.VisionReferenceImage
	actual    *visionpb.ReferenceImage
}

var _ directbase.Adapter = &ReferenceImageAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReferenceImageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReferenceImage", "name", a.id)

	req := &visionpb.GetReferenceImageRequest{
		Name: a.id.String(),
	}
	referenceimagepb, err := a.gcpClient.GetReferenceImage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReferenceImage %q: %w", a.id, err)
	}

	a.actual = referenceimagepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReferenceImageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReferenceImage", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VisionReferenceImageSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &visionpb.CreateReferenceImageRequest{
		Parent:           a.id.ParentString(),
		ReferenceImageId: a.id.ReferenceImage,
		ReferenceImage:   resource,
	}
	_, err := a.gcpClient.CreateReferenceImage(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReferenceImage %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ReferenceImage", "name", a.id)

	status := &krm.VisionReferenceImageStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReferenceImageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReferenceImage", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		VisionReferenceImageSpec_FromProto,
		VisionReferenceImageSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating VisionReferenceImage is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.VisionReferenceImageStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReferenceImageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VisionReferenceImage{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VisionReferenceImageSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.VisionReferenceImageGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReferenceImageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReferenceImage", "name", a.id)

	req := &visionpb.DeleteReferenceImageRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteReferenceImage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ReferenceImage, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ReferenceImage %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ReferenceImage", "name", a.id)
	return true, nil
}
