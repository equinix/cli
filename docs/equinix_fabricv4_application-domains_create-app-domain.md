## equinix fabricv4 application-domains create-app-domain

Create App Domain

### Synopsis

This API provides capability to create user's App Domain

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-domains create-app-domain [flags]
```

### Options

```
      --app-domain-post-request-additional-properties string           app-domain-post-request-additional-properties (JSON)
      --app-domain-post-request-description string                     Customer-provided App Domain description
      --app-domain-post-request-name string                            Customer-provided App Domain name
      --app-domain-post-request-project-additional-properties string   app-domain-post-request-project-additional-properties (JSON)
      --app-domain-post-request-project-project-id string              Subscriber-assigned project ID
      --app-domain-post-request-type string                            app-domain-post-request-type
  -h, --help                                                           help for create-app-domain
      --request string                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-domains](equinix_fabricv4_application-domains.md)	 - Manage application-domains resources

