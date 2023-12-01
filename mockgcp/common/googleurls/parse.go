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

package googleurls

import (
	"fmt"
	"net/url"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

type Link struct {
	// Host is the host in the url, if set
	Host string

	// Service is the service name (e.g. compute), usually the first component of the url
	Service string

	// Version is the API version specified
	Version string

	// Project is the id (or number) of a project
	Project string

	// Global is true if the global parameter is explicitly specified
	Global bool

	// Region is the GCP region
	Region string

	// ResourceType is the name of the resource (e.g. serviceAttachment, instance, disk)
	ResourceType string

	// ResourceName is the name of the resource
	ResourceName string
}

func isResource(s string) bool {
	items := []string{
		"serviceAttachments",
		"targetVpnGateways",
	}
	resources := sets.NewString(items...)
	return resources.Has(s)
}

func ParseURL(s string) (*Link, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing %q: %w", s, err)
	}

	path := strings.TrimPrefix(u.Path, "/")
	tokens := strings.Split(path, "/")

	l := &Link{}
	l.Host = u.Host

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if i == 0 && token == "compute" {
			l.Service = token
			continue
		}

		if i == 1 && (token == "v1" || token == "beta") {
			l.Version = token
			continue
		}

		if token == "global" {
			l.Global = true
			continue
		}

		if (i + 1) < len(tokens) {
			nextToken := tokens[i+1]
			if token == "projects" {
				l.Project = nextToken
				i++
				continue
			}
			if token == "regions" {
				l.Region = nextToken
				i++
				continue
			}

			if isResource(token) {
				l.ResourceType = token
				l.ResourceName = nextToken
				i++
				continue
			}
		}
		return nil, fmt.Errorf("unknown token %q in %q", token, s)
	}

	return l, nil
}
