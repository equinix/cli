## equinix fabricv4 application-services get-attached-app-links-by-app-service-id

Get attached App Links for App Service

### Synopsis

This API provides capability to retrieve App Links attached to an App Service.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-services get-attached-app-links-by-app-service-id [flags]
```

### Options

```
      --app-service-id string   App Service UUID (required)
  -h, --help                    help for get-attached-app-links-by-app-service-id
      --limit int               limit field
      --offset int              offset field
      --order string            order field
      --request string          JSON payload for additional optional fields not exposed as flags
      --state string            state field (JSON or string)
      --style string            style field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-services](equinix_fabricv4_application-services.md)	 - Manage application-services resources

