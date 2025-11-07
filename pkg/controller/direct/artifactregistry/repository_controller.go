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
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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

func (m *model) client(ctx context.Context) (*artifactregistry.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, fmt.Errorf("error getting REST client options: %w", err)
	}

	gcpClient, err := artifactregistry.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building artifact registry client: %w", err)
	}

	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ArtifactRegistryRepository{}
	copied := u.DeepCopy()
	if err := label.ComputeLabels(copied); err != nil {
		return nil, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copied.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := ResolveArtifactRegistryRepositoryRefs(ctx, reader, &obj); err != nil {
		return nil, err
	}

	id, err := NewArtifactRegistryRepositoryIdentity(ctx, reader, obj, copied)
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
	gcpClient *artifactregistry.Client
	desired   *krm.ArtifactRegistryRepository
	actual    *artifactregistrypb.Repository
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	req := &artifactregistrypb.GetRepositoryRequest{Name: a.id.FullyQualifiedName()}

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

	req := &artifactregistrypb.CreateRepositoryRequest{
		Parent:       a.id.Parent.String(),
		RepositoryId: a.id.ResourceID,
		Repository:   resource,
	}

	_, err := a.gcpClient.CreateRepository(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ArtifactRegistry repository %q: %w", a.id.FullyQualifiedName(), err)
	}

	log.V(2).Info("successfully created ArtifactRegistry repository", "name", a.id.FullyQualifiedName())
	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ArtifactRegistry repository", "name", a.id.FullyQualifiedName())

	if a.actual == nil {
		return fmt.Errorf("Update called without a prior call to Find")
	}

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ArtifactRegistryRepositorySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.FullyQualifiedName()

	// Validate critical immutable fields explicitly
	if a.desired.Spec.Format != "" && a.actual.Format.String() != "" && a.desired.Spec.Format != a.actual.Format.String() {
		return fmt.Errorf("field 'spec.format' is immutable and cannot be updated from %q to %q", a.actual.Format.String(), a.desired.Spec.Format)
	}

	if a.desired.Spec.Location != "" && a.actual.GetName() != "" {
		// Extract location from actual name path (projects/.../locations/LOCATION/repositories/...)
		if !strings.Contains(a.actual.GetName(), "locations/"+a.desired.Spec.Location+"/") {
			return fmt.Errorf("field 'spec.location' is immutable and cannot be updated")
		}
	}

	if a.desired.Spec.Mode != nil && a.actual.Mode.String() != "" && *a.desired.Spec.Mode != a.actual.Mode.String() {
		return fmt.Errorf("field 'spec.mode' is immutable and cannot be updated from %q to %q", a.actual.Mode.String(), *a.desired.Spec.Mode)
	}

	// Validate Maven config immutable fields
	if a.desired.Spec.MavenConfig != nil && a.actual.GetMavenConfig() != nil {
		if a.desired.Spec.MavenConfig.AllowSnapshotOverwrites != nil &&
			a.actual.GetMavenConfig().GetAllowSnapshotOverwrites() != *a.desired.Spec.MavenConfig.AllowSnapshotOverwrites {
			return fmt.Errorf("field 'spec.mavenConfig.allowSnapshotOverwrites' is immutable and cannot be updated from %v to %v",
				a.actual.GetMavenConfig().GetAllowSnapshotOverwrites(), *a.desired.Spec.MavenConfig.AllowSnapshotOverwrites)
		}
		if a.desired.Spec.MavenConfig.VersionPolicy != nil &&
			a.actual.GetMavenConfig().GetVersionPolicy().String() != *a.desired.Spec.MavenConfig.VersionPolicy {
			return fmt.Errorf("field 'spec.mavenConfig.versionPolicy' is immutable and cannot be updated from %q to %q",
				a.actual.GetMavenConfig().GetVersionPolicy().String(), *a.desired.Spec.MavenConfig.VersionPolicy)
		}
	}

	// Validate RemoteRepositoryConfig is immutable (entire block)
	if (a.desired.Spec.RemoteRepositoryConfig != nil) != (a.actual.GetRemoteRepositoryConfig() != nil) {
		return fmt.Errorf("field 'spec.remoteRepositoryConfig' is immutable and cannot be changed after creation")
	}

	// Use common.CompareProtoMessage to automatically detect changes and validate immutable fields
	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing desired and actual state: %w", err)
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.FullyQualifiedName())
		return nil
	}

	updateMask := &fieldmaskpb.FieldMask{}
	updateMask.Paths = sets.List(paths)

	req := &artifactregistrypb.UpdateRepositoryRequest{
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
	// Set status fields directly (flat structure for backward compatibility)
	if a.actual.CreateTime != nil {
		obj.Status.CreateTime = direct.PtrTo(a.actual.CreateTime.AsTime().Format(time.RFC3339))
	}
	if a.actual.UpdateTime != nil {
		obj.Status.UpdateTime = direct.PtrTo(a.actual.UpdateTime.AsTime().Format(time.RFC3339))
	}
	// Extract just the repository name from the full path
	if a.actual.Name != "" {
		parts := strings.Split(a.actual.Name, "/")
		if len(parts) > 0 {
			obj.Status.Name = direct.PtrTo(parts[len(parts)-1])
		}
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

	req := &artifactregistrypb.DeleteRepositoryRequest{Name: a.id.FullyQualifiedName()}

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
