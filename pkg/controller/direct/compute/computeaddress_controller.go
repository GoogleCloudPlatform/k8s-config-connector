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
	"reflect"

	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeAddressGVK, NewAddressModel)
}

func NewAddressModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &addressModel{config: config}, nil
}

type addressModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &addressModel{}

type addressAdapter struct {
	id                  *krm.AddressIdentity
	addressClient       *gcp.AddressesClient
	globalAddressClient *gcp.GlobalAddressesClient
	desired             *krm.ComputeAddress
	actual              *computepb.Address
	reader              client.Reader
}

var _ directbase.Adapter = &addressAdapter{}

func (m *addressModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeAddress{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AddressIdentity)

	addressAdapter := &addressAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get region
	region := id.ParentID.Location

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if region == "global" {
		globalAddressClient, err := gcpClient.newGlobalAddressClient(ctx)
		if err != nil {
			return nil, err
		}
		addressAdapter.globalAddressClient = globalAddressClient
	} else {
		addressClient, err := gcpClient.addressClient(ctx)
		if err != nil {
			return nil, err
		}
		addressAdapter.addressClient = addressClient
	}
	return addressAdapter, nil
}

func (m *addressModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *addressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeAddress", "name", a.id)

	var err error
	address := &computepb.Address{}
	address, err = a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeAddress %q: %w", a.id, err)
	}
	a.actual = address
	return true, nil
}

func (a *addressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeAddress", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	// Resolve refs if any
	if desired.Spec.NetworkRef != nil {
		if err := desired.Spec.NetworkRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return fmt.Errorf("resolving NetworkRef: %w", err)
		}
	}
	// SubnetworkRef uses standard refs, resolve manually if needed or if Normalize exists.
	// refs.ComputeSubnetworkRef DOES NOT have Normalize method generated usually unless it is in apis/refs.
	// We might need to handle it. But `computerefs` has `ResolveComputeSubnetwork`.
	if desired.Spec.SubnetworkRef != nil {
		normalized, err := refs.ResolveComputeSubnetwork(ctx, a.reader, a.desired, desired.Spec.SubnetworkRef)
		if err != nil {
			return fmt.Errorf("resolving SubnetworkRef: %w", err)
		}
		desired.Spec.SubnetworkRef = normalized
	}

	sanitizedLabels := label.NewGCPLabelsFromK8sLabels(desired.Labels)

	address := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Manual mapping for IpCollection as it is missing in the generator's proto definition
	if desired.Spec.IpCollection != nil {
		address.IpCollection = desired.Spec.IpCollection
	}
	if desired.Spec.IpCollectionRef != nil {
		if err := desired.Spec.IpCollectionRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return fmt.Errorf("resolving IpCollectionRef: %w", err)
		}
		address.IpCollection = direct.LazyPtr(desired.Spec.IpCollectionRef.External)
	}
	address.Name = direct.LazyPtr(a.id.ResourceID)
	address.Labels = sanitizedLabels

	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.InsertGlobalAddressRequest{
			AddressResource: address,
			Project:         a.id.ParentID.ProjectID,
		}
		op, err = a.globalAddressClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertAddressRequest{
			AddressResource: address,
			Region:          a.id.ParentID.Location,
			Project:         a.id.ParentID.ProjectID,
		}
		op, err = a.addressClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputeAddress %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeAddress %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeAddress", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeAddress %q: %w", a.id, err)
	}

	// Set status
	status := &krm.ComputeAddressStatus{
		LabelFingerprint:  created.LabelFingerprint,
		CreationTimestamp: created.CreationTimestamp,
		SelfLink:          created.SelfLink,
		Users:             created.Users,
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(u, status)
}

func (a *addressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	var err error

	if a.id.ResourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeAddress", "name", a.id.ResourceID)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	sanitizedLabels := label.NewGCPLabelsFromK8sLabels(desired.Labels)
	address := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Manual mapping for IpCollection as it is missing in the generator's proto definition
	if desired.Spec.IpCollection != nil {
		address.IpCollection = desired.Spec.IpCollection
	}
	address.Name = direct.LazyPtr(a.id.ResourceID)
	address.Labels = sanitizedLabels

	// Resolve refs (needed for comparison/mapping?) - maybe not for limited updates, but good practice.
	if desired.Spec.NetworkRef != nil {
		if err := desired.Spec.NetworkRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return fmt.Errorf("resolving NetworkRef: %w", err)
		}
	}
	if desired.Spec.SubnetworkRef != nil {
		normalized, err := refs.ResolveComputeSubnetwork(ctx, a.reader, a.desired, desired.Spec.SubnetworkRef)
		if err != nil {
			return fmt.Errorf("resolving SubnetworkRef: %w", err)
		}
		desired.Spec.SubnetworkRef = normalized
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// Address update is limited. Typically Labels and Description?
	// Check immutable fields? For now, we assume immutable unless we implement specific update logic.
	// Labels can be updated.

	if !reflect.DeepEqual(address.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, address.Labels)
		op, err := a.setLabels(ctx, a.actual.LabelFingerprint, address.Labels)
		if err != nil {
			return fmt.Errorf("updating ComputeAddress labels %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeAddress %s update labels failed: %w", a.id, err)
			}
		}
		log.V(2).Info("successfully updated ComputeAddress labels", "name", a.id)
	}

	structuredreporting.ReportDiff(ctx, report)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeAddress %q: %w", a.id.ResourceID, err)
	}

	status := &krm.ComputeAddressStatus{
		LabelFingerprint:  updated.LabelFingerprint,
		CreationTimestamp: updated.CreationTimestamp,
		SelfLink:          updated.SelfLink,
		Users:             updated.Users,
	}
	return setStatus(u, status)
}

func (a *addressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("address %q not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeAddressSpec_v1beta1_FromProto(mc, a.actual)
	// Manual mapping for IpCollection as it is missing in the generator's proto definition
	if a.actual.IpCollection != nil {
		spec.IpCollection = a.actual.IpCollection
		spec.IpCollectionRef = &v1beta1.ComputePublicDelegatedPrefixRef{
			External: *a.actual.IpCollection,
		}
	}
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting address spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ComputeAddressGVK)
	u.SetLabels(a.actual.Labels)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *addressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.ResourceID == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeAddress", "name", a.id.ResourceID)

	var err error
	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.DeleteGlobalAddressRequest{
			Address: a.id.ResourceID,
			Project: a.id.ParentID.ProjectID,
		}
		op, err = a.globalAddressClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteAddressRequest{
			Address: a.id.ResourceID,
			Region:  a.id.ParentID.Location,
			Project: a.id.ParentID.ProjectID,
		}
		op, err = a.addressClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputeAddress %s: %w", a.id.ResourceID, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeAddress %s delete failed: %w", a.id.ResourceID, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeAddress", "name", a.id.ResourceID)
	return true, nil
}

func (a *addressAdapter) get(ctx context.Context) (*computepb.Address, error) {
	if a.id.ParentID.Location == "global" {
		getReq := &computepb.GetGlobalAddressRequest{
			Address: a.id.ResourceID,
			Project: a.id.ParentID.ProjectID,
		}
		return a.globalAddressClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetAddressRequest{
			Address: a.id.ResourceID,
			Region:  a.id.ParentID.Location,
			Project: a.id.ParentID.ProjectID,
		}
		return a.addressClient.Get(ctx, getReq)
	}
}

func (a *addressAdapter) setLabels(ctx context.Context, fingerprint *string, labels map[string]string) (*gcp.Operation, error) {
	op := &gcp.Operation{}
	var err error
	if a.id.ParentID.Location == "global" {
		setLabelsReq := &computepb.SetLabelsGlobalAddressRequest{
			Resource:                       a.id.ResourceID,
			GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.ParentID.ProjectID,
		}
		op, err = a.globalAddressClient.SetLabels(ctx, setLabelsReq)
	} else {
		setLabelsReq := &computepb.SetLabelsAddressRequest{
			Resource:                       a.id.ResourceID,
			RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{LabelFingerprint: fingerprint, Labels: labels},
			Project:                        a.id.ParentID.ProjectID,
			Region:                         a.id.ParentID.Location,
		}
		op, err = a.addressClient.SetLabels(ctx, setLabelsReq)
	}
	return op, err
}
