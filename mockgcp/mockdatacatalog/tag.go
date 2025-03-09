// Copyright 2024 Google LLC
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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.Tag

package mockdatacatalog

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DataCatalogV1) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.Tag, error) {
	// Parent name format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}
	parentName, err := s.parseEntryName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Generate a fixed tag ID for predictability in tests.
	// The real API generates an ID, but for mock tests, a fixed one is easier.
	tagID := "mocktagid" // Use a predictable ID

	fqn := parentName.String() + "/tags/" + tagID

	obj := proto.Clone(req.Tag).(*pb.Tag)
	obj.Name = fqn

	// Validate the TagTemplate reference
	if obj.Template == "" {
		return nil, status.Errorf(codes.InvalidArgument, "tag template is required")
	}
	templateName, err := s.parseTagTemplateName(obj.Template)
	if err != nil {
		// Wrap the error to provide more context
		return nil, status.Errorf(codes.InvalidArgument, "invalid tag template name %q: %v", obj.Template, err)
	}

	// Fetch the TagTemplate to populate display names
	template := &pb.TagTemplate{}
	if err := s.storage.Get(ctx, templateName.String(), template); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.FailedPrecondition, "referenced tag template %q not found", obj.Template)
		}
		return nil, fmt.Errorf("fetching referenced tag template %q failed: %w", obj.Template, err)
	}

	obj.TemplateDisplayName = template.DisplayName
	for fieldID, field := range obj.Fields {
		if templateField, ok := template.Fields[fieldID]; ok {
			field.DisplayName = templateField.DisplayName
		} else {
			// This case might indicate inconsistency or a lenient create behavior.
			// For now, leave the display name empty if the field doesn't exist in the template.
			// Consider logging a warning or returning an error based on desired strictness.
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) UpdateTag(ctx context.Context, req *pb.UpdateTagRequest) (*pb.Tag, error) {
	reqName := req.GetTag().GetName()
	name, err := s.parseTagName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Tag{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement FieldMask application properly.
	// For now, rudimentary update that overwrites fields present in the request.
	// This assumes the request tag object contains only the fields to be updated.
	// A real implementation needs to merge based on updateMask.
	updateData := req.GetTag()

	// Example simple merge (overwriting fields if present in request)
	// This doesn't respect updateMask properly!
	for fieldID, fieldValue := range updateData.Fields {
		if existingField, ok := obj.Fields[fieldID]; ok {
			// Update existing field values based on type
			if fieldValue.GetStringValue() != "" {
				existingField.Kind = &pb.TagField_StringValue{StringValue: fieldValue.GetStringValue()}
			} else if fieldValue.GetBoolValue() { // Assuming bool defaults to false
				existingField.Kind = &pb.TagField_BoolValue{BoolValue: fieldValue.GetBoolValue()}
			} else if fieldValue.GetDoubleValue() != 0.0 {
				existingField.Kind = &pb.TagField_DoubleValue{DoubleValue: fieldValue.GetDoubleValue()}
			} else if fieldValue.GetTimestampValue() != nil {
				existingField.Kind = &pb.TagField_TimestampValue{TimestampValue: fieldValue.GetTimestampValue()}
			} else if enumVal := fieldValue.GetEnumValue(); enumVal != nil { // Check if the kind is EnumValue and get the inner message
				// Assign the oneof wrapper correctly
				existingField.Kind = &pb.TagField_EnumValue_{EnumValue: enumVal}
			} // Add other types as needed
		} else {
			// Field doesn't exist in the original tag? This might indicate an issue
			// or the updateMask intends to add it. A proper implementation needs the mask.
			// For now, let's just skip adding new fields during update.
			// Log this? return status.Errorf(codes.InvalidArgument, "attempting to update non-existent field %q without proper updateMask handling", fieldID)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Fetch the TagTemplate to populate display names
	templateName, err := s.parseTagTemplateName(obj.Template)
	if err != nil {
		// This should ideally not happen if the tag was created correctly referencing a valid template
		return nil, status.Errorf(codes.Internal, "error parsing template name from existing tag %q: %v", obj.Template, err)
	}
	template := &pb.TagTemplate{}
	if err := s.storage.Get(ctx, templateName.String(), template); err != nil {
		// This should ideally not happen if the template existed during tag creation
		return nil, status.Errorf(codes.Internal, "error fetching referenced tag template %q for existing tag: %v", obj.Template, err)
	}

	obj.TemplateDisplayName = template.DisplayName
	for fieldID, field := range obj.Fields {
		if templateField, ok := template.Fields[fieldID]; ok {
			field.DisplayName = templateField.DisplayName
		} else {
			// Field exists in tag but not in template - inconsistency?
			// Leave display name empty for now.
		}
	}

	return obj, nil
}

func (s *DataCatalogV1) ListTags(ctx context.Context, req *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	// Parent name format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}
	parentName, err := s.parseEntryName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentFQNPrefix := parentName.String() + "/tags/"

	response := &pb.ListTagsResponse{}

	// TODO: Implement pagination if needed. For now, list all matching tags.
	if err := s.storage.List(ctx, (*pb.Tag)(nil).ProtoReflect().Descriptor(), storage.ListOptions{}, func(obj proto.Message) error {
		tag, ok := obj.(*pb.Tag)
		if !ok {
			// This should not happen with properly typed storage, but handle defensively.
			return status.Errorf(codes.Internal, "unexpected object type found in storage: %T", obj)
		}

		// Filter by parent prefix
		if !strings.HasPrefix(tag.Name, parentFQNPrefix) {
			return nil // Skip this tag, it doesn't belong to the requested parent
		}
		// Clear fields that shouldn't be in the list response based on observed behavior.
		tag.TemplateDisplayName = ""
		for key := range tag.Fields {
			// DisplayName is part of the template, not the tag instance itself in this context.
			// Clear it as it shouldn't be returned in the Tag response based on observed behavior.
			if tag.Fields[key] != nil {
				tag.Fields[key].DisplayName = ""
			}
		}
		response.Tags = append(response.Tags, tag)
		return nil
	}); err != nil {
		return nil, err // Errors during listing (e.g., storage connection issues)
	}

	return response, nil
}

func (s *DataCatalogV1) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*emptypb.Empty, error) {
	name, err := s.parseTagName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deleted := &pb.Tag{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type tagName struct {
	Project    *projects.ProjectData
	Location   string
	EntryGroup string
	Entry      string
	Tag        string
}

// projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}
// projects/{project}/locations/{location}/tagTemplates/{tag_template}
func (n *tagName) String() string {
	// Tag name format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}/tags/{tag}
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s/tags/%s",
		n.Project.ID,
		n.Location,
		n.EntryGroup,
		n.Entry,
		n.Tag)
}

// parseTagName parses a string into a tagName.
// Format: projects/{project}/locations/{location}/entryGroups/{entry_group}/entries/{entry}/tags/{tag}
func (s *DataCatalogV1) parseTagName(name string) (*tagName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" && tokens[6] == "entries" && tokens[8] == "tags" {
		// Valid format
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q does not match pattern projects/*/locations/*/entryGroups/*/entries/*/tags/*", name)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		// This is how the GCP API seems to behave in practice
		return nil, status.Errorf(codes.NotFound, "project %q not found", tokens[1])
	}

	namePart := &tagName{
		Project:    project,
		Location:   tokens[3],
		EntryGroup: tokens[5],
		Entry:      tokens[7],
		Tag:        tokens[9],
	}

	return namePart, nil
}
