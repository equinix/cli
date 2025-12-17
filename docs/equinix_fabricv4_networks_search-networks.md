## equinix fabricv4 networks search-networks

Search Network

### Synopsis

The API provides capability to get list of user's Fabric Network using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 networks search-networks [flags]
```

### Options

```
  -h, --help                                                             help for search-networks
      --network-search-request-additional-properties string              network-search-request-additional-properties (JSON)
      --network-search-request-filter-additional-properties string       network-search-request-filter-additional-properties (JSON)
      --network-search-request-filter-and string                         network-search-request-filter-and (JSON array)
      --network-search-request-filter-operator string                    network-search-request-filter-operator
      --network-search-request-filter-or string                          network-search-request-filter-or (JSON array)
      --network-search-request-filter-property string                    network-search-request-filter-property
      --network-search-request-filter-values string                      network-search-request-filter-values (JSON array)
      --network-search-request-pagination-additional-properties string   network-search-request-pagination-additional-properties (JSON)
      --network-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --network-search-request-pagination-offset int                     Index of the first element.
      --network-search-request-sort string                               network-search-request-sort (JSON array)
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 networks](equinix_fabricv4_networks.md)	 - Manage networks resources

