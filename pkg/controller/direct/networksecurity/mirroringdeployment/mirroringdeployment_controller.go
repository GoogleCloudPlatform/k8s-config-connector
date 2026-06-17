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

package mirroringdeployment

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityMirroringDeploymentGVK, NewMirroringDeploymentModel)
}

func NewMirroringDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &mirroringDeploymentModel{config: *config}, nil
}

var _ directbase.Model = &mirroringDeploymentModel{}

type mirroringDeploymentModel struct {
	config config.ControllerConfig
}

func (m *mirroringDeploymentModel) client(ctx context.Context) (*gcp.MirroringClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewMirroringRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurity client: %w", err)
	}
	return gcpClient, err
}

func (m *mirroringDeploymentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityMirroringDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.NetworkSecurityMirroringDeploymentIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &mirroringDeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *mirroringDeploymentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type mirroringDeploymentAdapter struct {
	id        *krm.NetworkSecurityMirroringDeploymentIdentity
	gcpClient *gcp.MirroringClient
	desired   *krm.NetworkSecurityMirroringDeployment
	actual    *pb.MirroringDeployment
}

var _ directbase.Adapter = &mirroringDeploymentAdapter{}

func (a *mirroringDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkSecurityMirroringDeployment", "name", a.id)

	req := &pb.GetMirroringDeploymentRequest{Name: a.id.String()}
	obj, err := a.gcpClient.GetMirroringDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkSecurityMirroringDeployment %q: %w", a.id, err)
	}

	a.actual = obj
	return true, nil
}

func (a *mirroringDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkSecurityMirroringDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkSecurityMirroringDeploymentSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateMirroringDeploymentRequest{
		Parent:                fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		MirroringDeploymentId: a.id.Mirroring_deployment,
		MirroringDeployment:   desired,
	}
	op, err := a.gcpClient.CreateMirroringDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkSecurityMirroringDeployment %q: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for NetworkSecurityMirroringDeployment %q creation: %w", a.id, err)
	}

	log.V(2).Info("created NetworkSecurityMirroringDeployment", "name", a.id)

	status := &krm.NetworkSecurityMirroringDeploymentStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *mirroringDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityMirroringDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkSecurityMirroringDeploymentSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
		if !(len(a.desired.Spec.Labels) == 0 && len(a.actual.Labels) == 0) {
			updateMask.Paths = append(updateMask.Paths, "labels")
		}
	}

	desiredDescription := ""
	if a.desired.Spec.Description != nil {
		desiredDescription = *a.desired.Spec.Description
	}
	if desiredDescription != a.actual.Description {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for _, path := range updateMask.Paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateMirroringDeploymentRequest{
		UpdateMask:          updateMask,
		MirroringDeployment: desired,
	}
	req.MirroringDeployment.Name = a.id.String()

	op, err := a.gcpClient.UpdateMirroringDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("updating NetworkSecurityMirroringDeployment %q: %w", a.id, err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for NetworkSecurityMirroringDeployment %q update: %w", a.id, err)
	}

	log.V(2).Info("updated NetworkSecurityMirroringDeployment", "name", a.id)

	status := &krm.NetworkSecurityMirroringDeploymentStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.GetName())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *mirroringDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("NetworkSecurityMirroringDeployment %q not found", a.id)
	}

	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityMirroringDeploymentSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting NetworkSecurityMirroringDeployment spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Mirroring_deployment)
	u.SetGroupVersionKind(krm.NetworkSecurityMirroringDeploymentGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *mirroringDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkSecurityMirroringDeployment", "name", a.id)

	req := &pb.DeleteMirroringDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMirroringDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting NetworkSecurityMirroringDeployment %q: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for NetworkSecurityMirroringDeployment %q deletion: %w", a.id, err)
	}

	return true, nil
}
