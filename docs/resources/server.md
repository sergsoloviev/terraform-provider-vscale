# Resource `vscale_server`

## Example Usage

```hcl
resource "vscale_server" "ubuntu1604" {
  name     = "test1"
  keys     = ["key0"]
  image    = "ubuntu_16.04_64_001_master"
  location = "msk0"
  rplan    = "small"
  count    = 0
}
```

## Argument Reference

- **name** (String) Server name
- **keys** (List<String>) List of SSH key names
- **image** (String) Operation system image
- **location** (String) Data Center id (msk0, spb0)
- **rplan** (String) Tariff plan (small, medium, large, huge, monster)
- **count** (Int, Optional) Terraform meta-argument

#### Images list:
    "ubuntu_14.04_64_001_preseed",
    "ubuntu_16.04_64_001_master",
    "ubuntu_18.04_64_001_master",
    "ubuntu_20.04_64_001_master",
    "ubuntu_20.04_64_001_ajenti",
    "ubuntu_20.04_64_001_docker",
    "ubuntu_20.04_64_001_bitrix",
    "ubuntu_20.04_64_001_mongodb",
    "ubuntu_20.04_64_001_lamp",
    "ubuntu_20.04_64_001_nodejs",
    "ubuntu_20.04_64_001_gitlab",
    "ubuntu_20.04_64_001_gogs",
    "ubuntu_20.04_64_001_wordpress",
    "ubuntu_20.04_64_001_django",
    "ubuntu_20.04_64_001_redmine",
    "ubuntu_20.04_64_001_redis",
    "ubuntu_20.04_64_001_jenkins",

    "debian_8_64_001_master",
    "debian_9_64_001_master",
    "debian_10_64_001_master",
    "debian_10_64_001_fastpanel",

    "centos_7_64_001_master",
    "CentOS_8_64_001_master",

    "Fedora_27_64_001_master",
    "Fedora_28_64_001_master",
    "Fedora_29_64_001_master",
    "Fedora_30_64_001_master",
    "Fedora_31_64_001_master",
    "Fedora_32_64_001_master"
