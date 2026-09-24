## equinix fabric company-profiles detach-tag-from-profile

Detach Tag

### Synopsis

Detach a tag from a company profile <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabric company-profiles detach-tag-from-profile [flags]
```

### Options

```
      --company-profile-id string   Company Profile UUID (required)
  -h, --help                        help for detach-tag-from-profile
      --request string              JSON payload for request body fields
      --tag-id string               Tag UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabric company-profiles](equinix_fabric_company-profiles.md)	 - Manage company-profiles resources

