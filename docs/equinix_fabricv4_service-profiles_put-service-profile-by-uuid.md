## equinix fabricv4 service-profiles put-service-profile-by-uuid

Replace Profile

### Synopsis

This API request replaces a service profile definition

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles put-service-profile-by-uuid [flags]
```

### Options

```
  -h, --help                        help for put-service-profile-by-uuid
      --request string              JSON payload for request body. Available fields: if-match (string), service-profile-request (ServiceProfileRequest)
      --service-profile-id string   Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

