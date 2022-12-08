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
  title = "newtitle"
  password_recipe = {
    character_set = ["digits", "letters"]
    length        = 30
  }
  field_name  = "cellnumber2"
  field_type  = "phone"
  field_value = "12312341234"
}

output "new_secret" {
  value = onepassword-secrets-edu-ut-edition_secret.edu
}
