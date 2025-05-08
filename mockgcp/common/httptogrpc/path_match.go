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

type pathMatcher struct {
	components []componentMatcher
}

func newPathMatcher(httpPath string) (*pathMatcher, error) {
	var components []componentMatcher
	httpPath = strings.TrimPrefix(httpPath, "/")
	for httpPath != "" {
		if httpPath[0] == '{' {
			closeQuote := strings.Index(httpPath, "}")
			if closeQuote == -1 {
				return nil, fmt.Errorf("invalid httpPath %q", httpPath)
			}
			token := httpPath[1:closeQuote]
			components = append(components, newMatchWildcard(token))
			httpPath = httpPath[closeQuote+1:]
		} else {
			nextSlash := strings.Index(httpPath, "/")
			if nextSlash == -1 {
				nextSlash = len(httpPath)
			}
			token := httpPath[:nextSlash]
			components = append(components, newMatchLiteral(token))
			httpPath = httpPath[nextSlash:]
		}
		httpPath = strings.TrimPrefix(httpPath, "/")
	}

	// Configure termination criteria for ** wildcards
	for i, component := range components {
		switch component := component.(type) {
		case *matchWildcard:
			if len(component.patternTokens) == 1 && component.patternTokens[0] == "**" {
				if len(components) == i+1 {
					// Match end of path
					component.next = newMatchEnd()
					continue
				}
				nextComponent := components[i+1]
				if literal, ok := nextComponent.(*matchLiteral); ok {
					component.next = literal
				} else {
					return nil, fmt.Errorf("cannot compute termination token for ** wildcard")
				}
			}
		}
	}
	return &pathMatcher{components: components}, nil
}

func (m *pathMatcher) Match(tokens []string) (map[string]string, bool) {
	values := make(map[string]string)
	for _, submatcher := range m.components {
		ok, newTokens := submatcher.Match(tokens, values)
		if !ok {
			return nil, false
		}
		tokens = newTokens
	}
	if len(tokens) != 0 {
		return nil, false
	}
	return values, true
}

type componentMatcher interface {
	Match(tokens []string, matches map[string]string) (bool, []string)
}

type matchLiteral struct {
	literal string
}

func newMatchLiteral(literal string) *matchLiteral {
	return &matchLiteral{literal: literal}
}

func (m *matchLiteral) Match(tokens []string, matches map[string]string) (bool, []string) {
	if len(tokens) == 0 {
		return false, nil
	}
	if tokens[0] == m.literal {
		return true, tokens[1:]
	}
	return false, nil
}

type matchEnd struct {
}

func newMatchEnd() *matchEnd {
	return &matchEnd{}
}

func (m *matchEnd) Match(tokens []string, matches map[string]string) (bool, []string) {
	if len(tokens) == 0 {
		return true, nil
	}
	return false, nil
}

type matchWildcard struct {
	key           string
	patternTokens []string
	// next is used when matching ** wildcards, it terminates the match
	next componentMatcher
}

func newMatchWildcard(wildcard string) *matchWildcard {
	eqTokens := strings.Split(wildcard, "=")
	if len(eqTokens) == 2 {
		// e.g. {name=projects/*/locations/*/foo/*}
		key := eqTokens[0]
		pattern := eqTokens[1]

		patternTokens := strings.Split(pattern, "/")

		return &matchWildcard{
			key:           key,
			patternTokens: patternTokens,
		}
	} else if len(eqTokens) == 1 {
		// e.g. {name}
		key := eqTokens[0]

		return &matchWildcard{
			key:           key,
			patternTokens: []string{"*"},
		}
	} else {
		klog.Fatalf("unhandled wildcard token: %q", wildcard)
		return nil
	}
}

func (m *matchWildcard) Match(tokens []string, matches map[string]string) (bool, []string) {
	if len(m.patternTokens) == 1 && m.patternTokens[0] == "**" {
		// klog.Infof("WILDCARD match for %v", tokens)

		// Special case: match until we hit the terminal token
		var matched []string
		for {
			if ok, _ := m.next.Match(tokens, nil); ok {
				matches[m.key] = strings.Join(matched, "/")
				// klog.Infof("WILDCARD match for %v; matches %v", tokens, matches)
				return true, tokens
			}

			if len(tokens) == 0 {
				return false, nil
			}

			matched = append(matched, tokens[0])
			tokens = tokens[1:]
		}
	} else {
		if len(tokens) < len(m.patternTokens) {
			return false, nil
		}

		for i, patternToken := range m.patternTokens {
			if patternToken == "*" {
				continue
			}
			if tokens[i] != patternToken {
				return false, nil
			}
		}
	}

	matches[m.key] = strings.Join(tokens[:len(m.patternTokens)], "/")
	return true, tokens[len(m.patternTokens):]
}
