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
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.Enrollment
// crd.type: EventarcEnrollment
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.EventarcEnrollmentGVK, NewEnrollmentModel)
}

func NewEnrollmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &enrollmentModel{config: *config}, nil
}

var _ directbase.Model = &enrollmentModel{}

type enrollmentModel struct {
	config config.ControllerConfig
}

func (m *enrollmentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcEnrollment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get eventarc GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := EventarcEnrollmentSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &enrollmentAdapter{
		gcpClient: eventarcClient,
		id:        identity.(*krm.EventarcEnrollmentIdentity),
		desired:   desiredPb,
		reader:    reader,
	}, nil
}

func (m *enrollmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.EventarcEnrollmentIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	return &enrollmentAdapter{
		gcpClient: eventarcClient,
		id:        id,
	}, nil
}

type enrollmentAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcEnrollmentIdentity
	desired   *pb.Enrollment
	actual    *pb.Enrollment
	reader    client.Reader
}

var _ directbase.Adapter = &enrollmentAdapter{}

func (a *enrollmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc enrollment", "name", a.id.String())

	req := &pb.GetEnrollmentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEnrollment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc enrollment %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *enrollmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc enrollment", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.Enrollment)
	desired.Name = a.id.String()

	req := &pb.CreateEnrollmentRequest{
		Parent:       fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Enrollment:   desired,
		EnrollmentId: a.id.Enrollment,
	}

	op, err := a.gcpClient.CreateEnrollment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc enrollment %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc enrollment creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc enrollment", "name", a.id.String())

	// Fetch fully-populated resource after LRO completion
	latest, err := a.gcpClient.GetEnrollment(ctx, &pb.GetEnrollmentRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting eventarc enrollment after creation %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *enrollmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc enrollment", "name", a.id.String())

	diffs, updateMask, err := compareEventarcEnrollment(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*pb.Enrollment)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateEnrollmentRequest{
			Enrollment: desiredCopy,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateEnrollment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating eventarc enrollment %s: %w", a.id.String(), err)
		}

		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for eventarc enrollment update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated eventarc enrollment", "name", a.id.String())

		latest, err = a.gcpClient.GetEnrollment(ctx, &pb.GetEnrollmentRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting eventarc enrollment after update %s: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *enrollmentAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Enrollment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.EventarcEnrollmentStatus{}
	status.ObservedState = EventarcEnrollmentObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *enrollmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcEnrollment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *EventarcEnrollmentSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Enrollment)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Enrollment)
	u.SetNamespace(obj.Namespace) // This is required KCC controller code convention
	u.SetGroupVersionKind(krm.EventarcEnrollmentGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *enrollmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc enrollment", "name", a.id.String())

	req := &pb.DeleteEnrollmentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEnrollment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc enrollment not found", "name", a.id.String())
			return true, nil // Resource is gone, consider the delete successful.
		}
		return false, fmt.Errorf("deleting eventarc enrollment %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc enrollment deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc enrollment", "name", a.id.String())
	return true, nil
}

func compareEventarcEnrollment(ctx context.Context, actual, desired *pb.Enrollment) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, EventarcEnrollmentSpec_FromProto, EventarcEnrollmentSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
