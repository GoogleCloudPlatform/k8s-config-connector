package openapi

import (
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Callback Object
type Callback map[string]*PathItem

// Validate the values of Callback object.
func (callback Callback) Validate() error {
	for key, pathItem := range callback {
		if !matchRuntimeExpression(key) {
			return ErrRuntimeExprFormat
		}
		if err := pathItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

const (
	rfc5234Alpha     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rfc5234Digit     = "0123456789"
	rfc7230TChar     = "!#$%&'*+-.^_`|~" + rfc5234Digit + rfc5234Alpha
	rfc7159Unescaped = "\x20\x21\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30"
)

func matchRuntimeExpression(key string) bool {
	if key == "" {
		return false
	}
	for {
		ob := strings.IndexRune(key, '{')
		if ob == -1 {
			break
		}
		cb := strings.IndexRune(key, '}')
		if cb == -1 {
			break
		}
		expr := key[ob+1 : cb]
		key = key[cb+1:]
		if !strings.HasPrefix(expr, "$") {
			return false
		}
		if expr == "$url" || expr == "$method" || expr == "$statusCode" {
			continue
		}
		var source string
		if !strings.HasPrefix(expr, "$request.") {
			if !strings.HasPrefix(expr, "$response.") {
				return false
			}
			source = strings.TrimPrefix(expr, "$response.")
		} else {
			source = strings.TrimPrefix(expr, "$request.")
		}
		if len(source) == 0 {
			return false
		}
		var name string
		switch {
		case strings.HasPrefix(source, "header."):
			token := strings.TrimPrefix(source, "header.")
			if len(token) == 0 {
				return false
			}
			return len(strings.Trim(token, rfc7230TChar)) == 0
		case strings.HasPrefix(source, "body"):
			if strings.Contains(source, "#") {
				split := strings.Split(source, "#")
				if split[0] != "body" {
					return false
				}
				fragment := split[1]
				if !strings.HasPrefix(fragment, "/") {
					return false
				}
				continue
			}
		case strings.HasPrefix(source, "query."):
			name = strings.TrimPrefix(source, "query.")
		case strings.HasPrefix(source, "path."):
			name = strings.TrimPrefix(source, "path.")
		}
		if len(name) == 0 {
			return false
		}
	}
	return true
}
