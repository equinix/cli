## equinix metalv1 firmware-sets get-organization-firmware-sets

Get Organization's Firmware Sets

### Synopsis

Returns all firmware sets associated with the organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 firmware-sets get-organization-firmware-sets [flags]
```

### Options

```
  -h, --help             help for get-organization-firmware-sets
      --id string        Organization UUID (required)
      --page int         page field
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 firmware-sets](equinix_metalv1_firmware-sets.md)	 - Manage firmware-sets resources

