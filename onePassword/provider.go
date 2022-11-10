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
	_ provider.Provider = &onePasswordProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &onePasswordProvider{}
}

// onePasswordProvider is the provider implementation.
type onePasswordProvider struct{}

// Metadata returns the provider type name.
func (p *onePasswordProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "onepassword"
}

// onePasswordProviderModel maps provider schema data to a Go type.
type onePasswordProviderModel struct {
}

// GetSchema defines the provider-level schema for configuration data.
func (p *onePasswordProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}

// Configure prepares a onePassword API client for data sources and resources.
func (p *onePasswordProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *onePasswordProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewReferenceDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *onePasswordProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSecretResource,
	}
}
