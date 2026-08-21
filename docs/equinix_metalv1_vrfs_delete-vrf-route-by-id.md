## equinix metalv1 vrfs delete-vrf-route-by-id

Delete a VRF Route

### Synopsis

Trigger the deletion of a VRF Route resource. The status of the route will update to 'deleting', and the route resource will remain accessible while background operations remove the route from the network. Once the route has been removed from the network, the resource will be fully deleted.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs delete-vrf-route-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for delete-vrf-route-by-id
      --id string        VRF Route UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

