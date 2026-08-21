package resourcemanager

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

func ResourceIamCustomRole() *schema.Resource {
	return tpgresource.CombinedResource(
		map[string]*schema.Resource{
			tpgresource.ProjectType: ResourceGoogleProjectIamCustomRole(),
			tpgresource.OrgType:     ResourceGoogleOrganizationIamCustomRole(),
		}, func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			// FROM: {{project?}}#{{org_id?}}#{{role_id}}
			// TO:
			//   - project: projects/{{project?}}/roles/{{role_id}}
			//   - org:     organizations/{{org_id?}}/roles/{{role_id}}
			importID := d.Id()
			partitions := strings.Split(importID, "#")
			if len(partitions) != 3 {
				return nil, fmt.Errorf("expected 3 partitions in import ID, got %v", len(partitions))
			}
			roleID := partitions[2]
			for roleType, parentVal := range map[string]string{
				tpgresource.ProjectType: partitions[0],
				tpgresource.OrgType:     partitions[1],
			} {
				if parentVal == "" {
					continue
				}
				if err := d.Set(tpgresource.FieldForParentType(roleType), parentVal); err != nil {
					return nil, fmt.Errorf("error setting role parent ID: %w", err)
				}
				var id string
				switch roleType {
				case tpgresource.ProjectType:
					id = fmt.Sprintf("projects/%v/roles/%v", parentVal, roleID)
				case tpgresource.OrgType:
					id = fmt.Sprintf("organizations/%v/roles/%v", parentVal, roleID)
				default:
					return nil, fmt.Errorf("unknown role type in import ID")
				}
				d.SetId(id)
				return []*schema.ResourceData{d}, nil
			}
			return nil, fmt.Errorf("no role type specified")
		})
}
