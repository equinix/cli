## equinix metalv1 vrfs get-vrf-learned-routes

Retrieve learned L3 routes within the VRF

### Synopsis

Provides information about learned routes for the VRF. The VRF builds this information dynamically though BGP from other routers in the network.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs get-vrf-learned-routes [flags]
```

### Options

```
  -h, --help             help for get-vrf-learned-routes
      --id string        VRF UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

