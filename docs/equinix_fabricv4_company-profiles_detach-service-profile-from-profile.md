## equinix fabricv4 company-profiles detach-service-profile-from-profile

Detach Service Profile

### Synopsis

Detach a service profile from a company profile <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 company-profiles detach-service-profile-from-profile [flags]
```

### Options

```
      --company-profile-id string   Company Profile UUID (required)
  -h, --help                        help for detach-service-profile-from-profile
      --request string              JSON payload for request body fields
      --service-profile-id string   Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 company-profiles](equinix_fabricv4_company-profiles.md)	 - Manage company-profiles resources

