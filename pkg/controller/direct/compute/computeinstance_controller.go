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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeInstanceGVK, NewComputeInstanceModel)
}

func NewComputeInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeInstanceModel{config: config}, nil
}

type computeInstanceModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computeInstanceModel{}

type computeInstanceAdapter struct {
	id              *krm.ComputeInstanceIdentity
	instancesClient *gcp.InstancesClient
	desired         *computepb.Instance
	actual          *computepb.Instance
	reader          client.Reader
}

var _ directbase.Adapter = &computeInstanceAdapter{}

func (m *computeInstanceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resolvedIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedIdentity.(*krm.ComputeInstanceIdentity)

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	instancesClient, err := gcpClient.newInstancesClient(ctx)
	if err != nil {
		return nil, err
	}

	var desired *computepb.Instance
	if !op.IsDeleting() {
		// Resolve references
		if err := resolveComputeInstanceRefs(ctx, reader, obj); err != nil {
			return nil, fmt.Errorf("resolving references: %w", err)
		}

		mapCtx := &direct.MapContext{}
		desired = ComputeInstanceSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}

		// Ensure Name and Zone are set on the desired proto
		desired.Name = &id.Instance
		desired.Zone = &id.Zone
	}

	return &computeInstanceAdapter{
		id:              id,
		instancesClient: instancesClient,
		desired:         desired,
		reader:          reader,
	}, nil
}

func (m *computeInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *computeInstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeInstance", "name", a.id)

	req := &computepb.GetInstanceRequest{
		Project:  a.id.Project,
		Zone:     a.id.Zone,
		Instance: a.id.Instance,
	}
	instance, err := a.instancesClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeInstance %q: %w", a.id, err)
	}
	a.actual = instance
	return true, nil
}

func (a *computeInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeInstance", "name", a.id)

	req := &computepb.InsertInstanceRequest{
		Project:          a.id.Project,
		Zone:             a.id.Zone,
		InstanceResource: a.desired,
	}
	op, err := a.instancesClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeInstance %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for ComputeInstance %s creation: %w", a.id, err)
		}
	}

	log.V(2).Info("successfully created ComputeInstance", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInstance %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *computeInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeInstance", "name", a.id)

	// Determine if instance is running and if we need to stop it
	isStarted := false
	if a.actual.GetStatus() == "RUNNING" || a.actual.GetStatus() == "PROVISIONING" || a.actual.GetStatus() == "STAGING" {
		isStarted = true
	}

	needStop := false
	if a.desired.MachineType != nil && a.actual.MachineType != nil {
		actualMachineType := a.actual.GetMachineType()
		desiredMachineType := a.desired.GetMachineType()
		if !strings.HasSuffix(actualMachineType, "/"+desiredMachineType) && actualMachineType != desiredMachineType {
			needStop = true
		}
	}
	if len(a.desired.ServiceAccounts) > 0 && len(a.actual.ServiceAccounts) > 0 {
		if !reflect.DeepEqual(a.desired.ServiceAccounts[0].Scopes, a.actual.ServiceAccounts[0].Scopes) ||
			direct.ValueOf(a.desired.ServiceAccounts[0].Email) != direct.ValueOf(a.actual.ServiceAccounts[0].Email) {
			needStop = true
		}
	} else if len(a.desired.ServiceAccounts) != len(a.actual.ServiceAccounts) {
		needStop = true
	}

	if needStop && isStarted {
		if !allowStoppingForUpdate(updateOp.GetUnstructured()) {
			return fmt.Errorf("cannot update machineType or serviceAccount on running instance without cnrm.cloud.google.com/allow-stopping-for-update annotation set to true")
		}
		if err := a.stopInstance(ctx); err != nil {
			return fmt.Errorf("stopping instance for update: %w", err)
		}
		// Refresh actual state after stop
		refreshed, err := a.get(ctx)
		if err != nil {
			return err
		}
		a.actual = refreshed
	}

	// 1. Machine Type
	if a.desired.MachineType != nil {
		actualMachineType := a.actual.GetMachineType()
		desiredMachineType := a.desired.GetMachineType()
		if !strings.HasSuffix(actualMachineType, "/"+desiredMachineType) && actualMachineType != desiredMachineType {
			log.Info("updating machine type", "name", a.id, "from", actualMachineType, "to", desiredMachineType)
			machineTypeURL := desiredMachineType
			if !strings.Contains(machineTypeURL, "/") {
				machineTypeURL = fmt.Sprintf("zones/%s/machineTypes/%s", a.id.Zone, desiredMachineType)
			}
			req := &computepb.SetMachineTypeInstanceRequest{
				Project:  a.id.Project,
				Zone:     a.id.Zone,
				Instance: a.id.Instance,
				InstancesSetMachineTypeRequestResource: &computepb.InstancesSetMachineTypeRequest{
					MachineType: &machineTypeURL,
				},
			}
			op, err := a.instancesClient.SetMachineType(ctx, req)
			if err != nil {
				return fmt.Errorf("setting machine type: %w", err)
			}
			if err := op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting for machine type update: %w", err)
			}
		}
	}

	// 2. Service Account
	if len(a.desired.ServiceAccounts) > 0 {
		desiredSA := a.desired.ServiceAccounts[0]
		actualSAExists := len(a.actual.ServiceAccounts) > 0
		var actualSA *computepb.ServiceAccount
		if actualSAExists {
			actualSA = a.actual.ServiceAccounts[0]
		}
		if !actualSAExists || !reflect.DeepEqual(desiredSA.Scopes, actualSA.Scopes) || direct.ValueOf(desiredSA.Email) != direct.ValueOf(actualSA.Email) {
			log.Info("updating service account scopes", "name", a.id)
			req := &computepb.SetServiceAccountInstanceRequest{
				Project:  a.id.Project,
				Zone:     a.id.Zone,
				Instance: a.id.Instance,
				InstancesSetServiceAccountRequestResource: &computepb.InstancesSetServiceAccountRequest{
					Email:  desiredSA.Email,
					Scopes: desiredSA.Scopes,
				},
			}
			op, err := a.instancesClient.SetServiceAccount(ctx, req)
			if err != nil {
				return fmt.Errorf("setting service account: %w", err)
			}
			if err := op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting for service account update: %w", err)
			}
		}
	}

	// 3. Metadata
	desiredMetadata := a.desired.Metadata
	actualMetadata := a.actual.Metadata
	if desiredMetadata == nil {
		desiredMetadata = &computepb.Metadata{}
	}
	if actualMetadata == nil {
		actualMetadata = &computepb.Metadata{}
	}
	if !reflect.DeepEqual(desiredMetadata.Items, actualMetadata.Items) {
		log.Info("updating metadata", "name", a.id)
		req := &computepb.SetMetadataInstanceRequest{
			Project:  a.id.Project,
			Zone:     a.id.Zone,
			Instance: a.id.Instance,
			MetadataResource: &computepb.Metadata{
				Fingerprint: actualMetadata.Fingerprint,
				Items:       desiredMetadata.Items,
			},
		}
		op, err := a.instancesClient.SetMetadata(ctx, req)
		if err != nil {
			return fmt.Errorf("setting metadata: %w", err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for metadata update: %w", err)
		}
	}

	// 4. Tags
	desiredTags := a.desired.Tags
	actualTags := a.actual.Tags
	if desiredTags == nil {
		desiredTags = &computepb.Tags{}
	}
	if actualTags == nil {
		actualTags = &computepb.Tags{}
	}
	if !reflect.DeepEqual(desiredTags.Items, actualTags.Items) {
		log.Info("updating tags", "name", a.id)
		req := &computepb.SetTagsInstanceRequest{
			Project:  a.id.Project,
			Zone:     a.id.Zone,
			Instance: a.id.Instance,
			TagsResource: &computepb.Tags{
				Fingerprint: actualTags.Fingerprint,
				Items:       desiredTags.Items,
			},
		}
		op, err := a.instancesClient.SetTags(ctx, req)
		if err != nil {
			return fmt.Errorf("setting tags: %w", err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for tags update: %w", err)
		}
	}

	// 5. Scheduling
	if !reflect.DeepEqual(a.desired.Scheduling, a.actual.Scheduling) {
		log.Info("updating scheduling", "name", a.id)
		req := &computepb.SetSchedulingInstanceRequest{
			Project:            a.id.Project,
			Zone:               a.id.Zone,
			Instance:           a.id.Instance,
			SchedulingResource: a.desired.Scheduling,
		}
		op, err := a.instancesClient.SetScheduling(ctx, req)
		if err != nil {
			return fmt.Errorf("setting scheduling: %w", err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for scheduling update: %w", err)
		}
	}

	// 6. Labels
	desiredLabels := updateOp.GetUnstructured().GetLabels()
	actualLabels := a.actual.Labels
	if !reflect.DeepEqual(desiredLabels, actualLabels) {
		log.Info("updating labels", "name", a.id)
		req := &computepb.SetLabelsInstanceRequest{
			Project:  a.id.Project,
			Zone:     a.id.Zone,
			Instance: a.id.Instance,
			InstancesSetLabelsRequestResource: &computepb.InstancesSetLabelsRequest{
				LabelFingerprint: a.actual.LabelFingerprint,
				Labels:           desiredLabels,
			},
		}
		op, err := a.instancesClient.SetLabels(ctx, req)
		if err != nil {
			return fmt.Errorf("setting labels: %w", err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for labels update: %w", err)
		}
	}

	// Restart if stopped and was previously started
	if needStop && isStarted {
		if err := a.startInstance(ctx); err != nil {
			return fmt.Errorf("starting instance after update: %w", err)
		}
	}

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeInstance %s: %w", a.id, err)
	}

	// Compute any diffs for structured reporting
	paths, err := common.CompareProtoMessage(a.desired, updated, common.BasicDiff)
	if err == nil && len(paths) > 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *computeInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeInstance", "name", a.id)

	req := &computepb.DeleteInstanceRequest{
		Project:  a.id.Project,
		Zone:     a.id.Zone,
		Instance: a.id.Instance,
	}
	op, err := a.instancesClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeInstance %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for ComputeInstance %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ComputeInstance", "name", a.id)
	return true, nil
}

func (a *computeInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *computeInstanceAdapter) get(ctx context.Context) (*computepb.Instance, error) {
	req := &computepb.GetInstanceRequest{
		Project:  a.id.Project,
		Zone:     a.id.Zone,
		Instance: a.id.Instance,
	}
	return a.instancesClient.Get(ctx, req)
}

func (a *computeInstanceAdapter) stopInstance(ctx context.Context) error {
	log := klog.FromContext(ctx)
	log.Info("stopping ComputeInstance to apply updates", "name", a.id)
	req := &computepb.StopInstanceRequest{
		Project:  a.id.Project,
		Zone:     a.id.Zone,
		Instance: a.id.Instance,
	}
	op, err := a.instancesClient.Stop(ctx, req)
	if err != nil {
		return err
	}
	return op.Wait(ctx)
}

func (a *computeInstanceAdapter) startInstance(ctx context.Context) error {
	log := klog.FromContext(ctx)
	log.Info("restarting ComputeInstance after applying updates", "name", a.id)
	req := &computepb.StartInstanceRequest{
		Project:  a.id.Project,
		Zone:     a.id.Zone,
		Instance: a.id.Instance,
	}
	op, err := a.instancesClient.Start(ctx, req)
	if err != nil {
		return err
	}
	return op.Wait(ctx)
}

func (a *computeInstanceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := ComputeInstanceStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func allowStoppingForUpdate(u *unstructured.Unstructured) bool {
	val, _, _ := unstructured.NestedString(u.Object, "metadata", "annotations", "cnrm.cloud.google.com/allow-stopping-for-update")
	return val == "true"
}
