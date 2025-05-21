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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
)

type buckets struct {
	*MockService
	pb.UnimplementedBucketsServerServer
}

func (s *buckets) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetName()

	obj := &pb.Bucket{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	ret := proto.Clone(obj).(*pb.Bucket)

	projection := req.GetProjection()
	if projection == "" {
		projection = "noAcl"
	}
	switch projection {
	case "full":
	// full: Include all properties.
	// TODO: Verify storage.buckets.getIamPolicy permission (one day!)
	case "noAcl":
		// noAcl: Omit owner, acl, and defaultObjectAcl properties.
		ret.Acl = nil
		ret.DefaultObjectAcl = nil
		ret.Owner = nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid projection: %s", projection)
	}

	// httpmux.SetExpiresHeader(ctx, "Mon, 01 Jan 1990 00:00:00 GMT")

	return ret, nil
}

func (s *buckets) ListBuckets(ctx context.Context, req *pb.ListBucketsRequest) (*pb.Buckets, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	var buckets []*pb.Bucket

	bucketKind := (&pb.Bucket{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, bucketKind, storage.ListOptions{}, func(obj proto.Message) error {
		bucket := obj.(*pb.Bucket)

		// TODO: Some form of ACL?

		if bucket.GetProjectNumber() != uint64(project.Number) {
			return nil
		}

		buckets = append(buckets, bucket)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.Buckets{
		Items:         buckets,
		NextPageToken: nil,
		Kind:          PtrTo("storage#buckets"),
	}, nil
}

func (s *buckets) InsertBucket(ctx context.Context, req *pb.InsertBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetBucket().GetName()

	if req.GetProject() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "project is required")
	}
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}
	now := timestamppb.Now()

	obj := proto.Clone(req.GetBucket()).(*pb.Bucket)
	obj.Id = PtrTo(fqn)
	obj.Kind = PtrTo("storage#bucket")
	obj.Name = PtrTo(fqn)
	obj.ProjectNumber = PtrTo(uint64(project.Number))
	obj.Location = PtrTo("US")
	obj.LocationType = PtrTo("multi-region")
	obj.Rpo = PtrTo("DEFAULT")
	obj.SelfLink = PtrTo(fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s", obj.GetName()))
	obj.StorageClass = PtrTo("STANDARD")
	obj.TimeCreated = now
	obj.Updated = now
	obj.Metageneration = PtrTo(int64(1))

	obj.Etag = PtrTo(computeEtag(obj))
	if obj.Lifecycle != nil && proto.Equal(obj.Lifecycle, &pb.BucketLifecycle{}) {
		obj.Lifecycle = nil
	}

	iamConfiguration := obj.IamConfiguration
	if iamConfiguration == nil {
		iamConfiguration = &pb.BucketIamConfiguration{}
		obj.IamConfiguration = iamConfiguration
	}
	if iamConfiguration.PublicAccessPrevention == nil {
		iamConfiguration.PublicAccessPrevention = PtrTo("inherited")
	}
	bucketPolicyOnly := iamConfiguration.BucketPolicyOnly
	if bucketPolicyOnly == nil {
		bucketPolicyOnly = &pb.BucketPolicyOnly{}
		iamConfiguration.BucketPolicyOnly = bucketPolicyOnly
	}
	if bucketPolicyOnly.Enabled == nil {
		bucketPolicyOnly.Enabled = PtrTo(false)
	}
	ubla := iamConfiguration.UniformBucketLevelAccess
	if ubla == nil {
		ubla = &pb.UniformBucketLevelAccess{}
		iamConfiguration.UniformBucketLevelAccess = ubla
	}

	softDeletePolicy := obj.SoftDeletePolicy
	if softDeletePolicy == nil {
		softDeletePolicy = &pb.BucketSoftDeletePolicy{}
		obj.SoftDeletePolicy = softDeletePolicy
	}
	if softDeletePolicy.RetentionDurationSeconds == nil {
		defaultRetention := time.Hour * 7 * 24
		softDeletePolicy.RetentionDurationSeconds = PtrTo(int64(defaultRetention.Seconds()))
	}
	// TODO: Should be now
	softDeletePolicy.EffectiveTime = timestamppb.New(time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *buckets) populateDefaults(ctx context.Context, project *projects.ProjectData, obj *pb.Bucket) error {
	owner := obj.Owner
	if owner == nil {
		owner = &pb.BucketOwner{}
		obj.Owner = owner
	}
	if owner.Entity == nil {
		owner.Entity = PtrTo(fmt.Sprintf("project-owners-%d", project.Number))
	}

	if len(obj.Acl) == 0 {
		obj.Acl = append(obj.Acl, buildBucketACL(ctx, project, obj, "OWNER", "owners"))
		obj.Acl = append(obj.Acl, buildBucketACL(ctx, project, obj, "OWNER", "editors"))
		obj.Acl = append(obj.Acl, buildBucketACL(ctx, project, obj, "READER", "viewers"))
	}
	if len(obj.DefaultObjectAcl) == 0 {
		obj.DefaultObjectAcl = append(obj.DefaultObjectAcl, buildObjectACL(ctx, project, obj, "OWNER", "owners"))
		obj.DefaultObjectAcl = append(obj.DefaultObjectAcl, buildObjectACL(ctx, project, obj, "OWNER", "editors"))
		obj.DefaultObjectAcl = append(obj.DefaultObjectAcl, buildObjectACL(ctx, project, obj, "READER", "viewers"))
	}
	return nil
}

func buildBucketACL(ctx context.Context, project *projects.ProjectData, obj *pb.Bucket, role string, team string) *pb.BucketAccessControl {
	acl := &pb.BucketAccessControl{}
	acl.Bucket = PtrTo(obj.GetName())
	acl.Entity = PtrTo(fmt.Sprintf("project-%s-%d", team, project.Number))
	acl.Role = PtrTo(role)
	acl.Id = PtrTo(acl.GetBucket() + "/" + acl.GetEntity())
	acl.Kind = PtrTo("storage#bucketAccessControl")
	acl.SelfLink = PtrTo(fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s/acl/%s", obj.GetName(), acl.GetEntity()))
	acl.ProjectTeam = &pb.BucketAccessControlProjectTeam{
		ProjectNumber: PtrTo(fmt.Sprintf("%d", project.Number)),
		Team:          PtrTo(team),
	}
	acl.Etag = PtrTo(computeEtag(obj))
	return acl
}

func buildObjectACL(ctx context.Context, project *projects.ProjectData, obj *pb.Bucket, role string, team string) *pb.ObjectAccessControl {
	acl := &pb.ObjectAccessControl{}
	acl.Entity = PtrTo(fmt.Sprintf("project-%s-%d", team, project.Number))
	acl.Role = PtrTo(role)
	acl.Kind = PtrTo("storage#objectAccessControl")
	acl.ProjectTeam = &pb.ObjectAccessControlProjectTeam{
		ProjectNumber: PtrTo(fmt.Sprintf("%d", project.Number)),
		Team:          PtrTo(team),
	}
	acl.Etag = PtrTo(computeEtag(obj))
	return acl
}

func (s *buckets) PatchBucket(ctx context.Context, req *pb.PatchBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetName()

	obj := &pb.Bucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProjectByNumber(fmt.Sprintf("%d", obj.GetProjectNumber()))
	if err != nil {
		return nil, err
	}

	if patch := req.Bucket; patch != nil {
		if patch.Labels != nil {
			obj.Labels = patch.Labels
		}
		if patch.Lifecycle != nil {
			obj.Lifecycle = patch.Lifecycle
		}
		if patch.Versioning != nil {
			obj.Versioning = patch.Versioning
		}

		if patch.SoftDeletePolicy != nil {
			if patch.SoftDeletePolicy.RetentionDurationSeconds != nil {
				if obj.SoftDeletePolicy == nil {
					obj.SoftDeletePolicy = &pb.BucketSoftDeletePolicy{}
				}
				obj.SoftDeletePolicy.RetentionDurationSeconds = patch.SoftDeletePolicy.RetentionDurationSeconds

				// If the value is zero, we clear the effectiveTime (apparently)
				if obj.GetSoftDeletePolicy().GetRetentionDurationSeconds() == 0 {
					obj.SoftDeletePolicy.EffectiveTime = nil
				}
			}
		}
	}

	// Remove empty lifecycle (no rules)
	if obj.Lifecycle != nil && proto.Equal(obj.Lifecycle, &pb.BucketLifecycle{}) {
		obj.Lifecycle = nil
	}

	obj.Metageneration = PtrTo(int64(obj.GetMetageneration() + 1))

	if err := s.populateDefaults(ctx, project, obj); err != nil {
		return nil, err
	}

	obj.Etag = PtrTo(computeEtag(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *buckets) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*empty.Empty, error) {
	fqn := req.GetName()

	deletedObj := &pb.Bucket{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &empty.Empty{}, nil
}
