/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loaders

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

// HTTPRepository supports loading from http / https
type HTTPRepository struct {
	baseURL string
}

var _ Repository = &HTTPRepository{}

// NewHTTPRepository constructs an HTTPRepository
func NewHTTPRepository(baseURL string) *HTTPRepository {
	return &HTTPRepository{
		baseURL: baseURL,
	}
}

func (r *HTTPRepository) LoadChannel(ctx context.Context, name string) (*Channel, error) {
	if !allowedChannelName(name) {
		return nil, fmt.Errorf("invalid channel name: %q", name)
	}

	log := log.FromContext(ctx)
	log.WithValues("channel", name).WithValues("baseURL", r.baseURL).Info("loading channel")

	p := r.makeURL(name)
	b, err := r.readURL(ctx, p)
	if err != nil {
		log.WithValues("path", p).Error(err, "error reading channel")
		return nil, fmt.Errorf("error reading channel %s: %v", p, err)
	}

	channel := &Channel{}
	if err := yaml.Unmarshal(b, channel); err != nil {
		return nil, fmt.Errorf("error parsing channel %s: %v", p, err)
	}

	return channel, nil
}

func (r *HTTPRepository) LoadManifest(ctx context.Context, packageName string, id string) (map[string]string, error) {
	if !allowedManifestId(packageName) {
		return nil, fmt.Errorf("invalid package name: %q", id)
	}

	if !allowedManifestId(id) {
		return nil, fmt.Errorf("invalid manifest id: %q", id)
	}

	log := log.FromContext(ctx)
	log.WithValues("package", packageName).Info("loading package")

	p := r.makeURL("packages", packageName, id, "manifest.yaml")
	b, err := r.readURL(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("error reading package %s: %v", p, err)
	}
	result := map[string]string{
		p: string(b),
	}
	return result, nil
}

// makeURL joins the paths to the baseURL
func (r *HTTPRepository) makeURL(paths ...string) string {
	u := r.baseURL
	for _, path := range paths {
		if !strings.HasSuffix(u, "/") {
			u += "/"
		}
		u += path
	}
	return u
}

// readURL tries to fetch the specified url
func (r *HTTPRepository) readURL(ctx context.Context, url string) ([]byte, error) {
	log := log.FromContext(ctx)
	log.WithValues("url", url).Info("doing HTTP request")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching %q: %v", url, err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response for %q: %v", url, err)
	}

	if response.StatusCode == 200 {
		return body, nil
	}

	return nil, fmt.Errorf("unexpected response code %q fetching %q: %v", response.Status, url, string(body))
}
