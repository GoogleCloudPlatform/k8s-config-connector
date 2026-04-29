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
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

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
	registry.RegisterModel(krm.NetworkServicesLBRouteExtensionGVK, NewLBRouteExtensionModel)
}

func NewLBRouteExtensionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLBRouteExtension{config: *config}, nil
}

var _ directbase.Model = &modelLBRouteExtension{}

type modelLBRouteExtension struct {
	config config.ControllerConfig
}

func (m *modelLBRouteExtension) client(ctx context.Context) (*gcp.DepClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDepRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building LBRouteExtension client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLBRouteExtension) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	desired := &krm.NetworkServicesLBRouteExtension{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &desired); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", desired, err)
	}

	id, err := krm.NewLBRouteExtensionIdentity(ctx, reader, desired)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesLBRouteExtensionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &LBRouteExtensionAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desired:      desired,
		reader:       reader,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelLBRouteExtension) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LBRouteExtensionAdapter struct {
	id           *krm.LBRouteExtensionIdentity
	gcpClient    *gcp.DepClient
	desired      *krm.NetworkServicesLBRouteExtension
	reader       client.Reader
	actual       *networkservicespb.LbRouteExtension
	desiredProto *networkservicespb.LbRouteExtension
}

var _ directbase.Adapter = &LBRouteExtensionAdapter{}

func (a *LBRouteExtensionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LBRouteExtension", "name", a.id)

	req := &networkservicespb.GetLbRouteExtensionRequest{Name: a.id.String()}
	lbrouteextensionpb, err := a.gcpClient.GetLbRouteExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LBRouteExtension %q: %w", a.id, err)
	}

	a.actual = lbrouteextensionpb
	a.normalizeActual(a.actual)
	return true, nil
}

func (a *LBRouteExtensionAdapter) normalizeActual(obj *networkservicespb.LbRouteExtension) {
	if obj == nil {
		return
	}
	projectID := a.id.Parent().ProjectID
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

func (a *LBRouteExtensionAdapter) normalizeURL(url string, projectID string) string {
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

func (a *LBRouteExtensionAdapter) resolve(ctx context.Context) (*networkservicespb.LbRouteExtension, error) {
	reader := a.reader
	desired := a.desired
	projectID := a.id.Parent().ProjectID

	// Resolve references
	for _, ref := range desired.Spec.ForwardingRuleRefs {
		if ref == nil {
			continue
		}
		if err := ref.Normalize(ctx, reader, desired.GetNamespace()); err != nil {
			return nil, fmt.Errorf("resolving forwardingRuleRef: %w", err)
		}
		// GCP LBRouteExtension returns full URLs for forwarding rules.
		// ForwardingRuleRef.Normalize strips the prefix, so we must add it back.
		if ref.External != "" && !strings.HasPrefix(ref.External, "https://") {
			ref.External = "https://www.googleapis.com/compute/v1/" + ref.External
		}

		// Ensure the forwarding rule is in the same project as the LBRouteExtension.
		if refProject := common.ExtractProjectID(ref.External); refProject == "" || refProject != projectID {
			return nil, fmt.Errorf("cross-project references are not supported for LBRouteExtension: forwardingRule %q is in project %q, but LBRouteExtension is in project %q", ref.External, refProject, projectID)
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

				// Ensure the backend service is in the same project as the LBRouteExtension.
				if refProject := common.ExtractProjectID(external); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBRouteExtension: backendService %q is in project %q, but LBRouteExtension is in project %q", external, refProject, projectID)
				}
			}
			if extension.WasmPluginRef != nil {
				if err := extension.WasmPluginRef.Normalize(ctx, reader, desired.GetNamespace()); err != nil {
					return nil, fmt.Errorf("resolving wasmPluginRef: %w", err)
				}
				// Ensure the wasm plugin is in the same project as the LBRouteExtension.
				if refProject := common.ExtractProjectID(extension.WasmPluginRef.External); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBRouteExtension: wasmPlugin %q is in project %q, but LBRouteExtension is in project %q", extension.WasmPluginRef.External, refProject, projectID)
				}
			}
		}
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesLBRouteExtensionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set GCP Labels
	desiredProto.Labels = a.desiredProto.Labels

	return desiredProto, nil
}

func (a *LBRouteExtensionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LBRouteExtension", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.Clone(desiredProto).(*networkservicespb.LbRouteExtension)
	resource.Name = a.id.String()

	req := &networkservicespb.CreateLbRouteExtensionRequest{
		Parent:             a.id.Parent().String(),
		LbRouteExtensionId: a.id.ID(),
		LbRouteExtension:   resource,
	}
	op, err := a.gcpClient.CreateLbRouteExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LBRouteExtension %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("LBRouteExtension %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created LBRouteExtension", "name", a.id)

	status := &krm.NetworkServicesLBRouteExtensionStatus{}
	status.ObservedState = NetworkServicesLBRouteExtensionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *LBRouteExtensionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LBRouteExtension", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.Clone(desiredProto).(*networkservicespb.LbRouteExtension)
	resource.Name = a.id.String()

	diff, err := common.CompareProtoMessage(a.actual, resource, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing LBRouteExtension %s: %w", a.id, err)
	}

	updated := a.actual
	if diff.Len() == 0 {
		log.V(2).Info("no changes detected for LBRouteExtension", "name", a.id)
	} else {
		sortedPaths := diff.UnsortedList()
		sort.Strings(sortedPaths)
		updateMask := &fieldmaskpb.FieldMask{Paths: sortedPaths}

		req := &networkservicespb.UpdateLbRouteExtensionRequest{
			UpdateMask:       updateMask,
			LbRouteExtension: resource,
		}
		op, err := a.gcpClient.UpdateLbRouteExtension(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LBRouteExtension %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("LBRouteExtension %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated LBRouteExtension", "name", a.id)
	}

	status := &krm.NetworkServicesLBRouteExtensionStatus{}
	status.ObservedState = NetworkServicesLBRouteExtensionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *LBRouteExtensionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesLBRouteExtension{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesLBRouteExtensionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkServicesLBRouteExtensionGVK)

	u.Object = uObj
	return u, nil
}

func (a *LBRouteExtensionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LBRouteExtension", "name", a.id)

	req := &networkservicespb.DeleteLbRouteExtensionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteLbRouteExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LBRouteExtension, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting LBRouteExtension %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LBRouteExtension", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete LBRouteExtension %s: %w", a.id, err)
	}
	return true, nil
}
