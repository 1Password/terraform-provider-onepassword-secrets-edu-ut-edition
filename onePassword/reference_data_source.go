package onePassword

import (
	"bufio"
	"context"
	"log"
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
	// this is the 1Password CLI command to read
	cmd := exec.Command("op", "read", reference)

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

	// instantiate scanner and read standard output
	scanner := bufio.NewScanner(stdout)
	scanner.Scan()
	var output = string(scanner.Text())

	// instantiate scanner and read standard error
	scanner = bufio.NewScanner(stderr)
	scanner.Scan()
	var error_message = string(scanner.Text())

	// errors from CLI command we will be checking
	var doesnt_exist = "isn't an item in the \"" + data.Vault.Value + "\" vault"
	var duplicate = "More than one item matches \"" + data.Item.Value + "\""

	if strings.Contains(error_message, doesnt_exist) {
		var error_string = "Could not find item error."

		var start_index = strings.Index(error_message, string(data.Field.Value)) + len(string(data.Field.Value)) + 2
		var end_index = strings.Index(error_message, "vault.")
		var detail_string = "Details: " + error_message[start_index:end_index]

		resp.Diagnostics.AddError(error_string, detail_string)

	} else if strings.Contains(error_message, duplicate) {
		var error_string = "Duplicate Item in Vault error."

		var start_index = strings.Index(error_message, "More")
		var detail_string = error_message[start_index:] + "\n"

		for scanner.Scan() {
			error_message = string(scanner.Text())
			detail_string = detail_string + error_message + "\n"
		}

		resp.Diagnostics.AddError(error_string, detail_string[:strings.LastIndex(detail_string, "\n")])

	} else {
		var secret = strings.TrimSpace(output)
		data.Secret = types.StringValue(secret)
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

}
