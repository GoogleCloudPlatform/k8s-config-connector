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
// proto.service: google.cloud.dataform.v1.Dataform
// proto.message: google.cloud.dataform.v1.TeamFolder
// crd.type: DataformTeamFolder
// crd.version: v1alpha1

package dataform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataform/apiv1"
	dataformpb "cloud.google.com/go/dataform/apiv1/dataformpb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataformTeamFolderGVK, NewTeamFolderModel)
}

func NewTeamFolderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &teamFolderModel{config: *config}, nil
}

var _ directbase.Model = &teamFolderModel{}

type teamFolderModel struct {
	config config.ControllerConfig
}

func (m *teamFolderModel) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *teamFolderModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataformTeamFolder{}
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
	id := resolvedID.(*krm.DataformTeamFolderIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DataformTeamFolderSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &teamFolderAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *teamFolderModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type teamFolderAdapter struct {
	gcpClient *gcp.Client
	id        *krm.DataformTeamFolderIdentity
	desired   *dataformpb.TeamFolder
	actual    *dataformpb.TeamFolder
	reader    client.Reader
}

var _ directbase.Adapter = &teamFolderAdapter{}

func (a *teamFolderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataformTeamFolder", "name", a.id.String())

	req := &dataformpb.GetTeamFolderRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTeamFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataformTeamFolder %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *teamFolderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataformTeamFolder", "name", a.id.String())

	desired := proto.Clone(a.desired).(*dataformpb.TeamFolder)
	desired.Name = a.id.String()

	req := &dataformpb.CreateTeamFolderRequest{
		Parent:     a.id.ParentString(),
		TeamFolder: desired,
	}
	created, err := a.gcpClient.CreateTeamFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataformTeamFolder %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DataformTeamFolder in gcp", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *teamFolderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataformTeamFolder", "name", a.id.String())

	diffs, updateMask, err := compareTeamFolder(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	desired := proto.Clone(a.desired).(*dataformpb.TeamFolder)
	desired.Name = a.id.String()

	req := &dataformpb.UpdateTeamFolderRequest{
		TeamFolder: desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateTeamFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DataformTeamFolder %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareTeamFolder(ctx context.Context, actual, desired *dataformpb.TeamFolder) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DataformTeamFolderSpec_v1alpha1_FromProto, DataformTeamFolderSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *teamFolderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataformTeamFolder{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataformTeamFolderSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.PtrTo(a.id.Location)
	obj.Spec.ResourceID = direct.PtrTo(a.id.Team_folder)
	obj.Spec.ProjectRef = &apirefs.ProjectRef{External: "projects/" + a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Team_folder)
	u.SetGroupVersionKind(krm.DataformTeamFolderGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *teamFolderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataformTeamFolder", "name", a.id.String())

	req := &dataformpb.DeleteTeamFolderRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTeamFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DataformTeamFolder %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DataformTeamFolder", "name", a.id.String())
	return true, nil
}

func (a *teamFolderAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *dataformpb.TeamFolder) error {
	mapCtx := &direct.MapContext{}
	observedState := DataformTeamFolderObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status := &krm.DataformTeamFolderStatus{
		ObservedState: observedState,
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}
