// Copyright 2025 Google LLC
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

package preview

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"k8s.io/klog/v2"
)

// Object is the fields we need from a Kubernetes object.
type Object interface {
	GetNamespace() string
	GetName() string
}

// ClientOptions are the options for creating a StreamingClient.
type ClientOptions struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
}

// StreamingClient is a client for streaming Kubernetes API responses.
type StreamingClient struct {
	httpClient *http.Client
	baseURL    url.URL
}

// GroupVersionResource is a group, version, and resource.
type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}

// GroupKind is a group and kind.
type GroupKind struct {
	Group string
	Kind  string
}

type KubeClient interface {
	Get(ctx context.Context, typeInfo *typeInfo, namespace, name string, dest Object) error
	List(ctx context.Context, typeInfo *typeInfo, listener ListListener) error
	Watch(ctx context.Context, typeInfo *typeInfo, watchOptions WatchOptions, listener WatchListener) error
}

// NewStreamingClient creates a new StreamingClient.
func NewStreamingClient(opt ClientOptions) *StreamingClient {
	return &StreamingClient{
		httpClient: opt.HTTPClient,
		baseURL:    *opt.BaseURL,
	}
}

// ListMetadata is the metadata for a list operation.
type ListMetadata struct {
	APIVersion      string
	Kind            string
	ResourceVersion string
}

// ListListener is a listener for list operations.
type ListListener interface {
	OnListBegin(metadata ListMetadata)
	OnListObject(obj Object) error
	OnListEnd()
}

// userAgent returns the user agent for the StreamingClient.
func (c *StreamingClient) userAgent() string {
	return "StreamingClient" // TODO
}

// setHeaders sets the headers for the StreamingClient.
func (c *StreamingClient) setHeaders(request *http.Request) error {
	request.Header.Set("User-Agent", c.userAgent())
	return nil
}

// Get gets the requested object
func (c *StreamingClient) Get(ctx context.Context, typeInfo *typeInfo, namespace, name string, dest Object) error {
	log := klog.FromContext(ctx)

	u := c.resourceURL(typeInfo.gvr, namespace, name)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("building http request: %w", err)
	}
	if err := c.setHeaders(request); err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")

	log.Info("doing http request", "method", request.Method, "url", request.URL)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("sending http request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status from %v: %v", u, response.Status)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("reading response body from %v: %w", u, err)
	}

	if err := json.Unmarshal(b, dest); err != nil {
		return fmt.Errorf("decoding %T from %v: %w", dest, u, err)
	}

	return nil
}

// List lists the objects for the given type.
func (c *StreamingClient) List(ctx context.Context, typeInfo *typeInfo, listener ListListener) error {
	log := klog.FromContext(ctx)

	u := c.resourceURL(typeInfo.gvr, "", "")

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("building http request: %w", err)
	}
	if err := c.setHeaders(request); err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")

	log.Info("doing http request", "method", request.Method, "url", request.URL)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("sending http request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status from %v: %v", u, response.Status)
	}

	// TODO: Implement true streaming parsing (likely with https://github.com/golang/go/issues/71497)
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("reading response body from %v: %w", u, err)
	}

	type listT struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"metadata"`
		Items []json.RawMessage `json:"items"`
	}
	var list listT
	if err := json.Unmarshal(b, &list); err != nil {
		return fmt.Errorf("decoding response body from %v: %w", u, err)
	}

	listener.OnListBegin(ListMetadata{
		APIVersion:      list.APIVersion,
		Kind:            list.Kind,
		ResourceVersion: list.Metadata.ResourceVersion,
	})

	for _, item := range list.Items {
		t := typeInfo.factory()
		if err := json.Unmarshal(item, &t); err != nil {
			return fmt.Errorf("decoding %T from %v: %w", t, u, err)
		}
		if err := listener.OnListObject(t); err != nil {
			return err
		}
	}

	listener.OnListEnd()

	return nil
}

// WatchOptions are the options for a watch operation.
type WatchOptions struct {
	ResourceVersion     string
	AllowWatchBookmarks bool
}

// WatchListener is a listener for watch operations.
type WatchListener interface {
	OnWatchEvent(eventType string, object Object) error
}

// Watch watches the given type.
func (c *StreamingClient) Watch(ctx context.Context, typeInfo *typeInfo, watchOptions WatchOptions, listener WatchListener) error {
	log := klog.FromContext(ctx)

	u := c.resourceURL(typeInfo.gvr, "", "")

	q := u.Query()
	q.Set("watch", "true")
	if watchOptions.ResourceVersion != "" {
		q.Set("resourceVersion", watchOptions.ResourceVersion)
	}
	if watchOptions.AllowWatchBookmarks {
		q.Set("allowWatchBookmarks", "true")
	}
	u.RawQuery = q.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("building http request: %w", err)
	}
	if err := c.setHeaders(request); err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")

	log.Info("doing http request", "method", request.Method, "url", request.URL)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("sending http request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status from %v: %v", u, response.Status)
	}

	type eventT struct {
		Type   string          `json:"type"`
		Object json.RawMessage `json:"object"`
	}
	lineReader := &lineSplitReader{inner: response.Body}
	for {
		var event eventT

		jsonDecoder := json.NewDecoder(lineReader)
		if err := jsonDecoder.Decode(&event); err != nil {
			return fmt.Errorf("decoding event from %v: %w", u, err)
		}

		var object Object
		if len(event.Object) != 0 {
			object = typeInfo.factory()
			if err := json.Unmarshal(event.Object, object); err != nil {
				return fmt.Errorf("decoding %T from %v: %w", object, u, err)
			}
		}
		if err := listener.OnWatchEvent(event.Type, object); err != nil {
			return err
		}

		// reset lineEOF to read next line
		lineReader.lineEOF = false
	}
}

// lineSplitReader is a reader that splits lines from an inner reader.
// It is similar to io.LimitedReader, but it stops at a newline.
type lineSplitReader struct {
	inner    io.Reader
	buffer   []byte
	innerEOF bool

	lineEOF bool
}

// lineReaderBufferSize is the size of the buffer for the lineSplitReader.
const lineReaderBufferSize = 8192

// Read reads from the inner reader.
var _ io.Reader = &lineSplitReader{}

// Read reads a line from the inner reader.
// It is similar to io.LimitedReader, but it stops at a newline.
func (r *lineSplitReader) Read(p []byte) (int, error) {
	// Send io.EOF until we are reset
	if r.lineEOF {
		return 0, io.EOF
	}
	if len(r.buffer) == 0 {
		if r.buffer == nil {
			r.buffer = make([]byte, lineReaderBufferSize)
		}
		n, err := r.inner.Read(r.buffer[0:cap(r.buffer)])
		if err != nil {
			if err == io.EOF {
				r.innerEOF = true
			} else {
				return 0, err
			}
		}

		if n == 0 {
			if r.innerEOF {
				return 0, io.EOF
			}
		}

		r.buffer = r.buffer[:n]
		klog.Infof("buffer is %v", string(r.buffer))
	}

	// Only return up to the newline
	ret := r.buffer
	nl := bytes.IndexByte(r.buffer, '\n')
	if nl != -1 {
		ret = r.buffer[:nl+1]
	}

	// Return what fits into the target buffer
	n := copy(p, ret)
	r.buffer = r.buffer[n:]
	if len(r.buffer) == 0 {
		// release the buffer (maybe a sync pool would be faster)
		r.buffer = nil
	}
	klog.Infof("returning %v", string(p[:n]))

	// If we sent a newline, pretend this is EOF
	if nl != -1 && n == nl+1 {
		r.lineEOF = true
		return n, io.EOF
	}
	return n, nil
}

// resourceURL constructs a URL for the given GroupVersionResource.
func (c *StreamingClient) resourceURL(gvr GroupVersionResource, namespace, name string) *url.URL {
	u := &c.baseURL
	if gvr.Group != "" {
		u = u.JoinPath("apis", gvr.Group, gvr.Version)
	} else {
		u = u.JoinPath("api", gvr.Version)
	}
	if namespace != "" {
		u = u.JoinPath("namespaces", namespace)
	}
	if name != "" {
		u = u.JoinPath(gvr.Resource, name)
	} else {
		u = u.JoinPath(gvr.Resource)
	}
	return u
}
