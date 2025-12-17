## equinix fabricv4 cloud-routers get-gateway-attachment-to-cloud-router-by-uuid

Get Gateway Attachment details to a Cloud Router

### Synopsis

Get details of a Specific Gateway Attachment to a Cloud Router.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers get-gateway-attachment-to-cloud-router-by-uuid [flags]
```

### Options

```
      --gateway-id string   Gateway UUID (required)
  -h, --help                help for get-gateway-attachment-to-cloud-router-by-uuid
      --request string      JSON payload for request body fields
      --router-id string    Cloud Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

