# 1Password Terraform Framework Provider

Use the 1Password Terraform Framwork Provider to read, create, delete or edit items in your 1Password Vaults.

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-onepassword
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
    onepassword = {
      
      // Local provider - change once published
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword" {}

data "onepassword_reference" "edu" {
  vault = "test"
  item  = "login"
  field = "password"
}

output "login_secret" {
  value = data.onepassword_reference.edu.secret
}


```
