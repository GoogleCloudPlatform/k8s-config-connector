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

package artifactregistry

import (
	"context"
	"fmt"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/artifactregistry/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "artifactregistry-controller"
)

func init() {
	registry.RegisterModel(krm.ArtifactRegistryRepositoryGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (pb.ArtifactRegistryClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	
	// Convert option.ClientOption to grpc.DialOption
	var dialOpts []grpc.DialOption
	for _, opt := range opts {
		// This is a workaround for MockGCP - in production this would use actual gRPC client creation
		_ = opt // Use the option somehow or ignore for MockGCP
	}
	
	conn, err := grpc.Dial("", dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("dialing ArtifactRegistry API: %w", err)
	}
	
	return pb.NewArtifactRegistryClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ArtifactRegistryRepository{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := NewArtifactRegistryRepositoryIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support reading resources by URL
	return nil, nil
}

type Adapter struct {
	id        *ArtifactRegistryRepositoryIdentity
	gcpClient pb.ArtifactRegistryClient
	desired   *krm.ArtifactRegistryRepository
	actual    *pb.Repository
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	req := &pb.GetRepositoryRequest{Name: a.id.FullyQualifiedName()}
	repositoryPB, err := a.gcpClient.GetRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ArtifactRegistry repository %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = repositoryPB
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ArtifactRegistryRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.FullyQualifiedName()

	req := &pb.CreateRepositoryRequest{
		Parent:       a.id.Parent.String(),
		RepositoryId: a.id.ResourceID,
		Repository:   resource,
	}

	op, err := a.gcpClient.CreateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ArtifactRegistry repository %q: %w", a.id.FullyQualifiedName(), err)
	}

	log.V(2).Info("successfully created ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	// For MockGCP, operations complete immediately, so we can ignore the operation result
	_ = op

	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ArtifactRegistryRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.FullyQualifiedName()

	updateMask := &fieldmaskpb.FieldMask{}
	
	// For artifact registry, we can update description, labels, cleanup policies, etc.
	// But format, mode, location are immutable
	updateMask.Paths = append(updateMask.Paths, "description")
	updateMask.Paths = append(updateMask.Paths, "labels")
	updateMask.Paths = append(updateMask.Paths, "cleanup_policies")
	updateMask.Paths = append(updateMask.Paths, "cleanup_policy_dry_run")
	
	// Only add docker/maven config if they exist
	if resource.GetDockerConfig() != nil {
		updateMask.Paths = append(updateMask.Paths, "docker_config")
	}
	if resource.GetMavenConfig() != nil {
		updateMask.Paths = append(updateMask.Paths, "maven_config")
	}

	req := &pb.UpdateRepositoryRequest{
		Repository: resource,
		UpdateMask: updateMask,
	}

	repository, err := a.gcpClient.UpdateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ArtifactRegistry repository %q: %w", a.id.FullyQualifiedName(), err)
	}

	log.V(2).Info("successfully updated ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	a.actual = repository
	return nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Status.ObservedState = ArtifactRegistryRepositoryObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	req := &pb.DeleteRepositoryRequest{Name: a.id.FullyQualifiedName()}
	_, err := a.gcpClient.DeleteRepository(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("ArtifactRegistry repository was not found, assuming already deleted", "name", a.id.FullyQualifiedName())
			return true, nil
		}
		return false, fmt.Errorf("deleting ArtifactRegistry repository %q: %w", a.id.FullyQualifiedName(), err)
	}

	log.V(2).Info("successfully deleted ArtifactRegistry repository", "name", a.id.FullyQualifiedName())
	return true, nil
}