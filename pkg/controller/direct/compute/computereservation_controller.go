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
// proto.service: google.cloud.compute.v1.Reservations
// proto.message: google.cloud.compute.v1.Reservation
// crd.type: ComputeReservation
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeReservationGVK, NewReservationModel)
}

func NewReservationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &reservationModel{config: config}, nil
}

var _ directbase.Model = &reservationModel{}

type reservationModel struct {
	config *config.ControllerConfig
}

func (m *reservationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeReservation{}
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
	reservationsClient, err := gcpClient.newReservationsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ReservationAdapter{
		gcpClient:     reservationsClient,
		id:            id.(*krm.ComputeReservationIdentity),
		desired:       obj,
		reader:        reader,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *reservationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ReservationAdapter struct {
	gcpClient     *compute.ReservationsClient
	id            *krm.ComputeReservationIdentity
	desired       *krm.ComputeReservation
	actual        *computepb.Reservation
	reader        client.Reader
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &ReservationAdapter{}

// Find retrieves the GCP resource.
func (a *ReservationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Reservation", "name", a.id)

	req := &computepb.GetReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Reservation %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReservationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Reservation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	if err := ResolveComputeReservationRefs(ctx, a.reader, a.projectMapper, desired); err != nil {
		return err
	}
	resource := ComputeReservationSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Reservation)
	// match realGCP log: "zone": "projects/${projectId}/global/zones/us-central1-a"
	resource.Zone = direct.LazyPtr(fmt.Sprintf("projects/%s/global/zones/%s", a.id.Project, a.id.Zone))

	req := &computepb.InsertReservationRequest{
		Project:             a.id.Project,
		Zone:                a.id.Zone,
		ReservationResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Reservation %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute Reservation %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute Reservation in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting Reservation %s: %w", a.id, err)
	}

	status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ReservationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Reservation", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	if err := ResolveComputeReservationRefs(ctx, a.reader, a.projectMapper, desired); err != nil {
		return err
	}
	resource := ComputeReservationSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report, updateMask, err := compareComputeReservation(ctx, &desired.Spec, a.actual)
	if err != nil {
		return err
	}

	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	if slices.Contains(updateMask.GetPaths(), "specific_reservation") {
		req := &computepb.ResizeReservationRequest{
			Project:     a.id.Project,
			Zone:        a.id.Zone,
			Reservation: a.id.Reservation,
			ReservationsResizeRequestResource: &computepb.ReservationsResizeRequest{
				SpecificSkuCount: direct.PtrTo(int64(resource.GetSpecificReservation().GetCount())),
			},
		}
		op, err := a.gcpClient.Resize(ctx, req)
		if err != nil {
			return fmt.Errorf("resizing compute Reservation %s: %w", a.id.String(), err)
		}

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for resize of compute Reservation %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully resized compute Reservation", "name", a.id)
	}

	if slices.Contains(updateMask.GetPaths(), "share_settings") {
		actualShareSettings := a.actual.GetShareSettings()
		desiredShareSettings := resource.GetShareSettings()

		actualProjectMap := actualShareSettings.GetProjectMap()
		desiredProjectMap := desiredShareSettings.GetProjectMap()

		var addProjectPaths []string
		addedProjectMap := make(map[string]*computepb.ShareSettingsProjectConfig)
		for k, v := range desiredProjectMap {
			actualVal, ok := actualProjectMap[k]
			if !ok || actualVal.GetProjectId() != v.GetProjectId() {
				addProjectPaths = append(addProjectPaths, fmt.Sprintf("shareSettings.projectMap.%s", k))
				addedProjectMap[k] = v
			}
		}

		var removeProjectPaths []string
		for k := range actualProjectMap {
			if _, ok := desiredProjectMap[k]; !ok {
				removeProjectPaths = append(removeProjectPaths, fmt.Sprintf("shareSettings.projectMap.%s", k))
			}
		}

		// Perform additions first.
		// Cannot remove the single remaining shared project first for change. The API requires at least one shared project at all times.
		if len(addProjectPaths) > 0 {
			sort.Strings(addProjectPaths)
			req := &computepb.UpdateReservationRequest{
				Project:     a.id.Project,
				Zone:        a.id.Zone,
				Reservation: a.id.Reservation,
				Paths:       direct.LazyPtr(strings.Join(addProjectPaths, ",")),
				ReservationResource: &computepb.Reservation{
					Name: direct.LazyPtr(a.id.Reservation),
					ShareSettings: &computepb.ShareSettings{
						ShareType:  actualShareSettings.ShareType,
						ProjectMap: addedProjectMap,
					},
				},
			}

			op, err := a.gcpClient.Update(ctx, req)
			if err != nil {
				return fmt.Errorf("adding projects to shareSettings for compute Reservation %s: %w", a.id.String(), err)
			}
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for add projects to shareSettings of compute Reservation %s: %w", a.id.String(), err)
			}
			log.V(2).Info("successfully added projects to shareSettings of compute Reservation", "name", a.id)
		}

		// Perform removals next.
		if len(removeProjectPaths) > 0 {
			sort.Strings(removeProjectPaths)
			req := &computepb.UpdateReservationRequest{
				Project:     a.id.Project,
				Zone:        a.id.Zone,
				Reservation: a.id.Reservation,
				Paths:       direct.LazyPtr(strings.Join(removeProjectPaths, ",")),
				ReservationResource: &computepb.Reservation{
					Name: direct.LazyPtr(a.id.Reservation),
				},
			}

			op, err := a.gcpClient.Update(ctx, req)
			if err != nil {
				return fmt.Errorf("removing projects from shareSettings for compute Reservation %s: %w", a.id.String(), err)
			}
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for remove projects from shareSettings of compute Reservation %s: %w", a.id.String(), err)
			}
			log.V(2).Info("successfully removed projects from shareSettings of compute Reservation", "name", a.id)
		}
	}

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeReservation %s: %w", a.id, err)
	}

	status := ComputeReservationStatus_v1beta1_FromProto(mapCtx, updated)
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReservationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeReservation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *ComputeReservationSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// Note: project is not explicitly in spec but inferred from namespace, we set it if needed
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Reservation)
	u.SetGroupVersionKind(krm.ComputeReservationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReservationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Reservation", "name", a.id)

	req := &computepb.DeleteReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent compute Reservation, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting compute Reservation %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute Reservation %s: %w", a.id.String(), err)
		}
	}
	log.Info("successfully deleted compute Reservation", "name", a.id)

	return true, nil
}

func (a *ReservationAdapter) get(ctx context.Context) (*computepb.Reservation, error) {
	getReq := &computepb.GetReservationRequest{
		Project:     a.id.Project,
		Zone:        a.id.Zone,
		Reservation: a.id.Reservation,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeReservation %s: %w", a.id, err)
	}
	return resource, nil
}

func compareComputeReservation(ctx context.Context, desiredKRM *krm.ComputeReservationSpec, actual *computepb.Reservation) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	report, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		desiredKRM,
		actual,
		ComputeReservationSpec_v1beta1_FromProto,
		ComputeReservationSpec_v1beta1_ToProto,
		normalizeReservationProto,
	)
	if err != nil {
		return nil, nil, err
	}

	// Enforce immutability: only specific_reservation.count and share_settings.project_map can be updated.
	for _, field := range report.Fields {
		switch field.ID {
		case "specific_reservation":
			actualSR := protoMessageFromAny[*computepb.AllocationSpecificSKUReservation](field.Old)
			desiredSR := protoMessageFromAny[*computepb.AllocationSpecificSKUReservation](field.New)
			if !proto.Equal(desiredSR.GetInstanceProperties(), actualSR.GetInstanceProperties()) {
				subPaths, _, err := common.CompareProtoMessageStructuredDiff(
					desiredSR.GetInstanceProperties(),
					actualSR.GetInstanceProperties(),
					common.BasicDiff,
				)
				if err != nil {
					return nil, nil, err
				}
				for _, subPath := range sets.List(subPaths) {
					return nil, nil, fmt.Errorf("field %q is immutable", "specific_reservation.instance_properties."+subPath)
				}
				return nil, nil, fmt.Errorf("field %q is immutable", "specific_reservation.instance_properties")
			}
		case "share_settings":
			actualSS := protoMessageFromAny[*computepb.ShareSettings](field.Old)
			desiredSS := protoMessageFromAny[*computepb.ShareSettings](field.New)
			if actualSS.GetShareType() != desiredSS.GetShareType() {
				return nil, nil, fmt.Errorf("field %q is immutable", "share_settings.share_type")
			}
		default:
			return nil, nil, fmt.Errorf("field %q is immutable", field.ID)
		}
	}

	return report, updateMask, nil
}

func protoMessageFromAny[T proto.Message](v any) T {
	var zero T
	if pm, ok := v.(protoreflect.Message); ok {
		if msg, ok := pm.Interface().(T); ok {
			return msg
		}
	}
	if msg, ok := v.(T); ok {
		return msg
	}
	return zero
}

func normalizeReservationProto(_ context.Context, pb *computepb.Reservation) error {
	if pb.Zone != nil {
		zone := pb.GetZone()
		if idx := strings.LastIndex(zone, "/"); idx != -1 {
			pb.Zone = direct.LazyPtr(zone[idx+1:])
		}
	}
	if pb.ShareSettings == nil {
		pb.ShareSettings = &computepb.ShareSettings{
			ShareType: direct.LazyPtr("LOCAL"),
		}
	} else if pb.GetShareSettings().GetShareType() == "" {
		pb.ShareSettings.ShareType = direct.LazyPtr("LOCAL")
	}
	if pb.SpecificReservation != nil {
		pb.SpecificReservation.InUseCount = nil
	}
	return nil
}
