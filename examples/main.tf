terraform {
  required_providers {
    onepprovider = {
      source  = "hashicorp.com/edu/onepassword"

    }
  }
}
#
#provider "onepasswordprovider" {}
#

provider  "onepprovider" {}

data "onepprovider_items" "edu"{}

#output "edu_coffees" {
#  value = data.onepprovider_items.edu
#}
