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
      --direction string       direction field
  -h, --help                   help for get-connection-by-uuid
      --request string         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

