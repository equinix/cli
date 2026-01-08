## equinix metalv1 firmware-sets get-project-firmware-sets

Get Project's Firmware Sets

### Synopsis

Returns all firmware sets associated with the project or organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 firmware-sets get-project-firmware-sets [flags]
```

### Options

```
  -h, --help             help for get-project-firmware-sets
      --id string        Project UUID (required)
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

