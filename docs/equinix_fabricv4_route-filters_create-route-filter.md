## equinix fabricv4 route-filters create-route-filter

Create Route Filters

### Synopsis

This API provides capability to create a Route Filter

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters create-route-filter [flags]
```

### Options

```
  -h, --help                                                      help for create-route-filter
      --request string                                            JSON payload for additional optional fields not exposed as flags
      --route-filters-base-additional-properties string           route-filters-base-additional-properties (JSON)
      --route-filters-base-description string                     Customer-provided connection description
      --route-filters-base-name string                            route-filters-base-name
      --route-filters-base-project-additional-properties string   route-filters-base-project-additional-properties (JSON)
      --route-filters-base-project-project-id string              Subscriber-assigned project ID
      --route-filters-base-type string                            route-filters-base-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

