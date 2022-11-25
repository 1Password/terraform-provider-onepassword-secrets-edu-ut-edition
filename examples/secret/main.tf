terraform {
  required_providers {
    onepassword = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword" {}

resource "onepassword_secret" "edu" {
  vault = "test"
  title = "newtitle2"
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
  value = onepassword_secret.edu
}
