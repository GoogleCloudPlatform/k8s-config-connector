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
// proto.service: google.cloud.essentialcontacts.v1.EssentialContactsService
// proto.message: google.cloud.essentialcontacts.v1.Contact

package mockessentialcontacts

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/essentialcontacts/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *EssentialContactsV1) GetContact(ctx context.Context, req *pb.GetContactRequest) (*pb.Contact, error) {
	name, err := s.parseContactName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Contact{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EssentialContactsV1) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.Contact, error) {
	//name, err := common.NewResourceID(req.Parent, "", "")
	contactId := 7 // to match mock logs. server generates this id.
	reqName := fmt.Sprintf("%s/contacts/%d", req.GetParent(), contactId)
	name, err := s.parseContactName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetContact()).(*pb.Contact)
	obj.Name = name.GetName()
	obj.ValidateTime = timestamppb.Now()
	obj.ValidationState = pb.ValidationState_VALID

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EssentialContactsV1) UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.Contact, error) {
	name, err := s.parseContactName(req.GetContact().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Contact{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "notification_category_subscriptions", "notificationCategorySubscriptions":
			obj.NotificationCategorySubscriptions = req.GetContact().GetNotificationCategorySubscriptions()
		case "language_tag", "languageTag":
			obj.LanguageTag = req.GetContact().GetLanguageTag()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *EssentialContactsV1) DeleteContact(ctx context.Context, req *pb.DeleteContactRequest) (*emptypb.Empty, error) {
	name, err := s.parseContactName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Contact{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// ListContacts implements the List method for Essential Contacts
func (s *EssentialContactsV1) ListContacts(ctx context.Context, req *pb.ListContactsRequest) (*pb.ListContactsResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be specified")
	}

	var contacts []*pb.Contact
	if err := s.storage.List(ctx, (&pb.Contact{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parent + "/contacts/",
	}, func(obj proto.Message) error {
		contact := obj.(*pb.Contact)
		contacts = append(contacts, contact)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListContactsResponse{
		Contacts: contacts,
	}, nil
}

type contactName struct {
	Project *projects.ProjectData
	Contact string
}

// GetName returns the name using the project number instead of ID
func (n *contactName) GetName() string {
	return fmt.Sprintf("projects/%d/contacts/%s", n.Project.Number, n.Contact)
}

func (n *contactName) String() string {
	return fmt.Sprintf("projects/%s/contacts/%s", n.Project.ID, n.Contact)
}

// parseContactName parses a string into an contactName.
// The expected form is `projects/*/contacts/*`.
func (s *MockService) parseContactName(name string) (*contactName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "contacts" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &contactName{
			Project: project,
			Contact: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
