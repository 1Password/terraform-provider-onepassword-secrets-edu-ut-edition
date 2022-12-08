terraform {
  required_providers {
    onepassword-terraform-edu-ut-edition = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword-terraform-edu-ut-edition" {}

data "onepassword-terraform-edu-ut-edition_reference" "edu" {
  vault = "test"
  item  = "login"
  field = "password"
}

output "login_secret" {
  value = data.onepassword-terraform-edu-ut-edition_reference.edu.secret
}


