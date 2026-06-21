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
// proto.service: google.pubsub.v1.Publisher
// proto.message: google.pubsub.v1.Topic
// crd.type: PubSubTopic
// crd.version: v1beta1

package pubsub

import (
	"context"
	"fmt"

	api "cloud.google.com/go/pubsub/v2/apiv1"
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.PubSubTopicGVK, NewTopicModel)
}

func NewTopicModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &TopicModel{config: *config}, nil
}

var _ directbase.Model = &TopicModel{}

type TopicModel struct {
	config config.ControllerConfig
}

func (m *TopicModel) client(ctx context.Context, projectID string) (*api.TopicAdminClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewTopicAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building pubsub publisher client: %w", err)
	}

	return gcpClient, err
}

func (m *TopicModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.PubSubTopic{}
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
	topicId := id.(*krm.PubSubTopicIdentity)

	gcpClient, err := m.client(ctx, topicId.Project)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := PubSubTopicSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = topicId.String()
	desired.Labels = label.GCPLabels(obj)

	return &pubSubTopicAdapter{
		gcpClient: gcpClient,
		id:        topicId,
		desired:   desired,
	}, nil
}

func (m *TopicModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.PubSubTopicIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &pubSubTopicAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type pubSubTopicAdapter struct {
	gcpClient *api.TopicAdminClient
	id        *krm.PubSubTopicIdentity
	desired   *pb.Topic
	actual    *pb.Topic
}

var _ directbase.Adapter = &pubSubTopicAdapter{}

func (a *pubSubTopicAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting pubsub topic", "name", a.id)

	req := &pb.GetTopicRequest{Topic: a.id.String()}
	actual, err := a.gcpClient.GetTopic(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting pubsub topic %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *pubSubTopicAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating pubsub topic", "name", a.id)

	created, err := a.gcpClient.CreateTopic(ctx, a.desired)
	if err != nil {
		return fmt.Errorf("creating pubsub topic %s: %w", a.id.String(), err)
	}
	log.Info("successfully created pubsub topic in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *pubSubTopicAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating pubsub topic", "name", a.id)

	diffs, updateMask, err := compareTopic(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for pubsub topic", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateTopicRequest{
		Topic:      a.desired,
		UpdateMask: updateMask,
	}

	updatedTopic, err := a.gcpClient.UpdateTopic(ctx, req)
	if err != nil {
		return fmt.Errorf("updating pubsub topic %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated pubsub topic", "name", a.id)

	return a.updateStatus(ctx, updateOp, updatedTopic)
}

func (a *pubSubTopicAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting pubsub topic", "name", a.id)

	req := &pb.DeleteTopicRequest{Topic: a.id.String()}
	err := a.gcpClient.DeleteTopic(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting pubsub topic %s: %w", a.id, err)
	}
	log.Info("successfully deleted pubsub topic", "name", a.id)

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *pubSubTopicAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PubSubTopic{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *PubSubTopicSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Topic)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Topic)
	u.SetGroupVersionKind(krm.PubSubTopicGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	// Maintain compatibility with old export
	if a.actual.GetMessageRetentionDuration() != nil {
		seconds := a.actual.GetMessageRetentionDuration().GetSeconds()
		if err := unstructured.SetNestedField(u.Object, fmt.Sprintf("%ds", seconds), "spec", "messageRetentionDuration"); err != nil {
			return nil, err
		}
	}
	if obj.Spec.MessageStoragePolicy != nil {
		enforceInTransit := false
		if obj.Spec.MessageStoragePolicy.EnforceInTransit != nil {
			enforceInTransit = *obj.Spec.MessageStoragePolicy.EnforceInTransit
		}
		if err := unstructured.SetNestedField(u.Object, enforceInTransit, "spec", "messageStoragePolicy", "enforceInTransit"); err != nil {
			return nil, err
		}
	}

	return u, nil
}

func compareTopic(ctx context.Context, actual, desired *pb.Topic) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PubSubTopicSpec_FromProto, PubSubTopicSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.Labels = actual.Labels
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *pubSubTopicAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Topic) error {
	status := &krm.PubSubTopicStatus{}
	return op.UpdateStatus(ctx, status, nil)
}
