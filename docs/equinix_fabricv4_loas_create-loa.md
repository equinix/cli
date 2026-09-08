## equinix fabricv4 loas create-loa

Create Loa

### Synopsis

The API provides capability to create a new Loa

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas create-loa [flags]
```

### Options

```
      --create-loa-issue-loa-additional-properties string       create-loa-issue-loa-additional-properties (JSON)
      --create-loa-issue-loa-authorized-product-type string     create-loa-issue-loa-authorized-product-type
      --create-loa-issue-loa-demarcation-point string           create-loa-issue-loa-demarcation-point (JSON)
      --create-loa-issue-loa-description string                 Additional context about this LOA.
      --create-loa-issue-loa-expiration-date-time string        create-loa-issue-loa-expiration-date-time (JSON)
      --create-loa-issue-loa-name string                        A short, descriptive name for this LOA.
      --create-loa-issue-loa-requestor string                   create-loa-issue-loa-requestor (JSON)
      --create-loa-issue-loa-type string                        create-loa-issue-loa-type
      --create-loa-request-loa-additional-properties string     create-loa-request-loa-additional-properties (JSON)
      --create-loa-request-loa-authorized-product-type string   create-loa-request-loa-authorized-product-type
      --create-loa-request-loa-description string               Additional context about this LOA.
      --create-loa-request-loa-issuer string                    create-loa-request-loa-issuer (JSON)
      --create-loa-request-loa-location string                  create-loa-request-loa-location (JSON)
      --create-loa-request-loa-name string                      A short, descriptive name for this LOA.
      --create-loa-request-loa-type string                      create-loa-request-loa-type
  -h, --help                                                    help for create-loa
      --request string                                          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

