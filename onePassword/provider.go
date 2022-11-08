package onePassword

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &onepprovider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &onepprovider{}
}

// hashicupsProvider is the provider implementation.
type onepprovider struct{}

// Metadata returns the provider type name.
func (p *onepprovider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "onepprovider"
}

// hashicupsProviderModel maps provider schema data to a Go type.
type hashicupsProviderModel struct {
}

// GetSchema defines the provider-level schema for configuration data.
func (p *onepprovider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *onepprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

// DataSources defines the data sources implemented in the provider.
func (p *onepprovider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewItemsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *onepprovider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
