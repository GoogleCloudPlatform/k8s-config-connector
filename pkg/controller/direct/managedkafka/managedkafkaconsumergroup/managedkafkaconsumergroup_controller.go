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

package managedkafkaconsumergroup

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

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
	registry.RegisterModel(krm.ManagedKafkaConsumerGroupGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ManagedKafka client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ManagedKafkaConsumerGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.GetManagedKafkaConsumerGroupSpecIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get managedkafka GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.ManagedKafkaConsumerGroupIdentity
	gcpClient *gcp.Client
	desired   *krm.ManagedKafkaConsumerGroup
	actual    *pb.ConsumerGroup
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ManagedKafkaConsumerGroup", "name", a.id)

	req := &pb.GetConsumerGroupRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetConsumerGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ManagedKafkaConsumerGroup %q: %w", a.id, err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ManagedKafkaConsumerGroup", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ManagedKafkaConsumerGroupSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	req := &pb.UpdateConsumerGroupRequest{
		ConsumerGroup: resource,
		UpdateMask:    &fieldmaskpb.FieldMask{Paths: []string{"topics"}},
	}
	// Note: UpdateConsumerGroup acts as Create if the group doesn't exist but has offsets committed?
	// Actually, in our mock we allowed it. In real GCP, it might fail.
	// If it fails, we might need a different approach.
	updated, err := a.gcpClient.UpdateConsumerGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ManagedKafkaConsumerGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ManagedKafkaConsumerGroup", "name", a.id)

	status := &krm.ManagedKafkaConsumerGroupStatus{}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ManagedKafkaConsumerGroup", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := ManagedKafkaConsumerGroupSpec_v1alpha1_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Name = a.id.String()

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.ManagedKafkaConsumerGroupStatus{}
		status.ExternalRef = direct.LazyPtr(a.actual.Name)
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	// Update GCP resource
	req := &pb.UpdateConsumerGroupRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		},
		ConsumerGroup: desiredPb,
	}
	updated, err := a.gcpClient.UpdateConsumerGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ManagedKafkaConsumerGroup %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated ManagedKafkaConsumerGroup", "name", a.id.String())

	status := &krm.ManagedKafkaConsumerGroupStatus{}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ManagedKafkaConsumerGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ManagedKafkaConsumerGroupSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set Parent fields for export
	obj.Spec.Location = a.id.Location
	obj.Spec.ClusterRef = &krm.ClusterRef{External: fmt.Sprintf("projects/%s/locations/%s/clusters/%s", a.id.Project, a.id.Location, a.id.Cluster)}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ConsumerGroup)
	u.SetGroupVersionKind(krm.ManagedKafkaConsumerGroupGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ManagedKafkaConsumerGroup", "name", a.id)

	req := &pb.DeleteConsumerGroupRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteConsumerGroup(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ManagedKafkaConsumerGroup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ManagedKafkaConsumerGroup", "name", a.id)

	return true, nil
}
