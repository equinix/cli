## equinix fabricv4 loas get-loa-actions-by-uuid

Get Loa Action by Action ID

### Synopsis

This API provides capability to fetch action details

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas get-loa-actions-by-uuid [flags]
```

### Options

```
      --action-id string   Action UUID (required)
  -h, --help               help for get-loa-actions-by-uuid
      --loa-id string      Loa UUID (required)
      --request string     JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

