terraform {
  required_providers {
    onepassword = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword" {}

resource "onepassword_secret" "edu" {
  title = "newtitle"
  vault = "test"
  new_title="newtitle2"
  field_name="cellnumber"
  field_type="phone"
  field_value="12312341234"
}

output "new_secret" {
  value = onepassword_secret.edu
}