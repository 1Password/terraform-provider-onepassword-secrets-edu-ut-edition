package onePassword

import (
	"context"
	"os/exec"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &secretResource{}
)

// NewSecretResource is a helper function to simplify the provider implementation.
func NewSecretResource() resource.Resource {
	return &secretResource{}
}

// secretResource is the resource implementation.
type secretResource struct{}

type secretResourceModel struct {
	ID             types.String        `tfsdk:"id"`
	Title          types.String        `tfsdk:"title"`
	Vault          types.String        `tfsdk:"vault"`
	Created        types.String        `tfsdk:"created"`
	Updated        types.String        `tfsdk:"updated"`
	Favorite       types.String        `tfsdk:"favorite"`
	Version        types.String        `tfsdk:"version"`
	Category       types.String        `tfsdk:"category"`
	PasswordRecipe passwordRecipeModel `tfsdk:"password_recipe"`
}

type passwordRecipeModel struct {
	CharacterSet types.Set   `tfsdk:"character_set"`
	Length       types.Int64 `tfsdk:"length"`
}

// Metadata returns the resource type name.
func (r *secretResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

// GetSchema defines the schema for the resource.
func (r *secretResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "The id of the secret",
				Type:        types.StringType,
				Computed:    true,
			},
			"title": {
				Description: "The title of the secret",
				Type:        types.StringType,
				Optional:    true,
			},
			"vault": {
				Description: "The vault associated with the secret",
				Type:        types.StringType,
				Optional:    true,
			},
			"created": {
				Description: "The time the secret was created",
				Type:        types.StringType,
				Computed:    true,
			},
			"updated": {
				Description: "The time the secret was last updated",
				Type:        types.StringType,
				Computed:    true,
			},
			"favorite": {
				Description: "Whether the secret is favourited or not",
				Type:        types.StringType,
				Computed:    true,
			},
			"version": {
				Description: "The version of the secret",
				Type:        types.StringType,
				Computed:    true,
			},
			"category": {
				Description: "The category of the secret",
				Type:        types.StringType,
				// Optional:    true,
				Computed: true,
			},
			"password_recipe": {
				Description: "The password recipe for the secret",
				Optional:    true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"character_set": {
						Description: "The id of the secret",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
					},
					"length": {
						Description: "The title of the secret",
						Type:        types.Int64Type,
						Optional:    true,
					},
					// }),
				}),
				// PlanModifiers: tfsdk.AttributePlanModifiers{
				// 	tfsdk.UseStateForUnknown(),
				// },
			},
		},
	}, nil
}

// Create creates the resource and sets the initial Terraform state.
func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data secretResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	var characters string = ""
	for _, s := range data.PasswordRecipe.CharacterSet.Elements() {
		characters += s.String() + ","
	}

	var password_recipe_flag = "=" + characters + strconv.FormatInt(int64(data.PasswordRecipe.Length.Value), 10)

	out, err := exec.Command("op", "item", "create", "--category", "password", "--title", data.Title.Value, "--vault", data.Vault.Value, "--generate-password", password_recipe_flag).Output()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating secret",
			"Could not create secret, unexpected error: "+err.Error(),
		)
		return
	}

	var response = string(out)

	idLine := response[strings.Index(response, "ID:"):strings.Index(response, "Title")]
	id := strings.TrimSpace(strings.TrimPrefix(idLine, "ID:"))

	createdLine := response[strings.Index(response, "Created:"):strings.Index(response, "Updated")]
	created := strings.TrimSpace(strings.TrimPrefix(createdLine, "Created:"))

	updatedLine := response[strings.Index(response, "Updated:"):strings.Index(response, "Favorite")]
	updated := strings.TrimSpace(strings.TrimPrefix(updatedLine, "Updated:"))

	favoriteLine := response[strings.Index(response, "Favorite:"):strings.Index(response, "Version")]
	favorite := strings.TrimSpace(strings.TrimPrefix(favoriteLine, "Favorite:"))

	versionLine := response[strings.Index(response, "Version:"):strings.Index(response, "Category")]
	version := strings.TrimSpace(strings.TrimPrefix(versionLine, "Version:"))

	// categoryLine := response[strings.Index(response, "Category:"):]
	// category := strings.TrimSpace(strings.TrimPrefix(categoryLine, "Category:"))

	data.ID = types.StringValue(id)
	data.Created = types.StringValue(created)
	data.Updated = types.StringValue(updated)
	data.Favorite = types.StringValue(favorite)
	data.Version = types.StringValue(version)
	data.Category = types.StringValue("password")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *secretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data secretResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	// Typically resources will make external calls, however this example
	// omits any refreshed data updates for brevity.

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *secretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *secretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data secretResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	out, err := exec.Command("op", "item", "delete", data.Title.Value, "--vault", data.Vault.Value).Output()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting secret",
			"Could not delete secret, unexpected error: "+err.Error(),
		)
		return
	}
	out = out
	// var response = string(out)
	// Need to modify data
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

}
