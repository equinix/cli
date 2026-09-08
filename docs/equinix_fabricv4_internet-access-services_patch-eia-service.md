## equinix fabricv4 internet-access-services patch-eia-service

Patch Internet Access Service by UUID

### Synopsis

Patch Internet Access Service by UUID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 internet-access-services patch-eia-service [flags]
```

### Options

```
  -h, --help                                            help for patch-eia-service
      --internet-access-patch-operation-update string   internet-access-patch-operation-update field (JSON or string)
      --request string                                  JSON payload for additional optional fields not exposed as flags
      --uuid string                                     UUID of the EIA Service (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-access-services](equinix_fabricv4_internet-access-services.md)	 - Manage internet-access-services resources

