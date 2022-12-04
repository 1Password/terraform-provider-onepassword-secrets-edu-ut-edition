# setup go
sudo apt install golang-go

# verify go install
GOVERIFY=0
if go version ; then
  GOVERIFY=1
else
  GOVERIFY=0
fi

# setup terraform
sudo apt-get install unzip
wget https://releases.hashicorp.com/terraform/1.3.4/terraform_1.3.4_linux_amd64.zip
unzip terraform_1.3.4_linux_amd64.zip
sudo mv terraform /usr/local/bin
sudo rm terraform_1.3.4_linux_amd64.zip

# verify terraform installation
TERRAFORMVERIFY=0
if terraform --version ; then
  TERRAFORMVERIFY=1
else
  TERRAFORMVERIFY=0
fi

# override file
cd ~
GOPATH="$(go env GOPATH)"
sudo echo "provider_installation {
  dev_overrides {
      \"hashicorp.com/edu/onepassword\" = \"$GOPATH/bin\"
  }
  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}" > .terraformrc
# Download 1Password Dektop 8
wget https://downloads.1password.com/linux/debian/amd64/stable/1password-latest.deb
sudo dpkg -i 1password-latest.deb
sudo rm 1password-latest.deb

# Download 1Password CLI
wget https://downloads.1password.com/linux/debian/amd64/stable/1password-cli-amd64-latest.deb
sudo dpkg -i 1password-cli-amd64-latest.deb
sudo rm 1password-cli-amd64-latest.deb

# verify 1Password CLI installation
CLIVERIFY=0
if op --version ; then
  CLIVERIFY=1
else
  CLIVERIFY=0
fi

# verify all
if [ $GOVERIFY -eq 1 ]; then
  echo "GO SETUP PASSED"
else
  echo "GO SETUP FAILED"
fi

if [ $TERRAFORMVERIFY -eq 1 ]; then
  echo "TERRAFORM SETUP PASSED"
else
  echo "TERRAFORM SETUP FAILED"
fi

if [ $CLIVERIFY -eq 1 ]; then
  echo "1PASSWORD CLI SETUP PASSED"
else
  echo "1PASSWORD CLI SETUP FAILED"
fi

echo "YOU MUST MANUALLY CHECK IF .terraformrc FILE WAS SETUP PROPERLY, IT CAN BE FOUND IN YOUR ~ DIRECTORY (IT MAY BE HIDDEN)"
echo "YOU MUST MANUALLY CHECK IF 1PASSWORD WAS INSTALLED CORRECTLY"
