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

package mockcloudtalentsolution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

type Company struct {
	Name                                 string   `json:"name,omitempty"`
	DisplayName                          string   `json:"displayName,omitempty"`
	ExternalID                           string   `json:"externalId,omitempty"`
	Size                                 any      `json:"size,omitempty"`
	HeadquartersAddress                  string   `json:"headquartersAddress,omitempty"`
	HiringAgency                         bool     `json:"hiringAgency,omitempty"`
	EeoText                              string   `json:"eeoText,omitempty"`
	WebsiteURI                           string   `json:"websiteUri,omitempty"`
	CareerSiteURI                        string   `json:"careerSiteUri,omitempty"`
	ImageURI                             string   `json:"imageUri,omitempty"`
	KeywordSearchableJobCustomAttributes []string `json:"keywordSearchableJobCustomAttributes,omitempty"`
	DerivedInfo                          any      `json:"derivedInfo,omitempty"`
	Suspended                            bool     `json:"suspended,omitempty"`
}

type cloudTalentSolutionMock struct {
	mu        sync.Mutex
	companies map[string]*Company
}

var mockStore = &cloudTalentSolutionMock{
	companies: make(map[string]*Company),
}

func (s *MockService) registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/v4/projects/", mockStore.handleTalentRequest)
}

func (store *cloudTalentSolutionMock) handleTalentRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	// Regular expression for /v4/projects/{project}/tenants
	tenantsRegex := regexp.MustCompile(`^/v4/projects/([^/]+)/tenants$`)
	// Regular expression for /v4/projects/{project}/tenants/{tenant}/companies
	companiesRegex := regexp.MustCompile(`^/v4/projects/([^/]+)/tenants/([^/]+)/companies$`)
	// Regular expression for /v4/projects/{project}/tenants/{tenant}/companies/{company}
	companyRegex := regexp.MustCompile(`^/v4/projects/([^/]+)/tenants/([^/]+)/companies/([^/]+)$`)

	w.Header().Set("Content-Type", "application/json")

	if matches := tenantsRegex.FindStringSubmatch(path); len(matches) > 0 && method == "GET" {
		project := matches[1]
		tenantID := "mock-default-tenant-id"
		response := map[string]any{
			"tenants": []map[string]any{
				{
					"name":       fmt.Sprintf("projects/%s/tenants/%s", project, tenantID),
					"externalId": "default-tenant",
				},
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if matches := companiesRegex.FindStringSubmatch(path); len(matches) > 0 && method == "POST" {
		project := matches[1]
		tenant := matches[2]

		company := &Company{}
		if err := json.NewDecoder(r.Body).Decode(company); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		store.mu.Lock()
		defer store.mu.Unlock()

		// Use the name specified by the controller (derived from the KRM resource ID)
		fullName := company.Name
		if fullName == "" {
			companyID := fmt.Sprintf("mock-company-id-%s", company.ExternalID)
			fullName = fmt.Sprintf("projects/%s/tenants/%s/companies/%s", project, tenant, companyID)
			company.Name = fullName
		}

		company.DerivedInfo = map[string]any{}
		company.Suspended = false

		store.companies[fullName] = company

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
		return
	}

	if matches := companyRegex.FindStringSubmatch(path); len(matches) > 0 {
		project := matches[1]
		tenant := matches[2]
		companyID := matches[3]
		fullName := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", project, tenant, companyID)

		store.mu.Lock()
		defer store.mu.Unlock()

		company, exists := store.companies[fullName]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": map[string]any{
					"code":    404,
					"message": fmt.Sprintf("Company %s not found", fullName),
					"status":  "NOT_FOUND",
				},
			})
			return
		}

		switch method {
		case "GET":
			json.NewEncoder(w).Encode(company)
		case "PATCH":
			patchCompany := &Company{}
			if err := json.NewDecoder(r.Body).Decode(patchCompany); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// Update mutable fields
			if patchCompany.DisplayName != "" {
				company.DisplayName = patchCompany.DisplayName
			}
			if patchCompany.Size != nil {
				company.Size = patchCompany.Size
			}
			if patchCompany.HeadquartersAddress != "" {
				company.HeadquartersAddress = patchCompany.HeadquartersAddress
			}
			company.HiringAgency = patchCompany.HiringAgency
			if patchCompany.EeoText != "" {
				company.EeoText = patchCompany.EeoText
			}
			if patchCompany.WebsiteURI != "" {
				company.WebsiteURI = patchCompany.WebsiteURI
			}
			if patchCompany.CareerSiteURI != "" {
				company.CareerSiteURI = patchCompany.CareerSiteURI
			}
			if patchCompany.ImageURI != "" {
				company.ImageURI = patchCompany.ImageURI
			}
			if patchCompany.KeywordSearchableJobCustomAttributes != nil {
				company.KeywordSearchableJobCustomAttributes = patchCompany.KeywordSearchableJobCustomAttributes
			}
			json.NewEncoder(w).Encode(company)
		case "DELETE":
			delete(store.companies, fullName)
			json.NewEncoder(w).Encode(map[string]any{})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
