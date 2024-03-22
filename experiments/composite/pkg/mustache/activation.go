// Copyright 2024 Google LLC
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

package mustache

import "github.com/google/cel-go/interpreter"

const (
	ObjectVarName = "object"
	// OldObjectVarName                 = "oldObject"
	// ParamsVarName                    = "params"
	// RequestVarName                   = "request"
	// AuthorizerVarName                = "authorizer"
	// RequestResourceAuthorizerVarName = "authorizer.requestResource"
)

type Activation struct {
	Object  interface{}
	Version int
	Objects map[string]interface{}
	// object, oldObject, params, request, authorizer, requestResourceAuthorizer interface{}
}

// ResolveName returns a value from the activation by qualified name, or false if the name
// could not be found.
func (a *Activation) ResolveName(name string) (interface{}, bool) {
	switch name {
	case ObjectVarName:
		return a.Object, true
	case "version":
		return a.Version, true
	// case OldObjectVarName:
	// 	return a.oldObject, true
	// case ParamsVarName:
	// 	return a.params, true // params may be null
	// case RequestVarName:
	// 	return a.request, true
	// case AuthorizerVarName:
	// 	return a.authorizer, a.authorizer != nil
	// case RequestResourceAuthorizerVarName:
	// 	return a.requestResourceAuthorizer, a.requestResourceAuthorizer != nil
	default:
		v, found := a.Objects[name]
		if found {
			return v, true
		}
		return nil, false
	}
}

// Parent returns the parent of the current activation, may be nil.
// If non-nil, the parent will be searched during resolve calls.
func (a *Activation) Parent() interpreter.Activation {
	return nil
}
