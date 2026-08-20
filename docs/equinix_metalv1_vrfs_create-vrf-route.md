## equinix metalv1 vrfs create-vrf-route

Create a VRF route

### Synopsis

Create a route in a VRF. Currently only static default routes are supported.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs create-vrf-route [flags]
```

### Options

```
      --exclude string                                        exclude field (JSON or string)
  -h, --help                                                  help for create-vrf-route
      --id string                                             VRF UUID (required)
      --include string                                        include field (JSON or string)
      --request string                                        JSON payload for additional optional fields not exposed as flags
      --vrf-route-create-input-additional-properties string   vrf-route-create-input-additional-properties (JSON)
      --vrf-route-create-input-next_hop string                The IPv4 address within the VRF of the host that will handle this route
      --vrf-route-create-input-prefix string                  The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be \"0.0.0.0/0\"
      --vrf-route-create-input-tags string                    vrf-route-create-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

