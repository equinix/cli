## equinix fabricv4 streams get-stream-device

Execute get-stream-device operation

### Synopsis

Execute the get-stream-device operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 streams get-stream-device [flags]
```

### Options

```
  -h, --help               help for get-stream-device
      --request string     JSON payload for request body fields
      --stream-id string   stream-id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

