## equinix fabricv4 service-profiles get-service-profiles

Get all Profiles

### Synopsis

The API request returns all Equinix Fabric Service Profiles in accordance with the view point requested.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles get-service-profiles [flags]
```

### Options

```
  -h, --help                help for get-service-profiles
      --limit int           limit field
      --offset int          offset field
      --request string      JSON payload for additional optional fields not exposed as flags
      --view-point string   view-point field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

