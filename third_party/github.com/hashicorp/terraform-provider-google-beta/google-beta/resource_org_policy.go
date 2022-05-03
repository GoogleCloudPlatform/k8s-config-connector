package google

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrgPolicy() *schema.Resource {
	return combinedResource(
		map[string]*schema.Resource{
			projectType: resourceGoogleProjectOrganizationPolicy(),
			folderType:  resourceGoogleFolderOrganizationPolicy(),
			orgType:     resourceGoogleOrganizationPolicy(),
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
				projectType: partitions[0],
				folderType:  partitions[1],
				orgType:     partitions[2],
			} {
				if parentVal == "" {
					continue
				}
				if err := d.Set(fieldForParentType(parentType), parentVal); err != nil {
					return nil, fmt.Errorf("error setting parent ID: %w", err)
				}
				var id string
				switch parentType {
				case projectType:
					id = fmt.Sprintf("projects/%v:%v", parentVal, constraint)
				case folderType, orgType:
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
