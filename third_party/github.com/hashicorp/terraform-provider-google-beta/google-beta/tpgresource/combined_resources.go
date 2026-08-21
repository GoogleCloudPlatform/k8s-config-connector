package tpgresource

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ProjectType = "project"
	FolderType  = "folder"
	OrgType     = "organization"
)

var parentTypes = []string{ProjectType, FolderType, OrgType}

func CombinedResource(resources map[string]*schema.Resource, importStateFunc schema.StateFunc) *schema.Resource {
	return &schema.Resource{
		Schema: combinedSchema(resources),
		Create: func(d *schema.ResourceData, meta interface{}) error {
			if err := validateOnlyOneParent(d); err != nil {
				return err
			}
			parentType, err := getParentType(d)
			if err != nil {
				return err
			}
			if err := validateSubschema(d, parentType, resources); err != nil {
				return err
			}
			return resources[parentType].Create(d, meta)
		},
		Read: func(d *schema.ResourceData, meta interface{}) error {
			if err := validateOnlyOneParent(d); err != nil {
				return err
			}
			parentType, err := getParentType(d)
			if err != nil {
				return err
			}
			return resources[parentType].Read(d, meta)
		},
		Update: func(d *schema.ResourceData, meta interface{}) error {
			if err := validateOnlyOneParent(d); err != nil {
				return err
			}
			parentType, err := getParentType(d)
			if err != nil {
				return err
			}
			if err := validateSubschema(d, parentType, resources); err != nil {
				return err
			}
			return resources[parentType].Update(d, meta)
		},
		Delete: func(d *schema.ResourceData, meta interface{}) error {
			if err := validateOnlyOneParent(d); err != nil {
				return err
			}
			parentType, err := getParentType(d)
			if err != nil {
				return err
			}
			if err := validateSubschema(d, parentType, resources); err != nil {
				return err
			}
			return resources[parentType].Delete(d, meta)
		},
		Importer: &schema.ResourceImporter{
			State: importStateFunc,
		},
	}
}

func validateOnlyOneParent(d *schema.ResourceData) error {
	found := false
	for _, parentType := range parentTypes {
		if _, ok := d.GetOk(FieldForParentType(parentType)); ok {
			if found {
				return fmt.Errorf("only one of project, folder, or organization may be specified")
			}
			found = true
		}
	}
	return nil
}

func validateSubschema(d *schema.ResourceData, parentType string, resources map[string]*schema.Resource) error {
	// NOTE: This only checks at the root level.
	//
	// Ensure that all the fields specified on the resource are for the
	// appropriate subschema. As the schema.ResourceData type does not let us
	// list all set fields, must iterate the other schemas and check for
	// presence.
	for _, otherParent := range getOtherParentTypes(parentType) {
		if _, ok := resources[otherParent]; !ok {
			continue
		}
		for field, _ := range resources[otherParent].Schema {
			if _, ok := d.GetOk(field); ok {
				if _, ok := resources[parentType].Schema[field]; !ok {
					return fmt.Errorf("field %v cannot be set on resources with a parent of type %v", field, parentType)
				}
			}
		}
	}
	// Validates that required fields for this particular parent type are set.
	for field, s := range resources[parentType].Schema {
		if s.Required {
			if _, ok := d.GetOk(field); !ok {
				return fmt.Errorf("field %v is required for resources with a parent of type %v", field, parentType)
			}
		}
	}
	return nil
}

func FieldForParentType(parentType string) string {
	switch parentType {
	case ProjectType:
		return "project"
	case FolderType:
		return "folder"
	case OrgType:
		return "org_id"
	default:
		panic(fmt.Sprintf("unknown parent type %v", parentType))
	}
}

func combinedSchema(resources map[string]*schema.Resource) map[string]*schema.Schema {
	// combines the given resource schemas, setting required fields to optional
	// unless they are required on all the schemas.
	//
	// NOTE: Only sets fields optional at the root level.
	combinedSchema := make(map[string]*schema.Schema)
	for _, parentType := range parentTypes {
		if _, ok := resources[parentType]; !ok {
			continue
		}
		for field, s := range resources[parentType].Schema {
			combinedSchema[field] = s
			if s.Required {
				for _, otherParent := range getOtherParentTypes(parentType) {
					if _, ok := resources[otherParent]; !ok {
						continue
					}
					if !isRequiredForParent(field, otherParent, resources) {
						s.Required = false
						s.Optional = true
					}
				}
			}
		}
	}
	return combinedSchema
}

func getOtherParentTypes(parentType string) []string {
	// Returns a list of parent types that do NOT match the provided type.
	res := make([]string, 0)
	for _, parent := range parentTypes {
		if parent != parentType {
			res = append(res, parent)
		}
	}
	return res
}

func isRequiredForParent(field, parentType string, resources map[string]*schema.Resource) bool {
	s, ok := resources[parentType].Schema[field]
	return ok && s.Required
}

func getParentType(d *schema.ResourceData) (string, error) {
	for _, parentType := range parentTypes {
		if field, ok := d.GetOk(FieldForParentType(parentType)); ok && field.(string) != "" {
			return parentType, nil
		}
	}
	return "", fmt.Errorf("unknown parent type")
}
