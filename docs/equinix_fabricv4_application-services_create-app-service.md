## equinix fabricv4 application-services create-app-service

Create App Service

### Synopsis

This API provides capability to create user's App Service

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-services create-app-service [flags]
```

### Options

```
      --app-service-post-request-additional-properties string           app-service-post-request-additional-properties (JSON)
      --app-service-post-request-description string                     Customer-provided App Service description
      --app-service-post-request-endpoint string                        Accessible endpoint through this service
      --app-service-post-request-name string                            Customer-provided App Service name
      --app-service-post-request-project-additional-properties string   app-service-post-request-project-additional-properties (JSON)
      --app-service-post-request-project-project-id string              Subscriber-assigned project ID
      --app-service-post-request-source-domains string                  List of source domains from where traffic is allowed (JSON array)
      --app-service-post-request-type string                            app-service-post-request-type
  -h, --help                                                            help for create-app-service
      --request string                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-services](equinix_fabricv4_application-services.md)	 - Manage application-services resources

