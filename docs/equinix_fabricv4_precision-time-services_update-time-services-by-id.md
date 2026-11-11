## equinix fabricv4 precision-time-services update-time-services-by-id

Execute update-time-services-by-id operation

### Synopsis

Execute the update-time-services-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 precision-time-services update-time-services-by-id [flags]
```

### Options

```
  -h, --help                                     help for update-time-services-by-id
      --id string                                id (required)
      --precision-time-change-operation string   precision-time-change-operation field (JSON or string)
      --request string                           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 precision-time-services](equinix_fabricv4_precision-time-services.md)	 - Manage precision-time-services resources

