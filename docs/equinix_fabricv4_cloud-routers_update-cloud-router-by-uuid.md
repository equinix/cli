## equinix fabricv4 cloud-routers update-cloud-router-by-uuid

Update Routers

### Synopsis

This API provides capability to update user's Cloud Routers

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers update-cloud-router-by-uuid [flags]
```

### Options

```
      --cloud-router-change-operation string   cloud-router-change-operation field (required) (JSON or string)
  -h, --help                                   help for update-cloud-router-by-uuid
      --request string                         JSON payload for additional optional fields not exposed as flags
      --router-id string                       Cloud Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

