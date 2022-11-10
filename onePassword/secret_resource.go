package onePassword

import (
	"context"
	"os/exec"
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
	ID       types.String `tfsdk:"id"`
	Title    types.String `tfsdk:"title"`
	Vault    types.String `tfsdk:"vault"`
	Created  types.String `tfsdk:"created"`
	Updated  types.String `tfsdk:"updated"`
	Favorite types.String `tfsdk:"favorite"`
	Version  types.String `tfsdk:"version"`
	Category types.String `tfsdk:"category"`
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
				Required:    true,
			},
		},
	}, nil
}

// Create creates the resource and sets the initial Terraform state.
func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data secretResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	out, err := exec.Command("op", "item", "create", "--category", data.Category.Value, "--title", data.Title.Value, "--vault", data.Vault.Value).Output()
	var response = string(out)

	idLine := response[strings.Index(response, "ID:"):strings.Index(response, "Title")]
	id := strings.TrimSpace(strings.TrimPrefix(idLine, "ID:"))

	titleLine := response[strings.Index(response, "Title:"):strings.Index(response, "Vault")]
	title := strings.TrimSpace(strings.TrimPrefix(titleLine, "Title:"))

	vaultReferenceLine := response[strings.Index(response, "Vault:"):strings.Index(response, "Created")]
	vaultReference := strings.TrimSpace(strings.TrimPrefix(vaultReferenceLine, "Vault:"))
	vault := vaultReference[:strings.Index(vaultReference, " ")]

	createdLine := response[strings.Index(response, "Created:"):strings.Index(response, "Updated")]
	created := strings.TrimSpace(strings.TrimPrefix(createdLine, "Created:"))

	updatedLine := response[strings.Index(response, "Updated:"):strings.Index(response, "Favorite")]
	updated := strings.TrimSpace(strings.TrimPrefix(updatedLine, "Updated:"))

	favoriteLine := response[strings.Index(response, "Favorite:"):strings.Index(response, "Version")]
	favorite := strings.TrimSpace(strings.TrimPrefix(favoriteLine, "Favorite:"))

	versionLine := response[strings.Index(response, "Version:"):strings.Index(response, "Category")]
	version := strings.TrimSpace(strings.TrimPrefix(versionLine, "Version:"))

	categoryLine := response[strings.Index(response, "Category:"):]
	category := strings.TrimSpace(strings.TrimPrefix(categoryLine, "Category:"))

	data.ID = types.StringValue(id)
	data.Title = types.StringValue(title)
	data.Vault = types.StringValue(vault)
	data.Created = types.StringValue(created)
	data.Updated = types.StringValue(updated)
	data.Favorite = types.StringValue(favorite)
	data.Version = types.StringValue(version)
	data.Category = types.StringValue(category)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating secret",
			"Could not create secret, unexpected error: "+err.Error(),
		)
	}

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
}
