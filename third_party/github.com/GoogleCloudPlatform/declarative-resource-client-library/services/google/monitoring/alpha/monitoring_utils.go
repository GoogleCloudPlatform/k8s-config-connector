// Copyright 2024 Google LLC. All Rights Reserved.
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
// Package monitoring provides methods and types for managing monitoring GCP resources.
package alpha

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
)

func equalsMetricDescriptorValueType(m, n *MetricDescriptorValueTypeEnum) bool {
	mStr := dcl.ValueOrEmptyString(m)
	if mStr == "" {
		mStr = "STRING"
	}
	nStr := dcl.ValueOrEmptyString(n)
	if nStr == "" {
		nStr = "STRING"
	}
	return mStr == nStr
}

func equalsMetricDescriptorLabelsValueType(m, n *MetricDescriptorLabelsValueTypeEnum) bool {
	mStr := dcl.ValueOrEmptyString(m)
	if mStr == "" {
		mStr = "STRING"
	}
	nStr := dcl.ValueOrEmptyString(n)
	if nStr == "" {
		nStr = "STRING"
	}
	return mStr == nStr
}

func canonicalizeMetricDescriptorValueType(m, n any) bool {
	if m == nil && n == nil {
		return true
	}

	mVal, _ := m.(*MetricDescriptorValueTypeEnum)
	nVal, _ := n.(*MetricDescriptorValueTypeEnum)
	return equalsMetricDescriptorValueType(mVal, nVal)
}

func canonicalizeMetricDescriptorLabelsValueType(m, n any) bool {
	if m == nil && n == nil {
		return true
	}

	mVal, _ := m.(*MetricDescriptorLabelsValueTypeEnum)
	nVal, _ := n.(*MetricDescriptorLabelsValueTypeEnum)
	return equalsMetricDescriptorLabelsValueType(mVal, nVal)
}

// GetMonitoredProject is a custom method because projects are returned as project numbers instead of project ids.
func (c *Client) GetMonitoredProject(ctx context.Context, r *MonitoredProject) (*MonitoredProject, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getMonitoredProjectRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMonitoredProject(b, c, r)
	if err != nil {
		return nil, err
	}
	result.MetricsScope = r.MetricsScope
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMonitoredProjectNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) getMonitoredProjectRaw(ctx context.Context, r *MonitoredProject) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	// URL Normalize the resource to get project by short name.
	nr := r.urlNormalized()
	// Create a client with an empty base path so that it doesn't inherit the base path from the
	// monitoring client.
	cloudresourcemanagerCl := cloudresourcemanager.NewClient(c.Config.Clone(dcl.WithBasePath("")))
	project, err := cloudresourcemanagerCl.GetProject(ctx, &cloudresourcemanager.Project{
		Name: nr.Name,
	})
	if err != nil {
		c.Config.Logger.Warningf("Could not look up project %s: %v", *nr.Name, err)
	}

	isElement := r.customMatcher(ctx, dcl.ValueOrEmptyInt64(project.ProjectNumber), c)
	b, err = dcl.ExtractElementFromList(b, "monitoredProjects", isElement)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// This resource has a custom matcher to do a look for a match by either project id or project number.
func (r *MonitoredProject) customMatcher(ctx context.Context, projectNumber int64, c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMonitoredProject(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("Failed to unmarshal provided resource in matcher.")
			return false
		}
		// URL Normalize both resources to compare only the short names.
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
			return true
		}
		if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		}
		if *nr.Name == *ncr.Name {
			c.Config.Logger.Info("Both Name fields equal - considering equal.")
			return true
		}
		c.Config.Logger.Infof("Attempting to match %d with %s.", projectNumber, *ncr.Name)
		iName, err := strconv.ParseInt(*ncr.Name, 10, 64)
		if err != nil {
			c.Config.Logger.Warningf("Could not convert %s to int: %v", *ncr.Name, err)
			return false
		}
		return projectNumber == iName
	}
}
