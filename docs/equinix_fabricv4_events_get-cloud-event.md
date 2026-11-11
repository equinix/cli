## equinix fabricv4 events get-cloud-event

Execute get-cloud-event operation

### Synopsis

Execute the get-cloud-event operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 events get-cloud-event [flags]
```

### Options

```
  -h, --help[=true|false]   help for get-cloud-event
      --param-1 string      param-1 (required)
      --request string      JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 events](equinix_fabricv4_events.md)	 - Manage events resources

