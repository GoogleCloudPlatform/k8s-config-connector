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

package managedkafka

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/managedkafka/apiv1"
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ManagedKafkaTopicGVK, NewTopicModel)
}

func NewTopicModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTopic{config: *config}, nil
}

var _ directbase.Model = &modelTopic{}

type modelTopic struct {
	config config.ControllerConfig
}

func (m *modelTopic) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Topic client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTopic) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ManagedKafkaTopic{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTopicIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get managedkafka GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TopicAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTopic) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TopicAdapter struct {
	id        *krm.TopicIdentity
	gcpClient *gcp.Client
	desired   *krm.ManagedKafkaTopic
	actual    *pb.Topic
}

var _ directbase.Adapter = &TopicAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TopicAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Topic", "name", a.id)

	req := &pb.GetTopicRequest{Name: a.id.String()}
	topicpb, err := a.gcpClient.GetTopic(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Topic %q: %w", a.id, err)
	}

	a.actual = topicpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TopicAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Topic", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ManagedKafkaTopicSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateTopicRequest{
		Parent:  a.id.Parent().String(),
		TopicId: a.id.ID(),
		Topic:   resource,
	}
	created, err := a.gcpClient.CreateTopic(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Topic %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Topic", "name", a.id)

	status := &krm.ManagedKafkaTopicStatus{}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TopicAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Topic", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := ManagedKafkaTopicSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set the name field to ensure the GCP API can identity the resource during UpdateTopic().
	// This also prevents incorrect diffs, as the name field is not populated by ManagedKafkaTopicSpec_ToProto.
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.ManagedKafkaTopicStatus{}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateTopicRequest{
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		Topic:      desiredPb,
	}
	updated, err := a.gcpClient.UpdateTopic(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Topic %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Topic", "name", a.id.String())

	status := &krm.ManagedKafkaTopicStatus{}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TopicAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ManagedKafkaTopic{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ManagedKafkaTopicSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ClusterRef = &krm.ClusterRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ManagedKafkaTopicGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TopicAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Topic", "name", a.id)

	req := &pb.DeleteTopicRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTopic(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Topic %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Topic", "name", a.id)

	return true, nil
}
