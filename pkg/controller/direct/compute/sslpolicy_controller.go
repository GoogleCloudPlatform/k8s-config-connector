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

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeSSLPolicyGVK, NewSSLPolicyModel)
}

func NewSSLPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sslPolicyModel{config: config}, nil
}

type sslPolicyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &sslPolicyModel{}

type SSLPolicyAdapter struct {
	gcpClient *compute.SslPoliciesClient
	id        *krm.ComputeSSLPolicyIdentity
	desired   *pb.SslPolicy
	actual    *pb.SslPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &SSLPolicyAdapter{}

func (m *sslPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeSSLPolicy{}
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

	sslPoliciesClient, err := gcpClient.newSslPoliciesClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeSSLPolicySpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set name in proto
	desired.Name = direct.LazyPtr(id.(*krm.ComputeSSLPolicyIdentity).SslPolicy)

	// Set defaults
	if desired.MinTlsVersion == nil {
		desired.MinTlsVersion = direct.PtrTo("TLS_1_0")
	}
	if desired.Profile == nil {
		desired.Profile = direct.PtrTo("COMPATIBLE")
	}

	return &SSLPolicyAdapter{
		gcpClient: sslPoliciesClient,
		id:        id.(*krm.ComputeSSLPolicyIdentity),
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *sslPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *SSLPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeSSLPolicy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeSSLPolicy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *SSLPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeSSLPolicy", "name", a.id)

	req := &pb.InsertSslPolicyRequest{
		Project:           a.id.Project,
		SslPolicyResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeSSLPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeSSLPolicy %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeSSLPolicy", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSSLPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *SSLPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeSSLPolicy", "name", a.id)

	// Handle output-only fields from GCP
	a.assignGCPDefaults(a.desired, a.actual)

	paths, report, err := common.CompareProtoMessageStructuredDiff(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	// Perform the Patch update
	req := &pb.PatchSslPolicyRequest{
		Project:           a.id.Project,
		SslPolicy:         a.id.SslPolicy,
		SslPolicyResource: a.desired,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("patching ComputeSSLPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeSSLPolicy %s patch failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully updated ComputeSSLPolicy", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSSLPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *SSLPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeSSLPolicy", "name", a.id)

	req := &pb.DeleteSslPolicyRequest{
		Project:   a.id.Project,
		SslPolicy: a.id.SslPolicy,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeSSLPolicy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeSSLPolicy %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeSSLPolicy", "name", a.id)
	return true, nil
}

func (a *SSLPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeSSLPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeSSLPolicySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeSSLPolicyGVK)
	return u, nil
}

func (a *SSLPolicyAdapter) get(ctx context.Context) (*pb.SslPolicy, error) {
	req := &pb.GetSslPolicyRequest{
		Project:   a.id.Project,
		SslPolicy: a.id.SslPolicy,
	}
	return a.gcpClient.Get(ctx, req)
}

func (a *SSLPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SslPolicy) error {
	mapCtx := &direct.MapContext{}
	status := ComputeSSLPolicyStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *SSLPolicyAdapter) assignGCPDefaults(desired, actual *pb.SslPolicy) {
	desired.Id = actual.Id
	desired.Kind = actual.Kind
	desired.CreationTimestamp = actual.CreationTimestamp
	desired.EnabledFeatures = actual.EnabledFeatures
	desired.Fingerprint = actual.Fingerprint
	desired.SelfLink = actual.SelfLink
	desired.Warnings = actual.Warnings

	if len(desired.CustomFeatures) == 0 && len(actual.CustomFeatures) == 0 {
		desired.CustomFeatures = nil
		actual.CustomFeatures = nil
	}
}
