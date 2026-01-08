## equinix fabricv4 cloud-routers list-gateway-attachments-to-cloud-router

List Cloud Routers of a Gateway Attachment.

### Synopsis

Get all Cloud Routers attached on a Gateway.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers list-gateway-attachments-to-cloud-router [flags]
```

### Options

```
      --gateway-id string   Gateway UUID (required)
  -h, --help                help for list-gateway-attachments-to-cloud-router
      --limit int           limit field
      --offset int          offset field
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

