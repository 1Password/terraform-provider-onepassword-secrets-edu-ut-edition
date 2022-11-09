package onePassword

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &opItemDataSource{}

func NewopItemsDataSource() datasource.DataSource {
	return &opItemDataSource{}
}

// data source implementation.
type opItemDataSource struct {
	// client *onePassword.client
}

//	type opItemDataSourceModel struct {
//		Items []itemsModel `tfsdk:"items"`
//	}
type itemsModel struct {
	Vault types.String `tfsdk:"vault"`
	Item  types.String `tfsdk:"item"`
	Field types.String `tfsdk:"field"`
}

// Metadata returns the data source type name.
func (d *opItemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	// fmt.Println(req.ProviderTypeName)
	resp.TypeName = req.ProviderTypeName + "_items"
	//resp.TypeName = "onepprovider_items"
}

// GetSchema defines the schema for the data source.
func (d *opItemDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"vault": {
				Type: types.StringType,
				// Required: true,
				Computed: true,
				// Optional: true,
			},

			"item": {
				Description: "The name of the item to retrieve.",
				Type:        types.StringType,
				// Required:    true,
				Computed: true,
				// Optional: true,
			},
			"field": {
				Description: "The name of the field to retrieve.",
				Type:        types.StringType,
				// Required:    true,
				Computed: true,
				// Optional: true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (d *opItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data itemsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	data.Vault = types.StringValue("alam")
	data.Field = types.StringValue("email")
	data.Item = types.StringValue("pass")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
