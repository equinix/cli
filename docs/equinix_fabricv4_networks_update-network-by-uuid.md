## equinix fabricv4 networks update-network-by-uuid

Update Network By ID

### Synopsis

This API provides capability to update user's Fabric Network

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 networks update-network-by-uuid [flags]
```

### Options

```
  -h, --help                              help for update-network-by-uuid
      --network-change-operation string   network-change-operation field (required) (JSON or string)
      --network-id string                 Network UUID (required)
      --request string                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 networks](equinix_fabricv4_networks.md)	 - Manage networks resources

