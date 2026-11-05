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

package mocklogging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type logScopeEntity struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	CreateTime    string   `json:"createTime,omitempty"`
	UpdateTime    string   `json:"updateTime,omitempty"`
	ResourceNames []string `json:"resourceNames,omitempty"`
}

var logScopeURLRegex = regexp.MustCompile(`^/v2/([^/]+/[^/]+)/locations/([^/]+)/logScopes(?:/([^/?]+))?$`)

func (s *MockService) serveLogScope(w http.ResponseWriter, r *http.Request) bool {
	matches := logScopeURLRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		return false
	}

	parent := matches[1]
	location := matches[2]
	logScopeID := matches[3]

	if logScopeID == "" {
		switch r.Method {
		case http.MethodPost:
			s.createLogScope(w, r, parent, location)
			return true
		case http.MethodGet:
			s.listLogScopes(w, r, parent, location)
			return true
		default:
			return false
		}
	}

	fqn := fmt.Sprintf("%s/locations/%s/logScopes/%s", parent, location, logScopeID)

	switch r.Method {
	case http.MethodGet:
		s.getLogScope(w, r, fqn)
		return true
	case http.MethodPatch:
		s.patchLogScope(w, r, fqn)
		return true
	case http.MethodDelete:
		s.deleteLogScope(w, r, fqn)
		return true
	default:
		return false
	}
}

func (s *MockService) createLogScope(w http.ResponseWriter, r *http.Request, parent, location string) {
	logScopeID := r.URL.Query().Get("logScopeId")

	var req logScopeEntity
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeGCPError(w, http.StatusBadRequest, "INVALID_ARGUMENT", err.Error())
		return
	}

	if logScopeID == "" && req.Name != "" {
		parts := strings.Split(req.Name, "/")
		logScopeID = parts[len(parts)-1]
	}

	if logScopeID == "" {
		writeGCPError(w, http.StatusBadRequest, "INVALID_ARGUMENT", "logScopeId is required")
		return
	}

	fqn := fmt.Sprintf("%s/locations/%s/logScopes/%s", parent, location, logScopeID)

	s.logScopesMutex.Lock()
	defer s.logScopesMutex.Unlock()

	if _, found := s.logScopes[fqn]; found {
		writeGCPError(w, http.StatusConflict, "ALREADY_EXISTS", fmt.Sprintf("%s already exists", fqn))
		return
	}

	now := time.Now().UTC().Format(time.RFC3339Nano)
	req.Name = fqn
	req.CreateTime = now
	req.UpdateTime = now

	if req.ResourceNames != nil {
		rn := make([]string, len(req.ResourceNames))
		copy(rn, req.ResourceNames)
		req.ResourceNames = rn
	}

	s.logScopes[fqn] = &req
	writeJSONResponse(w, http.StatusOK, &req)
}

func (s *MockService) listLogScopes(w http.ResponseWriter, r *http.Request, parent, location string) {
	prefix := fmt.Sprintf("%s/locations/%s/logScopes/", parent, location)

	s.logScopesMutex.Lock()
	defer s.logScopesMutex.Unlock()

	var list []*logScopeEntity
	for name, item := range s.logScopes {
		if strings.HasPrefix(name, prefix) {
			list = append(list, item)
		}
	}

	writeJSONResponse(w, http.StatusOK, map[string]any{"logScopes": list})
}

func (s *MockService) getLogScope(w http.ResponseWriter, r *http.Request, fqn string) {
	s.logScopesMutex.Lock()
	defer s.logScopesMutex.Unlock()

	item, found := s.logScopes[fqn]
	if !found {
		writeGCPError(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("%s does not exist", fqn))
		return
	}

	writeJSONResponse(w, http.StatusOK, item)
}

func (s *MockService) patchLogScope(w http.ResponseWriter, r *http.Request, fqn string) {
	var req logScopeEntity
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeGCPError(w, http.StatusBadRequest, "INVALID_ARGUMENT", err.Error())
		return
	}

	s.logScopesMutex.Lock()
	defer s.logScopesMutex.Unlock()

	item, found := s.logScopes[fqn]
	if !found {
		writeGCPError(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("%s does not exist", fqn))
		return
	}

	updateMask := r.URL.Query().Get("updateMask")
	paths := strings.Split(updateMask, ",")
	for _, p := range paths {
		switch strings.TrimSpace(p) {
		case "description":
			item.Description = req.Description
		case "resource_names", "resourceNames":
			if req.ResourceNames != nil {
				rn := make([]string, len(req.ResourceNames))
				copy(rn, req.ResourceNames)
				item.ResourceNames = rn
			} else {
				item.ResourceNames = nil
			}
		}
	}

	now := time.Now().UTC().Format(time.RFC3339Nano)
	item.UpdateTime = now

	writeJSONResponse(w, http.StatusOK, item)
}

func (s *MockService) deleteLogScope(w http.ResponseWriter, r *http.Request, fqn string) {
	s.logScopesMutex.Lock()
	defer s.logScopesMutex.Unlock()

	_, found := s.logScopes[fqn]
	if !found {
		writeGCPError(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("%s does not exist", fqn))
		return
	}

	delete(s.logScopes, fqn)
	writeJSONResponse(w, http.StatusOK, map[string]any{})
}

func writeGCPHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Server", "ESF")
	w.Header()["Vary"] = []string{"Origin", "X-Origin", "Referer"}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Xss-Protection", "0")
}

func writeJSONResponse(w http.ResponseWriter, code int, resp any) {
	writeGCPHeaders(w)
	w.WriteHeader(code)
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func writeGCPError(w http.ResponseWriter, code int, statusText, message string) {
	writeGCPHeaders(w)
	w.WriteHeader(code)
	resp := map[string]any{
		"error": map[string]any{
			"code":    code,
			"message": message,
			"status":  statusText,
		},
	}
	b, _ := json.Marshal(resp)
	w.Write(b)
}
