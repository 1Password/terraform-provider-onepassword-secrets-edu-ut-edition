terraform {
  required_providers {
    onepassword-terraform-edu-ut-edition = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword-terraform-edu-ut-edition" {}

resource "onepassword-terraform-edu-ut-edition_secret" "edu" {
  vault = "test"
  title = "newtitle"
  password_recipe = {
    character_set = ["digits", "letters"]
    length        = 30
  }
  field_name="cellnumber2"
  field_type="phone"
  field_value="12312341234"
}

output "new_secret" {
  value = onepassword-terraform-edu-ut-edition_secret.edu
}