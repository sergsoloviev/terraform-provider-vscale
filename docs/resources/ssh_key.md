# Resource `vscale_ssh_key`

## Example Usage

```hcl
resource "vscale_ssh_key" "key" {
  name = "key0"
  key  = "ssh-rsa AAA...xYz user@comp.local"
}
```

## Argument Reference

- **name** (String) Key name
- **key** (String) Public SSH key
