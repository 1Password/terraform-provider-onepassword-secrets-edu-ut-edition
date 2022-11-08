package onePassword

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = &opItemDataSource{}
)

func NewItemsDataSource() datasource.DataSource {
	return &opItemDataSource{}
}

// data source implementation.
type opItemDataSource struct {
	// client *onePassword.client
}

type itemsDataSourceModel struct {
	Items []itemsModel `tfsdk:"items"`
}
type itemsModel struct {
	vault types.String `tfsdk:"vault"`
	item  types.String `tfsdk:"item"`
	field types.String `tfsdk:"field"`
}

// Metadata returns the data source type name.
func (d *opItemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_items"
}

// GetSchema defines the schema for the data source.
func (d *opItemDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"items": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"vault": {
						Type:     types.StringType,
						Required: true,
					},

					"item": {
						Description: "The name of the item to retrieve.",
						Type:        types.StringType,
						Required:    true,
						Computed:    true,
					},
					"field": {
						Description: "The name of the field to retrieve.",
						Type:        types.StringType,
						Required:    true,
						Computed:    true,
					},
				}),
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (d *opItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	//var state itemsDataSourceModel

	//coffees, err := d.client.GetCoffees()
	//if err != nil {
	//	resp.Diagnostics.AddError(
	//		"Unable to Read HashiCups Coffees",
	//		err.Error(),
	//	)
	//	return
	//}

	// Map response body to model
	//for _, coffee := range coffees {
	//	coffeeState := coffeesModel{
	//		ID:          types.Int64Value(int64(coffee.ID)),
	//		Name:        types.StringValue(coffee.Name),
	//		Teaser:      types.StringValue(coffee.Teaser),
	//		Description: types.StringValue(coffee.Description),
	//		Price:       types.Float64Value(coffee.Price),
	//		Image:       types.StringValue(coffee.Image),
	//	}
	//
	//	for _, ingredient := range coffee.Ingredient {
	//		coffeeState.Ingredients = append(coffeeState.Ingredients, coffeesIngredientsModel{
	//			ID: types.Int64Value(int64(ingredient.ID)),
	//		})
	//	}
	//
	//	state.Coffees = append(state.Coffees, coffeeState)
	//}
	//
	//// Set state
	//diags := resp.State.Set(ctx, &state)
	//resp.Diagnostics.Append(diags...)
	//if resp.Diagnostics.HasError() {
	//	return
	//}
}
