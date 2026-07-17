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
// proto.service: google.cloud.compute.v1.Firewalls
// proto.message: google.cloud.compute.v1.Firewall
// crd.type: ComputeFirewall
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeFirewallGVK, NewFirewallModel)
}

func NewFirewallModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firewallModel{config: config}, nil
}

var _ directbase.Model = &firewallModel{}

type firewallModel struct {
	config *config.ControllerConfig
}

func (m *firewallModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeFirewall{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve references.
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	firewallsClient, err := gcpClient.newFirewallsClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeFirewallSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = direct.LazyPtr(id.(*krm.ComputeFirewallIdentity).Firewall)

	return &FirewallAdapter{
		gcpClient: firewallsClient,
		id:        id.(*krm.ComputeFirewallIdentity),
		desired:   desired,
	}, nil
}

func (m *firewallModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FirewallAdapter struct {
	gcpClient *compute.FirewallsClient
	id        *krm.ComputeFirewallIdentity
	desired   *computepb.Firewall
	actual    *computepb.Firewall
}

var _ directbase.Adapter = &FirewallAdapter{}

func (a *FirewallAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeFirewall", "name", a.id)

	req := &computepb.GetFirewallRequest{
		Project:  a.id.Project,
		Firewall: a.id.Firewall,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeFirewall %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *FirewallAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeFirewall", "name", a.id)

	req := &computepb.InsertFirewallRequest{
		Project:          a.id.Project,
		FirewallResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeFirewall %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeFirewall %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created ComputeFirewall in gcp", "name", a.id)

	created, err := a.gcpClient.Get(ctx, &computepb.GetFirewallRequest{
		Project:  a.id.Project,
		Firewall: a.id.Firewall,
	})
	if err != nil {
		return fmt.Errorf("getting ComputeFirewall after creation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *FirewallAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeFirewall", "name", a.id)

	diffs, err := compareFirewall(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeFirewall", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &computepb.PatchFirewallRequest{
		Project:          a.id.Project,
		Firewall:         a.id.Firewall,
		FirewallResource: a.desired,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeFirewall %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeFirewall %s waiting for update: %w", a.id.String(), err)
	}

	updated, err := a.gcpClient.Get(ctx, &computepb.GetFirewallRequest{
		Project:  a.id.Project,
		Firewall: a.id.Firewall,
	})
	if err != nil {
		return fmt.Errorf("getting ComputeFirewall %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *FirewallAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeFirewall{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeFirewallSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeFirewallGVK)

	return u, nil
}

func (a *FirewallAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeFirewall", "name", a.id)

	req := &computepb.DeleteFirewallRequest{
		Project:  a.id.Project,
		Firewall: a.id.Firewall,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeFirewall %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeFirewall", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeFirewall %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *FirewallAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Firewall) error {
	mapCtx := &direct.MapContext{}
	status := ComputeFirewallStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareFirewall(ctx context.Context, actual, desired *computepb.Firewall) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeFirewallSpec_v1beta1_FromProto, ComputeFirewallSpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*computepb.Firewall)

	populateDefaults := func(obj *computepb.Firewall) {
		if obj.Priority == nil {
			obj.Priority = direct.LazyPtr(int32(1000))
		}
		if obj.Direction == nil {
			obj.Direction = direct.LazyPtr("INGRESS")
		}
		if obj.Disabled == nil {
			obj.Disabled = direct.LazyPtr(false)
		}
		if obj.LogConfig == nil {
			obj.LogConfig = &computepb.FirewallLogConfig{
				Enable: direct.LazyPtr(false),
			}
		} else if obj.LogConfig.Enable == nil {
			obj.LogConfig.Enable = direct.LazyPtr(false)
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
