## equinix fabricv4 service-profiles search-service-profiles

Profile Search

### Synopsis

Search service profiles by search criteria

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles search-service-profiles [flags]
```

### Options

```
  -h, --help                                                                             help for search-service-profiles
      --request string                                                                   JSON payload for additional optional fields not exposed as flags
      --service-profile-search-request-additional-properties string                      service-profile-search-request-additional-properties (required) (JSON)
      --service-profile-search-request-filter-service-profile-and-filter string          service-profile-search-request-filter-service-profile-and-filter (JSON)
      --service-profile-search-request-filter-service-profile-simple-expression string   service-profile-search-request-filter-service-profile-simple-expression (JSON)
      --service-profile-search-request-pagination-additional-properties string           service-profile-search-request-pagination-additional-properties (JSON)
      --service-profile-search-request-pagination-limit int                              Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --service-profile-search-request-pagination-offset int                             Index of the first element.
      --service-profile-search-request-sort string                                       service-profile-search-request-sort (JSON array)
      --view-point string                                                                view-point field (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

