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

package certificatemanager

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	directcommon "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerCertificateIssuanceConfigGVK, NewCertificateIssuanceConfigModel)
}

func NewCertificateIssuanceConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateIssuanceConfigModel{config: *config}, nil
}

var _ directbase.Model = &certificateIssuanceConfigModel{}

type certificateIssuanceConfigModel struct {
	config config.ControllerConfig
}

func (m *certificateIssuanceConfigModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CertificateIssuanceConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *certificateIssuanceConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerCertificateIssuanceConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve resource references using standard directcommon.NormalizeReferences
	if err := directcommon.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.CertificateManagerCertificateIssuanceConfigIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &CertificateIssuanceConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *certificateIssuanceConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type CertificateIssuanceConfigAdapter struct {
	id        *krm.CertificateManagerCertificateIssuanceConfigIdentity
	gcpClient *gcp.Client
	desired   *pb.CertificateIssuanceConfig
	actual    *pb.CertificateIssuanceConfig
}

var _ directbase.Adapter = &CertificateIssuanceConfigAdapter{}

func (a *CertificateIssuanceConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateIssuanceConfig", "name", a.id)

	req := &pb.GetCertificateIssuanceConfigRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetCertificateIssuanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateIssuanceConfig %q: %w", a.id, err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *CertificateIssuanceConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CertificateIssuanceConfig", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateCertificateIssuanceConfigRequest{
		Parent:                      parent,
		CertificateIssuanceConfigId: a.id.CertificateIssuanceConfig,
		CertificateIssuanceConfig:   a.desired,
	}
	op, err := a.gcpClient.CreateCertificateIssuanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CertificateIssuanceConfig %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting CertificateIssuanceConfig %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created CertificateIssuanceConfig", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateIssuanceConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateIssuanceConfig", "name", a.id.String())

	diffs, _, err := compareCertificateIssuanceConfig(ctx, a.actual, a.desired, a.id.Project)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		return fmt.Errorf("CertificateIssuanceConfig is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *CertificateIssuanceConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificateIssuanceConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{
		External: a.id.Project,
	}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.CertificateIssuanceConfig)
	u.SetGroupVersionKind(krm.CertificateManagerCertificateIssuanceConfigGVK)

	u.Object = uObj
	return u, nil
}

func (a *CertificateIssuanceConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CertificateIssuanceConfig", "name", a.id)

	req := &pb.DeleteCertificateIssuanceConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificateIssuanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CertificateIssuanceConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CertificateIssuanceConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CertificateIssuanceConfig", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CertificateIssuanceConfig %s: %w", a.id, err)
	}
	return true, nil
}

func (a *CertificateIssuanceConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CertificateIssuanceConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := CertificateManagerCertificateIssuanceConfigObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status := &krm.CertificateManagerCertificateIssuanceConfigStatus{
		ObservedState: observedState,
		ExternalRef:   direct.LazyPtr(a.id.String()),
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareCertificateIssuanceConfig(ctx context.Context, actual, desired *pb.CertificateIssuanceConfig, projectID string) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *pb.CertificateIssuanceConfig
	{
		mapCtx := &direct.MapContext{}
		spec := CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = CertificateManagerCertificateIssuanceConfigSpec_v1alpha1_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}

	// Normalize CAPool names to ignore project ID vs project number differences in diff.
	if config := desired.GetCertificateAuthorityConfig().GetCertificateAuthorityServiceConfig(); config != nil {
		config.CaPool = normalizeCAPool(config.CaPool, projectID)
	}
	if config := maskedActual.GetCertificateAuthorityConfig().GetCertificateAuthorityServiceConfig(); config != nil {
		config.CaPool = normalizeCAPool(config.CaPool, projectID)
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func normalizeCAPool(caPool string, projectID string) string {
	tokens := strings.Split(caPool, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		tokens[1] = projectID
	}
	return strings.Join(tokens, "/")
}
