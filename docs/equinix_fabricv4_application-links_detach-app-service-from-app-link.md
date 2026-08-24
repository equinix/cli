## equinix fabricv4 application-links detach-app-service-from-app-link

Detach App Service from App Link

### Synopsis

This API provides ability to detach App Service from App Link

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links detach-app-service-from-app-link [flags]
```

### Options

```
      --app-link-id string      App Link UUID (required)
      --app-service-id string   App Service UUID (required)
  -h, --help                    help for detach-app-service-from-app-link
      --request string          JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

