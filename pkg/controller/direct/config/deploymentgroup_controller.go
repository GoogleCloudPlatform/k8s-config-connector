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

package config

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/config/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/config/apiv1"
	pb "cloud.google.com/go/config/apiv1/configpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ConfigDeploymentGroupGVK, NewDeploymentGroupModel)
}

func NewDeploymentGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDeploymentGroup{config: *config}, nil
}

var _ directbase.Model = &modelDeploymentGroup{}

type modelDeploymentGroup struct {
	config config.ControllerConfig
}

func (m *modelDeploymentGroup) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentGroup client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelDeploymentGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeploymentGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	configID, ok := id.(*krm.ConfigDeploymentGroupIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	// Get config GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeploymentGroupAdapter{
		id:        configID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDeploymentGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DeploymentGroupAdapter struct {
	id        *krm.ConfigDeploymentGroupIdentity
	gcpClient *gcp.Client
	desired   *krm.ConfigDeploymentGroup
	actual    *pb.DeploymentGroup
}

var _ directbase.Adapter = &DeploymentGroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DeploymentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeploymentGroup", "name", a.id)

	req := &pb.GetDeploymentGroupRequest{Name: a.id.String()}
	deploymentgrouppb, err := a.gcpClient.GetDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeploymentGroup %q: %w", a.id, err)
	}

	a.actual = deploymentgrouppb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeploymentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeploymentGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ConfigDeploymentGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateDeploymentGroupRequest{
		Parent:            parent,
		DeploymentGroupId: a.id.DeploymentGroup,
		DeploymentGroup:   resource,
	}
	op, err := a.gcpClient.CreateDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeploymentGroup %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeploymentGroup %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeploymentGroup", "name", a.id)

	status := &krm.ConfigDeploymentGroupStatus{}
	status.ObservedState = ConfigDeploymentGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeploymentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeploymentGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := ConfigDeploymentGroupSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		desiredPb.Name = a.id.String()
		req := &pb.UpdateDeploymentGroupRequest{
			UpdateMask:      updateMask,
			DeploymentGroup: desiredPb,
		}
		op, err := a.gcpClient.UpdateDeploymentGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DeploymentGroup %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DeploymentGroup %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DeploymentGroup", "name", a.id)
	}

	status := &krm.ConfigDeploymentGroupStatus{}
	status.ObservedState = ConfigDeploymentGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DeploymentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeploymentGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeploymentGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.DeploymentGroup)
	u.SetGroupVersionKind(krm.ConfigDeploymentGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DeploymentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeploymentGroup", "name", a.id)

	req := &pb.DeleteDeploymentGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DeploymentGroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DeploymentGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DeploymentGroup", "name", a.id)

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete DeploymentGroup %s: %w", a.id, err)
	}
	return true, nil
}
