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
// proto.service: google.cloud.compute.v1.FutureReservations
// proto.message: google.cloud.compute.v1.FutureReservation
// crd.type: ComputeFutureReservation
// crd.version: v1alpha1

package compute

import (
	"reflect"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"context"
	"fmt"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeFutureReservationGVK, NewFutureReservationModel)
}

func NewFutureReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &futureReservationModel{config: config}, nil
}

var _ directbase.Model = &futureReservationModel{}

type futureReservationModel struct {
	config *config.ControllerConfig
}

func (m *futureReservationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeFutureReservation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	futureReservationClient, err := gcpClient.newFutureReservationsClient(ctx)
	if err != nil {
		return nil, err
	}

	desired := obj.DeepCopy()
	if err := ResolveComputeFutureReservationRefs(ctx, reader, m.config.ProjectMapper, desired); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeFutureReservationSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredProto.Name = direct.LazyPtr(id.(*krm.ComputeFutureReservationIdentity).FutureReservation)

	return &FutureReservationAdapter{
		gcpClient: futureReservationClient,
		id:        id.(*krm.ComputeFutureReservationIdentity),
		desired:   desiredProto,
	}, nil
}

func (m *futureReservationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FutureReservationAdapter struct {
	gcpClient *compute.FutureReservationsClient
	id        *v1alpha1.ComputeFutureReservationIdentity
	desired   *computepb.FutureReservation
	actual    *computepb.FutureReservation
}

var _ directbase.Adapter = &FutureReservationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FutureReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FutureReservation", "name", a.id)

	req := &computepb.GetFutureReservationRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		FutureReservation: a.id.FutureReservation,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FutureReservation %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := proto.CloneOf(a.desired)

	req := &computepb.InsertFutureReservationRequest{
		Project:                   a.id.Project,
		Zone:                      a.id.Zone,
		FutureReservationResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FutureReservation %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute FutureReservation %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute FutureReservation in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting FutureReservation %s: %w", a.id, err)
	}

	status := &krm.ComputeFutureReservationStatus{}
	status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FutureReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FutureReservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := proto.CloneOf(a.desired)

	updateMask := fieldmaskpb.FieldMask{}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// We need to check if AutoCreatedReservationsDeleteTime is set in actual as the GCP API.
	// GET operation does not populate AutoDeleteAutoCreatedReservations.
	if resource.GetAutoDeleteAutoCreatedReservations() == false && a.actual.AutoCreatedReservationsDeleteTime != nil {
		report.AddField("auto_delete_auto_created_reservations", "true", resource.GetAutoDeleteAutoCreatedReservations())
		updateMask.Paths = append(updateMask.Paths, "auto_delete_auto_created_reservations")
	}

	if resource.GetAutoDeleteAutoCreatedReservations() == true && a.actual.AutoCreatedReservationsDeleteTime == nil {
		report.AddField("auto_delete_auto_created_reservations", "false", resource.GetAutoDeleteAutoCreatedReservations())
		updateMask.Paths = append(updateMask.Paths, "auto_delete_auto_created_reservations")
	}

	// handle AutoCreatedReservationsDeleteTime calculated by duration
	if resource.AutoCreatedReservationsDuration != nil {
		end, err := calculateTimeFromDuration(resource.TimeWindow.StartTime, resource.AutoCreatedReservationsDuration)
		if err != nil {
			return err
		}
		if a.actual.AutoCreatedReservationsDeleteTime == nil {
			report.AddField("auto_created_reservations_duration", nil, resource.AutoCreatedReservationsDuration)
			updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_duration")
		} else if actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetAutoCreatedReservationsDeleteTime()); err == nil {
			calculatedDelete := end
			if !calculatedDelete.Equal(actualDelete) {
				report.AddField("auto_created_reservations_duration", a.actual.AutoCreatedReservationsDeleteTime, end)
				updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_duration")
			}
		} else {
			return fmt.Errorf("invalid auto_created_reservations_delete_time: %w", err)
		}
	}

	// compare AutoCreatedReservationsDeleteTime if exists in desired configuration
	if resource.AutoCreatedReservationsDeleteTime != nil {
		if isEqual, err := compareTimestamp(a.actual.GetAutoCreatedReservationsDeleteTime(), resource.GetAutoCreatedReservationsDeleteTime()); err != nil {
			return fmt.Errorf("invalid auto_created_reservations_delete_time: %w", err)
		} else if !isEqual {
			report.AddField("auto_created_reservations_delete_time", a.actual.AutoCreatedReservationsDeleteTime, resource.AutoCreatedReservationsDeleteTime)
			updateMask.Paths = append(updateMask.Paths, "auto_created_reservations_delete_time")
		}
	}
	if !reflect.DeepEqual(resource.TimeWindow, a.actual.TimeWindow) {
		// compare start time,  start time must be specified
		if isEqual, err := compareTimestamp(a.actual.GetTimeWindow().GetStartTime(), resource.GetTimeWindow().GetStartTime()); err != nil {
			return fmt.Errorf("invalid time_window.start_time: %w", err)
		} else if !isEqual {
			report.AddField("time_window.start_time", a.actual.GetTimeWindow().GetStartTime(), resource.GetTimeWindow().GetStartTime())
			updateMask.Paths = append(updateMask.Paths, "time_window.start_time")
		}
		// compare endtime if exists in desired configuration
		if resource.TimeWindow.EndTime != nil {
			if isEqual, err := compareTimestamp(a.actual.GetTimeWindow().GetEndTime(), resource.GetTimeWindow().GetEndTime()); err != nil {
				return fmt.Errorf("invalid time_window.end_time: %w", err)
			} else if !isEqual {
				report.AddField("time_window.end_time", a.actual.GetTimeWindow().GetEndTime(), resource.GetTimeWindow().GetEndTime())
				updateMask.Paths = append(updateMask.Paths, "time_window.end_time")
			}
		}
	}

	// handle endtime calculated by duration
	if resource.TimeWindow != nil && resource.TimeWindow.Duration != nil {
		end, err := calculateTimeFromDuration(resource.TimeWindow.StartTime, resource.TimeWindow.Duration)
		if err != nil {
			return err
		}
		if a.actual.TimeWindow.EndTime == nil {
			report.AddField("time_window.duration", nil, resource.TimeWindow.Duration)
			updateMask.Paths = append(updateMask.Paths, "time_window.duration")
		} else if actualDelete, err := time.Parse(time.RFC3339Nano, a.actual.GetTimeWindow().GetEndTime()); err == nil {
			calculatedDelete := end
			if !calculatedDelete.Equal(actualDelete) {
				report.AddField("time_window.duration", a.actual.TimeWindow.EndTime, end)
				updateMask.Paths = append(updateMask.Paths, "time_window.duration")
			}
		} else {
			return fmt.Errorf("invalid time_window.end_time: %w", err)
		}
	}

	// If the description is nil, after update operation GCP API converts it to "" (see http logs in fixtures tests).
	// Compare their values to avoid unnecessary diff.
	if resource.GetDescription() != a.actual.GetDescription() {
		report.AddField("description", a.actual.GetDescription(), resource.GetDescription())
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if desiredProperties, actualProperties := resource.GetSpecificSkuProperties(), a.actual.GetSpecificSkuProperties(); desiredProperties != nil && actualProperties != nil {
		if !reflect.DeepEqual(desiredProperties.TotalCount, actualProperties.TotalCount) {
			report.AddField("specific_sku_properties.total_count", actualProperties.TotalCount, desiredProperties.TotalCount)
			updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.total_count")
		}
		if ip, actualIP := desiredProperties.GetInstanceProperties(), actualProperties.GetInstanceProperties(); ip != nil && actualIP != nil {
			if !reflect.DeepEqual(ip.MachineType, actualIP.MachineType) {
				report.AddField("specific_sku_properties.instance_properties.machine_type", actualIP.MachineType, ip.MachineType)
				updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.instance_properties.machine_type")
			}
			if !reflect.DeepEqual(ip.MinCpuPlatform, actualIP.MinCpuPlatform) {
				report.AddField("specific_sku_properties.instance_properties.min_cpu_platform", actualIP.MinCpuPlatform, ip.MinCpuPlatform)
				updateMask.Paths = append(updateMask.Paths, "specific_sku_properties.instance_properties.min_cpu_platform")
			}
		}
	}

	desiredShareSettings := resource.GetShareSettings()
	if desiredShareSettings != nil && desiredShareSettings.GetShareType() == "LOCAL" {
		desiredShareSettings = nil
	}

	if !proto.Equal(desiredShareSettings, a.actual.GetShareSettings()) {
		report.AddField("share_settings", a.actual.GetShareSettings(), resource.GetShareSettings())
		updateMask.Paths = append(updateMask.Paths, "share_settings")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ComputeFutureReservationStatus{}
		status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	sort.Strings(updateMask.Paths)

	req := &computepb.UpdateFutureReservationRequest{
		Project:                   a.id.Project,
		Zone:                      a.id.Zone,
		FutureReservation:         a.id.FutureReservation,
		FutureReservationResource: resource,
		UpdateMask:                direct.PtrTo(strings.Join(updateMask.Paths, ",")),
	}
	op, err := a.gcpClient.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("updating compute FutureReservation %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of compute FutureReservation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated compute FutureReservation", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeFutureReservation %s: %w", a.id, err)
	}

	status := &krm.ComputeFutureReservationStatus{}
	status.ObservedState = ComputeFutureReservationObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FutureReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeFutureReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeFutureReservationSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Zone
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeFutureReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FutureReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FutureReservation", "name", a.id)

	req := &computepb.DeleteFutureReservationRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		FutureReservation: a.id.FutureReservation,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute FutureReservation %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute FutureReservation", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute FutureReservation %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *FutureReservationAdapter) get(ctx context.Context) (*computepb.FutureReservation, error) {
	getReq := &computepb.GetFutureReservationRequest{
		Project:           a.id.Project,
		Zone:              a.id.Zone,
		FutureReservation: a.id.FutureReservation,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeFutureReservation %s: %w", a.id, err)
	}
	return resource, nil
}

func calculateTimeFromDuration(startTime *string, duration *computepb.Duration) (time.Time, error) {
	start, err := time.Parse(time.RFC3339Nano, direct.ValueOf(startTime))
	if err != nil {
		return time.Time{}, status.Errorf(codes.InvalidArgument, "invalid start_time: %v", err)
	}

	d := time.Duration(0)
	if duration.Seconds != nil {
		d = d + time.Duration(direct.ValueOf(duration.Seconds))*time.Second
	}
	if duration.Nanos != nil {
		d = d + time.Duration(direct.ValueOf(duration.Nanos))*time.Nanosecond
	}

	return start.Add(d), nil
}

func compareTimestamp(actual, desired string) (bool, error) {
	// GCP API timestamp requirements:
	// Must be a string representation of a full-date or date-time in valid RFC3339 format with either 0 or 3 digits for fractional seconds
	actualTime, err := time.Parse(time.RFC3339Nano, actual)
	if err != nil {
		return false, err
	}
	desiredTime, err := time.Parse(time.RFC3339Nano, desired)
	if err != nil {
		return false, err
	}
	return actualTime.Equal(desiredTime), nil
}
