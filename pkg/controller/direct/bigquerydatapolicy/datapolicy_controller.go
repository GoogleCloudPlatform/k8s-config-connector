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

package bigquerydatapolicy

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/bigquery/datapolicies/apiv1beta1"
	pb "cloud.google.com/go/bigquery/datapolicies/apiv1beta1/datapoliciespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatapolicy/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "bigquerydatapolicy-controller"
)

func init() {
	registry.RegisterModel(krm.BigQueryDataPolicyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.DataPolicyClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataPolicyRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building bigquerydatapolicy client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryDataPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	identity, err := krm.NewDataPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	// Get bigquerydatapolicy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        identity,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DataPolicyIdentity
	gcpClient *gcp.DataPolicyClient
	desired   *krm.BigQueryDataPolicy
	actual    *pb.DataPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) normalizeReference(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.PolicyTagRef != nil {
		external, err := obj.Spec.PolicyTagRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.PolicyTagRef.External = external
	}
	return nil
}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryDataPolicy", "name", a.id.String())

	req := &pb.GetDataPolicyRequest{Name: a.id.String()}
	dataPolicy, err := a.gcpClient.GetDataPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryDataPolicy %q: %w", a.id.String(), err)
	}

	a.actual = dataPolicy
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigQueryDataPolicy", "name", a.id.String())

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BigQueryDataPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if resource.DataPolicyId == "" {
		resource.DataPolicyId = a.id.ID()
	}

	parent := a.id.Parent().String()
	req := &pb.CreateDataPolicyRequest{
		Parent:     parent,
		DataPolicy: resource,
	}

	created, err := a.gcpClient.CreateDataPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigQueryDataPolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created BigQueryDataPolicy", "name", created.Name)

	status := &krm.BigQueryDataPolicyStatus{}
	status.ObservedState = BigQueryDataPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := created.Name
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigQueryDataPolicy", "name", a.id.String())

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BigQueryDataPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.actual.Name

	paths := []string{"data_policy_type", "data_masking_policy", "policy_tag"}
	updateMask := &fieldmaskpb.FieldMask{Paths: paths}

	req := &pb.UpdateDataPolicyRequest{
		DataPolicy: resource,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateDataPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BigQueryDataPolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated BigQueryDataPolicy", "name", updated.Name)

	status := &krm.BigQueryDataPolicyStatus{}
	status.ObservedState = BigQueryDataPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := updated.Name
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryDataPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryDataPolicySpec_FromProto(mapCtx, a.actual))
	tokens := strings.Split(a.id.String(), "/datapolicies/")
	obj.Spec.ResourceID = direct.LazyPtr(tokens[len(tokens)-1])
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.Parent().String()
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: tokens[1]}
			obj.Spec.Location = tokens[3]
		}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigQueryDataPolicy", "name", a.id.String())

	req := &pb.DeleteDataPolicyRequest{Name: a.id.String()}
	if err := a.gcpClient.DeleteDataPolicy(ctx, req); err != nil {
		return false, fmt.Errorf("deleting BigQueryDataPolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted BigQueryDataPolicy", "name", a.id.String())
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
