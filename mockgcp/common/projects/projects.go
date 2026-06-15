// Copyright 2022 Google LLC
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

package projects

import (
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectStore interface {
	GetProject(project *ProjectName) (*ProjectData, error)
	GetProjectByID(projectID string) (*ProjectData, error)

	// GetProjectByNumber returns the project with the specified project number, or an error if not found.
	// Note that the project number must still be passed as a string, to keep terraform happy.
	GetProjectByNumber(projectNumberAsString string) (*ProjectData, error)

	// GetProjectByIDOrNumber will return the project by the id or number provided.
	GetProjectByIDOrNumber(projectIDOrNumber string) (*ProjectData, error)
}

type ProjectData struct {
	Number int64
	ID     string
}

type ProjectName struct {
	ProjectID     string
	ProjectNumber int64
	OriginalValue string
}

func (n *ProjectName) String() string {
	return "projects/" + n.OriginalValue
}

// ParseProjectName parses a string into a ProjectName.
// The expected form is projects/<projectIDOrNumber>
func ParseProjectName(name string) (*ProjectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		return ParseProjectIDOrNumber(tokens[1])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// ParseProjectIDOrNumber parses a string into a ProjectName.
// The expected form is <projectID> or <projectNumber> (without a projects/ prefix)
func ParseProjectIDOrNumber(s string) (*ProjectName, error) {
	name := &ProjectName{
		OriginalValue: s,
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		name.ProjectNumber = n
	} else {
		name.ProjectID = s
	}

	return name, nil
}
