// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package projects

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/concurrent"
)

// ProjectMapper maps between project IDs and project numbers.
// It uses a ProjectCache to minimize API calls.
type ProjectMapper struct {
	cache *ProjectCache
}

// ProjectCache is a simple in-memory cache for project ID to project number mappings.
type ProjectCache struct {
	client *resourcemanager.ProjectsClient

	expiration       time.Duration
	projectsByID     *concurrent.ExpiringMap[string, *cachedProjectInfo]
	projectsByNumber *concurrent.ExpiringMap[int64, *cachedProjectInfo]
}

// cachedProjectInfo holds the cached information for a project.
type cachedProjectInfo struct {
	projectID     string
	projectNumber int64
	timestamp     time.Time
}

// NewProjectCache creates a new ProjectCache with the specified expiration duration.
func NewProjectCache(client *resourcemanager.ProjectsClient, expiration time.Duration) *ProjectCache {
	isValid := func(value *cachedProjectInfo) bool {
		return time.Since(value.timestamp) < expiration
	}

	return &ProjectCache{
		client:           client,
		expiration:       expiration,
		projectsByID:     concurrent.NewExpiringMap[string, *cachedProjectInfo](isValid),
		projectsByNumber: concurrent.NewExpiringMap[int64, *cachedProjectInfo](isValid),
	}
}

// NewProjectMapper creates a new ProjectMapper with the given ProjectCache.
// This is used to efficiently map between project IDs and numbers.
func NewProjectMapper(cache *ProjectCache) *ProjectMapper {
	return &ProjectMapper{
		cache: cache,
	}
}

// ReplaceProjectNumberWithIDInLink replaces the project number with the project ID in a link.
// It assumes the link contains segments like "projects/{projectNumber}".
func (m *ProjectMapper) ReplaceProjectNumberWithIDInLink(ctx context.Context, link string) (string, error) {
	return replaceProjectNumberWithIDInLink(ctx, link, m.ReplaceProjectNumberWithID)
}

// projectRegexp matches "projects/" followed by one or more non-slash characters.
var projectRegexp = regexp.MustCompile(`projects/([^/]+)`)

func replaceProjectNumberWithIDInLink(ctx context.Context, link string, resolver func(context.Context, string) (string, error)) (string, error) {
	// FindAllStringSubmatchIndex returns a slice of start/end indices for the match and submatches.
	matches := projectRegexp.FindAllStringSubmatchIndex(link, -1)
	if matches == nil {
		return link, nil
	}

	var sb strings.Builder
	lastIndex := 0

	for _, match := range matches {
		// match[0], match[1] are start/end of the full match "projects/XYZ"
		// match[2], match[3] are start/end of the submatch "XYZ"

		fullStart, fullEnd := match[0], match[1]
		subStart, subEnd := match[2], match[3]

		// Append text before the match
		sb.WriteString(link[lastIndex:fullStart])

		projectSegment := link[subStart:subEnd]
		projectID, err := resolver(ctx, projectSegment)
		if err != nil {
			return "", err
		}

		sb.WriteString("projects/")
		sb.WriteString(projectID)

		lastIndex = fullEnd
	}

	sb.WriteString(link[lastIndex:])

	return sb.String(), nil
}

func (m *ProjectMapper) ReplaceProjectNumberWithID(ctx context.Context, projectID string) (string, error) {
	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		// Not a project number, no need to map
		return projectID, nil
	}

	info, err := m.cache.lookupByProjectNumber(ctx, projectNumber)
	if err != nil {
		return "", err
	}
	return info.projectID, nil
}

// LookupProjectNumber retrieves the project number for a given project ID.
// If the project ID is actually a project number, it returns it directly.
func (m *ProjectMapper) LookupProjectNumber(ctx context.Context, projectID string) (int64, error) {
	// Check if the project number is already a valid integer
	// If not, we need to look it up
	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		info, err := m.cache.lookupByProjectID(ctx, projectID)
		if err != nil {
			return 0, err
		}
		projectNumber = info.projectNumber
	}
	return projectNumber, nil
}

// lookupByProjectID retrieves the cached project info by project ID,
// or fetches it from the API if not present or expired.
func (c *ProjectCache) lookupByProjectID(ctx context.Context, projectID string) (*cachedProjectInfo, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project ID was unexpectedly empty")
	}

	lookup := func() (*cachedProjectInfo, error) {
		projectInfo, err := c.queryForProject(ctx, "projects/"+projectID)
		if err != nil {
			return nil, err
		}
		return projectInfo, nil
	}

	// Use ConcurrentMap's GetOrCompute to avoid duplicate fetches
	info, err := c.projectsByID.GetOrCompute(projectID, lookup)
	if err != nil {
		return nil, err
	}

	// Update the reverse mapping as well
	c.projectsByNumber.Set(info.projectNumber, info)

	return info, nil
}

// lookupByProjectNumber retrieves the cached project info by project number,
// or fetches it from the API if not present or expired.
func (c *ProjectCache) lookupByProjectNumber(ctx context.Context, projectNumber int64) (*cachedProjectInfo, error) {
	if projectNumber == 0 {
		return nil, fmt.Errorf("project number was unexpectedly zero")
	}

	lookup := func() (*cachedProjectInfo, error) {
		name := "projects/" + strconv.FormatInt(projectNumber, 10)
		projectInfo, err := c.queryForProject(ctx, name)
		if err != nil {
			return nil, err
		}
		return projectInfo, nil
	}

	// Use ConcurrentMap's GetOrCompute to avoid duplicate fetches
	info, err := c.projectsByNumber.GetOrCompute(projectNumber, lookup)
	if err != nil {
		return nil, err
	}

	// Update the reverse mapping as well
	c.projectsByID.Set(info.projectID, info)

	return info, nil
}

// queryForProject fetches the project information from the Cloud Resource Manager API.
func (c *ProjectCache) queryForProject(ctx context.Context, name string) (*cachedProjectInfo, error) {
	// We use a relatively short timeout because the concurrent map is probably holding a lock here.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req := &resourcemanagerpb.GetProjectRequest{
		Name: name,
	}
	project, err := c.client.GetProject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error getting project %q: %w", req.Name, err)
	}

	tokens := strings.Split(project.Name, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, fmt.Errorf("unexpected project name format: %q", project.Name)
	}

	projectNumber, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
	}

	return &cachedProjectInfo{
		projectID:     project.ProjectId,
		projectNumber: projectNumber,
		timestamp:     time.Now(),
	}, nil
}
