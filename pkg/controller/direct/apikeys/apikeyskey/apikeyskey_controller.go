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

package apikeyskey

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"k8s.io/klog/v2"

	pb "cnrm.googlesource.com/cnrm/mockgcp/generated/google/api/apikeys/v2"
	"cnrm.googlesource.com/cnrm/pkg/clients/generated/apis/apikeys/v1alpha1"
	"cnrm.googlesource.com/cnrm/pkg/controller/direct/debug"
	"cnrm.googlesource.com/cnrm/pkg/controller/direct/directbase"
	"cnrm.googlesource.com/cnrm/pkg/controller/direct/mappings"
	"cnrm.googlesource.com/cnrm/pkg/dcl/conversion"
	"cnrm.googlesource.com/cnrm/pkg/servicemapping/servicemappingloader"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/encoding/protojson"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Add creates a new IAM Policy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and start it when the Manager is started.
func Add(mgr manager.Manager, tfProvider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) error {
	gvk := schema.GroupVersionKind{
		Group:   "apikeys.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "APIKeysKey",
	}
	return directbase.Add(mgr, tfProvider, smLoader, converter, dclConfig, gvk, &model{})
}

type model struct {
}

func init() {
	mappings.Add(&v1alpha1.APIKeysKeySpec{}, &pb.Key{},
		mappings.Simple("displayName"),
		mappings.Simple("restrictions"),
	)
	// TODO: Auto convert reverse
	mappings.Add(&pb.Key{}, &v1alpha1.APIKeysKeySpec{},
		mappings.Simple("displayName"),
		mappings.Simple("restrictions"),
	)
	mappings.Add(&v1alpha1.KeyRestrictions{}, &pb.Restrictions{},
		mappings.Simple("apiTargets"),
	)
	mappings.Add(&pb.Restrictions{}, &v1alpha1.KeyRestrictions{},
		mappings.Simple("apiTargets"),
	)
	mappings.Add(&v1alpha1.KeyApiTargets{}, &pb.ApiTarget{},
		mappings.Simple("methods"),
		mappings.Simple("service"),
	)
	mappings.Add(&pb.ApiTarget{}, &v1alpha1.KeyApiTargets{},
		mappings.Simple("methods"),
		mappings.Simple("service"),
	)
}

type adapter struct {
	projectID string
	location  string
	keyID     string

	desired *v1alpha1.APIKeysKeySpec
	actual  *v1alpha1.APIKeysKeySpec

	gcp *gcpClient
}

func (*model) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	// TODO: Get from context?   We want our hooks though...
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
	}

	gcpClient := &gcpClient{
		httpClient: httpClient,
	}

	// TODO: Just fetch this object?
	obj := &v1alpha1.APIKeysKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	keyID := ValueOf(obj.Spec.ResourceID)
	if keyID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	location := "global"

	return &adapter{
		projectID: projectID,
		location:  location,
		keyID:     keyID,
		desired:   &obj.Spec,
		gcp:       gcpClient,
	}, nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	key, err := a.gcp.get(ctx, a.projectID, a.location, a.keyID)
	if IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	u := &v1alpha1.APIKeysKeySpec{}
	if err := mappings.Map(key, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *adapter) Delete(ctx context.Context) error {
	// TODO: Delete from status?
	err := a.gcp.delete(ctx, a.projectID, a.location, a.keyID)
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) BuildCreate(ctx context.Context) (*pb.CreateKeyRequest, error) {
	desired := &pb.Key{}
	if err := mappings.Map(a.desired, desired); err != nil {
		return nil, err
	}
	return &pb.CreateKeyRequest{Key: desired}, nil
}

func (a *adapter) Create(ctx context.Context) (*unstructured.Unstructured, error) {
	// desired := &UnstructuredKey{}
	// if err := mapKubeToCloud(a.desired, desired); err != nil {
	// 	return nil, err
	// }

	req, err := a.BuildCreate(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := a.gcp.create(ctx, a.projectID, a.location, a.keyID, req.Key); err != nil {
		return nil, err
	}
	// TODO: Return created object
	return nil, nil
}

func (a *adapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	// TODO: Return updated object

	desiredCloud := &pb.Key{}
	if err := mappings.Map(a.desired, desiredCloud); err != nil {
		return nil, err
	}

	// patch := &pb.Key{} //Data: make(map[string]interface{})}
	// if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
	// 	patch.Data["displayName"] = desiredCloud.["displayName"]
	// }
	// if !reflect.DeepEqual(a.desired.Restrictions, a.actual.Restrictions) {
	// 	patch.Data["restrictions"] = desiredCloud.Data["restrictions"]
	// }

	delta := &v1alpha1.APIKeysKeySpec{} //Data: make(map[string]interface{})}
	diffCount := 0
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		delta.DisplayName = a.desired.DisplayName
		diffCount++
	}
	if !reflect.DeepEqual(a.desired.Restrictions, a.actual.Restrictions) {
		delta.Restrictions = a.desired.Restrictions
		diffCount++
	}
	klog.Infof("desired: %v", debug.JSON(a.desired))
	klog.Infof("actual: %v", debug.JSON(a.actual))
	klog.Infof("delta: %v", debug.JSON(delta))

	if diffCount == 0 {
		// TODO: Log/warn/error?
		return nil, nil
	}

	patch := &pb.Key{}
	if err := mappings.Map(delta, patch); err != nil {
		return nil, err
	}
	_, err := a.gcp.patch(ctx, a.projectID, a.location, a.keyID, patch)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type gcpClient struct {
	httpClient *http.Client
}

func (c *gcpClient) buildURL(projectID, location, keyID string) (string, error) {
	url := "https://apikeys.googleapis.com/v2/projects/" + projectID + "/locations/" + location + "/keys"
	if keyID != "" {
		url += "/" + keyID
	}
	return url, nil
}

func (r *gcpClient) delete(ctx context.Context, projectID, location, keyID string) error {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return err
	}
	method := "DELETE"
	var body io.Reader

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}

	b, err := r.do(ctx, req)
	if err != nil {
		return err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(b))
	return nil
}

func (r *gcpClient) do(ctx context.Context, req *http.Request) ([]byte, error) {
	method := req.Method
	url := req.URL.String()

	httpClient := r.httpClient
	httpClient = transport_tpg.DefaultHTTPClientTransformer(ctx, httpClient)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from http request %s %s: %w", method, url, err)
	}
	defer resp.Body.Close()
	if err := googleapi.CheckResponse(resp); err != nil {
		return nil, err
		// return nil, fmt.Errorf("unexpected result from http request %s %s: %v", method, url, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http response from %s %s: %w", method, url, err)
	}

	return b, nil
}
func (r *gcpClient) get(ctx context.Context, projectID, location, keyID string) (*pb.Key, error) {
	// log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return nil, err
	}
	method := "GET"
	var body io.Reader

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	b, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// m := make(map[string]interface{})
	// if err := json.Unmarshal(b, &m); err != nil {
	// 	return nil, fmt.Errorf("error parsing http response from %s %s: %w", method, url, err)
	// }
	// return &Key{Data: m}, nil

	k := &pb.Key{}
	if err := protojson.Unmarshal(b, k); err != nil {
		return nil, fmt.Errorf("error parsing http response from %s %s: %w", method, url, err)
	}
	return k, nil
}

func (r *gcpClient) create(ctx context.Context, projectID, location, keyID string, key *pb.Key) (*string, error) {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, "")
	if err != nil {
		return nil, err
	}
	if keyID != "" {
		// TODO: encoding etc
		url += "?keyId=" + keyID
	}
	method := "POST"

	reqBody, err := protojson.Marshal(key)
	// reqBody, err := json.Marshal(key.Data)
	if err != nil {
		return nil, fmt.Errorf("error building request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	respBody, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(respBody))
	return nil, nil
}

func (r *gcpClient) patch(ctx context.Context, projectID, location, keyID string, key *pb.Key) (*string, error) {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return nil, err
	}
	method := "PATCH"

	reqBody, err := protojson.Marshal(key)
	// reqBody, err := json.Marshal(key.Data)
	if err != nil {
		return nil, fmt.Errorf("error building request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	respBody, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(respBody))
	return nil, nil
}

// IsNotFound reports whether err is the result of the
// server replying with http.StatusNotFound.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	ae, ok := err.(*googleapi.Error)
	return ok && ae.Code == http.StatusNotFound
}
