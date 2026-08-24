## equinix fabricv4 application-services update-app-service-by-uuid

Update App Service

### Synopsis

This API provides capability to update user's App Service

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-services update-app-service-by-uuid [flags]
```

### Options

```
      --app-service-change-operation string   app-service-change-operation field (JSON or string)
      --app-service-id string                 App Service UUID (required)
  -h, --help                                  help for update-app-service-by-uuid
      --request string                        JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-services](equinix_fabricv4_application-services.md)	 - Manage application-services resources

