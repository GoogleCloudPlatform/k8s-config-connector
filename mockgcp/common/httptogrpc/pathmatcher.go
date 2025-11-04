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

package httptogrpc

import (
	"fmt"
	"strings"

	"k8s.io/klog/v2"
)

// pathMatcher matches HTTP paths against gRPC HTTP annotations.
// For example, it can match a path like: v1/{name=projects/*/locations/*/instances/*}
type pathMatcher struct {
	// httpPath is the original HTTP path pattern (for debugging)
	httpPath string

	// components are the component matchers for this path expression
	components []componentMatcher
}

// newPathMatcher creates a new pathMatcher for the given HTTP path pattern.
func newPathMatcher(httpPath string) (*pathMatcher, error) {
	matcher := &pathMatcher{httpPath: httpPath}

	var components []componentMatcher
	httpPath = strings.TrimPrefix(httpPath, "/")
	for httpPath != "" {
		if httpPath[0] == '{' {
			closeQuote := strings.Index(httpPath, "}")
			if closeQuote == -1 {
				return nil, fmt.Errorf("invalid httpPath %q", httpPath)
			}
			token := httpPath[1:closeQuote]
			components = append(components, newMatchNamedComponent(token))
			httpPath = httpPath[closeQuote+1:]
		} else {
			nextSlash := strings.Index(httpPath, "/")
			if nextSlash == -1 {
				nextSlash = len(httpPath)
			}
			token := httpPath[:nextSlash]
			components = append(components, newMatchLiteralComponent(token))
			httpPath = httpPath[nextSlash:]
		}
		httpPath = strings.TrimPrefix(httpPath, "/")
	}

	variableLengthCount := 0
	for _, component := range components {
		if _, isFixedLength := component.FixedLength(); !isFixedLength {
			variableLengthCount++
		}
	}
	if variableLengthCount > 1 {
		return nil, fmt.Errorf("cannot handle multiple variable-length components in httpPath %q", httpPath)
	}

	matcher.components = components

	return matcher, nil
}

// Match attempts to match the given tokens against the path pattern.
// If matched, it returns the mapping of variable names to values, and true.
// If not matched, it returns nil and false.
func (m *pathMatcher) Match(tokens []string) (map[string]string, bool) {
	values := make(map[string]string)

	// First match from fixed-length components from the left
	leftComponentPos := 0
	leftTokenPos := 0
	for _, submatcher := range m.components {
		fixedLength, isFixed := submatcher.FixedLength()
		if !isFixed {
			// Stop at first variable-length component
			break
		}
		if leftTokenPos+fixedLength > len(tokens) {
			return nil, false
		}

		if !submatcher.Match(tokens[leftTokenPos:leftTokenPos+fixedLength], values) {
			return nil, false
		}
		leftTokenPos += fixedLength
		leftComponentPos++
	}

	// Then match from fixed-length components from the right
	rightComponentPos := len(m.components) - 1
	rightTokenPos := len(tokens)
	for rightComponentPos > leftComponentPos {
		submatcher := m.components[rightComponentPos]
		fixedLength, isFixed := submatcher.FixedLength()
		if !isFixed {
			// Stop at first variable-length component
			break
		}
		if rightTokenPos-fixedLength < leftTokenPos {
			return nil, false
		}
		if !submatcher.Match(tokens[rightTokenPos-fixedLength:rightTokenPos], values) {
			return nil, false
		}
		rightTokenPos -= fixedLength
		rightComponentPos--
	}

	// Then match a wildcard component if we have one
	if leftComponentPos-1 == rightComponentPos {
		// All components matched, but did we match all tokens?
		if leftTokenPos != rightTokenPos {
			return nil, false
		}
	} else if leftComponentPos == rightComponentPos {
		// One variable-length component to match
		wildcardMatcher := m.components[leftComponentPos]
		if !wildcardMatcher.Match(tokens[leftTokenPos:rightTokenPos], values) {
			return nil, false
		}
	} else {
		// Multiple variable-length components to match; not (yet?) handled
		klog.Fatalf("cannot handle multiple variable-length components for path: %v", m.httpPath)
		return nil, false
	}

	return values, true
}

// componentMatcher matches a component of a path expression
type componentMatcher interface {
	// Match attempts to match the given tokens
	// If matched, it populates matches with any named captures
	// It returns true if matched, false otherwise
	Match(tokens []string, matches map[string]string) bool

	// FixedLength returns the fixed length of this component, if applicable
	// If the component is variable length, the second return value is false
	// If the component is fixed length, the second return value is true
	FixedLength() (int, bool)
}

// matchLiteralComponent matches a single literal path component
type matchLiteralComponent struct {
	literal string
}

// newMatchLiteralComponent creates a new matchLiteralComponent for the given literal string.
func newMatchLiteralComponent(literal string) *matchLiteralComponent {
	return &matchLiteralComponent{literal: literal}
}

// Match attempts to match the given tokens against the literal.
func (m *matchLiteralComponent) Match(tokens []string, matches map[string]string) bool {
	if len(tokens) == 0 {
		return false
	}
	if tokens[0] == m.literal {
		return true
	}
	return false
}

// FixedLength returns the fixed length of this component.
func (m *matchLiteralComponent) FixedLength() (int, bool) {
	return 1, true
}

// matchNamedComponent matches a named path component, e.g. {name=projects/*/locations/*/instances/**}
type matchNamedComponent struct {
	// key is the name of the variable to capture
	key string

	// patternTokens are the tokens in the pattern, e.g. projects/*/locations/*/foo/* => ["projects", "*", "locations", "*", "foo", "*"]
	patternTokens []string
}

// newMatchNamedComponent creates a new matchNamedComponent for the given wildcard string.
func newMatchNamedComponent(wildcard string) *matchNamedComponent {
	eqTokens := strings.Split(wildcard, "=")
	if len(eqTokens) == 2 {
		// e.g. {name=projects/*/locations/*/foo/*}
		// e.g. {name=projects/*/databases/*/documents/*/**}
		key := eqTokens[0]
		pattern := eqTokens[1]

		patternTokens := strings.Split(pattern, "/")

		return &matchNamedComponent{
			key:           key,
			patternTokens: patternTokens,
		}
	} else if len(eqTokens) == 1 {
		// e.g. {name}
		key := eqTokens[0]

		return &matchNamedComponent{
			key:           key,
			patternTokens: []string{"*"},
		}
	} else {
		klog.Fatalf("unhandled wildcard token: %q", wildcard)
		return nil
	}
}

// FixedLength returns the fixed length of this component.
func (m *matchNamedComponent) FixedLength() (int, bool) {
	for _, token := range m.patternTokens {
		if token == "**" {
			return 0, false
		}
	}
	return len(m.patternTokens), true
}

// Match attempts to match the given tokens against the pattern.
func (m *matchNamedComponent) Match(tokens []string, matches map[string]string) bool {
	tokenPos := 0
	for i, patternToken := range m.patternTokens {
		if tokenPos >= len(tokens) {
			if patternToken == "**" {
				// Special case: ** can match zero tokens (i.e., empty)
				continue
			} else {
				return false
			}
		}

		switch patternToken {

		case "*":
			// Single wildcard matches any single token
			tokenPos++
			continue

		case "**":
			// Double wildcard matches any number of tokens until we hit the terminal token
			// The terminal token is populated into next
			// Make sure this is the last token
			if i != len(m.patternTokens)-1 {
				klog.Fatalf("invalid ** wildcard position in %v", m.patternTokens)
				return false
			}

			// We match all remaining tokens

		default:
			// Literal match
			if tokens[tokenPos] != patternToken {
				return false
			}
			tokenPos++
		}
	}

	matches[m.key] = strings.Join(tokens, "/")
	return true
}
