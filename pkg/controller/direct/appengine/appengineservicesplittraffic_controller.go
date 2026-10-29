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

// +tool:controller
// proto.service: google.appengine.v1.Services
// proto.message: google.appengine.v1.Service
// crd.type: AppEngineServiceSplitTraffic
// crd.version: v1alpha1

package appengine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	api "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AppEngineServiceSplitTrafficGVK, NewAppEngineServiceSplitTrafficModel, registry.CannotBeDeleted())
}

func NewAppEngineServiceSplitTrafficModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &splitTrafficModel{config: *config}, nil
}

type splitTrafficModel struct {
	config config.ControllerConfig
}

var _ directbase.Model = &splitTrafficModel{}

func (m *splitTrafficModel) client(ctx context.Context) (*api.ServicesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building appengine services client: %w", err)
	}

	return gcpClient, nil
}

func (m *splitTrafficModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineServiceSplitTraffic{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	trafficId := id.(*krm.AppEngineServiceSplitTrafficIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &splitTrafficAdapter{
		gcpClient: gcpClient,
		id:        trafficId,
		desired:   obj,
	}, nil
}

func (m *splitTrafficModel) AdapterForURL(ctx context.Context, rawURL string) (directbase.Adapter, error) {
	parsedURL, err := url.Parse(rawURL)
	if err == nil && parsedURL.Path != "" {
		rawURL = strings.TrimPrefix(parsedURL.Path, "/v1/")
		rawURL = strings.TrimPrefix(rawURL, "/")
	}
	id := &krm.AppEngineServiceSplitTrafficIdentity{}
	if err := id.FromExternal(rawURL); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &splitTrafficAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type splitTrafficAdapter struct {
	id        *krm.AppEngineServiceSplitTrafficIdentity
	gcpClient *api.ServicesClient
	desired   *krm.AppEngineServiceSplitTraffic
	actual    *pb.Service
}

var _ directbase.Adapter = &splitTrafficAdapter{}

func (a *splitTrafficAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AppEngineServiceSplitTraffic", "name", a.id.String())

	req := &pb.GetServiceRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AppEngineService %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *splitTrafficAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("cannot create AppEngineServiceSplitTraffic: parent AppEngine service %q does not exist; an App Engine service must be created by deploying an AppEngineVersion before configuring split traffic", a.id.String())
}

func (a *splitTrafficAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineServiceSplitTraffic", "name", a.id.String())

	mburMatch, err := updateOp.CompareMBUR(a.computeMBUR)
	if err != nil {
		return fmt.Errorf("comparing MBUR: %w", err)
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto,
		AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto,
		nil,
		false,
	)
	if err != nil {
		return err
	}

	if diffRes.Empty() && mburMatch {
		log.V(2).Info("no diff detected for AppEngineServiceSplitTraffic", "name", a.id.String())
		return a.updateStatus(ctx, updateOp)
	}

	if !mburMatch && diffRes.Empty() {
		diffRes.Diff.AddField("migrate_traffic", nil, direct.ValueOf(a.desired.Spec.MigrateTraffic))
	}

	diffRes.Diff.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffRes.Diff)

	if err := a.updateService(ctx, updateOp, diffRes.MergedDesired); err != nil {
		return err
	}

	if err := updateOp.SetMBUR(ctx, a.computeMBUR); err != nil {
		return fmt.Errorf("setting MBUR: %w", err)
	}

	return nil
}

func (a *splitTrafficAdapter) computeMBUR() (string, error) {
	if a.desired.Spec.MigrateTraffic == nil {
		return "", nil
	}
	bytes, err := json.Marshal(a.desired.Spec.MigrateTraffic)
	if err != nil {
		return "", fmt.Errorf("marshalling migrateTraffic: %w", err)
	}
	h := sha256.Sum256(bytes)
	return hex.EncodeToString(h[:]), nil
}

func (a *splitTrafficAdapter) updateService(ctx context.Context, op directbase.Operation, desiredPb *pb.Service) error {
	log := klog.FromContext(ctx)
	desiredPb.Id = a.id.Service

	req := &pb.UpdateServiceRequest{
		Name:           a.id.String(),
		Service:        desiredPb,
		UpdateMask:     &fieldmaskpb.FieldMask{Paths: []string{"split"}},
		MigrateTraffic: direct.ValueOf(a.desired.Spec.MigrateTraffic),
	}

	opLRO, err := a.gcpClient.UpdateService(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineServiceSplitTraffic %s: %w", a.id.String(), err)
	}

	latest, err := opLRO.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineServiceSplitTraffic %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated AppEngineServiceSplitTraffic", "name", a.id.String())

	// Fetch fully-populated service
	getReq := &pb.GetServiceRequest{Name: a.id.String()}
	refetched, err := a.gcpClient.GetService(ctx, getReq)
	if err != nil {
		refetched = latest
	}
	a.actual = refetched

	return a.updateStatus(ctx, op)
}

func (a *splitTrafficAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineServiceSplitTraffic is a no-op because split traffic cannot be deleted in GCP", "name", a.id.String())
	return true, nil
}

func (a *splitTrafficAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppEngineServiceSplitTraffic{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Project = &a.id.App
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Service)
	obj.Status.ExternalRef = direct.LazyPtr(a.id.String())

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Service)
	u.SetGroupVersionKind(krm.AppEngineServiceSplitTrafficGVK)

	return u, nil
}

func (a *splitTrafficAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	status := &krm.AppEngineServiceSplitTrafficStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
