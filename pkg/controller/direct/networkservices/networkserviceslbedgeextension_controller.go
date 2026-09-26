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
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/networkservices/apiv1"
	networkservicespb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesLBEdgeExtensionGVK, NewLBEdgeExtensionModel)
}

func NewLBEdgeExtensionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLBEdgeExtension{config: *config}, nil
}

var _ directbase.Model = &modelLBEdgeExtension{}

type modelLBEdgeExtension struct {
	config config.ControllerConfig
}

func (m *modelLBEdgeExtension) client(ctx context.Context) (*gcp.DepClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDepRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building LBEdgeExtension client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLBEdgeExtension) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	desired := &krm.NetworkServicesLBEdgeExtension{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &desired); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", desired, err)
	}

	id, err := krm.NewLBEdgeExtensionIdentity(ctx, reader, desired)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &LBEdgeExtensionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
		labels:    label.NewGCPLabelsFromK8sLabels(u.GetLabels()),
	}, nil
}

func (m *modelLBEdgeExtension) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LBEdgeExtensionAdapter struct {
	id        *krm.LBEdgeExtensionIdentity
	gcpClient *gcp.DepClient
	desired   *krm.NetworkServicesLBEdgeExtension
	reader    client.Reader
	actual    *networkservicespb.LbEdgeExtension
	labels    map[string]string
}

var _ directbase.Adapter = &LBEdgeExtensionAdapter{}

func (a *LBEdgeExtensionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LBEdgeExtension", "name", a.id)

	req := &networkservicespb.GetLbEdgeExtensionRequest{Name: a.id.String()}
	lbedgeextensionpb, err := a.gcpClient.GetLbEdgeExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LBEdgeExtension %q: %w", a.id, err)
	}

	a.actual = lbedgeextensionpb
	a.normalizeActual(a.actual)
	return true, nil
}

func (a *LBEdgeExtensionAdapter) normalizeActual(obj *networkservicespb.LbEdgeExtension) {
	if obj == nil {
		return
	}
	projectID := a.id.Project
	// GCP often returns project numbers in URLs. We normalize them to project IDs to match the desired state.
	for i, rule := range obj.ForwardingRules {
		obj.ForwardingRules[i] = a.normalizeURL(rule, projectID)
	}
	for _, chain := range obj.ExtensionChains {
		for _, extension := range chain.Extensions {
			extension.Service = a.normalizeURL(extension.Service, projectID)
		}
	}
}

func (a *LBEdgeExtensionAdapter) normalizeURL(url string, projectID string) string {
	if !strings.HasPrefix(url, "https://www.googleapis.com/compute/v1/projects/") &&
		!strings.HasPrefix(url, "https://compute.googleapis.com/compute/v1/projects/") {
		return url
	}
	// Format: https://[hostname]/compute/v1/projects/[projectID_or_number]/...
	tokens := strings.Split(url, "/")
	if len(tokens) < 7 {
		return url
	}
	// If it's a number (or just not the ID), replace it.
	// Since we know the project ID from the identity, we can safely substitute it.
	if tokens[6] != projectID {
		tokens[6] = projectID
	}
	return strings.Join(tokens, "/")
}

func (a *LBEdgeExtensionAdapter) resolve(ctx context.Context) (*networkservicespb.LbEdgeExtension, error) {
	reader := a.reader
	desired := a.desired
	projectID := a.id.Project

	// Resolve references
	for _, ref := range desired.Spec.ForwardingRuleRefs {
		if ref == nil {
			continue
		}
		if err := ref.Normalize(ctx, reader, desired.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving forwardingRuleRef: %w", err)
		}
		// GCP LBEdgeExtension returns full URLs for forwarding rules.
		// ForwardingRuleRef.Normalize strips the prefix, so we must add it back.
		if ref.External != "" && !strings.HasPrefix(ref.External, "https://") {
			ref.External = "https://www.googleapis.com/compute/v1/" + ref.External
		}

		// Ensure the forwarding rule is in the same project as the LBEdgeExtension.
		if refProject := common.ExtractProjectID(ref.External); refProject == "" || refProject != projectID {
			return nil, fmt.Errorf("cross-project references are not supported for LBEdgeExtension: forwardingRule %q is in project %q, but LBEdgeExtension is in project %q", ref.External, refProject, projectID)
		}
	}
	for i := range desired.Spec.ExtensionChains {
		chain := &desired.Spec.ExtensionChains[i]
		for j := range chain.Extensions {
			extension := &chain.Extensions[j]
			if extension.BackendServiceRef != nil {
				external, err := extension.BackendServiceRef.NormalizedExternal(ctx, reader, desired.GetNamespace())
				if err != nil {
					return nil, fmt.Errorf("resolving backendServiceRef: %w", err)
				}
				extension.BackendServiceRef.External = external
				extension.BackendServiceRef.Name = ""
				extension.BackendServiceRef.Namespace = ""

				// Ensure the backend service is in the same project as the LBEdgeExtension.
				if refProject := common.ExtractProjectID(external); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBEdgeExtension: backendService %q is in project %q, but LBEdgeExtension is in project %q", external, refProject, projectID)
				}
			}
			if extension.WasmPluginRef != nil {
				if err := extension.WasmPluginRef.Normalize(ctx, reader, desired.GetNamespace()); err != nil {
					return nil, fmt.Errorf("resolving wasmPluginRef: %w", err)
				}
				// Ensure the wasm plugin is in the same project as the LBEdgeExtension.
				if refProject := common.ExtractProjectID(extension.WasmPluginRef.External); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBEdgeExtension: wasmPlugin %q is in project %q, but LBEdgeExtension is in project %q", extension.WasmPluginRef.External, refProject, projectID)
				}
			}
		}
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesLBEdgeExtensionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set GCP Labels
	desiredProto.Labels = a.labels

	return desiredProto, nil
}

func (a *LBEdgeExtensionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LBEdgeExtension", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.CloneOf(desiredProto)
	resource.Name = a.id.String()

	req := &networkservicespb.CreateLbEdgeExtensionRequest{
		Parent:            fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		LbEdgeExtensionId: a.id.LbEdgeExtension,
		LbEdgeExtension:   resource,
	}
	op, err := a.gcpClient.CreateLbEdgeExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LBEdgeExtension %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("LBEdgeExtension %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created LBEdgeExtension", "name", a.id)

	status := &krm.NetworkServicesLBEdgeExtensionStatus{}
	status.ObservedState = NetworkServicesLBEdgeExtensionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *LBEdgeExtensionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LBEdgeExtension", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.CloneOf(desiredProto)
	resource.Name = a.id.String()

	diffs, updateMask, err := compareLBEdgeExtension(ctx, a.actual, resource)
	if err != nil {
		return fmt.Errorf("comparing LBEdgeExtension %s: %w", a.id, err)
	}

	updated := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for LBEdgeExtension", "name", a.id)
	} else {
		// Report exact diffs
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &networkservicespb.UpdateLbEdgeExtensionRequest{
			UpdateMask:      updateMask,
			LbEdgeExtension: resource,
		}
		op, err := a.gcpClient.UpdateLbEdgeExtension(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LBEdgeExtension %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("LBEdgeExtension %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated LBEdgeExtension", "name", a.id)
	}

	status := &krm.NetworkServicesLBEdgeExtensionStatus{}
	status.ObservedState = NetworkServicesLBEdgeExtensionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *LBEdgeExtensionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesLBEdgeExtension{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesLBEdgeExtensionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.LbEdgeExtension)
	u.SetGroupVersionKind(krm.NetworkServicesLBEdgeExtensionGVK)

	u.Object = uObj
	return u, nil
}

func (a *LBEdgeExtensionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LBEdgeExtension", "name", a.id)

	req := &networkservicespb.DeleteLbEdgeExtensionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteLbEdgeExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LBEdgeExtension, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting LBEdgeExtension %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LBEdgeExtension", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete LBEdgeExtension %s: %w", a.id, err)
	}
	return true, nil
}

func compareLBEdgeExtension(ctx context.Context, actual, desired *networkservicespb.LbEdgeExtension) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesLBEdgeExtensionSpec_FromProto, NetworkServicesLBEdgeExtensionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *networkservicespb.LbEdgeExtension) {
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
