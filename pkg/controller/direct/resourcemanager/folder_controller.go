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

package resourcemanager

import (
	"context"
	"errors"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FolderGVK, NewFolderModel)
}

func NewFolderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &FolderModel{config: config}, nil
}

var _ directbase.Model = &FolderModel{}

type FolderModel struct {
	config *config.ControllerConfig
}

func newFoldersClient(ctx context.Context, config *config.ControllerConfig) (*api.FoldersClient, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewFoldersRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building folders client: %w", err)
	}
	return client, err
}

func (m *FolderModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	foldersClient, err := newFoldersClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.Folder{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	var id *krm.FolderIdentity
	if obj.Spec.ResourceID != nil || obj.Status.Name != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.FolderIdentity)
	}

	var desired *pb.Folder
	{
		mapCtx := &direct.MapContext{}
		desired = FolderSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &FolderAdapter{
		id:            id,
		foldersClient: foldersClient,
		desired:       desired,
	}, nil
}

func (m *FolderModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudresourcemanager.googleapis.com/") {
		return nil, nil
	}

	id := &krm.FolderIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	foldersClient, err := newFoldersClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	return &FolderAdapter{
		id:            id,
		foldersClient: foldersClient,
	}, nil
}

type FolderAdapter struct {
	id            *krm.FolderIdentity
	foldersClient *api.FoldersClient
	desired       *pb.Folder
	actual        *pb.Folder
}

var _ directbase.Adapter = &FolderAdapter{}

func (a *FolderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(2).Info("getting Folder", "name", fqn)

	req := &pb.GetFolderRequest{Name: fqn}
	actual, err := a.foldersClient.GetFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Folder %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *FolderAdapter) findActiveFolderByDisplayName(ctx context.Context) (*pb.Folder, error) {
	query := fmt.Sprintf("state=ACTIVE AND parent=%s AND displayName=%q", a.desired.GetParent(), a.desired.GetDisplayName())
	req := &pb.SearchFoldersRequest{
		Query: query,
	}
	it := a.foldersClient.SearchFolders(ctx, req)
	for {
		folder, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return nil, fmt.Errorf("searching for Folder under %q: %w", a.desired.GetParent(), err)
		}
		if folder.GetDisplayName() == a.desired.GetDisplayName() {
			return folder, nil
		}
	}
	return nil, nil
}

func (a *FolderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)

	// Check if there's an ACTIVE folder with the given display_name in the
	// given parent first before trying to create a new folder. This allows
	// users to acquire existing folders by specifying the folder's
	// display_name and parent.
	existing, err := a.findActiveFolderByDisplayName(ctx)
	if err != nil {
		return err
	}
	if existing != nil {
		log.V(2).Info("found existing Folder with same displayName and parent, acquiring it", "name", existing.GetName())
		return a.setResourceIDAndStatus(ctx, createOp, existing)
	}

	log.V(2).Info("creating Folder")

	req := &pb.CreateFolderRequest{
		Folder: proto.CloneOf(a.desired),
	}

	op, err := a.foldersClient.CreateFolder(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Folder: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of Folder: %w", err)
	}
	log.V(2).Info("created Folder", "name", created.GetName())

	return a.setResourceIDAndStatus(ctx, createOp, created)
}

func (a *FolderAdapter) setResourceIDAndStatus(ctx context.Context, createOp *directbase.CreateOperation, folder *pb.Folder) error {
	resourceID := strings.TrimPrefix(folder.GetName(), "folders/")
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}
	return a.updateStatus(ctx, createOp, folder)
}

func (a *FolderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	diff, updateMask, err := compareFolder(ctx, a.actual, a.desired)
	if err != nil {
		return fmt.Errorf("getting changed fields for Folder %q: %w", fqn, err)
	}

	if !diff.HasDiff() {
		log.V(2).Info("no diff detected for Folder", "name", fqn)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diff)

	latest := a.actual

	hasDisplayNameChange := false
	for _, path := range updateMask.Paths {
		if path == "display_name" {
			hasDisplayNameChange = true
			break
		}
	}

	if hasDisplayNameChange {
		log.V(2).Info("updating Folder", "name", fqn)
		req := &pb.UpdateFolderRequest{
			Folder: &pb.Folder{
				Name:        fqn,
				DisplayName: a.desired.GetDisplayName(),
			},
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"display_name"}},
		}
		op, err := a.foldersClient.UpdateFolder(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Folder %q: %w", fqn, err)
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of Folder %s: %w", fqn, err)
		}
		log.V(2).Info("updated Folder display name", "name", fqn)
	}

	if a.desired.GetParent() != a.actual.GetParent() {
		log.V(2).Info("moving Folder", "name", fqn, "from", a.actual.GetParent(), "to", a.desired.GetParent())
		req := &pb.MoveFolderRequest{
			Name:              fqn,
			DestinationParent: a.desired.GetParent(),
		}
		op, err := a.foldersClient.MoveFolder(ctx, req)
		if err != nil {
			return fmt.Errorf("moving Folder %q: %w", fqn, err)
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for move of Folder %s: %w", fqn, err)
		}
		log.V(2).Info("moved Folder parent", "name", fqn)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *FolderAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Folder) error {
	mapCtx := &direct.MapContext{}
	status := FolderStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *FolderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Folder %q actual state is nil", a.id)
	}

	obj := &krm.Folder{}
	{
		mapCtx := &direct.MapContext{}
		spec := FolderSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.FolderGVK)
	if a.id != nil {
		obj.Name = a.id.Folder
	}

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting Folder to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

func (a *FolderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()
	log.V(2).Info("deleting Folder", "name", fqn)

	req := &pb.DeleteFolderRequest{Name: fqn}
	op, err := a.foldersClient.DeleteFolder(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Folder %s: %w", fqn, err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of Folder %s: %w", fqn, err)
	}

	return true, nil
}

func compareFolder(ctx context.Context, actual, desired *pb.Folder) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	actualMasked, err := mappers.OnlySpecFields(actual, FolderSpec_FromProto, FolderSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	return tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), actualMasked.ProtoReflect())
}
