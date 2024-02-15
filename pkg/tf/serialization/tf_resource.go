// Copyright 2022 Google LLC
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

package serialization

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

// InstanceStateToHCL converts a terraform.InstanceState into the HCL that represents the maximally hydrated form.  The generated
// HCL will not include any references to other resources, since that information cannot be reconstructed without a full view
// of all resources present..
func InstanceStateToHCL(state *terraform.InstanceState, info *terraform.InstanceInfo, provider *tfschema.Provider) (string, error) {
	providerSchema := provider.ResourcesMap[info.Type].Schema
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("resource %q %q {\n", info.Type, info.Id))
	if err := resourceOrSubresourceHCL(&buff, providerSchema, state.Attributes, 1); err != nil {
		return "", err
	}
	buff.WriteString("}\n")
	// str = fmt.Sprintf("resource %q %q {\n%s\n}\n", info.Type, info.Id, strings.TrimSpace(str))

	hcl := hclwrite.Format(buff.Bytes())

	return string(hcl), nil
}

func resourceOrSubresourceHCL(w io.Writer, schema map[string]*tfschema.Schema, attributes map[string]string, indent int) error {
	var errs []error
	writeLine := func(msg string, args ...any) {
		line := strings.Repeat("  ", indent) + fmt.Sprintf(msg, args...) + "\n"
		if _, err := w.Write([]byte(line)); err != nil {
			errs = append(errs, err)
		}
	}

	keys := make([]string, 0, len(schema))
	for k := range schema {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, aname := range keys {
		aschema := schema[aname]
		if !aschema.Optional && !aschema.Required {
			continue
		}
		switch aschema.Type {
		case tfschema.TypeFloat:
			fallthrough
		case tfschema.TypeInt:
			fallthrough
		case tfschema.TypeBool:
			if attributes[aname] == "" {
				continue
			}
			writeLine("%s = %s", aname, attributes[aname])
		case tfschema.TypeString:
			if attributes[aname] == "" {
				continue
			}
			writeLine("%s = %s", aname, stringLitToHCL(attributes[aname]))
		// for each non-primitive, find if there are any values and set them.
		case tfschema.TypeMap:
			if val, ok := attributes[aname+".%"]; !ok || val == "0" {
				continue
			}
			writeLine("%s = {", aname)
			m, err := mapFromPrefix(attributes, aname)
			if err != nil {
				return err
			}
			var keys []string
			for k := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				v := m[k]
				writeLine("%s = %q", k, v)
			}
			writeLine("}")
		case tfschema.TypeSet:
			fallthrough
		case tfschema.TypeList:
			if val, ok := attributes[aname+".#"]; !ok || val == "0" {
				continue
			}
			if subtype, ok := aschema.Elem.(*tfschema.Schema); ok {
				writeLine("%s = [", aname)
				l, err := listFromPrefix(attributes, aname, aschema.Type)
				if err != nil {
					return err
				}
				for _, v := range l {
					switch subtype.Type {
					case tfschema.TypeFloat:
						fallthrough
					case tfschema.TypeInt:
						fallthrough
					case tfschema.TypeBool:
						writeLine("%s,", v)
					case tfschema.TypeString:
						writeLine("%q,", v)
					}
				}
				writeLine("]")
			} else if subtype, ok := aschema.Elem.(*tfschema.Resource); ok {
				cnt, err := strconv.Atoi(attributes[aname+".#"])
				if err != nil {
					return fmt.Errorf("could not parse count of %s, %w", aname, err)
				}
				for i := 0; i < cnt; i++ {
					subAttrs, err := mapsFromPrefix(attributes, fmt.Sprintf("%s.%d", aname, i))
					if err != nil {
						return fmt.Errorf("could not get subresource attributes for %s: %w", aname, err)
					}

					writeLine("%s {", aname)
					if err := resourceOrSubresourceHCL(w, subtype.Schema, subAttrs, indent+1); err != nil {
						return err
					}
					writeLine("}")
				}
			}
		}
	}
	return errors.Join(errs...)
}

func mapsFromPrefix(attributes map[string]string, prefix string) (map[string]string, error) {
	a := make(map[string]string)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") {
			a[strings.TrimPrefix(k, prefix+".")] = v
		}
	}
	return a, nil
}

func listFromPrefix(attributes map[string]string, prefix string, listType tfschema.ValueType) ([]string, error) {
	size, err := strconv.Atoi(attributes[prefix+".#"])
	if err != nil {
		return nil, fmt.Errorf("could not parse size of list %s: %w", prefix, err)
	}
	switch listType {
	case tfschema.TypeList:
		return listListFromPrefix(attributes, prefix, size)
	case tfschema.TypeSet:
		return setListFromPrefix(attributes, prefix, size), nil
	default:
		return nil, fmt.Errorf("unhandled list type: %v", listType)
	}
}

func listListFromPrefix(attributes map[string]string, prefix string, size int) ([]string, error) {
	out := make([]string, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".#" {
			kparts := strings.Split(k, ".")
			c, err := strconv.Atoi(kparts[len(kparts)-1])
			if err != nil {
				return nil, fmt.Errorf("could not parse index of %s: %w", k, err)
			}
			out[c] = v
		}
	}
	return out, nil
}

func setListFromPrefix(attributes map[string]string, prefix string, size int) []string {
	out := make([]string, 0, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".#" {
			out = append(out, v)
		}
	}
	// sort the list so that the ordering is deterministic for tests
	sort.Strings(out)
	return out
}

func mapFromPrefix(attributes map[string]string, prefix string) (map[string]string, error) {
	size, err := strconv.Atoi(attributes[prefix+".%"])
	if err != nil {
		return nil, fmt.Errorf("could not parse size of map %s: %w", prefix, err)
	}
	out := make(map[string]string, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".%" {
			kparts := strings.Split(k, ".")
			out[kparts[len(kparts)-1]] = v
		}
	}
	return out, nil
}

// stringLitToHCL converts a string literal value into a quoted, HCL-compatible
// format. It currently uses the `TokensForValue()` function from the
// `hclwrite` library to stay consistent with the HCL spec, particularly in
// cases where special characters need to be escaped.
func stringLitToHCL(val string) string {
	return string(hclwrite.TokensForValue(cty.StringVal(val)).Bytes())
}
