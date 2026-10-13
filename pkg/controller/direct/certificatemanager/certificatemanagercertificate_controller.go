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

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerCertificateGVK, NewCertificateModel)
}

func NewCertificateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &certificateModel{config: *config}, nil
}

type certificateModel struct {
	config config.ControllerConfig
}

func (m *certificateModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %w", err)
	}
	return gcpClient, nil
}

func resolveCertificateSecrets(ctx context.Context, kube client.Reader, obj *krm.CertificateManagerCertificate) error {
	if obj.Spec.SelfManaged == nil {
		return nil
	}

	resolveSecretVal := func(secretKeyRef *k8sv1alpha1.SecretKeyRef) (*string, error) {
		if secretKeyRef == nil {
			return nil, nil
		}
		key := types.NamespacedName{
			Namespace: obj.Namespace,
			Name:      secretKeyRef.Name,
		}
		secret := &corev1.Secret{}
		if err := kube.Get(ctx, key, secret); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, k8s.NewSecretNotFoundError(key)
			}
			return nil, fmt.Errorf("error reading referenced Secret %v: %w", key, err)
		}
		valBytes, ok := secret.Data[secretKeyRef.Key]
		if !ok {
			return nil, fmt.Errorf("key %q not found in Secret %v", secretKeyRef.Key, key)
		}
		val := string(valBytes)
		return &val, nil
	}

	if obj.Spec.SelfManaged.PemPrivateKey != nil && obj.Spec.SelfManaged.PemPrivateKey.ValueFrom != nil {
		if obj.Spec.SelfManaged.PemPrivateKey.Value != nil {
			return fmt.Errorf("cannot specify both spec.selfManaged.pemPrivateKey.value and spec.selfManaged.pemPrivateKey.valueFrom")
		}
		val, err := resolveSecretVal(obj.Spec.SelfManaged.PemPrivateKey.ValueFrom.SecretKeyRef)
		if err != nil {
			return fmt.Errorf("resolving spec.selfManaged.pemPrivateKey: %w", err)
		}
		obj.Spec.SelfManaged.PemPrivateKey.Value = val
		obj.Spec.SelfManaged.PemPrivateKey.ValueFrom = nil
	}

	if obj.Spec.SelfManaged.PrivateKeyPem != nil && obj.Spec.SelfManaged.PrivateKeyPem.ValueFrom != nil {
		if obj.Spec.SelfManaged.PrivateKeyPem.Value != nil {
			return fmt.Errorf("cannot specify both spec.selfManaged.privateKeyPem.value and spec.selfManaged.privateKeyPem.valueFrom")
		}
		val, err := resolveSecretVal(obj.Spec.SelfManaged.PrivateKeyPem.ValueFrom.SecretKeyRef)
		if err != nil {
			return fmt.Errorf("resolving spec.selfManaged.privateKeyPem: %w", err)
		}
		obj.Spec.SelfManaged.PrivateKeyPem.Value = val
		obj.Spec.SelfManaged.PrivateKeyPem.ValueFrom = nil
	}

	return nil
}

func (m *certificateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerCertificate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if err := resolveCertificateSecrets(ctx, reader, obj); err != nil {
		return nil, err
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := CertificateManagerCertificateSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &CertificateAdapter{
		id:        identity.(*krm.CertificateManagerCertificateIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *certificateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CertificateManagerCertificateIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &CertificateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type CertificateAdapter struct {
	id        *krm.CertificateManagerCertificateIdentity
	gcpClient *gcp.Client
	desired   *pb.Certificate
	actual    *pb.Certificate
}

var _ directbase.Adapter = &CertificateAdapter{}

func (a *CertificateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CertificateManagerCertificate", "name", a.id.String())

	req := &pb.GetCertificateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CertificateManagerCertificate %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *CertificateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating CertificateManagerCertificate", "id", fqn)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateCertificateRequest{
		Parent:        parent,
		Certificate:   a.desired,
		CertificateId: a.id.Certificate,
	}
	op, err := a.gcpClient.CreateCertificate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Certificate %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Certificate %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Certificate", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CertificateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CertificateManagerCertificate", "name", a.id.String())

	diffs, updateMask, err := compareCertificate(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.CloneOf(a.desired)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateCertificateRequest{
			Certificate: desiredCopy,
			UpdateMask:  updateMask,
		}

		op, err := a.gcpClient.UpdateCertificate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CertificateManagerCertificate %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting CertificateManagerCertificate %s update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CertificateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Certificate) error {
	mapCtx := &direct.MapContext{}
	observedState := CertificateObservedStateStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.CertificateManagerCertificateStatus{}
	status.ObservedState = observedState

	return op.UpdateStatus(ctx, status, nil)
}

func (a *CertificateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerCertificate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CertificateManagerCertificateSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CertificateManagerCertificateGVK)
	return u, nil
}

func (a *CertificateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Certificate", "name", a.id.String())

	req := &pb.DeleteCertificateRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Certificate, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Certificate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Certificate", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Certificate %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareCertificate(ctx context.Context, actual, desired *pb.Certificate) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CertificateManagerCertificateSpec_v1beta1_FromProto, CertificateManagerCertificateSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	// The GCP Certificate resource type (managed vs self_managed) is immutable.
	// For self_managed certificates, GCP never returns the private key (pemPrivateKey) on GET calls.
	// Normalize maskedActual's pemPrivateKey to match desired so we don't trigger a false diff on the secret,
	// and avoid mutating desired in-place.
	if desiredSelfManaged := desired.GetSelfManaged(); desiredSelfManaged != nil {
		if actualSelfManaged := maskedActual.GetSelfManaged(); actualSelfManaged != nil {
			actualSelfManaged.PemPrivateKey = desiredSelfManaged.PemPrivateKey
		}
	}

	// For managed certificates, GCP populates server-computed state, authorization attempt info, etc.
	// Ensure that actual and desired managed fields only compare user-specified spec fields.
	if desiredManaged := desired.GetManaged(); desiredManaged != nil {
		if actualManaged := maskedActual.GetManaged(); actualManaged != nil {
			// If IssuanceConfig was not specified in desired KRM, normalize maskedActual to match desired.
			if desiredManaged.IssuanceConfig == "" {
				actualManaged.IssuanceConfig = desiredManaged.IssuanceConfig
			}
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
