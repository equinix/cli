## equinix fabricv4 internet-access-services patch-eia-service

Execute patch-eia-service operation

### Synopsis

Execute the patch-eia-service operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 internet-access-services patch-eia-service [flags]
```

### Options

```
  -h, --help                                            help for patch-eia-service
      --internet-access-patch-operation-update string   internet-access-patch-operation-update field (JSON or string)
      --param-1 string                                  param-1 (required)
      --request string                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-access-services](equinix_fabricv4_internet-access-services.md)	 - Manage internet-access-services resources

