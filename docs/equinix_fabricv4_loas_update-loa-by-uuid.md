## equinix fabricv4 loas update-loa-by-uuid

Update Loa

### Synopsis

The API provides capability to update Loa details by Loa ID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas update-loa-by-uuid [flags]
```

### Options

```
  -h, --help                           help for update-loa-by-uuid
      --loa-id string                  Loa UUID (required)
      --loa-replace-operation string   loa-replace-operation field (JSON or string)
      --request string                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

