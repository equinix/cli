## equinix fabricv4 service-profiles get-service-profiles

Get all Profiles

### Synopsis

The API request returns all Equinix Fabric Service Profiles in accordance with the view point requested.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles get-service-profiles [flags]
```

### Options

```
  -h, --help             help for get-service-profiles
      --request string   JSON payload for request body. Available fields: limit (int32), offset (int32), view-point (GetServiceProfilesViewPointParameter)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

