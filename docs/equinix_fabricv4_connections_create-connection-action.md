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
      --connection-action-request-additional-properties string        connection-action-request-additional-properties (required) (JSON)
      --connection-action-request-data-additional-properties string   connection-action-request-data-additional-properties (JSON)
      --connection-action-request-data-provider-bandwidth int         Authorization key bandwidth in Mbps
      --connection-action-request-data-z-side string                  connection-action-request-data-z-side (JSON)
      --connection-action-request-description string                  Connection rejection reason detail
      --connection-action-request-type string                         connection-action-request-type (required)
      --connection-id string                                          Connection Id (required)
  -h, --help                                                          help for create-connection-action
      --request string                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

