## equinix fabricv4 application-services search-attached-app-subscriptions-by-app-service-id

Search attached App Subscriptions

### Synopsis

The API provides capability to get list of App Subscriptions attached to an App Service using search criteria.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-services search-attached-app-subscriptions-by-app-service-id [flags]
```

### Options

```
      --app-service-attached-app-subscription-search-request-additional-properties string              app-service-attached-app-subscription-search-request-additional-properties (JSON)
      --app-service-attached-app-subscription-search-request-filter-additional-properties string       app-service-attached-app-subscription-search-request-filter-additional-properties (JSON)
      --app-service-attached-app-subscription-search-request-filter-and string                         app-service-attached-app-subscription-search-request-filter-and (JSON array)
      --app-service-attached-app-subscription-search-request-pagination-additional-properties string   app-service-attached-app-subscription-search-request-pagination-additional-properties (JSON)
      --app-service-attached-app-subscription-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-service-attached-app-subscription-search-request-pagination-offset int                     Index of the first element.
      --app-service-attached-app-subscription-search-request-sort string                               app-service-attached-app-subscription-search-request-sort (JSON array)
      --app-service-id string                                                                          App Service UUID (required)
  -h, --help                                                                                           help for search-attached-app-subscriptions-by-app-service-id
      --request string                                                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-services](equinix_fabricv4_application-services.md)	 - Manage application-services resources

