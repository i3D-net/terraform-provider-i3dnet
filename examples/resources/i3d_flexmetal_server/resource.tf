# Create a Flexmetal Server
resource "i3d_flexmetal_server" "my-server" {
  name          = "TerraFlex-Server"
  location      = "EU: Rotterdam"
  instance_type = "bm7.std.8"
  os = {
    slug = "ubuntu-2404-lts"
    kernel_params = [
      {
        key   = "KEY_A"
        value = "VALUE_A"
      }
    ]
  }
  ssh_key             = ["<YOUR-PUBLIC-SSH-KEY>"]
  post_install_script = "#!/bin/bash\necho \"Hi TerraFlex there!\" > /root/output.txt"
}

# Create a partitioned Flexmetal Server
resource "i3d_flexmetal_server" "my-partitioned-server" {
  name          = "TerraFlex-Server"
  location      = "EU: Rotterdam"
  instance_type = "bm7.std.8"
  os = {
    slug = "ubuntu-2404-lts"
    kernel_params = [
      {
        key   = "KEY_A"
        value = "VALUE_A"
      }
    ]
    partitions = [
      {
        "target" : "/boot",
        "filesystem" : "ext2",
        "size" : 4096
      },
      {
        "target" : "/",
        "filesystem" : "ext4",
        "size" : -1
      },
      {
        "target"     = "/custom",
        "filesystem" = "ext4",
        "size"       = 10240
      }
    ]
  }
  ssh_key             = ["<YOUR-PUBLIC-SSH-KEY>"]
  post_install_script = "#!/bin/bash\necho \"Hi TerraFlex there!\" > /root/output.txt"
}