## equinix fabricv4 application-subscriptions create-app-subscription

Create App Subscription

### Synopsis

This API provides capability to create user's App Subscription

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-subscriptions create-app-subscription [flags]
```

### Options

```
      --app-subscription-post-request-additional-properties string           app-subscription-post-request-additional-properties (JSON)
      --app-subscription-post-request-project-additional-properties string   app-subscription-post-request-project-additional-properties (JSON)
      --app-subscription-post-request-project-project-id string              Subscriber-assigned project ID
      --app-subscription-post-request-source-additional-properties string    app-subscription-post-request-source-additional-properties (JSON)
      --app-subscription-post-request-source-app-link string                 app-subscription-post-request-source-app-link (JSON)
      --app-subscription-post-request-source-ip-subnets string               List of IP subnets in CIDR notation (JSON array)
      --app-subscription-post-request-target-additional-properties string    app-subscription-post-request-target-additional-properties (JSON)
      --app-subscription-post-request-target-app-service string              app-subscription-post-request-target-app-service (JSON)
      --app-subscription-post-request-target-geo-scope string                Geo scope
      --app-subscription-post-request-target-prioritization string           app-subscription-post-request-target-prioritization
      --app-subscription-post-request-type string                            app-subscription-post-request-type
  -h, --help                                                                 help for create-app-subscription
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-subscriptions](equinix_fabricv4_application-subscriptions.md)	 - Manage application-subscriptions resources

