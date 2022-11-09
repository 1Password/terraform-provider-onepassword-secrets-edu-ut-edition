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
  # vault = "hi"
  # item  = "sk"
  # field = "sfj"
}
# data "fgdgd" "edu"{
#  vault = "hi"
#  item = "sk"
#  field = "sfj"
#}

output "edu_coffees" {
  value = data.onepprovider_items.edu
}
