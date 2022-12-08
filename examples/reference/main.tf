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
  item  = "login"
  field = "password"
}

output "login_secret" {
  value = data.onepassword-secrets-edu-ut-edition_reference.edu.secret
}


