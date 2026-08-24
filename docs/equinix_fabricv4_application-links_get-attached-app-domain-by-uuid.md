## equinix fabricv4 application-links get-attached-app-domain-by-uuid

Get attached App Domain for App Link

### Synopsis

This API provides ability to retrieve an App Domain attached to an App Link.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links get-attached-app-domain-by-uuid [flags]
```

### Options

```
      --app-domain-id string   App Domain UUID (required)
      --app-link-id string     App Link UUID (required)
  -h, --help                   help for get-attached-app-domain-by-uuid
      --request string         JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

