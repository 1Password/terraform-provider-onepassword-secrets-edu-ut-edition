---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "onepassword-secrets-edu-ut-edition_reference Data Source - onepassword-secrets-edu-ut-edition"
subcategory: ""
description: |-
  Reads a 1Password secret from a vault.
---

# onepassword-secrets-edu-ut-edition_reference (Data Source)

Reads a 1Password secret from a vault.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `field` (String) The name of the field of the secret to retrieve. Usually password.
- `item` (String) The name of the item to retrieve e.x Netflix.
- `vault` (String) The name of the vault from which the item is fetched. Must be present in 1Password account.

### Read-Only

- `secret` (String) The secret of the field item in the vault. The secret value is stored here.


