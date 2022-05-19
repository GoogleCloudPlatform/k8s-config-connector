// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package privateca contains methods and objects for handling privateca GCP resources.
package beta

import (
	"bytes"
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// Waits for the state of the certificate authority to be "ENABLED".
func (r *CertificateAuthority) waitForCertificateAuthorityEnabled(ctx context.Context, c *Client) error {
	return dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		nr, err := c.GetCertificateAuthority(ctx, r)
		if err != nil {
			return nil, err
		}
		if *nr.State == *CertificateAuthorityStateEnumRef("ENABLED") || *nr.State == *CertificateAuthorityStateEnumRef("STAGED") {
			return nil, nil
		}
		return &dcl.RetryDetails{}, dcl.OperationNotDone{}
	}, c.Config.RetryProvider)
}

// Disables the certificate authority so that it can be deleted.
func (r *CertificateAuthority) disableCertificateAuthority(ctx context.Context, c *Client) error {
	if *r.State != *CertificateAuthorityStateEnumRef("ENABLED") {
		// Only enabled certificate authorities need to be disabled before deletion.
		return nil
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"ca_pool":  dcl.ValueOrEmptyString(nr.CaPool),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	u := dcl.URL("projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}/certificateAuthorities/{{name}}:disable", "https://privateca.googleapis.com/v1/", c.Config.BasePath, params)
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for certificate authority to be disabled.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://privateca.googleapis.com/v1beta1/", "GET"); err != nil {
		return err
	}
	return nil
}

func (r *Certificate) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
		"caPool":               dcl.ValueOrEmptyString(nr.CaPool),
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"certificateAuthority": dcl.ValueOrEmptyString(nr.CertificateAuthority),
	}
	basePath := dcl.URL("projects/{{project}}/locations/{{location}}/caPools/{{caPool}}/certificates", nr.basePath(), userBasePath, params)

	var err error
	if dcl.ValueOrEmptyString(nr.CertificateAuthority) != "" {
		basePath, err = dcl.AddQueryParams(basePath, map[string]string{"issuingCertificateAuthorityId": dcl.ValueOrEmptyString(nr.CertificateAuthority)})
		if err != nil {
			return "", err
		}
	}
	if dcl.ValueOrEmptyString(nr.Name) != "" {
		// Need to set name to nil or else it'll trigger error messages.
		basePath, err = dcl.AddQueryParams(basePath, map[string]string{"certificateId": dcl.ValueOrEmptyString(nr.Name)})
		if err != nil {
			return "", err
		}
	}
	return basePath, nil
}

func flattenCertificateConfigX509ConfigCAOptions(_ *Client, i interface{}, _ *Certificate) *CertificateConfigX509ConfigCaOptions {
	if i == nil {
		return nil
	}
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	result := &CertificateConfigX509ConfigCaOptions{}

	isCA, ok := m["isCa"].(bool)
	if ok {
		result.IsCa = dcl.Bool(isCA)
		if !isCA {
			result.NonCa = dcl.Bool(true)
		}
	}

	if _, ok := m["maxIssuerPathLength"]; ok {
		pathLen := dcl.FlattenInteger(m["maxIssuerPathLength"])
		result.MaxIssuerPathLength = pathLen
		if dcl.ValueOrEmptyInt64(pathLen) == 0 {
			result.ZeroMaxIssuerPathLength = dcl.Bool(true)
		}
	}

	return result
}

func expandCertificateConfigX509ConfigCAOptions(_ *Client, caOptions *CertificateConfigX509ConfigCaOptions, _ *Certificate) (map[string]interface{}, error) {
	if caOptions == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	isCA := dcl.ValueOrEmptyBool(caOptions.IsCa)
	nonCA := dcl.ValueOrEmptyBool(caOptions.NonCa)
	zeroPathLength := dcl.ValueOrEmptyBool(caOptions.ZeroMaxIssuerPathLength)
	maxIssuerPathLength := dcl.ValueOrEmptyInt64(caOptions.MaxIssuerPathLength)

	if !isCA && !nonCA {
		return nil, nil
	} else if isCA && nonCA {
		return nil, fmt.Errorf("is_ca and non_ca are mutually exclusive")
	} else if isCA || nonCA {
		m["isCa"] = isCA
	}

	if zeroPathLength && maxIssuerPathLength > 0 {
		return nil, fmt.Errorf("max_issuer_path_length and zero_max_issuer_path_length are mutually exclusive")
	}
	if maxIssuerPathLength > 0 || zeroPathLength {
		m["maxIssuerPathLength"] = maxIssuerPathLength
	}

	return m, nil
}
