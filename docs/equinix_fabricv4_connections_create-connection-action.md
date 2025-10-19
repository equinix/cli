## equinix fabricv4 connections create-connection-action

Connection Actions

### Synopsis

This API provides capability to accept/reject user's virtual connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections create-connection-action [flags]
```

### Options

```
      --connection-id string   Connection Id (required)
  -h, --help                   help for create-connection-action
      --request string         JSON payload for request body. Available fields: connection-action-request (ConnectionActionRequest)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

