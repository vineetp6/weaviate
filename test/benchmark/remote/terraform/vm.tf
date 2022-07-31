resource "google_compute_instance" "default" {
  name         = "automated-loadtest"
  machine_type = "c2-standard-16"
  zone         = "us-central1-a"

  tags = ["automated-loadtest"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
      size = 100
    }
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral public IP
    }
  }

  metadata_startup_script = "echo hi > /test.txt"
}
