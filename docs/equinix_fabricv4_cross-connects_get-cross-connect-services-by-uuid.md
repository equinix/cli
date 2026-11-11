## equinix fabricv4 cross-connects get-cross-connect-services-by-uuid

Execute get-cross-connect-services-by-uuid operation

### Synopsis

Execute the get-cross-connect-services-by-uuid operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects get-cross-connect-services-by-uuid [flags]
```

### Options

```
  -h, --help             help for get-cross-connect-services-by-uuid
      --request string   JSON payload for request body fields
      --uuid string      uuid (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

