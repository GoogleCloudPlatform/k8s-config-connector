// Copyright 2026 High Five Labs
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

// Log Router
// settings (getSettings, updateSettings) as the Cloud Logging API documents
// them. Settings exist for every organization, folder and project and are
// read as defaults until written; only an organization's or a folder's can
// be updated. The service account IDs follow the Logging service agent's
// documented form; the key itself is not checked.

package mocklogging

import (
	"context"
	"strings"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// settingsName returns the settings' parent (organizations/{id},
// folders/{id} or projects/{id}) from a name with or without "/settings".
func settingsName(name string) (string, error) {
	parent := strings.TrimSuffix(name, "/settings")
	t := strings.Split(parent, "/")
	if len(t) != 2 || t[1] == "" {
		return "", status.Errorf(codes.InvalidArgument, "invalid settings name %q", name)
	}
	switch t[0] {
	case "organizations", "folders", "projects":
		return parent, nil
	}
	return "", status.Errorf(codes.InvalidArgument, "invalid settings name %q", name)
}

func (s *configServiceV2) readSettings(ctx context.Context, parent string) (*pb.Settings, error) {
	obj := &pb.Settings{}
	if err := s.storage.Get(ctx, parent+"/settings", obj); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		t := strings.Split(parent, "/")
		kind := map[string]string{"organizations": "org", "folders": "folder", "projects": "p"}[t[0]]
		obj = &pb.Settings{
			Name:                parent + "/settings",
			KmsServiceAccountId: "service-" + kind + "-" + t[1] + "@gcp-sa-logging.iam.gserviceaccount.com",
		}
	}
	return obj, nil
}

func (s *configServiceV2) GetSettings(ctx context.Context, req *pb.GetSettingsRequest) (*pb.Settings, error) {
	parent, err := settingsName(req.GetName())
	if err != nil {
		return nil, err
	}
	return s.readSettings(ctx, parent)
}

func (s *configServiceV2) UpdateSettings(ctx context.Context, req *pb.UpdateSettingsRequest) (*pb.Settings, error) {
	parent, err := settingsName(req.GetName())
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(parent, "projects/") {
		return nil, status.Errorf(codes.InvalidArgument, "settings can only be updated for organizations and folders, not %q", parent)
	}
	obj, err := s.readSettings(ctx, parent)
	if err != nil {
		return nil, err
	}
	in := req.GetSettings()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"kms_key_name", "storage_location", "disable_default_sink"}
	}
	for _, p := range paths {
		switch p {
		case "kms_key_name", "kmsKeyName":
			obj.KmsKeyName = in.GetKmsKeyName()
		case "storage_location", "storageLocation":
			obj.StorageLocation = in.GetStorageLocation()
		case "disable_default_sink", "disableDefaultSink":
			obj.DisableDefaultSink = in.GetDisableDefaultSink()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q in the update mask is not supported", p)
		}
	}
	obj = proto.CloneOf(obj)
	if err := s.storage.Update(ctx, parent+"/settings", obj); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		if err := s.storage.Create(ctx, parent+"/settings", obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}
