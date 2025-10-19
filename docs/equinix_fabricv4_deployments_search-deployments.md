## equinix fabricv4 deployments search-deployments

Search deployments

### Synopsis

The API provides capability to get list of user's deployment using search criteria.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 deployments search-deployments [flags]
```

### Options

```
      --deployment-search-request-additional-properties string          deployment-search-request-additional-properties (required) (JSON)
      --deployment-search-request-filter-additional-properties string   deployment-search-request-filter-additional-properties (JSON)
      --deployment-search-request-filter-and string                     deployment-search-request-filter-and (JSON array)
      --deployment-search-request-filter-or string                      deployment-search-request-filter-or (JSON array)
  -h, --help                                                            help for search-deployments
      --request string                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 deployments](equinix_fabricv4_deployments.md)	 - Manage deployments resources

