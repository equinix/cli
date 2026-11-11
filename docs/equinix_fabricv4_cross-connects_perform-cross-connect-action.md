## equinix fabricv4 cross-connects perform-cross-connect-action

Execute perform-cross-connect-action operation

### Synopsis

Execute the perform-cross-connect-action operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects perform-cross-connect-action [flags]
```

### Options

```
      --cross-connect-action-request-additional-properties string                      cross-connect-action-request-additional-properties (JSON)
      --cross-connect-action-request-data-additional-properties string                 cross-connect-action-request-data-additional-properties (JSON)
      --cross-connect-action-request-data-attachments string                           cross-connect-action-request-data-attachments (JSON array)
      --cross-connect-action-request-data-live-traffic-consent-provided[=true|false]   cross-connect-action-request-data-live-traffic-consent-provided
      --cross-connect-action-request-data-notes string                                 cross-connect-action-request-data-notes (JSON array)
      --cross-connect-action-request-data-notifications string                         cross-connect-action-request-data-notifications (JSON array)
      --cross-connect-action-request-data-order string                                 cross-connect-action-request-data-order (JSON)
      --cross-connect-action-request-data-schedule string                              cross-connect-action-request-data-schedule (JSON)
      --cross-connect-action-request-data-services string                              cross-connect-action-request-data-services (JSON array)
      --cross-connect-action-request-type string                                       cross-connect-action-request-type
  -h, --help[=true|false]                                                              help for perform-cross-connect-action
      --param-1 string                                                                 param-1 (required)
      --request string                                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

