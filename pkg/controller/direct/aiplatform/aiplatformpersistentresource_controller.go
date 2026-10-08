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

package aiplatform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
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
	registry.RegisterModel(krm.AIPlatformPersistentResourceGVK, NewAIPlatformPersistentResourceModel)
}

func NewAIPlatformPersistentResourceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &aiplatformPersistentResourceModel{config: config}, nil
}

var _ directbase.Model = &aiplatformPersistentResourceModel{}

type aiplatformPersistentResourceModel struct {
	config *config.ControllerConfig
}

func (m *aiplatformPersistentResourceModel) client(ctx context.Context, location string) (*gcp.PersistentResourceClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewPersistentResourceClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PersistentResourceClient client: %w", err)
	}
	return gcpClient, nil
}

func (m *aiplatformPersistentResourceModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.AIPlatformPersistentResource{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(reader.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	typedID, ok := id.(*krm.AIPlatformPersistentResourceIdentity)
	if !ok {
		return nil, fmt.Errorf("expected AIPlatformPersistentResourceIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx, typedID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := AIPlatformPersistentResourceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &AIPlatformPersistentResourceAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *aiplatformPersistentResourceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AIPlatformPersistentResourceIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &AIPlatformPersistentResourceAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AIPlatformPersistentResourceAdapter struct {
	id        *krm.AIPlatformPersistentResourceIdentity
	gcpClient *gcp.PersistentResourceClient
	desiredpb *pb.PersistentResource
	desired   *krm.AIPlatformPersistentResource
	actual    *pb.PersistentResource
}

var _ directbase.Adapter = &AIPlatformPersistentResourceAdapter{}

func (a *AIPlatformPersistentResourceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AIPlatformPersistentResource", "name", a.id.String())

	req := &pb.GetPersistentResourceRequest{
		Name: a.id.String(),
	}

	persistentResourcepb, err := a.gcpClient.GetPersistentResource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIPlatformPersistentResource %q: %w", a.id.String(), err)
	}

	a.actual = persistentResourcepb

	mapCtx := &direct.MapContext{}
	observedState := AIPlatformPersistentResourceObservedState_FromProto(mapCtx, persistentResourcepb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}

	if a.desired != nil {
		a.desired.Status.ObservedState = observedState
		a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return true, nil
}

func (a *AIPlatformPersistentResourceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIPlatformPersistentResource", "name", a.id.String())

	// Ensure Name is empty/set correctly according to GCP REST API patterns
	a.desiredpb.Name = ""

	req := &pb.CreatePersistentResourceRequest{
		Parent:               a.id.ParentString(),
		PersistentResourceId: a.id.PersistentResource,
		PersistentResource:   a.desiredpb,
	}

	op, err := a.gcpClient.CreatePersistentResource(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AIPlatformPersistentResource %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully started creation of AIPlatformPersistentResource", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformPersistentResource %q creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully completed creation of AIPlatformPersistentResource", "name", a.id.String())

	// Fetch fully-populated resource immediately after LRO success
	getReq := &pb.GetPersistentResourceRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetPersistentResource(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created AIPlatformPersistentResource %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *AIPlatformPersistentResourceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIPlatformPersistentResource", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, _, err := comparePersistentResource(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("AIPlatformPersistentResource is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *AIPlatformPersistentResourceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformPersistentResource{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformPersistentResourceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.PersistentResource

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.PersistentResource)
	u.SetGroupVersionKind(krm.AIPlatformPersistentResourceGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *AIPlatformPersistentResourceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.PersistentResource) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformPersistentResourceStatus{}
	status.ObservedState = AIPlatformPersistentResourceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AIPlatformPersistentResourceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIPlatformPersistentResource", "name", a.id.String())

	req := &pb.DeletePersistentResourceRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeletePersistentResource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AIPlatformPersistentResource %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of AIPlatformPersistentResource %q: %w", a.id.String(), err)
	}

	return true, nil
}

func comparePersistentResource(ctx context.Context, actual, desired *pb.PersistentResource) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AIPlatformPersistentResourceSpec_FromProto, AIPlatformPersistentResourceSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.PersistentResource)
	// If displayName is not specified in desired, but populated in actual, align them
	if clonedDesired.DisplayName == "" && maskedActual.DisplayName != "" {
		clonedDesired.DisplayName = maskedActual.DisplayName
	}
	// Same for resourcePool IDs and diskSpecs
	for i, pool := range clonedDesired.ResourcePools {
		if i < len(maskedActual.ResourcePools) {
			if pool.Id == "" && maskedActual.ResourcePools[i].Id != "" {
				pool.Id = maskedActual.ResourcePools[i].Id
			}
			if pool.DiskSpec == nil && maskedActual.ResourcePools[i].DiskSpec != nil {
				pool.DiskSpec = proto.Clone(maskedActual.ResourcePools[i].DiskSpec).(*pb.DiskSpec)
			} else if pool.DiskSpec != nil && maskedActual.ResourcePools[i].DiskSpec != nil {
				if pool.DiskSpec.BootDiskSizeGb == 0 && maskedActual.ResourcePools[i].DiskSpec.BootDiskSizeGb != 0 {
					pool.DiskSpec.BootDiskSizeGb = maskedActual.ResourcePools[i].DiskSpec.BootDiskSizeGb
				}
				if pool.DiskSpec.BootDiskType == "" && maskedActual.ResourcePools[i].DiskSpec.BootDiskType != "" {
					pool.DiskSpec.BootDiskType = maskedActual.ResourcePools[i].DiskSpec.BootDiskType
				}
			}
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
