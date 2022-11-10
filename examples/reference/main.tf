terraform {
  required_providers {
    onepassword = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword" {}

data "onepassword_reference" "edu" {
  vault = "test"
  item  = "login"
  field = "password"
}

output "login_secret" {
  value = data.onepassword_reference.edu.secret
}
