## equinix metalv1 v-r-fs update-vrf-route-by-id

Execute update-vrf-route-by-id operation

### Synopsis

Execute the update-vrf-route-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-r-fs update-vrf-route-by-id [flags]
```

### Options

```
      --exclude string                                        exclude field (JSON or string)
  -h, --help                                                  help for update-vrf-route-by-id
      --id string                                             id (required)
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

* [equinix metalv1 v-r-fs](equinix_metalv1_v-r-fs.md)	 - Manage v-r-fs resources

