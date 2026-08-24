## equinix fabricv4 application-links create-app-link

Create App Link

### Synopsis

This API provides capability to create user's App Link

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links create-app-link [flags]
```

### Options

```
      --app-link-post-request-additional-properties string           app-link-post-request-additional-properties (JSON)
      --app-link-post-request-bandwidth int                          App Link aggregated data transfer capacity in Mbps
      --app-link-post-request-description string                     Customer-provided App Link description
      --app-link-post-request-ipv4-address string                    AppLink IP address
      --app-link-post-request-name string                            Customer-provided App Link name
      --app-link-post-request-project-additional-properties string   app-link-post-request-project-additional-properties (JSON)
      --app-link-post-request-project-project-id string              Subscriber-assigned project ID
      --app-link-post-request-router-additional-properties string    app-link-post-request-router-additional-properties (JSON)
      --app-link-post-request-router-uuid string                     Cloud Router UUID
      --app-link-post-request-type string                            app-link-post-request-type
  -h, --help                                                         help for create-app-link
      --request string                                               JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

