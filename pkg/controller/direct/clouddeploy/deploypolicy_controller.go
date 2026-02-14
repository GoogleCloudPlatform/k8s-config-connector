// Copyright 2025 Google LLC
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

package clouddeploy

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/deploy/apiv1"
	clouddeploypb "cloud.google.com/go/deploy/apiv1/deploypb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudDeployDeployPolicyGVK, NewDeployPolicyModel)
}

func NewDeployPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDeployPolicy{config: *config}, nil
}

var _ directbase.Model = &modelDeployPolicy{}

type modelDeployPolicy struct {
	config config.ControllerConfig
}

func (m *modelDeployPolicy) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DeployPolicy client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDeployPolicy) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudDeployDeployPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewDeployPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get clouddeploy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeployPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDeployPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DeployPolicyAdapter struct {
	id        *krm.DeployPolicyIdentity
	gcpClient *gcp.CloudDeployClient
	desired   *krm.CloudDeployDeployPolicy
	actual    *clouddeploypb.DeployPolicy
}

var _ directbase.Adapter = &DeployPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DeployPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeployPolicy", "name", a.id)

	req := &clouddeploypb.GetDeployPolicyRequest{Name: a.id.String()}
	deploypolicypb, err := a.gcpClient.GetDeployPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeployPolicy %q: %w", a.id, err)
	}

	a.actual = deploypolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeployPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeployPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeployPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &clouddeploypb.CreateDeployPolicyRequest{
		Parent:         a.id.Parent().String(),
		DeployPolicy:   resource,
		DeployPolicyId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateDeployPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeployPolicy %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeployPolicy %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeployPolicy", "name", a.id)

	status := &krm.DeployPolicyStatus{}
	status.ObservedState = DeployPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeployPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeployPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := DeployPolicySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	req := &clouddeploypb.UpdateDeployPolicyRequest{
		UpdateMask:   updateMask,
		DeployPolicy: desiredPb,
	}
	op, err := a.gcpClient.UpdateDeployPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DeployPolicy %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeployPolicy %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated DeployPolicy", "name", a.id)

	status := &krm.DeployPolicyStatus{}
	status.ObservedState = DeployPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DeployPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployDeployPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DeployPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.CloudDeployDeployPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DeployPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeployPolicy", "name", a.id)

	req := &clouddeploypb.DeleteDeployPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDeployPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DeployPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DeployPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DeployPolicy", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete DeployPolicy %s: %w", a.id, err)
	}
	return true, nil
}
