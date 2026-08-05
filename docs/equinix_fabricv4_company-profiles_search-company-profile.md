## equinix fabricv4 company-profiles search-company-profile

Execute search-company-profile operation

### Synopsis

Execute the search-company-profile operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 company-profiles search-company-profile [flags]
```

### Options

```
      --company-profile-search-request-additional-properties string              company-profile-search-request-additional-properties (JSON)
      --company-profile-search-request-filter-additional-properties string       company-profile-search-request-filter-additional-properties (JSON)
      --company-profile-search-request-filter-and string                         company-profile-search-request-filter-and (JSON array)
      --company-profile-search-request-filter-operator string                    company-profile-search-request-filter-operator
      --company-profile-search-request-filter-or string                          company-profile-search-request-filter-or (JSON array)
      --company-profile-search-request-filter-property string                    company-profile-search-request-filter-property
      --company-profile-search-request-filter-values string                      company-profile-search-request-filter-values (JSON array)
      --company-profile-search-request-pagination-additional-properties string   company-profile-search-request-pagination-additional-properties (JSON)
      --company-profile-search-request-pagination-limit int                      Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
      --company-profile-search-request-pagination-next string                    URL relative to the next item in the response.
      --company-profile-search-request-pagination-offset int                     Index of the first item returned in the response. The default is 0.
      --company-profile-search-request-pagination-previous string                URL relative to the previous item in the response.
      --company-profile-search-request-pagination-total int                      Total number of elements returned.
      --company-profile-search-request-sort-additional-properties string         company-profile-search-request-sort-additional-properties (JSON)
      --company-profile-search-request-sort-direction string                     company-profile-search-request-sort-direction
      --company-profile-search-request-sort-property string                      company-profile-search-request-sort-property
  -h, --help                                                                     help for search-company-profile
      --request string                                                           JSON payload for additional optional fields not exposed as flags
      --view-point string                                                        view-point field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 company-profiles](equinix_fabricv4_company-profiles.md)	 - Manage company-profiles resources

