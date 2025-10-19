## equinix fabricv4 service-tokens create-service-token-action

ServiceToken Actions

### Synopsis

This API provides capability to accept/reject user's servicetokens

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens create-service-token-action [flags]
```

### Options

```
  -h, --help                      help for create-service-token-action
      --request string            JSON payload for request body. Available fields: service-token-action-request (ServiceTokenActionRequest)
      --service-token-id string   Service Token UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

