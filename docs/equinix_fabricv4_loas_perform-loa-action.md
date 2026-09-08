## equinix fabricv4 loas perform-loa-action

Loa Actions

### Synopsis

The API provides capability to perform actions on Loa

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas perform-loa-action [flags]
```

### Options

```
  -h, --help                                                   help for perform-loa-action
      --loa-action-request-additional-properties string        loa-action-request-additional-properties (JSON)
      --loa-action-request-data-additional-properties string   loa-action-request-data-additional-properties (JSON)
      --loa-action-request-data-demarcation-point string       loa-action-request-data-demarcation-point (JSON)
      --loa-action-request-data-expiration-date-time string    loa-action-request-data-expiration-date-time (JSON)
      --loa-action-request-data-portal-url string              Portal URL for the LOA to either accept from requestor <br> or authorize from the issuer.
      --loa-action-request-type string                         loa-action-request-type
      --loa-id string                                          Loa UUID (required)
      --request string                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

