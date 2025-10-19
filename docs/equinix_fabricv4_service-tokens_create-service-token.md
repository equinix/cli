## equinix fabricv4 service-tokens create-service-token

Create Service Token

### Synopsis

Create Service Tokens generates Equinix Fabric™ service tokens. These tokens authorize users to access protected resources and services.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens create-service-token [flags]
```

### Options

```
  -h, --help             help for create-service-token
      --request string   JSON payload for request body. Available fields: dry-run (bool), service-token (ServiceToken)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

