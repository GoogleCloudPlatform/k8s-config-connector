package resourcemanager

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

func ResourceOrgPolicy() *schema.Resource {
	return tpgresource.CombinedResource(
		map[string]*schema.Resource{
			tpgresource.ProjectType: ResourceGoogleProjectOrganizationPolicy(),
			tpgresource.FolderType:  ResourceGoogleFolderOrganizationPolicy(),
			tpgresource.OrgType:     ResourceGoogleOrganizationPolicy(),
		}, func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			// FROM: {{project?}}#{{folder?}}#{{org_id?}}#{{constraint}}
			// TO:
			//   - project: projects/{{project}}:{{constraint}}
			//   - folder:  {{folder}}/{{constraint}}
			//   - org:     {{org_id}}/{{constraint}}
			importID := d.Id()
			partitions := strings.Split(importID, "#")
			if len(partitions) != 4 {
				return nil, fmt.Errorf("expected 4 partitions in import ID, got %v", len(partitions))
			}
			constraint := partitions[3]
			for parentType, parentVal := range map[string]string{
				tpgresource.ProjectType: partitions[0],
				tpgresource.FolderType:  partitions[1],
				tpgresource.OrgType:     partitions[2],
			} {
				if parentVal == "" {
					continue
				}
				if err := d.Set(tpgresource.FieldForParentType(parentType), parentVal); err != nil {
					return nil, fmt.Errorf("error setting parent ID: %w", err)
				}
				var id string
				switch parentType {
				case tpgresource.ProjectType:
					id = fmt.Sprintf("projects/%v:%v", parentVal, constraint)
				case tpgresource.FolderType, tpgresource.OrgType:
					id = fmt.Sprintf("%v/%v", parentVal, constraint)
				default:
					return nil, fmt.Errorf("unknown policy type in import ID")
				}
				d.SetId(id)
				if err := d.Set("constraint", constraint); err != nil {
					return nil, fmt.Errorf("error setting constraint: %w", err)
				}
				return []*schema.ResourceData{d}, nil
			}
			return nil, fmt.Errorf("no policy type specified")
		})
}
