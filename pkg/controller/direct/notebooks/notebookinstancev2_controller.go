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

package notebooks

import (
	"context"
	"fmt"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/notebooks/apiv2"
	notebookspb "cloud.google.com/go/notebooks/apiv2/notebookspb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NotebookInstanceV2GVK, NewInstanceV2Model)
}

func NewInstanceV2Model(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstanceV2{config: *config}, nil
}

var _ directbase.Model = &modelInstanceV2{}

type modelInstanceV2 struct {
	config config.ControllerConfig
}

func (m *modelInstanceV2) client(ctx context.Context) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewNotebookClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building InstanceV2 client: %w", err)
	}
	return gcpClient, err
}

func (m *modelInstanceV2) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NotebookInstanceV2{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Manually resolve IAMServiceAccountRef
	if obj.Spec.GCESetup != nil {
		for i := range obj.Spec.GCESetup.ServiceAccounts {
			sa := &obj.Spec.GCESetup.ServiceAccounts[i]
			if sa.ServiceAccountRef != nil {
				klog.FromContext(ctx).Info("DEBUG: Before resolving", "name", sa.ServiceAccountRef.Name, "external", sa.ServiceAccountRef.External)
				if err := sa.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
					return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
				}
				klog.FromContext(ctx).Info("DEBUG: After resolving", "name", sa.ServiceAccountRef.Name, "external", sa.ServiceAccountRef.External)
			}
		}
	}

	identityObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identityObj.(*krm.NotebookInstanceV2Identity)

	// Get notebooks GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := Spec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &InstanceV2Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelInstanceV2) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.NotebookInstanceV2Identity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceV2Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type InstanceV2Adapter struct {
	id        *krm.NotebookInstanceV2Identity
	gcpClient *gcp.NotebookClient
	desired   *notebookspb.Instance
	actual    *notebookspb.Instance
}

var _ directbase.Adapter = &InstanceV2Adapter{}

// Find retrieves the GCP resource.
func (a *InstanceV2Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InstanceV2", "name", a.id)

	req := &notebookspb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InstanceV2 %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP.
func (a *InstanceV2Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InstanceV2", "name", a.id)

	req := &notebookspb.CreateInstanceRequest{
		Parent:     a.id.ParentString(),
		InstanceId: a.id.Instance,
		Instance:   a.desired,
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InstanceV2 %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("instanceV2 %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created InstanceV2", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

// Update updates the resource in GCP.
func (a *InstanceV2Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InstanceV2", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareNotebooksV2(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &notebookspb.UpdateInstanceRequest{
		Instance:   a.desired,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating InstanceV2 %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("instanceV2 %s waiting update: %w", a.id, err)
	}

	log.V(2).Info("successfully updated InstanceV2", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func compareNotebooksV2(ctx context.Context, actual, desired *notebookspb.Instance) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, Spec_FromProto, Spec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*notebookspb.Instance)

	populateDefaults := func(act, des *notebookspb.Instance) {
		actGCE := act.GetGceSetup()
		desGCE := des.GetGceSetup()
		if desGCE == nil {
			desGCE = &notebookspb.GceSetup{}
			des.Infrastructure = &notebookspb.Instance_GceSetup{GceSetup: desGCE}
		}
		if actGCE == nil {
			actGCE = &notebookspb.GceSetup{}
			act.Infrastructure = &notebookspb.Instance_GceSetup{GceSetup: actGCE}
		}

		// Copy/align ServiceAccounts if not specified in desired
		if len(desGCE.ServiceAccounts) == 0 && len(actGCE.ServiceAccounts) > 0 {
			desGCE.ServiceAccounts = actGCE.ServiceAccounts
		}
		// Copy/align BootDisk if not specified in desired
		if desGCE.BootDisk == nil && actGCE.BootDisk != nil {
			desGCE.BootDisk = actGCE.BootDisk
		}
		// Copy/align BootDisk KMS keys if not specified in desired
		if desGCE.BootDisk != nil && actGCE.BootDisk != nil {
			if desGCE.BootDisk.KmsKey == "" && actGCE.BootDisk.KmsKey != "" {
				desGCE.BootDisk.KmsKey = actGCE.BootDisk.KmsKey
				desGCE.BootDisk.DiskEncryption = actGCE.BootDisk.DiskEncryption
			}
		}
		// Copy/align DataDisks if not specified in desired
		if len(desGCE.DataDisks) == 0 && len(actGCE.DataDisks) > 0 {
			desGCE.DataDisks = actGCE.DataDisks
		}
		// Copy/align DataDisks KMS keys if not specified in desired
		if len(desGCE.DataDisks) > 0 && len(actGCE.DataDisks) == len(desGCE.DataDisks) {
			for i := range desGCE.DataDisks {
				desDisk := desGCE.DataDisks[i]
				actDisk := actGCE.DataDisks[i]
				if desDisk.KmsKey == "" && actDisk.KmsKey != "" {
					desDisk.KmsKey = actDisk.KmsKey
					desDisk.DiskEncryption = actDisk.DiskEncryption
				}
			}
		}
		// Copy/align NetworkInterfaces if not specified in desired
		if len(desGCE.NetworkInterfaces) == 0 && len(actGCE.NetworkInterfaces) > 0 {
			desGCE.NetworkInterfaces = actGCE.NetworkInterfaces
		}
		// Copy/align Image (VmImage / ContainerImage) if not specified in desired
		if desGCE.Image == nil && actGCE.Image != nil {
			desGCE.Image = actGCE.Image
		}
		// Copy/align VmImage fields if present on both sides
		desVM := desGCE.GetVmImage()
		actVM := actGCE.GetVmImage()
		if desVM != nil && actVM != nil {
			if desVM.Project == "" {
				desVM.Project = actVM.Project
			}
			if desVM.GetFamily() == "" && actVM.GetFamily() != "" {
				desVM.Image = &notebookspb.VmImage_Family{Family: actVM.GetFamily()}
			}
			if desVM.GetName() == "" && actVM.GetName() != "" {
				desVM.Image = &notebookspb.VmImage_Name{Name: actVM.GetName()}
			}
		}
		// Copy/align ContainerImage fields if present on both sides
		desContainer := desGCE.GetContainerImage()
		actContainer := actGCE.GetContainerImage()
		if desContainer != nil && actContainer != nil {
			if desContainer.Repository == "" {
				desContainer.Repository = actContainer.Repository
			}
			if desContainer.Tag == "" {
				desContainer.Tag = actContainer.Tag
			}
		}
		// Copy/align GpuDriverConfig, Tags, and Metadata (which are immutable on GCP)
		desGCE.GpuDriverConfig = actGCE.GpuDriverConfig
		desGCE.Tags = actGCE.Tags
		desGCE.Metadata = actGCE.Metadata
		// Copy/align DisablePublicIp if not specified in desired
		if !desGCE.DisablePublicIp && actGCE.DisablePublicIp {
			desGCE.DisablePublicIp = actGCE.DisablePublicIp
		}
		// Copy/align EnableIpForwarding if not specified in desired
		if !desGCE.EnableIpForwarding && actGCE.EnableIpForwarding {
			desGCE.EnableIpForwarding = actGCE.EnableIpForwarding
		}
		// Copy/align ShieldedInstanceConfig defaults
		if actGCE.ShieldedInstanceConfig == nil {
			actGCE.ShieldedInstanceConfig = &notebookspb.ShieldedInstanceConfig{
				EnableSecureBoot:          false,
				EnableVtpm:                true,
				EnableIntegrityMonitoring: true,
			}
		}
		if desGCE.ShieldedInstanceConfig == nil {
			desGCE.ShieldedInstanceConfig = &notebookspb.ShieldedInstanceConfig{
				EnableSecureBoot:          false,
				EnableVtpm:                true,
				EnableIntegrityMonitoring: true,
			}
		}
	}
	populateDefaults(maskedActual, clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *InstanceV2Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *notebookspb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NotebookInstanceV2Status{}
	status.ObservedState = NotebookInstanceV2ObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstanceV2Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NotebookInstanceV2{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(Spec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Instance)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Instance)
	u.SetGroupVersionKind(krm.NotebookInstanceV2GVK)

	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

// Delete the resource from GCP service.
func (a *InstanceV2Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InstanceV2", "name", a.id)

	req := &notebookspb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent InstanceV2, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting InstanceV2 %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted InstanceV2", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete InstanceV2 %s: %w", a.id, err)
	}
	return true, nil
}

func Spec_ToProto(mapCtx *direct.MapContext, in *krm.NotebookInstanceV2Spec) *notebookspb.Instance {
	if in == nil {
		return nil
	}
	out := NotebookInstanceV2Spec_v1alpha1_ToProto(mapCtx, in)
	if out == nil {
		return nil
	}
	// Manual mapping for ServiceAccounts
	if in.GCESetup != nil && out.GetGceSetup() != nil {
		outGCE := out.GetGceSetup()
		for i, sa := range in.GCESetup.ServiceAccounts {
			if i < len(outGCE.ServiceAccounts) && sa.ServiceAccountRef != nil {
				outGCE.ServiceAccounts[i].Email = sa.ServiceAccountRef.External
			}
		}
		// Manual mapping for GPUDriverConfig
		if in.GCESetup.GPUDriverConfig != nil {
			outGCE.GpuDriverConfig = &notebookspb.GPUDriverConfig{
				EnableGpuDriver:     direct.ValueOf(in.GCESetup.GPUDriverConfig.EnableGpuDriver),
				CustomGpuDriverPath: direct.ValueOf(in.GCESetup.GPUDriverConfig.CustomGpuDriverPath),
			}
		}
		// Manual mapping for DataDisks KMS keys
		for i, dd := range in.GCESetup.DataDisks {
			if i < len(outGCE.DataDisks) && dd.KmsKeyRef != nil {
				outGCE.DataDisks[i].KmsKey = dd.KmsKeyRef.External
				outGCE.DataDisks[i].DiskEncryption = notebookspb.DiskEncryption_CMEK
			}
		}
	}
	return out
}

func Spec_FromProto(mapCtx *direct.MapContext, in *notebookspb.Instance) *krm.NotebookInstanceV2Spec {
	if in == nil {
		return nil
	}
	out := NotebookInstanceV2Spec_v1alpha1_FromProto(mapCtx, in)
	if out == nil {
		return nil
	}
	// Manual mapping for ServiceAccounts
	if in.GetGceSetup() != nil && out.GCESetup != nil {
		inGCE := in.GetGceSetup()
		for i, sa := range inGCE.ServiceAccounts {
			if i < len(out.GCESetup.ServiceAccounts) && sa.GetEmail() != "" {
				out.GCESetup.ServiceAccounts[i].ServiceAccountRef = &refs.IAMServiceAccountRef{External: sa.GetEmail()}
			}
		}
		// Manual mapping for GPUDriverConfig
		if inGCE.GetGpuDriverConfig() != nil {
			out.GCESetup.GPUDriverConfig = &krm.InstanceGPUDriverConfig{
				EnableGpuDriver:     direct.LazyPtr(inGCE.GetGpuDriverConfig().GetEnableGpuDriver()),
				CustomGpuDriverPath: direct.LazyPtr(inGCE.GetGpuDriverConfig().GetCustomGpuDriverPath()),
			}
		}
		// Manual mapping for DataDisks KMS keys
		for i, dd := range inGCE.DataDisks {
			if i < len(out.GCESetup.DataDisks) && dd.GetKmsKey() != "" {
				out.GCESetup.DataDisks[i].KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: dd.GetKmsKey()}
			}
		}
	}
	return out
}
