## equinix fabricv4 internet-exchange-services create-exchange-service-action

Execute create-exchange-service-action operation

### Synopsis

Execute the create-exchange-service-action operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 internet-exchange-services create-exchange-service-action [flags]
```

### Options

```
      --exchange-service-action-request-additional-properties string   exchange-service-action-request-additional-properties (JSON)
      --exchange-service-action-request-type string                    exchange-service-action-request-type
  -h, --help[=true|false]                                              help for create-exchange-service-action
      --param-1 string                                                 param-1 (required)
      --request string                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-exchange-services](equinix_fabricv4_internet-exchange-services.md)	 - Manage internet-exchange-services resources

