terraform {
  required_providers {
    i3d = {
      source = "registry.terraform.io/i3D-net/i3d"
    }
  }
}

provider "i3d" {
  api_key  = "98brv7g34b9843rjfdde"
  base_url = "http://localhost:8081"
}