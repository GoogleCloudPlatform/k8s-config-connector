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

package identity

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"cloud.google.com/go/auth"
	"google.golang.org/grpc/credentials"
)

var errNoCredentials = errors.New("no GCP credentials for this call: it is not made under a context from ForObject")

// NewAuthCredentials returns credentials that authenticate each HTTP request
// with the credentials in its context; see ForObject. Use them with
// httptransport.AddAuthorizationMiddleware, not option.WithAuthCredentials,
// which can cache or ignore them.
func NewAuthCredentials() *auth.Credentials {
	return auth.NewCredentials(&auth.CredentialsOptions{TokenProvider: contextTokenProvider{}})
}

// NewPerRPCCredentials is the gRPC counterpart of NewAuthCredentials. Dial
// with option.WithoutAuthentication; the client library then omits the quota
// project, so pass it here.
func NewPerRPCCredentials(quotaProject string) credentials.PerRPCCredentials {
	return &perRPCCredentials{quotaProject: quotaProject}
}

type perRPCCredentials struct {
	quotaProject string
}

func (c *perRPCCredentials) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	token, err := contextTokenProvider{}.Token(ctx)
	if err != nil {
		return nil, err
	}
	md := map[string]string{"authorization": token.Type + " " + token.Value}
	if c.quotaProject != "" {
		md["x-goog-user-project"] = c.quotaProject
	}
	return md, nil
}

func (c *perRPCCredentials) RequireTransportSecurity() bool {
	return true
}

// BindCredentials returns a RoundTripper that sends each request whose context
// has no credentials with creds, through base. base must authenticate from the
// request's context, as a client with NewAuthCredentials does. Use it for
// libraries that make some calls without the context from ForObject.
func BindCredentials(base http.RoundTripper, creds *Credentials) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &boundTransport{base: base, creds: creds}
}

type boundTransport struct {
	base  http.RoundTripper
	creds *Credentials
}

func (t *boundTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if _, ok := CredentialsFromContext(req.Context()); !ok {
		req = req.WithContext(context.WithValue(req.Context(), credentialsKey{}, t.creds))
	}
	return t.base.RoundTrip(req)
}

// contextTokenProvider returns the token of the credentials in ctx.
type contextTokenProvider struct{}

func (contextTokenProvider) Token(ctx context.Context) (*auth.Token, error) {
	creds, ok := CredentialsFromContext(ctx)
	if !ok {
		return nil, errNoCredentials
	}
	token, err := creds.TokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("getting a GCP token for %s: %w", creds.ID, err)
	}
	return &auth.Token{Value: token.AccessToken, Type: token.Type(), Expiry: token.Expiry}, nil
}
