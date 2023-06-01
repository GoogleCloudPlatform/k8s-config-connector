// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"context"
	"fmt"

	"google.golang.org/api/firebase/v1beta1"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ datasource.DataSource              = &GoogleFirebaseAndroidAppConfigDataSource{}
	_ datasource.DataSourceWithConfigure = &GoogleFirebaseAndroidAppConfigDataSource{}
)

func NewGoogleFirebaseAndroidAppConfigDataSource() datasource.DataSource {
	return &GoogleFirebaseAndroidAppConfigDataSource{}
}

// GoogleFirebaseAndroidAppConfigDataSource defines the data source implementation
type GoogleFirebaseAndroidAppConfigDataSource struct {
	client  *firebase.Service
	project types.String
}

type GoogleFirebaseAndroidAppConfigModel struct {
	Id                 types.String `tfsdk:"id"`
	AppId              types.String `tfsdk:"app_id"`
	ConfigFilename     types.String `tfsdk:"config_filename"`
	ConfigFileContents types.String `tfsdk:"config_file_contents"`
	Project            types.String `tfsdk:"project"`
}

func (d *GoogleFirebaseAndroidAppConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firebase_android_app_config"
}

func (d *GoogleFirebaseAndroidAppConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "A Google Cloud Firebase Android application configuration",

		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Description:         "The id of the Firebase Android App.",
				MarkdownDescription: "The id of the Firebase Android App.",
				Required:            true,
			},

			"project": schema.StringAttribute{
				Description:         "The project id of the Firebase Android App.",
				MarkdownDescription: "The project id of the Firebase Android App.",
				Optional:            true,
			},

			"config_filename": schema.StringAttribute{
				Description:         "The filename that the configuration artifact for the AndroidApp is typically saved as.",
				MarkdownDescription: "The filename that the configuration artifact for the AndroidApp is typically saved as.",
				Computed:            true,
			},

			"config_file_contents": schema.StringAttribute{
				Description:         "The content of the XML configuration file as a base64-encoded string.",
				MarkdownDescription: "The content of the XML configuration file as a base64-encoded string.",
				Computed:            true,
			},

			"id": schema.StringAttribute{
				Description:         "Firebase Android App Config identifier",
				MarkdownDescription: "Firebase Android App Config identifier",
				Computed:            true,
			},
		},
	}
}

func (d *GoogleFirebaseAndroidAppConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	p, ok := req.ProviderData.(*frameworkProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *frameworkProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = p.NewFirebaseClient(p.userAgent, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	d.project = p.project
}

func (d *GoogleFirebaseAndroidAppConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GoogleFirebaseAndroidAppConfigModel
	var metaData *ProviderMetaModel

	// Read Provider meta into the meta model
	resp.Diagnostics.Append(req.ProviderMeta.Get(ctx, &metaData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	d.client.UserAgent = generateFrameworkUserAgentString(metaData, d.client.UserAgent)

	client := firebase.NewProjectsAndroidAppsService(d.client)

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Project = getProjectFramework(data.Project, d.project, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	appName := fmt.Sprintf("projects/%s/androidApps/%s/config", data.Project.ValueString(), data.AppId.ValueString())
	data.Id = types.StringValue(appName)

	clientResp, err := client.GetConfig(appName).Do()
	if err != nil {
		handleDatasourceNotFoundError(ctx, err, &resp.State, fmt.Sprintf("dataSourceFirebaseAndroidAppConfig %q", data.AppId.ValueString()), &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Trace(ctx, "read firebase android app config data source")

	data.ConfigFilename = types.StringValue(clientResp.ConfigFilename)
	data.ConfigFileContents = types.StringValue(clientResp.ConfigFileContents)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
