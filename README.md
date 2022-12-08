# 1Password Terraform Framework Provider

Use the 1Password Terraform Framwork Provider to read, create, delete or edit items in your 1Password Vaults.

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-onepassword-secrets-edu-ut-edition
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```
For a local provider add to the ```.terraformrc``` file the following line:
```
dev_overrides {
      "hashicorp.com/edu/onepassword" = "<GOBIN>"
  }
  ```
  Where <GOBIN> is the directory of the go bin (obtained from ```go env GOBIN```).
  For more information please refer to the official documentation at 
  https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider.
  
Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
      
## Usage
  Detailed documentation for using this provider can be found on the Terraform registry docs.
      
### Data Source Read
```shell
terraform {
  required_providers {
    onepassword-secrets-edu-ut-edition = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword-secrets-edu-ut-edition" {}

data "onepassword-secrets-edu-ut-edition_reference" "edu" {
  vault = "test"
  item  = "uber"
  field = "password"
}

output "login_secret" {
  value = data.onepassword-secrets-edu-ut-edition_reference.edu.secret
}


```
### Resource create edit and delete
For creation and editing (after creation) run 'terraform apply'.

For deletion run 'terraform destroy'.
```shell
terraform {
  required_providers {
    onepassword-secrets-edu-ut-edition = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword-secrets-edu-ut-edition" {}

resource "onepassword-secrets-edu-ut-edition_secret" "edu" {
  vault = "test"
  title = "newtitle3"
  password_recipe = {
    character_set = ["digits", "letters"]
    length        = 30
  }
  field_name="cellnumber"
  field_type="phone"
  field_value="123-1234-1234"
  delete_field=false
  update_password=true
}

output "new_secret" {
  value = onepassword-secrets-edu-ut-edition_secret.edu
}
```
