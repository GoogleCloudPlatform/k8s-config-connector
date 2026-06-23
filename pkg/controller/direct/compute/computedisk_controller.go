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
// proto.service: google.cloud.compute.v1.RegionDisks
// proto.message: google.cloud.compute.v1.Disk
// crd.type: ComputeDisk
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeDiskGVK, NewComputeDiskModel)
}

func NewComputeDiskModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeDiskModel{config: config}, nil
}

var _ directbase.Model = &computeDiskModel{}

type computeDiskModel struct {
	config *config.ControllerConfig
}

func (m *computeDiskModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeDisk{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	diskId := id.(*v1beta1.ComputeDiskIdentity)
	if err := diskId.Validate(); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	disksClient, err := gcpClient.newDisksClient(ctx)
	if err != nil {
		return nil, err
	}
	regionDisksClient, err := gcpClient.newRegionDisksClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeDiskSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(diskId.Disk)
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	// Normalize Type if specified by the user to be fully qualified projects/...
	if desired.Type != nil {
		val := *desired.Type
		if !strings.Contains(val, "/") {
			if diskId.IsZonal() {
				desired.Type = direct.PtrTo(fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", diskId.Project, diskId.Zone, val))
			} else {
				desired.Type = direct.PtrTo(fmt.Sprintf("projects/%s/regions/%s/diskTypes/%s", diskId.Project, diskId.Region, val))
			}
		}
	}

	return &ComputeDiskAdapter{
		disksClient:       disksClient,
		regionDisksClient: regionDisksClient,
		id:                diskId,
		desired:           desired,
		reader:            reader,
	}, nil
}

func (m *computeDiskModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeDiskAdapter struct {
	disksClient       *gcp.DisksClient
	regionDisksClient *gcp.RegionDisksClient
	id                *v1beta1.ComputeDiskIdentity
	desired           *computepb.Disk
	actual            *computepb.Disk
	reader            client.Reader
}

var _ directbase.Adapter = &ComputeDiskAdapter{}

func (a *ComputeDiskAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeDisk", "name", a.id)

	if a.id.IsZonal() {
		req := &computepb.GetDiskRequest{
			Project: a.id.Project,
			Zone:    a.id.Zone,
			Disk:    a.id.Disk,
		}
		actual, err := a.disksClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting ComputeDisk %q: %w", a.id, err)
		}
		a.actual = actual
		return true, nil
	} else if a.id.IsRegional() {
		req := &computepb.GetRegionDiskRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Disk:    a.id.Disk,
		}
		actual, err := a.regionDisksClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting ComputeRegionDisk %q: %w", a.id, err)
		}
		a.actual = actual
		return true, nil
	} else {
		return false, fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
	}
}

func (a *ComputeDiskAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeDisk", "name", a.id)

	if a.id.IsZonal() {
		req := &computepb.InsertDiskRequest{
			Project:      a.id.Project,
			Zone:         a.id.Zone,
			DiskResource: a.desired,
		}
		op, err := a.disksClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeDisk %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeDisk %s waiting creation: %w", a.id.String(), err)
		}
	} else if a.id.IsRegional() {
		req := &computepb.InsertRegionDiskRequest{
			Project:      a.id.Project,
			Region:       a.id.Region,
			DiskResource: a.desired,
		}
		op, err := a.regionDisksClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeRegionDisk %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeRegionDisk %s waiting creation: %w", a.id.String(), err)
		}
	} else {
		return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
	}

	log.Info("successfully created compute ComputeDisk in gcp", "name", a.id)

	// Get the created resource
	var created *computepb.Disk
	var err error
	if a.id.IsZonal() {
		created, err = a.disksClient.Get(ctx, &computepb.GetDiskRequest{
			Project: a.id.Project,
			Zone:    a.id.Zone,
			Disk:    a.id.Disk,
		})
	} else if a.id.IsRegional() {
		created, err = a.regionDisksClient.Get(ctx, &computepb.GetRegionDiskRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Disk:    a.id.Disk,
		})
	} else {
		return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
	}
	if err != nil {
		return fmt.Errorf("getting ComputeDisk %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeDiskAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeDisk", "name", a.id)

	diffs, _, err := compareComputeDisk(ctx, a.actual, a.desired, a.id)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		// Check if labels changed
		labelsChanged := false
		if len(a.actual.Labels) != len(a.desired.Labels) {
			labelsChanged = true
		} else {
			for k, v := range a.desired.Labels {
				if a.actual.Labels[k] != v {
					labelsChanged = true
					break
				}
			}
		}

		if labelsChanged {
			log.V(2).Info("updating ComputeDisk labels", "name", a.id)
			if a.id.IsZonal() {
				req := &computepb.SetLabelsDiskRequest{
					Project:  a.id.Project,
					Zone:     a.id.Zone,
					Resource: a.id.Disk,
					ZoneSetLabelsRequestResource: &computepb.ZoneSetLabelsRequest{
						Labels:           a.desired.Labels,
						LabelFingerprint: a.actual.LabelFingerprint,
					},
				}
				op, err := a.disksClient.SetLabels(ctx, req)
				if err != nil {
					return fmt.Errorf("updating ComputeDisk labels %s: %w", a.id, err)
				}
				if err = op.Wait(ctx); err != nil {
					return fmt.Errorf("compute ComputeDisk %s waiting labels update: %w", a.id, err)
				}
			} else if a.id.IsRegional() {
				req := &computepb.SetLabelsRegionDiskRequest{
					Project:  a.id.Project,
					Region:   a.id.Region,
					Resource: a.id.Disk,
					RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{
						Labels:           a.desired.Labels,
						LabelFingerprint: a.actual.LabelFingerprint,
					},
				}
				op, err := a.regionDisksClient.SetLabels(ctx, req)
				if err != nil {
					return fmt.Errorf("updating ComputeRegionDisk labels %s: %w", a.id, err)
				}
				if err = op.Wait(ctx); err != nil {
					return fmt.Errorf("compute ComputeRegionDisk %s waiting labels update: %w", a.id, err)
				}
			} else {
				return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
			}
		}

		// Check if size changed
		if a.desired.SizeGb != nil && direct.ValueOf(a.desired.SizeGb) != direct.ValueOf(a.actual.SizeGb) {
			log.V(2).Info("updating ComputeDisk size", "name", a.id, "old", direct.ValueOf(a.actual.SizeGb), "new", direct.ValueOf(a.desired.SizeGb))
			if a.id.IsZonal() {
				req := &computepb.ResizeDiskRequest{
					Project: a.id.Project,
					Zone:    a.id.Zone,
					Disk:    a.id.Disk,
					DisksResizeRequestResource: &computepb.DisksResizeRequest{
						SizeGb: a.desired.SizeGb,
					},
				}
				op, err := a.disksClient.Resize(ctx, req)
				if err != nil {
					return fmt.Errorf("resizing ComputeDisk %s: %w", a.id, err)
				}
				if err = op.Wait(ctx); err != nil {
					return fmt.Errorf("compute ComputeDisk %s waiting resize: %w", a.id, err)
				}
			} else if a.id.IsRegional() {
				req := &computepb.ResizeRegionDiskRequest{
					Project: a.id.Project,
					Region:  a.id.Region,
					Disk:    a.id.Disk,
					RegionDisksResizeRequestResource: &computepb.RegionDisksResizeRequest{
						SizeGb: a.desired.SizeGb,
					},
				}
				op, err := a.regionDisksClient.Resize(ctx, req)
				if err != nil {
					return fmt.Errorf("resizing ComputeRegionDisk %s: %w", a.id, err)
				}
				if err = op.Wait(ctx); err != nil {
					return fmt.Errorf("compute ComputeRegionDisk %s waiting resize: %w", a.id, err)
				}
			} else {
				return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
			}
		}

		// Check if resource policies changed
		toAdd, toRemove := diffResourcePolicies(a.actual.ResourcePolicies, a.desired.ResourcePolicies)
		if len(toAdd) > 0 || len(toRemove) > 0 {
			log.V(2).Info("updating ComputeDisk resource policies", "name", a.id, "toAdd", toAdd, "toRemove", toRemove)
			if a.id.IsZonal() {
				if len(toAdd) > 0 {
					req := &computepb.AddResourcePoliciesDiskRequest{
						Project: a.id.Project,
						Zone:    a.id.Zone,
						Disk:    a.id.Disk,
						DisksAddResourcePoliciesRequestResource: &computepb.DisksAddResourcePoliciesRequest{
							ResourcePolicies: toAdd,
						},
					}
					op, err := a.disksClient.AddResourcePolicies(ctx, req)
					if err != nil {
						return fmt.Errorf("adding resource policies to ComputeDisk %s: %w", a.id, err)
					}
					if err = op.Wait(ctx); err != nil {
						return fmt.Errorf("compute ComputeDisk %s waiting resource policies add: %w", a.id, err)
					}
				}
				if len(toRemove) > 0 {
					req := &computepb.RemoveResourcePoliciesDiskRequest{
						Project: a.id.Project,
						Zone:    a.id.Zone,
						Disk:    a.id.Disk,
						DisksRemoveResourcePoliciesRequestResource: &computepb.DisksRemoveResourcePoliciesRequest{
							ResourcePolicies: toRemove,
						},
					}
					op, err := a.disksClient.RemoveResourcePolicies(ctx, req)
					if err != nil {
						return fmt.Errorf("removing resource policies from ComputeDisk %s: %w", a.id, err)
					}
					if err = op.Wait(ctx); err != nil {
						return fmt.Errorf("compute ComputeDisk %s waiting resource policies remove: %w", a.id, err)
					}
				}
			} else if a.id.IsRegional() {
				if len(toAdd) > 0 {
					req := &computepb.AddResourcePoliciesRegionDiskRequest{
						Project: a.id.Project,
						Region:  a.id.Region,
						Disk:    a.id.Disk,
						RegionDisksAddResourcePoliciesRequestResource: &computepb.RegionDisksAddResourcePoliciesRequest{
							ResourcePolicies: toAdd,
						},
					}
					op, err := a.regionDisksClient.AddResourcePolicies(ctx, req)
					if err != nil {
						return fmt.Errorf("adding resource policies to ComputeRegionDisk %s: %w", a.id, err)
					}
					if err = op.Wait(ctx); err != nil {
						return fmt.Errorf("compute ComputeRegionDisk %s waiting resource policies add: %w", a.id, err)
					}
				}
				if len(toRemove) > 0 {
					req := &computepb.RemoveResourcePoliciesRegionDiskRequest{
						Project: a.id.Project,
						Region:  a.id.Region,
						Disk:    a.id.Disk,
						RegionDisksRemoveResourcePoliciesRequestResource: &computepb.RegionDisksRemoveResourcePoliciesRequest{
							ResourcePolicies: toRemove,
						},
					}
					op, err := a.regionDisksClient.RemoveResourcePolicies(ctx, req)
					if err != nil {
						return fmt.Errorf("removing resource policies from ComputeRegionDisk %s: %w", a.id, err)
					}
					if err = op.Wait(ctx); err != nil {
						return fmt.Errorf("compute ComputeRegionDisk %s waiting resource policies remove: %w", a.id, err)
					}
				}
			} else {
				return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
			}
		}

		// Retrieve the latest disk state after all updates
		var getErr error
		if a.id.IsZonal() {
			latest, getErr = a.disksClient.Get(ctx, &computepb.GetDiskRequest{
				Project: a.id.Project,
				Zone:    a.id.Zone,
				Disk:    a.id.Disk,
			})
		} else if a.id.IsRegional() {
			latest, getErr = a.regionDisksClient.Get(ctx, &computepb.GetRegionDiskRequest{
				Project: a.id.Project,
				Region:  a.id.Region,
				Disk:    a.id.Disk,
			})
		} else {
			return fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
		}
		if getErr != nil {
			return fmt.Errorf("getting ComputeDisk %s after update: %w", a.id, getErr)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ComputeDiskAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeDisk{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeDiskSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	if a.id.IsZonal() {
		obj.Spec.Location = a.id.Zone
	} else if a.id.IsRegional() {
		obj.Spec.Location = a.id.Region
	} else {
		return nil, fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeDiskGVK)
	return u, nil
}

func (a *ComputeDiskAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeDisk", "name", a.id)

	if a.id.IsZonal() {
		req := &computepb.DeleteDiskRequest{
			Project: a.id.Project,
			Zone:    a.id.Zone,
			Disk:    a.id.Disk,
		}
		op, err := a.disksClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				log.V(2).Info("skipping delete for non-existent ComputeDisk, assuming it was already deleted", "name", a.id.String())
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeDisk %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("ComputeDisk %s waiting deletion: %w", a.id, err)
		}
	} else if a.id.IsRegional() {
		req := &computepb.DeleteRegionDiskRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Disk:    a.id.Disk,
		}
		op, err := a.regionDisksClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				log.V(2).Info("skipping delete for non-existent ComputeRegionDisk, assuming it was already deleted", "name", a.id.String())
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeRegionDisk %s: %w", a.id, err)
		}
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("ComputeRegionDisk %s waiting deletion: %w", a.id, err)
		}
	} else {
		return false, fmt.Errorf("ComputeDisk %s is neither zonal nor regional", a.id)
	}

	log.V(2).Info("successfully deleted ComputeDisk", "name", a.id)
	return true, nil
}

func (a *ComputeDiskAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Disk) error {
	mapCtx := &direct.MapContext{}
	status := ComputeDiskStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func canonicalizeComputeURL(t *string) *string {
	if t == nil {
		return nil
	}
	trimmed := refs.TrimComputeURIPrefix(*t)
	return &trimmed
}

func canonicalizeComputeURLs(arr []string) []string {
	if arr == nil {
		return nil
	}
	out := make([]string, len(arr))
	for i, val := range arr {
		out[i] = refs.TrimComputeURIPrefix(val)
	}
	return out
}

func compareComputeDisk(ctx context.Context, actual, desired *computepb.Disk, id *v1beta1.ComputeDiskIdentity) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeDiskSpec_v1beta1_FromProto, ComputeDiskSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)
	clonedDesired.Name = actual.Name

	populateDefaults := func(obj *computepb.Disk) {
		// Populate physical_block_size_bytes default
		if obj.PhysicalBlockSizeBytes == nil {
			obj.PhysicalBlockSizeBytes = direct.PtrTo(int64(4096))
		}

		// Populate type default
		if obj.Type == nil {
			if id.IsZonal() {
				obj.Type = direct.PtrTo(fmt.Sprintf("projects/%s/zones/%s/diskTypes/pd-standard", id.Project, id.Zone))
			} else {
				obj.Type = direct.PtrTo(fmt.Sprintf("projects/%s/regions/%s/diskTypes/pd-standard", id.Project, id.Region))
			}
		} else {
			val := *obj.Type
			if !strings.Contains(val, "/") {
				if id.IsZonal() {
					obj.Type = direct.PtrTo(fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", id.Project, id.Zone, val))
				} else {
					obj.Type = direct.PtrTo(fmt.Sprintf("projects/%s/regions/%s/diskTypes/%s", id.Project, id.Region, val))
				}
			}
		}

		// Canonicalize URLs to projects/... format
		obj.Type = canonicalizeComputeURL(obj.Type)
		obj.SourceDisk = canonicalizeComputeURL(obj.SourceDisk)
		obj.SourceImage = canonicalizeComputeSourceImage(obj.SourceImage)
		obj.SourceSnapshot = canonicalizeComputeSnapshot(obj.SourceSnapshot)

		// Canonicalize ResourcePolicies
		for i, p := range obj.ResourcePolicies {
			if u := canonicalizeComputeURL(&p); u != nil {
				obj.ResourcePolicies[i] = *u
			}
		}
	}

	populateDefaults(clonedDesired)
	populateDefaults(maskedActual)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func canonicalizeComputeSourceImage(t *string) *string {
	if t == nil {
		return nil
	}
	val := *t
	// Wait, sourceImage is often returned by GCP with full path, or images/debian-11
	// Let's strip standard prefixes
	val = refs.TrimComputeURIPrefix(val)
	// If it was family/debian-11, etc., sometimes GCP returns bullseye-v...
	// Just return standard trimmed
	return &val
}

func canonicalizeComputeSnapshot(t *string) *string {
	if t == nil {
		return nil
	}
	val := refs.TrimComputeURIPrefix(*t)
	return &val
}

func diffResourcePolicies(actual, desired []string) (toAdd, toRemove []string) {
	actualMap := make(map[string]bool)
	for _, p := range actual {
		trimmed := refs.TrimComputeURIPrefix(p)
		actualMap[trimmed] = true
	}

	desiredMap := make(map[string]bool)
	for _, p := range desired {
		trimmed := refs.TrimComputeURIPrefix(p)
		desiredMap[trimmed] = true
	}

	for p := range desiredMap {
		if !actualMap[p] {
			toAdd = append(toAdd, p)
		}
	}

	for p := range actualMap {
		if !desiredMap[p] {
			toRemove = append(toRemove, p)
		}
	}

	return toAdd, toRemove
}
