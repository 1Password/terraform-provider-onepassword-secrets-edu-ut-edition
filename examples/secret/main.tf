terraform {
  required_providers {
    onepassword = {
      source = "hashicorp.com/edu/onepassword"

    }
  }
}

provider "onepassword" {}

resource "onepassword_secret" "edu" {
  vault    = "test"
  title    = "uber"
  category = "LOGIN"
}

output "new_secret" {
  value = onepassword_secret.edu
}
