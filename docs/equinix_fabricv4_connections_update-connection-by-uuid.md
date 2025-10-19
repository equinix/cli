## equinix fabricv4 connections update-connection-by-uuid

Update by ID

### Synopsis

Update Connection by ID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections update-connection-by-uuid [flags]
```

### Options

```
      --connection-id string   Connection Id (required)
  -h, --help                   help for update-connection-by-uuid
      --request string         JSON payload for request body. Available fields: connection-change-operation (ConnectionChangeOperation), dry-run (bool)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

