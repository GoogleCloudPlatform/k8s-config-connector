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
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ArtifactRegistryVPCSCConfigGVK, NewVPCSCConfigModel)
}

func NewVPCSCConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelVPCSCConfig{config: *config}, nil
}

var _ directbase.Model = &modelVPCSCConfig{}

type modelVPCSCConfig struct {
	config config.ControllerConfig
}

func (m *modelVPCSCConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building client: %w", err)
	}
	return gcpClient, err
}

func (m *modelVPCSCConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ArtifactRegistryVPCSCConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedId := id.(*krm.ArtifactRegistryVPCSCConfigIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VPCSCConfigAdapter{
		id:        typedId,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelVPCSCConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type VPCSCConfigAdapter struct {
	id        *krm.ArtifactRegistryVPCSCConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.ArtifactRegistryVPCSCConfig
}

var _ directbase.Adapter = &VPCSCConfigAdapter{}

func (a *VPCSCConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VPCSCConfig", "name", a.id.String())

	req := &pb.GetVPCSCConfigRequest{Name: a.id.String()}
	vpCscConfig, err := a.gcpClient.GetVPCSCConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VPCSCConfig %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	observedState := ArtifactRegistryVPCSCConfigObservedState_FromProto(mapCtx, vpCscConfig)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}

	a.desired.Status.ObservedState = observedState
	return true, nil
}

func (a *VPCSCConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VPCSCConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredProto := ArtifactRegistryVPCSCConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredProto.Name = a.id.String()

	req := &pb.UpdateVPCSCConfigRequest{
		VpcscConfig: desiredProto,
	}

	// For Create, we update the existing singleton resource (since VPCSCConfig is 1:1 with project/location).
	vpCscConfig, err := a.gcpClient.UpdateVPCSCConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VPCSCConfig %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VPCSCConfig", "name", a.id.String())

	status := &krm.ArtifactRegistryVPCSCConfigStatus{}
	status.ObservedState = ArtifactRegistryVPCSCConfigObservedState_FromProto(mapCtx, vpCscConfig)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *VPCSCConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VPCSCConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredProto := ArtifactRegistryVPCSCConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredProto.Name = a.id.String()

	req := &pb.UpdateVPCSCConfigRequest{
		VpcscConfig: desiredProto,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"vpcsc_policy"}},
	}

	vpCscConfig, err := a.gcpClient.UpdateVPCSCConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VPCSCConfig %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated VPCSCConfig", "name", a.id.String())

	status := &krm.ArtifactRegistryVPCSCConfigStatus{}
	status.ObservedState = ArtifactRegistryVPCSCConfigObservedState_FromProto(mapCtx, vpCscConfig)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *VPCSCConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("VPCSCConfigAdapter not initialized")
	}
	mapCtx := &direct.MapContext{}
	spec := ArtifactRegistryVPCSCConfigSpec_FromProto(mapCtx, ArtifactRegistryVPCSCConfigSpec_ToProto(mapCtx, &a.desired.Spec))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Clear out refs
	spec.ProjectRef = nil

	obj := &krm.ArtifactRegistryVPCSCConfig{}
	obj.Spec = *spec
	obj.Spec.Location = &a.id.Location
	obj.ObjectMeta.Name = "vpcscconfig"

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: u}, nil
}

func (a *VPCSCConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VPCSCConfig", "name", a.id.String())

	// Since VPCSCConfig is a singleton per project/location, "deleting" it means reverting to default (VPCSC_POLICY_UNSPECIFIED)
	req := &pb.UpdateVPCSCConfigRequest{
		VpcscConfig: &pb.VPCSCConfig{
			Name:        a.id.String(),
			VpcscPolicy: pb.VPCSCConfig_VPCSC_POLICY_UNSPECIFIED,
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"vpcsc_policy"}},
	}

	_, err := a.gcpClient.UpdateVPCSCConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("reverting VPCSCConfig to default %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully reverted VPCSCConfig to default", "name", a.id.String())
	return true, nil
}
