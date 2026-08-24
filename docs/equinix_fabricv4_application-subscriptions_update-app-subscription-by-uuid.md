## equinix fabricv4 application-subscriptions update-app-subscription-by-uuid

Update App Subscription

### Synopsis

This API provides capability to update user's App Subscription

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-subscriptions update-app-subscription-by-uuid [flags]
```

### Options

```
      --app-subscription-change-operation string   app-subscription-change-operation field (JSON or string)
      --app-subscription-id string                 App Subscription UUID (required)
  -h, --help                                       help for update-app-subscription-by-uuid
      --request string                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-subscriptions](equinix_fabricv4_application-subscriptions.md)	 - Manage application-subscriptions resources

