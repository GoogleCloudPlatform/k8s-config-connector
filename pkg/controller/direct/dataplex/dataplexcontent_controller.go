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
// proto.service: google.cloud.dataplex.v1.ContentService
// proto.message: google.cloud.dataplex.v1.Content
// crd.type: DataplexContent
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
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
	registry.RegisterModel(krm.DataplexContentGVK, NewContentModel)
}

func NewContentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &contentModel{config: config}, nil
}

var _ directbase.Model = &contentModel{}

type contentModel struct {
	config *config.ControllerConfig
}

func (m *contentModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexContent{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewContentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	//// normalize reference fields
	//if obj.Spec.LakeRef != nil {
	//	external, err := obj.Spec.LakeRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	//	if err != nil {
	//		return nil, err
	//	}
	//	obj.Spec.LakeRef.External = external
	//}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	contentClient, err := gcpClient.newContentClient(ctx)
	if err != nil {
		return nil, err
	}

	return &contentAdapter{
		gcpClient: contentClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *contentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type contentAdapter struct {
	gcpClient *gcp.ContentClient
	id        *krm.ContentIdentity
	desired   *krm.DataplexContent
	actual    *pb.Content
	reader    client.Reader
}

var _ directbase.Adapter = &contentAdapter{}

func (a *contentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex content", "name", a.id)

	req := &pb.GetContentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetContent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex content %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *contentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex content", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()

	resource := DataplexContentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &pb.CreateContentRequest{
		Parent:  a.id.Parent(),
		Content: resource,
	}
	created, err := a.gcpClient.CreateContent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex content %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created dataplex content in gcp", "name", a.id)

	status := &krm.DataplexContentStatus{}
	status.ObservedState = DataplexContentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *contentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex content", "name", a.id)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexContentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if desired.Spec.DataText != nil && !reflect.DeepEqual(resource.GetDataText(), a.actual.GetDataText()) {
		updateMask.Paths = append(updateMask.Paths, "data_text")
	}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if desired.Spec.Path != nil && !reflect.DeepEqual(resource.Path, a.actual.Path) {
		updateMask.Paths = append(updateMask.Paths, "path")
	}
	if desired.Spec.SQLScript != nil && !reflect.DeepEqual(resource.GetSqlScript(), a.actual.GetSqlScript()) {
		updateMask.Paths = append(updateMask.Paths, "sql_script")
	}
	if desired.Spec.Notebook != nil && !reflect.DeepEqual(resource.GetNotebook(), a.actual.GetNotebook()) {
		updateMask.Paths = append(updateMask.Paths, "notebook")
	}

	var updated *pb.Content
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		req := &pb.UpdateContentRequest{
			UpdateMask: updateMask,
			Content:    resource,
		}
		var err error
		updated, err = a.gcpClient.UpdateContent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex content %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex content", "name", a.id)
	}

	status := &krm.DataplexContentStatus{}
	status.ObservedState = DataplexContentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *contentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexContent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexContentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.LakeRef = &krm.LakeRef{External: a.id.Parent()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID()) // Use the path-like identifier as the KCC name
	u.SetGroupVersionKind(krm.DataplexContentGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *contentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex content", "name", a.id)

	req := &pb.DeleteContentRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteContent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent dataplex content, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex content %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex content", "name", a.id)

	return true, nil
}
