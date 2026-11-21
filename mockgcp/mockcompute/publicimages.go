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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// publicImage is an image published by a public GCP image project (such as debian-cloud or cos-cloud).
// Public images are not stored by the mock; they are served from publicImages instead.
type publicImage struct {
	project     string
	family      string
	name        string
	description string
	status      string
	diskSizeGb  int64
	// id is the image ID recorded on disks created from this image.
	id uint64
}

// Public images and families not listed here return NotFound, as a nonexistent image does on real GCP.
// Add an entry when a fixture or sample run against the mock needs another public image.
var publicImages = []publicImage{
	// The debian entries share the id 2443108620951880213 deliberately: it matches the value hardcoded in
	// disksv1.go and recorded in existing goldens (e.g. computenetworkendpointbasic, computemachineimageautogen).
	// Do not change it independently.
	{
		project:     "debian-cloud",
		family:      "debian-11",
		name:        "debian-11-bullseye-v20231010",
		description: "Debian, Debian GNU/Linux, debian-11, amd64",
		status:      "UP",
		id:          2443108620951880213,
	},
	{
		project:     "debian-cloud",
		family:      "debian-12",
		name:        "debian-12-bookworm-v20261023",
		description: "Debian, Debian GNU/Linux, debian-12, amd64",
		status:      "UP",
		id:          2443108620951880213,
	},
	{
		project:    "cos-cloud",
		family:     "cos-81-lts",
		name:       "cos-81-12871-1196-0",
		status:     "READY",
		diskSizeGb: 10,
		id:         1867341290458829913,
	},
	{
		// TODO: verify against real GCP (gcloud compute images describe-from-family cos-121-lts --project=cos-cloud)
		// and update name, id, status and diskSizeGb; all four are invented placeholders. Goldens that record this
		// image's name or id must be regenerated if they change.
		project:    "cos-cloud",
		family:     "cos-121-lts",
		name:       "cos-121-18867-90-42",
		status:     "READY",
		diskSizeGb: 10,
		id:         7519230484621035527,
	},
}

// isPublicImageProject returns true if project is a public image project served from publicImages.
func isPublicImageProject(project string) bool {
	for i := range publicImages {
		if publicImages[i].project == project {
			return true
		}
	}
	return false
}

func lookupPublicImageByFamily(project, family string) (*publicImage, bool) {
	for i := range publicImages {
		if publicImages[i].project == project && publicImages[i].family == family {
			return &publicImages[i], true
		}
	}
	return nil, false
}

func lookupPublicImageByName(project, name string) (*publicImage, bool) {
	for i := range publicImages {
		if publicImages[i].project == project && publicImages[i].name == name {
			return &publicImages[i], true
		}
	}
	return nil, false
}

// toProto returns the image as the images API returns it.
func (img *publicImage) toProto(ctx context.Context) *pb.Image {
	obj := &pb.Image{}
	obj.Kind = PtrTo("compute#image")
	obj.Name = PtrTo(img.name)
	obj.Family = PtrTo(img.family)
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, "projects/"+img.project+"/global/images/"+img.name))
	obj.Status = PtrTo(img.status)
	if img.description != "" {
		obj.Description = PtrTo(img.description)
	}
	if img.diskSizeGb != 0 {
		obj.DiskSizeGb = PtrTo(img.diskSizeGb)
	}
	return obj
}

// trimComputeURLPrefix strips the https://{www,compute}.googleapis.com/compute/<version>/ prefix from a
// compute URL, returning the resource path. Paths without the prefix are returned unchanged.
func trimComputeURLPrefix(ref string) string {
	for _, prefix := range []string{"https://www.googleapis.com/compute/", "https://compute.googleapis.com/compute/"} {
		if rest, ok := strings.CutPrefix(ref, prefix); ok {
			if i := strings.Index(rest, "/"); i != -1 {
				return rest[i+1:]
			}
			return ""
		}
	}
	return ref
}

func imageNotFound(fqn string) error {
	return status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
}

// findStoredImageByFamily returns the newest non-deprecated stored image in the project with the given family.
func (s *MockService) findStoredImageByFamily(ctx context.Context, projectID, family string) (*pb.Image, error) {
	var newest *pb.Image
	prefix := "projects/" + projectID + "/global/images/"
	if err := s.storage.List(ctx, (&pb.Image{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		img := obj.(*pb.Image)
		if img.GetFamily() != family || img.GetDeprecated() != nil {
			return nil
		}
		if newest == nil || img.GetId() > newest.GetId() {
			newest = img
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if newest == nil {
		return nil, imageNotFound(prefix + "family/" + family)
	}
	return proto.CloneOf(newest), nil
}

// resolveSourceImage resolves a source image reference, as used by a boot disk's initializeParams, to an image.
// Bare "<project>/<name>" references are tried as an image name first, then as a family.
func (s *MockService) resolveSourceImage(ctx context.Context, requestProject, ref string) (*pb.Image, error) {
	tokens := strings.Split(strings.TrimPrefix(trimComputeURLPrefix(ref), "/"), "/")
	switch {
	case len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "images" && tokens[4] == "family":
		return s.resolveImageFamily(ctx, tokens[1], tokens[5])
	case len(tokens) == 4 && tokens[0] == "global" && tokens[1] == "images" && tokens[2] == "family":
		return s.resolveImageFamily(ctx, requestProject, tokens[3])
	case len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "images":
		return s.resolveImageByName(ctx, tokens[1], tokens[4])
	case len(tokens) == 3 && tokens[0] == "global" && tokens[1] == "images":
		return s.resolveImageByName(ctx, requestProject, tokens[2])
	case len(tokens) == 2:
		img, err := s.resolveImageByName(ctx, tokens[0], tokens[1])
		if status.Code(err) == codes.NotFound {
			return s.resolveImageFamily(ctx, tokens[0], tokens[1])
		}
		return img, err
	}
	return nil, status.Errorf(codes.InvalidArgument, "source image %q is not valid", ref)
}

// resolveImageByName returns the image with the given name, including its ID.
func (s *MockService) resolveImageByName(ctx context.Context, project, name string) (*pb.Image, error) {
	if isPublicImageProject(project) {
		img, ok := lookupPublicImageByName(project, name)
		if !ok {
			return nil, imageNotFound("projects/" + project + "/global/images/" + name)
		}
		obj := img.toProto(ctx)
		obj.Id = PtrTo(img.id)
		return obj, nil
	}

	imageName, err := s.parseImageName(project, name)
	if err != nil {
		return nil, err
	}
	obj := &pb.Image{}
	if err := s.storage.Get(ctx, imageName.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, imageNotFound(imageName.String())
		}
		return nil, err
	}
	return obj, nil
}

// resolveImageFamily returns the image currently serving the given family, including its ID.
func (s *MockService) resolveImageFamily(ctx context.Context, project, family string) (*pb.Image, error) {
	if isPublicImageProject(project) {
		img, ok := lookupPublicImageByFamily(project, family)
		if !ok {
			return nil, imageNotFound(fmt.Sprintf("projects/%s/global/images/family/%s", project, family))
		}
		obj := img.toProto(ctx)
		obj.Id = PtrTo(img.id)
		return obj, nil
	}
	return s.findStoredImageByFamily(ctx, project, family)
}
