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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.BackupPlanAssociation
// crd.type: BackupDRBackupPlanAssociation
// crd.version: v1beta1

package backupdr

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/backupdr/apiv1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	gcpcompute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BackupDRBackupPlanAssociationGVK, NewBackupPlanAssociationModel)
}

func NewBackupPlanAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupPlanAssociation{config: *config}, nil
}

var _ directbase.Model = &modelBackupPlanAssociation{}

type modelBackupPlanAssociation struct {
	config config.ControllerConfig
}

func (m *modelBackupPlanAssociation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BackupDRBackupPlanAssociation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupPlanAssociationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Resolve SQLInstanceRef if present
	if obj.Spec.Resource != nil && obj.Spec.Resource.SQLInstanceRef != nil {
		sqlInst, err := refs.ResolveSQLInstanceRef(ctx, reader, obj, obj.Spec.Resource.SQLInstanceRef)
		if err != nil {
			return nil, fmt.Errorf("resolving SQLInstanceRef: %w", err)
		}
		// Cloud SQL instances in BackupDR must use the "projects/{project}/instances/{instance}" format
		obj.Spec.Resource.SQLInstanceRef.External = "projects/" + sqlInst.ProjectID + "/instances/" + sqlInst.SQLInstanceName
	}

	// Get backupdr GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	backupDRClient, err := gcpClient.newBackupDRClient(ctx)
	if err != nil {
		return nil, err
	}
	computeClient, err := gcpClient.newComputeInstancesClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := BackupDRBackupPlanAssociationSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &BackupPlanAssociationAdapter{
		id:            id,
		gcpClient:     backupDRClient,
		computeClient: computeClient,
		projectMapper: m.config.ProjectMapper,
		desired:       desired,
		reader:        reader,
	}, nil
}

func (m *modelBackupPlanAssociation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupPlanAssociationAdapter struct {
	id            *krm.BackupPlanAssociationIdentity
	gcpClient     *gcp.Client
	computeClient *gcpcompute.InstancesClient
	projectMapper *projects.ProjectMapper
	desired       *pb.BackupPlanAssociation
	actual        *pb.BackupPlanAssociation
	reader        client.Reader
}

var _ directbase.Adapter = &BackupPlanAssociationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupPlanAssociationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupPlanAssociation", "name", a.id)

	req := &pb.GetBackupPlanAssociationRequest{Name: a.id.String()}
	backupplanassociationpb, err := a.gcpClient.GetBackupPlanAssociation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackupPlanAssociation %q: %w", a.id, err)
	}

	a.actual = backupplanassociationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAssociationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupPlanAssociation", "name", a.id)

	req := &pb.CreateBackupPlanAssociationRequest{
		Parent:                  a.id.Parent().String(),
		BackupPlanAssociationId: a.id.ID(),
		BackupPlanAssociation:   a.desired,
	}
	op, err := a.gcpClient.CreateBackupPlanAssociation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BackupPlanAssociation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BackupPlanAssociation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackupPlanAssociation", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.BackupDRBackupPlanAssociationStatus{}
	status.ObservedState = BackupDRBackupPlanAssociationObservedState_v1beta1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAssociationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupPlanAssociation", "name", a.id)

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	actual, err := normalizeActualReferences(ctx, a.projectMapper, a.getComputeInstanceID, a.desired, a.actual)
	if err != nil {
		return err
	}

	paths := []string{}
	if !reflect.DeepEqual(a.desired.ResourceType, actual.ResourceType) {
		report.AddField("resource_type", actual.ResourceType, a.desired.ResourceType)
		paths = append(paths, "resource_type")
	}
	if !reflect.DeepEqual(a.desired.Resource, actual.Resource) {
		report.AddField("resource", actual.Resource, a.desired.Resource)
		paths = append(paths, "resource")
	}
	if !reflect.DeepEqual(a.desired.BackupPlan, actual.BackupPlan) {
		report.AddField("backup_plan", actual.BackupPlan, a.desired.BackupPlan)
		paths = append(paths, "backup_plan")
	}

	if len(paths) != 0 {
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("BackupDRBackupPlanAssociation is immutable and cannot be updated, fields: %v", paths)
	}

	// still need to update status (in the event of acquiring an existing resource)
	mapCtx := &direct.MapContext{}
	status := &krm.BackupDRBackupPlanAssociationStatus{}
	status.ObservedState = BackupDRBackupPlanAssociationObservedState_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// getComputeInstanceID returns the numeric ID of a Compute Engine instance.
func (a *BackupPlanAssociationAdapter) getComputeInstanceID(ctx context.Context, id *computev1beta1.ComputeInstanceIdentity) (uint64, error) {
	req := &computepb.GetInstanceRequest{
		Project:  id.Project,
		Zone:     id.Zone,
		Instance: id.Instance,
	}
	instance, err := a.computeClient.Get(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("getting ComputeInstance %q: %w", id, err)
	}
	return instance.GetId(), nil
}

// normalizeActualReferences returns a copy of actual where the backup plan and the resource are
// replaced with the desired values when they refer to the same GCP resources.
// The API returns these references with the project number instead of the project ID,
// and Compute Engine instances with their numeric ID instead of their name.
func normalizeActualReferences(ctx context.Context, projectMapper *projects.ProjectMapper, getComputeInstanceID func(context.Context, *computev1beta1.ComputeInstanceIdentity) (uint64, error), desired, actual *pb.BackupPlanAssociation) (*pb.BackupPlanAssociation, error) {
	normalized := proto.CloneOf(actual)

	if desired.GetBackupPlan() != actual.GetBackupPlan() {
		same, err := isSameResourceName(ctx, projectMapper, desired.GetBackupPlan(), actual.GetBackupPlan())
		if err != nil {
			return nil, fmt.Errorf("comparing backup plan: %w", err)
		}
		if same {
			normalized.BackupPlan = desired.GetBackupPlan()
		}
	}

	if desired.GetResource() != actual.GetResource() {
		var same bool
		var err error
		if actual.GetResourceType() == krm.ResourceType_ComputeInstance {
			same, err = isSameComputeInstance(ctx, projectMapper, getComputeInstanceID, desired.GetResource(), actual.GetResource())
		} else {
			same, err = isSameResourceName(ctx, projectMapper, desired.GetResource(), actual.GetResource())
		}
		if err != nil {
			return nil, fmt.Errorf("comparing resource: %w", err)
		}
		if same {
			normalized.Resource = desired.GetResource()
		}
	}

	return normalized, nil
}

// isSameResourceName returns true if both resource names are equal once project numbers are replaced with project IDs.
func isSameResourceName(ctx context.Context, projectMapper *projects.ProjectMapper, desired, actual string) (bool, error) {
	normalizedDesired, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, desired)
	if err != nil {
		return false, err
	}
	normalizedActual, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, actual)
	if err != nil {
		return false, err
	}
	return normalizedDesired == normalizedActual, nil
}

// isSameComputeInstance returns true if both names refer to the same Compute Engine instance.
// The project can be identified by its ID or number, and the instance by its name or numeric ID.
func isSameComputeInstance(ctx context.Context, projectMapper *projects.ProjectMapper, getComputeInstanceID func(context.Context, *computev1beta1.ComputeInstanceIdentity) (uint64, error), desired, actual string) (bool, error) {
	desiredID, err := computev1beta1.ParseComputeInstanceExternal(desired)
	if err != nil {
		// Not a known format, we cannot prove that this is the same instance.
		return false, nil
	}
	actualID, err := computev1beta1.ParseComputeInstanceExternal(actual)
	if err != nil {
		return false, nil
	}

	desiredProject, err := projectMapper.ReplaceProjectNumberWithID(ctx, desiredID.Project)
	if err != nil {
		return false, err
	}
	actualProject, err := projectMapper.ReplaceProjectNumberWithID(ctx, actualID.Project)
	if err != nil {
		return false, err
	}
	if desiredProject != actualProject || desiredID.Zone != actualID.Zone {
		return false, nil
	}
	if desiredID.Instance == actualID.Instance {
		return true, nil
	}

	// Instance names start with a letter, so a numeric value is an instance ID.
	if _, err := strconv.ParseUint(actualID.Instance, 10, 64); err != nil {
		return false, nil
	}
	instanceID, err := getComputeInstanceID(ctx, desiredID)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return strconv.FormatUint(instanceID, 10) == actualID.Instance, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupPlanAssociationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BackupDRBackupPlanAssociation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupDRBackupPlanAssociationSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BackupDRBackupPlanAssociationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupPlanAssociationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupPlanAssociation", "name", a.id)

	req := &pb.DeleteBackupPlanAssociationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupPlanAssociation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BackupPlanAssociation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupPlanAssociation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackupPlanAssociation", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupPlanAssociation %s: %w", a.id, err)
	}
	return true, nil
}
