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

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeProjectMetadataGVK, NewComputeProjectMetadataModel)
}

func NewComputeProjectMetadataModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeProjectMetadataModel{config: config}, nil
}

type computeProjectMetadataModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &computeProjectMetadataModel{}

func (m *computeProjectMetadataModel) client(ctx context.Context) (*compute.ProjectsClient, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	return gcpClient.newProjectsClient(ctx)
}

func (m *computeProjectMetadataModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeProjectMetadata{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ComputeProjectMetadataIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	projectsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeProjectMetadataAdapter{
		gcpClient: projectsClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeProjectMetadataModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeProjectMetadataIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	projectsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeProjectMetadataAdapter{
		gcpClient: projectsClient,
		id:        id,
	}, nil
}

type ComputeProjectMetadataAdapter struct {
	gcpClient *compute.ProjectsClient
	id        *krm.ComputeProjectMetadataIdentity
	desired   *krm.ComputeProjectMetadata
	actual    *pb.Metadata
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeProjectMetadataAdapter{}

func (a *ComputeProjectMetadataAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeProjectMetadata", "project", a.id.Project)

	req := &computepb.GetProjectRequest{
		Project: a.id.Project,
	}
	project, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting compute project %q: %w", a.id.Project, err)
	}

	a.actual = project.GetCommonInstanceMetadata()
	if a.actual == nil {
		a.actual = &computepb.Metadata{}
	}
	return true, nil
}

func (a *ComputeProjectMetadataAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := ComputeProjectMetadataSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return a.save(ctx, createOp.GetUnstructured(), createOp, desired)
}

func (a *ComputeProjectMetadataAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := ComputeProjectMetadataSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return a.save(ctx, updateOp.GetUnstructured(), updateOp, desired)
}

func (a *ComputeProjectMetadataAdapter) save(ctx context.Context, u *unstructured.Unstructured, op directbase.Operation, desired *pb.Metadata) error {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		found, err := a.Find(ctx)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("compute project %q not found", a.id.Project)
		}
	}

	diffs, err := compareComputeProjectMetadata(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		log.V(2).Info("updating ComputeProjectMetadata", "project", a.id.Project)
		diffs.Object = u
		structuredreporting.ReportDiff(ctx, diffs)

		metadataResource := proto.Clone(desired).(*pb.Metadata)
		metadataResource.Fingerprint = a.actual.Fingerprint

		req := &computepb.SetCommonInstanceMetadataProjectRequest{
			Project:          a.id.Project,
			MetadataResource: metadataResource,
		}
		callOp, err := a.gcpClient.SetCommonInstanceMetadata(ctx, req)
		if err != nil {
			return fmt.Errorf("setting common instance metadata for %q: %w", a.id.Project, err)
		}
		if err := callOp.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for setCommonInstanceMetadata for %q: %w", a.id.Project, err)
		}

		latestProject, err := a.gcpClient.Get(ctx, &computepb.GetProjectRequest{Project: a.id.Project})
		if err != nil {
			return fmt.Errorf("getting project after update %q: %w", a.id.Project, err)
		}
		a.actual = latestProject.GetCommonInstanceMetadata()
		if a.actual == nil {
			a.actual = &computepb.Metadata{}
		}
	}

	return a.updateStatus(ctx, op, a.actual)
}

func (a *ComputeProjectMetadataAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	mapCtx := &direct.MapContext{}
	spec := ComputeProjectMetadataSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.ComputeProjectMetadata{}
	obj.Spec = *spec
	obj.Name = a.id.Project
	obj.SetGroupVersionKind(krm.ComputeProjectMetadataGVK)

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}
	return &unstructured.Unstructured{Object: u}, nil
}

func (a *ComputeProjectMetadataAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeProjectMetadata", "project", a.id.Project)

	if a.actual == nil {
		found, err := a.Find(ctx)
		if err != nil {
			return false, err
		}
		if !found {
			return true, nil
		}
	}

	if len(a.actual.GetItems()) == 0 {
		return true, nil
	}

	emptyMetadata := &computepb.Metadata{
		Fingerprint: a.actual.Fingerprint,
	}
	req := &computepb.SetCommonInstanceMetadataProjectRequest{
		Project:          a.id.Project,
		MetadataResource: emptyMetadata,
	}
	op, err := a.gcpClient.SetCommonInstanceMetadata(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("clearing common instance metadata for %q: %w", a.id.Project, err)
	}

	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for clearing common instance metadata for %q: %w", a.id.Project, err)
	}

	log.V(2).Info("successfully cleared ComputeProjectMetadata", "project", a.id.Project)
	return true, nil
}

func compareComputeProjectMetadata(ctx context.Context, actual, desired *pb.Metadata) (*structuredreporting.Diff, error) {
	if actual == nil {
		actual = &pb.Metadata{}
	}
	if desired == nil {
		desired = &pb.Metadata{}
	}

	maskedActual, err := mappers.OnlySpecFields(actual, ComputeProjectMetadataSpec_FromProto, ComputeProjectMetadataSpec_ToProto)
	if err != nil {
		return nil, err
	}

	clonedDesired := proto.Clone(desired).(*pb.Metadata)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *ComputeProjectMetadataAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Metadata) error {
	mapCtx := &direct.MapContext{}
	status := ComputeProjectMetadataStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
