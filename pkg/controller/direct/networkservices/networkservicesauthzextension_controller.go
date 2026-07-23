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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesAuthzExtensionGVK, NewAuthzExtensionModel)
}

func NewAuthzExtensionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAuthzExtension{config: *config}, nil
}

var _ directbase.Model = &modelAuthzExtension{}

type modelAuthzExtension struct {
	config config.ControllerConfig
}

func (m *modelAuthzExtension) client(ctx context.Context) (*gcp.DepClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDepRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AuthzExtension client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAuthzExtension) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesAuthzExtension{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewAuthzExtensionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesAuthzExtensionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &AuthzExtensionAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desired:      obj,
		reader:       reader,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelAuthzExtension) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AuthzExtensionAdapter struct {
	id           *krm.AuthzExtensionIdentity
	gcpClient    *gcp.DepClient
	desired      *krm.NetworkServicesAuthzExtension
	reader       client.Reader
	actual       *pb.AuthzExtension
	desiredProto *pb.AuthzExtension
}

var _ directbase.Adapter = &AuthzExtensionAdapter{}

func (a *AuthzExtensionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AuthzExtension", "name", a.id)

	req := &pb.GetAuthzExtensionRequest{Name: a.id.String()}
	authzextensionpb, err := a.gcpClient.GetAuthzExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AuthzExtension %q: %w", a.id, err)
	}

	a.actual = authzextensionpb
	a.normalizeActual(a.actual)
	return true, nil
}

func (a *AuthzExtensionAdapter) resolve(ctx context.Context) (*pb.AuthzExtension, error) {
	reader := a.reader
	desired := a.desired
	projectID := a.id.Project

	if desired.Spec.BackendServiceRef != nil {
		external, err := desired.Spec.BackendServiceRef.NormalizedExternal(ctx, reader, desired.GetNamespace())
		if err != nil {
			return nil, fmt.Errorf("resolving backendServiceRef: %w", err)
		}
		desired.Spec.BackendServiceRef.External = external

		// Ensure the backend service is in the same project as the AuthzExtension.
		if refProject := common.ExtractProjectID(external); refProject == "" || refProject != projectID {
			return nil, fmt.Errorf("cross-project references are not supported for AuthzExtension: backendService %q is in project %q, but AuthzExtension is in project %q", external, refProject, projectID)
		}
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesAuthzExtensionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set GCP Labels
	desiredProto.Labels = a.desiredProto.Labels

	return desiredProto, nil
}

func (a *AuthzExtensionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AuthzExtension", "name", a.id)

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}

	resource := proto.CloneOf(desiredProto)
	resource.Name = a.id.String()

	req := &pb.CreateAuthzExtensionRequest{
		Parent:           a.id.ParentString(),
		AuthzExtensionId: a.id.AuthzExtension,
		AuthzExtension:   resource,
	}
	op, err := a.gcpClient.CreateAuthzExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AuthzExtension %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AuthzExtension %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AuthzExtension", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *AuthzExtensionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AuthzExtension", "name", a.id)

	desiredProto, err := a.resolve(ctx)
	if err != nil {
		return err
	}
	desiredProto.Name = a.id.String()

	diffs, updateMask, err := compareAuthzExtension(ctx, a.actual, desiredProto)
	if err != nil {
		return fmt.Errorf("comparing AuthzExtension %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for AuthzExtension", "name", a.id)
	} else {
		// Report exact diffs
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateAuthzExtensionRequest{
			UpdateMask:     updateMask,
			AuthzExtension: desiredProto,
		}
		op, err := a.gcpClient.UpdateAuthzExtension(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AuthzExtension %s: %w", a.id, err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("AuthzExtension %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated AuthzExtension", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AuthzExtensionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesAuthzExtension{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesAuthzExtensionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.AuthzExtension)
	u.SetGroupVersionKind(krm.NetworkServicesAuthzExtensionGVK)

	u.Object = uObj
	return u, nil
}

func (a *AuthzExtensionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AuthzExtension", "name", a.id)

	req := &pb.DeleteAuthzExtensionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAuthzExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent AuthzExtension, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AuthzExtension %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AuthzExtension", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete AuthzExtension %s: %w", a.id, err)
	}
	return true, nil
}

func compareAuthzExtension(ctx context.Context, actual, desired *pb.AuthzExtension) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesAuthzExtensionSpec_FromProto, NetworkServicesAuthzExtensionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.AuthzExtension) {
		// Add any server-side or GCP defaults if known, or leave empty
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *AuthzExtensionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AuthzExtension) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkServicesAuthzExtensionStatus{}
	status.ObservedState = NetworkServicesAuthzExtensionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AuthzExtensionAdapter) normalizeActual(obj *pb.AuthzExtension) {
	if obj == nil {
		return
	}
	projectID := a.id.Project
	if obj.Service != "" {
		obj.Service = a.normalizeURL(obj.Service, projectID)
	}
}

func (a *AuthzExtensionAdapter) normalizeURL(url string, projectID string) string {
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
