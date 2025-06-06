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

package codebot

import (
	"fmt"
	"regexp"
	"strings"
)

// Filter interface for tools to be included or excluded based on regex patterns.
type Filter interface {
	// Exclude methods determine if a tool should be excluded based on regex patterns.
	Exclude(tool Tool) bool
	// Include methods determine if a tool should be included based on regex patterns.
	Include(tool Tool) bool
}

var _ Filter = &CustomRegexFilter{}

func NewCustomRegexFilter(includeRegex, excludeRegex string) (*CustomRegexFilter, error) {
	if includeRegex == "" && excludeRegex == "" {
		return &CustomRegexFilter{}, nil
	}
	var incl, excl *regexp.Regexp
	var err error
	if includeRegex != "" {
		incl, err = regexp.Compile("(" + strings.TrimSpace(includeRegex) + ")")
		if err != nil {
			return nil, fmt.Errorf("invalid include regex: %s: %w", includeRegex, err)
		}
	}
	if excludeRegex != "" {
		excl, err = regexp.Compile("(" + strings.TrimSpace(excludeRegex) + ")")
		if err != nil {
			return nil, fmt.Errorf("invalid exclude regex: %s: %w", excludeRegex, err)
		}
	}

	return &CustomRegexFilter{
		includeRegex: incl,
		excludeRegex: excl,
	}, nil
}

// CustomRegexFilter implements the Filter interface.
// It allows for inclusion and exclusion of tools based on regex patterns.
// Note: If excludeRegex is not give, it will not exclude any tools.
type CustomRegexFilter struct {
	includeRegex *regexp.Regexp
	excludeRegex *regexp.Regexp
}

func (a *CustomRegexFilter) Exclude(tool Tool) bool {
	if a.excludeRegex == nil {
		return false
	}

	toolName := tool.BuildFunctionDefinition().Name
	return len(a.excludeRegex.FindStringSubmatch(toolName)) > 0
}

func (a *CustomRegexFilter) Include(tool Tool) bool {
	if a.includeRegex == nil {
		return true
	}
	toolName := tool.BuildFunctionDefinition().Name
	return len(a.includeRegex.FindStringSubmatch(toolName)) > 0
}
