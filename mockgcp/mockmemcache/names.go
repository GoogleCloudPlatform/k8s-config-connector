package mockmemcache

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Instance
}

func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &instanceName{
			Project:  project,
			Location: tokens[3],
			Instance: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
