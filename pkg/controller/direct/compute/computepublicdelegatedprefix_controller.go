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
	"reflect"

	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func init() {
	registry.RegisterModel(krm.ComputePublicDelegatedPrefixGVK, NewPublicDelegatedPrefixModel)
}

func NewPublicDelegatedPrefixModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &publicDelegatedPrefixModel{config: config}, nil
}

type publicDelegatedPrefixModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &publicDelegatedPrefixModel{}

type publicDelegatedPrefixAdapter struct {
	id                                *krm.PublicDelegatedPrefixIdentity
	publicDelegatedPrefixClient       *gcp.PublicDelegatedPrefixesClient
	globalPublicDelegatedPrefixClient *gcp.GlobalPublicDelegatedPrefixesClient
	desired                           *krm.ComputePublicDelegatedPrefix
	actual                            *computepb.PublicDelegatedPrefix
	reader                            client.Reader
}

var _ directbase.Adapter = &publicDelegatedPrefixAdapter{}

func (m *publicDelegatedPrefixModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputePublicDelegatedPrefix{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.PublicDelegatedPrefixIdentity)

	adapter := &publicDelegatedPrefixAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get region
	region := id.ParentID.Location

	// Get GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	if region == "global" {
		globalClient, err := gcpClient.newGlobalPublicDelegatedPrefixesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.globalPublicDelegatedPrefixClient = globalClient
	} else {
		regionalClient, err := gcpClient.publicDelegatedPrefixesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.publicDelegatedPrefixClient = regionalClient
	}
	return adapter, nil
}

func (m *publicDelegatedPrefixModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *publicDelegatedPrefixAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputePublicDelegatedPrefix", "name", a.id)

	var err error
	pdp := &computepb.PublicDelegatedPrefix{}
	pdp, err = a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputePublicDelegatedPrefix %q: %w", a.id, err)
	}
	a.actual = pdp
	return true, nil
}

func (a *publicDelegatedPrefixAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	var err error

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputePublicDelegatedPrefix", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	// No other refs to resolve for now based on spec (ProjectRef is handled by identity)

	pdp := ComputePublicDelegatedPrefixSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if desired.Spec.ParentPrefixRef != nil {
		parentPrefix, err := resolveParentPrefixRef(ctx, a.reader, a.desired, desired.Spec.ParentPrefixRef)
		if err != nil {
			return err
		}
		pdp.ParentPrefix = direct.LazyPtr(parentPrefix)
	}

	pdp.Name = direct.LazyPtr(a.id.ResourceID)

	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.InsertGlobalPublicDelegatedPrefixeRequest{
			PublicDelegatedPrefixResource: pdp,
			Project:                       a.id.ParentID.ProjectID,
		}
		op, err = a.globalPublicDelegatedPrefixClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertPublicDelegatedPrefixeRequest{
			PublicDelegatedPrefixResource: pdp,
			Region:                        a.id.ParentID.Location,
			Project:                       a.id.ParentID.ProjectID,
		}
		op, err = a.publicDelegatedPrefixClient.Insert(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating ComputePublicDelegatedPrefix %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputePublicDelegatedPrefix %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputePublicDelegatedPrefix", "name", a.id)

	// Get the created resource
	_, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputePublicDelegatedPrefix %q: %w", a.id, err)
	}

	// Set status
	status := &krm.ComputePublicDelegatedPrefixStatus{
		ObservedGeneration: direct.LazyPtr(u.GetGeneration()),
		ExternalRef:        direct.LazyPtr(a.id.String()),
	}
	return setPublicDelegatedPrefixStatus(u, status)
}

func (a *publicDelegatedPrefixAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	if a.id.ResourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputePublicDelegatedPrefix", "name", a.id.ResourceID)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	pdp := ComputePublicDelegatedPrefixSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if desired.Spec.ParentPrefixRef != nil {
		parentPrefix, err := resolveParentPrefixRef(ctx, a.reader, a.desired, desired.Spec.ParentPrefixRef)
		if err != nil {
			return err
		}
		pdp.ParentPrefix = direct.LazyPtr(parentPrefix)
	}

	pdp.Name = direct.LazyPtr(a.id.ResourceID)

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if !reflect.DeepEqual(pdp.Description, a.actual.Description) {
		report.AddField("description", a.actual.Description, pdp.Description)
		// Perform patch
		op := &gcp.Operation{}
		var err error
		if a.id.ParentID.Location == "global" {
			req := &computepb.PatchGlobalPublicDelegatedPrefixeRequest{
				PublicDelegatedPrefix:         a.id.ResourceID,
				PublicDelegatedPrefixResource: pdp,
				Project:                       a.id.ParentID.ProjectID,
			}
			op, err = a.globalPublicDelegatedPrefixClient.Patch(ctx, req)
		} else {
			req := &computepb.PatchPublicDelegatedPrefixeRequest{
				PublicDelegatedPrefix:         a.id.ResourceID,
				PublicDelegatedPrefixResource: pdp,
				Region:                        a.id.ParentID.Location,
				Project:                       a.id.ParentID.ProjectID,
			}
			op, err = a.publicDelegatedPrefixClient.Patch(ctx, req)
		}
		if err != nil {
			return fmt.Errorf("updating ComputePublicDelegatedPrefix %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputePublicDelegatedPrefix %s update failed: %w", a.id, err)
			}
		}
	}

	structuredreporting.ReportDiff(ctx, report)

	// Get the updated resource
	_, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputePublicDelegatedPrefix %q: %w", a.id.ResourceID, err)
	}

	status := &krm.ComputePublicDelegatedPrefixStatus{
		ObservedGeneration: direct.LazyPtr(updateOp.GetUnstructured().GetGeneration()),
		ExternalRef:        direct.LazyPtr(a.id.String()),
	}
	return setPublicDelegatedPrefixStatus(u, status)
}

func (a *publicDelegatedPrefixAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("publicDelegatedPrefix %q not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputePublicDelegatedPrefixSpec_v1beta1_FromProto(mc, a.actual)
	
	// Map returned string parent prefix back to ParentPrefixRef
	if a.actual.ParentPrefix != nil {
		spec.ParentPrefixRef = &krm.ComputePublicDelegatedPrefixParentPrefixRef{
			External: *a.actual.ParentPrefix,
		}
	}
	
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting publicDelegatedPrefix spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ComputePublicDelegatedPrefixGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *publicDelegatedPrefixAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.ResourceID == "" {
		return false, fmt.Errorf("resourceID is empty")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputePublicDelegatedPrefix", "name", a.id.ResourceID)

	var err error
	op := &gcp.Operation{}
	if a.id.ParentID.Location == "global" {
		req := &computepb.DeleteGlobalPublicDelegatedPrefixeRequest{
			PublicDelegatedPrefix: a.id.ResourceID,
			Project:               a.id.ParentID.ProjectID,
		}
		op, err = a.globalPublicDelegatedPrefixClient.Delete(ctx, req)
	} else {
		req := &computepb.DeletePublicDelegatedPrefixeRequest{
			PublicDelegatedPrefix: a.id.ResourceID,
			Region:                a.id.ParentID.Location,
			Project:               a.id.ParentID.ProjectID,
		}
		op, err = a.publicDelegatedPrefixClient.Delete(ctx, req)
	}
	if err != nil {
		return false, fmt.Errorf("deleting ComputePublicDelegatedPrefix %s: %w", a.id.ResourceID, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputePublicDelegatedPrefix %s delete failed: %w", a.id.ResourceID, err)
		}
	}
	log.V(2).Info("successfully deleted ComputePublicDelegatedPrefix", "name", a.id.ResourceID)
	return true, nil
}

func (a *publicDelegatedPrefixAdapter) get(ctx context.Context) (*computepb.PublicDelegatedPrefix, error) {
	if a.id.ParentID.Location == "global" {
		getReq := &computepb.GetGlobalPublicDelegatedPrefixeRequest{
			PublicDelegatedPrefix: a.id.ResourceID,
			Project:               a.id.ParentID.ProjectID,
		}
		return a.globalPublicDelegatedPrefixClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetPublicDelegatedPrefixeRequest{
			PublicDelegatedPrefix: a.id.ResourceID,
			Region:                a.id.ParentID.Location,
			Project:               a.id.ParentID.ProjectID,
		}
		return a.publicDelegatedPrefixClient.Get(ctx, getReq)
	}
}

func setPublicDelegatedPrefixStatus(u *unstructured.Unstructured, status interface{}) error {
	statusMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(status)
	if err != nil {
		return fmt.Errorf("converting status to unstructured: %w", err)
	}
	return unstructured.SetNestedField(u.Object, statusMap, "status")
}

func resolveParentPrefixRef(ctx context.Context, reader client.Reader, src client.Object, ref *krm.ComputePublicDelegatedPrefixParentPrefixRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return "", fmt.Errorf("cannot specify both name and external on parentPrefixRef")
		}
		return ref.External, nil
	}

	if ref.Name == "" {
		return "", fmt.Errorf("must specify either name or external on parentPrefixRef")
	}

	key := client.ObjectKey{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	kind := ref.Kind
	if kind == "" {
		kind = "ComputePublicAdvertisedPrefix"
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    kind,
	})

	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", kind, key, err)
	}

	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef on %s %s: %w", kind, key, err)
	}
	if actualExternalRef != "" {
		return actualExternalRef, nil
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return "", fmt.Errorf("reading status.selfLink on %s %s: %w", kind, key, err)
	}
	if selfLink != "" {
		return selfLink, nil
	}

	return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
}
