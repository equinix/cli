## equinix fabricv4 application-links update-app-service-attachment-to-app-link

Update App Service attachment to App Link

### Synopsis

This API provides ability to update the App Service attachment to App Link.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links update-app-service-attachment-to-app-link [flags]
```

### Options

```
      --app-link-app-service-attachment-change-operation string   app-link-app-service-attachment-change-operation field (JSON or string)
      --app-link-id string                                        App Link UUID (required)
      --app-service-id string                                     App Service UUID (required)
  -h, --help                                                      help for update-app-service-attachment-to-app-link
      --request string                                            JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

