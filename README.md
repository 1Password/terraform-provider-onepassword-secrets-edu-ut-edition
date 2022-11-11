# Terraform Provider Hashicups

This repo is a companion repo to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers) Learn collection. 

In the collection, you will use the HashiCups provider as a bridge between Terraform and the HashiCups API. Then, extend Terraform by recreating the HashiCups provider. By the end of this collection, you will be able to take these intuitions to create your own custom Terraform provider. 

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-hashicups
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

## Development Requirements

### Windows Installation

**VIDEO GUIDE AND DEMO:** https://youtu.be/h2WviO3VA5E

#### Setup GO

1. Download the `.msi` file from https://go.dev/doc/install
2. Run the `.msi` file and follow the installer instructions.
3. Open `powershell` and type `go version`, if you did the above steps correctly, this will return a version number.
4. Open the environment variable menu, you should be able to access this menu by typing "edit env" in your start menu and selecting the first option that shows up. For any environment variable that does not already exist you can use the `New...` button and for any environment variable with the incorrect value you can update it with the `Edit...` button.
    - In user variables set `GOPATH` as the folder path where you installed go
        - Example: `C:\Programs\Go`
    - In user variables set `GOBIN` as the bin folder path fond in your go installation
        - Example: `C:\Programs\Go\bin`
    - In user variables add the bin folder path found in your go installation to the `path` variable
        - Example: `C:\Programs\Go\bin`
    - In user variables if you see something like `%USERPROFILE%\go\bin` in the `path` variable delete it. This is very important and may cause issues later if you don't do this.

5. Click "OK" in all the menus to apply changes and close the environment variable windows.
6. Ensure all the environment variables were set correctly by opening up `powershell` and checking that the following commands return correct paths:
    - `go env GOPATH`
    - `go env GOBIN`
    - `$env:path`

#### Setup Terraform

1. Download the `AMD64` file from https://developer.hashicorp.com/terraform/downloads (assuming you are using a 64-bit based machine).
2. Unzip the downloaded file.
3. Open the extracted folder and move the `terraform.exe` to a spot where you'll remember, for me I created a folder called `Terraform` where my other programs are located and put the `.exe` in there.
4. Open the environment variable menu and in user variables add the path to the folder where `terraform.exe` is located to the `Path` variables.
    - Example: `C:\Programs\Terraform`
5. Click "OK" in all the menus to apply changes and close the environment variable windows.

6. Open `powershell` and type `terraform --version`, if everything was done correctly this will return a version number.

#### Setup 1Password 8 Desktop Application

1. Download the Windows version of 1Password 8 from https://1password.com/downloads/windows/
2. Run the file and follow setup instructions.
3. When it is done, launch the 1Password application and login.

#### Setup 1Password CLI

1. Download the most recent `amd64` file for Windows from https://app-updates.agilebits.com/product_history/CLI2
2. Unzip the downloaded file.
3. Open the extracted folder and move `op.exe` to a spot where you'll remember, for me I created a folder called `1Password CLI` where my other programs are and put the `.exe` in there.
4. Open the environment variable menu and in user variables add the path to the folder where `op.exe` is located to the `Path` variables.
    - Example: `C:\Programs\1Password CLI`
5. Click "OK" in all the menus to apply changes and close the environment variable windows.
6. Open `powershell` and type `op --version`, if everything was done correctly this will return a version number.

#### Link 1Password 8 Desktop Application and 1Password CLI

1. Open 1Password
2. Settings -> Security -> Check "Unlock using Windows Hello" and "Show Windows Hello prompt automatically".
3. Settings -> Developer -> Check "Connect with 1Passworde CLI".
4. Open `powershell` and type `$Env:OP_BIOMETRIC_UNLOCK_ENABLED = "true"`

4. Then type `op vault list`, you should be prompted to login.
    - You may receive a warning saying you need to lock 1Password to setup Windows Hello, open the 1Password 8 Desktop App and click the 3 dots in the top left of the application and press "Lock 1Password". Then try `op vault list` again. It might ask you to login to the 1Password app again. Do so then try `op vault list` once more.
5. After logging in a list of your vaults will be displayed (if you have any).

#### Setup RC File

1. Press the Windows Key and the R Key together (Win + R). This will open your `AppData\Roaming` folder.

2. Create a file called `terraform.rc`, ensure this is a RC File and not a normal text file.

3. Copy and paste the following text into the file:

   ```
   provider_installation {
   
     dev_overrides {
         "hashicorp.com/edu/onepassword" = "<PATH>"
         }
   
     # For all other providers, install them directly from their origin provider
     # registries as normal. If you omit this, Terraform will _only_ use
     # the dev_overrides block, and so no other providers will be available.
     direct {}
   }
   ```

4. Open `powershell` and type `go env GOBIN` and copy the path that is returned.
5. Change `<PATH>` in this file to the path you copied above and change the backslashes to double backslash (`\` -> `\\`) in order to avoid escape character errors.

6. Save the file and close it.

### Linux Installation [Ubuntu Based 64-Bit System]

**VIDEO GUIDE AND DEMO:** https://youtu.be/CC-_tp_LMCI

#### Setup Go, Terraform, and RC File

1. Run `bash setup.sh` (you can find `setup.sh` in the `installation_and_demo` folder of our repo)
    - Do not run this script as `sudo`, it may not set the `.terraformrc` file correctly
2. The script will tell you if everything was setup properly except for the `.terraformrc` file and the 1Password 8 Desktop application which you must check yourself.
3. Run `cd ~` followed by `cat .terraformrc.`
4. Open a new terminal and run `go env GOPATH`, ensure that the path in the `.terraformrc` file is the same, except with a `/bin` appended to the end.
    - If it is not, simply make the changes and and save the file.
5. Search your programs for 1Password and open it
    - Follow the login instructions and sign in
6. Once logged in click the 3 dots in the top left and click settings
    - Go to Security and ensure "Unlock using system authentication service" is **NOT** checked.
    - Go to Developer and ensure "Connect with 1Password CLI" it **NOT** checked.

### Mac Installation

**VIDEO GUIDE AND DEMO:** You can find a video in the `installation_and_demo` folder.

#### Setup GO

1. Go to https://go.dev/doc/install and download the `.pkg` file for macOS
2. Run the file and follow instructions

#### Setup Terraform

1. Run the following commands in terminal in order

   `brew tap hashicorp/tap`

   `brew install hashicorp/tap/terraform`

#### Setup 1Password 8 Desktop Application

1. Download the macOS version of 1Password 8 from https://1password.com/downloads/mac/
2. Run the file and follow setup instructions.
3. When it is done, launch the 1Password application and login.

#### Setup 1Password CLI

1. Open a terminal and run the following commands in order

   `brew install --cask 1password/tap/1password-cli`

   `op --version`

#### Link 1Password 8 Desktop Application and 1Password CLI

1. Open 1Password
2. Settings -> Security -> Select Touch ID
3. Settings -> Developer -> Check "Connect with 1Passworde CLI".
4. Open terminal and type `OP_BIOMETRIC_UNLOCK_ENABLED=true`

4. Then type `op vault list`, you should be prompted to login.
5. After logging in a list of your vaults will be displayed (if you have any).

#### Setup RC File

1. Open terminal and `cd ~`

2. Create a file called `.terraformrc` here, ensure this is a RC File and not a normal text file.

3. Copy and paste the following text into the file:

   ```
   provider_installation {
   
     dev_overrides {
         "hashicorp.com/edu/onepassword" = "<PATH>"
         }
   
     # For all other providers, install them directly from their origin provider
     # registries as normal. If you omit this, Terraform will _only_ use
     # the dev_overrides block, and so no other providers will be available.
     direct {}
   }
   ```

4. Open terminal and type `go env GOBIN` and copy the path that is returned.
    - If no path is returned set GOBIN as the return of `go env GOPATH` with `/bin` appended to the end
    - This can be done using the following command with path being as described above.

      `export GOBIN=<PATH>`
5. Change `<PATH>` in this file to the path you copied above.

6. Save the file and close it.

## Instructions

#### Clone the repo

Run `git clone https://github.com/1Password/cse301-red-onepassword-terraform-provder.git` into any directory you want.

#### Functionality

This program incorporates two functionalities, READ and CREATE. Each of which use a terraform (`.tf`) file for configuration. We have included two example terraform files. The one for READ can be found at `cse301-red-onepassword-terraform-provder\examples\reference\main.tf` and the one for CREATE can be found at `cse301-red-onepassword-terraform-provder\examples\secret\main.tf`.

- READ

    - If you open the `main.tf` file for read you can see in `data`, there are three objects:

      `vault`: name of the vault you want to read

      `item`: name of the item you want to read

      `field`: name of the field you want to read

    - In order to read the value of a field you input the name of your vault, the name of your item, and the name of the field into the terraform file

- CREATE

    - If you open the `main.tf` file for create you can see in `resource`, there are 3 objects:

      `vault`: name of the vault you want to create.

      `item`: name of the secret you want to create.

      `password_recipe`: password for the secret you want to create.
    - In order to create a secret you input the name of your vault, the name of your item, and the specifications (character_set and length) of the password recipe.

#### READ Example

1. Open the 1Password 8 Desktop App
2. Create a new Vault named "test" if you don't have one already

3. Create a new Login item in the vault named "login"
    - Set the username and password to whatever you want

4. Navigate to `cse301-red-onepassword-terraform-provder\examples\reference\main.tf`

5. Set the `data` part of the terraform file to the following:

   ```
   data "onepassword_reference" "edu" {
     vault = "test"
     item  = "login"
     field = "password"
   }
   ```

6. Save the terraform file.

7. Navigate to the `cse301-red-onepassword-terraform-provder` folder
8. Open a terminal here and run `go install .`  followed by `go mod tidy`
9. Navigate to `cse301-red-onepassword-terraform-provder\examples\reference`
10. Open a terminal here and run `terraform plan`, the value of the `password` field will then be returned to you, showing a successful read.
    - **NOTE: On Linux you must run `eval $(op signin)` first**. You will be prompted for the following information
        - url: should just be `my.1password.com`
        - email: the email you used for 1password
        - secret key: you can find this in your profile on the 1password website when you login
        - password: the password you use to login to your account.

#### CREATE Example

1. Open the 1Password 8 Desktop App

2. Create a new Vault named "test" if you don't have one already

3. Navigate to `cse301-red-onepassword-terraform-provder\examples\secret\main.tf`

4. Set the `resource` part of the terraform file to the following:

   ```
   resource "onepassword_secret" "edu" {
     vault = "test"
     title = "uber4"
     password_recipe = {
       character_set = ["digits", "letters"]
       length        = 30
     }
   }
   ```

5. Save the terraform file.
6. Navigate to the `cse301-red-onepassword-terraform-provder` folder
7. Open a terminal here and run `go install .`  followed by `go mod tidy`
8. Navigate to `cse301-red-onepassword-terraform-provder\examples\secret`
9. Open a terminal here and run `terraform apply`, if it asks you to type `yes` do so and press enter.
    - **NOTE: On Linux you must run `eval $(op signin)` first**. You will be prompted for the following information
        - url: should just be `my.1password.com`
        - email: the email you used for 1password
        - secret key: you can find this in your profile on the 1password website when you login
        - password: the password you use to login to your account.
10. Open the 1Password 8 desktop and look inside your vault named "test", you should see the new secret.

