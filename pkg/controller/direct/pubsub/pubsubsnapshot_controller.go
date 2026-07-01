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
// proto.service: google.pubsub.v1.Subscriber
// proto.message: google.pubsub.v1.Snapshot
// crd.type: PubSubSnapshot
// crd.version: v1beta1

package pubsub

import (
	"context"
	"fmt"
	"sort"

	api "cloud.google.com/go/pubsub/v2/apiv1"
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.PubSubSnapshotGVK, NewSnapshotModel)
}

func NewSnapshotModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &SnapshotModel{config: *config}, nil
}

var _ directbase.Model = &SnapshotModel{}

type SnapshotModel struct {
	config config.ControllerConfig
}

func (m *SnapshotModel) client(ctx context.Context, projectID string) (*api.SubscriptionAdminClient, error) {
	var opts []option.ClientOption

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewSubscriptionAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building pubsub snapshot client: %w", err)
	}

	return gcpClient, err
}

func (m *SnapshotModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.PubSubSnapshot{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	snapshotId := id.(*krm.PubSubSnapshotIdentity)

	gcpClient, err := m.client(ctx, snapshotId.Project)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := PubSubSnapshotSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = snapshotId.String()
	desired.Labels = label.GCPLabels(obj)

	// pubSubSubscriptionRef is required when creating a snapshot, but it is not part
	// of the standard pb.Snapshot proto structure returned from GET or updated via PATCH.
	// We track it separately here so that Create() can use it.
	var pubSubSubscriptionRefExternal string
	if obj.Spec.PubSubSubscriptionRef != nil {
		pubSubSubscriptionRefExternal = obj.Spec.PubSubSubscriptionRef.External
	}

	return &snapshotAdapter{
		gcpClient:                    gcpClient,
		id:                           snapshotId,
		desired:                      desired,
		desiredPubSubSubscriptionRef: pubSubSubscriptionRefExternal,
	}, nil
}

func (m *SnapshotModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.PubSubSnapshotIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &snapshotAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type snapshotAdapter struct {
	gcpClient                    *api.SubscriptionAdminClient
	id                           *krm.PubSubSnapshotIdentity
	desired                      *pb.Snapshot
	desiredPubSubSubscriptionRef string
	actual                       *pb.Snapshot
}

var _ directbase.Adapter = &snapshotAdapter{}

func (a *snapshotAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting pubsub snapshot", "name", a.id)

	req := &pb.GetSnapshotRequest{Snapshot: a.id.String()}
	actual, err := a.gcpClient.GetSnapshot(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting pubsub snapshot %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *snapshotAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating pubsub snapshot", "name", a.id)

	req := &pb.CreateSnapshotRequest{
		Name:         a.desired.Name,
		Subscription: a.desiredPubSubSubscriptionRef,
	}
	if a.desired.Labels != nil {
		req.Labels = a.desired.Labels
	}
	created, err := a.gcpClient.CreateSnapshot(ctx, req)
	if err != nil {
		return fmt.Errorf("creating pubsub snapshot %s: %w", a.id.String(), err)
	}
	log.Info("successfully created pubsub snapshot in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *snapshotAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating pubsub snapshot", "name", a.id)

	paths, diffs, updateMask, err := compareSnapshot(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	// Topic is immutable. If there's a diff on "topic", return an error.
	if paths.Has("topic") {
		return fmt.Errorf("topic is immutable and cannot be updated")
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for pubsub snapshot", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSnapshotRequest{
		Snapshot:   a.desired,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateSnapshot(ctx, req)
	if err != nil {
		return fmt.Errorf("updating pubsub snapshot %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated pubsub snapshot", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *snapshotAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting pubsub snapshot", "name", a.id)

	req := &pb.DeleteSnapshotRequest{Snapshot: a.id.String()}
	err := a.gcpClient.DeleteSnapshot(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting pubsub snapshot %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted pubsub snapshot", "name", a.id)

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *snapshotAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PubSubSnapshot{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *PubSubSnapshotSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Note: The exported object does not have the subscription reference (pubSubSubscriptionRef)
	// because the GCP GET/LIST APIs do not return it. Consequently, this exported object
	// cannot be used directly to create a new snapshot without manually specifying that field.
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Snapshot)
	obj.Spec.ProjectRef = &refv1beta1.ProjectRef{
		External: a.id.Project,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Snapshot)
	u.SetGroupVersionKind(krm.PubSubSnapshotGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func compareSnapshot(ctx context.Context, actual, desired *pb.Snapshot) (sets.Set[string], *structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PubSubSnapshotSpec_FromProto, PubSubSnapshotSpec_ToProto)
	if err != nil {
		return nil, nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.Labels = actual.Labels

	paths, diffs, err := common.CompareProtoMessageStructuredDiff(desired, maskedActual, common.BasicDiff)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("comparing spec: %w", err)
	}

	pathsList := paths.UnsortedList()
	sort.Strings(pathsList)
	updateMask := &fieldmaskpb.FieldMask{
		Paths: pathsList,
	}
	return paths, diffs, updateMask, nil
}

func (a *snapshotAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Snapshot) error {
	status := &krm.PubSubSnapshotStatus{}
	status.ExternalRef = direct.LazyPtr(latest.Name)

	mapCtx := &direct.MapContext{}
	status.ObservedState = PubSubSnapshotObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
