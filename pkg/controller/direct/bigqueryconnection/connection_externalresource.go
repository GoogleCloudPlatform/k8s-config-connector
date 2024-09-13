package bigqueryconnection

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the BigQueryConnectionConnection resource from the GCP service.
type BigQueryConnectionConnectionIdentity struct {
	Parent     *parent
	Connection string
}

type parent struct {
	Project  string
	Location string
}

func (p *parent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *BigQueryConnectionConnectionIdentity) FullyQualifiedName() string {
	// TODO(user): Edit the URL path
	return fmt.Sprintf("%s/connections/%s", c.Parent, c.Connection)
}

// AsExternalRef builds a externalRef from a BigQueryConnectionConnection
func (c *BigQueryConnectionConnectionIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a BigQueryConnectionConnectionIdentity from a `status.externalRef`
func asID(externalRef string) (*BigQueryConnectionConnectionIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/connections/<Connection>, got %s",
			serviceDomain, externalRef)
	}
	return &BigQueryConnectionConnectionIdentity{
		Parent:     &parent{Project: tokens[1], Location: tokens[3]},
		Connection: tokens[5],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the BigQueryConnectionConnection resource from the GCP service.
func BuildID(project, location, resourceID string) *BigQueryConnectionConnectionIdentity {
	// TODO(user): Build resource identity from resource components, i.e. project, location, resource id
	return &BigQueryConnectionConnectionIdentity{
		Parent:     &parent{Project: project, Location: location},
		Connection: resourceID,
	}
}
