- tutorial: https://terratest.gruntwork.io/docs/getting-started/quick-start/

:memo: 初回実行

```
% cd test
% go mod init "github.com/TaikiFnit/first_terratest"
% go mod tidy
% go test -v
=== RUN   TestTerraformBasicExample
=== PAUSE TestTerraformBasicExample
=== CONT  TestTerraformBasicExample
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 retry.go:91: terraform [init -upgrade=false -no-color]
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66: Running command terraform with args [init -upgrade=false -no-color]
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66: Initializing the backend...
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66: Initializing provider plugins...
TestTerraformBasicExample 2023-08-29T12:41:14+09:00 logger.go:66: - Finding latest version of hashicorp/local...
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: - Installing hashicorp/local v2.4.0...
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: - Installed hashicorp/local v2.4.0 (signed by HashiCorp)
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: Terraform has created a lock file .terraform.lock.hcl to record the provider
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: selections it made above. Include this file in your version control repository
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: so that Terraform can guarantee to make the same selections by default when
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: you run "terraform init" in the future.
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: Terraform has been successfully initialized!
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: You may now begin working with Terraform. Try running "terraform plan" to see
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: any changes that are required for your infrastructure. All Terraform commands
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: should now work.
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: If you ever set or change modules or backend configuration for Terraform,
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: rerun this command to reinitialize your working directory. If you forget, other
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: commands will detect it and remind you to do so if necessary.
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 retry.go:91: terraform [apply -input=false -auto-approve -var example=test -var example_list=["test"] -var example_map={"expected" = "test"} -var-file varfile.tfvars -no-color -lock=false]
TestTerraformBasicExample 2023-08-29T12:41:15+09:00 logger.go:66: Running command terraform with args [apply -input=false -auto-approve -var example=test -var example_list=["test"] -var example_map={"expected" = "test"} -var-file varfile.tfvars -no-color -lock=false]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Terraform used the selected providers to generate the following execution
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: plan. Resource actions are indicated with the following symbols:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + create
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Terraform will perform the following actions:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   # local_file.example will be created
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + resource "local_file" "example" {
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content              = "test + test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_base64sha256 = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_base64sha512 = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_md5          = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha1         = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha256       = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha512       = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + directory_permission = "0777"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + file_permission      = "0777"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + filename             = "example.txt"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + id                   = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:     }
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   # local_file.example2 will be created
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + resource "local_file" "example2" {
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content              = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_base64sha256 = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_base64sha512 = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_md5          = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha1         = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha256       = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + content_sha512       = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + directory_permission = "0777"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + file_permission      = "0777"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + filename             = "example2.txt"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + id                   = (known after apply)
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:     }
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Plan: 2 to add, 0 to change, 0 to destroy.
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Changes to Outputs:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + example      = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + example2     = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + example_list = [
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + "test",
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:     ]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   + example_map  = {
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:       + expected = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:     }
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: local_file.example: Creating...
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: local_file.example2: Creating...
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: local_file.example2: Creation complete after 0s [id=a94a8fe5ccb19ba61c4c0873d391e987982fbbd3]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: local_file.example: Creation complete after 0s [id=91dce2a64ab5d2abaeb58e8acbbdc71257ef5e3a]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Outputs:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: example = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: example2 = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: example_list = tolist([
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   "test",
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: ])
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: example_map = tomap({
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66:   "expected" = "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: })
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 retry.go:91: terraform [output -no-color -json example]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Running command terraform with args [output -no-color -json example]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 retry.go:91: terraform [output -no-color -json example2]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Running command terraform with args [output -no-color -json example2]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: "test"
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 retry.go:91: terraform [output -no-color -json example_list]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Running command terraform with args [output -no-color -json example_list]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: ["test"]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 retry.go:91: terraform [output -no-color -json example_map]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Running command terraform with args [output -no-color -json example_map]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: {"expected":"test"}
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 retry.go:91: terraform [destroy -auto-approve -input=false -var example=test -var example_list=["test"] -var example_map={"expected" = "test"} -var-file varfile.tfvars -no-color -lock=false]
TestTerraformBasicExample 2023-08-29T12:41:16+09:00 logger.go:66: Running command terraform with args [destroy -auto-approve -input=false -var example=test -var example_list=["test"] -var example_map={"expected" = "test"} -var-file varfile.tfvars -no-color -lock=false]
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example: Refreshing state... [id=91dce2a64ab5d2abaeb58e8acbbdc71257ef5e3a]
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example2: Refreshing state... [id=a94a8fe5ccb19ba61c4c0873d391e987982fbbd3]
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: Terraform used the selected providers to generate the following execution
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: plan. Resource actions are indicated with the following symbols:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - destroy
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: Terraform will perform the following actions:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   # local_file.example will be destroyed
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - resource "local_file" "example" {
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content              = "test + test" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_base64sha256 = "sqFbsGIbVxFDk8DU0kmoXCyFUMZLSa4hUvDVpoR5unQ=" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_base64sha512 = "qjgeGmwsluil0zjtevOqcsin5S54uevb5Ac1L7qbVgCi0QUIxD8OViCa6fOsE8T+olVxQpdcuuSO/P7rlzzsyQ==" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_md5          = "dd0b4f04c1757f805ee462a83ac5e57b" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha1         = "91dce2a64ab5d2abaeb58e8acbbdc71257ef5e3a" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha256       = "b2a15bb0621b57114393c0d4d249a85c2c8550c64b49ae2152f0d5a68479ba74" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha512       = "aa381e1a6c2c96e8a5d338ed7af3aa72c8a7e52e78b9ebdbe407352fba9b5600a2d10508c43f0e56209ae9f3ac13c4fea2557142975cbae48efcfeeb973cecc9" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - directory_permission = "0777" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - file_permission      = "0777" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - filename             = "example.txt" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - id                   = "91dce2a64ab5d2abaeb58e8acbbdc71257ef5e3a" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:     }
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   # local_file.example2 will be destroyed
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - resource "local_file" "example2" {
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content              = "test" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_base64sha256 = "n4bQgYhMfWWaL+qgxVrQFaO/TxsrC4Is0V1sFbDwCgg=" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_base64sha512 = "7iaw3Ur350mqGo7jwQrpkj9hiYB3Lkc/iBml1JQODbJ6wYX4oOHV+E+IvIh/1nsUNzLDBMxfqa2Ob1f1ACio/w==" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_md5          = "098f6bcd4621d373cade4e832627b4f6" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha1         = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha256       = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - content_sha512       = "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - directory_permission = "0777" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - file_permission      = "0777" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - filename             = "example2.txt" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - id                   = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:     }
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: Plan: 0 to add, 0 to change, 2 to destroy.
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: Changes to Outputs:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - example      = "test" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - example2     = "test" -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - example_list = [
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - "test",
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:     ] -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:   - example_map  = {
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:       - expected = "test"
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:     } -> null
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example2: Destroying... [id=a94a8fe5ccb19ba61c4c0873d391e987982fbbd3]
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example: Destroying... [id=91dce2a64ab5d2abaeb58e8acbbdc71257ef5e3a]
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example: Destruction complete after 0s
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: local_file.example2: Destruction complete after 0s
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66:
TestTerraformBasicExample 2023-08-29T12:41:17+09:00 logger.go:66: Destroy complete! Resources: 2 destroyed.
--- PASS: TestTerraformBasicExample (2.98s)
PASS
ok  	github.com/TaikiFnit/first_terratest	3.292s
```

すでにapply済みのリソースをそのままのstateでテストすると再apply&destroy

```
TestTerraformBasicExample 2023-08-29T12:51:57+09:00 logger.go:66: Plan: 2 to add, 0 to change, 2 to destroy.
TestTerraformBasicExample 2023-08-29T12:51:57+09:00 logger.go:66: Apply complete! Resources: 2 added, 0 changed, 2 destroyed.
...
TestTerraformBasicExample 2023-08-29T12:51:58+09:00 logger.go:66: Plan: 0 to add, 0 to change, 2 to destroy.
TestTerraformBasicExample 2023-08-29T12:51:58+09:00 logger.go:66: Destroy complete! Resources: 2 destroyed.
```
