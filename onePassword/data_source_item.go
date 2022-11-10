package onePassword

import (
	"context"
	"log"
	"os/exec"

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
}

type itemsModel struct {
	Vault  types.String `tfsdk:"vault"`
	Item   types.String `tfsdk:"item"`
	Field  types.String `tfsdk:"field"`
	Secret types.String `tfsdk:"secret"`
}

// Metadata returns the data source type name.
func (d *opItemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_items"
}

// GetSchema defines the schema for the data source.
func (d *opItemDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"vault": {
				Type:     types.StringType,
				Required: true,
			},

			"item": {
				Description: "The name of the item to retrieve.",
				Type:        types.StringType,
				Required:    true,
			},
			"field": {
				Description: "The name of the field to retrieve.",
				Type:        types.StringType,
				Required:    true,
			},
			"secret": {
				Description: "The secret of the field item in the vault",
				Type:        types.StringType,
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (d *opItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data itemsModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	var reference = "op://" + data.Vault.Value + "/" + data.Item.Value + "/" + data.Field.Value
	out, err := exec.Command("op", "read", reference).Output()

	if err != nil {
		log.Fatal(err)
	}

	// This might be causing the EOT in the output
	data.Secret = types.StringValue(string(out))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
