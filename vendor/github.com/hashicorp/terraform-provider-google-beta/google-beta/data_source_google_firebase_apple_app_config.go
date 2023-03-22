package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleFirebaseAppleAppConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleFirebaseAppleAppConfigRead,

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The id of the Firebase iOS App.`,
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The project id of the Firebase iOS App.`,
			},
			"config_filename": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The filename that the configuration artifact for the IosApp is typically saved as.`,
			},
			"config_file_contents": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The content of the XML configuration file as a base64-encoded string.`,
			},
		},
	}

}

func dataSourceGoogleFirebaseAppleAppConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("projects/%s/iosApps/%s", project, d.Get("app_id").(string))

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/iosApps/{{app_id}}/config")
	if err != nil {
		return err
	}

	res, err := SendRequest(config, "GET", project, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseAppleApp config %q", d.Id()))
	}

	if err = d.Set("config_filename", res["configFilename"]); err != nil {
		return err
	}

	if err = d.Set("config_file_contents", res["configFileContents"]); err != nil {
		return err
	}
	if err = d.Set("project", project); err != nil {
		return err
	}

	d.SetId(id)
	return nil
}
