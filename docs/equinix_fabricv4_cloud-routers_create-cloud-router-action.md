## equinix fabricv4 cloud-routers create-cloud-router-action

Create Route Table Action

### Synopsis

This API provides capability to refresh route table and bgp session summary information

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers create-cloud-router-action [flags]
```

### Options

```
      --cloud-router-action-request-additional-properties string              cloud-router-action-request-additional-properties (JSON)
      --cloud-router-action-request-connection-additional-properties string   cloud-router-action-request-connection-additional-properties (JSON)
      --cloud-router-action-request-connection-href string                    cloud-router-action-request-connection-href
      --cloud-router-action-request-connection-operation string               cloud-router-action-request-connection-operation (JSON)
      --cloud-router-action-request-connection-type string                    cloud-router-action-request-connection-type
      --cloud-router-action-request-connection-uuid string                    Connection UUID
      --cloud-router-action-request-type string                               cloud-router-action-request-type
  -h, --help                                                                  help for create-cloud-router-action
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --router-id string                                                      Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

