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
	registry.RegisterModel(krm.ComputeSecurityPolicyGVK, NewSecurityPolicyModel)
}

func NewSecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &securityPolicyModel{config: config}, nil
}

type securityPolicyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &securityPolicyModel{}

type SecurityPolicyAdapter struct {
	globalClient   *compute.SecurityPoliciesClient
	regionalClient *compute.RegionSecurityPoliciesClient
	id             *krm.ComputeSecurityPolicyIdentity
	desired        *pb.SecurityPolicy
	actual         *pb.SecurityPolicy
	reader         client.Reader
}

var _ directbase.Adapter = &SecurityPolicyAdapter{}

func (m *securityPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeSecurityPolicy{}
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

	var globalClient *compute.SecurityPoliciesClient
	var regionalClient *compute.RegionSecurityPoliciesClient
	identity := id.(*krm.ComputeSecurityPolicyIdentity)
	if !identity.IsGlobal() {
		regionalClient, err = gcpClient.newRegionSecurityPoliciesClient(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		globalClient, err = gcpClient.newSecurityPoliciesClient(ctx)
		if err != nil {
			return nil, err
		}
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeSecurityPolicySpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(identity.Name)

	return &SecurityPolicyAdapter{
		globalClient:   globalClient,
		regionalClient: regionalClient,
		id:             identity,
		desired:        desired,
		reader:         reader,
	}, nil
}

func (m *securityPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *SecurityPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeSecurityPolicy", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeSecurityPolicy %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *SecurityPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeSecurityPolicy", "name", a.id)

	if !a.id.IsGlobal() {
		req := &pb.InsertRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicyResource: a.desired,
		}
		op, err := a.regionalClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeSecurityPolicy %s create failed: %w", a.id, err)
			}
		}
	} else {
		req := &pb.InsertSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicyResource: a.desired,
		}
		op, err := a.globalClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("creating ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeSecurityPolicy %s create failed: %w", a.id, err)
			}
		}
	}
	log.V(2).Info("successfully created ComputeSecurityPolicy", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSecurityPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *SecurityPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeSecurityPolicy", "name", a.id)

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

	if !a.id.IsGlobal() {
		req := &pb.PatchRegionSecurityPolicyRequest{
			Project:                a.id.Project,
			Region:                 a.id.Region,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: a.desired,
		}
		op, err := a.regionalClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("patching ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeSecurityPolicy %s patch failed: %w", a.id, err)
			}
		}
	} else {
		req := &pb.PatchSecurityPolicyRequest{
			Project:                a.id.Project,
			SecurityPolicy:         a.id.Name,
			SecurityPolicyResource: a.desired,
		}
		op, err := a.globalClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("patching ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeSecurityPolicy %s patch failed: %w", a.id, err)
			}
		}
	}
	log.V(2).Info("successfully updated ComputeSecurityPolicy", "name", a.id)

	// Get the updated resource
	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSecurityPolicy %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *SecurityPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeSecurityPolicy", "name", a.id)

	if !a.id.IsGlobal() {
		req := &pb.DeleteRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		op, err := a.regionalClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return false, fmt.Errorf("waiting ComputeSecurityPolicy %s delete failed: %w", a.id, err)
			}
		}
	} else {
		req := &pb.DeleteSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		op, err := a.globalClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeSecurityPolicy %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return false, fmt.Errorf("waiting ComputeSecurityPolicy %s delete failed: %w", a.id, err)
			}
		}
	}
	log.V(2).Info("successfully deleted ComputeSecurityPolicy", "name", a.id)
	return true, nil
}

func (a *SecurityPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeSecurityPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeSecurityPolicySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeSecurityPolicyGVK)
	return u, nil
}

func (a *SecurityPolicyAdapter) get(ctx context.Context) (*pb.SecurityPolicy, error) {
	if !a.id.IsGlobal() {
		req := &pb.GetRegionSecurityPolicyRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			SecurityPolicy: a.id.Name,
		}
		return a.regionalClient.Get(ctx, req)
	} else {
		req := &pb.GetSecurityPolicyRequest{
			Project:        a.id.Project,
			SecurityPolicy: a.id.Name,
		}
		return a.globalClient.Get(ctx, req)
	}
}

func (a *SecurityPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SecurityPolicy) error {
	mapCtx := &direct.MapContext{}
	status := ComputeSecurityPolicyStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *SecurityPolicyAdapter) assignGCPDefaults(desired, actual *pb.SecurityPolicy) {
	desired.Id = actual.Id
	desired.Kind = actual.Kind
	desired.CreationTimestamp = actual.CreationTimestamp
	desired.Fingerprint = actual.Fingerprint
	desired.LabelFingerprint = actual.LabelFingerprint
	desired.SelfLink = actual.SelfLink
	// If region is us-central1, GCP API converts it to full url
	// https://www.googleapis.com/compute/v1/projects/projectID/regions/us-central1
	if actual.Region != nil {
		*actual.Region = lastComponent(*actual.Region)
	}
	// If type is unspecified, the default value is "CLOUD_ARMOR"
	if desired.Type == nil {
		desired.Type = actual.Type
	}
}

func ComputeSecurityPolicyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krm.ComputeSecurityPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSecurityPolicyStatus{}
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	return out
}
