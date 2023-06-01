// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package gkehub

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgdclresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceGkeHubFeature() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkeHubFeatureCreate,
		Read:   resourceGkeHubFeatureRead,
		Update: resourceGkeHubFeatureUpdate,
		Delete: resourceGkeHubFeatureDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkeHubFeatureImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "GCP labels for this Feature.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The full, unique name of this Feature resource",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was created.",
			},

			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was deleted.",
			},

			"resource_state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "State of the Feature resource itself.",
				Elem:        GkeHubFeatureResourceStateSchema(),
			},

			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The Hub-wide Feature state",
				Elem:        GkeHubFeatureStateSchema(),
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. When the Feature resource was last updated.",
			},
		},
	}
}

func GkeHubFeatureSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fleetobservability": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Fleet Observability spec.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecFleetobservabilitySchema(),
			},

			"multiclusteringress": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Multicluster Ingress-specific spec.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecMulticlusteringressSchema(),
			},
		},
	}
}

func GkeHubFeatureSpecFleetobservabilitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Fleet Observability Logging-specific spec.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecFleetobservabilityLoggingConfigSchema(),
			},
		},
	}
}

func GkeHubFeatureSpecFleetobservabilityLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specified if applying the default routing config to logs not specified in other configs.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSchema(),
			},

			"fleet_scope_logs_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specified if applying the routing config to all logs for all fleet scopes.",
				MaxItems:    1,
				Elem:        GkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSchema(),
			},
		},
	}
}

func GkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The logs routing mode Possible values: MODE_UNSPECIFIED, COPY, MOVE",
			},
		},
	}
}

func GkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The logs routing mode Possible values: MODE_UNSPECIFIED, COPY, MOVE",
			},
		},
	}
}

func GkeHubFeatureSpecMulticlusteringressSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_membership": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`",
			},
		},
	}
}

func GkeHubFeatureResourceStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"has_resources": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the Feature resource in the Hub API. Possible values: STATE_UNSPECIFIED, ENABLING, ACTIVE, DISABLING, UPDATING, SERVICE_UPDATING",
			},
		},
	}
}

func GkeHubFeatureStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The \"running state\" of the Feature in this Hub.",
				Elem:        GkeHubFeatureStateStateSchema(),
			},
		},
	}
}

func GkeHubFeatureStateStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The high-level, machine-readable status of this Feature. Possible values: CODE_UNSPECIFIED, OK, WARNING, ERROR",
			},

			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable description of the current status.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"",
			},
		},
	}
}

func resourceGkeHubFeatureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   tpgresource.CheckStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := tpgresource.ReplaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	directive := tpgdclresource.CreateDirective
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLGkeHubClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyFeature(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureRead(d, meta)
}

func resourceGkeHubFeatureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   tpgresource.CheckStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLGkeHubClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetFeature(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("GkeHubFeature %q", d.Id())
		return tpgdclresource.HandleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("spec", flattenGkeHubFeatureSpec(res.Spec)); err != nil {
		return fmt.Errorf("error setting spec in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("delete_time", res.DeleteTime); err != nil {
		return fmt.Errorf("error setting delete_time in state: %s", err)
	}
	if err = d.Set("resource_state", flattenGkeHubFeatureResourceState(res.ResourceState)); err != nil {
		return fmt.Errorf("error setting resource_state in state: %s", err)
	}
	if err = d.Set("state", flattenGkeHubFeatureState(res.State)); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceGkeHubFeatureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   tpgresource.CheckStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := tpgresource.ReplaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	directive := tpgdclresource.UpdateDirective
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLGkeHubClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutUpdate))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyFeature(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureRead(d, meta)
}

func resourceGkeHubFeatureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.Feature{
		Location: dcl.String(d.Get("location").(string)),
		Labels:   tpgresource.CheckStringMap(d.Get("labels")),
		Name:     dcl.String(d.Get("name").(string)),
		Project:  dcl.String(project),
		Spec:     expandGkeHubFeatureSpec(d.Get("spec")),
	}
	lockName, err := tpgresource.ReplaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	log.Printf("[DEBUG] Deleting Feature %q", d.Id())
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLGkeHubClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	if err := client.DeleteFeature(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Feature: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Feature %q", d.Id())
	return nil
}

func resourceGkeHubFeatureImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkeHubFeatureSpec(o interface{}) *gkehub.FeatureSpec {
	if o == nil {
		return gkehub.EmptyFeatureSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpec{
		Fleetobservability:  expandGkeHubFeatureSpecFleetobservability(obj["fleetobservability"]),
		Multiclusteringress: expandGkeHubFeatureSpecMulticlusteringress(obj["multiclusteringress"]),
	}
}

func flattenGkeHubFeatureSpec(obj *gkehub.FeatureSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"fleetobservability":  flattenGkeHubFeatureSpecFleetobservability(obj.Fleetobservability),
		"multiclusteringress": flattenGkeHubFeatureSpecMulticlusteringress(obj.Multiclusteringress),
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecFleetobservability(o interface{}) *gkehub.FeatureSpecFleetobservability {
	if o == nil {
		return gkehub.EmptyFeatureSpecFleetobservability
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpecFleetobservability
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecFleetobservability{
		LoggingConfig: expandGkeHubFeatureSpecFleetobservabilityLoggingConfig(obj["logging_config"]),
	}
}

func flattenGkeHubFeatureSpecFleetobservability(obj *gkehub.FeatureSpecFleetobservability) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"logging_config": flattenGkeHubFeatureSpecFleetobservabilityLoggingConfig(obj.LoggingConfig),
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecFleetobservabilityLoggingConfig(o interface{}) *gkehub.FeatureSpecFleetobservabilityLoggingConfig {
	if o == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecFleetobservabilityLoggingConfig{
		DefaultConfig:        expandGkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(obj["default_config"]),
		FleetScopeLogsConfig: expandGkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(obj["fleet_scope_logs_config"]),
	}
}

func flattenGkeHubFeatureSpecFleetobservabilityLoggingConfig(obj *gkehub.FeatureSpecFleetobservabilityLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"default_config":          flattenGkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(obj.DefaultConfig),
		"fleet_scope_logs_config": flattenGkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(obj.FleetScopeLogsConfig),
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(o interface{}) *gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if o == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfigDefaultConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfigDefaultConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{
		Mode: gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumRef(obj["mode"].(string)),
	}
}

func flattenGkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(obj *gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"mode": obj.Mode,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(o interface{}) *gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if o == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{
		Mode: gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumRef(obj["mode"].(string)),
	}
}

func flattenGkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(obj *gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"mode": obj.Mode,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureSpecMulticlusteringress(o interface{}) *gkehub.FeatureSpecMulticlusteringress {
	if o == nil {
		return gkehub.EmptyFeatureSpecMulticlusteringress
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return gkehub.EmptyFeatureSpecMulticlusteringress
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.String(obj["config_membership"].(string)),
	}
}

func flattenGkeHubFeatureSpecMulticlusteringress(obj *gkehub.FeatureSpecMulticlusteringress) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"config_membership": obj.ConfigMembership,
	}

	return []interface{}{transformed}

}

func flattenGkeHubFeatureResourceState(obj *gkehub.FeatureResourceState) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"has_resources": obj.HasResources,
		"state":         obj.State,
	}

	return []interface{}{transformed}

}

func flattenGkeHubFeatureState(obj *gkehub.FeatureState) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"state": flattenGkeHubFeatureStateState(obj.State),
	}

	return []interface{}{transformed}

}

func flattenGkeHubFeatureStateState(obj *gkehub.FeatureStateState) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"code":        obj.Code,
		"description": obj.Description,
		"update_time": obj.UpdateTime,
	}

	return []interface{}{transformed}

}
