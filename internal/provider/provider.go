package provider

import (
	"context"
	"github.com/d-strobel/terraform-provider-windows/internal/generate/provider_windows"
	"github.com/d-strobel/terraform-provider-windows/internal/provider/local"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure WindowsProvider satisfies various provider interfaces.
var _ provider.Provider = &WindowsProvider{}

// WindowsProvider defines the provider implementation.
type WindowsProvider struct {
	version string
}

// Metadata returns the provider metadata.
func (p *WindowsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "windows"
	resp.Version = p.version
}

// Schema returns the provider schema.
// It includes the schema generated by the terraform-plugin-framework code generator
// and a given description which is used in the provider documentation.
func (p *WindowsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = provider_windows.WindowsProviderSchema(ctx)
}

// Resources returns the provider resources.
// All resources must be returned as functions.
func (p *WindowsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		local.NewLocalGroupResource,
		local.NewLocalUserResource,
		local.NewLocalGroupMemberResource,
	}
}

// DataSources returns the provider data sources.
// All data sources must be returned as functions.
func (p *WindowsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		local.NewLocalGroupDataSource,
		local.NewLocalGroupsDataSource,
		local.NewLocalUserDataSource,
		local.NewLocalUsersDataSource,
		local.NewLocalGroupMembersDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &WindowsProvider{
			version: version,
		}
	}
}
