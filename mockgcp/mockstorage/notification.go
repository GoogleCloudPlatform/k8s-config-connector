// Copyright 2023 Google LLC
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

package mockstorage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

type notifications struct {
	*MockService
	pb.UnimplementedNotificationsServerServer
	id int
}

func (n *notifications) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.Notification, error) {
	fqn := fullyQualifiedName(req.GetBucket(), req.GetName())
	obj := &pb.Notification{}
	if err := n.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// The real GCP service stores what the request gives but returns the topic with the PubSub domain.
	obj.Topic = PtrTo("//pubsub.googleapis.com/" + ValueOf(obj.Topic))
	ret := proto.Clone(obj).(*pb.Notification)
	return ret, nil

}

func (n *notifications) InsertNotification(ctx context.Context, req *pb.InsertNotificationRequest) (*pb.Notification, error) {
	n.id += 1
	fqn := fullyQualifiedName(req.GetBucket(), fmt.Sprint(n.id))

	obj := proto.Clone(req.GetNotification()).(*pb.Notification)
	obj.Id = PtrTo(fmt.Sprint(n.id))
	obj.Etag = PtrTo(fields.ComputeWeakEtag(obj))
	obj.SelfLink = PtrTo("https://www.googleapis.com/storage/v1/b/" + req.GetBucket() + "/notificationConfigs/" + fmt.Sprint(n.id))
	obj.Kind = PtrTo("storage#notification")
	if err := n.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	// The real GCP service stores what the request gives but returns the topic with the PubSub domain.
	obj.Topic = PtrTo("//pubsub.googleapis.com/" + ValueOf(obj.Topic))
	return obj, nil
}

func (n *notifications) DeleteNotification(ctx context.Context, req *pb.DeleteNotificationRequest) (*empty.Empty, error) {
	fqn := fullyQualifiedName(req.GetBucket(), ValueOf(req.Name))
	deletedObj := &pb.Notification{}
	if err := n.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	httpmux.SetStatusCode(ctx, http.StatusNoContent)
	return &empty.Empty{}, nil
}

func fullyQualifiedName(bucket, notification string) string {
	return "storage/" + bucket + "/notificationConfigs/" + notification
}
