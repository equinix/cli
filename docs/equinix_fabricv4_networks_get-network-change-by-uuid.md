## equinix fabricv4 networks get-network-change-by-uuid

Get Change By ID

### Synopsis

This API provides capability to retrieve user's Fabric Network Change

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 networks get-network-change-by-uuid [flags]
```

### Options

```
      --change-id string    Network Change UUID (required)
  -h, --help                help for get-network-change-by-uuid
      --network-id string   Network UUID (required)
      --request string      JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 networks](equinix_fabricv4_networks.md)	 - Manage networks resources

