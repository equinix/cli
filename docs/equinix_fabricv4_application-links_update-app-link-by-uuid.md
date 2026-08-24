## equinix fabricv4 application-links update-app-link-by-uuid

Update App Link

### Synopsis

This API provides capability to update user's App Link

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links update-app-link-by-uuid [flags]
```

### Options

```
      --app-link-change-operation string   app-link-change-operation field (JSON or string)
      --app-link-id string                 App Link UUID (required)
  -h, --help                               help for update-app-link-by-uuid
      --request string                     JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

