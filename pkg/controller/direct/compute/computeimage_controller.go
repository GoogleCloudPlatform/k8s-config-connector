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

// +tool:controller
// proto.service: google.cloud.compute.v1.Images
// proto.message: google.cloud.compute.v1.Image
// crd.type: ComputeImage
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"
	"reflect"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeImageGVK, NewComputeImageModel)
}

func NewComputeImageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeImageModel{config: config}, nil
}

var _ directbase.Model = &computeImageModel{}

type computeImageModel struct {
	config *config.ControllerConfig
}

func (m *computeImageModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeImage{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	imagesClient, err := gcpClient.newImagesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeImageSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	resource.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &ComputeImageAdapter{
		gcpClient: imagesClient,
		id:        id.(*krm.ComputeImageIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *computeImageModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeImageAdapter struct {
	gcpClient *compute.ImagesClient
	id        *krm.ComputeImageIdentity
	desired   *computepb.Image
	actual    *computepb.Image
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeImageAdapter{}

func (a *ComputeImageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeImage", "name", a.id)

	req := &computepb.GetImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeImage %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeImageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeImage", "name", a.id)
	mapCtx := &direct.MapContext{}

	a.desired.Name = direct.LazyPtr(a.id.Image)

	req := &computepb.InsertImageRequest{
		Project:       a.id.Project,
		ImageResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeImage %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeImage %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeImage in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
	}

	status := ComputeImageStatus_v1beta1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeImageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeImage", "name", a.id)
	mapCtx := &direct.MapContext{}

	// Update labels if needed
	desiredLabels := updateOp.GetUnstructured().GetLabels()
	actualLabels := a.actual.Labels
	labelsChanged := !reflect.DeepEqual(desiredLabels, actualLabels)

	diffs, _, err := compareComputeImage(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.Image
	if !diffs.HasDiff() && !labelsChanged {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		if diffs.HasDiff() {
			structuredreporting.ReportDiff(ctx, diffs)

			a.desired.Name = direct.LazyPtr(a.id.Image)

			req := &computepb.PatchImageRequest{
				Project:       a.id.Project,
				Image:         a.id.Image,
				ImageResource: a.desired,
			}
			op, err := a.gcpClient.Patch(ctx, req)
			if err != nil {
				return fmt.Errorf("updating compute ComputeImage %s: %w", a.id.String(), err)
			}
			log.V(2).Info("successfully updated compute ComputeImage", "name", a.id.String())

			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("compute ComputeImage %s waiting for update: %w", a.id.String(), err)
			}
		}

		if labelsChanged {
			log.Info("updating labels", "name", a.id)
			req := &computepb.SetLabelsImageRequest{
				Project:  a.id.Project,
				Resource: a.id.Image,
				GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{
					LabelFingerprint: a.actual.LabelFingerprint,
					Labels:           desiredLabels,
				},
			}
			op, err := a.gcpClient.SetLabels(ctx, req)
			if err != nil {
				return fmt.Errorf("setting labels: %w", err)
			}
			if err := op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting for labels update: %w", err)
			}
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
		}
	}

	status := ComputeImageStatus_v1beta1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeImageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeImage{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeImageSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Image)
	u.SetGroupVersionKind(krm.ComputeImageGVK)

	u.Object = uObj
	u.SetLabels(a.actual.Labels)
	return u, nil
}

func (a *ComputeImageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeImage", "name", a.id)

	req := &computepb.DeleteImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeImage %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeImage", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeImage %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeImageAdapter) get(ctx context.Context) (*computepb.Image, error) {
	getReq := &computepb.GetImageRequest{
		Project: a.id.Project,
		Image:   a.id.Image,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeImage %s: %w", a.id, err)
	}
	return resource, nil
}

func compareComputeImage(ctx context.Context, actual, desired *computepb.Image) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeImageSpec_v1beta1_FromProto, ComputeImageSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
