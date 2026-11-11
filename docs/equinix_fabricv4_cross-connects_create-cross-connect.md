## equinix fabricv4 cross-connects create-cross-connect

Execute create-cross-connect operation

### Synopsis

Execute the create-cross-connect operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects create-cross-connect [flags]
```

### Options

```
      --create-cross-connect-a-side-additional-properties string              create-cross-connect-a-side-additional-properties (JSON)
      --create-cross-connect-a-side-demarcation-point string                  create-cross-connect-a-side-demarcation-point (JSON)
      --create-cross-connect-additional-properties string                     create-cross-connect-additional-properties (JSON)
      --create-cross-connect-attachments string                               create-cross-connect-attachments (JSON array)
      --create-cross-connect-bmmr-type string                                 create-cross-connect-bmmr-type
      --create-cross-connect-circuit-id string                                create-cross-connect-circuit-id
      --create-cross-connect-name string                                      create-cross-connect-name
      --create-cross-connect-notes string                                     create-cross-connect-notes (JSON array)
      --create-cross-connect-notifications string                             create-cross-connect-notifications (JSON array)
      --create-cross-connect-order-additional-properties string               create-cross-connect-order-additional-properties (JSON)
      --create-cross-connect-order-customer-reference-id string               create-cross-connect-order-customer-reference-id
      --create-cross-connect-order-order-line-item-number string              create-cross-connect-order-order-line-item-number
      --create-cross-connect-order-order-number string                        create-cross-connect-order-order-number
      --create-cross-connect-order-purchase-order string                      create-cross-connect-order-purchase-order (JSON)
      --create-cross-connect-redundancy-additional-properties string          create-cross-connect-redundancy-additional-properties (JSON)
      --create-cross-connect-redundancy-group string                          create-cross-connect-redundancy-group
      --create-cross-connect-redundancy-priority string                       create-cross-connect-redundancy-priority
      --create-cross-connect-schedule-additional-properties string            create-cross-connect-schedule-additional-properties (JSON)
      --create-cross-connect-schedule-auto-convert-to-standard[=true|false]   create-cross-connect-schedule-auto-convert-to-standard
      --create-cross-connect-schedule-code string                             create-cross-connect-schedule-code
      --create-cross-connect-schedule-readiness-date-time string              create-cross-connect-schedule-readiness-date-time (JSON)
      --create-cross-connect-schedule-requested-date-time string              create-cross-connect-schedule-requested-date-time (JSON)
      --create-cross-connect-schedule-type string                             create-cross-connect-schedule-type
      --create-cross-connect-services string                                  create-cross-connect-services (JSON array)
      --create-cross-connect-type string                                      create-cross-connect-type
      --create-cross-connect-z-side-additional-properties string              create-cross-connect-z-side-additional-properties (JSON)
      --create-cross-connect-z-side-demarcation-point string                  create-cross-connect-z-side-demarcation-point (JSON)
      --create-cross-connect-z-side-loa string                                create-cross-connect-z-side-loa (JSON)
      --create-cross-connect-z-side-location string                           create-cross-connect-z-side-location (JSON)
  -h, --help[=true|false]                                                     help for create-cross-connect
      --request string                                                        JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

