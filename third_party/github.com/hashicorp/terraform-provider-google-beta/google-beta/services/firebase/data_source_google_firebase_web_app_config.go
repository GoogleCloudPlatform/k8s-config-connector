// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase

import (
	"context"
	"fmt"

	"google.golang.org/api/firebase/v1beta1"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwmodels"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwtransport"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ datasource.DataSource              = &GoogleFirebaseWebAppConfigDataSource{}
	_ datasource.DataSourceWithConfigure = &GoogleFirebaseWebAppConfigDataSource{}
)

func NewGoogleFirebaseWebAppConfigDataSource() datasource.DataSource {
	return &GoogleFirebaseWebAppConfigDataSource{}
}

// GoogleFirebaseWebAppConfigDataSource defines the data source implementation
type GoogleFirebaseWebAppConfigDataSource struct {
	client  *firebase.Service
	project types.String
}

type GoogleFirebaseWebAppConfigModel struct {
	Id                types.String `tfsdk:"id"`
	WebAppId          types.String `tfsdk:"web_app_id"`
	ApiKey            types.String `tfsdk:"api_key"`
	AuthDomain        types.String `tfsdk:"auth_domain"`
	DatabaseUrl       types.String `tfsdk:"database_url"`
	LocationId        types.String `tfsdk:"location_id"`
	MeasurementId     types.String `tfsdk:"measurement_id"`
	MessagingSenderId types.String `tfsdk:"messaging_sender_id"`
	StorageBucket     types.String `tfsdk:"storage_bucket"`
	Project           types.String `tfsdk:"project"`
}

func (d *GoogleFirebaseWebAppConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firebase_web_app_config"
}

func (d *GoogleFirebaseWebAppConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "A Google Cloud Firebase web application configuration",

		Attributes: map[string]schema.Attribute{
			"web_app_id": schema.StringAttribute{
				Description:         "The id of the Firebase web App.",
				MarkdownDescription: "The id of the Firebase web App.",
				Required:            true,
			},

			"project": schema.StringAttribute{
				Description:         "The project id of the Firebase web App.",
				MarkdownDescription: "The project id of the Firebase web App.",
				Optional:            true,
			},

			"api_key": schema.StringAttribute{
				Description:         "The API key associated with the web App.",
				MarkdownDescription: "The API key associated with the web App.",
				Computed:            true,
			},

			"auth_domain": schema.StringAttribute{
				Description:         "The domain Firebase Auth configures for OAuth redirects, in the format `projectId.firebaseapp.com`",
				MarkdownDescription: "The domain Firebase Auth configures for OAuth redirects, in the format `projectId.firebaseapp.com`",
				Computed:            true,
			},

			"database_url": schema.StringAttribute{
				Description:         "The default Firebase Realtime Database URL.",
				MarkdownDescription: "The default Firebase Realtime Database URL.",
				Computed:            true,
			},

			"location_id": schema.StringAttribute{
				Description: "The ID of the project's default GCP resource location. The location is one of the available GCP resource locations. " +
					"This field is omitted if the default GCP resource location has not been finalized yet. To set your project's " +
					"default GCP resource location, call defaultLocation.finalize after you add Firebase services to your project.",
				MarkdownDescription: "The ID of the project's default GCP resource location. The location is one of the available GCP resource locations. " +
					"This field is omitted if the default GCP resource location has not been finalized yet. To set your project's " +
					"default GCP resource location, call defaultLocation.finalize after you add Firebase services to your project.",
				Computed: true,
			},

			"measurement_id": schema.StringAttribute{
				Description: "The unique Google-assigned identifier of the Google Analytics web stream associated with the Firebase Web App. " +
					"Firebase SDKs use this ID to interact with Google Analytics APIs. " +
					"This field is only present if the App is linked to a web stream in a Google Analytics App + Web property. " +
					"Learn more about this ID and Google Analytics web streams in the Analytics documentation. " +
					"To generate a measurementId and link the Web App with a Google Analytics web stream, call projects.addGoogleAnalytics.",
				MarkdownDescription: "The unique Google-assigned identifier of the Google Analytics web stream associated with the Firebase Web App. " +
					"Firebase SDKs use this ID to interact with Google Analytics APIs. " +
					"This field is only present if the App is linked to a web stream in a Google Analytics App + Web property. " +
					"Learn more about this ID and Google Analytics web streams in the Analytics documentation. " +
					"To generate a measurementId and link the Web App with a Google Analytics web stream, call projects.addGoogleAnalytics.",
				Computed: true,
			},

			"messaging_sender_id": schema.StringAttribute{
				Description:         "The sender ID for use with Firebase Cloud Messaging.",
				MarkdownDescription: "The sender ID for use with Firebase Cloud Messaging.",
				Computed:            true,
			},

			"storage_bucket": schema.StringAttribute{
				Description:         "The default Cloud Storage for Firebase storage bucket name.",
				MarkdownDescription: "The default Cloud Storage for Firebase storage bucket name.",
				Computed:            true,
			},

			"id": schema.StringAttribute{
				Description:         "Firebase Web App Config identifier",
				MarkdownDescription: "Firebase Web App Config identifier",
				Computed:            true,
			},
		},
	}
}

func (d *GoogleFirebaseWebAppConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	p, ok := req.ProviderData.(*fwtransport.FrameworkProviderConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *fwtransport.FrameworkProviderConfig, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = p.NewFirebaseClient(p.UserAgent, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	d.project = p.Project
}

func (d *GoogleFirebaseWebAppConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GoogleFirebaseWebAppConfigModel
	var metaData *fwmodels.ProviderMetaModel

	// Read Provider meta into the meta model
	resp.Diagnostics.Append(req.ProviderMeta.Get(ctx, &metaData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	d.client.UserAgent = fwtransport.GenerateFrameworkUserAgentString(metaData, d.client.UserAgent)

	client := firebase.NewProjectsWebAppsService(d.client)

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Project = fwresource.GetProjectFramework(data.Project, d.project, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	appName := fmt.Sprintf("projects/%s/webApps/%s/config", data.Project.ValueString(), data.WebAppId.ValueString())
	data.Id = data.WebAppId

	clientResp, err := client.GetConfig(appName).Do()
	if err != nil {
		fwtransport.HandleDatasourceNotFoundError(ctx, err, &resp.State, fmt.Sprintf("dataSourceFirebaseWebAppConfig %q", data.WebAppId.ValueString()), &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Trace(ctx, "read firebase web app config data source")

	data.ApiKey = types.StringValue(clientResp.ApiKey)
	data.AuthDomain = types.StringValue(clientResp.AuthDomain)
	data.DatabaseUrl = types.StringValue(clientResp.DatabaseURL)
	data.LocationId = types.StringValue(clientResp.LocationId)
	data.MeasurementId = types.StringValue(clientResp.MeasurementId)
	data.MessagingSenderId = types.StringValue(clientResp.MessagingSenderId)
	data.StorageBucket = types.StringValue(clientResp.StorageBucket)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
