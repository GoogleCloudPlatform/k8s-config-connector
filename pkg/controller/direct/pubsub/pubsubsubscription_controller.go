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
// proto.service: google.pubsub.v1.Subscriber
// proto.message: google.pubsub.v1.Subscription
// crd.type: PubSubSubscription
// crd.version: v1beta1

package pubsub

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "cloud.google.com/go/pubsub/v2/apiv1"
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.PubSubSubscriptionGVK, NewSubscriptionModel)
}

func NewSubscriptionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &SubscriptionModel{config: *config}, nil
}

var _ directbase.Model = &SubscriptionModel{}

type SubscriptionModel struct {
	config config.ControllerConfig
}

func (m *SubscriptionModel) client(ctx context.Context, projectID string) (*api.SubscriptionAdminClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewSubscriptionAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building pubsub subscriber client: %w", err)
	}

	return gcpClient, err
}

func (m *SubscriptionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.PubSubSubscription{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Support backward compatibility for legacy short-form topic external references.
	// We dynamically qualify the short name with the subscription's own project context
	// before invoking standard reference normalization.
	if obj.Spec.TopicRef != nil && obj.Spec.TopicRef.External != "" && !strings.Contains(obj.Spec.TopicRef.External, "/") {
		obj.Spec.TopicRef.External = "projects/" + projectID + "/topics/" + obj.Spec.TopicRef.External
	}

	// Support backward compatibility for legacy short-form dead letter topic external references.
	// We dynamically qualify the short name with the subscription's own project context
	// before invoking standard reference normalization.
	if obj.Spec.DeadLetterPolicy != nil && obj.Spec.DeadLetterPolicy.DeadLetterTopicRef != nil && obj.Spec.DeadLetterPolicy.DeadLetterTopicRef.External != "" && !strings.Contains(obj.Spec.DeadLetterPolicy.DeadLetterTopicRef.External, "/") {
		obj.Spec.DeadLetterPolicy.DeadLetterTopicRef.External = "projects/" + projectID + "/topics/" + obj.Spec.DeadLetterPolicy.DeadLetterTopicRef.External
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	subscriptionId := id.(*krm.PubSubSubscriptionIdentity)

	gcpClient, err := m.client(ctx, subscriptionId.Project)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := PubSubSubscriptionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(obj)

	return &pubSubSubscriptionAdapter{
		gcpClient: gcpClient,
		id:        subscriptionId,
		desired:   desired,
	}, nil
}

func (m *SubscriptionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type pubSubSubscriptionAdapter struct {
	gcpClient *api.SubscriptionAdminClient
	id        *krm.PubSubSubscriptionIdentity
	desired   *pb.Subscription
	actual    *pb.Subscription
}

var _ directbase.Adapter = &pubSubSubscriptionAdapter{}

func (a *pubSubSubscriptionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting pubsub subscription", "name", a.id)

	req := &pb.GetSubscriptionRequest{Subscription: a.id.String()}
	actual, err := a.gcpClient.GetSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting pubsub subscription %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *pubSubSubscriptionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating pubsub subscription", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	created, err := a.gcpClient.CreateSubscription(ctx, desired)
	if err != nil {
		return fmt.Errorf("creating pubsub subscription %s: %w", a.id.String(), err)
	}
	log.Info("successfully created pubsub subscription in gcp", "name", a.id)

	// CreateSubscription is eventually consistent. Poll GetSubscription for a sensible amount of time
	// (up to 2 minutes total, with backoff starting at 1 second and growing linearly by 1 second up to 10 seconds).
	pollCtx, pollCancel := context.WithTimeout(ctx, 2*time.Minute)
	defer pollCancel()

	backoff := 1 * time.Second
	for {
		req := &pb.GetSubscriptionRequest{Subscription: a.id.String()}
		_, err := a.gcpClient.GetSubscription(pollCtx, req)
		if err == nil {
			break
		}
		if direct.IsNotFound(err) {
			log.V(2).Info("subscription not found yet during eventual consistency check, retrying", "name", a.id, "backoff", backoff)
		} else {
			log.V(2).Info("error getting subscription during eventual consistency check, retrying", "name", a.id, "error", err, "backoff", backoff)
		}

		select {
		case <-pollCtx.Done():
			return fmt.Errorf("subscription %s created but timed out waiting for it to become visible: %w", a.id.String(), pollCtx.Err())
		case <-time.After(backoff):
		}

		backoff += 1 * time.Second
		if backoff > 10*time.Second {
			backoff = 10 * time.Second
		}
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *pubSubSubscriptionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating pubsub subscription", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	diffs, updateMask, err := compareSubscription(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for pubsub subscription", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSubscriptionRequest{
		Subscription: desired,
		UpdateMask:   updateMask,
	}

	updatedSubscription, err := a.gcpClient.UpdateSubscription(ctx, req)
	if err != nil {
		return fmt.Errorf("updating pubsub subscription %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated pubsub subscription", "name", a.id)

	return a.updateStatus(ctx, updateOp, updatedSubscription)
}

func (a *pubSubSubscriptionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting pubsub subscription", "name", a.id)

	req := &pb.DeleteSubscriptionRequest{Subscription: a.id.String()}
	err := a.gcpClient.DeleteSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting pubsub subscription %s: %w", a.id, err)
	}
	log.Info("successfully deleted pubsub subscription", "name", a.id)

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *pubSubSubscriptionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PubSubSubscription{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *PubSubSubscriptionSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Subscription)
	u.SetGroupVersionKind(krm.PubSubSubscriptionGVK)

	return u, nil
}

func compareSubscription(ctx context.Context, actual, desired *pb.Subscription) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PubSubSubscriptionSpec_FromProto, PubSubSubscriptionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Subscription) {
		if obj.AckDeadlineSeconds == 0 {
			obj.AckDeadlineSeconds = 10
		}
		if obj.ExpirationPolicy == nil {
			obj.ExpirationPolicy = &pb.ExpirationPolicy{
				Ttl: &durationpb.Duration{
					Seconds: 3600 * 24 * 31,
				},
			}
		}
		if obj.MessageRetentionDuration == nil {
			obj.MessageRetentionDuration = &durationpb.Duration{
				Seconds: 3600 * 24 * 7,
			}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *pubSubSubscriptionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Subscription) error {
	status := &krm.PubSubSubscriptionStatus{}
	return op.UpdateStatus(ctx, status, nil)
}
