## equinix fabricv4 connections get-connection-by-uuid

Get Connection by ID

### Synopsis

The API provides capability to get user's virtual connection details (Service Tokens, Access Points, Link Protocols, etc) by it's connection ID (UUID)

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections get-connection-by-uuid [flags]
```

### Options

```
      --connection-id string   Connection Id (required)
  -h, --help                   help for get-connection-by-uuid
      --request string         Raw JSON payload for optional request fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

