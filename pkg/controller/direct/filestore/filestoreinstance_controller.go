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
	"maps"
	"sort"
	"strings"

	"google.golang.org/api/option"
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
	registry.RegisterModel(krm.FilestoreInstanceGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

type model struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type filestoreInstanceAdapter struct {
	id *krm.FilestoreInstanceIdentity

	desiredKRM *krm.FilestoreInstance
	actual     *pb.Instance

	client *gcp.CloudFilestoreManagerClient
	kube   client.Reader
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &filestoreInstanceAdapter{}

func (m *model) client(ctx context.Context) (*gcp.CloudFilestoreManagerClient, error) {
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
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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
		id:         id.(*krm.FilestoreInstanceIdentity),
		desiredKRM: obj,
		client:     gcpClient,
		kube:       kube,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *filestoreInstanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Instance == "" {
		return false, nil
	}

	req := &pb.GetInstanceRequest{
		Name: a.id.String(),
	}
	filestoreInstance, err := a.client.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = filestoreInstance

	return true, nil
}

// Delete implements the Adapter interface.
func (a *filestoreInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FilestoreInstance", "name", a.id.String())

	req := &pb.DeleteInstanceRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting FilestoreInstance %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for FilestoreInstance delete %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *filestoreInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *filestoreInstanceAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.Instance, error) {
	if err := common.NormalizeReferences(ctx, a.kube, a.desiredKRM, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := FilestoreInstanceSpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Apply default values from DCL schema
	for _, fs := range desired.FileShares {
		for _, opt := range fs.NfsExportOptions {
			if opt.AccessMode == pb.NfsExportOptions_ACCESS_MODE_UNSPECIFIED {
				opt.AccessMode = pb.NfsExportOptions_READ_WRITE
			}
			if opt.SquashMode == pb.NfsExportOptions_SQUASH_MODE_UNSPECIFIED {
				opt.SquashMode = pb.NfsExportOptions_NO_ROOT_SQUASH
			}
		}
	}

	for _, net := range desired.Networks {
		if len(net.Modes) == 0 {
			net.Modes = []pb.NetworkConfig_AddressMode{pb.NetworkConfig_MODE_IPV4}
		}
	}

	labels := label.NewGCPLabelsFromK8sLabels(u.GetLabels())
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["managed-by-cnrm"] = "true"
	desired.Labels = labels

	return desired, nil
}

// Create implements the Adapter interface.
func (a *filestoreInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	req := &pb.CreateInstanceRequest{
		Parent:     a.id.ParentString(),
		InstanceId: a.id.Instance,
		Instance:   desired,
	}

	log.V(2).Info("making Filestore CreateInstance call", "request", req)

	op, err := a.client.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance: %w", err)
	}

	latest, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for FilestoreInstance create %s: %w", a.id.String(), err)
	}

	log.V(2).Info("created FilestoreInstance", "FilestoreInstance", latest)

	return a.updateStatus(ctx, createOp, latest)
}

// Update implements the Adapter interface.
func (a *filestoreInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	paths, updateMask, err := compareFilestoreInstance(ctx, a.actual, desired)
	if err != nil {
		return fmt.Errorf("comparing actual and desired FilestoreInstance: %w", err)
	}

	var latest *pb.Instance
	if len(paths) > 0 {
		report := &structuredreporting.Diff{Object: u}
		for _, path := range paths.UnsortedList() {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)

		req := &pb.UpdateInstanceRequest{
			Instance:   desired,
			UpdateMask: updateMask,
		}
		req.Instance.Name = a.id.String()

		log.V(2).Info("making Filestore UpdateInstance call", "request", req)

		op, err := a.client.UpdateInstance(ctx, req)
		if err != nil {
			return err
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for FilestoreInstance update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("updated FilestoreInstance", "FilestoreInstance", latest)
	} else {
		latest = a.actual
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *filestoreInstanceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := FilestoreInstanceStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.PtrTo(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

func compareFilestoreInstance(ctx context.Context, actual, desired *pb.Instance) (sets.Set[string], *fieldmaskpb.FieldMask, error) {
	var maskedActual *pb.Instance
	{
		mapCtx := &direct.MapContext{}
		spec := FilestoreInstanceSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = FilestoreInstanceSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}

	fillUnsetFields(desired, maskedActual)

	maskedActual.Name = desired.Name

	paths, err := common.CompareProtoMessage(desired, maskedActual, common.BasicDiff)
	if err != nil {
		return nil, nil, err
	}

	if len(paths) == 0 {
		return nil, nil, nil
	}

	pathsList := paths.UnsortedList()
	sort.Strings(pathsList)
	updateMask := &fieldmaskpb.FieldMask{
		Paths: pathsList,
	}
	return paths, updateMask, nil
}

func fillUnsetFields(desired, maskedActual *pb.Instance) {
	if desired.Description == "" {
		desired.Description = maskedActual.Description
	}
	if desired.Labels == nil && maskedActual.Labels != nil {
		desired.Labels = maps.Clone(maskedActual.Labels)
	}
	for _, df := range desired.FileShares {
		for _, af := range maskedActual.FileShares {
			if df.Name == af.Name {
				if df.CapacityGb == 0 {
					df.CapacityGb = af.CapacityGb
				}
				if df.Source == nil {
					df.Source = af.Source
				}
				if len(df.NfsExportOptions) == 0 {
					df.NfsExportOptions = af.NfsExportOptions
				}
			}
		}
	}
	for _, dn := range desired.Networks {
		for _, an := range maskedActual.Networks {
			if isSameNetwork(dn.Network, an.Network) {
				an.Network = dn.Network

				if len(dn.Modes) == 0 {
					dn.Modes = an.Modes
				}
				if dn.ReservedIpRange == "" {
					dn.ReservedIpRange = an.ReservedIpRange
				}
				if dn.ConnectMode == pb.NetworkConfig_CONNECT_MODE_UNSPECIFIED {
					dn.ConnectMode = an.ConnectMode
				}
			}
		}
	}
}

func isSameNetwork(a, b string) bool {
	if a == b {
		return true
	}
	if a == "" || b == "" {
		return false
	}
	if strings.HasSuffix(a, "/"+b) || strings.HasSuffix(b, "/"+a) {
		return true
	}
	return false
}
