## equinix fabricv4 internet-exchange-services update-exchange-service-by-id

Execute update-exchange-service-by-id operation

### Synopsis

Execute the update-exchange-service-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 internet-exchange-services update-exchange-service-by-id [flags]
```

### Options

```
      --exchange-service-replace-operation string   exchange-service-replace-operation field (JSON or string)
  -h, --help                                        help for update-exchange-service-by-id
      --id string                                   id (required)
      --request string                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-exchange-services](equinix_fabricv4_internet-exchange-services.md)	 - Manage internet-exchange-services resources

