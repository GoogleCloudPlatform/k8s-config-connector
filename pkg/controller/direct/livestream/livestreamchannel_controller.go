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

package livestream

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/video/livestream/apiv1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/livestream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LiveStreamChannelGVK, NewChannelModel)
}

func NewChannelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &channelModel{config: *config}, nil
}

var _ directbase.Model = &channelModel{}

type channelModel struct {
	config config.ControllerConfig
}

func (m *channelModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building livestream client: %w", err)
	}
	return gcpClient, nil
}

func (m *channelModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LiveStreamChannel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.LiveStreamChannelIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := LiveStreamChannelSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &channelAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       desired,
		projectMapper: m.config.ProjectMapper,
		model:         m,
	}, nil
}

func (m *channelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.LiveStreamChannelIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &channelAdapter{
		id:            id,
		gcpClient:     gcpClient,
		projectMapper: m.config.ProjectMapper,
		model:         m,
	}, nil
}

type channelAdapter struct {
	id            *krm.LiveStreamChannelIdentity
	gcpClient     *gcp.Client
	desired       *pb.Channel
	actual        *pb.Channel
	projectMapper *projects.ProjectMapper
	model         *channelModel
}

var _ directbase.Adapter = &channelAdapter{}

func (a *channelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding LiveStreamChannel", "id", a.id)

	req := &pb.GetChannelRequest{
		Name: a.id.String(),
	}
	channel, err := a.gcpClient.GetChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LiveStreamChannel %s: %w", a.id.String(), err)
	}

	a.actual = channel
	return true, nil
}

func (a *channelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LiveStreamChannel", "id", a.id)

	req := &pb.CreateChannelRequest{
		Parent:    a.id.ParentString(),
		ChannelId: a.id.Channel,
		Channel:   a.desired,
	}
	op, err := a.gcpClient.CreateChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LiveStreamChannel %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for LiveStreamChannel %s creation: %w", a.id.String(), err)
	}

	// Fetch fully-populated resource after creation
	refetched, err := a.gcpClient.GetChannel(ctx, &pb.GetChannelRequest{Name: a.id.String()})
	if err != nil {
		refetched = created
	}
	a.actual = refetched

	return a.updateStatus(ctx, createOp, refetched)
}

func (a *channelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LiveStreamChannel", "id", a.id)

	diffs, updateMask, err := a.compareChannel(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*pb.Channel)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateChannelRequest{
			Channel:    desiredCopy,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateChannel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LiveStreamChannel %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting LiveStreamChannel %s update: %w", a.id.String(), err)
		}

		// Fetch fully-populated resource after update
		refetched, err := a.gcpClient.GetChannel(ctx, &pb.GetChannelRequest{Name: a.id.String()})
		if err != nil {
			refetched = latest
		}
		latest = refetched
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *channelAdapter) compareChannel(ctx context.Context, actual, desired *pb.Channel) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LiveStreamChannelSpec_FromProto, LiveStreamChannelSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Channel)

	// Normalize project number to project ID in input attachments and static overlays
	if a.projectMapper != nil {
		for _, attachment := range maskedActual.InputAttachments {
			if attachment.Input != "" {
				if normalized, err := a.projectMapper.ReplaceProjectNumberWithIDInLink(ctx, attachment.Input); err == nil {
					attachment.Input = normalized
				}
			}
		}
		for _, attachment := range clonedDesired.InputAttachments {
			if attachment.Input != "" {
				if normalized, err := a.projectMapper.ReplaceProjectNumberWithIDInLink(ctx, attachment.Input); err == nil {
					attachment.Input = normalized
				}
			}
		}
		for _, overlay := range maskedActual.StaticOverlays {
			if overlay.Asset != "" {
				if normalized, err := a.projectMapper.ReplaceProjectNumberWithIDInLink(ctx, overlay.Asset); err == nil {
					overlay.Asset = normalized
				}
			}
		}
		for _, overlay := range clonedDesired.StaticOverlays {
			if overlay.Asset != "" {
				if normalized, err := a.projectMapper.ReplaceProjectNumberWithIDInLink(ctx, overlay.Asset); err == nil {
					overlay.Asset = normalized
				}
			}
		}
	}

	populateDefaults := func(obj *pb.Channel) {
		if obj.InputConfig == nil {
			obj.InputConfig = &pb.InputConfig{
				InputSwitchMode: pb.InputConfig_FAILOVER_PREFER_PRIMARY,
			}
		} else if obj.InputConfig.InputSwitchMode == pb.InputConfig_INPUT_SWITCH_MODE_UNSPECIFIED {
			obj.InputConfig.InputSwitchMode = pb.InputConfig_FAILOVER_PREFER_PRIMARY
		}
		if obj.LogConfig == nil {
			obj.LogConfig = &pb.LogConfig{
				LogSeverity: pb.LogConfig_OFF,
			}
		} else if obj.LogConfig.LogSeverity == pb.LogConfig_LOG_SEVERITY_UNSPECIFIED {
			obj.LogConfig.LogSeverity = pb.LogConfig_OFF
		}
		if obj.TimecodeConfig == nil {
			obj.TimecodeConfig = &pb.TimecodeConfig{
				Source: pb.TimecodeConfig_MEDIA_TIMESTAMP,
			}
		} else if obj.TimecodeConfig.Source == pb.TimecodeConfig_TIMECODE_SOURCE_UNSPECIFIED {
			obj.TimecodeConfig.Source = pb.TimecodeConfig_MEDIA_TIMESTAMP
		}
		if obj.RetentionConfig != nil && obj.RetentionConfig.RetentionWindowDuration == nil {
			obj.RetentionConfig = nil
		}
		for _, manifest := range obj.Manifests {
			if manifest.MaxSegmentCount == 0 {
				manifest.MaxSegmentCount = 5
			}
			if manifest.SegmentKeepDuration == nil {
				manifest.SegmentKeepDuration = &durationpb.Duration{Seconds: 60}
			}
		}
		for _, stream := range obj.ElementaryStreams {
			if vs := stream.GetVideoStream(); vs != nil && vs.GetH264() != nil {
				h264 := vs.GetH264()
				if h264.EntropyCoder == "" {
					h264.EntropyCoder = "cabac"
				}
				if h264.GopMode == nil {
					h264.GopMode = &pb.VideoStream_H264CodecSettings_GopDuration{
						GopDuration: &durationpb.Duration{Seconds: 2},
					}
				}
				if h264.Profile == "" {
					h264.Profile = "main"
				}
				if h264.VbvSizeBits == 0 && h264.BitrateBps > 0 {
					h264.VbvSizeBits = h264.BitrateBps
				}
				if h264.VbvFullnessBits == 0 && h264.BitrateBps > 0 {
					h264.VbvFullnessBits = int32(float64(h264.BitrateBps) * 0.9)
				}
			}
			if as := stream.GetAudioStream(); as != nil {
				if as.SampleRateHertz == 0 {
					as.SampleRateHertz = 48000
				}
			}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *channelAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Channel) error {
	mapCtx := &direct.MapContext{}
	status := krm.LiveStreamChannelStatus{}
	status.ObservedState = LiveStreamChannelObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *channelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LiveStreamChannel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LiveStreamChannelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Channel)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Channel)
	u.SetGroupVersionKind(krm.LiveStreamChannelGVK)

	return u, nil
}

func (a *channelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LiveStreamChannel", "id", a.id)

	req := &pb.DeleteChannelRequest{
		Name:  a.id.String(),
		Force: true,
	}
	op, err := a.gcpClient.DeleteChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LiveStreamChannel, assuming it was already deleted", "id", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LiveStreamChannel %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete LiveStreamChannel %s: %w", a.id.String(), err)
	}
	return true, nil
}
