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
// proto.service: google.cloud.compute.v1.Disks
// proto.message: google.cloud.compute.v1.AddResourcePoliciesDiskRequest
// crd.type: ComputeDiskResourcePolicyAttachment
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeDiskResourcePolicyAttachmentGVK, NewComputeDiskResourcePolicyAttachmentModel)
}

func NewComputeDiskResourcePolicyAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeDiskResourcePolicyAttachmentModel{config: config}, nil
}

var _ directbase.Model = &computeDiskResourcePolicyAttachmentModel{}

type computeDiskResourcePolicyAttachmentModel struct {
	config *config.ControllerConfig
}

func (m *computeDiskResourcePolicyAttachmentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeDiskResourcePolicyAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	attachmentId := id.(*krm.ComputeDiskResourcePolicyAttachmentIdentity)

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	disksClient, err := gcpClient.newDisksClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeDiskResourcePolicyAttachmentAdapter{
		disksClient: disksClient,
		id:          attachmentId,
		reader:      reader,
	}, nil
}

func (m *computeDiskResourcePolicyAttachmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeDiskResourcePolicyAttachmentAdapter struct {
	disksClient *gcp.DisksClient
	id          *krm.ComputeDiskResourcePolicyAttachmentIdentity
	actual      string
	reader      client.Reader
}

var _ directbase.Adapter = &ComputeDiskResourcePolicyAttachmentAdapter{}

func (a *ComputeDiskResourcePolicyAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeDiskResourcePolicyAttachment", "id", a.id)

	disk, err := a.disksClient.Get(ctx, &pb.GetDiskRequest{
		Project: a.id.Project,
		Zone:    a.id.Zone,
		Disk:    a.id.Disk,
	})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeDisk %s: %w", a.id.Disk, err)
	}

	for _, policyUrl := range disk.ResourcePolicies {
		tokens := strings.Split(policyUrl, "/")
		if len(tokens) == 0 {
			continue
		}
		policyName := tokens[len(tokens)-1]
		if policyName == a.id.Name {
			a.actual = policyUrl
			return true, nil
		}
	}

	return false, nil
}

func (a *ComputeDiskResourcePolicyAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeDiskResourcePolicyAttachment", "id", a.id)

	region := GetRegionFromZone(a.id.Zone)
	policyUrl := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", a.id.Project, region, a.id.Name)

	req := &pb.AddResourcePoliciesDiskRequest{
		Project: a.id.Project,
		Zone:    a.id.Zone,
		Disk:    a.id.Disk,
		DisksAddResourcePoliciesRequestResource: &pb.DisksAddResourcePoliciesRequest{
			ResourcePolicies: []string{policyUrl},
		},
	}

	op, err := a.disksClient.AddResourcePolicies(ctx, req)
	if err != nil {
		return fmt.Errorf("attaching resource policy %s to disk %s: %w", a.id.Name, a.id.Disk, err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for resource policy %s to attach to disk %s: %w", a.id.Name, a.id.Disk, err)
		}
	}

	log.V(2).Info("successfully attached resource policy to disk", "id", a.id)

	status := &krm.ComputeDiskResourcePolicyAttachmentStatus{}
	status.ObservedGeneration = direct.PtrTo(createOp.GetUnstructured().GetGeneration())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ComputeDiskResourcePolicyAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeDiskResourcePolicyAttachment", "id", a.id)
	return fmt.Errorf("ComputeDiskResourcePolicyAttachment is immutable and cannot be updated")
}

func (a *ComputeDiskResourcePolicyAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeDiskResourcePolicyAttachment", "id", a.id)

	region := GetRegionFromZone(a.id.Zone)
	policyUrl := fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", a.id.Project, region, a.id.Name)

	req := &pb.RemoveResourcePoliciesDiskRequest{
		Project: a.id.Project,
		Zone:    a.id.Zone,
		Disk:    a.id.Disk,
		DisksRemoveResourcePoliciesRequestResource: &pb.DisksRemoveResourcePoliciesRequest{
			ResourcePolicies: []string{policyUrl},
		},
	}

	op, err := a.disksClient.RemoveResourcePolicies(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("removing resource policy %s from disk %s: %w", a.id.Name, a.id.Disk, err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for resource policy %s to remove from disk %s: %w", a.id.Name, a.id.Disk, err)
		}
	}

	log.V(2).Info("successfully removed resource policy from disk", "id", a.id)
	return true, nil
}

func (a *ComputeDiskResourcePolicyAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == "" {
		return nil, fmt.Errorf("Find() not called or attachment not found")
	}

	u := &unstructured.Unstructured{}

	obj := &krm.ComputeDiskResourcePolicyAttachment{}
	obj.Spec.DiskRef = computev1beta1.ComputeDiskRef{
		External: a.id.Disk,
	}
	obj.Spec.ProjectRef = refs.ProjectRef{
		External: a.id.Project,
	}
	obj.Spec.Zone = a.id.Zone
	obj.Spec.ResourceID = &a.id.Name

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ComputeDiskResourcePolicyAttachmentGVK)
	return u, nil
}

// GetRegionFromZone returns the region from a zone for Google cloud.
// Zones are formatted as <region>-<zone_suffix>, e.g. "us-central1-f" -> "us-central1"
func GetRegionFromZone(zone string) string {
	parts := strings.Split(zone, "-")
	if len(parts) >= 2 {
		return strings.Join(parts[:len(parts)-1], "-")
	}
	return zone
}
