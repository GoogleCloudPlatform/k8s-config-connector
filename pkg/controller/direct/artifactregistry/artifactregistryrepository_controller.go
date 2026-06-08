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

package artifactregistry

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/artifactregistry/apiv1"
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ArtifactRegistryRepositoryGVK, NewArtifactRegistryRepositoryModel)
}

func NewArtifactRegistryRepositoryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelArtifactRegistryRepository{config: *config}, nil
}

var _ directbase.Model = &modelArtifactRegistryRepository{}

type modelArtifactRegistryRepository struct {
	config config.ControllerConfig
}

func (m *modelArtifactRegistryRepository) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ArtifactRegistryRepository client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelArtifactRegistryRepository) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ArtifactRegistryRepository{}
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

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ArtifactRegistryRepositorySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.GCPLabels(obj)

	return &ArtifactRegistryRepositoryAdapter{
		id:        id.(*krm.ArtifactRegistryRepositoryIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelArtifactRegistryRepository) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ArtifactRegistryRepositoryAdapter struct {
	id        *krm.ArtifactRegistryRepositoryIdentity
	gcpClient *gcp.Client
	desired   *pb.Repository
	actual    *pb.Repository
}

var _ directbase.Adapter = &ArtifactRegistryRepositoryAdapter{}

func (a *ArtifactRegistryRepositoryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ArtifactRegistryRepository", "name", a.id)

	req := &pb.GetRepositoryRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ArtifactRegistryRepository %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ArtifactRegistryRepositoryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ArtifactRegistryRepository", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	repositoryID := a.id.Repository

	req := &pb.CreateRepositoryRequest{
		Parent:       parent,
		Repository:   a.desired,
		RepositoryId: repositoryID,
	}
	op, err := a.gcpClient.CreateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ArtifactRegistryRepository %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ArtifactRegistryRepository %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ArtifactRegistryRepository", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ArtifactRegistryRepositoryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ArtifactRegistryRepository", "name", a.id.String())

	diffs, updateMask, err := compareRepository(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateRepositoryRequest{
			Repository: a.desired,
			UpdateMask: updateMask,
		}
		req.Repository.Name = a.id.String()

		updated, err := a.gcpClient.UpdateRepository(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ArtifactRegistryRepository %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ArtifactRegistryRepositoryAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Repository) error {
	mapCtx := &direct.MapContext{}
	status := ArtifactRegistryRepositoryStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ArtifactRegistryRepositoryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ArtifactRegistryRepository{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ArtifactRegistryRepositorySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ArtifactRegistryRepositoryGVK)
	return u, nil
}

func (a *ArtifactRegistryRepositoryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ArtifactRegistryRepository", "name", a.id)

	req := &pb.DeleteRepositoryRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ArtifactRegistryRepository, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ArtifactRegistryRepository %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("ArtifactRegistryRepository %s waiting deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ArtifactRegistryRepository", "name", a.id)
	return true, nil
}

func compareRepository(ctx context.Context, actual, desired *pb.Repository) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ArtifactRegistryRepositorySpec_FromProto, ArtifactRegistryRepositorySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	populateDefaults := func(obj *pb.Repository) {
		if obj.Mode == pb.Repository_MODE_UNSPECIFIED {
			obj.Mode = pb.Repository_STANDARD_REPOSITORY
		}
	}

	desired = proto.Clone(desired).(*pb.Repository)
	populateDefaults(desired)
	populateDefaults(maskedActual)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
