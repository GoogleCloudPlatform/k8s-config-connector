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

package networksecurityinterceptdeployment

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gcp "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityInterceptDeploymentGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.InterceptClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewInterceptClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityInterceptDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.NetworkSecurityInterceptDeploymentIdentity)

	mapCtx := &direct.MapContext{}
	copied := obj.DeepCopy()
	desired := NetworkSecurityInterceptDeploymentSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type Adapter struct {
	id        *krm.NetworkSecurityInterceptDeploymentIdentity
	gcpClient *gcp.InterceptClient
	desired   *pb.InterceptDeployment
	actual    *pb.InterceptDeployment
}

var _ directbase.Adapter = &Adapter{}

// Find retrieves the GCP resource.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InterceptDeployment", "name", a.id)

	req := &pb.GetInterceptDeploymentRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetInterceptDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InterceptDeployment %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

// Create creates the resource in GCP.
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InterceptDeployment", "name", a.id)

	req := &pb.CreateInterceptDeploymentRequest{
		Parent:                fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		InterceptDeploymentId: a.id.InterceptDeployment,
		InterceptDeployment:   a.desired,
	}
	op, err := a.gcpClient.CreateInterceptDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InterceptDeployment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of InterceptDeployment %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created InterceptDeployment", "name", a.id)

	status := &krm.NetworkSecurityInterceptDeploymentStatus{}
	if err := a.updateStatus(ctx, status, created); err != nil {
		return err
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InterceptDeployment", "name", a.id)

	updateReq := proto.Clone(a.desired).(*pb.InterceptDeployment)
	updateReq.Name = a.actual.Name

	paths, err := common.CompareProtoMessage(updateReq, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.NetworkSecurityInterceptDeploymentStatus{}
		if err := a.updateStatus(ctx, status, a.actual); err != nil {
			return err
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateInterceptDeploymentRequest{
		InterceptDeployment: updateReq,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		},
	}
	op, err := a.gcpClient.UpdateInterceptDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("updating InterceptDeployment %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("InterceptDeployment %s waiting update: %w", a.id, err)
	}
	log.Info("successfully updated InterceptDeployment", "name", a.id)

	status := &krm.NetworkSecurityInterceptDeploymentStatus{}
	if err := a.updateStatus(ctx, status, updated); err != nil {
		return err
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Delete deletes the resource in GCP.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InterceptDeployment", "name", a.id)

	req := &pb.DeleteInterceptDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInterceptDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting InterceptDeployment %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("waiting for deletion of InterceptDeployment %q: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted InterceptDeployment", "name", a.id)
	return true, nil
}

// Export generates the Config Connector resource from the GCP object.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityInterceptDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkSecurityInterceptDeploymentSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Re-inject project/location back to spec because mapper doesn't know about it
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, status *krm.NetworkSecurityInterceptDeploymentStatus, updated *pb.InterceptDeployment) error {
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkSecurityInterceptDeploymentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return nil
}
