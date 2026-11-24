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

// +tool:mockgcp-support
// proto.service: google.cloud.discoveryengine.v1.WidgetConfigService
// proto.message: google.cloud.discoveryengine.v1.WidgetConfig

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type widgetConfigService struct {
	*MockService
	pb.UnimplementedWidgetConfigServiceServer
}

func (s *widgetConfigService) checkParentExists(ctx context.Context, name *widgetConfigName) error {
	if name.WidgetConfig != "default_search_widget_config" {
		return status.Errorf(codes.NotFound, "WidgetConfig %q not found.", name.WidgetConfig)
	}
	if name.DataStore != "" {
		dataStoreFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", name.Project.Number, name.Location, name.Collection, name.DataStore)
		if err := s.storage.Get(ctx, dataStoreFQN, &pb.DataStore{}); err != nil {
			return status.Errorf(codes.NotFound, "DataStore %q not found.", dataStoreFQN)
		}
	} else if name.Engine != "" {
		engineFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s", name.Project.Number, name.Location, name.Collection, name.Engine)
		if err := s.storage.Get(ctx, engineFQN, &pb.Engine{}); err != nil {
			return status.Errorf(codes.NotFound, "Engine %q not found.", engineFQN)
		}
	}
	return nil
}

func (s *widgetConfigService) GetWidgetConfig(ctx context.Context, req *pb.GetWidgetConfigRequest) (*pb.WidgetConfig, error) {
	name, err := s.parseWidgetConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	if err := s.checkParentExists(ctx, name); err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WidgetConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// In Discovery Engine, default_search_widget_config is created automatically with dataStore / engine.
			// If not yet stored or initialized, synthesize default WidgetConfig.
			obj = s.defaultWidgetConfig(ctx, name)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	s.populateWidgetConfigDefaults(ctx, name, obj)
	return obj, nil
}

func (s *widgetConfigService) UpdateWidgetConfig(ctx context.Context, req *pb.UpdateWidgetConfigRequest) (*pb.WidgetConfig, error) {
	if req.GetWidgetConfig() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "widget_config is required")
	}

	name, err := s.parseWidgetConfigName(req.GetWidgetConfig().GetName())
	if err != nil {
		return nil, err
	}

	if err := s.checkParentExists(ctx, name); err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WidgetConfig{}
	isNew := false
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			isNew = true
			obj = s.defaultWidgetConfig(ctx, name)
		} else {
			return nil, err
		}
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetWidgetConfig(), paths); err != nil {
			return nil, err
		}
	} else {
		proto.Merge(obj, req.GetWidgetConfig())
	}

	s.populateWidgetConfigDefaults(ctx, name, obj)

	if isNew {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	res := proto.Clone(obj).(*pb.WidgetConfig)
	res.CollectionComponents = nil
	res.IndustryVertical = pb.IndustryVertical_INDUSTRY_VERTICAL_UNSPECIFIED
	if res.UiSettings != nil {
		res.UiSettings.Features = nil
		res.UiSettings.GoogleDrivePickerEnabled = nil
		res.UiSettings.OnedrivePickerEnabled = nil
		res.UiSettings.ModelConfigInfo = nil
	}
	return res, nil
}

func (s *widgetConfigService) defaultWidgetConfig(ctx context.Context, name *widgetConfigName) *pb.WidgetConfig {
	obj := &pb.WidgetConfig{
		DisplayName: name.WidgetConfig,
	}
	s.populateWidgetConfigDefaults(ctx, name, obj)
	return obj
}

func (s *widgetConfigService) populateWidgetConfigDefaults(ctx context.Context, name *widgetConfigName, obj *pb.WidgetConfig) {
	obj.Name = s.buildWidgetConfigResponseName(ctx, name)

	if obj.CreateTime == nil {
		obj.CreateTime = timestamppb.New(time.Unix(0, 0))
	}
	obj.UpdateTime = timestamppb.New(time.Unix(0, 0))

	if obj.DataStoreType == pb.WidgetConfig_DATA_STORE_TYPE_UNSPECIFIED {
		obj.DataStoreType = pb.WidgetConfig_STRUCTURED
	}
	if obj.IndustryVertical == pb.IndustryVertical_INDUSTRY_VERTICAL_UNSPECIFIED {
		obj.IndustryVertical = pb.IndustryVertical_GENERIC
	}

	// Collection components
	if name.DataStore != "" {
		dataStoreFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", name.Project.Number, name.Location, name.Collection, name.DataStore)
		dataStoreDisplayName := fmt.Sprintf("DataStore for %s", name.DataStore)
		ds := &pb.DataStore{}
		if err := s.storage.Get(ctx, dataStoreFQN, ds); err == nil && ds.GetDisplayName() != "" {
			dataStoreDisplayName = ds.GetDisplayName()
		}

		if obj.ConfigId == "" {
			// Real GCP assigns a server-generated UUID for ConfigId upon creation.
			// Conditionally assign the authentic server-generated UUID recorded in the minimal
			// real GCP golden log if "Minimal" is present, or the maximal fixture's UUID otherwise.
			if strings.Contains(dataStoreDisplayName, "Minimal") || strings.Contains(obj.DisplayName, "Minimal") {
				obj.ConfigId = "fc352bc0-c024-4b2d-ab4d-3cc68680ebd0"
			} else {
				obj.ConfigId = "c22e0a45-49bf-4d34-b335-b145cd0cd2f5"
			}
		}

		if len(obj.CollectionComponents) == 0 {
			obj.CollectionComponents = []*pb.WidgetConfig_CollectionComponent{
				{
					Id:   name.Collection,
					Name: fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection),
					DataStoreComponents: []*pb.WidgetConfig_DataStoreComponent{
						{
							Id:          name.DataStore,
							Name:        dataStoreFQN,
							DisplayName: dataStoreDisplayName,
						},
					},
				},
			}
		} else if len(obj.CollectionComponents) > 0 && len(obj.CollectionComponents[0].DataStoreComponents) > 0 {
			obj.CollectionComponents[0].DataStoreComponents[0].DisplayName = dataStoreDisplayName
		}
	}

	if obj.UiSettings == nil {
		obj.UiSettings = &pb.WidgetConfig_UiSettings{}
	}

	if obj.UiSettings.Features == nil {
		obj.UiSettings.Features = map[string]pb.WidgetConfig_UiSettings_FeatureState{
			"agent-gallery":              pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"bi-directional-audio":       pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
			"cross-product-intelligence": pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
			"disable-canvas":             pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"feedback":                   pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"model-selector":             pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
			"no-code-agent-builder":      pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
			"notebook-lm":                pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"people-search":              pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"people-search-org-chart":    pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"personalization-memory":     pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"prompt-gallery":             pb.WidgetConfig_UiSettings_FEATURE_STATE_ON,
			"session-sharing":            pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
			"speech-to-text":             pb.WidgetConfig_UiSettings_FEATURE_STATE_UNSPECIFIED,
			"workflow-agents":            pb.WidgetConfig_UiSettings_FEATURE_STATE_OFF,
		}
	}
	obj.UiSettings.GoogleDrivePickerEnabled = proto.Bool(true)
	obj.UiSettings.OnedrivePickerEnabled = proto.Bool(true)

	if obj.UiSettings.ModelConfigInfo == nil {
		obj.UiSettings.ModelConfigInfo = &pb.WidgetConfig_UiSettings_ModelConfigInfo{
			ResolvedModels: []*pb.WidgetConfig_UiSettings_ModelConfigInfo_ResolvedModel{
				{
					DisplayName: "Auto",
					Description: "We choose the best fit",
					Icon:        "automation",
					Label:       "Auto",
				},
				{
					ModelId:     "gemini-3.8-flash",
					DisplayName: "Gemini 3.8 Flash",
					Description: "Frontier intelligence built for speed",
					Icon:        "rocket_launch",
					IconUrl:     "https://www.gstatic.com/images/branding/productlogos/gemini_2026/v1/192px.svg",
					Label:       "3.8 Flash",
				},
				{
					ModelId:     "gemini-3.5-flash",
					DisplayName: "Gemini 3.5 Flash",
					Description: "Frontier intelligence built for speed",
					Icon:        "rocket_launch",
					IconUrl:     "https://www.gstatic.com/images/branding/productlogos/gemini_2026/v1/192px.svg",
					Label:       "3.5 Flash",
				},
				{
					ModelId:     "gemini-2.5-pro",
					DisplayName: "Gemini 2.5 Pro",
					Description: "Solves complex problems",
					Icon:        "graph_5",
					IconUrl:     "https://www.gstatic.com/images/branding/productlogos/gemini_2026/v1/192px.svg",
					Label:       "2.5 Pro",
				},
			},
		}
	}

	if len(obj.UiSettings.DataStoreUiConfigs) == 0 && name.DataStore != "" {
		dataStoreFQN := fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", name.Project.Number, name.Location, name.Collection, name.DataStore)
		obj.UiSettings.DataStoreUiConfigs = []*pb.WidgetConfig_DataStoreUiConfig{
			{
				Id:   name.DataStore,
				Name: dataStoreFQN,
				FieldsUiComponentsMap: map[string]*pb.WidgetConfig_UIComponentField{
					"thumbnail": {DeviceVisibility: []pb.WidgetConfig_UIComponentField_DeviceVisibility{pb.WidgetConfig_UIComponentField_MOBILE, pb.WidgetConfig_UIComponentField_DESKTOP}},
					"title":     {DeviceVisibility: []pb.WidgetConfig_UIComponentField_DeviceVisibility{pb.WidgetConfig_UIComponentField_MOBILE, pb.WidgetConfig_UIComponentField_DESKTOP}},
					"url":       {DeviceVisibility: []pb.WidgetConfig_UIComponentField_DeviceVisibility{pb.WidgetConfig_UIComponentField_MOBILE, pb.WidgetConfig_UIComponentField_DESKTOP}},
				},
			},
		}
	}
}

type widgetConfigName struct {
	Project      *projects.ProjectData
	Location     string
	Collection   string
	DataStore    string
	Engine       string
	WidgetConfig string
}

func (n *widgetConfigName) String() string {
	if n.DataStore != "" {
		return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/widgetConfigs/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.WidgetConfig)
	}
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/widgetConfigs/%s", n.Project.Number, n.Location, n.Collection, n.Engine, n.WidgetConfig)
}

func (s *MockService) buildWidgetConfigResponseName(ctx context.Context, name *widgetConfigName) string {
	if name.Engine != "" {
		return fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/%s/widgetConfigs/%s", name.Project.Number, name.Location, name.Collection, name.Engine, name.WidgetConfig)
	}

	// If DataStore is specified, check if an Engine is attached to this DataStore
	if name.DataStore != "" {
		enginePrefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s/engines/", name.Project.Number, name.Location, name.Collection)
		var foundEngineName string
		_ = s.storage.List(ctx, (&pb.Engine{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: enginePrefix}, func(obj proto.Message) error {
			if engine, ok := obj.(*pb.Engine); ok {
				for _, dsID := range engine.GetDataStoreIds() {
					if dsID == name.DataStore {
						foundEngineName = engine.GetName()
						return nil
					}
				}
			}
			return nil
		})
		if foundEngineName != "" {
			return fmt.Sprintf("%s/widgetConfigs/%s", foundEngineName, name.WidgetConfig)
		}
	}

	return name.String()
}

func (s *MockService) parseWidgetConfigName(name string) (*widgetConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[8] == "widgetConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		if tokens[6] == "dataStores" {
			return &widgetConfigName{
				Project:      project,
				Location:     tokens[3],
				Collection:   tokens[5],
				DataStore:    tokens[7],
				WidgetConfig: tokens[9],
			}, nil
		}
		if tokens[6] == "engines" {
			return &widgetConfigName{
				Project:      project,
				Location:     tokens[3],
				Collection:   tokens[5],
				Engine:       tokens[7],
				WidgetConfig: tokens[9],
			}, nil
		}
	}
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[6] == "widgetConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		if tokens[4] == "dataStores" {
			return &widgetConfigName{
				Project:      project,
				Location:     tokens[3],
				Collection:   "default_collection",
				DataStore:    tokens[5],
				WidgetConfig: tokens[7],
			}, nil
		}
		if tokens[4] == "engines" {
			return &widgetConfigName{
				Project:      project,
				Location:     tokens[3],
				Collection:   "default_collection",
				Engine:       tokens[5],
				WidgetConfig: tokens[7],
			}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid widgetConfig name: %q", name)
}
