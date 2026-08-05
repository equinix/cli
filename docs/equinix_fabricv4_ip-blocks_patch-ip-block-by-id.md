## equinix fabricv4 ip-blocks patch-ip-block-by-id

Execute patch-ip-block-by-id operation

### Synopsis

Execute the patch-ip-block-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 ip-blocks patch-ip-block-by-id [flags]
```

### Options

```
  -h, --help                                      help for patch-ip-block-by-id
      --id string                                 id (required)
      --patch-ip-block-request-body-item string   patch-ip-block-request-body-item field (JSON or string)
      --request string                            JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ip-blocks](equinix_fabricv4_ip-blocks.md)	 - Manage ip-blocks resources

