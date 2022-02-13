package arvanecc

import (
	"context"
	"os"

	"github.com/arvancloud/ecc-go-client"
	"github.com/arvancloud/ecc-go-client/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *api.APIClient
}

// GetSchema
func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"apikey": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"access_token": {
				Type:      types.StringType,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
		},
	}, nil
}

// Provider schema struct
type providerData struct {
	APIKey      types.String `tfsdk:"apikey"`
	AccessToken types.String `tfsdk:"access_token"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// APIKey for authentication Arvan APIs
	var apikey string

	// AccessToken for authentication Arvan APIs
	var token string

	if config.APIKey.Null {
		apikey = os.Getenv("ARVAN_APIKEY")
	} else {
		apikey = config.APIKey.Value
	}

	if apikey == "" {
		if config.AccessToken.Null {
			token = os.Getenv("ARVAN_ACCESS_TOKEN")
		} else {
			token = config.AccessToken.Value
		}

		if token == "" {
			resp.Diagnostics.AddError(
				"Unable to find token or APIKey",
				"APIKey or token must be present in config or environment variables (ARVAN_APIKEY, ARVAN_ACCESS_TOKEN) for authenticate your account in arvan provider",
			)
			return
		}
	}

	cfg := ecc.NewConfiguration()
	cfg.AccessToken = token
	cfg.APIKey = apikey

	c := api.NewAPIClient(cfg)

	p.client = c
	p.configured = true
}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"arvanecc_ssh_key": resourceSSHKeyType{},
		// "arvanecc_server":  resourceServerType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
