## equinix fabricv4 cross-connects create-diverse-cross-connect

Execute create-diverse-cross-connect operation

### Synopsis

Execute the create-diverse-cross-connect operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects create-diverse-cross-connect [flags]
```

### Options

```
      --diverse-cross-connect-request-additional-properties string   diverse-cross-connect-request-additional-properties (JSON)
      --diverse-cross-connect-request-data string                    diverse-cross-connect-request-data (JSON array)
  -h, --help[=true|false]                                            help for create-diverse-cross-connect
      --request string                                               JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

