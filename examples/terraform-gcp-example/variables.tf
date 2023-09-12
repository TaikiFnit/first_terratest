variable "gcp_project_id" {}
variable "instance_name" {
  type    = string
  default = "terratest-example"
}

variable "machine_type" {
  type    = string
  default = "f1-micro"
}

variable "zone" {
  type    = string
  default = "us-central-a"
}

variable "bucket_name" {
  type    = string
  default = "terratest-example-taikifnit-bucket"
}

variable "bucket_location" {
  type    = string
  default = "US"
}
