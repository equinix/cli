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
      --cloud-router-action-request-additional-properties string              cloud-router-action-request-additional-properties (required) (JSON)
      --cloud-router-action-request-connection-additional-properties string   cloud-router-action-request-connection-additional-properties (JSON)
      --cloud-router-action-request-connection-uuid string                    Connection UUID
      --cloud-router-action-request-type string                               cloud-router-action-request-type (required)
  -h, --help                                                                  help for create-cloud-router-action
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --router-id string                                                      Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

