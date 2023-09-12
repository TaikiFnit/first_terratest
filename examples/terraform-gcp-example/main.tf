terraform {
  required_version = ">= 0.12.26"
}

resource "google_compute_instance" "example" {
  project      = var.gcp_project_id
  name         = var.instance_name
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    network = "default"
    acces_config {

    }
  }
}

resource "google_storage_bucket" "example_bucket" {
  project  = var.gcp_project_id
  name     = var.bucket_name
  location = var.bucket_location
}
