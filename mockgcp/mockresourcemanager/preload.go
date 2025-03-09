package mockresourcemanager

import (
	"context"
	"fmt"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsPreload = &MockService{}

func (s *MockService) Preload(ctx context.Context) error {
	if err := s.preloadProjects(ctx); err != nil {
		return fmt.Errorf("preloading projects: %w", err)
	}
	return nil
}

func (s *MockService) preloadProjects(ctx context.Context) error {
	for _, project := range preloadProjects {
		fqn := "projects/" + project.ProjectId
		if err := s.storage.Create(ctx, fqn, project); err != nil {
			return fmt.Errorf("preloading project %q: %v", fqn, err)
		}
	}
	return nil
}

var preloadProjects = []*pb.Project{
	{
		ProjectId: "debian-cloud",
		Name:      "projects/10001",
		State:     pb.Project_ACTIVE,
	},
}
