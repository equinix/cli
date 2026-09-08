## equinix fabricv4 company-profiles search-company-profile

Search Company Profiles

### Synopsis

Search company profiles based on filter criteria <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 company-profiles search-company-profile [flags]
```

### Options

```
      --company-profile-search-request-additional-properties string              company-profile-search-request-additional-properties (JSON)
      --company-profile-search-request-filter-additional-properties string       company-profile-search-request-filter-additional-properties (JSON)
      --company-profile-search-request-filter-and string                         company-profile-search-request-filter-and (JSON array)
      --company-profile-search-request-pagination-additional-properties string   company-profile-search-request-pagination-additional-properties (JSON)
      --company-profile-search-request-pagination-limit int                      Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
      --company-profile-search-request-pagination-next string                    URL relative to the next item in the response.
      --company-profile-search-request-pagination-offset int                     Index of the first item returned in the response. The default is 0.
      --company-profile-search-request-pagination-previous string                URL relative to the previous item in the response.
      --company-profile-search-request-pagination-total int                      Total number of elements returned.
      --company-profile-search-request-sort-additional-properties string         company-profile-search-request-sort-additional-properties (JSON)
      --company-profile-search-request-sort-direction string                     company-profile-search-request-sort-direction
      --company-profile-search-request-sort-property /name                       Property to sort by. Supported values:  * /name - Company profile name  * `/state` - Company profile state  * `/changeLog/updatedDateTime` - Date and time the profile was last updated  * `/changeLog/createdDateTime` - Date and time the profile was created
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

