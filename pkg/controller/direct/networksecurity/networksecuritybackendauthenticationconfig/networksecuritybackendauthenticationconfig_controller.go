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

package networksecuritybackendauthenticationconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networksecurity/v1"
	api "google.golang.org/api/networksecurity/v1"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityBackendAuthenticationConfigGVK, NewBackendAuthenticationConfigModel)
}

func NewBackendAuthenticationConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackendAuthenticationConfig{config: *config}, nil
}

var _ directbase.Model = &modelBackendAuthenticationConfig{}

type modelBackendAuthenticationConfig struct {
	config config.ControllerConfig
}

func (m *modelBackendAuthenticationConfig) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BackendAuthenticationConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBackendAuthenticationConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityBackendAuthenticationConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackendAuthenticationConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networksecurity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BackendAuthenticationConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBackendAuthenticationConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type BackendAuthenticationConfigAdapter struct {
	id        *krm.BackendAuthenticationConfigIdentity
	gcpClient *api.Service
	desired   *krm.NetworkSecurityBackendAuthenticationConfig
	actual    *pb.BackendAuthenticationConfig
}

var _ directbase.Adapter = &BackendAuthenticationConfigAdapter{}

// Find retrieves the GCP resource.
func (a *BackendAuthenticationConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackendAuthenticationConfig", "name", a.id)

	actual, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackendAuthenticationConfig %q: %w", a.id, err)
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the status.
func (a *BackendAuthenticationConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackendAuthenticationConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityBackendAuthenticationConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &api.BackendAuthenticationConfig{}
	if err := convertProtoToAPI(resource, req); err != nil {
		return err
	}

	op, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Create(a.id.Parent().String(), req).BackendAuthenticationConfigId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating BackendAuthenticationConfig %s: %w", a.id, err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("BackendAuthenticationConfig %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackendAuthenticationConfig", "name", a.id)

	created, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created BackendAuthenticationConfig %q: %w", a.id, err)
	}

	var createdPB *pb.BackendAuthenticationConfig
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}

	status := &krm.NetworkSecurityBackendAuthenticationConfigStatus{}
	status.ObservedState = NetworkSecurityBackendAuthenticationConfigObservedState_FromProto(mapCtx, createdPB)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP.
func (a *BackendAuthenticationConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackendAuthenticationConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := NetworkSecurityBackendAuthenticationConfigSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	paths := []string{}
	if desiredPb.Description != "" && !reflect.DeepEqual(desiredPb.Description, a.actual.Description) {
		report.AddField("description", a.actual.Description, desiredPb.Description)
		paths = append(paths, "description")
	}
	if !reflect.DeepEqual(desiredPb.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, desiredPb.Labels)
		paths = append(paths, "labels")
	}
	if desiredPb.ClientCertificate != "" && !reflect.DeepEqual(desiredPb.ClientCertificate, a.actual.ClientCertificate) {
		report.AddField("client_certificate", a.actual.ClientCertificate, desiredPb.ClientCertificate)
		paths = append(paths, "client_certificate")
	}
	if desiredPb.TrustConfig != "" && !reflect.DeepEqual(desiredPb.TrustConfig, a.actual.TrustConfig) {
		report.AddField("trust_config", a.actual.TrustConfig, desiredPb.TrustConfig)
		paths = append(paths, "trust_config")
	}
	if desiredPb.WellKnownRoots != a.actual.WellKnownRoots {
		report.AddField("well_known_roots", a.actual.WellKnownRoots, desiredPb.WellKnownRoots)
		paths = append(paths, "well_known_roots")
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		structuredreporting.ReportDiff(ctx, report)

		req := &api.BackendAuthenticationConfig{}
		if err := convertProtoToAPI(desiredPb, req); err != nil {
			return err
		}

		op, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Patch(a.id.String(), req).UpdateMask(strings.Join(paths, ",")).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating BackendAuthenticationConfig %s: %w", a.id, err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("BackendAuthenticationConfig %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated BackendAuthenticationConfig", "name", a.id)

		updatedAPI, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Get(a.id.String()).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated BackendAuthenticationConfig %q: %w", a.id, err)
		}
		if err := convertAPIToProto(updatedAPI, &updated); err != nil {
			return err
		}
	}

	status := &krm.NetworkSecurityBackendAuthenticationConfigStatus{}
	status.ObservedState = NetworkSecurityBackendAuthenticationConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackendAuthenticationConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityBackendAuthenticationConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityBackendAuthenticationConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = &a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkSecurityBackendAuthenticationConfigGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service.
func (a *BackendAuthenticationConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackendAuthenticationConfig", "name", a.id)

	op, err := a.gcpClient.Projects.Locations.BackendAuthenticationConfigs.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackendAuthenticationConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackendAuthenticationConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackendAuthenticationConfig", "name", a.id)

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting delete BackendAuthenticationConfig %s: %w", a.id, err)
	}
	return true, nil
}

func (a *BackendAuthenticationConfigAdapter) waitForOperation(ctx context.Context, op *api.Operation) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		latest, err := a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}

		if latest.Done {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}

func convertProtoToAPI(u protoreflect.ProtoMessage, v any) error {
	if u == nil {
		return nil
	}

	j, err := protojson.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}

	if err := json.Unmarshal(j, v); err != nil {
		return fmt.Errorf("converting json to cloud API type: %w", err)
	}
	return nil
}

func convertAPIToProto[V protoreflect.ProtoMessage](u any, pV *V) error {
	if u == nil {
		return nil
	}

	j, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting API to json: %w", err)
	}

	var zero V
	messageType := reflect.TypeOf(zero).Elem()
	val := reflect.New(messageType).Interface().(protoreflect.ProtoMessage)

	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(j, val); err != nil {
		return fmt.Errorf("converting json to proto type: %w", err)
	}
	*pV = val.(V)
	return nil
}
