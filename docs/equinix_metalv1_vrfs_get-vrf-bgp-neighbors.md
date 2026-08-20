## equinix metalv1 vrfs get-vrf-bgp-neighbors

Retrieve BGP neighbor states for the VRF

### Synopsis

Provides BGP peering information such as the IP and state of the neighbor.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs get-vrf-bgp-neighbors [flags]
```

### Options

```
  -h, --help             help for get-vrf-bgp-neighbors
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

