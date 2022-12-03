package onePassword

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
	"runtime"
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
	ID    types.String `tfsdk:"id"`
	Title types.String `tfsdk:"title"`
	// tag            types.String        `tfsdk:"tag"`
	// url            types.String        `tfsdk:"url"`
	Vault          types.String        `tfsdk:"vault"`
	Created        types.String        `tfsdk:"created"`
	Updated        types.String        `tfsdk:"updated"`
	Favorite       types.String        `tfsdk:"favorite"`
	Version        types.String        `tfsdk:"version"`
	Category       types.String        `tfsdk:"category"`
	PasswordRecipe passwordRecipeModel `tfsdk:"password_recipe"`
	FieldName      types.String        `tfsdk:"field_name"`
	FieldType      types.String        `tfsdk:"field_type"`
	FieldValue     types.String        `tfsdk:"field_value"`
	DeleteField    types.Bool          `tfsdk:"delete_field"`
	UpdatePassword types.Bool          `tfsdk:"update_password"`
}

type passwordRecipeModel struct {
	CharacterSet types.Set   `tfsdk:"character_set"`
	Length       types.Int64 `tfsdk:"length"`
}

type updateResponse struct{}

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
			// "tag": { // s
			// 	Description: "The tags of the secret",
			// 	Type:        types.StringType,
			// 	Optional:    true,
			// },
			// "url": {
			// 	Description: "The url of the secret",
			// 	Type:        types.StringType,
			// 	Optional:    true,
			// },
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
			// for op item edit (update)
			"field_name": {
				Description: "name of field of the item for creation/update",
				Type:        types.StringType,
				Optional:    true,
			},
			"field_type": {
				Description: "type of field of the item for creation/update",
				Type:        types.StringType,
				Optional:    true,
			},
			"field_value": {
				Description: "value of field of the item for creation/update",
				Type:        types.StringType,
				Optional:    true,
			},
			"delete_field": {
				Description: "if true field of the item is deleted for update",
				Type:        types.BoolType,
				Optional:    true,
			},
			"update_password": {
				Description: "if true password of the item is re auto-generated using password recipe",
				Type:        types.BoolType,
				Optional:    true,
			},
		},
	}, nil
}

// Create creates the resource and sets the initial Terraform state.
func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data secretResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	var requested_password_length = int64(data.PasswordRecipe.Length.Value)
	if requested_password_length < 1 || requested_password_length > 64 {
		resp.Diagnostics.AddError("Password recipe not vaild error.", "Password length must be between 1 and 64 inclusive.")
		return
	}

	var characters string = ""
	for _, s := range data.PasswordRecipe.CharacterSet.Elements() {
		characters += s.String() + ","
	}

	characters = strings.ReplaceAll(characters, "\"", "")

	var password_recipe_flag = "=" + characters + strconv.FormatInt(requested_password_length, 10)
	var generate_password_string = "--generate-password" + password_recipe_flag

	cmd := exec.Command("op", "item", "create", "--category", "password", "--title", data.Title.Value, "--vault", data.Vault.Value, generate_password_string)

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// create shell script with op create command
		err := os.WriteFile("../../temp/linux_create.sh", []byte("op item create --category password --title "+data.Title.Value+" --vault "+data.Vault.Value+" "+generate_password_string), 0755)
		// detect if the script was not created properly
		if err != nil {
			// log the error
			log.Fatal("Error writing to shell script: %v", err)
		}

		// execute the shell script
		// The shell script is executed here instead of the normal command as biometrics is not triggered
		// using the normal command.
		cmd = exec.Command("/bin/sh", "../../temp/linux_create.sh")
	}

	// create standard output pipe and error check
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	// create standard error pipe and error check
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// run CLI command and error check the call itself not the output
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// instantiate scanner and read standard error
	scanner := bufio.NewScanner(stderr)
	scanner.Scan()
	var error_message = string(scanner.Text())

	// errors from CLU command we will be checking
	var vault_doesnt_exist = "\"" + data.Vault.Value + "\" isn't a vault in this account."
	var invalid_arguments_to_generate_password = "Value must be one of `letters,digits,symbols`"

	if strings.Contains(error_message, vault_doesnt_exist) {
		var error_string = "Vault does not exist error."

		var start_index = strings.Index(error_message, string(data.Vault.Value))
		var end_index = strings.Index(error_message, "account.") + len("account.")
		var detail_string = "Details: " + error_message[start_index:end_index]

		resp.Diagnostics.AddError(error_string, detail_string)
		return
	} else if strings.Contains(error_message, invalid_arguments_to_generate_password) {
		var error_string = "Invalid arguments in character set error."

		var start_index = strings.Index(error_message, "invalid argument")
		var detail_string = "Details: " + error_message[start_index:]

		resp.Diagnostics.AddError(error_string, detail_string)
		return
	} else {
		// instantiate scanner and read standard output
		var output = ""
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			var temp = string(scanner.Text())
			output = output + temp + "\n"
		}

		var response = output

		scanner = bufio.NewScanner(stdout)
		for scanner.Scan() {
			output = string(scanner.Text())
			response = response + output
		}

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

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// remove the shell script
		err = os.Remove("../../temp/linux_create.sh")
		// if an error occurs while deleting the shell script
		if err != nil {
			// log the error
			log.Fatal("Error deleting shell script: %v", err)
		}
	}
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

	// Read Terraform changed state data into the model
	var data secretResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read Terraform prior state data into the model
	var prevState secretResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prevState)...)

	// Different cases of update: execute commands by the optional fields that exist
	args := []string{}
	args = append(args, "item")
	args = append(args, "edit")

	args = append(args, prevState.ID.Value)
	args = append(args, "--vault")
	args = append(args, prevState.Vault.Value)

	// update title
	if data.Title.Value != "" {
		args = append(args, "title="+data.Title.Value)
	}

	// update password
	if data.UpdatePassword.Value == true {
		// TODO: error check if password recipe is given, if not return with error message

		var requested_password_length = int64(data.PasswordRecipe.Length.Value)
		if requested_password_length < 1 || requested_password_length > 64 {
			resp.Diagnostics.AddError("Password recipe not vaild error.", "Password length must be between 1 and 64 inclusive.")
			return
		}
		var characters string = ""
		for _, s := range data.PasswordRecipe.CharacterSet.Elements() {
			characters += s.String() + ","
		}
		characters = strings.ReplaceAll(characters, "\"", "")
		var password_recipe_flag = "=" + characters + strconv.FormatInt(requested_password_length, 10)
		var generate_password_string = "--generate-password" + password_recipe_flag

		args = append(args, generate_password_string)
	}

	if data.DeleteField.Value == true {
		// delete given field
		if data.FieldName.Value != "" {
			args = append(args, data.FieldName.Value+"[delete]")
		} else {
			resp.Diagnostics.AddError(
				"Error updating secret item:",
				"if you want to delete a field, please specify the field name!",
			)
			return
		}
	} else if data.FieldName.Value != "" && data.FieldType.Value != "" && data.FieldValue.Value != "" {
		// set field name, type, value
		args = append(args, data.FieldName.Value+"["+data.FieldType.Value+"]"+"="+data.FieldValue.Value)
	} else if data.FieldName.Value != "" && data.FieldType.Value != "" {
		// set field name, type
		args = append(args, data.FieldName.Value+"["+data.FieldType.Value+"]")
	} else if data.FieldName.Value != "" && data.FieldValue.Value != "" {
		// set field name, value
		args = append(args, data.FieldName.Value+"="+data.FieldValue.Value)
	} else if data.FieldType.Value != "" && data.FieldValue.Value != "" {
		resp.Diagnostics.AddError(
			"Error updating secret item:",
			"if you have entered field type and value, please specify field name too!",
		)
		return
	}

	out, err := exec.Command("op", args...).Output()

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// create shell script with op delete command
		file_err := os.WriteFile("../../temp/linux_update.sh", []byte("op "+strings.Join(args, " ")), 0755)
		// detect if the script was not created properly
		if file_err != nil {
			// log the error
			log.Fatal("Error writing to shell script: %v", file_err)
		}

		// execute the shell script
		// The shell script is executed here instead of the normal command as biometrics is not triggered
		// using the normal command.
		out, err = exec.Command("/bin/sh", "../../temp/linux_update.sh").Output()
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating secret",
			"Could not update secret, unexpected error: "+strings.Join(args, " "),
		)
		return
	}

	var response = string(out)

	data.Category = types.StringValue(prevState.Category.Value)
	data.Created = types.StringValue(prevState.Created.Value)
	data.Favorite = types.StringValue(prevState.Favorite.Value)
	data.ID = types.StringValue(prevState.ID.Value)

	updatedLine := response[strings.Index(response, "Updated:"):strings.Index(response, "Favorite")]
	updated := strings.TrimSpace(strings.TrimPrefix(updatedLine, "Updated:"))
	data.Updated = types.StringValue(updated)

	versionLine := response[strings.Index(response, "Version:"):strings.Index(response, "Category")]
	version := strings.TrimSpace(strings.TrimPrefix(versionLine, "Version:"))
	data.Version = types.StringValue(version)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// remove the shell script
		err = os.Remove("../../temp/linux_update.sh")
		// if an error occurs while deleting the shell script
		if err != nil {
			// log the error
			log.Fatal("Error deleting shell script: %v", err)
		}
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *secretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data secretResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	out, err := exec.Command("op", "item", "delete", data.Title.Value, "--vault", data.Vault.Value).Output()

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// create shell script with op delete command
		file_err := os.WriteFile("../../temp/linux_delete.sh", []byte("op item delete "+data.Title.Value+" --vault "+data.Vault.Value), 0755)
		// detect if the script was not created properly
		if file_err != nil {
			// log the error
			log.Fatal("Error writing to shell script: %v", file_err)
		}

		// execute the shell script
		// The shell script is executed here instead of the normal command as biometrics is not triggered
		// using the normal command.
		out, err = exec.Command("/bin/sh", "../../temp/linux_delete.sh").Output()
	}

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

	// if linux environment is detected
	if runtime.GOOS == "linux" {
		// remove the shell script
		err = os.Remove("../../temp/linux_delete.sh")
		// if an error occurs while deleting the shell script
		if err != nil {
			// log the error
			log.Fatal("Error deleting shell script: %v", err)
		}
	}
}
