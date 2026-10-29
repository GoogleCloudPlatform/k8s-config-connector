// Copyright 2025 Google LLC
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

// +generated:controller
// krm.group: cloudchannel.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudchannelChannelPartnerLink

package cloudchannel

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudchannel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/channel/apiv1"
	channelpb "cloud.google.com/go/channel/apiv1/channelpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudchannelChannelPartnerLinkGVK, NewChannelPartnerLinkModel)
}

func NewChannelPartnerLinkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelChannelPartnerLink{config: *config}, nil
}

var _ directbase.Model = &modelChannelPartnerLink{}

type modelChannelPartnerLink struct {
	config config.ControllerConfig
}

func (m *modelChannelPartnerLink) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudchannelChannelPartnerLink{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudchannelChannelPartnerLinkIdentity)

	// Get cloudchannel GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ChannelPartnerLinkAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelChannelPartnerLink) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudchannel.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudchannelChannelPartnerLinkIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ChannelPartnerLinkAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelChannelPartnerLink) client(ctx context.Context) (*gcp.CloudChannelClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudChannelRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ChannelPartnerLink client: %w", err)
	}
	return gcpClient, err
}

type ChannelPartnerLinkAdapter struct {
	id        *krm.CloudchannelChannelPartnerLinkIdentity
	gcpClient *gcp.CloudChannelClient
	desired   *krm.CloudchannelChannelPartnerLink
	actual    *channelpb.ChannelPartnerLink
}

var _ directbase.Adapter = &ChannelPartnerLinkAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ChannelPartnerLinkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ChannelPartnerLink", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &channelpb.GetChannelPartnerLinkRequest{
		Name: a.id.String(),
	}
	channelpartnerlinkpb, err := a.gcpClient.GetChannelPartnerLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ChannelPartnerLink %q: %w", a.id, err)
	}

	a.actual = channelpartnerlinkpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ChannelPartnerLinkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ChannelPartnerLink", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudchannelChannelPartnerLinkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &channelpb.CreateChannelPartnerLinkRequest{
		Parent:             a.id.ParentString(),
		ChannelPartnerLink: resource,
	}
	created, err := a.gcpClient.CreateChannelPartnerLink(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ChannelPartnerLink %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ChannelPartnerLink", "name", a.id)

	status := &krm.CloudchannelChannelPartnerLinkStatus{}
	status.ObservedState = CloudchannelChannelPartnerLinkObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.CloudchannelChannelPartnerLinkIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.ChannelPartnerLink = parsedID.ChannelPartnerLink
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ChannelPartnerLinkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ChannelPartnerLink", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudchannelChannelPartnerLinkSpec_FromProto,
		CloudchannelChannelPartnerLinkSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &channelpb.UpdateChannelPartnerLinkRequest{
			Name:               a.id.String(),
			UpdateMask:         updateMask,
			ChannelPartnerLink: resource,
		}
		updated, err = a.gcpClient.UpdateChannelPartnerLink(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ChannelPartnerLink %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ChannelPartnerLink", "name", a.id)
	}

	status := &krm.CloudchannelChannelPartnerLinkStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudchannelChannelPartnerLinkObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ChannelPartnerLinkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudchannelChannelPartnerLink{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudchannelChannelPartnerLinkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudchannelChannelPartnerLinkGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ChannelPartnerLinkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for ChannelPartnerLink, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
