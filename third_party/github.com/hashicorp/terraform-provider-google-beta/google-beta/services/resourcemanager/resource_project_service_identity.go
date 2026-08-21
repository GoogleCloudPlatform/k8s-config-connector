// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/serviceusage"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProjectServiceIdentity() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectServiceIdentityCreate,
		Read:   resourceProjectServiceIdentityRead,
		Delete: resourceProjectServiceIdentityDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceProjectServiceIdentityCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}:generateServiceIdentity")
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	billingProject := project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Service Identity: %s", err)
	}

	var opRes map[string]interface{}
	err = serviceusage.ServiceUsageOperationWaitTimeWithResponse(
		config, res, &opRes, billingProject, "Creating Service Identity", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished creating Service Identity %q: %#v", d.Id(), res)

	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/services/{{service}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// This API may not return the service identity's details, even if the relevant
	// Google API is configured for service identities.
	if emailVal, ok := opRes["email"]; ok {
		email, ok := emailVal.(string)
		if !ok {
			return fmt.Errorf("unexpected type for email: got %T, want string", email)
		}
		if err := d.Set("email", email); err != nil {
			return fmt.Errorf("Error setting email: %s", err)
		}
	}
	return nil
}

// There is no read endpoint for this API.
func resourceProjectServiceIdentityRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// There is no delete endpoint for this API.
func resourceProjectServiceIdentityDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
