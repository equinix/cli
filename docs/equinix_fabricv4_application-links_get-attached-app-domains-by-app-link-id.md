## equinix fabricv4 application-links get-attached-app-domains-by-app-link-id

Get attached App Domains for App Link

### Synopsis

This API provides capability to retrieve App Domains attached to an App Link.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links get-attached-app-domains-by-app-link-id [flags]
```

### Options

```
      --app-link-id string         App Link UUID (required)
      --attachment-status string   attachment-status field (JSON or string)
  -h, --help                       help for get-attached-app-domains-by-app-link-id
      --limit int                  limit field
      --offset int                 offset field
      --order string               order field
      --request string             JSON payload for additional optional fields not exposed as flags
      --style string               style field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

