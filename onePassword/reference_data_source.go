package onePassword

import (
	"context"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &referenceDataSource{}

func NewReferenceDataSource() datasource.DataSource {
	return &referenceDataSource{}
}

// data source implementation.
type referenceDataSource struct {
}

type referenceDataSourceModel struct {
	Vault  types.String `tfsdk:"vault"`
	Item   types.String `tfsdk:"item"`
	Field  types.String `tfsdk:"field"`
	Secret types.String `tfsdk:"secret"`
}

// Metadata returns the data source type name.
func (d *referenceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reference"
}

// GetSchema defines the schema for the data source.
func (d *referenceDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
func (d *referenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data referenceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	var reference = "op://" + data.Vault.Value + "/" + data.Item.Value + "/" + data.Field.Value
	out, err := exec.Command("op", "read", reference).Output()

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Secret Reference",
			err.Error(),
		)
		return
	}

	var response = string(out)
	var secret = strings.TrimSpace(response)
	data.Secret = types.StringValue(secret)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
