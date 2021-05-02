# vscale provider

```bash
export VSCALE_URL=https://api.vscale.io/v1
export VSCALE_TOKEN=aaa...zzz
```

## Example Usage

```hcl
terraform {
  required_version = ">= 0.15.0"
  required_providers {
    vscale = {
      source  = "sergsoloviev/vscale"
      version = "0.0.3"
    }
  }
}

provider "vscale" {}

```
