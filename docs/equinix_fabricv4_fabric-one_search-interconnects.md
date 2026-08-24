## equinix fabricv4 fabric-one search-interconnects

Search Interconnects

### Synopsis

The API provides capability to get list of user's Interconnects using search criteria, including optional filtering, pagination and sorting <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 fabric-one search-interconnects [flags]
```

### Options

```
  -h, --help                                                                  help for search-interconnects
      --interconnect-search-request-additional-properties string              interconnect-search-request-additional-properties (JSON)
      --interconnect-search-request-filter-additional-properties string       interconnect-search-request-filter-additional-properties (JSON)
      --interconnect-search-request-filter-and string                         interconnect-search-request-filter-and (JSON array)
      --interconnect-search-request-filter-operator string                    interconnect-search-request-filter-operator
      --interconnect-search-request-filter-or string                          interconnect-search-request-filter-or (JSON array)
      --interconnect-search-request-filter-property string                    interconnect-search-request-filter-property
      --interconnect-search-request-filter-values string                      interconnect-search-request-filter-values (JSON array)
      --interconnect-search-request-pagination-additional-properties string   interconnect-search-request-pagination-additional-properties (JSON)
      --interconnect-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --interconnect-search-request-pagination-offset int                     Index of the first element.
      --interconnect-search-request-sort string                               interconnect-search-request-sort (JSON array)
      --request string                                                        JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 fabric-one](equinix_fabricv4_fabric-one.md)	 - Manage fabric-one resources

