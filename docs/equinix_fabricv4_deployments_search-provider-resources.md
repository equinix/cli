## equinix fabricv4 deployments search-provider-resources

Search provider resources

### Synopsis

The API provides capability to get list of user's provider resources using search criteria, including optional filtering

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 deployments search-provider-resources [flags]
```

### Options

```
  -h, --help                                                          help for search-provider-resources
      --provider-search-request-additional-properties string          provider-search-request-additional-properties (JSON)
      --provider-search-request-filter-additional-properties string   provider-search-request-filter-additional-properties (JSON)
      --provider-search-request-filter-and string                     provider-search-request-filter-and (JSON array)
      --provider-search-request-filter-or string                      provider-search-request-filter-or (JSON array)
      --request string                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 deployments](equinix_fabricv4_deployments.md)	 - Manage deployments resources

