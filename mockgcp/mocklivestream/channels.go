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

package mocklivestream

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type channelName struct {
	Project  *projects.ProjectData
	Location string
	Channel  string
}

func (n *channelName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/channels/" + n.Channel
}

func (s *MockService) parseChannelName(name string) (*channelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "channels" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &channelName{
			Project:  project,
			Location: tokens[3],
			Channel:  tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) normalizeInputAttachmentLink(input string) string {
	tokens := strings.Split(input, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		if project, err := s.Projects.GetProjectByID(tokens[1]); err == nil {
			tokens[1] = strconv.FormatInt(project.Number, 10)
			return strings.Join(tokens, "/")
		}
	}
	return input
}

func (s *MockService) normalizeAssetLink(asset string) string {
	tokens := strings.Split(asset, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		if project, err := s.Projects.GetProjectByID(tokens[1]); err == nil {
			tokens[1] = strconv.FormatInt(project.Number, 10)
			return strings.Join(tokens, "/")
		}
	}
	return asset
}

func (s *MockService) populateChannelDefaults(obj *pb.Channel) {
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

	if obj.RetentionConfig == nil {
		obj.RetentionConfig = &pb.RetentionConfig{}
	}

	if obj.AutoTranscriptionConfig == nil {
		obj.AutoTranscriptionConfig = &pb.AutoTranscriptionConfig{
			DisplayTiming: pb.AutoTranscriptionConfig_ASYNC,
			QualityPreset: pb.AutoTranscriptionConfig_BALANCED_QUALITY,
		}
	}

	if len(obj.InputAttachments) > 0 && obj.ActiveInput == "" {
		obj.ActiveInput = obj.InputAttachments[0].Key
	}

	for _, attachment := range obj.InputAttachments {
		attachment.Input = s.normalizeInputAttachmentLink(attachment.Input)
	}

	for _, overlay := range obj.StaticOverlays {
		overlay.Asset = s.normalizeAssetLink(overlay.Asset)
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

func (s *LivestreamServer) GetChannel(ctx context.Context, req *pb.GetChannelRequest) (*pb.Channel, error) {
	name, err := s.parseChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Channel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *LivestreamServer) CreateChannel(ctx context.Context, req *pb.CreateChannelRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/channels/" + req.ChannelId
	name, err := s.parseChannelName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.Channel).(*pb.Channel)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.StreamingState = pb.Channel_STOPPED

	s.populateChannelDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "create",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *LivestreamServer) UpdateChannel(ctx context.Context, req *pb.UpdateChannelRequest) (*longrunning.Operation, error) {
	reqName := req.GetChannel().GetName()
	name, err := s.parseChannelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Channel{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	updated := proto.Clone(existing).(*pb.Channel)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetChannel()).(*pb.Channel)
		updated.Name = fqn
		updated.CreateTime = existing.CreateTime
		updated.StreamingState = existing.StreamingState
	} else {
		for _, path := range paths {
			switch path {
			case "elementary_streams", "elementaryStreams":
				updated.ElementaryStreams = req.GetChannel().GetElementaryStreams()
			case "input_config", "inputConfig":
				updated.InputConfig = req.GetChannel().GetInputConfig()
			case "log_config", "logConfig":
				updated.LogConfig = req.GetChannel().GetLogConfig()
			case "manifests":
				updated.Manifests = req.GetChannel().GetManifests()
			case "sprite_sheets", "spriteSheets":
				updated.SpriteSheets = req.GetChannel().GetSpriteSheets()
			case "labels":
				updated.Labels = req.GetChannel().GetLabels()
			case "input_attachments", "inputAttachments":
				updated.InputAttachments = req.GetChannel().GetInputAttachments()
			case "output":
				updated.Output = req.GetChannel().GetOutput()
			case "mux_streams", "muxStreams":
				updated.MuxStreams = req.GetChannel().GetMuxStreams()
			case "retention_config", "retentionConfig":
				updated.RetentionConfig = req.GetChannel().GetRetentionConfig()
			case "static_overlays", "staticOverlays":
				updated.StaticOverlays = req.GetChannel().GetStaticOverlays()
			case "timecode_config", "timecodeConfig":
				updated.TimecodeConfig = req.GetChannel().GetTimecodeConfig()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
			}
		}
	}

	s.populateChannelDefaults(updated)
	updated.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "update",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (s *LivestreamServer) DeleteChannel(ctx context.Context, req *pb.DeleteChannelRequest) (*longrunning.Operation, error) {
	name, err := s.parseChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Channel{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	opMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Verb:       "delete",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *LivestreamServer) ListChannels(ctx context.Context, req *pb.ListChannelsRequest) (*pb.ListChannelsResponse, error) {
	response := &pb.ListChannelsResponse{}
	findPrefix := req.Parent + "/channels/"

	channelKind := (&pb.Channel{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, channelKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		response.Channels = append(response.Channels, obj.(*pb.Channel))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
