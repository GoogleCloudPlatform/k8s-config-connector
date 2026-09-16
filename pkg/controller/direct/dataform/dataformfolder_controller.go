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
// proto.service: google.cloud.dataform.v1beta1.Dataform
// proto.message: google.cloud.dataform.v1beta1.Folder
// crd.type: DataformFolder
// crd.version: v1alpha1

package dataform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataform/apiv1beta1"
	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataformFolderGVK, NewFolderModel)
}

func NewFolderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &folderModel{config: *config}, nil
}

var _ directbase.Model = &folderModel{}

type folderModel struct {
	config config.ControllerConfig
}

func (m *folderModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataform client: %w", err)
	}

	return gcpClient, err
}

func (m *folderModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataformFolder{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resolvedID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedID.(*krm.DataformFolderIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DataformFolderSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &folderAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *folderModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type folderAdapter struct {
	gcpClient *gcp.Client
	id        *krm.DataformFolderIdentity
	desired   *dataformpb.Folder
	actual    *dataformpb.Folder
	reader    client.Reader
}

var _ directbase.Adapter = &folderAdapter{}

func (a *folderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataformFolder", "name", a.id.String())

	req := &dataformpb.GetFolderRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataformFolder %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *folderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataformFolder", "name", a.id.String())

	desired := proto.Clone(a.desired).(*dataformpb.Folder)
	desired.Name = a.id.String()
	desired.DisplayName = a.id.Folder // Use the Folder ID as displayName as there is no KRM field for it but it is a required field in GCP Folder

	req := &dataformpb.CreateFolderRequest{
		Parent:   a.id.ParentString(),
		Folder:   desired,
		FolderId: a.id.Folder,
	}
	created, err := a.gcpClient.CreateFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataformFolder %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DataformFolder in gcp", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *folderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataformFolder", "name", a.id.String())

	diffs, _, err := compareFolder(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	return fmt.Errorf("DataformFolder is immutable and cannot be updated")
}

func compareFolder(ctx context.Context, actual, desired *dataformpb.Folder) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DataformFolderSpec_v1alpha1_FromProto, DataformFolderSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *folderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataformFolder{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataformFolderSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.PtrTo(a.id.Location)
	obj.Spec.ResourceID = direct.PtrTo(a.id.Folder)
	obj.Spec.ProjectRef = &apirefs.ProjectRef{External: "projects/" + a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Folder)
	u.SetGroupVersionKind(krm.DataformFolderGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *folderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataformFolder", "name", a.id.String())

	req := &dataformpb.DeleteFolderRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DataformFolder %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DataformFolder", "name", a.id.String())
	return true, nil
}

func (a *folderAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *dataformpb.Folder) error {
	status := &krm.DataformFolderStatus{}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}
