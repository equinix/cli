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
      --limit int         limit field (required)
      --offset int        offset field (required)
      --presence string   presence field (required)
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 metros](equinix_fabricv4_metros.md)	 - Manage metros resources

