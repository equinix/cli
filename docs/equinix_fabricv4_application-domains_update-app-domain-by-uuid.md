## equinix fabricv4 application-domains update-app-domain-by-uuid

Update App Domain

### Synopsis

This API provides capability to update user's App Domain

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-domains update-app-domain-by-uuid [flags]
```

### Options

```
      --app-domain-change-operation string   app-domain-change-operation field (JSON or string)
      --app-domain-id string                 App Domain UUID (required)
  -h, --help                                 help for update-app-domain-by-uuid
      --request string                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-domains](equinix_fabricv4_application-domains.md)	 - Manage application-domains resources

