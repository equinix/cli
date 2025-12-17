## equinix fabricv4 cloud-routers get-cloud-router-actions-by-uuid

Get Route Table Action by ID

### Synopsis

This API provides capability to fetch action status

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers get-cloud-router-actions-by-uuid [flags]
```

### Options

```
      --action-id string   Action UUID (required)
  -h, --help               help for get-cloud-router-actions-by-uuid
      --request string     JSON payload for additional optional fields not exposed as flags
      --router-id string   Router UUID (required)
      --state string       state field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

