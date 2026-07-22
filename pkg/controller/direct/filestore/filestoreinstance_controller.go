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

package filestore

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/filestore/apiv1"
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.FilestoreInstanceGVK, newFilestoreInstanceModel)
}

func newFilestoreInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &filestoreInstanceModel{config: config}, nil
}

type filestoreInstanceModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &filestoreInstanceModel{}

type filestoreInstanceAdapter struct {
	id *krm.FilestoreInstanceIdentity

	desired *krm.FilestoreInstance
	actual  *pb.Instance

	reader client.Reader
	client *gcp.CloudFilestoreManagerClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &filestoreInstanceAdapter{}

func (m *filestoreInstanceModel) client(ctx context.Context) (*gcp.CloudFilestoreManagerClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudFilestoreManagerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FilestoreInstance client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *filestoreInstanceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.FilestoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	return &filestoreInstanceAdapter{
		id:      id.(*krm.FilestoreInstanceIdentity),
		desired: obj,
		reader:  kube,
		client:  gcpClient,
	}, nil
}

func (m *filestoreInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *filestoreInstanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Instance == "" {
		return false, nil
	}

	filestoreInstance, err := a.getActual(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = filestoreInstance

	return true, nil
}

func (a *filestoreInstanceAdapter) getActual(ctx context.Context) (*pb.Instance, error) {
	req := &pb.GetInstanceRequest{
		Name: a.id.String(),
	}
	return a.client.GetInstance(ctx, req)
}

// Delete implements the Adapter interface.
func (a *filestoreInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &pb.DeleteInstanceRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting filestoreInstance %s: %w", a.id.String(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for filestoreInstance delete %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *filestoreInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *filestoreInstanceAdapter) normalizeReferences(ctx context.Context) error {
	if err := common.NormalizeReferences(ctx, a.reader, a.desired, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}
	return nil
}

// Create implements the Adapter interface.
func (a *filestoreInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := FilestoreInstanceSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Clean KRM only labels and propagate to desired GCP state
	desiredPb.Labels = label.GCPLabels(a.desired)

	parent := a.id.ParentString()

	req := &pb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: a.id.Instance,
		Instance:   desiredPb,
	}

	log.V(0).Info("making filestore CreateInstance call", "request", req)

	op, err := a.client.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance: %w", err)
	}

	latest, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for filestoreInstance create %s: %w", a.id.String(), err)
	}

	log.V(0).Info("created filestoreInstance", "filestoreInstance", latest)

	return a.updateStatus(ctx, createOp, latest)
}

// Update implements the Adapter interface.
func (a *filestoreInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := FilestoreInstanceSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Clean KRM only labels and propagate to desired GCP state
	desiredPb.Labels = label.GCPLabels(a.desired)

	paths, updateMask, err := compareFilestoreInstance(ctx, a.actual, a.desired, desiredPb)
	if err != nil {
		return fmt.Errorf("comparing actual and desired FilestoreInstance: %w", err)
	}

	if len(paths) == 0 {
		log.V(2).Info("no diff detected, skipping update")
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	report := &structuredreporting.Diff{Object: u}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateInstanceRequest{
		Instance:   desiredPb,
		UpdateMask: updateMask,
	}
	req.Instance.Name = a.id.String()

	log.V(0).Info("making filestore UpdateInstance call", "request", req)

	op, err := a.client.UpdateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating instance: %w", err)
	}

	latest, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for filestoreInstance update %s: %w", a.id.String(), err)
	}

	log.V(0).Info("updated filestoreInstance", "filestoreInstance", latest)

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *filestoreInstanceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := FilestoreInstanceStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareFilestoreInstance(ctx context.Context, actual *pb.Instance, desiredKRM *krm.FilestoreInstance, desiredPb *pb.Instance) (sets.Set[string], *fieldmaskpb.FieldMask, error) {
	// 1. Identify unspecified spec fields in the desired state (in KRM), and assign the corresponding actual values to those fields.
	clonedDesired := proto.Clone(desiredPb).(*pb.Instance)

	if desiredKRM.Spec.Description == nil {
		clonedDesired.Description = actual.Description
	}
	if desiredKRM.Spec.Tier == nil {
		clonedDesired.Tier = actual.Tier
	}
	// For file shares:
	if len(desiredKRM.Spec.FileShares) == len(actual.FileShares) {
		for i := range desiredKRM.Spec.FileShares {
			krmShare := desiredKRM.Spec.FileShares[i]
			actShare := actual.FileShares[i]
			desShare := clonedDesired.FileShares[i]
			if krmShare.Name == nil {
				desShare.Name = actShare.Name
			}
			if krmShare.CapacityGb == nil {
				desShare.CapacityGb = actShare.CapacityGb
			}
			if len(krmShare.NfsExportOptions) == 0 && len(actShare.NfsExportOptions) > 0 {
				desShare.NfsExportOptions = actShare.NfsExportOptions
			}
			if krmShare.SourceBackupRef == nil && actShare.GetSourceBackup() != "" {
				desShare.Source = &pb.FileShareConfig_SourceBackup{SourceBackup: actShare.GetSourceBackup()}
			}
		}
	}
	// For networks:
	if len(desiredKRM.Spec.Networks) == len(actual.Networks) {
		for i := range desiredKRM.Spec.Networks {
			krmNet := desiredKRM.Spec.Networks[i]
			actNet := actual.Networks[i]
			desNet := clonedDesired.Networks[i]
			if krmNet.NetworkRef == nil {
				desNet.Network = actNet.Network
			}
			if len(krmNet.Modes) == 0 && len(actNet.Modes) > 0 {
				desNet.Modes = actNet.Modes
			}
			if krmNet.ReservedIPRange == nil {
				desNet.ReservedIpRange = actNet.ReservedIpRange
			}
			if len(krmNet.IpAddresses) == 0 && len(actNet.IpAddresses) > 0 {
				desNet.IpAddresses = actNet.IpAddresses
			}
		}
	}

	// 2. Use common.CompareProtoMessage to do the comparison
	paths, err := common.CompareProtoMessage(clonedDesired, actual, common.BasicDiff)
	if err != nil {
		return nil, nil, err
	}

	// 3. Build the updateMask from the detected diff paths, filtering to allowed fields (description, file_shares, labels)
	var maskPaths []string
	for path := range paths {
		if path == "description" || strings.HasPrefix(path, "file_shares") || path == "labels" {
			maskPaths = append(maskPaths, path)
		}
	}
	updateMask := &fieldmaskpb.FieldMask{Paths: maskPaths}

	return paths, updateMask, nil
}
