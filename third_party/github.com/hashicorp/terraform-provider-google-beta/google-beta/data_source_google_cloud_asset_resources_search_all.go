package google

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleCloudAssetResourcesSearchAll() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGoogleCloudAssetResourcesSearchAllRead,
		Schema: map[string]*schema.Schema{
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"additional_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"network_tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func datasourceGoogleCloudAssetResourcesSearchAllRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	params := make(map[string]string)
	results := make([]map[string]interface{}, 0)

	scope := d.Get("scope").(string)
	query := d.Get("query").(string)
	assetTypes := d.Get("asset_types").([]interface{})

	url := fmt.Sprintf("https://cloudasset.googleapis.com/v1p1beta1/%s/resources:searchAll", scope)
	params["query"] = query

	url, err = addArrayQueryParam(url, "asset_types", assetTypes)
	if err != nil {
		return fmt.Errorf("Error setting asset_types: %s", err)
	}

	for {
		url, err := transport_tpg.AddQueryParams(url, params)
		if err != nil {
			return err
		}

		res, err := transport_tpg.SendRequest(config, "GET", "", url, userAgent, nil)
		if err != nil {
			return fmt.Errorf("Error searching resources: %s", err)
		}

		pageResults := flattenDatasourceGoogleCloudAssetResourcesList(res["results"])
		results = append(results, pageResults...)

		pToken, ok := res["nextPageToken"]
		if ok && pToken != nil && pToken.(string) != "" {
			params["pageToken"] = pToken.(string)
		} else {
			break
		}
	}

	if err := d.Set("results", results); err != nil {
		return fmt.Errorf("Error searching resources: %s", err)
	}

	if err := d.Set("query", query); err != nil {
		return fmt.Errorf("Error setting query: %s", err)
	}

	if err := d.Set("asset_types", assetTypes); err != nil {
		return fmt.Errorf("Error setting asset_types: %s", err)
	}

	d.SetId(scope)

	return nil
}

func flattenDatasourceGoogleCloudAssetResourcesList(v interface{}) []map[string]interface{} {
	if v == nil {
		return make([]map[string]interface{}, 0)
	}

	ls := v.([]interface{})
	results := make([]map[string]interface{}, 0, len(ls))
	for _, raw := range ls {
		p := raw.(map[string]interface{})

		var mName, mAssetType, mProject, mDisplayName, mDescription, mAdditionalAttributes, mLocation, mLabels, mNetworkTags interface{}
		if pName, ok := p["name"]; ok {
			mName = pName
		}
		if pAssetType, ok := p["assetType"]; ok {
			mAssetType = pAssetType
		}
		if pProject, ok := p["project"]; ok {
			mProject = pProject
		}
		if pDisplayName, ok := p["displayName"]; ok {
			mDisplayName = pDisplayName
		}
		if pDescription, ok := p["description"]; ok {
			mDescription = pDescription
		}
		if pAdditionalAttributes, ok := p["additionalAttributes"]; ok {
			mAdditionalAttributes = pAdditionalAttributes
		}
		if pLocation, ok := p["location"]; ok {
			mLocation = pLocation
		}
		if pLabels, ok := p["labels"]; ok {
			mLabels = pLabels
		}
		if pNetworkTags, ok := p["networkTags"]; ok {
			mNetworkTags = pNetworkTags
		}
		results = append(results, map[string]interface{}{
			"name":                  mName,
			"asset_type":            mAssetType,
			"project":               mProject,
			"display_name":          mDisplayName,
			"description":           mDescription,
			"additional_attributes": mAdditionalAttributes,
			"location":              mLocation,
			"labels":                mLabels,
			"network_tags":          mNetworkTags,
		})
	}

	return results
}

func addArrayQueryParam(rawurl string, param string, values []interface{}) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for _, v := range values {
		q.Add(param, v.(string))
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
