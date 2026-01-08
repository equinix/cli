## equinix fabricv4 cloud-routers get-cloud-router-command

Get Command

### Synopsis

This API provides capability to fetch command using command Id

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers get-cloud-router-command [flags]
```

### Options

```
      --command-id string   Command UUID (required)
  -h, --help                help for get-cloud-router-command
      --request string      JSON payload for request body fields
      --router-id string    Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

