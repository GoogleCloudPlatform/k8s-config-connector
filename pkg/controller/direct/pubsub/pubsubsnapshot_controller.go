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
// crd.version: v1alpha1

package pubsub

import (
	"context"
	"fmt"

	api "cloud.google.com/go/pubsub/apiv1"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1alpha1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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

func (m *SnapshotModel) client(ctx context.Context, projectID string) (*api.SubscriberClient, error) {
	var opts []option.ClientOption

	config := m.config

	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewSubscriberRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building pubsub snapshot client: %w", err)
	}

	return gcpClient, err
}

func (m *SnapshotModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.PubSubSnapshot{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSnapshotIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// resolve subscription
	if obj.Spec.PubSubSubscriptionRef != nil {
		subscription, err := refsv1beta1.ResolvePubSubSubscription(ctx, reader, obj, obj.Spec.PubSubSubscriptionRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.PubSubSubscriptionRef.External = subscription.String()
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &snapshotAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *SnapshotModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type snapshotAdapter struct {
	gcpClient *api.SubscriberClient
	id        *v1alpha1.SnapshotIdentity
	desired   *krm.PubSubSnapshot
	actual    *pb.Snapshot
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

	desired := a.desired.DeepCopy()
	desired.Name = a.id.String()

	req := &pb.CreateSnapshotRequest{
		Name:         desired.Name,
		Subscription: desired.Spec.PubSubSubscriptionRef.External,
	}
	if desired.Labels != nil {
		req.Labels = desired.Labels
	}
	_, err := a.gcpClient.CreateSnapshot(ctx, req)
	if err != nil {
		return fmt.Errorf("creating pubsub snapshot %s: %w", a.id.String(), err)
	}
	log.Info("successfully created pubsub snapshot in gcp", "name", a.id)

	status := &krm.PubSubSnapshotStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// PubSubSnapshot does not support update.
func (a *snapshotAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return fmt.Errorf("update pubsub snapshot is not supported")
}

func (a *snapshotAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting pubsub snapshot", "name", a.id)

	req := &pb.DeleteSnapshotRequest{Snapshot: a.id.String()}
	err := a.gcpClient.DeleteSnapshot(ctx, req)
	if err != nil {
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
	obj.Spec = direct.ValueOf(PubSubSnapshotSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.PubSubSnapshotGVK)

	u.Object = uObj
	return u, nil
}
