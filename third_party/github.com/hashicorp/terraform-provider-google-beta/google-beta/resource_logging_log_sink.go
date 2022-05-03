package google

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLogSink() *schema.Resource {
	return combinedResource(
		map[string]*schema.Resource{
			projectType: resourceLoggingProjectSink(),
			folderType:  resourceLoggingFolderSink(),
			orgType:     resourceLoggingOrganizationSink(),
		}, loggingLogSinkImporter)
}

func loggingLogSinkImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	if strings.HasPrefix(id, "projects/") {
		return resourceLoggingSinkImportState("project")(d, meta)
	} else if strings.HasPrefix(id, "folders/") {
		return resourceLoggingSinkImportState("folder")(d, meta)
	} else if strings.HasPrefix(id, "organizations/") {
		return resourceLoggingSinkImportState("organization")(d, meta)
	}
	return kccImportIdImporter(d, meta)
}

func kccImportIdImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	// FROM: {{project?}}#{{folder?}}#{{org_id?}}#{{name}}
	// TO:
	//   - project: projects/{{project?}}/sinks/{{name}}
	//   - folder:  folders/{{folder?}}/sinks/{{name}}
	//   - org:     organizations/{{org_id?}}/sinks/{{name}}
	importID := d.Id()
	partitions := strings.Split(importID, "#")
	if len(partitions) != 4 {
		return nil, fmt.Errorf("expected 4 partitions in import ID, got %v", len(partitions))
	}
	name := partitions[3]
	for sinkType, parentVal := range map[string]string{
		projectType: partitions[0],
		folderType:  partitions[1],
		orgType:     partitions[2],
	} {
		if parentVal == "" {
			continue
		}
		if err := d.Set(fieldForParentType(sinkType), parentVal); err != nil {
			return nil, fmt.Errorf("error setting sink parent ID: %w", err)
		}
		var id string
		switch sinkType {
		case projectType:
			id = fmt.Sprintf("projects/%v/sinks/%v", parentVal, name)
		case folderType:
			id = fmt.Sprintf("folders/%v/sinks/%v", parentVal, name)
		case orgType:
			id = fmt.Sprintf("organizations/%v/sinks/%v", parentVal, name)
		default:
			return nil, fmt.Errorf("unknown sink type in import ID")
		}
		d.SetId(id)
		return []*schema.ResourceData{d}, nil
	}
	return nil, fmt.Errorf("no sink type specified")
}
