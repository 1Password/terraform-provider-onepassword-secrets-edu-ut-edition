terraform {
  required_providers {
    onepassword-secrets-edu-ut-edition = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword-secrets-edu-ut-edition" {}

data "onepassword-secrets-edu-ut-edition" "edu" {
  vault = "test"
  item  = "uber4"
  field = "password"
}

output "login_secret" {
  value = data.onepassword-secrets-edu-ut-edition.edu.secret
}


