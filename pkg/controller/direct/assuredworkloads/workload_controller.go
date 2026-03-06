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
// proto.service: google.cloud.assuredworkloads.v1.AssuredWorkloadsService
// proto.message: google.cloud.assuredworkloads.v1.Workload
// crd.type: AssuredWorkloadsWorkload
// crd.version: v1alpha1

package assuredworkloads

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/assuredworkloads/apiv1"
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.AssuredWorkloadsWorkloadGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AssuredWorkloadsWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.ProvisionedResourcesParentRef != nil {
		folder, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.ProvisionedResourcesParentRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.ProvisionedResourcesParentRef.External = "folders/" + folder.FolderID
	}

	id, err := krm.NewWorkloadIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newAssuredWorkloadsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &adapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.Client
	id        *krm.WorkloadIdentity
	desired   *krm.AssuredWorkloadsWorkload
	actual    *pb.Workload
	reader    client.Reader
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AssuredWorkloadsWorkload", "name", a.id)

	req := &pb.GetWorkloadRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AssuredWorkloadsWorkload %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AssuredWorkloadsWorkload", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateWorkloadRequest{
		Parent:     a.id.Parent().String(),
		Workload:   resource,
		ExternalId: a.id.WorkloadID,
	}
	op, err := a.gcpClient.CreateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AssuredWorkloadsWorkload %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AssuredWorkloadsWorkload %s creation: %w", a.id, err)
	}

	log.V(2).Info("successfully created AssuredWorkloadsWorkload", "name", a.id)

	status := &krm.AssuredWorkloadsWorkloadStatus{}
	status.ObservedState = AssuredWorkloadsWorkloadObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AssuredWorkloadsWorkload", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if direct.ValueOf(desired.Spec.DisplayName) != a.actual.DisplayName {
		paths = append(paths, "display_name")
	}

	desiredLabels := desired.Spec.Labels
	if desiredLabels == nil {
		desiredLabels = make(map[string]string)
	}
	actualLabels := a.actual.Labels
	if actualLabels == nil {
		actualLabels = make(map[string]string)
	}

	if !reflect.DeepEqual(desiredLabels, actualLabels) {
		paths = append(paths, "labels")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	req := &pb.UpdateWorkloadRequest{
		Workload:   resource,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}
	resource.Name = a.id.String()

	updated, err := a.gcpClient.UpdateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AssuredWorkloadsWorkload", "name", a.id)

	status := &krm.AssuredWorkloadsWorkloadStatus{}
	status.ObservedState = AssuredWorkloadsWorkloadObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO: Implement Export
	return nil, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AssuredWorkloadsWorkload", "name", a.id)

	req := &pb.DeleteWorkloadRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AssuredWorkloadsWorkload", "name", a.id)

	return true, nil
}
