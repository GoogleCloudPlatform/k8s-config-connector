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

	gcp "cloud.google.com/go/networkservices/apiv1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
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
	registry.RegisterModel(krm.NetworkServicesLBTrafficExtensionGVK, NewLBTrafficExtensionModel)
}

func NewLBTrafficExtensionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLBTrafficExtension{config: *config}, nil
}

var _ directbase.Model = &modelLBTrafficExtension{}

type modelLBTrafficExtension struct {
	config config.ControllerConfig
}

func (m *modelLBTrafficExtension) client(ctx context.Context) (*gcp.DepClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDepRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building LBTrafficExtension client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLBTrafficExtension) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesLBTrafficExtension{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewLBTrafficExtensionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesLBTrafficExtensionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &LBTrafficExtensionAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desired:      obj,
		reader:       reader,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelLBTrafficExtension) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LBTrafficExtensionAdapter struct {
	id           *krm.LBTrafficExtensionIdentity
	gcpClient    *gcp.DepClient
	desired      *krm.NetworkServicesLBTrafficExtension
	reader       client.Reader
	actual       *pb.LbTrafficExtension
	desiredProto *pb.LbTrafficExtension
}

var _ directbase.Adapter = &LBTrafficExtensionAdapter{}

func (a *LBTrafficExtensionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LBTrafficExtension", "name", a.id)

	req := &pb.GetLbTrafficExtensionRequest{Name: a.id.String()}
	lbtrafficextensionpb, err := a.gcpClient.GetLbTrafficExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LBTrafficExtension %q: %w", a.id, err)
	}

	a.actual = lbtrafficextensionpb
	a.normalizeActual(a.actual)
	return true, nil
}

func (a *LBTrafficExtensionAdapter) normalizeActual(obj *pb.LbTrafficExtension) {
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

func (a *LBTrafficExtensionAdapter) normalizeURL(url string, projectID string) string {
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

func (a *LBTrafficExtensionAdapter) resolve(ctx context.Context) (*pb.LbTrafficExtension, error) {
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
		// GCP LBTrafficExtension returns full URLs for forwarding rules.
		// ForwardingRuleRef.Normalize strips the prefix, so we must add it back.
		if ref.External != "" && !strings.HasPrefix(ref.External, "https://") {
			ref.External = "https://www.googleapis.com/compute/v1/" + ref.External
		}

		// Ensure the forwarding rule is in the same project as the LBTrafficExtension.
		if refProject := common.ExtractProjectID(ref.External); refProject == "" || refProject != projectID {
			return nil, fmt.Errorf("cross-project references are not supported for LBTrafficExtension: forwardingRule %q is in project %q, but LBTrafficExtension is in project %q", ref.External, refProject, projectID)
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

				// Ensure the backend service is in the same project as the LBTrafficExtension.
				if refProject := common.ExtractProjectID(external); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBTrafficExtension: backendService %q is in project %q, but LBTrafficExtension is in project %q", external, refProject, projectID)
				}
			}
			if extension.WasmPluginRef != nil {
				if err := extension.WasmPluginRef.Normalize(ctx, reader, desired.GetNamespace()); err != nil {
					return nil, fmt.Errorf("resolving wasmPluginRef: %w", err)
				}
				// Ensure the wasm plugin is in the same project as the LBTrafficExtension.
				if refProject := common.ExtractProjectID(extension.WasmPluginRef.External); refProject == "" || refProject != projectID {
					return nil, fmt.Errorf("cross-project references are not supported for LBTrafficExtension: wasmPlugin %q is in project %q, but LBTrafficExtension is in project %q", extension.WasmPluginRef.External, refProject, projectID)
				}
			}
		}
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesLBTrafficExtensionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set GCP Labels
	desiredProto.Labels = a.desiredProto.Labels

	return desiredProto, nil
}

func (a *LBTrafficExtensionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LBTrafficExtension", "name", a.id)

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.CloneOf(desiredProto)
	resource.Name = a.id.String()

	req := &pb.CreateLbTrafficExtensionRequest{
		Parent:               a.id.ParentString(),
		LbTrafficExtensionId: a.id.LbTrafficExtension,
		LbTrafficExtension:   resource,
	}
	op, err := a.gcpClient.CreateLbTrafficExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LBTrafficExtension %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("LBTrafficExtension %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created LBTrafficExtension", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *LBTrafficExtensionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LBTrafficExtension", "name", a.id)

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}
	desiredProto.Name = a.id.String()

	diffs, updateMask, err := compareLBTrafficExtension(ctx, a.actual, desiredProto)
	if err != nil {
		return fmt.Errorf("comparing LBTrafficExtension %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for LBTrafficExtension", "name", a.id)
	} else {
		// Report exact diffs
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateLbTrafficExtensionRequest{
			UpdateMask:         updateMask,
			LbTrafficExtension: desiredProto,
		}
		op, err := a.gcpClient.UpdateLbTrafficExtension(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LBTrafficExtension %s: %w", a.id, err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("LBTrafficExtension %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated LBTrafficExtension", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LBTrafficExtensionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesLBTrafficExtension{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesLBTrafficExtensionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.LbTrafficExtension)
	u.SetGroupVersionKind(krm.NetworkServicesLBTrafficExtensionGVK)

	u.Object = uObj
	return u, nil
}

func (a *LBTrafficExtensionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LBTrafficExtension", "name", a.id)

	req := &pb.DeleteLbTrafficExtensionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteLbTrafficExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LBTrafficExtension, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting LBTrafficExtension %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LBTrafficExtension", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete LBTrafficExtension %s: %w", a.id, err)
	}
	return true, nil
}

func compareLBTrafficExtension(ctx context.Context, actual, desired *pb.LbTrafficExtension) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesLBTrafficExtensionSpec_FromProto, NetworkServicesLBTrafficExtensionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.LbTrafficExtension) {
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

func (a *LBTrafficExtensionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.LbTrafficExtension) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkServicesLBTrafficExtensionStatus{}
	status.ObservedState = NetworkServicesLBTrafficExtensionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
