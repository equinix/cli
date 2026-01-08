## equinix fabricv4 metros get-metros

Get all Metros

### Synopsis

GET All Subscriber Metros with an option query parameter to return all Equinix Fabric metros in which the customer has a presence, as well as latency data for each location.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 metros get-metros [flags]
```

### Options

```
  -h, --help              help for get-metros
      --limit int         limit field
      --offset int        offset field
      --presence string   presence field
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 metros](equinix_fabricv4_metros.md)	 - Manage metros resources

