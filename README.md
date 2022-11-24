# Terraform Provider onepassword



## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-onepassword
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```
For a local provider add to the ```.terraformrc``` file the following line:
```
dev_overrides {
      "hashicorp.com/edu/onepassword" = "<GOBIN>"
  }
  ```
  Where <GOBIN> is the directory of the go bin (obtained from ```go env GOBIN```).
  
Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
