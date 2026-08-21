## equinix metalv1 vrfs update-vrf-route-by-id

Update a VRF Route

### Synopsis

Requests a VRF Route be redeployed across the network. Updating the prefix or next-hop address on a route is not currently supported.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs update-vrf-route-by-id [flags]
```

### Options

```
      --exclude string                                        exclude field (JSON or string)
  -h, --help                                                  help for update-vrf-route-by-id
      --id string                                             VRF Route UUID (required)
      --include string                                        include field (JSON or string)
      --request string                                        JSON payload for additional optional fields not exposed as flags
      --vrf-route-update-input-additional-properties string   vrf-route-update-input-additional-properties (JSON)
      --vrf-route-update-input-next_hop string                The IPv4 address within the VRF of the host that will handle this route
      --vrf-route-update-input-prefix string                  The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be \"0.0.0.0/0\"
      --vrf-route-update-input-tags string                    vrf-route-update-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

