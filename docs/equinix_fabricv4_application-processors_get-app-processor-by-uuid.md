## equinix fabricv4 application-processors get-app-processor-by-uuid

Execute get-app-processor-by-uuid operation

### Synopsis

Execute the get-app-processor-by-uuid operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 application-processors get-app-processor-by-uuid [flags]
```

### Options

```
  -h, --help[=true|false]   help for get-app-processor-by-uuid
      --request string      JSON payload for request body fields
      --uuid string         uuid (required)
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-processors](equinix_fabricv4_application-processors.md)	 - Manage application-processors resources

