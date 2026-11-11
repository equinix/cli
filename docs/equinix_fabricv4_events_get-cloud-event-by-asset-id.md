## equinix fabricv4 events get-cloud-event-by-asset-id

Execute get-cloud-event-by-asset-id operation

### Synopsis

Execute the get-cloud-event-by-asset-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 events get-cloud-event-by-asset-id [flags]
```

### Options

```
      --cloud-event-asset-type string   cloud-event-asset-type (required)
  -h, --help[=true|false]               help for get-cloud-event-by-asset-id
      --limit int                       limit field
      --offset int                      offset field
      --param-2 string                  param-2 (required)
      --request string                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 events](equinix_fabricv4_events.md)	 - Manage events resources

