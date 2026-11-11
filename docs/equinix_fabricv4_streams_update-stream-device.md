## equinix fabricv4 streams update-stream-device

Execute update-stream-device operation

### Synopsis

Execute the update-stream-device operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 streams update-stream-device [flags]
```

### Options

```
  -h, --help[=true|false]                                        help for update-stream-device
      --request string                                           JSON payload for additional optional fields not exposed as flags
      --stream-device-put-request-additional-properties string   stream-device-put-request-additional-properties (JSON)
      --stream-device-put-request-interfaces string              stream-device-put-request-interfaces (JSON array)
      --stream-device-put-request-metrics-enabled[=true|false]   stream-device-put-request-metrics-enabled
      --stream-id string                                         stream-id (required)
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

