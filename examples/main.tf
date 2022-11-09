terraform {
  required_providers {
    onepprovider = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}
#
#provider "onepasswordprovider" {}
#

provider "onepprovider" {}

data "onepprovider_items" "edu" {
  vault = "Test"
  item  = "Server"
  field = "Username"
}

output "edu_coffees" {
  value = data.onepprovider_items.edu.reference
}
