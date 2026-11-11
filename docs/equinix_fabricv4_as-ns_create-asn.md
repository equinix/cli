## equinix fabricv4 as-ns create-asn

Execute create-asn operation

### Synopsis

Execute the create-asn operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 as-ns create-asn [flags]
```

### Options

```
      --create-asn-request-additional-properties string                               create-asn-request-additional-properties (JSON)
      --create-asn-request-location-additional-properties string                      create-asn-request-location-additional-properties (JSON)
      --create-asn-request-location-ibx string                                        Deprecated
      --create-asn-request-location-metro-code string                                 create-asn-request-location-metro-code
      --create-asn-request-location-metro-href string                                 create-asn-request-location-metro-href
      --create-asn-request-location-metro-name string                                 create-asn-request-location-metro-name
      --create-asn-request-location-region string                                     create-asn-request-location-region
      --create-asn-request-number int                                                 create-asn-request-number
      --create-asn-request-ownership-challenge-details-additional-properties string   create-asn-request-ownership-challenge-details-additional-properties (JSON)
      --create-asn-request-ownership-challenge-details-type string                    create-asn-request-ownership-challenge-details-type
      --create-asn-request-project-additional-properties string                       create-asn-request-project-additional-properties (JSON)
      --create-asn-request-project-project-id string                                  Subscriber-assigned project ID
      --create-asn-request-range string                                               create-asn-request-range
      --create-asn-request-rir-region string                                          create-asn-request-rir-region
  -h, --help                                                                          help for create-asn
      --request string                                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 as-ns](equinix_fabricv4_as-ns.md)	 - Manage as-ns resources

