## equinix fabricv4 connections search-connections

Search connections

### Synopsis

The API provides capability to get list of user's virtual connections using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections search-connections [flags]
```

### Options

```
  -h, --help                                                     help for search-connections
      --request string                                           JSON payload for additional optional fields not exposed as flags
      --search-request-additional-properties string              search-request-additional-properties (required) (JSON)
      --search-request-filter-additional-properties string       search-request-filter-additional-properties (JSON)
      --search-request-filter-and string                         search-request-filter-and (JSON array)
      --search-request-filter-operator string                    search-request-filter-operator
      --search-request-filter-or string                          search-request-filter-or (JSON array)
      --search-request-filter-property string                    search-request-filter-property
      --search-request-filter-values string                      search-request-filter-values (JSON array)
      --search-request-pagination-additional-properties string   search-request-pagination-additional-properties (JSON)
      --search-request-pagination-limit int                      search-request-pagination-limit
      --search-request-pagination-offset int                     search-request-pagination-offset
      --search-request-sort string                               search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

