## equinix fabricv4 ip-blocks patch-ip-block-by-id

patch IP Block by UUID

### Synopsis

patch IP Block by UUID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ip-blocks patch-ip-block-by-id [flags]
```

### Options

```
  -h, --help                                      help for patch-ip-block-by-id
      --patch-ip-block-request-body-item string   patch-ip-block-request-body-item field (JSON or string)
      --request string                            JSON payload for additional optional fields not exposed as flags
      --uuid string                               UUID of the IP Block (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ip-blocks](equinix_fabricv4_ip-blocks.md)	 - Manage ip-blocks resources

