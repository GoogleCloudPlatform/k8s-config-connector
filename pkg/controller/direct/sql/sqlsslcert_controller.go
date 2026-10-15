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

package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func init() {
	registry.RegisterModel(krm.SQLSSLCertGVK, newSQLSSLCertModel)
}

func newSQLSSLCertModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sqlSSLCertModel{config: config}, nil
}

type sqlSSLCertModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &sqlSSLCertModel{}

func (m *sqlSSLCertModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	obj := &krm.SQLSSLCert{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("converting to %T failed: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	typedID, ok := id.(*krm.SQLSSLCertIdentity)
	if !ok {
		return nil, fmt.Errorf("expected SQLSSLCertIdentity, got %T", id)
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	return &sqlSSLCertAdapter{
		id:                  typedID,
		gcpClient:           gcpClient,
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		sqlSslCertsClient:   gcpClient.sqlSslCertsClient(),
		desired:             obj,
		initialStatus:       &obj.Status,
	}, nil
}

func (m *sqlSSLCertModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.SQLSSLCertIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	return &sqlSSLCertAdapter{
		id:                  id,
		gcpClient:           gcpClient,
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		sqlSslCertsClient:   gcpClient.sqlSslCertsClient(),
	}, nil
}

type sqlSSLCertAdapter struct {
	id                  *krm.SQLSSLCertIdentity
	gcpClient           *gcpClient
	sqlOperationsClient *api.OperationsService
	sqlSslCertsClient   *api.SslCertsService
	desired             *krm.SQLSSLCert
	actual              *pb.SslCert
	initialStatus       *krm.SQLSSLCertStatus
	privateKey          *string
	serverCaCert        *string
}

var _ directbase.Adapter = &sqlSSLCertAdapter{}

// Find retrieves the existing resource state from GCP.
func (a *sqlSSLCertAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard for service-generated ID
	if a.id.Sha1Fingerprint == "" {
		return false, nil
	}

	cert, err := a.sqlSslCertsClient.Get(a.id.Project, a.id.Instance, a.id.Sha1Fingerprint).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SQLSSLCert %q: %w", a.id.String(), err)
	}

	certpb, err := sslCertAPIToProto(cert)
	if err != nil {
		return false, err
	}
	a.actual = certpb
	return true, nil
}

// Create creates a new SQLSSLCert in GCP.
func (a *sqlSSLCertAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SQLSSLCert", "parent", a.id.ParentString())

	// Rule 3: Do not set identifier in request payload
	req := &api.SslCertsInsertRequest{
		CommonName: a.desired.Spec.CommonName,
	}

	resp, err := a.sqlSslCertsClient.Insert(a.id.Project, a.id.Instance, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating SQLSSLCert in %s: %w", a.id.ParentString(), err)
	}

	if resp.Operation != nil {
		if err := a.pollForLROCompletion(ctx, resp.Operation, "creating"); err != nil {
			return err
		}
	}

	if resp.ClientCert == nil || resp.ClientCert.CertInfo == nil {
		return fmt.Errorf("insert response missing client cert info")
	}

	fingerprint := resp.ClientCert.CertInfo.Sha1Fingerprint
	if fingerprint == "" {
		return fmt.Errorf("insert response client cert missing sha1Fingerprint")
	}

	if err := createOp.SetSpecResourceID(ctx, fingerprint); err != nil {
		return err
	}
	a.id.Sha1Fingerprint = fingerprint

	a.privateKey = direct.LazyPtr(resp.ClientCert.CertPrivateKey)
	if resp.ServerCaCert != nil {
		a.serverCaCert = direct.LazyPtr(resp.ServerCaCert.Cert)
	}

	latest, err := a.sqlSslCertsClient.Get(a.id.Project, a.id.Instance, a.id.Sha1Fingerprint).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("fetching newly created SQLSSLCert %q: %w", a.id.String(), err)
	}

	latestpb, err := sslCertAPIToProto(latest)
	if err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, latestpb)
}

// Update handles updating an existing SQLSSLCert. Since SQLSSLCert is immutable in GCP,
// any modifications to spec fields will trigger an error.
func (a *sqlSSLCertAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SQLSSLCert", "name", a.id.String())

	if a.actual == nil {
		return fmt.Errorf("actual state is missing in Update")
	}

	if a.actual.CommonName != a.desired.Spec.CommonName {
		return fmt.Errorf("SQLSSLCert is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

// Delete deletes the SQLSSLCert resource in GCP.
func (a *sqlSSLCertAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SQLSSLCert", "name", a.id.String())

	if a.id.Sha1Fingerprint == "" {
		return true, nil
	}

	op, err := a.sqlSslCertsClient.Delete(a.id.Project, a.id.Instance, a.id.Sha1Fingerprint).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting SQLSSLCert %q: %w", a.id.String(), err)
	}

	if op != nil {
		if err := a.pollForLROCompletion(ctx, op, "deleting"); err != nil {
			return false, err
		}
	}

	return true, nil
}

// Export exports the SQLSSLCert resource to a KRM unstructured object.
func (a *sqlSSLCertAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SQLSSLCert{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SQLSSLCertSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef.External = a.id.ParentString()
	obj.Spec.ResourceID = &a.id.Sha1Fingerprint

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.CommonName)
	u.SetGroupVersionKind(krm.SQLSSLCertGVK)
	return u, nil
}

func (a *sqlSSLCertAdapter) pollForLROCompletion(ctx context.Context, op *api.Operation, verb string) error {
	log := klog.FromContext(ctx)
	var err error

	pollingBackoff := gax.Backoff{
		Initial:    time.Second,
		Max:        time.Minute,
		Multiplier: 2,
	}
	for {
		log.V(2).Info("polling", "op", op)

		if op.Status == "DONE" {
			if op.Error != nil && len(op.Error.Errors) > 0 {
				var errMsgs []string
				for _, opErr := range op.Error.Errors {
					errMsgs = append(errMsgs, opErr.Message)
				}
				return fmt.Errorf("operation failed: %s", strings.Join(errMsgs, "; "))
			}
			break
		}
		if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
			return fmt.Errorf("waiting for SQLSSLCert %s %s failed: %w", a.id.String(), verb, err)
		}
		op, err = a.sqlOperationsClient.Get(a.id.Project, op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting SQLSSLCert %s %s operation %s failed: %w", a.id.String(), verb, op.Name, err)
		}
	}

	return nil
}

func (a *sqlSSLCertAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SslCert) error {
	mapCtx := &direct.MapContext{}
	status := SQLSSLCertStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Preserve privateKey if available from creation or previous status
	if a.privateKey != nil {
		status.PrivateKey = a.privateKey
	} else if a.initialStatus != nil && a.initialStatus.PrivateKey != nil {
		status.PrivateKey = a.initialStatus.PrivateKey
	}

	// Preserve serverCaCert if available from creation or previous status
	if a.serverCaCert != nil {
		status.ServerCaCert = a.serverCaCert
	} else if a.initialStatus != nil && a.initialStatus.ServerCaCert != nil {
		status.ServerCaCert = a.initialStatus.ServerCaCert
	}

	return op.UpdateStatus(ctx, status, nil)
}

func sslCertAPIToProto(in *api.SslCert) (*pb.SslCert, error) {
	if in == nil {
		return nil, nil
	}
	out := &pb.SslCert{
		Kind:             in.Kind,
		CertSerialNumber: in.CertSerialNumber,
		Cert:             in.Cert,
		CommonName:       in.CommonName,
		Sha1Fingerprint:  in.Sha1Fingerprint,
		Instance:         in.Instance,
		SelfLink:         in.SelfLink,
	}
	if in.CreateTime != "" {
		t, err := time.Parse(time.RFC3339Nano, in.CreateTime)
		if err != nil {
			t, err = time.Parse(time.RFC3339, in.CreateTime)
		}
		if err != nil {
			return nil, fmt.Errorf("parsing createTime %q: %w", in.CreateTime, err)
		}
		out.CreateTime = timestamppb.New(t)
	}
	if in.ExpirationTime != "" {
		t, err := time.Parse(time.RFC3339Nano, in.ExpirationTime)
		if err != nil {
			t, err = time.Parse(time.RFC3339, in.ExpirationTime)
		}
		if err != nil {
			return nil, fmt.Errorf("parsing expirationTime %q: %w", in.ExpirationTime, err)
		}
		out.ExpirationTime = timestamppb.New(t)
	}
	return out, nil
}
