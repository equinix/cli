## equinix fabricv4 application-links attach-app-service-to-app-link

Attach App Service to App Link

### Synopsis

This API provides ability to attach the user's App Service to App Link.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links attach-app-service-to-app-link [flags]
```

### Options

```
      --app-link-attach-service-request-additional-properties string   app-link-attach-service-request-additional-properties (JSON)
      --app-link-attach-service-request-destination-ip string          Target IP for forwarding API requests
      --app-link-attach-service-request-geo-scope string               Geo scope for the App Service
      --app-link-id string                                             App Link UUID (required)
      --app-service-id string                                          App Service UUID (required)
  -h, --help                                                           help for attach-app-service-to-app-link
      --request string                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

