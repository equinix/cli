## equinix fabric cloud-routers update-cloud-router-by-uuid

Update Routers

### Synopsis

This API provides capability to update user's Cloud Routers

Use --request flag to provide optional JSON payload fields.

```
equinix fabric cloud-routers update-cloud-router-by-uuid [flags]
```

### Options

```
      --cloud-router-change-operation string   cloud-router-change-operation field (JSON or string)
  -h, --help                                   help for update-cloud-router-by-uuid
      --request string                         JSON payload for additional optional fields not exposed as flags
      --router-id string                       Cloud Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabric cloud-routers](equinix_fabric_cloud-routers.md)	 - Manage cloud-routers resources

