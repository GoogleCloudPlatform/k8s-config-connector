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

// +tool:controller
// proto.service: google.cloud.compute.v1.Interconnects
// proto.message: google.cloud.compute.v1.Interconnect
// crd.type: ComputeInterconnect
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ComputeInterconnectGVK, NewInterconnectModel)
}

func NewInterconnectModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &interconnectModel{config: *config}, nil
}

var _ directbase.Model = &interconnectModel{}

type interconnectModel struct {
	config config.ControllerConfig
}

func (m *interconnectModel) Client(ctx context.Context, projectID string) (*compute.InterconnectsClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := compute.NewInterconnectsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building compute interconnect client: %w", err)
	}

	return gcpClient, err
}

func (m *interconnectModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeInterconnect{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInterconnectIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeInterconnectSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.Client(ctx, id.ProjectID)
	if err != nil {
		return nil, err
	}

	return &interconnectAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *interconnectModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type interconnectAdapter struct {
	gcpClient *compute.InterconnectsClient
	id        *v1alpha1.InterconnectIdentity
	desired   *computepb.Interconnect
	actual    *computepb.Interconnect
}

var _ directbase.Adapter = &interconnectAdapter{}

func (a *interconnectAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting compute interconnect", "name", a.id)

	req := &computepb.GetInterconnectRequest{
		Project:      a.id.ProjectID,
		Interconnect: a.id.Interconnect,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting compute interconnect %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *interconnectAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating compute interconnect", "name", a.id)

	req := &computepb.InsertInterconnectRequest{
		Project:              a.id.ProjectID,
		InterconnectResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating compute interconnect %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute interconnect %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute interconnect in gcp", "name", a.id)

	status := &krm.ComputeInterconnectStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeInterconnectObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// ComputeInterconnect support update labels.
func (a *interconnectAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating compute interconnect", "name", a.id)

	desiredpb := proto.Clone(a.desired).(*computepb.Interconnect)
	paths, err := common.CompareProtoMessage(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	req := &computepb.PatchInterconnectRequest{
		Project:              a.id.ProjectID,
		Interconnect:         a.id.Interconnect,
		InterconnectResource: desiredpb,
		UpdateMask:           &fieldmaskpb.FieldMask{Paths: paths},
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating compute interconnect %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated compute interconnect", "name", a.id.String())

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute interconnect %s waiting for update: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.ComputeInterconnectStatus{}
	status.ObservedState = ComputeInterconnectObservedState_FromProto(mapCtx, updated)
	if err := mapCtx.Err(); err != nil {
		return err
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *interconnectAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting compute interconnect", "name", a.id)

	req := &computepb.DeleteInterconnectRequest{
		Project:      a.id.ProjectID,
		Interconnect: a.id.Interconnect,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute interconnect %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute interconnect", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute interconnect %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *interconnectAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeInterconnect{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeInterconnectSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Interconnect)
	u.SetGroupVersionKind(krm.ComputeInterconnectGVK)

	u.Object = uObj
	return u, nil
}
