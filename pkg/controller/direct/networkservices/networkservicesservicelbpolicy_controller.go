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

package networkservices

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/networkservices/apiv1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesServiceLBPolicyGVK, NewNetworkServicesServiceLBPolicyModel)
}

func NewNetworkServicesServiceLBPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkServicesServiceLBPolicy{config: *config}, nil
}

var _ directbase.Model = &modelNetworkServicesServiceLBPolicy{}

type modelNetworkServicesServiceLBPolicy struct {
	config config.ControllerConfig
}

func (m *modelNetworkServicesServiceLBPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NetworkServicesServiceLBPolicy client: %w", err)
	}
	return gcpClient, err
}

func (m *modelNetworkServicesServiceLBPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesServiceLBPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewNetworkServicesServiceLBPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesServiceLBPolicySpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &NetworkServicesServiceLBPolicyAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelNetworkServicesServiceLBPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkServicesServiceLBPolicyAdapter struct {
	id           *krm.NetworkServicesServiceLBPolicyIdentity
	gcpClient    *gcp.Client
	desiredProto *pb.ServiceLbPolicy
	actual       *pb.ServiceLbPolicy
}

var _ directbase.Adapter = &NetworkServicesServiceLBPolicyAdapter{}

func (a *NetworkServicesServiceLBPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkServicesServiceLBPolicy", "name", a.id)

	req := &pb.GetServiceLbPolicyRequest{Name: a.id.String()}
	servicelbpolicypb, err := a.gcpClient.GetServiceLbPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkServicesServiceLBPolicy %q: %w", a.id, err)
	}

	a.actual = servicelbpolicypb
	return true, nil
}

func (a *NetworkServicesServiceLBPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkServicesServiceLBPolicy", "name", a.id)

	req := &pb.CreateServiceLbPolicyRequest{
		Parent:            a.id.ParentString(),
		ServiceLbPolicyId: a.id.ServiceLbPolicy,
		ServiceLbPolicy:   a.desiredProto,
	}
	op, err := a.gcpClient.CreateServiceLbPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkServicesServiceLBPolicy %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NetworkServicesServiceLBPolicy %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NetworkServicesServiceLBPolicy", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *NetworkServicesServiceLBPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkServicesServiceLBPolicy", "name", a.id)

	diffs, updateMask, err := compareServiceLBPolicy(ctx, a.actual, a.desiredProto)
	if err != nil {
		return fmt.Errorf("comparing NetworkServicesServiceLBPolicy %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for NetworkServicesServiceLBPolicy", "name", a.id)
	} else {
		// Report exact diffs
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateServiceLbPolicyRequest{
			UpdateMask:      updateMask,
			ServiceLbPolicy: a.desiredProto,
		}
		op, err := a.gcpClient.UpdateServiceLbPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkServicesServiceLBPolicy %s: %w", a.id, err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("NetworkServicesServiceLBPolicy %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated NetworkServicesServiceLBPolicy", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *NetworkServicesServiceLBPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesServiceLBPolicy{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesServiceLBPolicySpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	desired.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ServiceLbPolicy)
	u.SetGroupVersionKind(krm.NetworkServicesServiceLBPolicyGVK)

	return u, nil
}

func (a *NetworkServicesServiceLBPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkServicesServiceLBPolicy", "name", a.id)

	req := &pb.DeleteServiceLbPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteServiceLbPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent NetworkServicesServiceLBPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkServicesServiceLBPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NetworkServicesServiceLBPolicy", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NetworkServicesServiceLBPolicy %s: %w", a.id, err)
	}
	return true, nil
}

func compareServiceLBPolicy(ctx context.Context, actual, desired *pb.ServiceLbPolicy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesServiceLBPolicySpec_v1alpha1_FromProto, NetworkServicesServiceLBPolicySpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.ServiceLbPolicy) {
		// Add any server-side or GCP defaults if known, or leave empty
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *NetworkServicesServiceLBPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ServiceLbPolicy) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkServicesServiceLBPolicyStatus{}
	status.ObservedState = NetworkServicesServiceLBPolicyObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
