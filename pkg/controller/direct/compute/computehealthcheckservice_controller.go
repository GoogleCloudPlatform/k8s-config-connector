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

// +generated:controller
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ComputeHealthCheckService

package compute

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeHealthCheckServiceGVK, NewHealthCheckServiceModel)
}

func NewHealthCheckServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelHealthCheckService{config: *config}, nil
}

var _ directbase.Model = &modelHealthCheckService{}

type modelHealthCheckService struct {
	config config.ControllerConfig
}

func (m *modelHealthCheckService) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeHealthCheckService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := resolveHealthCheckServiceReferences(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("resolving references: %w", err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComputeHealthCheckServiceIdentity)

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.newRegionHealthCheckServicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &HealthCheckServiceAdapter{
		id:        id,
		gcpClient: client,
		desired:   obj,
		reader:    reader,
	}, nil
}

func resolveHealthCheckServiceReferences(ctx context.Context, reader client.Reader, obj *krm.ComputeHealthCheckService) error {
	defaultNamespace := obj.GetNamespace()
	for _, ref := range obj.Spec.HealthCheckRefs {
		if ref == nil {
			continue
		}
		if ref.External != "" {
			ref.External = apirefs.TrimComputeURIPrefix(ref.External)
			continue
		}
		if ref.Name == "" {
			continue
		}
		ns := ref.Namespace
		if ns == "" {
			ns = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(computev1beta1.ComputeHealthCheckGVK)
		key := types.NamespacedName{Namespace: ns, Name: ref.Name}
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("getting referenced ComputeHealthCheck %s: %w", key, err)
		}
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink == "" {
			externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
			selfLink = externalRef
		}
		if selfLink == "" {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		ref.External = apirefs.TrimComputeURIPrefix(selfLink)
	}
	return nil
}

func (m *modelHealthCheckService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//compute.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComputeHealthCheckServiceIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(&m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.newRegionHealthCheckServicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &HealthCheckServiceAdapter{
		id:        id,
		gcpClient: client,
	}, nil
}

type HealthCheckServiceAdapter struct {
	id        *krm.ComputeHealthCheckServiceIdentity
	gcpClient *gcp.RegionHealthCheckServicesClient
	desired   *krm.ComputeHealthCheckService
	actual    *computepb.HealthCheckService
	reader    client.Reader
}

var _ directbase.Adapter = &HealthCheckServiceAdapter{}

func (a *HealthCheckServiceAdapter) get(ctx context.Context) (*computepb.HealthCheckService, error) {
	req := &computepb.GetRegionHealthCheckServiceRequest{
		Project:            a.id.Project,
		Region:             a.id.Region,
		HealthCheckService: a.id.Healthcheckservice,
	}
	return a.gcpClient.Get(ctx, req)
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *HealthCheckServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HealthCheckService", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeHealthCheckService %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *HealthCheckServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HealthCheckService", "name", a.id)

	mapCtx := &direct.MapContext{}
	resource := ComputeHealthCheckServiceSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.Healthcheckservice)

	req := &computepb.InsertRegionHealthCheckServiceRequest{
		Project:                    a.id.Project,
		Region:                     a.id.Region,
		HealthCheckServiceResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeHealthCheckService %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for ComputeHealthCheckService %s creation: %w", a.id, err)
	}

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeHealthCheckService %s after creation: %w", a.id, err)
	}

	return a.setStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *HealthCheckServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HealthCheckService", "name", a.id)

	normalize := func(ctx context.Context, pb *computepb.HealthCheckService) error {
		if pb.HealthChecks != nil {
			pb.HealthChecks = canonicalizeComputeURLs(pb.HealthChecks)
		}
		if pb.NetworkEndpointGroups != nil {
			pb.NetworkEndpointGroups = canonicalizeComputeURLs(pb.NetworkEndpointGroups)
		}
		if pb.NotificationEndpoints != nil {
			pb.NotificationEndpoints = canonicalizeComputeURLs(pb.NotificationEndpoints)
		}
		return nil
	}

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeHealthCheckServiceSpec_v1alpha1_FromProto,
		ComputeHealthCheckServiceSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)

		resource.Name = direct.LazyPtr(a.id.Healthcheckservice)
		resource.Fingerprint = a.actual.Fingerprint

		req := &computepb.PatchRegionHealthCheckServiceRequest{
			Project:                    a.id.Project,
			Region:                     a.id.Region,
			HealthCheckService:         a.id.Healthcheckservice,
			HealthCheckServiceResource: resource,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("patching ComputeHealthCheckService %s: %w", a.id, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for ComputeHealthCheckService %s update: %w", a.id, err)
		}

		updated, err := a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeHealthCheckService %s after update: %w", a.id, err)
		}
		return a.setStatus(ctx, updateOp, updated)
	}

	return a.setStatus(ctx, updateOp, a.actual)
}

func (a *HealthCheckServiceAdapter) setStatus(ctx context.Context, op directbase.Operation, actual *computepb.HealthCheckService) error {
	status := &krm.ComputeHealthCheckServiceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComputeHealthCheckServiceObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *HealthCheckServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeHealthCheckService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeHealthCheckServiceSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Region

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Healthcheckservice)
	u.SetGroupVersionKind(krm.ComputeHealthCheckServiceGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *HealthCheckServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HealthCheckService", "name", a.id)

	req := &computepb.DeleteRegionHealthCheckServiceRequest{
		Project:            a.id.Project,
		Region:             a.id.Region,
		HealthCheckService: a.id.Healthcheckservice,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeHealthCheckService %s: %w", a.id, err)
	}
	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for ComputeHealthCheckService %s deletion: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeHealthCheckService", "name", a.id)
	return true, nil
}
