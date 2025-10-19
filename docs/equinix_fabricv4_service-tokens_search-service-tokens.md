## equinix fabricv4 service-tokens search-service-tokens

Search servicetokens

### Synopsis

The API provides capability to get list of user's servicetokens using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens search-service-tokens [flags]
```

### Options

```
  -h, --help             help for search-service-tokens
      --request string   JSON payload for request body. Available fields: limit (float32), offset (float32), service-token-search-request (ServiceTokenSearchRequest)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

