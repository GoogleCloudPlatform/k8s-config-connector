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

package compute

import (
	"context"
	"fmt"

	api "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddSubnetworkController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddSubnetworkController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.ComputeSubnetworkGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &networkModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m)
}

type subnetworkModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &subnetworkModel{}

type subnetworkAdapter struct {
	projectID string

	desired *krm.ComputeSubnetwork
	actual  *pb.Subnetwork

	*gcpClient
	subnetworks *api.SubnetworksClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &subnetworkAdapter{}

// AdapterForObject implements the Model interface.
func (m *subnetworkModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	subnetworks, err := m.newSubnetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.ComputeSubnetwork{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	desired := obj.Spec.DeepCopy()

	projectID := obj.Annotations["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}
	desired.ResourceID = &resourceID

	if desired.Region == "" {
		return nil, fmt.Errorf("unable to determine region")
	}

	networkURL := desired.NetworkRef.External
	// TODO: Refs
	if networkURL == "" {
		return nil, fmt.Errorf("unable to determine networkURL")
	}

	return &subnetworkAdapter{
		projectID:   projectID,
		desired:     obj,
		gcpClient:   m.gcpClient,
		subnetworks: subnetworks,
	}, nil
}

// Find implements the Adapter interface.
func (a *subnetworkAdapter) Find(ctx context.Context) (bool, error) {
	if ValueOf(a.desired.Spec.ResourceID) == "" {
		return false, nil
	}

	req := &pb.GetSubnetworkRequest{
		Project:    a.projectID,
		Region:     a.desired.Spec.Region,
		Subnetwork: *a.desired.Spec.ResourceID,
	}
	subnetwork, err := a.subnetworks.Get(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("subnetwork was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	a.actual = subnetwork

	return true, nil
}

// Delete implements the Adapter interface.
func (a *subnetworkAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteSubnetworkRequest{
		Project:    a.projectID,
		Region:     a.desired.Spec.Region,
		Subnetwork: *a.desired.Spec.ResourceID,
	}

	op, err := a.subnetworks.Delete(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting subnetwork: %w", err)
	}

	completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
	if err != nil {
		return false, fmt.Errorf("waiting for subnetwork deletion: %w", err)
	}
	// TODO: Move this check to wait?
	if completed.GetStatus() != pb.Operation_DONE {
		return false, fmt.Errorf("subnetwork deletion failed: %q", completed.GetStatus())
	}

	return true, nil
}

func mapComputeSubnetworkSpecToCloud(in *krm.ComputeSubnetworkSpec, out *pb.Subnetwork) error {
	subnetwork := &pb.Subnetwork{
		Description: in.Description,
		// EnableFlowLogs:          in.EnableFwLogs,
		IpCidrRange:             &in.IpCidrRange,
		Ipv6AccessType:          in.Ipv6AccessType,
		Name:                    in.ResourceID,
		Network:                 &in.NetworkRef.External,
		PrivateIpGoogleAccess:   in.PrivateIpGoogleAccess,
		PrivateIpv6GoogleAccess: in.PrivateIpv6GoogleAccess,
		Purpose:                 in.Purpose,
		Region:                  &in.Region,
		Role:                    in.Role,
		StackType:               in.StackType,
	}

	for _, secondaryRange := range in.SecondaryIpRange {
		in := secondaryRange
		out := &pb.SubnetworkSecondaryRange{
			RangeName:   &in.RangeName,
			IpCidrRange: &in.IpCidrRange,
		}
		subnetwork.SecondaryIpRanges = append(subnetwork.SecondaryIpRanges, out)
	}
	if in.LogConfig != nil {
		subnetwork.LogConfig = &pb.SubnetworkLogConfig{
			AggregationInterval: in.LogConfig.AggregationInterval,
			FilterExpr:          in.LogConfig.FilterExpr,
			Metadata:            in.LogConfig.Metadata,
			MetadataFields:      in.LogConfig.MetadataFields,
		}

		// TODO: Should be explicit?
		subnetwork.LogConfig.Enable = PtrTo(true)

		// TODO: Should be float32
		if in.LogConfig.FlowSampling != nil {
			flowSampling := float32(*in.LogConfig.FlowSampling)
			subnetwork.LogConfig.FlowSampling = &flowSampling
		}
	}

	return nil
}

// Create implements the Adapter interface.
func (a *subnetworkAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	subnetwork := &pb.Subnetwork{}
	if err := mapComputeSubnetworkSpecToCloud(&a.desired.Spec, subnetwork); err != nil {
		return err
	}

	req := &pb.InsertSubnetworkRequest{
		SubnetworkResource: subnetwork,
		Project:            a.projectID,
		Region:             a.desired.Spec.Region,
	}
	op, err := a.subnetworks.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating subnetwork: %w", err)
	}
	// TODO: Return created object status

	completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
	if err != nil {
		return fmt.Errorf("waiting for subnetwork creation: %w", err)
	}
	if completed.GetStatus() != pb.Operation_DONE {
		return fmt.Errorf("subnetwork creation failed: %q", completed.GetStatus())
	}

	// Can we get this from the operation?
	created, err := a.subnetworks.Get(ctx, &pb.GetSubnetworkRequest{Project: a.projectID, Region: a.desired.Spec.Region, Subnetwork: *a.desired.Spec.ResourceID})
	if err != nil {
		return fmt.Errorf("getting subnetwork after creation: %w", err)
	}
	log.V(2).Info("created subnetwork", "subnetwork", created)

	status := &krm.ComputeSubnetworkStatus{}
	if err := subnetworkStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func subnetworkStatusToKRM(in *pb.Subnetwork, out *krm.ComputeSubnetworkStatus) error {
	out.CreationTimestamp = in.CreationTimestamp
	out.ExternalIpv6Prefix = in.ExternalIpv6Prefix
	out.GatewayAddress = in.GatewayAddress
	out.InternalIpv6Prefix = in.InternalIpv6Prefix
	out.Ipv6CidrRange = in.Ipv6CidrRange
	out.SelfLink = (in.SelfLink)

	return nil
}

// Update implements the Adapter interface.
func (a *subnetworkAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Skip updates at the higher level if no changes?

	subnetwork := &pb.Subnetwork{}
	if err := mapComputeSubnetworkSpecToCloud(&a.desired.Spec, subnetwork); err != nil {
		return err
	}

	// TODO: Where/how do we want to enforce immutability?

	// TODO: Compute update mask, only with spec
	// updateMask := &fieldmaskpb.FieldMask{}
	shouldUpdate := true // len(updateMask.Paths) != 0
	update := subnetwork

	if shouldUpdate {
		req := &pb.PatchSubnetworkRequest{
			Region:             a.desired.Spec.Region,
			Subnetwork:         *a.desired.Spec.ResourceID,
			Project:            a.projectID,
			SubnetworkResource: update,
		}
		op, err := a.subnetworks.Patch(ctx, req)
		if err != nil {
			return err
		}

		completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
		if err != nil {
			return fmt.Errorf("waiting for subnetwork updte: %w", err)
		}
		// TODO: Move this check to wait?
		if completed.GetStatus() != pb.Operation_DONE {
			return fmt.Errorf("subnetwork update failed: %q", completed.GetStatus())
		}

	}

	// TODO: Return updated object status
	return nil
}
